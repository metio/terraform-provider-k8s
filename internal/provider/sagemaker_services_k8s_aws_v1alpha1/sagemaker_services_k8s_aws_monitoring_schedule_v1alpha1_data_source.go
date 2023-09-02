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
	_ datasource.DataSource              = &SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSourceData struct {
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
						EndTimeOffset                 *string    `tfsdk:"end_time_offset" json:"endTimeOffset,omitempty"`
						EndpointName                  *string    `tfsdk:"endpoint_name" json:"endpointName,omitempty"`
						FeaturesAttribute             *string    `tfsdk:"features_attribute" json:"featuresAttribute,omitempty"`
						InferenceAttribute            *string    `tfsdk:"inference_attribute" json:"inferenceAttribute,omitempty"`
						LocalPath                     *string    `tfsdk:"local_path" json:"localPath,omitempty"`
						ProbabilityAttribute          *string    `tfsdk:"probability_attribute" json:"probabilityAttribute,omitempty"`
						ProbabilityThresholdAttribute *big.Float `tfsdk:"probability_threshold_attribute" json:"probabilityThresholdAttribute,omitempty"`
						S3DataDistributionType        *string    `tfsdk:"s3_data_distribution_type" json:"s3DataDistributionType,omitempty"`
						S3InputMode                   *string    `tfsdk:"s3_input_mode" json:"s3InputMode,omitempty"`
						StartTimeOffset               *string    `tfsdk:"start_time_offset" json:"startTimeOffset,omitempty"`
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
				ScheduleExpression *string `tfsdk:"schedule_expression" json:"scheduleExpression,omitempty"`
			} `tfsdk:"schedule_config" json:"scheduleConfig,omitempty"`
		} `tfsdk:"monitoring_schedule_config" json:"monitoringScheduleConfig,omitempty"`
		MonitoringScheduleName *string `tfsdk:"monitoring_schedule_name" json:"monitoringScheduleName,omitempty"`
		Tags                   *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1"
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MonitoringSchedule is the Schema for the MonitoringSchedules API",
		MarkdownDescription: "MonitoringSchedule is the Schema for the MonitoringSchedules API",
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
				Description:         "MonitoringScheduleSpec defines the desired state of MonitoringSchedule.  A schedule for a model monitoring job. For information about model monitor, see Amazon SageMaker Model Monitor (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).",
				MarkdownDescription: "MonitoringScheduleSpec defines the desired state of MonitoringSchedule.  A schedule for a model monitoring job. For information about model monitor, see Amazon SageMaker Model Monitor (https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html).",
				Attributes: map[string]schema.Attribute{
					"monitoring_schedule_config": schema.SingleNestedAttribute{
						Description:         "The configuration object that specifies the monitoring schedule and defines the monitoring job.",
						MarkdownDescription: "The configuration object that specifies the monitoring schedule and defines the monitoring job.",
						Attributes: map[string]schema.Attribute{
							"monitoring_job_definition": schema.SingleNestedAttribute{
								Description:         "Defines the monitoring job.",
								MarkdownDescription: "Defines the monitoring job.",
								Attributes: map[string]schema.Attribute{
									"baseline_config": schema.SingleNestedAttribute{
										Description:         "Configuration for monitoring constraints and monitoring statistics. These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.",
										MarkdownDescription: "Configuration for monitoring constraints and monitoring statistics. These baseline resources are compared against the results of the current job from the series of jobs scheduled to collect data periodically.",
										Attributes: map[string]schema.Attribute{
											"baselining_job_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"constraints_resource": schema.SingleNestedAttribute{
												Description:         "The constraints resource for a monitoring job.",
												MarkdownDescription: "The constraints resource for a monitoring job.",
												Attributes: map[string]schema.Attribute{
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

											"statistics_resource": schema.SingleNestedAttribute{
												Description:         "The statistics resource for a monitoring job.",
												MarkdownDescription: "The statistics resource for a monitoring job.",
												Attributes: map[string]schema.Attribute{
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

									"environment": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"container_entrypoint": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"image_uri": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"post_analytics_processor_source_uri": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"record_preprocessor_source_uri": schema.StringAttribute{
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
															Optional:            false,
															Computed:            true,
														},

														"endpoint_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"features_attribute": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"inference_attribute": schema.StringAttribute{
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

														"probability_attribute": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"probability_threshold_attribute": types.NumberType{
															Description:         "",
															MarkdownDescription: "",
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

														"s3_input_mode": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"start_time_offset": schema.StringAttribute{
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

									"monitoring_output_config": schema.SingleNestedAttribute{
										Description:         "The output configuration for monitoring jobs.",
										MarkdownDescription: "The output configuration for monitoring jobs.",
										Attributes: map[string]schema.Attribute{
											"kms_key_id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"monitoring_outputs": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"s3_output": schema.SingleNestedAttribute{
															Description:         "Information about where and how you want to store the results of a monitoring job.",
															MarkdownDescription: "Information about where and how you want to store the results of a monitoring job.",
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

																"s3_upload_mode": schema.StringAttribute{
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
										},
										Required: false,
										Optional: false,
										Computed: true,
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"network_config": schema.SingleNestedAttribute{
										Description:         "Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.",
										MarkdownDescription: "Networking options for a job, such as network traffic encryption between containers, whether to allow inbound and outbound network calls to and from containers, and the VPC subnets and security groups to use for VPC-enabled jobs.",
										Attributes: map[string]schema.Attribute{
											"enable_inter_container_traffic_encryption": schema.BoolAttribute{
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

									"role_arn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"stopping_condition": schema.SingleNestedAttribute{
										Description:         "A time limit for how long the monitoring job is allowed to run before stopping.",
										MarkdownDescription: "A time limit for how long the monitoring job is allowed to run before stopping.",
										Attributes: map[string]schema.Attribute{
											"max_runtime_in_seconds": schema.Int64Attribute{
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

							"monitoring_job_definition_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"monitoring_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"schedule_config": schema.SingleNestedAttribute{
								Description:         "Configuration details about the monitoring schedule.",
								MarkdownDescription: "Configuration details about the monitoring schedule.",
								Attributes: map[string]schema.Attribute{
									"schedule_expression": schema.StringAttribute{
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

					"monitoring_schedule_name": schema.StringAttribute{
						Description:         "The name of the monitoring schedule. The name must be unique within an Amazon Web Services Region within an Amazon Web Services account.",
						MarkdownDescription: "The name of the monitoring schedule. The name must be unique within an Amazon Web Services Region within an Amazon Web Services account.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the Amazon Web Services Billing and Cost Management User Guide.",
						MarkdownDescription: "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-whatURL) in the Amazon Web Services Billing and Cost Management User Guide.",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_monitoring_schedule_v1alpha1")

	var data SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "MonitoringSchedule"}).
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

	var readResponse SagemakerServicesK8SAwsMonitoringScheduleV1Alpha1DataSourceData
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
	data.Kind = pointer.String("MonitoringSchedule")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
