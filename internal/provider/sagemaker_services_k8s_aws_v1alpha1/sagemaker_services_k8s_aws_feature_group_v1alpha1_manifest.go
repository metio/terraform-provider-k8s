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
	_ datasource.DataSource = &SagemakerServicesK8SAwsFeatureGroupV1Alpha1Manifest{}
)

func NewSagemakerServicesK8SAwsFeatureGroupV1Alpha1Manifest() datasource.DataSource {
	return &SagemakerServicesK8SAwsFeatureGroupV1Alpha1Manifest{}
}

type SagemakerServicesK8SAwsFeatureGroupV1Alpha1Manifest struct{}

type SagemakerServicesK8SAwsFeatureGroupV1Alpha1ManifestData struct {
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
		Description          *string `tfsdk:"description" json:"description,omitempty"`
		EventTimeFeatureName *string `tfsdk:"event_time_feature_name" json:"eventTimeFeatureName,omitempty"`
		FeatureDefinitions   *[]struct {
			FeatureName *string `tfsdk:"feature_name" json:"featureName,omitempty"`
			FeatureType *string `tfsdk:"feature_type" json:"featureType,omitempty"`
		} `tfsdk:"feature_definitions" json:"featureDefinitions,omitempty"`
		FeatureGroupName   *string `tfsdk:"feature_group_name" json:"featureGroupName,omitempty"`
		OfflineStoreConfig *struct {
			DataCatalogConfig *struct {
				Catalog   *string `tfsdk:"catalog" json:"catalog,omitempty"`
				Database  *string `tfsdk:"database" json:"database,omitempty"`
				TableName *string `tfsdk:"table_name" json:"tableName,omitempty"`
			} `tfsdk:"data_catalog_config" json:"dataCatalogConfig,omitempty"`
			DisableGlueTableCreation *bool `tfsdk:"disable_glue_table_creation" json:"disableGlueTableCreation,omitempty"`
			S3StorageConfig          *struct {
				KmsKeyID            *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
				ResolvedOutputS3URI *string `tfsdk:"resolved_output_s3_uri" json:"resolvedOutputS3URI,omitempty"`
				S3URI               *string `tfsdk:"s3_uri" json:"s3URI,omitempty"`
			} `tfsdk:"s3_storage_config" json:"s3StorageConfig,omitempty"`
		} `tfsdk:"offline_store_config" json:"offlineStoreConfig,omitempty"`
		OnlineStoreConfig *struct {
			EnableOnlineStore *bool `tfsdk:"enable_online_store" json:"enableOnlineStore,omitempty"`
			SecurityConfig    *struct {
				KmsKeyID *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
			} `tfsdk:"security_config" json:"securityConfig,omitempty"`
		} `tfsdk:"online_store_config" json:"onlineStoreConfig,omitempty"`
		RecordIdentifierFeatureName *string `tfsdk:"record_identifier_feature_name" json:"recordIdentifierFeatureName,omitempty"`
		RoleARN                     *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		Tags                        *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sagemaker_services_k8s_aws_feature_group_v1alpha1_manifest"
}

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FeatureGroup is the Schema for the FeatureGroups API",
		MarkdownDescription: "FeatureGroup is the Schema for the FeatureGroups API",
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
				Description:         "FeatureGroupSpec defines the desired state of FeatureGroup.  Amazon SageMaker Feature Store stores features in a collection called Feature Group. A Feature Group can be visualized as a table which has rows, with a unique identifier for each row where each column in the table is a feature. In principle, a Feature Group is composed of features and values per features.",
				MarkdownDescription: "FeatureGroupSpec defines the desired state of FeatureGroup.  Amazon SageMaker Feature Store stores features in a collection called Feature Group. A Feature Group can be visualized as a table which has rows, with a unique identifier for each row where each column in the table is a feature. In principle, a Feature Group is composed of features and values per features.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "A free-form description of a FeatureGroup.",
						MarkdownDescription: "A free-form description of a FeatureGroup.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"event_time_feature_name": schema.StringAttribute{
						Description:         "The name of the feature that stores the EventTime of a Record in a FeatureGroup.  An EventTime is a point in time when a new event occurs that corresponds to the creation or update of a Record in a FeatureGroup. All Records in the FeatureGroup must have a corresponding EventTime.  An EventTime can be a String or Fractional.  * Fractional: EventTime feature values must be a Unix timestamp in seconds.  * String: EventTime feature values must be an ISO-8601 string in the format. The following formats are supported yyyy-MM-dd'T'HH:mm:ssZ and yyyy-MM-dd'T'HH:mm:ss.SSSZ where yyyy, MM, and dd represent the year, month, and day respectively and HH, mm, ss, and if applicable, SSS represent the hour, month, second and milliseconds respsectively. 'T' and Z are constants.",
						MarkdownDescription: "The name of the feature that stores the EventTime of a Record in a FeatureGroup.  An EventTime is a point in time when a new event occurs that corresponds to the creation or update of a Record in a FeatureGroup. All Records in the FeatureGroup must have a corresponding EventTime.  An EventTime can be a String or Fractional.  * Fractional: EventTime feature values must be a Unix timestamp in seconds.  * String: EventTime feature values must be an ISO-8601 string in the format. The following formats are supported yyyy-MM-dd'T'HH:mm:ssZ and yyyy-MM-dd'T'HH:mm:ss.SSSZ where yyyy, MM, and dd represent the year, month, and day respectively and HH, mm, ss, and if applicable, SSS represent the hour, month, second and milliseconds respsectively. 'T' and Z are constants.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"feature_definitions": schema.ListNestedAttribute{
						Description:         "A list of Feature names and types. Name and Type is compulsory per Feature.  Valid feature FeatureTypes are Integral, Fractional and String.  FeatureNames cannot be any of the following: is_deleted, write_time, api_invocation_time  You can create up to 2,500 FeatureDefinitions per FeatureGroup.",
						MarkdownDescription: "A list of Feature names and types. Name and Type is compulsory per Feature.  Valid feature FeatureTypes are Integral, Fractional and String.  FeatureNames cannot be any of the following: is_deleted, write_time, api_invocation_time  You can create up to 2,500 FeatureDefinitions per FeatureGroup.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"feature_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"feature_type": schema.StringAttribute{
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

					"feature_group_name": schema.StringAttribute{
						Description:         "The name of the FeatureGroup. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account. The name:  * Must start and end with an alphanumeric character.  * Can only contain alphanumeric character and hyphens. Spaces are not allowed.",
						MarkdownDescription: "The name of the FeatureGroup. The name must be unique within an Amazon Web Services Region in an Amazon Web Services account. The name:  * Must start and end with an alphanumeric character.  * Can only contain alphanumeric character and hyphens. Spaces are not allowed.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"offline_store_config": schema.SingleNestedAttribute{
						Description:         "Use this to configure an OfflineFeatureStore. This parameter allows you to specify:  * The Amazon Simple Storage Service (Amazon S3) location of an OfflineStore.  * A configuration for an Amazon Web Services Glue or Amazon Web Services Hive data catalog.  * An KMS encryption key to encrypt the Amazon S3 location used for OfflineStore. If KMS encryption key is not specified, by default we encrypt all data at rest using Amazon Web Services KMS key. By defining your bucket-level key (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucket-key.html) for SSE, you can reduce Amazon Web Services KMS requests costs by up to 99 percent.  * Format for the offline store table. Supported formats are Glue (Default) and Apache Iceberg (https://iceberg.apache.org/).  To learn more about this parameter, see OfflineStoreConfig.",
						MarkdownDescription: "Use this to configure an OfflineFeatureStore. This parameter allows you to specify:  * The Amazon Simple Storage Service (Amazon S3) location of an OfflineStore.  * A configuration for an Amazon Web Services Glue or Amazon Web Services Hive data catalog.  * An KMS encryption key to encrypt the Amazon S3 location used for OfflineStore. If KMS encryption key is not specified, by default we encrypt all data at rest using Amazon Web Services KMS key. By defining your bucket-level key (https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucket-key.html) for SSE, you can reduce Amazon Web Services KMS requests costs by up to 99 percent.  * Format for the offline store table. Supported formats are Glue (Default) and Apache Iceberg (https://iceberg.apache.org/).  To learn more about this parameter, see OfflineStoreConfig.",
						Attributes: map[string]schema.Attribute{
							"data_catalog_config": schema.SingleNestedAttribute{
								Description:         "The meta data of the Glue table which serves as data catalog for the OfflineStore.",
								MarkdownDescription: "The meta data of the Glue table which serves as data catalog for the OfflineStore.",
								Attributes: map[string]schema.Attribute{
									"catalog": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"database": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"table_name": schema.StringAttribute{
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

							"disable_glue_table_creation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_storage_config": schema.SingleNestedAttribute{
								Description:         "The Amazon Simple Storage (Amazon S3) location and and security configuration for OfflineStore.",
								MarkdownDescription: "The Amazon Simple Storage (Amazon S3) location and and security configuration for OfflineStore.",
								Attributes: map[string]schema.Attribute{
									"kms_key_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resolved_output_s3_uri": schema.StringAttribute{
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

					"online_store_config": schema.SingleNestedAttribute{
						Description:         "You can turn the OnlineStore on or off by specifying True for the EnableOnlineStore flag in OnlineStoreConfig; the default value is False.  You can also include an Amazon Web Services KMS key ID (KMSKeyId) for at-rest encryption of the OnlineStore.",
						MarkdownDescription: "You can turn the OnlineStore on or off by specifying True for the EnableOnlineStore flag in OnlineStoreConfig; the default value is False.  You can also include an Amazon Web Services KMS key ID (KMSKeyId) for at-rest encryption of the OnlineStore.",
						Attributes: map[string]schema.Attribute{
							"enable_online_store": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_config": schema.SingleNestedAttribute{
								Description:         "The security configuration for OnlineStore.",
								MarkdownDescription: "The security configuration for OnlineStore.",
								Attributes: map[string]schema.Attribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"record_identifier_feature_name": schema.StringAttribute{
						Description:         "The name of the Feature whose value uniquely identifies a Record defined in the FeatureStore. Only the latest record per identifier value will be stored in the OnlineStore. RecordIdentifierFeatureName must be one of feature definitions' names.  You use the RecordIdentifierFeatureName to access data in a FeatureStore.  This name:  * Must start and end with an alphanumeric character.  * Can only contains alphanumeric characters, hyphens, underscores. Spaces are not allowed.",
						MarkdownDescription: "The name of the Feature whose value uniquely identifies a Record defined in the FeatureStore. Only the latest record per identifier value will be stored in the OnlineStore. RecordIdentifierFeatureName must be one of feature definitions' names.  You use the RecordIdentifierFeatureName to access data in a FeatureStore.  This name:  * Must start and end with an alphanumeric character.  * Can only contains alphanumeric characters, hyphens, underscores. Spaces are not allowed.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the OfflineStore if an OfflineStoreConfig is provided.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the OfflineStore if an OfflineStoreConfig is provided.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Tags used to identify Features in each FeatureGroup.",
						MarkdownDescription: "Tags used to identify Features in each FeatureGroup.",
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

func (r *SagemakerServicesK8SAwsFeatureGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sagemaker_services_k8s_aws_feature_group_v1alpha1_manifest")

	var model SagemakerServicesK8SAwsFeatureGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("sagemaker.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("FeatureGroup")

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