/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sagemaker_services_k8s_aws_v1alpha1

import (
	"context"
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
	_ datasource.DataSource = &SagemakerServicesK8SAwsUserProfileV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsUserProfileV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsUserProfileV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsUserProfileV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsUserProfileV1Alpha1ManifestData struct {
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
		DomainID                   *string `tfsdk:"domain_id" json:"domainID,omitempty"`
		SingleSignOnUserIdentifier *string `tfsdk:"single_sign_on_user_identifier" json:"singleSignOnUserIdentifier,omitempty"`
		SingleSignOnUserValue      *string `tfsdk:"single_sign_on_user_value" json:"singleSignOnUserValue,omitempty"`
		Tags                       *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		UserProfileName *string `tfsdk:"user_profile_name" json:"userProfileName,omitempty"`
		UserSettings    *struct {
			CodeEditorAppSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType               *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN         *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN          *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN   *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
					SageMakerImageVersionAlias *string `tfsdk:"sage_maker_image_version_alias" json:"sageMakerImageVersionAlias,omitempty"`
				} `tfsdk:"default_resource_spec" json:"defaultResourceSpec,omitempty"`
				LifecycleConfigARNs *[]string `tfsdk:"lifecycle_config_ar_ns" json:"lifecycleConfigARNs,omitempty"`
			} `tfsdk:"code_editor_app_settings" json:"codeEditorAppSettings,omitempty"`
			CustomFileSystemConfigs *[]struct {
				EfsFileSystemConfig *struct {
					FileSystemID   *string `tfsdk:"file_system_id" json:"fileSystemID,omitempty"`
					FileSystemPath *string `tfsdk:"file_system_path" json:"fileSystemPath,omitempty"`
				} `tfsdk:"efs_file_system_config" json:"efsFileSystemConfig,omitempty"`
			} `tfsdk:"custom_file_system_configs" json:"customFileSystemConfigs,omitempty"`
			CustomPosixUserConfig *struct {
				Gid *int64 `tfsdk:"gid" json:"gid,omitempty"`
				Uid *int64 `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"custom_posix_user_config" json:"customPosixUserConfig,omitempty"`
			DefaultLandingURI     *string `tfsdk:"default_landing_uri" json:"defaultLandingURI,omitempty"`
			ExecutionRole         *string `tfsdk:"execution_role" json:"executionRole,omitempty"`
			JupyterLabAppSettings *struct {
				CustomImages *[]struct {
					AppImageConfigName *string `tfsdk:"app_image_config_name" json:"appImageConfigName,omitempty"`
					ImageName          *string `tfsdk:"image_name" json:"imageName,omitempty"`
					ImageVersionNumber *int64  `tfsdk:"image_version_number" json:"imageVersionNumber,omitempty"`
				} `tfsdk:"custom_images" json:"customImages,omitempty"`
				DefaultResourceSpec *struct {
					InstanceType               *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN         *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN          *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN   *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
					SageMakerImageVersionAlias *string `tfsdk:"sage_maker_image_version_alias" json:"sageMakerImageVersionAlias,omitempty"`
				} `tfsdk:"default_resource_spec" json:"defaultResourceSpec,omitempty"`
				LifecycleConfigARNs *[]string `tfsdk:"lifecycle_config_ar_ns" json:"lifecycleConfigARNs,omitempty"`
			} `tfsdk:"jupyter_lab_app_settings" json:"jupyterLabAppSettings,omitempty"`
			JupyterServerAppSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType               *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN         *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN          *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN   *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
					SageMakerImageVersionAlias *string `tfsdk:"sage_maker_image_version_alias" json:"sageMakerImageVersionAlias,omitempty"`
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
					InstanceType               *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN         *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN          *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN   *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
					SageMakerImageVersionAlias *string `tfsdk:"sage_maker_image_version_alias" json:"sageMakerImageVersionAlias,omitempty"`
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
			SpaceStorageSettings *struct {
				DefaultEBSStorageSettings *struct {
					DefaultEBSVolumeSizeInGb *int64 `tfsdk:"default_ebs_volume_size_in_gb" json:"defaultEBSVolumeSizeInGb,omitempty"`
					MaximumEBSVolumeSizeInGb *int64 `tfsdk:"maximum_ebs_volume_size_in_gb" json:"maximumEBSVolumeSizeInGb,omitempty"`
				} `tfsdk:"default_ebs_storage_settings" json:"defaultEBSStorageSettings,omitempty"`
			} `tfsdk:"space_storage_settings" json:"spaceStorageSettings,omitempty"`
			StudioWebPortal        *string `tfsdk:"studio_web_portal" json:"studioWebPortal,omitempty"`
			TensorBoardAppSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType               *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN         *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN          *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN   *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
					SageMakerImageVersionAlias *string `tfsdk:"sage_maker_image_version_alias" json:"sageMakerImageVersionAlias,omitempty"`
				} `tfsdk:"default_resource_spec" json:"defaultResourceSpec,omitempty"`
			} `tfsdk:"tensor_board_app_settings" json:"tensorBoardAppSettings,omitempty"`
		} `tfsdk:"user_settings" json:"userSettings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "UserProfile is the Schema for the UserProfiles API",
		MarkdownDescription: "UserProfile is the Schema for the UserProfiles API",
		Attributes: map[string]schema.Attribute{
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
				Description:         "UserProfileSpec defines the desired state of UserProfile.",
				MarkdownDescription: "UserProfileSpec defines the desired state of UserProfile.",
				Attributes: map[string]schema.Attribute{
					"domain_id": schema.StringAttribute{
						Description:         "The ID of the associated Domain.",
						MarkdownDescription: "The ID of the associated Domain.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"single_sign_on_user_identifier": schema.StringAttribute{
						Description:         "A specifier for the type of value specified in SingleSignOnUserValue. Currently,the only supported value is 'UserName'. If the Domain's AuthMode is IAM IdentityCenter, this field is required. If the Domain's AuthMode is not IAM IdentityCenter, this field cannot be specified.",
						MarkdownDescription: "A specifier for the type of value specified in SingleSignOnUserValue. Currently,the only supported value is 'UserName'. If the Domain's AuthMode is IAM IdentityCenter, this field is required. If the Domain's AuthMode is not IAM IdentityCenter, this field cannot be specified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"single_sign_on_user_value": schema.StringAttribute{
						Description:         "The username of the associated Amazon Web Services Single Sign-On User forthis UserProfile. If the Domain's AuthMode is IAM Identity Center, this fieldis required, and must match a valid username of a user in your directory.If the Domain's AuthMode is not IAM Identity Center, this field cannot bespecified.",
						MarkdownDescription: "The username of the associated Amazon Web Services Single Sign-On User forthis UserProfile. If the Domain's AuthMode is IAM Identity Center, this fieldis required, and must match a valid username of a user in your directory.If the Domain's AuthMode is not IAM Identity Center, this field cannot bespecified.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Each tag consists of a key and an optional value. Tag keys must be uniqueper resource.Tags that you specify for the User Profile are also added to all Apps thatthe User Profile launches.",
						MarkdownDescription: "Each tag consists of a key and an optional value. Tag keys must be uniqueper resource.Tags that you specify for the User Profile are also added to all Apps thatthe User Profile launches.",
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

					"user_profile_name": schema.StringAttribute{
						Description:         "A name for the UserProfile. This value is not case sensitive.",
						MarkdownDescription: "A name for the UserProfile. This value is not case sensitive.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"user_settings": schema.SingleNestedAttribute{
						Description:         "A collection of settings.",
						MarkdownDescription: "A collection of settings.",
						Attributes: map[string]schema.Attribute{
							"code_editor_app_settings": schema.SingleNestedAttribute{
								Description:         "The Code Editor application settings.For more information about Code Editor, see Get started with Code Editorin Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/code-editor.html).",
								MarkdownDescription: "The Code Editor application settings.For more information about Code Editor, see Get started with Code Editorin Amazon SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/code-editor.html).",
								Attributes: map[string]schema.Attribute{
									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_alias": schema.StringAttribute{
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

									"lifecycle_config_ar_ns": schema.ListAttribute{
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

							"custom_file_system_configs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"efs_file_system_config": schema.SingleNestedAttribute{
											Description:         "The settings for assigning a custom Amazon EFS file system to a user profileor space for an Amazon SageMaker Domain.",
											MarkdownDescription: "The settings for assigning a custom Amazon EFS file system to a user profileor space for an Amazon SageMaker Domain.",
											Attributes: map[string]schema.Attribute{
												"file_system_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"file_system_path": schema.StringAttribute{
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

							"custom_posix_user_config": schema.SingleNestedAttribute{
								Description:         "Details about the POSIX identity that is used for file system operations.",
								MarkdownDescription: "Details about the POSIX identity that is used for file system operations.",
								Attributes: map[string]schema.Attribute{
									"gid": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"uid": schema.Int64Attribute{
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

							"default_landing_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"execution_role": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jupyter_lab_app_settings": schema.SingleNestedAttribute{
								Description:         "The settings for the JupyterLab application.",
								MarkdownDescription: "The settings for the JupyterLab application.",
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
													Optional:            true,
													Computed:            false,
												},

												"image_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_version_number": schema.Int64Attribute{
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

									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_alias": schema.StringAttribute{
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

									"lifecycle_config_ar_ns": schema.ListAttribute{
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

							"jupyter_server_app_settings": schema.SingleNestedAttribute{
								Description:         "The JupyterServer app settings.",
								MarkdownDescription: "The JupyterServer app settings.",
								Attributes: map[string]schema.Attribute{
									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_alias": schema.StringAttribute{
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

									"lifecycle_config_ar_ns": schema.ListAttribute{
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
													Optional:            true,
													Computed:            false,
												},

												"image_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_version_number": schema.Int64Attribute{
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

									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_alias": schema.StringAttribute{
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

									"lifecycle_config_ar_ns": schema.ListAttribute{
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

							"r_studio_server_pro_app_settings": schema.SingleNestedAttribute{
								Description:         "A collection of settings that configure user interaction with the RStudioServerProapp.",
								MarkdownDescription: "A collection of settings that configure user interaction with the RStudioServerProapp.",
								Attributes: map[string]schema.Attribute{
									"access_status": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_group": schema.StringAttribute{
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

							"security_groups": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sharing_settings": schema.SingleNestedAttribute{
								Description:         "Specifies options for sharing Amazon SageMaker Studio notebooks. These settingsare specified as part of DefaultUserSettings when the CreateDomain API iscalled, and as part of UserSettings when the CreateUserProfile API is called.When SharingSettings is not specified, notebook sharing isn't allowed.",
								MarkdownDescription: "Specifies options for sharing Amazon SageMaker Studio notebooks. These settingsare specified as part of DefaultUserSettings when the CreateDomain API iscalled, and as part of UserSettings when the CreateUserProfile API is called.When SharingSettings is not specified, notebook sharing isn't allowed.",
								Attributes: map[string]schema.Attribute{
									"notebook_output_option": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"s3_kms_key_id": schema.StringAttribute{
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

							"space_storage_settings": schema.SingleNestedAttribute{
								Description:         "The default storage settings for a private space.",
								MarkdownDescription: "The default storage settings for a private space.",
								Attributes: map[string]schema.Attribute{
									"default_ebs_storage_settings": schema.SingleNestedAttribute{
										Description:         "A collection of default EBS storage settings that applies to private spacescreated within a domain or user profile.",
										MarkdownDescription: "A collection of default EBS storage settings that applies to private spacescreated within a domain or user profile.",
										Attributes: map[string]schema.Attribute{
											"default_ebs_volume_size_in_gb": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"maximum_ebs_volume_size_in_gb": schema.Int64Attribute{
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

							"studio_web_portal": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tensor_board_app_settings": schema.SingleNestedAttribute{
								Description:         "The TensorBoard app settings.",
								MarkdownDescription: "The TensorBoard app settings.",
								Attributes: map[string]schema.Attribute{
									"default_resource_spec": schema.SingleNestedAttribute{
										Description:         "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										MarkdownDescription: "Specifies the ARN's of a SageMaker image and SageMaker image version, andthe instance type that the version runs on.",
										Attributes: map[string]schema.Attribute{
											"instance_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"lifecycle_config_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_arn": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sage_maker_image_version_alias": schema.StringAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SagemakerServicesK8SAwsUserProfileV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsUserProfileV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("UserProfile")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
