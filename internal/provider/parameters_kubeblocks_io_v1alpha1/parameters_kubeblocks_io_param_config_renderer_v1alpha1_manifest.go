/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package parameters_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &ParametersKubeblocksIoParamConfigRendererV1Alpha1Manifest{}
)

func NewParametersKubeblocksIoParamConfigRendererV1Alpha1Manifest() datasource.DataSource {
	return &ParametersKubeblocksIoParamConfigRendererV1Alpha1Manifest{}
}

type ParametersKubeblocksIoParamConfigRendererV1Alpha1Manifest struct{}

type ParametersKubeblocksIoParamConfigRendererV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ComponentDef *string `tfsdk:"component_def" json:"componentDef,omitempty"`
		Configs      *[]struct {
			FileFormatConfig *struct {
				Format    *string `tfsdk:"format" json:"format,omitempty"`
				IniConfig *struct {
					SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
				} `tfsdk:"ini_config" json:"iniConfig,omitempty"`
			} `tfsdk:"file_format_config" json:"fileFormatConfig,omitempty"`
			Name                  *string   `tfsdk:"name" json:"name,omitempty"`
			ReRenderResourceTypes *[]string `tfsdk:"re_render_resource_types" json:"reRenderResourceTypes,omitempty"`
			TemplateName          *string   `tfsdk:"template_name" json:"templateName,omitempty"`
		} `tfsdk:"configs" json:"configs,omitempty"`
		ParametersDefs *[]string `tfsdk:"parameters_defs" json:"parametersDefs,omitempty"`
		ServiceVersion *string   `tfsdk:"service_version" json:"serviceVersion,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ParametersKubeblocksIoParamConfigRendererV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_parameters_kubeblocks_io_param_config_renderer_v1alpha1_manifest"
}

func (r *ParametersKubeblocksIoParamConfigRendererV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ParamConfigRenderer is the Schema for the paramconfigrenderers API",
		MarkdownDescription: "ParamConfigRenderer is the Schema for the paramconfigrenderers API",
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
				Description:         "ParamConfigRendererSpec defines the desired state of ParamConfigRenderer",
				MarkdownDescription: "ParamConfigRendererSpec defines the desired state of ParamConfigRenderer",
				Attributes: map[string]schema.Attribute{
					"component_def": schema.StringAttribute{
						Description:         "Specifies the ComponentDefinition custom resource (CR) that defines the Component's characteristics and behavior.",
						MarkdownDescription: "Specifies the ComponentDefinition custom resource (CR) that defines the Component's characteristics and behavior.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"configs": schema.ListNestedAttribute{
						Description:         "Specifies the configuration files.",
						MarkdownDescription: "Specifies the configuration files.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"file_format_config": schema.SingleNestedAttribute{
									Description:         "Specifies the format of the configuration file and any associated parameters that are specific to the chosen format. Supported formats include 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties', and 'toml'. Each format may have its own set of parameters that can be configured. For instance, when using the 'ini' format, you can specify the section name. Example: ''' fileFormatConfig: format: ini iniConfig: sectionName: mysqld '''",
									MarkdownDescription: "Specifies the format of the configuration file and any associated parameters that are specific to the chosen format. Supported formats include 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties', and 'toml'. Each format may have its own set of parameters that can be configured. For instance, when using the 'ini' format, you can specify the section name. Example: ''' fileFormatConfig: format: ini iniConfig: sectionName: mysqld '''",
									Attributes: map[string]schema.Attribute{
										"format": schema.StringAttribute{
											Description:         "The config file format. Valid values are 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties' and 'toml'. Each format has its own characteristics and use cases. - ini: is a text-based content with a structure and syntax comprising key–value pairs for properties, reference wiki: https://en.wikipedia.org/wiki/INI_file - xml: refers to wiki: https://en.wikipedia.org/wiki/XML - yaml: supports for complex data types and structures. - json: refers to wiki: https://en.wikipedia.org/wiki/JSON - hcl: The HashiCorp Configuration Language (HCL) is a configuration language authored by HashiCorp, reference url: https://www.linode.com/docs/guides/introduction-to-hcl/ - dotenv: is a plain text file with simple key–value pairs, reference wiki: https://en.wikipedia.org/wiki/Configuration_file#MS-DOS - properties: a file extension mainly used in Java, reference wiki: https://en.wikipedia.org/wiki/.properties - toml: refers to wiki: https://en.wikipedia.org/wiki/TOML - props-plus: a file extension mainly used in Java, supports CamelCase(e.g: brokerMaxConnectionsPerIp)",
											MarkdownDescription: "The config file format. Valid values are 'ini', 'xml', 'yaml', 'json', 'hcl', 'dotenv', 'properties' and 'toml'. Each format has its own characteristics and use cases. - ini: is a text-based content with a structure and syntax comprising key–value pairs for properties, reference wiki: https://en.wikipedia.org/wiki/INI_file - xml: refers to wiki: https://en.wikipedia.org/wiki/XML - yaml: supports for complex data types and structures. - json: refers to wiki: https://en.wikipedia.org/wiki/JSON - hcl: The HashiCorp Configuration Language (HCL) is a configuration language authored by HashiCorp, reference url: https://www.linode.com/docs/guides/introduction-to-hcl/ - dotenv: is a plain text file with simple key–value pairs, reference wiki: https://en.wikipedia.org/wiki/Configuration_file#MS-DOS - properties: a file extension mainly used in Java, reference wiki: https://en.wikipedia.org/wiki/.properties - toml: refers to wiki: https://en.wikipedia.org/wiki/TOML - props-plus: a file extension mainly used in Java, supports CamelCase(e.g: brokerMaxConnectionsPerIp)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("xml", "ini", "yaml", "json", "hcl", "dotenv", "toml", "properties", "redis", "props-plus"),
											},
										},

										"ini_config": schema.SingleNestedAttribute{
											Description:         "Holds options specific to the 'ini' file format.",
											MarkdownDescription: "Holds options specific to the 'ini' file format.",
											Attributes: map[string]schema.Attribute{
												"section_name": schema.StringAttribute{
													Description:         "A string that describes the name of the ini section.",
													MarkdownDescription: "A string that describes the name of the ini section.",
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

								"name": schema.StringAttribute{
									Description:         "Specifies the config file name in the config template.",
									MarkdownDescription: "Specifies the config file name in the config template.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"re_render_resource_types": schema.ListAttribute{
									Description:         "Specifies whether the configuration needs to be re-rendered after v-scale or h-scale operations to reflect changes. In some scenarios, the configuration may need to be updated to reflect the changes in resource allocation or cluster topology. Examples: - Redis: adjust maxmemory after v-scale operation. - MySQL: increase max connections after v-scale operation. - Zookeeper: update zoo.cfg with new node addresses after h-scale operation.",
									MarkdownDescription: "Specifies whether the configuration needs to be re-rendered after v-scale or h-scale operations to reflect changes. In some scenarios, the configuration may need to be updated to reflect the changes in resource allocation or cluster topology. Examples: - Redis: adjust maxmemory after v-scale operation. - MySQL: increase max connections after v-scale operation. - Zookeeper: update zoo.cfg with new node addresses after h-scale operation.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"template_name": schema.StringAttribute{
									Description:         "Specifies the name of the referenced componentTemplateSpec.",
									MarkdownDescription: "Specifies the name of the referenced componentTemplateSpec.",
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

					"parameters_defs": schema.ListAttribute{
						Description:         "Specifies the ParametersDefinition custom resource (CR) that defines the Component parameter's schema and behavior.",
						MarkdownDescription: "Specifies the ParametersDefinition custom resource (CR) that defines the Component parameter's schema and behavior.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_version": schema.StringAttribute{
						Description:         "ServiceVersion specifies the version of the Service expected to be provisioned by this Component. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If no version is specified, the latest available version will be used.",
						MarkdownDescription: "ServiceVersion specifies the version of the Service expected to be provisioned by this Component. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If no version is specified, the latest available version will be used.",
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

func (r *ParametersKubeblocksIoParamConfigRendererV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_parameters_kubeblocks_io_param_config_renderer_v1alpha1_manifest")

	var model ParametersKubeblocksIoParamConfigRendererV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("parameters.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("ParamConfigRenderer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
