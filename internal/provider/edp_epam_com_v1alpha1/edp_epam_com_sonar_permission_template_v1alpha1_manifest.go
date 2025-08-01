/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package edp_epam_com_v1alpha1

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
	_ datasource.DataSource = &EdpEpamComSonarPermissionTemplateV1Alpha1Manifest{}
)

func NewEdpEpamComSonarPermissionTemplateV1Alpha1Manifest() datasource.DataSource {
	return &EdpEpamComSonarPermissionTemplateV1Alpha1Manifest{}
}

type EdpEpamComSonarPermissionTemplateV1Alpha1Manifest struct{}

type EdpEpamComSonarPermissionTemplateV1Alpha1ManifestData struct {
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
		Default           *bool                `tfsdk:"default" json:"default,omitempty"`
		Description       *string              `tfsdk:"description" json:"description,omitempty"`
		GroupsPermissions *map[string][]string `tfsdk:"groups_permissions" json:"groupsPermissions,omitempty"`
		Name              *string              `tfsdk:"name" json:"name,omitempty"`
		ProjectKeyPattern *string              `tfsdk:"project_key_pattern" json:"projectKeyPattern,omitempty"`
		SonarRef          *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"sonar_ref" json:"sonarRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EdpEpamComSonarPermissionTemplateV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_edp_epam_com_sonar_permission_template_v1alpha1_manifest"
}

func (r *EdpEpamComSonarPermissionTemplateV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SonarPermissionTemplate is the Schema for the sonar permission template API.",
		MarkdownDescription: "SonarPermissionTemplate is the Schema for the sonar permission template API.",
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
				Description:         "SonarPermissionTemplateSpec defines the desired state of SonarPermissionTemplate.",
				MarkdownDescription: "SonarPermissionTemplateSpec defines the desired state of SonarPermissionTemplate.",
				Attributes: map[string]schema.Attribute{
					"default": schema.BoolAttribute{
						Description:         "Default is a flag to set permission template as default. Only one permission template can be default. If several permission templates have default flag, the random one will be chosen. Default permission template can't be deleted. You need to set another permission template as default before.",
						MarkdownDescription: "Default is a flag to set permission template as default. Only one permission template can be default. If several permission templates have default flag, the random one will be chosen. Default permission template can't be deleted. You need to set another permission template as default before.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "Description of sonar permission template.",
						MarkdownDescription: "Description of sonar permission template.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups_permissions": schema.MapAttribute{
						Description:         "GroupsPermissions is a map of groups and permissions assigned to them.",
						MarkdownDescription: "GroupsPermissions is a map of groups and permissions assigned to them.",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is a name of permission template. Name should be unique across all permission templates. Do not edit this field after creation. Otherwise, the permission template will be recreated.",
						MarkdownDescription: "Name is a name of permission template. Name should be unique across all permission templates. Do not edit this field after creation. Otherwise, the permission template will be recreated.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(100),
						},
					},

					"project_key_pattern": schema.StringAttribute{
						Description:         "ProjectKeyPattern is key pattern. Must be a valid Java regular expression.",
						MarkdownDescription: "ProjectKeyPattern is key pattern. Must be a valid Java regular expression.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sonar_ref": schema.SingleNestedAttribute{
						Description:         "SonarRef is a reference to Sonar custom resource.",
						MarkdownDescription: "SonarRef is a reference to Sonar custom resource.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind specifies the kind of the Sonar resource.",
								MarkdownDescription: "Kind specifies the kind of the Sonar resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the name of the Sonar resource.",
								MarkdownDescription: "Name specifies the name of the Sonar resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
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

func (r *EdpEpamComSonarPermissionTemplateV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_edp_epam_com_sonar_permission_template_v1alpha1_manifest")

	var model EdpEpamComSonarPermissionTemplateV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("edp.epam.com/v1alpha1")
	model.Kind = pointer.String("SonarPermissionTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
