/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &AppsKubeblocksIoComponentClassDefinitionV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoComponentClassDefinitionV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoComponentClassDefinitionV1Alpha1Manifest{}
}

type AppsKubeblocksIoComponentClassDefinitionV1Alpha1Manifest struct{}

type AppsKubeblocksIoComponentClassDefinitionV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Groups *[]struct {
			Series *[]struct {
				Classes *[]struct {
					Args   *[]string `tfsdk:"args" json:"args,omitempty"`
					Cpu    *string   `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *string   `tfsdk:"memory" json:"memory,omitempty"`
					Name   *string   `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"classes" json:"classes,omitempty"`
				NamingTemplate *string `tfsdk:"naming_template" json:"namingTemplate,omitempty"`
			} `tfsdk:"series" json:"series,omitempty"`
			Template *string   `tfsdk:"template" json:"template,omitempty"`
			Vars     *[]string `tfsdk:"vars" json:"vars,omitempty"`
		} `tfsdk:"groups" json:"groups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoComponentClassDefinitionV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_component_class_definition_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoComponentClassDefinitionV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ComponentClassDefinition is the Schema for the componentclassdefinitions API",
		MarkdownDescription: "ComponentClassDefinition is the Schema for the componentclassdefinitions API",
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
				Description:         "ComponentClassDefinitionSpec defines the desired state of ComponentClassDefinition",
				MarkdownDescription: "ComponentClassDefinitionSpec defines the desired state of ComponentClassDefinition",
				Attributes: map[string]schema.Attribute{
					"groups": schema.ListNestedAttribute{
						Description:         "group defines a list of class series that conform to the same constraint.",
						MarkdownDescription: "group defines a list of class series that conform to the same constraint.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"series": schema.ListNestedAttribute{
									Description:         "series is a series of class definitions.",
									MarkdownDescription: "series is a series of class definitions.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"classes": schema.ListNestedAttribute{
												Description:         "classes are definitions of classes that come in two forms. In the first form, only ComponentClass.Args need to be defined, and the complete class definition is generated by rendering the ComponentClassGroup.Template and Name. In the second form, the Name, CPU and Memory must be defined.",
												MarkdownDescription: "classes are definitions of classes that come in two forms. In the first form, only ComponentClass.Args need to be defined, and the complete class definition is generated by rendering the ComponentClassGroup.Template and Name. In the second form, the Name, CPU and Memory must be defined.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"args": schema.ListAttribute{
															Description:         "args are variable's value",
															MarkdownDescription: "args are variable's value",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cpu": schema.StringAttribute{
															Description:         "the CPU of the class",
															MarkdownDescription: "the CPU of the class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"memory": schema.StringAttribute{
															Description:         "the memory of the class",
															MarkdownDescription: "the memory of the class",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "name is the class name",
															MarkdownDescription: "name is the class name",
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

											"naming_template": schema.StringAttribute{
												Description:         "namingTemplate is a template that uses the Go template syntax and allows for referencing variables defined in ComponentClassGroup.Template. This enables dynamic generation of class names. For example: name: 'general-{{ .cpu }}c{{ .memory }}g'",
												MarkdownDescription: "namingTemplate is a template that uses the Go template syntax and allows for referencing variables defined in ComponentClassGroup.Template. This enables dynamic generation of class names. For example: name: 'general-{{ .cpu }}c{{ .memory }}g'",
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

								"template": schema.StringAttribute{
									Description:         "template is a class definition template that uses the Go template syntax and allows for variable declaration. When defining a class in Series, specifying the variable's value is sufficient, as the complete class definition will be generated through rendering the template. For example: '''yaml template: | cpu: '{{ or .cpu 1 }}' memory: '{{ or .memory 4 }}Gi' '''",
									MarkdownDescription: "template is a class definition template that uses the Go template syntax and allows for variable declaration. When defining a class in Series, specifying the variable's value is sufficient, as the complete class definition will be generated through rendering the template. For example: '''yaml template: | cpu: '{{ or .cpu 1 }}' memory: '{{ or .memory 4 }}Gi' '''",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"vars": schema.ListAttribute{
									Description:         "vars defines the variables declared in the template and will be used to generating the complete class definition by render the template.",
									MarkdownDescription: "vars defines the variables declared in the template and will be used to generating the complete class definition by render the template.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoComponentClassDefinitionV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_component_class_definition_v1alpha1_manifest")

	var model AppsKubeblocksIoComponentClassDefinitionV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("ComponentClassDefinition")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
