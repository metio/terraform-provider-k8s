/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package oracle_db_anthosapis_com_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OracleDbAnthosapisComBackupScheduleV1Alpha1Manifest{}
)

func NewOracleDbAnthosapisComBackupScheduleV1Alpha1Manifest() datasource.DataSource {
	return &OracleDbAnthosapisComBackupScheduleV1Alpha1Manifest{}
}

type OracleDbAnthosapisComBackupScheduleV1Alpha1Manifest struct{}

type OracleDbAnthosapisComBackupScheduleV1Alpha1ManifestData struct {
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
		BackupLabels          *map[string]string `tfsdk:"backup_labels" json:"backupLabels,omitempty"`
		BackupRetentionPolicy *struct {
			BackupRetention *int64 `tfsdk:"backup_retention" json:"backupRetention,omitempty"`
		} `tfsdk:"backup_retention_policy" json:"backupRetentionPolicy,omitempty"`
		BackupSpec *struct {
			BackupItems         *[]string `tfsdk:"backup_items" json:"backupItems,omitempty"`
			Backupset           *bool     `tfsdk:"backupset" json:"backupset,omitempty"`
			CheckLogical        *bool     `tfsdk:"check_logical" json:"checkLogical,omitempty"`
			Compressed          *bool     `tfsdk:"compressed" json:"compressed,omitempty"`
			Dop                 *int64    `tfsdk:"dop" json:"dop,omitempty"`
			Filesperset         *int64    `tfsdk:"filesperset" json:"filesperset,omitempty"`
			GcsDir              *string   `tfsdk:"gcs_dir" json:"gcsDir,omitempty"`
			GcsPath             *string   `tfsdk:"gcs_path" json:"gcsPath,omitempty"`
			Instance            *string   `tfsdk:"instance" json:"instance,omitempty"`
			KeepDataOnDeletion  *bool     `tfsdk:"keep_data_on_deletion" json:"keepDataOnDeletion,omitempty"`
			Level               *int64    `tfsdk:"level" json:"level,omitempty"`
			LocalPath           *string   `tfsdk:"local_path" json:"localPath,omitempty"`
			Mode                *string   `tfsdk:"mode" json:"mode,omitempty"`
			SectionSize         *string   `tfsdk:"section_size" json:"sectionSize,omitempty"`
			SubType             *string   `tfsdk:"sub_type" json:"subType,omitempty"`
			TimeLimitMinutes    *int64    `tfsdk:"time_limit_minutes" json:"timeLimitMinutes,omitempty"`
			Type                *string   `tfsdk:"type" json:"type,omitempty"`
			VolumeSnapshotClass *string   `tfsdk:"volume_snapshot_class" json:"volumeSnapshotClass,omitempty"`
		} `tfsdk:"backup_spec" json:"backupSpec,omitempty"`
		Schedule                *string `tfsdk:"schedule" json:"schedule,omitempty"`
		StartingDeadlineSeconds *int64  `tfsdk:"starting_deadline_seconds" json:"startingDeadlineSeconds,omitempty"`
		Suspend                 *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OracleDbAnthosapisComBackupScheduleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest"
}

