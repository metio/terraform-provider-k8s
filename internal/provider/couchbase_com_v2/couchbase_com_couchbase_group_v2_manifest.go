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
	_ datasource.DataSource = &CouchbaseComCouchbaseGroupV2Manifest{}
)

func NewCouchbaseComCouchbaseGroupV2Manifest() datasource.DataSource {
	return &CouchbaseComCouchbaseGroupV2Manifest{}
}

type CouchbaseComCouchbaseGroupV2Manifest struct{}

type CouchbaseComCouchbaseGroupV2ManifestData struct {
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
		LdapGroupRef *string `tfsdk:"ldap_group_ref" json:"ldapGroupRef,omitempty"`
		Roles        *[]struct {
			Bucket  *string `tfsdk:"bucket" json:"bucket,omitempty"`
			Buckets *struct {
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
			} `tfsdk:"buckets" json:"buckets,omitempty"`
			Collections *struct {
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
			} `tfsdk:"collections" json:"collections,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Scopes *struct {
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
		} `tfsdk:"roles" json:"roles,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CouchbaseComCouchbaseGroupV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_couchbase_com_couchbase_group_v2_manifest"
}

func (r *CouchbaseComCouchbaseGroupV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CouchbaseGroup allows the automation of Couchbase group management.",
		MarkdownDescription: "CouchbaseGroup allows the automation of Couchbase group management.",
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
				Description:         "CouchbaseGroupSpec allows the specification of Couchbase group configuration.",
				MarkdownDescription: "CouchbaseGroupSpec allows the specification of Couchbase group configuration.",
				Attributes: map[string]schema.Attribute{
					"ldap_group_ref": schema.StringAttribute{
						Description:         "LDAPGroupRef is a reference to an LDAP group.",
						MarkdownDescription: "LDAPGroupRef is a reference to an LDAP group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"roles": schema.ListNestedAttribute{
						Description:         "Roles is a list of roles that this group is granted.",
						MarkdownDescription: "Roles is a list of roles that this group is granted.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"bucket": schema.StringAttribute{
									Description:         "Bucket name for bucket admin roles. When not specified for a role that can be scoped to a specific bucket, the role will apply to all buckets in the cluster. Deprecated: Couchbase Autonomous Operator 2.3",
									MarkdownDescription: "Bucket name for bucket admin roles. When not specified for a role that can be scoped to a specific bucket, the role will apply to all buckets in the cluster. Deprecated: Couchbase Autonomous Operator 2.3",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^\*$|^[a-zA-Z0-9-_%\.]+$`), ""),
									},
								},

								"buckets": schema.SingleNestedAttribute{
									Description:         "Bucket level access to apply to specified role. The bucket must exist. When not specified, the bucket field will be checked. If both are empty and the role can be scoped to a specific bucket, the role will apply to all buckets in the cluster",
									MarkdownDescription: "Bucket level access to apply to specified role. The bucket must exist. When not specified, the bucket field will be checked. If both are empty and the role can be scoped to a specific bucket, the role will apply to all buckets in the cluster",
									Attributes: map[string]schema.Attribute{
										"resources": schema.ListNestedAttribute{
											Description:         "Resources is an explicit list of named bucket resources that will be considered for inclusion in this role. If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
											MarkdownDescription: "Resources is an explicit list of named bucket resources that will be considered for inclusion in this role. If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"kind": schema.StringAttribute{
														Description:         "Kind indicates the kind of resource that is being referenced. A Role can only reference 'CouchbaseBucket' kind. This field defaults to 'CouchbaseBucket' if not specified.",
														MarkdownDescription: "Kind indicates the kind of resource that is being referenced. A Role can only reference 'CouchbaseBucket' kind. This field defaults to 'CouchbaseBucket' if not specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("CouchbaseBucket"),
														},
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the Kubernetes resource name that is being referenced.",
														MarkdownDescription: "Name is the name of the Kubernetes resource name that is being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector allows resources to be implicitly considered for inclusion in this role. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
											MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in this role. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
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

								"collections": schema.SingleNestedAttribute{
									Description:         "Collection level access to apply to the specified role. The collection must exist. When not specified, the role is subject to scope or bucket level access.",
									MarkdownDescription: "Collection level access to apply to the specified role. The collection must exist. When not specified, the role is subject to scope or bucket level access.",
									Attributes: map[string]schema.Attribute{
										"resources": schema.ListNestedAttribute{
											Description:         "Resources is an explicit list of named resources that will be considered for inclusion in this collection or collections. If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
											MarkdownDescription: "Resources is an explicit list of named resources that will be considered for inclusion in this collection or collections. If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"kind": schema.StringAttribute{
														Description:         "Kind indicates the kind of resource that is being referenced. A scope can only reference 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource kinds. This field defaults to 'CouchbaseCollection' if not specified.",
														MarkdownDescription: "Kind indicates the kind of resource that is being referenced. A scope can only reference 'CouchbaseCollection' and 'CouchbaseCollectionGroup' resource kinds. This field defaults to 'CouchbaseCollection' if not specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("CouchbaseCollection", "CouchbaseCollectionGroup"),
														},
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the Kubernetes resource name that is being referenced. Legal collection names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
														MarkdownDescription: "Name is the name of the Kubernetes resource name that is being referenced. Legal collection names have a maximum length of 251 characters and may be composed of any character from 'a-z', 'A-Z', '0-9' and '_-%'.",
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
											Description:         "Selector allows resources to be implicitly considered for inclusion in this collection or collections. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
											MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in this collection or collections. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
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

								"name": schema.StringAttribute{
									Description:         "Name of role.",
									MarkdownDescription: "Name of role.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("admin", "analytics_admin", "analytics_manager", "analytics_reader", "analytics_select", "backup_admin", "bucket_admin", "bucket_full_access", "cluster_admin", "data_backup", "data_dcp_reader", "data_monitoring", "data_reader", "data_writer", "eventing_admin", "external_stats_reader", "fts_admin", "fts_searcher", "mobile_sync_gateway", "mobile_sync_gateway_application", "mobile_sync_gateway_application_read_only", "mobile_sync_gateway_architect", "mobile_sync_gateway_dev_ops", "mobile_sync_gateway_replicator", "query_delete", "query_execute_external_functions", "query_execute_functions", "query_execute_global_external_functions", "query_execute_global_functions", "query_external_access", "query_insert", "query_manage_external_functions", "query_manage_functions", "query_manage_global_external_functions", "query_manage_global_functions", "query_manage_index", "query_select", "query_system_catalog", "query_update", "replication_admin", "replication_target", "ro_admin", "scope_admin", "security_admin", "security_admin_external", "security_admin_local", "views_admin", "views_reader", "eventing_manage_functions"),
									},
								},

								"scopes": schema.SingleNestedAttribute{
									Description:         "Scope level access to apply to specified role. The scope must exist. When not specified, the role will apply to selected bucket or all buckets in the cluster.",
									MarkdownDescription: "Scope level access to apply to specified role. The scope must exist. When not specified, the role will apply to selected bucket or all buckets in the cluster.",
									Attributes: map[string]schema.Attribute{
										"resources": schema.ListNestedAttribute{
											Description:         "Resources is an explicit list of named resources that will be considered for inclusion in this scope or scopes. If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
											MarkdownDescription: "Resources is an explicit list of named resources that will be considered for inclusion in this scope or scopes. If a resource reference doesn't match a resource, then no error conditions are raised due to undefined resource creation ordering and eventual consistency.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"kind": schema.StringAttribute{
														Description:         "Kind indicates the kind of resource that is being referenced. A scope can only reference 'CouchbaseScope' and 'CouchbaseScopeGroup' resource kinds. This field defaults to 'CouchbaseScope' if not specified.",
														MarkdownDescription: "Kind indicates the kind of resource that is being referenced. A scope can only reference 'CouchbaseScope' and 'CouchbaseScopeGroup' resource kinds. This field defaults to 'CouchbaseScope' if not specified.",
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
											Description:         "Selector allows resources to be implicitly considered for inclusion in this scope or scopes. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
											MarkdownDescription: "Selector allows resources to be implicitly considered for inclusion in this scope or scopes. More info: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.28/#labelselector-v1-meta",
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CouchbaseComCouchbaseGroupV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_couchbase_com_couchbase_group_v2_manifest")

	var model CouchbaseComCouchbaseGroupV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("couchbase.com/v2")
	model.Kind = pointer.String("CouchbaseGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
