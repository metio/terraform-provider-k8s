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

type LambdaServicesK8SAwsFunctionV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*LambdaServicesK8SAwsFunctionV1Alpha1Resource)(nil)
)

type LambdaServicesK8SAwsFunctionV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LambdaServicesK8SAwsFunctionV1Alpha1GoModel struct {
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
		Architectures *[]string `tfsdk:"architectures" yaml:"architectures,omitempty"`

		Code *struct {
			ImageURI *string `tfsdk:"image_uri" yaml:"imageURI,omitempty"`

			S3Bucket *string `tfsdk:"s3_bucket" yaml:"s3Bucket,omitempty"`

			S3Key *string `tfsdk:"s3_key" yaml:"s3Key,omitempty"`

			S3ObjectVersion *string `tfsdk:"s3_object_version" yaml:"s3ObjectVersion,omitempty"`

			ZipFile *string `tfsdk:"zip_file" yaml:"zipFile,omitempty"`
		} `tfsdk:"code" yaml:"code,omitempty"`

		CodeSigningConfigARN *string `tfsdk:"code_signing_config_arn" yaml:"codeSigningConfigARN,omitempty"`

		DeadLetterConfig *struct {
			TargetARN *string `tfsdk:"target_arn" yaml:"targetARN,omitempty"`
		} `tfsdk:"dead_letter_config" yaml:"deadLetterConfig,omitempty"`

		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		Environment *struct {
			Variables *map[string]string `tfsdk:"variables" yaml:"variables,omitempty"`
		} `tfsdk:"environment" yaml:"environment,omitempty"`

		EphemeralStorage *struct {
			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"ephemeral_storage" yaml:"ephemeralStorage,omitempty"`

		FileSystemConfigs *[]struct {
			Arn *string `tfsdk:"arn" yaml:"arn,omitempty"`

			LocalMountPath *string `tfsdk:"local_mount_path" yaml:"localMountPath,omitempty"`
		} `tfsdk:"file_system_configs" yaml:"fileSystemConfigs,omitempty"`

		Handler *string `tfsdk:"handler" yaml:"handler,omitempty"`

		ImageConfig *struct {
			Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

			EntryPoint *[]string `tfsdk:"entry_point" yaml:"entryPoint,omitempty"`

			WorkingDirectory *string `tfsdk:"working_directory" yaml:"workingDirectory,omitempty"`
		} `tfsdk:"image_config" yaml:"imageConfig,omitempty"`

		KmsKeyARN *string `tfsdk:"kms_key_arn" yaml:"kmsKeyARN,omitempty"`

		Layers *[]string `tfsdk:"layers" yaml:"layers,omitempty"`

		MemorySize *int64 `tfsdk:"memory_size" yaml:"memorySize,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		PackageType *string `tfsdk:"package_type" yaml:"packageType,omitempty"`

		Publish *bool `tfsdk:"publish" yaml:"publish,omitempty"`

		ReservedConcurrentExecutions *int64 `tfsdk:"reserved_concurrent_executions" yaml:"reservedConcurrentExecutions,omitempty"`

		Role *string `tfsdk:"role" yaml:"role,omitempty"`

		Runtime *string `tfsdk:"runtime" yaml:"runtime,omitempty"`

		Tags *map[string]string `tfsdk:"tags" yaml:"tags,omitempty"`

		Timeout *int64 `tfsdk:"timeout" yaml:"timeout,omitempty"`

		TracingConfig *struct {
			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
		} `tfsdk:"tracing_config" yaml:"tracingConfig,omitempty"`

		VpcConfig *struct {
			SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

			SubnetIDs *[]string `tfsdk:"subnet_i_ds" yaml:"subnetIDs,omitempty"`
		} `tfsdk:"vpc_config" yaml:"vpcConfig,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLambdaServicesK8SAwsFunctionV1Alpha1Resource() resource.Resource {
	return &LambdaServicesK8SAwsFunctionV1Alpha1Resource{}
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lambda_services_k8s_aws_function_v1alpha1"
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Function is the Schema for the Functions API",
		MarkdownDescription: "Function is the Schema for the Functions API",
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
				Description:         "FunctionSpec defines the desired state of Function.",
				MarkdownDescription: "FunctionSpec defines the desired state of Function.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"architectures": {
						Description:         "The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is x86_64.",
						MarkdownDescription: "The instruction set architecture that the function supports. Enter a string array with one of the valid values (arm64 or x86_64). The default value is x86_64.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"code": {
						Description:         "The code for the function.",
						MarkdownDescription: "The code for the function.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"image_uri": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_bucket": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_object_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"zip_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.Base64Validator(),
								},
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"code_signing_config_arn": {
						Description:         "To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.",
						MarkdownDescription: "To enable code signing for this function, specify the ARN of a code-signing configuration. A code-signing configuration includes a set of signing profiles, which define the trusted publishers for this function.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dead_letter_config": {
						Description:         "A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see Dead Letter Queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).",
						MarkdownDescription: "A dead letter queue configuration that specifies the queue or topic where Lambda sends asynchronous events when they fail processing. For more information, see Dead Letter Queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"target_arn": {
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

					"description": {
						Description:         "A description of the function.",
						MarkdownDescription: "A description of the function.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"environment": {
						Description:         "Environment variables that are accessible from function code during execution.",
						MarkdownDescription: "Environment variables that are accessible from function code during execution.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"variables": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ephemeral_storage": {
						Description:         "The size of the function’s /tmp directory in MB. The default value is 512, but can be any whole number between 512 and 10240 MB.",
						MarkdownDescription: "The size of the function’s /tmp directory in MB. The default value is 512, but can be any whole number between 512 and 10240 MB.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"size": {
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

					"file_system_configs": {
						Description:         "Connection settings for an Amazon EFS file system.",
						MarkdownDescription: "Connection settings for an Amazon EFS file system.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"local_mount_path": {
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

					"handler": {
						Description:         "The name of the method within your code that Lambda calls to execute your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see Programming Model (https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html).",
						MarkdownDescription: "The name of the method within your code that Lambda calls to execute your function. Handler is required if the deployment package is a .zip file archive. The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime. For more information, see Programming Model (https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_config": {
						Description:         "Container image configuration values (https://docs.aws.amazon.com/lambda/latest/dg/configuration-images.html#configuration-images-settings) that override the values in the container image Dockerfile.",
						MarkdownDescription: "Container image configuration values (https://docs.aws.amazon.com/lambda/latest/dg/configuration-images.html#configuration-images-settings) that override the values in the container image Dockerfile.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"command": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"entry_point": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"working_directory": {
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

					"kms_key_arn": {
						Description:         "The ARN of the Amazon Web Services Key Management Service (KMS) key that's used to encrypt your function's environment variables. If it's not provided, Lambda uses a default service key.",
						MarkdownDescription: "The ARN of the Amazon Web Services Key Management Service (KMS) key that's used to encrypt your function's environment variables. If it's not provided, Lambda uses a default service key.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"layers": {
						Description:         "A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.",
						MarkdownDescription: "A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add to the function's execution environment. Specify each layer by its ARN, including the version.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"memory_size": {
						Description:         "The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-memory.html) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.",
						MarkdownDescription: "The amount of memory available to the function (https://docs.aws.amazon.com/lambda/latest/dg/configuration-memory.html) at runtime. Increasing the function memory also increases its CPU allocation. The default value is 128 MB. The value can be any multiple of 1 MB.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": {
						Description:         "The name of the Lambda function.  Name formats  * Function name - my-function.  * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN - 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
						MarkdownDescription: "The name of the Lambda function.  Name formats  * Function name - my-function.  * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.  * Partial ARN - 123456789012:function:my-function.  The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"package_type": {
						Description:         "The type of deployment package. Set to Image for container image and set Zip for ZIP archive.",
						MarkdownDescription: "The type of deployment package. Set to Image for container image and set Zip for ZIP archive.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"publish": {
						Description:         "Set to true to publish the first version of the function during creation.",
						MarkdownDescription: "Set to true to publish the first version of the function during creation.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"reserved_concurrent_executions": {
						Description:         "The number of simultaneous executions to reserve for the function.",
						MarkdownDescription: "The number of simultaneous executions to reserve for the function.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"role": {
						Description:         "The Amazon Resource Name (ARN) of the function's execution role.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the function's execution role.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"runtime": {
						Description:         "The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.",
						MarkdownDescription: "The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html). Runtime is required if the deployment package is a .zip file archive.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.",
						MarkdownDescription: "A list of tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to apply to the function.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"timeout": {
						Description:         "The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For additional information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).",
						MarkdownDescription: "The amount of time (in seconds) that Lambda allows a function to run before stopping it. The default is 3 seconds. The maximum allowed value is 900 seconds. For additional information, see Lambda execution environment (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tracing_config": {
						Description:         "Set Mode to Active to sample and trace a subset of incoming requests with X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).",
						MarkdownDescription: "Set Mode to Active to sample and trace a subset of incoming requests with X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"mode": {
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

					"vpc_config": {
						Description:         "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can only access resources and the internet through that VPC. For more information, see VPC Settings (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).",
						MarkdownDescription: "For network connectivity to Amazon Web Services resources in a VPC, specify a list of security groups and subnets in the VPC. When you connect a function to a VPC, it can only access resources and the internet through that VPC. For more information, see VPC Settings (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"security_group_i_ds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subnet_i_ds": {
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
		},
	}, nil
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_lambda_services_k8s_aws_function_v1alpha1")

	var state LambdaServicesK8SAwsFunctionV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LambdaServicesK8SAwsFunctionV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("lambda.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Function")

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

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_lambda_services_k8s_aws_function_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_lambda_services_k8s_aws_function_v1alpha1")

	var state LambdaServicesK8SAwsFunctionV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LambdaServicesK8SAwsFunctionV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("lambda.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Function")

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

func (r *LambdaServicesK8SAwsFunctionV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_lambda_services_k8s_aws_function_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
