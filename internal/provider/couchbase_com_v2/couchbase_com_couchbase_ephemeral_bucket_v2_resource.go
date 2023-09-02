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
)

var (
	_ resource.Resource                = &CouchbaseComCouchbaseEphemeralBucketV2Resource{}
	_ resource.ResourceWithConfigure   = &CouchbaseComCouchbaseEphemeralBucketV2Resource{}
	_ resource.ResourceWithImportState = &CouchbaseComCouchbaseEphemeralBucketV2Resource{}
)

func NewCouchbaseComCouchbaseEphemeralBucketV2Resource() resource.Resource {
	return &CouchbaseComCouchbaseEphemeralBucketV2Resource{}
}

type CouchbaseComCouchbaseEphemeralBucketV2Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CouchbaseComCouchbaseEphemeralBucketV2ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_ephemeral_bucket_v2"
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The CouchbaseEphemeralBucket resource defines a set of documents in Couchbase server. A Couchbase client connects to and operates on a bucket, which provides independent management of a set documents and a security boundary for role based access control. A CouchbaseEphemeralBucket provides in-memory only storage and replication for documents contained by it.",
		MarkdownDescription: "The CouchbaseEphemeralBucket resource defines a set of documents in Couchbase server. A Couchbase client connects to and operates on a bucket, which provides independent management of a set documents and a security boundary for role based access control. A CouchbaseEphemeralBucket provides in-memory only storage and replication for documents contained by it.",
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

			"wait_for": schema.ListNestedAttribute{
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
				Description:         "CouchbaseEphemeralBucketSpec is the specification for an ephemeral Couchbase bucket resource, and allows the bucket to be customized.",
				MarkdownDescription: "CouchbaseEphemeralBucketSpec is the specification for an ephemeral Couchbase bucket resource, and allows the bucket to be customized.",
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

					"eviction_policy": schema.StringAttribute{
						Description:         "EvictionPolicy controls how Couchbase handles memory exhaustion.  No eviction means that Couchbase server will make this bucket read-only when memory is exhausted in order to avoid data loss.  NRU eviction will delete documents that haven't been used recently in order to free up memory. This field must be 'noEviction' or 'nruEviction', defaulting to 'noEviction'.",
						MarkdownDescription: "EvictionPolicy controls how Couchbase handles memory exhaustion.  No eviction means that Couchbase server will make this bucket read-only when memory is exhausted in order to avoid data loss.  NRU eviction will delete documents that haven't been used recently in order to free up memory. This field must be 'noEviction' or 'nruEviction', defaulting to 'noEviction'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("noEviction", "nruEviction"),
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
						Description:         "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded, documents will be evicted from memory defined by the eviction policy.  The memory quota is defined per Couchbase pod running the data service.  This field defaults to, and must be greater than or equal to 100Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						MarkdownDescription: "MemoryQuota is a memory limit to the size of a bucket.  When this limit is exceeded, documents will be evicted from memory defined by the eviction policy.  The memory quota is defined per Couchbase pod running the data service.  This field defaults to, and must be greater than or equal to 100Mi.  More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#resource-units-in-kubernetes",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$`), ""),
						},
					},

					"minimum_durability": schema.StringAttribute{
						Description:         "MiniumumDurability defines how durable a document write is by default, and can be made more durable by the client.  This feature enables ACID transactions. When none, Couchbase server will respond when the document is in memory, it will become eventually consistent across the cluster.  When majority, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster.  This field must be either 'none' or 'majority', defaulting to 'none'.",
						MarkdownDescription: "MiniumumDurability defines how durable a document write is by default, and can be made more durable by the client.  This feature enables ACID transactions. When none, Couchbase server will respond when the document is in memory, it will become eventually consistent across the cluster.  When majority, Couchbase server will respond when the document is replicated to at least half of the pods running the data service in the cluster.  This field must be either 'none' or 'majority', defaulting to 'none'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "majority"),
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_couchbase_com_couchbase_ephemeral_bucket_v2")

	var model CouchbaseComCouchbaseEphemeralBucketV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseEphemeralBucket")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "CouchbaseEphemeralBucket"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CouchbaseComCouchbaseEphemeralBucketV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_ephemeral_bucket_v2")

	var data CouchbaseComCouchbaseEphemeralBucketV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "CouchbaseEphemeralBucket"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CouchbaseComCouchbaseEphemeralBucketV2ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_couchbase_com_couchbase_ephemeral_bucket_v2")

	var model CouchbaseComCouchbaseEphemeralBucketV2ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseEphemeralBucket")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "CouchbaseEphemeralBucket"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CouchbaseComCouchbaseEphemeralBucketV2ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_couchbase_com_couchbase_ephemeral_bucket_v2")

	var data CouchbaseComCouchbaseEphemeralBucketV2ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "couchbase.com", Version: "v2", Resource: "CouchbaseEphemeralBucket"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *CouchbaseComCouchbaseEphemeralBucketV2Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
