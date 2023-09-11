/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package s3_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &S3ServicesK8SAwsBucketV1Alpha1Manifest{}
)

func NewS3ServicesK8SAwsBucketV1Alpha1Manifest() datasource.DataSource {
	return &S3ServicesK8SAwsBucketV1Alpha1Manifest{}
}

type S3ServicesK8SAwsBucketV1Alpha1Manifest struct{}

type S3ServicesK8SAwsBucketV1Alpha1ManifestData struct {
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
		Accelerate *struct {
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"accelerate" json:"accelerate,omitempty"`
		Acl       *string `tfsdk:"acl" json:"acl,omitempty"`
		Analytics *[]struct {
			Filter *struct {
				And *struct {
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Tags   *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"and" json:"and,omitempty"`
				Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Tag    *struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"filter" json:"filter,omitempty"`
			Id                   *string `tfsdk:"id" json:"id,omitempty"`
			StorageClassAnalysis *struct {
				DataExport *struct {
					Destination *struct {
						S3BucketDestination *struct {
							Bucket          *string `tfsdk:"bucket" json:"bucket,omitempty"`
							BucketAccountID *string `tfsdk:"bucket_account_id" json:"bucketAccountID,omitempty"`
							Format          *string `tfsdk:"format" json:"format,omitempty"`
							Prefix          *string `tfsdk:"prefix" json:"prefix,omitempty"`
						} `tfsdk:"s3_bucket_destination" json:"s3BucketDestination,omitempty"`
					} `tfsdk:"destination" json:"destination,omitempty"`
					OutputSchemaVersion *string `tfsdk:"output_schema_version" json:"outputSchemaVersion,omitempty"`
				} `tfsdk:"data_export" json:"dataExport,omitempty"`
			} `tfsdk:"storage_class_analysis" json:"storageClassAnalysis,omitempty"`
		} `tfsdk:"analytics" json:"analytics,omitempty"`
		Cors *struct {
			CorsRules *[]struct {
				AllowedHeaders *[]string `tfsdk:"allowed_headers" json:"allowedHeaders,omitempty"`
				AllowedMethods *[]string `tfsdk:"allowed_methods" json:"allowedMethods,omitempty"`
				AllowedOrigins *[]string `tfsdk:"allowed_origins" json:"allowedOrigins,omitempty"`
				ExposeHeaders  *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
				Id             *string   `tfsdk:"id" json:"id,omitempty"`
				MaxAgeSeconds  *int64    `tfsdk:"max_age_seconds" json:"maxAgeSeconds,omitempty"`
			} `tfsdk:"cors_rules" json:"corsRules,omitempty"`
		} `tfsdk:"cors" json:"cors,omitempty"`
		CreateBucketConfiguration *struct {
			LocationConstraint *string `tfsdk:"location_constraint" json:"locationConstraint,omitempty"`
		} `tfsdk:"create_bucket_configuration" json:"createBucketConfiguration,omitempty"`
		Encryption *struct {
			Rules *[]struct {
				ApplyServerSideEncryptionByDefault *struct {
					KmsMasterKeyID *string `tfsdk:"kms_master_key_id" json:"kmsMasterKeyID,omitempty"`
					SseAlgorithm   *string `tfsdk:"sse_algorithm" json:"sseAlgorithm,omitempty"`
				} `tfsdk:"apply_server_side_encryption_by_default" json:"applyServerSideEncryptionByDefault,omitempty"`
				BucketKeyEnabled *bool `tfsdk:"bucket_key_enabled" json:"bucketKeyEnabled,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"encryption" json:"encryption,omitempty"`
		GrantFullControl   *string `tfsdk:"grant_full_control" json:"grantFullControl,omitempty"`
		GrantRead          *string `tfsdk:"grant_read" json:"grantRead,omitempty"`
		GrantReadACP       *string `tfsdk:"grant_read_acp" json:"grantReadACP,omitempty"`
		GrantWrite         *string `tfsdk:"grant_write" json:"grantWrite,omitempty"`
		GrantWriteACP      *string `tfsdk:"grant_write_acp" json:"grantWriteACP,omitempty"`
		IntelligentTiering *[]struct {
			Filter *struct {
				And *struct {
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Tags   *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"and" json:"and,omitempty"`
				Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Tag    *struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"filter" json:"filter,omitempty"`
			Id       *string `tfsdk:"id" json:"id,omitempty"`
			Status   *string `tfsdk:"status" json:"status,omitempty"`
			Tierings *[]struct {
				AccessTier *string `tfsdk:"access_tier" json:"accessTier,omitempty"`
				Days       *int64  `tfsdk:"days" json:"days,omitempty"`
			} `tfsdk:"tierings" json:"tierings,omitempty"`
		} `tfsdk:"intelligent_tiering" json:"intelligentTiering,omitempty"`
		Inventory *[]struct {
			Destination *struct {
				S3BucketDestination *struct {
					AccountID  *string `tfsdk:"account_id" json:"accountID,omitempty"`
					Bucket     *string `tfsdk:"bucket" json:"bucket,omitempty"`
					Encryption *struct {
						SseKMS *struct {
							KeyID *string `tfsdk:"key_id" json:"keyID,omitempty"`
						} `tfsdk:"sse_kms" json:"sseKMS,omitempty"`
					} `tfsdk:"encryption" json:"encryption,omitempty"`
					Format *string `tfsdk:"format" json:"format,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
				} `tfsdk:"s3_bucket_destination" json:"s3BucketDestination,omitempty"`
			} `tfsdk:"destination" json:"destination,omitempty"`
			Filter *struct {
				Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
			} `tfsdk:"filter" json:"filter,omitempty"`
			Id                     *string   `tfsdk:"id" json:"id,omitempty"`
			IncludedObjectVersions *string   `tfsdk:"included_object_versions" json:"includedObjectVersions,omitempty"`
			IsEnabled              *bool     `tfsdk:"is_enabled" json:"isEnabled,omitempty"`
			OptionalFields         *[]string `tfsdk:"optional_fields" json:"optionalFields,omitempty"`
			Schedule               *struct {
				Frequency *string `tfsdk:"frequency" json:"frequency,omitempty"`
			} `tfsdk:"schedule" json:"schedule,omitempty"`
		} `tfsdk:"inventory" json:"inventory,omitempty"`
		Lifecycle *struct {
			Rules *[]struct {
				AbortIncompleteMultipartUpload *struct {
					DaysAfterInitiation *int64 `tfsdk:"days_after_initiation" json:"daysAfterInitiation,omitempty"`
				} `tfsdk:"abort_incomplete_multipart_upload" json:"abortIncompleteMultipartUpload,omitempty"`
				Expiration *struct {
					Date                      *string `tfsdk:"date" json:"date,omitempty"`
					Days                      *int64  `tfsdk:"days" json:"days,omitempty"`
					ExpiredObjectDeleteMarker *bool   `tfsdk:"expired_object_delete_marker" json:"expiredObjectDeleteMarker,omitempty"`
				} `tfsdk:"expiration" json:"expiration,omitempty"`
				Filter *struct {
					And *struct {
						ObjectSizeGreaterThan *int64  `tfsdk:"object_size_greater_than" json:"objectSizeGreaterThan,omitempty"`
						ObjectSizeLessThan    *int64  `tfsdk:"object_size_less_than" json:"objectSizeLessThan,omitempty"`
						Prefix                *string `tfsdk:"prefix" json:"prefix,omitempty"`
						Tags                  *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tags" json:"tags,omitempty"`
					} `tfsdk:"and" json:"and,omitempty"`
					ObjectSizeGreaterThan *int64  `tfsdk:"object_size_greater_than" json:"objectSizeGreaterThan,omitempty"`
					ObjectSizeLessThan    *int64  `tfsdk:"object_size_less_than" json:"objectSizeLessThan,omitempty"`
					Prefix                *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Tag                   *struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"filter" json:"filter,omitempty"`
				Id                          *string `tfsdk:"id" json:"id,omitempty"`
				NoncurrentVersionExpiration *struct {
					NewerNoncurrentVersions *int64 `tfsdk:"newer_noncurrent_versions" json:"newerNoncurrentVersions,omitempty"`
					NoncurrentDays          *int64 `tfsdk:"noncurrent_days" json:"noncurrentDays,omitempty"`
				} `tfsdk:"noncurrent_version_expiration" json:"noncurrentVersionExpiration,omitempty"`
				NoncurrentVersionTransitions *[]struct {
					NewerNoncurrentVersions *int64  `tfsdk:"newer_noncurrent_versions" json:"newerNoncurrentVersions,omitempty"`
					NoncurrentDays          *int64  `tfsdk:"noncurrent_days" json:"noncurrentDays,omitempty"`
					StorageClass            *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				} `tfsdk:"noncurrent_version_transitions" json:"noncurrentVersionTransitions,omitempty"`
				Prefix      *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Status      *string `tfsdk:"status" json:"status,omitempty"`
				Transitions *[]struct {
					Date         *string `tfsdk:"date" json:"date,omitempty"`
					Days         *int64  `tfsdk:"days" json:"days,omitempty"`
					StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				} `tfsdk:"transitions" json:"transitions,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"lifecycle" json:"lifecycle,omitempty"`
		Logging *struct {
			LoggingEnabled *struct {
				TargetBucket *string `tfsdk:"target_bucket" json:"targetBucket,omitempty"`
				TargetGrants *[]struct {
					Grantee *struct {
						DisplayName  *string `tfsdk:"display_name" json:"displayName,omitempty"`
						EmailAddress *string `tfsdk:"email_address" json:"emailAddress,omitempty"`
						Id           *string `tfsdk:"id" json:"id,omitempty"`
						Type_        *string `tfsdk:"type_" json:"type_,omitempty"`
						URI          *string `tfsdk:"u_ri" json:"uRI,omitempty"`
					} `tfsdk:"grantee" json:"grantee,omitempty"`
					Permission *string `tfsdk:"permission" json:"permission,omitempty"`
				} `tfsdk:"target_grants" json:"targetGrants,omitempty"`
				TargetPrefix *string `tfsdk:"target_prefix" json:"targetPrefix,omitempty"`
			} `tfsdk:"logging_enabled" json:"loggingEnabled,omitempty"`
		} `tfsdk:"logging" json:"logging,omitempty"`
		Metrics *[]struct {
			Filter *struct {
				AccessPointARN *string `tfsdk:"access_point_arn" json:"accessPointARN,omitempty"`
				And            *struct {
					AccessPointARN *string `tfsdk:"access_point_arn" json:"accessPointARN,omitempty"`
					Prefix         *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Tags           *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"and" json:"and,omitempty"`
				Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Tag    *struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"filter" json:"filter,omitempty"`
			Id *string `tfsdk:"id" json:"id,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Name         *string `tfsdk:"name" json:"name,omitempty"`
		Notification *struct {
			LambdaFunctionConfigurations *[]struct {
				Events *[]string `tfsdk:"events" json:"events,omitempty"`
				Filter *struct {
					Key *struct {
						FilterRules *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"filter_rules" json:"filterRules,omitempty"`
					} `tfsdk:"key" json:"key,omitempty"`
				} `tfsdk:"filter" json:"filter,omitempty"`
				Id                *string `tfsdk:"id" json:"id,omitempty"`
				LambdaFunctionARN *string `tfsdk:"lambda_function_arn" json:"lambdaFunctionARN,omitempty"`
			} `tfsdk:"lambda_function_configurations" json:"lambdaFunctionConfigurations,omitempty"`
			QueueConfigurations *[]struct {
				Events *[]string `tfsdk:"events" json:"events,omitempty"`
				Filter *struct {
					Key *struct {
						FilterRules *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"filter_rules" json:"filterRules,omitempty"`
					} `tfsdk:"key" json:"key,omitempty"`
				} `tfsdk:"filter" json:"filter,omitempty"`
				Id       *string `tfsdk:"id" json:"id,omitempty"`
				QueueARN *string `tfsdk:"queue_arn" json:"queueARN,omitempty"`
			} `tfsdk:"queue_configurations" json:"queueConfigurations,omitempty"`
			TopicConfigurations *[]struct {
				Events *[]string `tfsdk:"events" json:"events,omitempty"`
				Filter *struct {
					Key *struct {
						FilterRules *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"filter_rules" json:"filterRules,omitempty"`
					} `tfsdk:"key" json:"key,omitempty"`
				} `tfsdk:"filter" json:"filter,omitempty"`
				Id       *string `tfsdk:"id" json:"id,omitempty"`
				TopicARN *string `tfsdk:"topic_arn" json:"topicARN,omitempty"`
			} `tfsdk:"topic_configurations" json:"topicConfigurations,omitempty"`
		} `tfsdk:"notification" json:"notification,omitempty"`
		ObjectLockEnabledForBucket *bool   `tfsdk:"object_lock_enabled_for_bucket" json:"objectLockEnabledForBucket,omitempty"`
		ObjectOwnership            *string `tfsdk:"object_ownership" json:"objectOwnership,omitempty"`
		OwnershipControls          *struct {
			Rules *[]struct {
				ObjectOwnership *string `tfsdk:"object_ownership" json:"objectOwnership,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"ownership_controls" json:"ownershipControls,omitempty"`
		Policy            *string `tfsdk:"policy" json:"policy,omitempty"`
		PublicAccessBlock *struct {
			BlockPublicACLs       *bool `tfsdk:"block_public_ac_ls" json:"blockPublicACLs,omitempty"`
			BlockPublicPolicy     *bool `tfsdk:"block_public_policy" json:"blockPublicPolicy,omitempty"`
			IgnorePublicACLs      *bool `tfsdk:"ignore_public_ac_ls" json:"ignorePublicACLs,omitempty"`
			RestrictPublicBuckets *bool `tfsdk:"restrict_public_buckets" json:"restrictPublicBuckets,omitempty"`
		} `tfsdk:"public_access_block" json:"publicAccessBlock,omitempty"`
		Replication *struct {
			Role  *string `tfsdk:"role" json:"role,omitempty"`
			Rules *[]struct {
				DeleteMarkerReplication *struct {
					Status *string `tfsdk:"status" json:"status,omitempty"`
				} `tfsdk:"delete_marker_replication" json:"deleteMarkerReplication,omitempty"`
				Destination *struct {
					AccessControlTranslation *struct {
						Owner *string `tfsdk:"owner" json:"owner,omitempty"`
					} `tfsdk:"access_control_translation" json:"accessControlTranslation,omitempty"`
					Account                 *string `tfsdk:"account" json:"account,omitempty"`
					Bucket                  *string `tfsdk:"bucket" json:"bucket,omitempty"`
					EncryptionConfiguration *struct {
						ReplicaKMSKeyID *string `tfsdk:"replica_kms_key_id" json:"replicaKMSKeyID,omitempty"`
					} `tfsdk:"encryption_configuration" json:"encryptionConfiguration,omitempty"`
					Metrics *struct {
						EventThreshold *struct {
							Minutes *int64 `tfsdk:"minutes" json:"minutes,omitempty"`
						} `tfsdk:"event_threshold" json:"eventThreshold,omitempty"`
						Status *string `tfsdk:"status" json:"status,omitempty"`
					} `tfsdk:"metrics" json:"metrics,omitempty"`
					ReplicationTime *struct {
						Status *string `tfsdk:"status" json:"status,omitempty"`
						Time   *struct {
							Minutes *int64 `tfsdk:"minutes" json:"minutes,omitempty"`
						} `tfsdk:"time" json:"time,omitempty"`
					} `tfsdk:"replication_time" json:"replicationTime,omitempty"`
					StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				ExistingObjectReplication *struct {
					Status *string `tfsdk:"status" json:"status,omitempty"`
				} `tfsdk:"existing_object_replication" json:"existingObjectReplication,omitempty"`
				Filter *struct {
					And *struct {
						Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
						Tags   *[]struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tags" json:"tags,omitempty"`
					} `tfsdk:"and" json:"and,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Tag    *struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tag" json:"tag,omitempty"`
				} `tfsdk:"filter" json:"filter,omitempty"`
				Id                      *string `tfsdk:"id" json:"id,omitempty"`
				Prefix                  *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Priority                *int64  `tfsdk:"priority" json:"priority,omitempty"`
				SourceSelectionCriteria *struct {
					ReplicaModifications *struct {
						Status *string `tfsdk:"status" json:"status,omitempty"`
					} `tfsdk:"replica_modifications" json:"replicaModifications,omitempty"`
					SseKMSEncryptedObjects *struct {
						Status *string `tfsdk:"status" json:"status,omitempty"`
					} `tfsdk:"sse_kms_encrypted_objects" json:"sseKMSEncryptedObjects,omitempty"`
				} `tfsdk:"source_selection_criteria" json:"sourceSelectionCriteria,omitempty"`
				Status *string `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"rules" json:"rules,omitempty"`
		} `tfsdk:"replication" json:"replication,omitempty"`
		RequestPayment *struct {
			Payer *string `tfsdk:"payer" json:"payer,omitempty"`
		} `tfsdk:"request_payment" json:"requestPayment,omitempty"`
		Tagging *struct {
			TagSet *[]struct {
				Key   *string `tfsdk:"key" json:"key,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tag_set" json:"tagSet,omitempty"`
		} `tfsdk:"tagging" json:"tagging,omitempty"`
		Versioning *struct {
			Status *string `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"versioning" json:"versioning,omitempty"`
		Website *struct {
			ErrorDocument *struct {
				Key *string `tfsdk:"key" json:"key,omitempty"`
			} `tfsdk:"error_document" json:"errorDocument,omitempty"`
			IndexDocument *struct {
				Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
			} `tfsdk:"index_document" json:"indexDocument,omitempty"`
			RedirectAllRequestsTo *struct {
				HostName *string `tfsdk:"host_name" json:"hostName,omitempty"`
				Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
			} `tfsdk:"redirect_all_requests_to" json:"redirectAllRequestsTo,omitempty"`
			RoutingRules *[]struct {
				Condition *struct {
					HttpErrorCodeReturnedEquals *string `tfsdk:"http_error_code_returned_equals" json:"httpErrorCodeReturnedEquals,omitempty"`
					KeyPrefixEquals             *string `tfsdk:"key_prefix_equals" json:"keyPrefixEquals,omitempty"`
				} `tfsdk:"condition" json:"condition,omitempty"`
				Redirect *struct {
					HostName             *string `tfsdk:"host_name" json:"hostName,omitempty"`
					HttpRedirectCode     *string `tfsdk:"http_redirect_code" json:"httpRedirectCode,omitempty"`
					Protocol             *string `tfsdk:"protocol" json:"protocol,omitempty"`
					ReplaceKeyPrefixWith *string `tfsdk:"replace_key_prefix_with" json:"replaceKeyPrefixWith,omitempty"`
					ReplaceKeyWith       *string `tfsdk:"replace_key_with" json:"replaceKeyWith,omitempty"`
				} `tfsdk:"redirect" json:"redirect,omitempty"`
			} `tfsdk:"routing_rules" json:"routingRules,omitempty"`
		} `tfsdk:"website" json:"website,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *S3ServicesK8SAwsBucketV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_s3_services_k8s_aws_bucket_v1alpha1_manifest"
}

func (r *S3ServicesK8SAwsBucketV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Bucket is the Schema for the Buckets API",
		MarkdownDescription: "Bucket is the Schema for the Buckets API",
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
				Description:         "BucketSpec defines the desired state of Bucket.  In terms of implementation, a Bucket is a resource. An Amazon S3 bucket name is globally unique, and the namespace is shared by all Amazon Web Services accounts.",
				MarkdownDescription: "BucketSpec defines the desired state of Bucket.  In terms of implementation, a Bucket is a resource. An Amazon S3 bucket name is globally unique, and the namespace is shared by all Amazon Web Services accounts.",
				Attributes: map[string]schema.Attribute{
					"accelerate": schema.SingleNestedAttribute{
						Description:         "Container for setting the transfer acceleration state.",
						MarkdownDescription: "Container for setting the transfer acceleration state.",
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"acl": schema.StringAttribute{
						Description:         "The canned ACL to apply to the bucket.",
						MarkdownDescription: "The canned ACL to apply to the bucket.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"analytics": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"filter": schema.SingleNestedAttribute{
									Description:         "The filter used to describe a set of objects for analyses. A filter must have exactly one prefix, one tag, or one conjunction (AnalyticsAndOperator). If no filter is provided, all objects will be considered in any analysis.",
									MarkdownDescription: "The filter used to describe a set of objects for analyses. A filter must have exactly one prefix, one tag, or one conjunction (AnalyticsAndOperator). If no filter is provided, all objects will be considered in any analysis.",
									Attributes: map[string]schema.Attribute{
										"and": schema.SingleNestedAttribute{
											Description:         "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates in any combination, and an object must match all of the predicates for the filter to apply.",
											MarkdownDescription: "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates in any combination, and an object must match all of the predicates for the filter to apply.",
											Attributes: map[string]schema.Attribute{
												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tags": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
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

										"prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.SingleNestedAttribute{
											Description:         "A container of a key value name pair.",
											MarkdownDescription: "A container of a key value name pair.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"storage_class_analysis": schema.SingleNestedAttribute{
									Description:         "Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.",
									MarkdownDescription: "Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.",
									Attributes: map[string]schema.Attribute{
										"data_export": schema.SingleNestedAttribute{
											Description:         "Container for data related to the storage class analysis for an Amazon S3 bucket for export.",
											MarkdownDescription: "Container for data related to the storage class analysis for an Amazon S3 bucket for export.",
											Attributes: map[string]schema.Attribute{
												"destination": schema.SingleNestedAttribute{
													Description:         "Where to publish the analytics results.",
													MarkdownDescription: "Where to publish the analytics results.",
													Attributes: map[string]schema.Attribute{
														"s3_bucket_destination": schema.SingleNestedAttribute{
															Description:         "Contains information about where to publish the analytics results.",
															MarkdownDescription: "Contains information about where to publish the analytics results.",
															Attributes: map[string]schema.Attribute{
																"bucket": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"bucket_account_id": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"format": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"prefix": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"output_schema_version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cors": schema.SingleNestedAttribute{
						Description:         "Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon S3 User Guide.",
						MarkdownDescription: "Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon S3 User Guide.",
						Attributes: map[string]schema.Attribute{
							"cors_rules": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"allowed_headers": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allowed_methods": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allowed_origins": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"expose_headers": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_age_seconds": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

					"create_bucket_configuration": schema.SingleNestedAttribute{
						Description:         "The configuration information for the bucket.",
						MarkdownDescription: "The configuration information for the bucket.",
						Attributes: map[string]schema.Attribute{
							"location_constraint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"encryption": schema.SingleNestedAttribute{
						Description:         "Specifies the default server-side-encryption configuration.",
						MarkdownDescription: "Specifies the default server-side-encryption configuration.",
						Attributes: map[string]schema.Attribute{
							"rules": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"apply_server_side_encryption_by_default": schema.SingleNestedAttribute{
											Description:         "Describes the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied. If you don't specify a customer managed key at configuration, Amazon S3 automatically creates an Amazon Web Services KMS key in your Amazon Web Services account the first time that you add an object encrypted with SSE-KMS to a bucket. By default, Amazon S3 uses this KMS key for SSE-KMS. For more information, see PUT Bucket encryption (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTencryption.html) in the Amazon S3 API Reference.",
											MarkdownDescription: "Describes the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied. If you don't specify a customer managed key at configuration, Amazon S3 automatically creates an Amazon Web Services KMS key in your Amazon Web Services account the first time that you add an object encrypted with SSE-KMS to a bucket. By default, Amazon S3 uses this KMS key for SSE-KMS. For more information, see PUT Bucket encryption (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTencryption.html) in the Amazon S3 API Reference.",
											Attributes: map[string]schema.Attribute{
												"kms_master_key_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sse_algorithm": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"bucket_key_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

					"grant_full_control": schema.StringAttribute{
						Description:         "Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.",
						MarkdownDescription: "Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"grant_read": schema.StringAttribute{
						Description:         "Allows grantee to list the objects in the bucket.",
						MarkdownDescription: "Allows grantee to list the objects in the bucket.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"grant_read_acp": schema.StringAttribute{
						Description:         "Allows grantee to read the bucket ACL.",
						MarkdownDescription: "Allows grantee to read the bucket ACL.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"grant_write": schema.StringAttribute{
						Description:         "Allows grantee to create new objects in the bucket.  For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.",
						MarkdownDescription: "Allows grantee to create new objects in the bucket.  For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"grant_write_acp": schema.StringAttribute{
						Description:         "Allows grantee to write the ACL for the applicable bucket.",
						MarkdownDescription: "Allows grantee to write the ACL for the applicable bucket.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"intelligent_tiering": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"filter": schema.SingleNestedAttribute{
									Description:         "The Filter is used to identify objects that the S3 Intelligent-Tiering configuration applies to.",
									MarkdownDescription: "The Filter is used to identify objects that the S3 Intelligent-Tiering configuration applies to.",
									Attributes: map[string]schema.Attribute{
										"and": schema.SingleNestedAttribute{
											Description:         "A container for specifying S3 Intelligent-Tiering filters. The filters determine the subset of objects to which the rule applies.",
											MarkdownDescription: "A container for specifying S3 Intelligent-Tiering filters. The filters determine the subset of objects to which the rule applies.",
											Attributes: map[string]schema.Attribute{
												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tags": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
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

										"prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.SingleNestedAttribute{
											Description:         "A container of a key value name pair.",
											MarkdownDescription: "A container of a key value name pair.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"status": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tierings": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"access_tier": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"days": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"inventory": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"destination": schema.SingleNestedAttribute{
									Description:         "Specifies the inventory configuration for an Amazon S3 bucket.",
									MarkdownDescription: "Specifies the inventory configuration for an Amazon S3 bucket.",
									Attributes: map[string]schema.Attribute{
										"s3_bucket_destination": schema.SingleNestedAttribute{
											Description:         "Contains the bucket name, file format, bucket owner (optional), and prefix (optional) where inventory results are published.",
											MarkdownDescription: "Contains the bucket name, file format, bucket owner (optional), and prefix (optional) where inventory results are published.",
											Attributes: map[string]schema.Attribute{
												"account_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"bucket": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"encryption": schema.SingleNestedAttribute{
													Description:         "Contains the type of server-side encryption used to encrypt the inventory results.",
													MarkdownDescription: "Contains the type of server-side encryption used to encrypt the inventory results.",
													Attributes: map[string]schema.Attribute{
														"sse_kms": schema.SingleNestedAttribute{
															Description:         "Specifies the use of SSE-KMS to encrypt delivered inventory reports.",
															MarkdownDescription: "Specifies the use of SSE-KMS to encrypt delivered inventory reports.",
															Attributes: map[string]schema.Attribute{
																"key_id": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"format": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"filter": schema.SingleNestedAttribute{
									Description:         "Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria.",
									MarkdownDescription: "Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria.",
									Attributes: map[string]schema.Attribute{
										"prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"included_object_versions": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"is_enabled": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"optional_fields": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"schedule": schema.SingleNestedAttribute{
									Description:         "Specifies the schedule for generating inventory results.",
									MarkdownDescription: "Specifies the schedule for generating inventory results.",
									Attributes: map[string]schema.Attribute{
										"frequency": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"lifecycle": schema.SingleNestedAttribute{
						Description:         "Container for lifecycle rules. You can add as many as 1,000 rules.",
						MarkdownDescription: "Container for lifecycle rules. You can add as many as 1,000 rules.",
						Attributes: map[string]schema.Attribute{
							"rules": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"abort_incomplete_multipart_upload": schema.SingleNestedAttribute{
											Description:         "Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. For more information, see Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) in the Amazon S3 User Guide.",
											MarkdownDescription: "Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. For more information, see Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) in the Amazon S3 User Guide.",
											Attributes: map[string]schema.Attribute{
												"days_after_initiation": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"expiration": schema.SingleNestedAttribute{
											Description:         "Container for the expiration for the lifecycle of the object.",
											MarkdownDescription: "Container for the expiration for the lifecycle of the object.",
											Attributes: map[string]schema.Attribute{
												"date": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														validators.DateTime64Validator(),
													},
												},

												"days": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expired_object_delete_marker": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"filter": schema.SingleNestedAttribute{
											Description:         "The Filter is used to identify objects that a Lifecycle Rule applies to. A Filter must have exactly one of Prefix, Tag, or And specified.",
											MarkdownDescription: "The Filter is used to identify objects that a Lifecycle Rule applies to. A Filter must have exactly one of Prefix, Tag, or And specified.",
											Attributes: map[string]schema.Attribute{
												"and": schema.SingleNestedAttribute{
													Description:         "This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will apply to any object matching all of the predicates configured inside the And operator.",
													MarkdownDescription: "This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will apply to any object matching all of the predicates configured inside the And operator.",
													Attributes: map[string]schema.Attribute{
														"object_size_greater_than": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"object_size_less_than": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"prefix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tags": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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

												"object_size_greater_than": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"object_size_less_than": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tag": schema.SingleNestedAttribute{
													Description:         "A container of a key value name pair.",
													MarkdownDescription: "A container of a key value name pair.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"noncurrent_version_expiration": schema.SingleNestedAttribute{
											Description:         "Specifies when noncurrent object versions expire. Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.",
											MarkdownDescription: "Specifies when noncurrent object versions expire. Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.",
											Attributes: map[string]schema.Attribute{
												"newer_noncurrent_versions": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"noncurrent_days": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"noncurrent_version_transitions": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"newer_noncurrent_versions": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"noncurrent_days": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"storage_class": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"status": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"transitions": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"date": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.DateTime64Validator(),
														},
													},

													"days": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"storage_class": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
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

					"logging": schema.SingleNestedAttribute{
						Description:         "Container for logging status information.",
						MarkdownDescription: "Container for logging status information.",
						Attributes: map[string]schema.Attribute{
							"logging_enabled": schema.SingleNestedAttribute{
								Description:         "Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket. For more information, see PUT Bucket logging (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the Amazon S3 API Reference.",
								MarkdownDescription: "Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket. For more information, see PUT Bucket logging (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the Amazon S3 API Reference.",
								Attributes: map[string]schema.Attribute{
									"target_bucket": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_grants": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"grantee": schema.SingleNestedAttribute{
													Description:         "Container for the person being granted permissions.",
													MarkdownDescription: "Container for the person being granted permissions.",
													Attributes: map[string]schema.Attribute{
														"display_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"email_address": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type_": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"u_ri": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"permission": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_prefix": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"filter": schema.SingleNestedAttribute{
									Description:         "Specifies a metrics configuration filter. The metrics configuration only includes objects that meet the filter's criteria. A filter must be a prefix, an object tag, an access point ARN, or a conjunction (MetricsAndOperator). For more information, see PutBucketMetricsConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html).",
									MarkdownDescription: "Specifies a metrics configuration filter. The metrics configuration only includes objects that meet the filter's criteria. A filter must be a prefix, an object tag, an access point ARN, or a conjunction (MetricsAndOperator). For more information, see PutBucketMetricsConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html).",
									Attributes: map[string]schema.Attribute{
										"access_point_arn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"and": schema.SingleNestedAttribute{
											Description:         "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates, and an object must match all of the predicates in order for the filter to apply.",
											MarkdownDescription: "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates, and an object must match all of the predicates in order for the filter to apply.",
											Attributes: map[string]schema.Attribute{
												"access_point_arn": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tags": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
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

										"prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.SingleNestedAttribute{
											Description:         "A container of a key value name pair.",
											MarkdownDescription: "A container of a key value name pair.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the bucket to create.",
						MarkdownDescription: "The name of the bucket to create.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"notification": schema.SingleNestedAttribute{
						Description:         "A container for specifying the notification configuration of the bucket. If this element is empty, notifications are turned off for the bucket.",
						MarkdownDescription: "A container for specifying the notification configuration of the bucket. If this element is empty, notifications are turned off for the bucket.",
						Attributes: map[string]schema.Attribute{
							"lambda_function_configurations": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"events": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.SingleNestedAttribute{
											Description:         "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
											MarkdownDescription: "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
											Attributes: map[string]schema.Attribute{
												"key": schema.SingleNestedAttribute{
													Description:         "A container for object key name prefix and suffix filtering rules.",
													MarkdownDescription: "A container for object key name prefix and suffix filtering rules.",
													Attributes: map[string]schema.Attribute{
														"filter_rules": schema.ListNestedAttribute{
															Description:         "A list of containers for the key-value pair that defines the criteria for the filter rule.",
															MarkdownDescription: "A list of containers for the key-value pair that defines the criteria for the filter rule.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"id": schema.StringAttribute{
											Description:         "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
											MarkdownDescription: "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lambda_function_arn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"queue_configurations": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"events": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.SingleNestedAttribute{
											Description:         "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
											MarkdownDescription: "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
											Attributes: map[string]schema.Attribute{
												"key": schema.SingleNestedAttribute{
													Description:         "A container for object key name prefix and suffix filtering rules.",
													MarkdownDescription: "A container for object key name prefix and suffix filtering rules.",
													Attributes: map[string]schema.Attribute{
														"filter_rules": schema.ListNestedAttribute{
															Description:         "A list of containers for the key-value pair that defines the criteria for the filter rule.",
															MarkdownDescription: "A list of containers for the key-value pair that defines the criteria for the filter rule.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"id": schema.StringAttribute{
											Description:         "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
											MarkdownDescription: "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"queue_arn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"topic_configurations": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"events": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.SingleNestedAttribute{
											Description:         "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
											MarkdownDescription: "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
											Attributes: map[string]schema.Attribute{
												"key": schema.SingleNestedAttribute{
													Description:         "A container for object key name prefix and suffix filtering rules.",
													MarkdownDescription: "A container for object key name prefix and suffix filtering rules.",
													Attributes: map[string]schema.Attribute{
														"filter_rules": schema.ListNestedAttribute{
															Description:         "A list of containers for the key-value pair that defines the criteria for the filter rule.",
															MarkdownDescription: "A list of containers for the key-value pair that defines the criteria for the filter rule.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"id": schema.StringAttribute{
											Description:         "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
											MarkdownDescription: "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"topic_arn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

					"object_lock_enabled_for_bucket": schema.BoolAttribute{
						Description:         "Specifies whether you want S3 Object Lock to be enabled for the new bucket.",
						MarkdownDescription: "Specifies whether you want S3 Object Lock to be enabled for the new bucket.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"object_ownership": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ownership_controls": schema.SingleNestedAttribute{
						Description:         "The OwnershipControls (BucketOwnerEnforced, BucketOwnerPreferred, or ObjectWriter) that you want to apply to this Amazon S3 bucket.",
						MarkdownDescription: "The OwnershipControls (BucketOwnerEnforced, BucketOwnerPreferred, or ObjectWriter) that you want to apply to this Amazon S3 bucket.",
						Attributes: map[string]schema.Attribute{
							"rules": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"object_ownership": schema.StringAttribute{
											Description:         "The container element for object ownership for a bucket's ownership controls.  BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the bucket owner if the objects are uploaded with the bucket-owner-full-control canned ACL.  ObjectWriter - The uploading account will own the object if the object is uploaded with the bucket-owner-full-control canned ACL.  BucketOwnerEnforced - Access control lists (ACLs) are disabled and no longer affect permissions. The bucket owner automatically owns and has full control over every object in the bucket. The bucket only accepts PUT requests that don't specify an ACL or bucket owner full control ACLs, such as the bucket-owner-full-control canned ACL or an equivalent form of this ACL expressed in the XML format.",
											MarkdownDescription: "The container element for object ownership for a bucket's ownership controls.  BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the bucket owner if the objects are uploaded with the bucket-owner-full-control canned ACL.  ObjectWriter - The uploading account will own the object if the object is uploaded with the bucket-owner-full-control canned ACL.  BucketOwnerEnforced - Access control lists (ACLs) are disabled and no longer affect permissions. The bucket owner automatically owns and has full control over every object in the bucket. The bucket only accepts PUT requests that don't specify an ACL or bucket owner full control ACLs, such as the bucket-owner-full-control canned ACL or an equivalent form of this ACL expressed in the XML format.",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

					"policy": schema.StringAttribute{
						Description:         "The bucket policy as a JSON document.",
						MarkdownDescription: "The bucket policy as a JSON document.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"public_access_block": schema.SingleNestedAttribute{
						Description:         "The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see The Meaning of 'Public' (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the Amazon S3 User Guide.",
						MarkdownDescription: "The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see The Meaning of 'Public' (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the Amazon S3 User Guide.",
						Attributes: map[string]schema.Attribute{
							"block_public_ac_ls": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"block_public_policy": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ignore_public_ac_ls": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"restrict_public_buckets": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replication": schema.SingleNestedAttribute{
						Description:         "A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.",
						MarkdownDescription: "A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.",
						Attributes: map[string]schema.Attribute{
							"role": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rules": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"delete_marker_replication": schema.SingleNestedAttribute{
											Description:         "Specifies whether Amazon S3 replicates delete markers. If you specify a Filter in your replication configuration, you must also include a DeleteMarkerReplication element. If your Filter includes a Tag element, the DeleteMarkerReplication Status must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config).  For more information about delete marker replication, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html).  If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see Backward Compatibility (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations).",
											MarkdownDescription: "Specifies whether Amazon S3 replicates delete markers. If you specify a Filter in your replication configuration, you must also include a DeleteMarkerReplication element. If your Filter includes a Tag element, the DeleteMarkerReplication Status must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config).  For more information about delete marker replication, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html).  If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see Backward Compatibility (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations).",
											Attributes: map[string]schema.Attribute{
												"status": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"destination": schema.SingleNestedAttribute{
											Description:         "Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).",
											MarkdownDescription: "Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).",
											Attributes: map[string]schema.Attribute{
												"access_control_translation": schema.SingleNestedAttribute{
													Description:         "A container for information about access control for replicas.",
													MarkdownDescription: "A container for information about access control for replicas.",
													Attributes: map[string]schema.Attribute{
														"owner": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"account": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"bucket": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"encryption_configuration": schema.SingleNestedAttribute{
													Description:         "Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.",
													MarkdownDescription: "Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.",
													Attributes: map[string]schema.Attribute{
														"replica_kms_key_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"metrics": schema.SingleNestedAttribute{
													Description:         "A container specifying replication metrics-related settings enabling replication metrics and events.",
													MarkdownDescription: "A container specifying replication metrics-related settings enabling replication metrics and events.",
													Attributes: map[string]schema.Attribute{
														"event_threshold": schema.SingleNestedAttribute{
															Description:         "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",
															MarkdownDescription: "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",
															Attributes: map[string]schema.Attribute{
																"minutes": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"status": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"replication_time": schema.SingleNestedAttribute{
													Description:         "A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a Metrics block.",
													MarkdownDescription: "A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a Metrics block.",
													Attributes: map[string]schema.Attribute{
														"status": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"time": schema.SingleNestedAttribute{
															Description:         "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",
															MarkdownDescription: "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",
															Attributes: map[string]schema.Attribute{
																"minutes": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"storage_class": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"existing_object_replication": schema.SingleNestedAttribute{
											Description:         "Optional configuration to replicate existing source bucket objects. For more information, see Replicating Existing Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-what-is-isnot-replicated.html#existing-object-replication) in the Amazon S3 User Guide.",
											MarkdownDescription: "Optional configuration to replicate existing source bucket objects. For more information, see Replicating Existing Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-what-is-isnot-replicated.html#existing-object-replication) in the Amazon S3 User Guide.",
											Attributes: map[string]schema.Attribute{
												"status": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"filter": schema.SingleNestedAttribute{
											Description:         "A filter that identifies the subset of objects to which the replication rule applies. A Filter must specify exactly one Prefix, Tag, or an And child element.",
											MarkdownDescription: "A filter that identifies the subset of objects to which the replication rule applies. A Filter must specify exactly one Prefix, Tag, or an And child element.",
											Attributes: map[string]schema.Attribute{
												"and": schema.SingleNestedAttribute{
													Description:         "A container for specifying rule filters. The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter.  For example:  * If you specify both a Prefix and a Tag filter, wrap these filters in an And tag.  * If you specify a filter based on multiple tags, wrap the Tag elements in an And tag.",
													MarkdownDescription: "A container for specifying rule filters. The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter.  For example:  * If you specify both a Prefix and a Tag filter, wrap these filters in an And tag.  * If you specify a filter based on multiple tags, wrap the Tag elements in an And tag.",
													Attributes: map[string]schema.Attribute{
														"prefix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tags": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tag": schema.SingleNestedAttribute{
													Description:         "A container of a key value name pair.",
													MarkdownDescription: "A container of a key value name pair.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"prefix": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"priority": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_selection_criteria": schema.SingleNestedAttribute{
											Description:         "A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects. Currently, Amazon S3 supports only the filter that you can specify for objects created with server-side encryption using a customer managed key stored in Amazon Web Services Key Management Service (SSE-KMS).",
											MarkdownDescription: "A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects. Currently, Amazon S3 supports only the filter that you can specify for objects created with server-side encryption using a customer managed key stored in Amazon Web Services Key Management Service (SSE-KMS).",
											Attributes: map[string]schema.Attribute{
												"replica_modifications": schema.SingleNestedAttribute{
													Description:         "A filter that you can specify for selection for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when Filter is specified), you can specify this element and set the status to Enabled to replicate modifications on replicas.  If you don't specify the Filter element, Amazon S3 assumes that the replication configuration is the earlier version, V1. In the earlier version, this element is not allowed.",
													MarkdownDescription: "A filter that you can specify for selection for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when Filter is specified), you can specify this element and set the status to Enabled to replicate modifications on replicas.  If you don't specify the Filter element, Amazon S3 assumes that the replication configuration is the earlier version, V1. In the earlier version, this element is not allowed.",
													Attributes: map[string]schema.Attribute{
														"status": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"sse_kms_encrypted_objects": schema.SingleNestedAttribute{
													Description:         "A container for filter information for the selection of S3 objects encrypted with Amazon Web Services KMS.",
													MarkdownDescription: "A container for filter information for the selection of S3 objects encrypted with Amazon Web Services KMS.",
													Attributes: map[string]schema.Attribute{
														"status": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"status": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

					"request_payment": schema.SingleNestedAttribute{
						Description:         "Container for Payer.",
						MarkdownDescription: "Container for Payer.",
						Attributes: map[string]schema.Attribute{
							"payer": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tagging": schema.SingleNestedAttribute{
						Description:         "Container for the TagSet and Tag elements.",
						MarkdownDescription: "Container for the TagSet and Tag elements.",
						Attributes: map[string]schema.Attribute{
							"tag_set": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
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

					"versioning": schema.SingleNestedAttribute{
						Description:         "Container for setting the versioning state.",
						MarkdownDescription: "Container for setting the versioning state.",
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"website": schema.SingleNestedAttribute{
						Description:         "Container for the request.",
						MarkdownDescription: "Container for the request.",
						Attributes: map[string]schema.Attribute{
							"error_document": schema.SingleNestedAttribute{
								Description:         "The error information.",
								MarkdownDescription: "The error information.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"index_document": schema.SingleNestedAttribute{
								Description:         "Container for the Suffix element.",
								MarkdownDescription: "Container for the Suffix element.",
								Attributes: map[string]schema.Attribute{
									"suffix": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"redirect_all_requests_to": schema.SingleNestedAttribute{
								Description:         "Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.",
								MarkdownDescription: "Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.",
								Attributes: map[string]schema.Attribute{
									"host_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"protocol": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"routing_rules": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition": schema.SingleNestedAttribute{
											Description:         "A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the /docs folder, redirect to the /documents folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.",
											MarkdownDescription: "A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the /docs folder, redirect to the /documents folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.",
											Attributes: map[string]schema.Attribute{
												"http_error_code_returned_equals": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key_prefix_equals": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"redirect": schema.SingleNestedAttribute{
											Description:         "Specifies how requests are redirected. In the event of an error, you can specify a different error code to return.",
											MarkdownDescription: "Specifies how requests are redirected. In the event of an error, you can specify a different error code to return.",
											Attributes: map[string]schema.Attribute{
												"host_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_redirect_code": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replace_key_prefix_with": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replace_key_with": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *S3ServicesK8SAwsBucketV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_s3_services_k8s_aws_bucket_v1alpha1_manifest")

	var model S3ServicesK8SAwsBucketV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("s3.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Bucket")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
