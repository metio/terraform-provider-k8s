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
	_ datasource.DataSource = &SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1ManifestData struct {
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
		Autotune *struct {
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"autotune" json:"autotune,omitempty"`
		HyperParameterTuningJobConfig *struct {
			HyperParameterTuningJobObjective *struct {
				MetricName *string `tfsdk:"metric_name" json:"metricName,omitempty"`
				Type_      *string `tfsdk:"type_" json:"type_,omitempty"`
			} `tfsdk:"hyper_parameter_tuning_job_objective" json:"hyperParameterTuningJobObjective,omitempty"`
			ParameterRanges *struct {
				AutoParameters *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					ValueHint *string `tfsdk:"value_hint" json:"valueHint,omitempty"`
				} `tfsdk:"auto_parameters" json:"autoParameters,omitempty"`
				CategoricalParameterRanges *[]struct {
					Name   *string   `tfsdk:"name" json:"name,omitempty"`
					Values *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"categorical_parameter_ranges" json:"categoricalParameterRanges,omitempty"`
				ContinuousParameterRanges *[]struct {
					MaxValue    *string `tfsdk:"max_value" json:"maxValue,omitempty"`
					MinValue    *string `tfsdk:"min_value" json:"minValue,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					ScalingType *string `tfsdk:"scaling_type" json:"scalingType,omitempty"`
				} `tfsdk:"continuous_parameter_ranges" json:"continuousParameterRanges,omitempty"`
				IntegerParameterRanges *[]struct {
					MaxValue    *string `tfsdk:"max_value" json:"maxValue,omitempty"`
					MinValue    *string `tfsdk:"min_value" json:"minValue,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					ScalingType *string `tfsdk:"scaling_type" json:"scalingType,omitempty"`
				} `tfsdk:"integer_parameter_ranges" json:"integerParameterRanges,omitempty"`
			} `tfsdk:"parameter_ranges" json:"parameterRanges,omitempty"`
			ResourceLimits *struct {
				MaxNumberOfTrainingJobs *int64 `tfsdk:"max_number_of_training_jobs" json:"maxNumberOfTrainingJobs,omitempty"`
				MaxParallelTrainingJobs *int64 `tfsdk:"max_parallel_training_jobs" json:"maxParallelTrainingJobs,omitempty"`
			} `tfsdk:"resource_limits" json:"resourceLimits,omitempty"`
			Strategy                     *string `tfsdk:"strategy" json:"strategy,omitempty"`
			TrainingJobEarlyStoppingType *string `tfsdk:"training_job_early_stopping_type" json:"trainingJobEarlyStoppingType,omitempty"`
			TuningJobCompletionCriteria  *struct {
				TargetObjectiveMetricValue *float64 `tfsdk:"target_objective_metric_value" json:"targetObjectiveMetricValue,omitempty"`
			} `tfsdk:"tuning_job_completion_criteria" json:"tuningJobCompletionCriteria,omitempty"`
		} `tfsdk:"hyper_parameter_tuning_job_config" json:"hyperParameterTuningJobConfig,omitempty"`
		HyperParameterTuningJobName *string `tfsdk:"hyper_parameter_tuning_job_name" json:"hyperParameterTuningJobName,omitempty"`
		Tags                        *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TrainingJobDefinition *struct {
			AlgorithmSpecification *struct {
				AlgorithmName     *string `tfsdk:"algorithm_name" json:"algorithmName,omitempty"`
				MetricDefinitions *[]struct {
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
			DefinitionName                        *string `tfsdk:"definition_name" json:"definitionName,omitempty"`
			EnableInterContainerTrafficEncryption *bool   `tfsdk:"enable_inter_container_traffic_encryption" json:"enableInterContainerTrafficEncryption,omitempty"`
			EnableManagedSpotTraining             *bool   `tfsdk:"enable_managed_spot_training" json:"enableManagedSpotTraining,omitempty"`
			EnableNetworkIsolation                *bool   `tfsdk:"enable_network_isolation" json:"enableNetworkIsolation,omitempty"`
			HyperParameterRanges                  *struct {
				AutoParameters *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					ValueHint *string `tfsdk:"value_hint" json:"valueHint,omitempty"`
				} `tfsdk:"auto_parameters" json:"autoParameters,omitempty"`
				CategoricalParameterRanges *[]struct {
					Name   *string   `tfsdk:"name" json:"name,omitempty"`
					Values *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"categorical_parameter_ranges" json:"categoricalParameterRanges,omitempty"`
				ContinuousParameterRanges *[]struct {
					MaxValue    *string `tfsdk:"max_value" json:"maxValue,omitempty"`
					MinValue    *string `tfsdk:"min_value" json:"minValue,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					ScalingType *string `tfsdk:"scaling_type" json:"scalingType,omitempty"`
				} `tfsdk:"continuous_parameter_ranges" json:"continuousParameterRanges,omitempty"`
				IntegerParameterRanges *[]struct {
					MaxValue    *string `tfsdk:"max_value" json:"maxValue,omitempty"`
					MinValue    *string `tfsdk:"min_value" json:"minValue,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					ScalingType *string `tfsdk:"scaling_type" json:"scalingType,omitempty"`
				} `tfsdk:"integer_parameter_ranges" json:"integerParameterRanges,omitempty"`
			} `tfsdk:"hyper_parameter_ranges" json:"hyperParameterRanges,omitempty"`
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
			RoleARN               *string            `tfsdk:"role_arn" json:"roleARN,omitempty"`
			StaticHyperParameters *map[string]string `tfsdk:"static_hyper_parameters" json:"staticHyperParameters,omitempty"`
			StoppingCondition     *struct {
				MaxPendingTimeInSeconds *int64 `tfsdk:"max_pending_time_in_seconds" json:"maxPendingTimeInSeconds,omitempty"`
				MaxRuntimeInSeconds     *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
				MaxWaitTimeInSeconds    *int64 `tfsdk:"max_wait_time_in_seconds" json:"maxWaitTimeInSeconds,omitempty"`
			} `tfsdk:"stopping_condition" json:"stoppingCondition,omitempty"`
			TuningObjective *struct {
				MetricName *string `tfsdk:"metric_name" json:"metricName,omitempty"`
				Type_      *string `tfsdk:"type_" json:"type_,omitempty"`
			} `tfsdk:"tuning_objective" json:"tuningObjective,omitempty"`
			VpcConfig *struct {
				SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
				Subnets          *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
			} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
		} `tfsdk:"training_job_definition" json:"trainingJobDefinition,omitempty"`
		TrainingJobDefinitions *[]struct {
			AlgorithmSpecification *struct {
				AlgorithmName     *string `tfsdk:"algorithm_name" json:"algorithmName,omitempty"`
				MetricDefinitions *[]struct {
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
			DefinitionName                        *string `tfsdk:"definition_name" json:"definitionName,omitempty"`
			EnableInterContainerTrafficEncryption *bool   `tfsdk:"enable_inter_container_traffic_encryption" json:"enableInterContainerTrafficEncryption,omitempty"`
			EnableManagedSpotTraining             *bool   `tfsdk:"enable_managed_spot_training" json:"enableManagedSpotTraining,omitempty"`
			EnableNetworkIsolation                *bool   `tfsdk:"enable_network_isolation" json:"enableNetworkIsolation,omitempty"`
			HyperParameterRanges                  *struct {
				AutoParameters *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					ValueHint *string `tfsdk:"value_hint" json:"valueHint,omitempty"`
				} `tfsdk:"auto_parameters" json:"autoParameters,omitempty"`
				CategoricalParameterRanges *[]struct {
					Name   *string   `tfsdk:"name" json:"name,omitempty"`
					Values *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"categorical_parameter_ranges" json:"categoricalParameterRanges,omitempty"`
				ContinuousParameterRanges *[]struct {
					MaxValue    *string `tfsdk:"max_value" json:"maxValue,omitempty"`
					MinValue    *string `tfsdk:"min_value" json:"minValue,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					ScalingType *string `tfsdk:"scaling_type" json:"scalingType,omitempty"`
				} `tfsdk:"continuous_parameter_ranges" json:"continuousParameterRanges,omitempty"`
				IntegerParameterRanges *[]struct {
					MaxValue    *string `tfsdk:"max_value" json:"maxValue,omitempty"`
					MinValue    *string `tfsdk:"min_value" json:"minValue,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					ScalingType *string `tfsdk:"scaling_type" json:"scalingType,omitempty"`
				} `tfsdk:"integer_parameter_ranges" json:"integerParameterRanges,omitempty"`
			} `tfsdk:"hyper_parameter_ranges" json:"hyperParameterRanges,omitempty"`
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
			RoleARN               *string            `tfsdk:"role_arn" json:"roleARN,omitempty"`
			StaticHyperParameters *map[string]string `tfsdk:"static_hyper_parameters" json:"staticHyperParameters,omitempty"`
			StoppingCondition     *struct {
				MaxPendingTimeInSeconds *int64 `tfsdk:"max_pending_time_in_seconds" json:"maxPendingTimeInSeconds,omitempty"`
				MaxRuntimeInSeconds     *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
				MaxWaitTimeInSeconds    *int64 `tfsdk:"max_wait_time_in_seconds" json:"maxWaitTimeInSeconds,omitempty"`
			} `tfsdk:"stopping_condition" json:"stoppingCondition,omitempty"`
			TuningObjective *struct {
				MetricName *string `tfsdk:"metric_name" json:"metricName,omitempty"`
				Type_      *string `tfsdk:"type_" json:"type_,omitempty"`
			} `tfsdk:"tuning_objective" json:"tuningObjective,omitempty"`
			VpcConfig *struct {
				SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
				Subnets          *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
			} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
		} `tfsdk:"training_job_definitions" json:"trainingJobDefinitions,omitempty"`
		WarmStartConfig *struct {
			ParentHyperParameterTuningJobs *[]struct {
				HyperParameterTuningJobName *string `tfsdk:"hyper_parameter_tuning_job_name" json:"hyperParameterTuningJobName,omitempty"`
			} `tfsdk:"parent_hyper_parameter_tuning_jobs" json:"parentHyperParameterTuningJobs,omitempty"`
			WarmStartType *string `tfsdk:"warm_start_type" json:"warmStartType,omitempty"`
		} `tfsdk:"warm_start_config" json:"warmStartConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HyperParameterTuningJob is the Schema for the HyperParameterTuningJobs API",
		MarkdownDescription: "HyperParameterTuningJob is the Schema for the HyperParameterTuningJobs API",
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
				Description:         "HyperParameterTuningJobSpec defines the desired state of HyperParameterTuningJob.",
				MarkdownDescription: "HyperParameterTuningJobSpec defines the desired state of HyperParameterTuningJob.",
				Attributes: map[string]schema.Attribute{
					"autotune": schema.SingleNestedAttribute{
						Description:         "Configures SageMaker Automatic model tuning (AMT) to automatically find optimalparameters for the following fields:   * ParameterRanges (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-ParameterRanges):   The names and ranges of parameters that a hyperparameter tuning job can   optimize.   * ResourceLimits (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ResourceLimits.html):   The maximum resources that can be used for a training job. These resources   include the maximum number of training jobs, the maximum runtime of a   tuning job, and the maximum number of training jobs to run at the same   time.   * TrainingJobEarlyStoppingType (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-TrainingJobEarlyStoppingType):   A flag that specifies whether or not to use early stopping for training   jobs launched by a hyperparameter tuning job.   * RetryStrategy (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html#sagemaker-Type-HyperParameterTrainingJobDefinition-RetryStrategy):   The number of times to retry a training job.   * Strategy (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html):   Specifies how hyperparameter tuning chooses the combinations of hyperparameter   values to use for the training jobs that it launches.   * ConvergenceDetected (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ConvergenceDetected.html):   A flag to indicate that Automatic model tuning (AMT) has detected model   convergence.",
						MarkdownDescription: "Configures SageMaker Automatic model tuning (AMT) to automatically find optimalparameters for the following fields:   * ParameterRanges (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-ParameterRanges):   The names and ranges of parameters that a hyperparameter tuning job can   optimize.   * ResourceLimits (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ResourceLimits.html):   The maximum resources that can be used for a training job. These resources   include the maximum number of training jobs, the maximum runtime of a   tuning job, and the maximum number of training jobs to run at the same   time.   * TrainingJobEarlyStoppingType (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html#sagemaker-Type-HyperParameterTuningJobConfig-TrainingJobEarlyStoppingType):   A flag that specifies whether or not to use early stopping for training   jobs launched by a hyperparameter tuning job.   * RetryStrategy (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html#sagemaker-Type-HyperParameterTrainingJobDefinition-RetryStrategy):   The number of times to retry a training job.   * Strategy (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html):   Specifies how hyperparameter tuning chooses the combinations of hyperparameter   values to use for the training jobs that it launches.   * ConvergenceDetected (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ConvergenceDetected.html):   A flag to indicate that Automatic model tuning (AMT) has detected model   convergence.",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
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

					"hyper_parameter_tuning_job_config": schema.SingleNestedAttribute{
						Description:         "The HyperParameterTuningJobConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html)object that describes the tuning job, including the search strategy, theobjective metric used to evaluate training jobs, ranges of parameters tosearch, and resource limits for the tuning job. For more information, seeHow Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).",
						MarkdownDescription: "The HyperParameterTuningJobConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTuningJobConfig.html)object that describes the tuning job, including the search strategy, theobjective metric used to evaluate training jobs, ranges of parameters tosearch, and resource limits for the tuning job. For more information, seeHow Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).",
						Attributes: map[string]schema.Attribute{
							"hyper_parameter_tuning_job_objective": schema.SingleNestedAttribute{
								Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparametertuning uses the value of this metric to evaluate the training jobs it launches,and returns the training job that results in either the highest or lowestvalue for this metric, depending on the value you specify for the Type parameter.If you want to define a custom objective metric, see Define metrics and environmentvariables (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics-variables.html).",
								MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparametertuning uses the value of this metric to evaluate the training jobs it launches,and returns the training job that results in either the highest or lowestvalue for this metric, depending on the value you specify for the Type parameter.If you want to define a custom objective metric, see Define metrics and environmentvariables (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics-variables.html).",
								Attributes: map[string]schema.Attribute{
									"metric_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type_": schema.StringAttribute{
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

							"parameter_ranges": schema.SingleNestedAttribute{
								Description:         "Specifies ranges of integer, continuous, and categorical hyperparametersthat a hyperparameter tuning job searches. The hyperparameter tuning joblaunches training jobs with hyperparameter values within these ranges tofind the combination of values that result in the training job with the bestperformance as measured by the objective metric of the hyperparameter tuningjob.The maximum number of items specified for Array Members refers to the maximumnumber of hyperparameters for each range and also the maximum for the hyperparametertuning job itself. That is, the sum of the number of hyperparameters forall the ranges can't exceed the maximum number specified.",
								MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparametersthat a hyperparameter tuning job searches. The hyperparameter tuning joblaunches training jobs with hyperparameter values within these ranges tofind the combination of values that result in the training job with the bestperformance as measured by the objective metric of the hyperparameter tuningjob.The maximum number of items specified for Array Members refers to the maximumnumber of hyperparameters for each range and also the maximum for the hyperparametertuning job itself. That is, the sum of the number of hyperparameters forall the ranges can't exceed the maximum number specified.",
								Attributes: map[string]schema.Attribute{
									"auto_parameters": schema.ListNestedAttribute{
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

												"value_hint": schema.StringAttribute{
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

									"categorical_parameter_ranges": schema.ListNestedAttribute{
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

												"values": schema.ListAttribute{
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

									"continuous_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scaling_type": schema.StringAttribute{
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

									"integer_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scaling_type": schema.StringAttribute{
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

							"resource_limits": schema.SingleNestedAttribute{
								Description:         "Specifies the maximum number of training jobs and parallel training jobsthat a hyperparameter tuning job can launch.",
								MarkdownDescription: "Specifies the maximum number of training jobs and parallel training jobsthat a hyperparameter tuning job can launch.",
								Attributes: map[string]schema.Attribute{
									"max_number_of_training_jobs": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_parallel_training_jobs": schema.Int64Attribute{
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

							"strategy": schema.StringAttribute{
								Description:         "The strategy hyperparameter tuning uses to find the best combination of hyperparametersfor your model.",
								MarkdownDescription: "The strategy hyperparameter tuning uses to find the best combination of hyperparametersfor your model.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"training_job_early_stopping_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tuning_job_completion_criteria": schema.SingleNestedAttribute{
								Description:         "The job completion criteria.",
								MarkdownDescription: "The job completion criteria.",
								Attributes: map[string]schema.Attribute{
									"target_objective_metric_value": schema.Float64Attribute{
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

					"hyper_parameter_tuning_job_name": schema.StringAttribute{
						Description:         "The name of the tuning job. This name is the prefix for the names of alltraining jobs that this tuning job launches. The name must be unique withinthe same Amazon Web Services account and Amazon Web Services Region. Thename must have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and: + = @ _ % - (hyphen). The name is not case sensitive.",
						MarkdownDescription: "The name of the tuning job. This name is the prefix for the names of alltraining jobs that this tuning job launches. The name must be unique withinthe same Amazon Web Services account and Amazon Web Services Region. Thename must have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and: + = @ _ % - (hyphen). The name is not case sensitive.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon WebServices resources in different ways, for example, by purpose, owner, orenvironment. For more information, see Tagging Amazon Web Services Resources(https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).Tags that you specify for the tuning job are also added to all training jobsthat the tuning job launches.",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon WebServices resources in different ways, for example, by purpose, owner, orenvironment. For more information, see Tagging Amazon Web Services Resources(https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).Tags that you specify for the tuning job are also added to all training jobsthat the tuning job launches.",
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

					"training_job_definition": schema.SingleNestedAttribute{
						Description:         "The HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)object that describes the training jobs that this tuning job launches, includingstatic hyperparameters, input data configuration, output data configuration,resource configuration, and stopping condition.",
						MarkdownDescription: "The HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)object that describes the training jobs that this tuning job launches, includingstatic hyperparameters, input data configuration, output data configuration,resource configuration, and stopping condition.",
						Attributes: map[string]schema.Attribute{
							"algorithm_specification": schema.SingleNestedAttribute{
								Description:         "Specifies which training algorithm to use for training jobs that a hyperparametertuning job launches and the metrics to monitor.",
								MarkdownDescription: "Specifies which training algorithm to use for training jobs that a hyperparametertuning job launches and the metrics to monitor.",
								Attributes: map[string]schema.Attribute{
									"algorithm_name": schema.StringAttribute{
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
								Required: false,
								Optional: true,
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

							"definition_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_inter_container_traffic_encryption": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_managed_spot_training": schema.BoolAttribute{
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

							"hyper_parameter_ranges": schema.SingleNestedAttribute{
								Description:         "Specifies ranges of integer, continuous, and categorical hyperparametersthat a hyperparameter tuning job searches. The hyperparameter tuning joblaunches training jobs with hyperparameter values within these ranges tofind the combination of values that result in the training job with the bestperformance as measured by the objective metric of the hyperparameter tuningjob.The maximum number of items specified for Array Members refers to the maximumnumber of hyperparameters for each range and also the maximum for the hyperparametertuning job itself. That is, the sum of the number of hyperparameters forall the ranges can't exceed the maximum number specified.",
								MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparametersthat a hyperparameter tuning job searches. The hyperparameter tuning joblaunches training jobs with hyperparameter values within these ranges tofind the combination of values that result in the training job with the bestperformance as measured by the objective metric of the hyperparameter tuningjob.The maximum number of items specified for Array Members refers to the maximumnumber of hyperparameters for each range and also the maximum for the hyperparametertuning job itself. That is, the sum of the number of hyperparameters forall the ranges can't exceed the maximum number specified.",
								Attributes: map[string]schema.Attribute{
									"auto_parameters": schema.ListNestedAttribute{
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

												"value_hint": schema.StringAttribute{
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

									"categorical_parameter_ranges": schema.ListNestedAttribute{
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

												"values": schema.ListAttribute{
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

									"continuous_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scaling_type": schema.StringAttribute{
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

									"integer_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scaling_type": schema.StringAttribute{
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

							"input_data_config": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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
								Description:         "Provides information about how to store model training results (model artifacts).",
								MarkdownDescription: "Provides information about how to store model training results (model artifacts).",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_config": schema.SingleNestedAttribute{
								Description:         "Describes the resources, including machine learning (ML) compute instancesand ML storage volumes, to use for model training.",
								MarkdownDescription: "Describes the resources, including machine learning (ML) compute instancesand ML storage volumes, to use for model training.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"retry_strategy": schema.SingleNestedAttribute{
								Description:         "The retry strategy to use when a training job fails due to an InternalServerError.RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJobrequests. You can add the StoppingCondition parameter to the request to limitthe training time for the complete job.",
								MarkdownDescription: "The retry strategy to use when a training job fails due to an InternalServerError.RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJobrequests. You can add the StoppingCondition parameter to the request to limitthe training time for the complete job.",
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
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"static_hyper_parameters": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stopping_condition": schema.SingleNestedAttribute{
								Description:         "Specifies a limit to how long a model training job or model compilation jobcan run. It also specifies how long a managed spot training job has to complete.When the job reaches the time limit, SageMaker ends the training or compilationjob. Use this API to cap model training costs.To stop a training job, SageMaker sends the algorithm the SIGTERM signal,which delays job termination for 120 seconds. Algorithms can use this 120-secondwindow to save the model artifacts, so the results of training are not lost.The training algorithms provided by SageMaker automatically save the intermediateresults of a model training job when possible. This attempt to save artifactsis only a best effort case as model might not be in a state from which itcan be saved. For example, if training has just started, the model mightnot be ready to save. When saved, this intermediate data is a valid modelartifact. You can use it to create a model with CreateModel.The Neural Topic Model (NTM) currently does not support saving intermediatemodel artifacts. When training NTMs, make sure that the maximum runtime issufficient for the training job to complete.",
								MarkdownDescription: "Specifies a limit to how long a model training job or model compilation jobcan run. It also specifies how long a managed spot training job has to complete.When the job reaches the time limit, SageMaker ends the training or compilationjob. Use this API to cap model training costs.To stop a training job, SageMaker sends the algorithm the SIGTERM signal,which delays job termination for 120 seconds. Algorithms can use this 120-secondwindow to save the model artifacts, so the results of training are not lost.The training algorithms provided by SageMaker automatically save the intermediateresults of a model training job when possible. This attempt to save artifactsis only a best effort case as model might not be in a state from which itcan be saved. For example, if training has just started, the model mightnot be ready to save. When saved, this intermediate data is a valid modelartifact. You can use it to create a model with CreateModel.The Neural Topic Model (NTM) currently does not support saving intermediatemodel artifacts. When training NTMs, make sure that the maximum runtime issufficient for the training job to complete.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tuning_objective": schema.SingleNestedAttribute{
								Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparametertuning uses the value of this metric to evaluate the training jobs it launches,and returns the training job that results in either the highest or lowestvalue for this metric, depending on the value you specify for the Type parameter.If you want to define a custom objective metric, see Define metrics and environmentvariables (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics-variables.html).",
								MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparametertuning uses the value of this metric to evaluate the training jobs it launches,and returns the training job that results in either the highest or lowestvalue for this metric, depending on the value you specify for the Type parameter.If you want to define a custom objective metric, see Define metrics and environmentvariables (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics-variables.html).",
								Attributes: map[string]schema.Attribute{
									"metric_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type_": schema.StringAttribute{
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

					"training_job_definitions": schema.ListNestedAttribute{
						Description:         "A list of the HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)objects launched for this tuning job.",
						MarkdownDescription: "A list of the HyperParameterTrainingJobDefinition (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_HyperParameterTrainingJobDefinition.html)objects launched for this tuning job.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"algorithm_specification": schema.SingleNestedAttribute{
									Description:         "Specifies which training algorithm to use for training jobs that a hyperparametertuning job launches and the metrics to monitor.",
									MarkdownDescription: "Specifies which training algorithm to use for training jobs that a hyperparametertuning job launches and the metrics to monitor.",
									Attributes: map[string]schema.Attribute{
										"algorithm_name": schema.StringAttribute{
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
									Required: false,
									Optional: true,
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

								"definition_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enable_inter_container_traffic_encryption": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enable_managed_spot_training": schema.BoolAttribute{
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

								"hyper_parameter_ranges": schema.SingleNestedAttribute{
									Description:         "Specifies ranges of integer, continuous, and categorical hyperparametersthat a hyperparameter tuning job searches. The hyperparameter tuning joblaunches training jobs with hyperparameter values within these ranges tofind the combination of values that result in the training job with the bestperformance as measured by the objective metric of the hyperparameter tuningjob.The maximum number of items specified for Array Members refers to the maximumnumber of hyperparameters for each range and also the maximum for the hyperparametertuning job itself. That is, the sum of the number of hyperparameters forall the ranges can't exceed the maximum number specified.",
									MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparametersthat a hyperparameter tuning job searches. The hyperparameter tuning joblaunches training jobs with hyperparameter values within these ranges tofind the combination of values that result in the training job with the bestperformance as measured by the objective metric of the hyperparameter tuningjob.The maximum number of items specified for Array Members refers to the maximumnumber of hyperparameters for each range and also the maximum for the hyperparametertuning job itself. That is, the sum of the number of hyperparameters forall the ranges can't exceed the maximum number specified.",
									Attributes: map[string]schema.Attribute{
										"auto_parameters": schema.ListNestedAttribute{
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

													"value_hint": schema.StringAttribute{
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

										"categorical_parameter_ranges": schema.ListNestedAttribute{
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

													"values": schema.ListAttribute{
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

										"continuous_parameter_ranges": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"max_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scaling_type": schema.StringAttribute{
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

										"integer_parameter_ranges": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"max_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scaling_type": schema.StringAttribute{
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

								"input_data_config": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
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
									Description:         "Provides information about how to store model training results (model artifacts).",
									MarkdownDescription: "Provides information about how to store model training results (model artifacts).",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"resource_config": schema.SingleNestedAttribute{
									Description:         "Describes the resources, including machine learning (ML) compute instancesand ML storage volumes, to use for model training.",
									MarkdownDescription: "Describes the resources, including machine learning (ML) compute instancesand ML storage volumes, to use for model training.",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"retry_strategy": schema.SingleNestedAttribute{
									Description:         "The retry strategy to use when a training job fails due to an InternalServerError.RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJobrequests. You can add the StoppingCondition parameter to the request to limitthe training time for the complete job.",
									MarkdownDescription: "The retry strategy to use when a training job fails due to an InternalServerError.RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJobrequests. You can add the StoppingCondition parameter to the request to limitthe training time for the complete job.",
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
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"static_hyper_parameters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"stopping_condition": schema.SingleNestedAttribute{
									Description:         "Specifies a limit to how long a model training job or model compilation jobcan run. It also specifies how long a managed spot training job has to complete.When the job reaches the time limit, SageMaker ends the training or compilationjob. Use this API to cap model training costs.To stop a training job, SageMaker sends the algorithm the SIGTERM signal,which delays job termination for 120 seconds. Algorithms can use this 120-secondwindow to save the model artifacts, so the results of training are not lost.The training algorithms provided by SageMaker automatically save the intermediateresults of a model training job when possible. This attempt to save artifactsis only a best effort case as model might not be in a state from which itcan be saved. For example, if training has just started, the model mightnot be ready to save. When saved, this intermediate data is a valid modelartifact. You can use it to create a model with CreateModel.The Neural Topic Model (NTM) currently does not support saving intermediatemodel artifacts. When training NTMs, make sure that the maximum runtime issufficient for the training job to complete.",
									MarkdownDescription: "Specifies a limit to how long a model training job or model compilation jobcan run. It also specifies how long a managed spot training job has to complete.When the job reaches the time limit, SageMaker ends the training or compilationjob. Use this API to cap model training costs.To stop a training job, SageMaker sends the algorithm the SIGTERM signal,which delays job termination for 120 seconds. Algorithms can use this 120-secondwindow to save the model artifacts, so the results of training are not lost.The training algorithms provided by SageMaker automatically save the intermediateresults of a model training job when possible. This attempt to save artifactsis only a best effort case as model might not be in a state from which itcan be saved. For example, if training has just started, the model mightnot be ready to save. When saved, this intermediate data is a valid modelartifact. You can use it to create a model with CreateModel.The Neural Topic Model (NTM) currently does not support saving intermediatemodel artifacts. When training NTMs, make sure that the maximum runtime issufficient for the training job to complete.",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tuning_objective": schema.SingleNestedAttribute{
									Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparametertuning uses the value of this metric to evaluate the training jobs it launches,and returns the training job that results in either the highest or lowestvalue for this metric, depending on the value you specify for the Type parameter.If you want to define a custom objective metric, see Define metrics and environmentvariables (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics-variables.html).",
									MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparametertuning uses the value of this metric to evaluate the training jobs it launches,and returns the training job that results in either the highest or lowestvalue for this metric, depending on the value you specify for the Type parameter.If you want to define a custom objective metric, see Define metrics and environmentvariables (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics-variables.html).",
									Attributes: map[string]schema.Attribute{
										"metric_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type_": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"warm_start_config": schema.SingleNestedAttribute{
						Description:         "Specifies the configuration for starting the hyperparameter tuning job usingone or more previous tuning jobs as a starting point. The results of previoustuning jobs are used to inform which combinations of hyperparameters to searchover in the new tuning job.All training jobs launched by the new hyperparameter tuning job are evaluatedby using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHMas the WarmStartType value for the warm start configuration, the trainingjob that performs the best in the new tuning job is compared to the besttraining jobs from the parent tuning jobs. From these, the training job thatperforms the best as measured by the objective metric is returned as theoverall best training job.All training jobs launched by parent hyperparameter tuning jobs and the newhyperparameter tuning jobs count against the limit of training jobs for thetuning job.",
						MarkdownDescription: "Specifies the configuration for starting the hyperparameter tuning job usingone or more previous tuning jobs as a starting point. The results of previoustuning jobs are used to inform which combinations of hyperparameters to searchover in the new tuning job.All training jobs launched by the new hyperparameter tuning job are evaluatedby using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHMas the WarmStartType value for the warm start configuration, the trainingjob that performs the best in the new tuning job is compared to the besttraining jobs from the parent tuning jobs. From these, the training job thatperforms the best as measured by the objective metric is returned as theoverall best training job.All training jobs launched by parent hyperparameter tuning jobs and the newhyperparameter tuning jobs count against the limit of training jobs for thetuning job.",
						Attributes: map[string]schema.Attribute{
							"parent_hyper_parameter_tuning_jobs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hyper_parameter_tuning_job_name": schema.StringAttribute{
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

							"warm_start_type": schema.StringAttribute{
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
		},
	}
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("HyperParameterTuningJob")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
