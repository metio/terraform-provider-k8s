/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsUserProfileV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsUserProfileV1Alpha1DataSourceData struct {
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
		DomainID                   *string `tfsdk:"domain_id" json:"domainID,omitempty"`
		SingleSignOnUserIdentifier *string `tfsdk:"single_sign_on_user_identifier" json:"singleSignOnUserIdentifier,omitempty"`
		SingleSignOnUserValue      *string `tfsdk:"single_sign_on_user_value" json:"singleSignOnUserValue,omitempty"`
		Tags                       *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		UserProfileName *string `tfsdk:"user_profile_name" json:"userProfileName,omitempty"`
		UserSettings    *struct {
			ExecutionRole            *string `tfsdk:"execution_role" json:"executionRole,omitempty"`
			JupyterServerAppSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType             *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN       *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN        *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
				} `tfsdk:"default_resource_spec" json:"defaultResourceSpec,omitempty"`
				LifecycleConfigARNs *[]string `tfsdk:"lifecycle_config_ar_ns" json:"lifecycleConfigARNs,omitempty"`
			} `tfsdk:"jupyter_server_app_settings" json:"jupyterServerAppSettings,omitempty"`
			KernelGatewayAppSettings *struct {
				CustomImages *[]struct {
					AppImageConfigName *string `tfsdk:"app_image_config_name" json:"appImageConfigName,omitempty"`
					ImageName          *string `tfsdk:"image_name" json:"imageName,omitempty"`
					ImageVersionNumber *int64  `tfsdk:"image_version_number" json:"imageVersionNumber,omitempty"`
				} `tfsdk:"custom_images" json:"customImages,omitempty"`
				DefaultResourceSpec *struct {
					InstanceType             *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN       *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN        *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
				} `tfsdk:"default_resource_spec" json:"defaultResourceSpec,omitempty"`
				LifecycleConfigARNs *[]string `tfsdk:"lifecycle_config_ar_ns" json:"lifecycleConfigARNs,omitempty"`
			} `tfsdk:"kernel_gateway_app_settings" json:"kernelGatewayAppSettings,omitempty"`
			RStudioServerProAppSettings *struct {
				AccessStatus *string `tfsdk:"access_status" json:"accessStatus,omitempty"`
				UserGroup    *string `tfsdk:"user_group" json:"userGroup,omitempty"`
			} `tfsdk:"r_studio_server_pro_app_settings" json:"rStudioServerProAppSettings,omitempty"`
			SecurityGroups  *[]string `tfsdk:"security_groups" json:"securityGroups,omitempty"`
			SharingSettings *struct {
				NotebookOutputOption *string `tfsdk:"notebook_output_option" json:"notebookOutputOption,omitempty"`
				S3KMSKeyID           *string `tfsdk:"s3_kms_key_id" json:"s3KMSKeyID,omitempty"`
				S3OutputPath         *string `tfsdk:"s3_output_path" json:"s3OutputPath,omitempty"`
			} `tfsdk:"sharing_settings" json:"sharingSettings,omitempty"`
			TensorBoardAppSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType             *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN       *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN        *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
				} `tfsdk:"default_resource_spec" json:"defaultResourceSpec,omitempty"`
			} `tfsdk:"tensor_board_app_settings" json:"tensorBoardAppSettings,omitempty"`
		} `tfsdk:"user_settings" json:"userSettings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_user_profile_v1alpha1"
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "UserProfile is the Schema for the UserProfiles API",
		MarkdownDescription: "UserProfile is the Schema for the UserProfiles API",
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
				Description:         "UserProfileSpec defines the desired state of UserProfile.",
				MarkdownDescription: "UserProfileSpec defines the desired state of UserProfile.",
				Attributes: map[string]schema.Attribute{
					"domain_id": schema.StringAttribute{
						Description:         "The ID of the associated Domain.",
						MarkdownDescription: "The ID of the associated Domain.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"single_sign_on_user_identifier": schema.StringAttribute{
						Description:         "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is 'UserName'. If the Domain's AuthMode is IAM Identity Center, this field is required. If the Domain's AuthMode is not IAM Identity Center, this field cannot be specified.",
						MarkdownDescription: "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is 'UserName'. If the Domain's AuthMode is IAM Identity Center, this field is required. If the Domain's AuthMode is not IAM Identity Center, this field cannot be specified.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"single_sign_on_user_value": schema.StringAttribute{
						Description:         "The username of the associated Amazon Web Services Single Sign-On User for this UserProfile. If the Domain's AuthMode is IAM Identity Center, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not IAM Identity Center, this field cannot be specified.",
						MarkdownDescription: "The username of the associated Amazon Web Services Single Sign-On User for this UserProfile. If the Domain's AuthMode is IAM Identity Center, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not IAM Identity Center, this field cannot be specified.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Each tag consists of a key and an optional value. Tag keys must be unique per resource.  Tags that you specify for the User Profile are also added to all Apps that the User Profile launches.",
						MarkdownDescription: "Each tag consists of a key and an optional value. Tag keys must be unique per resource.  Tags that you specify for the User Profile are also added to all Apps that the User Profile launches.",
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

					"user_profile_name": schema.StringAttribute{
						Description:         "A name for the UserProfile. This value is not case sensitive.",
						MarkdownDescription: "A name for the UserProfile. This value is not case sensitive.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"user_settings": schema.SingleNestedAttribute{
						Description:         "A collection of settings.",
						MarkdownDescription: "A collection of settings.",
						Attributes: map[string]schema.Attribute{
							"execution_role": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jupyter_server_app_settings": schema.SingleNestedAttribute{
								Description:         "The JupyterServer app settings.",
								MarkdownDescription: "The JupyterServer app settings.",
								Attributes: map[string]schema.Attribute{
									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
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

									"lifecycle_config_ar_ns": schema.ListAttribute{
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

							"kernel_gateway_app_settings": schema.SingleNestedAttribute{
								Description:         "The KernelGateway app settings.",
								MarkdownDescription: "The KernelGateway app settings.",
								Attributes: map[string]schema.Attribute{
									"custom_images": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"app_image_config_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"image_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"image_version_number": schema.Int64Attribute{
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

									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
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

									"lifecycle_config_ar_ns": schema.ListAttribute{
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

							"r_studio_server_pro_app_settings": schema.SingleNestedAttribute{
								Description:         "A collection of settings that configure user interaction with the RStudioServerPro app. RStudioServerProAppSettings cannot be updated. The RStudioServerPro app must be deleted and a new one created to make any changes.",
								MarkdownDescription: "A collection of settings that configure user interaction with the RStudioServerPro app. RStudioServerProAppSettings cannot be updated. The RStudioServerPro app must be deleted and a new one created to make any changes.",
								Attributes: map[string]schema.Attribute{
									"access_status": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"user_group": schema.StringAttribute{
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

							"security_groups": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sharing_settings": schema.SingleNestedAttribute{
								Description:         "Specifies options for sharing SageMaker Studio notebooks. These settings are specified as part of DefaultUserSettings when the CreateDomain API is called, and as part of UserSettings when the CreateUserProfile API is called. When SharingSettings is not specified, notebook sharing isn't allowed.",
								MarkdownDescription: "Specifies options for sharing SageMaker Studio notebooks. These settings are specified as part of DefaultUserSettings when the CreateDomain API is called, and as part of UserSettings when the CreateUserProfile API is called. When SharingSettings is not specified, notebook sharing isn't allowed.",
								Attributes: map[string]schema.Attribute{
									"notebook_output_option": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"s3_kms_key_id": schema.StringAttribute{
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

							"tensor_board_app_settings": schema.SingleNestedAttribute{
								Description:         "The TensorBoard app settings.",
								MarkdownDescription: "The TensorBoard app settings.",
								Attributes: map[string]schema.Attribute{
									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1")

	var data SagemakerServicesK8SAwsUserProfileV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "userprofiles"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse SagemakerServicesK8SAwsUserProfileV1Alpha1DataSourceData
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
	data.Kind = pointer.String("UserProfile")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
