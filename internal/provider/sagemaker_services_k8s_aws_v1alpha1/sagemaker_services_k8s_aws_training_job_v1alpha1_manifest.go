/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &SagemakerServicesK8SAwsTrainingJobV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsTrainingJobV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsTrainingJobV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsTrainingJobV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsTrainingJobV1Alpha1ManifestData struct {
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
		HyperParameters  *map[string]string `tfsdk:"hyper_parameters" json:"hyperParameters,omitempty"`
		InfraCheckConfig *struct {
			EnableInfraCheck *bool `tfsdk:"enable_infra_check" json:"enableInfraCheck,omitempty"`
		} `tfsdk:"infra_check_config" json:"infraCheckConfig,omitempty"`
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
			CompressionType *string `tfsdk:"compression_type" json:"compressionType,omitempty"`
			KmsKeyID        *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			S3OutputPath    *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
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
		RemoteDebugConfig *struct {
			EnableRemoteDebug *bool `tfsdk:"enable_remote_debug" json:"enableRemoteDebug,omitempty"`
		} `tfsdk:"remote_debug_config" json:"remoteDebugConfig,omitempty"`
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
			MaxPendingTimeInSeconds *int64 `tfsdk:"max_pending_time_in_seconds" json:"maxPendingTimeInSeconds,omitempty"`
			MaxRuntimeInSeconds     *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
			MaxWaitTimeInSeconds    *int64 `tfsdk:"max_wait_time_in_seconds" json:"maxWaitTimeInSeconds,omitempty"`
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

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_training_job_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "TrainingJobSpec defines the desired state of TrainingJob.Contains information about a training job.",
				MarkdownDescription: "TrainingJobSpec defines the desired state of TrainingJob.Contains information about a training job.",
				Attributes: map[string]schema.Attribute{
					"algorithm_specification": schema.SingleNestedAttribute{
						Description:         "The registry path of the Docker image that contains the training algorithmand algorithm-specific metadata, including the input mode. For more informationabout algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).For information about providing your own algorithms, see Using Your Own Algorithmswith Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",
						MarkdownDescription: "The registry path of the Docker image that contains the training algorithmand algorithm-specific metadata, including the input mode. For more informationabout algorithms provided by SageMaker, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).For information about providing your own algorithms, see Using Your Own Algorithmswith Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms.html).",
						Attributes: map[string]schema.Attribute{
							"algorithm_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_sage_maker_metrics_time_series": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"regex": schema.StringAttribute{
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

							"training_image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"training_input_mode": schema.StringAttribute{
								Description:         "The training input mode that the algorithm supports. For more informationabout input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).Pipe modeIf an algorithm supports Pipe mode, Amazon SageMaker streams data directlyfrom Amazon S3 to the container.File modeIf an algorithm supports File mode, SageMaker downloads the training datafrom S3 to the provisioned ML storage volume, and mounts the directory tothe Docker volume for the training container.You must provision the ML storage volume with sufficient capacity to accommodatethe data downloaded from S3. In addition to the training data, the ML storagevolume also stores the output model. The algorithm container uses the MLstorage volume to also store intermediate information, if any.For distributed algorithms, training data is distributed uniformly. Yourtraining duration is predictable if the input data objects sizes are approximatelythe same. SageMaker does not split the files any further for model training.If the object sizes are skewed, training won't be optimal as the data distributionis also skewed when one host in a training cluster is overloaded, thus becominga bottleneck in training.FastFile modeIf an algorithm supports FastFile mode, SageMaker streams data directly fromS3 to the container with no code changes, and provides file system accessto the data. Users can author their training script to interact with thesefiles as if they were stored on disk.FastFile mode works best when the data is read sequentially. Augmented manifestfiles aren't supported. The startup time is lower when there are fewer filesin the S3 bucket provided.",
								MarkdownDescription: "The training input mode that the algorithm supports. For more informationabout input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).Pipe modeIf an algorithm supports Pipe mode, Amazon SageMaker streams data directlyfrom Amazon S3 to the container.File modeIf an algorithm supports File mode, SageMaker downloads the training datafrom S3 to the provisioned ML storage volume, and mounts the directory tothe Docker volume for the training container.You must provision the ML storage volume with sufficient capacity to accommodatethe data downloaded from S3. In addition to the training data, the ML storagevolume also stores the output model. The algorithm container uses the MLstorage volume to also store intermediate information, if any.For distributed algorithms, training data is distributed uniformly. Yourtraining duration is predictable if the input data objects sizes are approximatelythe same. SageMaker does not split the files any further for model training.If the object sizes are skewed, training won't be optimal as the data distributionis also skewed when one host in a training cluster is overloaded, thus becominga bottleneck in training.FastFile modeIf an algorithm supports FastFile mode, SageMaker streams data directly fromS3 to the container with no code changes, and provides file system accessto the data. Users can author their training script to interact with thesefiles as if they were stored on disk.FastFile mode works best when the data is read sequentially. Augmented manifestfiles aren't supported. The startup time is lower when there are fewer filesin the S3 bucket provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"checkpoint_config": schema.SingleNestedAttribute{
						Description:         "Contains information about the output location for managed spot trainingcheckpoint data.",
						MarkdownDescription: "Contains information about the output location for managed spot trainingcheckpoint data.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"debug_hook_config": schema.SingleNestedAttribute{
						Description:         "Configuration information for the Amazon SageMaker Debugger hook parameters,metric and tensor collections, and storage paths. To learn more about howto configure the DebugHookConfig parameter, see Use the SageMaker and DebuggerConfiguration API Operations to Create, Update, and Debug Your Training Job(https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html).",
						MarkdownDescription: "Configuration information for the Amazon SageMaker Debugger hook parameters,metric and tensor collections, and storage paths. To learn more about howto configure the DebugHookConfig parameter, see Use the SageMaker and DebuggerConfiguration API Operations to Create, Update, and Debug Your Training Job(https://docs.aws.amazon.com/sagemaker/latest/dg/debugger-createtrainingjob-api.html).",
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
											Optional:            true,
											Computed:            false,
										},

										"collection_parameters": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
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

							"hook_parameters": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
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

							"s3_output_path": schema.StringAttribute{
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

					"debug_rule_configurations": schema.ListNestedAttribute{
						Description:         "Configuration information for Amazon SageMaker Debugger rules for debuggingoutput tensors.",
						MarkdownDescription: "Configuration information for Amazon SageMaker Debugger rules for debuggingoutput tensors.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"instance_type": schema.StringAttribute{
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

								"rule_configuration_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rule_evaluator_image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rule_parameters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"s3_output_path": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_inter_container_traffic_encryption": schema.BoolAttribute{
						Description:         "To encrypt all communications between ML compute instances in distributedtraining, choose True. Encryption provides greater security for distributedtraining, but training might take longer. How long it takes depends on theamount of communication between compute instances, especially if you usea deep learning algorithm in distributed training. For more information,see Protect Communications Between ML Compute Instances in a DistributedTraining Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).",
						MarkdownDescription: "To encrypt all communications between ML compute instances in distributedtraining, choose True. Encryption provides greater security for distributedtraining, but training might take longer. How long it takes depends on theamount of communication between compute instances, especially if you usea deep learning algorithm in distributed training. For more information,see Protect Communications Between ML Compute Instances in a DistributedTraining Job (https://docs.aws.amazon.com/sagemaker/latest/dg/train-encrypt.html).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_managed_spot_training": schema.BoolAttribute{
						Description:         "To train models using managed spot training, choose True. Managed spot trainingprovides a fully managed and scalable infrastructure for training machinelearning models. this option is useful when training jobs can be interruptedand when there is flexibility when the training job is run.The complete and intermediate results of jobs are stored in an Amazon S3bucket, and can be used as a starting point to train models incrementally.Amazon SageMaker provides metrics and logs in CloudWatch. They can be usedto see when managed spot training jobs are running, interrupted, resumed,or completed.",
						MarkdownDescription: "To train models using managed spot training, choose True. Managed spot trainingprovides a fully managed and scalable infrastructure for training machinelearning models. this option is useful when training jobs can be interruptedand when there is flexibility when the training job is run.The complete and intermediate results of jobs are stored in an Amazon S3bucket, and can be used as a starting point to train models incrementally.Amazon SageMaker provides metrics and logs in CloudWatch. They can be usedto see when managed spot training jobs are running, interrupted, resumed,or completed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_network_isolation": schema.BoolAttribute{
						Description:         "Isolates the training container. No inbound or outbound network calls canbe made, except for calls between peers within a training cluster for distributedtraining. If you enable network isolation for training jobs that are configuredto use a VPC, SageMaker downloads and uploads customer data and model artifactsthrough the specified VPC, but the training container does not have networkaccess.",
						MarkdownDescription: "Isolates the training container. No inbound or outbound network calls canbe made, except for calls between peers within a training cluster for distributedtraining. If you enable network isolation for training jobs that are configuredto use a VPC, SageMaker downloads and uploads customer data and model artifactsthrough the specified VPC, but the training container does not have networkaccess.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"environment": schema.MapAttribute{
						Description:         "The environment variables to set in the Docker container.",
						MarkdownDescription: "The environment variables to set in the Docker container.",
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

					"hyper_parameters": schema.MapAttribute{
						Description:         "Algorithm-specific parameters that influence the quality of the model. Youset hyperparameters before you start the learning process. For a list ofhyperparameters for each training algorithm provided by SageMaker, see Algorithms(https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).You can specify a maximum of 100 hyperparameters. Each hyperparameter isa key-value pair. Each key and value is limited to 256 characters, as specifiedby the Length Constraint.Do not include any security-sensitive information including account accessIDs, secrets or tokens in any hyperparameter field. If the use of security-sensitivecredentials are detected, SageMaker will reject your training job requestand return an exception error.",
						MarkdownDescription: "Algorithm-specific parameters that influence the quality of the model. Youset hyperparameters before you start the learning process. For a list ofhyperparameters for each training algorithm provided by SageMaker, see Algorithms(https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).You can specify a maximum of 100 hyperparameters. Each hyperparameter isa key-value pair. Each key and value is limited to 256 characters, as specifiedby the Length Constraint.Do not include any security-sensitive information including account accessIDs, secrets or tokens in any hyperparameter field. If the use of security-sensitivecredentials are detected, SageMaker will reject your training job requestand return an exception error.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"infra_check_config": schema.SingleNestedAttribute{
						Description:         "Contains information about the infrastructure health check configurationfor the training job.",
						MarkdownDescription: "Contains information about the infrastructure health check configurationfor the training job.",
						Attributes: map[string]schema.Attribute{
							"enable_infra_check": schema.BoolAttribute{
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

					"input_data_config": schema.ListNestedAttribute{
						Description:         "An array of Channel objects. Each channel is a named input source. InputDataConfigdescribes the input data and its location.Algorithms can accept input data from one or more channels. For example,an algorithm might have two channels of input data, training_data and validation_data.The configuration for each channel provides the S3, EFS, or FSx locationwhere the input data is stored. It also provides information about the storeddata: the MIME type, compression method, and whether the data is wrappedin RecordIO format.Depending on the input mode that the algorithm supports, SageMaker eithercopies input data files from an S3 bucket to a local directory in the Dockercontainer, or makes it available as input streams. For example, if you specifyan EFS location, input data files are available as input streams. They donot need to be downloaded.Your input must be in the same Amazon Web Services region as your trainingjob.",
						MarkdownDescription: "An array of Channel objects. Each channel is a named input source. InputDataConfigdescribes the input data and its location.Algorithms can accept input data from one or more channels. For example,an algorithm might have two channels of input data, training_data and validation_data.The configuration for each channel provides the S3, EFS, or FSx locationwhere the input data is stored. It also provides information about the storeddata: the MIME type, compression method, and whether the data is wrappedin RecordIO format.Depending on the input mode that the algorithm supports, SageMaker eithercopies input data files from an S3 bucket to a local directory in the Dockercontainer, or makes it available as input streams. For example, if you specifyan EFS location, input data files are available as input streams. They donot need to be downloaded.Your input must be in the same Amazon Web Services region as your trainingjob.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"channel_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"compression_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"content_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
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
													Optional:            true,
													Computed:            false,
												},

												"file_system_access_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_system_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_system_type": schema.StringAttribute{
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

										"s3_data_source": schema.SingleNestedAttribute{
											Description:         "Describes the S3 data source.Your input bucket must be in the same Amazon Web Services region as yourtraining job.",
											MarkdownDescription: "Describes the S3 data source.Your input bucket must be in the same Amazon Web Services region as yourtraining job.",
											Attributes: map[string]schema.Attribute{
												"attribute_names": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"instance_group_names": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"input_mode": schema.StringAttribute{
									Description:         "The training input mode that the algorithm supports. For more informationabout input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).Pipe modeIf an algorithm supports Pipe mode, Amazon SageMaker streams data directlyfrom Amazon S3 to the container.File modeIf an algorithm supports File mode, SageMaker downloads the training datafrom S3 to the provisioned ML storage volume, and mounts the directory tothe Docker volume for the training container.You must provision the ML storage volume with sufficient capacity to accommodatethe data downloaded from S3. In addition to the training data, the ML storagevolume also stores the output model. The algorithm container uses the MLstorage volume to also store intermediate information, if any.For distributed algorithms, training data is distributed uniformly. Yourtraining duration is predictable if the input data objects sizes are approximatelythe same. SageMaker does not split the files any further for model training.If the object sizes are skewed, training won't be optimal as the data distributionis also skewed when one host in a training cluster is overloaded, thus becominga bottleneck in training.FastFile modeIf an algorithm supports FastFile mode, SageMaker streams data directly fromS3 to the container with no code changes, and provides file system accessto the data. Users can author their training script to interact with thesefiles as if they were stored on disk.FastFile mode works best when the data is read sequentially. Augmented manifestfiles aren't supported. The startup time is lower when there are fewer filesin the S3 bucket provided.",
									MarkdownDescription: "The training input mode that the algorithm supports. For more informationabout input modes, see Algorithms (https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html).Pipe modeIf an algorithm supports Pipe mode, Amazon SageMaker streams data directlyfrom Amazon S3 to the container.File modeIf an algorithm supports File mode, SageMaker downloads the training datafrom S3 to the provisioned ML storage volume, and mounts the directory tothe Docker volume for the training container.You must provision the ML storage volume with sufficient capacity to accommodatethe data downloaded from S3. In addition to the training data, the ML storagevolume also stores the output model. The algorithm container uses the MLstorage volume to also store intermediate information, if any.For distributed algorithms, training data is distributed uniformly. Yourtraining duration is predictable if the input data objects sizes are approximatelythe same. SageMaker does not split the files any further for model training.If the object sizes are skewed, training won't be optimal as the data distributionis also skewed when one host in a training cluster is overloaded, thus becominga bottleneck in training.FastFile modeIf an algorithm supports FastFile mode, SageMaker streams data directly fromS3 to the container with no code changes, and provides file system accessto the data. Users can author their training script to interact with thesefiles as if they were stored on disk.FastFile mode works best when the data is read sequentially. Augmented manifestfiles aren't supported. The startup time is lower when there are fewer filesin the S3 bucket provided.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"record_wrapper_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"shuffle_config": schema.SingleNestedAttribute{
									Description:         "A configuration for a shuffle option for input data in a channel. If youuse S3Prefix for S3DataType, the results of the S3 key prefix matches areshuffled. If you use ManifestFile, the order of the S3 object referencesin the ManifestFile is shuffled. If you use AugmentedManifestFile, the orderof the JSON lines in the AugmentedManifestFile is shuffled. The shufflingorder is determined using the Seed value.For Pipe input mode, when ShuffleConfig is specified shuffling is done atthe start of every epoch. With large datasets, this ensures that the orderof the training data is different for each epoch, and it helps reduce biasand possible overfitting. In a multi-node training job when ShuffleConfigis combined with S3DataDistributionType of ShardedByS3Key, the data is shuffledacross nodes so that the content sent to a particular node on the first epochmight be sent to a different node on the second epoch.",
									MarkdownDescription: "A configuration for a shuffle option for input data in a channel. If youuse S3Prefix for S3DataType, the results of the S3 key prefix matches areshuffled. If you use ManifestFile, the order of the S3 object referencesin the ManifestFile is shuffled. If you use AugmentedManifestFile, the orderof the JSON lines in the AugmentedManifestFile is shuffled. The shufflingorder is determined using the Seed value.For Pipe input mode, when ShuffleConfig is specified shuffling is done atthe start of every epoch. With large datasets, this ensures that the orderof the training data is different for each epoch, and it helps reduce biasand possible overfitting. In a multi-node training job when ShuffleConfigis combined with S3DataDistributionType of ShardedByS3Key, the data is shuffledacross nodes so that the content sent to a particular node on the first epochmight be sent to a different node on the second epoch.",
									Attributes: map[string]schema.Attribute{
										"seed": schema.Int64Attribute{
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

					"output_data_config": schema.SingleNestedAttribute{
						Description:         "Specifies the path to the S3 location where you want to store model artifacts.SageMaker creates subfolders for the artifacts.",
						MarkdownDescription: "Specifies the path to the S3 location where you want to store model artifacts.SageMaker creates subfolders for the artifacts.",
						Attributes: map[string]schema.Attribute{
							"compression_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
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

							"s3_output_path": schema.StringAttribute{
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

					"profiler_config": schema.SingleNestedAttribute{
						Description:         "Configuration information for Amazon SageMaker Debugger system monitoring,framework profiling, and storage paths.",
						MarkdownDescription: "Configuration information for Amazon SageMaker Debugger system monitoring,framework profiling, and storage paths.",
						Attributes: map[string]schema.Attribute{
							"profiling_interval_in_milliseconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"profiling_parameters": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_output_path": schema.StringAttribute{
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

					"profiler_rule_configurations": schema.ListNestedAttribute{
						Description:         "Configuration information for Amazon SageMaker Debugger rules for profilingsystem and framework metrics.",
						MarkdownDescription: "Configuration information for Amazon SageMaker Debugger rules for profilingsystem and framework metrics.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"instance_type": schema.StringAttribute{
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

								"rule_configuration_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rule_evaluator_image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rule_parameters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"s3_output_path": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_debug_config": schema.SingleNestedAttribute{
						Description:         "Configuration for remote debugging. To learn more about the remote debuggingfunctionality of SageMaker, see Access a training container through AmazonWeb Services Systems Manager (SSM) for remote debugging (https://docs.aws.amazon.com/sagemaker/latest/dg/train-remote-debugging.html).",
						MarkdownDescription: "Configuration for remote debugging. To learn more about the remote debuggingfunctionality of SageMaker, see Access a training container through AmazonWeb Services Systems Manager (SSM) for remote debugging (https://docs.aws.amazon.com/sagemaker/latest/dg/train-remote-debugging.html).",
						Attributes: map[string]schema.Attribute{
							"enable_remote_debug": schema.BoolAttribute{
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

					"resource_config": schema.SingleNestedAttribute{
						Description:         "The resources, including the ML compute instances and ML storage volumes,to use for model training.ML storage volumes store model artifacts and incremental states. Trainingalgorithms might also use ML storage volumes for scratch space. If you wantSageMaker to use the ML storage volume to store the training data, chooseFile as the TrainingInputMode in the algorithm specification. For distributedtraining algorithms, specify an instance count greater than 1.",
						MarkdownDescription: "The resources, including the ML compute instances and ML storage volumes,to use for model training.ML storage volumes store model artifacts and incremental states. Trainingalgorithms might also use ML storage volumes for scratch space. If you wantSageMaker to use the ML storage volume to store the training data, chooseFile as the TrainingInputMode in the algorithm specification. For distributedtraining algorithms, specify an instance count greater than 1.",
						Attributes: map[string]schema.Attribute{
							"instance_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"instance_group_name": schema.StringAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keep_alive_period_in_seconds": schema.Int64Attribute{
								Description:         "Optional. Customer requested period in seconds for which the Training clusteris kept alive after the job is finished.",
								MarkdownDescription: "Optional. Customer requested period in seconds for which the Training clusteris kept alive after the job is finished.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"retry_strategy": schema.SingleNestedAttribute{
						Description:         "The number of times to retry the job when the job fails due to an InternalServerError.",
						MarkdownDescription: "The number of times to retry the job when the job fails due to an InternalServerError.",
						Attributes: map[string]schema.Attribute{
							"maximum_retry_attempts": schema.Int64Attribute{
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

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume toperform tasks on your behalf.During model training, SageMaker needs your permission to read input datafrom an S3 bucket, download a Docker image that contains training code, writemodel artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, andpublish metrics to Amazon CloudWatch. You grant permissions for all of thesetasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).To be able to pass this role to SageMaker, the caller of this API must havethe iam:PassRole permission.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of an IAM role that SageMaker can assume toperform tasks on your behalf.During model training, SageMaker needs your permission to read input datafrom an S3 bucket, download a Docker image that contains training code, writemodel artifacts to an S3 bucket, write logs to Amazon CloudWatch Logs, andpublish metrics to Amazon CloudWatch. You grant permissions for all of thesetasks to an IAM role. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).To be able to pass this role to SageMaker, the caller of this API must havethe iam:PassRole permission.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"stopping_condition": schema.SingleNestedAttribute{
						Description:         "Specifies a limit to how long a model training job can run. It also specifieshow long a managed Spot training job has to complete. When the job reachesthe time limit, SageMaker ends the training job. Use this API to cap modeltraining costs.To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delaysjob termination for 120 seconds. Algorithms can use this 120-second windowto save the model artifacts, so the results of training are not lost.",
						MarkdownDescription: "Specifies a limit to how long a model training job can run. It also specifieshow long a managed Spot training job has to complete. When the job reachesthe time limit, SageMaker ends the training job. Use this API to cap modeltraining costs.To stop a job, SageMaker sends the algorithm the SIGTERM signal, which delaysjob termination for 120 seconds. Algorithms can use this 120-second windowto save the model artifacts, so the results of training are not lost.",
						Attributes: map[string]schema.Attribute{
							"max_pending_time_in_seconds": schema.Int64Attribute{
								Description:         "Maximum job scheduler pending time in seconds.",
								MarkdownDescription: "Maximum job scheduler pending time in seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_runtime_in_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_wait_time_in_seconds": schema.Int64Attribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon WebServices resources in different ways, for example, by purpose, owner, orenvironment. For more information, see Tagging Amazon Web Services Resources(https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon WebServices resources in different ways, for example, by purpose, owner, orenvironment. For more information, see Tagging Amazon Web Services Resources(https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
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

					"tensor_board_output_config": schema.SingleNestedAttribute{
						Description:         "Configuration of storage locations for the Amazon SageMaker Debugger TensorBoardoutput data.",
						MarkdownDescription: "Configuration of storage locations for the Amazon SageMaker Debugger TensorBoardoutput data.",
						Attributes: map[string]schema.Attribute{
							"local_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_output_path": schema.StringAttribute{
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

					"training_job_name": schema.StringAttribute{
						Description:         "The name of the training job. The name must be unique within an Amazon WebServices Region in an Amazon Web Services account.",
						MarkdownDescription: "The name of the training job. The name must be unique within an Amazon WebServices Region in an Amazon Web Services account.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"vpc_config": schema.SingleNestedAttribute{
						Description:         "A VpcConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html)object that specifies the VPC that you want your training job to connectto. Control access to and from your training container by configuring theVPC. For more information, see Protect Training Jobs by Using an Amazon VirtualPrivate Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
						MarkdownDescription: "A VpcConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html)object that specifies the VPC that you want your training job to connectto. Control access to and from your training container by configuring theVPC. For more information, see Protect Training Jobs by Using an Amazon VirtualPrivate Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
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
		},
	}
}

func (r *SagemakerServicesK8SAwsTrainingJobV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_training_job_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsTrainingJobV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("TrainingJob")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
