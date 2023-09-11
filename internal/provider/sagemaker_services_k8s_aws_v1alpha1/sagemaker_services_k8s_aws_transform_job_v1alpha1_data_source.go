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
	_ datasource.DataSource              = &SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsTransformJobV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsTransformJobV1Alpha1DataSourceData struct {
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
		BatchStrategy  *string `tfsdk:"batch_strategy" json:"batchStrategy,omitempty"`
		DataProcessing *struct {
			InputFilter  *string `tfsdk:"input_filter" json:"inputFilter,omitempty"`
			JoinSource   *string `tfsdk:"join_source" json:"joinSource,omitempty"`
			OutputFilter *string `tfsdk:"output_filter" json:"outputFilter,omitempty"`
		} `tfsdk:"data_processing" json:"dataProcessing,omitempty"`
		Environment      *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
		ExperimentConfig *struct {
			ExperimentName            *string `tfsdk:"experiment_name" json:"experimentName,omitempty"`
			TrialComponentDisplayName *string `tfsdk:"trial_component_display_name" json:"trialComponentDisplayName,omitempty"`
			TrialName                 *string `tfsdk:"trial_name" json:"trialName,omitempty"`
		} `tfsdk:"experiment_config" json:"experimentConfig,omitempty"`
		MaxConcurrentTransforms *int64 `tfsdk:"max_concurrent_transforms" json:"maxConcurrentTransforms,omitempty"`
		MaxPayloadInMB          *int64 `tfsdk:"max_payload_in_mb" json:"maxPayloadInMB,omitempty"`
		ModelClientConfig       *struct {
			InvocationsMaxRetries       *int64 `tfsdk:"invocations_max_retries" json:"invocationsMaxRetries,omitempty"`
			InvocationsTimeoutInSeconds *int64 `tfsdk:"invocations_timeout_in_seconds" json:"invocationsTimeoutInSeconds,omitempty"`
		} `tfsdk:"model_client_config" json:"modelClientConfig,omitempty"`
		ModelName *string `tfsdk:"model_name" json:"modelName,omitempty"`
		Tags      *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TransformInput *struct {
			CompressionType *string `tfsdk:"compression_type" json:"compressionType,omitempty"`
			ContentType     *string `tfsdk:"content_type" json:"contentType,omitempty"`
			DataSource      *struct {
				S3DataSource *struct {
					S3DataType *string `tfsdk:"s3_data_type" json:"s3DataType,omitempty"`
					S3URI      *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"s3_data_source" json:"s3DataSource,omitempty"`
			} `tfsdk:"data_source" json:"dataSource,omitempty"`
			SplitType *string `tfsdk:"split_type" json:"splitType,omitempty"`
		} `tfsdk:"transform_input" json:"transformInput,omitempty"`
		TransformJobName *string `tfsdk:"transform_job_name" json:"transformJobName,omitempty"`
		TransformOutput  *struct {
			Accept       *string `tfsdk:"accept" json:"accept,omitempty"`
			AssembleWith *string `tfsdk:"assemble_with" json:"assembleWith,omitempty"`
			KmsKeyID     *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			S3OutputPath *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
		} `tfsdk:"transform_output" json:"transformOutput,omitempty"`
		TransformResources *struct {
			InstanceCount  *int64  `tfsdk:"instance_count" json:"instanceCount,omitempty"`
			InstanceType   *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
			VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" json:"volumeKMSKeyID,omitempty"`
		} `tfsdk:"transform_resources" json:"transformResources,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_transform_job_v1alpha1"
}

func (r *SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "TransformJob is the Schema for the TransformJobs API",
		MarkdownDescription: "TransformJob is the Schema for the TransformJobs API",
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
				Description:         "TransformJobSpec defines the desired state of TransformJob.  A batch transform job. For information about SageMaker batch transform, see Use Batch Transform (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform.html).",
				MarkdownDescription: "TransformJobSpec defines the desired state of TransformJob.  A batch transform job. For information about SageMaker batch transform, see Use Batch Transform (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform.html).",
				Attributes: map[string]schema.Attribute{
					"batch_strategy": schema.StringAttribute{
						Description:         "Specifies the number of records to include in a mini-batch for an HTTP inference request. A record is a single unit of input data that inference can be made on. For example, a single line in a CSV file is a record.  To enable the batch strategy, you must set the SplitType property to Line, RecordIO, or TFRecord.  To use only one record when making an HTTP invocation request to a container, set BatchStrategy to SingleRecord and SplitType to Line.  To fit as many records in a mini-batch as can fit within the MaxPayloadInMB limit, set BatchStrategy to MultiRecord and SplitType to Line.",
						MarkdownDescription: "Specifies the number of records to include in a mini-batch for an HTTP inference request. A record is a single unit of input data that inference can be made on. For example, a single line in a CSV file is a record.  To enable the batch strategy, you must set the SplitType property to Line, RecordIO, or TFRecord.  To use only one record when making an HTTP invocation request to a container, set BatchStrategy to SingleRecord and SplitType to Line.  To fit as many records in a mini-batch as can fit within the MaxPayloadInMB limit, set BatchStrategy to MultiRecord and SplitType to Line.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"data_processing": schema.SingleNestedAttribute{
						Description:         "The data structure used to specify the data to be used for inference in a batch transform job and to associate the data that is relevant to the prediction results in the output. The input filter provided allows you to exclude input data that is not needed for inference in a batch transform job. The output filter provided allows you to include input data relevant to interpreting the predictions in the output from the job. For more information, see Associate Prediction Results with their Corresponding Input Records (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html).",
						MarkdownDescription: "The data structure used to specify the data to be used for inference in a batch transform job and to associate the data that is relevant to the prediction results in the output. The input filter provided allows you to exclude input data that is not needed for inference in a batch transform job. The output filter provided allows you to include input data relevant to interpreting the predictions in the output from the job. For more information, see Associate Prediction Results with their Corresponding Input Records (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html).",
						Attributes: map[string]schema.Attribute{
							"input_filter": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"join_source": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"output_filter": schema.StringAttribute{
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

					"environment": schema.MapAttribute{
						Description:         "The environment variables to set in the Docker container. We support up to 16 key and values entries in the map.",
						MarkdownDescription: "The environment variables to set in the Docker container. We support up to 16 key and values entries in the map.",
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

					"max_concurrent_transforms": schema.Int64Attribute{
						Description:         "The maximum number of parallel requests that can be sent to each instance in a transform job. If MaxConcurrentTransforms is set to 0 or left unset, Amazon SageMaker checks the optional execution-parameters to determine the settings for your chosen algorithm. If the execution-parameters endpoint is not enabled, the default value is 1. For more information on execution-parameters, see How Containers Serve Requests (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-batch-code.html#your-algorithms-batch-code-how-containe-serves-requests). For built-in algorithms, you don't need to set a value for MaxConcurrentTransforms.",
						MarkdownDescription: "The maximum number of parallel requests that can be sent to each instance in a transform job. If MaxConcurrentTransforms is set to 0 or left unset, Amazon SageMaker checks the optional execution-parameters to determine the settings for your chosen algorithm. If the execution-parameters endpoint is not enabled, the default value is 1. For more information on execution-parameters, see How Containers Serve Requests (https://docs.aws.amazon.com/sagemaker/latest/dg/your-algorithms-batch-code.html#your-algorithms-batch-code-how-containe-serves-requests). For built-in algorithms, you don't need to set a value for MaxConcurrentTransforms.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_payload_in_mb": schema.Int64Attribute{
						Description:         "The maximum allowed size of the payload, in MB. A payload is the data portion of a record (without metadata). The value in MaxPayloadInMB must be greater than, or equal to, the size of a single record. To estimate the size of a record in MB, divide the size of your dataset by the number of records. To ensure that the records fit within the maximum payload size, we recommend using a slightly larger value. The default value is 6 MB.  The value of MaxPayloadInMB cannot be greater than 100 MB. If you specify the MaxConcurrentTransforms parameter, the value of (MaxConcurrentTransforms * MaxPayloadInMB) also cannot exceed 100 MB.  For cases where the payload might be arbitrarily large and is transmitted using HTTP chunked encoding, set the value to 0. This feature works only in supported algorithms. Currently, Amazon SageMaker built-in algorithms do not support HTTP chunked encoding.",
						MarkdownDescription: "The maximum allowed size of the payload, in MB. A payload is the data portion of a record (without metadata). The value in MaxPayloadInMB must be greater than, or equal to, the size of a single record. To estimate the size of a record in MB, divide the size of your dataset by the number of records. To ensure that the records fit within the maximum payload size, we recommend using a slightly larger value. The default value is 6 MB.  The value of MaxPayloadInMB cannot be greater than 100 MB. If you specify the MaxConcurrentTransforms parameter, the value of (MaxConcurrentTransforms * MaxPayloadInMB) also cannot exceed 100 MB.  For cases where the payload might be arbitrarily large and is transmitted using HTTP chunked encoding, set the value to 0. This feature works only in supported algorithms. Currently, Amazon SageMaker built-in algorithms do not support HTTP chunked encoding.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"model_client_config": schema.SingleNestedAttribute{
						Description:         "Configures the timeout and maximum number of retries for processing a transform job invocation.",
						MarkdownDescription: "Configures the timeout and maximum number of retries for processing a transform job invocation.",
						Attributes: map[string]schema.Attribute{
							"invocations_max_retries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"invocations_timeout_in_seconds": schema.Int64Attribute{
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

					"model_name": schema.StringAttribute{
						Description:         "The name of the model that you want to use for the transform job. ModelName must be the name of an existing Amazon SageMaker model within an Amazon Web Services Region in an Amazon Web Services account.",
						MarkdownDescription: "The name of the model that you want to use for the transform job. ModelName must be the name of an existing Amazon SageMaker model within an Amazon Web Services Region in an Amazon Web Services account.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the Amazon Web Services Billing and Cost Management User Guide.",
						MarkdownDescription: "(Optional) An array of key-value pairs. For more information, see Using Cost Allocation Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the Amazon Web Services Billing and Cost Management User Guide.",
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

					"transform_input": schema.SingleNestedAttribute{
						Description:         "Describes the input source and the way the transform job consumes it.",
						MarkdownDescription: "Describes the input source and the way the transform job consumes it.",
						Attributes: map[string]schema.Attribute{
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
									"s3_data_source": schema.SingleNestedAttribute{
										Description:         "Describes the S3 data source.",
										MarkdownDescription: "Describes the S3 data source.",
										Attributes: map[string]schema.Attribute{
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

							"split_type": schema.StringAttribute{
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

					"transform_job_name": schema.StringAttribute{
						Description:         "The name of the transform job. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account.",
						MarkdownDescription: "The name of the transform job. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"transform_output": schema.SingleNestedAttribute{
						Description:         "Describes the results of the transform job.",
						MarkdownDescription: "Describes the results of the transform job.",
						Attributes: map[string]schema.Attribute{
							"accept": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"assemble_with": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

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

					"transform_resources": schema.SingleNestedAttribute{
						Description:         "Describes the resources, including ML instance types and ML instance count, to use for the transform job.",
						MarkdownDescription: "Describes the resources, including ML instance types and ML instance count, to use for the transform job.",
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

func (r *SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SagemakerServicesK8SAwsTransformJobV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_transform_job_v1alpha1")

	var data SagemakerServicesK8SAwsTransformJobV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "transformjobs"}).
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

	var readResponse SagemakerServicesK8SAwsTransformJobV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("TransformJob")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
