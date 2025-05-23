/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rds_services_k8s_aws_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &RdsServicesK8SAwsDbinstanceV1Alpha1Manifest{}
)

func NewRdsServicesK8SAwsDbinstanceV1Alpha1Manifest() datasource.DataSource {
	return &RdsServicesK8SAwsDbinstanceV1Alpha1Manifest{}
}

type RdsServicesK8SAwsDbinstanceV1Alpha1Manifest struct{}

type RdsServicesK8SAwsDbinstanceV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AllocatedStorage            *int64  `tfsdk:"allocated_storage" json:"allocatedStorage,omitempty"`
		AutoMinorVersionUpgrade     *bool   `tfsdk:"auto_minor_version_upgrade" json:"autoMinorVersionUpgrade,omitempty"`
		AvailabilityZone            *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
		BackupRetentionPeriod       *int64  `tfsdk:"backup_retention_period" json:"backupRetentionPeriod,omitempty"`
		BackupTarget                *string `tfsdk:"backup_target" json:"backupTarget,omitempty"`
		CaCertificateIdentifier     *string `tfsdk:"ca_certificate_identifier" json:"caCertificateIdentifier,omitempty"`
		CharacterSetName            *string `tfsdk:"character_set_name" json:"characterSetName,omitempty"`
		CopyTagsToSnapshot          *bool   `tfsdk:"copy_tags_to_snapshot" json:"copyTagsToSnapshot,omitempty"`
		CustomIAMInstanceProfile    *string `tfsdk:"custom_iam_instance_profile" json:"customIAMInstanceProfile,omitempty"`
		DbClusterIdentifier         *string `tfsdk:"db_cluster_identifier" json:"dbClusterIdentifier,omitempty"`
		DbClusterSnapshotIdentifier *string `tfsdk:"db_cluster_snapshot_identifier" json:"dbClusterSnapshotIdentifier,omitempty"`
		DbInstanceClass             *string `tfsdk:"db_instance_class" json:"dbInstanceClass,omitempty"`
		DbInstanceIdentifier        *string `tfsdk:"db_instance_identifier" json:"dbInstanceIdentifier,omitempty"`
		DbName                      *string `tfsdk:"db_name" json:"dbName,omitempty"`
		DbParameterGroupName        *string `tfsdk:"db_parameter_group_name" json:"dbParameterGroupName,omitempty"`
		DbParameterGroupRef         *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"db_parameter_group_ref" json:"dbParameterGroupRef,omitempty"`
		DbSnapshotIdentifier *string `tfsdk:"db_snapshot_identifier" json:"dbSnapshotIdentifier,omitempty"`
		DbSubnetGroupName    *string `tfsdk:"db_subnet_group_name" json:"dbSubnetGroupName,omitempty"`
		DbSubnetGroupRef     *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"db_subnet_group_ref" json:"dbSubnetGroupRef,omitempty"`
		DeletionProtection              *bool     `tfsdk:"deletion_protection" json:"deletionProtection,omitempty"`
		DestinationRegion               *string   `tfsdk:"destination_region" json:"destinationRegion,omitempty"`
		Domain                          *string   `tfsdk:"domain" json:"domain,omitempty"`
		DomainIAMRoleName               *string   `tfsdk:"domain_iam_role_name" json:"domainIAMRoleName,omitempty"`
		EnableCloudwatchLogsExports     *[]string `tfsdk:"enable_cloudwatch_logs_exports" json:"enableCloudwatchLogsExports,omitempty"`
		EnableCustomerOwnedIP           *bool     `tfsdk:"enable_customer_owned_ip" json:"enableCustomerOwnedIP,omitempty"`
		EnableIAMDatabaseAuthentication *bool     `tfsdk:"enable_iam_database_authentication" json:"enableIAMDatabaseAuthentication,omitempty"`
		Engine                          *string   `tfsdk:"engine" json:"engine,omitempty"`
		EngineVersion                   *string   `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		Iops                            *int64    `tfsdk:"iops" json:"iops,omitempty"`
		KmsKeyID                        *string   `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		KmsKeyRef                       *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"kms_key_ref" json:"kmsKeyRef,omitempty"`
		LicenseModel             *string `tfsdk:"license_model" json:"licenseModel,omitempty"`
		ManageMasterUserPassword *bool   `tfsdk:"manage_master_user_password" json:"manageMasterUserPassword,omitempty"`
		MasterUserPassword       *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"master_user_password" json:"masterUserPassword,omitempty"`
		MasterUserSecretKMSKeyID  *string `tfsdk:"master_user_secret_kms_key_id" json:"masterUserSecretKMSKeyID,omitempty"`
		MasterUserSecretKMSKeyRef *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"master_user_secret_kms_key_ref" json:"masterUserSecretKMSKeyRef,omitempty"`
		MasterUsername                     *string `tfsdk:"master_username" json:"masterUsername,omitempty"`
		MaxAllocatedStorage                *int64  `tfsdk:"max_allocated_storage" json:"maxAllocatedStorage,omitempty"`
		MonitoringInterval                 *int64  `tfsdk:"monitoring_interval" json:"monitoringInterval,omitempty"`
		MonitoringRoleARN                  *string `tfsdk:"monitoring_role_arn" json:"monitoringRoleARN,omitempty"`
		MultiAZ                            *bool   `tfsdk:"multi_az" json:"multiAZ,omitempty"`
		NcharCharacterSetName              *string `tfsdk:"nchar_character_set_name" json:"ncharCharacterSetName,omitempty"`
		NetworkType                        *string `tfsdk:"network_type" json:"networkType,omitempty"`
		OptionGroupName                    *string `tfsdk:"option_group_name" json:"optionGroupName,omitempty"`
		PerformanceInsightsEnabled         *bool   `tfsdk:"performance_insights_enabled" json:"performanceInsightsEnabled,omitempty"`
		PerformanceInsightsKMSKeyID        *string `tfsdk:"performance_insights_kms_key_id" json:"performanceInsightsKMSKeyID,omitempty"`
		PerformanceInsightsRetentionPeriod *int64  `tfsdk:"performance_insights_retention_period" json:"performanceInsightsRetentionPeriod,omitempty"`
		Port                               *int64  `tfsdk:"port" json:"port,omitempty"`
		PreSignedURL                       *string `tfsdk:"pre_signed_url" json:"preSignedURL,omitempty"`
		PreferredBackupWindow              *string `tfsdk:"preferred_backup_window" json:"preferredBackupWindow,omitempty"`
		PreferredMaintenanceWindow         *string `tfsdk:"preferred_maintenance_window" json:"preferredMaintenanceWindow,omitempty"`
		ProcessorFeatures                  *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"processor_features" json:"processorFeatures,omitempty"`
		PromotionTier              *int64  `tfsdk:"promotion_tier" json:"promotionTier,omitempty"`
		PubliclyAccessible         *bool   `tfsdk:"publicly_accessible" json:"publiclyAccessible,omitempty"`
		ReplicaMode                *string `tfsdk:"replica_mode" json:"replicaMode,omitempty"`
		SourceDBInstanceIdentifier *string `tfsdk:"source_db_instance_identifier" json:"sourceDBInstanceIdentifier,omitempty"`
		SourceRegion               *string `tfsdk:"source_region" json:"sourceRegion,omitempty"`
		StorageEncrypted           *bool   `tfsdk:"storage_encrypted" json:"storageEncrypted,omitempty"`
		StorageThroughput          *int64  `tfsdk:"storage_throughput" json:"storageThroughput,omitempty"`
		StorageType                *string `tfsdk:"storage_type" json:"storageType,omitempty"`
		Tags                       *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TdeCredentialARN            *string   `tfsdk:"tde_credential_arn" json:"tdeCredentialARN,omitempty"`
		TdeCredentialPassword       *string   `tfsdk:"tde_credential_password" json:"tdeCredentialPassword,omitempty"`
		Timezone                    *string   `tfsdk:"timezone" json:"timezone,omitempty"`
		UseDefaultProcessorFeatures *bool     `tfsdk:"use_default_processor_features" json:"useDefaultProcessorFeatures,omitempty"`
		VpcSecurityGroupIDs         *[]string `tfsdk:"vpc_security_group_i_ds" json:"vpcSecurityGroupIDs,omitempty"`
		VpcSecurityGroupRefs        *[]struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"vpc_security_group_refs" json:"vpcSecurityGroupRefs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_db_instance_v1alpha1_manifest"
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DBInstance is the Schema for the DBInstances API",
		MarkdownDescription: "DBInstance is the Schema for the DBInstances API",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
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
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "DBInstanceSpec defines the desired state of DBInstance. Contains the details of an Amazon RDS DB instance. This data type is used as a response element in the operations CreateDBInstance, CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance, PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3, RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",
				MarkdownDescription: "DBInstanceSpec defines the desired state of DBInstance. Contains the details of an Amazon RDS DB instance. This data type is used as a response element in the operations CreateDBInstance, CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance, PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3, RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",
				Attributes: map[string]schema.Attribute{
					"allocated_storage": schema.Int64Attribute{
						Description:         "The amount of storage in gibibytes (GiB) to allocate for the DB instance. This setting doesn't apply to Amazon Aurora DB instances. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume. Amazon RDS Custom Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server. * Provisioned IOPS storage (io1, io2): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server. RDS for Db2 Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. RDS for MariaDB Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 5 to 3072. RDS for MySQL Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 5 to 3072. RDS for Oracle Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 10 to 3072. RDS for PostgreSQL Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 5 to 3072. RDS for SQL Server Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Enterprise and Standard editions: Must be an integer from 20 to 16384. Web and Express editions: Must be an integer from 20 to 16384. * Provisioned IOPS storage (io1, io2): Enterprise and Standard editions: Must be an integer from 100 to 16384. Web and Express editions: Must be an integer from 100 to 16384. * Magnetic storage (standard): Enterprise and Standard editions: Must be an integer from 20 to 1024. Web and Express editions: Must be an integer from 20 to 1024.",
						MarkdownDescription: "The amount of storage in gibibytes (GiB) to allocate for the DB instance. This setting doesn't apply to Amazon Aurora DB instances. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume. Amazon RDS Custom Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server. * Provisioned IOPS storage (io1, io2): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server. RDS for Db2 Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. RDS for MariaDB Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 5 to 3072. RDS for MySQL Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 5 to 3072. RDS for Oracle Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 10 to 3072. RDS for PostgreSQL Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536. * Provisioned IOPS storage (io1, io2): Must be an integer from 100 to 65536. * Magnetic storage (standard): Must be an integer from 5 to 3072. RDS for SQL Server Constraints to the amount of storage for each storage type are the following: * General Purpose (SSD) storage (gp2, gp3): Enterprise and Standard editions: Must be an integer from 20 to 16384. Web and Express editions: Must be an integer from 20 to 16384. * Provisioned IOPS storage (io1, io2): Enterprise and Standard editions: Must be an integer from 100 to 16384. Web and Express editions: Must be an integer from 100 to 16384. * Magnetic storage (standard): Enterprise and Standard editions: Must be an integer from 20 to 1024. Web and Express editions: Must be an integer from 20 to 1024.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auto_minor_version_upgrade": schema.BoolAttribute{
						Description:         "Specifies whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically. If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgrade to false.",
						MarkdownDescription: "Specifies whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically. If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgrade to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"availability_zone": schema.StringAttribute{
						Description:         "The Availability Zone (AZ) where the database will be created. For information on Amazon Web Services Regions and Availability Zones, see Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html). For Amazon Aurora, each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one. Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web Services Region. Constraints: * The AvailabilityZone parameter can't be specified if the DB instance is a Multi-AZ deployment. * The specified Availability Zone must be in the same Amazon Web Services Region as the current endpoint. Example: us-east-1d",
						MarkdownDescription: "The Availability Zone (AZ) where the database will be created. For information on Amazon Web Services Regions and Availability Zones, see Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html). For Amazon Aurora, each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one. Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web Services Region. Constraints: * The AvailabilityZone parameter can't be specified if the DB instance is a Multi-AZ deployment. * The specified Availability Zone must be in the same Amazon Web Services Region as the current endpoint. Example: us-east-1d",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_retention_period": schema.Int64Attribute{
						Description:         "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. This setting doesn't apply to Amazon Aurora DB instances. The retention period for automated backups is managed by the DB cluster. Default: 1 Constraints: * Must be a value from 0 to 35. * Can't be set to 0 if the DB instance is a source to read replicas. * Can't be set to 0 for an RDS Custom for Oracle DB instance.",
						MarkdownDescription: "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups. This setting doesn't apply to Amazon Aurora DB instances. The retention period for automated backups is managed by the DB cluster. Default: 1 Constraints: * Must be a value from 0 to 35. * Can't be set to 0 if the DB instance is a source to read replicas. * Can't be set to 0 for an RDS Custom for Oracle DB instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_target": schema.StringAttribute{
						Description:         "The location for storing automated backups and manual snapshots. Valid Values: * outposts (Amazon Web Services Outposts) * region (Amazon Web Services Region) Default: region For more information, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.",
						MarkdownDescription: "The location for storing automated backups and manual snapshots. Valid Values: * outposts (Amazon Web Services Outposts) * region (Amazon Web Services Region) Default: region For more information, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ca_certificate_identifier": schema.StringAttribute{
						Description:         "The CA certificate identifier to use for the DB instance's server certificate. This setting doesn't apply to RDS Custom DB instances. For more information, see Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the Amazon Aurora User Guide.",
						MarkdownDescription: "The CA certificate identifier to use for the DB instance's server certificate. This setting doesn't apply to RDS Custom DB instances. For more information, see Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the Amazon Aurora User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"character_set_name": schema.StringAttribute{
						Description:         "For supported engines, the character set (CharacterSet) to associate the DB instance with. This setting doesn't apply to the following DB instances: * Amazon Aurora - The character set is managed by the DB cluster. For more information, see CreateDBCluster. * RDS Custom - However, if you need to change the character set, you can change it on the database itself.",
						MarkdownDescription: "For supported engines, the character set (CharacterSet) to associate the DB instance with. This setting doesn't apply to the following DB instances: * Amazon Aurora - The character set is managed by the DB cluster. For more information, see CreateDBCluster. * RDS Custom - However, if you need to change the character set, you can change it on the database itself.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"copy_tags_to_snapshot": schema.BoolAttribute{
						Description:         "Specifies whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied. This setting doesn't apply to Amazon Aurora DB instances. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.",
						MarkdownDescription: "Specifies whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied. This setting doesn't apply to Amazon Aurora DB instances. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_iam_instance_profile": schema.StringAttribute{
						Description:         "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance. This setting is required for RDS Custom. Constraints: * The profile must exist in your account. * The profile must have an IAM role that Amazon EC2 has permissions to assume. * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom. For the list of permissions required for the IAM role, see Configure IAM and your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the Amazon RDS User Guide.",
						MarkdownDescription: "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance. This setting is required for RDS Custom. Constraints: * The profile must exist in your account. * The profile must have an IAM role that Amazon EC2 has permissions to assume. * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom. For the list of permissions required for the IAM role, see Configure IAM and your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_identifier": schema.StringAttribute{
						Description:         "The identifier of the DB cluster that this DB instance will belong to. This setting doesn't apply to RDS Custom DB instances.",
						MarkdownDescription: "The identifier of the DB cluster that this DB instance will belong to. This setting doesn't apply to RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_snapshot_identifier": schema.StringAttribute{
						Description:         "The identifier for the Multi-AZ DB cluster snapshot to restore from. For more information on Multi-AZ DB clusters, see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the Amazon RDS User Guide. Constraints: * Must match the identifier of an existing Multi-AZ DB cluster snapshot. * Can't be specified when DBSnapshotIdentifier is specified. * Must be specified when DBSnapshotIdentifier isn't specified. * If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot. * Can't be the identifier of an Aurora DB cluster snapshot.",
						MarkdownDescription: "The identifier for the Multi-AZ DB cluster snapshot to restore from. For more information on Multi-AZ DB clusters, see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the Amazon RDS User Guide. Constraints: * Must match the identifier of an existing Multi-AZ DB cluster snapshot. * Can't be specified when DBSnapshotIdentifier is specified. * Must be specified when DBSnapshotIdentifier isn't specified. * If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot. * Can't be the identifier of an Aurora DB cluster snapshot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_instance_class": schema.StringAttribute{
						Description:         "The compute and memory capacity of the DB instance, for example db.m5.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines. For the full list of DB instance classes, and availability for your engine, see DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide or Aurora DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html) in the Amazon Aurora User Guide.",
						MarkdownDescription: "The compute and memory capacity of the DB instance, for example db.m5.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines. For the full list of DB instance classes, and availability for your engine, see DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide or Aurora DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html) in the Amazon Aurora User Guide.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_instance_identifier": schema.StringAttribute{
						Description:         "The identifier for this DB instance. This parameter is stored as a lowercase string. Constraints: * Must contain from 1 to 63 letters, numbers, or hyphens. * First character must be a letter. * Can't end with a hyphen or contain two consecutive hyphens. Example: mydbinstance",
						MarkdownDescription: "The identifier for this DB instance. This parameter is stored as a lowercase string. Constraints: * Must contain from 1 to 63 letters, numbers, or hyphens. * First character must be a letter. * Can't end with a hyphen or contain two consecutive hyphens. Example: mydbinstance",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_name": schema.StringAttribute{
						Description:         "The meaning of this parameter differs according to the database engine you use. Amazon Aurora MySQL The name of the database to create when the primary DB instance of the Aurora MySQL DB cluster is created. If this parameter isn't specified for an Aurora MySQL DB cluster, no database is created in the DB cluster. Constraints: * Must contain 1 to 64 alphanumeric characters. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the database engine. Amazon Aurora PostgreSQL The name of the database to create when the primary DB instance of the Aurora PostgreSQL DB cluster is created. A database named postgres is always created. If this parameter is specified, an additional database with this name is created. Constraints: * It must contain 1 to 63 alphanumeric characters. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0 to 9). * Can't be a word reserved by the database engine. Amazon RDS Custom for Oracle The Oracle System ID (SID) of the created RDS Custom DB instance. If you don't specify a value, the default value is ORCL for non-CDBs and RDSCDB for CDBs. Default: ORCL Constraints: * Must contain 1 to 8 alphanumeric characters. * Must contain a letter. * Can't be a word reserved by the database engine. Amazon RDS Custom for SQL Server Not applicable. Must be null. RDS for Db2 The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance. In some cases, we recommend that you don't add a database name. For more information, see Additional considerations (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-db-instance-prereqs.html#db2-prereqs-additional-considerations) in the Amazon RDS User Guide. Constraints: * Must contain 1 to 64 letters or numbers. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for MariaDB The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance. Constraints: * Must contain 1 to 64 letters or numbers. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for MySQL The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance. Constraints: * Must contain 1 to 64 letters or numbers. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for Oracle The Oracle System ID (SID) of the created DB instance. If you don't specify a value, the default value is ORCL. You can't specify the string null, or any other reserved word, for DBName. Default: ORCL Constraints: * Can't be longer than 8 characters. RDS for PostgreSQL The name of the database to create when the DB instance is created. A database named postgres is always created. If this parameter is specified, an additional database with this name is created. Constraints: * Must contain 1 to 63 letters, numbers, or underscores. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for SQL Server Not applicable. Must be null.",
						MarkdownDescription: "The meaning of this parameter differs according to the database engine you use. Amazon Aurora MySQL The name of the database to create when the primary DB instance of the Aurora MySQL DB cluster is created. If this parameter isn't specified for an Aurora MySQL DB cluster, no database is created in the DB cluster. Constraints: * Must contain 1 to 64 alphanumeric characters. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the database engine. Amazon Aurora PostgreSQL The name of the database to create when the primary DB instance of the Aurora PostgreSQL DB cluster is created. A database named postgres is always created. If this parameter is specified, an additional database with this name is created. Constraints: * It must contain 1 to 63 alphanumeric characters. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0 to 9). * Can't be a word reserved by the database engine. Amazon RDS Custom for Oracle The Oracle System ID (SID) of the created RDS Custom DB instance. If you don't specify a value, the default value is ORCL for non-CDBs and RDSCDB for CDBs. Default: ORCL Constraints: * Must contain 1 to 8 alphanumeric characters. * Must contain a letter. * Can't be a word reserved by the database engine. Amazon RDS Custom for SQL Server Not applicable. Must be null. RDS for Db2 The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance. In some cases, we recommend that you don't add a database name. For more information, see Additional considerations (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-db-instance-prereqs.html#db2-prereqs-additional-considerations) in the Amazon RDS User Guide. Constraints: * Must contain 1 to 64 letters or numbers. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for MariaDB The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance. Constraints: * Must contain 1 to 64 letters or numbers. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for MySQL The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance. Constraints: * Must contain 1 to 64 letters or numbers. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for Oracle The Oracle System ID (SID) of the created DB instance. If you don't specify a value, the default value is ORCL. You can't specify the string null, or any other reserved word, for DBName. Default: ORCL Constraints: * Can't be longer than 8 characters. RDS for PostgreSQL The name of the database to create when the DB instance is created. A database named postgres is always created. If this parameter is specified, an additional database with this name is created. Constraints: * Must contain 1 to 63 letters, numbers, or underscores. * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9). * Can't be a word reserved by the specified database engine. RDS for SQL Server Not applicable. Must be null.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_parameter_group_name": schema.StringAttribute{
						Description:         "The name of the DB parameter group to associate with this DB instance. If you don't specify a value, then Amazon RDS uses the default DB parameter group for the specified DB engine and version. This setting doesn't apply to RDS Custom DB instances. Constraints: * Must be 1 to 255 letters, numbers, or hyphens. * The first character must be a letter. * Can't end with a hyphen or contain two consecutive hyphens.",
						MarkdownDescription: "The name of the DB parameter group to associate with this DB instance. If you don't specify a value, then Amazon RDS uses the default DB parameter group for the specified DB engine and version. This setting doesn't apply to RDS Custom DB instances. Constraints: * Must be 1 to 255 letters, numbers, or hyphens. * The first character must be a letter. * Can't end with a hyphen or contain two consecutive hyphens.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_parameter_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"db_snapshot_identifier": schema.StringAttribute{
						Description:         "The identifier for the DB snapshot to restore from. Constraints: * Must match the identifier of an existing DB snapshot. * Can't be specified when DBClusterSnapshotIdentifier is specified. * Must be specified when DBClusterSnapshotIdentifier isn't specified. * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB snapshot.",
						MarkdownDescription: "The identifier for the DB snapshot to restore from. Constraints: * Must match the identifier of an existing DB snapshot. * Can't be specified when DBClusterSnapshotIdentifier is specified. * Must be specified when DBClusterSnapshotIdentifier isn't specified. * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB snapshot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_subnet_group_name": schema.StringAttribute{
						Description:         "A DB subnet group to associate with this DB instance. Constraints: * Must match the name of an existing DB subnet group. Example: mydbsubnetgroup",
						MarkdownDescription: "A DB subnet group to associate with this DB instance. Constraints: * Must match the name of an existing DB subnet group. Example: mydbsubnetgroup",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_subnet_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"deletion_protection": schema.BoolAttribute{
						Description:         "Specifies whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html). This setting doesn't apply to Amazon Aurora DB instances. You can enable or disable deletion protection for the DB cluster. For more information, see CreateDBCluster. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.",
						MarkdownDescription: "Specifies whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html). This setting doesn't apply to Amazon Aurora DB instances. You can enable or disable deletion protection for the DB cluster. For more information, see CreateDBCluster. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"destination_region": schema.StringAttribute{
						Description:         "DestinationRegion is used for presigning the request to a given region.",
						MarkdownDescription: "DestinationRegion is used for presigning the request to a given region.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"domain": schema.StringAttribute{
						Description:         "The Active Directory directory ID to create the DB instance in. Currently, you can create only Db2, MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances in an Active Directory Domain. For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (The domain is managed by the DB cluster.) * RDS Custom",
						MarkdownDescription: "The Active Directory directory ID to create the DB instance in. Currently, you can create only Db2, MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances in an Active Directory Domain. For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (The domain is managed by the DB cluster.) * RDS Custom",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"domain_iam_role_name": schema.StringAttribute{
						Description:         "The name of the IAM role to use when making API calls to the Directory Service. This setting doesn't apply to the following DB instances: * Amazon Aurora (The domain is managed by the DB cluster.) * RDS Custom",
						MarkdownDescription: "The name of the IAM role to use when making API calls to the Directory Service. This setting doesn't apply to the following DB instances: * Amazon Aurora (The domain is managed by the DB cluster.) * RDS Custom",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_cloudwatch_logs_exports": schema.ListAttribute{
						Description:         "The list of log types to enable for exporting to CloudWatch Logs. For more information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (CloudWatch Logs exports are managed by the DB cluster.) * RDS Custom The following values are valid for each DB engine: * RDS for Db2 - diag.log | notify.log * RDS for MariaDB - audit | error | general | slowquery * RDS for Microsoft SQL Server - agent | error * RDS for MySQL - audit | error | general | slowquery * RDS for Oracle - alert | audit | listener | trace | oemagent * RDS for PostgreSQL - postgresql | upgrade",
						MarkdownDescription: "The list of log types to enable for exporting to CloudWatch Logs. For more information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (CloudWatch Logs exports are managed by the DB cluster.) * RDS Custom The following values are valid for each DB engine: * RDS for Db2 - diag.log | notify.log * RDS for MariaDB - audit | error | general | slowquery * RDS for Microsoft SQL Server - agent | error * RDS for MySQL - audit | error | general | slowquery * RDS for Oracle - alert | audit | listener | trace | oemagent * RDS for PostgreSQL - postgresql | upgrade",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_customer_owned_ip": schema.BoolAttribute{
						Description:         "Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance. A CoIP provides local or external connectivity to resources in your Outpost subnets through your on-premises network. For some use cases, a CoIP can provide lower latency for connections to the DB instance from outside of its virtual private cloud (VPC) on your local network. For more information about RDS on Outposts, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide. For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing) in the Amazon Web Services Outposts User Guide.",
						MarkdownDescription: "Specifies whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance. A CoIP provides local or external connectivity to resources in your Outpost subnets through your on-premises network. For some use cases, a CoIP can provide lower latency for connections to the DB instance from outside of its virtual private cloud (VPC) on your local network. For more information about RDS on Outposts, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide. For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing) in the Amazon Web Services Outposts User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_iam_database_authentication": schema.BoolAttribute{
						Description:         "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled. For more information, see IAM Database Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (Mapping Amazon Web Services IAM accounts to database accounts is managed by the DB cluster.) * RDS Custom",
						MarkdownDescription: "Specifies whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled. For more information, see IAM Database Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (Mapping Amazon Web Services IAM accounts to database accounts is managed by the DB cluster.) * RDS Custom",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine": schema.StringAttribute{
						Description:         "The database engine to use for this DB instance. Not every database engine is available in every Amazon Web Services Region. Valid Values: * aurora-mysql (for Aurora MySQL DB instances) * aurora-postgresql (for Aurora PostgreSQL DB instances) * custom-oracle-ee (for RDS Custom for Oracle DB instances) * custom-oracle-ee-cdb (for RDS Custom for Oracle DB instances) * custom-oracle-se2 (for RDS Custom for Oracle DB instances) * custom-oracle-se2-cdb (for RDS Custom for Oracle DB instances) * custom-sqlserver-ee (for RDS Custom for SQL Server DB instances) * custom-sqlserver-se (for RDS Custom for SQL Server DB instances) * custom-sqlserver-web (for RDS Custom for SQL Server DB instances) * custom-sqlserver-dev (for RDS Custom for SQL Server DB instances) * db2-ae * db2-se * mariadb * mysql * oracle-ee * oracle-ee-cdb * oracle-se2 * oracle-se2-cdb * postgres * sqlserver-ee * sqlserver-se * sqlserver-ex * sqlserver-web",
						MarkdownDescription: "The database engine to use for this DB instance. Not every database engine is available in every Amazon Web Services Region. Valid Values: * aurora-mysql (for Aurora MySQL DB instances) * aurora-postgresql (for Aurora PostgreSQL DB instances) * custom-oracle-ee (for RDS Custom for Oracle DB instances) * custom-oracle-ee-cdb (for RDS Custom for Oracle DB instances) * custom-oracle-se2 (for RDS Custom for Oracle DB instances) * custom-oracle-se2-cdb (for RDS Custom for Oracle DB instances) * custom-sqlserver-ee (for RDS Custom for SQL Server DB instances) * custom-sqlserver-se (for RDS Custom for SQL Server DB instances) * custom-sqlserver-web (for RDS Custom for SQL Server DB instances) * custom-sqlserver-dev (for RDS Custom for SQL Server DB instances) * db2-ae * db2-se * mariadb * mysql * oracle-ee * oracle-ee-cdb * oracle-se2 * oracle-se2-cdb * postgres * sqlserver-ee * sqlserver-se * sqlserver-ex * sqlserver-web",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The version number of the database engine to use. This setting doesn't apply to Amazon Aurora DB instances. The version number of the database engine the DB instance uses is managed by the DB cluster. For a list of valid engine versions, use the DescribeDBEngineVersions operation. The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every Amazon Web Services Region. Amazon RDS Custom for Oracle A custom engine version (CEV) that you have previously created. This setting is required for RDS Custom for Oracle. The CEV name has the following format: 19.customized_string. A valid CEV name is 19.my_cev1. For more information, see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create) in the Amazon RDS User Guide. Amazon RDS Custom for SQL Server See RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html) in the Amazon RDS User Guide. RDS for Db2 For information, see Db2 on Amazon RDS versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Db2.html#Db2.Concepts.VersionMgmt) in the Amazon RDS User Guide. RDS for MariaDB For information, see MariaDB on Amazon RDS versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the Amazon RDS User Guide. RDS for Microsoft SQL Server For information, see Microsoft SQL Server versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the Amazon RDS User Guide. RDS for MySQL For information, see MySQL on Amazon RDS versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide. RDS for Oracle For information, see Oracle Database Engine release notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the Amazon RDS User Guide. RDS for PostgreSQL For information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.",
						MarkdownDescription: "The version number of the database engine to use. This setting doesn't apply to Amazon Aurora DB instances. The version number of the database engine the DB instance uses is managed by the DB cluster. For a list of valid engine versions, use the DescribeDBEngineVersions operation. The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every Amazon Web Services Region. Amazon RDS Custom for Oracle A custom engine version (CEV) that you have previously created. This setting is required for RDS Custom for Oracle. The CEV name has the following format: 19.customized_string. A valid CEV name is 19.my_cev1. For more information, see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create) in the Amazon RDS User Guide. Amazon RDS Custom for SQL Server See RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html) in the Amazon RDS User Guide. RDS for Db2 For information, see Db2 on Amazon RDS versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Db2.html#Db2.Concepts.VersionMgmt) in the Amazon RDS User Guide. RDS for MariaDB For information, see MariaDB on Amazon RDS versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the Amazon RDS User Guide. RDS for Microsoft SQL Server For information, see Microsoft SQL Server versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the Amazon RDS User Guide. RDS for MySQL For information, see MySQL on Amazon RDS versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide. RDS for Oracle For information, see Oracle Database Engine release notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the Amazon RDS User Guide. RDS for PostgreSQL For information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"iops": schema.Int64Attribute{
						Description:         "The amount of Provisioned IOPS (input/output operations per second) to initially allocate for the DB instance. For information about valid IOPS values, see Amazon RDS DB instance storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html) in the Amazon RDS User Guide. This setting doesn't apply to Amazon Aurora DB instances. Storage is managed by the DB cluster. Constraints: * For RDS for Db2, MariaDB, MySQL, Oracle, and PostgreSQL - Must be a multiple between .5 and 50 of the storage amount for the DB instance. * For RDS for SQL Server - Must be a multiple between 1 and 50 of the storage amount for the DB instance.",
						MarkdownDescription: "The amount of Provisioned IOPS (input/output operations per second) to initially allocate for the DB instance. For information about valid IOPS values, see Amazon RDS DB instance storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html) in the Amazon RDS User Guide. This setting doesn't apply to Amazon Aurora DB instances. Storage is managed by the DB cluster. Constraints: * For RDS for Db2, MariaDB, MySQL, Oracle, and PostgreSQL - Must be a multiple between .5 and 50 of the storage amount for the DB instance. * For RDS for SQL Server - Must be a multiple between 1 and 50 of the storage amount for the DB instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for an encrypted DB instance. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN. This setting doesn't apply to Amazon Aurora DB instances. The Amazon Web Services KMS key identifier is managed by the DB cluster. For more information, see CreateDBCluster. If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region. For Amazon RDS Custom, a KMS key is required for DB instances. For most RDS engines, if you leave this parameter empty while enabling StorageEncrypted, the engine uses the default KMS key. However, RDS Custom doesn't use the default key when this parameter is empty. You must explicitly specify a key.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for an encrypted DB instance. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN. This setting doesn't apply to Amazon Aurora DB instances. The Amazon Web Services KMS key identifier is managed by the DB cluster. For more information, see CreateDBCluster. If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region. For Amazon RDS Custom, a KMS key is required for DB instances. For most RDS engines, if you leave this parameter empty while enabling StorageEncrypted, the engine uses the default KMS key. However, RDS Custom doesn't use the default key when this parameter is empty. You must explicitly specify a key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"license_model": schema.StringAttribute{
						Description:         "The license model information for this DB instance. License models for RDS for Db2 require additional configuration. The Bring Your Own License (BYOL) model requires a custom parameter group and an Amazon Web Services License Manager self-managed license. The Db2 license through Amazon Web Services Marketplace model requires an Amazon Web Services Marketplace subscription. For more information, see Amazon RDS for Db2 licensing options (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-licensing.html) in the Amazon RDS User Guide. The default for RDS for Db2 is bring-your-own-license. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances. Valid Values: * RDS for Db2 - bring-your-own-license | marketplace-license * RDS for MariaDB - general-public-license * RDS for Microsoft SQL Server - license-included * RDS for MySQL - general-public-license * RDS for Oracle - bring-your-own-license | license-included * RDS for PostgreSQL - postgresql-license",
						MarkdownDescription: "The license model information for this DB instance. License models for RDS for Db2 require additional configuration. The Bring Your Own License (BYOL) model requires a custom parameter group and an Amazon Web Services License Manager self-managed license. The Db2 license through Amazon Web Services Marketplace model requires an Amazon Web Services Marketplace subscription. For more information, see Amazon RDS for Db2 licensing options (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-licensing.html) in the Amazon RDS User Guide. The default for RDS for Db2 is bring-your-own-license. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances. Valid Values: * RDS for Db2 - bring-your-own-license | marketplace-license * RDS for MariaDB - general-public-license * RDS for Microsoft SQL Server - license-included * RDS for MySQL - general-public-license * RDS for Oracle - bring-your-own-license | license-included * RDS for PostgreSQL - postgresql-license",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"manage_master_user_password": schema.BoolAttribute{
						Description:         "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager. For more information, see Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the Amazon RDS User Guide. Constraints: * Can't manage the master user password with Amazon Web Services Secrets Manager if MasterUserPassword is specified.",
						MarkdownDescription: "Specifies whether to manage the master user password with Amazon Web Services Secrets Manager. For more information, see Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the Amazon RDS User Guide. Constraints: * Can't manage the master user password with Amazon Web Services Secrets Manager if MasterUserPassword is specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"master_user_password": schema.SingleNestedAttribute{
						Description:         "The password for the master user. This setting doesn't apply to Amazon Aurora DB instances. The password for the master user is managed by the DB cluster. Constraints: * Can't be specified if ManageMasterUserPassword is turned on. * Can include any printable ASCII character except '/', ''', or '@'. For RDS for Oracle, can't include the '&' (ampersand) or the ''' (single quotes) character. Length Constraints: * RDS for Db2 - Must contain from 8 to 255 characters. * RDS for MariaDB - Must contain from 8 to 41 characters. * RDS for Microsoft SQL Server - Must contain from 8 to 128 characters. * RDS for MySQL - Must contain from 8 to 41 characters. * RDS for Oracle - Must contain from 8 to 30 characters. * RDS for PostgreSQL - Must contain from 8 to 128 characters.",
						MarkdownDescription: "The password for the master user. This setting doesn't apply to Amazon Aurora DB instances. The password for the master user is managed by the DB cluster. Constraints: * Can't be specified if ManageMasterUserPassword is turned on. * Can include any printable ASCII character except '/', ''', or '@'. For RDS for Oracle, can't include the '&' (ampersand) or the ''' (single quotes) character. Length Constraints: * RDS for Db2 - Must contain from 8 to 255 characters. * RDS for MariaDB - Must contain from 8 to 41 characters. * RDS for Microsoft SQL Server - Must contain from 8 to 128 characters. * RDS for MySQL - Must contain from 8 to 41 characters. * RDS for Oracle - Must contain from 8 to 30 characters. * RDS for PostgreSQL - Must contain from 8 to 128 characters.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "Key is the key within the secret",
								MarkdownDescription: "Key is the key within the secret",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"master_user_secret_kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager. This setting is valid only if the master user password is managed by RDS in Amazon Web Services Secrets Manager for the DB instance. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN. If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a different Amazon Web Services account, then you can't use the aws/secretsmanager KMS key to encrypt the secret, and you must use a customer managed KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager. This setting is valid only if the master user password is managed by RDS in Amazon Web Services Secrets Manager for the DB instance. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN. If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a different Amazon Web Services account, then you can't use the aws/secretsmanager KMS key to encrypt the secret, and you must use a customer managed KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"master_user_secret_kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"master_username": schema.StringAttribute{
						Description:         "The name for the master user. This setting doesn't apply to Amazon Aurora DB instances. The name for the master user is managed by the DB cluster. This setting is required for RDS DB instances. Constraints: * Must be 1 to 16 letters, numbers, or underscores. * First character must be a letter. * Can't be a reserved word for the chosen database engine.",
						MarkdownDescription: "The name for the master user. This setting doesn't apply to Amazon Aurora DB instances. The name for the master user is managed by the DB cluster. This setting is required for RDS DB instances. Constraints: * Must be 1 to 16 letters, numbers, or underscores. * First character must be a letter. * Can't be a reserved word for the chosen database engine.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_allocated_storage": schema.Int64Attribute{
						Description:         "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance. For more information about this setting, including limitations that apply to it, see Managing capacity automatically with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (Storage is managed by the DB cluster.) * RDS Custom",
						MarkdownDescription: "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance. For more information about this setting, including limitations that apply to it, see Managing capacity automatically with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the Amazon RDS User Guide. This setting doesn't apply to the following DB instances: * Amazon Aurora (Storage is managed by the DB cluster.) * RDS Custom",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitoring_interval": schema.Int64Attribute{
						Description:         "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify 0. If MonitoringRoleArn is specified, then you must set MonitoringInterval to a value other than 0. This setting doesn't apply to RDS Custom DB instances. Valid Values: 0 | 1 | 5 | 10 | 15 | 30 | 60 Default: 0",
						MarkdownDescription: "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify 0. If MonitoringRoleArn is specified, then you must set MonitoringInterval to a value other than 0. This setting doesn't apply to RDS Custom DB instances. Valid Values: 0 | 1 | 5 | 10 | 15 | 30 | 60 Default: 0",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitoring_role_arn": schema.StringAttribute{
						Description:         "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting Up and Enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide. If MonitoringInterval is set to a value other than 0, then you must supply a MonitoringRoleArn value. This setting doesn't apply to RDS Custom DB instances.",
						MarkdownDescription: "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting Up and Enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide. If MonitoringInterval is set to a value other than 0, then you must supply a MonitoringRoleArn value. This setting doesn't apply to RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"multi_az": schema.BoolAttribute{
						Description:         "Specifies whether the DB instance is a Multi-AZ deployment. You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ deployment. This setting doesn't apply to the following DB instances: * Amazon Aurora (DB instance Availability Zones (AZs) are managed by the DB cluster.) * RDS Custom",
						MarkdownDescription: "Specifies whether the DB instance is a Multi-AZ deployment. You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ deployment. This setting doesn't apply to the following DB instances: * Amazon Aurora (DB instance Availability Zones (AZs) are managed by the DB cluster.) * RDS Custom",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nchar_character_set_name": schema.StringAttribute{
						Description:         "The name of the NCHAR character set for the Oracle DB instance. This setting doesn't apply to RDS Custom DB instances.",
						MarkdownDescription: "The name of the NCHAR character set for the Oracle DB instance. This setting doesn't apply to RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network_type": schema.StringAttribute{
						Description:         "The network type of the DB instance. The network type is determined by the DBSubnetGroup specified for the DB instance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL). For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon RDS User Guide. Valid Values: IPV4 | DUAL",
						MarkdownDescription: "The network type of the DB instance. The network type is determined by the DBSubnetGroup specified for the DB instance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL). For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon RDS User Guide. Valid Values: IPV4 | DUAL",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"option_group_name": schema.StringAttribute{
						Description:         "The option group to associate the DB instance with. Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance after it is associated with a DB instance. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						MarkdownDescription: "The option group to associate the DB instance with. Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance after it is associated with a DB instance. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_enabled": schema.BoolAttribute{
						Description:         "Specifies whether to enable Performance Insights for the DB instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom DB instances.",
						MarkdownDescription: "Specifies whether to enable Performance Insights for the DB instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide. This setting doesn't apply to RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for encryption of Performance Insights data. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you don't specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region. This setting doesn't apply to RDS Custom DB instances.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for encryption of Performance Insights data. The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. If you don't specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region. This setting doesn't apply to RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_retention_period": schema.Int64Attribute{
						Description:         "The number of days to retain Performance Insights data. This setting doesn't apply to RDS Custom DB instances. Valid Values: * 7 * month * 31, where month is a number of months from 1-23. Examples: 93 (3 months * 31), 341 (11 months * 31), 589 (19 months * 31) * 731 Default: 7 days If you specify a retention period that isn't valid, such as 94, Amazon RDS returns an error.",
						MarkdownDescription: "The number of days to retain Performance Insights data. This setting doesn't apply to RDS Custom DB instances. Valid Values: * 7 * month * 31, where month is a number of months from 1-23. Examples: 93 (3 months * 31), 341 (11 months * 31), 589 (19 months * 31) * 731 Default: 7 days If you specify a retention period that isn't valid, such as 94, Amazon RDS returns an error.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"port": schema.Int64Attribute{
						Description:         "The port number on which the database accepts connections. This setting doesn't apply to Aurora DB instances. The port number is managed by the cluster. Valid Values: 1150-65535 Default: * RDS for Db2 - 50000 * RDS for MariaDB - 3306 * RDS for Microsoft SQL Server - 1433 * RDS for MySQL - 3306 * RDS for Oracle - 1521 * RDS for PostgreSQL - 5432 Constraints: * For RDS for Microsoft SQL Server, the value can't be 1234, 1434, 3260, 3343, 3389, 47001, or 49152-49156.",
						MarkdownDescription: "The port number on which the database accepts connections. This setting doesn't apply to Aurora DB instances. The port number is managed by the cluster. Valid Values: 1150-65535 Default: * RDS for Db2 - 50000 * RDS for MariaDB - 3306 * RDS for Microsoft SQL Server - 1433 * RDS for MySQL - 3306 * RDS for Oracle - 1521 * RDS for PostgreSQL - 5432 Constraints: * For RDS for Microsoft SQL Server, the value can't be 1234, 1434, 3260, 3343, 3389, 47001, or 49152-49156.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pre_signed_url": schema.StringAttribute{
						Description:         "When you are creating a read replica from one Amazon Web Services GovCloud (US) Region to another or from one China Amazon Web Services Region to another, the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplica API operation in the source Amazon Web Services Region that contains the source DB instance. This setting applies only to Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. It's ignored in other Amazon Web Services Regions. This setting applies only when replicating from a source DB instance. Source DB clusters aren't supported in Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. You must specify this parameter when you create an encrypted read replica from another Amazon Web Services Region by using the Amazon RDS API. Don't specify PreSignedUrl when you are creating an encrypted read replica in the same Amazon Web Services Region. The presigned URL must be a valid request for the CreateDBInstanceReadReplica API operation that can run in the source Amazon Web Services Region that contains the encrypted source DB instance. The presigned URL request must contain the following parameter values: * DestinationRegion - The Amazon Web Services Region that the encrypted read replica is created in. This Amazon Web Services Region is the same one where the CreateDBInstanceReadReplica operation is called that contains this presigned URL. For example, if you create an encrypted DB instance in the us-west-1 Amazon Web Services Region, from a source DB instance in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica operation in the us-east-1 Amazon Web Services Region and provide a presigned URL that contains a call to the CreateDBInstanceReadReplica operation in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion in the presigned URL must be set to the us-east-1 Amazon Web Services Region. * KmsKeyId - The KMS key identifier for the key to use to encrypt the read replica in the destination Amazon Web Services Region. This is the same identifier for both the CreateDBInstanceReadReplica operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL. * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted DB instance to be replicated. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are creating an encrypted read replica from a DB instance in the us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115. To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region. This setting doesn't apply to RDS Custom DB instances.",
						MarkdownDescription: "When you are creating a read replica from one Amazon Web Services GovCloud (US) Region to another or from one China Amazon Web Services Region to another, the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplica API operation in the source Amazon Web Services Region that contains the source DB instance. This setting applies only to Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. It's ignored in other Amazon Web Services Regions. This setting applies only when replicating from a source DB instance. Source DB clusters aren't supported in Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. You must specify this parameter when you create an encrypted read replica from another Amazon Web Services Region by using the Amazon RDS API. Don't specify PreSignedUrl when you are creating an encrypted read replica in the same Amazon Web Services Region. The presigned URL must be a valid request for the CreateDBInstanceReadReplica API operation that can run in the source Amazon Web Services Region that contains the encrypted source DB instance. The presigned URL request must contain the following parameter values: * DestinationRegion - The Amazon Web Services Region that the encrypted read replica is created in. This Amazon Web Services Region is the same one where the CreateDBInstanceReadReplica operation is called that contains this presigned URL. For example, if you create an encrypted DB instance in the us-west-1 Amazon Web Services Region, from a source DB instance in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica operation in the us-east-1 Amazon Web Services Region and provide a presigned URL that contains a call to the CreateDBInstanceReadReplica operation in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion in the presigned URL must be set to the us-east-1 Amazon Web Services Region. * KmsKeyId - The KMS key identifier for the key to use to encrypt the read replica in the destination Amazon Web Services Region. This is the same identifier for both the CreateDBInstanceReadReplica operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL. * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted DB instance to be replicated. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are creating an encrypted read replica from a DB instance in the us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115. To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html). If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region. This setting doesn't apply to RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_backup_window": schema.StringAttribute{
						Description:         "The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. For more information, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the Amazon RDS User Guide. This setting doesn't apply to Amazon Aurora DB instances. The daily time range for creating automated backups is managed by the DB cluster. Constraints: * Must be in the format hh24:mi-hh24:mi. * Must be in Universal Coordinated Time (UTC). * Must not conflict with the preferred maintenance window. * Must be at least 30 minutes.",
						MarkdownDescription: "The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. For more information, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the Amazon RDS User Guide. This setting doesn't apply to Amazon Aurora DB instances. The daily time range for creating automated backups is managed by the DB cluster. Constraints: * Must be in the format hh24:mi-hh24:mi. * Must be in Universal Coordinated Time (UTC). * Must not conflict with the preferred maintenance window. * Must be at least 30 minutes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_maintenance_window": schema.StringAttribute{
						Description:         "The time range each week during which system maintenance can occur. For more information, see Amazon RDS Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance) in the Amazon RDS User Guide. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week. Constraints: * Must be in the format ddd:hh24:mi-ddd:hh24:mi. * The day values must be mon | tue | wed | thu | fri | sat | sun. * Must be in Universal Coordinated Time (UTC). * Must not conflict with the preferred backup window. * Must be at least 30 minutes.",
						MarkdownDescription: "The time range each week during which system maintenance can occur. For more information, see Amazon RDS Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance) in the Amazon RDS User Guide. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week. Constraints: * Must be in the format ddd:hh24:mi-ddd:hh24:mi. * The day values must be mon | tue | wed | thu | fri | sat | sun. * Must be in Universal Coordinated Time (UTC). * Must not conflict with the preferred backup window. * Must be at least 30 minutes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"processor_features": schema.ListNestedAttribute{
						Description:         "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						MarkdownDescription: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"promotion_tier": schema.Int64Attribute{
						Description:         "The order of priority in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.AuroraHighAvailability.html#Aurora.Managing.FaultTolerance) in the Amazon Aurora User Guide. This setting doesn't apply to RDS Custom DB instances. Default: 1 Valid Values: 0 - 15",
						MarkdownDescription: "The order of priority in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.AuroraHighAvailability.html#Aurora.Managing.FaultTolerance) in the Amazon Aurora User Guide. This setting doesn't apply to RDS Custom DB instances. Default: 1 Valid Values: 0 - 15",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"publicly_accessible": schema.BoolAttribute{
						Description:         "Specifies whether the DB instance is publicly accessible. When the DB instance is publicly accessible and you connect from outside of the DB instance's virtual private cloud (VPC), its Domain Name System (DNS) endpoint resolves to the public IP address. When you connect from within the same VPC as the DB instance, the endpoint resolves to the private IP address. Access to the DB instance is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB instance doesn't permit it. When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a private IP address. Default: The default behavior varies depending on whether DBSubnetGroupName is specified. If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies: * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB instance is private. * If the default VPC in the target Region has an internet gateway attached to it, the DB instance is public. If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies: * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB instance is private. * If the subnets are part of a VPC that has an internet gateway attached to it, the DB instance is public.",
						MarkdownDescription: "Specifies whether the DB instance is publicly accessible. When the DB instance is publicly accessible and you connect from outside of the DB instance's virtual private cloud (VPC), its Domain Name System (DNS) endpoint resolves to the public IP address. When you connect from within the same VPC as the DB instance, the endpoint resolves to the private IP address. Access to the DB instance is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB instance doesn't permit it. When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a private IP address. Default: The default behavior varies depending on whether DBSubnetGroupName is specified. If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies: * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB instance is private. * If the default VPC in the target Region has an internet gateway attached to it, the DB instance is public. If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies: * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB instance is private. * If the subnets are part of a VPC that has an internet gateway attached to it, the DB instance is public.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_mode": schema.StringAttribute{
						Description:         "The open mode of the replica database: mounted or read-only. This parameter is only supported for Oracle DB instances. Mounted DB replicas are included in Oracle Database Enterprise Edition. The main use case for mounted replicas is cross-Region disaster recovery. The primary database doesn't use Active Data Guard to transmit information to the mounted replica. Because it doesn't accept user connections, a mounted replica can't serve a read-only workload. You can create a combination of mounted and read-only DB replicas for the same primary DB instance. For more information, see Working with Oracle Read Replicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide. For RDS Custom, you must specify this parameter and set it to mounted. The value won't be set by default. After replica creation, you can manage the open mode manually.",
						MarkdownDescription: "The open mode of the replica database: mounted or read-only. This parameter is only supported for Oracle DB instances. Mounted DB replicas are included in Oracle Database Enterprise Edition. The main use case for mounted replicas is cross-Region disaster recovery. The primary database doesn't use Active Data Guard to transmit information to the mounted replica. Because it doesn't accept user connections, a mounted replica can't serve a read-only workload. You can create a combination of mounted and read-only DB replicas for the same primary DB instance. For more information, see Working with Oracle Read Replicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide. For RDS Custom, you must specify this parameter and set it to mounted. The value won't be set by default. After replica creation, you can manage the open mode manually.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_db_instance_identifier": schema.StringAttribute{
						Description:         "The identifier of the DB instance that will act as the source for the read replica. Each DB instance can have up to 15 read replicas, with the exception of Oracle and SQL Server, which can have up to five. Constraints: * Must be the identifier of an existing Db2, MariaDB, MySQL, Oracle, PostgreSQL, or SQL Server DB instance. * Can't be specified if the SourceDBClusterIdentifier parameter is also specified. * For the limitations of Oracle read replicas, see Version and licensing considerations for RDS for Oracle replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.limitations.html#oracle-read-replicas.limitations.versions-and-licenses) in the Amazon RDS User Guide. * For the limitations of SQL Server read replicas, see Read replica limitations with SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.html#SQLServer.ReadReplicas.Limitations) in the Amazon RDS User Guide. * The specified DB instance must have automatic backups enabled, that is, its backup retention period must be greater than 0. * If the source DB instance is in the same Amazon Web Services Region as the read replica, specify a valid DB instance identifier. * If the source DB instance is in a different Amazon Web Services Region from the read replica, specify a valid DB instance ARN. For more information, see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing) in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS Custom, which don't support cross-Region replicas.",
						MarkdownDescription: "The identifier of the DB instance that will act as the source for the read replica. Each DB instance can have up to 15 read replicas, with the exception of Oracle and SQL Server, which can have up to five. Constraints: * Must be the identifier of an existing Db2, MariaDB, MySQL, Oracle, PostgreSQL, or SQL Server DB instance. * Can't be specified if the SourceDBClusterIdentifier parameter is also specified. * For the limitations of Oracle read replicas, see Version and licensing considerations for RDS for Oracle replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.limitations.html#oracle-read-replicas.limitations.versions-and-licenses) in the Amazon RDS User Guide. * For the limitations of SQL Server read replicas, see Read replica limitations with SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.html#SQLServer.ReadReplicas.Limitations) in the Amazon RDS User Guide. * The specified DB instance must have automatic backups enabled, that is, its backup retention period must be greater than 0. * If the source DB instance is in the same Amazon Web Services Region as the read replica, specify a valid DB instance identifier. * If the source DB instance is in a different Amazon Web Services Region from the read replica, specify a valid DB instance ARN. For more information, see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing) in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS Custom, which don't support cross-Region replicas.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_region": schema.StringAttribute{
						Description:         "SourceRegion is the source region where the resource exists. This is not sent over the wire and is only used for presigning. This value should always have the same region as the source ARN.",
						MarkdownDescription: "SourceRegion is the source region where the resource exists. This is not sent over the wire and is only used for presigning. This value should always have the same region as the source ARN.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_encrypted": schema.BoolAttribute{
						Description:         "Specifes whether the DB instance is encrypted. By default, it isn't encrypted. For RDS Custom DB instances, either enable this setting or leave it unset. Otherwise, Amazon RDS reports an error. This setting doesn't apply to Amazon Aurora DB instances. The encryption for DB instances is managed by the DB cluster.",
						MarkdownDescription: "Specifes whether the DB instance is encrypted. By default, it isn't encrypted. For RDS Custom DB instances, either enable this setting or leave it unset. Otherwise, Amazon RDS reports an error. This setting doesn't apply to Amazon Aurora DB instances. The encryption for DB instances is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_throughput": schema.Int64Attribute{
						Description:         "The storage throughput value for the DB instance. This setting applies only to the gp3 storage type. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						MarkdownDescription: "The storage throughput value for the DB instance. This setting applies only to the gp3 storage type. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_type": schema.StringAttribute{
						Description:         "The storage type to associate with the DB instance. If you specify io1, io2, or gp3, you must also include a value for the Iops parameter. This setting doesn't apply to Amazon Aurora DB instances. Storage is managed by the DB cluster. Valid Values: gp2 | gp3 | io1 | io2 | standard Default: io1, if the Iops parameter is specified. Otherwise, gp2.",
						MarkdownDescription: "The storage type to associate with the DB instance. If you specify io1, io2, or gp3, you must also include a value for the Iops parameter. This setting doesn't apply to Amazon Aurora DB instances. Storage is managed by the DB cluster. Valid Values: gp2 | gp3 | io1 | io2 | standard Default: io1, if the Iops parameter is specified. Otherwise, gp2.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Tags to assign to the DB instance.",
						MarkdownDescription: "Tags to assign to the DB instance.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tde_credential_arn": schema.StringAttribute{
						Description:         "The ARN from the key store with which to associate the instance for TDE encryption. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						MarkdownDescription: "The ARN from the key store with which to associate the instance for TDE encryption. This setting doesn't apply to Amazon Aurora or RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tde_credential_password": schema.StringAttribute{
						Description:         "The password for the given ARN from the key store in order to access the device. This setting doesn't apply to RDS Custom DB instances.",
						MarkdownDescription: "The password for the given ARN from the key store in order to access the device. This setting doesn't apply to RDS Custom DB instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timezone": schema.StringAttribute{
						Description:         "The time zone of the DB instance. The time zone parameter is currently supported only by RDS for Db2 (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-time-zone) and RDS for SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
						MarkdownDescription: "The time zone of the DB instance. The time zone parameter is currently supported only by RDS for Db2 (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/db2-time-zone) and RDS for SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_default_processor_features": schema.BoolAttribute{
						Description:         "Specifies whether the DB instance class of the DB instance uses its default processor features. This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "Specifies whether the DB instance class of the DB instance uses its default processor features. This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vpc_security_group_i_ds": schema.ListAttribute{
						Description:         "A list of Amazon EC2 VPC security groups to associate with this DB instance. This setting doesn't apply to Amazon Aurora DB instances. The associated list of EC2 VPC security groups is managed by the DB cluster. Default: The default EC2 VPC security group for the DB subnet group's VPC.",
						MarkdownDescription: "A list of Amazon EC2 VPC security groups to associate with this DB instance. This setting doesn't apply to Amazon Aurora DB instances. The associated list of EC2 VPC security groups is managed by the DB cluster. Default: The default EC2 VPC security group for the DB subnet group's VPC.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rds_services_k8s_aws_db_instance_v1alpha1_manifest")

	var model RdsServicesK8SAwsDbinstanceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBInstance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
