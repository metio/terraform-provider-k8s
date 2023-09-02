/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package expansion_gatekeeper_sh_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ExpansionGatekeeperShExpansionTemplateV1Alpha1Manifest{}
)

func NewExpansionGatekeeperShExpansionTemplateV1Alpha1Manifest() datasource.DataSource {
	return &ExpansionGatekeeperShExpansionTemplateV1Alpha1Manifest{}
}

type ExpansionGatekeeperShExpansionTemplateV1Alpha1Manifest struct{}

type ExpansionGatekeeperShExpansionTemplateV1Alpha1ManifestData struct {
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
		EnforcementAction *string `tfsdk:"enforcement_action" json:"enforcementAction,omitempty"`
		GeneratedGVK      *struct {
			Group   *string `tfsdk:"group" json:"group,omitempty"`
			Kind    *string `tfsdk:"kind" json:"kind,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"generated_gvk" json:"generatedGVK,omitempty"`
		TemplateSource *string `tfsdk:"template_source" json:"templateSource,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExpansionGatekeeperShExpansionTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_expansion_gatekeeper_sh_expansion_template_v1alpha1_manifest"
}

func (r *ExpansionGatekeeperShExpansionTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExpansionTemplate is the Schema for the ExpansionTemplate API.",
		MarkdownDescription: "ExpansionTemplate is the Schema for the ExpansionTemplate API.",
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
				Description:         "ExpansionTemplateSpec defines the desired state of ExpansionTemplate.",
				MarkdownDescription: "ExpansionTemplateSpec defines the desired state of ExpansionTemplate.",
				Attributes: map[string]schema.Attribute{
					"apply_to": schema.ListNestedAttribute{
						Description:         "ApplyTo lists the specific groups, versions and kinds of generator resources which will be expanded.",
						MarkdownDescription: "ApplyTo lists the specific groups, versions and kinds of generator resources which will be expanded.",
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

					"enforcement_action": schema.StringAttribute{
						Description:         "EnforcementAction specifies the enforcement action to be used for resources matching the ExpansionTemplate. Specifying an empty value will use the enforcement action specified by the Constraint in violation.",
						MarkdownDescription: "EnforcementAction specifies the enforcement action to be used for resources matching the ExpansionTemplate. Specifying an empty value will use the enforcement action specified by the Constraint in violation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"generated_gvk": schema.SingleNestedAttribute{
						Description:         "GeneratedGVK specifies the GVK of the resources which the generator resource creates.",
						MarkdownDescription: "GeneratedGVK specifies the GVK of the resources which the generator resource creates.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"template_source": schema.StringAttribute{
						Description:         "TemplateSource specifies the source field on the generator resource to use as the base for expanded resource. For Pod-creating generators, this is usually spec.template",
						MarkdownDescription: "TemplateSource specifies the source field on the generator resource to use as the base for expanded resource. For Pod-creating generators, this is usually spec.template",
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
	}
}

func (r *ExpansionGatekeeperShExpansionTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_expansion_gatekeeper_sh_expansion_template_v1alpha1_manifest")

	var model ExpansionGatekeeperShExpansionTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("expansion.gatekeeper.sh/v1alpha1")
	model.Kind = pointer.String("ExpansionTemplate")

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