func (r *OracleDbAnthosapisComBackupScheduleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupSchedule is the Schema for the backupschedules API.",
		MarkdownDescription: "BackupSchedule is the Schema for the backupschedules API.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"backup_labels": schema.MapAttribute{
						Description:         "BackupLabels define the desired labels that scheduled backups will be created with.",
						MarkdownDescription: "BackupLabels define the desired labels that scheduled backups will be created with.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_retention_policy": schema.SingleNestedAttribute{
						Description:         "BackupRetentionPolicy is the policy used to trigger automatic deletion of backups produced from this BackupSchedule.",
						MarkdownDescription: "BackupRetentionPolicy is the policy used to trigger automatic deletion of backups produced from this BackupSchedule.",
						Attributes: map[string]schema.Attribute{
							"backup_retention": schema.Int64Attribute{
								Description:         "BackupRetention is the number of successful backups to keep around. The default is 7. A value of 0 means 'do not delete backups based on count'. Max of 512 allows for ~21 days of hourly backups or ~1.4 years of daily backups.",
								MarkdownDescription: "BackupRetention is the number of successful backups to keep around. The default is 7. A value of 0 means 'do not delete backups based on count'. Max of 512 allows for ~21 days of hourly backups or ~1.4 years of daily backups.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(512),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup_spec": schema.SingleNestedAttribute{
						Description:         "BackupSpec defines the Backup that will be created on the provided schedule.",
						MarkdownDescription: "BackupSpec defines the Backup that will be created on the provided schedule.",
						Attributes: map[string]schema.Attribute{
							"backup_items": schema.ListAttribute{
								Description:         "For a Physical backup this slice can be used to indicate what PDBs, schemas, tablespaces or tables to back up.",
								MarkdownDescription: "For a Physical backup this slice can be used to indicate what PDBs, schemas, tablespaces or tables to back up.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backupset": schema.BoolAttribute{
								Description:         "For a Physical backup the choices are Backupset and Image Copies. Backupset is the default, but if Image Copies are required, flip this flag to false.",
								MarkdownDescription: "For a Physical backup the choices are Backupset and Image Copies. Backupset is the default, but if Image Copies are required, flip this flag to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"check_logical": schema.BoolAttribute{
								Description:         "For a Physical backup, optionally turn on an additional 'check logical' option. The default is off.",
								MarkdownDescription: "For a Physical backup, optionally turn on an additional 'check logical' option. The default is off.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"compressed": schema.BoolAttribute{
								Description:         "For a Physical backup, optionally turn on compression, by flipping this flag to true. The default is false.",
								MarkdownDescription: "For a Physical backup, optionally turn on compression, by flipping this flag to true. The default is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dop": schema.Int64Attribute{
								Description:         "For a Physical backup, optionally indicate a degree of parallelism also known as DOP.",
								MarkdownDescription: "For a Physical backup, optionally indicate a degree of parallelism also known as DOP.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(100),
								},
							},

							"filesperset": schema.Int64Attribute{
								Description:         "For a Physical backup, optionally specify filesperset. The default depends on a type of backup, generally 64.",
								MarkdownDescription: "For a Physical backup, optionally specify filesperset. The default depends on a type of backup, generally 64.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gcs_dir": schema.StringAttribute{
								Description:         "Similar to GcsPath but specify a Gcs directory. The backup sets of physical backup will be transferred to this GcsDir under a folder named .backup.Spec.Name. This field is usually set in .backupSchedule.Spec.backSpec to specify a GcsDir which all scheduled backups will be uploaded to. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
								MarkdownDescription: "Similar to GcsPath but specify a Gcs directory. The backup sets of physical backup will be transferred to this GcsDir under a folder named .backup.Spec.Name. This field is usually set in .backupSchedule.Spec.backSpec to specify a GcsDir which all scheduled backups will be uploaded to. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^gs:\/\/.+$`), ""),
								},
							},

							"gcs_path": schema.StringAttribute{
								Description:         "If set up ahead of time, the backup sets of a physical backup can be optionally transferred to a GCS bucket. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
								MarkdownDescription: "If set up ahead of time, the backup sets of a physical backup can be optionally transferred to a GCS bucket. A user is to ensure proper write access to the bucket from within the Oracle Operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^gs:\/\/.+$`), ""),
								},
							},

							"instance": schema.StringAttribute{
								Description:         "Instance is a name of an instance to take a backup for.",
								MarkdownDescription: "Instance is a name of an instance to take a backup for.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"keep_data_on_deletion": schema.BoolAttribute{
								Description:         "KeepDataOnDeletion defines whether to keep backup data when backup resource is removed. The default value is false.",
								MarkdownDescription: "KeepDataOnDeletion defines whether to keep backup data when backup resource is removed. The default value is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"level": schema.Int64Attribute{
								Description:         "For a Physical backup, optionally specify an incremental level. The default is 0 (the whole database).",
								MarkdownDescription: "For a Physical backup, optionally specify an incremental level. The default is 0 (the whole database).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"local_path": schema.StringAttribute{
								Description:         "For a Physical backup, optionally specify a local backup dir. If omitted, /u03/app/oracle/rman is assumed.",
								MarkdownDescription: "For a Physical backup, optionally specify a local backup dir. If omitted, /u03/app/oracle/rman is assumed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "Mode specifies how this backup will be managed by the operator. if it is not set, the operator tries to create a backup based on the specifications. if it is set to VerifyExists, the operator verifies the existence of a backup.",
								MarkdownDescription: "Mode specifies how this backup will be managed by the operator. if it is not set, the operator tries to create a backup based on the specifications. if it is set to VerifyExists, the operator verifies the existence of a backup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("VerifyExists"),
								},
							},

							"section_size": schema.StringAttribute{
								Description:         "For a Physical backup, optionally specify a section size in various units (K M G).",
								MarkdownDescription: "For a Physical backup, optionally specify a section size in various units (K M G).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sub_type": schema.StringAttribute{
								Description:         "Backup sub-type, which is only relevant for a Physical backup type (e.g. RMAN). If omitted, the default of Instance(Level) is assumed. Supported options at this point are: Instance or Database level backups.",
								MarkdownDescription: "Backup sub-type, which is only relevant for a Physical backup type (e.g. RMAN). If omitted, the default of Instance(Level) is assumed. Supported options at this point are: Instance or Database level backups.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Instance", "Database", "Tablespace", "Datafile"),
								},
							},

							"time_limit_minutes": schema.Int64Attribute{
								Description:         "For a Physical backup, optionally specify the time threshold. If a threshold is reached, the backup request would time out and error out. The threshold is expressed in minutes. Don't include the unit (minutes), just the integer.",
								MarkdownDescription: "For a Physical backup, optionally specify the time threshold. If a threshold is reached, the backup request would time out and error out. The threshold is expressed in minutes. Don't include the unit (minutes), just the integer.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type describes a type of a backup to take. Immutable. Available options are: - Snapshot: storage level disk snapshot. - Physical: database engine specific backup that relies on a redo stream / continuous archiving (WAL) and may allow a PITR. Examples include pg_backup, pgBackRest, mysqlbackup. A Physical backup may be file based or database block based (e.g. Oracle RMAN). - Logical: database engine specific backup that relies on running SQL statements, e.g. mysqldump, pg_dump, expdp. If not specified, the default of Snapshot is assumed.",
								MarkdownDescription: "Type describes a type of a backup to take. Immutable. Available options are: - Snapshot: storage level disk snapshot. - Physical: database engine specific backup that relies on a redo stream / continuous archiving (WAL) and may allow a PITR. Examples include pg_backup, pgBackRest, mysqlbackup. A Physical backup may be file based or database block based (e.g. Oracle RMAN). - Logical: database engine specific backup that relies on running SQL statements, e.g. mysqldump, pg_dump, expdp. If not specified, the default of Snapshot is assumed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Snapshot", "Physical", "Logical"),
								},
							},

							"volume_snapshot_class": schema.StringAttribute{
								Description:         "VolumeSnapshotClass points to a particular CSI driver and is used for taking a volume snapshot. If requested here at the Backup level, this setting overrides the platform default as well as the default set via the Config (global user preferences).",
								MarkdownDescription: "VolumeSnapshotClass points to a particular CSI driver and is used for taking a volume snapshot. If requested here at the Backup level, this setting overrides the platform default as well as the default set via the Config (global user preferences).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"schedule": schema.StringAttribute{
						Description:         "Schedule is a cron-style expression of the schedule on which Backup will be created. For allowed syntax, see en.wikipedia.org/wiki/Cron and godoc.org/github.com/robfig/cron.",
						MarkdownDescription: "Schedule is a cron-style expression of the schedule on which Backup will be created. For allowed syntax, see en.wikipedia.org/wiki/Cron and godoc.org/github.com/robfig/cron.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"starting_deadline_seconds": schema.Int64Attribute{
						Description:         "StartingDeadlineSeconds is an optional deadline in seconds for starting the backup creation if it misses scheduled time for any reason. The default is 30 seconds.",
						MarkdownDescription: "StartingDeadlineSeconds is an optional deadline in seconds for starting the backup creation if it misses scheduled time for any reason. The default is 30 seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend tells the controller to suspend operations - both creation of new Backup and retention actions. This will not have any effect on backups currently in progress. Default is false.",
						MarkdownDescription: "Suspend tells the controller to suspend operations - both creation of new Backup and retention actions. This will not have any effect on backups currently in progress. Default is false.",
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
	}
}

func (r *OracleDbAnthosapisComBackupScheduleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_oracle_db_anthosapis_com_backup_schedule_v1alpha1_manifest")

	var model OracleDbAnthosapisComBackupScheduleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("oracle.db.anthosapis.com/v1alpha1")
	model.Kind = pointer.String("BackupSchedule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
