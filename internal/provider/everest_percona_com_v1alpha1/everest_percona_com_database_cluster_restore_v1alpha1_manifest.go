/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package everest_percona_com_v1alpha1

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
	_ datasource.DataSource = &EverestPerconaComDatabaseClusterRestoreV1Alpha1Manifest{}
)

func NewEverestPerconaComDatabaseClusterRestoreV1Alpha1Manifest() datasource.DataSource {
	return &EverestPerconaComDatabaseClusterRestoreV1Alpha1Manifest{}
}

type EverestPerconaComDatabaseClusterRestoreV1Alpha1Manifest struct{}

type EverestPerconaComDatabaseClusterRestoreV1Alpha1ManifestData struct {
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
		DataSource *struct {
			BackupSource *struct {
				BackupStorageName *string `tfsdk:"backup_storage_name" json:"backupStorageName,omitempty"`
				Path              *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"backup_source" json:"backupSource,omitempty"`
			DbClusterBackupName *string `tfsdk:"db_cluster_backup_name" json:"dbClusterBackupName,omitempty"`
			Pitr                *struct {
				Date *string `tfsdk:"date" json:"date,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"pitr" json:"pitr,omitempty"`
		} `tfsdk:"data_source" json:"dataSource,omitempty"`
		DbClusterName *string `tfsdk:"db_cluster_name" json:"dbClusterName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EverestPerconaComDatabaseClusterRestoreV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_everest_percona_com_database_cluster_restore_v1alpha1_manifest"
}

func (r *EverestPerconaComDatabaseClusterRestoreV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DatabaseClusterRestore is the Schema for the databaseclusterrestores API.",
		MarkdownDescription: "DatabaseClusterRestore is the Schema for the databaseclusterrestores API.",
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
				Description:         "DatabaseClusterRestoreSpec defines the desired state of DatabaseClusterRestore.",
				MarkdownDescription: "DatabaseClusterRestoreSpec defines the desired state of DatabaseClusterRestore.",
				Attributes: map[string]schema.Attribute{
					"data_source": schema.SingleNestedAttribute{
						Description:         "DataSource defines a data source for restoration.",
						MarkdownDescription: "DataSource defines a data source for restoration.",
						Attributes: map[string]schema.Attribute{
							"backup_source": schema.SingleNestedAttribute{
								Description:         "BackupSource is the backup source to restore from",
								MarkdownDescription: "BackupSource is the backup source to restore from",
								Attributes: map[string]schema.Attribute{
									"backup_storage_name": schema.StringAttribute{
										Description:         "BackupStorageName is the name of the BackupStorage used for backups.",
										MarkdownDescription: "BackupStorageName is the name of the BackupStorage used for backups.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "Path is the path to the backup file/directory.",
										MarkdownDescription: "Path is the path to the backup file/directory.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_cluster_backup_name": schema.StringAttribute{
								Description:         "DBClusterBackupName is the name of the DB cluster backup to restore from",
								MarkdownDescription: "DBClusterBackupName is the name of the DB cluster backup to restore from",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pitr": schema.SingleNestedAttribute{
								Description:         "PITR is the point-in-time recovery configuration",
								MarkdownDescription: "PITR is the point-in-time recovery configuration",
								Attributes: map[string]schema.Attribute{
									"date": schema.StringAttribute{
										Description:         "Date is the UTC date to recover to. The accepted format: '2006-01-02T15:04:05Z'.",
										MarkdownDescription: "Date is the UTC date to recover to. The accepted format: '2006-01-02T15:04:05Z'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is the type of recovery.",
										MarkdownDescription: "Type is the type of recovery.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("date", "latest"),
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

					"db_cluster_name": schema.StringAttribute{
						Description:         "DBClusterName defines the cluster name to restore.",
						MarkdownDescription: "DBClusterName defines the cluster name to restore.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *EverestPerconaComDatabaseClusterRestoreV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_everest_percona_com_database_cluster_restore_v1alpha1_manifest")

	var model EverestPerconaComDatabaseClusterRestoreV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("everest.percona.com/v1alpha1")
	model.Kind = pointer.String("DatabaseClusterRestore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
