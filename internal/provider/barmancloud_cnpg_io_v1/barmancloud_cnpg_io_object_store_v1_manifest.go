/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package barmancloud_cnpg_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &BarmancloudCnpgIoObjectStoreV1Manifest{}
)

func NewBarmancloudCnpgIoObjectStoreV1Manifest() datasource.DataSource {
	return &BarmancloudCnpgIoObjectStoreV1Manifest{}
}

type BarmancloudCnpgIoObjectStoreV1Manifest struct{}

type BarmancloudCnpgIoObjectStoreV1ManifestData struct {
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
		Configuration *struct {
			AzureCredentials *struct {
				ConnectionString *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"connection_string" json:"connectionString,omitempty"`
				InheritFromAzureAD *bool `tfsdk:"inherit_from_azure_ad" json:"inheritFromAzureAD,omitempty"`
				StorageAccount     *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"storage_account" json:"storageAccount,omitempty"`
				StorageKey *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"storage_key" json:"storageKey,omitempty"`
				StorageSasToken *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"storage_sas_token" json:"storageSasToken,omitempty"`
			} `tfsdk:"azure_credentials" json:"azureCredentials,omitempty"`
			Data *struct {
				AdditionalCommandArgs *[]string `tfsdk:"additional_command_args" json:"additionalCommandArgs,omitempty"`
				Compression           *string   `tfsdk:"compression" json:"compression,omitempty"`
				Encryption            *string   `tfsdk:"encryption" json:"encryption,omitempty"`
				ImmediateCheckpoint   *bool     `tfsdk:"immediate_checkpoint" json:"immediateCheckpoint,omitempty"`
				Jobs                  *int64    `tfsdk:"jobs" json:"jobs,omitempty"`
			} `tfsdk:"data" json:"data,omitempty"`
			DestinationPath *string `tfsdk:"destination_path" json:"destinationPath,omitempty"`
			EndpointCA      *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"endpoint_ca" json:"endpointCA,omitempty"`
			EndpointURL       *string `tfsdk:"endpoint_url" json:"endpointURL,omitempty"`
			GoogleCredentials *struct {
				ApplicationCredentials *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"application_credentials" json:"applicationCredentials,omitempty"`
				GkeEnvironment *bool `tfsdk:"gke_environment" json:"gkeEnvironment,omitempty"`
			} `tfsdk:"google_credentials" json:"googleCredentials,omitempty"`
			HistoryTags   *map[string]string `tfsdk:"history_tags" json:"historyTags,omitempty"`
			S3Credentials *struct {
				AccessKeyId *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
				InheritFromIAMRole *bool `tfsdk:"inherit_from_iam_role" json:"inheritFromIAMRole,omitempty"`
				Region             *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"region" json:"region,omitempty"`
				SecretAccessKey *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
				SessionToken *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"session_token" json:"sessionToken,omitempty"`
			} `tfsdk:"s3_credentials" json:"s3Credentials,omitempty"`
			ServerName *string            `tfsdk:"server_name" json:"serverName,omitempty"`
			Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
			Wal        *struct {
				ArchiveAdditionalCommandArgs *[]string `tfsdk:"archive_additional_command_args" json:"archiveAdditionalCommandArgs,omitempty"`
				Compression                  *string   `tfsdk:"compression" json:"compression,omitempty"`
				Encryption                   *string   `tfsdk:"encryption" json:"encryption,omitempty"`
				MaxParallel                  *int64    `tfsdk:"max_parallel" json:"maxParallel,omitempty"`
				RestoreAdditionalCommandArgs *[]string `tfsdk:"restore_additional_command_args" json:"restoreAdditionalCommandArgs,omitempty"`
			} `tfsdk:"wal" json:"wal,omitempty"`
		} `tfsdk:"configuration" json:"configuration,omitempty"`
		InstanceSidecarConfiguration *struct {
			Env *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					FieldRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
					ResourceFieldRef *struct {
						ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
						Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
						Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					SecretKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"env" json:"env,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RetentionPolicyIntervalSeconds *int64 `tfsdk:"retention_policy_interval_seconds" json:"retentionPolicyIntervalSeconds,omitempty"`
		} `tfsdk:"instance_sidecar_configuration" json:"instanceSidecarConfiguration,omitempty"`
		RetentionPolicy *string `tfsdk:"retention_policy" json:"retentionPolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BarmancloudCnpgIoObjectStoreV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_barmancloud_cnpg_io_object_store_v1_manifest"
}

func (r *BarmancloudCnpgIoObjectStoreV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ObjectStore is the Schema for the objectstores API.",
		MarkdownDescription: "ObjectStore is the Schema for the objectstores API.",
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
				Description:         "Specification of the desired behavior of the ObjectStore. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired behavior of the ObjectStore. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Attributes: map[string]schema.Attribute{
					"configuration": schema.SingleNestedAttribute{
						Description:         "The configuration for the barman-cloud tool suite",
						MarkdownDescription: "The configuration for the barman-cloud tool suite",
						Attributes: map[string]schema.Attribute{
							"azure_credentials": schema.SingleNestedAttribute{
								Description:         "The credentials to use to upload data to Azure Blob Storage",
								MarkdownDescription: "The credentials to use to upload data to Azure Blob Storage",
								Attributes: map[string]schema.Attribute{
									"connection_string": schema.SingleNestedAttribute{
										Description:         "The connection string to be used",
										MarkdownDescription: "The connection string to be used",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"inherit_from_azure_ad": schema.BoolAttribute{
										Description:         "Use the Azure AD based authentication without providing explicitly the keys.",
										MarkdownDescription: "Use the Azure AD based authentication without providing explicitly the keys.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_account": schema.SingleNestedAttribute{
										Description:         "The storage account where to upload data",
										MarkdownDescription: "The storage account where to upload data",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_key": schema.SingleNestedAttribute{
										Description:         "The storage account key to be used in conjunction with the storage account name",
										MarkdownDescription: "The storage account key to be used in conjunction with the storage account name",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_sas_token": schema.SingleNestedAttribute{
										Description:         "A shared-access-signature to be used in conjunction with the storage account name",
										MarkdownDescription: "A shared-access-signature to be used in conjunction with the storage account name",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
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

							"data": schema.SingleNestedAttribute{
								Description:         "The configuration to be used to backup the data files When not defined, base backups files will be stored uncompressed and may be unencrypted in the object store, according to the bucket default policy.",
								MarkdownDescription: "The configuration to be used to backup the data files When not defined, base backups files will be stored uncompressed and may be unencrypted in the object store, according to the bucket default policy.",
								Attributes: map[string]schema.Attribute{
									"additional_command_args": schema.ListAttribute{
										Description:         "AdditionalCommandArgs represents additional arguments that can be appended to the 'barman-cloud-backup' command-line invocation. These arguments provide flexibility to customize the backup process further according to specific requirements or configurations. Example: In a scenario where specialized backup options are required, such as setting a specific timeout or defining custom behavior, users can use this field to specify additional command arguments. Note: It's essential to ensure that the provided arguments are valid and supported by the 'barman-cloud-backup' command, to avoid potential errors or unintended behavior during execution.",
										MarkdownDescription: "AdditionalCommandArgs represents additional arguments that can be appended to the 'barman-cloud-backup' command-line invocation. These arguments provide flexibility to customize the backup process further according to specific requirements or configurations. Example: In a scenario where specialized backup options are required, such as setting a specific timeout or defining custom behavior, users can use this field to specify additional command arguments. Note: It's essential to ensure that the provided arguments are valid and supported by the 'barman-cloud-backup' command, to avoid potential errors or unintended behavior during execution.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"compression": schema.StringAttribute{
										Description:         "Compress a backup file (a tar file per tablespace) while streaming it to the object store. Available options are empty string (no compression, default), 'gzip', 'bzip2', and 'snappy'.",
										MarkdownDescription: "Compress a backup file (a tar file per tablespace) while streaming it to the object store. Available options are empty string (no compression, default), 'gzip', 'bzip2', and 'snappy'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("bzip2", "gzip", "snappy"),
										},
									},

									"encryption": schema.StringAttribute{
										Description:         "Whenever to force the encryption of files (if the bucket is not already configured for that). Allowed options are empty string (use the bucket policy, default), 'AES256' and 'aws:kms'",
										MarkdownDescription: "Whenever to force the encryption of files (if the bucket is not already configured for that). Allowed options are empty string (use the bucket policy, default), 'AES256' and 'aws:kms'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("AES256", "aws:kms"),
										},
									},

									"immediate_checkpoint": schema.BoolAttribute{
										Description:         "Control whether the I/O workload for the backup initial checkpoint will be limited, according to the 'checkpoint_completion_target' setting on the PostgreSQL server. If set to true, an immediate checkpoint will be used, meaning PostgreSQL will complete the checkpoint as soon as possible. 'false' by default.",
										MarkdownDescription: "Control whether the I/O workload for the backup initial checkpoint will be limited, according to the 'checkpoint_completion_target' setting on the PostgreSQL server. If set to true, an immediate checkpoint will be used, meaning PostgreSQL will complete the checkpoint as soon as possible. 'false' by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"jobs": schema.Int64Attribute{
										Description:         "The number of parallel jobs to be used to upload the backup, defaults to 2",
										MarkdownDescription: "The number of parallel jobs to be used to upload the backup, defaults to 2",
										Required:            false,
										Optional:            true,
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

							"destination_path": schema.StringAttribute{
								Description:         "The path where to store the backup (i.e. s3://bucket/path/to/folder) this path, with different destination folders, will be used for WALs and for data",
								MarkdownDescription: "The path where to store the backup (i.e. s3://bucket/path/to/folder) this path, with different destination folders, will be used for WALs and for data",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"endpoint_ca": schema.SingleNestedAttribute{
								Description:         "EndpointCA store the CA bundle of the barman endpoint. Useful when using self-signed certificates to avoid errors with certificate issuer and barman-cloud-wal-archive",
								MarkdownDescription: "EndpointCA store the CA bundle of the barman endpoint. Useful when using self-signed certificates to avoid errors with certificate issuer and barman-cloud-wal-archive",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key to select",
										MarkdownDescription: "The key to select",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent.",
										MarkdownDescription: "Name of the referent.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"endpoint_url": schema.StringAttribute{
								Description:         "Endpoint to be used to upload data to the cloud, overriding the automatic endpoint discovery",
								MarkdownDescription: "Endpoint to be used to upload data to the cloud, overriding the automatic endpoint discovery",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"google_credentials": schema.SingleNestedAttribute{
								Description:         "The credentials to use to upload data to Google Cloud Storage",
								MarkdownDescription: "The credentials to use to upload data to Google Cloud Storage",
								Attributes: map[string]schema.Attribute{
									"application_credentials": schema.SingleNestedAttribute{
										Description:         "The secret containing the Google Cloud Storage JSON file with the credentials",
										MarkdownDescription: "The secret containing the Google Cloud Storage JSON file with the credentials",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"gke_environment": schema.BoolAttribute{
										Description:         "If set to true, will presume that it's running inside a GKE environment, default to false.",
										MarkdownDescription: "If set to true, will presume that it's running inside a GKE environment, default to false.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"history_tags": schema.MapAttribute{
								Description:         "HistoryTags is a list of key value pairs that will be passed to the Barman --history-tags option.",
								MarkdownDescription: "HistoryTags is a list of key value pairs that will be passed to the Barman --history-tags option.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_credentials": schema.SingleNestedAttribute{
								Description:         "The credentials to use to upload data to S3",
								MarkdownDescription: "The credentials to use to upload data to S3",
								Attributes: map[string]schema.Attribute{
									"access_key_id": schema.SingleNestedAttribute{
										Description:         "The reference to the access key id",
										MarkdownDescription: "The reference to the access key id",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"inherit_from_iam_role": schema.BoolAttribute{
										Description:         "Use the role based authentication without providing explicitly the keys.",
										MarkdownDescription: "Use the role based authentication without providing explicitly the keys.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"region": schema.SingleNestedAttribute{
										Description:         "The reference to the secret containing the region name",
										MarkdownDescription: "The reference to the secret containing the region name",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_access_key": schema.SingleNestedAttribute{
										Description:         "The reference to the secret access key",
										MarkdownDescription: "The reference to the secret access key",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"session_token": schema.SingleNestedAttribute{
										Description:         "The references to the session key",
										MarkdownDescription: "The references to the session key",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
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

							"server_name": schema.StringAttribute{
								Description:         "The server name on S3, the cluster name is used if this parameter is omitted",
								MarkdownDescription: "The server name on S3, the cluster name is used if this parameter is omitted",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tags": schema.MapAttribute{
								Description:         "Tags is a list of key value pairs that will be passed to the Barman --tags option.",
								MarkdownDescription: "Tags is a list of key value pairs that will be passed to the Barman --tags option.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wal": schema.SingleNestedAttribute{
								Description:         "The configuration for the backup of the WAL stream. When not defined, WAL files will be stored uncompressed and may be unencrypted in the object store, according to the bucket default policy.",
								MarkdownDescription: "The configuration for the backup of the WAL stream. When not defined, WAL files will be stored uncompressed and may be unencrypted in the object store, according to the bucket default policy.",
								Attributes: map[string]schema.Attribute{
									"archive_additional_command_args": schema.ListAttribute{
										Description:         "Additional arguments that can be appended to the 'barman-cloud-wal-archive' command-line invocation. These arguments provide flexibility to customize the WAL archive process further, according to specific requirements or configurations. Example: In a scenario where specialized backup options are required, such as setting a specific timeout or defining custom behavior, users can use this field to specify additional command arguments. Note: It's essential to ensure that the provided arguments are valid and supported by the 'barman-cloud-wal-archive' command, to avoid potential errors or unintended behavior during execution.",
										MarkdownDescription: "Additional arguments that can be appended to the 'barman-cloud-wal-archive' command-line invocation. These arguments provide flexibility to customize the WAL archive process further, according to specific requirements or configurations. Example: In a scenario where specialized backup options are required, such as setting a specific timeout or defining custom behavior, users can use this field to specify additional command arguments. Note: It's essential to ensure that the provided arguments are valid and supported by the 'barman-cloud-wal-archive' command, to avoid potential errors or unintended behavior during execution.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"compression": schema.StringAttribute{
										Description:         "Compress a WAL file before sending it to the object store. Available options are empty string (no compression, default), 'gzip', 'bzip2', 'lz4', 'snappy', 'xz', and 'zstd'.",
										MarkdownDescription: "Compress a WAL file before sending it to the object store. Available options are empty string (no compression, default), 'gzip', 'bzip2', 'lz4', 'snappy', 'xz', and 'zstd'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("bzip2", "gzip", "lz4", "snappy", "xz", "zstd"),
										},
									},

									"encryption": schema.StringAttribute{
										Description:         "Whenever to force the encryption of files (if the bucket is not already configured for that). Allowed options are empty string (use the bucket policy, default), 'AES256' and 'aws:kms'",
										MarkdownDescription: "Whenever to force the encryption of files (if the bucket is not already configured for that). Allowed options are empty string (use the bucket policy, default), 'AES256' and 'aws:kms'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("AES256", "aws:kms"),
										},
									},

									"max_parallel": schema.Int64Attribute{
										Description:         "Number of WAL files to be either archived in parallel (when the PostgreSQL instance is archiving to a backup object store) or restored in parallel (when a PostgreSQL standby is fetching WAL files from a recovery object store). If not specified, WAL files will be processed one at a time. It accepts a positive integer as a value - with 1 being the minimum accepted value.",
										MarkdownDescription: "Number of WAL files to be either archived in parallel (when the PostgreSQL instance is archiving to a backup object store) or restored in parallel (when a PostgreSQL standby is fetching WAL files from a recovery object store). If not specified, WAL files will be processed one at a time. It accepts a positive integer as a value - with 1 being the minimum accepted value.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"restore_additional_command_args": schema.ListAttribute{
										Description:         "Additional arguments that can be appended to the 'barman-cloud-wal-restore' command-line invocation. These arguments provide flexibility to customize the WAL restore process further, according to specific requirements or configurations. Example: In a scenario where specialized backup options are required, such as setting a specific timeout or defining custom behavior, users can use this field to specify additional command arguments. Note: It's essential to ensure that the provided arguments are valid and supported by the 'barman-cloud-wal-restore' command, to avoid potential errors or unintended behavior during execution.",
										MarkdownDescription: "Additional arguments that can be appended to the 'barman-cloud-wal-restore' command-line invocation. These arguments provide flexibility to customize the WAL restore process further, according to specific requirements or configurations. Example: In a scenario where specialized backup options are required, such as setting a specific timeout or defining custom behavior, users can use this field to specify additional command arguments. Note: It's essential to ensure that the provided arguments are valid and supported by the 'barman-cloud-wal-restore' command, to avoid potential errors or unintended behavior during execution.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"instance_sidecar_configuration": schema.SingleNestedAttribute{
						Description:         "The configuration for the sidecar that runs in the instance pods",
						MarkdownDescription: "The configuration for the sidecar that runs in the instance pods",
						Attributes: map[string]schema.Attribute{
							"env": schema.ListNestedAttribute{
								Description:         "The environment to be explicitly passed to the sidecar",
								MarkdownDescription: "The environment to be explicitly passed to the sidecar",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
											MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
											MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value_from": schema.SingleNestedAttribute{
											Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
											MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a ConfigMap.",
													MarkdownDescription: "Selects a key of a ConfigMap.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the ConfigMap or its key must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"field_path": schema.StringAttribute{
															Description:         "Path of the field to select in the specified API version.",
															MarkdownDescription: "Path of the field to select in the specified API version.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource_field_ref": schema.SingleNestedAttribute{
													Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
													Attributes: map[string]schema.Attribute{
														"container_name": schema.StringAttribute{
															Description:         "Container name: required for volumes, optional for env vars",
															MarkdownDescription: "Container name: required for volumes, optional for env vars",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"divisor": schema.StringAttribute{
															Description:         "Specifies the output format of the exposed resources, defaults to '1'",
															MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resource": schema.StringAttribute{
															Description:         "Required: resource to select",
															MarkdownDescription: "Required: resource to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret_key_ref": schema.SingleNestedAttribute{
													Description:         "Selects a key of a secret in the pod's namespace",
													MarkdownDescription: "Selects a key of a secret in the pod's namespace",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from. Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the Secret or its key must be defined",
															MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources define cpu/memory requests and limits for the sidecar that runs in the instance pods.",
								MarkdownDescription: "Resources define cpu/memory requests and limits for the sidecar that runs in the instance pods.",
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

							"retention_policy_interval_seconds": schema.Int64Attribute{
								Description:         "The retentionCheckInterval defines the frequency at which the system checks and enforces retention policies.",
								MarkdownDescription: "The retentionCheckInterval defines the frequency at which the system checks and enforces retention policies.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"retention_policy": schema.StringAttribute{
						Description:         "RetentionPolicy is the retention policy to be used for backups and WALs (i.e. '60d'). The retention policy is expressed in the form of 'XXu' where 'XX' is a positive integer and 'u' is in '[dwm]' - days, weeks, months.",
						MarkdownDescription: "RetentionPolicy is the retention policy to be used for backups and WALs (i.e. '60d'). The retention policy is expressed in the form of 'XXu' where 'XX' is a positive integer and 'u' is in '[dwm]' - days, weeks, months.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9][0-9]*[dwm]$`), ""),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *BarmancloudCnpgIoObjectStoreV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_barmancloud_cnpg_io_object_store_v1_manifest")

	var model BarmancloudCnpgIoObjectStoreV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("barmancloud.cnpg.io/v1")
	model.Kind = pointer.String("ObjectStore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
