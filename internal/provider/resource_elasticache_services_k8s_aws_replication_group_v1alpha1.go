/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource)(nil)
)

type ElasticacheServicesK8SAwsReplicationGroupV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ElasticacheServicesK8SAwsReplicationGroupV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AtRestEncryptionEnabled *bool `tfsdk:"at_rest_encryption_enabled" yaml:"atRestEncryptionEnabled,omitempty"`

		AuthToken *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"auth_token" yaml:"authToken,omitempty"`

		AutomaticFailoverEnabled *bool `tfsdk:"automatic_failover_enabled" yaml:"automaticFailoverEnabled,omitempty"`

		CacheNodeType *string `tfsdk:"cache_node_type" yaml:"cacheNodeType,omitempty"`

		CacheParameterGroupName *string `tfsdk:"cache_parameter_group_name" yaml:"cacheParameterGroupName,omitempty"`

		CacheSecurityGroupNames *[]string `tfsdk:"cache_security_group_names" yaml:"cacheSecurityGroupNames,omitempty"`

		CacheSubnetGroupName *string `tfsdk:"cache_subnet_group_name" yaml:"cacheSubnetGroupName,omitempty"`

		DataTieringEnabled *bool `tfsdk:"data_tiering_enabled" yaml:"dataTieringEnabled,omitempty"`

		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		Engine *string `tfsdk:"engine" yaml:"engine,omitempty"`

		EngineVersion *string `tfsdk:"engine_version" yaml:"engineVersion,omitempty"`

		KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

		LogDeliveryConfigurations *[]struct {
			DestinationDetails *struct {
				CloudWatchLogsDetails *struct {
					LogGroup *string `tfsdk:"log_group" yaml:"logGroup,omitempty"`
				} `tfsdk:"cloud_watch_logs_details" yaml:"cloudWatchLogsDetails,omitempty"`

				KinesisFirehoseDetails *struct {
					DeliveryStream *string `tfsdk:"delivery_stream" yaml:"deliveryStream,omitempty"`
				} `tfsdk:"kinesis_firehose_details" yaml:"kinesisFirehoseDetails,omitempty"`
			} `tfsdk:"destination_details" yaml:"destinationDetails,omitempty"`

			DestinationType *string `tfsdk:"destination_type" yaml:"destinationType,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			LogFormat *string `tfsdk:"log_format" yaml:"logFormat,omitempty"`

			LogType *string `tfsdk:"log_type" yaml:"logType,omitempty"`
		} `tfsdk:"log_delivery_configurations" yaml:"logDeliveryConfigurations,omitempty"`

		MultiAZEnabled *bool `tfsdk:"multi_az_enabled" yaml:"multiAZEnabled,omitempty"`

		NodeGroupConfiguration *[]struct {
			NodeGroupID *string `tfsdk:"node_group_id" yaml:"nodeGroupID,omitempty"`

			PrimaryAvailabilityZone *string `tfsdk:"primary_availability_zone" yaml:"primaryAvailabilityZone,omitempty"`

			PrimaryOutpostARN *string `tfsdk:"primary_outpost_arn" yaml:"primaryOutpostARN,omitempty"`

			ReplicaAvailabilityZones *[]string `tfsdk:"replica_availability_zones" yaml:"replicaAvailabilityZones,omitempty"`

			ReplicaCount *int64 `tfsdk:"replica_count" yaml:"replicaCount,omitempty"`

			ReplicaOutpostARNs *[]string `tfsdk:"replica_outpost_ar_ns" yaml:"replicaOutpostARNs,omitempty"`

			Slots *string `tfsdk:"slots" yaml:"slots,omitempty"`
		} `tfsdk:"node_group_configuration" yaml:"nodeGroupConfiguration,omitempty"`

		NotificationTopicARN *string `tfsdk:"notification_topic_arn" yaml:"notificationTopicARN,omitempty"`

		NumNodeGroups *int64 `tfsdk:"num_node_groups" yaml:"numNodeGroups,omitempty"`

		Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

		PreferredCacheClusterAZs *[]string `tfsdk:"preferred_cache_cluster_a_zs" yaml:"preferredCacheClusterAZs,omitempty"`

		PreferredMaintenanceWindow *string `tfsdk:"preferred_maintenance_window" yaml:"preferredMaintenanceWindow,omitempty"`

		PrimaryClusterID *string `tfsdk:"primary_cluster_id" yaml:"primaryClusterID,omitempty"`

		ReplicasPerNodeGroup *int64 `tfsdk:"replicas_per_node_group" yaml:"replicasPerNodeGroup,omitempty"`

		ReplicationGroupID *string `tfsdk:"replication_group_id" yaml:"replicationGroupID,omitempty"`

		SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

		SnapshotARNs *[]string `tfsdk:"snapshot_ar_ns" yaml:"snapshotARNs,omitempty"`

		SnapshotName *string `tfsdk:"snapshot_name" yaml:"snapshotName,omitempty"`

		SnapshotRetentionLimit *int64 `tfsdk:"snapshot_retention_limit" yaml:"snapshotRetentionLimit,omitempty"`

		SnapshotWindow *string `tfsdk:"snapshot_window" yaml:"snapshotWindow,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		TransitEncryptionEnabled *bool `tfsdk:"transit_encryption_enabled" yaml:"transitEncryptionEnabled,omitempty"`

		UserGroupIDs *[]string `tfsdk:"user_group_i_ds" yaml:"userGroupIDs,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource() resource.Resource {
	return &ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource{}
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_elasticache_services_k8s_aws_replication_group_v1alpha1"
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ReplicationGroup is the Schema for the ReplicationGroups API",
		MarkdownDescription: "ReplicationGroup is the Schema for the ReplicationGroups API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "ReplicationGroupSpec defines the desired state of ReplicationGroup.  Contains all of the attributes of a specific Redis replication group.",
				MarkdownDescription: "ReplicationGroupSpec defines the desired state of ReplicationGroup.  Contains all of the attributes of a specific Redis replication group.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"at_rest_encryption_enabled": {
						Description:         "A flag that enables encryption at rest when set to true.  You cannot modify the value of AtRestEncryptionEnabled after the replication group is created. To enable encryption at rest on a replication group you must set AtRestEncryptionEnabled to true when you create the replication group.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false",
						MarkdownDescription: "A flag that enables encryption at rest when set to true.  You cannot modify the value of AtRestEncryptionEnabled after the replication group is created. To enable encryption at rest on a replication group you must set AtRestEncryptionEnabled to true when you create the replication group.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auth_token": {
						Description:         "Reserved parameter. The password used to access a password protected server.  AuthToken can be specified only on replication groups where TransitEncryptionEnabled is true.  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.  Password constraints:  * Must be only printable ASCII characters.  * Must be at least 16 characters and no more than 128 characters in length.  * The only permitted printable special characters are !, &, #, $, ^, <, >, and -. Other printable special characters cannot be used in the AUTH token.  For more information, see AUTH password (http://redis.io/commands/AUTH) at http://redis.io/commands/AUTH.",
						MarkdownDescription: "Reserved parameter. The password used to access a password protected server.  AuthToken can be specified only on replication groups where TransitEncryptionEnabled is true.  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.  Password constraints:  * Must be only printable ASCII characters.  * Must be at least 16 characters and no more than 128 characters in length.  * The only permitted printable special characters are !, &, #, $, ^, <, >, and -. Other printable special characters cannot be used in the AUTH token.  For more information, see AUTH password (http://redis.io/commands/AUTH) at http://redis.io/commands/AUTH.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "Key is the key within the secret",
								MarkdownDescription: "Key is the key within the secret",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"name": {
								Description:         "Name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"automatic_failover_enabled": {
						Description:         "Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.  AutomaticFailoverEnabled must be enabled for Redis (cluster mode enabled) replication groups.  Default: false",
						MarkdownDescription: "Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.  AutomaticFailoverEnabled must be enabled for Redis (cluster mode enabled) replication groups.  Default: false",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_node_type": {
						Description:         "The compute and memory capacity of the nodes in the node group (shard).  The following node types are supported by ElastiCache. Generally speaking, the current generation types provide more memory and computational power at lower cost when compared to their equivalent previous generation counterparts.  * General purpose: Current generation: M6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward): cache.m6g.large, cache.m6g.xlarge, cache.m6g.2xlarge, cache.m6g.4xlarge, cache.m6g.8xlarge, cache.m6g.12xlarge, cache.m6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) M5 node types: cache.m5.large, cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge, cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge, cache.m4.4xlarge, cache.m4.10xlarge T4g node types (available only for Redis engine version 5.0.6 onward and Memcached engine version 1.5.16 onward): cache.t4g.micro, cache.t4g.small, cache.t4g.medium T3 node types: cache.t3.micro, cache.t3.small, cache.t3.medium T2 node types: cache.t2.micro, cache.t2.small, cache.t2.medium Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) T1 node types: cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large, cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge, cache.m3.2xlarge  * Compute optimized: Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) C1 node types: cache.c1.xlarge  * Memory optimized with data tiering: Current generation: R6gd node types (available only for Redis engine version 6.2 onward). cache.r6gd.xlarge, cache.r6gd.2xlarge, cache.r6gd.4xlarge, cache.r6gd.8xlarge, cache.r6gd.12xlarge, cache.r6gd.16xlarge  * Memory optimized: Current generation: R6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward). cache.r6g.large, cache.r6g.xlarge, cache.r6g.2xlarge, cache.r6g.4xlarge, cache.r6g.8xlarge, cache.r6g.12xlarge, cache.r6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) R5 node types: cache.r5.large, cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge, cache.r5.12xlarge, cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge, cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge, cache.r3.8xlarge  Additional node type info  * All current generation instance types are created in Amazon VPC by default.  * Redis append-only files (AOF) are not supported for T1 or T2 instances.  * Redis Multi-AZ with automatic failover is not supported on T1 instances.  * Redis configuration variables appendonly and appendfsync are not supported on Redis version 2.8.22 and later.",
						MarkdownDescription: "The compute and memory capacity of the nodes in the node group (shard).  The following node types are supported by ElastiCache. Generally speaking, the current generation types provide more memory and computational power at lower cost when compared to their equivalent previous generation counterparts.  * General purpose: Current generation: M6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward): cache.m6g.large, cache.m6g.xlarge, cache.m6g.2xlarge, cache.m6g.4xlarge, cache.m6g.8xlarge, cache.m6g.12xlarge, cache.m6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) M5 node types: cache.m5.large, cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge, cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge, cache.m4.4xlarge, cache.m4.10xlarge T4g node types (available only for Redis engine version 5.0.6 onward and Memcached engine version 1.5.16 onward): cache.t4g.micro, cache.t4g.small, cache.t4g.medium T3 node types: cache.t3.micro, cache.t3.small, cache.t3.medium T2 node types: cache.t2.micro, cache.t2.small, cache.t2.medium Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) T1 node types: cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large, cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge, cache.m3.2xlarge  * Compute optimized: Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) C1 node types: cache.c1.xlarge  * Memory optimized with data tiering: Current generation: R6gd node types (available only for Redis engine version 6.2 onward). cache.r6gd.xlarge, cache.r6gd.2xlarge, cache.r6gd.4xlarge, cache.r6gd.8xlarge, cache.r6gd.12xlarge, cache.r6gd.16xlarge  * Memory optimized: Current generation: R6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward). cache.r6g.large, cache.r6g.xlarge, cache.r6g.2xlarge, cache.r6g.4xlarge, cache.r6g.8xlarge, cache.r6g.12xlarge, cache.r6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) R5 node types: cache.r5.large, cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge, cache.r5.12xlarge, cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge, cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge, cache.r3.8xlarge  Additional node type info  * All current generation instance types are created in Amazon VPC by default.  * Redis append-only files (AOF) are not supported for T1 or T2 instances.  * Redis Multi-AZ with automatic failover is not supported on T1 instances.  * Redis configuration variables appendonly and appendfsync are not supported on Redis version 2.8.22 and later.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_parameter_group_name": {
						Description:         "The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.  If you are running Redis version 3.2.4 or later, only one node group (shard), and want to use a default parameter group, we recommend that you specify the parameter group by name.  * To create a Redis (cluster mode disabled) replication group, use CacheParameterGroupName=default.redis3.2.  * To create a Redis (cluster mode enabled) replication group, use CacheParameterGroupName=default.redis3.2.cluster.on.",
						MarkdownDescription: "The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.  If you are running Redis version 3.2.4 or later, only one node group (shard), and want to use a default parameter group, we recommend that you specify the parameter group by name.  * To create a Redis (cluster mode disabled) replication group, use CacheParameterGroupName=default.redis3.2.  * To create a Redis (cluster mode enabled) replication group, use CacheParameterGroupName=default.redis3.2.cluster.on.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_security_group_names": {
						Description:         "A list of cache security group names to associate with this replication group.",
						MarkdownDescription: "A list of cache security group names to associate with this replication group.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_subnet_group_name": {
						Description:         "The name of the cache subnet group to be used for the replication group.  If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start creating a cluster. For more information, see Subnets and Subnet Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html).",
						MarkdownDescription: "The name of the cache subnet group to be used for the replication group.  If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start creating a cluster. For more information, see Subnets and Subnet Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"data_tiering_enabled": {
						Description:         "Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes. For more information, see Data tiering (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/data-tiering.html).",
						MarkdownDescription: "Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes. For more information, see Data tiering (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/data-tiering.html).",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"description": {
						Description:         "A user-created description for the replication group.",
						MarkdownDescription: "A user-created description for the replication group.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"engine": {
						Description:         "The name of the cache engine to be used for the clusters in this replication group. Must be Redis.",
						MarkdownDescription: "The name of the cache engine to be used for the clusters in this replication group. Must be Redis.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"engine_version": {
						Description:         "The version number of the cache engine to be used for the clusters in this replication group. To view the supported cache engine versions, use the DescribeCacheEngineVersions operation.  Important: You can upgrade to a newer engine version (see Selecting a Cache Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)) in the ElastiCache User Guide, but you cannot downgrade to an earlier engine version. If you want to use an earlier engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine version.",
						MarkdownDescription: "The version number of the cache engine to be used for the clusters in this replication group. To view the supported cache engine versions, use the DescribeCacheEngineVersions operation.  Important: You can upgrade to a newer engine version (see Selecting a Cache Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)) in the ElastiCache User Guide, but you cannot downgrade to an earlier engine version. If you want to use an earlier engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine version.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kms_key_id": {
						Description:         "The ID of the KMS key used to encrypt the disk in the cluster.",
						MarkdownDescription: "The ID of the KMS key used to encrypt the disk in the cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_delivery_configurations": {
						Description:         "Specifies the destination, format and type of the logs.",
						MarkdownDescription: "Specifies the destination, format and type of the logs.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"destination_details": {
								Description:         "Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.",
								MarkdownDescription: "Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cloud_watch_logs_details": {
										Description:         "The configuration details of the CloudWatch Logs destination.",
										MarkdownDescription: "The configuration details of the CloudWatch Logs destination.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"log_group": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kinesis_firehose_details": {
										Description:         "The configuration details of the Kinesis Data Firehose destination.",
										MarkdownDescription: "The configuration details of the Kinesis Data Firehose destination.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"delivery_stream": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"destination_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_format": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"multi_az_enabled": {
						Description:         "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For more information, see Minimizing Downtime: Multi-AZ (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html).",
						MarkdownDescription: "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For more information, see Minimizing Downtime: Multi-AZ (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html).",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_group_configuration": {
						Description:         "A list of node group (shard) configuration options. Each node group (shard) configuration has the following members: PrimaryAvailabilityZone, ReplicaAvailabilityZones, ReplicaCount, and Slots.  If you're creating a Redis (cluster mode disabled) or a Redis (cluster mode enabled) replication group, you can use this parameter to individually configure each node group (shard), or you can omit this parameter. However, it is required when seeding a Redis (cluster mode enabled) cluster from a S3 rdb file. You must configure each node group (shard) using this parameter because you must specify the slots for each node group.",
						MarkdownDescription: "A list of node group (shard) configuration options. Each node group (shard) configuration has the following members: PrimaryAvailabilityZone, ReplicaAvailabilityZones, ReplicaCount, and Slots.  If you're creating a Redis (cluster mode disabled) or a Redis (cluster mode enabled) replication group, you can use this parameter to individually configure each node group (shard), or you can omit this parameter. However, it is required when seeding a Redis (cluster mode enabled) cluster from a S3 rdb file. You must configure each node group (shard) using this parameter because you must specify the slots for each node group.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"node_group_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"primary_availability_zone": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"primary_outpost_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_availability_zones": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_outpost_ar_ns": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"slots": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"notification_topic_arn": {
						Description:         "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.  The Amazon SNS topic owner must be the same as the cluster owner.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.  The Amazon SNS topic owner must be the same as the cluster owner.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"num_node_groups": {
						Description:         "An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group. For Redis (cluster mode disabled) either omit this parameter or set it to 1.  Default: 1",
						MarkdownDescription: "An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group. For Redis (cluster mode disabled) either omit this parameter or set it to 1.  Default: 1",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": {
						Description:         "The port number on which each member of the replication group accepts connections.",
						MarkdownDescription: "The port number on which each member of the replication group accepts connections.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"preferred_cache_cluster_a_zs": {
						Description:         "A list of EC2 Availability Zones in which the replication group's clusters are created. The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list.  This parameter is not used if there is more than one node group (shard). You should use NodeGroupConfiguration instead.  If you are creating your replication group in an Amazon VPC (recommended), you can only locate clusters in Availability Zones associated with the subnets in the selected subnet group.  The number of Availability Zones listed must equal the value of NumCacheClusters.  Default: system chosen Availability Zones.",
						MarkdownDescription: "A list of EC2 Availability Zones in which the replication group's clusters are created. The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list.  This parameter is not used if there is more than one node group (shard). You should use NodeGroupConfiguration instead.  If you are creating your replication group in an Amazon VPC (recommended), you can only locate clusters in Availability Zones associated with the subnets in the selected subnet group.  The number of Availability Zones listed must equal the value of NumCacheClusters.  Default: system chosen Availability Zones.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"preferred_maintenance_window": {
						Description:         "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid values for ddd are:  Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.  Valid values for ddd are:  * sun  * mon  * tue  * wed  * thu  * fri  * sat  Example: sun:23:00-mon:01:30",
						MarkdownDescription: "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid values for ddd are:  Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.  Valid values for ddd are:  * sun  * mon  * tue  * wed  * thu  * fri  * sat  Example: sun:23:00-mon:01:30",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"primary_cluster_id": {
						Description:         "The identifier of the cluster that serves as the primary for this replication group. This cluster must already exist and have a status of available.  This parameter is not required if NumCacheClusters, NumNodeGroups, or ReplicasPerNodeGroup is specified.",
						MarkdownDescription: "The identifier of the cluster that serves as the primary for this replication group. This cluster must already exist and have a status of available.  This parameter is not required if NumCacheClusters, NumNodeGroups, or ReplicasPerNodeGroup is specified.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas_per_node_group": {
						Description:         "An optional parameter that specifies the number of replica nodes in each node group (shard). Valid values are 0 to 5.",
						MarkdownDescription: "An optional parameter that specifies the number of replica nodes in each node group (shard). Valid values are 0 to 5.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replication_group_id": {
						Description:         "The replication group identifier. This parameter is stored as a lowercase string.  Constraints:  * A name must contain from 1 to 40 alphanumeric characters or hyphens.  * The first character must be a letter.  * A name cannot end with a hyphen or contain two consecutive hyphens.",
						MarkdownDescription: "The replication group identifier. This parameter is stored as a lowercase string.  Constraints:  * A name must contain from 1 to 40 alphanumeric characters or hyphens.  * The first character must be a letter.  * A name cannot end with a hyphen or contain two consecutive hyphens.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"security_group_i_ds": {
						Description:         "One or more Amazon VPC security groups associated with this replication group.  Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud (Amazon VPC).",
						MarkdownDescription: "One or more Amazon VPC security groups associated with this replication group.  Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud (Amazon VPC).",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"snapshot_ar_ns": {
						Description:         "A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new replication group. The Amazon S3 object name in the ARN cannot contain any commas. The new replication group will have the number of node groups (console: shards) specified by the parameter NumNodeGroups or the number of node groups configured by NodeGroupConfiguration regardless of the number of ARNs specified here.  Example of an Amazon S3 ARN: arn:aws:s3:::my_bucket/snapshot1.rdb",
						MarkdownDescription: "A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new replication group. The Amazon S3 object name in the ARN cannot contain any commas. The new replication group will have the number of node groups (console: shards) specified by the parameter NumNodeGroups or the number of node groups configured by NodeGroupConfiguration regardless of the number of ARNs specified here.  Example of an Amazon S3 ARN: arn:aws:s3:::my_bucket/snapshot1.rdb",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"snapshot_name": {
						Description:         "The name of a snapshot from which to restore data into the new replication group. The snapshot status changes to restoring while the new replication group is being created.",
						MarkdownDescription: "The name of a snapshot from which to restore data into the new replication group. The snapshot status changes to restoring while the new replication group is being created.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"snapshot_retention_limit": {
						Description:         "The number of days for which ElastiCache retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.  Default: 0 (i.e., automatic backups are disabled for this cluster).",
						MarkdownDescription: "The number of days for which ElastiCache retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.  Default: 0 (i.e., automatic backups are disabled for this cluster).",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"snapshot_window": {
						Description:         "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).  Example: 05:00-09:00  If you do not specify this parameter, ElastiCache automatically chooses an appropriate time range.",
						MarkdownDescription: "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).  Example: 05:00-09:00  If you do not specify this parameter, ElastiCache automatically chooses an appropriate time range.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "A list of tags to be added to this resource. Tags are comma-separated key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags as shown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue. Tags on replication groups will be replicated to all nodes.",
						MarkdownDescription: "A list of tags to be added to this resource. Tags are comma-separated key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags as shown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue. Tags on replication groups will be replicated to all nodes.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"transit_encryption_enabled": {
						Description:         "A flag that enables in-transit encryption when set to true.  You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.  This parameter is valid only if the Engine parameter is redis, the EngineVersion parameter is 3.2.6, 4.x or later, and the cluster is being created in an Amazon VPC.  If you enable in-transit encryption, you must also specify a value for CacheSubnetGroup.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.",
						MarkdownDescription: "A flag that enables in-transit encryption when set to true.  You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.  This parameter is valid only if the Engine parameter is redis, the EngineVersion parameter is 3.2.6, 4.x or later, and the cluster is being created in an Amazon VPC.  If you enable in-transit encryption, you must also specify a value for CacheSubnetGroup.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"user_group_i_ds": {
						Description:         "The user group to associate with the replication group.",
						MarkdownDescription: "The user group to associate with the replication group.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_elasticache_services_k8s_aws_replication_group_v1alpha1")

	var state ElasticacheServicesK8SAwsReplicationGroupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ElasticacheServicesK8SAwsReplicationGroupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("elasticache.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ReplicationGroup")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_elasticache_services_k8s_aws_replication_group_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_elasticache_services_k8s_aws_replication_group_v1alpha1")

	var state ElasticacheServicesK8SAwsReplicationGroupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ElasticacheServicesK8SAwsReplicationGroupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("elasticache.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ReplicationGroup")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_elasticache_services_k8s_aws_replication_group_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
