/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"regexp"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &CouchbaseComCouchbaseBackupV2Resource{}
	_ resource.ResourceWithConfigure   = &CouchbaseComCouchbaseBackupV2Resource{}
	_ resource.ResourceWithImportState = &CouchbaseComCouchbaseBackupV2Resource{}
)

func NewCouchbaseComCouchbaseBackupV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseBackupV2Resource{}
}

type CouchbaseComCouchbaseBackupV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CouchbaseComCouchbaseBackupV2ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

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
		EphemeralVolume        *bool  `tfsdk:"ephemeral_volume" json:"ephemeralVolume,omitempty"`
		FailedJobsHistoryLimit *int64 `tfsdk:"failed_jobs_history_limit" json:"failedJobsHistoryLimit,omitempty"`
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

func (r *CouchbaseComCouchbaseBackupV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_backup_v2"
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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

func (r *CouchbaseComCouchbaseBackupV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_backup_v2")

	var model CouchbaseComCouchbaseBackupV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseBackup")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackups"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CouchbaseComCouchbaseBackupV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_backup_v2")

	var data CouchbaseComCouchbaseBackupV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackups"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CouchbaseComCouchbaseBackupV2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_backup_v2")

	var model CouchbaseComCouchbaseBackupV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseBackup")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackups"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CouchbaseComCouchbaseBackupV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseBackupV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_backup_v2")

	var data CouchbaseComCouchbaseBackupV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackups"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackups"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *CouchbaseComCouchbaseBackupV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
