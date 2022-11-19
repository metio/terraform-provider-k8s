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

type RdsServicesK8SAwsDBInstanceV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*RdsServicesK8SAwsDBInstanceV1Alpha1Resource)(nil)
)

type RdsServicesK8SAwsDBInstanceV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type RdsServicesK8SAwsDBInstanceV1Alpha1GoModel struct {
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
		AllocatedStorage *int64 `tfsdk:"allocated_storage" yaml:"allocatedStorage,omitempty"`

		AutoMinorVersionUpgrade *bool `tfsdk:"auto_minor_version_upgrade" yaml:"autoMinorVersionUpgrade,omitempty"`

		AvailabilityZone *string `tfsdk:"availability_zone" yaml:"availabilityZone,omitempty"`

		BackupRetentionPeriod *int64 `tfsdk:"backup_retention_period" yaml:"backupRetentionPeriod,omitempty"`

		BackupTarget *string `tfsdk:"backup_target" yaml:"backupTarget,omitempty"`

		CharacterSetName *string `tfsdk:"character_set_name" yaml:"characterSetName,omitempty"`

		CopyTagsToSnapshot *bool `tfsdk:"copy_tags_to_snapshot" yaml:"copyTagsToSnapshot,omitempty"`

		CustomIAMInstanceProfile *string `tfsdk:"custom_iam_instance_profile" yaml:"customIAMInstanceProfile,omitempty"`

		DbClusterIdentifier *string `tfsdk:"db_cluster_identifier" yaml:"dbClusterIdentifier,omitempty"`

		DbInstanceClass *string `tfsdk:"db_instance_class" yaml:"dbInstanceClass,omitempty"`

		DbInstanceIdentifier *string `tfsdk:"db_instance_identifier" yaml:"dbInstanceIdentifier,omitempty"`

		DbName *string `tfsdk:"db_name" yaml:"dbName,omitempty"`

		DbParameterGroupName *string `tfsdk:"db_parameter_group_name" yaml:"dbParameterGroupName,omitempty"`

		DbParameterGroupRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"db_parameter_group_ref" yaml:"dbParameterGroupRef,omitempty"`

		DbSnapshotIdentifier *string `tfsdk:"db_snapshot_identifier" yaml:"dbSnapshotIdentifier,omitempty"`

		DbSubnetGroupName *string `tfsdk:"db_subnet_group_name" yaml:"dbSubnetGroupName,omitempty"`

		DbSubnetGroupRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"db_subnet_group_ref" yaml:"dbSubnetGroupRef,omitempty"`

		DeletionProtection *bool `tfsdk:"deletion_protection" yaml:"deletionProtection,omitempty"`

		DestinationRegion *string `tfsdk:"destination_region" yaml:"destinationRegion,omitempty"`

		Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

		DomainIAMRoleName *string `tfsdk:"domain_iam_role_name" yaml:"domainIAMRoleName,omitempty"`

		EnableCloudwatchLogsExports *[]string `tfsdk:"enable_cloudwatch_logs_exports" yaml:"enableCloudwatchLogsExports,omitempty"`

		EnableCustomerOwnedIP *bool `tfsdk:"enable_customer_owned_ip" yaml:"enableCustomerOwnedIP,omitempty"`

		EnableIAMDatabaseAuthentication *bool `tfsdk:"enable_iam_database_authentication" yaml:"enableIAMDatabaseAuthentication,omitempty"`

		Engine *string `tfsdk:"engine" yaml:"engine,omitempty"`

		EngineVersion *string `tfsdk:"engine_version" yaml:"engineVersion,omitempty"`

		Iops *int64 `tfsdk:"iops" yaml:"iops,omitempty"`

		KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

		KmsKeyRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"kms_key_ref" yaml:"kmsKeyRef,omitempty"`

		LicenseModel *string `tfsdk:"license_model" yaml:"licenseModel,omitempty"`

		MasterUserPassword *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"master_user_password" yaml:"masterUserPassword,omitempty"`

		MasterUsername *string `tfsdk:"master_username" yaml:"masterUsername,omitempty"`

		MaxAllocatedStorage *int64 `tfsdk:"max_allocated_storage" yaml:"maxAllocatedStorage,omitempty"`

		MonitoringInterval *int64 `tfsdk:"monitoring_interval" yaml:"monitoringInterval,omitempty"`

		MonitoringRoleARN *string `tfsdk:"monitoring_role_arn" yaml:"monitoringRoleARN,omitempty"`

		MultiAZ *bool `tfsdk:"multi_az" yaml:"multiAZ,omitempty"`

		NcharCharacterSetName *string `tfsdk:"nchar_character_set_name" yaml:"ncharCharacterSetName,omitempty"`

		NetworkType *string `tfsdk:"network_type" yaml:"networkType,omitempty"`

		OptionGroupName *string `tfsdk:"option_group_name" yaml:"optionGroupName,omitempty"`

		PerformanceInsightsEnabled *bool `tfsdk:"performance_insights_enabled" yaml:"performanceInsightsEnabled,omitempty"`

		PerformanceInsightsKMSKeyID *string `tfsdk:"performance_insights_kms_key_id" yaml:"performanceInsightsKMSKeyID,omitempty"`

		PerformanceInsightsRetentionPeriod *int64 `tfsdk:"performance_insights_retention_period" yaml:"performanceInsightsRetentionPeriod,omitempty"`

		Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

		PreSignedURL *string `tfsdk:"pre_signed_url" yaml:"preSignedURL,omitempty"`

		PreferredBackupWindow *string `tfsdk:"preferred_backup_window" yaml:"preferredBackupWindow,omitempty"`

		PreferredMaintenanceWindow *string `tfsdk:"preferred_maintenance_window" yaml:"preferredMaintenanceWindow,omitempty"`

		ProcessorFeatures *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"processor_features" yaml:"processorFeatures,omitempty"`

		PromotionTier *int64 `tfsdk:"promotion_tier" yaml:"promotionTier,omitempty"`

		PubliclyAccessible *bool `tfsdk:"publicly_accessible" yaml:"publiclyAccessible,omitempty"`

		ReplicaMode *string `tfsdk:"replica_mode" yaml:"replicaMode,omitempty"`

		SourceDBInstanceIdentifier *string `tfsdk:"source_db_instance_identifier" yaml:"sourceDBInstanceIdentifier,omitempty"`

		SourceRegion *string `tfsdk:"source_region" yaml:"sourceRegion,omitempty"`

		StorageEncrypted *bool `tfsdk:"storage_encrypted" yaml:"storageEncrypted,omitempty"`

		StorageType *string `tfsdk:"storage_type" yaml:"storageType,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		TdeCredentialARN *string `tfsdk:"tde_credential_arn" yaml:"tdeCredentialARN,omitempty"`

		TdeCredentialPassword *string `tfsdk:"tde_credential_password" yaml:"tdeCredentialPassword,omitempty"`

		Timezone *string `tfsdk:"timezone" yaml:"timezone,omitempty"`

		UseDefaultProcessorFeatures *bool `tfsdk:"use_default_processor_features" yaml:"useDefaultProcessorFeatures,omitempty"`

		VpcSecurityGroupIDs *[]string `tfsdk:"vpc_security_group_i_ds" yaml:"vpcSecurityGroupIDs,omitempty"`

		VpcSecurityGroupRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"vpc_security_group_refs" yaml:"vpcSecurityGroupRefs,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewRdsServicesK8SAwsDBInstanceV1Alpha1Resource() resource.Resource {
	return &RdsServicesK8SAwsDBInstanceV1Alpha1Resource{}
}

