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
	_ datasource.DataSource              = &SagemakerServicesK8SAwsModelV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsModelV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsModelV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsModelV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsModelV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsModelV1Alpha1DataSourceData struct {
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
		Containers *[]struct {
			ContainerHostname *string            `tfsdk:"container_hostname" json:"containerHostname,omitempty"`
			Environment       *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
			Image             *string            `tfsdk:"image" json:"image,omitempty"`
			ImageConfig       *struct {
				RepositoryAccessMode *string `tfsdk:"repository_access_mode" json:"repositoryAccessMode,omitempty"`
				RepositoryAuthConfig *struct {
					RepositoryCredentialsProviderARN *string `tfsdk:"repository_credentials_provider_arn" json:"repositoryCredentialsProviderARN,omitempty"`
				} `tfsdk:"repository_auth_config" json:"repositoryAuthConfig,omitempty"`
			} `tfsdk:"image_config" json:"imageConfig,omitempty"`
			InferenceSpecificationName *string `tfsdk:"inference_specification_name" json:"inferenceSpecificationName,omitempty"`
			Mode                       *string `tfsdk:"mode" json:"mode,omitempty"`
			ModelDataURL               *string `tfsdk:"model_data_url" json:"modelDataURL,omitempty"`
			ModelPackageName           *string `tfsdk:"model_package_name" json:"modelPackageName,omitempty"`
			MultiModelConfig           *struct {
				ModelCacheSetting *string `tfsdk:"model_cache_setting" json:"modelCacheSetting,omitempty"`
			} `tfsdk:"multi_model_config" json:"multiModelConfig,omitempty"`
		} `tfsdk:"containers" json:"containers,omitempty"`
		EnableNetworkIsolation   *bool   `tfsdk:"enable_network_isolation" json:"enableNetworkIsolation,omitempty"`
		ExecutionRoleARN         *string `tfsdk:"execution_role_arn" json:"executionRoleARN,omitempty"`
		InferenceExecutionConfig *struct {
			Mode *string `tfsdk:"mode" json:"mode,omitempty"`
		} `tfsdk:"inference_execution_config" json:"inferenceExecutionConfig,omitempty"`
		ModelName        *string `tfsdk:"model_name" json:"modelName,omitempty"`
		PrimaryContainer *struct {
			ContainerHostname *string            `tfsdk:"container_hostname" json:"containerHostname,omitempty"`
			Environment       *map[string]string `tfsdk:"environment" json:"environment,omitempty"`
			Image             *string            `tfsdk:"image" json:"image,omitempty"`
			ImageConfig       *struct {
				RepositoryAccessMode *string `tfsdk:"repository_access_mode" json:"repositoryAccessMode,omitempty"`
				RepositoryAuthConfig *struct {
					RepositoryCredentialsProviderARN *string `tfsdk:"repository_credentials_provider_arn" json:"repositoryCredentialsProviderARN,omitempty"`
				} `tfsdk:"repository_auth_config" json:"repositoryAuthConfig,omitempty"`
			} `tfsdk:"image_config" json:"imageConfig,omitempty"`
			InferenceSpecificationName *string `tfsdk:"inference_specification_name" json:"inferenceSpecificationName,omitempty"`
			Mode                       *string `tfsdk:"mode" json:"mode,omitempty"`
			ModelDataURL               *string `tfsdk:"model_data_url" json:"modelDataURL,omitempty"`
			ModelPackageName           *string `tfsdk:"model_package_name" json:"modelPackageName,omitempty"`
			MultiModelConfig           *struct {
				ModelCacheSetting *string `tfsdk:"model_cache_setting" json:"modelCacheSetting,omitempty"`
			} `tfsdk:"multi_model_config" json:"multiModelConfig,omitempty"`
		} `tfsdk:"primary_container" json:"primaryContainer,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcConfig *struct {
			SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
			Subnets          *[]string `tfsdk:"subnets" json:"subnets,omitempty"`
		} `tfsdk:"vpc_config" json:"vpcConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsModelV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_model_v1alpha1"
}

