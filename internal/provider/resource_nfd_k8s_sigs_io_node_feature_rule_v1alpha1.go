/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource)(nil)
)

type NfdK8SSigsIoNodeFeatureRuleV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NfdK8SSigsIoNodeFeatureRuleV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Rules *[]struct {
			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			LabelsTemplate *string `tfsdk:"labels_template" yaml:"labelsTemplate,omitempty"`

			MatchAny *[]struct {
				MatchFeatures *[]struct {
					Feature *string `tfsdk:"feature" yaml:"feature,omitempty"`

					MatchExpressions *struct {
						Op *string `tfsdk:"op" yaml:"op,omitempty"`

						Value *[]string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
				} `tfsdk:"match_features" yaml:"matchFeatures,omitempty"`
			} `tfsdk:"match_any" yaml:"matchAny,omitempty"`

			MatchFeatures *[]struct {
				Feature *string `tfsdk:"feature" yaml:"feature,omitempty"`

				MatchExpressions *struct {
					Op *string `tfsdk:"op" yaml:"op,omitempty"`

					Value *[]string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
			} `tfsdk:"match_features" yaml:"matchFeatures,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Vars *map[string]string `tfsdk:"vars" yaml:"vars,omitempty"`

			VarsTemplate *string `tfsdk:"vars_template" yaml:"varsTemplate,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource() resource.Resource {
	return &NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource{}
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nfd_k8s_sigs_io_node_feature_rule_v1alpha1"
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "NodeFeatureRule resource specifies a configuration for feature-based customization of node objects, such as node labeling.",
		MarkdownDescription: "NodeFeatureRule resource specifies a configuration for feature-based customization of node objects, such as node labeling.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "NodeFeatureRuleSpec describes a NodeFeatureRule.",
				MarkdownDescription: "NodeFeatureRuleSpec describes a NodeFeatureRule.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"rules": {
						Description:         "Rules is a list of node customization rules.",
						MarkdownDescription: "Rules is a list of node customization rules.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"labels": {
								Description:         "Labels to create if the rule matches.",
								MarkdownDescription: "Labels to create if the rule matches.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels_template": {
								Description:         "LabelsTemplate specifies a template to expand for dynamically generating multiple labels. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
								MarkdownDescription: "LabelsTemplate specifies a template to expand for dynamically generating multiple labels. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_any": {
								Description:         "MatchAny specifies a list of matchers one of which must match.",
								MarkdownDescription: "MatchAny specifies a list of matchers one of which must match.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_features": {
										Description:         "MatchFeatures specifies a set of matcher terms all of which must match.",
										MarkdownDescription: "MatchFeatures specifies a set of matcher terms all of which must match.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"feature": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"match_expressions": {
												Description:         "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
												MarkdownDescription: "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"op": {
														Description:         "Op is the operator to be applied.",
														MarkdownDescription: "Op is the operator to be applied.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("In", "NotIn", "InRegexp", "Exists", "DoesNotExist", "Gt", "Lt", "GtLt", "IsTrue", "IsFalse"),
														},
													},

													"value": {
														Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
														MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_features": {
								Description:         "MatchFeatures specifies a set of matcher terms all of which must match.",
								MarkdownDescription: "MatchFeatures specifies a set of matcher terms all of which must match.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"feature": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"match_expressions": {
										Description:         "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
										MarkdownDescription: "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"op": {
												Description:         "Op is the operator to be applied.",
												MarkdownDescription: "Op is the operator to be applied.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "InRegexp", "Exists", "DoesNotExist", "Gt", "Lt", "GtLt", "IsTrue", "IsFalse"),
												},
											},

											"value": {
												Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
												MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the rule.",
								MarkdownDescription: "Name of the rule.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"vars": {
								Description:         "Vars is the variables to store if the rule matches. Variables do not directly inflict any changes in the node object. However, they can be referenced from other rules enabling more complex rule hierarchies, without exposing intermediary output values as labels.",
								MarkdownDescription: "Vars is the variables to store if the rule matches. Variables do not directly inflict any changes in the node object. However, they can be referenced from other rules enabling more complex rule hierarchies, without exposing intermediary output values as labels.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vars_template": {
								Description:         "VarsTemplate specifies a template to expand for dynamically generating multiple variables. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
								MarkdownDescription: "VarsTemplate specifies a template to expand for dynamically generating multiple variables. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1")

	var state NfdK8SSigsIoNodeFeatureRuleV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NfdK8SSigsIoNodeFeatureRuleV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("nfd.k8s-sigs.io/v1alpha1")
	goModel.Kind = utilities.Ptr("NodeFeatureRule")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1")

	var state NfdK8SSigsIoNodeFeatureRuleV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NfdK8SSigsIoNodeFeatureRuleV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("nfd.k8s-sigs.io/v1alpha1")
	goModel.Kind = utilities.Ptr("NodeFeatureRule")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
