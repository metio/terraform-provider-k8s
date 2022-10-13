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

type DynamodbServicesK8SAwsTableV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*DynamodbServicesK8SAwsTableV1Alpha1Resource)(nil)
)

type DynamodbServicesK8SAwsTableV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type DynamodbServicesK8SAwsTableV1Alpha1GoModel struct {
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
		AttributeDefinitions *[]struct {
			AttributeName *string `tfsdk:"attribute_name" yaml:"attributeName,omitempty"`

			AttributeType *string `tfsdk:"attribute_type" yaml:"attributeType,omitempty"`
		} `tfsdk:"attribute_definitions" yaml:"attributeDefinitions,omitempty"`

		BillingMode *string `tfsdk:"billing_mode" yaml:"billingMode,omitempty"`

		GlobalSecondaryIndexes *[]struct {
			IndexName *string `tfsdk:"index_name" yaml:"indexName,omitempty"`

			KeySchema *[]struct {
				AttributeName *string `tfsdk:"attribute_name" yaml:"attributeName,omitempty"`

				KeyType *string `tfsdk:"key_type" yaml:"keyType,omitempty"`
			} `tfsdk:"key_schema" yaml:"keySchema,omitempty"`

			Projection *struct {
				NonKeyAttributes *[]string `tfsdk:"non_key_attributes" yaml:"nonKeyAttributes,omitempty"`

				ProjectionType *string `tfsdk:"projection_type" yaml:"projectionType,omitempty"`
			} `tfsdk:"projection" yaml:"projection,omitempty"`

			ProvisionedThroughput *struct {
				ReadCapacityUnits *int64 `tfsdk:"read_capacity_units" yaml:"readCapacityUnits,omitempty"`

				WriteCapacityUnits *int64 `tfsdk:"write_capacity_units" yaml:"writeCapacityUnits,omitempty"`
			} `tfsdk:"provisioned_throughput" yaml:"provisionedThroughput,omitempty"`
		} `tfsdk:"global_secondary_indexes" yaml:"globalSecondaryIndexes,omitempty"`

		KeySchema *[]struct {
			AttributeName *string `tfsdk:"attribute_name" yaml:"attributeName,omitempty"`

			KeyType *string `tfsdk:"key_type" yaml:"keyType,omitempty"`
		} `tfsdk:"key_schema" yaml:"keySchema,omitempty"`

		LocalSecondaryIndexes *[]struct {
			IndexName *string `tfsdk:"index_name" yaml:"indexName,omitempty"`

			KeySchema *[]struct {
				AttributeName *string `tfsdk:"attribute_name" yaml:"attributeName,omitempty"`

				KeyType *string `tfsdk:"key_type" yaml:"keyType,omitempty"`
			} `tfsdk:"key_schema" yaml:"keySchema,omitempty"`

			Projection *struct {
				NonKeyAttributes *[]string `tfsdk:"non_key_attributes" yaml:"nonKeyAttributes,omitempty"`

				ProjectionType *string `tfsdk:"projection_type" yaml:"projectionType,omitempty"`
			} `tfsdk:"projection" yaml:"projection,omitempty"`
		} `tfsdk:"local_secondary_indexes" yaml:"localSecondaryIndexes,omitempty"`

		ProvisionedThroughput *struct {
			ReadCapacityUnits *int64 `tfsdk:"read_capacity_units" yaml:"readCapacityUnits,omitempty"`

			WriteCapacityUnits *int64 `tfsdk:"write_capacity_units" yaml:"writeCapacityUnits,omitempty"`
		} `tfsdk:"provisioned_throughput" yaml:"provisionedThroughput,omitempty"`

		SseSpecification *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			KmsMasterKeyID *string `tfsdk:"kms_master_key_id" yaml:"kmsMasterKeyID,omitempty"`

			SseType *string `tfsdk:"sse_type" yaml:"sseType,omitempty"`
		} `tfsdk:"sse_specification" yaml:"sseSpecification,omitempty"`

		StreamSpecification *struct {
			StreamEnabled *bool `tfsdk:"stream_enabled" yaml:"streamEnabled,omitempty"`

			StreamViewType *string `tfsdk:"stream_view_type" yaml:"streamViewType,omitempty"`
		} `tfsdk:"stream_specification" yaml:"streamSpecification,omitempty"`

		TableClass *string `tfsdk:"table_class" yaml:"tableClass,omitempty"`

		TableName *string `tfsdk:"table_name" yaml:"tableName,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		TimeToLive *struct {
			AttributeName *string `tfsdk:"attribute_name" yaml:"attributeName,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"time_to_live" yaml:"timeToLive,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewDynamodbServicesK8SAwsTableV1Alpha1Resource() resource.Resource {
	return &DynamodbServicesK8SAwsTableV1Alpha1Resource{}
}

