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
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SagemakerServicesK8SAwsModelPackageV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsModelPackageV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsModelPackageV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsModelPackageV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsModelPackageV1Alpha1ManifestData struct {
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
		AdditionalInferenceSpecifications *[]struct {
			Containers *[]struct {
				AdditionalS3DataSource *struct {
					CompressionType *string `tfsdk:"compression_type" json:"compressionType,omitempty"`
					S3DataType      *string `tfsdk:"s3_data_type" json:"s3DataType,omitempty"`
					S3URI           *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"additional_s3_data_source" json:"additionalS3DataSource,omitempty"`
				ContainerHostname *string            `tfsdk:"container_hostname" json:"containerHostname,omitempty"`
				Environment       *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
				Framework         *string            `tfsdk:"framework" json:"framework,omitempty"`
				FrameworkVersion  *string            `tfsdk:"framework_version" json:"frameworkVersion,omitempty"`
				Image             *string            `tfsdk:"image" json:"image,omitempty"`
				ImageDigest       *string            `tfsdk:"image_digest" json:"imageDigest,omitempty"`
				ModelDataURL      *string            `tfsdk:"model_data_url" json:"modelDataURL,omitempty"`
				ModelInput        *struct {
					DataInputConfig *string `tfsdk:"data_input_config" json:"dataInputConfig,omitempty"`
				} `tfsdk:"model_input" json:"modelInput,omitempty"`
				NearestModelName *string `tfsdk:"nearest_model_name" json:"nearestModelName,omitempty"`
				ProductID        *string `tfsdk:"product_id" json:"productID,omitempty"`
			} `tfsdk:"containers" json:"containers,omitempty"`
			Description                             *string   `tfsdk:"description" json:"description,omitempty"`
			Name                                    *string   `tfsdk:"name" json:"name,omitempty"`
			SupportedContentTypes                   *[]string `tfsdk:"supported_content_types" json:"supportedContentTypes,omitempty"`
			SupportedRealtimeInferenceInstanceTypes *[]string `tfsdk:"supported_realtime_inference_instance_types" json:"supportedRealtimeInferenceInstanceTypes,omitempty"`
			SupportedResponseMIMETypes              *[]string `tfsdk:"supported_response_mime_types" json:"supportedResponseMIMETypes,omitempty"`
			SupportedTransformInstanceTypes         *[]string `tfsdk:"supported_transform_instance_types" json:"supportedTransformInstanceTypes,omitempty"`
		} `tfsdk:"additional_inference_specifications" json:"additionalInferenceSpecifications,omitempty"`
		ApprovalDescription        *string            `tfsdk:"approval_description" json:"approvalDescription,omitempty"`
		CertifyForMarketplace      *bool              `tfsdk:"certify_for_marketplace" json:"certifyForMarketplace,omitempty"`
		ClientToken                *string            `tfsdk:"client_token" json:"clientToken,omitempty"`
		CustomerMetadataProperties *map[string]string `tfsdk:"customer_metadata_properties" json:"customerMetadataProperties,omitempty"`
		Domain                     *string            `tfsdk:"domain" json:"domain,omitempty"`
		DriftCheckBaselines        *struct {
			Bias *struct {
				ConfigFile *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"config_file" json:"configFile,omitempty"`
				PostTrainingConstraints *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"post_training_constraints" json:"postTrainingConstraints,omitempty"`
				PreTrainingConstraints *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"pre_training_constraints" json:"preTrainingConstraints,omitempty"`
			} `tfsdk:"bias" json:"bias,omitempty"`
			Explainability *struct {
				ConfigFile *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"config_file" json:"configFile,omitempty"`
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"constraints" json:"constraints,omitempty"`
			} `tfsdk:"explainability" json:"explainability,omitempty"`
			ModelDataQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"constraints" json:"constraints,omitempty"`
				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"statistics" json:"statistics,omitempty"`
			} `tfsdk:"model_data_quality" json:"modelDataQuality,omitempty"`
			ModelQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"constraints" json:"constraints,omitempty"`
				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"statistics" json:"statistics,omitempty"`
			} `tfsdk:"model_quality" json:"modelQuality,omitempty"`
		} `tfsdk:"drift_check_baselines" json:"driftCheckBaselines,omitempty"`
		InferenceSpecification *struct {
			Containers *[]struct {
				AdditionalS3DataSource *struct {
					CompressionType *string `tfsdk:"compression_type" json:"compressionType,omitempty"`
					S3DataType      *string `tfsdk:"s3_data_type" json:"s3DataType,omitempty"`
					S3URI           *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"additional_s3_data_source" json:"additionalS3DataSource,omitempty"`
				ContainerHostname *string            `tfsdk:"container_hostname" json:"containerHostname,omitempty"`
				Environment       *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
				Framework         *string            `tfsdk:"framework" json:"framework,omitempty"`
				FrameworkVersion  *string            `tfsdk:"framework_version" json:"frameworkVersion,omitempty"`
				Image             *string            `tfsdk:"image" json:"image,omitempty"`
				ImageDigest       *string            `tfsdk:"image_digest" json:"imageDigest,omitempty"`
				ModelDataURL      *string            `tfsdk:"model_data_url" json:"modelDataURL,omitempty"`
				ModelInput        *struct {
					DataInputConfig *string `tfsdk:"data_input_config" json:"dataInputConfig,omitempty"`
				} `tfsdk:"model_input" json:"modelInput,omitempty"`
				NearestModelName *string `tfsdk:"nearest_model_name" json:"nearestModelName,omitempty"`
				ProductID        *string `tfsdk:"product_id" json:"productID,omitempty"`
			} `tfsdk:"containers" json:"containers,omitempty"`
			SupportedContentTypes                   *[]string `tfsdk:"supported_content_types" json:"supportedContentTypes,omitempty"`
			SupportedRealtimeInferenceInstanceTypes *[]string `tfsdk:"supported_realtime_inference_instance_types" json:"supportedRealtimeInferenceInstanceTypes,omitempty"`
			SupportedResponseMIMETypes              *[]string `tfsdk:"supported_response_mime_types" json:"supportedResponseMIMETypes,omitempty"`
			SupportedTransformInstanceTypes         *[]string `tfsdk:"supported_transform_instance_types" json:"supportedTransformInstanceTypes,omitempty"`
		} `tfsdk:"inference_specification" json:"inferenceSpecification,omitempty"`
		MetadataProperties *struct {
			CommitID    *string `tfsdk:"commit_id" json:"commitID,omitempty"`
			GeneratedBy *string `tfsdk:"generated_by" json:"generatedBy,omitempty"`
			ProjectID   *string `tfsdk:"project_id" json:"projectID,omitempty"`
			Repository  *string `tfsdk:"repository" json:"repository,omitempty"`
		} `tfsdk:"metadata_properties" json:"metadataProperties,omitempty"`
		ModelApprovalStatus *string `tfsdk:"model_approval_status" json:"modelApprovalStatus,omitempty"`
		ModelMetrics        *struct {
			Bias *struct {
				PostTrainingReport *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"post_training_report" json:"postTrainingReport,omitempty"`
				PreTrainingReport *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"pre_training_report" json:"preTrainingReport,omitempty"`
				Report *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"report" json:"report,omitempty"`
			} `tfsdk:"bias" json:"bias,omitempty"`
			Explainability *struct {
				Report *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"report" json:"report,omitempty"`
			} `tfsdk:"explainability" json:"explainability,omitempty"`
			ModelDataQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"constraints" json:"constraints,omitempty"`
				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"statistics" json:"statistics,omitempty"`
			} `tfsdk:"model_data_quality" json:"modelDataQuality,omitempty"`
			ModelQuality *struct {
				Constraints *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"constraints" json:"constraints,omitempty"`
				Statistics *struct {
					ContentDigest *string `tfsdk:"content_digest" json:"contentDigest,omitempty"`
					ContentType   *string `tfsdk:"content_type" json:"contentType,omitempty"`
					S3URI         *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
				} `tfsdk:"statistics" json:"statistics,omitempty"`
			} `tfsdk:"model_quality" json:"modelQuality,omitempty"`
		} `tfsdk:"model_metrics" json:"modelMetrics,omitempty"`
		ModelPackageDescription      *string `tfsdk:"model_package_description" json:"modelPackageDescription,omitempty"`
		ModelPackageGroupName        *string `tfsdk:"model_package_group_name" json:"modelPackageGroupName,omitempty"`
		ModelPackageName             *string `tfsdk:"model_package_name" json:"modelPackageName,omitempty"`
		SamplePayloadURL             *string `tfsdk:"sample_payload_url" json:"samplePayloadURL,omitempty"`
		SkipModelValidation          *string `tfsdk:"skip_model_validation" json:"skipModelValidation,omitempty"`
		SourceAlgorithmSpecification *struct {
			SourceAlgorithms *[]struct {
				AlgorithmName *string `tfsdk:"algorithm_name" json:"algorithmName,omitempty"`
				ModelDataURL  *string `tfsdk:"model_data_url" json:"modelDataURL,omitempty"`
			} `tfsdk:"source_algorithms" json:"sourceAlgorithms,omitempty"`
		} `tfsdk:"source_algorithm_specification" json:"sourceAlgorithmSpecification,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		Task                    *string `tfsdk:"task" json:"task,omitempty"`
		ValidationSpecification *struct {
			ValidationProfiles *[]struct {
				ProfileName            *string `tfsdk:"profile_name" json:"profileName,omitempty"`
				TransformJobDefinition *struct {
					BatchStrategy           *string            `tfsdk:"batch_strategy" json:"batchStrategy,omitempty"`
					Environment             *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
					MaxConcurrentTransforms *int64             `tfsdk:"max_concurrent_transforms" json:"maxConcurrentTransforms,omitempty"`
					MaxPayloadInMB          *int64             `tfsdk:"max_payload_in_mb" json:"maxPayloadInMB,omitempty"`
					TransformInput          *struct {
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
					TransformOutput *struct {
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
				} `tfsdk:"transform_job_definition" json:"transformJobDefinition,omitempty"`
			} `tfsdk:"validation_profiles" json:"validationProfiles,omitempty"`
			ValidationRole *string `tfsdk:"validation_role" json:"validationRole,omitempty"`
		} `tfsdk:"validation_specification" json:"validationSpecification,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_model_package_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ModelPackage is the Schema for the ModelPackages API",
		MarkdownDescription: "ModelPackage is the Schema for the ModelPackages API",
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
				Description:         "ModelPackageSpec defines the desired state of ModelPackage.A versioned model that can be deployed for SageMaker inference.",
				MarkdownDescription: "ModelPackageSpec defines the desired state of ModelPackage.A versioned model that can be deployed for SageMaker inference.",
				Attributes: map[string]schema.Attribute{
					"additional_inference_specifications": schema.ListNestedAttribute{
						Description:         "An array of additional Inference Specification objects. Each additional InferenceSpecification specifies artifacts based on this model package that can beused on inference endpoints. Generally used with SageMaker Neo to store thecompiled artifacts.",
						MarkdownDescription: "An array of additional Inference Specification objects. Each additional InferenceSpecification specifies artifacts based on this model package that can beused on inference endpoints. Generally used with SageMaker Neo to store thecompiled artifacts.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"containers": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"additional_s3_data_source": schema.SingleNestedAttribute{
												Description:         "A data source used for training or inference that is in addition to the inputdataset or model data.",
												MarkdownDescription: "A data source used for training or inference that is in addition to the inputdataset or model data.",
												Attributes: map[string]schema.Attribute{
													"compression_type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"s3_data_type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"s3_uri": schema.StringAttribute{
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

											"container_hostname": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"environment": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"framework": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"framework_version": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"model_data_url": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"model_input": schema.SingleNestedAttribute{
												Description:         "Input object for the model.",
												MarkdownDescription: "Input object for the model.",
												Attributes: map[string]schema.Attribute{
													"data_input_config": schema.StringAttribute{
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

											"nearest_model_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"product_id": schema.StringAttribute{
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

								"description": schema.StringAttribute{
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

								"supported_content_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"supported_realtime_inference_instance_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"supported_response_mime_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"supported_transform_instance_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
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

					"approval_description": schema.StringAttribute{
						Description:         "A description for the approval status of the model.",
						MarkdownDescription: "A description for the approval status of the model.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"certify_for_marketplace": schema.BoolAttribute{
						Description:         "Whether to certify the model package for listing on Amazon Web Services Marketplace.This parameter is optional for unversioned models, and does not apply toversioned models.",
						MarkdownDescription: "Whether to certify the model package for listing on Amazon Web Services Marketplace.This parameter is optional for unversioned models, and does not apply toversioned models.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_token": schema.StringAttribute{
						Description:         "A unique token that guarantees that the call to this API is idempotent.",
						MarkdownDescription: "A unique token that guarantees that the call to this API is idempotent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"customer_metadata_properties": schema.MapAttribute{
						Description:         "The metadata properties associated with the model package versions.",
						MarkdownDescription: "The metadata properties associated with the model package versions.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"domain": schema.StringAttribute{
						Description:         "The machine learning domain of your model package and its components. Commonmachine learning domains include computer vision and natural language processing.",
						MarkdownDescription: "The machine learning domain of your model package and its components. Commonmachine learning domains include computer vision and natural language processing.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"drift_check_baselines": schema.SingleNestedAttribute{
						Description:         "Represents the drift check baselines that can be used when the model monitoris set using the model package. For more information, see the topic on DriftDetection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection)in the Amazon SageMaker Developer Guide.",
						MarkdownDescription: "Represents the drift check baselines that can be used when the model monitoris set using the model package. For more information, see the topic on DriftDetection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection)in the Amazon SageMaker Developer Guide.",
						Attributes: map[string]schema.Attribute{
							"bias": schema.SingleNestedAttribute{
								Description:         "Represents the drift check bias baselines that can be used when the modelmonitor is set using the model package.",
								MarkdownDescription: "Represents the drift check bias baselines that can be used when the modelmonitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"config_file": schema.SingleNestedAttribute{
										Description:         "Contains details regarding the file source.",
										MarkdownDescription: "Contains details regarding the file source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"post_training_constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"pre_training_constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

							"explainability": schema.SingleNestedAttribute{
								Description:         "Represents the drift check explainability baselines that can be used whenthe model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check explainability baselines that can be used whenthe model monitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"config_file": schema.SingleNestedAttribute{
										Description:         "Contains details regarding the file source.",
										MarkdownDescription: "Contains details regarding the file source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

							"model_data_quality": schema.SingleNestedAttribute{
								Description:         "Represents the drift check data quality baselines that can be used when themodel monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check data quality baselines that can be used when themodel monitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

							"model_quality": schema.SingleNestedAttribute{
								Description:         "Represents the drift check model quality baselines that can be used whenthe model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check model quality baselines that can be used whenthe model monitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

					"inference_specification": schema.SingleNestedAttribute{
						Description:         "Specifies details about inference jobs that can be run with models basedon this model package, including the following:   * The Amazon ECR paths of containers that contain the inference code and   model artifacts.   * The instance types that the model package supports for transform jobs   and real-time endpoints used for inference.   * The input and output content formats that the model package supports   for inference.",
						MarkdownDescription: "Specifies details about inference jobs that can be run with models basedon this model package, including the following:   * The Amazon ECR paths of containers that contain the inference code and   model artifacts.   * The instance types that the model package supports for transform jobs   and real-time endpoints used for inference.   * The input and output content formats that the model package supports   for inference.",
						Attributes: map[string]schema.Attribute{
							"containers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"additional_s3_data_source": schema.SingleNestedAttribute{
											Description:         "A data source used for training or inference that is in addition to the inputdataset or model data.",
											MarkdownDescription: "A data source used for training or inference that is in addition to the inputdataset or model data.",
											Attributes: map[string]schema.Attribute{
												"compression_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"s3_data_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"s3_uri": schema.StringAttribute{
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

										"container_hostname": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"environment": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"framework": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"framework_version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image_digest": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"model_data_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"model_input": schema.SingleNestedAttribute{
											Description:         "Input object for the model.",
											MarkdownDescription: "Input object for the model.",
											Attributes: map[string]schema.Attribute{
												"data_input_config": schema.StringAttribute{
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

										"nearest_model_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"product_id": schema.StringAttribute{
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

							"supported_content_types": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supported_realtime_inference_instance_types": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supported_response_mime_types": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supported_transform_instance_types": schema.ListAttribute{
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

					"metadata_properties": schema.SingleNestedAttribute{
						Description:         "Metadata properties of the tracking entity, trial, or trial component.",
						MarkdownDescription: "Metadata properties of the tracking entity, trial, or trial component.",
						Attributes: map[string]schema.Attribute{
							"commit_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"generated_by": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"project_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
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

					"model_approval_status": schema.StringAttribute{
						Description:         "Whether the model is approved for deployment.This parameter is optional for versioned models, and does not apply to unversionedmodels.For versioned models, the value of this parameter must be set to Approvedto deploy the model.",
						MarkdownDescription: "Whether the model is approved for deployment.This parameter is optional for versioned models, and does not apply to unversionedmodels.For versioned models, the value of this parameter must be set to Approvedto deploy the model.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"model_metrics": schema.SingleNestedAttribute{
						Description:         "A structure that contains model metrics reports.",
						MarkdownDescription: "A structure that contains model metrics reports.",
						Attributes: map[string]schema.Attribute{
							"bias": schema.SingleNestedAttribute{
								Description:         "Contains bias metrics for a model.",
								MarkdownDescription: "Contains bias metrics for a model.",
								Attributes: map[string]schema.Attribute{
									"post_training_report": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"pre_training_report": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"report": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

							"explainability": schema.SingleNestedAttribute{
								Description:         "Contains explainability metrics for a model.",
								MarkdownDescription: "Contains explainability metrics for a model.",
								Attributes: map[string]schema.Attribute{
									"report": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

							"model_data_quality": schema.SingleNestedAttribute{
								Description:         "Data quality constraints and statistics for a model.",
								MarkdownDescription: "Data quality constraints and statistics for a model.",
								Attributes: map[string]schema.Attribute{
									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

							"model_quality": schema.SingleNestedAttribute{
								Description:         "Model quality statistics and constraints.",
								MarkdownDescription: "Model quality statistics and constraints.",
								Attributes: map[string]schema.Attribute{
									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"s3_uri": schema.StringAttribute{
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

					"model_package_description": schema.StringAttribute{
						Description:         "A description of the model package.",
						MarkdownDescription: "A description of the model package.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"model_package_group_name": schema.StringAttribute{
						Description:         "The name or Amazon Resource Name (ARN) of the model package group that thismodel version belongs to.This parameter is required for versioned models, and does not apply to unversionedmodels.",
						MarkdownDescription: "The name or Amazon Resource Name (ARN) of the model package group that thismodel version belongs to.This parameter is required for versioned models, and does not apply to unversionedmodels.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"model_package_name": schema.StringAttribute{
						Description:         "The name of the model package. The name must have 1 to 63 characters. Validcharacters are a-z, A-Z, 0-9, and - (hyphen).This parameter is required for unversioned models. It is not applicable toversioned models.",
						MarkdownDescription: "The name of the model package. The name must have 1 to 63 characters. Validcharacters are a-z, A-Z, 0-9, and - (hyphen).This parameter is required for unversioned models. It is not applicable toversioned models.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sample_payload_url": schema.StringAttribute{
						Description:         "The Amazon Simple Storage Service (Amazon S3) path where the sample payloadis stored. This path must point to a single gzip compressed tar archive (.tar.gzsuffix). This archive can hold multiple files that are all equally used inthe load test. Each file in the archive must satisfy the size constraintsof the InvokeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html#API_runtime_InvokeEndpoint_RequestSyntax)call.",
						MarkdownDescription: "The Amazon Simple Storage Service (Amazon S3) path where the sample payloadis stored. This path must point to a single gzip compressed tar archive (.tar.gzsuffix). This archive can hold multiple files that are all equally used inthe load test. Each file in the archive must satisfy the size constraintsof the InvokeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html#API_runtime_InvokeEndpoint_RequestSyntax)call.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"skip_model_validation": schema.StringAttribute{
						Description:         "Indicates if you want to skip model validation.",
						MarkdownDescription: "Indicates if you want to skip model validation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_algorithm_specification": schema.SingleNestedAttribute{
						Description:         "Details about the algorithm that was used to create the model package.",
						MarkdownDescription: "Details about the algorithm that was used to create the model package.",
						Attributes: map[string]schema.Attribute{
							"source_algorithms": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"algorithm_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"model_data_url": schema.StringAttribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "A list of key value pairs associated with the model. For more information,see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)in the Amazon Web Services General Reference Guide.If you supply ModelPackageGroupName, your model package belongs to the modelgroup you specify and uses the tags associated with the model group. In thiscase, you cannot supply a tag argument.",
						MarkdownDescription: "A list of key value pairs associated with the model. For more information,see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)in the Amazon Web Services General Reference Guide.If you supply ModelPackageGroupName, your model package belongs to the modelgroup you specify and uses the tags associated with the model group. In thiscase, you cannot supply a tag argument.",
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

					"task": schema.StringAttribute{
						Description:         "The machine learning task your model package accomplishes. Common machinelearning tasks include object detection and image classification. The followingtasks are supported by Inference Recommender: 'IMAGE_CLASSIFICATION' | 'OBJECT_DETECTION'| 'TEXT_GENERATION' |'IMAGE_SEGMENTATION' | 'FILL_MASK' | 'CLASSIFICATION'| 'REGRESSION' | 'OTHER'.Specify 'OTHER' if none of the tasks listed fit your use case.",
						MarkdownDescription: "The machine learning task your model package accomplishes. Common machinelearning tasks include object detection and image classification. The followingtasks are supported by Inference Recommender: 'IMAGE_CLASSIFICATION' | 'OBJECT_DETECTION'| 'TEXT_GENERATION' |'IMAGE_SEGMENTATION' | 'FILL_MASK' | 'CLASSIFICATION'| 'REGRESSION' | 'OTHER'.Specify 'OTHER' if none of the tasks listed fit your use case.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"validation_specification": schema.SingleNestedAttribute{
						Description:         "Specifies configurations for one or more transform jobs that SageMaker runsto test the model package.",
						MarkdownDescription: "Specifies configurations for one or more transform jobs that SageMaker runsto test the model package.",
						Attributes: map[string]schema.Attribute{
							"validation_profiles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"profile_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"transform_job_definition": schema.SingleNestedAttribute{
											Description:         "Defines the input needed to run a transform job using the inference specificationspecified in the algorithm.",
											MarkdownDescription: "Defines the input needed to run a transform job using the inference specificationspecified in the algorithm.",
											Attributes: map[string]schema.Attribute{
												"batch_strategy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"environment": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_concurrent_transforms": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_payload_in_mb": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"transform_input": schema.SingleNestedAttribute{
													Description:         "Describes the input source of a transform job and the way the transform jobconsumes it.",
													MarkdownDescription: "Describes the input source of a transform job and the way the transform jobconsumes it.",
													Attributes: map[string]schema.Attribute{
														"compression_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"content_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
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
																			Optional:            true,
																			Computed:            false,
																		},

																		"s3_uri": schema.StringAttribute{
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

														"split_type": schema.StringAttribute{
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

												"transform_output": schema.SingleNestedAttribute{
													Description:         "Describes the results of a transform job.",
													MarkdownDescription: "Describes the results of a transform job.",
													Attributes: map[string]schema.Attribute{
														"accept": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"assemble_with": schema.StringAttribute{
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

												"transform_resources": schema.SingleNestedAttribute{
													Description:         "Describes the resources, including ML instance types and ML instance count,to use for transform job.",
													MarkdownDescription: "Describes the resources, including ML instance types and ML instance count,to use for transform job.",
													Attributes: map[string]schema.Attribute{
														"instance_count": schema.Int64Attribute{
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

														"volume_kms_key_id": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"validation_role": schema.StringAttribute{
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
	}
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_model_package_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsModelPackageV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("ModelPackage")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
