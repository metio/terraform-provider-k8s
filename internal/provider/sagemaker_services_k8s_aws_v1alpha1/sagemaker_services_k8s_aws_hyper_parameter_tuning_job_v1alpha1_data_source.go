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
	_ datasource.DataSource              = &SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSourceData struct {
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
		HyperParameterTuningJobConfig *struct {
			HyperParameterTuningJobObjective *struct {
				MetricName *string `tfsdk:"metric_name" json:"metricName,omitempty"`
				Type_      *string `tfsdk:"type_" json:"type_,omitempty"`
			} `tfsdk:"hyper_parameter_tuning_job_objective" json:"hyperParameterTuningJobObjective,omitempty"`
			ParameterRanges *struct {
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
				TargetObjectiveMetricValue *big.Float `tfsdk:"target_objective_metric_value" json:"targetObjectiveMetricValue,omitempty"`
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
				KmsKeyID     *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
				S3OutputPath *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
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
				MaxRuntimeInSeconds  *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
				MaxWaitTimeInSeconds *int64 `tfsdk:"max_wait_time_in_seconds" json:"maxWaitTimeInSeconds,omitempty"`
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
				KmsKeyID     *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
				S3OutputPath *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
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
				MaxRuntimeInSeconds  *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
				MaxWaitTimeInSeconds *int64 `tfsdk:"max_wait_time_in_seconds" json:"maxWaitTimeInSeconds,omitempty"`
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

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1"
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HyperParameterTuningJob is the Schema for the HyperParameterTuningJobs API",
		MarkdownDescription: "HyperParameterTuningJob is the Schema for the HyperParameterTuningJobs API",
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
				Description:         "HyperParameterTuningJobSpec defines the desired state of HyperParameterTuningJob.",
				MarkdownDescription: "HyperParameterTuningJobSpec defines the desired state of HyperParameterTuningJob.",
				Attributes: map[string]schema.Attribute{
					"hyper_parameter_tuning_job_config": schema.SingleNestedAttribute{
						Description:         "The HyperParameterTuningJobConfig object that describes the tuning job, including the search strategy, the objective metric used to evaluate training jobs, ranges of parameters to search, and resource limits for the tuning job. For more information, see How Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).",
						MarkdownDescription: "The HyperParameterTuningJobConfig object that describes the tuning job, including the search strategy, the objective metric used to evaluate training jobs, ranges of parameters to search, and resource limits for the tuning job. For more information, see How Hyperparameter Tuning Works (https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-how-it-works.html).",
						Attributes: map[string]schema.Attribute{
							"hyper_parameter_tuning_job_objective": schema.SingleNestedAttribute{
								Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
								MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
								Attributes: map[string]schema.Attribute{
									"metric_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type_": schema.StringAttribute{
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

							"parameter_ranges": schema.SingleNestedAttribute{
								Description:         "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
								MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
								Attributes: map[string]schema.Attribute{
									"categorical_parameter_ranges": schema.ListNestedAttribute{
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

												"values": schema.ListAttribute{
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

									"continuous_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"scaling_type": schema.StringAttribute{
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

									"integer_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"scaling_type": schema.StringAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"resource_limits": schema.SingleNestedAttribute{
								Description:         "Specifies the maximum number of training jobs and parallel training jobs that a hyperparameter tuning job can launch.",
								MarkdownDescription: "Specifies the maximum number of training jobs and parallel training jobs that a hyperparameter tuning job can launch.",
								Attributes: map[string]schema.Attribute{
									"max_number_of_training_jobs": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_parallel_training_jobs": schema.Int64Attribute{
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

							"strategy": schema.StringAttribute{
								Description:         "The strategy hyperparameter tuning uses to find the best combination of hyperparameters for your model.",
								MarkdownDescription: "The strategy hyperparameter tuning uses to find the best combination of hyperparameters for your model.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"training_job_early_stopping_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tuning_job_completion_criteria": schema.SingleNestedAttribute{
								Description:         "The job completion criteria.",
								MarkdownDescription: "The job completion criteria.",
								Attributes: map[string]schema.Attribute{
									"target_objective_metric_value": types.NumberType{
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

					"hyper_parameter_tuning_job_name": schema.StringAttribute{
						Description:         "The name of the tuning job. This name is the prefix for the names of all training jobs that this tuning job launches. The name must be unique within the same Amazon Web Services account and Amazon Web Services Region. The name must have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and : + = @ _ % - (hyphen). The name is not case sensitive.",
						MarkdownDescription: "The name of the tuning job. This name is the prefix for the names of all training jobs that this tuning job launches. The name must be unique within the same Amazon Web Services account and Amazon Web Services Region. The name must have 1 to 32 characters. Valid characters are a-z, A-Z, 0-9, and : + = @ _ % - (hyphen). The name is not case sensitive.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).  Tags that you specify for the tuning job are also added to all training jobs that the tuning job launches.",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).  Tags that you specify for the tuning job are also added to all training jobs that the tuning job launches.",
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

					"training_job_definition": schema.SingleNestedAttribute{
						Description:         "The HyperParameterTrainingJobDefinition object that describes the training jobs that this tuning job launches, including static hyperparameters, input data configuration, output data configuration, resource configuration, and stopping condition.",
						MarkdownDescription: "The HyperParameterTrainingJobDefinition object that describes the training jobs that this tuning job launches, including static hyperparameters, input data configuration, output data configuration, resource configuration, and stopping condition.",
						Attributes: map[string]schema.Attribute{
							"algorithm_specification": schema.SingleNestedAttribute{
								Description:         "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",
								MarkdownDescription: "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",
								Attributes: map[string]schema.Attribute{
									"algorithm_name": schema.StringAttribute{
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

							"definition_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enable_inter_container_traffic_encryption": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enable_managed_spot_training": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enable_network_isolation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"hyper_parameter_ranges": schema.SingleNestedAttribute{
								Description:         "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
								MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
								Attributes: map[string]schema.Attribute{
									"categorical_parameter_ranges": schema.ListNestedAttribute{
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

												"values": schema.ListAttribute{
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

									"continuous_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"scaling_type": schema.StringAttribute{
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

									"integer_parameter_ranges": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"max_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min_value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"scaling_type": schema.StringAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
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
								Description:         "Provides information about how to store model training results (model artifacts).",
								MarkdownDescription: "Provides information about how to store model training results (model artifacts).",
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

							"resource_config": schema.SingleNestedAttribute{
								Description:         "Describes the resources, including machine learning (ML) compute instances and ML storage volumes, to use for model training.",
								MarkdownDescription: "Describes the resources, including machine learning (ML) compute instances and ML storage volumes, to use for model training.",
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
								Description:         "The retry strategy to use when a training job fails due to an InternalServerError. RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJob requests. You can add the StoppingCondition parameter to the request to limit the training time for the complete job.",
								MarkdownDescription: "The retry strategy to use when a training job fails due to an InternalServerError. RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJob requests. You can add the StoppingCondition parameter to the request to limit the training time for the complete job.",
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
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"static_hyper_parameters": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"stopping_condition": schema.SingleNestedAttribute{
								Description:         "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",
								MarkdownDescription: "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",
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

							"tuning_objective": schema.SingleNestedAttribute{
								Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
								MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
								Attributes: map[string]schema.Attribute{
									"metric_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type_": schema.StringAttribute{
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

							"vpc_config": schema.SingleNestedAttribute{
								Description:         "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
								MarkdownDescription: "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
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

					"training_job_definitions": schema.ListNestedAttribute{
						Description:         "A list of the HyperParameterTrainingJobDefinition objects launched for this tuning job.",
						MarkdownDescription: "A list of the HyperParameterTrainingJobDefinition objects launched for this tuning job.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"algorithm_specification": schema.SingleNestedAttribute{
									Description:         "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",
									MarkdownDescription: "Specifies which training algorithm to use for training jobs that a hyperparameter tuning job launches and the metrics to monitor.",
									Attributes: map[string]schema.Attribute{
										"algorithm_name": schema.StringAttribute{
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

								"definition_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"enable_inter_container_traffic_encryption": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"enable_managed_spot_training": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"enable_network_isolation": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"hyper_parameter_ranges": schema.SingleNestedAttribute{
									Description:         "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
									MarkdownDescription: "Specifies ranges of integer, continuous, and categorical hyperparameters that a hyperparameter tuning job searches. The hyperparameter tuning job launches training jobs with hyperparameter values within these ranges to find the combination of values that result in the training job with the best performance as measured by the objective metric of the hyperparameter tuning job.  The maximum number of items specified for Array Members refers to the maximum number of hyperparameters for each range and also the maximum for the hyperparameter tuning job itself. That is, the sum of the number of hyperparameters for all the ranges can't exceed the maximum number specified.",
									Attributes: map[string]schema.Attribute{
										"categorical_parameter_ranges": schema.ListNestedAttribute{
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

													"values": schema.ListAttribute{
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

										"continuous_parameter_ranges": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"max_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"min_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"scaling_type": schema.StringAttribute{
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

										"integer_parameter_ranges": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"max_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"min_value": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"scaling_type": schema.StringAttribute{
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
									},
									Required: false,
									Optional: false,
									Computed: true,
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
									Description:         "Provides information about how to store model training results (model artifacts).",
									MarkdownDescription: "Provides information about how to store model training results (model artifacts).",
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

								"resource_config": schema.SingleNestedAttribute{
									Description:         "Describes the resources, including machine learning (ML) compute instances and ML storage volumes, to use for model training.",
									MarkdownDescription: "Describes the resources, including machine learning (ML) compute instances and ML storage volumes, to use for model training.",
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
									Description:         "The retry strategy to use when a training job fails due to an InternalServerError. RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJob requests. You can add the StoppingCondition parameter to the request to limit the training time for the complete job.",
									MarkdownDescription: "The retry strategy to use when a training job fails due to an InternalServerError. RetryStrategy is specified as part of the CreateTrainingJob and CreateHyperParameterTuningJob requests. You can add the StoppingCondition parameter to the request to limit the training time for the complete job.",
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
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"static_hyper_parameters": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"stopping_condition": schema.SingleNestedAttribute{
									Description:         "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",
									MarkdownDescription: "Specifies a limit to how long a model training job or model compilation job can run. It also specifies how long a managed spot training job has to complete. When the job reaches the time limit, SageMaker ends the training or compilation job. Use this API to cap model training costs.  To stop a training job, SageMaker sends the algorithm the SIGTERM signal, which delays job termination for 120 seconds. Algorithms can use this 120-second window to save the model artifacts, so the results of training are not lost.  The training algorithms provided by SageMaker automatically save the intermediate results of a model training job when possible. This attempt to save artifacts is only a best effort case as model might not be in a state from which it can be saved. For example, if training has just started, the model might not be ready to save. When saved, this intermediate data is a valid model artifact. You can use it to create a model with CreateModel.  The Neural Topic Model (NTM) currently does not support saving intermediate model artifacts. When training NTMs, make sure that the maximum runtime is sufficient for the training job to complete.",
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

								"tuning_objective": schema.SingleNestedAttribute{
									Description:         "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
									MarkdownDescription: "Defines the objective metric for a hyperparameter tuning job. Hyperparameter tuning uses the value of this metric to evaluate the training jobs it launches, and returns the training job that results in either the highest or lowest value for this metric, depending on the value you specify for the Type parameter.",
									Attributes: map[string]schema.Attribute{
										"metric_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type_": schema.StringAttribute{
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

								"vpc_config": schema.SingleNestedAttribute{
									Description:         "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
									MarkdownDescription: "Specifies a VPC that your training jobs and hosted models have access to. Control access to and from your training and model containers by configuring the VPC. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Training Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/train-vpc.html).",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"warm_start_config": schema.SingleNestedAttribute{
						Description:         "Specifies the configuration for starting the hyperparameter tuning job using one or more previous tuning jobs as a starting point. The results of previous tuning jobs are used to inform which combinations of hyperparameters to search over in the new tuning job.  All training jobs launched by the new hyperparameter tuning job are evaluated by using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHM as the WarmStartType value for the warm start configuration, the training job that performs the best in the new tuning job is compared to the best training jobs from the parent tuning jobs. From these, the training job that performs the best as measured by the objective metric is returned as the overall best training job.  All training jobs launched by parent hyperparameter tuning jobs and the new hyperparameter tuning jobs count against the limit of training jobs for the tuning job.",
						MarkdownDescription: "Specifies the configuration for starting the hyperparameter tuning job using one or more previous tuning jobs as a starting point. The results of previous tuning jobs are used to inform which combinations of hyperparameters to search over in the new tuning job.  All training jobs launched by the new hyperparameter tuning job are evaluated by using the objective metric. If you specify IDENTICAL_DATA_AND_ALGORITHM as the WarmStartType value for the warm start configuration, the training job that performs the best in the new tuning job is compared to the best training jobs from the parent tuning jobs. From these, the training job that performs the best as measured by the objective metric is returned as the overall best training job.  All training jobs launched by parent hyperparameter tuning jobs and the new hyperparameter tuning jobs count against the limit of training jobs for the tuning job.",
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
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"warm_start_type": schema.StringAttribute{
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
		},
	}
}

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_hyper_parameter_tuning_job_v1alpha1")

	var data SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "HyperParameterTuningJob"}).
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

	var readResponse SagemakerServicesK8SAwsHyperParameterTuningJobV1Alpha1DataSourceData
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
	data.Kind = pointer.String("HyperParameterTuningJob")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
