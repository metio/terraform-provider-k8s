/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kyverno_io_v2alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &KyvernoIoPolicyExceptionV2Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &KyvernoIoPolicyExceptionV2Alpha1DataSource{}
)

func NewKyvernoIoPolicyExceptionV2Alpha1DataSource() datasource.DataSource {
	return &KyvernoIoPolicyExceptionV2Alpha1DataSource{}
}

type KyvernoIoPolicyExceptionV2Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KyvernoIoPolicyExceptionV2Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Background *bool `tfsdk:"background" json:"background,omitempty"`
		Exceptions *[]struct {
			PolicyName *string   `tfsdk:"policy_name" json:"policyName,omitempty"`
			RuleNames  *[]string `tfsdk:"rule_names" json:"ruleNames,omitempty"`
		} `tfsdk:"exceptions" json:"exceptions,omitempty"`
		Match *struct {
			All *[]struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"all" json:"all,omitempty"`
			Any *[]struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Resources    *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Kinds             *[]string          `tfsdk:"kinds" json:"kinds,omitempty"`
					Name              *string            `tfsdk:"name" json:"name,omitempty"`
					Names             *[]string          `tfsdk:"names" json:"names,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					Operations *[]string `tfsdk:"operations" json:"operations,omitempty"`
					Selector   *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Roles    *[]string `tfsdk:"roles" json:"roles,omitempty"`
				Subjects *[]struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"subjects" json:"subjects,omitempty"`
			} `tfsdk:"any" json:"any,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KyvernoIoPolicyExceptionV2Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kyverno_io_policy_exception_v2alpha1"
}

func (r *KyvernoIoPolicyExceptionV2Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PolicyException declares resources to be excluded from specified policies.",
		MarkdownDescription: "PolicyException declares resources to be excluded from specified policies.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec declares policy exception behaviors.",
				MarkdownDescription: "Spec declares policy exception behaviors.",
				Attributes: map[string]schema.Attribute{
					"background": schema.BoolAttribute{
						Description:         "Background controls if exceptions are applied to existing policies during a background scan. Optional. Default value is 'true'. The value must be set to 'false' if the policy rule uses variables that are only available in the admission review request (e.g. user name).",
						MarkdownDescription: "Background controls if exceptions are applied to existing policies during a background scan. Optional. Default value is 'true'. The value must be set to 'false' if the policy rule uses variables that are only available in the admission review request (e.g. user name).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"exceptions": schema.ListNestedAttribute{
						Description:         "Exceptions is a list policy/rules to be excluded",
						MarkdownDescription: "Exceptions is a list policy/rules to be excluded",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"policy_name": schema.StringAttribute{
									Description:         "PolicyName identifies the policy to which the exception is applied. The policy name uses the format <namespace>/<name> unless it references a ClusterPolicy.",
									MarkdownDescription: "PolicyName identifies the policy to which the exception is applied. The policy name uses the format <namespace>/<name> unless it references a ClusterPolicy.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"rule_names": schema.ListAttribute{
									Description:         "RuleNames identifies the rules to which the exception is applied.",
									MarkdownDescription: "RuleNames identifies the rules to which the exception is applied.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"match": schema.SingleNestedAttribute{
						Description:         "Match defines match clause used to check if a resource applies to the exception",
						MarkdownDescription: "Match defines match clause used to check if a resource applies to the exception",
						Attributes: map[string]schema.Attribute{
							"all": schema.ListNestedAttribute{
								Description:         "All allows specifying resources which will be ANDed",
								MarkdownDescription: "All allows specifying resources which will be ANDed",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"any": schema.ListNestedAttribute{
								Description:         "Any allows specifying resources which will be ORed",
								MarkdownDescription: "Any allows specifying resources which will be ORed",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_roles": schema.ListAttribute{
											Description:         "ClusterRoles is the list of cluster-wide role names for the user.",
											MarkdownDescription: "ClusterRoles is the list of cluster-wide role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "ResourceDescription contains information about the resource being created or modified.",
											MarkdownDescription: "ResourceDescription contains information about the resource being created or modified.",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													MarkdownDescription: "Annotations is a  map of annotations (key-value pairs of type string). Annotation keys and values support the wildcard characters '*' (matches zero or many characters) and '?' (matches at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"kinds": schema.ListAttribute{
													Description:         "Kinds is a list of resource kinds.",
													MarkdownDescription: "Kinds is a list of resource kinds.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													MarkdownDescription: "Name is the name of the resource. The name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character). NOTE: 'Name' is being deprecated in favor of 'Names'.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"names": schema.ListAttribute{
													Description:         "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Names are the names of the resources. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "NamespaceSelector is a label selector for the resource namespace. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character).Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"namespaces": schema.ListAttribute{
													Description:         "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													MarkdownDescription: "Namespaces is a list of namespaces names. Each name supports wildcard characters '*' (matches zero or many characters) and '?' (at least one character).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operations": schema.ListAttribute{
													Description:         "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													MarkdownDescription: "Operations can contain values ['CREATE, 'UPDATE', 'CONNECT', 'DELETE'], which are used to match a specific action.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													MarkdownDescription: "Selector is a label selector. Label keys and values in 'matchLabels' support the wildcard characters '*' (matches zero or many characters) and '?' (matches one character). Wildcards allows writing label selectors like ['storage.k8s.io/*': '*']. Note that using ['*' : '*'] matches any key and value but does not match an empty label set.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of namespaced role names for the user.",
											MarkdownDescription: "Roles is the list of namespaced role names for the user.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"subjects": schema.ListNestedAttribute{
											Description:         "Subjects is the list of subject names like users, user groups, and service accounts.",
											MarkdownDescription: "Subjects is the list of subject names like users, user groups, and service accounts.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														MarkdownDescription: "APIGroup holds the API group of the referenced subject. Defaults to '' for ServiceAccount subjects. Defaults to 'rbac.authorization.k8s.io' for User and Group subjects.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														MarkdownDescription: "Kind of object being referenced. Values defined by this API group are 'User', 'Group', and 'ServiceAccount'. If the Authorizer does not recognized the kind value, the Authorizer should report an error.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the object being referenced.",
														MarkdownDescription: "Name of the object being referenced.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														MarkdownDescription: "Namespace of the referenced object.  If the object kind is non-namespace, such as 'User' or 'Group', and this value is not empty the Authorizer should report an error.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *KyvernoIoPolicyExceptionV2Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KyvernoIoPolicyExceptionV2Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kyverno_io_policy_exception_v2alpha1")

	var data KyvernoIoPolicyExceptionV2Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kyverno.io", Version: "v2alpha1", Resource: "policyexceptions"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse KyvernoIoPolicyExceptionV2Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("kyverno.io/v2alpha1")
	data.Kind = pointer.String("PolicyException")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
