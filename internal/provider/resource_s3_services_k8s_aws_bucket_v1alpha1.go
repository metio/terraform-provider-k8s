/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type S3ServicesK8SAwsBucketV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*S3ServicesK8SAwsBucketV1Alpha1Resource)(nil)
)

type S3ServicesK8SAwsBucketV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type S3ServicesK8SAwsBucketV1Alpha1GoModel struct {
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
		Accelerate *struct {
			Status *string `tfsdk:"status" yaml:"status,omitempty"`
		} `tfsdk:"accelerate" yaml:"accelerate,omitempty"`

		Acl *string `tfsdk:"acl" yaml:"acl,omitempty"`

		Analytics *[]struct {
			Filter *struct {
				And *struct {
					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Tags *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tags" yaml:"tags,omitempty"`
				} `tfsdk:"and" yaml:"and,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				Tag *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tag" yaml:"tag,omitempty"`
			} `tfsdk:"filter" yaml:"filter,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			StorageClassAnalysis *struct {
				DataExport *struct {
					Destination *struct {
						S3BucketDestination *struct {
							Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

							BucketAccountID *string `tfsdk:"bucket_account_id" yaml:"bucketAccountID,omitempty"`

							Format *string `tfsdk:"format" yaml:"format,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
						} `tfsdk:"s3_bucket_destination" yaml:"s3BucketDestination,omitempty"`
					} `tfsdk:"destination" yaml:"destination,omitempty"`

					OutputSchemaVersion *string `tfsdk:"output_schema_version" yaml:"outputSchemaVersion,omitempty"`
				} `tfsdk:"data_export" yaml:"dataExport,omitempty"`
			} `tfsdk:"storage_class_analysis" yaml:"storageClassAnalysis,omitempty"`
		} `tfsdk:"analytics" yaml:"analytics,omitempty"`

		Cors *struct {
			CorsRules *[]struct {
				AllowedHeaders *[]string `tfsdk:"allowed_headers" yaml:"allowedHeaders,omitempty"`

				AllowedMethods *[]string `tfsdk:"allowed_methods" yaml:"allowedMethods,omitempty"`

				AllowedOrigins *[]string `tfsdk:"allowed_origins" yaml:"allowedOrigins,omitempty"`

				ExposeHeaders *[]string `tfsdk:"expose_headers" yaml:"exposeHeaders,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				MaxAgeSeconds *int64 `tfsdk:"max_age_seconds" yaml:"maxAgeSeconds,omitempty"`
			} `tfsdk:"cors_rules" yaml:"corsRules,omitempty"`
		} `tfsdk:"cors" yaml:"cors,omitempty"`

		CreateBucketConfiguration *struct {
			LocationConstraint *string `tfsdk:"location_constraint" yaml:"locationConstraint,omitempty"`
		} `tfsdk:"create_bucket_configuration" yaml:"createBucketConfiguration,omitempty"`

		Encryption *struct {
			Rules *[]struct {
				ApplyServerSideEncryptionByDefault *struct {
					KmsMasterKeyID *string `tfsdk:"kms_master_key_id" yaml:"kmsMasterKeyID,omitempty"`

					SseAlgorithm *string `tfsdk:"sse_algorithm" yaml:"sseAlgorithm,omitempty"`
				} `tfsdk:"apply_server_side_encryption_by_default" yaml:"applyServerSideEncryptionByDefault,omitempty"`

				BucketKeyEnabled *bool `tfsdk:"bucket_key_enabled" yaml:"bucketKeyEnabled,omitempty"`
			} `tfsdk:"rules" yaml:"rules,omitempty"`
		} `tfsdk:"encryption" yaml:"encryption,omitempty"`

		GrantFullControl *string `tfsdk:"grant_full_control" yaml:"grantFullControl,omitempty"`

		GrantRead *string `tfsdk:"grant_read" yaml:"grantRead,omitempty"`

		GrantReadACP *string `tfsdk:"grant_read_acp" yaml:"grantReadACP,omitempty"`

		GrantWrite *string `tfsdk:"grant_write" yaml:"grantWrite,omitempty"`

		GrantWriteACP *string `tfsdk:"grant_write_acp" yaml:"grantWriteACP,omitempty"`

		IntelligentTiering *[]struct {
			Filter *struct {
				And *struct {
					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Tags *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tags" yaml:"tags,omitempty"`
				} `tfsdk:"and" yaml:"and,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				Tag *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tag" yaml:"tag,omitempty"`
			} `tfsdk:"filter" yaml:"filter,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			Status *string `tfsdk:"status" yaml:"status,omitempty"`

			Tierings *[]struct {
				AccessTier *string `tfsdk:"access_tier" yaml:"accessTier,omitempty"`

				Days *int64 `tfsdk:"days" yaml:"days,omitempty"`
			} `tfsdk:"tierings" yaml:"tierings,omitempty"`
		} `tfsdk:"intelligent_tiering" yaml:"intelligentTiering,omitempty"`

		Inventory *[]struct {
			Destination *struct {
				S3BucketDestination *struct {
					AccountID *string `tfsdk:"account_id" yaml:"accountID,omitempty"`

					Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

					Encryption *struct {
						SseKMS *struct {
							KeyID *string `tfsdk:"key_id" yaml:"keyID,omitempty"`
						} `tfsdk:"sse_kms" yaml:"sseKMS,omitempty"`
					} `tfsdk:"encryption" yaml:"encryption,omitempty"`

					Format *string `tfsdk:"format" yaml:"format,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
				} `tfsdk:"s3_bucket_destination" yaml:"s3BucketDestination,omitempty"`
			} `tfsdk:"destination" yaml:"destination,omitempty"`

			Filter *struct {
				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
			} `tfsdk:"filter" yaml:"filter,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			IncludedObjectVersions *string `tfsdk:"included_object_versions" yaml:"includedObjectVersions,omitempty"`

			IsEnabled *bool `tfsdk:"is_enabled" yaml:"isEnabled,omitempty"`

			OptionalFields *[]string `tfsdk:"optional_fields" yaml:"optionalFields,omitempty"`

			Schedule *struct {
				Frequency *string `tfsdk:"frequency" yaml:"frequency,omitempty"`
			} `tfsdk:"schedule" yaml:"schedule,omitempty"`
		} `tfsdk:"inventory" yaml:"inventory,omitempty"`

		Lifecycle *struct {
			Rules *[]struct {
				AbortIncompleteMultipartUpload *struct {
					DaysAfterInitiation *int64 `tfsdk:"days_after_initiation" yaml:"daysAfterInitiation,omitempty"`
				} `tfsdk:"abort_incomplete_multipart_upload" yaml:"abortIncompleteMultipartUpload,omitempty"`

				Expiration *struct {
					Date *string `tfsdk:"date" yaml:"date,omitempty"`

					Days *int64 `tfsdk:"days" yaml:"days,omitempty"`

					ExpiredObjectDeleteMarker *bool `tfsdk:"expired_object_delete_marker" yaml:"expiredObjectDeleteMarker,omitempty"`
				} `tfsdk:"expiration" yaml:"expiration,omitempty"`

				Filter *struct {
					And *struct {
						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						Tags *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"tags" yaml:"tags,omitempty"`
					} `tfsdk:"and" yaml:"and,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Tag *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tag" yaml:"tag,omitempty"`
				} `tfsdk:"filter" yaml:"filter,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				NoncurrentVersionExpiration *struct {
					NoncurrentDays *int64 `tfsdk:"noncurrent_days" yaml:"noncurrentDays,omitempty"`
				} `tfsdk:"noncurrent_version_expiration" yaml:"noncurrentVersionExpiration,omitempty"`

				NoncurrentVersionTransitions *[]struct {
					NoncurrentDays *int64 `tfsdk:"noncurrent_days" yaml:"noncurrentDays,omitempty"`

					StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`
				} `tfsdk:"noncurrent_version_transitions" yaml:"noncurrentVersionTransitions,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				Status *string `tfsdk:"status" yaml:"status,omitempty"`

				Transitions *[]struct {
					Date *string `tfsdk:"date" yaml:"date,omitempty"`

					Days *int64 `tfsdk:"days" yaml:"days,omitempty"`

					StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`
				} `tfsdk:"transitions" yaml:"transitions,omitempty"`
			} `tfsdk:"rules" yaml:"rules,omitempty"`
		} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

		Logging *struct {
			LoggingEnabled *struct {
				TargetBucket *string `tfsdk:"target_bucket" yaml:"targetBucket,omitempty"`

				TargetGrants *[]struct {
					Grantee *struct {
						DisplayName *string `tfsdk:"display_name" yaml:"displayName,omitempty"`

						EmailAddress *string `tfsdk:"email_address" yaml:"emailAddress,omitempty"`

						Id *string `tfsdk:"id" yaml:"id,omitempty"`

						Type_ *string `tfsdk:"type_" yaml:"type_,omitempty"`

						URI *string `tfsdk:"u_ri" yaml:"uRI,omitempty"`
					} `tfsdk:"grantee" yaml:"grantee,omitempty"`

					Permission *string `tfsdk:"permission" yaml:"permission,omitempty"`
				} `tfsdk:"target_grants" yaml:"targetGrants,omitempty"`

				TargetPrefix *string `tfsdk:"target_prefix" yaml:"targetPrefix,omitempty"`
			} `tfsdk:"logging_enabled" yaml:"loggingEnabled,omitempty"`
		} `tfsdk:"logging" yaml:"logging,omitempty"`

		Metrics *[]struct {
			Filter *struct {
				AccessPointARN *string `tfsdk:"access_point_arn" yaml:"accessPointARN,omitempty"`

				And *struct {
					AccessPointARN *string `tfsdk:"access_point_arn" yaml:"accessPointARN,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Tags *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tags" yaml:"tags,omitempty"`
				} `tfsdk:"and" yaml:"and,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				Tag *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tag" yaml:"tag,omitempty"`
			} `tfsdk:"filter" yaml:"filter,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`
		} `tfsdk:"metrics" yaml:"metrics,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Notification *struct {
			LambdaFunctionConfigurations *[]struct {
				Events *[]string `tfsdk:"events" yaml:"events,omitempty"`

				Filter *struct {
					Key *struct {
						FilterRules *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"filter_rules" yaml:"filterRules,omitempty"`
					} `tfsdk:"key" yaml:"key,omitempty"`
				} `tfsdk:"filter" yaml:"filter,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				LambdaFunctionARN *string `tfsdk:"lambda_function_arn" yaml:"lambdaFunctionARN,omitempty"`
			} `tfsdk:"lambda_function_configurations" yaml:"lambdaFunctionConfigurations,omitempty"`

			QueueConfigurations *[]struct {
				Events *[]string `tfsdk:"events" yaml:"events,omitempty"`

				Filter *struct {
					Key *struct {
						FilterRules *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"filter_rules" yaml:"filterRules,omitempty"`
					} `tfsdk:"key" yaml:"key,omitempty"`
				} `tfsdk:"filter" yaml:"filter,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				QueueARN *string `tfsdk:"queue_arn" yaml:"queueARN,omitempty"`
			} `tfsdk:"queue_configurations" yaml:"queueConfigurations,omitempty"`

			TopicConfigurations *[]struct {
				Events *[]string `tfsdk:"events" yaml:"events,omitempty"`

				Filter *struct {
					Key *struct {
						FilterRules *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"filter_rules" yaml:"filterRules,omitempty"`
					} `tfsdk:"key" yaml:"key,omitempty"`
				} `tfsdk:"filter" yaml:"filter,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				TopicARN *string `tfsdk:"topic_arn" yaml:"topicARN,omitempty"`
			} `tfsdk:"topic_configurations" yaml:"topicConfigurations,omitempty"`
		} `tfsdk:"notification" yaml:"notification,omitempty"`

		ObjectLockEnabledForBucket *bool `tfsdk:"object_lock_enabled_for_bucket" yaml:"objectLockEnabledForBucket,omitempty"`

		OwnershipControls *struct {
			Rules *[]struct {
				ObjectOwnership *string `tfsdk:"object_ownership" yaml:"objectOwnership,omitempty"`
			} `tfsdk:"rules" yaml:"rules,omitempty"`
		} `tfsdk:"ownership_controls" yaml:"ownershipControls,omitempty"`

		Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`

		PublicAccessBlock *struct {
			BlockPublicACLs *bool `tfsdk:"block_public_ac_ls" yaml:"blockPublicACLs,omitempty"`

			BlockPublicPolicy *bool `tfsdk:"block_public_policy" yaml:"blockPublicPolicy,omitempty"`

			IgnorePublicACLs *bool `tfsdk:"ignore_public_ac_ls" yaml:"ignorePublicACLs,omitempty"`

			RestrictPublicBuckets *bool `tfsdk:"restrict_public_buckets" yaml:"restrictPublicBuckets,omitempty"`
		} `tfsdk:"public_access_block" yaml:"publicAccessBlock,omitempty"`

		Replication *struct {
			Role *string `tfsdk:"role" yaml:"role,omitempty"`

			Rules *[]struct {
				DeleteMarkerReplication *struct {
					Status *string `tfsdk:"status" yaml:"status,omitempty"`
				} `tfsdk:"delete_marker_replication" yaml:"deleteMarkerReplication,omitempty"`

				Destination *struct {
					AccessControlTranslation *struct {
						Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`
					} `tfsdk:"access_control_translation" yaml:"accessControlTranslation,omitempty"`

					Account *string `tfsdk:"account" yaml:"account,omitempty"`

					Bucket *string `tfsdk:"bucket" yaml:"bucket,omitempty"`

					EncryptionConfiguration *struct {
						ReplicaKMSKeyID *string `tfsdk:"replica_kms_key_id" yaml:"replicaKMSKeyID,omitempty"`
					} `tfsdk:"encryption_configuration" yaml:"encryptionConfiguration,omitempty"`

					Metrics *struct {
						EventThreshold *struct {
							Minutes *int64 `tfsdk:"minutes" yaml:"minutes,omitempty"`
						} `tfsdk:"event_threshold" yaml:"eventThreshold,omitempty"`

						Status *string `tfsdk:"status" yaml:"status,omitempty"`
					} `tfsdk:"metrics" yaml:"metrics,omitempty"`

					ReplicationTime *struct {
						Status *string `tfsdk:"status" yaml:"status,omitempty"`

						Time *struct {
							Minutes *int64 `tfsdk:"minutes" yaml:"minutes,omitempty"`
						} `tfsdk:"time" yaml:"time,omitempty"`
					} `tfsdk:"replication_time" yaml:"replicationTime,omitempty"`

					StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`
				} `tfsdk:"destination" yaml:"destination,omitempty"`

				ExistingObjectReplication *struct {
					Status *string `tfsdk:"status" yaml:"status,omitempty"`
				} `tfsdk:"existing_object_replication" yaml:"existingObjectReplication,omitempty"`

				Filter *struct {
					And *struct {
						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						Tags *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"tags" yaml:"tags,omitempty"`
					} `tfsdk:"and" yaml:"and,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Tag *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tag" yaml:"tag,omitempty"`
				} `tfsdk:"filter" yaml:"filter,omitempty"`

				Id *string `tfsdk:"id" yaml:"id,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

				SourceSelectionCriteria *struct {
					ReplicaModifications *struct {
						Status *string `tfsdk:"status" yaml:"status,omitempty"`
					} `tfsdk:"replica_modifications" yaml:"replicaModifications,omitempty"`

					SseKMSEncryptedObjects *struct {
						Status *string `tfsdk:"status" yaml:"status,omitempty"`
					} `tfsdk:"sse_kms_encrypted_objects" yaml:"sseKMSEncryptedObjects,omitempty"`
				} `tfsdk:"source_selection_criteria" yaml:"sourceSelectionCriteria,omitempty"`

				Status *string `tfsdk:"status" yaml:"status,omitempty"`
			} `tfsdk:"rules" yaml:"rules,omitempty"`
		} `tfsdk:"replication" yaml:"replication,omitempty"`

		RequestPayment *struct {
			Payer *string `tfsdk:"payer" yaml:"payer,omitempty"`
		} `tfsdk:"request_payment" yaml:"requestPayment,omitempty"`

		Tagging *struct {
			TagSet *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"tag_set" yaml:"tagSet,omitempty"`
		} `tfsdk:"tagging" yaml:"tagging,omitempty"`

		Versioning *struct {
			Status *string `tfsdk:"status" yaml:"status,omitempty"`
		} `tfsdk:"versioning" yaml:"versioning,omitempty"`

		Website *struct {
			ErrorDocument *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`
			} `tfsdk:"error_document" yaml:"errorDocument,omitempty"`

			IndexDocument *struct {
				Suffix *string `tfsdk:"suffix" yaml:"suffix,omitempty"`
			} `tfsdk:"index_document" yaml:"indexDocument,omitempty"`

			RedirectAllRequestsTo *struct {
				HostName *string `tfsdk:"host_name" yaml:"hostName,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
			} `tfsdk:"redirect_all_requests_to" yaml:"redirectAllRequestsTo,omitempty"`

			RoutingRules *[]struct {
				Condition *struct {
					HttpErrorCodeReturnedEquals *string `tfsdk:"http_error_code_returned_equals" yaml:"httpErrorCodeReturnedEquals,omitempty"`

					KeyPrefixEquals *string `tfsdk:"key_prefix_equals" yaml:"keyPrefixEquals,omitempty"`
				} `tfsdk:"condition" yaml:"condition,omitempty"`

				Redirect *struct {
					HostName *string `tfsdk:"host_name" yaml:"hostName,omitempty"`

					HttpRedirectCode *string `tfsdk:"http_redirect_code" yaml:"httpRedirectCode,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

					ReplaceKeyPrefixWith *string `tfsdk:"replace_key_prefix_with" yaml:"replaceKeyPrefixWith,omitempty"`

					ReplaceKeyWith *string `tfsdk:"replace_key_with" yaml:"replaceKeyWith,omitempty"`
				} `tfsdk:"redirect" yaml:"redirect,omitempty"`
			} `tfsdk:"routing_rules" yaml:"routingRules,omitempty"`
		} `tfsdk:"website" yaml:"website,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewS3ServicesK8SAwsBucketV1Alpha1Resource() resource.Resource {
	return &S3ServicesK8SAwsBucketV1Alpha1Resource{}
}

func (r *S3ServicesK8SAwsBucketV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_s3_services_k8s_aws_bucket_v1alpha1"
}

func (r *S3ServicesK8SAwsBucketV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Bucket is the Schema for the Buckets API",
		MarkdownDescription: "Bucket is the Schema for the Buckets API",
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
				Description:         "BucketSpec defines the desired state of Bucket.  In terms of implementation, a Bucket is a resource. An Amazon S3 bucket name is globally unique, and the namespace is shared by all Amazon Web Services accounts.",
				MarkdownDescription: "BucketSpec defines the desired state of Bucket.  In terms of implementation, a Bucket is a resource. An Amazon S3 bucket name is globally unique, and the namespace is shared by all Amazon Web Services accounts.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"accelerate": {
						Description:         "Container for setting the transfer acceleration state.",
						MarkdownDescription: "Container for setting the transfer acceleration state.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"status": {
								Description:         "",
								MarkdownDescription: "",

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

					"acl": {
						Description:         "The canned ACL to apply to the bucket.",
						MarkdownDescription: "The canned ACL to apply to the bucket.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"analytics": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"filter": {
								Description:         "The filter used to describe a set of objects for analyses. A filter must have exactly one prefix, one tag, or one conjunction (AnalyticsAndOperator). If no filter is provided, all objects will be considered in any analysis.",
								MarkdownDescription: "The filter used to describe a set of objects for analyses. A filter must have exactly one prefix, one tag, or one conjunction (AnalyticsAndOperator). If no filter is provided, all objects will be considered in any analysis.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"and": {
										Description:         "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates in any combination, and an object must match all of the predicates for the filter to apply.",
										MarkdownDescription: "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates in any combination, and an object must match all of the predicates for the filter to apply.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tags": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tag": {
										Description:         "A container of a key value name pair.",
										MarkdownDescription: "A container of a key value name pair.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage_class_analysis": {
								Description:         "Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.",
								MarkdownDescription: "Specifies data related to access patterns to be collected and made available to analyze the tradeoffs between different storage classes for an Amazon S3 bucket.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"data_export": {
										Description:         "Container for data related to the storage class analysis for an Amazon S3 bucket for export.",
										MarkdownDescription: "Container for data related to the storage class analysis for an Amazon S3 bucket for export.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"destination": {
												Description:         "Where to publish the analytics results.",
												MarkdownDescription: "Where to publish the analytics results.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"s3_bucket_destination": {
														Description:         "Contains information about where to publish the analytics results.",
														MarkdownDescription: "Contains information about where to publish the analytics results.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"bucket": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"bucket_account_id": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"format": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"output_schema_version": {
												Description:         "",
												MarkdownDescription: "",

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

					"cors": {
						Description:         "Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon S3 User Guide.",
						MarkdownDescription: "Describes the cross-origin access configuration for objects in an Amazon S3 bucket. For more information, see Enabling Cross-Origin Resource Sharing (https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html) in the Amazon S3 User Guide.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cors_rules": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"allowed_headers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allowed_methods": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allowed_origins": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expose_headers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_age_seconds": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
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

					"create_bucket_configuration": {
						Description:         "The configuration information for the bucket.",
						MarkdownDescription: "The configuration information for the bucket.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"location_constraint": {
								Description:         "",
								MarkdownDescription: "",

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

					"encryption": {
						Description:         "Specifies the default server-side-encryption configuration.",
						MarkdownDescription: "Specifies the default server-side-encryption configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"rules": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"apply_server_side_encryption_by_default": {
										Description:         "Describes the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied. For more information, see PUT Bucket encryption (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTencryption.html) in the Amazon S3 API Reference.",
										MarkdownDescription: "Describes the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied. For more information, see PUT Bucket encryption (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTencryption.html) in the Amazon S3 API Reference.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"kms_master_key_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sse_algorithm": {
												Description:         "",
												MarkdownDescription: "",

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

									"bucket_key_enabled": {
										Description:         "",
										MarkdownDescription: "",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grant_full_control": {
						Description:         "Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.",
						MarkdownDescription: "Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grant_read": {
						Description:         "Allows grantee to list the objects in the bucket.",
						MarkdownDescription: "Allows grantee to list the objects in the bucket.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grant_read_acp": {
						Description:         "Allows grantee to read the bucket ACL.",
						MarkdownDescription: "Allows grantee to read the bucket ACL.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grant_write": {
						Description:         "Allows grantee to create new objects in the bucket.  For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.",
						MarkdownDescription: "Allows grantee to create new objects in the bucket.  For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grant_write_acp": {
						Description:         "Allows grantee to write the ACL for the applicable bucket.",
						MarkdownDescription: "Allows grantee to write the ACL for the applicable bucket.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"intelligent_tiering": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"filter": {
								Description:         "The Filter is used to identify objects that the S3 Intelligent-Tiering configuration applies to.",
								MarkdownDescription: "The Filter is used to identify objects that the S3 Intelligent-Tiering configuration applies to.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"and": {
										Description:         "A container for specifying S3 Intelligent-Tiering filters. The filters determine the subset of objects to which the rule applies.",
										MarkdownDescription: "A container for specifying S3 Intelligent-Tiering filters. The filters determine the subset of objects to which the rule applies.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tags": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tag": {
										Description:         "A container of a key value name pair.",
										MarkdownDescription: "A container of a key value name pair.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"status": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tierings": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"access_tier": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"days": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
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

					"inventory": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"destination": {
								Description:         "Specifies the inventory configuration for an Amazon S3 bucket.",
								MarkdownDescription: "Specifies the inventory configuration for an Amazon S3 bucket.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"s3_bucket_destination": {
										Description:         "Contains the bucket name, file format, bucket owner (optional), and prefix (optional) where inventory results are published.",
										MarkdownDescription: "Contains the bucket name, file format, bucket owner (optional), and prefix (optional) where inventory results are published.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"account_id": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bucket": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"encryption": {
												Description:         "Contains the type of server-side encryption used to encrypt the inventory results.",
												MarkdownDescription: "Contains the type of server-side encryption used to encrypt the inventory results.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"sse_kms": {
														Description:         "Specifies the use of SSE-KMS to encrypt delivered inventory reports.",
														MarkdownDescription: "Specifies the use of SSE-KMS to encrypt delivered inventory reports.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key_id": {
																Description:         "",
																MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"format": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"filter": {
								Description:         "Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria.",
								MarkdownDescription: "Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"prefix": {
										Description:         "",
										MarkdownDescription: "",

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

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"included_object_versions": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"is_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"optional_fields": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"schedule": {
								Description:         "Specifies the schedule for generating inventory results.",
								MarkdownDescription: "Specifies the schedule for generating inventory results.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"frequency": {
										Description:         "",
										MarkdownDescription: "",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"lifecycle": {
						Description:         "Container for lifecycle rules. You can add as many as 1,000 rules.",
						MarkdownDescription: "Container for lifecycle rules. You can add as many as 1,000 rules.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"rules": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"abort_incomplete_multipart_upload": {
										Description:         "Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. For more information, see Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) in the Amazon S3 User Guide.",
										MarkdownDescription: "Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload. For more information, see Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Policy (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config) in the Amazon S3 User Guide.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"days_after_initiation": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expiration": {
										Description:         "Container for the expiration for the lifecycle of the object.",
										MarkdownDescription: "Container for the expiration for the lifecycle of the object.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"date": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.DateTime64Validator(),
												},
											},

											"days": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"expired_object_delete_marker": {
												Description:         "",
												MarkdownDescription: "",

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

									"filter": {
										Description:         "The Filter is used to identify objects that a Lifecycle Rule applies to. A Filter must have exactly one of Prefix, Tag, or And specified.",
										MarkdownDescription: "The Filter is used to identify objects that a Lifecycle Rule applies to. A Filter must have exactly one of Prefix, Tag, or And specified.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"and": {
												Description:         "This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will apply to any object matching all of the predicates configured inside the And operator.",
												MarkdownDescription: "This is used in a Lifecycle Rule Filter to apply a logical AND to two or more predicates. The Lifecycle Rule will apply to any object matching all of the predicates configured inside the And operator.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"prefix": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tags": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tag": {
												Description:         "A container of a key value name pair.",
												MarkdownDescription: "A container of a key value name pair.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"noncurrent_version_expiration": {
										Description:         "Specifies when noncurrent object versions expire. Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.",
										MarkdownDescription: "Specifies when noncurrent object versions expire. Upon expiration, Amazon S3 permanently deletes the noncurrent object versions. You set this lifecycle configuration action on a bucket that has versioning enabled (or suspended) to request that Amazon S3 delete noncurrent object versions at a specific period in the object's lifetime.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"noncurrent_days": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"noncurrent_version_transitions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"noncurrent_days": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_class": {
												Description:         "",
												MarkdownDescription: "",

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

									"prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"status": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"transitions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"date": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.DateTime64Validator(),
												},
											},

											"days": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_class": {
												Description:         "",
												MarkdownDescription: "",

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

					"logging": {
						Description:         "Container for logging status information.",
						MarkdownDescription: "Container for logging status information.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"logging_enabled": {
								Description:         "Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket. For more information, see PUT Bucket logging (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the Amazon S3 API Reference.",
								MarkdownDescription: "Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a bucket. For more information, see PUT Bucket logging (https://docs.aws.amazon.com/AmazonS3/latest/API/RESTBucketPUTlogging.html) in the Amazon S3 API Reference.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"target_bucket": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_grants": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"grantee": {
												Description:         "Container for the person being granted permissions.",
												MarkdownDescription: "Container for the person being granted permissions.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"display_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"email_address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"id": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type_": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"u_ri": {
														Description:         "",
														MarkdownDescription: "",

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

											"permission": {
												Description:         "",
												MarkdownDescription: "",

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

									"target_prefix": {
										Description:         "",
										MarkdownDescription: "",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"filter": {
								Description:         "Specifies a metrics configuration filter. The metrics configuration only includes objects that meet the filter's criteria. A filter must be a prefix, an object tag, an access point ARN, or a conjunction (MetricsAndOperator). For more information, see PutBucketMetricsConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html).",
								MarkdownDescription: "Specifies a metrics configuration filter. The metrics configuration only includes objects that meet the filter's criteria. A filter must be a prefix, an object tag, an access point ARN, or a conjunction (MetricsAndOperator). For more information, see PutBucketMetricsConfiguration (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"access_point_arn": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"and": {
										Description:         "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates, and an object must match all of the predicates in order for the filter to apply.",
										MarkdownDescription: "A conjunction (logical AND) of predicates, which is used in evaluating a metrics filter. The operator must have at least two predicates, and an object must match all of the predicates in order for the filter to apply.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_point_arn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tags": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tag": {
										Description:         "A container of a key value name pair.",
										MarkdownDescription: "A container of a key value name pair.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"id": {
								Description:         "",
								MarkdownDescription: "",

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

					"name": {
						Description:         "The name of the bucket to create.",
						MarkdownDescription: "The name of the bucket to create.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"notification": {
						Description:         "A container for specifying the notification configuration of the bucket. If this element is empty, notifications are turned off for the bucket.",
						MarkdownDescription: "A container for specifying the notification configuration of the bucket. If this element is empty, notifications are turned off for the bucket.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"lambda_function_configurations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"events": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filter": {
										Description:         "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
										MarkdownDescription: "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "A container for object key name prefix and suffix filtering rules.",
												MarkdownDescription: "A container for object key name prefix and suffix filtering rules.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"filter_rules": {
														Description:         "A list of containers for the key-value pair that defines the criteria for the filter rule.",
														MarkdownDescription: "A list of containers for the key-value pair that defines the criteria for the filter rule.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

									"id": {
										Description:         "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
										MarkdownDescription: "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"lambda_function_arn": {
										Description:         "",
										MarkdownDescription: "",

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

							"queue_configurations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"events": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filter": {
										Description:         "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
										MarkdownDescription: "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "A container for object key name prefix and suffix filtering rules.",
												MarkdownDescription: "A container for object key name prefix and suffix filtering rules.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"filter_rules": {
														Description:         "A list of containers for the key-value pair that defines the criteria for the filter rule.",
														MarkdownDescription: "A list of containers for the key-value pair that defines the criteria for the filter rule.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

									"id": {
										Description:         "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
										MarkdownDescription: "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"queue_arn": {
										Description:         "",
										MarkdownDescription: "",

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

							"topic_configurations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"events": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filter": {
										Description:         "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",
										MarkdownDescription: "Specifies object key name filtering rules. For information about key name filtering, see Configuring Event Notifications (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html) in the Amazon S3 User Guide.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "A container for object key name prefix and suffix filtering rules.",
												MarkdownDescription: "A container for object key name prefix and suffix filtering rules.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"filter_rules": {
														Description:         "A list of containers for the key-value pair that defines the criteria for the filter rule.",
														MarkdownDescription: "A list of containers for the key-value pair that defines the criteria for the filter rule.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

									"id": {
										Description:         "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",
										MarkdownDescription: "An optional unique identifier for configurations in a notification configuration. If you don't provide one, Amazon S3 will assign an ID.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"topic_arn": {
										Description:         "",
										MarkdownDescription: "",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"object_lock_enabled_for_bucket": {
						Description:         "Specifies whether you want S3 Object Lock to be enabled for the new bucket.",
						MarkdownDescription: "Specifies whether you want S3 Object Lock to be enabled for the new bucket.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ownership_controls": {
						Description:         "The OwnershipControls (BucketOwnerPreferred or ObjectWriter) that you want to apply to this Amazon S3 bucket.",
						MarkdownDescription: "The OwnershipControls (BucketOwnerPreferred or ObjectWriter) that you want to apply to this Amazon S3 bucket.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"rules": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"object_ownership": {
										Description:         "The container element for object ownership for a bucket's ownership controls.  BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the bucket owner if the objects are uploaded with the bucket-owner-full-control canned ACL.  ObjectWriter - The uploading account will own the object if the object is uploaded with the bucket-owner-full-control canned ACL.",
										MarkdownDescription: "The container element for object ownership for a bucket's ownership controls.  BucketOwnerPreferred - Objects uploaded to the bucket change ownership to the bucket owner if the objects are uploaded with the bucket-owner-full-control canned ACL.  ObjectWriter - The uploading account will own the object if the object is uploaded with the bucket-owner-full-control canned ACL.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"policy": {
						Description:         "The bucket policy as a JSON document.",
						MarkdownDescription: "The bucket policy as a JSON document.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"public_access_block": {
						Description:         "The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see The Meaning of 'Public' (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the Amazon S3 User Guide.",
						MarkdownDescription: "The PublicAccessBlock configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. For more information about when Amazon S3 considers a bucket or object public, see The Meaning of 'Public' (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status) in the Amazon S3 User Guide.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"block_public_ac_ls": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"block_public_policy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ignore_public_ac_ls": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"restrict_public_buckets": {
								Description:         "",
								MarkdownDescription: "",

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

					"replication": {
						Description:         "A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.",
						MarkdownDescription: "A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication configuration is 2 MB.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"role": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rules": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"delete_marker_replication": {
										Description:         "Specifies whether Amazon S3 replicates delete markers. If you specify a Filter in your replication configuration, you must also include a DeleteMarkerReplication element. If your Filter includes a Tag element, the DeleteMarkerReplication Status must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config).  For more information about delete marker replication, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html).  If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see Backward Compatibility (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations).",
										MarkdownDescription: "Specifies whether Amazon S3 replicates delete markers. If you specify a Filter in your replication configuration, you must also include a DeleteMarkerReplication element. If your Filter includes a Tag element, the DeleteMarkerReplication Status must be set to Disabled, because Amazon S3 does not support replicating delete markers for tag-based rules. For an example configuration, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-config-min-rule-config).  For more information about delete marker replication, see Basic Rule Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/delete-marker-replication.html).  If you are using an earlier version of the replication configuration, Amazon S3 handles replication of delete markers differently. For more information, see Backward Compatibility (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations).",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"status": {
												Description:         "",
												MarkdownDescription: "",

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

									"destination": {
										Description:         "Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).",
										MarkdownDescription: "Specifies information about where to publish analysis or configuration results for an Amazon S3 bucket and S3 Replication Time Control (S3 RTC).",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_control_translation": {
												Description:         "A container for information about access control for replicas.",
												MarkdownDescription: "A container for information about access control for replicas.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"owner": {
														Description:         "",
														MarkdownDescription: "",

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

											"account": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bucket": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"encryption_configuration": {
												Description:         "Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.",
												MarkdownDescription: "Specifies encryption-related information for an Amazon S3 bucket that is a destination for replicated objects.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"replica_kms_key_id": {
														Description:         "",
														MarkdownDescription: "",

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

											"metrics": {
												Description:         "A container specifying replication metrics-related settings enabling replication metrics and events.",
												MarkdownDescription: "A container specifying replication metrics-related settings enabling replication metrics and events.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"event_threshold": {
														Description:         "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",
														MarkdownDescription: "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"minutes": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"status": {
														Description:         "",
														MarkdownDescription: "",

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

											"replication_time": {
												Description:         "A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a Metrics block.",
												MarkdownDescription: "A container specifying S3 Replication Time Control (S3 RTC) related information, including whether S3 RTC is enabled and the time when all objects and operations on objects must be replicated. Must be specified together with a Metrics block.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"status": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"time": {
														Description:         "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",
														MarkdownDescription: "A container specifying the time value for S3 Replication Time Control (S3 RTC) and replication metrics EventThreshold.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"minutes": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
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

											"storage_class": {
												Description:         "",
												MarkdownDescription: "",

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

									"existing_object_replication": {
										Description:         "Optional configuration to replicate existing source bucket objects. For more information, see Replicating Existing Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-what-is-isnot-replicated.html#existing-object-replication) in the Amazon S3 User Guide.",
										MarkdownDescription: "Optional configuration to replicate existing source bucket objects. For more information, see Replicating Existing Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-what-is-isnot-replicated.html#existing-object-replication) in the Amazon S3 User Guide.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"status": {
												Description:         "",
												MarkdownDescription: "",

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

									"filter": {
										Description:         "A filter that identifies the subset of objects to which the replication rule applies. A Filter must specify exactly one Prefix, Tag, or an And child element.",
										MarkdownDescription: "A filter that identifies the subset of objects to which the replication rule applies. A Filter must specify exactly one Prefix, Tag, or an And child element.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"and": {
												Description:         "A container for specifying rule filters. The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter.  For example:     * If you specify both a Prefix and a Tag filter, wrap these filters in    an And tag.     * If you specify a filter based on multiple tags, wrap the Tag elements    in an And tag.",
												MarkdownDescription: "A container for specifying rule filters. The filters determine the subset of objects to which the rule applies. This element is required only if you specify more than one filter.  For example:     * If you specify both a Prefix and a Tag filter, wrap these filters in    an And tag.     * If you specify a filter based on multiple tags, wrap the Tag elements    in an And tag.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"prefix": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tags": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tag": {
												Description:         "A container of a key value name pair.",
												MarkdownDescription: "A container of a key value name pair.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_selection_criteria": {
										Description:         "A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects. Currently, Amazon S3 supports only the filter that you can specify for objects created with server-side encryption using a customer managed key stored in Amazon Web Services Key Management Service (SSE-KMS).",
										MarkdownDescription: "A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects. Currently, Amazon S3 supports only the filter that you can specify for objects created with server-side encryption using a customer managed key stored in Amazon Web Services Key Management Service (SSE-KMS).",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"replica_modifications": {
												Description:         "A filter that you can specify for selection for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when Filter is specified), you can specify this element and set the status to Enabled to replicate modifications on replicas.  If you don't specify the Filter element, Amazon S3 assumes that the replication configuration is the earlier version, V1. In the earlier version, this element is not allowed.",
												MarkdownDescription: "A filter that you can specify for selection for modifications on replicas. Amazon S3 doesn't replicate replica modifications by default. In the latest version of replication configuration (when Filter is specified), you can specify this element and set the status to Enabled to replicate modifications on replicas.  If you don't specify the Filter element, Amazon S3 assumes that the replication configuration is the earlier version, V1. In the earlier version, this element is not allowed.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"status": {
														Description:         "",
														MarkdownDescription: "",

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

											"sse_kms_encrypted_objects": {
												Description:         "A container for filter information for the selection of S3 objects encrypted with Amazon Web Services KMS.",
												MarkdownDescription: "A container for filter information for the selection of S3 objects encrypted with Amazon Web Services KMS.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"status": {
														Description:         "",
														MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"status": {
										Description:         "",
										MarkdownDescription: "",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"request_payment": {
						Description:         "Container for Payer.",
						MarkdownDescription: "Container for Payer.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"payer": {
								Description:         "",
								MarkdownDescription: "",

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

					"tagging": {
						Description:         "Container for the TagSet and Tag elements.",
						MarkdownDescription: "Container for the TagSet and Tag elements.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"tag_set": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"versioning": {
						Description:         "Container for setting the versioning state.",
						MarkdownDescription: "Container for setting the versioning state.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"status": {
								Description:         "",
								MarkdownDescription: "",

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

					"website": {
						Description:         "Container for the request.",
						MarkdownDescription: "Container for the request.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"error_document": {
								Description:         "The error information.",
								MarkdownDescription: "The error information.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "",
										MarkdownDescription: "",

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

							"index_document": {
								Description:         "Container for the Suffix element.",
								MarkdownDescription: "Container for the Suffix element.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"suffix": {
										Description:         "",
										MarkdownDescription: "",

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

							"redirect_all_requests_to": {
								Description:         "Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.",
								MarkdownDescription: "Specifies the redirect behavior of all requests to a website endpoint of an Amazon S3 bucket.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol": {
										Description:         "",
										MarkdownDescription: "",

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

							"routing_rules": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"condition": {
										Description:         "A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the /docs folder, redirect to the /documents folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.",
										MarkdownDescription: "A container for describing a condition that must be met for the specified redirect to apply. For example, 1. If request is for pages in the /docs folder, redirect to the /documents folder. 2. If request results in HTTP error 4xx, redirect request to another host where you might process the error.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"http_error_code_returned_equals": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key_prefix_equals": {
												Description:         "",
												MarkdownDescription: "",

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

									"redirect": {
										Description:         "Specifies how requests are redirected. In the event of an error, you can specify a different error code to return.",
										MarkdownDescription: "Specifies how requests are redirected. In the event of an error, you can specify a different error code to return.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_redirect_code": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"protocol": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"replace_key_prefix_with": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"replace_key_with": {
												Description:         "",
												MarkdownDescription: "",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *S3ServicesK8SAwsBucketV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_s3_services_k8s_aws_bucket_v1alpha1")

	var state S3ServicesK8SAwsBucketV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel S3ServicesK8SAwsBucketV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("s3.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Bucket")

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

func (r *S3ServicesK8SAwsBucketV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_s3_services_k8s_aws_bucket_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *S3ServicesK8SAwsBucketV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_s3_services_k8s_aws_bucket_v1alpha1")

	var state S3ServicesK8SAwsBucketV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel S3ServicesK8SAwsBucketV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("s3.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Bucket")

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

func (r *S3ServicesK8SAwsBucketV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_s3_services_k8s_aws_bucket_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
