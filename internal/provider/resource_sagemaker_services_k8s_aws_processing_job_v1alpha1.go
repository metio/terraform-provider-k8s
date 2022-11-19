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

type SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsProcessingJobV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsProcessingJobV1Alpha1GoModel struct {
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
		AppSpecification *struct {
			ContainerArguments *[]string `tfsdk:"container_arguments" yaml:"containerArguments,omitempty"`

			ContainerEntrypoint *[]string `tfsdk:"container_entrypoint" yaml:"containerEntrypoint,omitempty"`

			ImageURI *string `tfsdk:"image_uri" yaml:"imageURI,omitempty"`
		} `tfsdk:"app_specification" yaml:"appSpecification,omitempty"`

		Environment *map[string]string `tfsdk:"environment" yaml:"environment,omitempty"`

		ExperimentConfig *struct {
			ExperimentName *string `tfsdk:"experiment_name" yaml:"experimentName,omitempty"`

			TrialComponentDisplayName *string `tfsdk:"trial_component_display_name" yaml:"trialComponentDisplayName,omitempty"`

			TrialName *string `tfsdk:"trial_name" yaml:"trialName,omitempty"`
		} `tfsdk:"experiment_config" yaml:"experimentConfig,omitempty"`

		NetworkConfig *struct {
			EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" yaml:"enableInterContainerTrafficEncryption,omitempty"`

			EnableNetworkIsolation *bool `tfsdk:"enable_network_isolation" yaml:"enableNetworkIsolation,omitempty"`

			VpcConfig *struct {
				SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

				Subnets *[]string `tfsdk:"subnets" yaml:"subnets,omitempty"`
			} `tfsdk:"vpc_config" yaml:"vpcConfig,omitempty"`
		} `tfsdk:"network_config" yaml:"networkConfig,omitempty"`

		ProcessingInputs *[]struct {
			AppManaged *bool `tfsdk:"app_managed" yaml:"appManaged,omitempty"`

			DatasetDefinition *struct {
				AthenaDatasetDefinition *struct {
					Catalog *string `tfsdk:"catalog" yaml:"catalog,omitempty"`

					Database *string `tfsdk:"database" yaml:"database,omitempty"`

					KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

					OutputCompression *string `tfsdk:"output_compression" yaml:"outputCompression,omitempty"`

					OutputFormat *string `tfsdk:"output_format" yaml:"outputFormat,omitempty"`

					OutputS3URI *string `tfsdk:"output_s3_uri" yaml:"outputS3URI,omitempty"`

					QueryString *string `tfsdk:"query_string" yaml:"queryString,omitempty"`

					WorkGroup *string `tfsdk:"work_group" yaml:"workGroup,omitempty"`
				} `tfsdk:"athena_dataset_definition" yaml:"athenaDatasetDefinition,omitempty"`

				DataDistributionType *string `tfsdk:"data_distribution_type" yaml:"dataDistributionType,omitempty"`

				InputMode *string `tfsdk:"input_mode" yaml:"inputMode,omitempty"`

				LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

				RedshiftDatasetDefinition *struct {
					ClusterID *string `tfsdk:"cluster_id" yaml:"clusterID,omitempty"`

					ClusterRoleARN *string `tfsdk:"cluster_role_arn" yaml:"clusterRoleARN,omitempty"`

					Database *string `tfsdk:"database" yaml:"database,omitempty"`

					DbUser *string `tfsdk:"db_user" yaml:"dbUser,omitempty"`

					KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

					OutputCompression *string `tfsdk:"output_compression" yaml:"outputCompression,omitempty"`

					OutputFormat *string `tfsdk:"output_format" yaml:"outputFormat,omitempty"`

					OutputS3URI *string `tfsdk:"output_s3_uri" yaml:"outputS3URI,omitempty"`

					QueryString *string `tfsdk:"query_string" yaml:"queryString,omitempty"`
				} `tfsdk:"redshift_dataset_definition" yaml:"redshiftDatasetDefinition,omitempty"`
			} `tfsdk:"dataset_definition" yaml:"datasetDefinition,omitempty"`

			InputName *string `tfsdk:"input_name" yaml:"inputName,omitempty"`

			S3Input *struct {
				LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

				S3CompressionType *string `tfsdk:"s3_compression_type" yaml:"s3CompressionType,omitempty"`

				S3DataDistributionType *string `tfsdk:"s3_data_distribution_type" yaml:"s3DataDistributionType,omitempty"`

				S3DataType *string `tfsdk:"s3_data_type" yaml:"s3DataType,omitempty"`

				S3InputMode *string `tfsdk:"s3_input_mode" yaml:"s3InputMode,omitempty"`

				S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
			} `tfsdk:"s3_input" yaml:"s3Input,omitempty"`
		} `tfsdk:"processing_inputs" yaml:"processingInputs,omitempty"`

		ProcessingJobName *string `tfsdk:"processing_job_name" yaml:"processingJobName,omitempty"`

		ProcessingOutputConfig *struct {
			KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

			Outputs *[]struct {
				AppManaged *bool `tfsdk:"app_managed" yaml:"appManaged,omitempty"`

				FeatureStoreOutput *struct {
					FeatureGroupName *string `tfsdk:"feature_group_name" yaml:"featureGroupName,omitempty"`
				} `tfsdk:"feature_store_output" yaml:"featureStoreOutput,omitempty"`

				OutputName *string `tfsdk:"output_name" yaml:"outputName,omitempty"`

				S3Output *struct {
					LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`

					S3UploadMode *string `tfsdk:"s3_upload_mode" yaml:"s3UploadMode,omitempty"`
				} `tfsdk:"s3_output" yaml:"s3Output,omitempty"`
			} `tfsdk:"outputs" yaml:"outputs,omitempty"`
		} `tfsdk:"processing_output_config" yaml:"processingOutputConfig,omitempty"`

		ProcessingResources *struct {
			ClusterConfig *struct {
				InstanceCount *int64 `tfsdk:"instance_count" yaml:"instanceCount,omitempty"`

				InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

				VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" yaml:"volumeKMSKeyID,omitempty"`

				VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
			} `tfsdk:"cluster_config" yaml:"clusterConfig,omitempty"`
		} `tfsdk:"processing_resources" yaml:"processingResources,omitempty"`

		RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`

		StoppingCondition *struct {
			MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" yaml:"maxRuntimeInSeconds,omitempty"`
		} `tfsdk:"stopping_condition" yaml:"stoppingCondition,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsProcessingJobV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_processing_job_v1alpha1"
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ProcessingJob is the Schema for the ProcessingJobs API",
		MarkdownDescription: "ProcessingJob is the Schema for the ProcessingJobs API",
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
				Description:         "ProcessingJobSpec defines the desired state of ProcessingJob.  An Amazon SageMaker processing job that is used to analyze data and evaluate models. For more information, see Process Data and Evaluate Models (https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html).",
				MarkdownDescription: "ProcessingJobSpec defines the desired state of ProcessingJob.  An Amazon SageMaker processing job that is used to analyze data and evaluate models. For more information, see Process Data and Evaluate Models (https://docs.aws.amazon.com/sagemaker/latest/dg/processing-job.html).",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"app_specification": {
						Description:         "Configures the processing job to run a specified Docker container image.",
						MarkdownDescription: "Configures the processing job to run a specified Docker container image.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"container_arguments": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"container_entrypoint": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_uri": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"environment": {
						Description:         "The environment variables to set in the Docker container. Up to 100 key and values entries in the map are supported.",
						MarkdownDescription: "The environment variables to set in the Docker container. Up to 100 key and values entries in the map are supported.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"experiment_config": {
						Description:         "Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:  * CreateProcessingJob  * CreateTrainingJob  * CreateTransformJob",
						MarkdownDescription: "Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:  * CreateProcessingJob  * CreateTrainingJob  * CreateTransformJob",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"experiment_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"trial_component_display_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"trial_name": {
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

					"network_config": {
						Description:         "Networking options for a processing job, such as whether to allow inbound and outbound network calls to and from processing containers, and the VPC subnets and security groups to use for VPC-enabled processing jobs.",
						MarkdownDescription: "Networking options for a processing job, such as whether to allow inbound and outbound network calls to and from processing containers, and the VPC subnets and security groups to use for VPC-enabled processing jobs.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable_inter_container_traffic_encryption": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_network_isolation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vpc_config": {
								Description:         "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
								MarkdownDescription: "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"security_group_i_ds": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subnets": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

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

					"processing_inputs": {
						Description:         "An array of inputs configuring the data to download into the processing container.",
						MarkdownDescription: "An array of inputs configuring the data to download into the processing container.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"app_managed": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dataset_definition": {
								Description:         "Configuration for Dataset Definition inputs. The Dataset Definition input must specify exactly one of either AthenaDatasetDefinition or RedshiftDatasetDefinition types.",
								MarkdownDescription: "Configuration for Dataset Definition inputs. The Dataset Definition input must specify exactly one of either AthenaDatasetDefinition or RedshiftDatasetDefinition types.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"athena_dataset_definition": {
										Description:         "Configuration for Athena Dataset Definition input.",
										MarkdownDescription: "Configuration for Athena Dataset Definition input.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"catalog": {
												Description:         "The name of the data catalog used in Athena query execution.",
												MarkdownDescription: "The name of the data catalog used in Athena query execution.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"database": {
												Description:         "The name of the database used in the Athena query execution.",
												MarkdownDescription: "The name of the database used in the Athena query execution.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kms_key_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"output_compression": {
												Description:         "The compression used for Athena query results.",
												MarkdownDescription: "The compression used for Athena query results.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"output_format": {
												Description:         "The data storage format for Athena query results.",
												MarkdownDescription: "The data storage format for Athena query results.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"output_s3_uri": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_string": {
												Description:         "The SQL query statements, to be executed.",
												MarkdownDescription: "The SQL query statements, to be executed.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"work_group": {
												Description:         "The name of the workgroup in which the Athena query is being started.",
												MarkdownDescription: "The name of the workgroup in which the Athena query is being started.",

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

									"data_distribution_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"input_mode": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"local_path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"redshift_dataset_definition": {
										Description:         "Configuration for Redshift Dataset Definition input.",
										MarkdownDescription: "Configuration for Redshift Dataset Definition input.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_id": {
												Description:         "The Redshift cluster Identifier.",
												MarkdownDescription: "The Redshift cluster Identifier.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cluster_role_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"database": {
												Description:         "The name of the Redshift database used in Redshift query execution.",
												MarkdownDescription: "The name of the Redshift database used in Redshift query execution.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"db_user": {
												Description:         "The database user name used in Redshift query execution.",
												MarkdownDescription: "The database user name used in Redshift query execution.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kms_key_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"output_compression": {
												Description:         "The compression used for Redshift query results.",
												MarkdownDescription: "The compression used for Redshift query results.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"output_format": {
												Description:         "The data storage format for Redshift query results.",
												MarkdownDescription: "The data storage format for Redshift query results.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"output_s3_uri": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_string": {
												Description:         "The SQL query statements to be executed.",
												MarkdownDescription: "The SQL query statements to be executed.",

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

							"input_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_input": {
								Description:         "Configuration for downloading input data from Amazon S3 into the processing container.",
								MarkdownDescription: "Configuration for downloading input data from Amazon S3 into the processing container.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"local_path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3_compression_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3_data_distribution_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3_data_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3_input_mode": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3_uri": {
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

					"processing_job_name": {
						Description:         "The name of the processing job. The name must be unique within an Amazon Web Services Region in the Amazon Web Services account.",
						MarkdownDescription: "The name of the processing job. The name must be unique within an Amazon Web Services Region in the Amazon Web Services account.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"processing_output_config": {
						Description:         "Output configuration for the processing job.",
						MarkdownDescription: "Output configuration for the processing job.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"kms_key_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"outputs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"app_managed": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"feature_store_output": {
										Description:         "Configuration for processing job outputs in Amazon SageMaker Feature Store.",
										MarkdownDescription: "Configuration for processing job outputs in Amazon SageMaker Feature Store.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"feature_group_name": {
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

									"output_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3_output": {
										Description:         "Configuration for uploading output data to Amazon S3 from the processing container.",
										MarkdownDescription: "Configuration for uploading output data to Amazon S3 from the processing container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"local_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_upload_mode": {
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

					"processing_resources": {
						Description:         "Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job. In distributed training, you specify more than one instance.",
						MarkdownDescription: "Identifies the resources, ML compute instances, and ML storage volumes to deploy for a processing job. In distributed training, you specify more than one instance.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cluster_config": {
								Description:         "Configuration for the cluster used to run a processing job.",
								MarkdownDescription: "Configuration for the cluster used to run a processing job.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"instance_count": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"instance_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_kms_key_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_size_in_gb": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"role_arn": {
						Description:         "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"stopping_condition": {
						Description:         "The time limit for how long the processing job is allowed to run.",
						MarkdownDescription: "The time limit for how long the processing job is allowed to run.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_runtime_in_seconds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the Amazon Web Services Billing and Cost Management User Guide.",
						MarkdownDescription: "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the Amazon Web Services Billing and Cost Management User Guide.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_processing_job_v1alpha1")

	var state SagemakerServicesK8SAwsProcessingJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsProcessingJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ProcessingJob")

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

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_processing_job_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_processing_job_v1alpha1")

	var state SagemakerServicesK8SAwsProcessingJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsProcessingJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ProcessingJob")

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

func (r *SagemakerServicesK8SAwsProcessingJobV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_processing_job_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
