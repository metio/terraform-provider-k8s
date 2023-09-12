/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mutations_gatekeeper_sh_v1beta1

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &MutationsGatekeeperShModifySetV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &MutationsGatekeeperShModifySetV1Beta1DataSource{}
)

func NewMutationsGatekeeperShModifySetV1Beta1DataSource() datasource.DataSource {
	return &MutationsGatekeeperShModifySetV1Beta1DataSource{}
}

type MutationsGatekeeperShModifySetV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type MutationsGatekeeperShModifySetV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ApplyTo *[]struct {
			Groups   *[]string `tfsdk:"groups" json:"groups,omitempty"`
			Kinds    *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
			Versions *[]string `tfsdk:"versions" json:"versions,omitempty"`
		} `tfsdk:"apply_to" json:"applyTo,omitempty"`
		Location *string `tfsdk:"location" json:"location,omitempty"`
		Match    *struct {
			ExcludedNamespaces *[]string `tfsdk:"excluded_namespaces" json:"excludedNamespaces,omitempty"`
			Kinds              *[]struct {
				ApiGroups *[]string `tfsdk:"api_groups" json:"apiGroups,omitempty"`
				Kinds     *[]string `tfsdk:"kinds" json:"kinds,omitempty"`
			} `tfsdk:"kinds" json:"kinds,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name              *string `tfsdk:"name" json:"name,omitempty"`
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			Namespaces *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
			Scope      *string   `tfsdk:"scope" json:"scope,omitempty"`
			Source     *string   `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
		Parameters *struct {
			Operation *string `tfsdk:"operation" json:"operation,omitempty"`
			PathTests *[]struct {
				Condition *string `tfsdk:"condition" json:"condition,omitempty"`
				SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			} `tfsdk:"path_tests" json:"pathTests,omitempty"`
			Values *map[string]string `tfsdk:"values" json:"values,omitempty"`
		} `tfsdk:"parameters" json:"parameters,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MutationsGatekeeperShModifySetV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_mutations_gatekeeper_sh_modify_set_v1beta1"
}

func (r *MutationsGatekeeperShModifySetV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ModifySet allows the user to modify non-keyed lists, such as the list of arguments to a container.",
		MarkdownDescription: "ModifySet allows the user to modify non-keyed lists, such as the list of arguments to a container.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ModifySetSpec defines the desired state of ModifySet.",
				MarkdownDescription: "ModifySetSpec defines the desired state of ModifySet.",
				Attributes: map[string]schema.Attribute{
					"apply_to": schema.ListNestedAttribute{
						Description:         "ApplyTo lists the specific groups, versions and kinds a mutation will be applied to. This is necessary because every mutation implies part of an object schema and object schemas are associated with specific GVKs.",
						MarkdownDescription: "ApplyTo lists the specific groups, versions and kinds a mutation will be applied to. This is necessary because every mutation implies part of an object schema and object schemas are associated with specific GVKs.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"groups": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kinds": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"versions": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"location": schema.StringAttribute{
						Description:         "Location describes the path to be mutated, for example: 'spec.containers[name: main].args'.",
						MarkdownDescription: "Location describes the path to be mutated, for example: 'spec.containers[name: main].args'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"match": schema.SingleNestedAttribute{
						Description:         "Match allows the user to limit which resources get mutated. Individual match criteria are AND-ed together. An undefined match criteria matches everything.",
						MarkdownDescription: "Match allows the user to limit which resources get mutated. Individual match criteria are AND-ed together. An undefined match criteria matches everything.",
						Attributes: map[string]schema.Attribute{
							"excluded_namespaces": schema.ListAttribute{
								Description:         "ExcludedNamespaces is a list of namespace names. If defined, a constraint only applies to resources not in a listed namespace. ExcludedNamespaces also supports a prefix or suffix based glob.  For example, 'excludedNamespaces: [kube-*]' matches both 'kube-system' and 'kube-public', and 'excludedNamespaces: [*-system]' matches both 'kube-system' and 'gatekeeper-system'.",
								MarkdownDescription: "ExcludedNamespaces is a list of namespace names. If defined, a constraint only applies to resources not in a listed namespace. ExcludedNamespaces also supports a prefix or suffix based glob.  For example, 'excludedNamespaces: [kube-*]' matches both 'kube-system' and 'kube-public', and 'excludedNamespaces: [*-system]' matches both 'kube-system' and 'gatekeeper-system'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kinds": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"api_groups": schema.ListAttribute{
											Description:         "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
											MarkdownDescription: "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"kinds": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"label_selector": schema.SingleNestedAttribute{
								Description:         "LabelSelector is the combination of two optional fields: 'matchLabels' and 'matchExpressions'.  These two fields provide different methods of selecting or excluding k8s objects based on the label keys and values included in object metadata.  All selection expressions from both sections are ANDed to determine if an object meets the cumulative requirements of the selector.",
								MarkdownDescription: "LabelSelector is the combination of two optional fields: 'matchLabels' and 'matchExpressions'.  These two fields provide different methods of selecting or excluding k8s objects based on the label keys and values included in object metadata.  All selection expressions from both sections are ANDed to determine if an object meets the cumulative requirements of the selector.",
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

							"name": schema.StringAttribute{
								Description:         "Name is the name of an object.  If defined, it will match against objects with the specified name.  Name also supports a prefix or suffix glob.  For example, 'name: pod-*' would match both 'pod-a' and 'pod-b', and 'name: *-pod' would match both 'a-pod' and 'b-pod'.",
								MarkdownDescription: "Name is the name of an object.  If defined, it will match against objects with the specified name.  Name also supports a prefix or suffix glob.  For example, 'name: pod-*' would match both 'pod-a' and 'pod-b', and 'name: *-pod' would match both 'a-pod' and 'b-pod'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace_selector": schema.SingleNestedAttribute{
								Description:         "NamespaceSelector is a label selector against an object's containing namespace or the object itself, if the object is a namespace.",
								MarkdownDescription: "NamespaceSelector is a label selector against an object's containing namespace or the object itself, if the object is a namespace.",
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
								Description:         "Namespaces is a list of namespace names. If defined, a constraint only applies to resources in a listed namespace.  Namespaces also supports a prefix or suffix based glob.  For example, 'namespaces: [kube-*]' matches both 'kube-system' and 'kube-public', and 'namespaces: [*-system]' matches both 'kube-system' and 'gatekeeper-system'.",
								MarkdownDescription: "Namespaces is a list of namespace names. If defined, a constraint only applies to resources in a listed namespace.  Namespaces also supports a prefix or suffix based glob.  For example, 'namespaces: [kube-*]' matches both 'kube-system' and 'kube-public', and 'namespaces: [*-system]' matches both 'kube-system' and 'gatekeeper-system'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"scope": schema.StringAttribute{
								Description:         "Scope determines if cluster-scoped and/or namespaced-scoped resources are matched.  Accepts '*', 'Cluster', or 'Namespaced'. (defaults to '*')",
								MarkdownDescription: "Scope determines if cluster-scoped and/or namespaced-scoped resources are matched.  Accepts '*', 'Cluster', or 'Namespaced'. (defaults to '*')",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"source": schema.StringAttribute{
								Description:         "Source determines whether generated or original resources are matched. Accepts 'Generated'|'Original'|'All' (defaults to 'All'). A value of 'Generated' will only match generated resources, while 'Original' will only match regular resources.",
								MarkdownDescription: "Source determines whether generated or original resources are matched. Accepts 'Generated'|'Original'|'All' (defaults to 'All'). A value of 'Generated' will only match generated resources, while 'Original' will only match regular resources.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"parameters": schema.SingleNestedAttribute{
						Description:         "Parameters define the behavior of the mutator.",
						MarkdownDescription: "Parameters define the behavior of the mutator.",
						Attributes: map[string]schema.Attribute{
							"operation": schema.StringAttribute{
								Description:         "Operation describes whether values should be merged in ('merge'), or pruned ('prune'). Default value is 'merge'",
								MarkdownDescription: "Operation describes whether values should be merged in ('merge'), or pruned ('prune'). Default value is 'merge'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"path_tests": schema.ListNestedAttribute{
								Description:         "PathTests are a series of existence tests that can be checked before a mutation is applied",
								MarkdownDescription: "PathTests are a series of existence tests that can be checked before a mutation is applied",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition": schema.StringAttribute{
											Description:         "Condition describes whether the path either MustExist or MustNotExist in the original object",
											MarkdownDescription: "Condition describes whether the path either MustExist or MustNotExist in the original object",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"values": schema.MapAttribute{
								Description:         "Values describes the values provided to the operation as 'values.fromList'.",
								MarkdownDescription: "Values describes the values provided to the operation as 'values.fromList'.",
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
		},
	}
}

func (r *MutationsGatekeeperShModifySetV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *MutationsGatekeeperShModifySetV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_mutations_gatekeeper_sh_modify_set_v1beta1")

	var data MutationsGatekeeperShModifySetV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "mutations.gatekeeper.sh", Version: "v1beta1", Resource: "modifyset"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse MutationsGatekeeperShModifySetV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("mutations.gatekeeper.sh/v1beta1")
	data.Kind = pointer.String("ModifySet")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
