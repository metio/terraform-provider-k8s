/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rules_kubeedge_io_v1

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
	_ datasource.DataSource = &RulesKubeedgeIoRuleV1Manifest{}
)

func NewRulesKubeedgeIoRuleV1Manifest() datasource.DataSource {
	return &RulesKubeedgeIoRuleV1Manifest{}
}

type RulesKubeedgeIoRuleV1Manifest struct{}

type RulesKubeedgeIoRuleV1ManifestData struct {
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
		Source         *string            `tfsdk:"source" json:"source,omitempty"`
		SourceResource *map[string]string `tfsdk:"source_resource" json:"sourceResource,omitempty"`
		Target         *string            `tfsdk:"target" json:"target,omitempty"`
		TargetResource *map[string]string `tfsdk:"target_resource" json:"targetResource,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RulesKubeedgeIoRuleV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rules_kubeedge_io_rule_v1_manifest"
}

func (r *RulesKubeedgeIoRuleV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"source": schema.StringAttribute{
						Description:         "source is a string value representing where the messages come from. Itsvalue is the same with ruleendpoint name. For example, my-rest or my-eventbus.",
						MarkdownDescription: "source is a string value representing where the messages come from. Itsvalue is the same with ruleendpoint name. For example, my-rest or my-eventbus.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source_resource": schema.MapAttribute{
						Description:         "sourceResource is a map representing the resource info of source. For restrule-endpoint type its value is {'path':'/test'}. For eventbus ruleendpoint type itsvalue is {'topic':'<user define string>','node_name':'edge-node'}",
						MarkdownDescription: "sourceResource is a map representing the resource info of source. For restrule-endpoint type its value is {'path':'/test'}. For eventbus ruleendpoint type itsvalue is {'topic':'<user define string>','node_name':'edge-node'}",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"target": schema.StringAttribute{
						Description:         "target is a string value representing where the messages go to. its value isthe same with ruleendpoint name. For example, my-eventbus or my-rest or my-servicebus.",
						MarkdownDescription: "target is a string value representing where the messages go to. its value isthe same with ruleendpoint name. For example, my-eventbus or my-rest or my-servicebus.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"target_resource": schema.MapAttribute{
						Description:         "targetResource is a map representing the resource info of target. For restrule-endpoint type its value is {'resource':'http://a.com'}. For eventbus ruleendpointtype its value is {'topic':'/test'}. For servicebus rule-endpoint type its value is{'path':'/request_path'}.",
						MarkdownDescription: "targetResource is a map representing the resource info of target. For restrule-endpoint type its value is {'resource':'http://a.com'}. For eventbus ruleendpointtype its value is {'topic':'/test'}. For servicebus rule-endpoint type its value is{'path':'/request_path'}.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *RulesKubeedgeIoRuleV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rules_kubeedge_io_rule_v1_manifest")

	var model RulesKubeedgeIoRuleV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("rules.kubeedge.io/v1")
	model.Kind = pointer.String("Rule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
