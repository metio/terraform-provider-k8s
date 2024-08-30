/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

import (
	"context"
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
	_ datasource.DataSource = &CouchbaseComCouchbaseEphemeralBucketV2Manifest{}
)

func NewCouchbaseComCouchbaseEphemeralBucketV2Manifest() datasource.DataSource {
	return &CouchbaseComCouchbaseEphemeralBucketV2Manifest{}
}

type CouchbaseComCouchbaseEphemeralBucketV2Manifest struct{}

type CouchbaseComCouchbaseEphemeralBucketV2ManifestData struct {
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
		EvictionPolicy     *string `tfsdk:"eviction_policy" json:"evictionPolicy,omitempty"`
		IoPriority         *string `tfsdk:"io_priority" json:"ioPriority,omitempty"`
		MaxTTL             *string `tfsdk:"max_ttl" json:"maxTTL,omitempty"`
		MemoryQuota        *string `tfsdk:"memory_quota" json:"memoryQuota,omitempty"`
		MinimumDurability  *string `tfsdk:"minimum_durability" json:"minimumDurability,omitempty"`
		Name               *string `tfsdk:"name" json:"name,omitempty"`
		Rank               *int64  `tfsdk:"rank" json:"rank,omitempty"`
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
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_ephemeral_bucket_v2_manifest"
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The CouchbaseEphemeralBucket resource defines a set of documents in Couchbase server.A Couchbase client connects to and operates on a bucket, which provides independentmanagement of a set documents and a security boundary for role based access control.A CouchbaseEphemeralBucket provides in-memory only storage and replication for documentscontained by it.",
		MarkdownDescription: "The CouchbaseEphemeralBucket resource defines a set of documents in Couchbase server.A Couchbase client connects to and operates on a bucket, which provides independentmanagement of a set documents and a security boundary for role based access control.A CouchbaseEphemeralBucket provides in-memory only storage and replication for documentscontained by it.",
		Attributes: map[string]schema.Attribute{
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
				Description:         "CouchbaseEphemeralBucketSpec is the specification for an ephemeral Couchbase bucketresource, and allows the bucket to be customized.",
				MarkdownDescription: "CouchbaseEphemeralBucketSpec is the specification for an ephemeral Couchbase bucketresource, and allows the bucket to be customized.",
				Attributes: map[string]schema.Attribute{
					"compression_mode": schema.StringAttribute{
						Description:         "CompressionMode defines how Couchbase server handles document compression.  Whenoff, documents are stored in memory, and transferred to the client uncompressed.When passive, documents are stored compressed in memory, and transferred to theclient compressed when requested.  When active, documents are stored compressesin memory and when transferred to the client.  This field must be 'off', 'passive'or 'active', defaulting to 'passive'.  Be aware 'off' in YAML 1.2 is a boolean, somust be quoted as a string in configuration files.",
						MarkdownDescription: "CompressionMode defines how Couchbase server handles document compression.  Whenoff, documents are stored in memory, and transferred to the client uncompressed.When passive, documents are stored compressed in memory, and transferred to theclient compressed when requested.  When active, documents are stored compressesin memory and when transferred to the client.  This field must be 'off', 'passive'or 'active', defaulting to 'passive'.  Be aware 'off' in YAML 1.2 is a boolean, somust be quoted as a string in configuration files.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("off", "passive", "active"),
						},
					},

					"conflict_resolution": schema.StringAttribute{
						Description:         "ConflictResolution defines how XDCR handles concurrent write conflicts.  Sequence numberbased resolution selects the document with the highest sequence number as the most recent.Timestamp based resolution selects the document that was written to most recently as themost recent.  This field must be 'seqno' (sequence based), or 'lww' (timestamp based),defaulting to 'seqno'.",
						MarkdownDescription: "ConflictResolution defines how XDCR handles concurrent write conflicts.  Sequence numberbased resolution selects the document with the highest sequence number as the most recent.Timestamp based resolution selects the document that was written to most recently as themost recent.  This field must be 'seqno' (sequence based), or 'lww' (timestamp based),defaulting to 'seqno'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("seqno", "lww"),
						},
					},

					"enable_flush": schema.BoolAttribute{
						Description:         "EnableFlush defines whether a client can delete all documents in a bucket.This field defaults to false.",
						MarkdownDescription: "EnableFlush defines whether a client can delete all documents in a bucket.This field defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"eviction_policy": schema.StringAttribute{
						Description:         "EvictionPolicy controls how Couchbase handles memory exhaustion.  No eviction meansthat Couchbase server will make this bucket read-only when memory is exhausted inorder to avoid data loss.  NRU eviction will delete documents that haven't been usedrecently in order to free up memory. This field must be 'noEviction' or 'nruEviction',defaulting to 'noEviction'.",
						MarkdownDescription: "EvictionPolicy controls how Couchbase handles memory exhaustion.  No eviction meansthat Couchbase server will make this bucket read-only when memory is exhausted inorder to avoid data loss.  NRU eviction will delete documents that haven't been usedrecently in order to free up memory. This field must be 'noEviction' or 'nruEviction',defaulting to 'noEviction'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("noEviction", "nruEviction"),
						},
					},

					"io_priority": schema.StringAttribute{
						Description:         "IOPriority controls how many threads a bucket has, per pod, to process reads and writes.This field must be 'low' or 'high', defaulting to 'low'.  Modification of this field willcause a temporary service disruption as threads are restarted.",
						MarkdownDescription: "IOPriority controls how many threads a bucket has, per pod, to process reads and writes.This field must be 'low' or 'high', defaulting to 'low'.  Modification of this field willcause a temporary service disruption as threads are restarted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("low", "high"),
						},
					},

					"max_ttl": schema.StringAttribute{
						Description:         "MaxTTL defines how long a document is permitted to exist for, withoutmodification, until it is automatically deleted.  This is a default and maximumtime-to-live and may be set to a lower value by the client.  If the client specifiesa higher value, then it is truncated to the maximum durability.  Documents areremoved by Couchbase, after they have expired, when either accessed, the expirypager is run, or the bucket is compacted.  When set to 0, then documents are notexpired by default.  This field must be a duration in the range 0-2147483648s,defaulting to 0.  More info:https://golang.org/pkg/time/#ParseDuration",
						MarkdownDescription: "MaxTTL defines how long a document is permitted to exist for, withoutmodification, until it is automatically deleted.  This is a default and maximumtime-to-live and may be set to a lower value by the client.  If the client specifiesa higher value, then it is truncated to the maximum durability.  Documents areremoved by Couchbase, after they have expired, when either accessed, the expirypager is run, or the bucket is compacted.  When set to 0, then documents are notexpired by default.  This field must be a duration in the range 0-2147483648s,defaulting to 0.  More info:https://golang.org/pkg/time/#ParseDuration",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"memory_quota": schema.StringAttribute{
						Description:         "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded,documents will be evicted from memory defined by the eviction policy.  The memory quotais defined per Couchbase pod running the data service.  This field defaults to, and mustbe greater than or equal to 100Mi.  More info:https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						MarkdownDescription: "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded,documents will be evicted from memory defined by the eviction policy.  The memory quotais defined per Couchbase pod running the data service.  This field defaults to, and mustbe greater than or equal to 100Mi.  More info:https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
						},
					},

					"minimum_durability": schema.StringAttribute{
						Description:         "MiniumumDurability defines how durable a document write is by default, and canbe made more durable by the client.  This feature enables ACID transactions.When none, Couchbase server will respond when the document is in memory, it willbecome eventually consistent across the cluster.  When majority, Couchbase server willrespond when the document is replicated to at least half of the pods running thedata service in the cluster.  This field must be either 'none' or 'majority',defaulting to 'none'.",
						MarkdownDescription: "MiniumumDurability defines how durable a document write is by default, and canbe made more durable by the client.  This feature enables ACID transactions.When none, Couchbase server will respond when the document is in memory, it willbecome eventually consistent across the cluster.  When majority, Couchbase server willrespond when the document is replicated to at least half of the pods running thedata service in the cluster.  This field must be either 'none' or 'majority',defaulting to 'none'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "majority"),
						},
					},

					"name": schema.StringAttribute{
						Description:         "Name is the name of the bucket within Couchbase server.  By default the Operatorwill use the 'metadata.name' field to define the bucket name.  The 'metadata.name'field only supports a subset of the supported character set.  When specified, thisfield overrides 'metadata.name'.  Legal bucket names have a maximum length of 100characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						MarkdownDescription: "Name is the name of the bucket within Couchbase server.  By default the Operatorwill use the 'metadata.name' field to define the bucket name.  The 'metadata.name'field only supports a subset of the supported character set.  When specified, thisfield overrides 'metadata.name'.  Legal bucket names have a maximum length of 100characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '-_%.'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(100),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-_%\.]{1,100}$`), ""),
						},
					},

					"rank": schema.Int64Attribute{
						Description:         "Rank determines the bucket’s place in the order in which the rebalance processhandles the buckets on the cluster. The higher a bucket’s assigned integer(in relation to the integers assigned other buckets), the sooner in therebalance process the bucket is handled. This assignment of rank allows acluster’s most mission-critical data to be rebalanced with top priority.This option is only supported for Couchbase Server 7.6.0+.",
						MarkdownDescription: "Rank determines the bucket’s place in the order in which the rebalance processhandles the buckets on the cluster. The higher a bucket’s assigned integer(in relation to the integers assigned other buckets), the sooner in therebalance process the bucket is handled. This assignment of rank allows acluster’s most mission-critical data to be rebalanced with top priority.This option is only supported for Couchbase Server 7.6.0+.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(1000),
						},
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas defines how many copies of documents Couchbase server maintains.  This directlyaffects how fault tolerant a Couchbase cluster is.  With a single replica, the clustercan tolerate one data pod going down and still service requests without data loss.  Thenumber of replicas also affect memory use.  With a single replica, the effective memoryquota for documents is halved, with two replicas it is one third.  The number of replicasmust be between 0 and 3, defaulting to 1.",
						MarkdownDescription: "Replicas defines how many copies of documents Couchbase server maintains.  This directlyaffects how fault tolerant a Couchbase cluster is.  With a single replica, the clustercan tolerate one data pod going down and still service requests without data loss.  Thenumber of replicas also affect memory use.  With a single replica, the effective memoryquota for documents is halved, with two replicas it is one third.  The number of replicasmust be between 0 and 3, defaulting to 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(3),
						},
					},

					"scopes": schema.SingleNestedAttribute{
						Description:         "Scopes defines whether the Operator manages scopes for the bucket or not, andthe set of scopes defined for the bucket.",
						MarkdownDescription: "Scopes defines whether the Operator manages scopes for the bucket or not, andthe set of scopes defined for the bucket.",
						Attributes: map[string]schema.Attribute{
							"managed": schema.BoolAttribute{
								Description:         "Managed defines whether scopes are managed for this bucket.This field is 'false' by default, and the Operator will take no actions thatwill affect scopes and collections in this bucket.  The default scope andcollection will be present.  When set to 'true', the Operator will manageuser defined scopes, and optionally, their collections as defined by the'CouchbaseScope', 'CouchbaseScopeGroup', 'CouchbaseCollection' and'CouchbaseCollectionGroup' resource documentation.  If this field is set to'false' while the  already managed, then the Operator will leave whateverconfiguration is already present.",
								MarkdownDescription: "Managed defines whether scopes are managed for this bucket.This field is 'false' by default, and the Operator will take no actions thatwill affect scopes and collections in this bucket.  The default scope andcollection will be present.  When set to 'true', the Operator will manageuser defined scopes, and optionally, their collections as defined by the'CouchbaseScope', 'CouchbaseScopeGroup', 'CouchbaseCollection' and'CouchbaseCollectionGroup' resource documentation.  If this field is set to'false' while the  already managed, then the Operator will leave whateverconfiguration is already present.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.ListNestedAttribute{
								Description:         "Resources is an explicit list of named resources that will be consideredfor inclusion in this bucket.  If a resource reference doesn'tmatch a resource, then no error conditions are raised due to undefinedresource creation ordering and eventual consistency.",
								MarkdownDescription: "Resources is an explicit list of named resources that will be consideredfor inclusion in this bucket.  If a resource reference doesn'tmatch a resource, then no error conditions are raised due to undefinedresource creation ordering and eventual consistency.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind indicates the kind of resource that is being referenced.  A scopecan only reference 'CouchbaseScope' and 'CouchbaseScopeGroup'resource kinds.  This field defaults to 'CouchbaseScope' if notspecified.",
											MarkdownDescription: "Kind indicates the kind of resource that is being referenced.  A scopecan only reference 'CouchbaseScope' and 'CouchbaseScopeGroup'resource kinds.  This field defaults to 'CouchbaseScope' if notspecified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("CouchbaseScope", "CouchbaseScopeGroup"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the Kubernetes resource name that is being referenced.Legal scope names have a maximum length of 251characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
											MarkdownDescription: "Name is the name of the Kubernetes resource name that is being referenced.Legal scope names have a maximum length of 251characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
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
								Description:         "Selector allows resources to be implicitly considered for inclusion in thisbucket.  More info:https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
								MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in thisbucket.  More info:https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
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
													Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_ephemeral_bucket_v2_manifest")

	var model CouchbaseComCouchbaseEphemeralBucketV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseEphemeralBucket")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
