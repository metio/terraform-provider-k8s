/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &SagemakerServicesK8SAwsProcessingJobV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsProcessingJobV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsProcessingJobV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsProcessingJobV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsProcessingJobV1Alpha1ManifestData struct {
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
		AppSpecification *struct {
			ContainerArguments  *[]string `tfsdk:"container_arguments" json:"containerArguments,omitempty"`
			ContainerEntrypoint *[]string `tfsdk:"container_entrypoint" json:"containerEntrypoint,omitempty"`
			ImageURI            *string   `tfsdk:"image_uri" json:"imageURI,omitempty"`
		} `tfsdk:"app_specification" json:"appSpecification,omitempty"`
		Environment      *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
		ExperimentConfig *struct {
			ExperimentName            *string `tfsdk:"experiment_name" json:"experimentName,omitempty"`
			TrialComponentDisplayName *string `tfsdk:"trial_component_display_name" json:"trialComponentDisplayName,omitempty"`
			TrialName                 *string `tfsdk:"trial_name" json:"trialName,omitempty"`
		} `tfsdk:"experiment_config" json:"experimentConfig,omitempty"`
		NetworkConfig *struct {
			EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" json:"enableInterContainerTrafficEncryption,omitempty"`
			EnableNetworkIsolation                *bool `tfsdk:"enable_network_isolation" json:"enableNetworkIsolation,omitempty"`
			VpcConfig                             *struct {
				SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
				Subnets          *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
			} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
		} `tfsdk:"network_config" json:"networkConfig,omitempty"`
		ProcessingInputs *[]struct {
			AppManaged        *bool `tfsdk:"app_managed" json:"appManaged,omitempty"`
			DatasetDefinition *struct {
				AthenaDatasetDefinition *struct {
					Catalog           *string `tfsdk:"catalog" json:"catalog,omitempty"`
					Database          *string `tfsdk:"database" json:"database,omitempty"`
					KmsKeyID          *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
					OutputCompression *string `tfsdk:"output_compression" json:"outputCompression,omitempty"`
					OutputFormat      *string `tfsdk:"output_format" json:"outputFormat,omitempty"`
					OutputS3URI       *string `tfsdk:"output_s3_uri" json:"outputS3URI,omitempty"`
					QueryString       *string `tfsdk:"query_string" json:"queryString,omitempty"`
					WorkGroup         *string `tfsdk:"work_group" json:"workGroup,omitempty"`
				} `tfsdk:"athena_dataset_definition" json:"athenaDatasetDefinition,omitempty"`
				DataDistributionType      *string `tfsdk:"data_distribution_type" json:"dataDistributionType,omitempty"`
				InputMode                 *string `tfsdk:"input_mode" json:"inputMode,omitempty"`
				LocalPath                 *string `tfsdk:"local_path" json:"localPath,omitempty"`
				RedshiftDatasetDefinition *struct {
					ClusterID         *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
					ClusterRoleARN    *string `tfsdk:"cluster_role_arn" json:"clusterRoleARN,omitempty"`
					Database          *string `tfsdk:"database" json:"database,omitempty"`
					DbUser            *string `tfsdk:"db_user" json:"dbUser,omitempty"`
					KmsKeyID          *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
					OutputCompression *string `tfsdk:"output_compression" json:"outputCompression,omitempty"`
					OutputFormat      *string `tfsdk:"output_format" json:"outputFormat,omitempty"`
					OutputS3URI       *string `tfsdk:"output_s3_uri" json:"outputS3URI,omitempty"`
					QueryString       *string `tfsdk:"query_string" json:"queryString,omitempty"`
				} `tfsdk:"redshift_dataset_definition" json:"redshiftDatasetDefinition,omitempty"`
			} `tfsdk:"dataset_definition" json:"datasetDefinition,omitempty"`
			InputName *string `tfsdk:"input_name" json:"inputName,omitempty"`
			S3Input   *struct {
				LocalPath              *string `tfsdk:"local_path" json:"localPath,omitempty"`
				S3CompressionType      *string `tfsdk:"s3_compression_type" json:"s3CompressionType,omitempty"`
				S3DataDistributionType *string `tfsdk:"s3_data_distribution_type" json:"s3DataDistributionType,omitempty"`
				S3DataType             *string `tfsdk:"s3_data_type" json:"s3DataType,omitempty"`
				S3InputMode            *string `tfsdk:"s3_input_mode" json:"s3InputMode,omitempty"`
				S3URI                  *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
			} `tfsdk:"s3_input" json:"s3Input,omitempty"`
		} `tfsdk:"processing_inputs" json:"processingInputs,omitempty"`
		ProcessingJobName      *string `tfsdk:"processing_job_name" json:"processingJobName,omitempty"`
		ProcessingOutputConfig *struct {
			KmsKeyID *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			Outputs  *[]struct {
				AppManaged         *bool `tfsdk:"app_managed" json:"appManaged,omitempty"`
				FeatureStoreOutput *struct {
					FeatureGroupName *string `tfsdk:"feature_group_name" json:"featureGroupName,omitempty"`
				} `tfsdk:"feature_store_output" json:"featureStoreOutput,omitempty"`
				OutputName *string `tfsdk:"output_name" json:"outputName,omitempty"`
				S3Output   *struct {
					LocalPath    *string `tfsdk:"local_path" json:"localPath,omitempty"`
					S3URI        *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
					S3UploadMode *string `tfsdk:"s3_upload_mode" json:"s3UploadMode,omitempty"`
				} `tfsdk:"s3_output" json:"s3Output,omitempty"`
			} `tfsdk:"outputs" json:"outputs,omitempty"`
		} `tfsdk:"processing_output_config" json:"processingOutputConfig,omitempty"`
		ProcessingResources *struct {
			ClusterConfig *struct {
				InstanceCount  *int64  `tfsdk:"instance_count" json:"instanceCount,omitempty"`
				InstanceType   *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
				VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" json:"volumeKMSKeyID,omitempty"`
				VolumeSizeInGB *int64  `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
			} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
		} `tfsdk:"processing_resources" json:"processingResources,omitempty"`
		RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		StoppingCondition *struct {
			MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
		} `tfsdk:"stopping_condition" json:"stoppingCondition,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_processing_job_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ProcessingJob is the Schema for the ProcessingJobs API",
		MarkdownDescription: "ProcessingJob is the Schema for the ProcessingJobs API",
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
				Description:         "ProcessingJobSpec defines the desired state of ProcessingJob.An Amazon SageMaker processing job that is used to analyze data and evaluatemodels. For more information, see Process Data and Evaluate Models (https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html).",
				MarkdownDescription: "ProcessingJobSpec defines the desired state of ProcessingJob.An Amazon SageMaker processing job that is used to analyze data and evaluatemodels. For more information, see Process Data and Evaluate Models (https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html).",
				Attributes: map[string]schema.Attribute{
					"app_specification": schema.SingleNestedAttribute{
						Description:         "Configures the processing job to run a specified Docker container image.",
						MarkdownDescription: "Configures the processing job to run a specified Docker container image.",
						Attributes: map[string]schema.Attribute{
							"container_arguments": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"container_entrypoint": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"environment": schema.MapAttribute{
						Description:         "The environment variables to set in the Docker container. Up to 100 key andvalues entries in the map are supported.",
						MarkdownDescription: "The environment variables to set in the Docker container. Up to 100 key andvalues entries in the map are supported.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"experiment_config": schema.SingleNestedAttribute{
						Description:         "Associates a SageMaker job as a trial component with an experiment and trial.Specified when you call the following APIs:   * CreateProcessingJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html)   * CreateTrainingJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html)   * CreateTransformJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html)",
						MarkdownDescription: "Associates a SageMaker job as a trial component with an experiment and trial.Specified when you call the following APIs:   * CreateProcessingJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html)   * CreateTrainingJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html)   * CreateTransformJob (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html)",
						Attributes: map[string]schema.Attribute{
							"experiment_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trial_component_display_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trial_name": schema.StringAttribute{
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

					"network_config": schema.SingleNestedAttribute{
						Description:         "Networking options for a processing job, such as whether to allow inboundand outbound network calls to and from processing containers, and the VPCsubnets and security groups to use for VPC-enabled processing jobs.",
						MarkdownDescription: "Networking options for a processing job, such as whether to allow inboundand outbound network calls to and from processing containers, and the VPCsubnets and security groups to use for VPC-enabled processing jobs.",
						Attributes: map[string]schema.Attribute{
							"enable_inter_container_traffic_encryption": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_network_isolation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vpc_config": schema.SingleNestedAttribute{
								Description:         "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs,hosted models, and compute resources have access to. You can control accessto and from your resources by configuring a VPC. For more information, seeGive SageMaker Access to Resources in your Amazon VPC (https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html).",
								MarkdownDescription: "Specifies an Amazon Virtual Private Cloud (VPC) that your SageMaker jobs,hosted models, and compute resources have access to. You can control accessto and from your resources by configuring a VPC. For more information, seeGive SageMaker Access to Resources in your Amazon VPC (https://docs.aws.amazon.com/sagemaker/latest/dg/infrastructure-give-access.html).",
								Attributes: map[string]schema.Attribute{
									"security_group_i_ds": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"subnets": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"processing_inputs": schema.ListNestedAttribute{
						Description:         "An array of inputs configuring the data to download into the processing container.",
						MarkdownDescription: "An array of inputs configuring the data to download into the processing container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"app_managed": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"dataset_definition": schema.SingleNestedAttribute{
									Description:         "Configuration for Dataset Definition inputs. The Dataset Definition inputmust specify exactly one of either AthenaDatasetDefinition or RedshiftDatasetDefinitiontypes.",
									MarkdownDescription: "Configuration for Dataset Definition inputs. The Dataset Definition inputmust specify exactly one of either AthenaDatasetDefinition or RedshiftDatasetDefinitiontypes.",
									Attributes: map[string]schema.Attribute{
										"athena_dataset_definition": schema.SingleNestedAttribute{
											Description:         "Configuration for Athena Dataset Definition input.",
											MarkdownDescription: "Configuration for Athena Dataset Definition input.",
											Attributes: map[string]schema.Attribute{
												"catalog": schema.StringAttribute{
													Description:         "The name of the data catalog used in Athena query execution.",
													MarkdownDescription: "The name of the data catalog used in Athena query execution.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"database": schema.StringAttribute{
													Description:         "The name of the database used in the Athena query execution.",
													MarkdownDescription: "The name of the database used in the Athena query execution.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kms_key_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"output_compression": schema.StringAttribute{
													Description:         "The compression used for Athena query results.",
													MarkdownDescription: "The compression used for Athena query results.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"output_format": schema.StringAttribute{
													Description:         "The data storage format for Athena query results.",
													MarkdownDescription: "The data storage format for Athena query results.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"output_s3_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"query_string": schema.StringAttribute{
													Description:         "The SQL query statements, to be executed.",
													MarkdownDescription: "The SQL query statements, to be executed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"work_group": schema.StringAttribute{
													Description:         "The name of the workgroup in which the Athena query is being started.",
													MarkdownDescription: "The name of the workgroup in which the Athena query is being started.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"data_distribution_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"input_mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"local_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"redshift_dataset_definition": schema.SingleNestedAttribute{
											Description:         "Configuration for Redshift Dataset Definition input.",
											MarkdownDescription: "Configuration for Redshift Dataset Definition input.",
											Attributes: map[string]schema.Attribute{
												"cluster_id": schema.StringAttribute{
													Description:         "The Redshift cluster Identifier.",
													MarkdownDescription: "The Redshift cluster Identifier.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cluster_role_arn": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"database": schema.StringAttribute{
													Description:         "The name of the Redshift database used in Redshift query execution.",
													MarkdownDescription: "The name of the Redshift database used in Redshift query execution.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"db_user": schema.StringAttribute{
													Description:         "The database user name used in Redshift query execution.",
													MarkdownDescription: "The database user name used in Redshift query execution.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kms_key_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"output_compression": schema.StringAttribute{
													Description:         "The compression used for Redshift query results.",
													MarkdownDescription: "The compression used for Redshift query results.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"output_format": schema.StringAttribute{
													Description:         "The data storage format for Redshift query results.",
													MarkdownDescription: "The data storage format for Redshift query results.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"output_s3_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"query_string": schema.StringAttribute{
													Description:         "The SQL query statements to be executed.",
													MarkdownDescription: "The SQL query statements to be executed.",
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

								"input_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"s3_input": schema.SingleNestedAttribute{
									Description:         "Configuration for downloading input data from Amazon S3 into the processingcontainer.",
									MarkdownDescription: "Configuration for downloading input data from Amazon S3 into the processingcontainer.",
									Attributes: map[string]schema.Attribute{
										"local_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_compression_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_data_distribution_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_data_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_input_mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_uri": schema.StringAttribute{
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

					"processing_job_name": schema.StringAttribute{
						Description:         "The name of the processing job. The name must be unique within an AmazonWeb Services Region in the Amazon Web Services account.",
						MarkdownDescription: "The name of the processing job. The name must be unique within an AmazonWeb Services Region in the Amazon Web Services account.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"processing_output_config": schema.SingleNestedAttribute{
						Description:         "Output configuration for the processing job.",
						MarkdownDescription: "Output configuration for the processing job.",
						Attributes: map[string]schema.Attribute{
							"kms_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"outputs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"app_managed": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"feature_store_output": schema.SingleNestedAttribute{
											Description:         "Configuration for processing job outputs in Amazon SageMaker Feature Store.",
											MarkdownDescription: "Configuration for processing job outputs in Amazon SageMaker Feature Store.",
											Attributes: map[string]schema.Attribute{
												"feature_group_name": schema.StringAttribute{
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

										"output_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_output": schema.SingleNestedAttribute{
											Description:         "Configuration for uploading output data to Amazon S3 from the processingcontainer.",
											MarkdownDescription: "Configuration for uploading output data to Amazon S3 from the processingcontainer.",
											Attributes: map[string]schema.Attribute{
												"local_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"s3_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"s3_upload_mode": schema.StringAttribute{
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

					"processing_resources": schema.SingleNestedAttribute{
						Description:         "Identifies the resources, ML compute instances, and ML storage volumes todeploy for a processing job. In distributed training, you specify more thanone instance.",
						MarkdownDescription: "Identifies the resources, ML compute instances, and ML storage volumes todeploy for a processing job. In distributed training, you specify more thanone instance.",
						Attributes: map[string]schema.Attribute{
							"cluster_config": schema.SingleNestedAttribute{
								Description:         "Configuration for the cluster used to run a processing job.",
								MarkdownDescription: "Configuration for the cluster used to run a processing job.",
								Attributes: map[string]schema.Attribute{
									"instance_count": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"instance_type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_kms_key_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_size_in_gb": schema.Int64Attribute{
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assumeto perform tasks on your behalf.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assumeto perform tasks on your behalf.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"stopping_condition": schema.SingleNestedAttribute{
						Description:         "The time limit for how long the processing job is allowed to run.",
						MarkdownDescription: "The time limit for how long the processing job is allowed to run.",
						Attributes: map[string]schema.Attribute{
							"max_runtime_in_seconds": schema.Int64Attribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "(Optional) An array of key-value pairs. For more information, see Using CostAllocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)in the Amazon Web Services Billing and Cost Management User Guide.",
						MarkdownDescription: "(Optional) An array of key-value pairs. For more information, see Using CostAllocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL)in the Amazon Web Services Billing and Cost Management User Guide.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_processing_job_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsProcessingJobV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("ProcessingJob")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
