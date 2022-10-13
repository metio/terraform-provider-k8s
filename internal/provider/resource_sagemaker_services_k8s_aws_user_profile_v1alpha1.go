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

type SagemakerServicesK8SAwsUserProfileV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsUserProfileV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsUserProfileV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsUserProfileV1Alpha1GoModel struct {
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
		DomainID *string `tfsdk:"domain_id" yaml:"domainID,omitempty"`

		SingleSignOnUserIdentifier *string `tfsdk:"single_sign_on_user_identifier" yaml:"singleSignOnUserIdentifier,omitempty"`

		SingleSignOnUserValue *string `tfsdk:"single_sign_on_user_value" yaml:"singleSignOnUserValue,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		UserProfileName *string `tfsdk:"user_profile_name" yaml:"userProfileName,omitempty"`

		UserSettings *struct {
			ExecutionRole *string `tfsdk:"execution_role" yaml:"executionRole,omitempty"`

			JupyterServerAppSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

					LifecycleConfigARN *string `tfsdk:"lifecycle_config_arn" yaml:"lifecycleConfigARN,omitempty"`

					SageMakerImageARN *string `tfsdk:"sage_maker_image_arn" yaml:"sageMakerImageARN,omitempty"`

					SageMakerImageVersionARN *string `tfsdk:"sage_maker_image_version_arn" yaml:"sageMakerImageVersionARN,omitempty"`
				} `tfsdk:"default_resource_spec" yaml:"defaultResourceSpec,omitempty"`

				LifecycleConfigARNs *[]string `tfsdk:"lifecycle_config_ar_ns" yaml:"lifecycleConfigARNs,omitempty"`
			} `tfsdk:"jupyter_server_app_settings" yaml:"jupyterServerAppSettings,omitempty"`

			KernelGatewayAppSettings *struct {
				CustomImages *[]struct {
					AppImageConfigName *string `tfsdk:"app_image_config_name" yaml:"appImageConfigName,omitempty"`

					ImageName *string `tfsdk:"image_name" yaml:"imageName,omitempty"`

					ImageVersionNumber *int64 `tfsdk:"image_version_number" yaml:"imageVersionNumber,omitempty"`
				} `tfsdk:"custom_images" yaml:"customImages,omitempty"`

				DefaultResourceSpec *struct {
					InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

					LifecycleConfigARN *string `tfsdk:"lifecycle_config_arn" yaml:"lifecycleConfigARN,omitempty"`

					SageMakerImageARN *string `tfsdk:"sage_maker_image_arn" yaml:"sageMakerImageARN,omitempty"`

					SageMakerImageVersionARN *string `tfsdk:"sage_maker_image_version_arn" yaml:"sageMakerImageVersionARN,omitempty"`
				} `tfsdk:"default_resource_spec" yaml:"defaultResourceSpec,omitempty"`

				LifecycleConfigARNs *[]string `tfsdk:"lifecycle_config_ar_ns" yaml:"lifecycleConfigARNs,omitempty"`
			} `tfsdk:"kernel_gateway_app_settings" yaml:"kernelGatewayAppSettings,omitempty"`

			RStudioServerProAppSettings *struct {
				AccessStatus *string `tfsdk:"access_status" yaml:"accessStatus,omitempty"`

				UserGroup *string `tfsdk:"user_group" yaml:"userGroup,omitempty"`
			} `tfsdk:"r_studio_server_pro_app_settings" yaml:"rStudioServerProAppSettings,omitempty"`

			SecurityGroups *[]string `tfsdk:"security_groups" yaml:"securityGroups,omitempty"`

			SharingSettings *struct {
				NotebookOutputOption *string `tfsdk:"notebook_output_option" yaml:"notebookOutputOption,omitempty"`

				S3KMSKeyID *string `tfsdk:"s3_kms_key_id" yaml:"s3KMSKeyID,omitempty"`

				S3OutputPath *string `tfsdk:"s3_output_path" yaml:"s3OutputPath,omitempty"`
			} `tfsdk:"sharing_settings" yaml:"sharingSettings,omitempty"`

			TensorBoardAppSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

					LifecycleConfigARN *string `tfsdk:"lifecycle_config_arn" yaml:"lifecycleConfigARN,omitempty"`

					SageMakerImageARN *string `tfsdk:"sage_maker_image_arn" yaml:"sageMakerImageARN,omitempty"`

					SageMakerImageVersionARN *string `tfsdk:"sage_maker_image_version_arn" yaml:"sageMakerImageVersionARN,omitempty"`
				} `tfsdk:"default_resource_spec" yaml:"defaultResourceSpec,omitempty"`
			} `tfsdk:"tensor_board_app_settings" yaml:"tensorBoardAppSettings,omitempty"`
		} `tfsdk:"user_settings" yaml:"userSettings,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsUserProfileV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsUserProfileV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_user_profile_v1alpha1"
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "UserProfile is the Schema for the UserProfiles API",
		MarkdownDescription: "UserProfile is the Schema for the UserProfiles API",
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
				Description:         "UserProfileSpec defines the desired state of UserProfile.",
				MarkdownDescription: "UserProfileSpec defines the desired state of UserProfile.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"domain_id": {
						Description:         "The ID of the associated Domain.",
						MarkdownDescription: "The ID of the associated Domain.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"single_sign_on_user_identifier": {
						Description:         "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is 'UserName'. If the Domain's AuthMode is Amazon Web Services SSO, this field is required. If the Domain's AuthMode is not Amazon Web Services SSO, this field cannot be specified.",
						MarkdownDescription: "A specifier for the type of value specified in SingleSignOnUserValue. Currently, the only supported value is 'UserName'. If the Domain's AuthMode is Amazon Web Services SSO, this field is required. If the Domain's AuthMode is not Amazon Web Services SSO, this field cannot be specified.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"single_sign_on_user_value": {
						Description:         "The username of the associated Amazon Web Services Single Sign-On User for this UserProfile. If the Domain's AuthMode is Amazon Web Services SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not Amazon Web Services SSO, this field cannot be specified.",
						MarkdownDescription: "The username of the associated Amazon Web Services Single Sign-On User for this UserProfile. If the Domain's AuthMode is Amazon Web Services SSO, this field is required, and must match a valid username of a user in your directory. If the Domain's AuthMode is not Amazon Web Services SSO, this field cannot be specified.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "Each tag consists of a key and an optional value. Tag keys must be unique per resource.  Tags that you specify for the User Profile are also added to all Apps that the User Profile launches.",
						MarkdownDescription: "Each tag consists of a key and an optional value. Tag keys must be unique per resource.  Tags that you specify for the User Profile are also added to all Apps that the User Profile launches.",

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

					"user_profile_name": {
						Description:         "A name for the UserProfile. This value is not case sensitive.",
						MarkdownDescription: "A name for the UserProfile. This value is not case sensitive.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"user_settings": {
						Description:         "A collection of settings.",
						MarkdownDescription: "A collection of settings.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"execution_role": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jupyter_server_app_settings": {
								Description:         "The JupyterServer app settings.",
								MarkdownDescription: "The JupyterServer app settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default_resource_spec": {
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"instance_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lifecycle_config_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sage_maker_image_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sage_maker_image_version_arn": {
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

									"lifecycle_config_ar_ns": {
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

							"kernel_gateway_app_settings": {
								Description:         "The KernelGateway app settings.",
								MarkdownDescription: "The KernelGateway app settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"custom_images": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"app_image_config_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_version_number": {
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

									"default_resource_spec": {
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"instance_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lifecycle_config_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sage_maker_image_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sage_maker_image_version_arn": {
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

									"lifecycle_config_ar_ns": {
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

							"r_studio_server_pro_app_settings": {
								Description:         "A collection of settings that configure user interaction with the RStudioServerPro app. RStudioServerProAppSettings cannot be updated. The RStudioServerPro app must be deleted and a new one created to make any changes.",
								MarkdownDescription: "A collection of settings that configure user interaction with the RStudioServerPro app. RStudioServerProAppSettings cannot be updated. The RStudioServerPro app must be deleted and a new one created to make any changes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"access_status": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_group": {
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

							"security_groups": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sharing_settings": {
								Description:         "Specifies options for sharing SageMaker Studio notebooks. These settings are specified as part of DefaultUserSettings when the CreateDomain API is called, and as part of UserSettings when the CreateUserProfile API is called. When SharingSettings is not specified, notebook sharing isn't allowed.",
								MarkdownDescription: "Specifies options for sharing SageMaker Studio notebooks. These settings are specified as part of DefaultUserSettings when the CreateDomain API is called, and as part of UserSettings when the CreateUserProfile API is called. When SharingSettings is not specified, notebook sharing isn't allowed.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"notebook_output_option": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"s3_kms_key_id": {
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

							"tensor_board_app_settings": {
								Description:         "The TensorBoard app settings.",
								MarkdownDescription: "The TensorBoard app settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default_resource_spec": {
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, and the instance type that the version runs on.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"instance_type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lifecycle_config_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sage_maker_image_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sage_maker_image_version_arn": {
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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1")

	var state SagemakerServicesK8SAwsUserProfileV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsUserProfileV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("UserProfile")

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

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1")

	var state SagemakerServicesK8SAwsUserProfileV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsUserProfileV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("UserProfile")

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

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
