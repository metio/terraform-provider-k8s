/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package elasticache_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource{}
)

func NewElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource() datasource.DataSource {
	return &ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource{}
}

type ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AtRestEncryptionEnabled *bool `tfsdk:"at_rest_encryption_enabled" json:"atRestEncryptionEnabled,omitempty"`
		AuthToken               *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"auth_token" json:"authToken,omitempty"`
		AutomaticFailoverEnabled  *bool     `tfsdk:"automatic_failover_enabled" json:"automaticFailoverEnabled,omitempty"`
		CacheNodeType             *string   `tfsdk:"cache_node_type" json:"cacheNodeType,omitempty"`
		CacheParameterGroupName   *string   `tfsdk:"cache_parameter_group_name" json:"cacheParameterGroupName,omitempty"`
		CacheSecurityGroupNames   *[]string `tfsdk:"cache_security_group_names" json:"cacheSecurityGroupNames,omitempty"`
		CacheSubnetGroupName      *string   `tfsdk:"cache_subnet_group_name" json:"cacheSubnetGroupName,omitempty"`
		DataTieringEnabled        *bool     `tfsdk:"data_tiering_enabled" json:"dataTieringEnabled,omitempty"`
		Description               *string   `tfsdk:"description" json:"description,omitempty"`
		Engine                    *string   `tfsdk:"engine" json:"engine,omitempty"`
		EngineVersion             *string   `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		KmsKeyID                  *string   `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		LogDeliveryConfigurations *[]struct {
			DestinationDetails *struct {
				CloudWatchLogsDetails *struct {
					LogGroup *string `tfsdk:"log_group" json:"logGroup,omitempty"`
				} `tfsdk:"cloud_watch_logs_details" json:"cloudWatchLogsDetails,omitempty"`
				KinesisFirehoseDetails *struct {
					DeliveryStream *string `tfsdk:"delivery_stream" json:"deliveryStream,omitempty"`
				} `tfsdk:"kinesis_firehose_details" json:"kinesisFirehoseDetails,omitempty"`
			} `tfsdk:"destination_details" json:"destinationDetails,omitempty"`
			DestinationType *string `tfsdk:"destination_type" json:"destinationType,omitempty"`
			Enabled         *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			LogFormat       *string `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogType         *string `tfsdk:"log_type" json:"logType,omitempty"`
		} `tfsdk:"log_delivery_configurations" json:"logDeliveryConfigurations,omitempty"`
		MultiAZEnabled         *bool `tfsdk:"multi_az_enabled" json:"multiAZEnabled,omitempty"`
		NodeGroupConfiguration *[]struct {
			NodeGroupID              *string   `tfsdk:"node_group_id" json:"nodeGroupID,omitempty"`
			PrimaryAvailabilityZone  *string   `tfsdk:"primary_availability_zone" json:"primaryAvailabilityZone,omitempty"`
			PrimaryOutpostARN        *string   `tfsdk:"primary_outpost_arn" json:"primaryOutpostARN,omitempty"`
			ReplicaAvailabilityZones *[]string `tfsdk:"replica_availability_zones" json:"replicaAvailabilityZones,omitempty"`
			ReplicaCount             *int64    `tfsdk:"replica_count" json:"replicaCount,omitempty"`
			ReplicaOutpostARNs       *[]string `tfsdk:"replica_outpost_ar_ns" json:"replicaOutpostARNs,omitempty"`
			Slots                    *string   `tfsdk:"slots" json:"slots,omitempty"`
		} `tfsdk:"node_group_configuration" json:"nodeGroupConfiguration,omitempty"`
		NotificationTopicARN       *string   `tfsdk:"notification_topic_arn" json:"notificationTopicARN,omitempty"`
		NumNodeGroups              *int64    `tfsdk:"num_node_groups" json:"numNodeGroups,omitempty"`
		Port                       *int64    `tfsdk:"port" json:"port,omitempty"`
		PreferredCacheClusterAZs   *[]string `tfsdk:"preferred_cache_cluster_a_zs" json:"preferredCacheClusterAZs,omitempty"`
		PreferredMaintenanceWindow *string   `tfsdk:"preferred_maintenance_window" json:"preferredMaintenanceWindow,omitempty"`
		PrimaryClusterID           *string   `tfsdk:"primary_cluster_id" json:"primaryClusterID,omitempty"`
		ReplicasPerNodeGroup       *int64    `tfsdk:"replicas_per_node_group" json:"replicasPerNodeGroup,omitempty"`
		ReplicationGroupID         *string   `tfsdk:"replication_group_id" json:"replicationGroupID,omitempty"`
		SecurityGroupIDs           *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
		SnapshotARNs               *[]string `tfsdk:"snapshot_ar_ns" json:"snapshotARNs,omitempty"`
		SnapshotName               *string   `tfsdk:"snapshot_name" json:"snapshotName,omitempty"`
		SnapshotRetentionLimit     *int64    `tfsdk:"snapshot_retention_limit" json:"snapshotRetentionLimit,omitempty"`
		SnapshotWindow             *string   `tfsdk:"snapshot_window" json:"snapshotWindow,omitempty"`
		Tags                       *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TransitEncryptionEnabled *bool     `tfsdk:"transit_encryption_enabled" json:"transitEncryptionEnabled,omitempty"`
		UserGroupIDs             *[]string `tfsdk:"user_group_i_ds" json:"userGroupIDs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_elasticache_services_k8s_aws_replication_group_v1alpha1"
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ReplicationGroup is the Schema for the ReplicationGroups API",
		MarkdownDescription: "ReplicationGroup is the Schema for the ReplicationGroups API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ReplicationGroupSpec defines the desired state of ReplicationGroup.  Contains all of the attributes of a specific Redis replication group.",
				MarkdownDescription: "ReplicationGroupSpec defines the desired state of ReplicationGroup.  Contains all of the attributes of a specific Redis replication group.",
				Attributes: map[string]schema.Attribute{
					"at_rest_encryption_enabled": schema.BoolAttribute{
						Description:         "A flag that enables encryption at rest when set to true.  You cannot modify the value of AtRestEncryptionEnabled after the replication group is created. To enable encryption at rest on a replication group you must set AtRestEncryptionEnabled to true when you create the replication group.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false",
						MarkdownDescription: "A flag that enables encryption at rest when set to true.  You cannot modify the value of AtRestEncryptionEnabled after the replication group is created. To enable encryption at rest on a replication group you must set AtRestEncryptionEnabled to true when you create the replication group.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"auth_token": schema.SingleNestedAttribute{
						Description:         "Reserved parameter. The password used to access a password protected server.  AuthToken can be specified only on replication groups where TransitEncryptionEnabled is true.  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.  Password constraints:  * Must be only printable ASCII characters.  * Must be at least 16 characters and no more than 128 characters in length.  * The only permitted printable special characters are !, &, #, $, ^, <, >, and -. Other printable special characters cannot be used in the AUTH token.  For more information, see AUTH password (http://redis.io/commands/AUTH) at http://redis.io/commands/AUTH.",
						MarkdownDescription: "Reserved parameter. The password used to access a password protected server.  AuthToken can be specified only on replication groups where TransitEncryptionEnabled is true.  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.  Password constraints:  * Must be only printable ASCII characters.  * Must be at least 16 characters and no more than 128 characters in length.  * The only permitted printable special characters are !, &, #, $, ^, <, >, and -. Other printable special characters cannot be used in the AUTH token.  For more information, see AUTH password (http://redis.io/commands/AUTH) at http://redis.io/commands/AUTH.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Key is the key within the secret",
								MarkdownDescription: "Key is the key within the secret",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"automatic_failover_enabled": schema.BoolAttribute{
						Description:         "Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.  AutomaticFailoverEnabled must be enabled for Redis (cluster mode enabled) replication groups.  Default: false",
						MarkdownDescription: "Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary fails.  AutomaticFailoverEnabled must be enabled for Redis (cluster mode enabled) replication groups.  Default: false",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cache_node_type": schema.StringAttribute{
						Description:         "The compute and memory capacity of the nodes in the node group (shard).  The following node types are supported by ElastiCache. Generally speaking, the current generation types provide more memory and computational power at lower cost when compared to their equivalent previous generation counterparts.  * General purpose: Current generation: M6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward): cache.m6g.large, cache.m6g.xlarge, cache.m6g.2xlarge, cache.m6g.4xlarge, cache.m6g.8xlarge, cache.m6g.12xlarge, cache.m6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) M5 node types: cache.m5.large, cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge, cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge, cache.m4.4xlarge, cache.m4.10xlarge T4g node types (available only for Redis engine version 5.0.6 onward and Memcached engine version 1.5.16 onward): cache.t4g.micro, cache.t4g.small, cache.t4g.medium T3 node types: cache.t3.micro, cache.t3.small, cache.t3.medium T2 node types: cache.t2.micro, cache.t2.small, cache.t2.medium Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) T1 node types: cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large, cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge, cache.m3.2xlarge  * Compute optimized: Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) C1 node types: cache.c1.xlarge  * Memory optimized with data tiering: Current generation: R6gd node types (available only for Redis engine version 6.2 onward). cache.r6gd.xlarge, cache.r6gd.2xlarge, cache.r6gd.4xlarge, cache.r6gd.8xlarge, cache.r6gd.12xlarge, cache.r6gd.16xlarge  * Memory optimized: Current generation: R6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward). cache.r6g.large, cache.r6g.xlarge, cache.r6g.2xlarge, cache.r6g.4xlarge, cache.r6g.8xlarge, cache.r6g.12xlarge, cache.r6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) R5 node types: cache.r5.large, cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge, cache.r5.12xlarge, cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge, cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge, cache.r3.8xlarge  Additional node type info  * All current generation instance types are created in Amazon VPC by default.  * Redis append-only files (AOF) are not supported for T1 or T2 instances.  * Redis Multi-AZ with automatic failover is not supported on T1 instances.  * Redis configuration variables appendonly and appendfsync are not supported on Redis version 2.8.22 and later.",
						MarkdownDescription: "The compute and memory capacity of the nodes in the node group (shard).  The following node types are supported by ElastiCache. Generally speaking, the current generation types provide more memory and computational power at lower cost when compared to their equivalent previous generation counterparts.  * General purpose: Current generation: M6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward): cache.m6g.large, cache.m6g.xlarge, cache.m6g.2xlarge, cache.m6g.4xlarge, cache.m6g.8xlarge, cache.m6g.12xlarge, cache.m6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) M5 node types: cache.m5.large, cache.m5.xlarge, cache.m5.2xlarge, cache.m5.4xlarge, cache.m5.12xlarge, cache.m5.24xlarge M4 node types: cache.m4.large, cache.m4.xlarge, cache.m4.2xlarge, cache.m4.4xlarge, cache.m4.10xlarge T4g node types (available only for Redis engine version 5.0.6 onward and Memcached engine version 1.5.16 onward): cache.t4g.micro, cache.t4g.small, cache.t4g.medium T3 node types: cache.t3.micro, cache.t3.small, cache.t3.medium T2 node types: cache.t2.micro, cache.t2.small, cache.t2.medium Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) T1 node types: cache.t1.micro M1 node types: cache.m1.small, cache.m1.medium, cache.m1.large, cache.m1.xlarge M3 node types: cache.m3.medium, cache.m3.large, cache.m3.xlarge, cache.m3.2xlarge  * Compute optimized: Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) C1 node types: cache.c1.xlarge  * Memory optimized with data tiering: Current generation: R6gd node types (available only for Redis engine version 6.2 onward). cache.r6gd.xlarge, cache.r6gd.2xlarge, cache.r6gd.4xlarge, cache.r6gd.8xlarge, cache.r6gd.12xlarge, cache.r6gd.16xlarge  * Memory optimized: Current generation: R6g node types (available only for Redis engine version 5.0.6 onward and for Memcached engine version 1.5.16 onward). cache.r6g.large, cache.r6g.xlarge, cache.r6g.2xlarge, cache.r6g.4xlarge, cache.r6g.8xlarge, cache.r6g.12xlarge, cache.r6g.16xlarge For region availability, see Supported Node Types (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheNodes.SupportedTypes.html#CacheNodes.SupportedTypesByRegion) R5 node types: cache.r5.large, cache.r5.xlarge, cache.r5.2xlarge, cache.r5.4xlarge, cache.r5.12xlarge, cache.r5.24xlarge R4 node types: cache.r4.large, cache.r4.xlarge, cache.r4.2xlarge, cache.r4.4xlarge, cache.r4.8xlarge, cache.r4.16xlarge Previous generation: (not recommended. Existing clusters are still supported but creation of new clusters is not supported for these types.) M2 node types: cache.m2.xlarge, cache.m2.2xlarge, cache.m2.4xlarge R3 node types: cache.r3.large, cache.r3.xlarge, cache.r3.2xlarge, cache.r3.4xlarge, cache.r3.8xlarge  Additional node type info  * All current generation instance types are created in Amazon VPC by default.  * Redis append-only files (AOF) are not supported for T1 or T2 instances.  * Redis Multi-AZ with automatic failover is not supported on T1 instances.  * Redis configuration variables appendonly and appendfsync are not supported on Redis version 2.8.22 and later.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cache_parameter_group_name": schema.StringAttribute{
						Description:         "The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.  If you are running Redis version 3.2.4 or later, only one node group (shard), and want to use a default parameter group, we recommend that you specify the parameter group by name.  * To create a Redis (cluster mode disabled) replication group, use CacheParameterGroupName=default.redis3.2.  * To create a Redis (cluster mode enabled) replication group, use CacheParameterGroupName=default.redis3.2.cluster.on.",
						MarkdownDescription: "The name of the parameter group to associate with this replication group. If this argument is omitted, the default cache parameter group for the specified engine is used.  If you are running Redis version 3.2.4 or later, only one node group (shard), and want to use a default parameter group, we recommend that you specify the parameter group by name.  * To create a Redis (cluster mode disabled) replication group, use CacheParameterGroupName=default.redis3.2.  * To create a Redis (cluster mode enabled) replication group, use CacheParameterGroupName=default.redis3.2.cluster.on.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cache_security_group_names": schema.ListAttribute{
						Description:         "A list of cache security group names to associate with this replication group.",
						MarkdownDescription: "A list of cache security group names to associate with this replication group.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cache_subnet_group_name": schema.StringAttribute{
						Description:         "The name of the cache subnet group to be used for the replication group.  If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start creating a cluster. For more information, see Subnets and Subnet Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html).",
						MarkdownDescription: "The name of the cache subnet group to be used for the replication group.  If you're going to launch your cluster in an Amazon VPC, you need to create a subnet group before you start creating a cluster. For more information, see Subnets and Subnet Groups (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SubnetGroups.html).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"data_tiering_enabled": schema.BoolAttribute{
						Description:         "Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes. For more information, see Data tiering (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/data-tiering.html).",
						MarkdownDescription: "Enables data tiering. Data tiering is only supported for replication groups using the r6gd node type. This parameter must be set to true when using r6gd nodes. For more information, see Data tiering (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/data-tiering.html).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"description": schema.StringAttribute{
						Description:         "A user-created description for the replication group.",
						MarkdownDescription: "A user-created description for the replication group.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"engine": schema.StringAttribute{
						Description:         "The name of the cache engine to be used for the clusters in this replication group. Must be Redis.",
						MarkdownDescription: "The name of the cache engine to be used for the clusters in this replication group. Must be Redis.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The version number of the cache engine to be used for the clusters in this replication group. To view the supported cache engine versions, use the DescribeCacheEngineVersions operation.  Important: You can upgrade to a newer engine version (see Selecting a Cache Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)) in the ElastiCache User Guide, but you cannot downgrade to an earlier engine version. If you want to use an earlier engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine version.",
						MarkdownDescription: "The version number of the cache engine to be used for the clusters in this replication group. To view the supported cache engine versions, use the DescribeCacheEngineVersions operation.  Important: You can upgrade to a newer engine version (see Selecting a Cache Engine and Version (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/SelectEngine.html#VersionManagement)) in the ElastiCache User Guide, but you cannot downgrade to an earlier engine version. If you want to use an earlier engine version, you must delete the existing cluster or replication group and create it anew with the earlier engine version.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The ID of the KMS key used to encrypt the disk in the cluster.",
						MarkdownDescription: "The ID of the KMS key used to encrypt the disk in the cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"log_delivery_configurations": schema.ListNestedAttribute{
						Description:         "Specifies the destination, format and type of the logs.",
						MarkdownDescription: "Specifies the destination, format and type of the logs.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"destination_details": schema.SingleNestedAttribute{
									Description:         "Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.",
									MarkdownDescription: "Configuration details of either a CloudWatch Logs destination or Kinesis Data Firehose destination.",
									Attributes: map[string]schema.Attribute{
										"cloud_watch_logs_details": schema.SingleNestedAttribute{
											Description:         "The configuration details of the CloudWatch Logs destination.",
											MarkdownDescription: "The configuration details of the CloudWatch Logs destination.",
											Attributes: map[string]schema.Attribute{
												"log_group": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"kinesis_firehose_details": schema.SingleNestedAttribute{
											Description:         "The configuration details of the Kinesis Data Firehose destination.",
											MarkdownDescription: "The configuration details of the Kinesis Data Firehose destination.",
											Attributes: map[string]schema.Attribute{
												"delivery_stream": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"destination_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"enabled": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"log_format": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"log_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"multi_az_enabled": schema.BoolAttribute{
						Description:         "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For more information, see Minimizing Downtime: Multi-AZ (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html).",
						MarkdownDescription: "A flag indicating if you have Multi-AZ enabled to enhance fault tolerance. For more information, see Minimizing Downtime: Multi-AZ (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_group_configuration": schema.ListNestedAttribute{
						Description:         "A list of node group (shard) configuration options. Each node group (shard) configuration has the following members: PrimaryAvailabilityZone, ReplicaAvailabilityZones, ReplicaCount, and Slots.  If you're creating a Redis (cluster mode disabled) or a Redis (cluster mode enabled) replication group, you can use this parameter to individually configure each node group (shard), or you can omit this parameter. However, it is required when seeding a Redis (cluster mode enabled) cluster from a S3 rdb file. You must configure each node group (shard) using this parameter because you must specify the slots for each node group.",
						MarkdownDescription: "A list of node group (shard) configuration options. Each node group (shard) configuration has the following members: PrimaryAvailabilityZone, ReplicaAvailabilityZones, ReplicaCount, and Slots.  If you're creating a Redis (cluster mode disabled) or a Redis (cluster mode enabled) replication group, you can use this parameter to individually configure each node group (shard), or you can omit this parameter. However, it is required when seeding a Redis (cluster mode enabled) cluster from a S3 rdb file. You must configure each node group (shard) using this parameter because you must specify the slots for each node group.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"node_group_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"primary_availability_zone": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"primary_outpost_arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"replica_availability_zones": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"replica_count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"replica_outpost_ar_ns": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"slots": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"notification_topic_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.  The Amazon SNS topic owner must be the same as the cluster owner.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (SNS) topic to which notifications are sent.  The Amazon SNS topic owner must be the same as the cluster owner.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"num_node_groups": schema.Int64Attribute{
						Description:         "An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group. For Redis (cluster mode disabled) either omit this parameter or set it to 1.  Default: 1",
						MarkdownDescription: "An optional parameter that specifies the number of node groups (shards) for this Redis (cluster mode enabled) replication group. For Redis (cluster mode disabled) either omit this parameter or set it to 1.  Default: 1",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"port": schema.Int64Attribute{
						Description:         "The port number on which each member of the replication group accepts connections.",
						MarkdownDescription: "The port number on which each member of the replication group accepts connections.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"preferred_cache_cluster_a_zs": schema.ListAttribute{
						Description:         "A list of EC2 Availability Zones in which the replication group's clusters are created. The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list.  This parameter is not used if there is more than one node group (shard). You should use NodeGroupConfiguration instead.  If you are creating your replication group in an Amazon VPC (recommended), you can only locate clusters in Availability Zones associated with the subnets in the selected subnet group.  The number of Availability Zones listed must equal the value of NumCacheClusters.  Default: system chosen Availability Zones.",
						MarkdownDescription: "A list of EC2 Availability Zones in which the replication group's clusters are created. The order of the Availability Zones in the list is the order in which clusters are allocated. The primary cluster is created in the first AZ in the list.  This parameter is not used if there is more than one node group (shard). You should use NodeGroupConfiguration instead.  If you are creating your replication group in an Amazon VPC (recommended), you can only locate clusters in Availability Zones associated with the subnets in the selected subnet group.  The number of Availability Zones listed must equal the value of NumCacheClusters.  Default: system chosen Availability Zones.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"preferred_maintenance_window": schema.StringAttribute{
						Description:         "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid values for ddd are:  Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.  Valid values for ddd are:  * sun  * mon  * tue  * wed  * thu  * fri  * sat  Example: sun:23:00-mon:01:30",
						MarkdownDescription: "Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Valid values for ddd are:  Specifies the weekly time range during which maintenance on the cluster is performed. It is specified as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period.  Valid values for ddd are:  * sun  * mon  * tue  * wed  * thu  * fri  * sat  Example: sun:23:00-mon:01:30",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"primary_cluster_id": schema.StringAttribute{
						Description:         "The identifier of the cluster that serves as the primary for this replication group. This cluster must already exist and have a status of available.  This parameter is not required if NumCacheClusters, NumNodeGroups, or ReplicasPerNodeGroup is specified.",
						MarkdownDescription: "The identifier of the cluster that serves as the primary for this replication group. This cluster must already exist and have a status of available.  This parameter is not required if NumCacheClusters, NumNodeGroups, or ReplicasPerNodeGroup is specified.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replicas_per_node_group": schema.Int64Attribute{
						Description:         "An optional parameter that specifies the number of replica nodes in each node group (shard). Valid values are 0 to 5.",
						MarkdownDescription: "An optional parameter that specifies the number of replica nodes in each node group (shard). Valid values are 0 to 5.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replication_group_id": schema.StringAttribute{
						Description:         "The replication group identifier. This parameter is stored as a lowercase string.  Constraints:  * A name must contain from 1 to 40 alphanumeric characters or hyphens.  * The first character must be a letter.  * A name cannot end with a hyphen or contain two consecutive hyphens.",
						MarkdownDescription: "The replication group identifier. This parameter is stored as a lowercase string.  Constraints:  * A name must contain from 1 to 40 alphanumeric characters or hyphens.  * The first character must be a letter.  * A name cannot end with a hyphen or contain two consecutive hyphens.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"security_group_i_ds": schema.ListAttribute{
						Description:         "One or more Amazon VPC security groups associated with this replication group.  Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud (Amazon VPC).",
						MarkdownDescription: "One or more Amazon VPC security groups associated with this replication group.  Use this parameter only when you are creating a replication group in an Amazon Virtual Private Cloud (Amazon VPC).",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"snapshot_ar_ns": schema.ListAttribute{
						Description:         "A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new replication group. The Amazon S3 object name in the ARN cannot contain any commas. The new replication group will have the number of node groups (console: shards) specified by the parameter NumNodeGroups or the number of node groups configured by NodeGroupConfiguration regardless of the number of ARNs specified here.  Example of an Amazon S3 ARN: arn:aws:s3:::my_bucket/snapshot1.rdb",
						MarkdownDescription: "A list of Amazon Resource Names (ARN) that uniquely identify the Redis RDB snapshot files stored in Amazon S3. The snapshot files are used to populate the new replication group. The Amazon S3 object name in the ARN cannot contain any commas. The new replication group will have the number of node groups (console: shards) specified by the parameter NumNodeGroups or the number of node groups configured by NodeGroupConfiguration regardless of the number of ARNs specified here.  Example of an Amazon S3 ARN: arn:aws:s3:::my_bucket/snapshot1.rdb",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"snapshot_name": schema.StringAttribute{
						Description:         "The name of a snapshot from which to restore data into the new replication group. The snapshot status changes to restoring while the new replication group is being created.",
						MarkdownDescription: "The name of a snapshot from which to restore data into the new replication group. The snapshot status changes to restoring while the new replication group is being created.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"snapshot_retention_limit": schema.Int64Attribute{
						Description:         "The number of days for which ElastiCache retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.  Default: 0 (i.e., automatic backups are disabled for this cluster).",
						MarkdownDescription: "The number of days for which ElastiCache retains automatic snapshots before deleting them. For example, if you set SnapshotRetentionLimit to 5, a snapshot that was taken today is retained for 5 days before being deleted.  Default: 0 (i.e., automatic backups are disabled for this cluster).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"snapshot_window": schema.StringAttribute{
						Description:         "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).  Example: 05:00-09:00  If you do not specify this parameter, ElastiCache automatically chooses an appropriate time range.",
						MarkdownDescription: "The daily time range (in UTC) during which ElastiCache begins taking a daily snapshot of your node group (shard).  Example: 05:00-09:00  If you do not specify this parameter, ElastiCache automatically chooses an appropriate time range.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of tags to be added to this resource. Tags are comma-separated key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags as shown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue. Tags on replication groups will be replicated to all nodes.",
						MarkdownDescription: "A list of tags to be added to this resource. Tags are comma-separated key,value pairs (e.g. Key=myKey, Value=myKeyValue. You can include multiple tags as shown following: Key=myKey, Value=myKeyValue Key=mySecondKey, Value=mySecondKeyValue. Tags on replication groups will be replicated to all nodes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"transit_encryption_enabled": schema.BoolAttribute{
						Description:         "A flag that enables in-transit encryption when set to true.  You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.  This parameter is valid only if the Engine parameter is redis, the EngineVersion parameter is 3.2.6, 4.x or later, and the cluster is being created in an Amazon VPC.  If you enable in-transit encryption, you must also specify a value for CacheSubnetGroup.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.",
						MarkdownDescription: "A flag that enables in-transit encryption when set to true.  You cannot modify the value of TransitEncryptionEnabled after the cluster is created. To enable in-transit encryption on a cluster you must set TransitEncryptionEnabled to true when you create a cluster.  This parameter is valid only if the Engine parameter is redis, the EngineVersion parameter is 3.2.6, 4.x or later, and the cluster is being created in an Amazon VPC.  If you enable in-transit encryption, you must also specify a value for CacheSubnetGroup.  Required: Only available when creating a replication group in an Amazon VPC using redis version 3.2.6, 4.x or later.  Default: false  For HIPAA compliance, you must specify TransitEncryptionEnabled as true, an AuthToken, and a CacheSubnetGroup.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"user_group_i_ds": schema.ListAttribute{
						Description:         "The user group to associate with the replication group.",
						MarkdownDescription: "The user group to associate with the replication group.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_elasticache_services_k8s_aws_replication_group_v1alpha1")

	var data ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "elasticache.services.k8s.aws", Version: "v1alpha1", Resource: "replicationgroups"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ElasticacheServicesK8SAwsReplicationGroupV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("elasticache.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("ReplicationGroup")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
