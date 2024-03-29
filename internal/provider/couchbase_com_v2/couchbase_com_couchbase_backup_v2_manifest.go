/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

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
	_ datasource.DataSource = &CouchbaseComCouchbaseBackupV2Manifest{}
)

func NewCouchbaseComCouchbaseBackupV2Manifest() datasource.DataSource {
	return &CouchbaseComCouchbaseBackupV2Manifest{}
}

type CouchbaseComCouchbaseBackupV2Manifest struct{}

type CouchbaseComCouchbaseBackupV2ManifestData struct {
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
		AutoScaling *struct {
			IncrementPercent *int64  `tfsdk:"increment_percent" json:"incrementPercent,omitempty"`
			Limit            *string `tfsdk:"limit" json:"limit,omitempty"`
			ThresholdPercent *int64  `tfsdk:"threshold_percent" json:"thresholdPercent,omitempty"`
		} `tfsdk:"auto_scaling" json:"autoScaling,omitempty"`
		BackoffLimit    *int64  `tfsdk:"backoff_limit" json:"backoffLimit,omitempty"`
		BackupRetention *string `tfsdk:"backup_retention" json:"backupRetention,omitempty"`
		Data            *struct {
			Exclude *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
			Include *[]string `tfsdk:"include" json:"include,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		DefaultRecoveryMethod  *string `tfsdk:"default_recovery_method" json:"defaultRecoveryMethod,omitempty"`
		EphemeralVolume        *bool   `tfsdk:"ephemeral_volume" json:"ephemeralVolume,omitempty"`
		FailedJobsHistoryLimit *int64  `tfsdk:"failed_jobs_history_limit" json:"failedJobsHistoryLimit,omitempty"`
		Full                   *struct {
			Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
		} `tfsdk:"full" json:"full,omitempty"`
		Incremental *struct {
			Schedule *string `tfsdk:"schedule" json:"schedule,omitempty"`
		} `tfsdk:"incremental" json:"incremental,omitempty"`
		LogRetention *string `tfsdk:"log_retention" json:"logRetention,omitempty"`
		ObjectStore  *struct {
			Endpoint *struct {
				Secret         *string `tfsdk:"secret" json:"secret,omitempty"`
				Url            *string `tfsdk:"url" json:"url,omitempty"`
				UseVirtualPath *bool   `tfsdk:"use_virtual_path" json:"useVirtualPath,omitempty"`
			} `tfsdk:"endpoint" json:"endpoint,omitempty"`
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
			Uri    *string `tfsdk:"uri" json:"uri,omitempty"`
			UseIAM *bool   `tfsdk:"use_iam" json:"useIAM,omitempty"`
		} `tfsdk:"object_store" json:"objectStore,omitempty"`
		S3bucket *string `tfsdk:"s3bucket" json:"s3bucket,omitempty"`
		Services *struct {
			Analytics        *bool `tfsdk:"analytics" json:"analytics,omitempty"`
			BucketConfig     *bool `tfsdk:"bucket_config" json:"bucketConfig,omitempty"`
			BucketQuery      *bool `tfsdk:"bucket_query" json:"bucketQuery,omitempty"`
			ClusterAnalytics *bool `tfsdk:"cluster_analytics" json:"clusterAnalytics,omitempty"`
			ClusterQuery     *bool `tfsdk:"cluster_query" json:"clusterQuery,omitempty"`
			Data             *bool `tfsdk:"data" json:"data,omitempty"`
			Eventing         *bool `tfsdk:"eventing" json:"eventing,omitempty"`
			FtsAliases       *bool `tfsdk:"fts_aliases" json:"ftsAliases,omitempty"`
			FtsIndexes       *bool `tfsdk:"fts_indexes" json:"ftsIndexes,omitempty"`
			GsIndexes        *bool `tfsdk:"gs_indexes" json:"gsIndexes,omitempty"`
			Views            *bool `tfsdk:"views" json:"views,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		Size                       *string `tfsdk:"size" json:"size,omitempty"`
		StorageClassName           *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
		Strategy                   *string `tfsdk:"strategy" json:"strategy,omitempty"`
		SuccessfulJobsHistoryLimit *int64  `tfsdk:"successful_jobs_history_limit" json:"successfulJobsHistoryLimit,omitempty"`
		Threads                    *int64  `tfsdk:"threads" json:"threads,omitempty"`
		TtlSecondsAfterFinished    *int64  `tfsdk:"ttl_seconds_after_finished" json:"ttlSecondsAfterFinished,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseBackupV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_backup_v2_manifest"
}

func (r *CouchbaseComCouchbaseBackupV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CouchbaseBackup allows automatic backup of all data from a Couchbase cluster into persistent storage.",
		MarkdownDescription: "CouchbaseBackup allows automatic backup of all data from a Couchbase cluster into persistent storage.",
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
				Description:         "CouchbaseBackupSpec is allows the specification of how a Couchbase backup is configured, including when backups are performed, how long they are retained for, and where they are backed up to.",
				MarkdownDescription: "CouchbaseBackupSpec is allows the specification of how a Couchbase backup is configured, including when backups are performed, how long they are retained for, and where they are backed up to.",
				Attributes: map[string]schema.Attribute{
					"auto_scaling": schema.SingleNestedAttribute{
						Description:         "AutoScaling allows the volume size to be dynamically increased. When specified, the backup volume will start with an initial size as defined by 'spec.size', and increase as required.",
						MarkdownDescription: "AutoScaling allows the volume size to be dynamically increased. When specified, the backup volume will start with an initial size as defined by 'spec.size', and increase as required.",
						Attributes: map[string]schema.Attribute{
							"increment_percent": schema.Int64Attribute{
								Description:         "IncrementPercent controls how much the volume is increased each time the threshold is exceeded, upto a maximum as defined by the limit. This field defaults to 20 if not specified.",
								MarkdownDescription: "IncrementPercent controls how much the volume is increased each time the threshold is exceeded, upto a maximum as defined by the limit. This field defaults to 20 if not specified.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"limit": schema.StringAttribute{
								Description:         "Limit imposes a hard limit on the size we can autoscale to.  When not specified no bounds are imposed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								MarkdownDescription: "Limit imposes a hard limit on the size we can autoscale to.  When not specified no bounds are imposed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
								},
							},

							"threshold_percent": schema.Int64Attribute{
								Description:         "ThresholdPercent determines the point at which a volume is autoscaled. This represents the percentage of free space remaining on the volume, when less than this threshold, it will trigger a volume expansion. For example, if the volume is 100Gi, and the threshold 20%, then a resize will be triggered when the used capacity exceeds 80Gi, and free space is less than 20Gi.  This field defaults to 20 if not specified.",
								MarkdownDescription: "ThresholdPercent determines the point at which a volume is autoscaled. This represents the percentage of free space remaining on the volume, when less than this threshold, it will trigger a volume expansion. For example, if the volume is 100Gi, and the threshold 20%, then a resize will be triggered when the used capacity exceeds 80Gi, and free space is less than 20Gi.  This field defaults to 20 if not specified.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(99),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"backoff_limit": schema.Int64Attribute{
						Description:         "Number of times a backup job should try to execute. Once it hits the BackoffLimit it will not run until the next scheduled job.",
						MarkdownDescription: "Number of times a backup job should try to execute. Once it hits the BackoffLimit it will not run until the next scheduled job.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup_retention": schema.StringAttribute{
						Description:         "Number of hours to hold backups for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "Number of hours to hold backups for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"data": schema.SingleNestedAttribute{
						Description:         "Data allows control over what key-value/document data is included in the backup.  By default, all data is included.  Modifications to this field will only take effect on the next full backup.",
						MarkdownDescription: "Data allows control over what key-value/document data is included in the backup.  By default, all data is included.  Modifications to this field will only take effect on the next full backup.",
						Attributes: map[string]schema.Attribute{
							"exclude": schema.ListAttribute{
								Description:         "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",
								MarkdownDescription: "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include": schema.ListAttribute{
								Description:         "Include defines the buckets, scopes or collections that are included in the backup. When this field is set, it implies that by default nothing will be backed up, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",
								MarkdownDescription: "Include defines the buckets, scopes or collections that are included in the backup. When this field is set, it implies that by default nothing will be backed up, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",
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

					"default_recovery_method": schema.StringAttribute{
						Description:         "DefaultRecoveryMethod specifies how cbbackupmgr should recover from broken backup/restore attempts.",
						MarkdownDescription: "DefaultRecoveryMethod specifies how cbbackupmgr should recover from broken backup/restore attempts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "resume", "purge"),
						},
					},

					"ephemeral_volume": schema.BoolAttribute{
						Description:         "EphemeralVolume sets backup to use an ephemeral volume instead of a persistent volume. This is used when backing up to a remote cloud provider, where a persistent volume is not needed.",
						MarkdownDescription: "EphemeralVolume sets backup to use an ephemeral volume instead of a persistent volume. This is used when backing up to a remote cloud provider, where a persistent volume is not needed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failed_jobs_history_limit": schema.Int64Attribute{
						Description:         "Amount of failed jobs to keep.",
						MarkdownDescription: "Amount of failed jobs to keep.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"full": schema.SingleNestedAttribute{
						Description:         "Full is the schedule on when to take full backups. Used in Full/Incremental and FullOnly backup strategies.",
						MarkdownDescription: "Full is the schedule on when to take full backups. Used in Full/Incremental and FullOnly backup strategies.",
						Attributes: map[string]schema.Attribute{
							"schedule": schema.StringAttribute{
								Description:         "Schedule takes a cron schedule in string format.",
								MarkdownDescription: "Schedule takes a cron schedule in string format.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"incremental": schema.SingleNestedAttribute{
						Description:         "Incremental is the schedule on when to take incremental backups. Used in Full/Incremental backup strategies.",
						MarkdownDescription: "Incremental is the schedule on when to take incremental backups. Used in Full/Incremental backup strategies.",
						Attributes: map[string]schema.Attribute{
							"schedule": schema.StringAttribute{
								Description:         "Schedule takes a cron schedule in string format.",
								MarkdownDescription: "Schedule takes a cron schedule in string format.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_retention": schema.StringAttribute{
						Description:         "Number of hours to hold script logs for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "Number of hours to hold script logs for, everything older will be deleted.  More info: https://golang.org/pkg/time/#ParseDuration",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"object_store": schema.SingleNestedAttribute{
						Description:         "ObjectStore allows for backing up to a remote cloud storage.",
						MarkdownDescription: "ObjectStore allows for backing up to a remote cloud storage.",
						Attributes: map[string]schema.Attribute{
							"endpoint": schema.SingleNestedAttribute{
								Description:         "Endpoint contains the configuration for connecting to a custom Azure/S3/GCP compliant object store. If set will override 'CouchbaseCluster.spec.backup.objectEndpoint' See https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-cloud.html#compatible-object-stores",
								MarkdownDescription: "Endpoint contains the configuration for connecting to a custom Azure/S3/GCP compliant object store. If set will override 'CouchbaseCluster.spec.backup.objectEndpoint' See https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-cloud.html#compatible-object-stores",
								Attributes: map[string]schema.Attribute{
									"secret": schema.StringAttribute{
										Description:         "The name of the secret, in this namespace, that contains the CA certificate for verification of a TLS endpoint The secret must have the key with the name 'tls.crt'",
										MarkdownDescription: "The name of the secret, in this namespace, that contains the CA certificate for verification of a TLS endpoint The secret must have the key with the name 'tls.crt'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "The host/address of the custom object endpoint.",
										MarkdownDescription: "The host/address of the custom object endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"use_virtual_path": schema.BoolAttribute{
										Description:         "UseVirtualPath will force the AWS SDK to use the new virtual style paths which are often required by S3 compatible object stores.",
										MarkdownDescription: "UseVirtualPath will force the AWS SDK to use the new virtual style paths which are often required by S3 compatible object stores.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret": schema.StringAttribute{
								Description:         "ObjStoreSecret must contain two fields, access-key-id, secret-access-key and optionally either region or refresh-token. These correspond to the fields used by cbbackupmgr https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-backup.html#optional-2",
								MarkdownDescription: "ObjStoreSecret must contain two fields, access-key-id, secret-access-key and optionally either region or refresh-token. These correspond to the fields used by cbbackupmgr https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-backup.html#optional-2",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uri": schema.StringAttribute{
								Description:         "URI is a reference to a remote object store. This is the prefix of the object store and the bucket name. i.e s3://bucket, az://bucket or gs://bucket.",
								MarkdownDescription: "URI is a reference to a remote object store. This is the prefix of the object store and the bucket name. i.e s3://bucket, az://bucket or gs://bucket.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(az|s3|gs)://.{3,}$`), ""),
								},
							},

							"use_iam": schema.BoolAttribute{
								Description:         "Whether to allow the backup SDK to attempt to authenticate using the instance metadata api. If set, will override 'CouchbaseCluster.spec.backup.useIAM'.",
								MarkdownDescription: "Whether to allow the backup SDK to attempt to authenticate using the instance metadata api. If set, will override 'CouchbaseCluster.spec.backup.useIAM'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"s3bucket": schema.StringAttribute{
						Description:         "DEPRECATED - by spec.objectStore.uri Name of S3 bucket to backup to. If non-empty this overrides local backup.",
						MarkdownDescription: "DEPRECATED - by spec.objectStore.uri Name of S3 bucket to backup to. If non-empty this overrides local backup.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^s3://[a-z0-9-\.\/]{3,63}$`), ""),
						},
					},

					"services": schema.SingleNestedAttribute{
						Description:         "Services allows control over what services are included in the backup. By default, all service data and metadata are included.  Modifications to this field will only take effect on the next full backup.",
						MarkdownDescription: "Services allows control over what services are included in the backup. By default, all service data and metadata are included.  Modifications to this field will only take effect on the next full backup.",
						Attributes: map[string]schema.Attribute{
							"analytics": schema.BoolAttribute{
								Description:         "Analytics enables the backup of analytics data. This field defaults to 'true'.",
								MarkdownDescription: "Analytics enables the backup of analytics data. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bucket_config": schema.BoolAttribute{
								Description:         "BucketConfig enables the backup of bucket configuration. This field defaults to 'true'.",
								MarkdownDescription: "BucketConfig enables the backup of bucket configuration. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bucket_query": schema.BoolAttribute{
								Description:         "BucketQuery enables the backup of query metadata for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "BucketQuery enables the backup of query metadata for all buckets. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_analytics": schema.BoolAttribute{
								Description:         "ClusterAnalytics enables the backup of cluster-wide analytics data, for example synonyms. This field defaults to 'true'.",
								MarkdownDescription: "ClusterAnalytics enables the backup of cluster-wide analytics data, for example synonyms. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_query": schema.BoolAttribute{
								Description:         "ClusterQuery enables the backup of cluster level query metadata. This field defaults to 'true'.",
								MarkdownDescription: "ClusterQuery enables the backup of cluster level query metadata. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"data": schema.BoolAttribute{
								Description:         "Data enables the backup of key-value data/documents for all buckets. This can be further refined with the couchbasebackups.spec.data configuration. This field defaults to 'true'.",
								MarkdownDescription: "Data enables the backup of key-value data/documents for all buckets. This can be further refined with the couchbasebackups.spec.data configuration. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eventing": schema.BoolAttribute{
								Description:         "Eventing enables the backup of eventing service metadata. This field defaults to 'true'.",
								MarkdownDescription: "Eventing enables the backup of eventing service metadata. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fts_aliases": schema.BoolAttribute{
								Description:         "FTSAliases enables the backup of full-text search alias definitions. This field defaults to 'true'.",
								MarkdownDescription: "FTSAliases enables the backup of full-text search alias definitions. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fts_indexes": schema.BoolAttribute{
								Description:         "FTSIndexes enables the backup of full-text search index definitions for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "FTSIndexes enables the backup of full-text search index definitions for all buckets. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gs_indexes": schema.BoolAttribute{
								Description:         "GSIndexes enables the backup of global secondary index definitions for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "GSIndexes enables the backup of global secondary index definitions for all buckets. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"views": schema.BoolAttribute{
								Description:         "Views enables the backup of view definitions for all buckets. This field defaults to 'true'.",
								MarkdownDescription: "Views enables the backup of view definitions for all buckets. This field defaults to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"size": schema.StringAttribute{
						Description:         "Size allows the specification of a backup persistent volume, when using volume based backup. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						MarkdownDescription: "Size allows the specification of a backup persistent volume, when using volume based backup. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
						},
					},

					"storage_class_name": schema.StringAttribute{
						Description:         "Name of StorageClass to use.",
						MarkdownDescription: "Name of StorageClass to use.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"strategy": schema.StringAttribute{
						Description:         "Strategy defines how to perform backups.  'full_only' will only perform full backups, and you must define a schedule in the 'spec.full' field.  'full_incremental' will perform periodic full backups, and incremental backups in between.  You must define full and incremental schedules in the 'spec.full' and 'spec.incremental' fields respectively.  Care should be taken to ensure full and incremental schedules do not overlap, taking into account the backup time, as this will cause failures as the jobs attempt to mount the same backup volume. This field default to 'full_incremental'. Info: https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-strategies.html",
						MarkdownDescription: "Strategy defines how to perform backups.  'full_only' will only perform full backups, and you must define a schedule in the 'spec.full' field.  'full_incremental' will perform periodic full backups, and incremental backups in between.  You must define full and incremental schedules in the 'spec.full' and 'spec.incremental' fields respectively.  Care should be taken to ensure full and incremental schedules do not overlap, taking into account the backup time, as this will cause failures as the jobs attempt to mount the same backup volume. This field default to 'full_incremental'. Info: https://docs.couchbase.com/server/current/backup-restore/cbbackupmgr-strategies.html",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("full_incremental", "full_only"),
						},
					},

					"successful_jobs_history_limit": schema.Int64Attribute{
						Description:         "Amount of successful jobs to keep.",
						MarkdownDescription: "Amount of successful jobs to keep.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"threads": schema.Int64Attribute{
						Description:         "How many threads to use during the backup.  This field defaults to 1.",
						MarkdownDescription: "How many threads to use during the backup.  This field defaults to 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"ttl_seconds_after_finished": schema.Int64Attribute{
						Description:         "Amount of time to elapse before a completed job is deleted.",
						MarkdownDescription: "Amount of time to elapse before a completed job is deleted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
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

func (r *CouchbaseComCouchbaseBackupV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_backup_v2_manifest")

	var model CouchbaseComCouchbaseBackupV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseBackup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
