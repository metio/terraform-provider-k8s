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
	_ datasource.DataSource = &NfdK8SSigsIoNodeFeatureRuleV1Alpha1Manifest{}
)

func NewNfdK8SSigsIoNodeFeatureRuleV1Alpha1Manifest() datasource.DataSource {
	return &NfdK8SSigsIoNodeFeatureRuleV1Alpha1Manifest{}
}

type NfdK8SSigsIoNodeFeatureRuleV1Alpha1Manifest struct{}

type NfdK8SSigsIoNodeFeatureRuleV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Rules *[]struct {
			Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			LabelsTemplate *string            `tfsdk:"labels_template" json:"labelsTemplate,omitempty"`
			MatchAny       *[]struct {
				MatchFeatures *[]struct {
					Feature          *string `tfsdk:"feature" json:"feature,omitempty"`
					MatchExpressions *struct {
						Op    *string   `tfsdk:"op" json:"op,omitempty"`
						Value *[]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				} `tfsdk:"match_features" json:"matchFeatures,omitempty"`
			} `tfsdk:"match_any" json:"matchAny,omitempty"`
			MatchFeatures *[]struct {
				Feature          *string `tfsdk:"feature" json:"feature,omitempty"`
				MatchExpressions *struct {
					Op    *string   `tfsdk:"op" json:"op,omitempty"`
					Value *[]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			} `tfsdk:"match_features" json:"matchFeatures,omitempty"`
			Name         *string            `tfsdk:"name" json:"name,omitempty"`
			Vars         *map[string]string `tfsdk:"vars" json:"vars,omitempty"`
			VarsTemplate *string            `tfsdk:"vars_template" json:"varsTemplate,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest"
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeFeatureRule resource specifies a configuration for feature-based customization of node objects, such as node labeling.",
		MarkdownDescription: "NodeFeatureRule resource specifies a configuration for feature-based customization of node objects, such as node labeling.",
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
				Description:         "NodeFeatureRuleSpec describes a NodeFeatureRule.",
				MarkdownDescription: "NodeFeatureRuleSpec describes a NodeFeatureRule.",
				Attributes: map[string]schema.Attribute{
					"rules": schema.ListNestedAttribute{
						Description:         "Rules is a list of node customization rules.",
						MarkdownDescription: "Rules is a list of node customization rules.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"labels": schema.MapAttribute{
									Description:         "Labels to create if the rule matches.",
									MarkdownDescription: "Labels to create if the rule matches.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels_template": schema.StringAttribute{
									Description:         "LabelsTemplate specifies a template to expand for dynamically generating multiple labels. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
									MarkdownDescription: "LabelsTemplate specifies a template to expand for dynamically generating multiple labels. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

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
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"match_expressions": schema.SingleNestedAttribute{
															Description:         "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
															MarkdownDescription: "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
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
															Required: true,
															Optional: false,
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
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"match_expressions": schema.SingleNestedAttribute{
												Description:         "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
												MarkdownDescription: "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
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

								"name": schema.StringAttribute{
									Description:         "Name of the rule.",
									MarkdownDescription: "Name of the rule.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"vars": schema.MapAttribute{
									Description:         "Vars is the variables to store if the rule matches. Variables do not directly inflict any changes in the node object. However, they can be referenced from other rules enabling more complex rule hierarchies, without exposing intermediary output values as labels.",
									MarkdownDescription: "Vars is the variables to store if the rule matches. Variables do not directly inflict any changes in the node object. However, they can be referenced from other rules enabling more complex rule hierarchies, without exposing intermediary output values as labels.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"vars_template": schema.StringAttribute{
									Description:         "VarsTemplate specifies a template to expand for dynamically generating multiple variables. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
									MarkdownDescription: "VarsTemplate specifies a template to expand for dynamically generating multiple variables. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
									Required:            false,
									Optional:            true,
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

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1_manifest")

	var model NfdK8SSigsIoNodeFeatureRuleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("nfd.k8s-sigs.io/v1alpha1")
	model.Kind = pointer.String("NodeFeatureRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
