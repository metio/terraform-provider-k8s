/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

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
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &HazelcastComCronHotBackupV1Alpha1Manifest{}
)

func NewHazelcastComCronHotBackupV1Alpha1Manifest() datasource.DataSource {
	return &HazelcastComCronHotBackupV1Alpha1Manifest{}
}

type HazelcastComCronHotBackupV1Alpha1Manifest struct{}

type HazelcastComCronHotBackupV1Alpha1ManifestData struct {
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
		FailedHotBackupsHistoryLimit *int64 `tfsdk:"failed_hot_backups_history_limit" json:"failedHotBackupsHistoryLimit,omitempty"`
		HotBackupTemplate            *struct {
			Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec     *struct {
				BucketURI             *string `tfsdk:"bucket_uri" json:"bucketURI,omitempty"`
				HazelcastResourceName *string `tfsdk:"hazelcast_resource_name" json:"hazelcastResourceName,omitempty"`
				Secret                *string `tfsdk:"secret" json:"secret,omitempty"`
				SecretName            *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"hot_backup_template" json:"hotBackupTemplate,omitempty"`
		Schedule                         *string `tfsdk:"schedule" json:"schedule,omitempty"`
		SuccessfulHotBackupsHistoryLimit *int64  `tfsdk:"successful_hot_backups_history_limit" json:"successfulHotBackupsHistoryLimit,omitempty"`
		Suspend                          *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *HazelcastComCronHotBackupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hazelcast_com_cron_hot_backup_v1alpha1_manifest"
}

func (r *HazelcastComCronHotBackupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CronHotBackup is the Schema for the cronhotbackups API",
		MarkdownDescription: "CronHotBackup is the Schema for the cronhotbackups API",
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
				Description:         "CronHotBackupSpec defines the desired state of CronHotBackup",
				MarkdownDescription: "CronHotBackupSpec defines the desired state of CronHotBackup",
				Attributes: map[string]schema.Attribute{
					"failed_hot_backups_history_limit": schema.Int64Attribute{
						Description:         "The number of failed finished hot backups to retain.",
						MarkdownDescription: "The number of failed finished hot backups to retain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"hot_backup_template": schema.SingleNestedAttribute{
						Description:         "Specifies the hot backup that will be created when executing a CronHotBackup.",
						MarkdownDescription: "Specifies the hot backup that will be created when executing a CronHotBackup.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.MapAttribute{
								Description:         "Standard object's metadata of the hot backups created from this template.",
								MarkdownDescription: "Standard object's metadata of the hot backups created from this template.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "Specification of the desired behavior of the hot backup.",
								MarkdownDescription: "Specification of the desired behavior of the hot backup.",
								Attributes: map[string]schema.Attribute{
									"bucket_uri": schema.StringAttribute{
										Description:         "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										MarkdownDescription: "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hazelcast_resource_name": schema.StringAttribute{
										Description:         "HazelcastResourceName defines the name of the Hazelcast resource",
										MarkdownDescription: "HazelcastResourceName defines the name of the Hazelcast resource",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"secret": schema.StringAttribute{
										Description:         "secret is a deprecated alias for secretName.",
										MarkdownDescription: "secret is a deprecated alias for secretName.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "Name of the secret with credentials for cloud providers.",
										MarkdownDescription: "Name of the secret with credentials for cloud providers.",
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
						Required: true,
						Optional: false,
						Computed: false,
					},

					"schedule": schema.StringAttribute{
						Description:         "Schedule contains a crontab-like expression that defines the schedule in which HotBackup will be started. If the Schedule is empty the HotBackup will start only once when applied. --- Several pre-defined schedules in place of a cron expression can be used. Entry                  | Description                                | Equivalent To -----                  | -----------                                | ------------- @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 * @monthly               | Run once a month, midnight, first of month | 0 0 1 * * @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0 @daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * * @hourly                | Run once an hour, beginning of hour        | 0 * * * *",
						MarkdownDescription: "Schedule contains a crontab-like expression that defines the schedule in which HotBackup will be started. If the Schedule is empty the HotBackup will start only once when applied. --- Several pre-defined schedules in place of a cron expression can be used. Entry                  | Description                                | Equivalent To -----                  | -----------                                | ------------- @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 * @monthly               | Run once a month, midnight, first of month | 0 0 1 * * @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0 @daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * * @hourly                | Run once an hour, beginning of hour        | 0 * * * *",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"successful_hot_backups_history_limit": schema.Int64Attribute{
						Description:         "The number of successful finished hot backups to retain.",
						MarkdownDescription: "The number of successful finished hot backups to retain.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"suspend": schema.BoolAttribute{
						Description:         "When true, CronHotBackup will stop creating HotBackup CRs until it is disabled",
						MarkdownDescription: "When true, CronHotBackup will stop creating HotBackup CRs until it is disabled",
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
	}
}

func (r *HazelcastComCronHotBackupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_hazelcast_com_cron_hot_backup_v1alpha1_manifest")

	var model HazelcastComCronHotBackupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	model.Kind = pointer.String("CronHotBackup")

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