func (r *RdsServicesK8SAwsDBInstanceV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rds_services_k8s_aws_db_instance_v1alpha1"
}

func (r *RdsServicesK8SAwsDBInstanceV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "DBInstance is the Schema for the DBInstances API",
		MarkdownDescription: "DBInstance is the Schema for the DBInstances API",
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
				Description:         "DBInstanceSpec defines the desired state of DBInstance.  Contains the details of an Amazon RDS DB instance.  This data type is used as a response element in the operations CreateDBInstance, CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance, PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3, RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",
				MarkdownDescription: "DBInstanceSpec defines the desired state of DBInstance.  Contains the details of an Amazon RDS DB instance.  This data type is used as a response element in the operations CreateDBInstance, CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance, PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3, RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"allocated_storage": {
						Description:         "The amount of storage in gibibytes (GiB) to allocate for the DB instance.  Type: Integer  Amazon Aurora  Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.  Amazon RDS Custom  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  * Provisioned IOPS storage (io1): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  MySQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  MariaDB  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  PostgreSQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  Oracle  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 10 to 3072.  SQL Server  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Enterprise and Standard editions: Must be an integer from 20 to 16384. Web and Express editions: Must be an integer from 20 to 16384.  * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must be an integer from 100 to 16384. Web and Express editions: Must be an integer from 100 to 16384.  * Magnetic storage (standard): Enterprise and Standard editions: Must be an integer from 20 to 1024. Web and Express editions: Must be an integer from 20 to 1024.",
						MarkdownDescription: "The amount of storage in gibibytes (GiB) to allocate for the DB instance.  Type: Integer  Amazon Aurora  Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.  Amazon RDS Custom  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  * Provisioned IOPS storage (io1): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  MySQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  MariaDB  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  PostgreSQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  Oracle  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 10 to 3072.  SQL Server  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2): Enterprise and Standard editions: Must be an integer from 20 to 16384. Web and Express editions: Must be an integer from 20 to 16384.  * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must be an integer from 100 to 16384. Web and Express editions: Must be an integer from 100 to 16384.  * Magnetic storage (standard): Enterprise and Standard editions: Must be an integer from 20 to 1024. Web and Express editions: Must be an integer from 20 to 1024.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auto_minor_version_upgrade": {
						Description:         "A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.  If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgrade to false.",
						MarkdownDescription: "A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.  If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgrade to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"availability_zone": {
						Description:         "The Availability Zone (AZ) where the database will be created. For information on Amazon Web Services Regions and Availability Zones, see Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).  Amazon Aurora  Each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.  Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web Services Region.  Example: us-east-1d  Constraint: The AvailabilityZone parameter can't be specified if the DB instance is a Multi-AZ deployment. The specified Availability Zone must be in the same Amazon Web Services Region as the current endpoint.",
						MarkdownDescription: "The Availability Zone (AZ) where the database will be created. For information on Amazon Web Services Regions and Availability Zones, see Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).  Amazon Aurora  Each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.  Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web Services Region.  Example: us-east-1d  Constraint: The AvailabilityZone parameter can't be specified if the DB instance is a Multi-AZ deployment. The specified Availability Zone must be in the same Amazon Web Services Region as the current endpoint.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup_retention_period": {
						Description:         "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.  Amazon Aurora  Not applicable. The retention period for automated backups is managed by the DB cluster.  Default: 1  Constraints:  * Must be a value from 0 to 35  * Can't be set to 0 if the DB instance is a source to read replicas  * Can't be set to 0 for an RDS Custom for Oracle DB instance",
						MarkdownDescription: "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.  Amazon Aurora  Not applicable. The retention period for automated backups is managed by the DB cluster.  Default: 1  Constraints:  * Must be a value from 0 to 35  * Can't be set to 0 if the DB instance is a source to read replicas  * Can't be set to 0 for an RDS Custom for Oracle DB instance",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup_target": {
						Description:         "Specifies where automated backups and manual snapshots are stored.  Possible values are outposts (Amazon Web Services Outposts) and region (Amazon Web Services Region). The default is region.  For more information, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.",
						MarkdownDescription: "Specifies where automated backups and manual snapshots are stored.  Possible values are outposts (Amazon Web Services Outposts) and region (Amazon Web Services Region). The default is region.  For more information, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"character_set_name": {
						Description:         "For supported engines, this value indicates that the DB instance should be associated with the specified CharacterSet.  This setting doesn't apply to RDS Custom. However, if you need to change the character set, you can change it on the database itself.  Amazon Aurora  Not applicable. The character set is managed by the DB cluster. For more information, see CreateDBCluster.",
						MarkdownDescription: "For supported engines, this value indicates that the DB instance should be associated with the specified CharacterSet.  This setting doesn't apply to RDS Custom. However, if you need to change the character set, you can change it on the database itself.  Amazon Aurora  Not applicable. The character set is managed by the DB cluster. For more information, see CreateDBCluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"copy_tags_to_snapshot": {
						Description:         "A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.  Amazon Aurora  Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.",
						MarkdownDescription: "A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.  Amazon Aurora  Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_iam_instance_profile": {
						Description:         "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance. The instance profile must meet the following requirements:  * The profile must exist in your account.  * The profile must have an IAM role that Amazon EC2 has permissions to assume.  * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom.  For the list of permissions required for the IAM role, see Configure IAM and your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the Amazon RDS User Guide.  This setting is required for RDS Custom.",
						MarkdownDescription: "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance. The instance profile must meet the following requirements:  * The profile must exist in your account.  * The profile must have an IAM role that Amazon EC2 has permissions to assume.  * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom.  For the list of permissions required for the IAM role, see Configure IAM and your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the Amazon RDS User Guide.  This setting is required for RDS Custom.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"db_cluster_identifier": {
						Description:         "The identifier of the DB cluster that the instance will belong to.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The identifier of the DB cluster that the instance will belong to.  This setting doesn't apply to RDS Custom.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"db_instance_class": {
						Description:         "The compute and memory capacity of the DB instance, for example db.m5.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines. For the full list of DB instance classes, and availability for your engine, see DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide or Aurora DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html) in the Amazon Aurora User Guide.",
						MarkdownDescription: "The compute and memory capacity of the DB instance, for example db.m5.large. Not all DB instance classes are available in all Amazon Web Services Regions, or for all database engines. For the full list of DB instance classes, and availability for your engine, see DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.DBInstanceClass.html) in the Amazon RDS User Guide or Aurora DB instance classes (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.DBInstanceClass.html) in the Amazon Aurora User Guide.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"db_instance_identifier": {
						Description:         "The DB instance identifier. This parameter is stored as a lowercase string.  Constraints:  * Must contain from 1 to 63 letters, numbers, or hyphens.  * First character must be a letter.  * Can't end with a hyphen or contain two consecutive hyphens.  Example: mydbinstance",
						MarkdownDescription: "The DB instance identifier. This parameter is stored as a lowercase string.  Constraints:  * Must contain from 1 to 63 letters, numbers, or hyphens.  * First character must be a letter.  * Can't end with a hyphen or contain two consecutive hyphens.  Example: mydbinstance",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"db_name": {
						Description:         "The meaning of this parameter differs according to the database engine you use.  MySQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  MariaDB  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  PostgreSQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, a database named postgres is created in the DB instance.  Constraints:  * Must contain 1 to 63 letters, numbers, or underscores.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  Oracle  The Oracle System ID (SID) of the created DB instance. If you specify null, the default value ORCL is used. You can't specify the string NULL, or any other reserved word, for DBName.  Default: ORCL  Constraints:  * Can't be longer than 8 characters  Amazon RDS Custom for Oracle  The Oracle System ID (SID) of the created RDS Custom DB instance. If you don't specify a value, the default value is ORCL.  Default: ORCL  Constraints:  * It must contain 1 to 8 alphanumeric characters.  * It must contain a letter.  * It can't be a word reserved by the database engine.  Amazon RDS Custom for SQL Server  Not applicable. Must be null.  SQL Server  Not applicable. Must be null.  Amazon Aurora MySQL  The name of the database to create when the primary DB instance of the Aurora MySQL DB cluster is created. If this parameter isn't specified for an Aurora MySQL DB cluster, no database is created in the DB cluster.  Constraints:  * It must contain 1 to 64 alphanumeric characters.  * It can't be a word reserved by the database engine.  Amazon Aurora PostgreSQL  The name of the database to create when the primary DB instance of the Aurora PostgreSQL DB cluster is created. If this parameter isn't specified for an Aurora PostgreSQL DB cluster, a database named postgres is created in the DB cluster.  Constraints:  * It must contain 1 to 63 alphanumeric characters.  * It must begin with a letter or an underscore. Subsequent characters can be letters, underscores, or digits (0 to 9).  * It can't be a word reserved by the database engine.",
						MarkdownDescription: "The meaning of this parameter differs according to the database engine you use.  MySQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  MariaDB  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  PostgreSQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, a database named postgres is created in the DB instance.  Constraints:  * Must contain 1 to 63 letters, numbers, or underscores.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  Oracle  The Oracle System ID (SID) of the created DB instance. If you specify null, the default value ORCL is used. You can't specify the string NULL, or any other reserved word, for DBName.  Default: ORCL  Constraints:  * Can't be longer than 8 characters  Amazon RDS Custom for Oracle  The Oracle System ID (SID) of the created RDS Custom DB instance. If you don't specify a value, the default value is ORCL.  Default: ORCL  Constraints:  * It must contain 1 to 8 alphanumeric characters.  * It must contain a letter.  * It can't be a word reserved by the database engine.  Amazon RDS Custom for SQL Server  Not applicable. Must be null.  SQL Server  Not applicable. Must be null.  Amazon Aurora MySQL  The name of the database to create when the primary DB instance of the Aurora MySQL DB cluster is created. If this parameter isn't specified for an Aurora MySQL DB cluster, no database is created in the DB cluster.  Constraints:  * It must contain 1 to 64 alphanumeric characters.  * It can't be a word reserved by the database engine.  Amazon Aurora PostgreSQL  The name of the database to create when the primary DB instance of the Aurora PostgreSQL DB cluster is created. If this parameter isn't specified for an Aurora PostgreSQL DB cluster, a database named postgres is created in the DB cluster.  Constraints:  * It must contain 1 to 63 alphanumeric characters.  * It must begin with a letter or an underscore. Subsequent characters can be letters, underscores, or digits (0 to 9).  * It can't be a word reserved by the database engine.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"db_parameter_group_name": {
						Description:         "The name of the DB parameter group to associate with this DB instance. If you do not specify a value, then the default DB parameter group for the specified DB engine and version is used.  This setting doesn't apply to RDS Custom.  Constraints:  * Must be 1 to 255 letters, numbers, or hyphens.  * First character must be a letter  * Can't end with a hyphen or contain two consecutive hyphens",
						MarkdownDescription: "The name of the DB parameter group to associate with this DB instance. If you do not specify a value, then the default DB parameter group for the specified DB engine and version is used.  This setting doesn't apply to RDS Custom.  Constraints:  * Must be 1 to 255 letters, numbers, or hyphens.  * First character must be a letter  * Can't end with a hyphen or contain two consecutive hyphens",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"db_parameter_group_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"db_snapshot_identifier": {
						Description:         "The identifier for the DB snapshot to restore from.  Constraints:  * Must match the identifier of an existing DBSnapshot.  * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB snapshot.",
						MarkdownDescription: "The identifier for the DB snapshot to restore from.  Constraints:  * Must match the identifier of an existing DBSnapshot.  * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB snapshot.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"db_subnet_group_name": {
						Description:         "A DB subnet group to associate with this DB instance.  Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.  Example: mydbsubnetgroup",
						MarkdownDescription: "A DB subnet group to associate with this DB instance.  Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.  Example: mydbsubnetgroup",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"db_subnet_group_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"deletion_protection": {
						Description:         "A value that indicates whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).  Amazon Aurora  Not applicable. You can enable or disable deletion protection for the DB cluster. For more information, see CreateDBCluster. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.",
						MarkdownDescription: "A value that indicates whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).  Amazon Aurora  Not applicable. You can enable or disable deletion protection for the DB cluster. For more information, see CreateDBCluster. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"destination_region": {
						Description:         "DestinationRegion is used for presigning the request to a given region.",
						MarkdownDescription: "DestinationRegion is used for presigning the request to a given region.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"domain": {
						Description:         "The Active Directory directory ID to create the DB instance in. Currently, only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.  For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",
						MarkdownDescription: "The Active Directory directory ID to create the DB instance in. Currently, only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.  For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"domain_iam_role_name": {
						Description:         "Specify the name of the IAM role to be used when making API calls to the Directory Service.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",
						MarkdownDescription: "Specify the name of the IAM role to be used when making API calls to the Directory Service.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_cloudwatch_logs_exports": {
						Description:         "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine. For more information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. CloudWatch Logs exports are managed by the DB cluster.  RDS Custom  Not applicable.  MariaDB  Possible values are audit, error, general, and slowquery.  Microsoft SQL Server  Possible values are agent and error.  MySQL  Possible values are audit, error, general, and slowquery.  Oracle  Possible values are alert, audit, listener, trace, and oemagent.  PostgreSQL  Possible values are postgresql and upgrade.",
						MarkdownDescription: "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine. For more information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. CloudWatch Logs exports are managed by the DB cluster.  RDS Custom  Not applicable.  MariaDB  Possible values are audit, error, general, and slowquery.  Microsoft SQL Server  Possible values are agent and error.  MySQL  Possible values are audit, error, general, and slowquery.  Oracle  Possible values are alert, audit, listener, trace, and oemagent.  PostgreSQL  Possible values are postgresql and upgrade.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_customer_owned_ip": {
						Description:         "A value that indicates whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance.  A CoIP provides local or external connectivity to resources in your Outpost subnets through your on-premises network. For some use cases, a CoIP can provide lower latency for connections to the DB instance from outside of its virtual private cloud (VPC) on your local network.  For more information about RDS on Outposts, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.  For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing) in the Amazon Web Services Outposts User Guide.",
						MarkdownDescription: "A value that indicates whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance.  A CoIP provides local or external connectivity to resources in your Outpost subnets through your on-premises network. For some use cases, a CoIP can provide lower latency for connections to the DB instance from outside of its virtual private cloud (VPC) on your local network.  For more information about RDS on Outposts, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.  For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing) in the Amazon Web Services Outposts User Guide.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_iam_database_authentication": {
						Description:         "A value that indicates whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled.  For more information, see IAM Database Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Mapping Amazon Web Services IAM accounts to database accounts is managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled.  For more information, see IAM Database Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Mapping Amazon Web Services IAM accounts to database accounts is managed by the DB cluster.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"engine": {
						Description:         "The name of the database engine to be used for this instance.  Not every database engine is available for every Amazon Web Services Region.  Valid Values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * custom-oracle-ee (for RDS Custom for Oracle instances)  * custom-sqlserver-ee (for RDS Custom for SQL Server instances)  * custom-sqlserver-se (for RDS Custom for SQL Server instances)  * custom-sqlserver-web (for RDS Custom for SQL Server instances)  * mariadb  * mysql  * oracle-ee  * oracle-ee-cdb  * oracle-se2  * oracle-se2-cdb  * postgres  * sqlserver-ee  * sqlserver-se  * sqlserver-ex  * sqlserver-web",
						MarkdownDescription: "The name of the database engine to be used for this instance.  Not every database engine is available for every Amazon Web Services Region.  Valid Values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * custom-oracle-ee (for RDS Custom for Oracle instances)  * custom-sqlserver-ee (for RDS Custom for SQL Server instances)  * custom-sqlserver-se (for RDS Custom for SQL Server instances)  * custom-sqlserver-web (for RDS Custom for SQL Server instances)  * mariadb  * mysql  * oracle-ee  * oracle-ee-cdb  * oracle-se2  * oracle-se2-cdb  * postgres  * sqlserver-ee  * sqlserver-se  * sqlserver-ex  * sqlserver-web",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"engine_version": {
						Description:         "The version number of the database engine to use.  For a list of valid engine versions, use the DescribeDBEngineVersions operation.  The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every Amazon Web Services Region.  Amazon Aurora  Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.  Amazon RDS Custom for Oracle  A custom engine version (CEV) that you have previously created. This setting is required for RDS Custom for Oracle. The CEV name has the following format: 19.customized_string . An example identifier is 19.my_cev1. For more information, see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create) in the Amazon RDS User Guide.  Amazon RDS Custom for SQL Server  See RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html) in the Amazon RDS User Guide.  MariaDB  For information, see MariaDB on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Microsoft SQL Server  For information, see Microsoft SQL Server Versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the Amazon RDS User Guide.  MySQL  For information, see MySQL on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Oracle  For information, see Oracle Database Engine Release Notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the Amazon RDS User Guide.  PostgreSQL  For information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.",
						MarkdownDescription: "The version number of the database engine to use.  For a list of valid engine versions, use the DescribeDBEngineVersions operation.  The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every Amazon Web Services Region.  Amazon Aurora  Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.  Amazon RDS Custom for Oracle  A custom engine version (CEV) that you have previously created. This setting is required for RDS Custom for Oracle. The CEV name has the following format: 19.customized_string . An example identifier is 19.my_cev1. For more information, see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create) in the Amazon RDS User Guide.  Amazon RDS Custom for SQL Server  See RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html) in the Amazon RDS User Guide.  MariaDB  For information, see MariaDB on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Microsoft SQL Server  For information, see Microsoft SQL Server Versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the Amazon RDS User Guide.  MySQL  For information, see MySQL on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Oracle  For information, see Oracle Database Engine Release Notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the Amazon RDS User Guide.  PostgreSQL  For information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iops": {
						Description:         "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for the DB instance. For information about valid Iops values, see Amazon RDS Provisioned IOPS storage to improve performance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the Amazon RDS User Guide.  Constraints: For MariaDB, MySQL, Oracle, and PostgreSQL DB instances, must be a multiple between .5 and 50 of the storage amount for the DB instance. For SQL Server DB instances, must be a multiple between 1 and 50 of the storage amount for the DB instance.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for the DB instance. For information about valid Iops values, see Amazon RDS Provisioned IOPS storage to improve performance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#USER_PIOPS) in the Amazon RDS User Guide.  Constraints: For MariaDB, MySQL, Oracle, and PostgreSQL DB instances, must be a multiple between .5 and 50 of the storage amount for the DB instance. For SQL Server DB instances, must be a multiple between 1 and 50 of the storage amount for the DB instance.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kms_key_id": {
						Description:         "The Amazon Web Services KMS key identifier for an encrypted DB instance.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  Amazon Aurora  Not applicable. The Amazon Web Services KMS key identifier is managed by the DB cluster. For more information, see CreateDBCluster.  If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Amazon RDS Custom  A KMS key is required for RDS Custom instances. For most RDS engines, if you leave this parameter empty while enabling StorageEncrypted, the engine uses the default KMS key. However, RDS Custom doesn't use the default key when this parameter is empty. You must explicitly specify a key.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for an encrypted DB instance.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  Amazon Aurora  Not applicable. The Amazon Web Services KMS key identifier is managed by the DB cluster. For more information, see CreateDBCluster.  If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Amazon RDS Custom  A KMS key is required for RDS Custom instances. For most RDS engines, if you leave this parameter empty while enabling StorageEncrypted, the engine uses the default KMS key. However, RDS Custom doesn't use the default key when this parameter is empty. You must explicitly specify a key.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kms_key_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"license_model": {
						Description:         "License model information for this DB instance.  Valid values: license-included | bring-your-own-license | general-public-license  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "License model information for this DB instance.  Valid values: license-included | bring-your-own-license | general-public-license  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"master_user_password": {
						Description:         "The password for the master user. The password can include any printable ASCII character except '/', ''', or '@'.  Amazon Aurora  Not applicable. The password for the master user is managed by the DB cluster.  MariaDB  Constraints: Must contain from 8 to 41 characters.  Microsoft SQL Server  Constraints: Must contain from 8 to 128 characters.  MySQL  Constraints: Must contain from 8 to 41 characters.  Oracle  Constraints: Must contain from 8 to 30 characters.  PostgreSQL  Constraints: Must contain from 8 to 128 characters.",
						MarkdownDescription: "The password for the master user. The password can include any printable ASCII character except '/', ''', or '@'.  Amazon Aurora  Not applicable. The password for the master user is managed by the DB cluster.  MariaDB  Constraints: Must contain from 8 to 41 characters.  Microsoft SQL Server  Constraints: Must contain from 8 to 128 characters.  MySQL  Constraints: Must contain from 8 to 41 characters.  Oracle  Constraints: Must contain from 8 to 30 characters.  PostgreSQL  Constraints: Must contain from 8 to 128 characters.",

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

					"master_username": {
						Description:         "The name for the master user.  Amazon Aurora  Not applicable. The name for the master user is managed by the DB cluster.  Amazon RDS  Constraints:  * Required.  * Must be 1 to 16 letters, numbers, or underscores.  * First character must be a letter.  * Can't be a reserved word for the chosen database engine.",
						MarkdownDescription: "The name for the master user.  Amazon Aurora  Not applicable. The name for the master user is managed by the DB cluster.  Amazon RDS  Constraints:  * Required.  * Must be 1 to 16 letters, numbers, or underscores.  * First character must be a letter.  * Can't be a reserved word for the chosen database engine.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_allocated_storage": {
						Description:         "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.  For more information about this setting, including limitations that apply to it, see Managing capacity automatically with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.  For more information about this setting, including limitations that apply to it, see Managing capacity automatically with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring_interval": {
						Description:         "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify 0. The default is 0.  If MonitoringRoleArn is specified, then you must set MonitoringInterval to a value other than 0.  This setting doesn't apply to RDS Custom.  Valid Values: 0, 1, 5, 10, 15, 30, 60",
						MarkdownDescription: "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify 0. The default is 0.  If MonitoringRoleArn is specified, then you must set MonitoringInterval to a value other than 0.  This setting doesn't apply to RDS Custom.  Valid Values: 0, 1, 5, 10, 15, 30, 60",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring_role_arn": {
						Description:         "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting Up and Enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide.  If MonitoringInterval is set to a value other than 0, then you must supply a MonitoringRoleArn value.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting Up and Enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide.  If MonitoringInterval is set to a value other than 0, then you must supply a MonitoringRoleArn value.  This setting doesn't apply to RDS Custom.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"multi_az": {
						Description:         "A value that indicates whether the DB instance is a Multi-AZ deployment. You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ deployment.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. DB instance Availability Zones (AZs) are managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether the DB instance is a Multi-AZ deployment. You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ deployment.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. DB instance Availability Zones (AZs) are managed by the DB cluster.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nchar_character_set_name": {
						Description:         "The name of the NCHAR character set for the Oracle DB instance.  This parameter doesn't apply to RDS Custom.",
						MarkdownDescription: "The name of the NCHAR character set for the Oracle DB instance.  This parameter doesn't apply to RDS Custom.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_type": {
						Description:         "The network type of the DB instance.  Valid values:  * IPV4  * DUAL  The network type is determined by the DBSubnetGroup specified for the DB instance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL).  For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon RDS User Guide.",
						MarkdownDescription: "The network type of the DB instance.  Valid values:  * IPV4  * DUAL  The network type is determined by the DBSubnetGroup specified for the DB instance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL).  For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon RDS User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"option_group_name": {
						Description:         "A value that indicates that the DB instance should be associated with the specified option group.  Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance after it is associated with a DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "A value that indicates that the DB instance should be associated with the specified option group.  Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance after it is associated with a DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"performance_insights_enabled": {
						Description:         "A value that indicates whether to enable Performance Insights for the DB instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "A value that indicates whether to enable Performance Insights for the DB instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"performance_insights_kms_key_id": {
						Description:         "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.  If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.  If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  This setting doesn't apply to RDS Custom.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"performance_insights_retention_period": {
						Description:         "The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:  * 7  * month * 31, where month is a number of months from 1-23  * 731  For example, the following values are valid:  * 93 (3 months * 31)  * 341 (11 months * 31)  * 589 (19 months * 31)  * 731  If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:  * 7  * month * 31, where month is a number of months from 1-23  * 731  For example, the following values are valid:  * 93 (3 months * 31)  * 341 (11 months * 31)  * 589 (19 months * 31)  * 731  If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.  This setting doesn't apply to RDS Custom.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"port": {
						Description:         "The port number on which the database accepts connections.  MySQL  Default: 3306  Valid values: 1150-65535  Type: Integer  MariaDB  Default: 3306  Valid values: 1150-65535  Type: Integer  PostgreSQL  Default: 5432  Valid values: 1150-65535  Type: Integer  Oracle  Default: 1521  Valid values: 1150-65535  SQL Server  Default: 1433  Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and 49152-49156.  Amazon Aurora  Default: 3306  Valid values: 1150-65535  Type: Integer",
						MarkdownDescription: "The port number on which the database accepts connections.  MySQL  Default: 3306  Valid values: 1150-65535  Type: Integer  MariaDB  Default: 3306  Valid values: 1150-65535  Type: Integer  PostgreSQL  Default: 5432  Valid values: 1150-65535  Type: Integer  Oracle  Default: 1521  Valid values: 1150-65535  SQL Server  Default: 1433  Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and 49152-49156.  Amazon Aurora  Default: 3306  Valid values: 1150-65535  Type: Integer",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pre_signed_url": {
						Description:         "When you are creating a read replica from one Amazon Web Services GovCloud (US) Region to another or from one China Amazon Web Services Region to another, the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplica API operation in the source Amazon Web Services Region that contains the source DB instance.  This setting applies only to Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. It's ignored in other Amazon Web Services Regions.  You must specify this parameter when you create an encrypted read replica from another Amazon Web Services Region by using the Amazon RDS API. Don't specify PreSignedUrl when you are creating an encrypted read replica in the same Amazon Web Services Region.  The presigned URL must be a valid request for the CreateDBInstanceReadReplica API operation that can run in the source Amazon Web Services Region that contains the encrypted source DB instance. The presigned URL request must contain the following parameter values:  * DestinationRegion - The Amazon Web Services Region that the encrypted read replica is created in. This Amazon Web Services Region is the same one where the CreateDBInstanceReadReplica operation is called that contains this presigned URL. For example, if you create an encrypted DB instance in the us-west-1 Amazon Web Services Region, from a source DB instance in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica operation in the us-east-1 Amazon Web Services Region and provide a presigned URL that contains a call to the CreateDBInstanceReadReplica operation in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion in the presigned URL must be set to the us-east-1 Amazon Web Services Region.  * KmsKeyId - The KMS key identifier for the key to use to encrypt the read replica in the destination Amazon Web Services Region. This is the same identifier for both the CreateDBInstanceReadReplica operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL.  * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted DB instance to be replicated. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are creating an encrypted read replica from a DB instance in the us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115.  To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).  If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region.  SourceRegion isn't supported for SQL Server, because Amazon RDS for SQL Server doesn't support cross-Region read replicas.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "When you are creating a read replica from one Amazon Web Services GovCloud (US) Region to another or from one China Amazon Web Services Region to another, the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplica API operation in the source Amazon Web Services Region that contains the source DB instance.  This setting applies only to Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. It's ignored in other Amazon Web Services Regions.  You must specify this parameter when you create an encrypted read replica from another Amazon Web Services Region by using the Amazon RDS API. Don't specify PreSignedUrl when you are creating an encrypted read replica in the same Amazon Web Services Region.  The presigned URL must be a valid request for the CreateDBInstanceReadReplica API operation that can run in the source Amazon Web Services Region that contains the encrypted source DB instance. The presigned URL request must contain the following parameter values:  * DestinationRegion - The Amazon Web Services Region that the encrypted read replica is created in. This Amazon Web Services Region is the same one where the CreateDBInstanceReadReplica operation is called that contains this presigned URL. For example, if you create an encrypted DB instance in the us-west-1 Amazon Web Services Region, from a source DB instance in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica operation in the us-east-1 Amazon Web Services Region and provide a presigned URL that contains a call to the CreateDBInstanceReadReplica operation in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion in the presigned URL must be set to the us-east-1 Amazon Web Services Region.  * KmsKeyId - The KMS key identifier for the key to use to encrypt the read replica in the destination Amazon Web Services Region. This is the same identifier for both the CreateDBInstanceReadReplica operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL.  * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted DB instance to be replicated. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are creating an encrypted read replica from a DB instance in the us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115.  To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).  If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region.  SourceRegion isn't supported for SQL Server, because Amazon RDS for SQL Server doesn't support cross-Region read replicas.  This setting doesn't apply to RDS Custom.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"preferred_backup_window": {
						Description:         "The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. For more information, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. The daily time range for creating automated backups is managed by the DB cluster.  Constraints:  * Must be in the format hh24:mi-hh24:mi.  * Must be in Universal Coordinated Time (UTC).  * Must not conflict with the preferred maintenance window.  * Must be at least 30 minutes.",
						MarkdownDescription: "The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. For more information, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. The daily time range for creating automated backups is managed by the DB cluster.  Constraints:  * Must be in the format hh24:mi-hh24:mi.  * Must be in Universal Coordinated Time (UTC).  * Must not conflict with the preferred maintenance window.  * Must be at least 30 minutes.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"preferred_maintenance_window": {
						Description:         "The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC). For more information, see Amazon RDS Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).  Format: ddd:hh24:mi-ddd:hh24:mi  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week.  Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.  Constraints: Minimum 30-minute window.",
						MarkdownDescription: "The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC). For more information, see Amazon RDS Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).  Format: ddd:hh24:mi-ddd:hh24:mi  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week.  Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.  Constraints: Minimum 30-minute window.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"processor_features": {
						Description:         "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
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

					"promotion_tier": {
						Description:         "A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the Amazon Aurora User Guide.  This setting doesn't apply to RDS Custom.  Default: 1  Valid Values: 0 - 15",
						MarkdownDescription: "A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the Amazon Aurora User Guide.  This setting doesn't apply to RDS Custom.  Default: 1  Valid Values: 0 - 15",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"publicly_accessible": {
						Description:         "A value that indicates whether the DB instance is publicly accessible.  When the DB instance is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB instance's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB instance's VPC. Access to the DB instance is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB instance doesn't permit it.  When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a private IP address.  Default: The default behavior varies depending on whether DBSubnetGroupName is specified.  If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies:  * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB instance is private.  * If the default VPC in the target Region has an internet gateway attached to it, the DB instance is public.  If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies:  * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB instance is private.  * If the subnets are part of a VPC that has an internet gateway attached to it, the DB instance is public.",
						MarkdownDescription: "A value that indicates whether the DB instance is publicly accessible.  When the DB instance is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB instance's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB instance's VPC. Access to the DB instance is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB instance doesn't permit it.  When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a private IP address.  Default: The default behavior varies depending on whether DBSubnetGroupName is specified.  If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies:  * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB instance is private.  * If the default VPC in the target Region has an internet gateway attached to it, the DB instance is public.  If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies:  * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB instance is private.  * If the subnets are part of a VPC that has an internet gateway attached to it, the DB instance is public.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replica_mode": {
						Description:         "The open mode of the replica database: mounted or read-only.  This parameter is only supported for Oracle DB instances.  Mounted DB replicas are included in Oracle Database Enterprise Edition. The main use case for mounted replicas is cross-Region disaster recovery. The primary database doesn't use Active Data Guard to transmit information to the mounted replica. Because it doesn't accept user connections, a mounted replica can't serve a read-only workload.  You can create a combination of mounted and read-only DB replicas for the same primary DB instance. For more information, see Working with Oracle Read Replicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide.  For RDS Custom, you must specify this parameter and set it to mounted. The value won't be set by default. After replica creation, you can manage the open mode manually.",
						MarkdownDescription: "The open mode of the replica database: mounted or read-only.  This parameter is only supported for Oracle DB instances.  Mounted DB replicas are included in Oracle Database Enterprise Edition. The main use case for mounted replicas is cross-Region disaster recovery. The primary database doesn't use Active Data Guard to transmit information to the mounted replica. Because it doesn't accept user connections, a mounted replica can't serve a read-only workload.  You can create a combination of mounted and read-only DB replicas for the same primary DB instance. For more information, see Working with Oracle Read Replicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide.  For RDS Custom, you must specify this parameter and set it to mounted. The value won't be set by default. After replica creation, you can manage the open mode manually.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_db_instance_identifier": {
						Description:         "The identifier of the DB instance that will act as the source for the read replica. Each DB instance can have up to five read replicas.  Constraints:  * Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL, or SQL Server DB instance.  * Can specify a DB instance that is a MySQL read replica only if the source is running MySQL 5.6 or later.  * For the limitations of Oracle read replicas, see Read Replica Limitations with Oracle (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide.  * For the limitations of SQL Server read replicas, see Read Replica Limitations with Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.Limitations.html) in the Amazon RDS User Guide.  * Can specify a PostgreSQL DB instance only if the source is running PostgreSQL 9.3.5 or later (9.4.7 and higher for cross-Region replication).  * The specified DB instance must have automatic backups enabled, that is, its backup retention period must be greater than 0.  * If the source DB instance is in the same Amazon Web Services Region as the read replica, specify a valid DB instance identifier.  * If the source DB instance is in a different Amazon Web Services Region from the read replica, specify a valid DB instance ARN. For more information, see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing) in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS Custom, which don't support cross-Region replicas.",
						MarkdownDescription: "The identifier of the DB instance that will act as the source for the read replica. Each DB instance can have up to five read replicas.  Constraints:  * Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL, or SQL Server DB instance.  * Can specify a DB instance that is a MySQL read replica only if the source is running MySQL 5.6 or later.  * For the limitations of Oracle read replicas, see Read Replica Limitations with Oracle (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide.  * For the limitations of SQL Server read replicas, see Read Replica Limitations with Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.Limitations.html) in the Amazon RDS User Guide.  * Can specify a PostgreSQL DB instance only if the source is running PostgreSQL 9.3.5 or later (9.4.7 and higher for cross-Region replication).  * The specified DB instance must have automatic backups enabled, that is, its backup retention period must be greater than 0.  * If the source DB instance is in the same Amazon Web Services Region as the read replica, specify a valid DB instance identifier.  * If the source DB instance is in a different Amazon Web Services Region from the read replica, specify a valid DB instance ARN. For more information, see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing) in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS Custom, which don't support cross-Region replicas.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_region": {
						Description:         "SourceRegion is the source region where the resource exists. This is not sent over the wire and is only used for presigning. This value should always have the same region as the source ARN.",
						MarkdownDescription: "SourceRegion is the source region where the resource exists. This is not sent over the wire and is only used for presigning. This value should always have the same region as the source ARN.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_encrypted": {
						Description:         "A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.  For RDS Custom instances, either set this parameter to true or leave it unset. If you set this parameter to false, RDS reports an error.  Amazon Aurora  Not applicable. The encryption for DB instances is managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.  For RDS Custom instances, either set this parameter to true or leave it unset. If you set this parameter to false, RDS reports an error.  Amazon Aurora  Not applicable. The encryption for DB instances is managed by the DB cluster.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_type": {
						Description:         "Specifies the storage type to be associated with the DB instance.  Valid values: standard | gp2 | io1  If you specify io1, you must also include a value for the Iops parameter.  Default: io1 if the Iops parameter is specified, otherwise gp2  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "Specifies the storage type to be associated with the DB instance.  Valid values: standard | gp2 | io1  If you specify io1, you must also include a value for the Iops parameter.  Default: io1 if the Iops parameter is specified, otherwise gp2  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "Tags to assign to the DB instance.",
						MarkdownDescription: "Tags to assign to the DB instance.",

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

					"tde_credential_arn": {
						Description:         "The ARN from the key store with which to associate the instance for TDE encryption.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "The ARN from the key store with which to associate the instance for TDE encryption.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tde_credential_password": {
						Description:         "The password for the given ARN from the key store in order to access the device.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The password for the given ARN from the key store in order to access the device.  This setting doesn't apply to RDS Custom.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timezone": {
						Description:         "The time zone of the DB instance. The time zone parameter is currently supported only by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
						MarkdownDescription: "The time zone of the DB instance. The time zone parameter is currently supported only by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"use_default_processor_features": {
						Description:         "A value that indicates whether the DB instance class of the DB instance uses its default processor features.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "A value that indicates whether the DB instance class of the DB instance uses its default processor features.  This setting doesn't apply to RDS Custom.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vpc_security_group_i_ds": {
						Description:         "A list of Amazon EC2 VPC security groups to associate with this DB instance.  Amazon Aurora  Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster.  Default: The default EC2 VPC security group for the DB subnet group's VPC.",
						MarkdownDescription: "A list of Amazon EC2 VPC security groups to associate with this DB instance.  Amazon Aurora  Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster.  Default: The default EC2 VPC security group for the DB subnet group's VPC.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vpc_security_group_refs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *RdsServicesK8SAwsDBInstanceV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")

	var state RdsServicesK8SAwsDBInstanceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel RdsServicesK8SAwsDBInstanceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("rds.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("DBInstance")

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

func (r *RdsServicesK8SAwsDBInstanceV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *RdsServicesK8SAwsDBInstanceV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")

	var state RdsServicesK8SAwsDBInstanceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel RdsServicesK8SAwsDBInstanceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("rds.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("DBInstance")

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

func (r *RdsServicesK8SAwsDBInstanceV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
