/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package v2_edp_epam_com_v1alpha1

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
	_ datasource.DataSource = &V2EdpEpamComTemplateV1Alpha1Manifest{}
)

func NewV2EdpEpamComTemplateV1Alpha1Manifest() datasource.DataSource {
	return &V2EdpEpamComTemplateV1Alpha1Manifest{}
}

type V2EdpEpamComTemplateV1Alpha1Manifest struct{}

type V2EdpEpamComTemplateV1Alpha1ManifestData struct {
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
		BuildTool   *string `tfsdk:"build_tool" json:"buildTool,omitempty"`
		Category    *string `tfsdk:"category" json:"category,omitempty"`
		Description *string `tfsdk:"description" json:"description,omitempty"`
		DisplayName *string `tfsdk:"display_name" json:"displayName,omitempty"`
		Framework   *string `tfsdk:"framework" json:"framework,omitempty"`
		Icon        *[]struct {
			Base64data *string `tfsdk:"base64data" json:"base64data,omitempty"`
			Mediatype  *string `tfsdk:"mediatype" json:"mediatype,omitempty"`
		} `tfsdk:"icon" json:"icon,omitempty"`
		Keywords    *[]string `tfsdk:"keywords" json:"keywords,omitempty"`
		Language    *string   `tfsdk:"language" json:"language,omitempty"`
		Maintainers *[]struct {
			Email *string `tfsdk:"email" json:"email,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"maintainers" json:"maintainers,omitempty"`
		Maturity      *string `tfsdk:"maturity" json:"maturity,omitempty"`
		MinEDPVersion *string `tfsdk:"min_edp_version" json:"minEDPVersion,omitempty"`
		Source        *string `tfsdk:"source" json:"source,omitempty"`
		Type          *string `tfsdk:"type" json:"type,omitempty"`
		Version       *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *V2EdpEpamComTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_v2_edp_epam_com_template_v1alpha1_manifest"
}

func (r *V2EdpEpamComTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Template is the Schema for the templates API.",
		MarkdownDescription: "Template is the Schema for the templates API.",
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
				Description:         "TemplateSpec defines the desired state of Template.",
				MarkdownDescription: "TemplateSpec defines the desired state of Template.",
				Attributes: map[string]schema.Attribute{
					"build_tool": schema.StringAttribute{
						Description:         "The build tool used to build the component from the template.",
						MarkdownDescription: "The build tool used to build the component from the template.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"category": schema.StringAttribute{
						Description:         "Category is the category of the template.",
						MarkdownDescription: "Category is the category of the template.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "The description of the template.",
						MarkdownDescription: "The description of the template.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"display_name": schema.StringAttribute{
						Description:         "The name of the template.",
						MarkdownDescription: "The name of the template.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"framework": schema.StringAttribute{
						Description:         "The framework used to build the component from the template.",
						MarkdownDescription: "The framework used to build the component from the template.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"icon": schema.ListNestedAttribute{
						Description:         "The icon for this template.",
						MarkdownDescription: "The icon for this template.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"base64data": schema.StringAttribute{
									Description:         "A base64 encoded PNG, JPEG or SVG image.",
									MarkdownDescription: "A base64 encoded PNG, JPEG or SVG image.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"mediatype": schema.StringAttribute{
									Description:         "The media type of the image. E.g image/svg+xml, image/png, image/jpeg.",
									MarkdownDescription: "The media type of the image. E.g image/svg+xml, image/png, image/jpeg.",
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

					"keywords": schema.ListAttribute{
						Description:         "A list of keywords describing the template.",
						MarkdownDescription: "A list of keywords describing the template.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"language": schema.StringAttribute{
						Description:         "The programming language used to build the component from the template.",
						MarkdownDescription: "The programming language used to build the component from the template.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"maintainers": schema.ListNestedAttribute{
						Description:         "A list of organizational entities maintaining the Template.",
						MarkdownDescription: "A list of organizational entities maintaining the Template.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"email": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"maturity": schema.StringAttribute{
						Description:         "The level of maturity the template has achieved at this version. Options include planning, pre-alpha, alpha, beta, stable, mature, inactive, and deprecated.",
						MarkdownDescription: "The level of maturity the template has achieved at this version. Options include planning, pre-alpha, alpha, beta, stable, mature, inactive, and deprecated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("planning", "pre-alpha", "alpha", "beta", "stable", "mature", "inactive", "deprecated"),
						},
					},

					"min_edp_version": schema.StringAttribute{
						Description:         "MinEDPVersion is the minimum EDP version that this template is compatible with.",
						MarkdownDescription: "MinEDPVersion is the minimum EDP version that this template is compatible with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source": schema.StringAttribute{
						Description:         "A repository containing the source code for the template.",
						MarkdownDescription: "A repository containing the source code for the template.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "The type of the template, e.g application, library, autotest, infrastructure, etc.",
						MarkdownDescription: "The type of the template, e.g application, library, autotest, infrastructure, etc.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "Version is the version of the template.",
						MarkdownDescription: "Version is the version of the template.",
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

func (r *V2EdpEpamComTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_v2_edp_epam_com_template_v1alpha1_manifest")

	var model V2EdpEpamComTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("v2.edp.epam.com/v1alpha1")
	model.Kind = pointer.String("Template")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
