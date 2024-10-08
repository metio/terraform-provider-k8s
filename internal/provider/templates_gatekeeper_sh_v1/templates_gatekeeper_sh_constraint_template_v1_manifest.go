/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package templates_gatekeeper_sh_v1

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
	_ datasource.DataSource = &TemplatesGatekeeperShConstraintTemplateV1Manifest{}
)

func NewTemplatesGatekeeperShConstraintTemplateV1Manifest() datasource.DataSource {
	return &TemplatesGatekeeperShConstraintTemplateV1Manifest{}
}

type TemplatesGatekeeperShConstraintTemplateV1Manifest struct{}

type TemplatesGatekeeperShConstraintTemplateV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Crd *struct {
			Spec *struct {
				Names *struct {
					Kind       *string   `tfsdk:"kind" json:"kind,omitempty"`
					ShortNames *[]string `tfsdk:"short_names" json:"shortNames,omitempty"`
				} `tfsdk:"names" json:"names,omitempty"`
				Validation *struct {
					LegacySchema    *bool              `tfsdk:"legacy_schema" json:"legacySchema,omitempty"`
					OpenAPIV3Schema *map[string]string `tfsdk:"open_apiv3_schema" json:"openAPIV3Schema,omitempty"`
				} `tfsdk:"validation" json:"validation,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"crd" json:"crd,omitempty"`
		Targets *[]struct {
			Code *[]struct {
				Engine *string            `tfsdk:"engine" json:"engine,omitempty"`
				Source *map[string]string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"code" json:"code,omitempty"`
			Libs   *[]string `tfsdk:"libs" json:"libs,omitempty"`
			Rego   *string   `tfsdk:"rego" json:"rego,omitempty"`
			Target *string   `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"targets" json:"targets,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TemplatesGatekeeperShConstraintTemplateV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_templates_gatekeeper_sh_constraint_template_v1_manifest"
}

func (r *TemplatesGatekeeperShConstraintTemplateV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ConstraintTemplate is the Schema for the constrainttemplates API",
		MarkdownDescription: "ConstraintTemplate is the Schema for the constrainttemplates API",
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
				Description:         "ConstraintTemplateSpec defines the desired state of ConstraintTemplate.",
				MarkdownDescription: "ConstraintTemplateSpec defines the desired state of ConstraintTemplate.",
				Attributes: map[string]schema.Attribute{
					"crd": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"names": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"kind": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"short_names": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"validation": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"legacy_schema": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"open_apiv3_schema": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"targets": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"code": schema.ListNestedAttribute{
									Description:         "The source code options for the constraint template. 'Rego' can only be specified in one place (either here or in the 'rego' field)",
									MarkdownDescription: "The source code options for the constraint template. 'Rego' can only be specified in one place (either here or in the 'rego' field)",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"engine": schema.StringAttribute{
												Description:         "The engine used to evaluate the code. Example: 'Rego'. Required.",
												MarkdownDescription: "The engine used to evaluate the code. Example: 'Rego'. Required.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"source": schema.MapAttribute{
												Description:         "The source code for the template. Required.",
												MarkdownDescription: "The source code for the template. Required.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"libs": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"rego": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target": schema.StringAttribute{
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
	}
}

func (r *TemplatesGatekeeperShConstraintTemplateV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_templates_gatekeeper_sh_constraint_template_v1_manifest")

	var model TemplatesGatekeeperShConstraintTemplateV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("templates.gatekeeper.sh/v1")
	model.Kind = pointer.String("ConstraintTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
