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

type SagemakerServicesK8SAwsModelPackageV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsModelPackageV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsModelPackageV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsModelPackageV1Alpha1GoModel struct {
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
		AdditionalInferenceSpecifications *[]struct {
			Containers *[]struct {
				ContainerHostname *string `tfsdk:"container_hostname" yaml:"containerHostname,omitempty"`

				Environment *map[string]string `tfsdk:"environment" yaml:"environment,omitempty"`

				Framework *string `tfsdk:"framework" yaml:"framework,omitempty"`

				FrameworkVersion *string `tfsdk:"framework_version" yaml:"frameworkVersion,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImageDigest *string `tfsdk:"image_digest" yaml:"imageDigest,omitempty"`

				ModelDataURL *string `tfsdk:"model_data_url" yaml:"modelDataURL,omitempty"`

				ModelInput *struct {
					DataInputConfig *string `tfsdk:"data_input_config" yaml:"dataInputConfig,omitempty"`
				} `tfsdk:"model_input" yaml:"modelInput,omitempty"`

				NearestModelName *string `tfsdk:"nearest_model_name" yaml:"nearestModelName,omitempty"`

				ProductID *string `tfsdk:"product_id" yaml:"productID,omitempty"`
			} `tfsdk:"containers" yaml:"containers,omitempty"`

			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			SupportedContentTypes *[]string `tfsdk:"supported_content_types" yaml:"supportedContentTypes,omitempty"`

			SupportedRealtimeInferenceInstanceTypes *[]string `tfsdk:"supported_realtime_inference_instance_types" yaml:"supportedRealtimeInferenceInstanceTypes,omitempty"`

			SupportedResponseMIMETypes *[]string `tfsdk:"supported_response_mime_types" yaml:"supportedResponseMIMETypes,omitempty"`

			SupportedTransformInstanceTypes *[]string `tfsdk:"supported_transform_instance_types" yaml:"supportedTransformInstanceTypes,omitempty"`
		} `tfsdk:"additional_inference_specifications" yaml:"additionalInferenceSpecifications,omitempty"`

		ApprovalDescription *string `tfsdk:"approval_description" yaml:"approvalDescription,omitempty"`

		CertifyForMarketplace *bool `tfsdk:"certify_for_marketplace" yaml:"certifyForMarketplace,omitempty"`

		ClientToken *string `tfsdk:"client_token" yaml:"clientToken,omitempty"`

		CustomerMetadataProperties *map[string]string `tfsdk:"customer_metadata_properties" yaml:"customerMetadataProperties,omitempty"`

		Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

		DriftCheckBaselines *struct {
			Bias *struct {
				ConfigFile *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"config_file" yaml:"configFile,omitempty"`

				PostTrainingConstraints *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"post_training_constraints" yaml:"postTrainingConstraints,omitempty"`

				PreTrainingConstraints *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"pre_training_constraints" yaml:"preTrainingConstraints,omitempty"`
			} `tfsdk:"bias" yaml:"bias,omitempty"`

			Explainability *struct {
				ConfigFile *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"config_file" yaml:"configFile,omitempty"`

				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"constraints" yaml:"constraints,omitempty"`
			} `tfsdk:"explainability" yaml:"explainability,omitempty"`

			ModelDataQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"constraints" yaml:"constraints,omitempty"`

				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"statistics" yaml:"statistics,omitempty"`
			} `tfsdk:"model_data_quality" yaml:"modelDataQuality,omitempty"`

			ModelQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"constraints" yaml:"constraints,omitempty"`

				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"statistics" yaml:"statistics,omitempty"`
			} `tfsdk:"model_quality" yaml:"modelQuality,omitempty"`
		} `tfsdk:"drift_check_baselines" yaml:"driftCheckBaselines,omitempty"`

		InferenceSpecification *struct {
			Containers *[]struct {
				ContainerHostname *string `tfsdk:"container_hostname" yaml:"containerHostname,omitempty"`

				Environment *map[string]string `tfsdk:"environment" yaml:"environment,omitempty"`

				Framework *string `tfsdk:"framework" yaml:"framework,omitempty"`

				FrameworkVersion *string `tfsdk:"framework_version" yaml:"frameworkVersion,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImageDigest *string `tfsdk:"image_digest" yaml:"imageDigest,omitempty"`

				ModelDataURL *string `tfsdk:"model_data_url" yaml:"modelDataURL,omitempty"`

				ModelInput *struct {
					DataInputConfig *string `tfsdk:"data_input_config" yaml:"dataInputConfig,omitempty"`
				} `tfsdk:"model_input" yaml:"modelInput,omitempty"`

				NearestModelName *string `tfsdk:"nearest_model_name" yaml:"nearestModelName,omitempty"`

				ProductID *string `tfsdk:"product_id" yaml:"productID,omitempty"`
			} `tfsdk:"containers" yaml:"containers,omitempty"`

			SupportedContentTypes *[]string `tfsdk:"supported_content_types" yaml:"supportedContentTypes,omitempty"`

			SupportedRealtimeInferenceInstanceTypes *[]string `tfsdk:"supported_realtime_inference_instance_types" yaml:"supportedRealtimeInferenceInstanceTypes,omitempty"`

			SupportedResponseMIMETypes *[]string `tfsdk:"supported_response_mime_types" yaml:"supportedResponseMIMETypes,omitempty"`

			SupportedTransformInstanceTypes *[]string `tfsdk:"supported_transform_instance_types" yaml:"supportedTransformInstanceTypes,omitempty"`
		} `tfsdk:"inference_specification" yaml:"inferenceSpecification,omitempty"`

		MetadataProperties *struct {
			CommitID *string `tfsdk:"commit_id" yaml:"commitID,omitempty"`

			GeneratedBy *string `tfsdk:"generated_by" yaml:"generatedBy,omitempty"`

			ProjectID *string `tfsdk:"project_id" yaml:"projectID,omitempty"`

			Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`
		} `tfsdk:"metadata_properties" yaml:"metadataProperties,omitempty"`

		ModelApprovalStatus *string `tfsdk:"model_approval_status" yaml:"modelApprovalStatus,omitempty"`

		ModelMetrics *struct {
			Bias *struct {
				PostTrainingReport *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"post_training_report" yaml:"postTrainingReport,omitempty"`

				PreTrainingReport *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"pre_training_report" yaml:"preTrainingReport,omitempty"`

				Report *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"report" yaml:"report,omitempty"`
			} `tfsdk:"bias" yaml:"bias,omitempty"`

			Explainability *struct {
				Report *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"report" yaml:"report,omitempty"`
			} `tfsdk:"explainability" yaml:"explainability,omitempty"`

			ModelDataQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"constraints" yaml:"constraints,omitempty"`

				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"statistics" yaml:"statistics,omitempty"`
			} `tfsdk:"model_data_quality" yaml:"modelDataQuality,omitempty"`

			ModelQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"constraints" yaml:"constraints,omitempty"`

				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" yaml:"contentDigest,omitempty"`

					ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

					S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
				} `tfsdk:"statistics" yaml:"statistics,omitempty"`
			} `tfsdk:"model_quality" yaml:"modelQuality,omitempty"`
		} `tfsdk:"model_metrics" yaml:"modelMetrics,omitempty"`

		ModelPackageDescription *string `tfsdk:"model_package_description" yaml:"modelPackageDescription,omitempty"`

		ModelPackageGroupName *string `tfsdk:"model_package_group_name" yaml:"modelPackageGroupName,omitempty"`

		ModelPackageName *string `tfsdk:"model_package_name" yaml:"modelPackageName,omitempty"`

		SamplePayloadURL *string `tfsdk:"sample_payload_url" yaml:"samplePayloadURL,omitempty"`

		SourceAlgorithmSpecification *struct {
			SourceAlgorithms *[]struct {
				AlgorithmName *string `tfsdk:"algorithm_name" yaml:"algorithmName,omitempty"`

				ModelDataURL *string `tfsdk:"model_data_url" yaml:"modelDataURL,omitempty"`
			} `tfsdk:"source_algorithms" yaml:"sourceAlgorithms,omitempty"`
		} `tfsdk:"source_algorithm_specification" yaml:"sourceAlgorithmSpecification,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		Task *string `tfsdk:"task" yaml:"task,omitempty"`

		ValidationSpecification *struct {
			ValidationProfiles *[]struct {
				ProfileName *string `tfsdk:"profile_name" yaml:"profileName,omitempty"`

				TransformJobDefinition *struct {
					BatchStrategy *string `tfsdk:"batch_strategy" yaml:"batchStrategy,omitempty"`

					Environment *map[string]string `tfsdk:"environment" yaml:"environment,omitempty"`

					MaxConcurrentTransforms *int64 `tfsdk:"max_concurrent_transforms" yaml:"maxConcurrentTransforms,omitempty"`

					MaxPayloadInMB *int64 `tfsdk:"max_payload_in_mb" yaml:"maxPayloadInMB,omitempty"`

					TransformInput *struct {
						CompressionType *string `tfsdk:"compression_type" yaml:"compressionType,omitempty"`

						ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

						DataSource *struct {
							S3DataSource *struct {
								S3DataType *string `tfsdk:"s3_data_type" yaml:"s3DataType,omitempty"`

								S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
							} `tfsdk:"s3_data_source" yaml:"s3DataSource,omitempty"`
						} `tfsdk:"data_source" yaml:"dataSource,omitempty"`

						SplitType *string `tfsdk:"split_type" yaml:"splitType,omitempty"`
					} `tfsdk:"transform_input" yaml:"transformInput,omitempty"`

					TransformOutput *struct {
						Accept *string `tfsdk:"accept" yaml:"accept,omitempty"`

						AssembleWith *string `tfsdk:"assemble_with" yaml:"assembleWith,omitempty"`

						KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

						S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`
					} `tfsdk:"transform_output" yaml:"transformOutput,omitempty"`

					TransformResources *struct {
						InstanceCount *int64 `tfsdk:"instance_count" yaml:"instanceCount,omitempty"`

						InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

						VolumeKMSKeyID *string `tfsdk:"volume_kms_key_id" yaml:"volumeKMSKeyID,omitempty"`
					} `tfsdk:"transform_resources" yaml:"transformResources,omitempty"`
				} `tfsdk:"transform_job_definition" yaml:"transformJobDefinition,omitempty"`
			} `tfsdk:"validation_profiles" yaml:"validationProfiles,omitempty"`

			ValidationRole *string `tfsdk:"validation_role" yaml:"validationRole,omitempty"`
		} `tfsdk:"validation_specification" yaml:"validationSpecification,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsModelPackageV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsModelPackageV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_model_package_v1alpha1"
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ModelPackage is the Schema for the ModelPackages API",
		MarkdownDescription: "ModelPackage is the Schema for the ModelPackages API",
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
				Description:         "ModelPackageSpec defines the desired state of ModelPackage.  A versioned model that can be deployed for SageMaker inference.",
				MarkdownDescription: "ModelPackageSpec defines the desired state of ModelPackage.  A versioned model that can be deployed for SageMaker inference.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"additional_inference_specifications": {
						Description:         "An array of additional Inference Specification objects. Each additional Inference Specification specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.",
						MarkdownDescription: "An array of additional Inference Specification objects. Each additional Inference Specification specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"containers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"container_hostname": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"environment": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"framework": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"framework_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_digest": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"model_data_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"model_input": {
										Description:         "Input object for the model.",
										MarkdownDescription: "Input object for the model.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"data_input_config": {
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

									"nearest_model_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"product_id": {
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
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_content_types": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_realtime_inference_instance_types": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_response_mime_types": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_transform_instance_types": {
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

					"approval_description": {
						Description:         "A description for the approval status of the model.",
						MarkdownDescription: "A description for the approval status of the model.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"certify_for_marketplace": {
						Description:         "Whether to certify the model package for listing on Amazon Web Services Marketplace.  This parameter is optional for unversioned models, and does not apply to versioned models.",
						MarkdownDescription: "Whether to certify the model package for listing on Amazon Web Services Marketplace.  This parameter is optional for unversioned models, and does not apply to versioned models.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"client_token": {
						Description:         "A unique token that guarantees that the call to this API is idempotent.",
						MarkdownDescription: "A unique token that guarantees that the call to this API is idempotent.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"customer_metadata_properties": {
						Description:         "The metadata properties associated with the model package versions.",
						MarkdownDescription: "The metadata properties associated with the model package versions.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"domain": {
						Description:         "The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.",
						MarkdownDescription: "The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"drift_check_baselines": {
						Description:         "Represents the drift check baselines that can be used when the model monitor is set using the model package. For more information, see the topic on Drift Detection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection) in the Amazon SageMaker Developer Guide.",
						MarkdownDescription: "Represents the drift check baselines that can be used when the model monitor is set using the model package. For more information, see the topic on Drift Detection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection) in the Amazon SageMaker Developer Guide.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bias": {
								Description:         "Represents the drift check bias baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check bias baselines that can be used when the model monitor is set using the model package.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_file": {
										Description:         "Contains details regarding the file source.",
										MarkdownDescription: "Contains details regarding the file source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"post_training_constraints": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"pre_training_constraints": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

							"explainability": {
								Description:         "Represents the drift check explainability baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check explainability baselines that can be used when the model monitor is set using the model package.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_file": {
										Description:         "Contains details regarding the file source.",
										MarkdownDescription: "Contains details regarding the file source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"constraints": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

							"model_data_quality": {
								Description:         "Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"constraints": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"statistics": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

							"model_quality": {
								Description:         "Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"constraints": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"statistics": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"inference_specification": {
						Description:         "Specifies details about inference jobs that can be run with models based on this model package, including the following:  * The Amazon ECR paths of containers that contain the inference code and model artifacts.  * The instance types that the model package supports for transform jobs and real-time endpoints used for inference.  * The input and output content formats that the model package supports for inference.",
						MarkdownDescription: "Specifies details about inference jobs that can be run with models based on this model package, including the following:  * The Amazon ECR paths of containers that contain the inference code and model artifacts.  * The instance types that the model package supports for transform jobs and real-time endpoints used for inference.  * The input and output content formats that the model package supports for inference.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"containers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"container_hostname": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"environment": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"framework": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"framework_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_digest": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"model_data_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"model_input": {
										Description:         "Input object for the model.",
										MarkdownDescription: "Input object for the model.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"data_input_config": {
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

									"nearest_model_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"product_id": {
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

							"supported_content_types": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_realtime_inference_instance_types": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_response_mime_types": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"supported_transform_instance_types": {
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

					"metadata_properties": {
						Description:         "Metadata properties of the tracking entity, trial, or trial component.",
						MarkdownDescription: "Metadata properties of the tracking entity, trial, or trial component.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"commit_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"generated_by": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"project_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"repository": {
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

					"model_approval_status": {
						Description:         "Whether the model is approved for deployment.  This parameter is optional for versioned models, and does not apply to unversioned models.  For versioned models, the value of this parameter must be set to Approved to deploy the model.",
						MarkdownDescription: "Whether the model is approved for deployment.  This parameter is optional for versioned models, and does not apply to unversioned models.  For versioned models, the value of this parameter must be set to Approved to deploy the model.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"model_metrics": {
						Description:         "A structure that contains model metrics reports.",
						MarkdownDescription: "A structure that contains model metrics reports.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bias": {
								Description:         "Contains bias metrics for a model.",
								MarkdownDescription: "Contains bias metrics for a model.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"post_training_report": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"pre_training_report": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"report": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

							"explainability": {
								Description:         "Contains explainability metrics for a model.",
								MarkdownDescription: "Contains explainability metrics for a model.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"report": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

							"model_data_quality": {
								Description:         "Data quality constraints and statistics for a model.",
								MarkdownDescription: "Data quality constraints and statistics for a model.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"constraints": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"statistics": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

							"model_quality": {
								Description:         "Model quality statistics and constraints.",
								MarkdownDescription: "Model quality statistics and constraints.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"constraints": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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

									"statistics": {
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_digest": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"content_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"s3_uri": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"model_package_description": {
						Description:         "A description of the model package.",
						MarkdownDescription: "A description of the model package.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"model_package_group_name": {
						Description:         "The name or Amazon Resource Name (ARN) of the model package group that this model version belongs to.  This parameter is required for versioned models, and does not apply to unversioned models.",
						MarkdownDescription: "The name or Amazon Resource Name (ARN) of the model package group that this model version belongs to.  This parameter is required for versioned models, and does not apply to unversioned models.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"model_package_name": {
						Description:         "The name of the model package. The name must have 1 to 63 characters. Valid characters are a-z, A-Z, 0-9, and - (hyphen).  This parameter is required for unversioned models. It is not applicable to versioned models.",
						MarkdownDescription: "The name of the model package. The name must have 1 to 63 characters. Valid characters are a-z, A-Z, 0-9, and - (hyphen).  This parameter is required for unversioned models. It is not applicable to versioned models.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sample_payload_url": {
						Description:         "The Amazon Simple Storage Service (Amazon S3) path where the sample payload are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix).",
						MarkdownDescription: "The Amazon Simple Storage Service (Amazon S3) path where the sample payload are stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"source_algorithm_specification": {
						Description:         "Details about the algorithm that was used to create the model package.",
						MarkdownDescription: "Details about the algorithm that was used to create the model package.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"source_algorithms": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"algorithm_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"model_data_url": {
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

					"tags": {
						Description:         "A list of key value pairs associated with the model. For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon Web Services General Reference Guide.",
						MarkdownDescription: "A list of key value pairs associated with the model. For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon Web Services General Reference Guide.",

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

					"task": {
						Description:         "The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification. The following tasks are supported by Inference Recommender: 'IMAGE_CLASSIFICATION' | 'OBJECT_DETECTION' | 'TEXT_GENERATION' |'IMAGE_SEGMENTATION' | 'FILL_MASK' | 'CLASSIFICATION' | 'REGRESSION' | 'OTHER'.  Specify 'OTHER' if none of the tasks listed fit your use case.",
						MarkdownDescription: "The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification. The following tasks are supported by Inference Recommender: 'IMAGE_CLASSIFICATION' | 'OBJECT_DETECTION' | 'TEXT_GENERATION' |'IMAGE_SEGMENTATION' | 'FILL_MASK' | 'CLASSIFICATION' | 'REGRESSION' | 'OTHER'.  Specify 'OTHER' if none of the tasks listed fit your use case.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"validation_specification": {
						Description:         "Specifies configurations for one or more transform jobs that SageMaker runs to test the model package.",
						MarkdownDescription: "Specifies configurations for one or more transform jobs that SageMaker runs to test the model package.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"validation_profiles": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"profile_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"transform_job_definition": {
										Description:         "Defines the input needed to run a transform job using the inference specification specified in the algorithm.",
										MarkdownDescription: "Defines the input needed to run a transform job using the inference specification specified in the algorithm.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"batch_strategy": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"environment": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_concurrent_transforms": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_payload_in_mb": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"transform_input": {
												Description:         "Describes the input source of a transform job and the way the transform job consumes it.",
												MarkdownDescription: "Describes the input source of a transform job and the way the transform job consumes it.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"compression_type": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"content_type": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"data_source": {
														Description:         "Describes the location of the channel data.",
														MarkdownDescription: "Describes the location of the channel data.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"s3_data_source": {
																Description:         "Describes the S3 data source.",
																MarkdownDescription: "Describes the S3 data source.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"s3_data_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"s3_uri": {
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

													"split_type": {
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

											"transform_output": {
												Description:         "Describes the results of a transform job.",
												MarkdownDescription: "Describes the results of a transform job.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"accept": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"assemble_with": {
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

											"transform_resources": {
												Description:         "Describes the resources, including ML instance types and ML instance count, to use for transform job.",
												MarkdownDescription: "Describes the resources, including ML instance types and ML instance count, to use for transform job.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"instance_count": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

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

													"volume_kms_key_id": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"validation_role": {
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

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_model_package_v1alpha1")

	var state SagemakerServicesK8SAwsModelPackageV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsModelPackageV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ModelPackage")

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

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_model_package_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_model_package_v1alpha1")

	var state SagemakerServicesK8SAwsModelPackageV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsModelPackageV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("ModelPackage")

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

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_model_package_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
