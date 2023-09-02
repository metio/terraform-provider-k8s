/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package hazelcast_com_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &HazelcastComCronHotBackupV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &HazelcastComCronHotBackupV1Alpha1DataSource{}
)

func NewHazelcastComCronHotBackupV1Alpha1DataSource() datasource.DataSource {
	return &HazelcastComCronHotBackupV1Alpha1DataSource{}
}

type HazelcastComCronHotBackupV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type HazelcastComCronHotBackupV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *HazelcastComCronHotBackupV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_hazelcast_com_cron_hot_backup_v1alpha1"
}

func (r *HazelcastComCronHotBackupV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "Specification of the desired behavior of the hot backup.",
								MarkdownDescription: "Specification of the desired behavior of the hot backup.",
								Attributes: map[string]schema.Attribute{
									"bucket_uri": schema.StringAttribute{
										Description:         "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										MarkdownDescription: "URL of the bucket to download HotBackup folders. AWS S3, GCP Bucket and Azure Blob storage buckets are supported. Example bucket URIs: - AWS S3     -> s3://bucket-name/path/to/folder - GCP Bucket -> gs://bucket-name/path/to/folder - Azure Blob -> azblob://bucket-name/path/to/folder",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"hazelcast_resource_name": schema.StringAttribute{
										Description:         "HazelcastResourceName defines the name of the Hazelcast resource",
										MarkdownDescription: "HazelcastResourceName defines the name of the Hazelcast resource",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret": schema.StringAttribute{
										Description:         "secret is a deprecated alias for secretName.",
										MarkdownDescription: "secret is a deprecated alias for secretName.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "Name of the secret with credentials for cloud providers.",
										MarkdownDescription: "Name of the secret with credentials for cloud providers.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"schedule": schema.StringAttribute{
						Description:         "Schedule contains a crontab-like expression that defines the schedule in which HotBackup will be started. If the Schedule is empty the HotBackup will start only once when applied. --- Several pre-defined schedules in place of a cron expression can be used. Entry                  | Description                                | Equivalent To -----                  | -----------                                | ------------- @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 * @monthly               | Run once a month, midnight, first of month | 0 0 1 * * @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0 @daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * * @hourly                | Run once an hour, beginning of hour        | 0 * * * *",
						MarkdownDescription: "Schedule contains a crontab-like expression that defines the schedule in which HotBackup will be started. If the Schedule is empty the HotBackup will start only once when applied. --- Several pre-defined schedules in place of a cron expression can be used. Entry                  | Description                                | Equivalent To -----                  | -----------                                | ------------- @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 1 1 * @monthly               | Run once a month, midnight, first of month | 0 0 1 * * @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 * * 0 @daily (or @midnight)  | Run once a day, midnight                   | 0 0 * * * @hourly                | Run once an hour, beginning of hour        | 0 * * * *",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"successful_hot_backups_history_limit": schema.Int64Attribute{
						Description:         "The number of successful finished hot backups to retain.",
						MarkdownDescription: "The number of successful finished hot backups to retain.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"suspend": schema.BoolAttribute{
						Description:         "When true, CronHotBackup will stop creating HotBackup CRs until it is disabled",
						MarkdownDescription: "When true, CronHotBackup will stop creating HotBackup CRs until it is disabled",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *HazelcastComCronHotBackupV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *HazelcastComCronHotBackupV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_hazelcast_com_cron_hot_backup_v1alpha1")

	var data HazelcastComCronHotBackupV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "hazelcast.com", Version: "v1alpha1", Resource: "CronHotBackup"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse HazelcastComCronHotBackupV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("hazelcast.com/v1alpha1")
	data.Kind = pointer.String("CronHotBackup")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
