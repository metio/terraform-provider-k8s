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

type LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource)(nil)
)

type LambdaServicesK8SAwsEventSourceMappingV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LambdaServicesK8SAwsEventSourceMappingV1Alpha1GoModel struct {
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
		AmazonManagedKafkaEventSourceConfig *struct {
			ConsumerGroupID *string `tfsdk:"consumer_group_id" yaml:"consumerGroupID,omitempty"`
		} `tfsdk:"amazon_managed_kafka_event_source_config" yaml:"amazonManagedKafkaEventSourceConfig,omitempty"`

		BatchSize *int64 `tfsdk:"batch_size" yaml:"batchSize,omitempty"`

		BisectBatchOnFunctionError *bool `tfsdk:"bisect_batch_on_function_error" yaml:"bisectBatchOnFunctionError,omitempty"`

		DestinationConfig *struct {
			OnFailure *struct {
				Destination *string `tfsdk:"destination" yaml:"destination,omitempty"`
			} `tfsdk:"on_failure" yaml:"onFailure,omitempty"`

			OnSuccess *struct {
				Destination *string `tfsdk:"destination" yaml:"destination,omitempty"`
			} `tfsdk:"on_success" yaml:"onSuccess,omitempty"`
		} `tfsdk:"destination_config" yaml:"destinationConfig,omitempty"`

		Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

		EventSourceARN *string `tfsdk:"event_source_arn" yaml:"eventSourceARN,omitempty"`

		FilterCriteria *struct {
			Filters *[]struct {
				Pattern *string `tfsdk:"pattern" yaml:"pattern,omitempty"`
			} `tfsdk:"filters" yaml:"filters,omitempty"`
		} `tfsdk:"filter_criteria" yaml:"filterCriteria,omitempty"`

		FunctionName *string `tfsdk:"function_name" yaml:"functionName,omitempty"`

		FunctionRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"function_ref" yaml:"functionRef,omitempty"`

		FunctionResponseTypes *[]string `tfsdk:"function_response_types" yaml:"functionResponseTypes,omitempty"`

		MaximumBatchingWindowInSeconds *int64 `tfsdk:"maximum_batching_window_in_seconds" yaml:"maximumBatchingWindowInSeconds,omitempty"`

		MaximumRecordAgeInSeconds *int64 `tfsdk:"maximum_record_age_in_seconds" yaml:"maximumRecordAgeInSeconds,omitempty"`

		MaximumRetryAttempts *int64 `tfsdk:"maximum_retry_attempts" yaml:"maximumRetryAttempts,omitempty"`

		ParallelizationFactor *int64 `tfsdk:"parallelization_factor" yaml:"parallelizationFactor,omitempty"`

		QueueRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"queue_refs" yaml:"queueRefs,omitempty"`

		Queues *[]string `tfsdk:"queues" yaml:"queues,omitempty"`

		SelfManagedEventSource *struct {
			Endpoints *map[string][]string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`
		} `tfsdk:"self_managed_event_source" yaml:"selfManagedEventSource,omitempty"`

		SelfManagedKafkaEventSourceConfig *struct {
			ConsumerGroupID *string `tfsdk:"consumer_group_id" yaml:"consumerGroupID,omitempty"`
		} `tfsdk:"self_managed_kafka_event_source_config" yaml:"selfManagedKafkaEventSourceConfig,omitempty"`

		SourceAccessConfigurations *[]struct {
			Type_ *string `tfsdk:"type_" yaml:"type_,omitempty"`

			URI *string `tfsdk:"u_ri" yaml:"uRI,omitempty"`
		} `tfsdk:"source_access_configurations" yaml:"sourceAccessConfigurations,omitempty"`

		StartingPosition *string `tfsdk:"starting_position" yaml:"startingPosition,omitempty"`

		StartingPositionTimestamp *string `tfsdk:"starting_position_timestamp" yaml:"startingPositionTimestamp,omitempty"`

		Topics *[]string `tfsdk:"topics" yaml:"topics,omitempty"`

		TumblingWindowInSeconds *int64 `tfsdk:"tumbling_window_in_seconds" yaml:"tumblingWindowInSeconds,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource() resource.Resource {
	return &LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource{}
}

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lambda_services_k8s_aws_event_source_mapping_v1alpha1"
}

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "EventSourceMapping is the Schema for the EventSourceMappings API",
		MarkdownDescription: "EventSourceMapping is the Schema for the EventSourceMappings API",
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
				Description:         "EventSourceMappingSpec defines the desired state of EventSourceMapping.",
				MarkdownDescription: "EventSourceMappingSpec defines the desired state of EventSourceMapping.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"amazon_managed_kafka_event_source_config": {
						Description:         "Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.",
						MarkdownDescription: "Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"consumer_group_id": {
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

					"batch_size": {
						Description:         "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).  * Amazon Kinesis - Default 100. Max 10,000.  * Amazon DynamoDB Streams - Default 100. Max 10,000.  * Amazon Simple Queue Service - Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.  * Amazon Managed Streaming for Apache Kafka - Default 100. Max 10,000.  * Self-managed Apache Kafka - Default 100. Max 10,000.  * Amazon MQ (ActiveMQ and RabbitMQ) - Default 100. Max 10,000.",
						MarkdownDescription: "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).  * Amazon Kinesis - Default 100. Max 10,000.  * Amazon DynamoDB Streams - Default 100. Max 10,000.  * Amazon Simple Queue Service - Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.  * Amazon Managed Streaming for Apache Kafka - Default 100. Max 10,000.  * Self-managed Apache Kafka - Default 100. Max 10,000.  * Amazon MQ (ActiveMQ and RabbitMQ) - Default 100. Max 10,000.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bisect_batch_on_function_error": {
						Description:         "(Streams only) If the function returns an error, split the batch in two and retry.",
						MarkdownDescription: "(Streams only) If the function returns an error, split the batch in two and retry.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"destination_config": {
						Description:         "(Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
						MarkdownDescription: "(Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"on_failure": {
								Description:         "A destination for events that failed processing.",
								MarkdownDescription: "A destination for events that failed processing.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"destination": {
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

							"on_success": {
								Description:         "A destination for events that were processed successfully.",
								MarkdownDescription: "A destination for events that were processed successfully.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"destination": {
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

					"enabled": {
						Description:         "When true, the event source mapping is active. When false, Lambda pauses polling and invocation.  Default: True",
						MarkdownDescription: "When true, the event source mapping is active. When false, Lambda pauses polling and invocation.  Default: True",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"event_source_arn": {
						Description:         "The Amazon Resource Name (ARN) of the event source.  * Amazon Kinesis - The ARN of the data stream or a stream consumer.  * Amazon DynamoDB Streams - The ARN of the stream.  * Amazon Simple Queue Service - The ARN of the queue.  * Amazon Managed Streaming for Apache Kafka - The ARN of the cluster.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the event source.  * Amazon Kinesis - The ARN of the data stream or a stream consumer.  * Amazon DynamoDB Streams - The ARN of the stream.  * Amazon Simple Queue Service - The ARN of the queue.  * Amazon Managed Streaming for Apache Kafka - The ARN of the cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"filter_criteria": {
						Description:         "(Streams and Amazon SQS) An object that defines the filter criteria that determine whether Lambda should process an event. For more information, see Lambda event filtering (https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).",
						MarkdownDescription: "(Streams and Amazon SQS) An object that defines the filter criteria that determine whether Lambda should process an event. For more information, see Lambda event filtering (https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"filters": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"pattern": {
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

					"function_name": {
						Description:         "The name of the Lambda function.  Name formats  * Function name - MyFunction.  * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.  * Version or Alias ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.  * Partial ARN - 123456789012:function:MyFunction.  The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.",
						MarkdownDescription: "The name of the Lambda function.  Name formats  * Function name - MyFunction.  * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction.  * Version or Alias ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.  * Partial ARN - 123456789012:function:MyFunction.  The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"function_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"function_response_types": {
						Description:         "(Streams and Amazon SQS) A list of current response type enums applied to the event source mapping.",
						MarkdownDescription: "(Streams and Amazon SQS) A list of current response type enums applied to the event source mapping.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"maximum_batching_window_in_seconds": {
						Description:         "(Streams and Amazon SQS standard queues) The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.  Default: 0  Related setting: When you set BatchSize to a value greater than 10, you must set MaximumBatchingWindowInSeconds to at least 1.",
						MarkdownDescription: "(Streams and Amazon SQS standard queues) The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function.  Default: 0  Related setting: When you set BatchSize to a value greater than 10, you must set MaximumBatchingWindowInSeconds to at least 1.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"maximum_record_age_in_seconds": {
						Description:         "(Streams only) Discard records older than the specified age. The default value is infinite (-1).",
						MarkdownDescription: "(Streams only) Discard records older than the specified age. The default value is infinite (-1).",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"maximum_retry_attempts": {
						Description:         "(Streams only) Discard records after the specified number of retries. The default value is infinite (-1). When set to infinite (-1), failed records are retried until the record expires.",
						MarkdownDescription: "(Streams only) Discard records after the specified number of retries. The default value is infinite (-1). When set to infinite (-1), failed records are retried until the record expires.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"parallelization_factor": {
						Description:         "(Streams only) The number of batches to process from each shard concurrently.",
						MarkdownDescription: "(Streams only) The number of batches to process from each shard concurrently.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"queue_refs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"queues": {
						Description:         "(MQ) The name of the Amazon MQ broker destination queue to consume.",
						MarkdownDescription: "(MQ) The name of the Amazon MQ broker destination queue to consume.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"self_managed_event_source": {
						Description:         "The self-managed Apache Kafka cluster to receive records from.",
						MarkdownDescription: "The self-managed Apache Kafka cluster to receive records from.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"endpoints": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"self_managed_kafka_event_source_config": {
						Description:         "Specific configuration settings for a self-managed Apache Kafka event source.",
						MarkdownDescription: "Specific configuration settings for a self-managed Apache Kafka event source.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"consumer_group_id": {
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

					"source_access_configurations": {
						Description:         "An array of authentication protocols or VPC components required to secure your event source.",
						MarkdownDescription: "An array of authentication protocols or VPC components required to secure your event source.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"type_": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"u_ri": {
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

					"starting_position": {
						Description:         "The position in a stream from which to start reading. Required for Amazon Kinesis, Amazon DynamoDB, and Amazon MSK Streams sources. AT_TIMESTAMP is supported only for Amazon Kinesis streams.",
						MarkdownDescription: "The position in a stream from which to start reading. Required for Amazon Kinesis, Amazon DynamoDB, and Amazon MSK Streams sources. AT_TIMESTAMP is supported only for Amazon Kinesis streams.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"starting_position_timestamp": {
						Description:         "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.",
						MarkdownDescription: "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							validators.DateTime64Validator(),
						},
					},

					"topics": {
						Description:         "The name of the Kafka topic.",
						MarkdownDescription: "The name of the Kafka topic.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tumbling_window_in_seconds": {
						Description:         "(Streams only) The duration in seconds of a processing window. The range is between 1 second and 900 seconds.",
						MarkdownDescription: "(Streams only) The duration in seconds of a processing window. The range is between 1 second and 900 seconds.",

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
		},
	}, nil
}

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_lambda_services_k8s_aws_event_source_mapping_v1alpha1")

	var state LambdaServicesK8SAwsEventSourceMappingV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LambdaServicesK8SAwsEventSourceMappingV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("lambda.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("EventSourceMapping")

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

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lambda_services_k8s_aws_event_source_mapping_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_lambda_services_k8s_aws_event_source_mapping_v1alpha1")

	var state LambdaServicesK8SAwsEventSourceMappingV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LambdaServicesK8SAwsEventSourceMappingV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("lambda.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("EventSourceMapping")

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

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_lambda_services_k8s_aws_event_source_mapping_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
