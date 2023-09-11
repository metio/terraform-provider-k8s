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
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &LambdaServicesK8SAwsFunctionV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &LambdaServicesK8SAwsFunctionV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &LambdaServicesK8SAwsFunctionV1Alpha1Resource{}
)

func NewLambdaServicesK8SAwsFunctionV1Alpha1Resource() resource.Resource {
	return &LambdaServicesK8SAwsFunctionV1Alpha1Resource{}
}

type LambdaServicesK8SAwsFunctionV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type LambdaServicesK8SAwsFunctionV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

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

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_lambda_services_k8s_aws_function_v1alpha1"
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
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
						Optional:            true,
						Computed:            false,
					},

					"code": schema.SingleNestedAttribute{
						Description:         "The code for the function.",
						MarkdownDescription: "The code for the function.",
						Attributes: map[string]schema.Attribute{
							"image_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_bucket": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
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

							"s3_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_object_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zip_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"code_signing_config_arn": schema.StringAttribute{
						Description:         "To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.",
						MarkdownDescription: "To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dead_letter_config": schema.SingleNestedAttribute{
						Description:         "A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see Dead-letter queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).",
						MarkdownDescription: "A dead-letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see Dead-letter queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#invocation-dlq).",
						Attributes: map[string]schema.Attribute{
							"target_arn": schema.StringAttribute{
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

					"description": schema.StringAttribute{
						Description:         "A description of the function.",
						MarkdownDescription: "A description of the function.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ephemeral_storage": schema.SingleNestedAttribute{
						Description:         "The size of the function's /tmp directory in MB. The default value is 512, but can be any whole number between 512 and 10,240 MB.",
						MarkdownDescription: "The size of the function's /tmp directory in MB. The default value is 512, but can be any whole number between 512 and 10,240 MB.",
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

					"file_system_configs": schema.ListNestedAttribute{
						Description:         "Connection settings for an Amazon EFS file system.",
						MarkdownDescription: "Connection settings for an Amazon EFS file system.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"local_mount_path": schema.StringAttribute{
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
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"on_success": schema.SingleNestedAttribute{
										Description:         "A destination for events that were processed successfully.",
										MarkdownDescription: "A destination for events that were processed successfully.",
										Attributes: map[string]schema.Attribute{
											"destination": schema.StringAttribute{
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

							"function_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maximum_event_age_in_seconds": schema.Int64Attribute{
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

							"qualifier": schema.StringAttribute{
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

					"handler": schema.StringAttribute{
						Description:         "The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see Lambda programming model (https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).",
						MarkdownDescription: "The name of the method within your code that Lambda calls to run your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see Lambda programming model (https://docs.aws.amazon.com/lambda/latest/dg/foundation-progmodel.html).",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
								Optional:            true,
								Computed:            false,
							},

							"entry_point": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"working_directory": schema.StringAttribute{
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

					"kms_key_arn": schema.StringAttribute{
						Description:         "The ARN of the Key Management Service (KMS) key that's used to encrypt your function's environment variables. If it's not provided, Lambda uses a default service key.",
						MarkdownDescription: "The ARN of the Key Management Service (KMS) key that's used to encrypt your function's environment variables. If it's not provided, Lambda uses a default service key.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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

					"layers": schema.ListAttribute{
						Description:         "A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.",
						MarkdownDescription: "A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"memory_size": schema.Int64Attribute{
						Description:         "The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.",
						MarkdownDescription: "The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-function-common.html#configuration-memory-console) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the Lambda function.  Name formats  * Function name – my-function.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN – 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						MarkdownDescription: "The name of the Lambda function.  Name formats  * Function name – my-function.  * Function ARN – arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN – 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"package_type": schema.StringAttribute{
						Description:         "The type of deployment package. Set to Image for container image and set to Zip for .zip file archive.",
						MarkdownDescription: "The type of deployment package. Set to Image for container image and set to Zip for .zip file archive.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"publish": schema.BoolAttribute{
						Description:         "Set to true to publish the first version of the function during creation.",
						MarkdownDescription: "Set to true to publish the first version of the function during creation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"reserved_concurrent_executions": schema.Int64Attribute{
						Description:         "The number of simultaneous executions to reserve for the function.",
						MarkdownDescription: "The number of simultaneous executions to reserve for the function.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"role": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the function's execution role.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the function's execution role.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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

					"runtime": schema.StringAttribute{
						Description:         "The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.",
						MarkdownDescription: "The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snap_start": schema.SingleNestedAttribute{
						Description:         "The function's SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.",
						MarkdownDescription: "The function's SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html) setting.",
						Attributes: map[string]schema.Attribute{
							"apply_on": schema.StringAttribute{
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

					"tags": schema.MapAttribute{
						Description:         "A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.",
						MarkdownDescription: "A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"timeout": schema.Int64Attribute{
						Description:         "The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).",
						MarkdownDescription: "The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For more information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tracing_config": schema.SingleNestedAttribute{
						Description:         "Set Mode to Active to sample and trace a subset of incoming requests with X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).",
						MarkdownDescription: "Set Mode to Active to sample and trace a subset of incoming requests with X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
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

					"vpc_config": schema.SingleNestedAttribute{
						Description:         "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can access resources and the internet only through that VPC. For more information, see Configuring a Lambda function to access resources in a VPC (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).",
						MarkdownDescription: "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can access resources and the internet only through that VPC. For more information, see Configuring a Lambda function to access resources in a VPC (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).",
						Attributes: map[string]schema.Attribute{
							"security_group_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
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

							"subnet_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_lambda_services_k8s_aws_function_v1alpha1")

	var model LambdaServicesK8SAwsFunctionV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("lambda.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Function")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lambda.services.k8s.aws", Version: "v1alpha1", Resource: "functions"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LambdaServicesK8SAwsFunctionV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lambda_services_k8s_aws_function_v1alpha1")

	var data LambdaServicesK8SAwsFunctionV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lambda.services.k8s.aws", Version: "v1alpha1", Resource: "functions"}).
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

	var readResponse LambdaServicesK8SAwsFunctionV1Alpha1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_lambda_services_k8s_aws_function_v1alpha1")

	var model LambdaServicesK8SAwsFunctionV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("lambda.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Function")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lambda.services.k8s.aws", Version: "v1alpha1", Resource: "functions"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LambdaServicesK8SAwsFunctionV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_lambda_services_k8s_aws_function_v1alpha1")

	var data LambdaServicesK8SAwsFunctionV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "lambda.services.k8s.aws", Version: "v1alpha1", Resource: "functions"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
