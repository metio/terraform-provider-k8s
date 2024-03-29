/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package keyspaces_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &KeyspacesServicesK8SAwsTableV1Alpha1Manifest{}
)

func NewKeyspacesServicesK8SAwsTableV1Alpha1Manifest() datasource.DataSource {
	return &KeyspacesServicesK8SAwsTableV1Alpha1Manifest{}
}

type KeyspacesServicesK8SAwsTableV1Alpha1Manifest struct{}

type KeyspacesServicesK8SAwsTableV1Alpha1ManifestData struct {
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
		CapacitySpecification *struct {
			ReadCapacityUnits  *int64  `tfsdk:"read_capacity_units" json:"readCapacityUnits,omitempty"`
			ThroughputMode     *string `tfsdk:"throughput_mode" json:"throughputMode,omitempty"`
			WriteCapacityUnits *int64  `tfsdk:"write_capacity_units" json:"writeCapacityUnits,omitempty"`
		} `tfsdk:"capacity_specification" json:"capacitySpecification,omitempty"`
		ClientSideTimestamps *struct {
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"client_side_timestamps" json:"clientSideTimestamps,omitempty"`
		Comment *struct {
			Message *string `tfsdk:"message" json:"message,omitempty"`
		} `tfsdk:"comment" json:"comment,omitempty"`
		DefaultTimeToLive       *int64 `tfsdk:"default_time_to_live" json:"defaultTimeToLive,omitempty"`
		EncryptionSpecification *struct {
			KmsKeyIdentifier *string `tfsdk:"kms_key_identifier" json:"kmsKeyIdentifier,omitempty"`
			Type_            *string `tfsdk:"type_" json:"type_,omitempty"`
		} `tfsdk:"encryption_specification" json:"encryptionSpecification,omitempty"`
		KeyspaceName        *string `tfsdk:"keyspace_name" json:"keyspaceName,omitempty"`
		PointInTimeRecovery *struct {
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"point_in_time_recovery" json:"pointInTimeRecovery,omitempty"`
		SchemaDefinition *struct {
			AllColumns *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
			} `tfsdk:"all_columns" json:"allColumns,omitempty"`
			ClusteringKeys *[]struct {
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				OrderBy *string `tfsdk:"order_by" json:"orderBy,omitempty"`
			} `tfsdk:"clustering_keys" json:"clusteringKeys,omitempty"`
			PartitionKeys *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"partition_keys" json:"partitionKeys,omitempty"`
			StaticColumns *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"static_columns" json:"staticColumns,omitempty"`
		} `tfsdk:"schema_definition" json:"schemaDefinition,omitempty"`
		TableName *string `tfsdk:"table_name" json:"tableName,omitempty"`
		Tags      *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		Ttl *struct {
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"ttl" json:"ttl,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KeyspacesServicesK8SAwsTableV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_keyspaces_services_k8s_aws_table_v1alpha1_manifest"
}

func (r *KeyspacesServicesK8SAwsTableV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Table is the Schema for the Tables API",
		MarkdownDescription: "Table is the Schema for the Tables API",
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
				Description:         "TableSpec defines the desired state of Table.",
				MarkdownDescription: "TableSpec defines the desired state of Table.",
				Attributes: map[string]schema.Attribute{
					"capacity_specification": schema.SingleNestedAttribute{
						Description:         "Specifies the read/write throughput capacity mode for the table. The optionsare:   * throughputMode:PAY_PER_REQUEST and   * throughputMode:PROVISIONED - Provisioned capacity mode requires readCapacityUnits   and writeCapacityUnits as input.The default is throughput_mode:PAY_PER_REQUEST.For more information, see Read/write capacity modes (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)in the Amazon Keyspaces Developer Guide.",
						MarkdownDescription: "Specifies the read/write throughput capacity mode for the table. The optionsare:   * throughputMode:PAY_PER_REQUEST and   * throughputMode:PROVISIONED - Provisioned capacity mode requires readCapacityUnits   and writeCapacityUnits as input.The default is throughput_mode:PAY_PER_REQUEST.For more information, see Read/write capacity modes (https://docs.aws.amazon.com/keyspaces/latest/devguide/ReadWriteCapacityMode.html)in the Amazon Keyspaces Developer Guide.",
						Attributes: map[string]schema.Attribute{
							"read_capacity_units": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"throughput_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_capacity_units": schema.Int64Attribute{
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

					"client_side_timestamps": schema.SingleNestedAttribute{
						Description:         "Enables client-side timestamps for the table. By default, the setting isdisabled. You can enable client-side timestamps with the following option:   * status: 'enabled'Once client-side timestamps are enabled for a table, this setting cannotbe disabled.",
						MarkdownDescription: "Enables client-side timestamps for the table. By default, the setting isdisabled. You can enable client-side timestamps with the following option:   * status: 'enabled'Once client-side timestamps are enabled for a table, this setting cannotbe disabled.",
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
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

					"comment": schema.SingleNestedAttribute{
						Description:         "This parameter allows to enter a description of the table.",
						MarkdownDescription: "This parameter allows to enter a description of the table.",
						Attributes: map[string]schema.Attribute{
							"message": schema.StringAttribute{
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

					"default_time_to_live": schema.Int64Attribute{
						Description:         "The default Time to Live setting in seconds for the table.For more information, see Setting the default TTL value for a table (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl)in the Amazon Keyspaces Developer Guide.",
						MarkdownDescription: "The default Time to Live setting in seconds for the table.For more information, see Setting the default TTL value for a table (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL-how-it-works.html#ttl-howitworks_default_ttl)in the Amazon Keyspaces Developer Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encryption_specification": schema.SingleNestedAttribute{
						Description:         "Specifies how the encryption key for encryption at rest is managed for thetable. You can choose one of the following KMS key (KMS key):   * type:AWS_OWNED_KMS_KEY - This key is owned by Amazon Keyspaces.   * type:CUSTOMER_MANAGED_KMS_KEY - This key is stored in your account and   is created, owned, and managed by you. This option requires the kms_key_identifier   of the KMS key in Amazon Resource Name (ARN) format as input.The default is type:AWS_OWNED_KMS_KEY.For more information, see Encryption at rest (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)in the Amazon Keyspaces Developer Guide.",
						MarkdownDescription: "Specifies how the encryption key for encryption at rest is managed for thetable. You can choose one of the following KMS key (KMS key):   * type:AWS_OWNED_KMS_KEY - This key is owned by Amazon Keyspaces.   * type:CUSTOMER_MANAGED_KMS_KEY - This key is stored in your account and   is created, owned, and managed by you. This option requires the kms_key_identifier   of the KMS key in Amazon Resource Name (ARN) format as input.The default is type:AWS_OWNED_KMS_KEY.For more information, see Encryption at rest (https://docs.aws.amazon.com/keyspaces/latest/devguide/EncryptionAtRest.html)in the Amazon Keyspaces Developer Guide.",
						Attributes: map[string]schema.Attribute{
							"kms_key_identifier": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type_": schema.StringAttribute{
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

					"keyspace_name": schema.StringAttribute{
						Description:         "The name of the keyspace that the table is going to be created in.",
						MarkdownDescription: "The name of the keyspace that the table is going to be created in.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"point_in_time_recovery": schema.SingleNestedAttribute{
						Description:         "Specifies if pointInTimeRecovery is enabled or disabled for the table. Theoptions are:   * status=ENABLED   * status=DISABLEDIf it's not specified, the default is status=DISABLED.For more information, see Point-in-time recovery (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html)in the Amazon Keyspaces Developer Guide.",
						MarkdownDescription: "Specifies if pointInTimeRecovery is enabled or disabled for the table. Theoptions are:   * status=ENABLED   * status=DISABLEDIf it's not specified, the default is status=DISABLED.For more information, see Point-in-time recovery (https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery.html)in the Amazon Keyspaces Developer Guide.",
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
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

					"schema_definition": schema.SingleNestedAttribute{
						Description:         "The schemaDefinition consists of the following parameters.For each column to be created:   * name - The name of the column.   * type - An Amazon Keyspaces data type. For more information, see Data   types (https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types)   in the Amazon Keyspaces Developer Guide.The primary key of the table consists of the following columns:   * partitionKeys - The partition key can be a single column, or it can   be a compound value composed of two or more columns. The partition key   portion of the primary key is required and determines how Amazon Keyspaces   stores your data.   * name - The name of each partition key column.   * clusteringKeys - The optional clustering column portion of your primary   key determines how the data is clustered and sorted within each partition.   * name - The name of the clustering column.   * orderBy - Sets the ascendant (ASC) or descendant (DESC) order modifier.   To define a column as static use staticColumns - Static columns store   values that are shared by all rows in the same partition:   * name - The name of the column.   * type - An Amazon Keyspaces data type.",
						MarkdownDescription: "The schemaDefinition consists of the following parameters.For each column to be created:   * name - The name of the column.   * type - An Amazon Keyspaces data type. For more information, see Data   types (https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types)   in the Amazon Keyspaces Developer Guide.The primary key of the table consists of the following columns:   * partitionKeys - The partition key can be a single column, or it can   be a compound value composed of two or more columns. The partition key   portion of the primary key is required and determines how Amazon Keyspaces   stores your data.   * name - The name of each partition key column.   * clusteringKeys - The optional clustering column portion of your primary   key determines how the data is clustered and sorted within each partition.   * name - The name of the clustering column.   * orderBy - Sets the ascendant (ASC) or descendant (DESC) order modifier.   To define a column as static use staticColumns - Static columns store   values that are shared by all rows in the same partition:   * name - The name of the column.   * type - An Amazon Keyspaces data type.",
						Attributes: map[string]schema.Attribute{
							"all_columns": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type_": schema.StringAttribute{
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

							"clustering_keys": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"order_by": schema.StringAttribute{
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

							"partition_keys": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

							"static_columns": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"table_name": schema.StringAttribute{
						Description:         "The name of the table.",
						MarkdownDescription: "The name of the table.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of key-value pair tags to be attached to the resource.For more information, see Adding tags and labels to Amazon Keyspaces resources(https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)in the Amazon Keyspaces Developer Guide.",
						MarkdownDescription: "A list of key-value pair tags to be attached to the resource.For more information, see Adding tags and labels to Amazon Keyspaces resources(https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html)in the Amazon Keyspaces Developer Guide.",
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

					"ttl": schema.SingleNestedAttribute{
						Description:         "Enables Time to Live custom settings for the table. The options are:   * status:enabled   * status:disabledThe default is status:disabled. After ttl is enabled, you can't disable itfor the table.For more information, see Expiring data by using Amazon Keyspaces Time toLive (TTL) (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html)in the Amazon Keyspaces Developer Guide.",
						MarkdownDescription: "Enables Time to Live custom settings for the table. The options are:   * status:enabled   * status:disabledThe default is status:disabled. After ttl is enabled, you can't disable itfor the table.For more information, see Expiring data by using Amazon Keyspaces Time toLive (TTL) (https://docs.aws.amazon.com/keyspaces/latest/devguide/TTL.html)in the Amazon Keyspaces Developer Guide.",
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
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

func (r *KeyspacesServicesK8SAwsTableV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_keyspaces_services_k8s_aws_table_v1alpha1_manifest")

	var model KeyspacesServicesK8SAwsTableV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("keyspaces.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Table")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
