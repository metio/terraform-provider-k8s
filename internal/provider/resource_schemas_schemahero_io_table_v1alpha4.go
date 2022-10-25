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

type SchemasSchemaheroIoTableV1Alpha4Resource struct{}

var (
	_ resource.Resource = (*SchemasSchemaheroIoTableV1Alpha4Resource)(nil)
)

type SchemasSchemaheroIoTableV1Alpha4TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SchemasSchemaheroIoTableV1Alpha4GoModel struct {
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
		Database *string `tfsdk:"database" yaml:"database,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Requires *[]string `tfsdk:"requires" yaml:"requires,omitempty"`

		Schema *struct {
			Cassandra *struct {
				ClusteringOrder *struct {
					Column *string `tfsdk:"column" yaml:"column,omitempty"`

					IsDescending *bool `tfsdk:"is_descending" yaml:"isDescending,omitempty"`
				} `tfsdk:"clustering_order" yaml:"clusteringOrder,omitempty"`

				Columns *[]struct {
					IsStatic *bool `tfsdk:"is_static" yaml:"isStatic,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"columns" yaml:"columns,omitempty"`

				IsDeleted *bool `tfsdk:"is_deleted" yaml:"isDeleted,omitempty"`

				PrimaryKey *[]string `tfsdk:"primary_key" yaml:"primaryKey,omitempty"`

				Properties *struct {
					BloomFilterFPChance *string `tfsdk:"bloom_filter_fp_chance" yaml:"bloomFilterFPChance,omitempty"`

					Caching *map[string]string `tfsdk:"caching" yaml:"caching,omitempty"`

					Comment *string `tfsdk:"comment" yaml:"comment,omitempty"`

					Compaction *map[string]string `tfsdk:"compaction" yaml:"compaction,omitempty"`

					Compression *map[string]string `tfsdk:"compression" yaml:"compression,omitempty"`

					CrcCheckChance *string `tfsdk:"crc_check_chance" yaml:"crcCheckChance,omitempty"`

					DcLocalReadRepairChance *string `tfsdk:"dc_local_read_repair_chance" yaml:"dcLocalReadRepairChance,omitempty"`

					DefaultTTL *int64 `tfsdk:"default_ttl" yaml:"defaultTTL,omitempty"`

					GcGraceSeconds *int64 `tfsdk:"gc_grace_seconds" yaml:"gcGraceSeconds,omitempty"`

					MaxIndexInterval *int64 `tfsdk:"max_index_interval" yaml:"maxIndexInterval,omitempty"`

					MemtableFlushPeriodMs *int64 `tfsdk:"memtable_flush_period_ms" yaml:"memtableFlushPeriodMs,omitempty"`

					MinIndexInterval *int64 `tfsdk:"min_index_interval" yaml:"minIndexInterval,omitempty"`

					ReadRepairChance *string `tfsdk:"read_repair_chance" yaml:"readRepairChance,omitempty"`

					SpeculativeRetry *string `tfsdk:"speculative_retry" yaml:"speculativeRetry,omitempty"`
				} `tfsdk:"properties" yaml:"properties,omitempty"`
			} `tfsdk:"cassandra" yaml:"cassandra,omitempty"`

			Cockroachdb *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" yaml:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" yaml:"attributes,omitempty"`

					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" yaml:"notNull,omitempty"`
					} `tfsdk:"constraints" yaml:"constraints,omitempty"`

					Default *string `tfsdk:"default" yaml:"default,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"columns" yaml:"columns,omitempty"`

				ForeignKeys *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					OnDelete *string `tfsdk:"on_delete" yaml:"onDelete,omitempty"`

					References *struct {
						Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

						Table *string `tfsdk:"table" yaml:"table,omitempty"`
					} `tfsdk:"references" yaml:"references,omitempty"`
				} `tfsdk:"foreign_keys" yaml:"foreignKeys,omitempty"`

				Indexes *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					IsUnique *bool `tfsdk:"is_unique" yaml:"isUnique,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"indexes" yaml:"indexes,omitempty"`

				IsDeleted *bool `tfsdk:"is_deleted" yaml:"isDeleted,omitempty"`

				Json_triggers *[]struct {
					Arguments *[]string `tfsdk:"arguments" yaml:"arguments,omitempty"`

					Condition *string `tfsdk:"condition" yaml:"condition,omitempty"`

					ConstraintTrigger *bool `tfsdk:"constraint_trigger" yaml:"constraintTrigger,omitempty"`

					Events *[]string `tfsdk:"events" yaml:"events,omitempty"`

					ExecuteProcedure *string `tfsdk:"execute_procedure" yaml:"executeProcedure,omitempty"`

					ForEachRun *bool `tfsdk:"for_each_run" yaml:"forEachRun,omitempty"`

					ForEachStatement *bool `tfsdk:"for_each_statement" yaml:"forEachStatement,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"json_triggers" yaml:"json:triggers,omitempty"`

				PrimaryKey *[]string `tfsdk:"primary_key" yaml:"primaryKey,omitempty"`
			} `tfsdk:"cockroachdb" yaml:"cockroachdb,omitempty"`

			Mysql *struct {
				Collation *string `tfsdk:"collation" yaml:"collation,omitempty"`

				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" yaml:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" yaml:"attributes,omitempty"`

					Charset *string `tfsdk:"charset" yaml:"charset,omitempty"`

					Collation *string `tfsdk:"collation" yaml:"collation,omitempty"`

					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" yaml:"notNull,omitempty"`
					} `tfsdk:"constraints" yaml:"constraints,omitempty"`

					Default *string `tfsdk:"default" yaml:"default,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"columns" yaml:"columns,omitempty"`

				DefaultCharset *string `tfsdk:"default_charset" yaml:"defaultCharset,omitempty"`

				ForeignKeys *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					OnDelete *string `tfsdk:"on_delete" yaml:"onDelete,omitempty"`

					References *struct {
						Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

						Table *string `tfsdk:"table" yaml:"table,omitempty"`
					} `tfsdk:"references" yaml:"references,omitempty"`
				} `tfsdk:"foreign_keys" yaml:"foreignKeys,omitempty"`

				Indexes *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					IsUnique *bool `tfsdk:"is_unique" yaml:"isUnique,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"indexes" yaml:"indexes,omitempty"`

				IsDeleted *bool `tfsdk:"is_deleted" yaml:"isDeleted,omitempty"`

				PrimaryKey *[]string `tfsdk:"primary_key" yaml:"primaryKey,omitempty"`
			} `tfsdk:"mysql" yaml:"mysql,omitempty"`

			Postgres *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" yaml:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" yaml:"attributes,omitempty"`

					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" yaml:"notNull,omitempty"`
					} `tfsdk:"constraints" yaml:"constraints,omitempty"`

					Default *string `tfsdk:"default" yaml:"default,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"columns" yaml:"columns,omitempty"`

				ForeignKeys *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					OnDelete *string `tfsdk:"on_delete" yaml:"onDelete,omitempty"`

					References *struct {
						Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

						Table *string `tfsdk:"table" yaml:"table,omitempty"`
					} `tfsdk:"references" yaml:"references,omitempty"`
				} `tfsdk:"foreign_keys" yaml:"foreignKeys,omitempty"`

				Indexes *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					IsUnique *bool `tfsdk:"is_unique" yaml:"isUnique,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"indexes" yaml:"indexes,omitempty"`

				IsDeleted *bool `tfsdk:"is_deleted" yaml:"isDeleted,omitempty"`

				Json_triggers *[]struct {
					Arguments *[]string `tfsdk:"arguments" yaml:"arguments,omitempty"`

					Condition *string `tfsdk:"condition" yaml:"condition,omitempty"`

					ConstraintTrigger *bool `tfsdk:"constraint_trigger" yaml:"constraintTrigger,omitempty"`

					Events *[]string `tfsdk:"events" yaml:"events,omitempty"`

					ExecuteProcedure *string `tfsdk:"execute_procedure" yaml:"executeProcedure,omitempty"`

					ForEachRun *bool `tfsdk:"for_each_run" yaml:"forEachRun,omitempty"`

					ForEachStatement *bool `tfsdk:"for_each_statement" yaml:"forEachStatement,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"json_triggers" yaml:"json:triggers,omitempty"`

				PrimaryKey *[]string `tfsdk:"primary_key" yaml:"primaryKey,omitempty"`
			} `tfsdk:"postgres" yaml:"postgres,omitempty"`

			Rqlite *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" yaml:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" yaml:"attributes,omitempty"`

					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" yaml:"notNull,omitempty"`
					} `tfsdk:"constraints" yaml:"constraints,omitempty"`

					Default *string `tfsdk:"default" yaml:"default,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"columns" yaml:"columns,omitempty"`

				ForeignKeys *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					OnDelete *string `tfsdk:"on_delete" yaml:"onDelete,omitempty"`

					References *struct {
						Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

						Table *string `tfsdk:"table" yaml:"table,omitempty"`
					} `tfsdk:"references" yaml:"references,omitempty"`
				} `tfsdk:"foreign_keys" yaml:"foreignKeys,omitempty"`

				Indexes *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					IsUnique *bool `tfsdk:"is_unique" yaml:"isUnique,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"indexes" yaml:"indexes,omitempty"`

				IsDeleted *bool `tfsdk:"is_deleted" yaml:"isDeleted,omitempty"`

				PrimaryKey *[]string `tfsdk:"primary_key" yaml:"primaryKey,omitempty"`

				Strict *bool `tfsdk:"strict" yaml:"strict,omitempty"`
			} `tfsdk:"rqlite" yaml:"rqlite,omitempty"`

			Sqlite *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" yaml:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" yaml:"attributes,omitempty"`

					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" yaml:"notNull,omitempty"`
					} `tfsdk:"constraints" yaml:"constraints,omitempty"`

					Default *string `tfsdk:"default" yaml:"default,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"columns" yaml:"columns,omitempty"`

				ForeignKeys *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					OnDelete *string `tfsdk:"on_delete" yaml:"onDelete,omitempty"`

					References *struct {
						Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

						Table *string `tfsdk:"table" yaml:"table,omitempty"`
					} `tfsdk:"references" yaml:"references,omitempty"`
				} `tfsdk:"foreign_keys" yaml:"foreignKeys,omitempty"`

				Indexes *[]struct {
					Columns *[]string `tfsdk:"columns" yaml:"columns,omitempty"`

					IsUnique *bool `tfsdk:"is_unique" yaml:"isUnique,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"indexes" yaml:"indexes,omitempty"`

				IsDeleted *bool `tfsdk:"is_deleted" yaml:"isDeleted,omitempty"`

				PrimaryKey *[]string `tfsdk:"primary_key" yaml:"primaryKey,omitempty"`
			} `tfsdk:"sqlite" yaml:"sqlite,omitempty"`
		} `tfsdk:"schema" yaml:"schema,omitempty"`

		SeedData *struct {
			Rows *[]struct {
				Columns *[]struct {
					Column *string `tfsdk:"column" yaml:"column,omitempty"`

					Value *struct {
						Int *int64 `tfsdk:"int" yaml:"int,omitempty"`

						Str *string `tfsdk:"str" yaml:"str,omitempty"`
					} `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"columns" yaml:"columns,omitempty"`
			} `tfsdk:"rows" yaml:"rows,omitempty"`
		} `tfsdk:"seed_data" yaml:"seedData,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSchemasSchemaheroIoTableV1Alpha4Resource() resource.Resource {
	return &SchemasSchemaheroIoTableV1Alpha4Resource{}
}

func (r *SchemasSchemaheroIoTableV1Alpha4Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_schemas_schemahero_io_table_v1alpha4"
}

func (r *SchemasSchemaheroIoTableV1Alpha4Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Table is the Schema for the tables API",
		MarkdownDescription: "Table is the Schema for the tables API",
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
				Description:         "TableSpec defines the desired state of Table",
				MarkdownDescription: "TableSpec defines the desired state of Table",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"database": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"requires": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"schema": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cassandra": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"clustering_order": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"column": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"is_descending": {
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

									"columns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"is_static": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"is_deleted": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"primary_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"bloom_filter_fp_chance": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"caching": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"comment": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"compaction": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"compression": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"crc_check_chance": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dc_local_read_repair_chance": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"default_ttl": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"gc_grace_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_index_interval": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memtable_flush_period_ms": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_index_interval": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_repair_chance": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"speculative_retry": {
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

							"cockroachdb": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"columns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"attributes": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"auto_increment": {
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

											"constraints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"not_null": {
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

											"default": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"on_delete": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"references": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"columns": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"table": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"is_unique": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
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

									"is_deleted": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_triggers": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"arguments": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"condition": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"constraint_trigger": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"events": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"execute_procedure": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"for_each_run": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"for_each_statement": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
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

									"primary_key": {
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

							"mysql": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"collation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"columns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"attributes": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"auto_increment": {
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

											"charset": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"collation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"constraints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"not_null": {
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

											"default": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_charset": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"on_delete": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"references": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"columns": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"table": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"is_unique": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
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

									"is_deleted": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"primary_key": {
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

							"postgres": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"columns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"attributes": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"auto_increment": {
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

											"constraints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"not_null": {
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

											"default": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"on_delete": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"references": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"columns": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"table": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"is_unique": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
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

									"is_deleted": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_triggers": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"arguments": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"condition": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"constraint_trigger": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"events": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"execute_procedure": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"for_each_run": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"for_each_statement": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
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

									"primary_key": {
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

							"rqlite": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"columns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"attributes": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"auto_increment": {
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

											"constraints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"not_null": {
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

											"default": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"on_delete": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"references": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"columns": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"table": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"is_unique": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
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

									"is_deleted": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"primary_key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"strict": {
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

							"sqlite": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"columns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"attributes": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"auto_increment": {
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

											"constraints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"not_null": {
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

											"default": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"on_delete": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"references": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"columns": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"table": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"columns": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"is_unique": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
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

									"is_deleted": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"primary_key": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"seed_data": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"rows": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"columns": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"column": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"int": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"str": {
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
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
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

func (r *SchemasSchemaheroIoTableV1Alpha4Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_schemas_schemahero_io_table_v1alpha4")

	var state SchemasSchemaheroIoTableV1Alpha4TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SchemasSchemaheroIoTableV1Alpha4GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("schemas.schemahero.io/v1alpha4")
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

func (r *SchemasSchemaheroIoTableV1Alpha4Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_schemas_schemahero_io_table_v1alpha4")
	// NO-OP: All data is already in Terraform state
}

func (r *SchemasSchemaheroIoTableV1Alpha4Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_schemas_schemahero_io_table_v1alpha4")

	var state SchemasSchemaheroIoTableV1Alpha4TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SchemasSchemaheroIoTableV1Alpha4GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("schemas.schemahero.io/v1alpha4")
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

func (r *SchemasSchemaheroIoTableV1Alpha4Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_schemas_schemahero_io_table_v1alpha4")
	// NO-OP: Terraform removes the state automatically for us
}
