/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rds_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &RdsServicesK8SAwsDbinstanceV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &RdsServicesK8SAwsDbinstanceV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &RdsServicesK8SAwsDbinstanceV1Alpha1Resource{}
)

func NewRdsServicesK8SAwsDbinstanceV1Alpha1Resource() resource.Resource {
	return &RdsServicesK8SAwsDbinstanceV1Alpha1Resource{}
}

type RdsServicesK8SAwsDbinstanceV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_db_instance_v1alpha1"
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DBInstance is the Schema for the DBInstances API",
		MarkdownDescription: "DBInstance is the Schema for the DBInstances API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "DBInstanceSpec defines the desired state of DBInstance.  Contains the details of an Amazon RDS DB instance.  This data type is used as a response element in the operations CreateDBInstance, CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance, PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3, RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",
				MarkdownDescription: "DBInstanceSpec defines the desired state of DBInstance.  Contains the details of an Amazon RDS DB instance.  This data type is used as a response element in the operations CreateDBInstance, CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances, ModifyDBInstance, PromoteReadReplica, RebootDBInstance, RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3, RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.",
				Attributes: map[string]schema.Attribute{
					"allocated_storage": schema.Int64Attribute{
						Description:         "The amount of storage in gibibytes (GiB) to allocate for the DB instance.  Type: Integer  Amazon Aurora  Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.  Amazon RDS Custom  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  * Provisioned IOPS storage (io1): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  MySQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  MariaDB  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  PostgreSQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  Oracle  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 10 to 3072.  SQL Server  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Enterprise and Standard editions: Must be an integer from 20 to 16384. Web and Express editions: Must be an integer from 20 to 16384.  * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must be an integer from 100 to 16384. Web and Express editions: Must be an integer from 100 to 16384.  * Magnetic storage (standard): Enterprise and Standard editions: Must be an integer from 20 to 1024. Web and Express editions: Must be an integer from 20 to 1024.",
						MarkdownDescription: "The amount of storage in gibibytes (GiB) to allocate for the DB instance.  Type: Integer  Amazon Aurora  Not applicable. Aurora cluster volumes automatically grow as the amount of data in your database increases, though you are only charged for the space that you use in an Aurora cluster volume.  Amazon RDS Custom  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  * Provisioned IOPS storage (io1): Must be an integer from 40 to 65536 for RDS Custom for Oracle, 16384 for RDS Custom for SQL Server.  MySQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  MariaDB  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  PostgreSQL  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 5 to 3072.  Oracle  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Must be an integer from 20 to 65536.  * Provisioned IOPS storage (io1): Must be an integer from 100 to 65536.  * Magnetic storage (standard): Must be an integer from 10 to 3072.  SQL Server  Constraints to the amount of storage for each storage type are the following:  * General Purpose (SSD) storage (gp2, gp3): Enterprise and Standard editions: Must be an integer from 20 to 16384. Web and Express editions: Must be an integer from 20 to 16384.  * Provisioned IOPS storage (io1): Enterprise and Standard editions: Must be an integer from 100 to 16384. Web and Express editions: Must be an integer from 100 to 16384.  * Magnetic storage (standard): Enterprise and Standard editions: Must be an integer from 20 to 1024. Web and Express editions: Must be an integer from 20 to 1024.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auto_minor_version_upgrade": schema.BoolAttribute{
						Description:         "A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.  If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgrade to false.",
						MarkdownDescription: "A value that indicates whether minor engine upgrades are applied automatically to the DB instance during the maintenance window. By default, minor engine upgrades are applied automatically.  If you create an RDS Custom DB instance, you must set AutoMinorVersionUpgrade to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"availability_zone": schema.StringAttribute{
						Description:         "The Availability Zone (AZ) where the database will be created. For information on Amazon Web Services Regions and Availability Zones, see Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).  Amazon Aurora  Each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.  Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web Services Region.  Example: us-east-1d  Constraint: The AvailabilityZone parameter can't be specified if the DB instance is a Multi-AZ deployment. The specified Availability Zone must be in the same Amazon Web Services Region as the current endpoint.",
						MarkdownDescription: "The Availability Zone (AZ) where the database will be created. For information on Amazon Web Services Regions and Availability Zones, see Regions and Availability Zones (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).  Amazon Aurora  Each Aurora DB cluster hosts copies of its storage in three separate Availability Zones. Specify one of these Availability Zones. Aurora automatically chooses an appropriate Availability Zone if you don't specify one.  Default: A random, system-chosen Availability Zone in the endpoint's Amazon Web Services Region.  Example: us-east-1d  Constraint: The AvailabilityZone parameter can't be specified if the DB instance is a Multi-AZ deployment. The specified Availability Zone must be in the same Amazon Web Services Region as the current endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_retention_period": schema.Int64Attribute{
						Description:         "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.  Amazon Aurora  Not applicable. The retention period for automated backups is managed by the DB cluster.  Default: 1  Constraints:  * Must be a value from 0 to 35  * Can't be set to 0 if the DB instance is a source to read replicas  * Can't be set to 0 for an RDS Custom for Oracle DB instance",
						MarkdownDescription: "The number of days for which automated backups are retained. Setting this parameter to a positive number enables backups. Setting this parameter to 0 disables automated backups.  Amazon Aurora  Not applicable. The retention period for automated backups is managed by the DB cluster.  Default: 1  Constraints:  * Must be a value from 0 to 35  * Can't be set to 0 if the DB instance is a source to read replicas  * Can't be set to 0 for an RDS Custom for Oracle DB instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_target": schema.StringAttribute{
						Description:         "Specifies where automated backups and manual snapshots are stored.  Possible values are outposts (Amazon Web Services Outposts) and region (Amazon Web Services Region). The default is region.  For more information, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.",
						MarkdownDescription: "Specifies where automated backups and manual snapshots are stored.  Possible values are outposts (Amazon Web Services Outposts) and region (Amazon Web Services Region). The default is region.  For more information, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ca_certificate_identifier": schema.StringAttribute{
						Description:         "Specifies the CA certificate identifier to use for the DB instance’s server certificate.  This setting doesn't apply to RDS Custom.  For more information, see Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the Amazon Aurora User Guide.",
						MarkdownDescription: "Specifies the CA certificate identifier to use for the DB instance’s server certificate.  This setting doesn't apply to RDS Custom.  For more information, see Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the Amazon Aurora User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"character_set_name": schema.StringAttribute{
						Description:         "For supported engines, this value indicates that the DB instance should be associated with the specified CharacterSet.  This setting doesn't apply to RDS Custom. However, if you need to change the character set, you can change it on the database itself.  Amazon Aurora  Not applicable. The character set is managed by the DB cluster. For more information, see CreateDBCluster.",
						MarkdownDescription: "For supported engines, this value indicates that the DB instance should be associated with the specified CharacterSet.  This setting doesn't apply to RDS Custom. However, if you need to change the character set, you can change it on the database itself.  Amazon Aurora  Not applicable. The character set is managed by the DB cluster. For more information, see CreateDBCluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"copy_tags_to_snapshot": schema.BoolAttribute{
						Description:         "A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.  Amazon Aurora  Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.",
						MarkdownDescription: "A value that indicates whether to copy tags from the DB instance to snapshots of the DB instance. By default, tags are not copied.  Amazon Aurora  Not applicable. Copying tags to snapshots is managed by the DB cluster. Setting this value for an Aurora DB instance has no effect on the DB cluster setting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_iam_instance_profile": schema.StringAttribute{
						Description:         "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance. The instance profile must meet the following requirements:  * The profile must exist in your account.  * The profile must have an IAM role that Amazon EC2 has permissions to assume.  * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom.  For the list of permissions required for the IAM role, see Configure IAM and your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the Amazon RDS User Guide.  This setting is required for RDS Custom.",
						MarkdownDescription: "The instance profile associated with the underlying Amazon EC2 instance of an RDS Custom DB instance. The instance profile must meet the following requirements:  * The profile must exist in your account.  * The profile must have an IAM role that Amazon EC2 has permissions to assume.  * The instance profile name and the associated IAM role name must start with the prefix AWSRDSCustom.  For the list of permissions required for the IAM role, see Configure IAM and your VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-setup-orcl.html#custom-setup-orcl.iam-vpc) in the Amazon RDS User Guide.  This setting is required for RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_identifier": schema.StringAttribute{
						Description:         "The identifier of the DB cluster that the instance will belong to.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The identifier of the DB cluster that the instance will belong to.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_snapshot_identifier": schema.StringAttribute{
						Description:         "The identifier for the RDS for MySQL Multi-AZ DB cluster snapshot to restore from.  For more information on Multi-AZ DB clusters, see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the Amazon RDS User Guide.  Constraints:  * Must match the identifier of an existing Multi-AZ DB cluster snapshot.  * Can't be specified when DBSnapshotIdentifier is specified.  * Must be specified when DBSnapshotIdentifier isn't specified.  * If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot.  * Can't be the identifier of an Aurora DB cluster snapshot.  * Can't be the identifier of an RDS for PostgreSQL Multi-AZ DB cluster snapshot.",
						MarkdownDescription: "The identifier for the RDS for MySQL Multi-AZ DB cluster snapshot to restore from.  For more information on Multi-AZ DB clusters, see Multi-AZ DB cluster deployments (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html) in the Amazon RDS User Guide.  Constraints:  * Must match the identifier of an existing Multi-AZ DB cluster snapshot.  * Can't be specified when DBSnapshotIdentifier is specified.  * Must be specified when DBSnapshotIdentifier isn't specified.  * If you are restoring from a shared manual Multi-AZ DB cluster snapshot, the DBClusterSnapshotIdentifier must be the ARN of the shared snapshot.  * Can't be the identifier of an Aurora DB cluster snapshot.  * Can't be the identifier of an RDS for PostgreSQL Multi-AZ DB cluster snapshot.",
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
						Description:         "The DB instance identifier. This parameter is stored as a lowercase string.  Constraints:  * Must contain from 1 to 63 letters, numbers, or hyphens.  * First character must be a letter.  * Can't end with a hyphen or contain two consecutive hyphens.  Example: mydbinstance",
						MarkdownDescription: "The DB instance identifier. This parameter is stored as a lowercase string.  Constraints:  * Must contain from 1 to 63 letters, numbers, or hyphens.  * First character must be a letter.  * Can't end with a hyphen or contain two consecutive hyphens.  Example: mydbinstance",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_name": schema.StringAttribute{
						Description:         "The meaning of this parameter differs according to the database engine you use.  MySQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  MariaDB  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  PostgreSQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, a database named postgres is created in the DB instance.  Constraints:  * Must contain 1 to 63 letters, numbers, or underscores.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  Oracle  The Oracle System ID (SID) of the created DB instance. If you specify null, the default value ORCL is used. You can't specify the string NULL, or any other reserved word, for DBName.  Default: ORCL  Constraints:  * Can't be longer than 8 characters  Amazon RDS Custom for Oracle  The Oracle System ID (SID) of the created RDS Custom DB instance. If you don't specify a value, the default value is ORCL.  Default: ORCL  Constraints:  * It must contain 1 to 8 alphanumeric characters.  * It must contain a letter.  * It can't be a word reserved by the database engine.  Amazon RDS Custom for SQL Server  Not applicable. Must be null.  SQL Server  Not applicable. Must be null.  Amazon Aurora MySQL  The name of the database to create when the primary DB instance of the Aurora MySQL DB cluster is created. If this parameter isn't specified for an Aurora MySQL DB cluster, no database is created in the DB cluster.  Constraints:  * It must contain 1 to 64 alphanumeric characters.  * It can't be a word reserved by the database engine.  Amazon Aurora PostgreSQL  The name of the database to create when the primary DB instance of the Aurora PostgreSQL DB cluster is created. If this parameter isn't specified for an Aurora PostgreSQL DB cluster, a database named postgres is created in the DB cluster.  Constraints:  * It must contain 1 to 63 alphanumeric characters.  * It must begin with a letter. Subsequent characters can be letters, underscores, or digits (0 to 9).  * It can't be a word reserved by the database engine.",
						MarkdownDescription: "The meaning of this parameter differs according to the database engine you use.  MySQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  MariaDB  The name of the database to create when the DB instance is created. If this parameter isn't specified, no database is created in the DB instance.  Constraints:  * Must contain 1 to 64 letters or numbers.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  PostgreSQL  The name of the database to create when the DB instance is created. If this parameter isn't specified, a database named postgres is created in the DB instance.  Constraints:  * Must contain 1 to 63 letters, numbers, or underscores.  * Must begin with a letter. Subsequent characters can be letters, underscores, or digits (0-9).  * Can't be a word reserved by the specified database engine  Oracle  The Oracle System ID (SID) of the created DB instance. If you specify null, the default value ORCL is used. You can't specify the string NULL, or any other reserved word, for DBName.  Default: ORCL  Constraints:  * Can't be longer than 8 characters  Amazon RDS Custom for Oracle  The Oracle System ID (SID) of the created RDS Custom DB instance. If you don't specify a value, the default value is ORCL.  Default: ORCL  Constraints:  * It must contain 1 to 8 alphanumeric characters.  * It must contain a letter.  * It can't be a word reserved by the database engine.  Amazon RDS Custom for SQL Server  Not applicable. Must be null.  SQL Server  Not applicable. Must be null.  Amazon Aurora MySQL  The name of the database to create when the primary DB instance of the Aurora MySQL DB cluster is created. If this parameter isn't specified for an Aurora MySQL DB cluster, no database is created in the DB cluster.  Constraints:  * It must contain 1 to 64 alphanumeric characters.  * It can't be a word reserved by the database engine.  Amazon Aurora PostgreSQL  The name of the database to create when the primary DB instance of the Aurora PostgreSQL DB cluster is created. If this parameter isn't specified for an Aurora PostgreSQL DB cluster, a database named postgres is created in the DB cluster.  Constraints:  * It must contain 1 to 63 alphanumeric characters.  * It must begin with a letter. Subsequent characters can be letters, underscores, or digits (0 to 9).  * It can't be a word reserved by the database engine.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_parameter_group_name": schema.StringAttribute{
						Description:         "The name of the DB parameter group to associate with this DB instance. If you do not specify a value, then the default DB parameter group for the specified DB engine and version is used.  This setting doesn't apply to RDS Custom.  Constraints:  * It must be 1 to 255 letters, numbers, or hyphens.  * The first character must be a letter.  * It can't end with a hyphen or contain two consecutive hyphens.",
						MarkdownDescription: "The name of the DB parameter group to associate with this DB instance. If you do not specify a value, then the default DB parameter group for the specified DB engine and version is used.  This setting doesn't apply to RDS Custom.  Constraints:  * It must be 1 to 255 letters, numbers, or hyphens.  * The first character must be a letter.  * It can't end with a hyphen or contain two consecutive hyphens.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_parameter_group_ref": schema.SingleNestedAttribute{
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
						Description:         "The identifier for the DB snapshot to restore from.  Constraints:  * Must match the identifier of an existing DBSnapshot.  * Can't be specified when DBClusterSnapshotIdentifier is specified.  * Must be specified when DBClusterSnapshotIdentifier isn't specified.  * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB snapshot.",
						MarkdownDescription: "The identifier for the DB snapshot to restore from.  Constraints:  * Must match the identifier of an existing DBSnapshot.  * Can't be specified when DBClusterSnapshotIdentifier is specified.  * Must be specified when DBClusterSnapshotIdentifier isn't specified.  * If you are restoring from a shared manual DB snapshot, the DBSnapshotIdentifier must be the ARN of the shared DB snapshot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_subnet_group_name": schema.StringAttribute{
						Description:         "A DB subnet group to associate with this DB instance.  Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.  Example: mydbsubnetgroup",
						MarkdownDescription: "A DB subnet group to associate with this DB instance.  Constraints: Must match the name of an existing DBSubnetGroup. Must not be default.  Example: mydbsubnetgroup",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
						Description:         "A value that indicates whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).  Amazon Aurora  Not applicable. You can enable or disable deletion protection for the DB cluster. For more information, see CreateDBCluster. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.",
						MarkdownDescription: "A value that indicates whether the DB instance has deletion protection enabled. The database can't be deleted when deletion protection is enabled. By default, deletion protection isn't enabled. For more information, see Deleting a DB Instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_DeleteInstance.html).  Amazon Aurora  Not applicable. You can enable or disable deletion protection for the DB cluster. For more information, see CreateDBCluster. DB instances in a DB cluster can be deleted even when deletion protection is enabled for the DB cluster.",
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
						Description:         "The Active Directory directory ID to create the DB instance in. Currently, only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.  For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",
						MarkdownDescription: "The Active Directory directory ID to create the DB instance in. Currently, only MySQL, Microsoft SQL Server, Oracle, and PostgreSQL DB instances can be created in an Active Directory Domain.  For more information, see Kerberos Authentication (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/kerberos-authentication.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"domain_iam_role_name": schema.StringAttribute{
						Description:         "Specify the name of the IAM role to be used when making API calls to the Directory Service.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",
						MarkdownDescription: "Specify the name of the IAM role to be used when making API calls to the Directory Service.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. The domain is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_cloudwatch_logs_exports": schema.ListAttribute{
						Description:         "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine. For more information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. CloudWatch Logs exports are managed by the DB cluster.  RDS Custom  Not applicable.  MariaDB  Possible values are audit, error, general, and slowquery.  Microsoft SQL Server  Possible values are agent and error.  MySQL  Possible values are audit, error, general, and slowquery.  Oracle  Possible values are alert, audit, listener, trace, and oemagent.  PostgreSQL  Possible values are postgresql and upgrade.",
						MarkdownDescription: "The list of log types that need to be enabled for exporting to CloudWatch Logs. The values in the list depend on the DB engine. For more information, see Publishing Database Logs to Amazon CloudWatch Logs (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.html#USER_LogAccess.Procedural.UploadtoCloudWatch) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. CloudWatch Logs exports are managed by the DB cluster.  RDS Custom  Not applicable.  MariaDB  Possible values are audit, error, general, and slowquery.  Microsoft SQL Server  Possible values are agent and error.  MySQL  Possible values are audit, error, general, and slowquery.  Oracle  Possible values are alert, audit, listener, trace, and oemagent.  PostgreSQL  Possible values are postgresql and upgrade.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_customer_owned_ip": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance.  A CoIP provides local or external connectivity to resources in your Outpost subnets through your on-premises network. For some use cases, a CoIP can provide lower latency for connections to the DB instance from outside of its virtual private cloud (VPC) on your local network.  For more information about RDS on Outposts, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.  For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing) in the Amazon Web Services Outposts User Guide.",
						MarkdownDescription: "A value that indicates whether to enable a customer-owned IP address (CoIP) for an RDS on Outposts DB instance.  A CoIP provides local or external connectivity to resources in your Outpost subnets through your on-premises network. For some use cases, a CoIP can provide lower latency for connections to the DB instance from outside of its virtual private cloud (VPC) on your local network.  For more information about RDS on Outposts, see Working with Amazon RDS on Amazon Web Services Outposts (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-on-outposts.html) in the Amazon RDS User Guide.  For more information about CoIPs, see Customer-owned IP addresses (https://docs.aws.amazon.com/outposts/latest/userguide/routing.html#ip-addressing) in the Amazon Web Services Outposts User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_iam_database_authentication": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled.  For more information, see IAM Database Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Mapping Amazon Web Services IAM accounts to database accounts is managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether to enable mapping of Amazon Web Services Identity and Access Management (IAM) accounts to database accounts. By default, mapping isn't enabled.  For more information, see IAM Database Authentication for MySQL and PostgreSQL (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.IAMDBAuth.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Mapping Amazon Web Services IAM accounts to database accounts is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine": schema.StringAttribute{
						Description:         "The name of the database engine to be used for this instance.  Not every database engine is available for every Amazon Web Services Region.  Valid Values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * custom-oracle-ee (for RDS Custom for Oracle instances)  * custom-sqlserver-ee (for RDS Custom for SQL Server instances)  * custom-sqlserver-se (for RDS Custom for SQL Server instances)  * custom-sqlserver-web (for RDS Custom for SQL Server instances)  * mariadb  * mysql  * oracle-ee  * oracle-ee-cdb  * oracle-se2  * oracle-se2-cdb  * postgres  * sqlserver-ee  * sqlserver-se  * sqlserver-ex  * sqlserver-web",
						MarkdownDescription: "The name of the database engine to be used for this instance.  Not every database engine is available for every Amazon Web Services Region.  Valid Values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * custom-oracle-ee (for RDS Custom for Oracle instances)  * custom-sqlserver-ee (for RDS Custom for SQL Server instances)  * custom-sqlserver-se (for RDS Custom for SQL Server instances)  * custom-sqlserver-web (for RDS Custom for SQL Server instances)  * mariadb  * mysql  * oracle-ee  * oracle-ee-cdb  * oracle-se2  * oracle-se2-cdb  * postgres  * sqlserver-ee  * sqlserver-se  * sqlserver-ex  * sqlserver-web",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The version number of the database engine to use.  For a list of valid engine versions, use the DescribeDBEngineVersions operation.  The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every Amazon Web Services Region.  Amazon Aurora  Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.  Amazon RDS Custom for Oracle  A custom engine version (CEV) that you have previously created. This setting is required for RDS Custom for Oracle. The CEV name has the following format: 19.customized_string. A valid CEV name is 19.my_cev1. For more information, see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create) in the Amazon RDS User Guide.  Amazon RDS Custom for SQL Server  See RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html) in the Amazon RDS User Guide.  MariaDB  For information, see MariaDB on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Microsoft SQL Server  For information, see Microsoft SQL Server Versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the Amazon RDS User Guide.  MySQL  For information, see MySQL on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Oracle  For information, see Oracle Database Engine Release Notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the Amazon RDS User Guide.  PostgreSQL  For information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.",
						MarkdownDescription: "The version number of the database engine to use.  For a list of valid engine versions, use the DescribeDBEngineVersions operation.  The following are the database engines and links to information about the major and minor versions that are available with Amazon RDS. Not every database engine is available for every Amazon Web Services Region.  Amazon Aurora  Not applicable. The version number of the database engine to be used by the DB instance is managed by the DB cluster.  Amazon RDS Custom for Oracle  A custom engine version (CEV) that you have previously created. This setting is required for RDS Custom for Oracle. The CEV name has the following format: 19.customized_string. A valid CEV name is 19.my_cev1. For more information, see Creating an RDS Custom for Oracle DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-creating.html#custom-creating.create) in the Amazon RDS User Guide.  Amazon RDS Custom for SQL Server  See RDS Custom for SQL Server general requirements (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/custom-reqs-limits-MS.html) in the Amazon RDS User Guide.  MariaDB  For information, see MariaDB on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MariaDB.html#MariaDB.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Microsoft SQL Server  For information, see Microsoft SQL Server Versions on Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.VersionSupport) in the Amazon RDS User Guide.  MySQL  For information, see MySQL on Amazon RDS Versions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_MySQL.html#MySQL.Concepts.VersionMgmt) in the Amazon RDS User Guide.  Oracle  For information, see Oracle Database Engine Release Notes (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.Oracle.PatchComposition.html) in the Amazon RDS User Guide.  PostgreSQL  For information, see Amazon RDS for PostgreSQL versions and extensions (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_PostgreSQL.html#PostgreSQL.Concepts) in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"iops": schema.Int64Attribute{
						Description:         "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for the DB instance. For information about valid IOPS values, see Amazon RDS DB instance storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html) in the Amazon RDS User Guide.  Constraints: For MariaDB, MySQL, Oracle, and PostgreSQL DB instances, must be a multiple between .5 and 50 of the storage amount for the DB instance. For SQL Server DB instances, must be a multiple between 1 and 50 of the storage amount for the DB instance.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "The amount of Provisioned IOPS (input/output operations per second) to be initially allocated for the DB instance. For information about valid IOPS values, see Amazon RDS DB instance storage (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html) in the Amazon RDS User Guide.  Constraints: For MariaDB, MySQL, Oracle, and PostgreSQL DB instances, must be a multiple between .5 and 50 of the storage amount for the DB instance. For SQL Server DB instances, must be a multiple between 1 and 50 of the storage amount for the DB instance.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for an encrypted DB instance.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  Amazon Aurora  Not applicable. The Amazon Web Services KMS key identifier is managed by the DB cluster. For more information, see CreateDBCluster.  If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Amazon RDS Custom  A KMS key is required for RDS Custom instances. For most RDS engines, if you leave this parameter empty while enabling StorageEncrypted, the engine uses the default KMS key. However, RDS Custom doesn't use the default key when this parameter is empty. You must explicitly specify a key.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for an encrypted DB instance.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  Amazon Aurora  Not applicable. The Amazon Web Services KMS key identifier is managed by the DB cluster. For more information, see CreateDBCluster.  If StorageEncrypted is enabled, and you do not specify a value for the KmsKeyId parameter, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  Amazon RDS Custom  A KMS key is required for RDS Custom instances. For most RDS engines, if you leave this parameter empty while enabling StorageEncrypted, the engine uses the default KMS key. However, RDS Custom doesn't use the default key when this parameter is empty. You must explicitly specify a key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
						Description:         "License model information for this DB instance.  Valid values: license-included | bring-your-own-license | general-public-license  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "License model information for this DB instance.  Valid values: license-included | bring-your-own-license | general-public-license  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"manage_master_user_password": schema.BoolAttribute{
						Description:         "A value that indicates whether to manage the master user password with Amazon Web Services Secrets Manager.  For more information, see Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the Amazon RDS User Guide.  Constraints:  * Can't manage the master user password with Amazon Web Services Secrets Manager if MasterUserPassword is specified.",
						MarkdownDescription: "A value that indicates whether to manage the master user password with Amazon Web Services Secrets Manager.  For more information, see Password management with Amazon Web Services Secrets Manager (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-secrets-manager.html) in the Amazon RDS User Guide.  Constraints:  * Can't manage the master user password with Amazon Web Services Secrets Manager if MasterUserPassword is specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"master_user_password": schema.SingleNestedAttribute{
						Description:         "The password for the master user. The password can include any printable ASCII character except '/', ''', or '@'.  Amazon Aurora  Not applicable. The password for the master user is managed by the DB cluster.  Constraints: Can't be specified if ManageMasterUserPassword is turned on.  MariaDB  Constraints: Must contain from 8 to 41 characters.  Microsoft SQL Server  Constraints: Must contain from 8 to 128 characters.  MySQL  Constraints: Must contain from 8 to 41 characters.  Oracle  Constraints: Must contain from 8 to 30 characters.  PostgreSQL  Constraints: Must contain from 8 to 128 characters.",
						MarkdownDescription: "The password for the master user. The password can include any printable ASCII character except '/', ''', or '@'.  Amazon Aurora  Not applicable. The password for the master user is managed by the DB cluster.  Constraints: Can't be specified if ManageMasterUserPassword is turned on.  MariaDB  Constraints: Must contain from 8 to 41 characters.  Microsoft SQL Server  Constraints: Must contain from 8 to 128 characters.  MySQL  Constraints: Must contain from 8 to 41 characters.  Oracle  Constraints: Must contain from 8 to 30 characters.  PostgreSQL  Constraints: Must contain from 8 to 128 characters.",
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
						Description:         "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.  This setting is valid only if the master user password is managed by RDS in Amazon Web Services Secrets Manager for the DB instance.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a different Amazon Web Services account, then you can't use the aws/secretsmanager KMS key to encrypt the secret, and you must use a customer managed KMS key.  There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier to encrypt a secret that is automatically generated and managed in Amazon Web Services Secrets Manager.  This setting is valid only if the master user password is managed by RDS in Amazon Web Services Secrets Manager for the DB instance.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key. To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.  If you don't specify MasterUserSecretKmsKeyId, then the aws/secretsmanager KMS key is used to encrypt the secret. If the secret is in a different Amazon Web Services account, then you can't use the aws/secretsmanager KMS key to encrypt the secret, and you must use a customer managed KMS key.  There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
						Description:         "The name for the master user.  Amazon Aurora  Not applicable. The name for the master user is managed by the DB cluster.  Amazon RDS  Constraints:  * Required.  * Must be 1 to 16 letters, numbers, or underscores.  * First character must be a letter.  * Can't be a reserved word for the chosen database engine.",
						MarkdownDescription: "The name for the master user.  Amazon Aurora  Not applicable. The name for the master user is managed by the DB cluster.  Amazon RDS  Constraints:  * Required.  * Must be 1 to 16 letters, numbers, or underscores.  * First character must be a letter.  * Can't be a reserved word for the chosen database engine.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_allocated_storage": schema.Int64Attribute{
						Description:         "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.  For more information about this setting, including limitations that apply to it, see Managing capacity automatically with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "The upper limit in gibibytes (GiB) to which Amazon RDS can automatically scale the storage of the DB instance.  For more information about this setting, including limitations that apply to it, see Managing capacity automatically with Amazon RDS storage autoscaling (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PIOPS.StorageTypes.html#USER_PIOPS.Autoscaling) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitoring_interval": schema.Int64Attribute{
						Description:         "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify 0. The default is 0.  If MonitoringRoleArn is specified, then you must set MonitoringInterval to a value other than 0.  This setting doesn't apply to RDS Custom.  Valid Values: 0, 1, 5, 10, 15, 30, 60",
						MarkdownDescription: "The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collection of Enhanced Monitoring metrics, specify 0. The default is 0.  If MonitoringRoleArn is specified, then you must set MonitoringInterval to a value other than 0.  This setting doesn't apply to RDS Custom.  Valid Values: 0, 1, 5, 10, 15, 30, 60",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitoring_role_arn": schema.StringAttribute{
						Description:         "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting Up and Enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide.  If MonitoringInterval is set to a value other than 0, then you must supply a MonitoringRoleArn value.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to Amazon CloudWatch Logs. For example, arn:aws:iam:123456789012:role/emaccess. For information on creating a monitoring role, see Setting Up and Enabling Enhanced Monitoring (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html#USER_Monitoring.OS.Enabling) in the Amazon RDS User Guide.  If MonitoringInterval is set to a value other than 0, then you must supply a MonitoringRoleArn value.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"multi_az": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB instance is a Multi-AZ deployment. You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ deployment.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. DB instance Availability Zones (AZs) are managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether the DB instance is a Multi-AZ deployment. You can't set the AvailabilityZone parameter if the DB instance is a Multi-AZ deployment.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable. DB instance Availability Zones (AZs) are managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nchar_character_set_name": schema.StringAttribute{
						Description:         "The name of the NCHAR character set for the Oracle DB instance.  This parameter doesn't apply to RDS Custom.",
						MarkdownDescription: "The name of the NCHAR character set for the Oracle DB instance.  This parameter doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network_type": schema.StringAttribute{
						Description:         "The network type of the DB instance.  Valid values:  * IPV4  * DUAL  The network type is determined by the DBSubnetGroup specified for the DB instance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL).  For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon RDS User Guide.",
						MarkdownDescription: "The network type of the DB instance.  Valid values:  * IPV4  * DUAL  The network type is determined by the DBSubnetGroup specified for the DB instance. A DBSubnetGroup can support only the IPv4 protocol or the IPv4 and the IPv6 protocols (DUAL).  For more information, see Working with a DB instance in a VPC (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html) in the Amazon RDS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"option_group_name": schema.StringAttribute{
						Description:         "A value that indicates that the DB instance should be associated with the specified option group.  Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance after it is associated with a DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "A value that indicates that the DB instance should be associated with the specified option group.  Permanent options, such as the TDE option for Oracle Advanced Security TDE, can't be removed from an option group. Also, that option group can't be removed from a DB instance after it is associated with a DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_enabled": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable Performance Insights for the DB instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "A value that indicates whether to enable Performance Insights for the DB instance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PerfInsights.html) in the Amazon RDS User Guide.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.  If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The Amazon Web Services KMS key identifier for encryption of Performance Insights data.  The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.  If you do not specify a value for PerformanceInsightsKMSKeyId, then Amazon RDS uses your default KMS key. There is a default KMS key for your Amazon Web Services account. Your Amazon Web Services account has a different default KMS key for each Amazon Web Services Region.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_retention_period": schema.Int64Attribute{
						Description:         "The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:  * 7  * month * 31, where month is a number of months from 1-23  * 731  For example, the following values are valid:  * 93 (3 months * 31)  * 341 (11 months * 31)  * 589 (19 months * 31)  * 731  If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The number of days to retain Performance Insights data. The default is 7 days. The following values are valid:  * 7  * month * 31, where month is a number of months from 1-23  * 731  For example, the following values are valid:  * 93 (3 months * 31)  * 341 (11 months * 31)  * 589 (19 months * 31)  * 731  If you specify a retention period such as 94, which isn't a valid value, RDS issues an error.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"port": schema.Int64Attribute{
						Description:         "The port number on which the database accepts connections.  MySQL  Default: 3306  Valid values: 1150-65535  Type: Integer  MariaDB  Default: 3306  Valid values: 1150-65535  Type: Integer  PostgreSQL  Default: 5432  Valid values: 1150-65535  Type: Integer  Oracle  Default: 1521  Valid values: 1150-65535  SQL Server  Default: 1433  Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and 49152-49156.  Amazon Aurora  Default: 3306  Valid values: 1150-65535  Type: Integer",
						MarkdownDescription: "The port number on which the database accepts connections.  MySQL  Default: 3306  Valid values: 1150-65535  Type: Integer  MariaDB  Default: 3306  Valid values: 1150-65535  Type: Integer  PostgreSQL  Default: 5432  Valid values: 1150-65535  Type: Integer  Oracle  Default: 1521  Valid values: 1150-65535  SQL Server  Default: 1433  Valid values: 1150-65535 except 1234, 1434, 3260, 3343, 3389, 47001, and 49152-49156.  Amazon Aurora  Default: 3306  Valid values: 1150-65535  Type: Integer",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pre_signed_url": schema.StringAttribute{
						Description:         "When you are creating a read replica from one Amazon Web Services GovCloud (US) Region to another or from one China Amazon Web Services Region to another, the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplica API operation in the source Amazon Web Services Region that contains the source DB instance.  This setting applies only to Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. It's ignored in other Amazon Web Services Regions.  This setting applies only when replicating from a source DB instance. Source DB clusters aren't supported in Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions.  You must specify this parameter when you create an encrypted read replica from another Amazon Web Services Region by using the Amazon RDS API. Don't specify PreSignedUrl when you are creating an encrypted read replica in the same Amazon Web Services Region.  The presigned URL must be a valid request for the CreateDBInstanceReadReplica API operation that can run in the source Amazon Web Services Region that contains the encrypted source DB instance. The presigned URL request must contain the following parameter values:  * DestinationRegion - The Amazon Web Services Region that the encrypted read replica is created in. This Amazon Web Services Region is the same one where the CreateDBInstanceReadReplica operation is called that contains this presigned URL. For example, if you create an encrypted DB instance in the us-west-1 Amazon Web Services Region, from a source DB instance in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica operation in the us-east-1 Amazon Web Services Region and provide a presigned URL that contains a call to the CreateDBInstanceReadReplica operation in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion in the presigned URL must be set to the us-east-1 Amazon Web Services Region.  * KmsKeyId - The KMS key identifier for the key to use to encrypt the read replica in the destination Amazon Web Services Region. This is the same identifier for both the CreateDBInstanceReadReplica operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL.  * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted DB instance to be replicated. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are creating an encrypted read replica from a DB instance in the us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115.  To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).  If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region.  SourceRegion isn't supported for SQL Server, because Amazon RDS for SQL Server doesn't support cross-Region read replicas.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "When you are creating a read replica from one Amazon Web Services GovCloud (US) Region to another or from one China Amazon Web Services Region to another, the URL that contains a Signature Version 4 signed request for the CreateDBInstanceReadReplica API operation in the source Amazon Web Services Region that contains the source DB instance.  This setting applies only to Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions. It's ignored in other Amazon Web Services Regions.  This setting applies only when replicating from a source DB instance. Source DB clusters aren't supported in Amazon Web Services GovCloud (US) Regions and China Amazon Web Services Regions.  You must specify this parameter when you create an encrypted read replica from another Amazon Web Services Region by using the Amazon RDS API. Don't specify PreSignedUrl when you are creating an encrypted read replica in the same Amazon Web Services Region.  The presigned URL must be a valid request for the CreateDBInstanceReadReplica API operation that can run in the source Amazon Web Services Region that contains the encrypted source DB instance. The presigned URL request must contain the following parameter values:  * DestinationRegion - The Amazon Web Services Region that the encrypted read replica is created in. This Amazon Web Services Region is the same one where the CreateDBInstanceReadReplica operation is called that contains this presigned URL. For example, if you create an encrypted DB instance in the us-west-1 Amazon Web Services Region, from a source DB instance in the us-east-2 Amazon Web Services Region, then you call the CreateDBInstanceReadReplica operation in the us-east-1 Amazon Web Services Region and provide a presigned URL that contains a call to the CreateDBInstanceReadReplica operation in the us-west-2 Amazon Web Services Region. For this example, the DestinationRegion in the presigned URL must be set to the us-east-1 Amazon Web Services Region.  * KmsKeyId - The KMS key identifier for the key to use to encrypt the read replica in the destination Amazon Web Services Region. This is the same identifier for both the CreateDBInstanceReadReplica operation that is called in the destination Amazon Web Services Region, and the operation contained in the presigned URL.  * SourceDBInstanceIdentifier - The DB instance identifier for the encrypted DB instance to be replicated. This identifier must be in the Amazon Resource Name (ARN) format for the source Amazon Web Services Region. For example, if you are creating an encrypted read replica from a DB instance in the us-west-2 Amazon Web Services Region, then your SourceDBInstanceIdentifier looks like the following example: arn:aws:rds:us-west-2:123456789012:instance:mysql-instance1-20161115.  To learn how to generate a Signature Version 4 signed request, see Authenticating Requests: Using Query Parameters (Amazon Web Services Signature Version 4) (https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html) and Signature Version 4 Signing Process (https://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).  If you are using an Amazon Web Services SDK tool or the CLI, you can specify SourceRegion (or --source-region for the CLI) instead of specifying PreSignedUrl manually. Specifying SourceRegion autogenerates a presigned URL that is a valid request for the operation that can run in the source Amazon Web Services Region.  SourceRegion isn't supported for SQL Server, because Amazon RDS for SQL Server doesn't support cross-Region read replicas.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_backup_window": schema.StringAttribute{
						Description:         "The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. For more information, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. The daily time range for creating automated backups is managed by the DB cluster.  Constraints:  * Must be in the format hh24:mi-hh24:mi.  * Must be in Universal Coordinated Time (UTC).  * Must not conflict with the preferred maintenance window.  * Must be at least 30 minutes.",
						MarkdownDescription: "The daily time range during which automated backups are created if automated backups are enabled, using the BackupRetentionPeriod parameter. The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region. For more information, see Backup window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_WorkingWithAutomatedBackups.html#USER_WorkingWithAutomatedBackups.BackupWindow) in the Amazon RDS User Guide.  Amazon Aurora  Not applicable. The daily time range for creating automated backups is managed by the DB cluster.  Constraints:  * Must be in the format hh24:mi-hh24:mi.  * Must be in Universal Coordinated Time (UTC).  * Must not conflict with the preferred maintenance window.  * Must be at least 30 minutes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_maintenance_window": schema.StringAttribute{
						Description:         "The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC). For more information, see Amazon RDS Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).  Format: ddd:hh24:mi-ddd:hh24:mi  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week.  Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.  Constraints: Minimum 30-minute window.",
						MarkdownDescription: "The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC). For more information, see Amazon RDS Maintenance Window (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_UpgradeDBInstance.Maintenance.html#Concepts.DBMaintenance).  Format: ddd:hh24:mi-ddd:hh24:mi  The default is a 30-minute window selected at random from an 8-hour block of time for each Amazon Web Services Region, occurring on a random day of the week.  Valid Days: Mon, Tue, Wed, Thu, Fri, Sat, Sun.  Constraints: Minimum 30-minute window.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"processor_features": schema.ListNestedAttribute{
						Description:         "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "The number of CPU cores and the number of threads per core for the DB instance class of the DB instance.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
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
						Description:         "A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the Amazon Aurora User Guide.  This setting doesn't apply to RDS Custom.  Default: 1  Valid Values: 0 - 15",
						MarkdownDescription: "A value that specifies the order in which an Aurora Replica is promoted to the primary instance after a failure of the existing primary instance. For more information, see Fault Tolerance for an Aurora DB Cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Backups.html#Aurora.Managing.FaultTolerance) in the Amazon Aurora User Guide.  This setting doesn't apply to RDS Custom.  Default: 1  Valid Values: 0 - 15",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"publicly_accessible": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB instance is publicly accessible.  When the DB instance is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB instance's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB instance's VPC. Access to the DB instance is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB instance doesn't permit it.  When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a private IP address.  Default: The default behavior varies depending on whether DBSubnetGroupName is specified.  If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies:  * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB instance is private.  * If the default VPC in the target Region has an internet gateway attached to it, the DB instance is public.  If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies:  * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB instance is private.  * If the subnets are part of a VPC that has an internet gateway attached to it, the DB instance is public.",
						MarkdownDescription: "A value that indicates whether the DB instance is publicly accessible.  When the DB instance is publicly accessible, its Domain Name System (DNS) endpoint resolves to the private IP address from within the DB instance's virtual private cloud (VPC). It resolves to the public IP address from outside of the DB instance's VPC. Access to the DB instance is ultimately controlled by the security group it uses. That public access is not permitted if the security group assigned to the DB instance doesn't permit it.  When the DB instance isn't publicly accessible, it is an internal DB instance with a DNS name that resolves to a private IP address.  Default: The default behavior varies depending on whether DBSubnetGroupName is specified.  If DBSubnetGroupName isn't specified, and PubliclyAccessible isn't specified, the following applies:  * If the default VPC in the target Region doesn’t have an internet gateway attached to it, the DB instance is private.  * If the default VPC in the target Region has an internet gateway attached to it, the DB instance is public.  If DBSubnetGroupName is specified, and PubliclyAccessible isn't specified, the following applies:  * If the subnets are part of a VPC that doesn’t have an internet gateway attached to it, the DB instance is private.  * If the subnets are part of a VPC that has an internet gateway attached to it, the DB instance is public.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_mode": schema.StringAttribute{
						Description:         "The open mode of the replica database: mounted or read-only.  This parameter is only supported for Oracle DB instances.  Mounted DB replicas are included in Oracle Database Enterprise Edition. The main use case for mounted replicas is cross-Region disaster recovery. The primary database doesn't use Active Data Guard to transmit information to the mounted replica. Because it doesn't accept user connections, a mounted replica can't serve a read-only workload.  You can create a combination of mounted and read-only DB replicas for the same primary DB instance. For more information, see Working with Oracle Read Replicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide.  For RDS Custom, you must specify this parameter and set it to mounted. The value won't be set by default. After replica creation, you can manage the open mode manually.",
						MarkdownDescription: "The open mode of the replica database: mounted or read-only.  This parameter is only supported for Oracle DB instances.  Mounted DB replicas are included in Oracle Database Enterprise Edition. The main use case for mounted replicas is cross-Region disaster recovery. The primary database doesn't use Active Data Guard to transmit information to the mounted replica. Because it doesn't accept user connections, a mounted replica can't serve a read-only workload.  You can create a combination of mounted and read-only DB replicas for the same primary DB instance. For more information, see Working with Oracle Read Replicas for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.html) in the Amazon RDS User Guide.  For RDS Custom, you must specify this parameter and set it to mounted. The value won't be set by default. After replica creation, you can manage the open mode manually.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_db_instance_identifier": schema.StringAttribute{
						Description:         "The identifier of the DB instance that will act as the source for the read replica. Each DB instance can have up to 15 read replicas, with the exception of Oracle and SQL Server, which can have up to five.  Constraints:  * Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL, or SQL Server DB instance.  * Can't be specified if the SourceDBClusterIdentifier parameter is also specified.  * For the limitations of Oracle read replicas, see Version and licensing considerations for RDS for Oracle replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.limitations.html#oracle-read-replicas.limitations.versions-and-licenses) in the Amazon RDS User Guide.  * For the limitations of SQL Server read replicas, see Read replica limitations with SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.html#SQLServer.ReadReplicas.Limitations) in the Amazon RDS User Guide.  * The specified DB instance must have automatic backups enabled, that is, its backup retention period must be greater than 0.  * If the source DB instance is in the same Amazon Web Services Region as the read replica, specify a valid DB instance identifier.  * If the source DB instance is in a different Amazon Web Services Region from the read replica, specify a valid DB instance ARN. For more information, see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing) in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS Custom, which don't support cross-Region replicas.",
						MarkdownDescription: "The identifier of the DB instance that will act as the source for the read replica. Each DB instance can have up to 15 read replicas, with the exception of Oracle and SQL Server, which can have up to five.  Constraints:  * Must be the identifier of an existing MySQL, MariaDB, Oracle, PostgreSQL, or SQL Server DB instance.  * Can't be specified if the SourceDBClusterIdentifier parameter is also specified.  * For the limitations of Oracle read replicas, see Version and licensing considerations for RDS for Oracle replicas (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-read-replicas.limitations.html#oracle-read-replicas.limitations.versions-and-licenses) in the Amazon RDS User Guide.  * For the limitations of SQL Server read replicas, see Read replica limitations with SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/SQLServer.ReadReplicas.html#SQLServer.ReadReplicas.Limitations) in the Amazon RDS User Guide.  * The specified DB instance must have automatic backups enabled, that is, its backup retention period must be greater than 0.  * If the source DB instance is in the same Amazon Web Services Region as the read replica, specify a valid DB instance identifier.  * If the source DB instance is in a different Amazon Web Services Region from the read replica, specify a valid DB instance ARN. For more information, see Constructing an ARN for Amazon RDS (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing) in the Amazon RDS User Guide. This doesn't apply to SQL Server or RDS Custom, which don't support cross-Region replicas.",
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
						Description:         "A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.  For RDS Custom instances, either set this parameter to true or leave it unset. If you set this parameter to false, RDS reports an error.  Amazon Aurora  Not applicable. The encryption for DB instances is managed by the DB cluster.",
						MarkdownDescription: "A value that indicates whether the DB instance is encrypted. By default, it isn't encrypted.  For RDS Custom instances, either set this parameter to true or leave it unset. If you set this parameter to false, RDS reports an error.  Amazon Aurora  Not applicable. The encryption for DB instances is managed by the DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_throughput": schema.Int64Attribute{
						Description:         "Specifies the storage throughput value for the DB instance.  This setting applies only to the gp3 storage type.  This setting doesn't apply to RDS Custom or Amazon Aurora.",
						MarkdownDescription: "Specifies the storage throughput value for the DB instance.  This setting applies only to the gp3 storage type.  This setting doesn't apply to RDS Custom or Amazon Aurora.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_type": schema.StringAttribute{
						Description:         "Specifies the storage type to be associated with the DB instance.  Valid values: gp2 | gp3 | io1 | standard  If you specify io1 or gp3, you must also include a value for the Iops parameter.  Default: io1 if the Iops parameter is specified, otherwise gp2  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
						MarkdownDescription: "Specifies the storage type to be associated with the DB instance.  Valid values: gp2 | gp3 | io1 | standard  If you specify io1 or gp3, you must also include a value for the Iops parameter.  Default: io1 if the Iops parameter is specified, otherwise gp2  Amazon Aurora  Not applicable. Storage is managed by the DB cluster.",
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
						Description:         "The ARN from the key store with which to associate the instance for TDE encryption.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						MarkdownDescription: "The ARN from the key store with which to associate the instance for TDE encryption.  This setting doesn't apply to RDS Custom.  Amazon Aurora  Not applicable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tde_credential_password": schema.StringAttribute{
						Description:         "The password for the given ARN from the key store in order to access the device.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "The password for the given ARN from the key store in order to access the device.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timezone": schema.StringAttribute{
						Description:         "The time zone of the DB instance. The time zone parameter is currently supported only by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
						MarkdownDescription: "The time zone of the DB instance. The time zone parameter is currently supported only by Microsoft SQL Server (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_default_processor_features": schema.BoolAttribute{
						Description:         "A value that indicates whether the DB instance class of the DB instance uses its default processor features.  This setting doesn't apply to RDS Custom.",
						MarkdownDescription: "A value that indicates whether the DB instance class of the DB instance uses its default processor features.  This setting doesn't apply to RDS Custom.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vpc_security_group_i_ds": schema.ListAttribute{
						Description:         "A list of Amazon EC2 VPC security groups to associate with this DB instance.  Amazon Aurora  Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster.  Default: The default EC2 VPC security group for the DB subnet group's VPC.",
						MarkdownDescription: "A list of Amazon EC2 VPC security groups to associate with this DB instance.  Amazon Aurora  Not applicable. The associated list of EC2 VPC security groups is managed by the DB cluster.  Default: The default EC2 VPC security group for the DB subnet group's VPC.",
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

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")

	var model RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBInstance")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "rds.services.k8s.aws", Version: "v1alpha1", Resource: "dbinstances"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")

	var data RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "rds.services.k8s.aws", Version: "v1alpha1", Resource: "dbinstances"}).
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

	var readResponse RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")

	var model RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBInstance")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "rds.services.k8s.aws", Version: "v1alpha1", Resource: "dbinstances"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_rds_services_k8s_aws_db_instance_v1alpha1")

	var data RdsServicesK8SAwsDbinstanceV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "rds.services.k8s.aws", Version: "v1alpha1", Resource: "dbinstances"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "rds.services.k8s.aws", Version: "v1alpha1", Resource: "dbinstances"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *RdsServicesK8SAwsDbinstanceV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
