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
	_ datasource.DataSource = &SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1ManifestData struct {
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
		MonitoringScheduleConfig *struct {
			MonitoringJobDefinition *struct {
				BaselineConfig *struct {
					BaseliningJobName   *string `tfsdk:"baselining_job_name" json:"baseliningJobName,omitempty"`
					ConstraintsResource *struct {
						S3URI *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
					} `tfsdk:"constraints_resource" json:"constraintsResource,omitempty"`
					StatisticsResource *struct {
						S3URI *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
					} `tfsdk:"statistics_resource" json:"statisticsResource,omitempty"`
				} `tfsdk:"baseline_config" json:"baselineConfig,omitempty"`
				Environment                *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
				MonitoringAppSpecification *struct {
					ContainerArguments              *[]string `tfsdk:"container_arguments" json:"containerArguments,omitempty"`
					ContainerEntrypoint             *[]string `tfsdk:"container_entrypoint" json:"containerEntrypoint,omitempty"`
					ImageURI                        *string   `tfsdk:"image_uri" json:"imageURI,omitempty"`
					PostAnalyticsProcessorSourceURI *string   `tfsdk:"post_analytics_processor_source_uri" json:"postAnalyticsProcessorSourceURI,omitempty"`
					RecordPreprocessorSourceURI     *string   `tfsdk:"record_preprocessor_source_uri" json:"recordPreprocessorSourceURI,omitempty"`
				} `tfsdk:"monitoring_app_specification" json:"monitoringAppSpecification,omitempty"`
				MonitoringInputs *[]struct {
					EndpointInput *struct {
						EndTimeOffset                 *string  `tfsdk:"end_time_offset" json:"endTimeOffset,omitempty"`
						EndpointName                  *string  `tfsdk:"endpoint_name" json:"endpointName,omitempty"`
						ExcludeFeaturesAttribute      *string  `tfsdk:"exclude_features_attribute" json:"excludeFeaturesAttribute,omitempty"`
						FeaturesAttribute             *string  `tfsdk:"features_attribute" json:"featuresAttribute,omitempty"`
						InferenceAttribute            *string  `tfsdk:"inference_attribute" json:"inferenceAttribute,omitempty"`
						LocalPath                     *string  `tfsdk:"local_path" json:"localPath,omitempty"`
						ProbabilityAttribute          *string  `tfsdk:"probability_attribute" json:"probabilityAttribute,omitempty"`
						ProbabilityThresholdAttribute *float64 `tfsdk:"probability_threshold_attribute" json:"probabilityThresholdAttribute,omitempty"`
						S3DataDistributionType        *string  `tfsdk:"s3_data_distribution_type" json:"s3DataDistributionType,omitempty"`
						S3InputMode                   *string  `tfsdk:"s3_input_mode" json:"s3InputMode,omitempty"`
						StartTimeOffset               *string  `tfsdk:"start_time_offset" json:"startTimeOffset,omitempty"`
					} `tfsdk:"endpoint_input" json:"endpointInput,omitempty"`
				} `tfsdk:"monitoring_inputs" json:"monitoringInputs,omitempty"`
				MonitoringOutputConfig *struct {
					KmsKeyID          *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
					MonitoringOutputs *[]struct {
						S3Output *struct {
							LocalPath    *string `tfsdk:"local_path" json:"localPath,omitempty"`
							S3URI        *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
							S3UploadMode *string `tfsdk:"s3_upload_mode" json:"s3UploadMode,omitempty"`
						} `tfsdk:"s3_output" json:"s3Output,omitempty"`
					} `tfsdk:"monitoring_outputs" json:"monitoringOutputs,omitempty"`
				} `tfsdk:"monitoring_output_config" json:"monitoringOutputConfig,omitempty"`
				MonitoringResources *struct {
					ClusterConfig *struct {
						InstanceCount  *int64  `tfsdk:"instance_count" json:"instanceCount,omitempty"`
						InstanceType   *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
						VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" json:"volumeKMSKeyID,omitempty"`
						VolumeSizeInGB *int64  `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
					} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
				} `tfsdk:"monitoring_resources" json:"monitoringResources,omitempty"`
				NetworkConfig *struct {
					EnableInterContainerTrafficEncryption *bool `tfsdk:"enable_inter_container_traffic_encryption" json:"enableInterContainerTrafficEncryption,omitempty"`
					EnableNetworkIsolation                *bool `tfsdk:"enable_network_isolation" json:"enableNetworkIsolation,omitempty"`
					VpcConfig                             *struct {
						SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
						Subnets          *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
					} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
				} `tfsdk:"network_config" json:"networkConfig,omitempty"`
				RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
				StoppingCondition *struct {
					MaxRuntimeInSeconds *int64 `tfsdk:"max_runtime_in_seconds" json:"maxRuntimeInSeconds,omitempty"`
				} `tfsdk:"stopping_condition" json:"stoppingCondition,omitempty"`
			} `tfsdk:"monitoring_job_definition" json:"monitoringJobDefinition,omitempty"`
			MonitoringJobDefinitionName *string `tfsdk:"monitoring_job_definition_name" json:"monitoringJobDefinitionName,omitempty"`
			MonitoringType              *string `tfsdk:"monitoring_type" json:"monitoringType,omitempty"`
			ScheduleConfig              *struct {
				DataAnalysisEndTime   *string `tfsdk:"data_analysis_end_time" json:"dataAnalysisEndTime,omitempty"`
				DataAnalysisStartTime *string `tfsdk:"data_analysis_start_time" json:"dataAnalysisStartTime,omitempty"`
				ScheduleExpression    *string `tfsdk:"schedule_expression" json:"scheduleExpression,omitempty"`
			} `tfsdk:"schedule_config" json:"scheduleConfig,omitempty"`
		} `tfsdk:"monitoring_schedule_config" json:"monitoringScheduleConfig,omitempty"`
		MonitoringScheduleName *string `tfsdk:"monitoring_schedule_name" json:"monitoringScheduleName,omitempty"`
		Tags                   *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MonitoringSchedule is the Schema for the MonitoringSchedules API",
		MarkdownDescription: "MonitoringSchedule is the Schema for the MonitoringSchedules API",
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
				Description:         "MonitoringScheduleSpec defines the desired state of MonitoringSchedule.A schedule for a model monitoring job. For information about model monitor,see Amazon SageMaker Model Monitor (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).",
				MarkdownDescription: "MonitoringScheduleSpec defines the desired state of MonitoringSchedule.A schedule for a model monitoring job. For information about model monitor,see Amazon SageMaker Model Monitor (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).",
				Attributes: map[string]schema.Attribute{
					"monitoring_schedule_config": schema.SingleNestedAttribute{
						Description:         "The configuration object that specifies the monitoring schedule and definesthe monitoring job.",
						MarkdownDescription: "The configuration object that specifies the monitoring schedule and definesthe monitoring job.",
						Attributes: map[string]schema.Attribute{
							"monitoring_job_definition": schema.SingleNestedAttribute{
								Description:         "Defines the monitoring job.",
								MarkdownDescription: "Defines the monitoring job.",
								Attributes: map[string]schema.Attribute{
									"baseline_config": schema.SingleNestedAttribute{
										Description:         "Configuration for monitoring constraints and monitoring statistics. Thesebaseline resources are compared against the results of the current job fromthe series of jobs scheduled to collect data periodically.",
										MarkdownDescription: "Configuration for monitoring constraints and monitoring statistics. Thesebaseline resources are compared against the results of the current job fromthe series of jobs scheduled to collect data periodically.",
										Attributes: map[string]schema.Attribute{
											"baselining_job_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"constraints_resource": schema.SingleNestedAttribute{
												Description:         "The constraints resource for a monitoring job.",
												MarkdownDescription: "The constraints resource for a monitoring job.",
												Attributes: map[string]schema.Attribute{
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

											"statistics_resource": schema.SingleNestedAttribute{
												Description:         "The statistics resource for a monitoring job.",
												MarkdownDescription: "The statistics resource for a monitoring job.",
												Attributes: map[string]schema.Attribute{
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

									"environment": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"monitoring_app_specification": schema.SingleNestedAttribute{
										Description:         "Container image configuration object for the monitoring job.",
										MarkdownDescription: "Container image configuration object for the monitoring job.",
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

											"post_analytics_processor_source_uri": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"record_preprocessor_source_uri": schema.StringAttribute{
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

									"monitoring_inputs": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"endpoint_input": schema.SingleNestedAttribute{
													Description:         "Input object for the endpoint",
													MarkdownDescription: "Input object for the endpoint",
													Attributes: map[string]schema.Attribute{
														"end_time_offset": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"endpoint_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"exclude_features_attribute": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"features_attribute": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"inference_attribute": schema.StringAttribute{
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

														"probability_attribute": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"probability_threshold_attribute": schema.Float64Attribute{
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

														"s3_input_mode": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"start_time_offset": schema.StringAttribute{
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

									"monitoring_output_config": schema.SingleNestedAttribute{
										Description:         "The output configuration for monitoring jobs.",
										MarkdownDescription: "The output configuration for monitoring jobs.",
										Attributes: map[string]schema.Attribute{
											"kms_key_id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"monitoring_outputs": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"s3_output": schema.SingleNestedAttribute{
															Description:         "Information about where and how you want to store the results of a monitoringjob.",
															MarkdownDescription: "Information about where and how you want to store the results of a monitoringjob.",
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

									"monitoring_resources": schema.SingleNestedAttribute{
										Description:         "Identifies the resources to deploy for a monitoring job.",
										MarkdownDescription: "Identifies the resources to deploy for a monitoring job.",
										Attributes: map[string]schema.Attribute{
											"cluster_config": schema.SingleNestedAttribute{
												Description:         "Configuration for the cluster used to run model monitoring jobs.",
												MarkdownDescription: "Configuration for the cluster used to run model monitoring jobs.",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network_config": schema.SingleNestedAttribute{
										Description:         "Networking options for a job, such as network traffic encryption betweencontainers, whether to allow inbound and outbound network calls to and fromcontainers, and the VPC subnets and security groups to use for VPC-enabledjobs.",
										MarkdownDescription: "Networking options for a job, such as network traffic encryption betweencontainers, whether to allow inbound and outbound network calls to and fromcontainers, and the VPC subnets and security groups to use for VPC-enabledjobs.",
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

									"role_arn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"stopping_condition": schema.SingleNestedAttribute{
										Description:         "A time limit for how long the monitoring job is allowed to run before stopping.",
										MarkdownDescription: "A time limit for how long the monitoring job is allowed to run before stopping.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"monitoring_job_definition_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"monitoring_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"schedule_config": schema.SingleNestedAttribute{
								Description:         "Configuration details about the monitoring schedule.",
								MarkdownDescription: "Configuration details about the monitoring schedule.",
								Attributes: map[string]schema.Attribute{
									"data_analysis_end_time": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"data_analysis_start_time": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"schedule_expression": schema.StringAttribute{
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

					"monitoring_schedule_name": schema.StringAttribute{
						Description:         "The name of the monitoring schedule. The name must be unique within an AmazonWeb Services Region within an Amazon Web Services account.",
						MarkdownDescription: "The name of the monitoring schedule. The name must be unique within an AmazonWeb Services Region within an Amazon Web Services account.",
						Required:            true,
						Optional:            false,
						Computed:            false,
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

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("MonitoringSchedule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
