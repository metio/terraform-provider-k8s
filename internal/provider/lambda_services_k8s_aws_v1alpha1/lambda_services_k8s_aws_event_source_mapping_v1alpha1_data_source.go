/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package lambda_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource              = &LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource{}
)

func NewLambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource() datasource.DataSource {
	return &LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource{}
}

type LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSourceData struct {
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
		AmazonManagedKafkaEventSourceConfig *struct {
			ConsumerGroupID *string `tfsdk:"consumer_group_id" json:"consumerGroupID,omitempty"`
		} `tfsdk:"amazon_managed_kafka_event_source_config" json:"amazonManagedKafkaEventSourceConfig,omitempty"`
		BatchSize                  *int64 `tfsdk:"batch_size" json:"batchSize,omitempty"`
		BisectBatchOnFunctionError *bool  `tfsdk:"bisect_batch_on_function_error" json:"bisectBatchOnFunctionError,omitempty"`
		DestinationConfig          *struct {
			OnFailure *struct {
				Destination *string `tfsdk:"destination" json:"destination,omitempty"`
			} `tfsdk:"on_failure" json:"onFailure,omitempty"`
			OnSuccess *struct {
				Destination *string `tfsdk:"destination" json:"destination,omitempty"`
			} `tfsdk:"on_success" json:"onSuccess,omitempty"`
		} `tfsdk:"destination_config" json:"destinationConfig,omitempty"`
		Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
		EventSourceARN *string `tfsdk:"event_source_arn" json:"eventSourceARN,omitempty"`
		EventSourceRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"event_source_ref" json:"eventSourceRef,omitempty"`
		FilterCriteria *struct {
			Filters *[]struct {
				Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
			} `tfsdk:"filters" json:"filters,omitempty"`
		} `tfsdk:"filter_criteria" json:"filterCriteria,omitempty"`
		FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
		FunctionRef  *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"function_ref" json:"functionRef,omitempty"`
		FunctionResponseTypes          *[]string `tfsdk:"function_response_types" json:"functionResponseTypes,omitempty"`
		MaximumBatchingWindowInSeconds *int64    `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
		MaximumRecordAgeInSeconds      *int64    `tfsdk:"maximum_record_age_in_seconds" json:"maximumRecordAgeInSeconds,omitempty"`
		MaximumRetryAttempts           *int64    `tfsdk:"maximum_retry_attempts" json:"maximumRetryAttempts,omitempty"`
		ParallelizationFactor          *int64    `tfsdk:"parallelization_factor" json:"parallelizationFactor,omitempty"`
		QueueRefs                      *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"queue_refs" json:"queueRefs,omitempty"`
		Queues        *[]string `tfsdk:"queues" json:"queues,omitempty"`
		ScalingConfig *struct {
			MaximumConcurrency *int64 `tfsdk:"maximum_concurrency" json:"maximumConcurrency,omitempty"`
		} `tfsdk:"scaling_config" json:"scalingConfig,omitempty"`
		SelfManagedEventSource *struct {
			Endpoints *map[string][]string `tfsdk:"endpoints" json:"endpoints,omitempty"`
		} `tfsdk:"self_managed_event_source" json:"selfManagedEventSource,omitempty"`
		SelfManagedKafkaEventSourceConfig *struct {
			ConsumerGroupID *string `tfsdk:"consumer_group_id" json:"consumerGroupID,omitempty"`
		} `tfsdk:"self_managed_kafka_event_source_config" json:"selfManagedKafkaEventSourceConfig,omitempty"`
		SourceAccessConfigurations *[]struct {
			Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
			URI   *string `tfsdk:"u_ri" json:"uRI,omitempty"`
		} `tfsdk:"source_access_configurations" json:"sourceAccessConfigurations,omitempty"`
		StartingPosition          *string   `tfsdk:"starting_position" json:"startingPosition,omitempty"`
		StartingPositionTimestamp *string   `tfsdk:"starting_position_timestamp" json:"startingPositionTimestamp,omitempty"`
		Topics                    *[]string `tfsdk:"topics" json:"topics,omitempty"`
		TumblingWindowInSeconds   *int64    `tfsdk:"tumbling_window_in_seconds" json:"tumblingWindowInSeconds,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lambda_services_k8s_aws_event_source_mapping_v1alpha1"
}

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EventSourceMapping is the Schema for the EventSourceMappings API",
		MarkdownDescription: "EventSourceMapping is the Schema for the EventSourceMappings API",
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
				Description:         "EventSourceMappingSpec defines the desired state of EventSourceMapping.",
				MarkdownDescription: "EventSourceMappingSpec defines the desired state of EventSourceMapping.",
				Attributes: map[string]schema.Attribute{
					"amazon_managed_kafka_event_source_config": schema.SingleNestedAttribute{
						Description:         "Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.",
						MarkdownDescription: "Specific configuration settings for an Amazon Managed Streaming for Apache Kafka (Amazon MSK) event source.",
						Attributes: map[string]schema.Attribute{
							"consumer_group_id": schema.StringAttribute{
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

					"batch_size": schema.Int64Attribute{
						Description:         "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).  * Amazon Kinesis – Default 100. Max 10,000.  * Amazon DynamoDB Streams – Default 100. Max 10,000.  * Amazon Simple Queue Service – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.  * Amazon Managed Streaming for Apache Kafka – Default 100. Max 10,000.  * Self-managed Apache Kafka – Default 100. Max 10,000.  * Amazon MQ (ActiveMQ and RabbitMQ) – Default 100. Max 10,000.",
						MarkdownDescription: "The maximum number of records in each batch that Lambda pulls from your stream or queue and sends to your function. Lambda passes all of the records in the batch to the function in a single call, up to the payload limit for synchronous invocation (6 MB).  * Amazon Kinesis – Default 100. Max 10,000.  * Amazon DynamoDB Streams – Default 100. Max 10,000.  * Amazon Simple Queue Service – Default 10. For standard queues the max is 10,000. For FIFO queues the max is 10.  * Amazon Managed Streaming for Apache Kafka – Default 100. Max 10,000.  * Self-managed Apache Kafka – Default 100. Max 10,000.  * Amazon MQ (ActiveMQ and RabbitMQ) – Default 100. Max 10,000.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"bisect_batch_on_function_error": schema.BoolAttribute{
						Description:         "(Streams only) If the function returns an error, split the batch in two and retry.",
						MarkdownDescription: "(Streams only) If the function returns an error, split the batch in two and retry.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"destination_config": schema.SingleNestedAttribute{
						Description:         "(Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
						MarkdownDescription: "(Streams only) An Amazon SQS queue or Amazon SNS topic destination for discarded records.",
						Attributes: map[string]schema.Attribute{
							"on_failure": schema.SingleNestedAttribute{
								Description:         "A destination for events that failed processing.",
								MarkdownDescription: "A destination for events that failed processing.",
								Attributes: map[string]schema.Attribute{
									"destination": schema.StringAttribute{
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

							"on_success": schema.SingleNestedAttribute{
								Description:         "A destination for events that were processed successfully.",
								MarkdownDescription: "A destination for events that were processed successfully.",
								Attributes: map[string]schema.Attribute{
									"destination": schema.StringAttribute{
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

					"enabled": schema.BoolAttribute{
						Description:         "When true, the event source mapping is active. When false, Lambda pauses polling and invocation.  Default: True",
						MarkdownDescription: "When true, the event source mapping is active. When false, Lambda pauses polling and invocation.  Default: True",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"event_source_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the event source.  * Amazon Kinesis – The ARN of the data stream or a stream consumer.  * Amazon DynamoDB Streams – The ARN of the stream.  * Amazon Simple Queue Service – The ARN of the queue.  * Amazon Managed Streaming for Apache Kafka – The ARN of the cluster.  * Amazon MQ – The ARN of the broker.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the event source.  * Amazon Kinesis – The ARN of the data stream or a stream consumer.  * Amazon DynamoDB Streams – The ARN of the stream.  * Amazon Simple Queue Service – The ARN of the queue.  * Amazon Managed Streaming for Apache Kafka – The ARN of the cluster.  * Amazon MQ – The ARN of the broker.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"event_source_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"filter_criteria": schema.SingleNestedAttribute{
						Description:         "An object that defines the filter criteria that determine whether Lambda should process an event. For more information, see Lambda event filtering (https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).",
						MarkdownDescription: "An object that defines the filter criteria that determine whether Lambda should process an event. For more information, see Lambda event filtering (https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html).",
						Attributes: map[string]schema.Attribute{
							"filters": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"pattern": schema.StringAttribute{
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

					"function_name": schema.StringAttribute{
						Description:         "The name of the Lambda function.  Name formats  * Function name – MyFunction.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:MyFunction.  * Version or Alias ARN – arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.  * Partial ARN – 123456789012:function:MyFunction.  The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.",
						MarkdownDescription: "The name of the Lambda function.  Name formats  * Function name – MyFunction.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:MyFunction.  * Version or Alias ARN – arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.  * Partial ARN – 123456789012:function:MyFunction.  The length constraint applies only to the full ARN. If you specify only the function name, it's limited to 64 characters in length.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"function_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

					"function_response_types": schema.ListAttribute{
						Description:         "(Streams and Amazon SQS) A list of current response type enums applied to the event source mapping.",
						MarkdownDescription: "(Streams and Amazon SQS) A list of current response type enums applied to the event source mapping.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"maximum_batching_window_in_seconds": schema.Int64Attribute{
						Description:         "The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function. You can configure MaximumBatchingWindowInSeconds to any value from 0 seconds to 300 seconds in increments of seconds.  For streams and Amazon SQS event sources, the default batching window is 0 seconds. For Amazon MSK, Self-managed Apache Kafka, and Amazon MQ event sources, the default batching window is 500 ms. Note that because you can only change MaximumBatchingWindowInSeconds in increments of seconds, you cannot revert back to the 500 ms default batching window after you have changed it. To restore the default batching window, you must create a new event source mapping.  Related setting: For streams and Amazon SQS event sources, when you set BatchSize to a value greater than 10, you must set MaximumBatchingWindowInSeconds to at least 1.",
						MarkdownDescription: "The maximum amount of time, in seconds, that Lambda spends gathering records before invoking the function. You can configure MaximumBatchingWindowInSeconds to any value from 0 seconds to 300 seconds in increments of seconds.  For streams and Amazon SQS event sources, the default batching window is 0 seconds. For Amazon MSK, Self-managed Apache Kafka, and Amazon MQ event sources, the default batching window is 500 ms. Note that because you can only change MaximumBatchingWindowInSeconds in increments of seconds, you cannot revert back to the 500 ms default batching window after you have changed it. To restore the default batching window, you must create a new event source mapping.  Related setting: For streams and Amazon SQS event sources, when you set BatchSize to a value greater than 10, you must set MaximumBatchingWindowInSeconds to at least 1.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"maximum_record_age_in_seconds": schema.Int64Attribute{
						Description:         "(Streams only) Discard records older than the specified age. The default value is infinite (-1).",
						MarkdownDescription: "(Streams only) Discard records older than the specified age. The default value is infinite (-1).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"maximum_retry_attempts": schema.Int64Attribute{
						Description:         "(Streams only) Discard records after the specified number of retries. The default value is infinite (-1). When set to infinite (-1), failed records are retried until the record expires.",
						MarkdownDescription: "(Streams only) Discard records after the specified number of retries. The default value is infinite (-1). When set to infinite (-1), failed records are retried until the record expires.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"parallelization_factor": schema.Int64Attribute{
						Description:         "(Streams only) The number of batches to process from each shard concurrently.",
						MarkdownDescription: "(Streams only) The number of batches to process from each shard concurrently.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"queue_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

					"queues": schema.ListAttribute{
						Description:         "(MQ) The name of the Amazon MQ broker destination queue to consume.",
						MarkdownDescription: "(MQ) The name of the Amazon MQ broker destination queue to consume.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scaling_config": schema.SingleNestedAttribute{
						Description:         "(Amazon SQS only) The scaling configuration for the event source. For more information, see Configuring maximum concurrency for Amazon SQS event sources (https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency).",
						MarkdownDescription: "(Amazon SQS only) The scaling configuration for the event source. For more information, see Configuring maximum concurrency for Amazon SQS event sources (https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency).",
						Attributes: map[string]schema.Attribute{
							"maximum_concurrency": schema.Int64Attribute{
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

					"self_managed_event_source": schema.SingleNestedAttribute{
						Description:         "The self-managed Apache Kafka cluster to receive records from.",
						MarkdownDescription: "The self-managed Apache Kafka cluster to receive records from.",
						Attributes: map[string]schema.Attribute{
							"endpoints": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"self_managed_kafka_event_source_config": schema.SingleNestedAttribute{
						Description:         "Specific configuration settings for a self-managed Apache Kafka event source.",
						MarkdownDescription: "Specific configuration settings for a self-managed Apache Kafka event source.",
						Attributes: map[string]schema.Attribute{
							"consumer_group_id": schema.StringAttribute{
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

					"source_access_configurations": schema.ListNestedAttribute{
						Description:         "An array of authentication protocols or VPC components required to secure your event source.",
						MarkdownDescription: "An array of authentication protocols or VPC components required to secure your event source.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type_": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"u_ri": schema.StringAttribute{
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

					"starting_position": schema.StringAttribute{
						Description:         "The position in a stream from which to start reading. Required for Amazon Kinesis, Amazon DynamoDB, and Amazon MSK Streams sources. AT_TIMESTAMP is supported only for Amazon Kinesis streams.",
						MarkdownDescription: "The position in a stream from which to start reading. Required for Amazon Kinesis, Amazon DynamoDB, and Amazon MSK Streams sources. AT_TIMESTAMP is supported only for Amazon Kinesis streams.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"starting_position_timestamp": schema.StringAttribute{
						Description:         "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.",
						MarkdownDescription: "With StartingPosition set to AT_TIMESTAMP, the time from which to start reading.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"topics": schema.ListAttribute{
						Description:         "The name of the Kafka topic.",
						MarkdownDescription: "The name of the Kafka topic.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tumbling_window_in_seconds": schema.Int64Attribute{
						Description:         "(Streams only) The duration in seconds of a processing window. The range is between 1 second and 900 seconds.",
						MarkdownDescription: "(Streams only) The duration in seconds of a processing window. The range is between 1 second and 900 seconds.",
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
	}
}

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_lambda_services_k8s_aws_event_source_mapping_v1alpha1")

	var data LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lambda.services.k8s.aws", Version: "v1alpha1", Resource: "eventsourcemappings"}).
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

	var readResponse LambdaServicesK8SAwsEventSourceMappingV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("lambda.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("EventSourceMapping")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
