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
	_ resource.Resource                = &CouchbaseComCouchbaseBackupRestoreV2Resource{}
	_ resource.ResourceWithConfigure   = &CouchbaseComCouchbaseBackupRestoreV2Resource{}
	_ resource.ResourceWithImportState = &CouchbaseComCouchbaseBackupRestoreV2Resource{}
)

func NewCouchbaseComCouchbaseBackupRestoreV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseBackupRestoreV2Resource{}
}

type CouchbaseComCouchbaseBackupRestoreV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CouchbaseComCouchbaseBackupRestoreV2ResourceData struct {
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

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_backup_restore_v2"
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_backup_restore_v2")

	var model CouchbaseComCouchbaseBackupRestoreV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseBackupRestore")

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
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackuprestores"}).
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

	var readResponse CouchbaseComCouchbaseBackupRestoreV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_backup_restore_v2")

	var data CouchbaseComCouchbaseBackupRestoreV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackuprestores"}).
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

	var readResponse CouchbaseComCouchbaseBackupRestoreV2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_backup_restore_v2")

	var model CouchbaseComCouchbaseBackupRestoreV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseBackupRestore")

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
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackuprestores"}).
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

	var readResponse CouchbaseComCouchbaseBackupRestoreV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_backup_restore_v2")

	var data CouchbaseComCouchbaseBackupRestoreV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackuprestores"}).
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
				Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "couchbasebackuprestores"}).
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

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
