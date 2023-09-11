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
	_ datasource.DataSource              = &SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSourceData struct {
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
		JobDefinitionName *string `tfsdk:"job_definition_name" json:"jobDefinitionName,omitempty"`
		JobResources      *struct {
			ClusterConfig *struct {
				InstanceCount  *int64  `tfsdk:"instance_count" json:"instanceCount,omitempty"`
				InstanceType   *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
				VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" json:"volumeKMSKeyID,omitempty"`
				VolumeSizeInGB *int64  `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
			} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
		} `tfsdk:"job_resources" json:"jobResources,omitempty"`
		ModelBiasAppSpecification *struct {
			ConfigURI   *string            `tfsdk:"config_uri" json:"configURI,omitempty"`
			Environment *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
			ImageURI    *string            `tfsdk:"image_uri" json:"imageURI,omitempty"`
		} `tfsdk:"model_bias_app_specification" json:"modelBiasAppSpecification,omitempty"`
		ModelBiasBaselineConfig *struct {
			BaseliningJobName   *string `tfsdk:"baselining_job_name" json:"baseliningJobName,omitempty"`
			ConstraintsResource *struct {
				S3URI *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
			} `tfsdk:"constraints_resource" json:"constraintsResource,omitempty"`
		} `tfsdk:"model_bias_baseline_config" json:"modelBiasBaselineConfig,omitempty"`
		ModelBiasJobInput *struct {
			EndpointInput *struct {
				EndTimeOffset                 *string  `tfsdk:"end_time_offset" json:"endTimeOffset,omitempty"`
				EndpointName                  *string  `tfsdk:"endpoint_name" json:"endpointName,omitempty"`
				FeaturesAttribute             *string  `tfsdk:"features_attribute" json:"featuresAttribute,omitempty"`
				InferenceAttribute            *string  `tfsdk:"inference_attribute" json:"inferenceAttribute,omitempty"`
				LocalPath                     *string  `tfsdk:"local_path" json:"localPath,omitempty"`
				ProbabilityAttribute          *string  `tfsdk:"probability_attribute" json:"probabilityAttribute,omitempty"`
				ProbabilityThresholdAttribute *float64 `tfsdk:"probability_threshold_attribute" json:"probabilityThresholdAttribute,omitempty"`
				S3DataDistributionType        *string  `tfsdk:"s3_data_distribution_type" json:"s3DataDistributionType,omitempty"`
				S3InputMode                   *string  `tfsdk:"s3_input_mode" json:"s3InputMode,omitempty"`
				StartTimeOffset               *string  `tfsdk:"start_time_offset" json:"startTimeOffset,omitempty"`
			} `tfsdk:"endpoint_input" json:"endpointInput,omitempty"`
			GroundTruthS3Input *struct {
				S3URI *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
			} `tfsdk:"ground_truth_s3_input" json:"groundTruthS3Input,omitempty"`
		} `tfsdk:"model_bias_job_input" json:"modelBiasJobInput,omitempty"`
		ModelBiasJobOutputConfig *struct {
			KmsKeyID          *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			MonitoringOutputs *[]struct {
				S3Output *struct {
					LocalPath    *string `tfsdk:"local_path" json:"localPath,omitempty"`
					S3URI        *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
					S3UploadMode *string `tfsdk:"s3_upload_mode" json:"s3UploadMode,omitempty"`
				} `tfsdk:"s3_output" json:"s3Output,omitempty"`
			} `tfsdk:"monitoring_outputs" json:"monitoringOutputs,omitempty"`
		} `tfsdk:"model_bias_job_output_config" json:"modelBiasJobOutputConfig,omitempty"`
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
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_model_bias_job_definition_v1alpha1"
}

func (r *SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ModelBiasJobDefinition is the Schema for the ModelBiasJobDefinitions API",
		MarkdownDescription: "ModelBiasJobDefinition is the Schema for the ModelBiasJobDefinitions API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ModelBiasJobDefinitionSpec defines the desired state of ModelBiasJobDefinition.",
				MarkdownDescription: "ModelBiasJobDefinitionSpec defines the desired state of ModelBiasJobDefinition.",
				Attributes: map[string]schema.Attribute{
					"job_definition_name": schema.StringAttribute{
						Description:         "The name of the bias job definition. The name must be unique within an Amazon Web Services Region in the Amazon Web Services account.",
						MarkdownDescription: "The name of the bias job definition. The name must be unique within an Amazon Web Services Region in the Amazon Web Services account.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"job_resources": schema.SingleNestedAttribute{
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

					"model_bias_app_specification": schema.SingleNestedAttribute{
						Description:         "Configures the model bias job to run a specified Docker container image.",
						MarkdownDescription: "Configures the model bias job to run a specified Docker container image.",
						Attributes: map[string]schema.Attribute{
							"config_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"environment": schema.MapAttribute{
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"model_bias_baseline_config": schema.SingleNestedAttribute{
						Description:         "The baseline configuration for a model bias job.",
						MarkdownDescription: "The baseline configuration for a model bias job.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"model_bias_job_input": schema.SingleNestedAttribute{
						Description:         "Inputs for the model bias job.",
						MarkdownDescription: "Inputs for the model bias job.",
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

									"probability_threshold_attribute": schema.Float64Attribute{
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

							"ground_truth_s3_input": schema.SingleNestedAttribute{
								Description:         "The ground truth labels for the dataset used for the monitoring job.",
								MarkdownDescription: "The ground truth labels for the dataset used for the monitoring job.",
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

					"model_bias_job_output_config": schema.SingleNestedAttribute{
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

					"network_config": schema.SingleNestedAttribute{
						Description:         "Networking options for a model bias job.",
						MarkdownDescription: "Networking options for a model bias job.",
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
						Description:         "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.",
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

func (r *SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_model_bias_job_definition_v1alpha1")

	var data SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "modelbiasjobdefinitions"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse SagemakerServicesK8SAwsModelBiasJobDefinitionV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("ModelBiasJobDefinition")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
