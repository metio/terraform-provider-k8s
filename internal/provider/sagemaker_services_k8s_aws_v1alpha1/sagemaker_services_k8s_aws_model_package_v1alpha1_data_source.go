/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource              = &SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsModelPackageV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsModelPackageV1Alpha1DataSourceData struct {
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
		AdditionalInferenceSpecifications *[]struct {
			Containers *[]struct {
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

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_model_package_v1alpha1"
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "ModelPackageSpec defines the desired state of ModelPackage.  A versioned model that can be deployed for SageMaker inference.",
				MarkdownDescription: "ModelPackageSpec defines the desired state of ModelPackage.  A versioned model that can be deployed for SageMaker inference.",
				Attributes: map[string]schema.Attribute{
					"additional_inference_specifications": schema.ListNestedAttribute{
						Description:         "An array of additional Inference Specification objects. Each additional Inference Specification specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.",
						MarkdownDescription: "An array of additional Inference Specification objects. Each additional Inference Specification specifies artifacts based on this model package that can be used on inference endpoints. Generally used with SageMaker Neo to store the compiled artifacts.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"containers": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"container_hostname": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"environment": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"framework": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"framework_version": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"image_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"model_data_url": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"model_input": schema.SingleNestedAttribute{
												Description:         "Input object for the model.",
												MarkdownDescription: "Input object for the model.",
												Attributes: map[string]schema.Attribute{
													"data_input_config": schema.StringAttribute{
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

											"nearest_model_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"product_id": schema.StringAttribute{
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

								"description": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"supported_content_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"supported_realtime_inference_instance_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"supported_response_mime_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"supported_transform_instance_types": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
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

					"approval_description": schema.StringAttribute{
						Description:         "A description for the approval status of the model.",
						MarkdownDescription: "A description for the approval status of the model.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"certify_for_marketplace": schema.BoolAttribute{
						Description:         "Whether to certify the model package for listing on Amazon Web Services Marketplace.  This parameter is optional for unversioned models, and does not apply to versioned models.",
						MarkdownDescription: "Whether to certify the model package for listing on Amazon Web Services Marketplace.  This parameter is optional for unversioned models, and does not apply to versioned models.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"client_token": schema.StringAttribute{
						Description:         "A unique token that guarantees that the call to this API is idempotent.",
						MarkdownDescription: "A unique token that guarantees that the call to this API is idempotent.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"customer_metadata_properties": schema.MapAttribute{
						Description:         "The metadata properties associated with the model package versions.",
						MarkdownDescription: "The metadata properties associated with the model package versions.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"domain": schema.StringAttribute{
						Description:         "The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.",
						MarkdownDescription: "The machine learning domain of your model package and its components. Common machine learning domains include computer vision and natural language processing.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"drift_check_baselines": schema.SingleNestedAttribute{
						Description:         "Represents the drift check baselines that can be used when the model monitor is set using the model package. For more information, see the topic on Drift Detection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection) in the Amazon SageMaker Developer Guide.",
						MarkdownDescription: "Represents the drift check baselines that can be used when the model monitor is set using the model package. For more information, see the topic on Drift Detection against Previous Baselines in SageMaker Pipelines (https://docs.aws.amazon.com/sagemaker/latest/dg/pipelines-quality-clarify-baseline-lifecycle.html#pipelines-quality-clarify-baseline-drift-detection) in the Amazon SageMaker Developer Guide.",
						Attributes: map[string]schema.Attribute{
							"bias": schema.SingleNestedAttribute{
								Description:         "Represents the drift check bias baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check bias baselines that can be used when the model monitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"config_file": schema.SingleNestedAttribute{
										Description:         "Contains details regarding the file source.",
										MarkdownDescription: "Contains details regarding the file source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"post_training_constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"pre_training_constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

							"explainability": schema.SingleNestedAttribute{
								Description:         "Represents the drift check explainability baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check explainability baselines that can be used when the model monitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"config_file": schema.SingleNestedAttribute{
										Description:         "Contains details regarding the file source.",
										MarkdownDescription: "Contains details regarding the file source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

							"model_data_quality": schema.SingleNestedAttribute{
								Description:         "Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check data quality baselines that can be used when the model monitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

							"model_quality": schema.SingleNestedAttribute{
								Description:         "Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.",
								MarkdownDescription: "Represents the drift check model quality baselines that can be used when the model monitor is set using the model package.",
								Attributes: map[string]schema.Attribute{
									"constraints": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"inference_specification": schema.SingleNestedAttribute{
						Description:         "Specifies details about inference jobs that can be run with models based on this model package, including the following:  * The Amazon ECR paths of containers that contain the inference code and model artifacts.  * The instance types that the model package supports for transform jobs and real-time endpoints used for inference.  * The input and output content formats that the model package supports for inference.",
						MarkdownDescription: "Specifies details about inference jobs that can be run with models based on this model package, including the following:  * The Amazon ECR paths of containers that contain the inference code and model artifacts.  * The instance types that the model package supports for transform jobs and real-time endpoints used for inference.  * The input and output content formats that the model package supports for inference.",
						Attributes: map[string]schema.Attribute{
							"containers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"container_hostname": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"environment": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"framework": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"framework_version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image_digest": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"model_data_url": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"model_input": schema.SingleNestedAttribute{
											Description:         "Input object for the model.",
											MarkdownDescription: "Input object for the model.",
											Attributes: map[string]schema.Attribute{
												"data_input_config": schema.StringAttribute{
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

										"nearest_model_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"product_id": schema.StringAttribute{
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

							"supported_content_types": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"supported_realtime_inference_instance_types": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"supported_response_mime_types": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"supported_transform_instance_types": schema.ListAttribute{
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

					"metadata_properties": schema.SingleNestedAttribute{
						Description:         "Metadata properties of the tracking entity, trial, or trial component.",
						MarkdownDescription: "Metadata properties of the tracking entity, trial, or trial component.",
						Attributes: map[string]schema.Attribute{
							"commit_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"generated_by": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"project_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"repository": schema.StringAttribute{
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

					"model_approval_status": schema.StringAttribute{
						Description:         "Whether the model is approved for deployment.  This parameter is optional for versioned models, and does not apply to unversioned models.  For versioned models, the value of this parameter must be set to Approved to deploy the model.",
						MarkdownDescription: "Whether the model is approved for deployment.  This parameter is optional for versioned models, and does not apply to unversioned models.  For versioned models, the value of this parameter must be set to Approved to deploy the model.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"pre_training_report": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"report": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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

									"statistics": schema.SingleNestedAttribute{
										Description:         "Details about the metrics source.",
										MarkdownDescription: "Details about the metrics source.",
										Attributes: map[string]schema.Attribute{
											"content_digest": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"content_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"s3_uri": schema.StringAttribute{
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"model_package_description": schema.StringAttribute{
						Description:         "A description of the model package.",
						MarkdownDescription: "A description of the model package.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"model_package_group_name": schema.StringAttribute{
						Description:         "The name or Amazon Resource Name (ARN) of the model package group that this model version belongs to.  This parameter is required for versioned models, and does not apply to unversioned models.",
						MarkdownDescription: "The name or Amazon Resource Name (ARN) of the model package group that this model version belongs to.  This parameter is required for versioned models, and does not apply to unversioned models.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"model_package_name": schema.StringAttribute{
						Description:         "The name of the model package. The name must have 1 to 63 characters. Valid characters are a-z, A-Z, 0-9, and - (hyphen).  This parameter is required for unversioned models. It is not applicable to versioned models.",
						MarkdownDescription: "The name of the model package. The name must have 1 to 63 characters. Valid characters are a-z, A-Z, 0-9, and - (hyphen).  This parameter is required for unversioned models. It is not applicable to versioned models.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"sample_payload_url": schema.StringAttribute{
						Description:         "The Amazon Simple Storage Service (Amazon S3) path where the sample payload is stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix). This archive can hold multiple files that are all equally used in the load test. Each file in the archive must satisfy the size constraints of the InvokeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html#API_runtime_InvokeEndpoint_RequestSyntax) call.",
						MarkdownDescription: "The Amazon Simple Storage Service (Amazon S3) path where the sample payload is stored. This path must point to a single gzip compressed tar archive (.tar.gz suffix). This archive can hold multiple files that are all equally used in the load test. Each file in the archive must satisfy the size constraints of the InvokeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html#API_runtime_InvokeEndpoint_RequestSyntax) call.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
											Optional:            false,
											Computed:            true,
										},

										"model_data_url": schema.StringAttribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "A list of key value pairs associated with the model. For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon Web Services General Reference Guide.",
						MarkdownDescription: "A list of key value pairs associated with the model. For more information, see Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon Web Services General Reference Guide.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
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

					"task": schema.StringAttribute{
						Description:         "The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification. The following tasks are supported by Inference Recommender: 'IMAGE_CLASSIFICATION' | 'OBJECT_DETECTION' | 'TEXT_GENERATION' |'IMAGE_SEGMENTATION' | 'FILL_MASK' | 'CLASSIFICATION' | 'REGRESSION' | 'OTHER'.  Specify 'OTHER' if none of the tasks listed fit your use case.",
						MarkdownDescription: "The machine learning task your model package accomplishes. Common machine learning tasks include object detection and image classification. The following tasks are supported by Inference Recommender: 'IMAGE_CLASSIFICATION' | 'OBJECT_DETECTION' | 'TEXT_GENERATION' |'IMAGE_SEGMENTATION' | 'FILL_MASK' | 'CLASSIFICATION' | 'REGRESSION' | 'OTHER'.  Specify 'OTHER' if none of the tasks listed fit your use case.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"validation_specification": schema.SingleNestedAttribute{
						Description:         "Specifies configurations for one or more transform jobs that SageMaker runs to test the model package.",
						MarkdownDescription: "Specifies configurations for one or more transform jobs that SageMaker runs to test the model package.",
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
											Optional:            false,
											Computed:            true,
										},

										"transform_job_definition": schema.SingleNestedAttribute{
											Description:         "Defines the input needed to run a transform job using the inference specification specified in the algorithm.",
											MarkdownDescription: "Defines the input needed to run a transform job using the inference specification specified in the algorithm.",
											Attributes: map[string]schema.Attribute{
												"batch_strategy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"environment": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_concurrent_transforms": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_payload_in_mb": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"transform_input": schema.SingleNestedAttribute{
													Description:         "Describes the input source of a transform job and the way the transform job consumes it.",
													MarkdownDescription: "Describes the input source of a transform job and the way the transform job consumes it.",
													Attributes: map[string]schema.Attribute{
														"compression_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"content_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
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
																			Optional:            false,
																			Computed:            true,
																		},

																		"s3_uri": schema.StringAttribute{
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

														"split_type": schema.StringAttribute{
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

												"transform_output": schema.SingleNestedAttribute{
													Description:         "Describes the results of a transform job.",
													MarkdownDescription: "Describes the results of a transform job.",
													Attributes: map[string]schema.Attribute{
														"accept": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"assemble_with": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"kms_key_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"s3_output_path": schema.StringAttribute{
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

												"transform_resources": schema.SingleNestedAttribute{
													Description:         "Describes the resources, including ML instance types and ML instance count, to use for transform job.",
													MarkdownDescription: "Describes the resources, including ML instance types and ML instance count, to use for transform job.",
													Attributes: map[string]schema.Attribute{
														"instance_count": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"instance_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_kms_key_id": schema.StringAttribute{
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
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"validation_role": schema.StringAttribute{
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
		},
	}
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *SagemakerServicesK8SAwsModelPackageV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_model_package_v1alpha1")

	var data SagemakerServicesK8SAwsModelPackageV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "ModelPackage"}).
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

	var readResponse SagemakerServicesK8SAwsModelPackageV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("ModelPackage")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
