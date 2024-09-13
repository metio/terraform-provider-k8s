/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package postgresql_cnpg_io_v1

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
	_ datasource.DataSource = &PostgresqlCnpgIoBackupV1Manifest{}
)

func NewPostgresqlCnpgIoBackupV1Manifest() datasource.DataSource {
	return &PostgresqlCnpgIoBackupV1Manifest{}
}

type PostgresqlCnpgIoBackupV1Manifest struct{}

type PostgresqlCnpgIoBackupV1ManifestData struct {
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
		Cluster *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cluster" json:"cluster,omitempty"`
		Method              *string `tfsdk:"method" json:"method,omitempty"`
		Online              *bool   `tfsdk:"online" json:"online,omitempty"`
		OnlineConfiguration *struct {
			ImmediateCheckpoint *bool `tfsdk:"immediate_checkpoint" json:"immediateCheckpoint,omitempty"`
			WaitForArchive      *bool `tfsdk:"wait_for_archive" json:"waitForArchive,omitempty"`
		} `tfsdk:"online_configuration" json:"onlineConfiguration,omitempty"`
		PluginConfiguration *struct {
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		} `tfsdk:"plugin_configuration" json:"pluginConfiguration,omitempty"`
		Target *string `tfsdk:"target" json:"target,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PostgresqlCnpgIoBackupV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_postgresql_cnpg_io_backup_v1_manifest"
}

func (r *PostgresqlCnpgIoBackupV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Backup is the Schema for the backups API",
		MarkdownDescription: "Backup is the Schema for the backups API",
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
				Description:         "Specification of the desired behavior of the backup. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired behavior of the backup. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Attributes: map[string]schema.Attribute{
					"cluster": schema.SingleNestedAttribute{
						Description:         "The cluster to backup",
						MarkdownDescription: "The cluster to backup",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"method": schema.StringAttribute{
						Description:         "The backup method to be used, possible options are 'barmanObjectStore', 'volumeSnapshot' or 'plugin'. Defaults to: 'barmanObjectStore'.",
						MarkdownDescription: "The backup method to be used, possible options are 'barmanObjectStore', 'volumeSnapshot' or 'plugin'. Defaults to: 'barmanObjectStore'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("barmanObjectStore", "volumeSnapshot", "plugin"),
						},
					},

					"online": schema.BoolAttribute{
						Description:         "Whether the default type of backup with volume snapshots is online/hot ('true', default) or offline/cold ('false') Overrides the default setting specified in the cluster field '.spec.backup.volumeSnapshot.online'",
						MarkdownDescription: "Whether the default type of backup with volume snapshots is online/hot ('true', default) or offline/cold ('false') Overrides the default setting specified in the cluster field '.spec.backup.volumeSnapshot.online'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"online_configuration": schema.SingleNestedAttribute{
						Description:         "Configuration parameters to control the online/hot backup with volume snapshots Overrides the default settings specified in the cluster '.backup.volumeSnapshot.onlineConfiguration' stanza",
						MarkdownDescription: "Configuration parameters to control the online/hot backup with volume snapshots Overrides the default settings specified in the cluster '.backup.volumeSnapshot.onlineConfiguration' stanza",
						Attributes: map[string]schema.Attribute{
							"immediate_checkpoint": schema.BoolAttribute{
								Description:         "Control whether the I/O workload for the backup initial checkpoint will be limited, according to the 'checkpoint_completion_target' setting on the PostgreSQL server. If set to true, an immediate checkpoint will be used, meaning PostgreSQL will complete the checkpoint as soon as possible. 'false' by default.",
								MarkdownDescription: "Control whether the I/O workload for the backup initial checkpoint will be limited, according to the 'checkpoint_completion_target' setting on the PostgreSQL server. If set to true, an immediate checkpoint will be used, meaning PostgreSQL will complete the checkpoint as soon as possible. 'false' by default.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_for_archive": schema.BoolAttribute{
								Description:         "If false, the function will return immediately after the backup is completed, without waiting for WAL to be archived. This behavior is only useful with backup software that independently monitors WAL archiving. Otherwise, WAL required to make the backup consistent might be missing and make the backup useless. By default, or when this parameter is true, pg_backup_stop will wait for WAL to be archived when archiving is enabled. On a standby, this means that it will wait only when archive_mode = always. If write activity on the primary is low, it may be useful to run pg_switch_wal on the primary in order to trigger an immediate segment switch.",
								MarkdownDescription: "If false, the function will return immediately after the backup is completed, without waiting for WAL to be archived. This behavior is only useful with backup software that independently monitors WAL archiving. Otherwise, WAL required to make the backup consistent might be missing and make the backup useless. By default, or when this parameter is true, pg_backup_stop will wait for WAL to be archived when archiving is enabled. On a standby, this means that it will wait only when archive_mode = always. If write activity on the primary is low, it may be useful to run pg_switch_wal on the primary in order to trigger an immediate segment switch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"plugin_configuration": schema.SingleNestedAttribute{
						Description:         "Configuration parameters passed to the plugin managing this backup",
						MarkdownDescription: "Configuration parameters passed to the plugin managing this backup",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the plugin managing this backup",
								MarkdownDescription: "Name is the name of the plugin managing this backup",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"parameters": schema.MapAttribute{
								Description:         "Parameters are the configuration parameters passed to the backup plugin for this backup",
								MarkdownDescription: "Parameters are the configuration parameters passed to the backup plugin for this backup",
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

					"target": schema.StringAttribute{
						Description:         "The policy to decide which instance should perform this backup. If empty, it defaults to 'cluster.spec.backup.target'. Available options are empty string, 'primary' and 'prefer-standby'. 'primary' to have backups run always on primary instances, 'prefer-standby' to have backups run preferably on the most updated standby, if available.",
						MarkdownDescription: "The policy to decide which instance should perform this backup. If empty, it defaults to 'cluster.spec.backup.target'. Available options are empty string, 'primary' and 'prefer-standby'. 'primary' to have backups run always on primary instances, 'prefer-standby' to have backups run preferably on the most updated standby, if available.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("primary", "prefer-standby"),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *PostgresqlCnpgIoBackupV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_postgresql_cnpg_io_backup_v1_manifest")

	var model PostgresqlCnpgIoBackupV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("postgresql.cnpg.io/v1")
	model.Kind = pointer.String("Backup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
