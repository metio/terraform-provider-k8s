/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package mutations_gatekeeper_sh_v1alpha1

import (
	"context"
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
	_ datasource.DataSource = &MutationsGatekeeperShAssignV1Alpha1Manifest{}
)

func NewMutationsGatekeeperShAssignV1Alpha1Manifest() datasource.DataSource {
	return &MutationsGatekeeperShAssignV1Alpha1Manifest{}
}

type MutationsGatekeeperShAssignV1Alpha1Manifest struct{}

type MutationsGatekeeperShAssignV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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
			Assign *struct {
				ExternalData *struct {
					DataSource    *string `tfsdk:"data_source" json:"dataSource,omitempty"`
					Default       *string `tfsdk:"default" json:"default,omitempty"`
					FailurePolicy *string `tfsdk:"failure_policy" json:"failurePolicy,omitempty"`
					Provider      *string `tfsdk:"provider" json:"provider,omitempty"`
				} `tfsdk:"external_data" json:"externalData,omitempty"`
				FromMetadata *struct {
					Field *string `tfsdk:"field" json:"field,omitempty"`
				} `tfsdk:"from_metadata" json:"fromMetadata,omitempty"`
				Value *map[string]string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"assign" json:"assign,omitempty"`
			PathTests *[]struct {
				Condition *string `tfsdk:"condition" json:"condition,omitempty"`
				SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			} `tfsdk:"path_tests" json:"pathTests,omitempty"`
		} `tfsdk:"parameters" json:"parameters,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MutationsGatekeeperShAssignV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_mutations_gatekeeper_sh_assign_v1alpha1_manifest"
}

func (r *MutationsGatekeeperShAssignV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Assign is the Schema for the assign API.",
		MarkdownDescription: "Assign is the Schema for the assign API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "AssignSpec defines the desired state of Assign.",
				MarkdownDescription: "AssignSpec defines the desired state of Assign.",
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
									Optional:            true,
									Computed:            false,
								},

								"kinds": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"versions": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"location": schema.StringAttribute{
						Description:         "Location describes the path to be mutated, for example: 'spec.containers[name: main]'.",
						MarkdownDescription: "Location describes the path to be mutated, for example: 'spec.containers[name: main]'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
								Optional:            true,
								Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"kinds": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
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

							"name": schema.StringAttribute{
								Description:         "Name is the name of an object.  If defined, it will match against objects with the specified name.  Name also supports a prefix or suffix glob.  For example, 'name: pod-*' would match both 'pod-a' and 'pod-b', and 'name: *-pod' would match both 'a-pod' and 'b-pod'.",
								MarkdownDescription: "Name is the name of an object.  If defined, it will match against objects with the specified name.  Name also supports a prefix or suffix glob.  For example, 'name: pod-*' would match both 'pod-a' and 'pod-b', and 'name: *-pod' would match both 'a-pod' and 'b-pod'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\*|\*-)?[a-z0-9]([-:a-z0-9]*[a-z0-9])?(\*|-\*)?$`), ""),
								},
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

							"namespaces": schema.ListAttribute{
								Description:         "Namespaces is a list of namespace names. If defined, a constraint only applies to resources in a listed namespace.  Namespaces also supports a prefix or suffix based glob.  For example, 'namespaces: [kube-*]' matches both 'kube-system' and 'kube-public', and 'namespaces: [*-system]' matches both 'kube-system' and 'gatekeeper-system'.",
								MarkdownDescription: "Namespaces is a list of namespace names. If defined, a constraint only applies to resources in a listed namespace.  Namespaces also supports a prefix or suffix based glob.  For example, 'namespaces: [kube-*]' matches both 'kube-system' and 'kube-public', and 'namespaces: [*-system]' matches both 'kube-system' and 'gatekeeper-system'.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scope": schema.StringAttribute{
								Description:         "Scope determines if cluster-scoped and/or namespaced-scoped resources are matched.  Accepts '*', 'Cluster', or 'Namespaced'. (defaults to '*')",
								MarkdownDescription: "Scope determines if cluster-scoped and/or namespaced-scoped resources are matched.  Accepts '*', 'Cluster', or 'Namespaced'. (defaults to '*')",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source": schema.StringAttribute{
								Description:         "Source determines whether generated or original resources are matched. Accepts 'Generated'|'Original'|'All' (defaults to 'All'). A value of 'Generated' will only match generated resources, while 'Original' will only match regular resources.",
								MarkdownDescription: "Source determines whether generated or original resources are matched. Accepts 'Generated'|'Original'|'All' (defaults to 'All'). A value of 'Generated' will only match generated resources, while 'Original' will only match regular resources.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("All", "Generated", "Original"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"parameters": schema.SingleNestedAttribute{
						Description:         "Parameters define the behavior of the mutator.",
						MarkdownDescription: "Parameters define the behavior of the mutator.",
						Attributes: map[string]schema.Attribute{
							"assign": schema.SingleNestedAttribute{
								Description:         "Assign.value holds the value to be assigned",
								MarkdownDescription: "Assign.value holds the value to be assigned",
								Attributes: map[string]schema.Attribute{
									"external_data": schema.SingleNestedAttribute{
										Description:         "ExternalData describes the external data provider to be used for mutation.",
										MarkdownDescription: "ExternalData describes the external data provider to be used for mutation.",
										Attributes: map[string]schema.Attribute{
											"data_source": schema.StringAttribute{
												Description:         "DataSource specifies where to extract the data that will be sent to the external data provider as parameters.",
												MarkdownDescription: "DataSource specifies where to extract the data that will be sent to the external data provider as parameters.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ValueAtLocation", "Username"),
												},
											},

											"default": schema.StringAttribute{
												Description:         "Default specifies the default value to use when the external data provider returns an error and the failure policy is set to 'UseDefault'.",
												MarkdownDescription: "Default specifies the default value to use when the external data provider returns an error and the failure policy is set to 'UseDefault'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"failure_policy": schema.StringAttribute{
												Description:         "FailurePolicy specifies the policy to apply when the external data provider returns an error.",
												MarkdownDescription: "FailurePolicy specifies the policy to apply when the external data provider returns an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("UseDefault", "Ignore", "Fail"),
												},
											},

											"provider": schema.StringAttribute{
												Description:         "Provider is the name of the external data provider.",
												MarkdownDescription: "Provider is the name of the external data provider.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_metadata": schema.SingleNestedAttribute{
										Description:         "FromMetadata assigns a value from the specified metadata field.",
										MarkdownDescription: "FromMetadata assigns a value from the specified metadata field.",
										Attributes: map[string]schema.Attribute{
											"field": schema.StringAttribute{
												Description:         "Field specifies which metadata field provides the assigned value. Valid fields are 'namespace' and 'name'.",
												MarkdownDescription: "Field specifies which metadata field provides the assigned value. Valid fields are 'namespace' and 'name'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.MapAttribute{
										Description:         "Value is a constant value that will be assigned to 'location'",
										MarkdownDescription: "Value is a constant value that will be assigned to 'location'",
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

							"path_tests": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition": schema.StringAttribute{
											Description:         "Condition describes whether the path either MustExist or MustNotExist in the original object",
											MarkdownDescription: "Condition describes whether the path either MustExist or MustNotExist in the original object",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("MustExist", "MustNotExist"),
											},
										},

										"sub_path": schema.StringAttribute{
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
		},
	}
}

func (r *MutationsGatekeeperShAssignV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_mutations_gatekeeper_sh_assign_v1alpha1_manifest")

	var model MutationsGatekeeperShAssignV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("mutations.gatekeeper.sh/v1alpha1")
	model.Kind = pointer.String("Assign")

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