func (r *DynamodbServicesK8SAwsTableV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dynamodb_services_k8s_aws_table_v1alpha1"
}

func (r *DynamodbServicesK8SAwsTableV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Table is the Schema for the Tables API",
		MarkdownDescription: "Table is the Schema for the Tables API",
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
				Description:         "TableSpec defines the desired state of Table.",
				MarkdownDescription: "TableSpec defines the desired state of Table.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"attribute_definitions": {
						Description:         "An array of attributes that describe the key schema for the table and indexes.",
						MarkdownDescription: "An array of attributes that describe the key schema for the table and indexes.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"attribute_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"attribute_type": {
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

					"billing_mode": {
						Description:         "Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.  * PROVISIONED - We recommend using PROVISIONED for predictable workloads. PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).  * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).",
						MarkdownDescription: "Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.  * PROVISIONED - We recommend using PROVISIONED for predictable workloads. PROVISIONED sets the billing mode to Provisioned Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).  * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"global_secondary_indexes": {
						Description:         "One or more global secondary indexes (the maximum is 20) to be created on the table. Each global secondary index in the array includes the following:  * IndexName - The name of the global secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the global secondary index.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.  * ProvisionedThroughput - The provisioned throughput settings for the global secondary index, consisting of read and write capacity units.",
						MarkdownDescription: "One or more global secondary indexes (the maximum is 20) to be created on the table. Each global secondary index in the array includes the following:  * IndexName - The name of the global secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the global secondary index.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.  * ProvisionedThroughput - The provisioned throughput settings for the global secondary index, consisting of read and write capacity units.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"index_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key_schema": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"attribute_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_type": {
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

							"projection": {
								Description:         "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
								MarkdownDescription: "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"non_key_attributes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"projection_type": {
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

							"provisioned_throughput": {
								Description:         "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",
								MarkdownDescription: "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"read_capacity_units": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"write_capacity_units": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"key_schema": {
						Description:         "Specifies the attributes that make up the primary key for a table or an index. The attributes in KeySchema must also be defined in the AttributeDefinitions array. For more information, see Data Model (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html) in the Amazon DynamoDB Developer Guide.  Each KeySchemaElement in the array is composed of:  * AttributeName - The name of this key attribute.  * KeyType - The role that the key attribute will assume: HASH - partition key RANGE - sort key  The partition key of an item is also known as its hash attribute. The term 'hash attribute' derives from the DynamoDB usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.  The sort key of an item is also known as its range attribute. The term 'range attribute' derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.  For a simple primary key (partition key), you must provide exactly one element with a KeyType of HASH.  For a composite primary key (partition key and sort key), you must provide exactly two elements, in this order: The first element must have a KeyType of HASH, and the second element must have a KeyType of RANGE.  For more information, see Working with Tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key) in the Amazon DynamoDB Developer Guide.",
						MarkdownDescription: "Specifies the attributes that make up the primary key for a table or an index. The attributes in KeySchema must also be defined in the AttributeDefinitions array. For more information, see Data Model (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html) in the Amazon DynamoDB Developer Guide.  Each KeySchemaElement in the array is composed of:  * AttributeName - The name of this key attribute.  * KeyType - The role that the key attribute will assume: HASH - partition key RANGE - sort key  The partition key of an item is also known as its hash attribute. The term 'hash attribute' derives from the DynamoDB usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.  The sort key of an item is also known as its range attribute. The term 'range attribute' derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.  For a simple primary key (partition key), you must provide exactly one element with a KeyType of HASH.  For a composite primary key (partition key and sort key), you must provide exactly two elements, in this order: The first element must have a KeyType of HASH, and the second element must have a KeyType of RANGE.  For more information, see Working with Tables (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key) in the Amazon DynamoDB Developer Guide.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"attribute_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key_type": {
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

					"local_secondary_indexes": {
						Description:         "One or more local secondary indexes (the maximum is 5) to be created on the table. Each index is scoped to a given partition key value. There is a 10 GB size limit per partition key value; otherwise, the size of a local secondary index is unconstrained.  Each local secondary index in the array includes the following:  * IndexName - The name of the local secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the local secondary index. The key schema must begin with the same partition key as the table.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.",
						MarkdownDescription: "One or more local secondary indexes (the maximum is 5) to be created on the table. Each index is scoped to a given partition key value. There is a 10 GB size limit per partition key value; otherwise, the size of a local secondary index is unconstrained.  Each local secondary index in the array includes the following:  * IndexName - The name of the local secondary index. Must be unique only for this table.  * KeySchema - Specifies the key schema for the local secondary index. The key schema must begin with the same partition key as the table.  * Projection - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of: ProjectionType - One of the following: KEYS_ONLY - Only the index and primary keys are projected into the index. INCLUDE - Only the specified table attributes are projected into the index. The list of projected attributes is in NonKeyAttributes. ALL - All of the table attributes are projected into the index. NonKeyAttributes - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in NonKeyAttributes, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"index_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key_schema": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"attribute_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_type": {
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

							"projection": {
								Description:         "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",
								MarkdownDescription: "Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"non_key_attributes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"projection_type": {
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

					"provisioned_throughput": {
						Description:         "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  If you set BillingMode as PROVISIONED, you must specify this property. If you set BillingMode as PAY_PER_REQUEST, you cannot specify this property.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",
						MarkdownDescription: "Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the UpdateTable operation.  If you set BillingMode as PROVISIONED, you must specify this property. If you set BillingMode as PAY_PER_REQUEST, you cannot specify this property.  For current minimum and maximum provisioned throughput values, see Service, Account, and Table Quotas (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the Amazon DynamoDB Developer Guide.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"read_capacity_units": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"write_capacity_units": {
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

					"sse_specification": {
						Description:         "Represents the settings used to enable server-side encryption.",
						MarkdownDescription: "Represents the settings used to enable server-side encryption.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kms_master_key_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sse_type": {
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

					"stream_specification": {
						Description:         "The settings for DynamoDB Streams on the table. These settings consist of:  * StreamEnabled - Indicates whether DynamoDB Streams is to be enabled (true) or disabled (false).  * StreamViewType - When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values for StreamViewType are: KEYS_ONLY - Only the key attributes of the modified item are written to the stream. NEW_IMAGE - The entire item, as it appears after it was modified, is written to the stream. OLD_IMAGE - The entire item, as it appeared before it was modified, is written to the stream. NEW_AND_OLD_IMAGES - Both the new and the old item images of the item are written to the stream.",
						MarkdownDescription: "The settings for DynamoDB Streams on the table. These settings consist of:  * StreamEnabled - Indicates whether DynamoDB Streams is to be enabled (true) or disabled (false).  * StreamViewType - When an item in the table is modified, StreamViewType determines what information is written to the table's stream. Valid values for StreamViewType are: KEYS_ONLY - Only the key attributes of the modified item are written to the stream. NEW_IMAGE - The entire item, as it appears after it was modified, is written to the stream. OLD_IMAGE - The entire item, as it appeared before it was modified, is written to the stream. NEW_AND_OLD_IMAGES - Both the new and the old item images of the item are written to the stream.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"stream_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"stream_view_type": {
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

					"table_class": {
						Description:         "The table class of the new table. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.",
						MarkdownDescription: "The table class of the new table. Valid values are STANDARD and STANDARD_INFREQUENT_ACCESS.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"table_name": {
						Description:         "The name of the table to create.",
						MarkdownDescription: "The name of the table to create.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": {
						Description:         "A list of key-value pairs to label the table. For more information, see Tagging for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html).",
						MarkdownDescription: "A list of key-value pairs to label the table. For more information, see Tagging for DynamoDB (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html).",

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

					"time_to_live": {
						Description:         "Represents the settings used to enable or disable Time to Live for the specified table.",
						MarkdownDescription: "Represents the settings used to enable or disable Time to Live for the specified table.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"attribute_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

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

func (r *DynamodbServicesK8SAwsTableV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_dynamodb_services_k8s_aws_table_v1alpha1")

	var state DynamodbServicesK8SAwsTableV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DynamodbServicesK8SAwsTableV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("dynamodb.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Table")

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

func (r *DynamodbServicesK8SAwsTableV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dynamodb_services_k8s_aws_table_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *DynamodbServicesK8SAwsTableV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_dynamodb_services_k8s_aws_table_v1alpha1")

	var state DynamodbServicesK8SAwsTableV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DynamodbServicesK8SAwsTableV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("dynamodb.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Table")

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

func (r *DynamodbServicesK8SAwsTableV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_dynamodb_services_k8s_aws_table_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
