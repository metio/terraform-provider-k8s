/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rds_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &RdsServicesK8SAwsDbclusterV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &RdsServicesK8SAwsDbclusterV1Alpha1DataSource{}
)

func NewRdsServicesK8SAwsDbclusterV1Alpha1DataSource() datasource.DataSource {
	return &RdsServicesK8SAwsDbclusterV1Alpha1DataSource{}
}

type RdsServicesK8SAwsDbclusterV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type RdsServicesK8SAwsDbclusterV1Alpha1DataSourceData struct {
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
		AllocatedStorage            *int64    `tfsdk:"allocated_storage" json:"allocatedStorage,omitempty"`
		AutoMinorVersionUpgrade     *bool     `tfsdk:"auto_minor_version_upgrade" json:"autoMinorVersionUpgrade,omitempty"`
		AvailabilityZones           *[]string `tfsdk:"availability_zones" json:"availabilityZones,omitempty"`
		BacktrackWindow             *int64    `tfsdk:"backtrack_window" json:"backtrackWindow,omitempty"`
		BackupRetentionPeriod       *int64    `tfsdk:"backup_retention_period" json:"backupRetentionPeriod,omitempty"`
		CharacterSetName            *string   `tfsdk:"character_set_name" json:"characterSetName,omitempty"`
		CopyTagsToSnapshot          *bool     `tfsdk:"copy_tags_to_snapshot" json:"copyTagsToSnapshot,omitempty"`
		DatabaseName                *string   `tfsdk:"database_name" json:"databaseName,omitempty"`
		DbClusterIdentifier         *string   `tfsdk:"db_cluster_identifier" json:"dbClusterIdentifier,omitempty"`
		DbClusterInstanceClass      *string   `tfsdk:"db_cluster_instance_class" json:"dbClusterInstanceClass,omitempty"`
		DbClusterParameterGroupName *string   `tfsdk:"db_cluster_parameter_group_name" json:"dbClusterParameterGroupName,omitempty"`
		DbClusterParameterGroupRef  *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"db_cluster_parameter_group_ref" json:"dbClusterParameterGroupRef,omitempty"`
		DbSubnetGroupName *string `tfsdk:"db_subnet_group_name" json:"dbSubnetGroupName,omitempty"`
		DbSubnetGroupRef  *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"db_subnet_group_ref" json:"dbSubnetGroupRef,omitempty"`
		DbSystemID                      *string   `tfsdk:"db_system_id" json:"dbSystemID,omitempty"`
		DeletionProtection              *bool     `tfsdk:"deletion_protection" json:"deletionProtection,omitempty"`
		DestinationRegion               *string   `tfsdk:"destination_region" json:"destinationRegion,omitempty"`
		Domain                          *string   `tfsdk:"domain" json:"domain,omitempty"`
		DomainIAMRoleName               *string   `tfsdk:"domain_iam_role_name" json:"domainIAMRoleName,omitempty"`
		EnableCloudwatchLogsExports     *[]string `tfsdk:"enable_cloudwatch_logs_exports" json:"enableCloudwatchLogsExports,omitempty"`
		EnableGlobalWriteForwarding     *bool     `tfsdk:"enable_global_write_forwarding" json:"enableGlobalWriteForwarding,omitempty"`
		EnableHTTPEndpoint              *bool     `tfsdk:"enable_http_endpoint" json:"enableHTTPEndpoint,omitempty"`
		EnableIAMDatabaseAuthentication *bool     `tfsdk:"enable_iam_database_authentication" json:"enableIAMDatabaseAuthentication,omitempty"`
		EnablePerformanceInsights       *bool     `tfsdk:"enable_performance_insights" json:"enablePerformanceInsights,omitempty"`
		Engine                          *string   `tfsdk:"engine" json:"engine,omitempty"`
		EngineMode                      *string   `tfsdk:"engine_mode" json:"engineMode,omitempty"`
		EngineVersion                   *string   `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		GlobalClusterIdentifier         *string   `tfsdk:"global_cluster_identifier" json:"globalClusterIdentifier,omitempty"`
		Iops                            *int64    `tfsdk:"iops" json:"iops,omitempty"`
		KmsKeyID                        *string   `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		KmsKeyRef                       *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"kms_key_ref" json:"kmsKeyRef,omitempty"`
		ManageMasterUserPassword *bool `tfsdk:"manage_master_user_password" json:"manageMasterUserPassword,omitempty"`
		MasterUserPassword       *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"master_user_password" json:"masterUserPassword,omitempty"`
		MasterUserSecretKMSKeyID  *string `tfsdk:"master_user_secret_kms_key_id" json:"masterUserSecretKMSKeyID,omitempty"`
		MasterUserSecretKMSKeyRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"master_user_secret_kms_key_ref" json:"masterUserSecretKMSKeyRef,omitempty"`
		MasterUsername                     *string `tfsdk:"master_username" json:"masterUsername,omitempty"`
		MonitoringInterval                 *int64  `tfsdk:"monitoring_interval" json:"monitoringInterval,omitempty"`
		MonitoringRoleARN                  *string `tfsdk:"monitoring_role_arn" json:"monitoringRoleARN,omitempty"`
		NetworkType                        *string `tfsdk:"network_type" json:"networkType,omitempty"`
		OptionGroupName                    *string `tfsdk:"option_group_name" json:"optionGroupName,omitempty"`
		PerformanceInsightsKMSKeyID        *string `tfsdk:"performance_insights_kms_key_id" json:"performanceInsightsKMSKeyID,omitempty"`
		PerformanceInsightsRetentionPeriod *int64  `tfsdk:"performance_insights_retention_period" json:"performanceInsightsRetentionPeriod,omitempty"`
		Port                               *int64  `tfsdk:"port" json:"port,omitempty"`
		PreSignedURL                       *string `tfsdk:"pre_signed_url" json:"preSignedURL,omitempty"`
		PreferredBackupWindow              *string `tfsdk:"preferred_backup_window" json:"preferredBackupWindow,omitempty"`
		PreferredMaintenanceWindow         *string `tfsdk:"preferred_maintenance_window" json:"preferredMaintenanceWindow,omitempty"`
		PubliclyAccessible                 *bool   `tfsdk:"publicly_accessible" json:"publiclyAccessible,omitempty"`
		ReplicationSourceIdentifier        *string `tfsdk:"replication_source_identifier" json:"replicationSourceIdentifier,omitempty"`
		ScalingConfiguration               *struct {
			AutoPause             *bool   `tfsdk:"auto_pause" json:"autoPause,omitempty"`
			MaxCapacity           *int64  `tfsdk:"max_capacity" json:"maxCapacity,omitempty"`
			MinCapacity           *int64  `tfsdk:"min_capacity" json:"minCapacity,omitempty"`
			SecondsBeforeTimeout  *int64  `tfsdk:"seconds_before_timeout" json:"secondsBeforeTimeout,omitempty"`
			SecondsUntilAutoPause *int64  `tfsdk:"seconds_until_auto_pause" json:"secondsUntilAutoPause,omitempty"`
			TimeoutAction         *string `tfsdk:"timeout_action" json:"timeoutAction,omitempty"`
		} `tfsdk:"scaling_configuration" json:"scalingConfiguration,omitempty"`
		ServerlessV2ScalingConfiguration *struct {
			MaxCapacity *float64 `tfsdk:"max_capacity" json:"maxCapacity,omitempty"`
			MinCapacity *float64 `tfsdk:"min_capacity" json:"minCapacity,omitempty"`
		} `tfsdk:"serverless_v2_scaling_configuration" json:"serverlessV2ScalingConfiguration,omitempty"`
		SnapshotIdentifier *string `tfsdk:"snapshot_identifier" json:"snapshotIdentifier,omitempty"`
		SourceRegion       *string `tfsdk:"source_region" json:"sourceRegion,omitempty"`
		StorageEncrypted   *bool   `tfsdk:"storage_encrypted" json:"storageEncrypted,omitempty"`
		StorageType        *string `tfsdk:"storage_type" json:"storageType,omitempty"`
		Tags               *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcSecurityGroupIDs  *[]string `tfsdk:"vpc_security_group_i_ds" json:"vpcSecurityGroupIDs,omitempty"`
		VpcSecurityGroupRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"vpc_security_group_refs" json:"vpcSecurityGroupRefs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RdsServicesK8SAwsDbclusterV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_db_cluster_v1alpha1"
}

func (r *RdsServicesK8SAwsDbclusterV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DBCluster is the Schema for the DBClusters API",
		MarkdownDescription: "DBCluster is the Schema for the DBClusters API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "DBClusterSpec defines the desired state of DBCluster.  Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster.  For an Amazon Aurora DB cluster, this data type is used as a response element in the operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster, ModifyDBCluster, PromoteReadReplicaDBCluster, RestoreDBClusterFromS3, RestoreDBClusterFromSnapshot, RestoreDBClusterToPointInTime, StartDBCluster, and StopDBCluster.  For a Multi-AZ DB cluster, this data type is used as a response element in the operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster, ModifyDBCluster, RebootDBCluster, RestoreDBClusterFromSnapshot, and RestoreDBClusterToPointInTime.  For more information on Amazon Aurora DB clusters, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html) in the Amazon Aurora User Guide.  For more information on Multi-AZ DB clusters, see Multi-AZ deployments with two readable standby DB instances (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the Amazon RDS User Guide.",
				MarkdownDescription: "DBClusterSpec defines the desired state of DBCluster.  Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster.  For an Amazon Aurora DB cluster, this data type is used as a response element in the operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster, ModifyDBCluster, PromoteReadReplicaDBCluster, RestoreDBClusterFromS3, RestoreDBClusterFromSnapshot, RestoreDBClusterToPointInTime, StartDBCluster, and StopDBCluster.  For a Multi-AZ DB cluster, this data type is used as a response element in the operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster, ModifyDBCluster, RebootDBCluster, RestoreDBClusterFromSnapshot, and RestoreDBClusterToPointInTime.  For more information on Amazon Aurora DB clusters, see What is Amazon Aurora? (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html) in the Amazon Aurora User Guide.  For more information on Multi-AZ DB clusters, see Multi-AZ deployments with two readable standby DB instances (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the Amazon RDS User Guide.",
				Attributes: map[string]schema.Attribute{
					"allocated_storage": schema.Int64Attribute{
						Description:         "The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.  This setting is required to create a Multi-AZ DB cluster.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "The amount of storage in gibibytes (GiB) to allocate to each DB instance in the Multi-AZ DB cluster.  This setting is required to create a Multi-AZ DB cluster.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"auto_minor_version_upgrade": schema.BoolAttribute{
						Description:         "A value that indicates whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window. By default, minor engine upgrades are applied automatically.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "A value that indicates whether minor engine upgrades are applied automatically to the DB cluster during the maintenance window. By default, minor engine upgrades are applied automatically.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"availability_zones": schema.ListAttribute{
						Description:         "A list of Availability Zones (AZs) where DB instances in the DB cluster can be created.  For information on Amazon Web Services Regions and Availability Zones, see Choosing the Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "A list of Availability Zones (AZs) where DB instances in the DB cluster can be created.  For information on Amazon Web Services Regions and Availability Zones, see Choosing the Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.RegionsAndAvailabilityZones.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"backtrack_window": schema.Int64Attribute{
						Description:         "The target backtrack window, in seconds. To disable backtracking, set this value to 0.  Default: 0  Constraints:  * If specified, this value must be set to a number from 0 to 259,200 (72 hours).  Valid for: Aurora MySQL DB clusters only",
						MarkdownDescription: "The target backtrack window, in seconds. To disable backtracking, set this value to 0.  Default: 0  Constraints:  * If specified, this value must be set to a number from 0 to 259,200 (72 hours).  Valid for: Aurora MySQL DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"backup_retention_period": schema.Int64Attribute{
						Description:         "The number of days for which automated backups are retained.  Default: 1  Constraints:  * Must be a value from 1 to 35  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The number of days for which automated backups are retained.  Default: 1  Constraints:  * Must be a value from 1 to 35  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"character_set_name": schema.StringAttribute{
						Description:         "A value that indicates that the DB cluster should be associated with the specified CharacterSet.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "A value that indicates that the DB cluster should be associated with the specified CharacterSet.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"copy_tags_to_snapshot": schema.BoolAttribute{
						Description:         "A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default is not to copy them.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "A value that indicates whether to copy all tags from the DB cluster to snapshots of the DB cluster. The default is not to copy them.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"database_name": schema.StringAttribute{
						Description:         "The name for your database of up to 64 alphanumeric characters. If you do not provide a name, Amazon RDS doesn't create a database in the DB cluster you are creating.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The name for your database of up to 64 alphanumeric characters. If you do not provide a name, Amazon RDS doesn't create a database in the DB cluster you are creating.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"db_cluster_identifier": schema.StringAttribute{
						Description:         "The DB cluster identifier. This parameter is stored as a lowercase string.  Constraints:  * Must contain from 1 to 63 letters, numbers, or hyphens.  * First character must be a letter.  * Can't end with a hyphen or contain two consecutive hyphens.  Example: my-cluster1  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The DB cluster identifier. This parameter is stored as a lowercase string.  Constraints:  * Must contain from 1 to 63 letters, numbers, or hyphens.  * First character must be a letter.  * Can't end with a hyphen or contain two consecutive hyphens.  Example: my-cluster1  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"db_cluster_instance_class": schema.StringAttribute{
						Description:         "The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example db.m6gd.xlarge. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.  For the full list of DB instance classes and availability for your engine, see DB instance class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide.  This setting is required to create a Multi-AZ DB cluster.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "The compute and memory capacity of each DB instance in the Multi-AZ DB cluster, for example db.m6gd.xlarge. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines.  For the full list of DB instance classes and availability for your engine, see DB instance class (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide.  This setting is required to create a Multi-AZ DB cluster.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"db_cluster_parameter_group_name": schema.StringAttribute{
						Description:         "The name of the DB cluster parameter group to associate with this DB cluster. If you do not specify a value, then the default DB cluster parameter group for the specified DB engine and version is used.  Constraints:  * If supplied, must match the name of an existing DB cluster parameter group.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The name of the DB cluster parameter group to associate with this DB cluster. If you do not specify a value, then the default DB cluster parameter group for the specified DB engine and version is used.  Constraints:  * If supplied, must match the name of an existing DB cluster parameter group.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"db_cluster_parameter_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"db_subnet_group_name": schema.StringAttribute{
						Description:         "A DB subnet group to associate with this DB cluster.  This setting is required to create a Multi-AZ DB cluster.  Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.  Example: mydbsubnetgroup  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "A DB subnet group to associate with this DB cluster.  This setting is required to create a Multi-AZ DB cluster.  Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.  Example: mydbsubnetgroup  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"db_subnet_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"db_system_id": schema.StringAttribute{
						Description:         "Reserved for future use.",
						MarkdownDescription: "Reserved for future use.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"deletion_protection": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "A value that indicates whether the DB cluster has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"destination_region": schema.StringAttribute{
						Description:         "DestinationRegion is used for presigning the request to a given region.",
						MarkdownDescription: "DestinationRegion is used for presigning the request to a given region.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"domain": schema.StringAttribute{
						Description:         "The Active Directory directory ID to create the DB cluster in.  For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.  For more information, see Kerberos authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "The Active Directory directory ID to create the DB cluster in.  For Amazon Aurora DB clusters, Amazon RDS can use Kerberos authentication to authenticate users that connect to the DB cluster.  For more information, see Kerberos authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/kerberos-authentication.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"domain_iam_role_name": schema.StringAttribute{
						Description:         "Specify the name of the IAM role to be used when making API calls to the Directory Service.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "Specify the name of the IAM role to be used when making API calls to the Directory Service.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_cloudwatch_logs_exports": schema.ListAttribute{
						Description:         "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used.  RDS for MySQL  Possible values are error, general, and slowquery.  RDS for PostgreSQL  Possible values are postgresql and upgrade.  Aurora MySQL  Possible values are audit, error, general, and slowquery.  Aurora PostgreSQL  Possible value is postgresql.  For more information about exporting CloudWatch Logs for Amazon RDS, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide.  For more information about exporting CloudWatch Logs for Amazon Aurora, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine being used.  RDS for MySQL  Possible values are error, general, and slowquery.  RDS for PostgreSQL  Possible values are postgresql and upgrade.  Aurora MySQL  Possible values are audit, error, general, and slowquery.  Aurora PostgreSQL  Possible value is postgresql.  For more information about exporting CloudWatch Logs for Amazon RDS, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide.  For more information about exporting CloudWatch Logs for Amazon Aurora, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_global_write_forwarding": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable this DB cluster to forward write operations to the primary cluster of an Aurora global database (GlobalCluster). By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database.  You can set this value only on Aurora DB clusters that are members of an Aurora global database. With this parameter enabled, a secondary cluster can forward writes to the current primary cluster and the resulting changes are replicated back to this cluster. For the primary DB cluster of an Aurora global database, this value is used immediately if the primary is demoted by the FailoverGlobalCluster API operation, but it does nothing until then.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "A value that indicates whether to enable this DB cluster to forward write operations to the primary cluster of an Aurora global database (GlobalCluster). By default, write operations are not allowed on Aurora DB clusters that are secondary clusters in an Aurora global database.  You can set this value only on Aurora DB clusters that are members of an Aurora global database. With this parameter enabled, a secondary cluster can forward writes to the current primary cluster and the resulting changes are replicated back to this cluster. For the primary DB cluster of an Aurora global database, this value is used immediately if the primary is demoted by the FailoverGlobalCluster API operation, but it does nothing until then.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_http_endpoint": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable the HTTP endpoint for an Aurora Serverless v1 DB cluster. By default, the HTTP endpoint is disabled.  When enabled, the HTTP endpoint provides a connectionless web service API for running SQL queries on the Aurora Serverless v1 DB cluster. You can also query your database from inside the RDS console with the query editor.  For more information, see Using the Data API for Aurora Serverless v1 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "A value that indicates whether to enable the HTTP endpoint for an Aurora Serverless v1 DB cluster. By default, the HTTP endpoint is disabled.  When enabled, the HTTP endpoint provides a connectionless web service API for running SQL queries on the Aurora Serverless v1 DB cluster. You can also query your database from inside the RDS console with the query editor.  For more information, see Using the Data API for Aurora Serverless v1 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/data-api.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_iam_database_authentication": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled.  For more information, see IAM Database Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "A value that indicates whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled.  For more information, see IAM Database Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_performance_insights": schema.BoolAttribute{
						Description:         "A value that indicates whether to turn on Performance Insights for the DB cluster.  For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "A value that indicates whether to turn on Performance Insights for the DB cluster.  For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"engine": schema.StringAttribute{
						Description:         "The name of the database engine to be used for this DB cluster.  Valid Values:  * aurora-mysql  * aurora-postgresql  * mysql  * postgres  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The name of the database engine to be used for this DB cluster.  Valid Values:  * aurora-mysql  * aurora-postgresql  * mysql  * postgres  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"engine_mode": schema.StringAttribute{
						Description:         "The DB engine mode of the DB cluster, either provisioned or serverless.  The serverless engine mode only applies for Aurora Serverless v1 DB clusters.  Limitations and requirements apply to some DB engine modes. For more information, see the following sections in the Amazon Aurora User Guide:  * Limitations of Aurora Serverless v1 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)  * Requirements for Aurora Serverless v2 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.requirements.html)  * Limitations of parallel query (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)  * Limitations of Aurora global databases (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)  Valid for: Aurora DB clusters only",
						MarkdownDescription: "The DB engine mode of the DB cluster, either provisioned or serverless.  The serverless engine mode only applies for Aurora Serverless v1 DB clusters.  Limitations and requirements apply to some DB engine modes. For more information, see the following sections in the Amazon Aurora User Guide:  * Limitations of Aurora Serverless v1 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.html#aurora-serverless.limitations)  * Requirements for Aurora Serverless v2 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.requirements.html)  * Limitations of parallel query (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-parallel-query.html#aurora-mysql-parallel-query-limitations)  * Limitations of Aurora global databases (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database.limitations)  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The version number of the database engine to use.  To list all of the available engine versions for Aurora MySQL version 2 (5.7-compatible) and version 3 (MySQL 8.0-compatible), use the following command:  aws rds describe-db-engine-versions --engine aurora-mysql --query 'DBEngineVersions[].EngineVersion'  You can supply either 5.7 or 8.0 to use the default engine version for Aurora MySQL version 2 or version 3, respectively.  To list all of the available engine versions for Aurora PostgreSQL, use the following command:  aws rds describe-db-engine-versions --engine aurora-postgresql --query 'DBEngineVersions[].EngineVersion'  To list all of the available engine versions for RDS for MySQL, use the following command:  aws rds describe-db-engine-versions --engine mysql --query 'DBEngineVersions[].EngineVersion'  To list all of the available engine versions for RDS for PostgreSQL, use the following command:  aws rds describe-db-engine-versions --engine postgres --query 'DBEngineVersions[].EngineVersion'  Aurora MySQL  For information, see Database engine updates for Amazon Aurora MySQL (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) in the Amazon Aurora User Guide.  Aurora PostgreSQL  For information, see Amazon Aurora PostgreSQL releases and engine versions (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html) in the Amazon Aurora User Guide.  MySQL  For information, see Amazon RDS for MySQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide.  PostgreSQL  For information, see Amazon RDS for PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The version number of the database engine to use.  To list all of the available engine versions for Aurora MySQL version 2 (5.7-compatible) and version 3 (MySQL 8.0-compatible), use the following command:  aws rds describe-db-engine-versions --engine aurora-mysql --query 'DBEngineVersions[].EngineVersion'  You can supply either 5.7 or 8.0 to use the default engine version for Aurora MySQL version 2 or version 3, respectively.  To list all of the available engine versions for Aurora PostgreSQL, use the following command:  aws rds describe-db-engine-versions --engine aurora-postgresql --query 'DBEngineVersions[].EngineVersion'  To list all of the available engine versions for RDS for MySQL, use the following command:  aws rds describe-db-engine-versions --engine mysql --query 'DBEngineVersions[].EngineVersion'  To list all of the available engine versions for RDS for PostgreSQL, use the following command:  aws rds describe-db-engine-versions --engine postgres --query 'DBEngineVersions[].EngineVersion'  Aurora MySQL  For information, see Database engine updates for Amazon Aurora MySQL (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Updates.html) in the Amazon Aurora User Guide.  Aurora PostgreSQL  For information, see Amazon Aurora PostgreSQL releases and engine versions (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraPostgreSQL.Updates.20180305.html) in the Amazon Aurora User Guide.  MySQL  For information, see Amazon RDS for MySQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide.  PostgreSQL  For information, see Amazon RDS for PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"global_cluster_identifier": schema.StringAttribute{
						Description:         "The global cluster ID of an Aurora cluster that becomes the primary cluster in the new global database cluster.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "The global cluster ID of an Aurora cluster that becomes the primary cluster in the new global database cluster.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"iops": schema.Int64Attribute{
						Description:         "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.  For information about valid IOPS values, see Provisioned IOPS storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the Amazon RDS User Guide.  This setting is required to create a Multi-AZ DB cluster.  Constraints: Must be a multiple between .5 and 50 of the storage amount for the DB cluster.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for each DB instance in the Multi-AZ DB cluster.  For information about valid IOPS values, see Provisioned IOPS storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the Amazon RDS User Guide.  This setting is required to create a Multi-AZ DB cluster.  Constraints: Must be a multiple between .5 and 50 of the storage amount for the DB cluster.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for an encrypted DB cluster.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  When a KMS key isn't specified in KmsKeyId:  * If ReplicationSourceIdentifier identifies an encrypted source, then Amazon RDS will use the KMS key used to encrypt the source. Otherwise, Amazon RDS will use your default KMS key.  * If the StorageEncrypted parameter is enabled and ReplicationSourceIdentifier isn't specified, then Amazon RDS will use your default KMS key.  There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  If you create a read replica of an encrypted DB cluster in another Amazon Web Services Region, you must set KmsKeyId to a KMS key identifier that is valid in the destination Amazon Web Services Region. This KMS key is used to encrypt the read replica in that Amazon Web Services Region.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for an encrypted DB cluster.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  When a KMS key isn't specified in KmsKeyId:  * If ReplicationSourceIdentifier identifies an encrypted source, then Amazon RDS will use the KMS key used to encrypt the source. Otherwise, Amazon RDS will use your default KMS key.  * If the StorageEncrypted parameter is enabled and ReplicationSourceIdentifier isn't specified, then Amazon RDS will use your default KMS key.  There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  If you create a read replica of an encrypted DB cluster in another Amazon Web Services Region, you must set KmsKeyId to a KMS key identifier that is valid in the destination Amazon Web Services Region. This KMS key is used to encrypt the read replica in that Amazon Web Services Region.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"manage_master_user_password": schema.BoolAttribute{
						Description:         "A value that indicates whether to manage the master user password with Amazon Web Services Secrets Manager.  For more information, see Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the Amazon RDS User Guide and Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the Amazon Aurora User Guide.  Constraints:  * Can't manage the master user password with Amazon Web Services Secrets Manager if MasterUserPassword is specified.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "A value that indicates whether to manage the master user password with Amazon Web Services Secrets Manager.  For more information, see Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the Amazon RDS User Guide and Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/rds-secrets-manager.html) in the Amazon Aurora User Guide.  Constraints:  * Can't manage the master user password with Amazon Web Services Secrets Manager if MasterUserPassword is specified.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"master_user_password": schema.SingleNestedAttribute{
						Description:         "The password for the master database user. This password can contain any printable ASCII character except '/', ''', or '@'.  Constraints:  * Must contain from 8 to 41 characters.  * Can't be specified if ManageMasterUserPassword is turned on.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The password for the master database user. This password can contain any printable ASCII character except '/', ''', or '@'.  Constraints:  * Must contain from 8 to 41 characters.  * Can't be specified if ManageMasterUserPassword is turned on.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
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

					"master_user_secret_kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.  This setting is valid only if the master user password is managed by RDS in Amazon Web Services Secrets Manager for the DB cluster.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a different Amazon Web Services account, then you can't use the aws/secretsmanager KMS key to encrypt the secret, and you must use a customer managed KMS key.  There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.  This setting is valid only if the master user password is managed by RDS in Amazon Web Services Secrets Manager for the DB cluster.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a different Amazon Web Services account, then you can't use the aws/secretsmanager KMS key to encrypt the secret, and you must use a customer managed KMS key.  There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"master_user_secret_kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"master_username": schema.StringAttribute{
						Description:         "The name of the master user for the DB cluster.  Constraints:  * Must be 1 to 16 letters or numbers.  * First character must be a letter.  * Can't be a reserved word for the chosen database engine.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The name of the master user for the DB cluster.  Constraints:  * Must be 1 to 16 letters or numbers.  * First character must be a letter.  * Can't be a reserved word for the chosen database engine.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"monitoring_interval": schema.Int64Attribute{
						Description:         "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster. To turn off collecting Enhanced Monitoring metrics, specify 0. The default is 0.  If MonitoringRoleArn is specified, also set MonitoringInterval to a value other than 0.  Valid Values: 0, 1, 5, 10, 15, 30, 60  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB cluster. To turn off collecting Enhanced Monitoring metrics, specify 0. The default is 0.  If MonitoringRoleArn is specified, also set MonitoringInterval to a value other than 0.  Valid Values: 0, 1, 5, 10, 15, 30, 60  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"monitoring_role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting up and enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide.  If MonitoringInterval is set to a value other than 0, supply a MonitoringRoleArn value.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "The Amazon Resource Name (ARN) for the IAM role that permits RDS to send Enhanced Monitoring metrics to Amazon CloudWatch Logs. An example is arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting up and enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide.  If MonitoringInterval is set to a value other than 0, supply a MonitoringRoleArn value.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"network_type": schema.StringAttribute{
						Description:         "The network type of the DB cluster.  Valid values:  * IPV4  * DUAL  The network type is determined by the DBSubnetGroup specified for the DB cluster. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL).  For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "The network type of the DB cluster.  Valid values:  * IPV4  * DUAL  The network type is determined by the DBSubnetGroup specified for the DB cluster. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL).  For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon Aurora User Guide.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"option_group_name": schema.StringAttribute{
						Description:         "A value that indicates that the DB cluster should be associated with the specified option group.  DB clusters are associated with a default option group that can't be modified.",
						MarkdownDescription: "A value that indicates that the DB cluster should be associated with the specified option group.  DB clusters are associated with a default option group that can't be modified.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"performance_insights_kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.  If you don't specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.  If you don't specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"performance_insights_retention_period": schema.Int64Attribute{
						Description:         "The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:  * 7  * month * 31, where month is a number of months from 1-23  * 731  For example, the following values are valid:  * 93 (3 months * 31)  * 341 (11 months * 31)  * 589 (19 months * 31)  * 731  If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:  * 7  * month * 31, where month is a number of months from 1-23  * 731  For example, the following values are valid:  * 93 (3 months * 31)  * 341 (11 months * 31)  * 589 (19 months * 31)  * 731  If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"port": schema.Int64Attribute{
						Description:         "The port number on which the instances in the DB cluster accept connections.  RDS for MySQL and Aurora MySQL  Default: 3306  Valid values: 1150-65535  RDS for PostgreSQL and Aurora PostgreSQL  Default: 5432  Valid values: 1150-65535  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The port number on which the instances in the DB cluster accept connections.  RDS for MySQL and Aurora MySQL  Default: 3306  Valid values: 1150-65535  RDS for PostgreSQL and Aurora PostgreSQL  Default: 5432  Valid values: 1150-65535  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pre_signed_url": schema.StringAttribute{
						Description:         "When you are replicating a DB cluster from one Amazon Web Services GovCloud (US) Region to another, an URL that contains a Signature Version 4 signed request for the CreateDBCluster operation to be called in the source Amazon Web Services Region where the DB cluster is replicated from. Specify PreSignedUrl only when you are performing cross-Region replication from an encrypted DB cluster.  The presigned URL must be a valid request for the CreateDBCluster API operation that can run in the source Amazon Web Services Region that contains the encrypted DB cluster to copy.  The presigned URL request must contain the following parameter values:  * KmsKeyId - The KMS key identifier for the KMS key to use to encrypt the copy of the DB cluster in the destination Amazon Web Services Region. This should refer to the same KMS key for both the CreateDBCluster operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL.  * DestinationRegion - The name of the Amazon Web Services Region that Aurora read replica will be created in.  * ReplicationSourceIdentifier - The DB cluster identifier for the encrypted DB cluster to be copied. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are copying an encrypted DB cluster from the us-west-2 Amazon Web Services Region, then your ReplicationSourceIdentifier would look like Example: arn:aws:rds:us-west-2:123456789012:cluster:aurora-cluster1.  To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).  If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "When you are replicating a DB cluster from one Amazon Web Services GovCloud (US) Region to another, an URL that contains a Signature Version 4 signed request for the CreateDBCluster operation to be called in the source Amazon Web Services Region where the DB cluster is replicated from. Specify PreSignedUrl only when you are performing cross-Region replication from an encrypted DB cluster.  The presigned URL must be a valid request for the CreateDBCluster API operation that can run in the source Amazon Web Services Region that contains the encrypted DB cluster to copy.  The presigned URL request must contain the following parameter values:  * KmsKeyId - The KMS key identifier for the KMS key to use to encrypt the copy of the DB cluster in the destination Amazon Web Services Region. This should refer to the same KMS key for both the CreateDBCluster operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL.  * DestinationRegion - The name of the Amazon Web Services Region that Aurora read replica will be created in.  * ReplicationSourceIdentifier - The DB cluster identifier for the encrypted DB cluster to be copied. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are copying an encrypted DB cluster from the us-west-2 Amazon Web Services Region, then your ReplicationSourceIdentifier would look like Example: arn:aws:rds:us-west-2:123456789012:cluster:aurora-cluster1.  To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).  If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region.  Valid for: Aurora DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"preferred_backup_window": schema.StringAttribute{
						Description:         "The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. To view the time blocks available, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the Amazon Aurora User Guide.  Constraints:  * Must be in the format hh24:mi-hh24:mi.  * Must be in Universal Coordinated Time (UTC).  * Must not conflict with the preferred maintenance window.  * Must be at least 30 minutes.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The daily time range during which automated backups are created if automated backups are enabled using the BackupRetentionPeriod parameter.  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. To view the time blocks available, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.Backups.BackupWindow) in the Amazon Aurora User Guide.  Constraints:  * Must be in the format hh24:mi-hh24:mi.  * Must be in Universal Coordinated Time (UTC).  * Must not conflict with the preferred maintenance window.  * Must be at least 30 minutes.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"preferred_maintenance_window": schema.StringAttribute{
						Description:         "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).  Format: ddd:hh24:mi-ddd:hh24:mi  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week. To see the time blocks available, see Adjusting the Preferred DB Cluster Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora) in the Amazon Aurora User Guide.  Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.  Constraints: Minimum 30-minute window.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The weekly time range during which system maintenance can occur, in Universal Coordinated Time (UTC).  Format: ddd:hh24:mi-ddd:hh24:mi  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week. To see the time blocks available, see Adjusting the Preferred DB Cluster Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_UpgradeDBInstance.Maintenance.html#AdjustingTheMaintenanceWindow.Aurora) in the Amazon Aurora User Guide.  Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.  Constraints: Minimum 30-minute window.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"publicly_accessible": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB cluster is publicly accessible.  When the DB cluster is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB cluster's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB cluster's VPC. Access to the DB cluster is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.  When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.  Default: The default behavior varies depending on whether DBSubnetGroupName is specified.  If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies:  * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB cluster is private.  * If the default VPC in the target Region has an internet gateway attached to it, the DB cluster is public.  If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies:  * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB cluster is private.  * If the subnets are part of a VPC that has an internet gateway attached to it, the DB cluster is public.  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "A value that indicates whether the DB cluster is publicly accessible.  When the DB cluster is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB cluster's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB cluster's VPC. Access to the DB cluster is ultimately controlled by the security group it uses. That public access isn't permitted if the security group assigned to the DB cluster doesn't permit it.  When the DB cluster isn't publicly accessible, it is an internal DB cluster with a DNS name that resolves to a private IP address.  Default: The default behavior varies depending on whether DBSubnetGroupName is specified.  If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies:  * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB cluster is private.  * If the default VPC in the target Region has an internet gateway attached to it, the DB cluster is public.  If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies:  * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB cluster is private.  * If the subnets are part of a VPC that has an internet gateway attached to it, the DB cluster is public.  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replication_source_identifier": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the source DB instance or DB cluster if this DB cluster is created as a read replica.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scaling_configuration": schema.SingleNestedAttribute{
						Description:         "For DB clusters in serverless DB engine mode, the scaling properties of the DB cluster.  Valid for: Aurora DB clusters only",
						MarkdownDescription: "For DB clusters in serverless DB engine mode, the scaling properties of the DB cluster.  Valid for: Aurora DB clusters only",
						Attributes: map[string]schema.Attribute{
							"auto_pause": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_capacity": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_capacity": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"seconds_before_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"seconds_until_auto_pause": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout_action": schema.StringAttribute{
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

					"serverless_v2_scaling_configuration": schema.SingleNestedAttribute{
						Description:         "Contains the scaling configuration of an Aurora Serverless v2 DB cluster.  For more information, see Using Amazon Aurora Serverless v2 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html) in the Amazon Aurora User Guide.",
						MarkdownDescription: "Contains the scaling configuration of an Aurora Serverless v2 DB cluster.  For more information, see Using Amazon Aurora Serverless v2 (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2.html) in the Amazon Aurora User Guide.",
						Attributes: map[string]schema.Attribute{
							"max_capacity": schema.Float64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_capacity": schema.Float64Attribute{
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

					"snapshot_identifier": schema.StringAttribute{
						Description:         "The identifier for the DB snapshot or DB cluster snapshot to restore from.  You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.  Constraints:  * Must match the identifier of an existing Snapshot.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "The identifier for the DB snapshot or DB cluster snapshot to restore from.  You can use either the name or the Amazon Resource Name (ARN) to specify a DB cluster snapshot. However, you can use only the ARN to specify a DB snapshot.  Constraints:  * Must match the identifier of an existing Snapshot.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"source_region": schema.StringAttribute{
						Description:         "SourceRegion is the source region where the resource exists. This is not sent over the wire and is only used for presigning. This value should always have the same region as the source ARN.",
						MarkdownDescription: "SourceRegion is the source region where the resource exists. This is not sent over the wire and is only used for presigning. This value should always have the same region as the source ARN.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"storage_encrypted": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB cluster is encrypted.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "A value that indicates whether the DB cluster is encrypted.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"storage_type": schema.StringAttribute{
						Description:         "Specifies the storage type to be associated with the DB cluster.  This setting is required to create a Multi-AZ DB cluster.  Valid values: io1  When specified, a value for the Iops parameter is required.  Default: io1  Valid for: Multi-AZ DB clusters only",
						MarkdownDescription: "Specifies the storage type to be associated with the DB cluster.  This setting is required to create a Multi-AZ DB cluster.  Valid values: io1  When specified, a value for the Iops parameter is required.  Default: io1  Valid for: Multi-AZ DB clusters only",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Tags to assign to the DB cluster.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "Tags to assign to the DB cluster.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
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

					"vpc_security_group_i_ds": schema.ListAttribute{
						Description:         "A list of EC2 VPC security groups to associate with this DB cluster.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						MarkdownDescription: "A list of EC2 VPC security groups to associate with this DB cluster.  Valid for: Aurora DB clusters and Multi-AZ DB clusters",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"vpc_security_group_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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
		},
	}
}

func (r *RdsServicesK8SAwsDbclusterV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *RdsServicesK8SAwsDbclusterV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_rds_services_k8s_aws_db_cluster_v1alpha1")

	var data RdsServicesK8SAwsDbclusterV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "rds.services.k8s.aws", Version: "v1alpha1", Resource: "dbclusters"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse RdsServicesK8SAwsDbclusterV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("DBCluster")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
