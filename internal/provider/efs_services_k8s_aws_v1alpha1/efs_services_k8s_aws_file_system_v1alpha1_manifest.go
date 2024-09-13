/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package efs_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &EfsServicesK8SAwsFileSystemV1Alpha1Manifest{}
)

func NewEfsServicesK8SAwsFileSystemV1Alpha1Manifest() datasource.DataSource {
	return &EfsServicesK8SAwsFileSystemV1Alpha1Manifest{}
}

type EfsServicesK8SAwsFileSystemV1Alpha1Manifest struct{}

type EfsServicesK8SAwsFileSystemV1Alpha1ManifestData struct {
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
		AvailabilityZoneName *string `tfsdk:"availability_zone_name" json:"availabilityZoneName,omitempty"`
		Backup               *bool   `tfsdk:"backup" json:"backup,omitempty"`
		BackupPolicy         *struct {
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"backup_policy" json:"backupPolicy,omitempty"`
		Encrypted            *bool `tfsdk:"encrypted" json:"encrypted,omitempty"`
		FileSystemProtection *struct {
			ReplicationOverwriteProtection *string `tfsdk:"replication_overwrite_protection" json:"replicationOverwriteProtection,omitempty"`
		} `tfsdk:"file_system_protection" json:"fileSystemProtection,omitempty"`
		KmsKeyID  *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		KmsKeyRef *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"kms_key_ref" json:"kmsKeyRef,omitempty"`
		LifecyclePolicies *[]struct {
			TransitionToArchive             *string `tfsdk:"transition_to_archive" json:"transitionToArchive,omitempty"`
			TransitionToIA                  *string `tfsdk:"transition_to_ia" json:"transitionToIA,omitempty"`
			TransitionToPrimaryStorageClass *string `tfsdk:"transition_to_primary_storage_class" json:"transitionToPrimaryStorageClass,omitempty"`
		} `tfsdk:"lifecycle_policies" json:"lifecyclePolicies,omitempty"`
		PerformanceMode              *string  `tfsdk:"performance_mode" json:"performanceMode,omitempty"`
		Policy                       *string  `tfsdk:"policy" json:"policy,omitempty"`
		ProvisionedThroughputInMiBps *float64 `tfsdk:"provisioned_throughput_in_mi_bps" json:"provisionedThroughputInMiBps,omitempty"`
		Tags                         *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		ThroughputMode *string `tfsdk:"throughput_mode" json:"throughputMode,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EfsServicesK8SAwsFileSystemV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_efs_services_k8s_aws_file_system_v1alpha1_manifest"
}

func (r *EfsServicesK8SAwsFileSystemV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FileSystem is the Schema for the FileSystems API",
		MarkdownDescription: "FileSystem is the Schema for the FileSystems API",
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
				Description:         "FileSystemSpec defines the desired state of FileSystem.",
				MarkdownDescription: "FileSystemSpec defines the desired state of FileSystem.",
				Attributes: map[string]schema.Attribute{
					"availability_zone_name": schema.StringAttribute{
						Description:         "Used to create a One Zone file system. It specifies the Amazon Web Services Availability Zone in which to create the file system. Use the format us-east-1a to specify the Availability Zone. For more information about One Zone file systems, see Using EFS storage classes (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the Amazon EFS User Guide. One Zone file systems are not available in all Availability Zones in Amazon Web Services Regions where Amazon EFS is available.",
						MarkdownDescription: "Used to create a One Zone file system. It specifies the Amazon Web Services Availability Zone in which to create the file system. Use the format us-east-1a to specify the Availability Zone. For more information about One Zone file systems, see Using EFS storage classes (https://docs.aws.amazon.com/efs/latest/ug/storage-classes.html) in the Amazon EFS User Guide. One Zone file systems are not available in all Availability Zones in Amazon Web Services Regions where Amazon EFS is available.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup": schema.BoolAttribute{
						Description:         "Specifies whether automatic backups are enabled on the file system that you are creating. Set the value to true to enable automatic backups. If you are creating a One Zone file system, automatic backups are enabled by default. For more information, see Automatic backups (https://docs.aws.amazon.com/efs/latest/ug/awsbackup.html#automatic-backups) in the Amazon EFS User Guide. Default is false. However, if you specify an AvailabilityZoneName, the default is true. Backup is not available in all Amazon Web Services Regions where Amazon EFS is available.",
						MarkdownDescription: "Specifies whether automatic backups are enabled on the file system that you are creating. Set the value to true to enable automatic backups. If you are creating a One Zone file system, automatic backups are enabled by default. For more information, see Automatic backups (https://docs.aws.amazon.com/efs/latest/ug/awsbackup.html#automatic-backups) in the Amazon EFS User Guide. Default is false. However, if you specify an AvailabilityZoneName, the default is true. Backup is not available in all Amazon Web Services Regions where Amazon EFS is available.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_policy": schema.SingleNestedAttribute{
						Description:         "The backup policy included in the PutBackupPolicy request.",
						MarkdownDescription: "The backup policy included in the PutBackupPolicy request.",
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
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

					"encrypted": schema.BoolAttribute{
						Description:         "A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying an existing Key Management Service key (KMS key). If you don't specify a KMS key, then the default KMS key for Amazon EFS, /aws/elasticfilesystem, is used to protect the encrypted file system.",
						MarkdownDescription: "A Boolean value that, if true, creates an encrypted file system. When creating an encrypted file system, you have the option of specifying an existing Key Management Service key (KMS key). If you don't specify a KMS key, then the default KMS key for Amazon EFS, /aws/elasticfilesystem, is used to protect the encrypted file system.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"file_system_protection": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"replication_overwrite_protection": schema.StringAttribute{
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

					"kms_key_id": schema.StringAttribute{
						Description:         "The ID of the KMS key that you want to use to protect the encrypted file system. This parameter is required only if you want to use a non-default KMS key. If this parameter is not specified, the default KMS key for Amazon EFS is used. You can specify a KMS key ID using the following formats: * Key ID - A unique identifier of the key, for example 1234abcd-12ab-34cd-56ef-1234567890ab. * ARN - An Amazon Resource Name (ARN) for the key, for example arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab. * Key alias - A previously created display name for a key, for example alias/projectKey1. * Key alias ARN - An ARN for a key alias, for example arn:aws:kms:us-west-2:444455556666:alias/projectKey1. If you use KmsKeyId, you must set the CreateFileSystemRequest$Encrypted parameter to true. EFS accepts only symmetric KMS keys. You cannot use asymmetric KMS keys with Amazon EFS file systems.",
						MarkdownDescription: "The ID of the KMS key that you want to use to protect the encrypted file system. This parameter is required only if you want to use a non-default KMS key. If this parameter is not specified, the default KMS key for Amazon EFS is used. You can specify a KMS key ID using the following formats: * Key ID - A unique identifier of the key, for example 1234abcd-12ab-34cd-56ef-1234567890ab. * ARN - An Amazon Resource Name (ARN) for the key, for example arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab. * Key alias - A previously created display name for a key, for example alias/projectKey1. * Key alias ARN - An ARN for a key alias, for example arn:aws:kms:us-west-2:444455556666:alias/projectKey1. If you use KmsKeyId, you must set the CreateFileSystemRequest$Encrypted parameter to true. EFS accepts only symmetric KMS keys. You cannot use asymmetric KMS keys with Amazon EFS file systems.",
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

					"lifecycle_policies": schema.ListNestedAttribute{
						Description:         "An array of LifecyclePolicy objects that define the file system's LifecycleConfiguration object. A LifecycleConfiguration object informs EFS Lifecycle management of the following: * TransitionToIA – When to move files in the file system from primary storage (Standard storage class) into the Infrequent Access (IA) storage. * TransitionToArchive – When to move files in the file system from their current storage class (either IA or Standard storage) into the Archive storage. File systems cannot transition into Archive storage before transitioning into IA storage. Therefore, TransitionToArchive must either not be set or must be later than TransitionToIA. The Archive storage class is available only for file systems that use the Elastic Throughput mode and the General Purpose Performance mode. * TransitionToPrimaryStorageClass – Whether to move files in the file system back to primary storage (Standard storage class) after they are accessed in IA or Archive storage. When using the put-lifecycle-configuration CLI command or the PutLifecycleConfiguration API action, Amazon EFS requires that each LifecyclePolicy object have only a single transition. This means that in a request body, LifecyclePolicies must be structured as an array of LifecyclePolicy objects, one object for each storage transition. See the example requests in the following section for more information.",
						MarkdownDescription: "An array of LifecyclePolicy objects that define the file system's LifecycleConfiguration object. A LifecycleConfiguration object informs EFS Lifecycle management of the following: * TransitionToIA – When to move files in the file system from primary storage (Standard storage class) into the Infrequent Access (IA) storage. * TransitionToArchive – When to move files in the file system from their current storage class (either IA or Standard storage) into the Archive storage. File systems cannot transition into Archive storage before transitioning into IA storage. Therefore, TransitionToArchive must either not be set or must be later than TransitionToIA. The Archive storage class is available only for file systems that use the Elastic Throughput mode and the General Purpose Performance mode. * TransitionToPrimaryStorageClass – Whether to move files in the file system back to primary storage (Standard storage class) after they are accessed in IA or Archive storage. When using the put-lifecycle-configuration CLI command or the PutLifecycleConfiguration API action, Amazon EFS requires that each LifecyclePolicy object have only a single transition. This means that in a request body, LifecyclePolicies must be structured as an array of LifecyclePolicy objects, one object for each storage transition. See the example requests in the following section for more information.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"transition_to_archive": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transition_to_ia": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"transition_to_primary_storage_class": schema.StringAttribute{
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

					"performance_mode": schema.StringAttribute{
						Description:         "The Performance mode of the file system. We recommend generalPurpose performance mode for all file systems. File systems using the maxIO performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The maxIO mode is not supported on One Zone file systems. Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems. Default is generalPurpose.",
						MarkdownDescription: "The Performance mode of the file system. We recommend generalPurpose performance mode for all file systems. File systems using the maxIO performance mode can scale to higher levels of aggregate throughput and operations per second with a tradeoff of slightly higher latencies for most file operations. The performance mode can't be changed after the file system has been created. The maxIO mode is not supported on One Zone file systems. Due to the higher per-operation latencies with Max I/O, we recommend using General Purpose performance mode for all file systems. Default is generalPurpose.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy": schema.StringAttribute{
						Description:         "The FileSystemPolicy that you're creating. Accepts a JSON formatted policy definition. EFS file system policies have a 20,000 character limit. To find out more about the elements that make up a file system policy, see EFS Resource-based Policies (https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies).",
						MarkdownDescription: "The FileSystemPolicy that you're creating. Accepts a JSON formatted policy definition. EFS file system policies have a 20,000 character limit. To find out more about the elements that make up a file system policy, see EFS Resource-based Policies (https://docs.aws.amazon.com/efs/latest/ug/access-control-overview.html#access-control-manage-access-intro-resource-policies).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provisioned_throughput_in_mi_bps": schema.Float64Attribute{
						Description:         "The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ThroughputMode is set to provisioned. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact Amazon Web Services Support. For more information, see Amazon EFS quotas that you can increase (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the Amazon EFS User Guide.",
						MarkdownDescription: "The throughput, measured in mebibytes per second (MiBps), that you want to provision for a file system that you're creating. Required if ThroughputMode is set to provisioned. Valid values are 1-3414 MiBps, with the upper limit depending on Region. To increase this limit, contact Amazon Web Services Support. For more information, see Amazon EFS quotas that you can increase (https://docs.aws.amazon.com/efs/latest/ug/limits.html#soft-limits) in the Amazon EFS User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a 'Key':'Name','Value':'{value}' key-value pair. Each key must be unique. For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon Web Services General Reference Guide.",
						MarkdownDescription: "Use to create one or more tags associated with the file system. Each tag is a user-defined key-value pair. Name your file system on creation by including a 'Key':'Name','Value':'{value}' key-value pair. Each key must be unique. For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon Web Services General Reference Guide.",
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

					"throughput_mode": schema.StringAttribute{
						Description:         "Specifies the throughput mode for the file system. The mode can be bursting, provisioned, or elastic. If you set ThroughputMode to provisioned, you must also set a value for ProvisionedThroughputInMibps. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see Specifying throughput with provisioned mode (https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the Amazon EFS User Guide. Default is bursting.",
						MarkdownDescription: "Specifies the throughput mode for the file system. The mode can be bursting, provisioned, or elastic. If you set ThroughputMode to provisioned, you must also set a value for ProvisionedThroughputInMibps. After you create the file system, you can decrease your file system's Provisioned throughput or change between the throughput modes, with certain time restrictions. For more information, see Specifying throughput with provisioned mode (https://docs.aws.amazon.com/efs/latest/ug/performance.html#provisioned-throughput) in the Amazon EFS User Guide. Default is bursting.",
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
	}
}

func (r *EfsServicesK8SAwsFileSystemV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_efs_services_k8s_aws_file_system_v1alpha1_manifest")

	var model EfsServicesK8SAwsFileSystemV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("efs.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("FileSystem")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
