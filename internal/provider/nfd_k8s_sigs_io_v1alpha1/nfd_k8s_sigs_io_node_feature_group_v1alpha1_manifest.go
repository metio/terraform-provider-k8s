/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package nfd_k8s_sigs_io_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &NfdK8SSigsIoNodeFeatureGroupV1Alpha1Manifest{}
)

func NewNfdK8SSigsIoNodeFeatureGroupV1Alpha1Manifest() datasource.DataSource {
	return &NfdK8SSigsIoNodeFeatureGroupV1Alpha1Manifest{}
}

type NfdK8SSigsIoNodeFeatureGroupV1Alpha1Manifest struct{}

type NfdK8SSigsIoNodeFeatureGroupV1Alpha1ManifestData struct {
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
		FeatureGroupRules *[]struct {
			MatchAny *[]struct {
				MatchFeatures *[]struct {
					Feature          *string `tfsdk:"feature" json:"feature,omitempty"`
					MatchExpressions *struct {
						Op    *string   `tfsdk:"op" json:"op,omitempty"`
						Value *[]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchName *struct {
						Op    *string   `tfsdk:"op" json:"op,omitempty"`
						Value *[]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"match_name" json:"matchName,omitempty"`
				} `tfsdk:"match_features" json:"matchFeatures,omitempty"`
			} `tfsdk:"match_any" json:"matchAny,omitempty"`
			MatchFeatures *[]struct {
				Feature          *string `tfsdk:"feature" json:"feature,omitempty"`
				MatchExpressions *struct {
					Op    *string   `tfsdk:"op" json:"op,omitempty"`
					Value *[]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchName *struct {
					Op    *string   `tfsdk:"op" json:"op,omitempty"`
					Value *[]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"match_name" json:"matchName,omitempty"`
			} `tfsdk:"match_features" json:"matchFeatures,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"feature_group_rules" json:"featureGroupRules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NfdK8SSigsIoNodeFeatureGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_nfd_k8s_sigs_io_node_feature_group_v1alpha1_manifest"
}

func (r *NfdK8SSigsIoNodeFeatureGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeFeatureGroup resource holds Node pools by featureGroup",
		MarkdownDescription: "NodeFeatureGroup resource holds Node pools by featureGroup",
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
				Description:         "Spec defines the rules to be evaluated.",
				MarkdownDescription: "Spec defines the rules to be evaluated.",
				Attributes: map[string]schema.Attribute{
					"feature_group_rules": schema.ListNestedAttribute{
						Description:         "List of rules to evaluate to determine nodes that belong in this group.",
						MarkdownDescription: "List of rules to evaluate to determine nodes that belong in this group.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match_any": schema.ListNestedAttribute{
									Description:         "MatchAny specifies a list of matchers one of which must match.",
									MarkdownDescription: "MatchAny specifies a list of matchers one of which must match.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_features": schema.ListNestedAttribute{
												Description:         "MatchFeatures specifies a set of matcher terms all of which must match.",
												MarkdownDescription: "MatchFeatures specifies a set of matcher terms all of which must match.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"feature": schema.StringAttribute{
															Description:         "Feature is the name of the feature set to match against.",
															MarkdownDescription: "Feature is the name of the feature set to match against.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"match_expressions": schema.SingleNestedAttribute{
															Description:         "MatchExpressions is the set of per-element expressions evaluated. These match against the value of the specified elements.",
															MarkdownDescription: "MatchExpressions is the set of per-element expressions evaluated. These match against the value of the specified elements.",
															Attributes: map[string]schema.Attribute{
																"op": schema.StringAttribute{
																	Description:         "Op is the operator to be applied.",
																	MarkdownDescription: "Op is the operator to be applied.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("In", "NotIn", "InRegexp", "Exists", "DoesNotExist", "Gt", "Lt", "GtLt", "IsTrue", "IsFalse"),
																	},
																},

																"value": schema.ListAttribute{
																	Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
																	MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
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

														"match_name": schema.SingleNestedAttribute{
															Description:         "MatchName in an expression that is matched against the name of each element in the feature set.",
															MarkdownDescription: "MatchName in an expression that is matched against the name of each element in the feature set.",
															Attributes: map[string]schema.Attribute{
																"op": schema.StringAttribute{
																	Description:         "Op is the operator to be applied.",
																	MarkdownDescription: "Op is the operator to be applied.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("In", "NotIn", "InRegexp", "Exists", "DoesNotExist", "Gt", "Lt", "GtLt", "IsTrue", "IsFalse"),
																	},
																},

																"value": schema.ListAttribute{
																	Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
																	MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
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
												},
												Required: true,
												Optional: false,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"match_features": schema.ListNestedAttribute{
									Description:         "MatchFeatures specifies a set of matcher terms all of which must match.",
									MarkdownDescription: "MatchFeatures specifies a set of matcher terms all of which must match.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"feature": schema.StringAttribute{
												Description:         "Feature is the name of the feature set to match against.",
												MarkdownDescription: "Feature is the name of the feature set to match against.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"match_expressions": schema.SingleNestedAttribute{
												Description:         "MatchExpressions is the set of per-element expressions evaluated. These match against the value of the specified elements.",
												MarkdownDescription: "MatchExpressions is the set of per-element expressions evaluated. These match against the value of the specified elements.",
												Attributes: map[string]schema.Attribute{
													"op": schema.StringAttribute{
														Description:         "Op is the operator to be applied.",
														MarkdownDescription: "Op is the operator to be applied.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("In", "NotIn", "InRegexp", "Exists", "DoesNotExist", "Gt", "Lt", "GtLt", "IsTrue", "IsFalse"),
														},
													},

													"value": schema.ListAttribute{
														Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
														MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
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

											"match_name": schema.SingleNestedAttribute{
												Description:         "MatchName in an expression that is matched against the name of each element in the feature set.",
												MarkdownDescription: "MatchName in an expression that is matched against the name of each element in the feature set.",
												Attributes: map[string]schema.Attribute{
													"op": schema.StringAttribute{
														Description:         "Op is the operator to be applied.",
														MarkdownDescription: "Op is the operator to be applied.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("In", "NotIn", "InRegexp", "Exists", "DoesNotExist", "Gt", "Lt", "GtLt", "IsTrue", "IsFalse"),
														},
													},

													"value": schema.ListAttribute{
														Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
														MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the rule.",
									MarkdownDescription: "Name of the rule.",
									Required:            true,
									Optional:            false,
									Computed:            false,
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

func (r *NfdK8SSigsIoNodeFeatureGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_nfd_k8s_sigs_io_node_feature_group_v1alpha1_manifest")

	var model NfdK8SSigsIoNodeFeatureGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("nfd.k8s-sigs.io/v1alpha1")
	model.Kind = pointer.String("NodeFeatureGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
