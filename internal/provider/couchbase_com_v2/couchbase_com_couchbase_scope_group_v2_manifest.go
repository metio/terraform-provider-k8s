/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package couchbase_com_v2

import (
	"context"
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
	_ datasource.DataSource = &CouchbaseComCouchbaseScopeGroupV2Manifest{}
)

func NewCouchbaseComCouchbaseScopeGroupV2Manifest() datasource.DataSource {
	return &CouchbaseComCouchbaseScopeGroupV2Manifest{}
}

type CouchbaseComCouchbaseScopeGroupV2Manifest struct{}

type CouchbaseComCouchbaseScopeGroupV2ManifestData struct {
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
		Collections *struct {
			Managed                   *bool `tfsdk:"managed" json:"managed,omitempty"`
			PreserveDefaultCollection *bool `tfsdk:"preserve_default_collection" json:"preserveDefaultCollection,omitempty"`
			Resources                 *[]struct {
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
		} `tfsdk:"collections" json:"collections,omitempty"`
		Names *[]string `tfsdk:"names" json:"names,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseScopeGroupV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_scope_group_v2_manifest"
}

func (r *CouchbaseComCouchbaseScopeGroupV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CouchbaseScopeGroup represents a logical unit of data storage that sits between buckets andcollections e.g. a bucket may contain multiple scopes, and a scope may contain multiplecollections.  At present, scopes are not nested, so provide only a single level ofabstraction.  Scopes provide a coarser grained basis for role-based access control (RBAC)and cross-datacenter replication (XDCR) than collections, but finer that buckets.In order to be considered by the Operator, a scope must be referenced by either a'CouchbaseBucket' or 'CouchbaseEphemeralBucket' resource.Unlike 'CouchbaseScope' resources, scope groups represents multiple scopes, with the samecommon set of collections, to be expressed as a single resource, minimizing requiredconfiguration and Kubernetes API traffic.  It also forms the basis of Couchbase RBACsecurity boundaries.",
		MarkdownDescription: "CouchbaseScopeGroup represents a logical unit of data storage that sits between buckets andcollections e.g. a bucket may contain multiple scopes, and a scope may contain multiplecollections.  At present, scopes are not nested, so provide only a single level ofabstraction.  Scopes provide a coarser grained basis for role-based access control (RBAC)and cross-datacenter replication (XDCR) than collections, but finer that buckets.In order to be considered by the Operator, a scope must be referenced by either a'CouchbaseBucket' or 'CouchbaseEphemeralBucket' resource.Unlike 'CouchbaseScope' resources, scope groups represents multiple scopes, with the samecommon set of collections, to be expressed as a single resource, minimizing requiredconfiguration and Kubernetes API traffic.  It also forms the basis of Couchbase RBACsecurity boundaries.",
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
				Description:         "Spec defines the desired state of the resource.",
				MarkdownDescription: "Spec defines the desired state of the resource.",
				Attributes: map[string]schema.Attribute{
					"collections": schema.SingleNestedAttribute{
						Description:         "Collections defines how to collate collections included in this scope or scope group.Any of the provided methods may be used to collate a set of collections tomanage.  Collated collections must have unique names, otherwise it isconsidered ambiguous, and an error condition.",
						MarkdownDescription: "Collections defines how to collate collections included in this scope or scope group.Any of the provided methods may be used to collate a set of collections tomanage.  Collated collections must have unique names, otherwise it isconsidered ambiguous, and an error condition.",
						Attributes: map[string]schema.Attribute{
							"managed": schema.BoolAttribute{
								Description:         "Managed indicates whether collections within this scope are managed.If not then you can dynamically create and delete collections withthe Couchbase UI or SDKs.",
								MarkdownDescription: "Managed indicates whether collections within this scope are managed.If not then you can dynamically create and delete collections withthe Couchbase UI or SDKs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"preserve_default_collection": schema.BoolAttribute{
								Description:         "PreserveDefaultCollection indicates whether the Operator should manage thedefault collection within the default scope.  The default collection canbe deleted, but can not be recreated by Couchbase Server.  By setting thisfield to 'true', the Operator will implicitly manage the default collectionwithin the default scope.  The default collection cannot be modified andwill have no document time-to-live (TTL).  When set to 'false', the operatorwill not manage the default collection, which will be deleted and cannot beused or recreated.",
								MarkdownDescription: "PreserveDefaultCollection indicates whether the Operator should manage thedefault collection within the default scope.  The default collection canbe deleted, but can not be recreated by Couchbase Server.  By setting thisfield to 'true', the Operator will implicitly manage the default collectionwithin the default scope.  The default collection cannot be modified andwill have no document time-to-live (TTL).  When set to 'false', the operatorwill not manage the default collection, which will be deleted and cannot beused or recreated.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.ListNestedAttribute{
								Description:         "Resources is an explicit list of named resources that will be consideredfor inclusion in this scope or scopes.  If a resource reference doesn'tmatch a resource, then no error conditions are raised due to undefinedresource creation ordering and eventual consistency.",
								MarkdownDescription: "Resources is an explicit list of named resources that will be consideredfor inclusion in this scope or scopes.  If a resource reference doesn'tmatch a resource, then no error conditions are raised due to undefinedresource creation ordering and eventual consistency.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "Kind indicates the kind of resource that is being referenced.  A scopecan only reference 'CouchbaseCollection' and 'CouchbaseCollectionGroup'resource kinds.  This field defaults to 'CouchbaseCollection' if notspecified.",
											MarkdownDescription: "Kind indicates the kind of resource that is being referenced.  A scopecan only reference 'CouchbaseCollection' and 'CouchbaseCollectionGroup'resource kinds.  This field defaults to 'CouchbaseCollection' if notspecified.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("CouchbaseCollection", "CouchbaseCollectionGroup"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name of the Kubernetes resource name that is being referenced.Legal collection names have a maximum length of 251characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
											MarkdownDescription: "Name is the name of the Kubernetes resource name that is being referenced.Legal collection names have a maximum length of 251characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
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
								Description:         "Selector allows resources to be implicitly considered for inclusion in thisscope or scopes.  More info:https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
								MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in thisscope or scopes.  More info:https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
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

					"names": schema.ListAttribute{
						Description:         "Names specifies the names of the scopes.  Unlike CouchbaseScope, whichspecifies a single scope, a scope group specifies multiple, and thescope group must specify at least one scope name.Any scope names specified must be unique.Scope names must be 1-251 characters in length, contain only [a-zA-Z0-9_-%]and not start with either _ or %.",
						MarkdownDescription: "Names specifies the names of the scopes.  Unlike CouchbaseScope, whichspecifies a single scope, a scope group specifies multiple, and thescope group must specify at least one scope name.Any scope names specified must be unique.Scope names must be 1-251 characters in length, contain only [a-zA-Z0-9_-%]and not start with either _ or %.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CouchbaseComCouchbaseScopeGroupV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_scope_group_v2_manifest")

	var model CouchbaseComCouchbaseScopeGroupV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseScopeGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
