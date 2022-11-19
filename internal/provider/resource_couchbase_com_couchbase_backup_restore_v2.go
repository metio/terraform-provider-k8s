/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type CouchbaseComCouchbaseBackupRestoreV2Resource struct{}

var (
	_ resource.Resource = (*CouchbaseComCouchbaseBackupRestoreV2Resource)(nil)
)

type CouchbaseComCouchbaseBackupRestoreV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CouchbaseComCouchbaseBackupRestoreV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		BackoffLimit *int64 `tfsdk:"backoff_limit" yaml:"backoffLimit,omitempty"`

		Backup *string `tfsdk:"backup" yaml:"backup,omitempty"`

		Buckets utilities.Dynamic `tfsdk:"buckets" yaml:"buckets,omitempty"`

		Data *struct {
			Exclude *[]string `tfsdk:"exclude" yaml:"exclude,omitempty"`

			FilterKeys *string `tfsdk:"filter_keys" yaml:"filterKeys,omitempty"`

			FilterValues *string `tfsdk:"filter_values" yaml:"filterValues,omitempty"`

			Include *[]string `tfsdk:"include" yaml:"include,omitempty"`

			Map *[]struct {
				Source *string `tfsdk:"source" yaml:"source,omitempty"`

				Target *string `tfsdk:"target" yaml:"target,omitempty"`
			} `tfsdk:"map" yaml:"map,omitempty"`
		} `tfsdk:"data" yaml:"data,omitempty"`

		End *struct {
			Int *int64 `tfsdk:"int" yaml:"int,omitempty"`

			Str *string `tfsdk:"str" yaml:"str,omitempty"`
		} `tfsdk:"end" yaml:"end,omitempty"`

		ForceUpdates *bool `tfsdk:"force_updates" yaml:"forceUpdates,omitempty"`

		LogRetention *string `tfsdk:"log_retention" yaml:"logRetention,omitempty"`

		Repo *string `tfsdk:"repo" yaml:"repo,omitempty"`

		S3bucket *string `tfsdk:"s3bucket" yaml:"s3bucket,omitempty"`

		Services *struct {
			Analytics *bool `tfsdk:"analytics" yaml:"analytics,omitempty"`

			BucketConfig *bool `tfsdk:"bucket_config" yaml:"bucketConfig,omitempty"`

			BucketQuery *bool `tfsdk:"bucket_query" yaml:"bucketQuery,omitempty"`

			ClusterAnalytics *bool `tfsdk:"cluster_analytics" yaml:"clusterAnalytics,omitempty"`

			ClusterQuery *bool `tfsdk:"cluster_query" yaml:"clusterQuery,omitempty"`

			Data *bool `tfsdk:"data" yaml:"data,omitempty"`

			Eventing *bool `tfsdk:"eventing" yaml:"eventing,omitempty"`

			FtAlias *bool `tfsdk:"ft_alias" yaml:"ftAlias,omitempty"`

			FtIndex *bool `tfsdk:"ft_index" yaml:"ftIndex,omitempty"`

			GsiIndex *bool `tfsdk:"gsi_index" yaml:"gsiIndex,omitempty"`

			Views *bool `tfsdk:"views" yaml:"views,omitempty"`
		} `tfsdk:"services" yaml:"services,omitempty"`

		Start *struct {
			Int *int64 `tfsdk:"int" yaml:"int,omitempty"`

			Str *string `tfsdk:"str" yaml:"str,omitempty"`
		} `tfsdk:"start" yaml:"start,omitempty"`

		Threads *int64 `tfsdk:"threads" yaml:"threads,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCouchbaseComCouchbaseBackupRestoreV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseBackupRestoreV2Resource{}
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_couchbase_com_couchbase_backup_restore_v2"
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CouchbaseBackupRestore allows the restoration of all Couchbase cluster data from a CouchbaseBackup resource.",
		MarkdownDescription: "CouchbaseBackupRestore allows the restoration of all Couchbase cluster data from a CouchbaseBackup resource.",
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
				Description:         "CouchbaseBackupRestoreSpec allows the specification of data restoration to be configured.  This includes the backup and repository to restore data from, and the time range of data to be restored.",
				MarkdownDescription: "CouchbaseBackupRestoreSpec allows the specification of data restoration to be configured.  This includes the backup and repository to restore data from, and the time range of data to be restored.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"backoff_limit": {
						Description:         "Number of times the restore job should try to execute.",
						MarkdownDescription: "Number of times the restore job should try to execute.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup": {
						Description:         "The backup resource name associated with this restore, or the backup PVC name to restore from.",
						MarkdownDescription: "The backup resource name associated with this restore, or the backup PVC name to restore from.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"buckets": {
						Description:         "DEPRECATED - by spec.data. Specific buckets can be explicitly included or excluded in the restore, as well as bucket mappings.  This field is now ignored.",
						MarkdownDescription: "DEPRECATED - by spec.data. Specific buckets can be explicitly included or excluded in the restore, as well as bucket mappings.  This field is now ignored.",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"data": {
						Description:         "Data allows control over what key-value/document data is included in the restore.  By default, all data is included.",
						MarkdownDescription: "Data allows control over what key-value/document data is included in the restore.  By default, all data is included.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"exclude": {
								Description:         "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",
								MarkdownDescription: "Exclude defines the buckets, scopes or collections that are excluded from the backup. When this field is set, it implies that by default everything will be backed up, and data items can be explicitly excluded.  You may define an exclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Excluded data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as included items.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"filter_keys": {
								Description:         "FilterKeys only restores documents whose names match the provided regular expression.",
								MarkdownDescription: "FilterKeys only restores documents whose names match the provided regular expression.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"filter_values": {
								Description:         "FilterValues only restores documents whose values match the provided regular expression.",
								MarkdownDescription: "FilterValues only restores documents whose values match the provided regular expression.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"include": {
								Description:         "Include defines the buckets, scopes or collections that are included in the restore. When this field is set, it implies that by default nothing will be restored, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",
								MarkdownDescription: "Include defines the buckets, scopes or collections that are included in the restore. When this field is set, it implies that by default nothing will be restored, and data items must be explicitly included.  You may define an inclusion as a bucket -- 'my-bucket', a scope -- 'my-bucket.my-scope', or a collection -- 'my-bucket.my-scope.my-collection'. Buckets may contain periods, and therefore must be escaped -- 'my.bucket.my-scope', as period is the separator used to delimit scopes and collections.  Included data cannot overlap e.g. specifying 'my-bucket' and 'my-bucket.my-scope' is illegal.  This field cannot be used at the same time as excluded items.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"map": {
								Description:         "Map allows data items in the restore to be remapped to a different named container. Buckets can be remapped to other buckets e.g. 'source=target', scopes and collections can be remapped to other scopes and collections within the same bucket only e.g. 'bucket.scope=bucket.other' or 'bucket.scope.collection=bucket.scope.other'.  Map sources may only be specified once, and may not overlap.",
								MarkdownDescription: "Map allows data items in the restore to be remapped to a different named container. Buckets can be remapped to other buckets e.g. 'source=target', scopes and collections can be remapped to other scopes and collections within the same bucket only e.g. 'bucket.scope=bucket.other' or 'bucket.scope.collection=bucket.scope.other'.  Map sources may only be specified once, and may not overlap.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"source": {
										Description:         "Source defines the data source of the mapping, this may be either a bucket, scope or collection.",
										MarkdownDescription: "Source defines the data source of the mapping, this may be either a bucket, scope or collection.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:[a-zA-Z0-9\-_%]|\\.){1,100}(\._default(\._default)?|\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29}(\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29})?)?$`), ""),
										},
									},

									"target": {
										Description:         "Target defines the data target of the mapping, this may be either a bucket, scope or collection, and must refer to the same type as the restore source.",
										MarkdownDescription: "Target defines the data target of the mapping, this may be either a bucket, scope or collection, and must refer to the same type as the restore source.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:[a-zA-Z0-9\-_%]|\\.){1,100}(\._default(\._default)?|\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29}(\.[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,29})?)?$`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"end": {
						Description:         "End denotes the last backup to restore from.  Omitting this field will only restore the backup referenced by start.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",
						MarkdownDescription: "End denotes the last backup to restore from.  Omitting this field will only restore the backup referenced by start.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"int": {
								Description:         "Int references a relative backup by index.",
								MarkdownDescription: "Int references a relative backup by index.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"str": {
								Description:         "Str references an absolute backup by name.",
								MarkdownDescription: "Str references an absolute backup by name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"force_updates": {
						Description:         "Forces data in the Couchbase cluster to be overwritten even if the data in the cluster is newer than the restore",
						MarkdownDescription: "Forces data in the Couchbase cluster to be overwritten even if the data in the cluster is newer than the restore",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_retention": {
						Description:         "Number of hours to hold restore script logs for, everything older will be deleted. More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "Number of hours to hold restore script logs for, everything older will be deleted. More info: https://golang.org/pkg/time/#ParseDuration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"repo": {
						Description:         "Repo is the backup folder to restore from.  If no repository is specified, the backup container will choose the latest.",
						MarkdownDescription: "Repo is the backup folder to restore from.  If no repository is specified, the backup container will choose the latest.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"s3bucket": {
						Description:         "Name of S3 bucket to restore from. If non-empty this overrides local backup.",
						MarkdownDescription: "Name of S3 bucket to restore from. If non-empty this overrides local backup.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^s3://[a-z0-9-\.]{3,63}$`), ""),
						},
					},

					"services": {
						Description:         "This list accepts a certain set of parameters that will disable that data and prevent it being restored.",
						MarkdownDescription: "This list accepts a certain set of parameters that will disable that data and prevent it being restored.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"analytics": {
								Description:         "Analytics restores analytics datasets from the backup.  This field defaults to true.",
								MarkdownDescription: "Analytics restores analytics datasets from the backup.  This field defaults to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bucket_config": {
								Description:         "BucketConfig restores all bucket configuration settings. If you are restoring to cluster with managed buckets, then this option may conflict with existing bucket settings, and the results are undefined, so avoid use.  This option is intended for use with unmanaged buckets.  Note that bucket durability settings are not restored in versions less than and equal to 1.1.0, and will need to be manually applied.  This field defaults to false.",
								MarkdownDescription: "BucketConfig restores all bucket configuration settings. If you are restoring to cluster with managed buckets, then this option may conflict with existing bucket settings, and the results are undefined, so avoid use.  This option is intended for use with unmanaged buckets.  Note that bucket durability settings are not restored in versions less than and equal to 1.1.0, and will need to be manually applied.  This field defaults to false.",

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
								Description:         "Data restores document data from the backup.  This field defaults to true.",
								MarkdownDescription: "Data restores document data from the backup.  This field defaults to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"eventing": {
								Description:         "Eventing restores eventing functions from the backup.  This field defaults to true.",
								MarkdownDescription: "Eventing restores eventing functions from the backup.  This field defaults to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ft_alias": {
								Description:         "FTAlias restores full-text search aliases from the backup.  This field defaults to true.",
								MarkdownDescription: "FTAlias restores full-text search aliases from the backup.  This field defaults to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ft_index": {
								Description:         "FTIndex restores full-text search indexes from the backup.  This field defaults to true.",
								MarkdownDescription: "FTIndex restores full-text search indexes from the backup.  This field defaults to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gsi_index": {
								Description:         "GSIIndex restores document indexes from the backup.  This field defaults to true.",
								MarkdownDescription: "GSIIndex restores document indexes from the backup.  This field defaults to true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"views": {
								Description:         "Views restores views from the backup.  This field defaults to true.",
								MarkdownDescription: "Views restores views from the backup.  This field defaults to true.",

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

					"start": {
						Description:         "Start denotes the first backup to restore from.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",
						MarkdownDescription: "Start denotes the first backup to restore from.  This may be specified as an integer index (starting from 1), a string specifying a short date DD-MM-YYYY, the backup name, or one of either 'start' or 'oldest' keywords.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"int": {
								Description:         "Int references a relative backup by index.",
								MarkdownDescription: "Int references a relative backup by index.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"str": {
								Description:         "Str references an absolute backup by name.",
								MarkdownDescription: "Str references an absolute backup by name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"threads": {
						Description:         "How many threads to use during the restore.",
						MarkdownDescription: "How many threads to use during the restore.",

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

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_backup_restore_v2")

	var state CouchbaseComCouchbaseBackupRestoreV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseBackupRestoreV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseBackupRestore")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_backup_restore_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_backup_restore_v2")

	var state CouchbaseComCouchbaseBackupRestoreV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseBackupRestoreV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseBackupRestore")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CouchbaseComCouchbaseBackupRestoreV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_backup_restore_v2")
	// NO-OP: Terraform removes the state automatically for us
}
