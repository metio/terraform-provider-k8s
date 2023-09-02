/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AlgorithmSpecification *struct {
			AlgorithmName                    *string `tfsdk:"algorithm_name" json:"algorithmName,omitempty"`
			EnableSageMakerMetricsTimeSeries *bool   `tfsdk:"enable_sage_maker_metrics_time_series" json:"enableSageMakerMetricsTimeSeries,omitempty"`
			MetricDefinitions                *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Regex *string `tfsdk:"regex" json:"regex,omitempty"`
			} `tfsdk:"metric_definitions" json:"metricDefinitions,omitempty"`
			TrainingImage     *string `tfsdk:"training_image" json:"trainingImage,omitempty"`
			TrainingInputMode *string `tfsdk:"training_input_mode" json:"trainingInputMode,omitempty"`
		} `tfsdk:"algorithm_specification" json:"algorithmSpecification,omitempty"`
		CheckpointConfig *struct {
			LocalPath *string `tfsdk:"local_path" json:"localPath,omitempty"`
			S3URI     *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
		} `tfsdk:"checkpoint_config" json:"checkpointConfig,omitempty"`
		DebugHookConfig *struct {
			CollectionConfigurations *[]struct {
				CollectionName       *string            `tfsdk:"collection_name" json:"collectionName,omitempty"`
				CollectionParameters *map[string]string `tfsdk:"collection_parameters" json:"collectionParameters,omitempty"`
			} `tfsdk:"collection_configurations" json:"collectionConfigurations,omitempty"`
			HookParameters *map[string]string `tfsdk:"hook_parameters" json:"hookParameters,omitempty"`
			LocalPath      *string            `tfsdk:"local_path" json:"localPath,omitempty"`
			S3OutputPath   *string            `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
		} `tfsdk:"debug_hook_config" json:"debugHookConfig,omitempty"`
		DebugRuleConfigurations *[]struct {
			InstanceType          *string            `tfsdk:"instance_type" json:"instanceType,omitempty"`
			LocalPath             *string            `tfsdk:"local_path" json:"localPath,omitempty"`
			RuleConfigurationName *string            `tfsdk:"rule_configuration_name" json:"ruleConfigurationName,omitempty"`
			RuleEvaluatorImage    *string            `tfsdk:"rule_evaluator_image" json:"ruleEvaluatorImage,omitempty"`
			RuleParameters        *map[string]string `tfsdk:"rule_parameters" json:"ruleParameters,omitempty"`
			S3OutputPath          *string            `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
			VolumeSizeInGB        *int64             `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
		} `tfsdk:"debug_rule_configurations" json:"debugRuleConfigurations,omitempty"`
		EnableInterContainerTrafficEncryption *bool              `tfsdk:"enable_inter_container_traffic_encryption" json:"enableInterContainerTrafficEncryption,omitempty"`
		EnableManagedSpotTraining             *bool              `tfsdk:"enable_managed_spot_training" json:"enableManagedSpotTraining,omitempty"`
		EnableNetworkIsolation                *bool              `tfsdk:"enable_network_isolation" json:"enableNetworkIsolation,omitempty"`
		Environment                           *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
		ExperimentConfig                      *struct {
			ExperimentName            *string `tfsdk:"experiment_name" json:"experimentName,omitempty"`
			TrialComponentDisplayName *string `tfsdk:"trial_component_display_name" json:"trialComponentDisplayName,omitempty"`
			TrialName                 *string `tfsdk:"trial_name" json:"trialName,omitempty"`
		} `tfsdk:"experiment_config" json:"experimentConfig,omitempty"`
		HyperParameters *map[string]string `tfsdk:"hyper_parameters" json:"hyperParameters,omitempty"`
		InputDataConfig *[]struct {
			ChannelName     *string `tfsdk:"channel_name" json:"channelName,omitempty"`
			CompressionType *string `tfsdk:"compression_type" json:"compressionType,omitempty"`
			ContentType     *string `tfsdk:"content_type" json:"contentType,omitempty"`
			DataSource      *struct {
				FileSystemDataSource *struct {
					DirectoryPath        *string `tfsdk:"directory_path" json:"directoryPath,omitempty"`
					FileSystemAccessMode *string `tfsdk:"file_system_access_mode" json:"fileSystemAccessMode,omitempty"`
					FileSystemID         *string `tfsdk:"file_system_id" json:"fileSystemID,omitempty"`
					FileSystemType       *string `tfsdk:"file_system_type" json:"fileSystemType,omitempty"`
				} `tfsdk:"file_system_data_source" json:"fileSystemDataSource,omitempty"`
				S3DataSource *struct {
					AttributeNames         *[]string `tfsdk:"attribute_names" json:"attributeNames,omitempty"`
					InstanceGroupNames     *[]string `tfsdk:"instance_group_names" json:"instanceGroupNames,omitempty"`
					S3DataDistributionType *string   `tfsdk:"s3_data_distribution_type" json:"s3DataDistributionType,omitempty"`
					S3DataType             *string   `tfsdk:"s3_data_type" json:"s3DataType,omitempty"`
					S3URI                  *string   `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"s3_data_source" json:"s3DataSource,omitempty"`
			} `tfsdk:"data_source" json:"dataSource,omitempty"`
			InputMode         *string `tfsdk:"input_mode" json:"inputMode,omitempty"`
			RecordWrapperType *string `tfsdk:"record_wrapper_type" json:"recordWrapperType,omitempty"`
			ShuffleConfig     *struct {
				Seed *int64 `tfsdk:"seed" json:"seed,omitempty"`
			} `tfsdk:"shuffle_config" json:"shuffleConfig,omitempty"`
		} `tfsdk:"input_data_config" json:"inputDataConfig,omitempty"`
		OutputDataConfig *struct {
			KmsKeyID     *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			S3OutputPath *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
		} `tfsdk:"output_data_config" json:"outputDataConfig,omitempty"`
		ProfilerConfig *struct {
			ProfilingIntervalInMilliseconds *int64             `tfsdk:"profiling_interval_in_milliseconds" json:"profilingIntervalInMilliseconds,omitempty"`
			ProfilingParameters             *map[string]string `tfsdk:"profiling_parameters" json:"profilingParameters,omitempty"`
			S3OutputPath                    *string            `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
		} `tfsdk:"profiler_config" json:"profilerConfig,omitempty"`
		ProfilerRuleConfigurations *[]struct {
			InstanceType          *string            `tfsdk:"instance_type" json:"instanceType,omitempty"`
			LocalPath             *string            `tfsdk:"local_path" json:"localPath,omitempty"`
			RuleConfigurationName *string            `tfsdk:"rule_configuration_name" json:"ruleConfigurationName,omitempty"`
			RuleEvaluatorImage    *string            `tfsdk:"rule_evaluator_image" json:"ruleEvaluatorImage,omitempty"`
			RuleParameters        *map[string]string `tfsdk:"rule_parameters" json:"ruleParameters,omitempty"`
			S3OutputPath          *string            `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
			VolumeSizeInGB        *int64             `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
		} `tfsdk:"profiler_rule_configurations" json:"profilerRuleConfigurations,omitempty"`
		ResourceConfig *struct {
			InstanceCount  *int64 `tfsdk:"instance_count" json:"instanceCount,omitempty"`
			InstanceGroups *[]struct {
				InstanceCount     *int64  `tfsdk:"instance_count" json:"instanceCount,omitempty"`
				InstanceGroupName *string `tfsdk:"instance_group_name" json:"instanceGroupName,omitempty"`
				InstanceType      *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
			} `tfsdk:"instance_groups" json:"instanceGroups,omitempty"`
			InstanceType             *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
			KeepAlivePeriodInSeconds *int64  `tfsdk:"keep_alive_period_in_seconds" json:"keepAlivePeriodInSeconds,omitempty"`
			VolumeKMSKeyID           *string `tfsdk:"volume_kms_key_id" json:"volumeKMSKeyID,omitempty"`
			VolumeSizeInGB           *int64  `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
		} `tfsdk:"resource_config" json:"resourceConfig,omitempty"`
		RetryStrategy *struct {
			MaximumRetryAttempts *int64 `tfsdk:"maximum_retry_attempts" json:"maximumRetryAttempts,omitempty"`
		} `tfsdk:"retry_strategy" json:"retryStrategy,omitempty"`
		RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		StoppingCondition *struct {
			MaxRuntimeInSeconds  *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
			MaxWaitTimeInSeconds *int64 `tfsdk:"max_wait_time_in_seconds" json:"maxWaitTimeInSeconds,omitempty"`
		} `tfsdk:"stopping_condition" json:"stoppingCondition,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TensorBoardOutputConfig *struct {
			LocalPath    *string `tfsdk:"local_path" json:"localPath,omitempty"`
			S3OutputPath *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
		} `tfsdk:"tensor_board_output_config" json:"tensorBoardOutputConfig,omitempty"`
		TrainingJobName *string `tfsdk:"training_job_name" json:"trainingJobName,omitempty"`
		VpcConfig       *struct {
			SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
			Subnets          *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
		} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_training_job_v1alpha1"
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TrainingJob is the Schema for the TrainingJobs API",
		MarkdownDescription: "TrainingJob is the Schema for the TrainingJobs API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "TrainingJobSpec defines the desired state of TrainingJob.  Contains information about a training job.",
				MarkdownDescription: "TrainingJobSpec defines the desired state of TrainingJob.  Contains information about a training job.",
				Attributes: map[string]schema.Attribute{
					"algorithm_specification": schema.SingleNestedAttribute{
						Description:         "The registry path of the Docker image that contains the training algorithm and algorithm-specific metadata, including the input mode. For more information about algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html). For information about providing your own algorithms, see Using Your Own Algorithms with Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",
						MarkdownDescription: "The registry path of the Docker image that contains the training algorithm and algorithm-specific metadata, including the input mode. For more information about algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html). For information about providing your own algorithms, see Using Your Own Algorithms with Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",
						Attributes: map[string]schema.Attribute{
							"algorithm_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enable_sage_maker_metrics_time_series": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metric_definitions": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"regex": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"training_image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"training_input_mode": schema.StringAttribute{
								Description:         "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",
								MarkdownDescription: "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"checkpoint_config": schema.SingleNestedAttribute{
						Description:         "Contains information about the output location for managed spot training checkpoint data.",
						MarkdownDescription: "Contains information about the output location for managed spot training checkpoint data.",
						Attributes: map[string]schema.Attribute{
							"local_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"debug_hook_config": schema.SingleNestedAttribute{
						Description:         "Configuration information for the Amazon SageMaker Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the DebugHookConfig parameter, see Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html).",
						MarkdownDescription: "Configuration information for the Amazon SageMaker Debugger hook parameters, metric and tensor collections, and storage paths. To learn more about how to configure the DebugHookConfig parameter, see Use the SageMaker and Debugger Configuration API Operations to Create, Update, and Debug Your Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html).",
						Attributes: map[string]schema.Attribute{
							"collection_configurations": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"collection_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"collection_parameters": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"hook_parameters": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"local_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_output_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"debug_rule_configurations": schema.ListNestedAttribute{
						Description:         "Configuration information for Amazon SageMaker Debugger rules for debugging output tensors.",
						MarkdownDescription: "Configuration information for Amazon SageMaker Debugger rules for debugging output tensors.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"instance_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"local_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rule_configuration_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rule_evaluator_image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rule_parameters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"s3_output_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"volume_size_in_gb": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"enable_inter_container_traffic_encryption": schema.BoolAttribute{
						Description:         "To encrypt all communications between ML compute instances in distributed training, choose True. Encryption provides greater security for distributed training, but training might take longer. How long it takes depends on the amount of communication between compute instances, especially if you use a deep learning algorithm in distributed training. For more information, see Protect Communications Between ML Compute Instances in a Distributed Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).",
						MarkdownDescription: "To encrypt all communications between ML compute instances in distributed training, choose True. Encryption provides greater security for distributed training, but training might take longer. How long it takes depends on the amount of communication between compute instances, especially if you use a deep learning algorithm in distributed training. For more information, see Protect Communications Between ML Compute Instances in a Distributed Training Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_managed_spot_training": schema.BoolAttribute{
						Description:         "To train models using managed spot training, choose True. Managed spot training provides a fully managed and scalable infrastructure for training machine learning models. this option is useful when training jobs can be interrupted and when there is flexibility when the training job is run.  The complete and intermediate results of jobs are stored in an Amazon S3 bucket, and can be used as a starting point to train models incrementally. Amazon SageMaker provides metrics and logs in CloudWatch. They can be used to see when managed spot training jobs are running, interrupted, resumed, or completed.",
						MarkdownDescription: "To train models using managed spot training, choose True. Managed spot training provides a fully managed and scalable infrastructure for training machine learning models. this option is useful when training jobs can be interrupted and when there is flexibility when the training job is run.  The complete and intermediate results of jobs are stored in an Amazon S3 bucket, and can be used as a starting point to train models incrementally. Amazon SageMaker provides metrics and logs in CloudWatch. They can be used to see when managed spot training jobs are running, interrupted, resumed, or completed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_network_isolation": schema.BoolAttribute{
						Description:         "Isolates the training container. No inbound or outbound network calls can be made, except for calls between peers within a training cluster for distributed training. If you enable network isolation for training jobs that are configured to use a VPC, SageMaker downloads and uploads customer data and model artifacts through the specified VPC, but the training container does not have network access.",
						MarkdownDescription: "Isolates the training container. No inbound or outbound network calls can be made, except for calls between peers within a training cluster for distributed training. If you enable network isolation for training jobs that are configured to use a VPC, SageMaker downloads and uploads customer data and model artifacts through the specified VPC, but the training container does not have network access.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"environment": schema.MapAttribute{
						Description:         "The environment variables to set in the Docker container.",
						MarkdownDescription: "The environment variables to set in the Docker container.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"experiment_config": schema.SingleNestedAttribute{
						Description:         "Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:  * CreateProcessingJob  * CreateTrainingJob  * CreateTransformJob",
						MarkdownDescription: "Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:  * CreateProcessingJob  * CreateTrainingJob  * CreateTransformJob",
						Attributes: map[string]schema.Attribute{
							"experiment_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"trial_component_display_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"trial_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"hyper_parameters": schema.MapAttribute{
						Description:         "Algorithm-specific parameters that influence the quality of the model. You set hyperparameters before you start the learning process. For a list of hyperparameters for each training algorithm provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  You can specify a maximum of 100 hyperparameters. Each hyperparameter is a key-value pair. Each key and value is limited to 256 characters, as specified by the Length Constraint.  Do not include any security-sensitive information including account access IDs, secrets or tokens in any hyperparameter field. If the use of security-sensitive credentials are detected, SageMaker will reject your training job request and return an exception error.",
						MarkdownDescription: "Algorithm-specific parameters that influence the quality of the model. You set hyperparameters before you start the learning process. For a list of hyperparameters for each training algorithm provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  You can specify a maximum of 100 hyperparameters. Each hyperparameter is a key-value pair. Each key and value is limited to 256 characters, as specified by the Length Constraint.  Do not include any security-sensitive information including account access IDs, secrets or tokens in any hyperparameter field. If the use of security-sensitive credentials are detected, SageMaker will reject your training job request and return an exception error.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"input_data_config": schema.ListNestedAttribute{
						Description:         "An array of Channel objects. Each channel is a named input source. InputDataConfig describes the input data and its location.  Algorithms can accept input data from one or more channels. For example, an algorithm might have two channels of input data, training_data and validation_data. The configuration for each channel provides the S3, EFS, or FSx location where the input data is stored. It also provides information about the stored data: the MIME type, compression method, and whether the data is wrapped in RecordIO format.  Depending on the input mode that the algorithm supports, SageMaker either copies input data files from an S3 bucket to a local directory in the Docker container, or makes it available as input streams. For example, if you specify an EFS location, input data files are available as input streams. They do not need to be downloaded.",
						MarkdownDescription: "An array of Channel objects. Each channel is a named input source. InputDataConfig describes the input data and its location.  Algorithms can accept input data from one or more channels. For example, an algorithm might have two channels of input data, training_data and validation_data. The configuration for each channel provides the S3, EFS, or FSx location where the input data is stored. It also provides information about the stored data: the MIME type, compression method, and whether the data is wrapped in RecordIO format.  Depending on the input mode that the algorithm supports, SageMaker either copies input data files from an S3 bucket to a local directory in the Docker container, or makes it available as input streams. For example, if you specify an EFS location, input data files are available as input streams. They do not need to be downloaded.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"channel_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"compression_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"data_source": schema.SingleNestedAttribute{
									Description:         "Describes the location of the channel data.",
									MarkdownDescription: "Describes the location of the channel data.",
									Attributes: map[string]schema.Attribute{
										"file_system_data_source": schema.SingleNestedAttribute{
											Description:         "Specifies a file system data source for a channel.",
											MarkdownDescription: "Specifies a file system data source for a channel.",
											Attributes: map[string]schema.Attribute{
												"directory_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"file_system_access_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"file_system_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"file_system_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"s3_data_source": schema.SingleNestedAttribute{
											Description:         "Describes the S3 data source.",
											MarkdownDescription: "Describes the S3 data source.",
											Attributes: map[string]schema.Attribute{
												"attribute_names": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"instance_group_names": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"s3_data_distribution_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"s3_data_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"s3_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"input_mode": schema.StringAttribute{
									Description:         "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",
									MarkdownDescription: "The training input mode that the algorithm supports. For more information about input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).  Pipe mode  If an algorithm supports Pipe mode, Amazon SageMaker streams data directly from Amazon S3 to the container.  File mode  If an algorithm supports File mode, SageMaker downloads the training data from S3 to the provisioned ML storage volume, and mounts the directory to the Docker volume for the training container.  You must provision the ML storage volume with sufficient capacity to accommodate the data downloaded from S3. In addition to the training data, the ML storage volume also stores the output model. The algorithm container uses the ML storage volume to also store intermediate information, if any.  For distributed algorithms, training data is distributed uniformly. Your training duration is predictable if the input data objects sizes are approximately the same. SageMaker does not split the files any further for model training. If the object sizes are skewed, training won't be optimal as the data distribution is also skewed when one host in a training cluster is overloaded, thus becoming a bottleneck in training.  FastFile mode  If an algorithm supports FastFile mode, SageMaker streams data directly from S3 to the container with no code changes, and provides file system access to the data. Users can author their training script to interact with these files as if they were stored on disk.  FastFile mode works best when the data is read sequentially. Augmented manifest files aren't supported. The startup time is lower when there are fewer files in the S3 bucket provided.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"record_wrapper_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"shuffle_config": schema.SingleNestedAttribute{
									Description:         "A configuration for a shuffle option for input data in a channel. If you use S3Prefix for S3DataType, the results of the S3 key prefix matches are shuffled. If you use ManifestFile, the order of the S3 object references in the ManifestFile is shuffled. If you use AugmentedManifestFile, the order of the JSON lines in the AugmentedManifestFile is shuffled. The shuffling order is determined using the Seed value.  For Pipe input mode, when ShuffleConfig is specified shuffling is done at the start of every epoch. With large datasets, this ensures that the order of the training data is different for each epoch, and it helps reduce bias and possible overfitting. In a multi-node training job when ShuffleConfig is combined with S3DataDistributionType of ShardedByS3Key, the data is shuffled across nodes so that the content sent to a particular node on the first epoch might be sent to a different node on the second epoch.",
									MarkdownDescription: "A configuration for a shuffle option for input data in a channel. If you use S3Prefix for S3DataType, the results of the S3 key prefix matches are shuffled. If you use ManifestFile, the order of the S3 object references in the ManifestFile is shuffled. If you use AugmentedManifestFile, the order of the JSON lines in the AugmentedManifestFile is shuffled. The shuffling order is determined using the Seed value.  For Pipe input mode, when ShuffleConfig is specified shuffling is done at the start of every epoch. With large datasets, this ensures that the order of the training data is different for each epoch, and it helps reduce bias and possible overfitting. In a multi-node training job when ShuffleConfig is combined with S3DataDistributionType of ShardedByS3Key, the data is shuffled across nodes so that the content sent to a particular node on the first epoch might be sent to a different node on the second epoch.",
									Attributes: map[string]schema.Attribute{
										"seed": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"output_data_config": schema.SingleNestedAttribute{
						Description:         "Specifies the path to the S3 location where you want to store model artifacts. SageMaker creates subfolders for the artifacts.",
						MarkdownDescription: "Specifies the path to the S3 location where you want to store model artifacts. SageMaker creates subfolders for the artifacts.",
						Attributes: map[string]schema.Attribute{
							"kms_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_output_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"profiler_config": schema.SingleNestedAttribute{
						Description:         "Configuration information for Amazon SageMaker Debugger system monitoring, framework profiling, and storage paths.",
						MarkdownDescription: "Configuration information for Amazon SageMaker Debugger system monitoring, framework profiling, and storage paths.",
						Attributes: map[string]schema.Attribute{
							"profiling_interval_in_milliseconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"profiling_parameters": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_output_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"profiler_rule_configurations": schema.ListNestedAttribute{
						Description:         "Configuration information for Amazon SageMaker Debugger rules for profiling system and framework metrics.",
						MarkdownDescription: "Configuration information for Amazon SageMaker Debugger rules for profiling system and framework metrics.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"instance_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"local_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rule_configuration_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rule_evaluator_image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rule_parameters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"s3_output_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"volume_size_in_gb": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"resource_config": schema.SingleNestedAttribute{
						Description:         "The resources, including the ML compute instances and ML storage volumes, to use for model training.  ML storage volumes store model artifacts and incremental states. Training algorithms might also use ML storage volumes for scratch space. If you want SageMaker to use the ML storage volume to store the training data, choose File as the TrainingInputMode in the algorithm specification. For distributed training algorithms, specify an instance count greater than 1.",
						MarkdownDescription: "The resources, including the ML compute instances and ML storage volumes, to use for model training.  ML storage volumes store model artifacts and incremental states. Training algorithms might also use ML storage volumes for scratch space. If you want SageMaker to use the ML storage volume to store the training data, choose File as the TrainingInputMode in the algorithm specification. For distributed training algorithms, specify an instance count greater than 1.",
						Attributes: map[string]schema.Attribute{
							"instance_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"instance_groups": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"instance_count": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"instance_group_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"instance_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"instance_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"keep_alive_period_in_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"volume_kms_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"volume_size_in_gb": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"retry_strategy": schema.SingleNestedAttribute{
						Description:         "The number of times to retry the job when the job fails due to an InternalServerError.",
						MarkdownDescription: "The number of times to retry the job when the job fails due to an InternalServerError.",
						Attributes: map[string]schema.Attribute{
							"maximum_retry_attempts": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to perform tasks on your behalf.  During model training, SageMaker needs your permission to read input data from an S3 bucket, download a Docker image that contains training code, write model artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, and publish metrics to Amazon CloudWatch. You grant permissions for all of these tasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).  To be able to pass this role to SageMaker, the caller of this API must have the iam:PassRole permission.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume to perform tasks on your behalf.  During model training, SageMaker needs your permission to read input data from an S3 bucket, download a Docker image that contains training code, write model artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, and publish metrics to Amazon CloudWatch. You grant permissions for all of these tasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).  To be able to pass this role to SageMaker, the caller of this API must have the iam:PassRole permission.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"stopping_condition": schema.SingleNestedAttribute{
						Description:         "Specifies a limit to how long a model training job can run. It also specifies how long a managed Spot training job has to complete. When the job reaches the time limit, SageMaker ends the training job. Use this API to cap model training costs.  To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.",
						MarkdownDescription: "Specifies a limit to how long a model training job can run. It also specifies how long a managed Spot training job has to complete. When the job reaches the time limit, SageMaker ends the training job. Use this API to cap model training costs.  To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.",
						Attributes: map[string]schema.Attribute{
							"max_runtime_in_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_wait_time_in_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tensor_board_output_config": schema.SingleNestedAttribute{
						Description:         "Configuration of storage locations for the Amazon SageMaker Debugger TensorBoard output data.",
						MarkdownDescription: "Configuration of storage locations for the Amazon SageMaker Debugger TensorBoard output data.",
						Attributes: map[string]schema.Attribute{
							"local_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_output_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"training_job_name": schema.StringAttribute{
						Description:         "The name of the training job. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account.",
						MarkdownDescription: "The name of the training job. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"vpc_config": schema.SingleNestedAttribute{
						Description:         "A VpcConfig object that specifies the VPC that you want your training job to connect to. Control access to and from your training container by configuring the VPC. For more information, see Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
						MarkdownDescription: "A VpcConfig object that specifies the VPC that you want your training job to connect to. Control access to and from your training container by configuring the VPC. For more information, see Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
						Attributes: map[string]schema.Attribute{
							"security_group_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"subnets": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_training_job_v1alpha1")

	var data SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "TrainingJob"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse SagemakerServicesK8SAwsTrainingJobV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("TrainingJob")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
