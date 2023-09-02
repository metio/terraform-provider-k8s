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
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CouchbaseComCouchbaseBucketV2Manifest{}
)

func NewCouchbaseComCouchbaseBucketV2Manifest() datasource.DataSource {
	return &CouchbaseComCouchbaseBucketV2Manifest{}
}

type CouchbaseComCouchbaseBucketV2Manifest struct{}

type CouchbaseComCouchbaseBucketV2ManifestData struct {
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
		CompressionMode    *string `tfsdk:"compression_mode" json:"compressionMode,omitempty"`
		ConflictResolution *string `tfsdk:"conflict_resolution" json:"conflictResolution,omitempty"`
		EnableFlush        *bool   `tfsdk:"enable_flush" json:"enableFlush,omitempty"`
		EnableIndexReplica *bool   `tfsdk:"enable_index_replica" json:"enableIndexReplica,omitempty"`
		EvictionPolicy     *string `tfsdk:"eviction_policy" json:"evictionPolicy,omitempty"`
		IoPriority         *string `tfsdk:"io_priority" json:"ioPriority,omitempty"`
		MaxTTL             *string `tfsdk:"max_ttl" json:"maxTTL,omitempty"`
		MemoryQuota        *string `tfsdk:"memory_quota" json:"memoryQuota,omitempty"`
		MinimumDurability  *string `tfsdk:"minimum_durability" json:"minimumDurability,omitempty"`
		Name               *string `tfsdk:"name" json:"name,omitempty"`
		Replicas           *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		Scopes             *struct {
			Managed   *bool `tfsdk:"managed" json:"managed,omitempty"`
			Resources *[]struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"scopes" json:"scopes,omitempty"`
		StorageBackend *string `tfsdk:"storage_backend" json:"storageBackend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseBucketV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_bucket_v2_manifest"
}

func (r *CouchbaseComCouchbaseBucketV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The CouchbaseBucket resource defines a set of documents in Couchbase server. A Couchbase client connects to and operates on a bucket, which provides independent management of a set documents and a security boundary for role based access control. A CouchbaseBucket provides replication and persistence for documents contained by it.",
		MarkdownDescription: "The CouchbaseBucket resource defines a set of documents in Couchbase server. A Couchbase client connects to and operates on a bucket, which provides independent management of a set documents and a security boundary for role based access control. A CouchbaseBucket provides replication and persistence for documents contained by it.",
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
				Description:         "CouchbaseBucketSpec is the specification for a Couchbase bucket resource, and allows the bucket to be customized.",
				MarkdownDescription: "CouchbaseBucketSpec is the specification for a Couchbase bucket resource, and allows the bucket to be customized.",
				Attributes: map[string]schema.Attribute{
					"compression_mode": schema.StringAttribute{
						Description:         "CompressionMode defines how Couchbase server handles document compression.  When off, documents are stored in memory, and transferred to the client uncompressed. When passive, documents are stored compressed in memory, and transferred to the client compressed when requested.  When active, documents are stored compresses in memory and when transferred to the client.  This field must be 'off', 'passive' or 'active', defaulting to 'passive'.  Be aware 'off' in YAML 1.2 is a boolean, so must be quoted as a string in configuration files.",
						MarkdownDescription: "CompressionMode defines how Couchbase server handles document compression.  When off, documents are stored in memory, and transferred to the client uncompressed. When passive, documents are stored compressed in memory, and transferred to the client compressed when requested.  When active, documents are stored compresses in memory and when transferred to the client.  This field must be 'off', 'passive' or 'active', defaulting to 'passive'.  Be aware 'off' in YAML 1.2 is a boolean, so must be quoted as a string in configuration files.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("off", "passive", "active"),
						},
					},

					"conflict_resolution": schema.StringAttribute{
						Description:         "ConflictResolution defines how XDCR handles concurrent write conflicts.  Sequence number based resolution selects the document with the highest sequence number as the most recent. Timestamp based resolution selects the document that was written to most recently as the most recent.  This field must be 'seqno' (sequence based), or 'lww' (timestamp based), defaulting to 'seqno'.",
						MarkdownDescription: "ConflictResolution defines how XDCR handles concurrent write conflicts.  Sequence number based resolution selects the document with the highest sequence number as the most recent. Timestamp based resolution selects the document that was written to most recently as the most recent.  This field must be 'seqno' (sequence based), or 'lww' (timestamp based), defaulting to 'seqno'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("seqno", "lww"),
						},
					},

					"enable_flush": schema.BoolAttribute{
						Description:         "EnableFlush defines whether a client can delete all documents in a bucket. This field defaults to false.",
						MarkdownDescription: "EnableFlush defines whether a client can delete all documents in a bucket. This field defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_index_replica": schema.BoolAttribute{
						Description:         "EnableIndexReplica defines whether indexes for this bucket are replicated. This field defaults to false.",
						MarkdownDescription: "EnableIndexReplica defines whether indexes for this bucket are replicated. This field defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"eviction_policy": schema.StringAttribute{
						Description:         "EvictionPolicy controls how Couchbase handles memory exhaustion.  Value only eviction flushes documents to disk but maintains document metadata in memory in order to improve query performance.  Full eviction removes all data from memory after the document is flushed to disk.  This field must be 'valueOnly' or 'fullEviction', defaulting to 'valueOnly'.",
						MarkdownDescription: "EvictionPolicy controls how Couchbase handles memory exhaustion.  Value only eviction flushes documents to disk but maintains document metadata in memory in order to improve query performance.  Full eviction removes all data from memory after the document is flushed to disk.  This field must be 'valueOnly' or 'fullEviction', defaulting to 'valueOnly'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("valueOnly", "fullEviction"),
						},
					},

					"io_priority": schema.StringAttribute{
						Description:         "IOPriority controls how many threads a bucket has, per pod, to process reads and writes. This field must be 'low' or 'high', defaulting to 'low'.  Modification of this field will cause a temporary service disruption as threads are restarted.",
						MarkdownDescription: "IOPriority controls how many threads a bucket has, per pod, to process reads and writes. This field must be 'low' or 'high', defaulting to 'low'.  Modification of this field will cause a temporary service disruption as threads are restarted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("low", "high"),
						},
					},

					"max_ttl": schema.StringAttribute{
						Description:         "MaxTTL defines how long a document is permitted to exist for, without modification, until it is automatically deleted.  This is a default and maximum time-to-live and may be set to a lower value by the client.  If the client specifies a higher value, then it is truncated to the maximum durability.  Documents are removed by Couchbase, after they have expired, when either accessed, the expiry pager is run, or the bucket is compacted.  When set to 0, then documents are not expired by default.  This field must be a duration in the range 0-2147483648s, defaulting to 0.  More info: https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "MaxTTL defines how long a document is permitted to exist for, without modification, until it is automatically deleted.  This is a default and maximum time-to-live and may be set to a lower value by the client.  If the client specifies a higher value, then it is truncated to the maximum durability.  Documents are removed by Couchbase, after they have expired, when either accessed, the expiry pager is run, or the bucket is compacted.  When set to 0, then documents are not expired by default.  This field must be a duration in the range 0-2147483648s, defaulting to 0.  More info: https://golang.org/pkg/time/#ParseDuration",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"memory_quota": schema.StringAttribute{
						Description:         "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded, documents will be evicted from memory to disk as defined by the eviction policy.  The memory quota is defined per Couchbase pod running the data service.  This field defaults to, and must be greater than or equal to 100Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						MarkdownDescription: "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded, documents will be evicted from memory to disk as defined by the eviction policy.  The memory quota is defined per Couchbase pod running the data service.  This field defaults to, and must be greater than or equal to 100Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
						},
					},

					"minimum_durability": schema.StringAttribute{
						Description:         "MiniumumDurability defines how durable a document write is by default, and can be made more durable by the client.  This feature enables ACID transactions. When none, Couchbase server will respond when the document is in memory, it will become eventually consistent across the cluster.  When majority, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster.  When majorityAndPersistActive, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster and the document has been persisted to disk on the document master pod.  When persistToMajority, Couchbase server will respond when the document is replicated and persisted to disk on at least half of the pods running the data service in the cluster.  This field must be either 'none', 'majority', 'majorityAndPersistActive' or 'persistToMajority', defaulting to 'none'.",
						MarkdownDescription: "MiniumumDurability defines how durable a document write is by default, and can be made more durable by the client.  This feature enables ACID transactions. When none, Couchbase server will respond when the document is in memory, it will become eventually consistent across the cluster.  When majority, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster.  When majorityAndPersistActive, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster and the document has been persisted to disk on the document master pod.  When persistToMajority, Couchbase server will respond when the document is replicated and persisted to disk on at least half of the pods running the data service in the cluster.  This field must be either 'none', 'majority', 'majorityAndPersistActive' or 'persistToMajority', defaulting to 'none'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "majority", "majorityAndPersistActive", "persistToMajority"),
						},
					},

					"name": schema.StringAttribute{
						Description:         "Name is the name of the bucket within Couchbase server.  By default the Operator will use the 'metadata.name' field to define the bucket name.  The 'metadata.name' field only supports a subset of the supported character set.  When specified, this field overrides 'metadata.name'.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "Name is the name of the bucket within Couchbase server.  By default the Operator will use the 'metadata.name' field to define the bucket name.  The 'metadata.name' field only supports a subset of the supported character set.  When specified, this field overrides 'metadata.name'.  Legal bucket names have a maximum length of 100 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(100),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_%\.]{1,100}$`), ""),
						},
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas defines how many copies of documents Couchbase server maintains.  This directly affects how fault tolerant a Couchbase cluster is.  With a single replica, the cluster can tolerate one data pod going down and still service requests without data loss.  The number of replicas also affect memory use.  With a single replica, the effective memory quota for documents is halved, with two replicas it is one third.  The number of replicas must be between 0 and 3, defaulting to 1.",
						MarkdownDescription: "Replicas defines how many copies of documents Couchbase server maintains.  This directly affects how fault tolerant a Couchbase cluster is.  With a single replica, the cluster can tolerate one data pod going down and still service requests without data loss.  The number of replicas also affect memory use.  With a single replica, the effective memory quota for documents is halved, with two replicas it is one third.  The number of replicas must be between 0 and 3, defaulting to 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(3),
						},
					},

					"scopes": schema.SingleNestedAttribute{
						Description:         "Scopes defines whether the Operator manages scopes for the bucket or not, and the set of scopes defined for the bucket.",
						MarkdownDescription: "Scopes defines whether the Operator manages scopes for the bucket or not, and the set of scopes defined for the bucket.",
						Attributes: map[string]schema.Attribute{
							"managed": schema.BoolAttribute{
								Description:         "Managed defines whether scopes are managed for this bucket. This field is 'false' by default, and the Operator will take no actions that will affect scopes and collections in this bucket.  The default scope and collection will be present.  When set to 'true', the Operator will manage user defined scopes, and optionally, their collections as defined by the 'CouchbaseScope', 'CouchbaseScopeGroup', 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource documentation.  If this field is set to 'false' while the  already managed, then the Operator will leave whatever configuration is already present.",
								MarkdownDescription: "Managed defines whether scopes are managed for this bucket. This field is 'false' by default, and the Operator will take no actions that will affect scopes and collections in this bucket.  The default scope and collection will be present.  When set to 'true', the Operator will manage user defined scopes, and optionally, their collections as defined by the 'CouchbaseScope', 'CouchbaseScopeGroup', 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource documentation.  If this field is set to 'false' while the  already managed, then the Operator will leave whatever configuration is already present.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.ListNestedAttribute{
								Description:         "Resources is an explicit list of named resources that will be considered for inclusion in this bucket.  If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
								MarkdownDescription: "Resources is an explicit list of named resources that will be considered for inclusion in this bucket.  If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind indicates the kind of resource that is being referenced.  A scope can only reference 'CouchbaseScope' and 'CouchbaseScopeGroup' resource kinds.  This field defaults to 'CouchbaseScope' if not specified.",
											MarkdownDescription: "Kind indicates the kind of resource that is being referenced.  A scope can only reference 'CouchbaseScope' and 'CouchbaseScopeGroup' resource kinds.  This field defaults to 'CouchbaseScope' if not specified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("CouchbaseScope", "CouchbaseScopeGroup"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the Kubernetes resource name that is being referenced. Legal scope names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
											MarkdownDescription: "Name is the name of the Kubernetes resource name that is being referenced. Legal scope names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(251),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9\-][a-zA-Z0-9\-%_]{0,250}$`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector allows resources to be implicitly considered for inclusion in this bucket.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta",
								MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in this bucket.  More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#labelselector-v1-meta",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_backend": schema.StringAttribute{
						Description:         "StorageBackend to be assigned to and used by the bucket. Only valid for Couchbase Server 7.0.0 onward. Two different backend storage mechanisms can be used - 'couchstore' or 'magma', defaulting to 'couchstore'. This cannot be edited after bucket creation. Note: 'magma' is only valid for Couchbase Server 7.1.0 onward.",
						MarkdownDescription: "StorageBackend to be assigned to and used by the bucket. Only valid for Couchbase Server 7.0.0 onward. Two different backend storage mechanisms can be used - 'couchstore' or 'magma', defaulting to 'couchstore'. This cannot be edited after bucket creation. Note: 'magma' is only valid for Couchbase Server 7.1.0 onward.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("couchstore", "magma"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CouchbaseComCouchbaseBucketV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_bucket_v2_manifest")

	var model CouchbaseComCouchbaseBucketV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseBucket")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
