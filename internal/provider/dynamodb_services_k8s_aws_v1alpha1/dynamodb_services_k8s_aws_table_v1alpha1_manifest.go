/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package dynamodb_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &DynamodbServicesK8SAwsTableV1Alpha1Manifest{}
)

func NewDynamodbServicesK8SAwsTableV1Alpha1Manifest() datasource.DataSource {
	return &DynamodbServicesK8SAwsTableV1Alpha1Manifest{}
}

type DynamodbServicesK8SAwsTableV1Alpha1Manifest struct{}

type DynamodbServicesK8SAwsTableV1Alpha1ManifestData struct {
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
		AttributeDefinitions *[]struct {
			AttributeName *string `tfsdk:"attribute_name" json:"attributeName,omitempty"`
			AttributeType *string `tfsdk:"attribute_type" json:"attributeType,omitempty"`
		} `tfsdk:"attribute_definitions" json:"attributeDefinitions,omitempty"`
		BillingMode       *string `tfsdk:"billing_mode" json:"billingMode,omitempty"`
		ContinuousBackups *struct {
			PointInTimeRecoveryEnabled *bool `tfsdk:"point_in_time_recovery_enabled" json:"pointInTimeRecoveryEnabled,omitempty"`
		} `tfsdk:"continuous_backups" json:"continuousBackups,omitempty"`
		GlobalSecondaryIndexes *[]struct {
			IndexName *string `tfsdk:"index_name" json:"indexName,omitempty"`
			KeySchema *[]struct {
				AttributeName *string `tfsdk:"attribute_name" json:"attributeName,omitempty"`
				KeyType       *string `tfsdk:"key_type" json:"keyType,omitempty"`
			} `tfsdk:"key_schema" json:"keySchema,omitempty"`
			Projection *struct {
				NonKeyAttributes *[]string `tfsdk:"non_key_attributes" json:"nonKeyAttributes,omitempty"`
				ProjectionType   *string   `tfsdk:"projection_type" json:"projectionType,omitempty"`
			} `tfsdk:"projection" json:"projection,omitempty"`
			ProvisionedThroughput *struct {
				ReadCapacityUnits  *int64 `tfsdk:"read_capacity_units" json:"readCapacityUnits,omitempty"`
				WriteCapacityUnits *int64 `tfsdk:"write_capacity_units" json:"writeCapacityUnits,omitempty"`
			} `tfsdk:"provisioned_throughput" json:"provisionedThroughput,omitempty"`
		} `tfsdk:"global_secondary_indexes" json:"globalSecondaryIndexes,omitempty"`
		KeySchema *[]struct {
			AttributeName *string `tfsdk:"attribute_name" json:"attributeName,omitempty"`
			KeyType       *string `tfsdk:"key_type" json:"keyType,omitempty"`
		} `tfsdk:"key_schema" json:"keySchema,omitempty"`
		LocalSecondaryIndexes *[]struct {
			IndexName *string `tfsdk:"index_name" json:"indexName,omitempty"`
			KeySchema *[]struct {
				AttributeName *string `tfsdk:"attribute_name" json:"attributeName,omitempty"`
				KeyType       *string `tfsdk:"key_type" json:"keyType,omitempty"`
			} `tfsdk:"key_schema" json:"keySchema,omitempty"`
			Projection *struct {
				NonKeyAttributes *[]string `tfsdk:"non_key_attributes" json:"nonKeyAttributes,omitempty"`
				ProjectionType   *string   `tfsdk:"projection_type" json:"projectionType,omitempty"`
			} `tfsdk:"projection" json:"projection,omitempty"`
		} `tfsdk:"local_secondary_indexes" json:"localSecondaryIndexes,omitempty"`
		ProvisionedThroughput *struct {
			ReadCapacityUnits  *int64 `tfsdk:"read_capacity_units" json:"readCapacityUnits,omitempty"`
			WriteCapacityUnits *int64 `tfsdk:"write_capacity_units" json:"writeCapacityUnits,omitempty"`
		} `tfsdk:"provisioned_throughput" json:"provisionedThroughput,omitempty"`
		SseSpecification *struct {
			Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			KmsMasterKeyID *string `tfsdk:"kms_master_key_id" json:"kmsMasterKeyID,omitempty"`
			SseType        *string `tfsdk:"sse_type" json:"sseType,omitempty"`
		} `tfsdk:"sse_specification" json:"sseSpecification,omitempty"`
		StreamSpecification *struct {
			StreamEnabled  *bool   `tfsdk:"stream_enabled" json:"streamEnabled,omitempty"`
			StreamViewType *string `tfsdk:"stream_view_type" json:"streamViewType,omitempty"`
		} `tfsdk:"stream_specification" json:"streamSpecification,omitempty"`
		TableClass *string `tfsdk:"table_class" json:"tableClass,omitempty"`
		TableName  *string `tfsdk:"table_name" json:"tableName,omitempty"`
		Tags       *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TimeToLive *struct {
			AttributeName *string `tfsdk:"attribute_name" json:"attributeName,omitempty"`
			Enabled       *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DynamodbServicesK8SAwsTableV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dynamodb_services_k8s_aws_table_v1alpha1_manifest"
}

func (r *DynamodbServicesK8SAwsTableV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
					"attribute_definitions": schema.ListNestedAttribute{
						Description:         "An array of attributes that describe the key schema for the table and indexes.",
						MarkdownDescription: "An array of attributes that describe the key schema for the table and indexes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"attribute_type": schema.StringAttribute{
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

					"billing_mode": schema.StringAttribute{
						Description:         "Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.  * PROVISIONED - We recommend using PROVISIONED for predictable workloads. PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).  * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).",
						MarkdownDescription: "Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.  * PROVISIONED - We recommend using PROVISIONED for predictable workloads. PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).  * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"continuous_backups": schema.SingleNestedAttribute{
						Description:         "Represents the settings used to enable point in time recovery.",
						MarkdownDescription: "Represents the settings used to enable point in time recovery.",
						Attributes: map[string]schema.Attribute{
							"point_in_time_recovery_enabled": schema.BoolAttribute{
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

					"global_secondary_indexes": schema.ListNestedAttribute{
						Description:         "One or more global secondary indexes (the maximum is 20) to be created on the table. Each global secondary index in the array includes the following:  * IndexName - The name of the global secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the global secondary index.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.  * ProvisionedThroughput - The provisioned throughput settings for the global secondary index, consisting of read and write capacity units.",
						MarkdownDescription: "One or more global secondary indexes (the maximum is 20) to be created on the table. Each global secondary index in the array includes the following:  * IndexName - The name of the global secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the global secondary index.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.  * ProvisionedThroughput - The provisioned throughput settings for the global secondary index, consisting of read and write capacity units.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"index_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key_schema": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_type": schema.StringAttribute{
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

								"projection": schema.SingleNestedAttribute{
									Description:         "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
									MarkdownDescription: "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
									Attributes: map[string]schema.Attribute{
										"non_key_attributes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"projection_type": schema.StringAttribute{
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

								"provisioned_throughput": schema.SingleNestedAttribute{
									Description:         "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",
									MarkdownDescription: "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",
									Attributes: map[string]schema.Attribute{
										"read_capacity_units": schema.Int64Attribute{
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"key_schema": schema.ListNestedAttribute{
						Description:         "Specifies the attributes that make up the primary key for a table or an index. The attributes in KeySchema must also be defined in the AttributeDefinitions array. For more information, see Data Model (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html) in the Amazon DynamoDB Developer Guide.  Each KeySchemaElement in the array is composed of:  * AttributeName - The name of this key attribute.  * KeyType - The role that the key attribute will assume: HASH - partition key RANGE - sort key  The partition key of an item is also known as its hash attribute. The term 'hash attribute' derives from the DynamoDB usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.  The sort key of an item is also known as its range attribute. The term 'range attribute' derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.  For a simple primary key (partition key), you must provide exactly one element with a KeyType of HASH.  For a composite primary key (partition key and sort key), you must provide exactly two elements, in this order: The first element must have a KeyType of HASH, and the second element must have a KeyType of RANGE.  For more information, see Working with Tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key) in the Amazon DynamoDB Developer Guide.",
						MarkdownDescription: "Specifies the attributes that make up the primary key for a table or an index. The attributes in KeySchema must also be defined in the AttributeDefinitions array. For more information, see Data Model (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html) in the Amazon DynamoDB Developer Guide.  Each KeySchemaElement in the array is composed of:  * AttributeName - The name of this key attribute.  * KeyType - The role that the key attribute will assume: HASH - partition key RANGE - sort key  The partition key of an item is also known as its hash attribute. The term 'hash attribute' derives from the DynamoDB usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.  The sort key of an item is also known as its range attribute. The term 'range attribute' derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.  For a simple primary key (partition key), you must provide exactly one element with a KeyType of HASH.  For a composite primary key (partition key and sort key), you must provide exactly two elements, in this order: The first element must have a KeyType of HASH, and the second element must have a KeyType of RANGE.  For more information, see Working with Tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key) in the Amazon DynamoDB Developer Guide.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"attribute_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key_type": schema.StringAttribute{
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

					"local_secondary_indexes": schema.ListNestedAttribute{
						Description:         "One or more local secondary indexes (the maximum is 5) to be created on the table. Each index is scoped to a given partition key value. There is a 10 GB size limit per partition key value; otherwise, the size of a local secondary index is unconstrained.  Each local secondary index in the array includes the following:  * IndexName - The name of the local secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the local secondary index. The key schema must begin with the same partition key as the table.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.",
						MarkdownDescription: "One or more local secondary indexes (the maximum is 5) to be created on the table. Each index is scoped to a given partition key value. There is a 10 GB size limit per partition key value; otherwise, the size of a local secondary index is unconstrained.  Each local secondary index in the array includes the following:  * IndexName - The name of the local secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the local secondary index. The key schema must begin with the same partition key as the table.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"index_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key_schema": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"attribute_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_type": schema.StringAttribute{
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

								"projection": schema.SingleNestedAttribute{
									Description:         "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
									MarkdownDescription: "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
									Attributes: map[string]schema.Attribute{
										"non_key_attributes": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"projection_type": schema.StringAttribute{
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

					"provisioned_throughput": schema.SingleNestedAttribute{
						Description:         "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  If you set BillingMode as PROVISIONED, you must specify this property. If you set BillingMode as PAY_PER_REQUEST, you cannot specify this property.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",
						MarkdownDescription: "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  If you set BillingMode as PROVISIONED, you must specify this property. If you set BillingMode as PAY_PER_REQUEST, you cannot specify this property.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",
						Attributes: map[string]schema.Attribute{
							"read_capacity_units": schema.Int64Attribute{
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

					"sse_specification": schema.SingleNestedAttribute{
						Description:         "Represents the settings used to enable server-side encryption.",
						MarkdownDescription: "Represents the settings used to enable server-side encryption.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kms_master_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sse_type": schema.StringAttribute{
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

					"stream_specification": schema.SingleNestedAttribute{
						Description:         "The settings for DynamoDB Streams on the table. These settings consist of:  * StreamEnabled - Indicates whether DynamoDB Streams is to be enabled (true) or disabled (false).  * StreamViewType - When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values for StreamViewType are: KEYS_ONLY - Only the key attributes of the modified item are written to the stream. NEW_IMAGE - The entire item, as it appears after it was modified, is written to the stream. OLD_IMAGE - The entire item, as it appeared before it was modified, is written to the stream. NEW_AND_OLD_IMAGES - Both the new and the old item images of the item are written to the stream.",
						MarkdownDescription: "The settings for DynamoDB Streams on the table. These settings consist of:  * StreamEnabled - Indicates whether DynamoDB Streams is to be enabled (true) or disabled (false).  * StreamViewType - When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values for StreamViewType are: KEYS_ONLY - Only the key attributes of the modified item are written to the stream. NEW_IMAGE - The entire item, as it appears after it was modified, is written to the stream. OLD_IMAGE - The entire item, as it appeared before it was modified, is written to the stream. NEW_AND_OLD_IMAGES - Both the new and the old item images of the item are written to the stream.",
						Attributes: map[string]schema.Attribute{
							"stream_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stream_view_type": schema.StringAttribute{
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

					"table_class": schema.StringAttribute{
						Description:         "The table class of the new table. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.",
						MarkdownDescription: "The table class of the new table. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"table_name": schema.StringAttribute{
						Description:         "The name of the table to create.",
						MarkdownDescription: "The name of the table to create.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of key-value pairs to label the table. For more information, see Tagging for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html).",
						MarkdownDescription: "A list of key-value pairs to label the table. For more information, see Tagging for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html).",
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

					"time_to_live": schema.SingleNestedAttribute{
						Description:         "Represents the settings used to enable or disable Time to Live for the specified table.",
						MarkdownDescription: "Represents the settings used to enable or disable Time to Live for the specified table.",
						Attributes: map[string]schema.Attribute{
							"attribute_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
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

func (r *DynamodbServicesK8SAwsTableV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dynamodb_services_k8s_aws_table_v1alpha1_manifest")

	var model DynamodbServicesK8SAwsTableV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("dynamodb.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Table")

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
