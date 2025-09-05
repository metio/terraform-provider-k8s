/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package everest_percona_com_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &EverestPerconaComDatabaseClusterV1Alpha1Manifest{}
)

func NewEverestPerconaComDatabaseClusterV1Alpha1Manifest() datasource.DataSource {
	return &EverestPerconaComDatabaseClusterV1Alpha1Manifest{}
}

type EverestPerconaComDatabaseClusterV1Alpha1Manifest struct{}

type EverestPerconaComDatabaseClusterV1Alpha1ManifestData struct {
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
		AllowUnsafeConfiguration *bool `tfsdk:"allow_unsafe_configuration" json:"allowUnsafeConfiguration,omitempty"`
		Backup                   *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Pitr    *struct {
				BackupStorageName *string `tfsdk:"backup_storage_name" json:"backupStorageName,omitempty"`
				Enabled           *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				UploadIntervalSec *int64  `tfsdk:"upload_interval_sec" json:"uploadIntervalSec,omitempty"`
			} `tfsdk:"pitr" json:"pitr,omitempty"`
			Schedules *[]struct {
				BackupStorageName *string `tfsdk:"backup_storage_name" json:"backupStorageName,omitempty"`
				Enabled           *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				RetentionCopies   *int64  `tfsdk:"retention_copies" json:"retentionCopies,omitempty"`
				Schedule          *string `tfsdk:"schedule" json:"schedule,omitempty"`
			} `tfsdk:"schedules" json:"schedules,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		DataSource *struct {
			BackupSource *struct {
				BackupStorageName *string `tfsdk:"backup_storage_name" json:"backupStorageName,omitempty"`
				Path              *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"backup_source" json:"backupSource,omitempty"`
			DataImport *struct {
				Config           *map[string]string `tfsdk:"config" json:"config,omitempty"`
				DataImporterName *string            `tfsdk:"data_importer_name" json:"dataImporterName,omitempty"`
				Source           *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
					S3   *struct {
						AccessKeyId           *string `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
						Bucket                *string `tfsdk:"bucket" json:"bucket,omitempty"`
						CredentialsSecretName *string `tfsdk:"credentials_secret_name" json:"credentialsSecretName,omitempty"`
						EndpointURL           *string `tfsdk:"endpoint_url" json:"endpointURL,omitempty"`
						ForcePathStyle        *bool   `tfsdk:"force_path_style" json:"forcePathStyle,omitempty"`
						Region                *string `tfsdk:"region" json:"region,omitempty"`
						SecretAccessKey       *string `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
						VerifyTLS             *bool   `tfsdk:"verify_tls" json:"verifyTLS,omitempty"`
					} `tfsdk:"s3" json:"s3,omitempty"`
				} `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"data_import" json:"dataImport,omitempty"`
			DbClusterBackupName *string `tfsdk:"db_cluster_backup_name" json:"dbClusterBackupName,omitempty"`
			Pitr                *struct {
				Date *string `tfsdk:"date" json:"date,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"pitr" json:"pitr,omitempty"`
		} `tfsdk:"data_source" json:"dataSource,omitempty"`
		Engine *struct {
			Config    *string `tfsdk:"config" json:"config,omitempty"`
			CrVersion *string `tfsdk:"cr_version" json:"crVersion,omitempty"`
			Replicas  *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Storage *struct {
				Class *string `tfsdk:"class" json:"class,omitempty"`
				Size  *string `tfsdk:"size" json:"size,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			Type            *string `tfsdk:"type" json:"type,omitempty"`
			UserSecretsName *string `tfsdk:"user_secrets_name" json:"userSecretsName,omitempty"`
			Version         *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"engine" json:"engine,omitempty"`
		Monitoring *struct {
			MonitoringConfigName *string `tfsdk:"monitoring_config_name" json:"monitoringConfigName,omitempty"`
			Resources            *struct {
				Claims *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"monitoring" json:"monitoring,omitempty"`
		Paused                  *bool   `tfsdk:"paused" json:"paused,omitempty"`
		PodSchedulingPolicyName *string `tfsdk:"pod_scheduling_policy_name" json:"podSchedulingPolicyName,omitempty"`
		Proxy                   *struct {
			Config *string `tfsdk:"config" json:"config,omitempty"`
			Expose *struct {
				IpSourceRanges *[]string `tfsdk:"ip_source_ranges" json:"ipSourceRanges,omitempty"`
				Type           *string   `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"expose" json:"expose,omitempty"`
			Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"proxy" json:"proxy,omitempty"`
		Sharding *struct {
			ConfigServer *struct {
				Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"config_server" json:"configServer,omitempty"`
			Enabled *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
			Shards  *int64 `tfsdk:"shards" json:"shards,omitempty"`
		} `tfsdk:"sharding" json:"sharding,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EverestPerconaComDatabaseClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_everest_percona_com_database_cluster_v1alpha1_manifest"
}

func (r *EverestPerconaComDatabaseClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DatabaseCluster is the Schema for the databaseclusters API.",
		MarkdownDescription: "DatabaseCluster is the Schema for the databaseclusters API.",
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
				Description:         "DatabaseClusterSpec defines the desired state of DatabaseCluster.",
				MarkdownDescription: "DatabaseClusterSpec defines the desired state of DatabaseCluster.",
				Attributes: map[string]schema.Attribute{
					"allow_unsafe_configuration": schema.BoolAttribute{
						Description:         "AllowUnsafeConfiguration field used to ensure that the user can create configurations unfit for production use. Deprecated: AllowUnsafeConfiguration will not be supported in the future releases.",
						MarkdownDescription: "AllowUnsafeConfiguration field used to ensure that the user can create configurations unfit for production use. Deprecated: AllowUnsafeConfiguration will not be supported in the future releases.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup": schema.SingleNestedAttribute{
						Description:         "Backup is the backup specification",
						MarkdownDescription: "Backup is the backup specification",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled is a flag to enable backups Deprecated. Please use db.spec.backup.schedules[].enabled to control each schedule separately and db.spec.backup.pitr.enabled to control PITR.",
								MarkdownDescription: "Enabled is a flag to enable backups Deprecated. Please use db.spec.backup.schedules[].enabled to control each schedule separately and db.spec.backup.pitr.enabled to control PITR.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pitr": schema.SingleNestedAttribute{
								Description:         "PITR is the configuration of the point in time recovery",
								MarkdownDescription: "PITR is the configuration of the point in time recovery",
								Attributes: map[string]schema.Attribute{
									"backup_storage_name": schema.StringAttribute{
										Description:         "BackupStorageName is the name of the BackupStorage where the PITR is enabled The BackupStorage must be created in the same namespace as the DatabaseCluster.",
										MarkdownDescription: "BackupStorageName is the name of the BackupStorage where the PITR is enabled The BackupStorage must be created in the same namespace as the DatabaseCluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled is a flag to enable PITR",
										MarkdownDescription: "Enabled is a flag to enable PITR",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"upload_interval_sec": schema.Int64Attribute{
										Description:         "UploadIntervalSec number of seconds between the binlogs uploads",
										MarkdownDescription: "UploadIntervalSec number of seconds between the binlogs uploads",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"schedules": schema.ListNestedAttribute{
								Description:         "Schedules is a list of backup schedules",
								MarkdownDescription: "Schedules is a list of backup schedules",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"backup_storage_name": schema.StringAttribute{
											Description:         "BackupStorageName is the name of the BackupStorage CR that defines the storage location. The BackupStorage must be created in the same namespace as the DatabaseCluster.",
											MarkdownDescription: "BackupStorageName is the name of the BackupStorage CR that defines the storage location. The BackupStorage must be created in the same namespace as the DatabaseCluster.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"enabled": schema.BoolAttribute{
											Description:         "Enabled is a flag to enable the schedule",
											MarkdownDescription: "Enabled is a flag to enable the schedule",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the schedule",
											MarkdownDescription: "Name is the name of the schedule",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"retention_copies": schema.Int64Attribute{
											Description:         "RetentionCopies is the number of backup copies to retain",
											MarkdownDescription: "RetentionCopies is the number of backup copies to retain",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"schedule": schema.StringAttribute{
											Description:         "Schedule is the cron schedule",
											MarkdownDescription: "Schedule is the cron schedule",
											Required:            true,
											Optional:            false,
											Computed:            false,
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

					"data_source": schema.SingleNestedAttribute{
						Description:         "DataSource defines a data source for bootstraping a new cluster",
						MarkdownDescription: "DataSource defines a data source for bootstraping a new cluster",
						Attributes: map[string]schema.Attribute{
							"backup_source": schema.SingleNestedAttribute{
								Description:         "BackupSource is the backup source to restore from",
								MarkdownDescription: "BackupSource is the backup source to restore from",
								Attributes: map[string]schema.Attribute{
									"backup_storage_name": schema.StringAttribute{
										Description:         "BackupStorageName is the name of the BackupStorage used for backups. The BackupStorage must be created in the same namespace as the DatabaseCluster.",
										MarkdownDescription: "BackupStorageName is the name of the BackupStorage used for backups. The BackupStorage must be created in the same namespace as the DatabaseCluster.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is the path to the backup file/directory.",
										MarkdownDescription: "Path is the path to the backup file/directory.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_import": schema.SingleNestedAttribute{
								Description:         "DataImport allows importing data from an external backup source.",
								MarkdownDescription: "DataImport allows importing data from an external backup source.",
								Attributes: map[string]schema.Attribute{
									"config": schema.MapAttribute{
										Description:         "Config defines the configuration for the data import job. These options are specific to the DataImporter being used and must conform to the schema defined in the DataImporter's .spec.config.openAPIV3Schema.",
										MarkdownDescription: "Config defines the configuration for the data import job. These options are specific to the DataImporter being used and must conform to the schema defined in the DataImporter's .spec.config.openAPIV3Schema.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"data_importer_name": schema.StringAttribute{
										Description:         "DataImporterName is the data importer to use for the import.",
										MarkdownDescription: "DataImporterName is the data importer to use for the import.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"source": schema.SingleNestedAttribute{
										Description:         "Source is the source of the data to import.",
										MarkdownDescription: "Source is the source of the data to import.",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path is the path to the directory to import the data from. This may be a path to a file or a directory, depending on the data importer. Only absolute file paths are allowed. Leading and trailing '/' are optional.",
												MarkdownDescription: "Path is the path to the directory to import the data from. This may be a path to a file or a directory, depending on the data importer. Only absolute file paths are allowed. Leading and trailing '/' are optional.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"s3": schema.SingleNestedAttribute{
												Description:         "S3 contains the S3 information for the data import.",
												MarkdownDescription: "S3 contains the S3 information for the data import.",
												Attributes: map[string]schema.Attribute{
													"access_key_id": schema.StringAttribute{
														Description:         "AccessKeyID allows specifying the S3 access key ID inline. It is provided as a write-only input field for convenience. When this field is set, a webhook writes this value in the Secret specified by 'credentialsSecretName' and empties this field. This field is not stored in the API.",
														MarkdownDescription: "AccessKeyID allows specifying the S3 access key ID inline. It is provided as a write-only input field for convenience. When this field is set, a webhook writes this value in the Secret specified by 'credentialsSecretName' and empties this field. This field is not stored in the API.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bucket": schema.StringAttribute{
														Description:         "Bucket is the name of the S3 bucket.",
														MarkdownDescription: "Bucket is the name of the S3 bucket.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"credentials_secret_name": schema.StringAttribute{
														Description:         "CredentialsSecreName is the reference to the secret containing the S3 credentials. The Secret must contain the keys 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY'.",
														MarkdownDescription: "CredentialsSecreName is the reference to the secret containing the S3 credentials. The Secret must contain the keys 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"endpoint_url": schema.StringAttribute{
														Description:         "EndpointURL is an endpoint URL of backup storage.",
														MarkdownDescription: "EndpointURL is an endpoint URL of backup storage.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"force_path_style": schema.BoolAttribute{
														Description:         "ForcePathStyle is set to use path-style URLs. If unspecified, the default value is false.",
														MarkdownDescription: "ForcePathStyle is set to use path-style URLs. If unspecified, the default value is false.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "Region is the region of the S3 bucket.",
														MarkdownDescription: "Region is the region of the S3 bucket.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"secret_access_key": schema.StringAttribute{
														Description:         "SecretAccessKey allows specifying the S3 secret access key inline. It is provided as a write-only input field for convenience. When this field is set, a webhook writes this value in the Secret specified by 'credentialsSecretName' and empties this field. This field is not stored in the API.",
														MarkdownDescription: "SecretAccessKey allows specifying the S3 secret access key inline. It is provided as a write-only input field for convenience. When this field is set, a webhook writes this value in the Secret specified by 'credentialsSecretName' and empties this field. This field is not stored in the API.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"verify_tls": schema.BoolAttribute{
														Description:         "VerifyTLS is set to ensure TLS/SSL verification. If unspecified, the default value is true.",
														MarkdownDescription: "VerifyTLS is set to ensure TLS/SSL verification. If unspecified, the default value is true.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_cluster_backup_name": schema.StringAttribute{
								Description:         "DBClusterBackupName is the name of the DB cluster backup to restore from",
								MarkdownDescription: "DBClusterBackupName is the name of the DB cluster backup to restore from",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pitr": schema.SingleNestedAttribute{
								Description:         "PITR is the point-in-time recovery configuration",
								MarkdownDescription: "PITR is the point-in-time recovery configuration",
								Attributes: map[string]schema.Attribute{
									"date": schema.StringAttribute{
										Description:         "Date is the UTC date to recover to. The accepted format: '2006-01-02T15:04:05Z'.",
										MarkdownDescription: "Date is the UTC date to recover to. The accepted format: '2006-01-02T15:04:05Z'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is the type of recovery.",
										MarkdownDescription: "Type is the type of recovery.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("date", "latest"),
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

					"engine": schema.SingleNestedAttribute{
						Description:         "Engine is the database engine specification",
						MarkdownDescription: "Engine is the database engine specification",
						Attributes: map[string]schema.Attribute{
							"config": schema.StringAttribute{
								Description:         "Config is the engine configuration",
								MarkdownDescription: "Config is the engine configuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cr_version": schema.StringAttribute{
								Description:         "CRVersion is the desired version of the CR to use with the underlying operator. If unspecified, everest-operator will use the same version as the operator. NOTE: Updating this property post installation may lead to a restart of the cluster.",
								MarkdownDescription: "CRVersion is the desired version of the CR to use with the underlying operator. If unspecified, everest-operator will use the same version as the operator. NOTE: Updating this property post installation may lead to a restart of the cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the number of engine replicas",
								MarkdownDescription: "Replicas is the number of engine replicas",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources are the resource limits for each engine replica. If not set, resource limits are not imposed",
								MarkdownDescription: "Resources are the resource limits for each engine replica. If not set, resource limits are not imposed",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "CPU is the CPU resource requirements",
										MarkdownDescription: "CPU is the CPU resource requirements",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "Memory is the memory resource requirements",
										MarkdownDescription: "Memory is the memory resource requirements",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage": schema.SingleNestedAttribute{
								Description:         "Storage is the engine storage configuration",
								MarkdownDescription: "Storage is the engine storage configuration",
								Attributes: map[string]schema.Attribute{
									"class": schema.StringAttribute{
										Description:         "Class is the storage class to use for the persistent volume claim",
										MarkdownDescription: "Class is the storage class to use for the persistent volume claim",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"size": schema.StringAttribute{
										Description:         "Size is the size of the persistent volume claim",
										MarkdownDescription: "Size is the size of the persistent volume claim",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the engine type",
								MarkdownDescription: "Type is the engine type",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("pxc", "postgresql", "psmdb"),
								},
							},

							"user_secrets_name": schema.StringAttribute{
								Description:         "UserSecretsName is the name of the secret containing the user secrets",
								MarkdownDescription: "UserSecretsName is the name of the secret containing the user secrets",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version is the engine version",
								MarkdownDescription: "Version is the engine version",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"monitoring": schema.SingleNestedAttribute{
						Description:         "Monitoring is the monitoring configuration",
						MarkdownDescription: "Monitoring is the monitoring configuration",
						Attributes: map[string]schema.Attribute{
							"monitoring_config_name": schema.StringAttribute{
								Description:         "MonitoringConfigName is the name of a monitoringConfig CR. The MonitoringConfig must be created in the same namespace as the DatabaseCluster.",
								MarkdownDescription: "MonitoringConfigName is the name of a monitoringConfig CR. The MonitoringConfig must be created in the same namespace as the DatabaseCluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources defines resource limitations for the monitoring.",
								MarkdownDescription: "Resources defines resource limitations for the monitoring.",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"request": schema.StringAttribute{
													Description:         "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
													MarkdownDescription: "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
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

					"paused": schema.BoolAttribute{
						Description:         "Paused is a flag to stop the cluster",
						MarkdownDescription: "Paused is a flag to stop the cluster",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_scheduling_policy_name": schema.StringAttribute{
						Description:         "PodSchedulingPolicyName is the name of the PodSchedulingPolicy CR that defines rules for DB cluster pods allocation across the cluster.",
						MarkdownDescription: "PodSchedulingPolicyName is the name of the PodSchedulingPolicy CR that defines rules for DB cluster pods allocation across the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"proxy": schema.SingleNestedAttribute{
						Description:         "Proxy is the proxy specification. If not set, an appropriate proxy specification will be applied for the given engine. A common use case for setting this field is to control the external access to the database cluster.",
						MarkdownDescription: "Proxy is the proxy specification. If not set, an appropriate proxy specification will be applied for the given engine. A common use case for setting this field is to control the external access to the database cluster.",
						Attributes: map[string]schema.Attribute{
							"config": schema.StringAttribute{
								Description:         "Config is the proxy configuration",
								MarkdownDescription: "Config is the proxy configuration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"expose": schema.SingleNestedAttribute{
								Description:         "Expose is the proxy expose configuration",
								MarkdownDescription: "Expose is the proxy expose configuration",
								Attributes: map[string]schema.Attribute{
									"ip_source_ranges": schema.ListAttribute{
										Description:         "IPSourceRanges is the list of IP source ranges (CIDR notation) to allow access from. If not set, there is no limitations",
										MarkdownDescription: "IPSourceRanges is the list of IP source ranges (CIDR notation) to allow access from. If not set, there is no limitations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is the expose type, can be internal or external",
										MarkdownDescription: "Type is the expose type, can be internal or external",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("internal", "external"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the number of proxy replicas",
								MarkdownDescription: "Replicas is the number of proxy replicas",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources are the resource limits for each proxy replica. If not set, resource limits are not imposed",
								MarkdownDescription: "Resources are the resource limits for each proxy replica. If not set, resource limits are not imposed",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "CPU is the CPU resource requirements",
										MarkdownDescription: "CPU is the CPU resource requirements",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "Memory is the memory resource requirements",
										MarkdownDescription: "Memory is the memory resource requirements",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the proxy type",
								MarkdownDescription: "Type is the proxy type",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("mongos", "haproxy", "proxysql", "pgbouncer"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sharding": schema.SingleNestedAttribute{
						Description:         "Sharding is the sharding configuration. PSMDB-only",
						MarkdownDescription: "Sharding is the sharding configuration. PSMDB-only",
						Attributes: map[string]schema.Attribute{
							"config_server": schema.SingleNestedAttribute{
								Description:         "ConfigServer represents the sharding configuration server settings",
								MarkdownDescription: "ConfigServer represents the sharding configuration server settings",
								Attributes: map[string]schema.Attribute{
									"replicas": schema.Int64Attribute{
										Description:         "Replicas is the amount of configServers",
										MarkdownDescription: "Replicas is the amount of configServers",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enabled defines if the sharding is enabled",
								MarkdownDescription: "Enabled defines if the sharding is enabled",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"shards": schema.Int64Attribute{
								Description:         "Shards defines the number of shards",
								MarkdownDescription: "Shards defines the number of shards",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
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

func (r *EverestPerconaComDatabaseClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_everest_percona_com_database_cluster_v1alpha1_manifest")

	var model EverestPerconaComDatabaseClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("everest.percona.com/v1alpha1")
	model.Kind = pointer.String("DatabaseCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
