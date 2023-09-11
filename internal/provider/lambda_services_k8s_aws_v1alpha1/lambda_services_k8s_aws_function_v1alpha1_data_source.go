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
	_ datasource.DataSource              = &LambdaServicesK8SAwsFunctionV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &LambdaServicesK8SAwsFunctionV1Alpha1DataSource{}
)

func NewLambdaServicesK8SAwsFunctionV1Alpha1DataSource() datasource.DataSource {
	return &LambdaServicesK8SAwsFunctionV1Alpha1DataSource{}
}

type LambdaServicesK8SAwsFunctionV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type LambdaServicesK8SAwsFunctionV1Alpha1DataSourceData struct {
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
		Architectures *[]string `tfsdk:"architectures" json:"architectures,omitempty"`
		Code          *struct {
			ImageURI    *string `tfsdk:"image_uri" json:"imageURI,omitempty"`
			S3Bucket    *string `tfsdk:"s3_bucket" json:"s3Bucket,omitempty"`
			S3BucketRef *struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"s3_bucket_ref" json:"s3BucketRef,omitempty"`
			S3Key           *string `tfsdk:"s3_key" json:"s3Key,omitempty"`
			S3ObjectVersion *string `tfsdk:"s3_object_version" json:"s3ObjectVersion,omitempty"`
			ZipFile         *string `tfsdk:"zip_file" json:"zipFile,omitempty"`
		} `tfsdk:"code" json:"code,omitempty"`
		CodeSigningConfigARN *string `tfsdk:"code_signing_config_arn" json:"codeSigningConfigARN,omitempty"`
		DeadLetterConfig     *struct {
			TargetARN *string `tfsdk:"target_arn" json:"targetARN,omitempty"`
		} `tfsdk:"dead_letter_config" json:"deadLetterConfig,omitempty"`
		Description *string `tfsdk:"description" json:"description,omitempty"`
		Environment *struct {
			Variables *map[string]string `tfsdk:"variables" json:"variables,omitempty"`
		} `tfsdk:"environment" json:"environment,omitempty"`
		EphemeralStorage *struct {
			Size *int64 `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"ephemeral_storage" json:"ephemeralStorage,omitempty"`
		FileSystemConfigs *[]struct {
			Arn            *string `tfsdk:"arn" json:"arn,omitempty"`
			LocalMountPath *string `tfsdk:"local_mount_path" json:"localMountPath,omitempty"`
		} `tfsdk:"file_system_configs" json:"fileSystemConfigs,omitempty"`
		FunctionEventInvokeConfig *struct {
			DestinationConfig *struct {
				OnFailure *struct {
					Destination *string `tfsdk:"destination" json:"destination,omitempty"`
				} `tfsdk:"on_failure" json:"onFailure,omitempty"`
				OnSuccess *struct {
					Destination *string `tfsdk:"destination" json:"destination,omitempty"`
				} `tfsdk:"on_success" json:"onSuccess,omitempty"`
			} `tfsdk:"destination_config" json:"destinationConfig,omitempty"`
			FunctionName             *string `tfsdk:"function_name" json:"functionName,omitempty"`
			MaximumEventAgeInSeconds *int64  `tfsdk:"maximum_event_age_in_seconds" json:"maximumEventAgeInSeconds,omitempty"`
			MaximumRetryAttempts     *int64  `tfsdk:"maximum_retry_attempts" json:"maximumRetryAttempts,omitempty"`
			Qualifier                *string `tfsdk:"qualifier" json:"qualifier,omitempty"`
		} `tfsdk:"function_event_invoke_config" json:"functionEventInvokeConfig,omitempty"`
		Handler     *string `tfsdk:"handler" json:"handler,omitempty"`
		ImageConfig *struct {
			Command          *[]string `tfsdk:"command" json:"command,omitempty"`
			EntryPoint       *[]string `tfsdk:"entry_point" json:"entryPoint,omitempty"`
			WorkingDirectory *string   `tfsdk:"working_directory" json:"workingDirectory,omitempty"`
		} `tfsdk:"image_config" json:"imageConfig,omitempty"`
		KmsKeyARN *string `tfsdk:"kms_key_arn" json:"kmsKeyARN,omitempty"`
		KmsKeyRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"kms_key_ref" json:"kmsKeyRef,omitempty"`
		Layers                       *[]string `tfsdk:"layers" json:"layers,omitempty"`
		MemorySize                   *int64    `tfsdk:"memory_size" json:"memorySize,omitempty"`
		Name                         *string   `tfsdk:"name" json:"name,omitempty"`
		PackageType                  *string   `tfsdk:"package_type" json:"packageType,omitempty"`
		Publish                      *bool     `tfsdk:"publish" json:"publish,omitempty"`
		ReservedConcurrentExecutions *int64    `tfsdk:"reserved_concurrent_executions" json:"reservedConcurrentExecutions,omitempty"`
		Role                         *string   `tfsdk:"role" json:"role,omitempty"`
		RoleRef                      *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"role_ref" json:"roleRef,omitempty"`
		Runtime   *string `tfsdk:"runtime" json:"runtime,omitempty"`
		SnapStart *struct {
			ApplyOn *string `tfsdk:"apply_on" json:"applyOn,omitempty"`
		} `tfsdk:"snap_start" json:"snapStart,omitempty"`
		Tags          *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		Timeout       *int64             `tfsdk:"timeout" json:"timeout,omitempty"`
		TracingConfig *struct {
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"tracing_config" json:"tracingConfig,omitempty"`
		VpcConfig *struct {
			SecurityGroupIDs  *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
			SecurityGroupRefs *[]struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"security_group_refs" json:"securityGroupRefs,omitempty"`
			SubnetIDs  *[]string `tfsdk:"subnet_i_ds" json:"subnetIDs,omitempty"`
			SubnetRefs *[]struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"subnet_refs" json:"subnetRefs,omitempty"`
		} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lambda_services_k8s_aws_function_v1alpha1"
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Function is the Schema for the Functions API",
		MarkdownDescription: "Function is the Schema for the Functions API",
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
				Description:         "FunctionSpec defines the desired state of Function.",
				MarkdownDescription: "FunctionSpec defines the desired state of Function.",
				Attributes: map[string]schema.Attribute{
					"architectures": schema.ListAttribute{
						Description:         "The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is x86_64.",
						MarkdownDescription: "The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is x86_64.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"code": schema.SingleNestedAttribute{
						Description:         "The code for the function.",
						MarkdownDescription: "The code for the function.",
						Attributes: map[string]schema.Attribute{
							"image_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_bucket": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_bucket_ref": schema.SingleNestedAttribute{
								Description:         "Reference field for S3Bucket",
								MarkdownDescription: "Reference field for S3Bucket",
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

							"s3_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"s3_object_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"zip_file": schema.StringAttribute{
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

					"code_signing_config_arn": schema.StringAttribute{
						Description:         "To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.",
						MarkdownDescription: "To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"dead_letter_config": schema.SingleNestedAttribute{
						Description:         "A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see Dead-letter queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).",
						MarkdownDescription: "A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see Dead-letter queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).",
						Attributes: map[string]schema.Attribute{
							"target_arn": schema.StringAttribute{
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

					"description": schema.StringAttribute{
						Description:         "A description of the function.",
						MarkdownDescription: "A description of the function.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"environment": schema.SingleNestedAttribute{
						Description:         "Environment variables that are accessible from function code during execution.",
						MarkdownDescription: "Environment variables that are accessible from function code during execution.",
						Attributes: map[string]schema.Attribute{
							"variables": schema.MapAttribute{
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

					"ephemeral_storage": schema.SingleNestedAttribute{
						Description:         "The size of the function's /tmp directory in MB. The default value is 512, but can be any whole number between 512 and 10,240 MB.",
						MarkdownDescription: "The size of the function's /tmp directory in MB. The default value is 512, but can be any whole number between 512 and 10,240 MB.",
						Attributes: map[string]schema.Attribute{
							"size": schema.Int64Attribute{
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

					"file_system_configs": schema.ListNestedAttribute{
						Description:         "Connection settings for an Amazon EFS file system.",
						MarkdownDescription: "Connection settings for an Amazon EFS file system.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"local_mount_path": schema.StringAttribute{
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

					"function_event_invoke_config": schema.SingleNestedAttribute{
						Description:         "Configures options for asynchronous invocation on a function.  - DestinationConfig A destination for events after they have been sent to a function for processing.  Types of Destinations: Function - The Amazon Resource Name (ARN) of a Lambda function. Queue - The ARN of a standard SQS queue. Topic - The ARN of a standard SNS topic. Event Bus - The ARN of an Amazon EventBridge event bus.  - MaximumEventAgeInSeconds The maximum age of a request that Lambda sends to a function for processing.  - MaximumRetryAttempts The maximum number of times to retry when the function returns an error.",
						MarkdownDescription: "Configures options for asynchronous invocation on a function.  - DestinationConfig A destination for events after they have been sent to a function for processing.  Types of Destinations: Function - The Amazon Resource Name (ARN) of a Lambda function. Queue - The ARN of a standard SQS queue. Topic - The ARN of a standard SNS topic. Event Bus - The ARN of an Amazon EventBridge event bus.  - MaximumEventAgeInSeconds The maximum age of a request that Lambda sends to a function for processing.  - MaximumRetryAttempts The maximum number of times to retry when the function returns an error.",
						Attributes: map[string]schema.Attribute{
							"destination_config": schema.SingleNestedAttribute{
								Description:         "A configuration object that specifies the destination of an event after Lambda processes it.",
								MarkdownDescription: "A configuration object that specifies the destination of an event after Lambda processes it.",
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

							"function_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"maximum_event_age_in_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"maximum_retry_attempts": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"qualifier": schema.StringAttribute{
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

					"handler": schema.StringAttribute{
						Description:         "The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see Lambda programming model (https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).",
						MarkdownDescription: "The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see Lambda programming model (https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_config": schema.SingleNestedAttribute{
						Description:         "Container image configuration values (https://docs.aws.amazon.com/lambda/latest/dg/configuration-images.html#configuration-images-settings) that override the values in the container image Dockerfile.",
						MarkdownDescription: "Container image configuration values (https://docs.aws.amazon.com/lambda/latest/dg/configuration-images.html#configuration-images-settings) that override the values in the container image Dockerfile.",
						Attributes: map[string]schema.Attribute{
							"command": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"entry_point": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"working_directory": schema.StringAttribute{
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

					"kms_key_arn": schema.StringAttribute{
						Description:         "The ARN of the Key Management Service (KMS) key that's used to encrypt your function's environment variables. If it's not provided, Lambda uses a default service key.",
						MarkdownDescription: "The ARN of the Key Management Service (KMS) key that's used to encrypt your function's environment variables. If it's not provided, Lambda uses a default service key.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kms_key_ref": schema.SingleNestedAttribute{
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

					"layers": schema.ListAttribute{
						Description:         "A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.",
						MarkdownDescription: "A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"memory_size": schema.Int64Attribute{
						Description:         "The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.",
						MarkdownDescription: "The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the Lambda function.  Name formats  * Function name – my-function.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN – 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						MarkdownDescription: "The name of the Lambda function.  Name formats  * Function name – my-function.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN – 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"package_type": schema.StringAttribute{
						Description:         "The type of deployment package. Set to Image for container image and set to Zip for .zip file archive.",
						MarkdownDescription: "The type of deployment package. Set to Image for container image and set to Zip for .zip file archive.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"publish": schema.BoolAttribute{
						Description:         "Set to true to publish the first version of the function during creation.",
						MarkdownDescription: "Set to true to publish the first version of the function during creation.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"reserved_concurrent_executions": schema.Int64Attribute{
						Description:         "The number of simultaneous executions to reserve for the function.",
						MarkdownDescription: "The number of simultaneous executions to reserve for the function.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"role": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the function's execution role.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the function's execution role.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"role_ref": schema.SingleNestedAttribute{
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

					"runtime": schema.StringAttribute{
						Description:         "The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.",
						MarkdownDescription: "The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"snap_start": schema.SingleNestedAttribute{
						Description:         "The function's SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.",
						MarkdownDescription: "The function's SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.",
						Attributes: map[string]schema.Attribute{
							"apply_on": schema.StringAttribute{
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

					"tags": schema.MapAttribute{
						Description:         "A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.",
						MarkdownDescription: "A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"timeout": schema.Int64Attribute{
						Description:         "The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).",
						MarkdownDescription: "The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tracing_config": schema.SingleNestedAttribute{
						Description:         "Set Mode to Active to sample and trace a subset of incoming requests with X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).",
						MarkdownDescription: "Set Mode to Active to sample and trace a subset of incoming requests with X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
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
						Description:         "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can access resources and the internet only through that VPC. For more information, see Configuring a Lambda function to access resources in a VPC (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).",
						MarkdownDescription: "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can access resources and the internet only through that VPC. For more information, see Configuring a Lambda function to access resources in a VPC (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).",
						Attributes: map[string]schema.Attribute{
							"security_group_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"security_group_refs": schema.ListNestedAttribute{
								Description:         "Reference field for SecurityGroupIDs",
								MarkdownDescription: "Reference field for SecurityGroupIDs",
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

							"subnet_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"subnet_refs": schema.ListNestedAttribute{
								Description:         "Reference field for SubnetIDs",
								MarkdownDescription: "Reference field for SubnetIDs",
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

func (r *LambdaServicesK8SAwsFunctionV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *LambdaServicesK8SAwsFunctionV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_lambda_services_k8s_aws_function_v1alpha1")

	var data LambdaServicesK8SAwsFunctionV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lambda.services.k8s.aws", Version: "v1alpha1", Resource: "functions"}).
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

	var readResponse LambdaServicesK8SAwsFunctionV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("lambda.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("Function")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
