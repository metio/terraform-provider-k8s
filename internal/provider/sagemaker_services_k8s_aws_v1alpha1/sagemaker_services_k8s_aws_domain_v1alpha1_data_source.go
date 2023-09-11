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
	_ datasource.DataSource              = &SagemakerServicesK8SAwsDomainV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SagemakerServicesK8SAwsDomainV1Alpha1DataSource{}
)

func NewSagemakerServicesK8SAwsDomainV1Alpha1DataSource() datasource.DataSource {
	return &SagemakerServicesK8SAwsDomainV1Alpha1DataSource{}
}

type SagemakerServicesK8SAwsDomainV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SagemakerServicesK8SAwsDomainV1Alpha1DataSourceData struct {
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
		AppNetworkAccessType       *string `tfsdk:"app_network_access_type" json:"appNetworkAccessType,omitempty"`
		AppSecurityGroupManagement *string `tfsdk:"app_security_group_management" json:"appSecurityGroupManagement,omitempty"`
		AuthMode                   *string `tfsdk:"auth_mode" json:"authMode,omitempty"`
		DefaultUserSettings        *struct {
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
		} `tfsdk:"default_user_settings" json:"defaultUserSettings,omitempty"`
		DomainName     *string `tfsdk:"domain_name" json:"domainName,omitempty"`
		DomainSettings *struct {
			RStudioServerProDomainSettings *struct {
				DefaultResourceSpec *struct {
					InstanceType             *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
					LifecycleConfigARN       *string `tfsdk:"lifecycle_config_arn" json:"lifecycleConfigARN,omitempty"`
					SageMakerImageARN        *string `tfsdk:"sage_maker_image_arn" json:"sageMakerImageARN,omitempty"`
					SageMakerImageVersionARN *string `tfsdk:"sage_maker_image_version_arn" json:"sageMakerImageVersionARN,omitempty"`
				} `tfsdk:"default_resource_spec" json:"defaultResourceSpec,omitempty"`
				DomainExecutionRoleARN   *string `tfsdk:"domain_execution_role_arn" json:"domainExecutionRoleARN,omitempty"`
				RStudioConnectURL        *string `tfsdk:"r_studio_connect_url" json:"rStudioConnectURL,omitempty"`
				RStudioPackageManagerURL *string `tfsdk:"r_studio_package_manager_url" json:"rStudioPackageManagerURL,omitempty"`
			} `tfsdk:"r_studio_server_pro_domain_settings" json:"rStudioServerProDomainSettings,omitempty"`
			SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
		} `tfsdk:"domain_settings" json:"domainSettings,omitempty"`
		HomeEFSFileSystemKMSKeyID *string   `tfsdk:"home_efs_file_system_kms_key_id" json:"homeEFSFileSystemKMSKeyID,omitempty"`
		KmsKeyID                  *string   `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
		SubnetIDs                 *[]string `tfsdk:"subnet_i_ds" json:"subnetIDs,omitempty"`
		Tags                      *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcID *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsDomainV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_domain_v1alpha1"
}

func (r *SagemakerServicesK8SAwsDomainV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Domain is the Schema for the Domains API",
		MarkdownDescription: "Domain is the Schema for the Domains API",
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
				Description:         "DomainSpec defines the desired state of Domain.",
				MarkdownDescription: "DomainSpec defines the desired state of Domain.",
				Attributes: map[string]schema.Attribute{
					"app_network_access_type": schema.StringAttribute{
						Description:         "Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.  * PublicInternetOnly - Non-EFS traffic is through a VPC managed by Amazon SageMaker, which allows direct internet access  * VpcOnly - All Studio traffic is through the specified VPC and subnets",
						MarkdownDescription: "Specifies the VPC used for non-EFS traffic. The default value is PublicInternetOnly.  * PublicInternetOnly - Non-EFS traffic is through a VPC managed by Amazon SageMaker, which allows direct internet access  * VpcOnly - All Studio traffic is through the specified VPC and subnets",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"app_security_group_management": schema.StringAttribute{
						Description:         "The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.",
						MarkdownDescription: "The entity that creates and manages the required security groups for inter-app communication in VPCOnly mode. Required when CreateDomain.AppNetworkAccessType is VPCOnly and DomainSettings.RStudioServerProDomainSettings.DomainExecutionRoleArn is provided.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"auth_mode": schema.StringAttribute{
						Description:         "The mode of authentication that members use to access the domain.",
						MarkdownDescription: "The mode of authentication that members use to access the domain.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"default_user_settings": schema.SingleNestedAttribute{
						Description:         "The default settings to use to create a user profile when UserSettings isn't specified in the call to the CreateUserProfile API.  SecurityGroups is aggregated when specified in both calls. For all other settings in UserSettings, the values specified in CreateUserProfile take precedence over those specified in CreateDomain.",
						MarkdownDescription: "The default settings to use to create a user profile when UserSettings isn't specified in the call to the CreateUserProfile API.  SecurityGroups is aggregated when specified in both calls. For all other settings in UserSettings, the values specified in CreateUserProfile take precedence over those specified in CreateDomain.",
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

					"domain_name": schema.StringAttribute{
						Description:         "A name for the domain.",
						MarkdownDescription: "A name for the domain.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"domain_settings": schema.SingleNestedAttribute{
						Description:         "A collection of Domain settings.",
						MarkdownDescription: "A collection of Domain settings.",
						Attributes: map[string]schema.Attribute{
							"r_studio_server_pro_domain_settings": schema.SingleNestedAttribute{
								Description:         "A collection of settings that configure the RStudioServerPro Domain-level app.",
								MarkdownDescription: "A collection of settings that configure the RStudioServerPro Domain-level app.",
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

									"domain_execution_role_arn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"r_studio_connect_url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"r_studio_package_manager_url": schema.StringAttribute{
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

							"security_group_i_ds": schema.ListAttribute{
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

					"home_efs_file_system_kms_key_id": schema.StringAttribute{
						Description:         "Use KmsKeyId.",
						MarkdownDescription: "Use KmsKeyId.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"kms_key_id": schema.StringAttribute{
						Description:         "SageMaker uses Amazon Web Services KMS to encrypt the EFS volume attached to the domain with an Amazon Web Services managed key by default. For more control, specify a customer managed key.",
						MarkdownDescription: "SageMaker uses Amazon Web Services KMS to encrypt the EFS volume attached to the domain with an Amazon Web Services managed key by default. For more control, specify a customer managed key.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"subnet_i_ds": schema.ListAttribute{
						Description:         "The VPC subnets that Studio uses for communication.",
						MarkdownDescription: "The VPC subnets that Studio uses for communication.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Tags to associated with the Domain. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.  Tags that you specify for the Domain are also added to all Apps that the Domain launches.",
						MarkdownDescription: "Tags to associated with the Domain. Each tag consists of a key and an optional value. Tag keys must be unique per resource. Tags are searchable using the Search API.  Tags that you specify for the Domain are also added to all Apps that the Domain launches.",
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

					"vpc_id": schema.StringAttribute{
						Description:         "The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
						MarkdownDescription: "The ID of the Amazon Virtual Private Cloud (VPC) that Studio uses for communication.",
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
	}
}

func (r *SagemakerServicesK8SAwsDomainV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SagemakerServicesK8SAwsDomainV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_sagemaker_services_k8s_aws_domain_v1alpha1")

	var data SagemakerServicesK8SAwsDomainV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "sagemaker.services.k8s.aws", Version: "v1alpha1", Resource: "domains"}).
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

	var readResponse SagemakerServicesK8SAwsDomainV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("Domain")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