func (r *SagemakerServicesK8SAwsModelV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Model is the Schema for the Models API",
		MarkdownDescription: "Model is the Schema for the Models API",
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
				Description:         "ModelSpec defines the desired state of Model.  The properties of a model as returned by the Search API.",
				MarkdownDescription: "ModelSpec defines the desired state of Model.  The properties of a model as returned by the Search API.",
				Attributes: map[string]schema.Attribute{
					"containers": schema.ListNestedAttribute{
						Description:         "Specifies the containers in the inference pipeline.",
						MarkdownDescription: "Specifies the containers in the inference pipeline.",
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

								"image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"image_config": schema.SingleNestedAttribute{
									Description:         "Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).",
									MarkdownDescription: "Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).",
									Attributes: map[string]schema.Attribute{
										"repository_access_mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"repository_auth_config": schema.SingleNestedAttribute{
											Description:         "Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field of the ImageConfig object that you passed to a call to CreateModel and the private Docker registry where the model image is hosted requires authentication.",
											MarkdownDescription: "Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field of the ImageConfig object that you passed to a call to CreateModel and the private Docker registry where the model image is hosted requires authentication.",
											Attributes: map[string]schema.Attribute{
												"repository_credentials_provider_arn": schema.StringAttribute{
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

								"inference_specification_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"mode": schema.StringAttribute{
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

								"model_package_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"multi_model_config": schema.SingleNestedAttribute{
									Description:         "Specifies additional configuration for hosting multi-model endpoints.",
									MarkdownDescription: "Specifies additional configuration for hosting multi-model endpoints.",
									Attributes: map[string]schema.Attribute{
										"model_cache_setting": schema.StringAttribute{
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

					"enable_network_isolation": schema.BoolAttribute{
						Description:         "Isolates the model container. No inbound or outbound network calls can be made to or from the model container.",
						MarkdownDescription: "Isolates the model container. No inbound or outbound network calls can be made to or from the model container.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"execution_role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to access model artifacts and docker image for deployment on ML compute instances or for batch transform jobs. Deploying on ML compute instances is part of model hosting. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).  To be able to pass this role to SageMaker, the caller of this API must have the iam:PassRole permission.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to access model artifacts and docker image for deployment on ML compute instances or for batch transform jobs. Deploying on ML compute instances is part of model hosting. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html).  To be able to pass this role to SageMaker, the caller of this API must have the iam:PassRole permission.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"inference_execution_config": schema.SingleNestedAttribute{
						Description:         "Specifies details of how containers in a multi-container endpoint are called.",
						MarkdownDescription: "Specifies details of how containers in a multi-container endpoint are called.",
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

					"model_name": schema.StringAttribute{
						Description:         "The name of the new model.",
						MarkdownDescription: "The name of the new model.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"primary_container": schema.SingleNestedAttribute{
						Description:         "The location of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.",
						MarkdownDescription: "The location of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.",
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

							"image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_config": schema.SingleNestedAttribute{
								Description:         "Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).",
								MarkdownDescription: "Specifies whether the model container is in Amazon ECR or a private Docker registry accessible from your Amazon Virtual Private Cloud (VPC).",
								Attributes: map[string]schema.Attribute{
									"repository_access_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"repository_auth_config": schema.SingleNestedAttribute{
										Description:         "Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field of the ImageConfig object that you passed to a call to CreateModel and the private Docker registry where the model image is hosted requires authentication.",
										MarkdownDescription: "Specifies an authentication configuration for the private docker registry where your model image is hosted. Specify a value for this property only if you specified Vpc as the value for the RepositoryAccessMode field of the ImageConfig object that you passed to a call to CreateModel and the private Docker registry where the model image is hosted requires authentication.",
										Attributes: map[string]schema.Attribute{
											"repository_credentials_provider_arn": schema.StringAttribute{
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

							"inference_specification_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mode": schema.StringAttribute{
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

							"model_package_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"multi_model_config": schema.SingleNestedAttribute{
								Description:         "Specifies additional configuration for hosting multi-model endpoints.",
								MarkdownDescription: "Specifies additional configuration for hosting multi-model endpoints.",
								Attributes: map[string]schema.Attribute{
									"model_cache_setting": schema.StringAttribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
						MarkdownDescription: "An array of key-value pairs. You can use tags to categorize your Amazon Web Services resources in different ways, for example, by purpose, owner, or environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html).",
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

					"vpc_config": schema.SingleNestedAttribute{
						Description:         "A VpcConfig object that specifies the VPC that you want your model to connect to. Control access to and from your model container by configuring the VPC. VpcConfig is used in hosting services and in batch transform. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html).",
						MarkdownDescription: "A VpcConfig object that specifies the VPC that you want your model to connect to. Control access to and from your model container by configuring the VPC. VpcConfig is used in hosting services and in batch transform. For more information, see Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html) and Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html).",
						Attributes: map[string]schema.Attribute{
							"security_group_i_ds": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"subnets": schema.ListAttribute{
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SagemakerServicesK8SAwsModelV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SagemakerServicesK8SAwsModelV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_model_v1alpha1")

	var data SagemakerServicesK8SAwsModelV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "models"}).
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

	var readResponse SagemakerServicesK8SAwsModelV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("Model")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
