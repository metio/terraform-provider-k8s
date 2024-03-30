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
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"db_parameter_group_ref" json:"dbParameterGroupRef,omitempty"`
		DbSnapshotIdentifier *string `tfsdk:"db_snapshot_identifier" json:"dbSnapshotIdentifier,omitempty"`
		DbSubnetGroupName    *string `tfsdk:"db_subnet_group_name" json:"dbSubnetGroupName,omitempty"`
		DbSubnetGroupRef     *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
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
				Name *string `tfsdk:"name" json:"name,omitempty"`
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
				Name *string `tfsdk:"name" json:"name,omitempty"`
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
				Name *string `tfsdk:"name" json:"name,omitempty"`
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
				Description:         "DBInstanceSpec defines the desired state of DBInstance.Contains the details of an Amazon RDS DB instance.This data type is used as a response element in the operations CreateDBInstance,CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance,PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3,RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",
				MarkdownDescription: "DBInstanceSpec defines the desired state of DBInstance.Contains the details of an Amazon RDS DB instance.This data type is used as a response element in the operations CreateDBInstance,CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance,PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3,RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",
				Attributes: map[string]schema.Attribute{
					"allocated_storage": schema.Int64Attribute{
						Description:         "The amount of storage in gibibytes (GiB) to allocate for the DB instance.Type: IntegerAmazon AuroraNot applicable. Aurora cluster volumes automatically grow as the amount ofdata in your database increases, though you are only charged for the spacethat you use in an Aurora cluster volume.Amazon RDS CustomConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 40   to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.   * Provisioned IOPS storage (io1): Must be an integer from 40 to 65536   for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.MySQLConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 5 to 3072.MariaDBConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 5 to 3072.PostgreSQLConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 5 to 3072.OracleConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 10 to 3072.SQL ServerConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Enterprise and Standard editions:   Must be an integer from 20 to 16384. Web and Express editions: Must be   an integer from 20 to 16384.   * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must   be an integer from 100 to 16384. Web and Express editions: Must be an   integer from 100 to 16384.   * Magnetic storage (standard): Enterprise and Standard editions: Must   be an integer from 20 to 1024. Web and Express editions: Must be an integer   from 20 to 1024.",
						MarkdownDescription: "The amount of storage in gibibytes (GiB) to allocate for the DB instance.Type: IntegerAmazon AuroraNot applicable. Aurora cluster volumes automatically grow as the amount ofdata in your database increases, though you are only charged for the spacethat you use in an Aurora cluster volume.Amazon RDS CustomConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 40   to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.   * Provisioned IOPS storage (io1): Must be an integer from 40 to 65536   for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.MySQLConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 5 to 3072.MariaDBConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 5 to 3072.PostgreSQLConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 5 to 3072.OracleConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20   to 65536.   * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.   * Magnetic storage (standard): Must be an integer from 10 to 3072.SQL ServerConstraints to the amount of storage for each storage type are the following:   * General Purpose (SSD) storage (gp2, gp3): Enterprise and Standard editions:   Must be an integer from 20 to 16384. Web and Express editions: Must be   an integer from 20 to 16384.   * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must   be an integer from 100 to 16384. Web and Express editions: Must be an   integer from 100 to 16384.   * Magnetic storage (standard): Enterprise and Standard editions: Must   be an integer from 20 to 1024. Web and Express editions: Must be an integer   from 20 to 1024.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auto_minor_version_upgrade": schema.BoolAttribute{
						Description:         "A value that indicates whether minor engine upgrades are applied automaticallyto the DB instance during the maintenance window. By default, minor engineupgrades are applied automatically.If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgradeto false.",
						MarkdownDescription: "A value that indicates whether minor engine upgrades are applied automaticallyto the DB instance during the maintenance window. By default, minor engineupgrades are applied automatically.If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgradeto false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"availability_zone": schema.StringAttribute{
						Description:         "The Availability Zone (AZ) where the database will be created. For informationon Amazon Web Services Regions and Availability Zones, see Regions and AvailabilityZones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).Amazon AuroraEach Aurora DB cluster hosts copies of its storage in three separate AvailabilityZones. Specify one of these Availability Zones. Aurora automatically choosesan appropriate Availability Zone if you don't specify one.Default: A random, system-chosen Availability Zone in the endpoint's AmazonWeb Services Region.Example: us-east-1dConstraint: The AvailabilityZone parameter can't be specified if the DB instanceis a Multi-AZ deployment. The specified Availability Zone must be in thesame Amazon Web Services Region as the current endpoint.",
						MarkdownDescription: "The Availability Zone (AZ) where the database will be created. For informationon Amazon Web Services Regions and Availability Zones, see Regions and AvailabilityZones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).Amazon AuroraEach Aurora DB cluster hosts copies of its storage in three separate AvailabilityZones. Specify one of these Availability Zones. Aurora automatically choosesan appropriate Availability Zone if you don't specify one.Default: A random, system-chosen Availability Zone in the endpoint's AmazonWeb Services Region.Example: us-east-1dConstraint: The AvailabilityZone parameter can't be specified if the DB instanceis a Multi-AZ deployment. The specified Availability Zone must be in thesame Amazon Web Services Region as the current endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_retention_period": schema.Int64Attribute{
						Description:         "The number of days for which automated backups are retained. Setting thisparameter to a positive number enables backups. Setting this parameter to0 disables automated backups.Amazon AuroraNot applicable. The retention period for automated backups is managed bythe DB cluster.Default: 1Constraints:   * Must be a value from 0 to 35   * Can't be set to 0 if the DB instance is a source to read replicas   * Can't be set to 0 for an RDS Custom for Oracle DB instance",
						MarkdownDescription: "The number of days for which automated backups are retained. Setting thisparameter to a positive number enables backups. Setting this parameter to0 disables automated backups.Amazon AuroraNot applicable. The retention period for automated backups is managed bythe DB cluster.Default: 1Constraints:   * Must be a value from 0 to 35   * Can't be set to 0 if the DB instance is a source to read replicas   * Can't be set to 0 for an RDS Custom for Oracle DB instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_target": schema.StringAttribute{
						Description:         "Specifies where automated backups and manual snapshots are stored.Possible values are outposts (Amazon Web Services Outposts) and region (AmazonWeb Services Region). The default is region.For more information, see Working with Amazon RDS on Amazon Web ServicesOutposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)in the Amazon RDS User Guide.",
						MarkdownDescription: "Specifies where automated backups and manual snapshots are stored.Possible values are outposts (Amazon Web Services Outposts) and region (AmazonWeb Services Region). The default is region.For more information, see Working with Amazon RDS on Amazon Web ServicesOutposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ca_certificate_identifier": schema.StringAttribute{
						Description:         "Specifies the CA certificate identifier to use for the DB instance’s servercertificate.This setting doesn't apply to RDS Custom.For more information, see Using SSL/TLS to encrypt a connection to a DB instance(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection toa DB cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html)in the Amazon Aurora User Guide.",
						MarkdownDescription: "Specifies the CA certificate identifier to use for the DB instance’s servercertificate.This setting doesn't apply to RDS Custom.For more information, see Using SSL/TLS to encrypt a connection to a DB instance(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection toa DB cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html)in the Amazon Aurora User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"character_set_name": schema.StringAttribute{
						Description:         "For supported engines, this value indicates that the DB instance should beassociated with the specified CharacterSet.This setting doesn't apply to RDS Custom. However, if you need to changethe character set, you can change it on the database itself.Amazon AuroraNot applicable. The character set is managed by the DB cluster. For moreinformation, see CreateDBCluster.",
						MarkdownDescription: "For supported engines, this value indicates that the DB instance should beassociated with the specified CharacterSet.This setting doesn't apply to RDS Custom. However, if you need to changethe character set, you can change it on the database itself.Amazon AuroraNot applicable. The character set is managed by the DB cluster. For moreinformation, see CreateDBCluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"copy_tags_to_snapshot": schema.BoolAttribute{
						Description:         "A value that indicates whether to copy tags from the DB instance to snapshotsof the DB instance. By default, tags are not copied.Amazon AuroraNot applicable. Copying tags to snapshots is managed by the DB cluster. Settingthis value for an Aurora DB instance has no effect on the DB cluster setting.",
						MarkdownDescription: "A value that indicates whether to copy tags from the DB instance to snapshotsof the DB instance. By default, tags are not copied.Amazon AuroraNot applicable. Copying tags to snapshots is managed by the DB cluster. Settingthis value for an Aurora DB instance has no effect on the DB cluster setting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_iam_instance_profile": schema.StringAttribute{
						Description:         "The instance profile associated with the underlying Amazon EC2 instance ofan RDS Custom DB instance. The instance profile must meet the following requirements:   * The profile must exist in your account.   * The profile must have an IAM role that Amazon EC2 has permissions to   assume.   * The instance profile name and the associated IAM role name must start   with the prefix AWSRDSCustom.For the list of permissions required for the IAM role, see Configure IAMand your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc)in the Amazon RDS User Guide.This setting is required for RDS Custom.",
						MarkdownDescription: "The instance profile associated with the underlying Amazon EC2 instance ofan RDS Custom DB instance. The instance profile must meet the following requirements:   * The profile must exist in your account.   * The profile must have an IAM role that Amazon EC2 has permissions to   assume.   * The instance profile name and the associated IAM role name must start   with the prefix AWSRDSCustom.For the list of permissions required for the IAM role, see Configure IAMand your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc)in the Amazon RDS User Guide.This setting is required for RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_identifier": schema.StringAttribute{
						Description:         "The identifier of the DB cluster that the instance will belong to.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The identifier of the DB cluster that the instance will belong to.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_snapshot_identifier": schema.StringAttribute{
						Description:         "The identifier for the RDS for MySQL Multi-AZ DB cluster snapshot to restorefrom.For more information on Multi-AZ DB clusters, see Multi-AZ DB cluster deployments(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)in the Amazon RDS User Guide.Constraints:   * Must match the identifier of an existing Multi-AZ DB cluster snapshot.   * Can't be specified when DBSnapshotIdentifier is specified.   * Must be specified when DBSnapshotIdentifier isn't specified.   * If you are restoring from a shared manual Multi-AZ DB cluster snapshot,   the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot.   * Can't be the identifier of an Aurora DB cluster snapshot.   * Can't be the identifier of an RDS for PostgreSQL Multi-AZ DB cluster   snapshot.",
						MarkdownDescription: "The identifier for the RDS for MySQL Multi-AZ DB cluster snapshot to restorefrom.For more information on Multi-AZ DB clusters, see Multi-AZ DB cluster deployments(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)in the Amazon RDS User Guide.Constraints:   * Must match the identifier of an existing Multi-AZ DB cluster snapshot.   * Can't be specified when DBSnapshotIdentifier is specified.   * Must be specified when DBSnapshotIdentifier isn't specified.   * If you are restoring from a shared manual Multi-AZ DB cluster snapshot,   the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot.   * Can't be the identifier of an Aurora DB cluster snapshot.   * Can't be the identifier of an RDS for PostgreSQL Multi-AZ DB cluster   snapshot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_instance_class": schema.StringAttribute{
						Description:         "The compute and memory capacity of the DB instance, for example db.m5.large.Not all DB instance classes are available in all Amazon Web Services Regions,or for all database engines. For the full list of DB instance classes, andavailability for your engine, see DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)in the Amazon RDS User Guide or Aurora DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html)in the Amazon Aurora User Guide.",
						MarkdownDescription: "The compute and memory capacity of the DB instance, for example db.m5.large.Not all DB instance classes are available in all Amazon Web Services Regions,or for all database engines. For the full list of DB instance classes, andavailability for your engine, see DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html)in the Amazon RDS User Guide or Aurora DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html)in the Amazon Aurora User Guide.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_instance_identifier": schema.StringAttribute{
						Description:         "The DB instance identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * First character must be a letter.   * Can't end with a hyphen or contain two consecutive hyphens.Example: mydbinstance",
						MarkdownDescription: "The DB instance identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * First character must be a letter.   * Can't end with a hyphen or contain two consecutive hyphens.Example: mydbinstance",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_name": schema.StringAttribute{
						Description:         "The meaning of this parameter differs according to the database engine youuse.MySQLThe name of the database to create when the DB instance is created. If thisparameter isn't specified, no database is created in the DB instance.Constraints:   * Must contain 1 to 64 letters or numbers.   * Must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0-9).   * Can't be a word reserved by the specified database engineMariaDBThe name of the database to create when the DB instance is created. If thisparameter isn't specified, no database is created in the DB instance.Constraints:   * Must contain 1 to 64 letters or numbers.   * Must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0-9).   * Can't be a word reserved by the specified database enginePostgreSQLThe name of the database to create when the DB instance is created. If thisparameter isn't specified, a database named postgres is created in the DBinstance.Constraints:   * Must contain 1 to 63 letters, numbers, or underscores.   * Must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0-9).   * Can't be a word reserved by the specified database engineOracleThe Oracle System ID (SID) of the created DB instance. If you specify null,the default value ORCL is used. You can't specify the string NULL, or anyother reserved word, for DBName.Default: ORCLConstraints:   * Can't be longer than 8 charactersAmazon RDS Custom for OracleThe Oracle System ID (SID) of the created RDS Custom DB instance. If youdon't specify a value, the default value is ORCL.Default: ORCLConstraints:   * It must contain 1 to 8 alphanumeric characters.   * It must contain a letter.   * It can't be a word reserved by the database engine.Amazon RDS Custom for SQL ServerNot applicable. Must be null.SQL ServerNot applicable. Must be null.Amazon Aurora MySQLThe name of the database to create when the primary DB instance of the AuroraMySQL DB cluster is created. If this parameter isn't specified for an AuroraMySQL DB cluster, no database is created in the DB cluster.Constraints:   * It must contain 1 to 64 alphanumeric characters.   * It can't be a word reserved by the database engine.Amazon Aurora PostgreSQLThe name of the database to create when the primary DB instance of the AuroraPostgreSQL DB cluster is created. If this parameter isn't specified for anAurora PostgreSQL DB cluster, a database named postgres is created in theDB cluster.Constraints:   * It must contain 1 to 63 alphanumeric characters.   * It must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0 to 9).   * It can't be a word reserved by the database engine.",
						MarkdownDescription: "The meaning of this parameter differs according to the database engine youuse.MySQLThe name of the database to create when the DB instance is created. If thisparameter isn't specified, no database is created in the DB instance.Constraints:   * Must contain 1 to 64 letters or numbers.   * Must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0-9).   * Can't be a word reserved by the specified database engineMariaDBThe name of the database to create when the DB instance is created. If thisparameter isn't specified, no database is created in the DB instance.Constraints:   * Must contain 1 to 64 letters or numbers.   * Must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0-9).   * Can't be a word reserved by the specified database enginePostgreSQLThe name of the database to create when the DB instance is created. If thisparameter isn't specified, a database named postgres is created in the DBinstance.Constraints:   * Must contain 1 to 63 letters, numbers, or underscores.   * Must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0-9).   * Can't be a word reserved by the specified database engineOracleThe Oracle System ID (SID) of the created DB instance. If you specify null,the default value ORCL is used. You can't specify the string NULL, or anyother reserved word, for DBName.Default: ORCLConstraints:   * Can't be longer than 8 charactersAmazon RDS Custom for OracleThe Oracle System ID (SID) of the created RDS Custom DB instance. If youdon't specify a value, the default value is ORCL.Default: ORCLConstraints:   * It must contain 1 to 8 alphanumeric characters.   * It must contain a letter.   * It can't be a word reserved by the database engine.Amazon RDS Custom for SQL ServerNot applicable. Must be null.SQL ServerNot applicable. Must be null.Amazon Aurora MySQLThe name of the database to create when the primary DB instance of the AuroraMySQL DB cluster is created. If this parameter isn't specified for an AuroraMySQL DB cluster, no database is created in the DB cluster.Constraints:   * It must contain 1 to 64 alphanumeric characters.   * It can't be a word reserved by the database engine.Amazon Aurora PostgreSQLThe name of the database to create when the primary DB instance of the AuroraPostgreSQL DB cluster is created. If this parameter isn't specified for anAurora PostgreSQL DB cluster, a database named postgres is created in theDB cluster.Constraints:   * It must contain 1 to 63 alphanumeric characters.   * It must begin with a letter. Subsequent characters can be letters, underscores,   or digits (0 to 9).   * It can't be a word reserved by the database engine.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_parameter_group_name": schema.StringAttribute{
						Description:         "The name of the DB parameter group to associate with this DB instance. Ifyou do not specify a value, then the default DB parameter group for the specifiedDB engine and version is used.This setting doesn't apply to RDS Custom.Constraints:   * It must be 1 to 255 letters, numbers, or hyphens.   * The first character must be a letter.   * It can't end with a hyphen or contain two consecutive hyphens.",
						MarkdownDescription: "The name of the DB parameter group to associate with this DB instance. Ifyou do not specify a value, then the default DB parameter group for the specifiedDB engine and version is used.This setting doesn't apply to RDS Custom.Constraints:   * It must be 1 to 255 letters, numbers, or hyphens.   * The first character must be a letter.   * It can't end with a hyphen or contain two consecutive hyphens.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_parameter_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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
						Description:         "The identifier for the DB snapshot to restore from.Constraints:   * Must match the identifier of an existing DBSnapshot.   * Can't be specified when DBClusterSnapshotIdentifier is specified.   * Must be specified when DBClusterSnapshotIdentifier isn't specified.   * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier   must be the ARN of the shared DB snapshot.",
						MarkdownDescription: "The identifier for the DB snapshot to restore from.Constraints:   * Must match the identifier of an existing DBSnapshot.   * Can't be specified when DBClusterSnapshotIdentifier is specified.   * Must be specified when DBClusterSnapshotIdentifier isn't specified.   * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier   must be the ARN of the shared DB snapshot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_subnet_group_name": schema.StringAttribute{
						Description:         "A DB subnet group to associate with this DB instance.Constraints: Must match the name of an existing DBSubnetGroup. Must not bedefault.Example: mydbsubnetgroup",
						MarkdownDescription: "A DB subnet group to associate with this DB instance.Constraints: Must match the name of an existing DBSubnetGroup. Must not bedefault.Example: mydbsubnetgroup",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_subnet_group_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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
						Description:         "A value that indicates whether the DB instance has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default,deletion protection isn't enabled. For more information, see Deleting a DBInstance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).Amazon AuroraNot applicable. You can enable or disable deletion protection for the DBcluster. For more information, see CreateDBCluster. DB instances in a DBcluster can be deleted even when deletion protection is enabled for the DBcluster.",
						MarkdownDescription: "A value that indicates whether the DB instance has deletion protection enabled.The database can't be deleted when deletion protection is enabled. By default,deletion protection isn't enabled. For more information, see Deleting a DBInstance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).Amazon AuroraNot applicable. You can enable or disable deletion protection for the DBcluster. For more information, see CreateDBCluster. DB instances in a DBcluster can be deleted even when deletion protection is enabled for the DBcluster.",
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
						Description:         "The Active Directory directory ID to create the DB instance in. Currently,only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances canbe created in an Active Directory Domain.For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. The domain is managed by the DB cluster.",
						MarkdownDescription: "The Active Directory directory ID to create the DB instance in. Currently,only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances canbe created in an Active Directory Domain.For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. The domain is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"domain_iam_role_name": schema.StringAttribute{
						Description:         "Specify the name of the IAM role to be used when making API calls to theDirectory Service.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. The domain is managed by the DB cluster.",
						MarkdownDescription: "Specify the name of the IAM role to be used when making API calls to theDirectory Service.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. The domain is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_cloudwatch_logs_exports": schema.ListAttribute{
						Description:         "The list of log types that need to be enabled for exporting to CloudWatchLogs. The values in the list depend on the DB engine. For more information,see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)in the Amazon RDS User Guide.Amazon AuroraNot applicable. CloudWatch Logs exports are managed by the DB cluster.RDS CustomNot applicable.MariaDBPossible values are audit, error, general, and slowquery.Microsoft SQL ServerPossible values are agent and error.MySQLPossible values are audit, error, general, and slowquery.OraclePossible values are alert, audit, listener, trace, and oemagent.PostgreSQLPossible values are postgresql and upgrade.",
						MarkdownDescription: "The list of log types that need to be enabled for exporting to CloudWatchLogs. The values in the list depend on the DB engine. For more information,see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch)in the Amazon RDS User Guide.Amazon AuroraNot applicable. CloudWatch Logs exports are managed by the DB cluster.RDS CustomNot applicable.MariaDBPossible values are audit, error, general, and slowquery.Microsoft SQL ServerPossible values are agent and error.MySQLPossible values are audit, error, general, and slowquery.OraclePossible values are alert, audit, listener, trace, and oemagent.PostgreSQLPossible values are postgresql and upgrade.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_customer_owned_ip": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable a customer-owned IP address (CoIP)for an RDS on Outposts DB instance.A CoIP provides local or external connectivity to resources in your Outpostsubnets through your on-premises network. For some use cases, a CoIP canprovide lower latency for connections to the DB instance from outside ofits virtual private cloud (VPC) on your local network.For more information about RDS on Outposts, see Working with Amazon RDS onAmazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)in the Amazon RDS User Guide.For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing)in the Amazon Web Services Outposts User Guide.",
						MarkdownDescription: "A value that indicates whether to enable a customer-owned IP address (CoIP)for an RDS on Outposts DB instance.A CoIP provides local or external connectivity to resources in your Outpostsubnets through your on-premises network. For some use cases, a CoIP canprovide lower latency for connections to the DB instance from outside ofits virtual private cloud (VPC) on your local network.For more information about RDS on Outposts, see Working with Amazon RDS onAmazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html)in the Amazon RDS User Guide.For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing)in the Amazon Web Services Outposts User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_iam_database_authentication": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable mapping of Amazon Web Services Identityand Access Management (IAM) accounts to database accounts. By default, mappingisn't enabled.For more information, see IAM Database Authentication for MySQL and PostgreSQL(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. Mapping Amazon Web Services IAM accounts to database accountsis managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether to enable mapping of Amazon Web Services Identityand Access Management (IAM) accounts to database accounts. By default, mappingisn't enabled.For more information, see IAM Database Authentication for MySQL and PostgreSQL(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. Mapping Amazon Web Services IAM accounts to database accountsis managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine": schema.StringAttribute{
						Description:         "The name of the database engine to be used for this instance.Not every database engine is available for every Amazon Web Services Region.Valid Values:   * aurora (for MySQL 5.6-compatible Aurora)   * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)   * aurora-postgresql   * custom-oracle-ee (for RDS Custom for Oracle instances)   * custom-sqlserver-ee (for RDS Custom for SQL Server instances)   * custom-sqlserver-se (for RDS Custom for SQL Server instances)   * custom-sqlserver-web (for RDS Custom for SQL Server instances)   * mariadb   * mysql   * oracle-ee   * oracle-ee-cdb   * oracle-se2   * oracle-se2-cdb   * postgres   * sqlserver-ee   * sqlserver-se   * sqlserver-ex   * sqlserver-web",
						MarkdownDescription: "The name of the database engine to be used for this instance.Not every database engine is available for every Amazon Web Services Region.Valid Values:   * aurora (for MySQL 5.6-compatible Aurora)   * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)   * aurora-postgresql   * custom-oracle-ee (for RDS Custom for Oracle instances)   * custom-sqlserver-ee (for RDS Custom for SQL Server instances)   * custom-sqlserver-se (for RDS Custom for SQL Server instances)   * custom-sqlserver-web (for RDS Custom for SQL Server instances)   * mariadb   * mysql   * oracle-ee   * oracle-ee-cdb   * oracle-se2   * oracle-se2-cdb   * postgres   * sqlserver-ee   * sqlserver-se   * sqlserver-ex   * sqlserver-web",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The version number of the database engine to use.For a list of valid engine versions, use the DescribeDBEngineVersions operation.The following are the database engines and links to information about themajor and minor versions that are available with Amazon RDS. Not every databaseengine is available for every Amazon Web Services Region.Amazon AuroraNot applicable. The version number of the database engine to be used by theDB instance is managed by the DB cluster.Amazon RDS Custom for OracleA custom engine version (CEV) that you have previously created. This settingis required for RDS Custom for Oracle. The CEV name has the following format:19.customized_string. A valid CEV name is 19.my_cev1. For more information,see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create)in the Amazon RDS User Guide.Amazon RDS Custom for SQL ServerSee RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html)in the Amazon RDS User Guide.MariaDBFor information, see MariaDB on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt)in the Amazon RDS User Guide.Microsoft SQL ServerFor information, see Microsoft SQL Server Versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport)in the Amazon RDS User Guide.MySQLFor information, see MySQL on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt)in the Amazon RDS User Guide.OracleFor information, see Oracle Database Engine Release Notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html)in the Amazon RDS User Guide.PostgreSQLFor information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts)in the Amazon RDS User Guide.",
						MarkdownDescription: "The version number of the database engine to use.For a list of valid engine versions, use the DescribeDBEngineVersions operation.The following are the database engines and links to information about themajor and minor versions that are available with Amazon RDS. Not every databaseengine is available for every Amazon Web Services Region.Amazon AuroraNot applicable. The version number of the database engine to be used by theDB instance is managed by the DB cluster.Amazon RDS Custom for OracleA custom engine version (CEV) that you have previously created. This settingis required for RDS Custom for Oracle. The CEV name has the following format:19.customized_string. A valid CEV name is 19.my_cev1. For more information,see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create)in the Amazon RDS User Guide.Amazon RDS Custom for SQL ServerSee RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html)in the Amazon RDS User Guide.MariaDBFor information, see MariaDB on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt)in the Amazon RDS User Guide.Microsoft SQL ServerFor information, see Microsoft SQL Server Versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport)in the Amazon RDS User Guide.MySQLFor information, see MySQL on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt)in the Amazon RDS User Guide.OracleFor information, see Oracle Database Engine Release Notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html)in the Amazon RDS User Guide.PostgreSQLFor information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts)in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"iops": schema.Int64Attribute{
						Description:         "The amount of Provisioned IOPS (input/output operations per second) to beinitially allocated for the DB instance. For information about valid IOPSvalues, see Amazon RDS DB instance storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html)in the Amazon RDS User Guide.Constraints: For MariaDB, MySQL, Oracle, and PostgreSQL DB instances, mustbe a multiple between .5 and 50 of the storage amount for the DB instance.For SQL Server DB instances, must be a multiple between 1 and 50 of the storageamount for the DB instance.Amazon AuroraNot applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "The amount of Provisioned IOPS (input/output operations per second) to beinitially allocated for the DB instance. For information about valid IOPSvalues, see Amazon RDS DB instance storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html)in the Amazon RDS User Guide.Constraints: For MariaDB, MySQL, Oracle, and PostgreSQL DB instances, mustbe a multiple between .5 and 50 of the storage amount for the DB instance.For SQL Server DB instances, must be a multiple between 1 and 50 of the storageamount for the DB instance.Amazon AuroraNot applicable. Storage is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for an encrypted DB instance.The Amazon Web Services KMS key identifier is the key ARN, key ID, aliasARN, or alias name for the KMS key. To use a KMS key in a different AmazonWeb Services account, specify the key ARN or alias ARN.Amazon AuroraNot applicable. The Amazon Web Services KMS key identifier is managed bythe DB cluster. For more information, see CreateDBCluster.If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyIdparameter, then Amazon RDS uses your default KMS key. There is a defaultKMS key for your Amazon Web Services account. Your Amazon Web Services accounthas a different default KMS key for each Amazon Web Services Region.Amazon RDS CustomA KMS key is required for RDS Custom instances. For most RDS engines, ifyou leave this parameter empty while enabling StorageEncrypted, the engineuses the default KMS key. However, RDS Custom doesn't use the default keywhen this parameter is empty. You must explicitly specify a key.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for an encrypted DB instance.The Amazon Web Services KMS key identifier is the key ARN, key ID, aliasARN, or alias name for the KMS key. To use a KMS key in a different AmazonWeb Services account, specify the key ARN or alias ARN.Amazon AuroraNot applicable. The Amazon Web Services KMS key identifier is managed bythe DB cluster. For more information, see CreateDBCluster.If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyIdparameter, then Amazon RDS uses your default KMS key. There is a defaultKMS key for your Amazon Web Services account. Your Amazon Web Services accounthas a different default KMS key for each Amazon Web Services Region.Amazon RDS CustomA KMS key is required for RDS Custom instances. For most RDS engines, ifyou leave this parameter empty while enabling StorageEncrypted, the engineuses the default KMS key. However, RDS Custom doesn't use the default keywhen this parameter is empty. You must explicitly specify a key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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
						Description:         "License model information for this DB instance.Valid values: license-included | bring-your-own-license | general-public-licenseThis setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
						MarkdownDescription: "License model information for this DB instance.Valid values: license-included | bring-your-own-license | general-public-licenseThis setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"manage_master_user_password": schema.BoolAttribute{
						Description:         "A value that indicates whether to manage the master user password with AmazonWeb Services Secrets Manager.For more information, see Password management with Amazon Web Services SecretsManager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)in the Amazon RDS User Guide.Constraints:   * Can't manage the master user password with Amazon Web Services Secrets   Manager if MasterUserPassword is specified.",
						MarkdownDescription: "A value that indicates whether to manage the master user password with AmazonWeb Services Secrets Manager.For more information, see Password management with Amazon Web Services SecretsManager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html)in the Amazon RDS User Guide.Constraints:   * Can't manage the master user password with Amazon Web Services Secrets   Manager if MasterUserPassword is specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"master_user_password": schema.SingleNestedAttribute{
						Description:         "The password for the master user. The password can include any printableASCII character except '/', ''', or '@'.Amazon AuroraNot applicable. The password for the master user is managed by the DB cluster.Constraints: Can't be specified if ManageMasterUserPassword is turned on.MariaDBConstraints: Must contain from 8 to 41 characters.Microsoft SQL ServerConstraints: Must contain from 8 to 128 characters.MySQLConstraints: Must contain from 8 to 41 characters.OracleConstraints: Must contain from 8 to 30 characters.PostgreSQLConstraints: Must contain from 8 to 128 characters.",
						MarkdownDescription: "The password for the master user. The password can include any printableASCII character except '/', ''', or '@'.Amazon AuroraNot applicable. The password for the master user is managed by the DB cluster.Constraints: Can't be specified if ManageMasterUserPassword is turned on.MariaDBConstraints: Must contain from 8 to 41 characters.Microsoft SQL ServerConstraints: Must contain from 8 to 128 characters.MySQLConstraints: Must contain from 8 to 41 characters.OracleConstraints: Must contain from 8 to 30 characters.PostgreSQLConstraints: Must contain from 8 to 128 characters.",
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
						Description:         "The Amazon Web Services KMS key identifier to encrypt a secret that is automaticallygenerated and managed in Amazon Web Services Secrets Manager.This setting is valid only if the master user password is managed by RDSin Amazon Web Services Secrets Manager for the DB instance.The Amazon Web Services KMS key identifier is the key ARN, key ID, aliasARN, or alias name for the KMS key. To use a KMS key in a different AmazonWeb Services account, specify the key ARN or alias ARN.If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanagerKMS key is used to encrypt the secret. If the secret is in a different AmazonWeb Services account, then you can't use the aws/secretsmanager KMS key toencrypt the secret, and you must use a customer managed KMS key.There is a default KMS key for your Amazon Web Services account. Your AmazonWeb Services account has a different default KMS key for each Amazon WebServices Region.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier to encrypt a secret that is automaticallygenerated and managed in Amazon Web Services Secrets Manager.This setting is valid only if the master user password is managed by RDSin Amazon Web Services Secrets Manager for the DB instance.The Amazon Web Services KMS key identifier is the key ARN, key ID, aliasARN, or alias name for the KMS key. To use a KMS key in a different AmazonWeb Services account, specify the key ARN or alias ARN.If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanagerKMS key is used to encrypt the secret. If the secret is in a different AmazonWeb Services account, then you can't use the aws/secretsmanager KMS key toencrypt the secret, and you must use a customer managed KMS key.There is a default KMS key for your Amazon Web Services account. Your AmazonWeb Services account has a different default KMS key for each Amazon WebServices Region.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"master_user_secret_kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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
						Description:         "The name for the master user.Amazon AuroraNot applicable. The name for the master user is managed by the DB cluster.Amazon RDSConstraints:   * Required.   * Must be 1 to 16 letters, numbers, or underscores.   * First character must be a letter.   * Can't be a reserved word for the chosen database engine.",
						MarkdownDescription: "The name for the master user.Amazon AuroraNot applicable. The name for the master user is managed by the DB cluster.Amazon RDSConstraints:   * Required.   * Must be 1 to 16 letters, numbers, or underscores.   * First character must be a letter.   * Can't be a reserved word for the chosen database engine.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_allocated_storage": schema.Int64Attribute{
						Description:         "The upper limit in gibibytes (GiB) to which Amazon RDS can automaticallyscale the storage of the DB instance.For more information about this setting, including limitations that applyto it, see Managing capacity automatically with Amazon RDS storage autoscaling(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "The upper limit in gibibytes (GiB) to which Amazon RDS can automaticallyscale the storage of the DB instance.For more information about this setting, including limitations that applyto it, see Managing capacity automatically with Amazon RDS storage autoscaling(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. Storage is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitoring_interval": schema.Int64Attribute{
						Description:         "The interval, in seconds, between points when Enhanced Monitoring metricsare collected for the DB instance. To disable collection of Enhanced Monitoringmetrics, specify 0. The default is 0.If MonitoringRoleArn is specified, then you must set MonitoringInterval toa value other than 0.This setting doesn't apply to RDS Custom.Valid Values: 0, 1, 5, 10, 15, 30, 60",
						MarkdownDescription: "The interval, in seconds, between points when Enhanced Monitoring metricsare collected for the DB instance. To disable collection of Enhanced Monitoringmetrics, specify 0. The default is 0.If MonitoringRoleArn is specified, then you must set MonitoringInterval toa value other than 0.This setting doesn't apply to RDS Custom.Valid Values: 0, 1, 5, 10, 15, 30, 60",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitoring_role_arn": schema.StringAttribute{
						Description:         "The ARN for the IAM role that permits RDS to send enhanced monitoring metricsto Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess.For information on creating a monitoring role, see Setting Up and EnablingEnhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling)in the Amazon RDS User Guide.If MonitoringInterval is set to a value other than 0, then you must supplya MonitoringRoleArn value.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The ARN for the IAM role that permits RDS to send enhanced monitoring metricsto Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess.For information on creating a monitoring role, see Setting Up and EnablingEnhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling)in the Amazon RDS User Guide.If MonitoringInterval is set to a value other than 0, then you must supplya MonitoringRoleArn value.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"multi_az": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB instance is a Multi-AZ deployment.You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZdeployment.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. DB instance Availability Zones (AZs) are managed by the DBcluster.",
						MarkdownDescription: "A value that indicates whether the DB instance is a Multi-AZ deployment.You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZdeployment.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable. DB instance Availability Zones (AZs) are managed by the DBcluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nchar_character_set_name": schema.StringAttribute{
						Description:         "The name of the NCHAR character set for the Oracle DB instance.This parameter doesn't apply to RDS Custom.",
						MarkdownDescription: "The name of the NCHAR character set for the Oracle DB instance.This parameter doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network_type": schema.StringAttribute{
						Description:         "The network type of the DB instance.Valid values:   * IPV4   * DUALThe network type is determined by the DBSubnetGroup specified for the DBinstance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4and the IPv6 protocols (DUAL).For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)in the Amazon RDS User Guide.",
						MarkdownDescription: "The network type of the DB instance.Valid values:   * IPV4   * DUALThe network type is determined by the DBSubnetGroup specified for the DBinstance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4and the IPv6 protocols (DUAL).For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html)in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"option_group_name": schema.StringAttribute{
						Description:         "A value that indicates that the DB instance should be associated with thespecified option group.Permanent options, such as the TDE option for Oracle Advanced Security TDE,can't be removed from an option group. Also, that option group can't be removedfrom a DB instance after it is associated with a DB instance.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
						MarkdownDescription: "A value that indicates that the DB instance should be associated with thespecified option group.Permanent options, such as the TDE option for Oracle Advanced Security TDE,can't be removed from an option group. Also, that option group can't be removedfrom a DB instance after it is associated with a DB instance.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_enabled": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable Performance Insights for the DBinstance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "A value that indicates whether to enable Performance Insights for the DBinstance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html)in the Amazon RDS User Guide.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for encryption of PerformanceInsights data.The Amazon Web Services KMS key identifier is the key ARN, key ID, aliasARN, or alias name for the KMS key.If you do not specify a value for PerformanceInsightsKMSKeyId, then AmazonRDS uses your default KMS key. There is a default KMS key for your AmazonWeb Services account. Your Amazon Web Services account has a different defaultKMS key for each Amazon Web Services Region.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for encryption of PerformanceInsights data.The Amazon Web Services KMS key identifier is the key ARN, key ID, aliasARN, or alias name for the KMS key.If you do not specify a value for PerformanceInsightsKMSKeyId, then AmazonRDS uses your default KMS key. There is a default KMS key for your AmazonWeb Services account. Your Amazon Web Services account has a different defaultKMS key for each Amazon Web Services Region.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_retention_period": schema.Int64Attribute{
						Description:         "The number of days to retain Performance Insights data. The default is 7days. The following values are valid:   * 7   * month * 31, where month is a number of months from 1-23   * 731For example, the following values are valid:   * 93 (3 months * 31)   * 341 (11 months * 31)   * 589 (19 months * 31)   * 731If you specify a retention period such as 94, which isn't a valid value,RDS issues an error.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The number of days to retain Performance Insights data. The default is 7days. The following values are valid:   * 7   * month * 31, where month is a number of months from 1-23   * 731For example, the following values are valid:   * 93 (3 months * 31)   * 341 (11 months * 31)   * 589 (19 months * 31)   * 731If you specify a retention period such as 94, which isn't a valid value,RDS issues an error.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"port": schema.Int64Attribute{
						Description:         "The port number on which the database accepts connections.MySQLDefault: 3306Valid values: 1150-65535Type: IntegerMariaDBDefault: 3306Valid values: 1150-65535Type: IntegerPostgreSQLDefault: 5432Valid values: 1150-65535Type: IntegerOracleDefault: 1521Valid values: 1150-65535SQL ServerDefault: 1433Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and49152-49156.Amazon AuroraDefault: 3306Valid values: 1150-65535Type: Integer",
						MarkdownDescription: "The port number on which the database accepts connections.MySQLDefault: 3306Valid values: 1150-65535Type: IntegerMariaDBDefault: 3306Valid values: 1150-65535Type: IntegerPostgreSQLDefault: 5432Valid values: 1150-65535Type: IntegerOracleDefault: 1521Valid values: 1150-65535SQL ServerDefault: 1433Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and49152-49156.Amazon AuroraDefault: 3306Valid values: 1150-65535Type: Integer",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pre_signed_url": schema.StringAttribute{
						Description:         "When you are creating a read replica from one Amazon Web Services GovCloud(US) Region to another or from one China Amazon Web Services Region to another,the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplicaAPI operation in the source Amazon Web Services Region that contains thesource DB instance.This setting applies only to Amazon Web Services GovCloud (US) Regions andChina Amazon Web Services Regions. It's ignored in other Amazon Web ServicesRegions.This setting applies only when replicating from a source DB instance. SourceDB clusters aren't supported in Amazon Web Services GovCloud (US) Regionsand China Amazon Web Services Regions.You must specify this parameter when you create an encrypted read replicafrom another Amazon Web Services Region by using the Amazon RDS API. Don'tspecify PreSignedUrl when you are creating an encrypted read replica in thesame Amazon Web Services Region.The presigned URL must be a valid request for the CreateDBInstanceReadReplicaAPI operation that can run in the source Amazon Web Services Region thatcontains the encrypted source DB instance. The presigned URL request mustcontain the following parameter values:   * DestinationRegion - The Amazon Web Services Region that the encrypted   read replica is created in. This Amazon Web Services Region is the same   one where the CreateDBInstanceReadReplica operation is called that contains   this presigned URL. For example, if you create an encrypted DB instance   in the us-west-1 Amazon Web Services Region, from a source DB instance   in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica   operation in the us-east-1 Amazon Web Services Region and provide a presigned   URL that contains a call to the CreateDBInstanceReadReplica operation   in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion   in the presigned URL must be set to the us-east-1 Amazon Web Services   Region.   * KmsKeyId - The KMS key identifier for the key to use to encrypt the   read replica in the destination Amazon Web Services Region. This is the   same identifier for both the CreateDBInstanceReadReplica operation that   is called in the destination Amazon Web Services Region, and the operation   contained in the presigned URL.   * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted   DB instance to be replicated. This identifier must be in the Amazon Resource   Name (ARN) format for the source Amazon Web Services Region. For example,   if you are creating an encrypted read replica from a DB instance in the   us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier   looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115.To learn how to generate a Signature Version 4 signed request, see AuthenticatingRequests: Using Query Parameters (Amazon Web Services Signature Version 4)(https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).If you are using an Amazon Web Services SDK tool or the CLI, you can specifySourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrlmanually. Specifying SourceRegion autogenerates a presigned URL that is avalid request for the operation that can run in the source Amazon Web ServicesRegion.SourceRegion isn't supported for SQL Server, because Amazon RDS for SQL Serverdoesn't support cross-Region read replicas.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "When you are creating a read replica from one Amazon Web Services GovCloud(US) Region to another or from one China Amazon Web Services Region to another,the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplicaAPI operation in the source Amazon Web Services Region that contains thesource DB instance.This setting applies only to Amazon Web Services GovCloud (US) Regions andChina Amazon Web Services Regions. It's ignored in other Amazon Web ServicesRegions.This setting applies only when replicating from a source DB instance. SourceDB clusters aren't supported in Amazon Web Services GovCloud (US) Regionsand China Amazon Web Services Regions.You must specify this parameter when you create an encrypted read replicafrom another Amazon Web Services Region by using the Amazon RDS API. Don'tspecify PreSignedUrl when you are creating an encrypted read replica in thesame Amazon Web Services Region.The presigned URL must be a valid request for the CreateDBInstanceReadReplicaAPI operation that can run in the source Amazon Web Services Region thatcontains the encrypted source DB instance. The presigned URL request mustcontain the following parameter values:   * DestinationRegion - The Amazon Web Services Region that the encrypted   read replica is created in. This Amazon Web Services Region is the same   one where the CreateDBInstanceReadReplica operation is called that contains   this presigned URL. For example, if you create an encrypted DB instance   in the us-west-1 Amazon Web Services Region, from a source DB instance   in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica   operation in the us-east-1 Amazon Web Services Region and provide a presigned   URL that contains a call to the CreateDBInstanceReadReplica operation   in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion   in the presigned URL must be set to the us-east-1 Amazon Web Services   Region.   * KmsKeyId - The KMS key identifier for the key to use to encrypt the   read replica in the destination Amazon Web Services Region. This is the   same identifier for both the CreateDBInstanceReadReplica operation that   is called in the destination Amazon Web Services Region, and the operation   contained in the presigned URL.   * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted   DB instance to be replicated. This identifier must be in the Amazon Resource   Name (ARN) format for the source Amazon Web Services Region. For example,   if you are creating an encrypted read replica from a DB instance in the   us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier   looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115.To learn how to generate a Signature Version 4 signed request, see AuthenticatingRequests: Using Query Parameters (Amazon Web Services Signature Version 4)(https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html)and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).If you are using an Amazon Web Services SDK tool or the CLI, you can specifySourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrlmanually. Specifying SourceRegion autogenerates a presigned URL that is avalid request for the operation that can run in the source Amazon Web ServicesRegion.SourceRegion isn't supported for SQL Server, because Amazon RDS for SQL Serverdoesn't support cross-Region read replicas.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_backup_window": schema.StringAttribute{
						Description:         "The daily time range during which automated backups are created if automatedbackups are enabled, using the BackupRetentionPeriod parameter. The defaultis a 30-minute window selected at random from an 8-hour block of time foreach Amazon Web Services Region. For more information, see Backup window(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow)in the Amazon RDS User Guide.Amazon AuroraNot applicable. The daily time range for creating automated backups is managedby the DB cluster.Constraints:   * Must be in the format hh24:mi-hh24:mi.   * Must be in Universal Coordinated Time (UTC).   * Must not conflict with the preferred maintenance window.   * Must be at least 30 minutes.",
						MarkdownDescription: "The daily time range during which automated backups are created if automatedbackups are enabled, using the BackupRetentionPeriod parameter. The defaultis a 30-minute window selected at random from an 8-hour block of time foreach Amazon Web Services Region. For more information, see Backup window(https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow)in the Amazon RDS User Guide.Amazon AuroraNot applicable. The daily time range for creating automated backups is managedby the DB cluster.Constraints:   * Must be in the format hh24:mi-hh24:mi.   * Must be in Universal Coordinated Time (UTC).   * Must not conflict with the preferred maintenance window.   * Must be at least 30 minutes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_maintenance_window": schema.StringAttribute{
						Description:         "The time range each week during which system maintenance can occur, in UniversalCoordinated Time (UTC). For more information, see Amazon RDS MaintenanceWindow (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.Constraints: Minimum 30-minute window.",
						MarkdownDescription: "The time range each week during which system maintenance can occur, in UniversalCoordinated Time (UTC). For more information, see Amazon RDS MaintenanceWindow (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.Constraints: Minimum 30-minute window.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"processor_features": schema.ListNestedAttribute{
						Description:         "The number of CPU cores and the number of threads per core for the DB instanceclass of the DB instance.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
						MarkdownDescription: "The number of CPU cores and the number of threads per core for the DB instanceclass of the DB instance.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
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
						Description:         "A value that specifies the order in which an Aurora Replica is promoted tothe primary instance after a failure of the existing primary instance. Formore information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance)in the Amazon Aurora User Guide.This setting doesn't apply to RDS Custom.Default: 1Valid Values: 0 - 15",
						MarkdownDescription: "A value that specifies the order in which an Aurora Replica is promoted tothe primary instance after a failure of the existing primary instance. Formore information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance)in the Amazon Aurora User Guide.This setting doesn't apply to RDS Custom.Default: 1Valid Values: 0 - 15",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"publicly_accessible": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB instance is publicly accessible.When the DB instance is publicly accessible, its Domain Name System (DNS)endpoint resolves to the private IP address from within the DB instance'svirtual private cloud (VPC). It resolves to the public IP address from outsideof the DB instance's VPC. Access to the DB instance is ultimately controlledby the security group it uses. That public access is not permitted if thesecurity group assigned to the DB instance doesn't permit it.When the DB instance isn't publicly accessible, it is an internal DB instancewith a DNS name that resolves to a private IP address.Default: The default behavior varies depending on whether DBSubnetGroupNameis specified.If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified,the following applies:   * If the default VPC in the target Region doesn’t have an internet gateway   attached to it, the DB instance is private.   * If the default VPC in the target Region has an internet gateway attached   to it, the DB instance is public.If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified,the following applies:   * If the subnets are part of a VPC that doesn’t have an internet gateway   attached to it, the DB instance is private.   * If the subnets are part of a VPC that has an internet gateway attached   to it, the DB instance is public.",
						MarkdownDescription: "A value that indicates whether the DB instance is publicly accessible.When the DB instance is publicly accessible, its Domain Name System (DNS)endpoint resolves to the private IP address from within the DB instance'svirtual private cloud (VPC). It resolves to the public IP address from outsideof the DB instance's VPC. Access to the DB instance is ultimately controlledby the security group it uses. That public access is not permitted if thesecurity group assigned to the DB instance doesn't permit it.When the DB instance isn't publicly accessible, it is an internal DB instancewith a DNS name that resolves to a private IP address.Default: The default behavior varies depending on whether DBSubnetGroupNameis specified.If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified,the following applies:   * If the default VPC in the target Region doesn’t have an internet gateway   attached to it, the DB instance is private.   * If the default VPC in the target Region has an internet gateway attached   to it, the DB instance is public.If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified,the following applies:   * If the subnets are part of a VPC that doesn’t have an internet gateway   attached to it, the DB instance is private.   * If the subnets are part of a VPC that has an internet gateway attached   to it, the DB instance is public.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_mode": schema.StringAttribute{
						Description:         "The open mode of the replica database: mounted or read-only.This parameter is only supported for Oracle DB instances.Mounted DB replicas are included in Oracle Database Enterprise Edition. Themain use case for mounted replicas is cross-Region disaster recovery. Theprimary database doesn't use Active Data Guard to transmit information tothe mounted replica. Because it doesn't accept user connections, a mountedreplica can't serve a read-only workload.You can create a combination of mounted and read-only DB replicas for thesame primary DB instance. For more information, see Working with Oracle ReadReplicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html)in the Amazon RDS User Guide.For RDS Custom, you must specify this parameter and set it to mounted. Thevalue won't be set by default. After replica creation, you can manage theopen mode manually.",
						MarkdownDescription: "The open mode of the replica database: mounted or read-only.This parameter is only supported for Oracle DB instances.Mounted DB replicas are included in Oracle Database Enterprise Edition. Themain use case for mounted replicas is cross-Region disaster recovery. Theprimary database doesn't use Active Data Guard to transmit information tothe mounted replica. Because it doesn't accept user connections, a mountedreplica can't serve a read-only workload.You can create a combination of mounted and read-only DB replicas for thesame primary DB instance. For more information, see Working with Oracle ReadReplicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html)in the Amazon RDS User Guide.For RDS Custom, you must specify this parameter and set it to mounted. Thevalue won't be set by default. After replica creation, you can manage theopen mode manually.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_db_instance_identifier": schema.StringAttribute{
						Description:         "The identifier of the DB instance that will act as the source for the readreplica. Each DB instance can have up to 15 read replicas, with the exceptionof Oracle and SQL Server, which can have up to five.Constraints:   * Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL,   or SQL Server DB instance.   * Can't be specified if the SourceDBClusterIdentifier parameter is also   specified.   * For the limitations of Oracle read replicas, see Version and licensing   considerations for RDS for Oracle replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.limitations.html#oracle-read-replicas.limitations.versions-and-licenses)   in the Amazon RDS User Guide.   * For the limitations of SQL Server read replicas, see Read replica limitations   with SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.html#SQLServer.ReadReplicas.Limitations)   in the Amazon RDS User Guide.   * The specified DB instance must have automatic backups enabled, that   is, its backup retention period must be greater than 0.   * If the source DB instance is in the same Amazon Web Services Region   as the read replica, specify a valid DB instance identifier.   * If the source DB instance is in a different Amazon Web Services Region   from the read replica, specify a valid DB instance ARN. For more information,   see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)   in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS   Custom, which don't support cross-Region replicas.",
						MarkdownDescription: "The identifier of the DB instance that will act as the source for the readreplica. Each DB instance can have up to 15 read replicas, with the exceptionof Oracle and SQL Server, which can have up to five.Constraints:   * Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL,   or SQL Server DB instance.   * Can't be specified if the SourceDBClusterIdentifier parameter is also   specified.   * For the limitations of Oracle read replicas, see Version and licensing   considerations for RDS for Oracle replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.limitations.html#oracle-read-replicas.limitations.versions-and-licenses)   in the Amazon RDS User Guide.   * For the limitations of SQL Server read replicas, see Read replica limitations   with SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.html#SQLServer.ReadReplicas.Limitations)   in the Amazon RDS User Guide.   * The specified DB instance must have automatic backups enabled, that   is, its backup retention period must be greater than 0.   * If the source DB instance is in the same Amazon Web Services Region   as the read replica, specify a valid DB instance identifier.   * If the source DB instance is in a different Amazon Web Services Region   from the read replica, specify a valid DB instance ARN. For more information,   see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)   in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS   Custom, which don't support cross-Region replicas.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_region": schema.StringAttribute{
						Description:         "SourceRegion is the source region where the resource exists. This is notsent over the wire and is only used for presigning. This value should alwayshave the same region as the source ARN.",
						MarkdownDescription: "SourceRegion is the source region where the resource exists. This is notsent over the wire and is only used for presigning. This value should alwayshave the same region as the source ARN.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_encrypted": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB instance is encrypted. By default,it isn't encrypted.For RDS Custom instances, either set this parameter to true or leave it unset.If you set this parameter to false, RDS reports an error.Amazon AuroraNot applicable. The encryption for DB instances is managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether the DB instance is encrypted. By default,it isn't encrypted.For RDS Custom instances, either set this parameter to true or leave it unset.If you set this parameter to false, RDS reports an error.Amazon AuroraNot applicable. The encryption for DB instances is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_throughput": schema.Int64Attribute{
						Description:         "Specifies the storage throughput value for the DB instance.This setting applies only to the gp3 storage type.This setting doesn't apply to RDS Custom or Amazon Aurora.",
						MarkdownDescription: "Specifies the storage throughput value for the DB instance.This setting applies only to the gp3 storage type.This setting doesn't apply to RDS Custom or Amazon Aurora.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_type": schema.StringAttribute{
						Description:         "Specifies the storage type to be associated with the DB instance.Valid values: gp2 | gp3 | io1 | standardIf you specify io1 or gp3, you must also include a value for the Iops parameter.Default: io1 if the Iops parameter is specified, otherwise gp2Amazon AuroraNot applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "Specifies the storage type to be associated with the DB instance.Valid values: gp2 | gp3 | io1 | standardIf you specify io1 or gp3, you must also include a value for the Iops parameter.Default: io1 if the Iops parameter is specified, otherwise gp2Amazon AuroraNot applicable. Storage is managed by the DB cluster.",
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
						Description:         "The ARN from the key store with which to associate the instance for TDE encryption.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
						MarkdownDescription: "The ARN from the key store with which to associate the instance for TDE encryption.This setting doesn't apply to RDS Custom.Amazon AuroraNot applicable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tde_credential_password": schema.StringAttribute{
						Description:         "The password for the given ARN from the key store in order to access thedevice.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The password for the given ARN from the key store in order to access thedevice.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timezone": schema.StringAttribute{
						Description:         "The time zone of the DB instance. The time zone parameter is currently supportedonly by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
						MarkdownDescription: "The time zone of the DB instance. The time zone parameter is currently supportedonly by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_default_processor_features": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB instance class of the DB instance usesits default processor features.This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "A value that indicates whether the DB instance class of the DB instance usesits default processor features.This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vpc_security_group_i_ds": schema.ListAttribute{
						Description:         "A list of Amazon EC2 VPC security groups to associate with this DB instance.Amazon AuroraNot applicable. The associated list of EC2 VPC security groups is managedby the DB cluster.Default: The default EC2 VPC security group for the DB subnet group's VPC.",
						MarkdownDescription: "A list of Amazon EC2 VPC security groups to associate with this DB instance.Amazon AuroraNot applicable. The associated list of EC2 VPC security groups is managedby the DB cluster.Default: The default EC2 VPC security group for the DB subnet group's VPC.",
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
									Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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
