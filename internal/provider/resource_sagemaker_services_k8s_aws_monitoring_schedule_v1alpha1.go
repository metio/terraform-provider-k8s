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

type SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1GoModel struct {
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
		MonitoringScheduleConfig *struct {
			MonitoringJobDefinition *struct {
				BaselineConfig *struct {
					BaseliningJobName *string `tfsdk:"baselining_job_name" yaml:"baseliningJobName,omitempty"`

					ConstraintsResource *struct {
						S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
					} `tfsdk:"constraints_resource" yaml:"constraintsResource,omitempty"`

					StatisticsResource *struct {
						S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
					} `tfsdk:"statistics_resource" yaml:"statisticsResource,omitempty"`
				} `tfsdk:"baseline_config" yaml:"baselineConfig,omitempty"`

				Environment *map[string]string `tfsdk:"environment" yaml:"environment,omitempty"`

				MonitoringAppSpecification *struct {
					ContainerArguments *[]string `tfsdk:"container_arguments" yaml:"containerArguments,omitempty"`

					ContainerEntrypoint *[]string `tfsdk:"container_entrypoint" yaml:"containerEntrypoint,omitempty"`

					ImageURI *string `tfsdk:"image_uri" yaml:"imageURI,omitempty"`

					PostAnalyticsProcessorSourceURI *string `tfsdk:"post_analytics_processor_source_uri" yaml:"postAnalyticsProcessorSourceURI,omitempty"`

					RecordPreprocessorSourceURI *string `tfsdk:"record_preprocessor_source_uri" yaml:"recordPreprocessorSourceURI,omitempty"`
				} `tfsdk:"monitoring_app_specification" yaml:"monitoringAppSpecification,omitempty"`

				MonitoringInputs *[]struct {
					EndpointInput *struct {
						EndTimeOffset *string `tfsdk:"end_time_offset" yaml:"endTimeOffset,omitempty"`

						EndpointName *string `tfsdk:"endpoint_name" yaml:"endpointName,omitempty"`

						FeaturesAttribute *string `tfsdk:"features_attribute" yaml:"featuresAttribute,omitempty"`

						InferenceAttribute *string `tfsdk:"inference_attribute" yaml:"inferenceAttribute,omitempty"`

						LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

						ProbabilityAttribute *string `tfsdk:"probability_attribute" yaml:"probabilityAttribute,omitempty"`

						ProbabilityThresholdAttribute utilities.DynamicNumber `tfsdk:"probability_threshold_attribute" yaml:"probabilityThresholdAttribute,omitempty"`

						S3DataDistributionType *string `tfsdk:"s3_data_distribution_type" yaml:"s3DataDistributionType,omitempty"`

						S3InputMode *string `tfsdk:"s3_input_mode" yaml:"s3InputMode,omitempty"`

						StartTimeOffset *string `tfsdk:"start_time_offset" yaml:"startTimeOffset,omitempty"`
					} `tfsdk:"endpoint_input" yaml:"endpointInput,omitempty"`
				} `tfsdk:"monitoring_inputs" yaml:"monitoringInputs,omitempty"`

				MonitoringOutputConfig *struct {
					KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

					MonitoringOutputs *[]struct {
						S3Output *struct {
							LocalPath *string `tfsdk:"local_path" yaml:"localPath,omitempty"`

							S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`

							S3UploadMode *string `tfsdk:"s3_upload_mode" yaml:"s3UploadMode,omitempty"`
						} `tfsdk:"s3_output" yaml:"s3Output,omitempty"`
					} `tfsdk:"monitoring_outputs" yaml:"monitoringOutputs,omitempty"`
				} `tfsdk:"monitoring_output_config" yaml:"monitoringOutputConfig,omitempty"`

				MonitoringResources *struct {
					ClusterConfig *struct {
						InstanceCount *int64 `tfsdk:"instance_count" yaml:"instanceCount,omitempty"`

						InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

						VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" yaml:"volumeKMSKeyID,omitempty"`

						VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
					} `tfsdk:"cluster_config" yaml:"clusterConfig,omitempty"`
				} `tfsdk:"monitoring_resources" yaml:"monitoringResources,omitempty"`

				NetworkConfig *struct {
					EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" yaml:"enableInterContainerTrafficEncryption,omitempty"`

					EnableNetworkIsolation *bool `tfsdk:"enable_network_isolation" yaml:"enableNetworkIsolation,omitempty"`

					VpcConfig *struct {
						SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

						Subnets *[]string `tfsdk:"subnets" yaml:"subnets,omitempty"`
					} `tfsdk:"vpc_config" yaml:"vpcConfig,omitempty"`
				} `tfsdk:"network_config" yaml:"networkConfig,omitempty"`

				RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`

				StoppingCondition *struct {
					MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" yaml:"maxRuntimeInSeconds,omitempty"`
				} `tfsdk:"stopping_condition" yaml:"stoppingCondition,omitempty"`
			} `tfsdk:"monitoring_job_definition" yaml:"monitoringJobDefinition,omitempty"`

			MonitoringJobDefinitionName *string `tfsdk:"monitoring_job_definition_name" yaml:"monitoringJobDefinitionName,omitempty"`

			MonitoringType *string `tfsdk:"monitoring_type" yaml:"monitoringType,omitempty"`

			ScheduleConfig *struct {
				ScheduleExpression *string `tfsdk:"schedule_expression" yaml:"scheduleExpression,omitempty"`
			} `tfsdk:"schedule_config" yaml:"scheduleConfig,omitempty"`
		} `tfsdk:"monitoring_schedule_config" yaml:"monitoringScheduleConfig,omitempty"`

		MonitoringScheduleName *string `tfsdk:"monitoring_schedule_name" yaml:"monitoringScheduleName,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1"
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "MonitoringSchedule is the Schema for the MonitoringSchedules API",
		MarkdownDescription: "MonitoringSchedule is the Schema for the MonitoringSchedules API",
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
				Description:         "MonitoringScheduleSpec defines the desired state of MonitoringSchedule.  A schedule for a model monitoring job. For information about model monitor, see Amazon SageMaker Model Monitor (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).",
				MarkdownDescription: "MonitoringScheduleSpec defines the desired state of MonitoringSchedule.  A schedule for a model monitoring job. For information about model monitor, see Amazon SageMaker Model Monitor (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"monitoring_schedule_config": {
						Description:         "The configuration object that specifies the monitoring schedule and defines the monitoring job.",
						MarkdownDescription: "The configuration object that specifies the monitoring schedule and defines the monitoring job.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"monitoring_job_definition": {
								Description:         "Defines the monitoring job.",
								MarkdownDescription: "Defines the monitoring job.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"baseline_config": {
										Description:         "Configuration for monitoring constraints and monitoring statistics. These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.",
										MarkdownDescription: "Configuration for monitoring constraints and monitoring statistics. These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"baselining_job_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"constraints_resource": {
												Description:         "The constraints resource for a monitoring job.",
												MarkdownDescription: "The constraints resource for a monitoring job.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"statistics_resource": {
												Description:         "The statistics resource for a monitoring job.",
												MarkdownDescription: "The statistics resource for a monitoring job.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

									"environment": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"monitoring_app_specification": {
										Description:         "Container image configuration object for the monitoring job.",
										MarkdownDescription: "Container image configuration object for the monitoring job.",

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

											"post_analytics_processor_source_uri": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"record_preprocessor_source_uri": {
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

									"monitoring_inputs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"endpoint_input": {
												Description:         "Input object for the endpoint",
												MarkdownDescription: "Input object for the endpoint",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"end_time_offset": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"endpoint_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"features_attribute": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"inference_attribute": {
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

													"probability_attribute": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"probability_threshold_attribute": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.DynamicNumberType{},

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

													"s3_input_mode": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"start_time_offset": {
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

									"monitoring_output_config": {
										Description:         "The output configuration for monitoring jobs.",
										MarkdownDescription: "The output configuration for monitoring jobs.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"kms_key_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"monitoring_outputs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"s3_output": {
														Description:         "Information about where and how you want to store the results of a monitoring job.",
														MarkdownDescription: "Information about where and how you want to store the results of a monitoring job.",

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

									"monitoring_resources": {
										Description:         "Identifies the resources to deploy for a monitoring job.",
										MarkdownDescription: "Identifies the resources to deploy for a monitoring job.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cluster_config": {
												Description:         "Configuration for the cluster used to run model monitoring jobs.",
												MarkdownDescription: "Configuration for the cluster used to run model monitoring jobs.",

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

										Required: false,
										Optional: true,
										Computed: false,
									},

									"network_config": {
										Description:         "Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.",
										MarkdownDescription: "Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.",

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

									"role_arn": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"stopping_condition": {
										Description:         "A time limit for how long the monitoring job is allowed to run before stopping.",
										MarkdownDescription: "A time limit for how long the monitoring job is allowed to run before stopping.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"monitoring_job_definition_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"monitoring_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"schedule_config": {
								Description:         "Configuration details about the monitoring schedule.",
								MarkdownDescription: "Configuration details about the monitoring schedule.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"schedule_expression": {
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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"monitoring_schedule_name": {
						Description:         "The name of the monitoring schedule. The name must be unique within an Amazon Web Services Region within an Amazon Web Services account.",
						MarkdownDescription: "The name of the monitoring schedule. The name must be unique within an Amazon Web Services Region within an Amazon Web Services account.",

						Type: types.StringType,

						Required: true,
						Optional: false,
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

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1")

	var state SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("MonitoringSchedule")

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

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1")

	var state SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("MonitoringSchedule")

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

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
