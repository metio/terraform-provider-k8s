/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package dataprotection_kubeblocks_io_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &DataprotectionKubeblocksIoBackupScheduleV1Alpha1Manifest{}
)

func NewDataprotectionKubeblocksIoBackupScheduleV1Alpha1Manifest() datasource.DataSource {
	return &DataprotectionKubeblocksIoBackupScheduleV1Alpha1Manifest{}
}

type DataprotectionKubeblocksIoBackupScheduleV1Alpha1Manifest struct{}

type DataprotectionKubeblocksIoBackupScheduleV1Alpha1ManifestData struct {
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
		BackupPolicyName *string `tfsdk:"backup_policy_name" json:"backupPolicyName,omitempty"`
		Schedules        *[]struct {
			BackupMethod    *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
			CronExpression  *string `tfsdk:"cron_expression" json:"cronExpression,omitempty"`
			Enabled         *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			RetentionPeriod *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		} `tfsdk:"schedules" json:"schedules,omitempty"`
		StartingDeadlineMinutes *int64 `tfsdk:"starting_deadline_minutes" json:"startingDeadlineMinutes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataprotectionKubeblocksIoBackupScheduleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_dataprotection_kubeblocks_io_backup_schedule_v1alpha1_manifest"
}

func (r *DataprotectionKubeblocksIoBackupScheduleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BackupSchedule is the Schema for the backupschedules API.",
		MarkdownDescription: "BackupSchedule is the Schema for the backupschedules API.",
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
				Description:         "BackupScheduleSpec defines the desired state of BackupSchedule.",
				MarkdownDescription: "BackupScheduleSpec defines the desired state of BackupSchedule.",
				Attributes: map[string]schema.Attribute{
					"backup_policy_name": schema.StringAttribute{
						Description:         "Specifies the backupPolicy to be applied for the 'schedules'.",
						MarkdownDescription: "Specifies the backupPolicy to be applied for the 'schedules'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"schedules": schema.ListNestedAttribute{
						Description:         "Defines the list of backup schedules.",
						MarkdownDescription: "Defines the list of backup schedules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backup_method": schema.StringAttribute{
									Description:         "Specifies the backup method name that is defined in backupPolicy.",
									MarkdownDescription: "Specifies the backup method name that is defined in backupPolicy.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"cron_expression": schema.StringAttribute{
									Description:         "Specifies the cron expression for the schedule. The timezone is in UTC. see https://en.wikipedia.org/wiki/Cron.",
									MarkdownDescription: "Specifies the cron expression for the schedule. The timezone is in UTC. see https://en.wikipedia.org/wiki/Cron.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"enabled": schema.BoolAttribute{
									Description:         "Specifies whether the backup schedule is enabled or not.",
									MarkdownDescription: "Specifies whether the backup schedule is enabled or not.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"retention_period": schema.StringAttribute{
									Description:         "Determines the duration for which the backup should be kept. KubeBlocks will remove all backups that are older than the RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  You can also combine the above durations. For example: 30d12h30m",
									MarkdownDescription: "Determines the duration for which the backup should be kept. KubeBlocks will remove all backups that are older than the RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  You can also combine the above durations. For example: 30d12h30m",
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

					"starting_deadline_minutes": schema.Int64Attribute{
						Description:         "Defines the deadline in minutes for starting the backup workload if it misses its scheduled time for any reason.",
						MarkdownDescription: "Defines the deadline in minutes for starting the backup workload if it misses its scheduled time for any reason.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(1440),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *DataprotectionKubeblocksIoBackupScheduleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_dataprotection_kubeblocks_io_backup_schedule_v1alpha1_manifest")

	var model DataprotectionKubeblocksIoBackupScheduleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("dataprotection.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("BackupSchedule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
