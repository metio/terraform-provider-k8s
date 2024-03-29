/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package api_kubemod_io_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &ApiKubemodIoModRuleV1Beta1Manifest{}
)

func NewApiKubemodIoModRuleV1Beta1Manifest() datasource.DataSource {
	return &ApiKubemodIoModRuleV1Beta1Manifest{}
}

type ApiKubemodIoModRuleV1Beta1Manifest struct{}

type ApiKubemodIoModRuleV1Beta1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		AdmissionOperations *[]string `tfsdk:"admission_operations" json:"admissionOperations,omitempty"`
		ExecutionTier       *int64    `tfsdk:"execution_tier" json:"executionTier,omitempty"`
		Match               *[]struct {
			MatchFor    *string   `tfsdk:"match_for" json:"matchFor,omitempty"`
			MatchRegex  *string   `tfsdk:"match_regex" json:"matchRegex,omitempty"`
			MatchValue  *string   `tfsdk:"match_value" json:"matchValue,omitempty"`
			MatchValues *[]string `tfsdk:"match_values" json:"matchValues,omitempty"`
			Negate      *bool     `tfsdk:"negate" json:"negate,omitempty"`
			Select      *string   `tfsdk:"select" json:"select,omitempty"`
		} `tfsdk:"match" json:"match,omitempty"`
		Patch *[]struct {
			Op     *string `tfsdk:"op" json:"op,omitempty"`
			Path   *string `tfsdk:"path" json:"path,omitempty"`
			Select *string `tfsdk:"select" json:"select,omitempty"`
			Value  *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"patch" json:"patch,omitempty"`
		RejectMessage        *string `tfsdk:"reject_message" json:"rejectMessage,omitempty"`
		TargetNamespaceRegex *string `tfsdk:"target_namespace_regex" json:"targetNamespaceRegex,omitempty"`
		Type                 *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApiKubemodIoModRuleV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_api_kubemod_io_mod_rule_v1beta1_manifest"
}

func (r *ApiKubemodIoModRuleV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ModRule is the Schema for the modrules API",
		MarkdownDescription: "ModRule is the Schema for the modrules API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "ModRuleSpec defines the desired state of ModRule",
				MarkdownDescription: "ModRuleSpec defines the desired state of ModRule",
				Attributes: map[string]schema.Attribute{
					"admission_operations": schema.ListAttribute{
						Description:         "AdmissionOperations specifies which admission hook operations this ModRule applies to. Valid values are: - 'CREATE' - the rule applies to all matching resources as they are created. - 'UPDATE' - the rule applies to all matching resources as they are updated. - 'DELETE' - the rule applies to all matching resources as they are deleted. By default, a ModRule applies to all admission operations.",
						MarkdownDescription: "AdmissionOperations specifies which admission hook operations this ModRule applies to. Valid values are: - 'CREATE' - the rule applies to all matching resources as they are created. - 'UPDATE' - the rule applies to all matching resources as they are updated. - 'DELETE' - the rule applies to all matching resources as they are deleted. By default, a ModRule applies to all admission operations.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"execution_tier": schema.Int64Attribute{
						Description:         "ExecutionTier is a value between -32767 and 32766. ExecutionTier controls when this ModRule will be executed as it relates to the other ModRules loaded in the system. ModRules are matched and executed in tiers, starting with the lowest tier. The results of executing all ModRules in a tier are passed as input to the ModRules in the next tier. This cascading execution continues until the highest tier of ModRules has been executed. ModRules in the same tier are executed in indeterminate order.",
						MarkdownDescription: "ExecutionTier is a value between -32767 and 32766. ExecutionTier controls when this ModRule will be executed as it relates to the other ModRules loaded in the system. ModRules are matched and executed in tiers, starting with the lowest tier. The results of executing all ModRules in a tier are passed as input to the ModRules in the next tier. This cascading execution continues until the highest tier of ModRules has been executed. ModRules in the same tier are executed in indeterminate order.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"match": schema.ListNestedAttribute{
						Description:         "Match is a list of match items which consist of select queries and expected match values or regular expressions. When all match items for an object are positive, the rule is in effect.",
						MarkdownDescription: "Match is a list of match items which consist of select queries and expected match values or regular expressions. When all match items for an object are positive, the rule is in effect.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match_for": schema.StringAttribute{
									Description:         "MatchFor instructs how to match the results against the match... requirements. Valid values are: - 'Any' - the match is considered positive if any of the results of select have a match. - 'All' - the match is considered positive only if all of the results of select have a match.",
									MarkdownDescription: "MatchFor instructs how to match the results against the match... requirements. Valid values are: - 'Any' - the match is considered positive if any of the results of select have a match. - 'All' - the match is considered positive only if all of the results of select have a match.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Any", "All"),
									},
								},

								"match_regex": schema.StringAttribute{
									Description:         "MatchRegex specifies the regular expression to compare the result of Select by. The match is considered positive if at least one of the results of evaluating the select query yields a match when compared to value.",
									MarkdownDescription: "MatchRegex specifies the regular expression to compare the result of Select by. The match is considered positive if at least one of the results of evaluating the select query yields a match when compared to value.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_value": schema.StringAttribute{
									Description:         "MatchValue specifies the exact value to match the result of Select by. The match is considered positive if at least one of the results of evaluating the select query yields a match when compared to matchValue.",
									MarkdownDescription: "MatchValue specifies the exact value to match the result of Select by. The match is considered positive if at least one of the results of evaluating the select query yields a match when compared to matchValue.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match_values": schema.ListAttribute{
									Description:         "MatchValues specifies a list of values to match the result of Select by. The match is considered positive if at least one of the results of evaluating the select query yields a match when compared to any of the values in the array.",
									MarkdownDescription: "MatchValues specifies a list of values to match the result of Select by. The match is considered positive if at least one of the results of evaluating the select query yields a match when compared to any of the values in the array.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"negate": schema.BoolAttribute{
									Description:         "Negate indicates whether the match result should be to inverted. Defaults to false.",
									MarkdownDescription: "Negate indicates whether the match result should be to inverted. Defaults to false.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"select": schema.StringAttribute{
									Description:         "Select is a JSONPath query expression: https://goessner.net/articles/JsonPath/ which yields zero or more values. If no match value or regex is specified, if the query yields a non-empty result, the match is considered positive.",
									MarkdownDescription: "Select is a JSONPath query expression: https://goessner.net/articles/JsonPath/ which yields zero or more values. If no match value or regex is specified, if the query yields a non-empty result, the match is considered positive.",
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

					"patch": schema.ListNestedAttribute{
						Description:         "Patch is a list of patch operations to perform on the matching resources at the time of creation. The value part of a patch operation can be a golang template which accepts the resource as its context. This field must be provided for ModRules of type 'patch'",
						MarkdownDescription: "Patch is a list of patch operations to perform on the matching resources at the time of creation. The value part of a patch operation can be a golang template which accepts the resource as its context. This field must be provided for ModRules of type 'patch'",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"op": schema.StringAttribute{
									Description:         "Operation is the type of JSON Path operation to perform against the target element.",
									MarkdownDescription: "Operation is the type of JSON Path operation to perform against the target element.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("add", "replace", "remove"),
									},
								},

								"path": schema.StringAttribute{
									Description:         "Path is the JSON path to the target element.",
									MarkdownDescription: "Path is the JSON path to the target element.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"select": schema.StringAttribute{
									Description:         "Optional JSONPath query expression: https://goessner.net/articles/JsonPath/ used to construct path. A patch operation is created for each result of the query. A placeholder is created for each wildcard and filter in the expression. These placeholders can be used when constructing 'path'. For example, if select is '$.spec.containers[*].ports[?@.containerPort == 80]' placeholder #0 will point to the index of 'containers' and #1 will point to the index of 'ports'. This allows us to define paths such as '/spec/template/spec/containers/#0/securityContext'",
									MarkdownDescription: "Optional JSONPath query expression: https://goessner.net/articles/JsonPath/ used to construct path. A patch operation is created for each result of the query. A placeholder is created for each wildcard and filter in the expression. These placeholders can be used when constructing 'path'. For example, if select is '$.spec.containers[*].ports[?@.containerPort == 80]' placeholder #0 will point to the index of 'containers' and #1 will point to the index of 'ports'. This allows us to define paths such as '/spec/template/spec/containers/#0/securityContext'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the JSON representation of the modification. The value is a golang template which is evaluated against the context of the target resource. KubeMod performs some analysis of the result of the template evaluation in order to infer its JSON type: - If the value matches the format of a JavaScript number, it is considered to be a number. - If the value matches a boolean literal (true/false), it is considered to be a boolean literal. - If the value matches 'null', it is considered to be null. - If the value is surrounded by double-quotes, it is considered to be a string. - If the value is surrounded by brackets, it is considered to be a JSON array. - If the value is surrounded by curly braces, it is considered to be a JSON object. - If none of the above is true, the value is considered to be a string.",
									MarkdownDescription: "Value is the JSON representation of the modification. The value is a golang template which is evaluated against the context of the target resource. KubeMod performs some analysis of the result of the template evaluation in order to infer its JSON type: - If the value matches the format of a JavaScript number, it is considered to be a number. - If the value matches a boolean literal (true/false), it is considered to be a boolean literal. - If the value matches 'null', it is considered to be null. - If the value is surrounded by double-quotes, it is considered to be a string. - If the value is surrounded by brackets, it is considered to be a JSON array. - If the value is surrounded by curly braces, it is considered to be a JSON object. - If none of the above is true, the value is considered to be a string.",
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

					"reject_message": schema.StringAttribute{
						Description:         "RejectMessage is an optional message displayed when a resource is rejected by a Reject ModRule. The field is a Golang template evaluated in the context of the object being rejected.",
						MarkdownDescription: "RejectMessage is an optional message displayed when a resource is rejected by a Reject ModRule. The field is a Golang template evaluated in the context of the object being rejected.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_namespace_regex": schema.StringAttribute{
						Description:         "TargetNamespaceRegex is optional and only applies to ModRules in 'kubemod-system' namespace. Its usage enables cluster-wide matching of namespaced resources.",
						MarkdownDescription: "TargetNamespaceRegex is optional and only applies to ModRules in 'kubemod-system' namespace. Its usage enables cluster-wide matching of namespaced resources.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type describes the type of a ModRule. Valid values are: - 'Patch' - the rule performs modifications on all the matching resources as they are created. - 'Reject' - the rule rejects the creation of all matching resources.",
						MarkdownDescription: "Type describes the type of a ModRule. Valid values are: - 'Patch' - the rule performs modifications on all the matching resources as they are created. - 'Reject' - the rule rejects the creation of all matching resources.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Patch", "Reject"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ApiKubemodIoModRuleV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_api_kubemod_io_mod_rule_v1beta1_manifest")

	var model ApiKubemodIoModRuleV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("api.kubemod.io/v1beta1")
	model.Kind = pointer.String("ModRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
