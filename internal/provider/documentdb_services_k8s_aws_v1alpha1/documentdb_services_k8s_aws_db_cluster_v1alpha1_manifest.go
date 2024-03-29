/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package documentdb_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &DocumentdbServicesK8SAwsDbclusterV1Alpha1Manifest{}
)

func NewDocumentdbServicesK8SAwsDbclusterV1Alpha1Manifest() datasource.DataSource {
	return &DocumentdbServicesK8SAwsDbclusterV1Alpha1Manifest{}
}

type DocumentdbServicesK8SAwsDbclusterV1Alpha1Manifest struct{}

type DocumentdbServicesK8SAwsDbclusterV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		AvailabilityZones           *[]string `tfsdk:"availability_zones" json:"availabilityZones,omitempty"`
		BackupRetentionPeriod       *int64    `tfsdk:"backup_retention_period" json:"backupRetentionPeriod,omitempty"`
		DbClusterIdentifier         *string   `tfsdk:"db_cluster_identifier" json:"dbClusterIdentifier,omitempty"`
		DbClusterParameterGroupName *string   `tfsdk:"db_cluster_parameter_group_name" json:"dbClusterParameterGroupName,omitempty"`
		DbSubnetGroupName           *string   `tfsdk:"db_subnet_group_name" json:"dbSubnetGroupName,omitempty"`
		DbSubnetGroupRef            *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"db_subnet_group_ref" json:"dbSubnetGroupRef,omitempty"`
		DeletionProtection          *bool     `tfsdk:"deletion_protection" json:"deletionProtection,omitempty"`
		DestinationRegion           *string   `tfsdk:"destination_region" json:"destinationRegion,omitempty"`
		EnableCloudwatchLogsExports *[]string `tfsdk:"enable_cloudwatch_logs_exports" json:"enableCloudwatchLogsExports,omitempty"`
		Engine                      *string   `tfsdk:"engine" json:"engine,omitempty"`
		EngineVersion               *string   `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		GlobalClusterIdentifier     *string   `tfsdk:"global_cluster_identifier" json:"globalClusterIdentifier,omitempty"`
		KmsKeyID                    *string   `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		KmsKeyRef                   *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"kms_key_ref" json:"kmsKeyRef,omitempty"`
		MasterUserPassword *struct {
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"master_user_password" json:"masterUserPassword,omitempty"`
		MasterUsername             *string `tfsdk:"master_username" json:"masterUsername,omitempty"`
		Port                       *int64  `tfsdk:"port" json:"port,omitempty"`
		PreSignedURL               *string `tfsdk:"pre_signed_url" json:"preSignedURL,omitempty"`
		PreferredBackupWindow      *string `tfsdk:"preferred_backup_window" json:"preferredBackupWindow,omitempty"`
		PreferredMaintenanceWindow *string `tfsdk:"preferred_maintenance_window" json:"preferredMaintenanceWindow,omitempty"`
		SnapshotIdentifier         *string `tfsdk:"snapshot_identifier" json:"snapshotIdentifier,omitempty"`
		SourceRegion               *string `tfsdk:"source_region" json:"sourceRegion,omitempty"`
		StorageEncrypted           *bool   `tfsdk:"storage_encrypted" json:"storageEncrypted,omitempty"`
		StorageType                *string `tfsdk:"storage_type" json:"storageType,omitempty"`
		Tags                       *[]struct {
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

func (r *DocumentdbServicesK8SAwsDbclusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest"
}

func (r *DocumentdbServicesK8SAwsDbclusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "DBClusterSpec defines the desired state of DBCluster.Detailed information about a cluster.",
				MarkdownDescription: "DBClusterSpec defines the desired state of DBCluster.Detailed information about a cluster.",
				Attributes: map[string]schema.Attribute{
					"availability_zones": schema.ListAttribute{
						Description:         "A list of Amazon EC2 Availability Zones that instances in the cluster canbe created in.",
						MarkdownDescription: "A list of Amazon EC2 Availability Zones that instances in the cluster canbe created in.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_retention_period": schema.Int64Attribute{
						Description:         "The number of days for which automated backups are retained. You must specifya minimum value of 1.Default: 1Constraints:   * Must be a value from 1 to 35.",
						MarkdownDescription: "The number of days for which automated backups are retained. You must specifya minimum value of 1.Default: 1Constraints:   * Must be a value from 1 to 35.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_identifier": schema.StringAttribute{
						Description:         "The cluster identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * The first character must be a letter.   * Cannot end with a hyphen or contain two consecutive hyphens.Example: my-cluster",
						MarkdownDescription: "The cluster identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * The first character must be a letter.   * Cannot end with a hyphen or contain two consecutive hyphens.Example: my-cluster",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_cluster_parameter_group_name": schema.StringAttribute{
						Description:         "The name of the cluster parameter group to associate with this cluster.",
						MarkdownDescription: "The name of the cluster parameter group to associate with this cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_subnet_group_name": schema.StringAttribute{
						Description:         "A subnet group to associate with this cluster.Constraints: Must match the name of an existing DBSubnetGroup. Must not bedefault.Example: mySubnetgroup",
						MarkdownDescription: "A subnet group to associate with this cluster.Constraints: Must match the name of an existing DBSubnetGroup. Must not bedefault.Example: mySubnetgroup",
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
						Description:         "Specifies whether this cluster can be deleted. If DeletionProtection is enabled,the cluster cannot be deleted unless it is modified and DeletionProtectionis disabled. DeletionProtection protects clusters from being accidentallydeleted.",
						MarkdownDescription: "Specifies whether this cluster can be deleted. If DeletionProtection is enabled,the cluster cannot be deleted unless it is modified and DeletionProtectionis disabled. DeletionProtection protects clusters from being accidentallydeleted.",
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

					"enable_cloudwatch_logs_exports": schema.ListAttribute{
						Description:         "A list of log types that need to be enabled for exporting to Amazon CloudWatchLogs. You can enable audit logs or profiler logs. For more information, seeAuditing Amazon DocumentDB Events (https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html)and Profiling Amazon DocumentDB Operations (https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html).",
						MarkdownDescription: "A list of log types that need to be enabled for exporting to Amazon CloudWatchLogs. You can enable audit logs or profiler logs. For more information, seeAuditing Amazon DocumentDB Events (https://docs.aws.amazon.com/documentdb/latest/developerguide/event-auditing.html)and Profiling Amazon DocumentDB Operations (https://docs.aws.amazon.com/documentdb/latest/developerguide/profiling.html).",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine": schema.StringAttribute{
						Description:         "The name of the database engine to be used for this cluster.Valid values: docdb",
						MarkdownDescription: "The name of the database engine to be used for this cluster.Valid values: docdb",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The version number of the database engine to use. The --engine-version willdefault to the latest major engine version. For production workloads, werecommend explicitly declaring this parameter with the intended major engineversion.",
						MarkdownDescription: "The version number of the database engine to use. The --engine-version willdefault to the latest major engine version. For production workloads, werecommend explicitly declaring this parameter with the intended major engineversion.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"global_cluster_identifier": schema.StringAttribute{
						Description:         "The cluster identifier of the new global cluster.",
						MarkdownDescription: "The cluster identifier of the new global cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The KMS key identifier for an encrypted cluster.The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryptionkey. If you are creating a cluster using the same Amazon Web Services accountthat owns the KMS encryption key that is used to encrypt the new cluster,you can use the KMS key alias instead of the ARN for the KMS encryption key.If an encryption key is not specified in KmsKeyId:   * If the StorageEncrypted parameter is true, Amazon DocumentDB uses your   default encryption key.KMS creates the default encryption key for your Amazon Web Services account.Your Amazon Web Services account has a different default encryption key foreach Amazon Web Services Regions.",
						MarkdownDescription: "The KMS key identifier for an encrypted cluster.The KMS key identifier is the Amazon Resource Name (ARN) for the KMS encryptionkey. If you are creating a cluster using the same Amazon Web Services accountthat owns the KMS encryption key that is used to encrypt the new cluster,you can use the KMS key alias instead of the ARN for the KMS encryption key.If an encryption key is not specified in KmsKeyId:   * If the StorageEncrypted parameter is true, Amazon DocumentDB uses your   default encryption key.KMS creates the default encryption key for your Amazon Web Services account.Your Amazon Web Services account has a different default encryption key foreach Amazon Web Services Regions.",
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

					"master_user_password": schema.SingleNestedAttribute{
						Description:         "The password for the master database user. This password can contain anyprintable ASCII character except forward slash (/), double quote ('), orthe 'at' symbol (@).Constraints: Must contain from 8 to 100 characters.",
						MarkdownDescription: "The password for the master database user. This password can contain anyprintable ASCII character except forward slash (/), double quote ('), orthe 'at' symbol (@).Constraints: Must contain from 8 to 100 characters.",
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

					"master_username": schema.StringAttribute{
						Description:         "The name of the master user for the cluster.Constraints:   * Must be from 1 to 63 letters or numbers.   * The first character must be a letter.   * Cannot be a reserved word for the chosen database engine.",
						MarkdownDescription: "The name of the master user for the cluster.Constraints:   * Must be from 1 to 63 letters or numbers.   * The first character must be a letter.   * Cannot be a reserved word for the chosen database engine.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"port": schema.Int64Attribute{
						Description:         "The port number on which the instances in the cluster accept connections.",
						MarkdownDescription: "The port number on which the instances in the cluster accept connections.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pre_signed_url": schema.StringAttribute{
						Description:         "Not currently supported.",
						MarkdownDescription: "Not currently supported.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_backup_window": schema.StringAttribute{
						Description:         "The daily time range during which automated backups are created if automatedbackups are enabled using the BackupRetentionPeriod parameter.The default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region.Constraints:   * Must be in the format hh24:mi-hh24:mi.   * Must be in Universal Coordinated Time (UTC).   * Must not conflict with the preferred maintenance window.   * Must be at least 30 minutes.",
						MarkdownDescription: "The daily time range during which automated backups are created if automatedbackups are enabled using the BackupRetentionPeriod parameter.The default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region.Constraints:   * Must be in the format hh24:mi-hh24:mi.   * Must be in Universal Coordinated Time (UTC).   * Must not conflict with the preferred maintenance window.   * Must be at least 30 minutes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preferred_maintenance_window": schema.StringAttribute{
						Description:         "The weekly time range during which system maintenance can occur, in UniversalCoordinated Time (UTC).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid days: Mon, Tue, Wed, Thu, Fri, Sat, SunConstraints: Minimum 30-minute window.",
						MarkdownDescription: "The weekly time range during which system maintenance can occur, in UniversalCoordinated Time (UTC).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid days: Mon, Tue, Wed, Thu, Fri, Sat, SunConstraints: Minimum 30-minute window.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot_identifier": schema.StringAttribute{
						Description:         "The identifier for the snapshot or cluster snapshot to restore from.You can use either the name or the Amazon Resource Name (ARN) to specifya cluster snapshot. However, you can use only the ARN to specify a snapshot.Constraints:   * Must match the identifier of an existing snapshot.",
						MarkdownDescription: "The identifier for the snapshot or cluster snapshot to restore from.You can use either the name or the Amazon Resource Name (ARN) to specifya cluster snapshot. However, you can use only the ARN to specify a snapshot.Constraints:   * Must match the identifier of an existing snapshot.",
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
						Description:         "Specifies whether the cluster is encrypted.",
						MarkdownDescription: "Specifies whether the cluster is encrypted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_type": schema.StringAttribute{
						Description:         "The storage type to associate with the DB cluster.For information on storage types for Amazon DocumentDB clusters, see Clusterstorage configurations in the Amazon DocumentDB Developer Guide.Valid values for storage type - standard | iopt1Default value is standardWhen you create a DocumentDB DB cluster with the storage type set to iopt1,the storage type is returned in the response. The storage type isn't returnedwhen you set it to standard.",
						MarkdownDescription: "The storage type to associate with the DB cluster.For information on storage types for Amazon DocumentDB clusters, see Clusterstorage configurations in the Amazon DocumentDB Developer Guide.Valid values for storage type - standard | iopt1Default value is standardWhen you create a DocumentDB DB cluster with the storage type set to iopt1,the storage type is returned in the response. The storage type isn't returnedwhen you set it to standard.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "The tags to be assigned to the cluster.",
						MarkdownDescription: "The tags to be assigned to the cluster.",
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

					"vpc_security_group_i_ds": schema.ListAttribute{
						Description:         "A list of EC2 VPC security groups to associate with this cluster.",
						MarkdownDescription: "A list of EC2 VPC security groups to associate with this cluster.",
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

func (r *DocumentdbServicesK8SAwsDbclusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_documentdb_services_k8s_aws_db_cluster_v1alpha1_manifest")

	var model DocumentdbServicesK8SAwsDbclusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("documentdb.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
