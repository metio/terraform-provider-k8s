/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package nfd_k8s_sigs_io_v1alpha1

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
	_ datasource.DataSource              = &NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource{}
)

func NewNfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource() datasource.DataSource {
	return &NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource{}
}

type NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_nfd_k8s_sigs_io_node_feature_rule_v1alpha1"
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeFeatureRule resource specifies a configuration for feature-based customization of node objects, such as node labeling.",
		MarkdownDescription: "NodeFeatureRule resource specifies a configuration for feature-based customization of node objects, such as node labeling.",
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
									Optional:            false,
									Computed:            true,
								},

								"labels_template": schema.StringAttribute{
									Description:         "LabelsTemplate specifies a template to expand for dynamically generating multiple labels. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
									MarkdownDescription: "LabelsTemplate specifies a template to expand for dynamically generating multiple labels. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
									Required:            false,
									Optional:            false,
									Computed:            true,
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
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"match_expressions": schema.SingleNestedAttribute{
															Description:         "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
															MarkdownDescription: "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
															Attributes: map[string]schema.Attribute{
																"op": schema.StringAttribute{
																	Description:         "Op is the operator to be applied.",
																	MarkdownDescription: "Op is the operator to be applied.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.ListAttribute{
																	Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
																	MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
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

								"match_features": schema.ListNestedAttribute{
									Description:         "MatchFeatures specifies a set of matcher terms all of which must match.",
									MarkdownDescription: "MatchFeatures specifies a set of matcher terms all of which must match.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"feature": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"match_expressions": schema.SingleNestedAttribute{
												Description:         "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
												MarkdownDescription: "MatchExpressionSet contains a set of MatchExpressions, each of which is evaluated against a set of input values.",
												Attributes: map[string]schema.Attribute{
													"op": schema.StringAttribute{
														Description:         "Op is the operator to be applied.",
														MarkdownDescription: "Op is the operator to be applied.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value": schema.ListAttribute{
														Description:         "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
														MarkdownDescription: "Value is the list of values that the operand evaluates the input against. Value should be empty if the operator is Exists, DoesNotExist, IsTrue or IsFalse. Value should contain exactly one element if the operator is Gt or Lt and exactly two elements if the operator is GtLt. In other cases Value should contain at least one element.",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the rule.",
									MarkdownDescription: "Name of the rule.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"vars": schema.MapAttribute{
									Description:         "Vars is the variables to store if the rule matches. Variables do not directly inflict any changes in the node object. However, they can be referenced from other rules enabling more complex rule hierarchies, without exposing intermediary output values as labels.",
									MarkdownDescription: "Vars is the variables to store if the rule matches. Variables do not directly inflict any changes in the node object. However, they can be referenced from other rules enabling more complex rule hierarchies, without exposing intermediary output values as labels.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"vars_template": schema.StringAttribute{
									Description:         "VarsTemplate specifies a template to expand for dynamically generating multiple variables. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
									MarkdownDescription: "VarsTemplate specifies a template to expand for dynamically generating multiple variables. Data (after template expansion) must be keys with an optional value (<key>[=<value>]) separated by newlines.",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1")

	var data NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "nfd.k8s-sigs.io", Version: "v1alpha1", Resource: "nodefeaturerules"}).
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

	var readResponse NfdK8SSigsIoNodeFeatureRuleV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("nfd.k8s-sigs.io/v1alpha1")
	data.Kind = pointer.String("NodeFeatureRule")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
