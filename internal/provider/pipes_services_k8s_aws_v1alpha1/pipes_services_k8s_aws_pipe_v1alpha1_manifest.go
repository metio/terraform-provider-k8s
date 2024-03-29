/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package pipes_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &PipesServicesK8SAwsPipeV1Alpha1Manifest{}
)

func NewPipesServicesK8SAwsPipeV1Alpha1Manifest() datasource.DataSource {
	return &PipesServicesK8SAwsPipeV1Alpha1Manifest{}
}

type PipesServicesK8SAwsPipeV1Alpha1Manifest struct{}

type PipesServicesK8SAwsPipeV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Description          *string `tfsdk:"description" json:"description,omitempty"`
		DesiredState         *string `tfsdk:"desired_state" json:"desiredState,omitempty"`
		Enrichment           *string `tfsdk:"enrichment" json:"enrichment,omitempty"`
		EnrichmentParameters *struct {
			HttpParameters *struct {
				HeaderParameters      *map[string]string `tfsdk:"header_parameters" json:"headerParameters,omitempty"`
				PathParameterValues   *[]string          `tfsdk:"path_parameter_values" json:"pathParameterValues,omitempty"`
				QueryStringParameters *map[string]string `tfsdk:"query_string_parameters" json:"queryStringParameters,omitempty"`
			} `tfsdk:"http_parameters" json:"httpParameters,omitempty"`
			InputTemplate *string `tfsdk:"input_template" json:"inputTemplate,omitempty"`
		} `tfsdk:"enrichment_parameters" json:"enrichmentParameters,omitempty"`
		Name             *string `tfsdk:"name" json:"name,omitempty"`
		RoleARN          *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		Source           *string `tfsdk:"source" json:"source,omitempty"`
		SourceParameters *struct {
			ActiveMQBrokerParameters *struct {
				BatchSize   *int64 `tfsdk:"batch_size" json:"batchSize,omitempty"`
				Credentials *struct {
					BasicAuth *string `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				MaximumBatchingWindowInSeconds *int64  `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
				QueueName                      *string `tfsdk:"queue_name" json:"queueName,omitempty"`
			} `tfsdk:"active_mq_broker_parameters" json:"activeMQBrokerParameters,omitempty"`
			DynamoDBStreamParameters *struct {
				BatchSize        *int64 `tfsdk:"batch_size" json:"batchSize,omitempty"`
				DeadLetterConfig *struct {
					Arn *string `tfsdk:"arn" json:"arn,omitempty"`
				} `tfsdk:"dead_letter_config" json:"deadLetterConfig,omitempty"`
				MaximumBatchingWindowInSeconds *int64  `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
				MaximumRecordAgeInSeconds      *int64  `tfsdk:"maximum_record_age_in_seconds" json:"maximumRecordAgeInSeconds,omitempty"`
				MaximumRetryAttempts           *int64  `tfsdk:"maximum_retry_attempts" json:"maximumRetryAttempts,omitempty"`
				OnPartialBatchItemFailure      *string `tfsdk:"on_partial_batch_item_failure" json:"onPartialBatchItemFailure,omitempty"`
				ParallelizationFactor          *int64  `tfsdk:"parallelization_factor" json:"parallelizationFactor,omitempty"`
				StartingPosition               *string `tfsdk:"starting_position" json:"startingPosition,omitempty"`
			} `tfsdk:"dynamo_db_stream_parameters" json:"dynamoDBStreamParameters,omitempty"`
			FilterCriteria *struct {
				Filters *[]struct {
					Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
				} `tfsdk:"filters" json:"filters,omitempty"`
			} `tfsdk:"filter_criteria" json:"filterCriteria,omitempty"`
			KinesisStreamParameters *struct {
				BatchSize        *int64 `tfsdk:"batch_size" json:"batchSize,omitempty"`
				DeadLetterConfig *struct {
					Arn *string `tfsdk:"arn" json:"arn,omitempty"`
				} `tfsdk:"dead_letter_config" json:"deadLetterConfig,omitempty"`
				MaximumBatchingWindowInSeconds *int64  `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
				MaximumRecordAgeInSeconds      *int64  `tfsdk:"maximum_record_age_in_seconds" json:"maximumRecordAgeInSeconds,omitempty"`
				MaximumRetryAttempts           *int64  `tfsdk:"maximum_retry_attempts" json:"maximumRetryAttempts,omitempty"`
				OnPartialBatchItemFailure      *string `tfsdk:"on_partial_batch_item_failure" json:"onPartialBatchItemFailure,omitempty"`
				ParallelizationFactor          *int64  `tfsdk:"parallelization_factor" json:"parallelizationFactor,omitempty"`
				StartingPosition               *string `tfsdk:"starting_position" json:"startingPosition,omitempty"`
				StartingPositionTimestamp      *string `tfsdk:"starting_position_timestamp" json:"startingPositionTimestamp,omitempty"`
			} `tfsdk:"kinesis_stream_parameters" json:"kinesisStreamParameters,omitempty"`
			ManagedStreamingKafkaParameters *struct {
				BatchSize       *int64  `tfsdk:"batch_size" json:"batchSize,omitempty"`
				ConsumerGroupID *string `tfsdk:"consumer_group_id" json:"consumerGroupID,omitempty"`
				Credentials     *struct {
					ClientCertificateTLSAuth *string `tfsdk:"client_certificate_tls_auth" json:"clientCertificateTLSAuth,omitempty"`
					SaslSCRAM512Auth         *string `tfsdk:"sasl_scram512_auth" json:"saslSCRAM512Auth,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				MaximumBatchingWindowInSeconds *int64  `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
				StartingPosition               *string `tfsdk:"starting_position" json:"startingPosition,omitempty"`
				TopicName                      *string `tfsdk:"topic_name" json:"topicName,omitempty"`
			} `tfsdk:"managed_streaming_kafka_parameters" json:"managedStreamingKafkaParameters,omitempty"`
			RabbitMQBrokerParameters *struct {
				BatchSize   *int64 `tfsdk:"batch_size" json:"batchSize,omitempty"`
				Credentials *struct {
					BasicAuth *string `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				MaximumBatchingWindowInSeconds *int64  `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
				QueueName                      *string `tfsdk:"queue_name" json:"queueName,omitempty"`
				VirtualHost                    *string `tfsdk:"virtual_host" json:"virtualHost,omitempty"`
			} `tfsdk:"rabbit_mq_broker_parameters" json:"rabbitMQBrokerParameters,omitempty"`
			SelfManagedKafkaParameters *struct {
				AdditionalBootstrapServers *[]string `tfsdk:"additional_bootstrap_servers" json:"additionalBootstrapServers,omitempty"`
				BatchSize                  *int64    `tfsdk:"batch_size" json:"batchSize,omitempty"`
				ConsumerGroupID            *string   `tfsdk:"consumer_group_id" json:"consumerGroupID,omitempty"`
				Credentials                *struct {
					BasicAuth                *string `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
					ClientCertificateTLSAuth *string `tfsdk:"client_certificate_tls_auth" json:"clientCertificateTLSAuth,omitempty"`
					SaslSCRAM256Auth         *string `tfsdk:"sasl_scram256_auth" json:"saslSCRAM256Auth,omitempty"`
					SaslSCRAM512Auth         *string `tfsdk:"sasl_scram512_auth" json:"saslSCRAM512Auth,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				MaximumBatchingWindowInSeconds *int64  `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
				ServerRootCaCertificate        *string `tfsdk:"server_root_ca_certificate" json:"serverRootCaCertificate,omitempty"`
				StartingPosition               *string `tfsdk:"starting_position" json:"startingPosition,omitempty"`
				TopicName                      *string `tfsdk:"topic_name" json:"topicName,omitempty"`
				Vpc                            *struct {
					SecurityGroup *[]string `tfsdk:"security_group" json:"securityGroup,omitempty"`
					Subnets       *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
				} `tfsdk:"vpc" json:"vpc,omitempty"`
			} `tfsdk:"self_managed_kafka_parameters" json:"selfManagedKafkaParameters,omitempty"`
			SqsQueueParameters *struct {
				BatchSize                      *int64 `tfsdk:"batch_size" json:"batchSize,omitempty"`
				MaximumBatchingWindowInSeconds *int64 `tfsdk:"maximum_batching_window_in_seconds" json:"maximumBatchingWindowInSeconds,omitempty"`
			} `tfsdk:"sqs_queue_parameters" json:"sqsQueueParameters,omitempty"`
		} `tfsdk:"source_parameters" json:"sourceParameters,omitempty"`
		Tags             *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		Target           *string            `tfsdk:"target" json:"target,omitempty"`
		TargetParameters *struct {
			BatchJobParameters *struct {
				ArrayProperties *struct {
					Size *int64 `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"array_properties" json:"arrayProperties,omitempty"`
				ContainerOverrides *struct {
					Command     *[]string `tfsdk:"command" json:"command,omitempty"`
					Environment *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"environment" json:"environment,omitempty"`
					InstanceType         *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					ResourceRequirements *[]struct {
						Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"resource_requirements" json:"resourceRequirements,omitempty"`
				} `tfsdk:"container_overrides" json:"containerOverrides,omitempty"`
				DependsOn *[]struct {
					JobID *string `tfsdk:"job_id" json:"jobID,omitempty"`
					Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
				} `tfsdk:"depends_on" json:"dependsOn,omitempty"`
				JobDefinition *string            `tfsdk:"job_definition" json:"jobDefinition,omitempty"`
				JobName       *string            `tfsdk:"job_name" json:"jobName,omitempty"`
				Parameters    *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
				RetryStrategy *struct {
					Attempts *int64 `tfsdk:"attempts" json:"attempts,omitempty"`
				} `tfsdk:"retry_strategy" json:"retryStrategy,omitempty"`
			} `tfsdk:"batch_job_parameters" json:"batchJobParameters,omitempty"`
			CloudWatchLogsParameters *struct {
				LogStreamName *string `tfsdk:"log_stream_name" json:"logStreamName,omitempty"`
				Timestamp     *string `tfsdk:"timestamp" json:"timestamp,omitempty"`
			} `tfsdk:"cloud_watch_logs_parameters" json:"cloudWatchLogsParameters,omitempty"`
			EcsTaskParameters *struct {
				CapacityProviderStrategy *[]struct {
					Base             *int64  `tfsdk:"base" json:"base,omitempty"`
					CapacityProvider *string `tfsdk:"capacity_provider" json:"capacityProvider,omitempty"`
					Weight           *int64  `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"capacity_provider_strategy" json:"capacityProviderStrategy,omitempty"`
				EnableECSManagedTags *bool   `tfsdk:"enable_ecs_managed_tags" json:"enableECSManagedTags,omitempty"`
				EnableExecuteCommand *bool   `tfsdk:"enable_execute_command" json:"enableExecuteCommand,omitempty"`
				Group                *string `tfsdk:"group" json:"group,omitempty"`
				LaunchType           *string `tfsdk:"launch_type" json:"launchType,omitempty"`
				NetworkConfiguration *struct {
					AwsVPCConfiguration *struct {
						AssignPublicIP *string   `tfsdk:"assign_public_ip" json:"assignPublicIP,omitempty"`
						SecurityGroups *[]string `tfsdk:"security_groups" json:"securityGroups,omitempty"`
						Subnets        *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
					} `tfsdk:"aws_vpc_configuration" json:"awsVPCConfiguration,omitempty"`
				} `tfsdk:"network_configuration" json:"networkConfiguration,omitempty"`
				Overrides *struct {
					ContainerOverrides *[]struct {
						Command     *[]string `tfsdk:"command" json:"command,omitempty"`
						Cpu         *int64    `tfsdk:"cpu" json:"cpu,omitempty"`
						Environment *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"environment" json:"environment,omitempty"`
						EnvironmentFiles *[]struct {
							Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"environment_files" json:"environmentFiles,omitempty"`
						Memory               *int64  `tfsdk:"memory" json:"memory,omitempty"`
						MemoryReservation    *int64  `tfsdk:"memory_reservation" json:"memoryReservation,omitempty"`
						Name                 *string `tfsdk:"name" json:"name,omitempty"`
						ResourceRequirements *[]struct {
							Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"resource_requirements" json:"resourceRequirements,omitempty"`
					} `tfsdk:"container_overrides" json:"containerOverrides,omitempty"`
					Cpu              *string `tfsdk:"cpu" json:"cpu,omitempty"`
					EphemeralStorage *struct {
						SizeInGiB *int64 `tfsdk:"size_in_gi_b" json:"sizeInGiB,omitempty"`
					} `tfsdk:"ephemeral_storage" json:"ephemeralStorage,omitempty"`
					ExecutionRoleARN              *string `tfsdk:"execution_role_arn" json:"executionRoleARN,omitempty"`
					InferenceAcceleratorOverrides *[]struct {
						DeviceName *string `tfsdk:"device_name" json:"deviceName,omitempty"`
						DeviceType *string `tfsdk:"device_type" json:"deviceType,omitempty"`
					} `tfsdk:"inference_accelerator_overrides" json:"inferenceAcceleratorOverrides,omitempty"`
					Memory      *string `tfsdk:"memory" json:"memory,omitempty"`
					TaskRoleARN *string `tfsdk:"task_role_arn" json:"taskRoleARN,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				PlacementConstraints *[]struct {
					Expression *string `tfsdk:"expression" json:"expression,omitempty"`
					Type_      *string `tfsdk:"type_" json:"type_,omitempty"`
				} `tfsdk:"placement_constraints" json:"placementConstraints,omitempty"`
				PlacementStrategy *[]struct {
					Field *string `tfsdk:"field" json:"field,omitempty"`
					Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
				} `tfsdk:"placement_strategy" json:"placementStrategy,omitempty"`
				PlatformVersion *string `tfsdk:"platform_version" json:"platformVersion,omitempty"`
				PropagateTags   *string `tfsdk:"propagate_tags" json:"propagateTags,omitempty"`
				ReferenceID     *string `tfsdk:"reference_id" json:"referenceID,omitempty"`
				Tags            *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tags" json:"tags,omitempty"`
				TaskCount         *int64  `tfsdk:"task_count" json:"taskCount,omitempty"`
				TaskDefinitionARN *string `tfsdk:"task_definition_arn" json:"taskDefinitionARN,omitempty"`
			} `tfsdk:"ecs_task_parameters" json:"ecsTaskParameters,omitempty"`
			EventBridgeEventBusParameters *struct {
				DetailType *string   `tfsdk:"detail_type" json:"detailType,omitempty"`
				EndpointID *string   `tfsdk:"endpoint_id" json:"endpointID,omitempty"`
				Resources  *[]string `tfsdk:"resources" json:"resources,omitempty"`
				Source     *string   `tfsdk:"source" json:"source,omitempty"`
				Time       *string   `tfsdk:"time" json:"time,omitempty"`
			} `tfsdk:"event_bridge_event_bus_parameters" json:"eventBridgeEventBusParameters,omitempty"`
			HttpParameters *struct {
				HeaderParameters      *map[string]string `tfsdk:"header_parameters" json:"headerParameters,omitempty"`
				PathParameterValues   *[]string          `tfsdk:"path_parameter_values" json:"pathParameterValues,omitempty"`
				QueryStringParameters *map[string]string `tfsdk:"query_string_parameters" json:"queryStringParameters,omitempty"`
			} `tfsdk:"http_parameters" json:"httpParameters,omitempty"`
			InputTemplate           *string `tfsdk:"input_template" json:"inputTemplate,omitempty"`
			KinesisStreamParameters *struct {
				PartitionKey *string `tfsdk:"partition_key" json:"partitionKey,omitempty"`
			} `tfsdk:"kinesis_stream_parameters" json:"kinesisStreamParameters,omitempty"`
			LambdaFunctionParameters *struct {
				InvocationType *string `tfsdk:"invocation_type" json:"invocationType,omitempty"`
			} `tfsdk:"lambda_function_parameters" json:"lambdaFunctionParameters,omitempty"`
			RedshiftDataParameters *struct {
				Database         *string   `tfsdk:"database" json:"database,omitempty"`
				DbUser           *string   `tfsdk:"db_user" json:"dbUser,omitempty"`
				SecretManagerARN *string   `tfsdk:"secret_manager_arn" json:"secretManagerARN,omitempty"`
				Sqls             *[]string `tfsdk:"sqls" json:"sqls,omitempty"`
				StatementName    *string   `tfsdk:"statement_name" json:"statementName,omitempty"`
				WithEvent        *bool     `tfsdk:"with_event" json:"withEvent,omitempty"`
			} `tfsdk:"redshift_data_parameters" json:"redshiftDataParameters,omitempty"`
			SageMakerPipelineParameters *struct {
				PipelineParameterList *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"pipeline_parameter_list" json:"pipelineParameterList,omitempty"`
			} `tfsdk:"sage_maker_pipeline_parameters" json:"sageMakerPipelineParameters,omitempty"`
			SqsQueueParameters *struct {
				MessageDeduplicationID *string `tfsdk:"message_deduplication_id" json:"messageDeduplicationID,omitempty"`
				MessageGroupID         *string `tfsdk:"message_group_id" json:"messageGroupID,omitempty"`
			} `tfsdk:"sqs_queue_parameters" json:"sqsQueueParameters,omitempty"`
			StepFunctionStateMachineParameters *struct {
				InvocationType *string `tfsdk:"invocation_type" json:"invocationType,omitempty"`
			} `tfsdk:"step_function_state_machine_parameters" json:"stepFunctionStateMachineParameters,omitempty"`
		} `tfsdk:"target_parameters" json:"targetParameters,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PipesServicesK8SAwsPipeV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_pipes_services_k8s_aws_pipe_v1alpha1_manifest"
}

func (r *PipesServicesK8SAwsPipeV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Pipe is the Schema for the Pipes API",
		MarkdownDescription: "Pipe is the Schema for the Pipes API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "PipeSpec defines the desired state of Pipe.An object that represents a pipe. Amazon EventBridgePipes connect event sourcesto targets and reduces the need for specialized knowledge and integrationcode.",
				MarkdownDescription: "PipeSpec defines the desired state of Pipe.An object that represents a pipe. Amazon EventBridgePipes connect event sourcesto targets and reduces the need for specialized knowledge and integrationcode.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "A description of the pipe.",
						MarkdownDescription: "A description of the pipe.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"desired_state": schema.StringAttribute{
						Description:         "The state the pipe should be in.",
						MarkdownDescription: "The state the pipe should be in.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enrichment": schema.StringAttribute{
						Description:         "The ARN of the enrichment resource.",
						MarkdownDescription: "The ARN of the enrichment resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enrichment_parameters": schema.SingleNestedAttribute{
						Description:         "The parameters required to set up enrichment on your pipe.",
						MarkdownDescription: "The parameters required to set up enrichment on your pipe.",
						Attributes: map[string]schema.Attribute{
							"http_parameters": schema.SingleNestedAttribute{
								Description:         "These are custom parameter to be used when the target is an API Gateway RESTAPIs or EventBridge ApiDestinations. In the latter case, these are mergedwith any InvocationParameters specified on the Connection, with any valuesfrom the Connection taking precedence.",
								MarkdownDescription: "These are custom parameter to be used when the target is an API Gateway RESTAPIs or EventBridge ApiDestinations. In the latter case, these are mergedwith any InvocationParameters specified on the Connection, with any valuesfrom the Connection taking precedence.",
								Attributes: map[string]schema.Attribute{
									"header_parameters": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path_parameter_values": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_string_parameters": schema.MapAttribute{
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

							"input_template": schema.StringAttribute{
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

					"name": schema.StringAttribute{
						Description:         "The name of the pipe.",
						MarkdownDescription: "The name of the pipe.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"role_arn": schema.StringAttribute{
						Description:         "The ARN of the role that allows the pipe to send data to the target.",
						MarkdownDescription: "The ARN of the role that allows the pipe to send data to the target.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source": schema.StringAttribute{
						Description:         "The ARN of the source resource.",
						MarkdownDescription: "The ARN of the source resource.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source_parameters": schema.SingleNestedAttribute{
						Description:         "The parameters required to set up a source for your pipe.",
						MarkdownDescription: "The parameters required to set up a source for your pipe.",
						Attributes: map[string]schema.Attribute{
							"active_mq_broker_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using an Active MQ broker as a source.",
								MarkdownDescription: "The parameters for using an Active MQ broker as a source.",
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials": schema.SingleNestedAttribute{
										Description:         "The Secrets Manager secret that stores your broker credentials.",
										MarkdownDescription: "The Secrets Manager secret that stores your broker credentials.",
										Attributes: map[string]schema.Attribute{
											"basic_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum_batching_window_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"queue_name": schema.StringAttribute{
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

							"dynamo_db_stream_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a DynamoDB stream as a source.",
								MarkdownDescription: "The parameters for using a DynamoDB stream as a source.",
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dead_letter_config": schema.SingleNestedAttribute{
										Description:         "A DeadLetterConfig object that contains information about a dead-letter queueconfiguration.",
										MarkdownDescription: "A DeadLetterConfig object that contains information about a dead-letter queueconfiguration.",
										Attributes: map[string]schema.Attribute{
											"arn": schema.StringAttribute{
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

									"maximum_batching_window_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_record_age_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_retry_attempts": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"on_partial_batch_item_failure": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parallelization_factor": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"starting_position": schema.StringAttribute{
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

							"filter_criteria": schema.SingleNestedAttribute{
								Description:         "The collection of event patterns used to filter events. For more information,see Events and Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)in the Amazon EventBridge User Guide.",
								MarkdownDescription: "The collection of event patterns used to filter events. For more information,see Events and Event Patterns (https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-and-event-patterns.html)in the Amazon EventBridge User Guide.",
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

							"kinesis_stream_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a Kinesis stream as a source.",
								MarkdownDescription: "The parameters for using a Kinesis stream as a source.",
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dead_letter_config": schema.SingleNestedAttribute{
										Description:         "A DeadLetterConfig object that contains information about a dead-letter queueconfiguration.",
										MarkdownDescription: "A DeadLetterConfig object that contains information about a dead-letter queueconfiguration.",
										Attributes: map[string]schema.Attribute{
											"arn": schema.StringAttribute{
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

									"maximum_batching_window_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_record_age_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_retry_attempts": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"on_partial_batch_item_failure": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parallelization_factor": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"starting_position": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"starting_position_timestamp": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.DateTime64Validator(),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"managed_streaming_kafka_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using an MSK stream as a source.",
								MarkdownDescription: "The parameters for using an MSK stream as a source.",
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"consumer_group_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials": schema.SingleNestedAttribute{
										Description:         "The Secrets Manager secret that stores your stream credentials.",
										MarkdownDescription: "The Secrets Manager secret that stores your stream credentials.",
										Attributes: map[string]schema.Attribute{
											"client_certificate_tls_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sasl_scram512_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum_batching_window_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"starting_position": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"topic_name": schema.StringAttribute{
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

							"rabbit_mq_broker_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a Rabbit MQ broker as a source.",
								MarkdownDescription: "The parameters for using a Rabbit MQ broker as a source.",
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials": schema.SingleNestedAttribute{
										Description:         "The Secrets Manager secret that stores your broker credentials.",
										MarkdownDescription: "The Secrets Manager secret that stores your broker credentials.",
										Attributes: map[string]schema.Attribute{
											"basic_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum_batching_window_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"queue_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"virtual_host": schema.StringAttribute{
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

							"self_managed_kafka_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a self-managed Apache Kafka stream as a source.",
								MarkdownDescription: "The parameters for using a self-managed Apache Kafka stream as a source.",
								Attributes: map[string]schema.Attribute{
									"additional_bootstrap_servers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"batch_size": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"consumer_group_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credentials": schema.SingleNestedAttribute{
										Description:         "The Secrets Manager secret that stores your stream credentials.",
										MarkdownDescription: "The Secrets Manager secret that stores your stream credentials.",
										Attributes: map[string]schema.Attribute{
											"basic_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"client_certificate_tls_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sasl_scram256_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sasl_scram512_auth": schema.StringAttribute{
												Description:         "// Optional SecretManager ARN which stores the database credentials",
												MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum_batching_window_in_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server_root_ca_certificate": schema.StringAttribute{
										Description:         "// Optional SecretManager ARN which stores the database credentials",
										MarkdownDescription: "// Optional SecretManager ARN which stores the database credentials",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"starting_position": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"topic_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vpc": schema.SingleNestedAttribute{
										Description:         "This structure specifies the VPC subnets and security groups for the stream,and whether a public IP address is to be used.",
										MarkdownDescription: "This structure specifies the VPC subnets and security groups for the stream,and whether a public IP address is to be used.",
										Attributes: map[string]schema.Attribute{
											"security_group": schema.ListAttribute{
												Description:         "List of SecurityGroupId.",
												MarkdownDescription: "List of SecurityGroupId.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"subnets": schema.ListAttribute{
												Description:         "List of SubnetId.",
												MarkdownDescription: "List of SubnetId.",
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

							"sqs_queue_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a Amazon SQS stream as a source.",
								MarkdownDescription: "The parameters for using a Amazon SQS stream as a source.",
								Attributes: map[string]schema.Attribute{
									"batch_size": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_batching_window_in_seconds": schema.Int64Attribute{
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

					"tags": schema.MapAttribute{
						Description:         "The list of key-value pairs to associate with the pipe.",
						MarkdownDescription: "The list of key-value pairs to associate with the pipe.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target": schema.StringAttribute{
						Description:         "The ARN of the target resource.",
						MarkdownDescription: "The ARN of the target resource.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"target_parameters": schema.SingleNestedAttribute{
						Description:         "The parameters required to set up a target for your pipe.",
						MarkdownDescription: "The parameters required to set up a target for your pipe.",
						Attributes: map[string]schema.Attribute{
							"batch_job_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using an Batch job as a target.",
								MarkdownDescription: "The parameters for using an Batch job as a target.",
								Attributes: map[string]schema.Attribute{
									"array_properties": schema.SingleNestedAttribute{
										Description:         "The array properties for the submitted job, such as the size of the array.The array size can be between 2 and 10,000. If you specify array propertiesfor a job, it becomes an array job. This parameter is used only if the targetis an Batch job.",
										MarkdownDescription: "The array properties for the submitted job, such as the size of the array.The array size can be between 2 and 10,000. If you specify array propertiesfor a job, it becomes an array job. This parameter is used only if the targetis an Batch job.",
										Attributes: map[string]schema.Attribute{
											"size": schema.Int64Attribute{
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

									"container_overrides": schema.SingleNestedAttribute{
										Description:         "The overrides that are sent to a container.",
										MarkdownDescription: "The overrides that are sent to a container.",
										Attributes: map[string]schema.Attribute{
											"command": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"environment": schema.ListNestedAttribute{
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

											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource_requirements": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"type_": schema.StringAttribute{
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

									"depends_on": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"job_id": schema.StringAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"job_definition": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"job_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parameters": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_strategy": schema.SingleNestedAttribute{
										Description:         "The retry strategy that's associated with a job. For more information, seeAutomated job retries (https://docs.aws.amazon.com/batch/latest/userguide/job_retries.html)in the Batch User Guide.",
										MarkdownDescription: "The retry strategy that's associated with a job. For more information, seeAutomated job retries (https://docs.aws.amazon.com/batch/latest/userguide/job_retries.html)in the Batch User Guide.",
										Attributes: map[string]schema.Attribute{
											"attempts": schema.Int64Attribute{
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

							"cloud_watch_logs_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using an CloudWatch Logs log stream as a target.",
								MarkdownDescription: "The parameters for using an CloudWatch Logs log stream as a target.",
								Attributes: map[string]schema.Attribute{
									"log_stream_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timestamp": schema.StringAttribute{
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

							"ecs_task_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using an Amazon ECS task as a target.",
								MarkdownDescription: "The parameters for using an Amazon ECS task as a target.",
								Attributes: map[string]schema.Attribute{
									"capacity_provider_strategy": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"base": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"capacity_provider": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"weight": schema.Int64Attribute{
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

									"enable_ecs_managed_tags": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_execute_command": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"group": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"launch_type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"network_configuration": schema.SingleNestedAttribute{
										Description:         "This structure specifies the network configuration for an Amazon ECS task.",
										MarkdownDescription: "This structure specifies the network configuration for an Amazon ECS task.",
										Attributes: map[string]schema.Attribute{
											"aws_vpc_configuration": schema.SingleNestedAttribute{
												Description:         "This structure specifies the VPC subnets and security groups for the task,and whether a public IP address is to be used. This structure is relevantonly for ECS tasks that use the awsvpc network mode.",
												MarkdownDescription: "This structure specifies the VPC subnets and security groups for the task,and whether a public IP address is to be used. This structure is relevantonly for ECS tasks that use the awsvpc network mode.",
												Attributes: map[string]schema.Attribute{
													"assign_public_ip": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups": schema.ListAttribute{
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

									"overrides": schema.SingleNestedAttribute{
										Description:         "The overrides that are associated with a task.",
										MarkdownDescription: "The overrides that are associated with a task.",
										Attributes: map[string]schema.Attribute{
											"container_overrides": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cpu": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"environment": schema.ListNestedAttribute{
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

														"environment_files": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"type_": schema.StringAttribute{
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

														"memory": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"memory_reservation": schema.Int64Attribute{
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

														"resource_requirements": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"type_": schema.StringAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"cpu": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ephemeral_storage": schema.SingleNestedAttribute{
												Description:         "The amount of ephemeral storage to allocate for the task. This parameteris used to expand the total amount of ephemeral storage available, beyondthe default amount, for tasks hosted on Fargate. For more information, seeFargate task storage (https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html)in the Amazon ECS User Guide for Fargate.This parameter is only supported for tasks hosted on Fargate using Linuxplatform version 1.4.0 or later. This parameter is not supported for Windowscontainers on Fargate.",
												MarkdownDescription: "The amount of ephemeral storage to allocate for the task. This parameteris used to expand the total amount of ephemeral storage available, beyondthe default amount, for tasks hosted on Fargate. For more information, seeFargate task storage (https://docs.aws.amazon.com/AmazonECS/latest/userguide/using_data_volumes.html)in the Amazon ECS User Guide for Fargate.This parameter is only supported for tasks hosted on Fargate using Linuxplatform version 1.4.0 or later. This parameter is not supported for Windowscontainers on Fargate.",
												Attributes: map[string]schema.Attribute{
													"size_in_gi_b": schema.Int64Attribute{
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

											"execution_role_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"inference_accelerator_overrides": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"device_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"device_type": schema.StringAttribute{
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

											"memory": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"task_role_arn": schema.StringAttribute{
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

									"placement_constraints": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"expression": schema.StringAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"placement_strategy": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"field": schema.StringAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"platform_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"propagate_tags": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"reference_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tags": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
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

									"task_count": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"task_definition_arn": schema.StringAttribute{
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

							"event_bridge_event_bus_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using an EventBridge event bus as a target.",
								MarkdownDescription: "The parameters for using an EventBridge event bus as a target.",
								Attributes: map[string]schema.Attribute{
									"detail_type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"endpoint_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"time": schema.StringAttribute{
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

							"http_parameters": schema.SingleNestedAttribute{
								Description:         "These are custom parameter to be used when the target is an API Gateway RESTAPIs or EventBridge ApiDestinations.",
								MarkdownDescription: "These are custom parameter to be used when the target is an API Gateway RESTAPIs or EventBridge ApiDestinations.",
								Attributes: map[string]schema.Attribute{
									"header_parameters": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path_parameter_values": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_string_parameters": schema.MapAttribute{
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

							"input_template": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kinesis_stream_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a Kinesis stream as a source.",
								MarkdownDescription: "The parameters for using a Kinesis stream as a source.",
								Attributes: map[string]schema.Attribute{
									"partition_key": schema.StringAttribute{
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

							"lambda_function_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a Lambda function as a target.",
								MarkdownDescription: "The parameters for using a Lambda function as a target.",
								Attributes: map[string]schema.Attribute{
									"invocation_type": schema.StringAttribute{
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

							"redshift_data_parameters": schema.SingleNestedAttribute{
								Description:         "These are custom parameters to be used when the target is a Amazon Redshiftcluster to invoke the Amazon Redshift Data API ExecuteStatement.",
								MarkdownDescription: "These are custom parameters to be used when the target is a Amazon Redshiftcluster to invoke the Amazon Redshift Data API ExecuteStatement.",
								Attributes: map[string]schema.Attribute{
									"database": schema.StringAttribute{
										Description:         "// Redshift Database",
										MarkdownDescription: "// Redshift Database",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"db_user": schema.StringAttribute{
										Description:         "// Database user name",
										MarkdownDescription: "// Database user name",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_manager_arn": schema.StringAttribute{
										Description:         "// For targets, can either specify an ARN or a jsonpath pointing to the ARN.",
										MarkdownDescription: "// For targets, can either specify an ARN or a jsonpath pointing to the ARN.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sqls": schema.ListAttribute{
										Description:         "// A list of SQLs.",
										MarkdownDescription: "// A list of SQLs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"statement_name": schema.StringAttribute{
										Description:         "// A name for Redshift DataAPI statement which can be used as filter of //ListStatement.",
										MarkdownDescription: "// A name for Redshift DataAPI statement which can be used as filter of //ListStatement.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"with_event": schema.BoolAttribute{
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

							"sage_maker_pipeline_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a SageMaker pipeline as a target.",
								MarkdownDescription: "The parameters for using a SageMaker pipeline as a target.",
								Attributes: map[string]schema.Attribute{
									"pipeline_parameter_list": schema.ListNestedAttribute{
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

							"sqs_queue_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a Amazon SQS stream as a source.",
								MarkdownDescription: "The parameters for using a Amazon SQS stream as a source.",
								Attributes: map[string]schema.Attribute{
									"message_deduplication_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"message_group_id": schema.StringAttribute{
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

							"step_function_state_machine_parameters": schema.SingleNestedAttribute{
								Description:         "The parameters for using a Step Functions state machine as a target.",
								MarkdownDescription: "The parameters for using a Step Functions state machine as a target.",
								Attributes: map[string]schema.Attribute{
									"invocation_type": schema.StringAttribute{
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *PipesServicesK8SAwsPipeV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_pipes_services_k8s_aws_pipe_v1alpha1_manifest")

	var model PipesServicesK8SAwsPipeV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("pipes.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Pipe")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
