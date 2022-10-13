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

type SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1GoModel struct {
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
		HyperParameterTuningJobConfig *struct {
			HyperParameterTuningJobObjective *struct {
				MetricName *string `tfsdk:"metric_name" yaml:"metricName,omitempty"`

				Type_ *string `tfsdk:"type_" yaml:"type_,omitempty"`
			} `tfsdk:"hyper_parameter_tuning_job_objective" yaml:"hyperParameterTuningJobObjective,omitempty"`

			ParameterRanges *struct {
				CategoricalParameterRanges *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"categorical_parameter_ranges" yaml:"categoricalParameterRanges,omitempty"`

				ContinuousParameterRanges *[]struct {
					MaxValue *string `tfsdk:"max_value" yaml:"maxValue,omitempty"`

					MinValue *string `tfsdk:"min_value" yaml:"minValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ScalingType *string `tfsdk:"scaling_type" yaml:"scalingType,omitempty"`
				} `tfsdk:"continuous_parameter_ranges" yaml:"continuousParameterRanges,omitempty"`

				IntegerParameterRanges *[]struct {
					MaxValue *string `tfsdk:"max_value" yaml:"maxValue,omitempty"`

					MinValue *string `tfsdk:"min_value" yaml:"minValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ScalingType *string `tfsdk:"scaling_type" yaml:"scalingType,omitempty"`
				} `tfsdk:"integer_parameter_ranges" yaml:"integerParameterRanges,omitempty"`
			} `tfsdk:"parameter_ranges" yaml:"parameterRanges,omitempty"`

			ResourceLimits *struct {
				MaxNumberOfTrainingJobs *int64 `tfsdk:"max_number_of_training_jobs" yaml:"maxNumberOfTrainingJobs,omitempty"`

				MaxParallelTrainingJobs *int64 `tfsdk:"max_parallel_training_jobs" yaml:"maxParallelTrainingJobs,omitempty"`
			} `tfsdk:"resource_limits" yaml:"resourceLimits,omitempty"`

			Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

			TrainingJobEarlyStoppingType *string `tfsdk:"training_job_early_stopping_type" yaml:"trainingJobEarlyStoppingType,omitempty"`

			TuningJobCompletionCriteria *struct {
				TargetObjectiveMetricValue *float64 `tfsdk:"target_objective_metric_value" yaml:"targetObjectiveMetricValue,omitempty"`
			} `tfsdk:"tuning_job_completion_criteria" yaml:"tuningJobCompletionCriteria,omitempty"`
		} `tfsdk:"hyper_parameter_tuning_job_config" yaml:"hyperParameterTuningJobConfig,omitempty"`

		HyperParameterTuningJobName *string `tfsdk:"hyper_parameter_tuning_job_name" yaml:"hyperParameterTuningJobName,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		TrainingJobDefinition *struct {
			AlgorithmSpecification *struct {
				AlgorithmName *string `tfsdk:"algorithm_name" yaml:"algorithmName,omitempty"`

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

			DefinitionName *string `tfsdk:"definition_name" yaml:"definitionName,omitempty"`

			EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" yaml:"enableInterContainerTrafficEncryption,omitempty"`

			EnableManagedSpotTraining *bool `tfsdk:"enable_managed_spot_training" yaml:"enableManagedSpotTraining,omitempty"`

			EnableNetworkIsolation *bool `tfsdk:"enable_network_isolation" yaml:"enableNetworkIsolation,omitempty"`

			HyperParameterRanges *struct {
				CategoricalParameterRanges *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"categorical_parameter_ranges" yaml:"categoricalParameterRanges,omitempty"`

				ContinuousParameterRanges *[]struct {
					MaxValue *string `tfsdk:"max_value" yaml:"maxValue,omitempty"`

					MinValue *string `tfsdk:"min_value" yaml:"minValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ScalingType *string `tfsdk:"scaling_type" yaml:"scalingType,omitempty"`
				} `tfsdk:"continuous_parameter_ranges" yaml:"continuousParameterRanges,omitempty"`

				IntegerParameterRanges *[]struct {
					MaxValue *string `tfsdk:"max_value" yaml:"maxValue,omitempty"`

					MinValue *string `tfsdk:"min_value" yaml:"minValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ScalingType *string `tfsdk:"scaling_type" yaml:"scalingType,omitempty"`
				} `tfsdk:"integer_parameter_ranges" yaml:"integerParameterRanges,omitempty"`
			} `tfsdk:"hyper_parameter_ranges" yaml:"hyperParameterRanges,omitempty"`

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

			ResourceConfig *struct {
				InstanceCount *int64 `tfsdk:"instance_count" yaml:"instanceCount,omitempty"`

				InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

				VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" yaml:"volumeKMSKeyID,omitempty"`

				VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
			} `tfsdk:"resource_config" yaml:"resourceConfig,omitempty"`

			RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`

			StaticHyperParameters *map[string]string `tfsdk:"static_hyper_parameters" yaml:"staticHyperParameters,omitempty"`

			StoppingCondition *struct {
				MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" yaml:"maxRuntimeInSeconds,omitempty"`

				MaxWaitTimeInSeconds *int64 `tfsdk:"max_wait_time_in_seconds" yaml:"maxWaitTimeInSeconds,omitempty"`
			} `tfsdk:"stopping_condition" yaml:"stoppingCondition,omitempty"`

			TuningObjective *struct {
				MetricName *string `tfsdk:"metric_name" yaml:"metricName,omitempty"`

				Type_ *string `tfsdk:"type_" yaml:"type_,omitempty"`
			} `tfsdk:"tuning_objective" yaml:"tuningObjective,omitempty"`

			VpcConfig *struct {
				SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

				Subnets *[]string `tfsdk:"subnets" yaml:"subnets,omitempty"`
			} `tfsdk:"vpc_config" yaml:"vpcConfig,omitempty"`
		} `tfsdk:"training_job_definition" yaml:"trainingJobDefinition,omitempty"`

		TrainingJobDefinitions *[]struct {
			AlgorithmSpecification *struct {
				AlgorithmName *string `tfsdk:"algorithm_name" yaml:"algorithmName,omitempty"`

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

			DefinitionName *string `tfsdk:"definition_name" yaml:"definitionName,omitempty"`

			EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" yaml:"enableInterContainerTrafficEncryption,omitempty"`

			EnableManagedSpotTraining *bool `tfsdk:"enable_managed_spot_training" yaml:"enableManagedSpotTraining,omitempty"`

			EnableNetworkIsolation *bool `tfsdk:"enable_network_isolation" yaml:"enableNetworkIsolation,omitempty"`

			HyperParameterRanges *struct {
				CategoricalParameterRanges *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"categorical_parameter_ranges" yaml:"categoricalParameterRanges,omitempty"`

				ContinuousParameterRanges *[]struct {
					MaxValue *string `tfsdk:"max_value" yaml:"maxValue,omitempty"`

					MinValue *string `tfsdk:"min_value" yaml:"minValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ScalingType *string `tfsdk:"scaling_type" yaml:"scalingType,omitempty"`
				} `tfsdk:"continuous_parameter_ranges" yaml:"continuousParameterRanges,omitempty"`

				IntegerParameterRanges *[]struct {
					MaxValue *string `tfsdk:"max_value" yaml:"maxValue,omitempty"`

					MinValue *string `tfsdk:"min_value" yaml:"minValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ScalingType *string `tfsdk:"scaling_type" yaml:"scalingType,omitempty"`
				} `tfsdk:"integer_parameter_ranges" yaml:"integerParameterRanges,omitempty"`
			} `tfsdk:"hyper_parameter_ranges" yaml:"hyperParameterRanges,omitempty"`

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

			ResourceConfig *struct {
				InstanceCount *int64 `tfsdk:"instance_count" yaml:"instanceCount,omitempty"`

				InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

				VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" yaml:"volumeKMSKeyID,omitempty"`

				VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
			} `tfsdk:"resource_config" yaml:"resourceConfig,omitempty"`

			RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`

			StaticHyperParameters *map[string]string `tfsdk:"static_hyper_parameters" yaml:"staticHyperParameters,omitempty"`

			StoppingCondition *struct {
				MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" yaml:"maxRuntimeInSeconds,omitempty"`

				MaxWaitTimeInSeconds *int64 `tfsdk:"max_wait_time_in_seconds" yaml:"maxWaitTimeInSeconds,omitempty"`
			} `tfsdk:"stopping_condition" yaml:"stoppingCondition,omitempty"`

			TuningObjective *struct {
				MetricName *string `tfsdk:"metric_name" yaml:"metricName,omitempty"`

				Type_ *string `tfsdk:"type_" yaml:"type_,omitempty"`
			} `tfsdk:"tuning_objective" yaml:"tuningObjective,omitempty"`

			VpcConfig *struct {
				SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

				Subnets *[]string `tfsdk:"subnets" yaml:"subnets,omitempty"`
			} `tfsdk:"vpc_config" yaml:"vpcConfig,omitempty"`
		} `tfsdk:"training_job_definitions" yaml:"trainingJobDefinitions,omitempty"`

		WarmStartConfig *struct {
			ParentHyperParameterTuningJobs *[]struct {
				HyperParameterTuningJobName *string `tfsdk:"hyper_parameter_tuning_job_name" yaml:"hyperParameterTuningJobName,omitempty"`
			} `tfsdk:"parent_hyper_parameter_tuning_jobs" yaml:"parentHyperParameterTuningJobs,omitempty"`

			WarmStartType *string `tfsdk:"warm_start_type" yaml:"warmStartType,omitempty"`
		} `tfsdk:"warm_start_config" yaml:"warmStartConfig,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1"
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "HyperParameterTuningJob is the Schema for the HyperParameterTuningJobs API",
		MarkdownDescription: "HyperParameterTuningJob is the Schema for the HyperParameterTuningJobs API",
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
				Description:         "HyperParameterTuningJobSpec defines the desired state of HyperParameterTuningJob.",
				MarkdownDescription: "HyperParameterTuningJobSpec defines the desired state of HyperParameterTuningJob.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"hyper_parameter_tuning_job_config": {
						Description:         "The HyperParameterTuningJobConfig object that describes the tuning job, including the search strategy, the objective metric used to evaluate training jobs, ranges of parameters to search, and resource limits for the tuning job. For more information, see How Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).",
						MarkdownDescription: "The HyperParameterTuningJobConfig object that describes the tuning job, including the search strategy, the objective metric used to evaluate training jobs, ranges of parameters to search, and resource limits for the tuning job. For more information, see How Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"hyper_parameter_tuning_job_objective": {
								Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
								MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metric_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type_": {
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

							"parameter_ranges": {
								Description:         "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
								MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"categorical_parameter_ranges": {
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

											"values": {
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

									"continuous_parameter_ranges": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"max_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scaling_type": {
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

									"integer_parameter_ranges": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"max_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scaling_type": {
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

							"resource_limits": {
								Description:         "Specifies the maximum number of training jobs and parallel training jobs that a hyperparameter tuning job can launch.",
								MarkdownDescription: "Specifies the maximum number of training jobs and parallel training jobs that a hyperparameter tuning job can launch.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"max_number_of_training_jobs": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_parallel_training_jobs": {
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

							"strategy": {
								Description:         "The strategy hyperparameter tuning uses to find the best combination of hyperparameters for your model. Currently, the only supported value is Bayesian.",
								MarkdownDescription: "The strategy hyperparameter tuning uses to find the best combination of hyperparameters for your model. Currently, the only supported value is Bayesian.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"training_job_early_stopping_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tuning_job_completion_criteria": {
								Description:         "The job completion criteria.",
								MarkdownDescription: "The job completion criteria.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"target_objective_metric_value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.NumberType,

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

					"hyper_parameter_tuning_job_name": {
						Description:         "The name of the tuning job. This name is the prefix for the names of all training jobs that this tuning job launches. The name must be unique within the same Amazon Web Services account and Amazon Web Services Region. The name must have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and : + = @ _ % - (hyphen). The name is not case sensitive.",
						MarkdownDescription: "The name of the tuning job. This name is the prefix for the names of all training jobs that this tuning job launches. The name must be unique within the same Amazon Web Services account and Amazon Web Services Region. The name must have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and : + = @ _ % - (hyphen). The name is not case sensitive.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": {
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).  Tags that you specify for the tuning job are also added to all training jobs that the tuning job launches.",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).  Tags that you specify for the tuning job are also added to all training jobs that the tuning job launches.",

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

					"training_job_definition": {
						Description:         "The HyperParameterTrainingJobDefinition object that describes the training jobs that this tuning job launches, including static hyperparameters, input data configuration, output data configuration, resource configuration, and stopping condition.",
						MarkdownDescription: "The HyperParameterTrainingJobDefinition object that describes the training jobs that this tuning job launches, including static hyperparameters, input data configuration, output data configuration, resource configuration, and stopping condition.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"algorithm_specification": {
								Description:         "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",
								MarkdownDescription: "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"algorithm_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

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

								Required: false,
								Optional: true,
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

							"definition_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_inter_container_traffic_encryption": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_managed_spot_training": {
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

							"hyper_parameter_ranges": {
								Description:         "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
								MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"categorical_parameter_ranges": {
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

											"values": {
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

									"continuous_parameter_ranges": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"max_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scaling_type": {
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

									"integer_parameter_ranges": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"max_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scaling_type": {
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

							"input_data_config": {
								Description:         "",
								MarkdownDescription: "",

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
								Description:         "Provides information about how to store model training results (model artifacts).",
								MarkdownDescription: "Provides information about how to store model training results (model artifacts).",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_config": {
								Description:         "Describes the resources, including ML compute instances and ML storage volumes, to use for model training.",
								MarkdownDescription: "Describes the resources, including ML compute instances and ML storage volumes, to use for model training.",

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

							"role_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"static_hyper_parameters": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"stopping_condition": {
								Description:         "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",
								MarkdownDescription: "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tuning_objective": {
								Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
								MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metric_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type_": {
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

					"training_job_definitions": {
						Description:         "A list of the HyperParameterTrainingJobDefinition objects launched for this tuning job.",
						MarkdownDescription: "A list of the HyperParameterTrainingJobDefinition objects launched for this tuning job.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"algorithm_specification": {
								Description:         "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",
								MarkdownDescription: "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"algorithm_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

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

								Required: false,
								Optional: true,
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

							"definition_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_inter_container_traffic_encryption": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_managed_spot_training": {
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

							"hyper_parameter_ranges": {
								Description:         "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
								MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"categorical_parameter_ranges": {
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

											"values": {
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

									"continuous_parameter_ranges": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"max_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scaling_type": {
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

									"integer_parameter_ranges": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"max_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scaling_type": {
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

							"input_data_config": {
								Description:         "",
								MarkdownDescription: "",

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
								Description:         "Provides information about how to store model training results (model artifacts).",
								MarkdownDescription: "Provides information about how to store model training results (model artifacts).",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_config": {
								Description:         "Describes the resources, including ML compute instances and ML storage volumes, to use for model training.",
								MarkdownDescription: "Describes the resources, including ML compute instances and ML storage volumes, to use for model training.",

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

							"role_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"static_hyper_parameters": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"stopping_condition": {
								Description:         "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",
								MarkdownDescription: "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tuning_objective": {
								Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
								MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metric_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type_": {
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

					"warm_start_config": {
						Description:         "Specifies the configuration for starting the hyperparameter tuning job using one or more previous tuning jobs as a starting point. The results of previous tuning jobs are used to inform which combinations of hyperparameters to search over in the new tuning job.  All training jobs launched by the new hyperparameter tuning job are evaluated by using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHM as the WarmStartType value for the warm start configuration, the training job that performs the best in the new tuning job is compared to the best training jobs from the parent tuning jobs. From these, the training job that performs the best as measured by the objective metric is returned as the overall best training job.  All training jobs launched by parent hyperparameter tuning jobs and the new hyperparameter tuning jobs count against the limit of training jobs for the tuning job.",
						MarkdownDescription: "Specifies the configuration for starting the hyperparameter tuning job using one or more previous tuning jobs as a starting point. The results of previous tuning jobs are used to inform which combinations of hyperparameters to search over in the new tuning job.  All training jobs launched by the new hyperparameter tuning job are evaluated by using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHM as the WarmStartType value for the warm start configuration, the training job that performs the best in the new tuning job is compared to the best training jobs from the parent tuning jobs. From these, the training job that performs the best as measured by the objective metric is returned as the overall best training job.  All training jobs launched by parent hyperparameter tuning jobs and the new hyperparameter tuning jobs count against the limit of training jobs for the tuning job.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"parent_hyper_parameter_tuning_jobs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"hyper_parameter_tuning_job_name": {
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

							"warm_start_type": {
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

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1")

	var state SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("HyperParameterTuningJob")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1")

	var state SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("HyperParameterTuningJob")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
