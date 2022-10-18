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

type CouchbaseComCouchbaseBucketV2Resource struct{}

var (
	_ resource.Resource = (*CouchbaseComCouchbaseBucketV2Resource)(nil)
)

type CouchbaseComCouchbaseBucketV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CouchbaseComCouchbaseBucketV2GoModel struct {
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
		CompressionMode *string `tfsdk:"compression_mode" yaml:"compressionMode,omitempty"`

		ConflictResolution *string `tfsdk:"conflict_resolution" yaml:"conflictResolution,omitempty"`

		EnableFlush *bool `tfsdk:"enable_flush" yaml:"enableFlush,omitempty"`

		EnableIndexReplica *bool `tfsdk:"enable_index_replica" yaml:"enableIndexReplica,omitempty"`

		EvictionPolicy *string `tfsdk:"eviction_policy" yaml:"evictionPolicy,omitempty"`

		IoPriority *string `tfsdk:"io_priority" yaml:"ioPriority,omitempty"`

		MaxTTL *string `tfsdk:"max_ttl" yaml:"maxTTL,omitempty"`

		MemoryQuota utilities.IntOrString `tfsdk:"memory_quota" yaml:"memoryQuota,omitempty"`

		MinimumDurability *string `tfsdk:"minimum_durability" yaml:"minimumDurability,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		Scopes *struct {
			Managed *bool `tfsdk:"managed" yaml:"managed,omitempty"`

			Resources *[]struct {
				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Selector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"selector" yaml:"selector,omitempty"`
		} `tfsdk:"scopes" yaml:"scopes,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCouchbaseComCouchbaseBucketV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseBucketV2Resource{}
}

func (r *CouchbaseComCouchbaseBucketV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_couchbase_com_couchbase_bucket_v2"
}

func (r *CouchbaseComCouchbaseBucketV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "The CouchbaseBucket resource defines a set of documents in Couchbase server. A Couchbase client connects to and operates on a bucket, which provides independent management of a set documents and a security boundary for role based access control. A CouchbaseBucket provides replication and persistence for documents contained by it.",
		MarkdownDescription: "The CouchbaseBucket resource defines a set of documents in Couchbase server. A Couchbase client connects to and operates on a bucket, which provides independent management of a set documents and a security boundary for role based access control. A CouchbaseBucket provides replication and persistence for documents contained by it.",
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
				Description:         "CouchbaseBucketSpec is the specification for a Couchbase bucket resource, and allows the bucket to be customized.",
				MarkdownDescription: "CouchbaseBucketSpec is the specification for a Couchbase bucket resource, and allows the bucket to be customized.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"compression_mode": {
						Description:         "CompressionMode defines how Couchbase server handles document compression.  When off, documents are stored in memory, and transferred to the client uncompressed. When passive, documents are stored compressed in memory, and transferred to the client compressed when requested.  When active, documents are stored compresses in memory and when transferred to the client.  This field must be 'off', 'passive' or 'active', defaulting to 'passive'.  Be aware 'off' in YAML 1.2 is a boolean, so must be quoted as a string in configuration files.",
						MarkdownDescription: "CompressionMode defines how Couchbase server handles document compression.  When off, documents are stored in memory, and transferred to the client uncompressed. When passive, documents are stored compressed in memory, and transferred to the client compressed when requested.  When active, documents are stored compresses in memory and when transferred to the client.  This field must be 'off', 'passive' or 'active', defaulting to 'passive'.  Be aware 'off' in YAML 1.2 is a boolean, so must be quoted as a string in configuration files.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("off", "passive", "active"),
						},
					},

					"conflict_resolution": {
						Description:         "ConflictResolution defines how XDCR handles concurrent write conflicts.  Sequence number based resolution selects the document with the highest sequence number as the most recent. Timestamp based resolution selects the document that was written to most recently as the most recent.  This field must be 'seqno' (sequence based), or 'lww' (timestamp based), defaulting to 'seqno'.",
						MarkdownDescription: "ConflictResolution defines how XDCR handles concurrent write conflicts.  Sequence number based resolution selects the document with the highest sequence number as the most recent. Timestamp based resolution selects the document that was written to most recently as the most recent.  This field must be 'seqno' (sequence based), or 'lww' (timestamp based), defaulting to 'seqno'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("seqno", "lww"),
						},
					},

					"enable_flush": {
						Description:         "EnableFlush defines whether a client can delete all documents in a bucket. This field defaults to false.",
						MarkdownDescription: "EnableFlush defines whether a client can delete all documents in a bucket. This field defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_index_replica": {
						Description:         "EnableIndexReplica defines whether indexes for this bucket are replicated. This field defaults to false.",
						MarkdownDescription: "EnableIndexReplica defines whether indexes for this bucket are replicated. This field defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"eviction_policy": {
						Description:         "EvictionPolicy controls how Couchbase handles memory exhaustion.  Value only eviction flushes documents to disk but maintains document metadata in memory in order to improve query performance.  Full eviction removes all data from memory after the document is flushed to disk.  This field must be 'valueOnly' or 'fullEviction', defaulting to 'valueOnly'.",
						MarkdownDescription: "EvictionPolicy controls how Couchbase handles memory exhaustion.  Value only eviction flushes documents to disk but maintains document metadata in memory in order to improve query performance.  Full eviction removes all data from memory after the document is flushed to disk.  This field must be 'valueOnly' or 'fullEviction', defaulting to 'valueOnly'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("valueOnly", "fullEviction"),
						},
					},

					"io_priority": {
						Description:         "IOPriority controls how many threads a bucket has, per pod, to process reads and writes. This field must be 'low' or 'high', defaulting to 'low'.  Modification of this field will cause a temporary service disruption as threads are restarted.",
						MarkdownDescription: "IOPriority controls how many threads a bucket has, per pod, to process reads and writes. This field must be 'low' or 'high', defaulting to 'low'.  Modification of this field will cause a temporary service disruption as threads are restarted.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("low", "high"),
						},
					},

					"max_ttl": {
						Description:         "MaxTTL defines how long a document is permitted to exist for, without modification, until it is automatically deleted.  This is a default and maximum time-to-live and may be set to a lower value by the client.  If the client specifies a higher value, then it is truncated to the maximum durability.  Documents are removed by Couchbase, after they have expired, when either accessed, the expiry pager is run, or the bucket is compacted.  When set to 0, then documents are not expired by default.  This field must be a duration in the range 0-2147483648s, defaulting to 0.  More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "MaxTTL defines how long a document is permitted to exist for, without modification, until it is automatically deleted.  This is a default and maximum time-to-live and may be set to a lower value by the client.  If the client specifies a higher value, then it is truncated to the maximum durability.  Documents are removed by Couchbase, after they have expired, when either accessed, the expiry pager is run, or the bucket is compacted.  When set to 0, then documents are not expired by default.  This field must be a duration in the range 0-2147483648s, defaulting to 0.  More info: https://golang.org/pkg/time/#ParseDuration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"memory_quota": {
						Description:         "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded, documents will be evicted from memory to disk as defined by the eviction policy.  The memory quota is defined per Couchbase pod running the data service.  This field defaults to, and must be greater than or equal to 100Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						MarkdownDescription: "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded, documents will be evicted from memory to disk as defined by the eviction policy.  The memory quota is defined per Couchbase pod running the data service.  This field defaults to, and must be greater than or equal to 100Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",

						Type: utilities.IntOrStringType{},

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							validators.RegexValidator(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`)),
						},
					},

					"minimum_durability": {
						Description:         "MiniumumDurability defines how durable a document write is by default, and can be made more durable by the client.  This feature enables ACID transactions. When none, Couchbase server will respond when the document is in memory, it will become eventually consistent across the cluster.  When majority, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster.  When majorityAndPersistActive, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster and the document has been persisted to disk on the document master pod.  When persistToMajority, Couchbase server will respond when the document is replicated and persisted to disk on at least half of the pods running the data service in the cluster.  This field must be either 'none', 'majority', 'majorityAndPersistActive' or 'persistToMajority', defaulting to 'none'.",
						MarkdownDescription: "MiniumumDurability defines how durable a document write is by default, and can be made more durable by the client.  This feature enables ACID transactions. When none, Couchbase server will respond when the document is in memory, it will become eventually consistent across the cluster.  When majority, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster.  When majorityAndPersistActive, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster and the document has been persisted to disk on the document master pod.  When persistToMajority, Couchbase server will respond when the document is replicated and persisted to disk on at least half of the pods running the data service in the cluster.  This field must be either 'none', 'majority', 'majorityAndPersistActive' or 'persistToMajority', defaulting to 'none'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("none", "majority", "majorityAndPersistActive", "persistToMajority"),
						},
					},

					"name": {
						Description:         "Name is the name of the bucket within Couchbase server.  By default the Operator will use the 'metadata.name' field to define the bucket name.  The 'metadata.name' field only supports a subset of the supported character set.  When specified, this field overrides 'metadata.name'.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "Name is the name of the bucket within Couchbase server.  By default the Operator will use the 'metadata.name' field to define the bucket name.  The 'metadata.name' field only supports a subset of the supported character set.  When specified, this field overrides 'metadata.name'.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtMost(100),

							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_%\.]{1,100}$`), ""),
						},
					},

					"replicas": {
						Description:         "Replicas defines how many copies of documents Couchbase server maintains.  This directly affects how fault tolerant a Couchbase cluster is.  With a single replica, the cluster can tolerate one data pod going down and still service requests without data loss.  The number of replicas also affect memory use.  With a single replica, the effective memory quota for documents is halved, with two replicas it is one third.  The number of replicas must be between 0 and 3, defaulting to 1.",
						MarkdownDescription: "Replicas defines how many copies of documents Couchbase server maintains.  This directly affects how fault tolerant a Couchbase cluster is.  With a single replica, the cluster can tolerate one data pod going down and still service requests without data loss.  The number of replicas also affect memory use.  With a single replica, the effective memory quota for documents is halved, with two replicas it is one third.  The number of replicas must be between 0 and 3, defaulting to 1.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(0),

							int64validator.AtMost(3),
						},
					},

					"scopes": {
						Description:         "Scopes defines whether the Operator manages scopes for the bucket or not, and the set of scopes defined for the bucket.",
						MarkdownDescription: "Scopes defines whether the Operator manages scopes for the bucket or not, and the set of scopes defined for the bucket.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"managed": {
								Description:         "Managed defines whether scopes are managed for this bucket. This field is 'false' by default, and the Operator will take no actions that will affect scopes and collections in this bucket.  The default scope and collection will be present.  When set to 'true', the Operator will manage user defined scopes, and optionally, their collections as defined by the 'CouchbaseScope', 'CouchbaseScopeGroup', 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource documentation.  If this field is set to 'false' while the  already managed, then the Operator will leave whatever configuration is already present.",
								MarkdownDescription: "Managed defines whether scopes are managed for this bucket. This field is 'false' by default, and the Operator will take no actions that will affect scopes and collections in this bucket.  The default scope and collection will be present.  When set to 'true', the Operator will manage user defined scopes, and optionally, their collections as defined by the 'CouchbaseScope', 'CouchbaseScopeGroup', 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource documentation.  If this field is set to 'false' while the  already managed, then the Operator will leave whatever configuration is already present.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources is an explicit list of named resources that will be considered for inclusion in this bucket.  If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
								MarkdownDescription: "Resources is an explicit list of named resources that will be considered for inclusion in this bucket.  If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"kind": {
										Description:         "Kind indicates the kind of resource that is being referenced.  A scope can only reference 'CouchbaseScope' and 'CouchbaseScopeGroup' resource kinds.  This field defaults to 'CouchbaseScope' if not specified.",
										MarkdownDescription: "Kind indicates the kind of resource that is being referenced.  A scope can only reference 'CouchbaseScope' and 'CouchbaseScopeGroup' resource kinds.  This field defaults to 'CouchbaseScope' if not specified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("CouchbaseScope", "CouchbaseScopeGroup"),
										},
									},

									"name": {
										Description:         "Name is the name of the Kubernetes resource name that is being referenced. Legal scope names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
										MarkdownDescription: "Name is the name of the Kubernetes resource name that is being referenced. Legal scope names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),

											stringvalidator.LengthAtMost(251),

											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,250}$`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": {
								Description:         "Selector allows resources to be implicitly considered for inclusion in this bucket.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta",
								MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in this bucket.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

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

func (r *CouchbaseComCouchbaseBucketV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_bucket_v2")

	var state CouchbaseComCouchbaseBucketV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseBucketV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseBucket")

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

func (r *CouchbaseComCouchbaseBucketV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_bucket_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CouchbaseComCouchbaseBucketV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_bucket_v2")

	var state CouchbaseComCouchbaseBucketV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CouchbaseComCouchbaseBucketV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("couchbase.com/v2")
	goModel.Kind = utilities.Ptr("CouchbaseBucket")

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

func (r *CouchbaseComCouchbaseBucketV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_bucket_v2")
	// NO-OP: Terraform removes the state automatically for us
}
