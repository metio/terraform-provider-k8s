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

type SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsTrainingJobV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsTrainingJobV1Alpha1GoModel struct {
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
		AlgorithmSpecification *struct {
			AlgorithmName *string `tfsdk:"algorithm_name" yaml:"algorithmName,omitempty"`

			EnableSageMakerMetricsTimeSeries *bool `tfsdk:"enable_sage_maker_metrics_time_series" yaml:"enableSageMakerMetricsTimeSeries,omitempty"`

			MetricDefinitions *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
			} `tfsdk:"metric_definitions" yaml:"metricDefinitions,omitempty"`

			TrainingImage *string `tfsdk:"training_image" yaml:"trainingImage,omitempty"`

			TrainingInputMode *string `tfsdk:"training_input_mode" yaml:"trainingInputMode,omitempty"`
		} `tfsdk:"algorithm_specification" yaml:"algorithmSpecification,omitempty"`

		CheckpointConfig *struct {
			LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

			S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
		} `tfsdk:"checkpoint_config" yaml:"checkpointConfig,omitempty"`

		DebugHookConfig *struct {
			CollectionConfigurations *[]struct {
				CollectionName *string `tfsdk:"collection_name" yaml:"collectionName,omitempty"`

				CollectionParameters *map[string]string `tfsdk:"collection_parameters" yaml:"collectionParameters,omitempty"`
			} `tfsdk:"collection_configurations" yaml:"collectionConfigurations,omitempty"`

			HookParameters *map[string]string `tfsdk:"hook_parameters" yaml:"hookParameters,omitempty"`

			LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

			S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`
		} `tfsdk:"debug_hook_config" yaml:"debugHookConfig,omitempty"`

		DebugRuleConfigurations *[]struct {
			InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

			LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

			RuleConfigurationName *string `tfsdk:"rule_configuration_name" yaml:"ruleConfigurationName,omitempty"`

			RuleEvaluatorImage *string `tfsdk:"rule_evaluator_image" yaml:"ruleEvaluatorImage,omitempty"`

			RuleParameters *map[string]string `tfsdk:"rule_parameters" yaml:"ruleParameters,omitempty"`

			S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`

			VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
		} `tfsdk:"debug_rule_configurations" yaml:"debugRuleConfigurations,omitempty"`

		EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" yaml:"enableInterContainerTrafficEncryption,omitempty"`

		EnableManagedSpotTraining *bool `tfsdk:"enable_managed_spot_training" yaml:"enableManagedSpotTraining,omitempty"`

		EnableNetworkIsolation *bool `tfsdk:"enable_network_isolation" yaml:"enableNetworkIsolation,omitempty"`

		Environment *map[string]string `tfsdk:"environment" yaml:"environment,omitempty"`

		ExperimentConfig *struct {
			ExperimentName *string `tfsdk:"experiment_name" yaml:"experimentName,omitempty"`

			TrialComponentDisplayName *string `tfsdk:"trial_component_display_name" yaml:"trialComponentDisplayName,omitempty"`

			TrialName *string `tfsdk:"trial_name" yaml:"trialName,omitempty"`
		} `tfsdk:"experiment_config" yaml:"experimentConfig,omitempty"`

		HyperParameters *map[string]string `tfsdk:"hyper_parameters" yaml:"hyperParameters,omitempty"`

		InputDataConfig *[]struct {
			ChannelName *string `tfsdk:"channel_name" yaml:"channelName,omitempty"`

			CompressionType *string `tfsdk:"compression_type" yaml:"compressionType,omitempty"`

			ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

			DataSource *struct {
				FileSystemDataSource *struct {
					DirectoryPath *string `tfsdk:"directory_path" yaml:"directoryPath,omitempty"`

					FileSystemAccessMode *string `tfsdk:"file_system_access_mode" yaml:"fileSystemAccessMode,omitempty"`

					FileSystemID *string `tfsdk:"file_system_id" yaml:"fileSystemID,omitempty"`

					FileSystemType *string `tfsdk:"file_system_type" yaml:"fileSystemType,omitempty"`
				} `tfsdk:"file_system_data_source" yaml:"fileSystemDataSource,omitempty"`

				S3DataSource *struct {
					AttributeNames *[]string `tfsdk:"attribute_names" yaml:"attributeNames,omitempty"`

					S3DataDistributionType *string `tfsdk:"s3_data_distribution_type" yaml:"s3DataDistributionType,omitempty"`

					S3DataType *string `tfsdk:"s3_data_type" yaml:"s3DataType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"s3_data_source" yaml:"s3DataSource,omitempty"`
			} `tfsdk:"data_source" yaml:"dataSource,omitempty"`

			InputMode *string `tfsdk:"input_mode" yaml:"inputMode,omitempty"`

			RecordWrapperType *string `tfsdk:"record_wrapper_type" yaml:"recordWrapperType,omitempty"`

			ShuffleConfig *struct {
				Seed *int64 `tfsdk:"seed" yaml:"seed,omitempty"`
			} `tfsdk:"shuffle_config" yaml:"shuffleConfig,omitempty"`
		} `tfsdk:"input_data_config" yaml:"inputDataConfig,omitempty"`

		OutputDataConfig *struct {
			KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

			S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`
		} `tfsdk:"output_data_config" yaml:"outputDataConfig,omitempty"`

		ProfilerConfig *struct {
			ProfilingIntervalInMilliseconds *int64 `tfsdk:"profiling_interval_in_milliseconds" yaml:"profilingIntervalInMilliseconds,omitempty"`

			ProfilingParameters *map[string]string `tfsdk:"profiling_parameters" yaml:"profilingParameters,omitempty"`

			S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`
		} `tfsdk:"profiler_config" yaml:"profilerConfig,omitempty"`

		ProfilerRuleConfigurations *[]struct {
			InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

			LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

			RuleConfigurationName *string `tfsdk:"rule_configuration_name" yaml:"ruleConfigurationName,omitempty"`

			RuleEvaluatorImage *string `tfsdk:"rule_evaluator_image" yaml:"ruleEvaluatorImage,omitempty"`

			RuleParameters *map[string]string `tfsdk:"rule_parameters" yaml:"ruleParameters,omitempty"`

			S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`

			VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
		} `tfsdk:"profiler_rule_configurations" yaml:"profilerRuleConfigurations,omitempty"`

		ResourceConfig *struct {
			InstanceCount *int64 `tfsdk:"instance_count" yaml:"instanceCount,omitempty"`

			InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

			VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" yaml:"volumeKMSKeyID,omitempty"`

			VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
		} `tfsdk:"resource_config" yaml:"resourceConfig,omitempty"`

		RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`

		StoppingCondition *struct {
			MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" yaml:"maxRuntimeInSeconds,omitempty"`

			MaxWaitTimeInSeconds *int64 `tfsdk:"max_wait_time_in_seconds" yaml:"maxWaitTimeInSeconds,omitempty"`
		} `tfsdk:"stopping_condition" yaml:"stoppingCondition,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		TensorBoardOutputConfig *struct {
			LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

			S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`
		} `tfsdk:"tensor_board_output_config" yaml:"tensorBoardOutputConfig,omitempty"`

		TrainingJobName *string `tfsdk:"training_job_name" yaml:"trainingJobName,omitempty"`

		VpcConfig *struct {
			SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

			Subnets *[]string `tfsdk:"subnets" yaml:"subnets,omitempty"`
		} `tfsdk:"vpc_config" yaml:"vpcConfig,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsTrainingJobV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_training_job_v1alpha1"
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "TrainingJob is the Schema for the TrainingJobs API",
		MarkdownDescription: "TrainingJob is the Schema for the TrainingJobs API",
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
				Description:         "TrainingJobSpec defines the desired state of TrainingJob.  Contains information about a training job.",
				MarkdownDescription: "TrainingJobSpec defines the desired state of TrainingJob.  Contains information about a training job.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"algorithm_specification": {
						Description:         "The registry path of the Docker image that contains the training algorithm and algorithm-specific metadata, including the input mode. For more information about algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html). For information about providing your own algorithms, see Using Your Own Algorithms with Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",
						MarkdownDescription: "The registry path of the Docker image that contains the training algorithm and algorithm-specific metadata, including the input mode. For more information about algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html). For information about providing your own algorithms, see Using Your Own Algorithms with Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"algorithm_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_sage_maker_metrics_time_series": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metric_definitions": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"regex": {
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

							"training_image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"training_input_mode": {
								Description:         "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",
								MarkdownDescription: "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",

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

					"checkpoint_config": {
						Description:         "Contains information about the output location for managed spot training checkpoint data.",
						MarkdownDescription: "Contains information about the output location for managed spot training checkpoint data.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"debug_hook_config": {
						Description:         "Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the DebugHookConfig parameter, see Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html).",
						MarkdownDescription: "Configuration information for the Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the DebugHookConfig parameter, see Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"collection_configurations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"collection_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"collection_parameters": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hook_parameters": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

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

							"s3_output_path": {
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

					"debug_rule_configurations": {
						Description:         "Configuration information for Debugger rules for debugging output tensors.",
						MarkdownDescription: "Configuration information for Debugger rules for debugging output tensors.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"instance_type": {
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

							"rule_configuration_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rule_evaluator_image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rule_parameters": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_output_path": {
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

					"enable_inter_container_traffic_encryption": {
						Description:         "To encrypt all communications between ML compute instances in distributed training, choose True. Encryption provides greater security for distributed training, but training might take longer. How long it takes depends on the amount of communication between compute instances, especially if you use a deep learning algorithm in distributed training. For more information, see Protect Communications Between ML Compute Instances in a Distributed Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).",
						MarkdownDescription: "To encrypt all communications between ML compute instances in distributed training, choose True. Encryption provides greater security for distributed training, but training might take longer. How long it takes depends on the amount of communication between compute instances, especially if you use a deep learning algorithm in distributed training. For more information, see Protect Communications Between ML Compute Instances in a Distributed Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_managed_spot_training": {
						Description:         "To train models using managed spot training, choose True. Managed spot training provides a fully managed and scalable infrastructure for training machine learning models. this option is useful when training jobs can be interrupted and when there is flexibility when the training job is run.  The complete and intermediate results of jobs are stored in an Amazon S3 bucket, and can be used as a starting point to train models incrementally. Amazon SageMaker provides metrics and logs in CloudWatch. They can be used to see when managed spot training jobs are running, interrupted, resumed, or completed.",
						MarkdownDescription: "To train models using managed spot training, choose True. Managed spot training provides a fully managed and scalable infrastructure for training machine learning models. this option is useful when training jobs can be interrupted and when there is flexibility when the training job is run.  The complete and intermediate results of jobs are stored in an Amazon S3 bucket, and can be used as a starting point to train models incrementally. Amazon SageMaker provides metrics and logs in CloudWatch. They can be used to see when managed spot training jobs are running, interrupted, resumed, or completed.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_network_isolation": {
						Description:         "Isolates the training container. No inbound or outbound network calls can be made, except for calls between peers within a training cluster for distributed training. If you enable network isolation for training jobs that are configured to use a VPC, SageMaker downloads and uploads customer data and model artifacts through the specified VPC, but the training container does not have network access.",
						MarkdownDescription: "Isolates the training container. No inbound or outbound network calls can be made, except for calls between peers within a training cluster for distributed training. If you enable network isolation for training jobs that are configured to use a VPC, SageMaker downloads and uploads customer data and model artifacts through the specified VPC, but the training container does not have network access.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"environment": {
						Description:         "The environment variables to set in the Docker container.",
						MarkdownDescription: "The environment variables to set in the Docker container.",

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

					"hyper_parameters": {
						Description:         "Algorithm-specific parameters that influence the quality of the model. You set hyperparameters before you start the learning process. For a list of hyperparameters for each training algorithm provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  You can specify a maximum of 100 hyperparameters. Each hyperparameter is a key-value pair. Each key and value is limited to 256 characters, as specified by the Length Constraint.",
						MarkdownDescription: "Algorithm-specific parameters that influence the quality of the model. You set hyperparameters before you start the learning process. For a list of hyperparameters for each training algorithm provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  You can specify a maximum of 100 hyperparameters. Each hyperparameter is a key-value pair. Each key and value is limited to 256 characters, as specified by the Length Constraint.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"input_data_config": {
						Description:         "An array of Channel objects. Each channel is a named input source. InputDataConfig describes the input data and its location.  Algorithms can accept input data from one or more channels. For example, an algorithm might have two channels of input data, training_data and validation_data. The configuration for each channel provides the S3, EFS, or FSx location where the input data is stored. It also provides information about the stored data: the MIME type, compression method, and whether the data is wrapped in RecordIO format.  Depending on the input mode that the algorithm supports, SageMaker either copies input data files from an S3 bucket to a local directory in the Docker container, or makes it available as input streams. For example, if you specify an EFS location, input data files are available as input streams. They do not need to be downloaded.",
						MarkdownDescription: "An array of Channel objects. Each channel is a named input source. InputDataConfig describes the input data and its location.  Algorithms can accept input data from one or more channels. For example, an algorithm might have two channels of input data, training_data and validation_data. The configuration for each channel provides the S3, EFS, or FSx location where the input data is stored. It also provides information about the stored data: the MIME type, compression method, and whether the data is wrapped in RecordIO format.  Depending on the input mode that the algorithm supports, SageMaker either copies input data files from an S3 bucket to a local directory in the Docker container, or makes it available as input streams. For example, if you specify an EFS location, input data files are available as input streams. They do not need to be downloaded.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"channel_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"compression_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_source": {
								Description:         "Describes the location of the channel data.",
								MarkdownDescription: "Describes the location of the channel data.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"file_system_data_source": {
										Description:         "Specifies a file system data source for a channel.",
										MarkdownDescription: "Specifies a file system data source for a channel.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"directory_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_system_access_mode": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_system_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"file_system_type": {
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

									"s3_data_source": {
										Description:         "Describes the S3 data source.",
										MarkdownDescription: "Describes the S3 data source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"attribute_names": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

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

							"input_mode": {
								Description:         "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",
								MarkdownDescription: "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"record_wrapper_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"shuffle_config": {
								Description:         "A configuration for a shuffle option for input data in a channel. If you use S3Prefix for S3DataType, the results of the S3 key prefix matches are shuffled. If you use ManifestFile, the order of the S3 object references in the ManifestFile is shuffled. If you use AugmentedManifestFile, the order of the JSON lines in the AugmentedManifestFile is shuffled. The shuffling order is determined using the Seed value.  For Pipe input mode, when ShuffleConfig is specified shuffling is done at the start of every epoch. With large datasets, this ensures that the order of the training data is different for each epoch, and it helps reduce bias and possible overfitting. In a multi-node training job when ShuffleConfig is combined with S3DataDistributionType of ShardedByS3Key, the data is shuffled across nodes so that the content sent to a particular node on the first epoch might be sent to a different node on the second epoch.",
								MarkdownDescription: "A configuration for a shuffle option for input data in a channel. If you use S3Prefix for S3DataType, the results of the S3 key prefix matches are shuffled. If you use ManifestFile, the order of the S3 object references in the ManifestFile is shuffled. If you use AugmentedManifestFile, the order of the JSON lines in the AugmentedManifestFile is shuffled. The shuffling order is determined using the Seed value.  For Pipe input mode, when ShuffleConfig is specified shuffling is done at the start of every epoch. With large datasets, this ensures that the order of the training data is different for each epoch, and it helps reduce bias and possible overfitting. In a multi-node training job when ShuffleConfig is combined with S3DataDistributionType of ShardedByS3Key, the data is shuffled across nodes so that the content sent to a particular node on the first epoch might be sent to a different node on the second epoch.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"seed": {
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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"output_data_config": {
						Description:         "Specifies the path to the S3 location where you want to store model artifacts. SageMaker creates subfolders for the artifacts.",
						MarkdownDescription: "Specifies the path to the S3 location where you want to store model artifacts. SageMaker creates subfolders for the artifacts.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"kms_key_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_output_path": {
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

					"profiler_config": {
						Description:         "Configuration information for Debugger system monitoring, framework profiling, and storage paths.",
						MarkdownDescription: "Configuration information for Debugger system monitoring, framework profiling, and storage paths.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"profiling_interval_in_milliseconds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"profiling_parameters": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_output_path": {
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

					"profiler_rule_configurations": {
						Description:         "Configuration information for Debugger rules for profiling system and framework metrics.",
						MarkdownDescription: "Configuration information for Debugger rules for profiling system and framework metrics.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"instance_type": {
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

							"rule_configuration_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rule_evaluator_image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rule_parameters": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_output_path": {
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

					"resource_config": {
						Description:         "The resources, including the ML compute instances and ML storage volumes, to use for model training.  ML storage volumes store model artifacts and incremental states. Training algorithms might also use ML storage volumes for scratch space. If you want SageMaker to use the ML storage volume to store the training data, choose File as the TrainingInputMode in the algorithm specification. For distributed training algorithms, specify an instance count greater than 1.",
						MarkdownDescription: "The resources, including the ML compute instances and ML storage volumes, to use for model training.  ML storage volumes store model artifacts and incremental states. Training algorithms might also use ML storage volumes for scratch space. If you want SageMaker to use the ML storage volume to store the training data, choose File as the TrainingInputMode in the algorithm specification. For distributed training algorithms, specify an instance count greater than 1.",

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"role_arn": {
						Description:         "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to perform tasks on your behalf.  During model training, SageMaker needs your permission to read input data from an S3 bucket, download a Docker image that contains training code, write model artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, and publish metrics to Amazon CloudWatch. You grant permissions for all of these tasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).  To be able to pass this role to SageMaker, the caller of this API must have the iam:PassRole permission.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to perform tasks on your behalf.  During model training, SageMaker needs your permission to read input data from an S3 bucket, download a Docker image that contains training code, write model artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, and publish metrics to Amazon CloudWatch. You grant permissions for all of these tasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).  To be able to pass this role to SageMaker, the caller of this API must have the iam:PassRole permission.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"stopping_condition": {
						Description:         "Specifies a limit to how long a model training job can run. It also specifies how long a managed Spot training job has to complete. When the job reaches the time limit, SageMaker ends the training job. Use this API to cap model training costs.  To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.",
						MarkdownDescription: "Specifies a limit to how long a model training job can run. It also specifies how long a managed Spot training job has to complete. When the job reaches the time limit, SageMaker ends the training job. Use this API to cap model training costs.  To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_runtime_in_seconds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_wait_time_in_seconds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": {
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",

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

					"tensor_board_output_config": {
						Description:         "Configuration of storage locations for the Debugger TensorBoard output data.",
						MarkdownDescription: "Configuration of storage locations for the Debugger TensorBoard output data.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"local_path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_output_path": {
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

					"training_job_name": {
						Description:         "The name of the training job. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account.",
						MarkdownDescription: "The name of the training job. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"vpc_config": {
						Description:         "A VpcConfig object that specifies the VPC that you want your training job to connect to. Control access to and from your training container by configuring the VPC. For more information, see Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
						MarkdownDescription: "A VpcConfig object that specifies the VPC that you want your training job to connect to. Control access to and from your training container by configuring the VPC. For more information, see Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",

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
		},
	}, nil
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_training_job_v1alpha1")

	var state SagemakerServicesK8SAwsTrainingJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsTrainingJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("TrainingJob")

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

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_training_job_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_training_job_v1alpha1")

	var state SagemakerServicesK8SAwsTrainingJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsTrainingJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("TrainingJob")

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

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_training_job_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
