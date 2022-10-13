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

type SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource)(nil)
)

type SagemakerServicesK8SAwsFeatureGroupV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SagemakerServicesK8SAwsFeatureGroupV1Alpha1GoModel struct {
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
		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		EventTimeFeatureName *string `tfsdk:"event_time_feature_name" yaml:"eventTimeFeatureName,omitempty"`

		FeatureDefinitions *[]struct {
			FeatureName *string `tfsdk:"feature_name" yaml:"featureName,omitempty"`

			FeatureType *string `tfsdk:"feature_type" yaml:"featureType,omitempty"`
		} `tfsdk:"feature_definitions" yaml:"featureDefinitions,omitempty"`

		FeatureGroupName *string `tfsdk:"feature_group_name" yaml:"featureGroupName,omitempty"`

		OfflineStoreConfig *struct {
			DataCatalogConfig *struct {
				Catalog *string `tfsdk:"catalog" yaml:"catalog,omitempty"`

				Database *string `tfsdk:"database" yaml:"database,omitempty"`

				TableName *string `tfsdk:"table_name" yaml:"tableName,omitempty"`
			} `tfsdk:"data_catalog_config" yaml:"dataCatalogConfig,omitempty"`

			DisableGlueTableCreation *bool `tfsdk:"disable_glue_table_creation" yaml:"disableGlueTableCreation,omitempty"`

			S3StorageConfig *struct {
				KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

				ResolvedOutputS3URI *string `tfsdk:"resolved_output_s3_uri" yaml:"resolvedOutputS3URI,omitempty"`

				S3URI *string `tfsdk:"s3_uri" yaml:"s3URI,omitempty"`
			} `tfsdk:"s3_storage_config" yaml:"s3StorageConfig,omitempty"`
		} `tfsdk:"offline_store_config" yaml:"offlineStoreConfig,omitempty"`

		OnlineStoreConfig *struct {
			EnableOnlineStore *bool `tfsdk:"enable_online_store" yaml:"enableOnlineStore,omitempty"`

			SecurityConfig *struct {
				KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`
			} `tfsdk:"security_config" yaml:"securityConfig,omitempty"`
		} `tfsdk:"online_store_config" yaml:"onlineStoreConfig,omitempty"`

		RecordIdentifierFeatureName *string `tfsdk:"record_identifier_feature_name" yaml:"recordIdentifierFeatureName,omitempty"`

		RoleARN *string `tfsdk:"role_arn" yaml:"roleARN,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource() resource.Resource {
	return &SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource{}
}

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sagemaker_services_k8s_aws_feature_group_v1alpha1"
}

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "FeatureGroup is the Schema for the FeatureGroups API",
		MarkdownDescription: "FeatureGroup is the Schema for the FeatureGroups API",
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
				Description:         "FeatureGroupSpec defines the desired state of FeatureGroup.  Amazon SageMaker Feature Store stores features in a collection called Feature Group. A Feature Group can be visualized as a table which has rows, with a unique identifier for each row where each column in the table is a feature. In principle, a Feature Group is composed of features and values per features.",
				MarkdownDescription: "FeatureGroupSpec defines the desired state of FeatureGroup.  Amazon SageMaker Feature Store stores features in a collection called Feature Group. A Feature Group can be visualized as a table which has rows, with a unique identifier for each row where each column in the table is a feature. In principle, a Feature Group is composed of features and values per features.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"description": {
						Description:         "A free-form description of a FeatureGroup.",
						MarkdownDescription: "A free-form description of a FeatureGroup.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"event_time_feature_name": {
						Description:         "The name of the feature that stores the EventTime of a Record in a FeatureGroup.  An EventTime is a point in time when a new event occurs that corresponds to the creation or update of a Record in a FeatureGroup. All Records in the FeatureGroup must have a corresponding EventTime.  An EventTime can be a String or Fractional.  * Fractional: EventTime feature values must be a Unix timestamp in seconds.  * String: EventTime feature values must be an ISO-8601 string in the format. The following formats are supported yyyy-MM-dd'T'HH:mm:ssZ and yyyy-MM-dd'T'HH:mm:ss.SSSZ where yyyy, MM, and dd represent the year, month, and day respectively and HH, mm, ss, and if applicable, SSS represent the hour, month, second and milliseconds respsectively. 'T' and Z are constants.",
						MarkdownDescription: "The name of the feature that stores the EventTime of a Record in a FeatureGroup.  An EventTime is a point in time when a new event occurs that corresponds to the creation or update of a Record in a FeatureGroup. All Records in the FeatureGroup must have a corresponding EventTime.  An EventTime can be a String or Fractional.  * Fractional: EventTime feature values must be a Unix timestamp in seconds.  * String: EventTime feature values must be an ISO-8601 string in the format. The following formats are supported yyyy-MM-dd'T'HH:mm:ssZ and yyyy-MM-dd'T'HH:mm:ss.SSSZ where yyyy, MM, and dd represent the year, month, and day respectively and HH, mm, ss, and if applicable, SSS represent the hour, month, second and milliseconds respsectively. 'T' and Z are constants.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"feature_definitions": {
						Description:         "A list of Feature names and types. Name and Type is compulsory per Feature.  Valid feature FeatureTypes are Integral, Fractional and String.  FeatureNames cannot be any of the following: is_deleted, write_time, api_invocation_time  You can create up to 2,500 FeatureDefinitions per FeatureGroup.",
						MarkdownDescription: "A list of Feature names and types. Name and Type is compulsory per Feature.  Valid feature FeatureTypes are Integral, Fractional and String.  FeatureNames cannot be any of the following: is_deleted, write_time, api_invocation_time  You can create up to 2,500 FeatureDefinitions per FeatureGroup.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"feature_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"feature_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"feature_group_name": {
						Description:         "The name of the FeatureGroup. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account. The name:  * Must start and end with an alphanumeric character.  * Can only contain alphanumeric character and hyphens. Spaces are not allowed.",
						MarkdownDescription: "The name of the FeatureGroup. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account. The name:  * Must start and end with an alphanumeric character.  * Can only contain alphanumeric character and hyphens. Spaces are not allowed.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"offline_store_config": {
						Description:         "Use this to configure an OfflineFeatureStore. This parameter allows you to specify:  * The Amazon Simple Storage Service (Amazon S3) location of an OfflineStore.  * A configuration for an Amazon Web Services Glue or Amazon Web Services Hive data catalog.  * An KMS encryption key to encrypt the Amazon S3 location used for OfflineStore. If KMS encryption key is not specified, by default we encrypt all data at rest using Amazon Web Services KMS key. By defining your bucket-level key (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucket-key.html) for SSE, you can reduce Amazon Web Services KMS requests costs by up to 99 percent.  To learn more about this parameter, see OfflineStoreConfig.",
						MarkdownDescription: "Use this to configure an OfflineFeatureStore. This parameter allows you to specify:  * The Amazon Simple Storage Service (Amazon S3) location of an OfflineStore.  * A configuration for an Amazon Web Services Glue or Amazon Web Services Hive data catalog.  * An KMS encryption key to encrypt the Amazon S3 location used for OfflineStore. If KMS encryption key is not specified, by default we encrypt all data at rest using Amazon Web Services KMS key. By defining your bucket-level key (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucket-key.html) for SSE, you can reduce Amazon Web Services KMS requests costs by up to 99 percent.  To learn more about this parameter, see OfflineStoreConfig.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"data_catalog_config": {
								Description:         "The meta data of the Glue table which serves as data catalog for the OfflineStore.",
								MarkdownDescription: "The meta data of the Glue table which serves as data catalog for the OfflineStore.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"catalog": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"database": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"table_name": {
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

							"disable_glue_table_creation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"s3_storage_config": {
								Description:         "The Amazon Simple Storage (Amazon S3) location and and security configuration for OfflineStore.",
								MarkdownDescription: "The Amazon Simple Storage (Amazon S3) location and and security configuration for OfflineStore.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"kms_key_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resolved_output_s3_uri": {
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

					"online_store_config": {
						Description:         "You can turn the OnlineStore on or off by specifying True for the EnableOnlineStore flag in OnlineStoreConfig; the default value is False.  You can also include an Amazon Web Services KMS key ID (KMSKeyId) for at-rest encryption of the OnlineStore.",
						MarkdownDescription: "You can turn the OnlineStore on or off by specifying True for the EnableOnlineStore flag in OnlineStoreConfig; the default value is False.  You can also include an Amazon Web Services KMS key ID (KMSKeyId) for at-rest encryption of the OnlineStore.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable_online_store": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_config": {
								Description:         "The security configuration for OnlineStore.",
								MarkdownDescription: "The security configuration for OnlineStore.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"record_identifier_feature_name": {
						Description:         "The name of the Feature whose value uniquely identifies a Record defined in the FeatureStore. Only the latest record per identifier value will be stored in the OnlineStore. RecordIdentifierFeatureName must be one of feature definitions' names.  You use the RecordIdentifierFeatureName to access data in a FeatureStore.  This name:  * Must start and end with an alphanumeric character.  * Can only contains alphanumeric characters, hyphens, underscores. Spaces are not allowed.",
						MarkdownDescription: "The name of the Feature whose value uniquely identifies a Record defined in the FeatureStore. Only the latest record per identifier value will be stored in the OnlineStore. RecordIdentifierFeatureName must be one of feature definitions' names.  You use the RecordIdentifierFeatureName to access data in a FeatureStore.  This name:  * Must start and end with an alphanumeric character.  * Can only contains alphanumeric characters, hyphens, underscores. Spaces are not allowed.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"role_arn": {
						Description:         "The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the OfflineStore if an OfflineStoreConfig is provided.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the OfflineStore if an OfflineStoreConfig is provided.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "Tags used to identify Features in each FeatureGroup.",
						MarkdownDescription: "Tags used to identify Features in each FeatureGroup.",

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

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sagemaker_services_k8s_aws_feature_group_v1alpha1")

	var state SagemakerServicesK8SAwsFeatureGroupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsFeatureGroupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("FeatureGroup")

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

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_feature_group_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sagemaker_services_k8s_aws_feature_group_v1alpha1")

	var state SagemakerServicesK8SAwsFeatureGroupV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SagemakerServicesK8SAwsFeatureGroupV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sagemaker.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("FeatureGroup")

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

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sagemaker_services_k8s_aws_feature_group_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
