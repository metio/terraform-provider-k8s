/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SagemakerServicesK8SAwsEndpointConfigV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsEndpointConfigV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsEndpointConfigV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsEndpointConfigV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsEndpointConfigV1Alpha1ManifestData struct {
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
		AsyncInferenceConfig *struct {
			ClientConfig *struct {
				MaxConcurrentInvocationsPerInstance *int64 `tfsdk:"max_concurrent_invocations_per_instance" json:"maxConcurrentInvocationsPerInstance,omitempty"`
			} `tfsdk:"client_config" json:"clientConfig,omitempty"`
			OutputConfig *struct {
				KmsKeyID           *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
				NotificationConfig *struct {
					ErrorTopic   *string `tfsdk:"error_topic" json:"errorTopic,omitempty"`
					SuccessTopic *string `tfsdk:"success_topic" json:"successTopic,omitempty"`
				} `tfsdk:"notification_config" json:"notificationConfig,omitempty"`
				S3OutputPath *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
			} `tfsdk:"output_config" json:"outputConfig,omitempty"`
		} `tfsdk:"async_inference_config" json:"asyncInferenceConfig,omitempty"`
		DataCaptureConfig *struct {
			CaptureContentTypeHeader *struct {
				CsvContentTypes  *[]string `tfsdk:"csv_content_types" json:"csvContentTypes,omitempty"`
				JsonContentTypes *[]string `tfsdk:"json_content_types" json:"jsonContentTypes,omitempty"`
			} `tfsdk:"capture_content_type_header" json:"captureContentTypeHeader,omitempty"`
			CaptureOptions *[]struct {
				CaptureMode *string `tfsdk:"capture_mode" json:"captureMode,omitempty"`
			} `tfsdk:"capture_options" json:"captureOptions,omitempty"`
			DestinationS3URI          *string `tfsdk:"destination_s3_uri" json:"destinationS3URI,omitempty"`
			EnableCapture             *bool   `tfsdk:"enable_capture" json:"enableCapture,omitempty"`
			InitialSamplingPercentage *int64  `tfsdk:"initial_sampling_percentage" json:"initialSamplingPercentage,omitempty"`
			KmsKeyID                  *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		} `tfsdk:"data_capture_config" json:"dataCaptureConfig,omitempty"`
		EndpointConfigName *string `tfsdk:"endpoint_config_name" json:"endpointConfigName,omitempty"`
		KmsKeyID           *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		ProductionVariants *[]struct {
			AcceleratorType                             *string `tfsdk:"accelerator_type" json:"acceleratorType,omitempty"`
			ContainerStartupHealthCheckTimeoutInSeconds *int64  `tfsdk:"container_startup_health_check_timeout_in_seconds" json:"containerStartupHealthCheckTimeoutInSeconds,omitempty"`
			CoreDumpConfig                              *struct {
				DestinationS3URI *string `tfsdk:"destination_s3_uri" json:"destinationS3URI,omitempty"`
				KmsKeyID         *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			} `tfsdk:"core_dump_config" json:"coreDumpConfig,omitempty"`
			EnableSSMAccess                   *bool    `tfsdk:"enable_ssm_access" json:"enableSSMAccess,omitempty"`
			InitialInstanceCount              *int64   `tfsdk:"initial_instance_count" json:"initialInstanceCount,omitempty"`
			InitialVariantWeight              *float64 `tfsdk:"initial_variant_weight" json:"initialVariantWeight,omitempty"`
			InstanceType                      *string  `tfsdk:"instance_type" json:"instanceType,omitempty"`
			ModelDataDownloadTimeoutInSeconds *int64   `tfsdk:"model_data_download_timeout_in_seconds" json:"modelDataDownloadTimeoutInSeconds,omitempty"`
			ModelName                         *string  `tfsdk:"model_name" json:"modelName,omitempty"`
			ServerlessConfig                  *struct {
				MaxConcurrency *int64 `tfsdk:"max_concurrency" json:"maxConcurrency,omitempty"`
				MemorySizeInMB *int64 `tfsdk:"memory_size_in_mb" json:"memorySizeInMB,omitempty"`
			} `tfsdk:"serverless_config" json:"serverlessConfig,omitempty"`
			VariantName    *string `tfsdk:"variant_name" json:"variantName,omitempty"`
			VolumeSizeInGB *int64  `tfsdk:"volume_size_in_gb" json:"volumeSizeInGB,omitempty"`
		} `tfsdk:"production_variants" json:"productionVariants,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_endpoint_config_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EndpointConfig is the Schema for the EndpointConfigs API",
		MarkdownDescription: "EndpointConfig is the Schema for the EndpointConfigs API",
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
				Description:         "EndpointConfigSpec defines the desired state of EndpointConfig.",
				MarkdownDescription: "EndpointConfigSpec defines the desired state of EndpointConfig.",
				Attributes: map[string]schema.Attribute{
					"async_inference_config": schema.SingleNestedAttribute{
						Description:         "Specifies configuration for how an endpoint performs asynchronous inference. This is a required field in order for your Endpoint to be invoked using InvokeEndpointAsync (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpointAsync.html).",
						MarkdownDescription: "Specifies configuration for how an endpoint performs asynchronous inference. This is a required field in order for your Endpoint to be invoked using InvokeEndpointAsync (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpointAsync.html).",
						Attributes: map[string]schema.Attribute{
							"client_config": schema.SingleNestedAttribute{
								Description:         "Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.",
								MarkdownDescription: "Configures the behavior of the client used by SageMaker to interact with the model container during asynchronous inference.",
								Attributes: map[string]schema.Attribute{
									"max_concurrent_invocations_per_instance": schema.Int64Attribute{
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

							"output_config": schema.SingleNestedAttribute{
								Description:         "Specifies the configuration for asynchronous inference invocation outputs.",
								MarkdownDescription: "Specifies the configuration for asynchronous inference invocation outputs.",
								Attributes: map[string]schema.Attribute{
									"kms_key_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"notification_config": schema.SingleNestedAttribute{
										Description:         "Specifies the configuration for notifications of inference results for asynchronous inference.",
										MarkdownDescription: "Specifies the configuration for notifications of inference results for asynchronous inference.",
										Attributes: map[string]schema.Attribute{
											"error_topic": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_topic": schema.StringAttribute{
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

									"s3_output_path": schema.StringAttribute{
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

					"data_capture_config": schema.SingleNestedAttribute{
						Description:         "Configuration to control how SageMaker captures inference data.",
						MarkdownDescription: "Configuration to control how SageMaker captures inference data.",
						Attributes: map[string]schema.Attribute{
							"capture_content_type_header": schema.SingleNestedAttribute{
								Description:         "Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.",
								MarkdownDescription: "Configuration specifying how to treat different headers. If no headers are specified SageMaker will by default base64 encode when capturing the data.",
								Attributes: map[string]schema.Attribute{
									"csv_content_types": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"json_content_types": schema.ListAttribute{
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

							"capture_options": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"capture_mode": schema.StringAttribute{
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

							"destination_s3_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_capture": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"initial_sampling_percentage": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kms_key_id": schema.StringAttribute{
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

					"endpoint_config_name": schema.StringAttribute{
						Description:         "The name of the endpoint configuration. You specify this name in a CreateEndpoint request.",
						MarkdownDescription: "The name of the endpoint configuration. You specify this name in a CreateEndpoint request.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.  The KmsKeyId can be any of the following formats:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  * Alias name: alias/ExampleAlias  * Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias  The KMS key policy must grant permission to the IAM role that you specify in your CreateEndpoint, UpdateEndpoint requests. For more information, refer to the Amazon Web Services Key Management Service section Using Key Policies in Amazon Web Services KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)  Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a KmsKeyId when using an instance type with local storage. If any of the models that you specify in the ProductionVariants parameter use nitro-based instances with local storage, do not specify a value for the KmsKeyId parameter. If you specify a value for KmsKeyId when using any nitro-based instances with local storage, the call to CreateEndpointConfig fails.  For a list of instance types that support local instance storage, see Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes).  For more information about local instance storage encryption, see SSD Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html).",
						MarkdownDescription: "The Amazon Resource Name (ARN) of a Amazon Web Services Key Management Service key that SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.  The KmsKeyId can be any of the following formats:  * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab  * Key ARN: arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab  * Alias name: alias/ExampleAlias  * Alias name ARN: arn:aws:kms:us-west-2:111122223333:alias/ExampleAlias  The KMS key policy must grant permission to the IAM role that you specify in your CreateEndpoint, UpdateEndpoint requests. For more information, refer to the Amazon Web Services Key Management Service section Using Key Policies in Amazon Web Services KMS (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)  Certain Nitro-based instances include local storage, dependent on the instance type. Local storage volumes are encrypted using a hardware module on the instance. You can't request a KmsKeyId when using an instance type with local storage. If any of the models that you specify in the ProductionVariants parameter use nitro-based instances with local storage, do not specify a value for the KmsKeyId parameter. If you specify a value for KmsKeyId when using any nitro-based instances with local storage, the call to CreateEndpointConfig fails.  For a list of instance types that support local instance storage, see Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#instance-store-volumes).  For more information about local instance storage encryption, see SSD Instance Store Volumes (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ssd-instance-store.html).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"production_variants": schema.ListNestedAttribute{
						Description:         "An array of ProductionVariant objects, one for each model that you want to host at this endpoint.",
						MarkdownDescription: "An array of ProductionVariant objects, one for each model that you want to host at this endpoint.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"accelerator_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"container_startup_health_check_timeout_in_seconds": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"core_dump_config": schema.SingleNestedAttribute{
									Description:         "Specifies configuration for a core dump from the model container when the process crashes.",
									MarkdownDescription: "Specifies configuration for a core dump from the model container when the process crashes.",
									Attributes: map[string]schema.Attribute{
										"destination_s3_uri": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kms_key_id": schema.StringAttribute{
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

								"enable_ssm_access": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"initial_instance_count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"initial_variant_weight": schema.Float64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"instance_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"model_data_download_timeout_in_seconds": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"model_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"serverless_config": schema.SingleNestedAttribute{
									Description:         "Specifies the serverless configuration for an endpoint variant.",
									MarkdownDescription: "Specifies the serverless configuration for an endpoint variant.",
									Attributes: map[string]schema.Attribute{
										"max_concurrency": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"memory_size_in_mb": schema.Int64Attribute{
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

								"variant_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"volume_size_in_gb": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SagemakerServicesK8SAwsEndpointConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_endpoint_config_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsEndpointConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("EndpointConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
