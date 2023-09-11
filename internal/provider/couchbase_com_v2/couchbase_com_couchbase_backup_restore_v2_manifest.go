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
	_ datasource.DataSource = &CouchbaseComCouchbaseBackupRestoreV2Manifest{}
)

func NewCouchbaseComCouchbaseBackupRestoreV2Manifest() datasource.DataSource {
	return &CouchbaseComCouchbaseBackupRestoreV2Manifest{}
}

type CouchbaseComCouchbaseBackupRestoreV2Manifest struct{}

type CouchbaseComCouchbaseBackupRestoreV2ManifestData struct {
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
		BackoffLimit *int64             `tfsdk:"backoff_limit" json:"backoffLimit,omitempty"`
		Backup       *string            `tfsdk:"backup" json:"backup,omitempty"`
		Buckets      *map[string]string `tfsdk:"buckets" json:"buckets,omitempty"`
		Data         *struct {
			Exclude      *[]string `tfsdk:"exclude" json:"exclude,omitempty"`
			FilterKeys   *string   `tfsdk:"filter_keys" json:"filterKeys,omitempty"`
			FilterValues *string   `tfsdk:"filter_values" json:"filterValues,omitempty"`
			Include      *[]string `tfsdk:"include" json:"include,omitempty"`
			Map          *[]struct {
				Source *string `tfsdk:"source" json:"source,omitempty"`
				Target *string `tfsdk:"target" json:"target,omitempty"`
			} `tfsdk:"map" json:"map,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		End *struct {
			Int *int64  `tfsdk:"int" json:"int,omitempty"`
			Str *string `tfsdk:"str" json:"str,omitempty"`
		} `tfsdk:"end" json:"end,omitempty"`
		ForceUpdates *bool   `tfsdk:"force_updates" json:"forceUpdates,omitempty"`
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
		Repo     *string `tfsdk:"repo" json:"repo,omitempty"`
		S3bucket *string `tfsdk:"s3bucket" json:"s3bucket,omitempty"`
		Services *struct {
			Analytics        *bool `tfsdk:"analytics" json:"analytics,omitempty"`
			BucketConfig     *bool `tfsdk:"bucket_config" json:"bucketConfig,omitempty"`
			BucketQuery      *bool `tfsdk:"bucket_query" json:"bucketQuery,omitempty"`
			ClusterAnalytics *bool `tfsdk:"cluster_analytics" json:"clusterAnalytics,omitempty"`
			ClusterQuery     *bool `tfsdk:"cluster_query" json:"clusterQuery,omitempty"`
			Data             *bool `tfsdk:"data" json:"data,omitempty"`
			Eventing         *bool `tfsdk:"eventing" json:"eventing,omitempty"`
			FtAlias          *bool `tfsdk:"ft_alias" json:"ftAlias,omitempty"`
			FtIndex          *bool `tfsdk:"ft_index" json:"ftIndex,omitempty"`
			GsiIndex         *bool `tfsdk:"gsi_index" json:"gsiIndex,omitempty"`
			Views            *bool `tfsdk:"views" json:"views,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		StagingVolume *struct {
			Size             *string `tfsdk:"size" json:"size,omitempty"`
			StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
		} `tfsdk:"staging_volume" json:"stagingVolume,omitempty"`
		Start *struct {
			Int *int64  `tfsdk:"int" json:"int,omitempty"`
			Str *string `tfsdk:"str" json:"str,omitempty"`
		} `tfsdk:"start" json:"start,omitempty"`
		Threads                 *int64 `tfsdk:"threads" json:"threads,omitempty"`
		TtlSecondsAfterFinished *int64 `tfsdk:"ttl_seconds_after_finished" json:"ttlSecondsAfterFinished,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_backup_restore_v2_manifest"
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CouchbaseBackupRestore allows the restoration of all Couchbase cluster data from a CouchbaseBackup resource.",
		MarkdownDescription: "CouchbaseBackupRestore allows the restoration of all Couchbase cluster data from a CouchbaseBackup resource.",
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
				Description:         "CouchbaseBackupRestoreSpec allows the specification of data restoration to be configured.  This includes the backup and repository to restore data from, and the time range of data to be restored.",
				MarkdownDescription: "CouchbaseBackupRestoreSpec allows the specification of data restoration to be configured.  This includes the backup and repository to restore data from, and the time range of data to be restored.",
				Attributes: map[string]schema.Attribute{
					"backoff_limit": schema.Int64Attribute{
						Description:         "Number of times the restore job should try to execute.",
						MarkdownDescription: "Number of times the restore job should try to execute.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"backup": schema.StringAttribute{
						Description:         "The backup resource name associated with this restore, or the backup PVC name to restore from.",
						MarkdownDescription: "The backup resource name associated with this restore, or the backup PVC name to restore from.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"buckets": schema.MapAttribute{
						Description:         "DEPRECATED - by spec.data. Specific buckets can be explicitly included or excluded in the restore, as well as bucket mappings.  This field is now ignored.",
						MarkdownDescription: "DEPRECATED - by spec.data. Specific buckets can be explicitly included or excluded in the restore, as well as bucket mappings.  This field is now ignored.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"data": schema.SingleNestedAttribute{
						Description:         "Data allows control over what key-value/document data is included in the restore.  By default, all data is included.",
						MarkdownDescription: "Data allows control over what key-value/document data is included in the restore.  By default, all data is included.",
						Attributes: map[string]schema.Attribute{
							"exclude": schema.ListAttribute{
								Description:         "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",
								MarkdownDescription: "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"filter_keys": schema.StringAttribute{
								Description:         "FilterKeys only restores documents whose names match the provided regular expression.",
								MarkdownDescription: "FilterKeys only restores documents whose names match the provided regular expression.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"filter_values": schema.StringAttribute{
								Description:         "FilterValues only restores documents whose values match the provided regular expression.",
								MarkdownDescription: "FilterValues only restores documents whose values match the provided regular expression.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"include": schema.ListAttribute{
								Description:         "Include defines the buckets, scopes or collections that are included in the restore. When this field is set, it implies that by default nothing will be restored, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",
								MarkdownDescription: "Include defines the buckets, scopes or collections that are included in the restore. When this field is set, it implies that by default nothing will be restored, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"map": schema.ListNestedAttribute{
								Description:         "Map allows data items in the restore to be remapped to a different named container. Buckets can be remapped to other buckets e.g. 'source=target', scopes and collections can be remapped to other scopes and collections within the same bucket only e.g. 'bucket.scope=bucket.other' or 'bucket.scope.collection=bucket.scope.other'.  Map sources may only be specified once, and may not overlap.",
								MarkdownDescription: "Map allows data items in the restore to be remapped to a different named container. Buckets can be remapped to other buckets e.g. 'source=target', scopes and collections can be remapped to other scopes and collections within the same bucket only e.g. 'bucket.scope=bucket.other' or 'bucket.scope.collection=bucket.scope.other'.  Map sources may only be specified once, and may not overlap.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"source": schema.StringAttribute{
											Description:         "Source defines the data source of the mapping, this may be either a bucket, scope or collection.",
											MarkdownDescription: "Source defines the data source of the mapping, this may be either a bucket, scope or collection.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(?:[a-zA-Z0-9\-_%]|\\.){1,100}(\._default(\._default)?|\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29}(\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29})?)?$`), ""),
											},
										},

										"target": schema.StringAttribute{
											Description:         "Target defines the data target of the mapping, this may be either a bucket, scope or collection, and must refer to the same type as the restore source.",
											MarkdownDescription: "Target defines the data target of the mapping, this may be either a bucket, scope or collection, and must refer to the same type as the restore source.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(?:[a-zA-Z0-9\-_%]|\\.){1,100}(\._default(\._default)?|\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29}(\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29})?)?$`), ""),
											},
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

					"end": schema.SingleNestedAttribute{
						Description:         "End denotes the last backup to restore from.  Omitting this field will only restore the backup referenced by start.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",
						MarkdownDescription: "End denotes the last backup to restore from.  Omitting this field will only restore the backup referenced by start.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",
						Attributes: map[string]schema.Attribute{
							"int": schema.Int64Attribute{
								Description:         "Int references a relative backup by index.",
								MarkdownDescription: "Int references a relative backup by index.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"str": schema.StringAttribute{
								Description:         "Str references an absolute backup by name.",
								MarkdownDescription: "Str references an absolute backup by name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"force_updates": schema.BoolAttribute{
						Description:         "Forces data in the Couchbase cluster to be overwritten even if the data in the cluster is newer than the restore",
						MarkdownDescription: "Forces data in the Couchbase cluster to be overwritten even if the data in the cluster is newer than the restore",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_retention": schema.StringAttribute{
						Description:         "Number of hours to hold restore script logs for, everything older will be deleted. More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "Number of hours to hold restore script logs for, everything older will be deleted. More info: https://golang.org/pkg/time/#ParseDuration",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"object_store": schema.SingleNestedAttribute{
						Description:         "The remote destination for backup.",
						MarkdownDescription: "The remote destination for backup.",
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

					"repo": schema.StringAttribute{
						Description:         "Repo is the backup folder to restore from.  If no repository is specified, the backup container will choose the latest.",
						MarkdownDescription: "Repo is the backup folder to restore from.  If no repository is specified, the backup container will choose the latest.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"s3bucket": schema.StringAttribute{
						Description:         "DEPRECATED - by spec.objectStore.uri Name of S3 bucket to restore from. If non-empty this overrides local backup.",
						MarkdownDescription: "DEPRECATED - by spec.objectStore.uri Name of S3 bucket to restore from. If non-empty this overrides local backup.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^s3://[a-z0-9-\.\/]{3,63}$`), ""),
						},
					},

					"services": schema.SingleNestedAttribute{
						Description:         "This list accepts a certain set of parameters that will disable that data and prevent it being restored.",
						MarkdownDescription: "This list accepts a certain set of parameters that will disable that data and prevent it being restored.",
						Attributes: map[string]schema.Attribute{
							"analytics": schema.BoolAttribute{
								Description:         "Analytics restores analytics datasets from the backup.  This field defaults to true.",
								MarkdownDescription: "Analytics restores analytics datasets from the backup.  This field defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"bucket_config": schema.BoolAttribute{
								Description:         "BucketConfig restores all bucket configuration settings. If you are restoring to cluster with managed buckets, then this option may conflict with existing bucket settings, and the results are undefined, so avoid use.  This option is intended for use with unmanaged buckets.  Note that bucket durability settings are not restored in versions less than and equal to 1.1.0, and will need to be manually applied.  This field defaults to false.",
								MarkdownDescription: "BucketConfig restores all bucket configuration settings. If you are restoring to cluster with managed buckets, then this option may conflict with existing bucket settings, and the results are undefined, so avoid use.  This option is intended for use with unmanaged buckets.  Note that bucket durability settings are not restored in versions less than and equal to 1.1.0, and will need to be manually applied.  This field defaults to false.",
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
								Description:         "Data restores document data from the backup.  This field defaults to true.",
								MarkdownDescription: "Data restores document data from the backup.  This field defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"eventing": schema.BoolAttribute{
								Description:         "Eventing restores eventing functions from the backup.  This field defaults to true.",
								MarkdownDescription: "Eventing restores eventing functions from the backup.  This field defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ft_alias": schema.BoolAttribute{
								Description:         "FTAlias restores full-text search aliases from the backup.  This field defaults to true.",
								MarkdownDescription: "FTAlias restores full-text search aliases from the backup.  This field defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ft_index": schema.BoolAttribute{
								Description:         "FTIndex restores full-text search indexes from the backup.  This field defaults to true.",
								MarkdownDescription: "FTIndex restores full-text search indexes from the backup.  This field defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gsi_index": schema.BoolAttribute{
								Description:         "GSIIndex restores document indexes from the backup.  This field defaults to true.",
								MarkdownDescription: "GSIIndex restores document indexes from the backup.  This field defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"views": schema.BoolAttribute{
								Description:         "Views restores views from the backup.  This field defaults to true.",
								MarkdownDescription: "Views restores views from the backup.  This field defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"staging_volume": schema.SingleNestedAttribute{
						Description:         "StagingVolume contains configuration related to the ephemeral volume used as staging when restoring from a cloud backup.",
						MarkdownDescription: "StagingVolume contains configuration related to the ephemeral volume used as staging when restoring from a cloud backup.",
						Attributes: map[string]schema.Attribute{
							"size": schema.StringAttribute{
								Description:         "Size allows the specification of a staging volume. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes The ephemeral volume will only be used when restoring from a cloud provider, if the backup job was created using ephemeral storage. Otherwise the restore job will share a staging volume with the backup job.",
								MarkdownDescription: "Size allows the specification of a staging volume. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes The ephemeral volume will only be used when restoring from a cloud provider, if the backup job was created using ephemeral storage. Otherwise the restore job will share a staging volume with the backup job.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"start": schema.SingleNestedAttribute{
						Description:         "Start denotes the first backup to restore from.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",
						MarkdownDescription: "Start denotes the first backup to restore from.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",
						Attributes: map[string]schema.Attribute{
							"int": schema.Int64Attribute{
								Description:         "Int references a relative backup by index.",
								MarkdownDescription: "Int references a relative backup by index.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"str": schema.StringAttribute{
								Description:         "Str references an absolute backup by name.",
								MarkdownDescription: "Str references an absolute backup by name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"threads": schema.Int64Attribute{
						Description:         "How many threads to use during the restore.",
						MarkdownDescription: "How many threads to use during the restore.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"ttl_seconds_after_finished": schema.Int64Attribute{
						Description:         "Number of seconds to elapse before a completed job is deleted.",
						MarkdownDescription: "Number of seconds to elapse before a completed job is deleted.",
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

func (r *CouchbaseComCouchbaseBackupRestoreV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_backup_restore_v2_manifest")

	var model CouchbaseComCouchbaseBackupRestoreV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseBackupRestore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
