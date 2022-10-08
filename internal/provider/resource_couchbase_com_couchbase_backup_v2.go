/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type CouchbaseComCouchbaseBackupV2Resource struct{}

var (
	_ resource.Resource = (*CouchbaseComCouchbaseBackupV2Resource)(nil)
)

type CouchbaseComCouchbaseBackupV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CouchbaseComCouchbaseBackupV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AutoScaling *struct {
			IncrementPercent *int64 `tfsdk:"increment_percent" yaml:"incrementPercent,omitempty"`

			Limit *string `tfsdk:"limit" yaml:"limit,omitempty"`

			ThresholdPercent *int64 `tfsdk:"threshold_percent" yaml:"thresholdPercent,omitempty"`
		} `tfsdk:"auto_scaling" yaml:"autoScaling,omitempty"`

		BackoffLimit *int64 `tfsdk:"backoff_limit" yaml:"backoffLimit,omitempty"`

		BackupRetention *string `tfsdk:"backup_retention" yaml:"backupRetention,omitempty"`

		Data *struct {
			Exclude *[]string `tfsdk:"exclude" yaml:"exclude,omitempty"`

			Include *[]string `tfsdk:"include" yaml:"include,omitempty"`
		} `tfsdk:"data" yaml:"data,omitempty"`

		FailedJobsHistoryLimit *int64 `tfsdk:"failed_jobs_history_limit" yaml:"failedJobsHistoryLimit,omitempty"`

		Full *struct {
			Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`
		} `tfsdk:"full" yaml:"full,omitempty"`

		Incremental *struct {
			Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`
		} `tfsdk:"incremental" yaml:"incremental,omitempty"`

		LogRetention *string `tfsdk:"log_retention" yaml:"logRetention,omitempty"`

		S3bucket *string `tfsdk:"s3bucket" yaml:"s3bucket,omitempty"`

		Services *struct {
			Analytics *bool `tfsdk:"analytics" yaml:"analytics,omitempty"`

			BucketConfig *bool `tfsdk:"bucket_config" yaml:"bucketConfig,omitempty"`

			BucketQuery *bool `tfsdk:"bucket_query" yaml:"bucketQuery,omitempty"`

			ClusterAnalytics *bool `tfsdk:"cluster_analytics" yaml:"clusterAnalytics,omitempty"`

			ClusterQuery *bool `tfsdk:"cluster_query" yaml:"clusterQuery,omitempty"`

			Data *bool `tfsdk:"data" yaml:"data,omitempty"`

			Eventing *bool `tfsdk:"eventing" yaml:"eventing,omitempty"`

			FtsAliases *bool `tfsdk:"fts_aliases" yaml:"ftsAliases,omitempty"`

			FtsIndexes *bool `tfsdk:"fts_indexes" yaml:"ftsIndexes,omitempty"`

			GsIndexes *bool `tfsdk:"gs_indexes" yaml:"gsIndexes,omitempty"`

			Views *bool `tfsdk:"views" yaml:"views,omitempty"`
		} `tfsdk:"services" yaml:"services,omitempty"`

		Size *string `tfsdk:"size" yaml:"size,omitempty"`

		StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

		Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

		SuccessfulJobsHistoryLimit *int64 `tfsdk:"successful_jobs_history_limit" yaml:"successfulJobsHistoryLimit,omitempty"`

		Threads *int64 `tfsdk:"threads" yaml:"threads,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCouchbaseComCouchbaseBackupV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseBackupV2Resource{}
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_couchbase_com_couchbase_backup_v2"
}

func (r *CouchbaseComCouchbaseBackupV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CouchbaseBackup allows automatic backup of all data from a Couchbase cluster into persistent storage.",
		MarkdownDescription: "CouchbaseBackup allows automatic backup of all data from a Couchbase cluster into persistent storage.",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "CouchbaseBackupSpec is allows the specification of how a Couchbase backup is configured, including when backups are performed, how long they are retained for, and where they are backed up to.",
				MarkdownDescription: "CouchbaseBackupSpec is allows the specification of how a Couchbase backup is configured, including when backups are performed, how long they are retained for, and where they are backed up to.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"auto_scaling": {
						Description:         "AutoScaling allows the volume size to be dynamically increased. When specified, the backup volume will start with an initial size as defined by 'spec.size', and increase as required.",
						MarkdownDescription: "AutoScaling allows the volume size to be dynamically increased. When specified, the backup volume will start with an initial size as defined by 'spec.size', and increase as required.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"increment_percent": {
								Description:         "IncrementPercent controls how much the volume is increased each time the threshold is exceeded, upto a maximum as defined by the limit. This field defaults to 20 if not specified.",
								MarkdownDescription: "IncrementPercent controls how much the volume is increased each time the threshold is exceeded, upto a maximum as defined by the limit. This field defaults to 20 if not specified.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"limit": {
								Description:         "Limit imposes a hard limit on the size we can autoscale to.  When not specified no bounds are imposed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "Limit imposes a hard limit on the size we can autoscale to.  When not specified no bounds are imposed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"threshold_percent": {
								Description:         "ThresholdPercent determines the point at which a volume is autoscaled. This represents the percentage of free space remaining on the volume, when less than this threshold, it will trigger a volume expansion. For example, if the volume is 100Gi, and the threshold 20%, then a resize will be triggered when the used capacity exceeds 80Gi, and free space is less than 20Gi.  This field defaults to 20 if not specified.",
								MarkdownDescription: "ThresholdPercent determines the point at which a volume is autoscaled. This represents the percentage of free space remaining on the volume, when less than this threshold, it will trigger a volume expansion. For example, if the volume is 100Gi, and the threshold 20%, then a resize will be triggered when the used capacity exceeds 80Gi, and free space is less than 20Gi.  This field defaults to 20 if not specified.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(99),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backoff_limit": {
						Description:         "Number of times a backup job should try to execute. Once it hits the BackoffLimit it will not run until the next scheduled job.",
						MarkdownDescription: "Number of times a backup job should try to execute. Once it hits the BackoffLimit it will not run until the next scheduled job.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup_retention": {
						Description:         "Number of hours to hold backups for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "Number of hours to hold backups for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"data": {
						Description:         "Data allows control over what key-value/document data is included in the backup.  By default, all data is included.  Modifications to this field will only take effect on the next full backup.",
						MarkdownDescription: "Data allows control over what key-value/document data is included in the backup.  By default, all data is included.  Modifications to this field will only take effect on the next full backup.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"exclude": {
								Description:         "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",
								MarkdownDescription: "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"include": {
								Description:         "Include defines the buckets, scopes or collections that are included in the backup. When this field is set, it implies that by default nothing will be backed up, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",
								MarkdownDescription: "Include defines the buckets, scopes or collections that are included in the backup. When this field is set, it implies that by default nothing will be backed up, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",

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

					"failed_jobs_history_limit": {
						Description:         "Amount of failed jobs to keep.",
						MarkdownDescription: "Amount of failed jobs to keep.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(0),
						},
					},

					"full": {
						Description:         "Full is the schedule on when to take full backups. Used in Full/Incremental and FullOnly backup strategies.",
						MarkdownDescription: "Full is the schedule on when to take full backups. Used in Full/Incremental and FullOnly backup strategies.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"schedule": {
								Description:         "Schedule takes a cron schedule in string format.",
								MarkdownDescription: "Schedule takes a cron schedule in string format.",

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

					"incremental": {
						Description:         "Incremental is the schedule on when to take incremental backups. Used in Full/Incremental backup strategies.",
						MarkdownDescription: "Incremental is the schedule on when to take incremental backups. Used in Full/Incremental backup strategies.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"schedule": {
								Description:         "Schedule takes a cron schedule in string format.",
								MarkdownDescription: "Schedule takes a cron schedule in string format.",

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

					"log_retention": {
						Description:         "Number of hours to hold script logs for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "Number of hours to hold script logs for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"s3bucket": {
						Description:         "Name of S3 bucket to backup to. If non-empty this overrides local backup.",
						MarkdownDescription: "Name of S3 bucket to backup to. If non-empty this overrides local backup.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"services": {
						Description:         "Services allows control over what services are included in the backup. By default, all service data and metadata are included.  Modifications to this field will only take effect on the next full backup.",
						MarkdownDescription: "Services allows control over what services are included in the backup. By default, all service data and metadata are included.  Modifications to this field will only take effect on the next full backup.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"analytics": {
								Description:         "Analytics enables the backup of analytics data. This field defaults to 'true'.",
								MarkdownDescription: "Analytics enables the backup of analytics data. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bucket_config": {
								Description:         "BucketConfig enables the backup of bucket configuration. This field defaults to 'true'.",
								MarkdownDescription: "BucketConfig enables the backup of bucket configuration. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bucket_query": {
								Description:         "BucketQuery enables the backup of query metadata for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "BucketQuery enables the backup of query metadata for all buckets. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_analytics": {
								Description:         "ClusterAnalytics enables the backup of cluster-wide analytics data, for example synonyms. This field defaults to 'true'.",
								MarkdownDescription: "ClusterAnalytics enables the backup of cluster-wide analytics data, for example synonyms. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_query": {
								Description:         "ClusterQuery enables the backup of cluster level query metadata. This field defaults to 'true'.",
								MarkdownDescription: "ClusterQuery enables the backup of cluster level query metadata. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data": {
								Description:         "Data enables the backup of key-value data/documents for all buckets. This can be further refined with the couchbasebackups.spec.data configuration. This field defaults to 'true'.",
								MarkdownDescription: "Data enables the backup of key-value data/documents for all buckets. This can be further refined with the couchbasebackups.spec.data configuration. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"eventing": {
								Description:         "Eventing enables the backup of eventing service metadata. This field defaults to 'true'.",
								MarkdownDescription: "Eventing enables the backup of eventing service metadata. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fts_aliases": {
								Description:         "FTSAliases enables the backup of full-text search alias definitions. This field defaults to 'true'.",
								MarkdownDescription: "FTSAliases enables the backup of full-text search alias definitions. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fts_indexes": {
								Description:         "FTSIndexes enables the backup of full-text search index definitions for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "FTSIndexes enables the backup of full-text search index definitions for all buckets. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gs_indexes": {
								Description:         "GSIndexes enables the backup of global secondary index definitions for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "GSIndexes enables the backup of global secondary index definitions for all buckets. This field defaults to 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"views": {
								Description:         "Views enables the backup of view definitions for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "Views enables the backup of view definitions for all buckets. This field defaults to 'true'.",

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

					"size": {
						Description:         "Size allows the specification of a backup persistent volume, when using volume based backup. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						MarkdownDescription: "Size allows the specification of a backup persistent volume, when using volume based backup. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_class_name": {
						Description:         "Name of StorageClass to use.",
						MarkdownDescription: "Name of StorageClass to use.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"strategy": {
						Description:         "Strategy defines how to perform backups.  'full_only' will only perform full backups, and you must define a schedule in the 'spec.full' field.  'full_incremental' will perform periodic full backups, and incremental backups in between.  You must define full and incremental schedules in the 'spec.full' and 'spec.incremental' fields respectively.  Care should be taken to ensure full and incremental schedules do not overlap, taking into account the backup time, as this will cause failures as the jobs attempt to mount the same backup volume. This field default to 'full_incremental'. Info: https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-strategies.html",
						MarkdownDescription: "Strategy defines how to perform backups.  'full_only' will only perform full backups, and you must define a schedule in the 'spec.full' field.  'full_incremental' will perform periodic full backups, and incremental backups in between.  You must define full and incremental schedules in the 'spec.full' and 'spec.incremental' fields respectively.  Care should be taken to ensure full and incremental schedules do not overlap, taking into account the backup time, as this will cause failures as the jobs attempt to mount the same backup volume. This field default to 'full_incremental'. Info: https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-strategies.html",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"successful_jobs_history_limit": {
						Description:         "Amount of successful jobs to keep.",
						MarkdownDescription: "Amount of successful jobs to keep.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(0),
						},
					},

					"threads": {
						Description:         "How many threads to use during the backup.  This field defaults to 1.",
						MarkdownDescription: "How many threads to use during the backup.  This field defaults to 1.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_backup_v2")

	var state CouchbaseComCouchbaseBackupV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseBackupV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseBackup")

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

func (r *CouchbaseComCouchbaseBackupV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_backup_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_backup_v2")

	var state CouchbaseComCouchbaseBackupV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseBackupV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseBackup")

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

func (r *CouchbaseComCouchbaseBackupV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_backup_v2")
	// NO-OP: Terraform removes the state automatically for us
}
