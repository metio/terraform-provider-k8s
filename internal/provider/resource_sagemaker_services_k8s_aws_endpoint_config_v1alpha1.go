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

type SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsEndpointConfigV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsEndpointConfigV1Alpha1GoModel struct {
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
		AsyncInferenceConfig *struct {
			ClientConfig *struct {
				MaxConcurrentInvocationsPerInstance *int64 `tfsdk:"max_concurrent_invocations_per_instance" yaml:"maxConcurrentInvocationsPerInstance,omitempty"`
			} `tfsdk:"client_config" yaml:"clientConfig,omitempty"`

			OutputConfig *struct {
				KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

				NotificationConfig *struct {
					ErrorTopic *string `tfsdk:"error_topic" yaml:"errorTopic,omitempty"`

					SuccessTopic *string `tfsdk:"success_topic" yaml:"successTopic,omitempty"`
				} `tfsdk:"notification_config" yaml:"notificationConfig,omitempty"`

				S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`
			} `tfsdk:"output_config" yaml:"outputConfig,omitempty"`
		} `tfsdk:"async_inference_config" yaml:"asyncInferenceConfig,omitempty"`

		DataCaptureConfig *struct {
			CaptureContentTypeHeader *struct {
				CsvContentTypes *[]string `tfsdk:"csv_content_types" yaml:"csvContentTypes,omitempty"`

				JsonContentTypes *[]string `tfsdk:"json_content_types" yaml:"jsonContentTypes,omitempty"`
			} `tfsdk:"capture_content_type_header" yaml:"captureContentTypeHeader,omitempty"`

			CaptureOptions *[]struct {
				CaptureMode *string `tfsdk:"capture_mode" yaml:"captureMode,omitempty"`
			} `tfsdk:"capture_options" yaml:"captureOptions,omitempty"`

			DestinationS3URI *string `tfsdk:"destination_s3_uri" yaml:"destinationS3URI,omitempty"`

			EnableCapture *bool `tfsdk:"enable_capture" yaml:"enableCapture,omitempty"`

			InitialSamplingPercentage *int64 `tfsdk:"initial_sampling_percentage" yaml:"initialSamplingPercentage,omitempty"`

			KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`
		} `tfsdk:"data_capture_config" yaml:"dataCaptureConfig,omitempty"`

		EndpointConfigName *string `tfsdk:"endpoint_config_name" yaml:"endpointConfigName,omitempty"`

		KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

		ProductionVariants *[]struct {
			AcceleratorType *string `tfsdk:"accelerator_type" yaml:"acceleratorType,omitempty"`

			ContainerStartupHealthCheckTimeoutInSeconds *int64 `tfsdk:"container_startup_health_check_timeout_in_seconds" yaml:"containerStartupHealthCheckTimeoutInSeconds,omitempty"`

			CoreDumpConfig *struct {
				DestinationS3URI *string `tfsdk:"destination_s3_uri" yaml:"destinationS3URI,omitempty"`

				KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`
			} `tfsdk:"core_dump_config" yaml:"coreDumpConfig,omitempty"`

			InitialInstanceCount *int64 `tfsdk:"initial_instance_count" yaml:"initialInstanceCount,omitempty"`

			InitialVariantWeight *float64 `tfsdk:"initial_variant_weight" yaml:"initialVariantWeight,omitempty"`

			InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

			ModelDataDownloadTimeoutInSeconds *int64 `tfsdk:"model_data_download_timeout_in_seconds" yaml:"modelDataDownloadTimeoutInSeconds,omitempty"`

			ModelName *string `tfsdk:"model_name" yaml:"modelName,omitempty"`

			VariantName *string `tfsdk:"variant_name" yaml:"variantName,omitempty"`

			VolumeSizeInGB *int64 `tfsdk:"volume_size_in_gb" yaml:"volumeSizeInGB,omitempty"`
		} `tfsdk:"production_variants" yaml:"productionVariants,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_endpoint_config_v1alpha1"
}

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "EndpointConfig is the Schema for the EndpointConfigs API",
		MarkdownDescription: "EndpointConfig is the Schema for the EndpointConfigs API",
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
				Description:         "EndpointConfigSpec defines the desired state of EndpointConfig.",
				MarkdownDescription: "EndpointConfigSpec defines the desired state of EndpointConfig.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"async_inference_config": {
						Description:         "Specifies configuration for how an endpoint performs asynchronous inference. This is a required field in order for your Endpoint to be invoked using InvokeEndpointAsync (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpointAsync.html).",
						MarkdownDescription: "Specifies configuration for how an endpoint performs asynchronous inference. This is a required field in order for your Endpoint to be invoked using InvokeEndpointAsync (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpointAsync.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"client_config": {
								Description:         "Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.",
								MarkdownDescription: "Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"max_concurrent_invocations_per_instance": {
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

							"output_config": {
								Description:         "Specifies the configuration for asynchronous inference invocation outputs.",
								MarkdownDescription: "Specifies the configuration for asynchronous inference invocation outputs.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"kms_key_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"notification_config": {
										Description:         "Specifies the configuration for notifications of inference results for asynchronous inference.",
										MarkdownDescription: "Specifies the configuration for notifications of inference results for asynchronous inference.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"error_topic": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"success_topic": {
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

									"s3_output_path": {
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

					"data_capture_config": {
						Description:         "Configuration to control how SageMaker captures inference data.",
						MarkdownDescription: "Configuration to control how SageMaker captures inference data.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"capture_content_type_header": {
								Description:         "Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.",
								MarkdownDescription: "Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"csv_content_types": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_content_types": {
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

							"capture_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"capture_mode": {
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

							"destination_s3_uri": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_capture": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_sampling_percentage": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kms_key_id": {
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

					"endpoint_config_name": {
						Description:         "The name of the endpoint configuration. You specify this name in a CreateEndpoint request.",
						MarkdownDescription: "The name of the endpoint configuration. You specify this name in a CreateEndpoint request.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"kms_key_id": {
						Description:         "The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.  The KmsKeyId can be any of the following formats:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  * Alias name: alias/ExampleAlias  * Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias  The KMS key policy must grant permission to the IAM role that you specify in your CreateEndpoint, UpdateEndpoint requests. For more information, refer to the Amazon Web Services Key Management Service section Using Key Policies in Amazon Web Services KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)  Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a KmsKeyId when using an instance type with local storage. If any of the models that you specify in the ProductionVariants parameter use nitro-based instances with local storage, do not specify a value for the KmsKeyId parameter. If you specify a value for KmsKeyId when using any nitro-based instances with local storage, the call to CreateEndpointConfig fails.  For a list of instance types that support local instance storage, see Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes).  For more information about local instance storage encryption, see SSD Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html).",
						MarkdownDescription: "The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.  The KmsKeyId can be any of the following formats:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  * Alias name: alias/ExampleAlias  * Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias  The KMS key policy must grant permission to the IAM role that you specify in your CreateEndpoint, UpdateEndpoint requests. For more information, refer to the Amazon Web Services Key Management Service section Using Key Policies in Amazon Web Services KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)  Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a KmsKeyId when using an instance type with local storage. If any of the models that you specify in the ProductionVariants parameter use nitro-based instances with local storage, do not specify a value for the KmsKeyId parameter. If you specify a value for KmsKeyId when using any nitro-based instances with local storage, the call to CreateEndpointConfig fails.  For a list of instance types that support local instance storage, see Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes).  For more information about local instance storage encryption, see SSD Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"production_variants": {
						Description:         "An list of ProductionVariant objects, one for each model that you want to host at this endpoint.",
						MarkdownDescription: "An list of ProductionVariant objects, one for each model that you want to host at this endpoint.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"accelerator_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"container_startup_health_check_timeout_in_seconds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"core_dump_config": {
								Description:         "Specifies configuration for a core dump from the model container when the process crashes.",
								MarkdownDescription: "Specifies configuration for a core dump from the model container when the process crashes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"destination_s3_uri": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kms_key_id": {
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

							"initial_instance_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_variant_weight": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.NumberType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"model_data_download_timeout_in_seconds": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"model_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"variant_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_size_in_gb": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": {
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
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
		},
	}, nil
}

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1")

	var state SagemakerServicesK8SAwsEndpointConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsEndpointConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("EndpointConfig")

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

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1")

	var state SagemakerServicesK8SAwsEndpointConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsEndpointConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("EndpointConfig")

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

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
