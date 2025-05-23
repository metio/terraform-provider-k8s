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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ParametersKubeblocksIoParameterV1Alpha1Manifest{}
)

func NewParametersKubeblocksIoParameterV1Alpha1Manifest() datasource.DataSource {
	return &ParametersKubeblocksIoParameterV1Alpha1Manifest{}
}

type ParametersKubeblocksIoParameterV1Alpha1Manifest struct{}

type ParametersKubeblocksIoParameterV1Alpha1ManifestData struct {
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
		ClusterName         *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ComponentParameters *[]struct {
			ComponentName       *string            `tfsdk:"component_name" json:"componentName,omitempty"`
			Parameters          *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			UserConfigTemplates *struct {
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Policy      *string `tfsdk:"policy" json:"policy,omitempty"`
				TemplateRef *string `tfsdk:"template_ref" json:"templateRef,omitempty"`
			} `tfsdk:"user_config_templates" json:"userConfigTemplates,omitempty"`
		} `tfsdk:"component_parameters" json:"componentParameters,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ParametersKubeblocksIoParameterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_parameters_kubeblocks_io_parameter_v1alpha1_manifest"
}

func (r *ParametersKubeblocksIoParameterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Parameter is the Schema for the parameters API",
		MarkdownDescription: "Parameter is the Schema for the parameters API",
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
				Description:         "ParameterSpec defines the desired state of Parameter",
				MarkdownDescription: "ParameterSpec defines the desired state of Parameter",
				Attributes: map[string]schema.Attribute{
					"cluster_name": schema.StringAttribute{
						Description:         "Specifies the name of the Cluster resource that this operation is targeting.",
						MarkdownDescription: "Specifies the name of the Cluster resource that this operation is targeting.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"component_parameters": schema.ListNestedAttribute{
						Description:         "Lists ComponentParametersSpec objects, each specifying a Component and its parameters and template updates.",
						MarkdownDescription: "Lists ComponentParametersSpec objects, each specifying a Component and its parameters and template updates.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"parameters": schema.MapAttribute{
									Description:         "Specifies the user-defined configuration template or parameters.",
									MarkdownDescription: "Specifies the user-defined configuration template or parameters.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"user_config_templates": schema.SingleNestedAttribute{
									Description:         "Specifies the user-defined configuration template. When provided, the 'importTemplateRef' overrides the default configuration template specified in 'configSpec.templateRef'. This allows users to customize the configuration template according to their specific requirements.",
									MarkdownDescription: "Specifies the user-defined configuration template. When provided, the 'importTemplateRef' overrides the default configuration template specified in 'configSpec.templateRef'. This allows users to customize the configuration template according to their specific requirements.",
									Attributes: map[string]schema.Attribute{
										"namespace": schema.StringAttribute{
											Description:         "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
											MarkdownDescription: "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"policy": schema.StringAttribute{
											Description:         "Defines the strategy for merging externally imported templates into component templates.",
											MarkdownDescription: "Defines the strategy for merging externally imported templates into component templates.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("patch", "replace", "none"),
											},
										},

										"template_ref": schema.StringAttribute{
											Description:         "Specifies the name of the referenced configuration template ConfigMap object.",
											MarkdownDescription: "Specifies the name of the referenced configuration template ConfigMap object.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
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

func (r *ParametersKubeblocksIoParameterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_parameters_kubeblocks_io_parameter_v1alpha1_manifest")

	var model ParametersKubeblocksIoParameterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("parameters.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Parameter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
