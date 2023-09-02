/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package schemas_schemahero_io_v1alpha4

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
	_ datasource.DataSource = &SchemasSchemaheroIoTableV1Alpha4Manifest{}
)

func NewSchemasSchemaheroIoTableV1Alpha4Manifest() datasource.DataSource {
	return &SchemasSchemaheroIoTableV1Alpha4Manifest{}
}

type SchemasSchemaheroIoTableV1Alpha4Manifest struct{}

type SchemasSchemaheroIoTableV1Alpha4ManifestData struct {
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
		Database *string   `tfsdk:"database" json:"database,omitempty"`
		Name     *string   `tfsdk:"name" json:"name,omitempty"`
		Requires *[]string `tfsdk:"requires" json:"requires,omitempty"`
		Schema   *struct {
			Cassandra *struct {
				ClusteringOrder *struct {
					Column       *string `tfsdk:"column" json:"column,omitempty"`
					IsDescending *bool   `tfsdk:"is_descending" json:"isDescending,omitempty"`
				} `tfsdk:"clustering_order" json:"clusteringOrder,omitempty"`
				Columns *[]struct {
					IsStatic *bool   `tfsdk:"is_static" json:"isStatic,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Type     *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
				IsDeleted  *bool     `tfsdk:"is_deleted" json:"isDeleted,omitempty"`
				PrimaryKey *[]string `tfsdk:"primary_key" json:"primaryKey,omitempty"`
				Properties *struct {
					BloomFilterFPChance     *string            `tfsdk:"bloom_filter_fp_chance" json:"bloomFilterFPChance,omitempty"`
					Caching                 *map[string]string `tfsdk:"caching" json:"caching,omitempty"`
					Comment                 *string            `tfsdk:"comment" json:"comment,omitempty"`
					Compaction              *map[string]string `tfsdk:"compaction" json:"compaction,omitempty"`
					Compression             *map[string]string `tfsdk:"compression" json:"compression,omitempty"`
					CrcCheckChance          *string            `tfsdk:"crc_check_chance" json:"crcCheckChance,omitempty"`
					DcLocalReadRepairChance *string            `tfsdk:"dc_local_read_repair_chance" json:"dcLocalReadRepairChance,omitempty"`
					DefaultTTL              *int64             `tfsdk:"default_ttl" json:"defaultTTL,omitempty"`
					GcGraceSeconds          *int64             `tfsdk:"gc_grace_seconds" json:"gcGraceSeconds,omitempty"`
					MaxIndexInterval        *int64             `tfsdk:"max_index_interval" json:"maxIndexInterval,omitempty"`
					MemtableFlushPeriodMs   *int64             `tfsdk:"memtable_flush_period_ms" json:"memtableFlushPeriodMs,omitempty"`
					MinIndexInterval        *int64             `tfsdk:"min_index_interval" json:"minIndexInterval,omitempty"`
					ReadRepairChance        *string            `tfsdk:"read_repair_chance" json:"readRepairChance,omitempty"`
					SpeculativeRetry        *string            `tfsdk:"speculative_retry" json:"speculativeRetry,omitempty"`
				} `tfsdk:"properties" json:"properties,omitempty"`
			} `tfsdk:"cassandra" json:"cassandra,omitempty"`
			Cockroachdb *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" json:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" json:"attributes,omitempty"`
					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" json:"notNull,omitempty"`
					} `tfsdk:"constraints" json:"constraints,omitempty"`
					Default *string `tfsdk:"default" json:"default,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
				ForeignKeys *[]struct {
					Columns    *[]string `tfsdk:"columns" json:"columns,omitempty"`
					Name       *string   `tfsdk:"name" json:"name,omitempty"`
					OnDelete   *string   `tfsdk:"on_delete" json:"onDelete,omitempty"`
					References *struct {
						Columns *[]string `tfsdk:"columns" json:"columns,omitempty"`
						Table   *string   `tfsdk:"table" json:"table,omitempty"`
					} `tfsdk:"references" json:"references,omitempty"`
				} `tfsdk:"foreign_keys" json:"foreignKeys,omitempty"`
				Indexes *[]struct {
					Columns  *[]string `tfsdk:"columns" json:"columns,omitempty"`
					IsUnique *bool     `tfsdk:"is_unique" json:"isUnique,omitempty"`
					Name     *string   `tfsdk:"name" json:"name,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"indexes" json:"indexes,omitempty"`
				IsDeleted     *bool `tfsdk:"is_deleted" json:"isDeleted,omitempty"`
				Json_triggers *[]struct {
					Arguments         *[]string `tfsdk:"arguments" json:"arguments,omitempty"`
					Condition         *string   `tfsdk:"condition" json:"condition,omitempty"`
					ConstraintTrigger *bool     `tfsdk:"constraint_trigger" json:"constraintTrigger,omitempty"`
					Events            *[]string `tfsdk:"events" json:"events,omitempty"`
					ExecuteProcedure  *string   `tfsdk:"execute_procedure" json:"executeProcedure,omitempty"`
					ForEachRun        *bool     `tfsdk:"for_each_run" json:"forEachRun,omitempty"`
					ForEachStatement  *bool     `tfsdk:"for_each_statement" json:"forEachStatement,omitempty"`
					Name              *string   `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"json_triggers" json:"json:triggers,omitempty"`
				PrimaryKey *[]string `tfsdk:"primary_key" json:"primaryKey,omitempty"`
			} `tfsdk:"cockroachdb" json:"cockroachdb,omitempty"`
			Mysql *struct {
				Collation *string `tfsdk:"collation" json:"collation,omitempty"`
				Columns   *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" json:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" json:"attributes,omitempty"`
					Charset     *string `tfsdk:"charset" json:"charset,omitempty"`
					Collation   *string `tfsdk:"collation" json:"collation,omitempty"`
					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" json:"notNull,omitempty"`
					} `tfsdk:"constraints" json:"constraints,omitempty"`
					Default *string `tfsdk:"default" json:"default,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
				DefaultCharset *string `tfsdk:"default_charset" json:"defaultCharset,omitempty"`
				ForeignKeys    *[]struct {
					Columns    *[]string `tfsdk:"columns" json:"columns,omitempty"`
					Name       *string   `tfsdk:"name" json:"name,omitempty"`
					OnDelete   *string   `tfsdk:"on_delete" json:"onDelete,omitempty"`
					References *struct {
						Columns *[]string `tfsdk:"columns" json:"columns,omitempty"`
						Table   *string   `tfsdk:"table" json:"table,omitempty"`
					} `tfsdk:"references" json:"references,omitempty"`
				} `tfsdk:"foreign_keys" json:"foreignKeys,omitempty"`
				Indexes *[]struct {
					Columns  *[]string `tfsdk:"columns" json:"columns,omitempty"`
					IsUnique *bool     `tfsdk:"is_unique" json:"isUnique,omitempty"`
					Name     *string   `tfsdk:"name" json:"name,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"indexes" json:"indexes,omitempty"`
				IsDeleted  *bool     `tfsdk:"is_deleted" json:"isDeleted,omitempty"`
				PrimaryKey *[]string `tfsdk:"primary_key" json:"primaryKey,omitempty"`
			} `tfsdk:"mysql" json:"mysql,omitempty"`
			Postgres *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" json:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" json:"attributes,omitempty"`
					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" json:"notNull,omitempty"`
					} `tfsdk:"constraints" json:"constraints,omitempty"`
					Default *string `tfsdk:"default" json:"default,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
				ForeignKeys *[]struct {
					Columns    *[]string `tfsdk:"columns" json:"columns,omitempty"`
					Name       *string   `tfsdk:"name" json:"name,omitempty"`
					OnDelete   *string   `tfsdk:"on_delete" json:"onDelete,omitempty"`
					References *struct {
						Columns *[]string `tfsdk:"columns" json:"columns,omitempty"`
						Table   *string   `tfsdk:"table" json:"table,omitempty"`
					} `tfsdk:"references" json:"references,omitempty"`
				} `tfsdk:"foreign_keys" json:"foreignKeys,omitempty"`
				Indexes *[]struct {
					Columns  *[]string `tfsdk:"columns" json:"columns,omitempty"`
					IsUnique *bool     `tfsdk:"is_unique" json:"isUnique,omitempty"`
					Name     *string   `tfsdk:"name" json:"name,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"indexes" json:"indexes,omitempty"`
				IsDeleted     *bool `tfsdk:"is_deleted" json:"isDeleted,omitempty"`
				Json_triggers *[]struct {
					Arguments         *[]string `tfsdk:"arguments" json:"arguments,omitempty"`
					Condition         *string   `tfsdk:"condition" json:"condition,omitempty"`
					ConstraintTrigger *bool     `tfsdk:"constraint_trigger" json:"constraintTrigger,omitempty"`
					Events            *[]string `tfsdk:"events" json:"events,omitempty"`
					ExecuteProcedure  *string   `tfsdk:"execute_procedure" json:"executeProcedure,omitempty"`
					ForEachRun        *bool     `tfsdk:"for_each_run" json:"forEachRun,omitempty"`
					ForEachStatement  *bool     `tfsdk:"for_each_statement" json:"forEachStatement,omitempty"`
					Name              *string   `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"json_triggers" json:"json:triggers,omitempty"`
				PrimaryKey *[]string `tfsdk:"primary_key" json:"primaryKey,omitempty"`
			} `tfsdk:"postgres" json:"postgres,omitempty"`
			Rqlite *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" json:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" json:"attributes,omitempty"`
					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" json:"notNull,omitempty"`
					} `tfsdk:"constraints" json:"constraints,omitempty"`
					Default *string `tfsdk:"default" json:"default,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
				ForeignKeys *[]struct {
					Columns    *[]string `tfsdk:"columns" json:"columns,omitempty"`
					Name       *string   `tfsdk:"name" json:"name,omitempty"`
					OnDelete   *string   `tfsdk:"on_delete" json:"onDelete,omitempty"`
					References *struct {
						Columns *[]string `tfsdk:"columns" json:"columns,omitempty"`
						Table   *string   `tfsdk:"table" json:"table,omitempty"`
					} `tfsdk:"references" json:"references,omitempty"`
				} `tfsdk:"foreign_keys" json:"foreignKeys,omitempty"`
				Indexes *[]struct {
					Columns  *[]string `tfsdk:"columns" json:"columns,omitempty"`
					IsUnique *bool     `tfsdk:"is_unique" json:"isUnique,omitempty"`
					Name     *string   `tfsdk:"name" json:"name,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"indexes" json:"indexes,omitempty"`
				IsDeleted  *bool     `tfsdk:"is_deleted" json:"isDeleted,omitempty"`
				PrimaryKey *[]string `tfsdk:"primary_key" json:"primaryKey,omitempty"`
				Strict     *bool     `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"rqlite" json:"rqlite,omitempty"`
			Sqlite *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" json:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" json:"attributes,omitempty"`
					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" json:"notNull,omitempty"`
					} `tfsdk:"constraints" json:"constraints,omitempty"`
					Default *string `tfsdk:"default" json:"default,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
				ForeignKeys *[]struct {
					Columns    *[]string `tfsdk:"columns" json:"columns,omitempty"`
					Name       *string   `tfsdk:"name" json:"name,omitempty"`
					OnDelete   *string   `tfsdk:"on_delete" json:"onDelete,omitempty"`
					References *struct {
						Columns *[]string `tfsdk:"columns" json:"columns,omitempty"`
						Table   *string   `tfsdk:"table" json:"table,omitempty"`
					} `tfsdk:"references" json:"references,omitempty"`
				} `tfsdk:"foreign_keys" json:"foreignKeys,omitempty"`
				Indexes *[]struct {
					Columns  *[]string `tfsdk:"columns" json:"columns,omitempty"`
					IsUnique *bool     `tfsdk:"is_unique" json:"isUnique,omitempty"`
					Name     *string   `tfsdk:"name" json:"name,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"indexes" json:"indexes,omitempty"`
				IsDeleted  *bool     `tfsdk:"is_deleted" json:"isDeleted,omitempty"`
				PrimaryKey *[]string `tfsdk:"primary_key" json:"primaryKey,omitempty"`
				Strict     *bool     `tfsdk:"strict" json:"strict,omitempty"`
			} `tfsdk:"sqlite" json:"sqlite,omitempty"`
			Timescaledb *struct {
				Columns *[]struct {
					Attributes *struct {
						AutoIncrement *bool `tfsdk:"auto_increment" json:"autoIncrement,omitempty"`
					} `tfsdk:"attributes" json:"attributes,omitempty"`
					Constraints *struct {
						NotNull *bool `tfsdk:"not_null" json:"notNull,omitempty"`
					} `tfsdk:"constraints" json:"constraints,omitempty"`
					Default *string `tfsdk:"default" json:"default,omitempty"`
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Type    *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
				ForeignKeys *[]struct {
					Columns    *[]string `tfsdk:"columns" json:"columns,omitempty"`
					Name       *string   `tfsdk:"name" json:"name,omitempty"`
					OnDelete   *string   `tfsdk:"on_delete" json:"onDelete,omitempty"`
					References *struct {
						Columns *[]string `tfsdk:"columns" json:"columns,omitempty"`
						Table   *string   `tfsdk:"table" json:"table,omitempty"`
					} `tfsdk:"references" json:"references,omitempty"`
				} `tfsdk:"foreign_keys" json:"foreignKeys,omitempty"`
				Hypertable *struct {
					AssociatedSchemaName  *string `tfsdk:"associated_schema_name" json:"associatedSchemaName,omitempty"`
					AssociatedTablePrefix *string `tfsdk:"associated_table_prefix" json:"associatedTablePrefix,omitempty"`
					ChunkTimeInterval     *string `tfsdk:"chunk_time_interval" json:"chunkTimeInterval,omitempty"`
					Compression           *struct {
						Interval  *string `tfsdk:"interval" json:"interval,omitempty"`
						SegmentBy *string `tfsdk:"segment_by" json:"segmentBy,omitempty"`
					} `tfsdk:"compression" json:"compression,omitempty"`
					CreateDefaultIndexes *bool     `tfsdk:"create_default_indexes" json:"createDefaultIndexes,omitempty"`
					DataNodes            *[]string `tfsdk:"data_nodes" json:"dataNodes,omitempty"`
					IfNotExists          *bool     `tfsdk:"if_not_exists" json:"ifNotExists,omitempty"`
					MigrateData          *bool     `tfsdk:"migrate_data" json:"migrateData,omitempty"`
					NumberPartitions     *int64    `tfsdk:"number_partitions" json:"numberPartitions,omitempty"`
					PartitioningColumn   *string   `tfsdk:"partitioning_column" json:"partitioningColumn,omitempty"`
					PartitioningFunc     *string   `tfsdk:"partitioning_func" json:"partitioningFunc,omitempty"`
					ReplicationFactor    *int64    `tfsdk:"replication_factor" json:"replicationFactor,omitempty"`
					Retention            *struct {
						Interval *string `tfsdk:"interval" json:"interval,omitempty"`
					} `tfsdk:"retention" json:"retention,omitempty"`
					TimeColumnName       *string `tfsdk:"time_column_name" json:"timeColumnName,omitempty"`
					TimePartitioningFunc *string `tfsdk:"time_partitioning_func" json:"timePartitioningFunc,omitempty"`
				} `tfsdk:"hypertable" json:"hypertable,omitempty"`
				Indexes *[]struct {
					Columns  *[]string `tfsdk:"columns" json:"columns,omitempty"`
					IsUnique *bool     `tfsdk:"is_unique" json:"isUnique,omitempty"`
					Name     *string   `tfsdk:"name" json:"name,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"indexes" json:"indexes,omitempty"`
				IsDeleted  *bool     `tfsdk:"is_deleted" json:"isDeleted,omitempty"`
				PrimaryKey *[]string `tfsdk:"primary_key" json:"primaryKey,omitempty"`
				Triggers   *[]struct {
					Arguments         *[]string `tfsdk:"arguments" json:"arguments,omitempty"`
					Condition         *string   `tfsdk:"condition" json:"condition,omitempty"`
					ConstraintTrigger *bool     `tfsdk:"constraint_trigger" json:"constraintTrigger,omitempty"`
					Events            *[]string `tfsdk:"events" json:"events,omitempty"`
					ExecuteProcedure  *string   `tfsdk:"execute_procedure" json:"executeProcedure,omitempty"`
					ForEachRun        *bool     `tfsdk:"for_each_run" json:"forEachRun,omitempty"`
					ForEachStatement  *bool     `tfsdk:"for_each_statement" json:"forEachStatement,omitempty"`
					Name              *string   `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"triggers" json:"triggers,omitempty"`
			} `tfsdk:"timescaledb" json:"timescaledb,omitempty"`
		} `tfsdk:"schema" json:"schema,omitempty"`
		SeedData *struct {
			Rows *[]struct {
				Columns *[]struct {
					Column *string `tfsdk:"column" json:"column,omitempty"`
					Value  *struct {
						Int *int64  `tfsdk:"int" json:"int,omitempty"`
						Str *string `tfsdk:"str" json:"str,omitempty"`
					} `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"columns" json:"columns,omitempty"`
			} `tfsdk:"rows" json:"rows,omitempty"`
		} `tfsdk:"seed_data" json:"seedData,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SchemasSchemaheroIoTableV1Alpha4Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_schemas_schemahero_io_table_v1alpha4_manifest"
}

func (r *SchemasSchemaheroIoTableV1Alpha4Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Table is the Schema for the tables API",
		MarkdownDescription: "Table is the Schema for the tables API",
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
				Description:         "TableSpec defines the desired state of Table",
				MarkdownDescription: "TableSpec defines the desired state of Table",
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"requires": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"schema": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cassandra": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"clustering_order": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"column": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"is_descending": schema.BoolAttribute{
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

									"columns": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"is_static": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"is_deleted": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"primary_key": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"properties": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"bloom_filter_fp_chance": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"caching": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"comment": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"compaction": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"compression": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"crc_check_chance": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"dc_local_read_repair_chance": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"default_ttl": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"gc_grace_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_index_interval": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"memtable_flush_period_ms": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"min_index_interval": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_repair_chance": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"speculative_retry": schema.StringAttribute{
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

							"cockroachdb": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"columns": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auto_increment": schema.BoolAttribute{
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

												"constraints": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"not_null": schema.BoolAttribute{
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

												"default": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_delete": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"references": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"columns": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"table": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"is_unique": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
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

									"is_deleted": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"json_triggers": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"arguments": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"condition": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"constraint_trigger": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"execute_procedure": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"for_each_run": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"for_each_statement": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

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

									"primary_key": schema.ListAttribute{
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

							"mysql": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"collation": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"columns": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auto_increment": schema.BoolAttribute{
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

												"charset": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"collation": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"constraints": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"not_null": schema.BoolAttribute{
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

												"default": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"default_charset": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"foreign_keys": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_delete": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"references": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"columns": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"table": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"is_unique": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
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

									"is_deleted": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"primary_key": schema.ListAttribute{
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

							"postgres": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"columns": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auto_increment": schema.BoolAttribute{
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

												"constraints": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"not_null": schema.BoolAttribute{
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

												"default": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_delete": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"references": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"columns": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"table": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"is_unique": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
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

									"is_deleted": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"json_triggers": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"arguments": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"condition": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"constraint_trigger": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"execute_procedure": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"for_each_run": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"for_each_statement": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

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

									"primary_key": schema.ListAttribute{
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

							"rqlite": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"columns": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auto_increment": schema.BoolAttribute{
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

												"constraints": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"not_null": schema.BoolAttribute{
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

												"default": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_delete": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"references": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"columns": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"table": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"is_unique": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
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

									"is_deleted": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"primary_key": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"strict": schema.BoolAttribute{
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

							"sqlite": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"columns": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auto_increment": schema.BoolAttribute{
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

												"constraints": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"not_null": schema.BoolAttribute{
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

												"default": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_delete": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"references": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"columns": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"table": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"indexes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"is_unique": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
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

									"is_deleted": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"primary_key": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"strict": schema.BoolAttribute{
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

							"timescaledb": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"columns": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"attributes": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auto_increment": schema.BoolAttribute{
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

												"constraints": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"not_null": schema.BoolAttribute{
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

												"default": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"foreign_keys": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"on_delete": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"references": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"columns": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"table": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"hypertable": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"associated_schema_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"associated_table_prefix": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"chunk_time_interval": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"compression": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"segment_by": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"create_default_indexes": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"data_nodes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"if_not_exists": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"migrate_data": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"number_partitions": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"partitioning_column": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"partitioning_func": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replication_factor": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retention": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"time_column_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_partitioning_func": schema.StringAttribute{
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

									"indexes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"columns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"is_unique": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
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

									"is_deleted": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"primary_key": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"triggers": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"arguments": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"condition": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"constraint_trigger": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"events": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"execute_procedure": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"for_each_run": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"for_each_statement": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"seed_data": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"rows": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"columns": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"column": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"int": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"str": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},
									},
								},
								Required: true,
								Optional: false,
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

func (r *SchemasSchemaheroIoTableV1Alpha4Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_schemas_schemahero_io_table_v1alpha4_manifest")

	var model SchemasSchemaheroIoTableV1Alpha4ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("schemas.schemahero.io/v1alpha4")
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
