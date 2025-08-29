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
	_ datasource.DataSource = &ParametersKubeblocksIoComponentParameterV1Alpha1Manifest{}
)

func NewParametersKubeblocksIoComponentParameterV1Alpha1Manifest() datasource.DataSource {
	return &ParametersKubeblocksIoComponentParameterV1Alpha1Manifest{}
}

type ParametersKubeblocksIoComponentParameterV1Alpha1Manifest struct{}

type ParametersKubeblocksIoComponentParameterV1Alpha1ManifestData struct {
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
		ClusterName       *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ComponentName     *string `tfsdk:"component_name" json:"componentName,omitempty"`
		ConfigItemDetails *[]struct {
			ConfigFileParams *struct {
				Content    *string            `tfsdk:"content" json:"content,omitempty"`
				Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			} `tfsdk:"config_file_params" json:"configFileParams,omitempty"`
			ConfigSpec *struct {
				DefaultMode         *int64  `tfsdk:"default_mode" json:"defaultMode,omitempty"`
				ExternalManaged     *bool   `tfsdk:"external_managed" json:"externalManaged,omitempty"`
				Name                *string `tfsdk:"name" json:"name,omitempty"`
				Namespace           *string `tfsdk:"namespace" json:"namespace,omitempty"`
				RestartOnFileChange *bool   `tfsdk:"restart_on_file_change" json:"restartOnFileChange,omitempty"`
				Template            *string `tfsdk:"template" json:"template,omitempty"`
				VolumeName          *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"config_spec" json:"configSpec,omitempty"`
			Name                *string            `tfsdk:"name" json:"name,omitempty"`
			Payload             *map[string]string `tfsdk:"payload" json:"payload,omitempty"`
			UserConfigTemplates *struct {
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Policy      *string `tfsdk:"policy" json:"policy,omitempty"`
				TemplateRef *string `tfsdk:"template_ref" json:"templateRef,omitempty"`
			} `tfsdk:"user_config_templates" json:"userConfigTemplates,omitempty"`
		} `tfsdk:"config_item_details" json:"configItemDetails,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ParametersKubeblocksIoComponentParameterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_parameters_kubeblocks_io_component_parameter_v1alpha1_manifest"
}

func (r *ParametersKubeblocksIoComponentParameterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ComponentParameter is the Schema for the componentparameters API",
		MarkdownDescription: "ComponentParameter is the Schema for the componentparameters API",
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
				Description:         "ComponentParameterSpec defines the desired state of ComponentConfiguration",
				MarkdownDescription: "ComponentParameterSpec defines the desired state of ComponentConfiguration",
				Attributes: map[string]schema.Attribute{
					"cluster_name": schema.StringAttribute{
						Description:         "Specifies the name of the Cluster that this configuration is associated with.",
						MarkdownDescription: "Specifies the name of the Cluster that this configuration is associated with.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"component_name": schema.StringAttribute{
						Description:         "Represents the name of the Component that this configuration pertains to.",
						MarkdownDescription: "Represents the name of the Component that this configuration pertains to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"config_item_details": schema.ListNestedAttribute{
						Description:         "ConfigItemDetails is an array of ConfigTemplateItemDetail objects. Each ConfigTemplateItemDetail corresponds to a configuration template, which is a ConfigMap that contains multiple configuration files. Each configuration file is stored as a key-value pair within the ConfigMap. The ConfigTemplateItemDetail includes information such as: - The configuration template (a ConfigMap) - The corresponding ConfigConstraint (constraints and validation rules for the configuration) - Volume mounts (for mounting the configuration files)",
						MarkdownDescription: "ConfigItemDetails is an array of ConfigTemplateItemDetail objects. Each ConfigTemplateItemDetail corresponds to a configuration template, which is a ConfigMap that contains multiple configuration files. Each configuration file is stored as a key-value pair within the ConfigMap. The ConfigTemplateItemDetail includes information such as: - The configuration template (a ConfigMap) - The corresponding ConfigConstraint (constraints and validation rules for the configuration) - Volume mounts (for mounting the configuration files)",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_file_params": schema.SingleNestedAttribute{
									Description:         "Specifies the user-defined configuration parameters. When provided, the parameter values in 'configFileParams' override the default configuration parameters. This allows users to override the default configuration according to their specific needs.",
									MarkdownDescription: "Specifies the user-defined configuration parameters. When provided, the parameter values in 'configFileParams' override the default configuration parameters. This allows users to override the default configuration according to their specific needs.",
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Description:         "Holds the configuration keys and values. This field is a workaround for issues found in kubebuilder and code-generator. Refer to https://github.com/kubernetes-sigs/kubebuilder/issues/528 and https://github.com/kubernetes/code-generator/issues/50 for more details. Represents the content of the configuration file.",
											MarkdownDescription: "Holds the configuration keys and values. This field is a workaround for issues found in kubebuilder and code-generator. Refer to https://github.com/kubernetes-sigs/kubebuilder/issues/528 and https://github.com/kubernetes/code-generator/issues/50 for more details. Represents the content of the configuration file.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parameters": schema.MapAttribute{
											Description:         "Represents the updated parameters for a single configuration file.",
											MarkdownDescription: "Represents the updated parameters for a single configuration file.",
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

								"config_spec": schema.SingleNestedAttribute{
									Description:         "Specifies the name of the configuration template (a ConfigMap), ConfigConstraint, and other miscellaneous options. The configuration template is a ConfigMap that contains multiple configuration files. Each configuration file is stored as a key-value pair within the ConfigMap. ConfigConstraint allows defining constraints and validation rules for configuration parameters. It ensures that the configuration adheres to certain requirements and limitations.",
									MarkdownDescription: "Specifies the name of the configuration template (a ConfigMap), ConfigConstraint, and other miscellaneous options. The configuration template is a ConfigMap that contains multiple configuration files. Each configuration file is stored as a key-value pair within the ConfigMap. ConfigConstraint allows defining constraints and validation rules for configuration parameters. It ensures that the configuration adheres to certain requirements and limitations.",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "The operator attempts to set default file permissions (0444). Must be specified as an octal value between 0000 and 0777 (inclusive), or as a decimal value between 0 and 511 (inclusive). YAML supports both octal and decimal values for file permissions. Please note that this setting only affects the permissions of the files themselves. Directories within the specified path are not impacted by this setting. It's important to be aware that this setting might conflict with other options that influence the file mode, such as fsGroup. In such cases, the resulting file mode may have additional bits set. Refers to documents of k8s.ConfigMapVolumeSource.defaultMode for more information.",
											MarkdownDescription: "The operator attempts to set default file permissions (0444). Must be specified as an octal value between 0000 and 0777 (inclusive), or as a decimal value between 0 and 511 (inclusive). YAML supports both octal and decimal values for file permissions. Please note that this setting only affects the permissions of the files themselves. Directories within the specified path are not impacted by this setting. It's important to be aware that this setting might conflict with other options that influence the file mode, such as fsGroup. In such cases, the resulting file mode may have additional bits set. Refers to documents of k8s.ConfigMapVolumeSource.defaultMode for more information.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"external_managed": schema.BoolAttribute{
											Description:         "ExternalManaged indicates whether the configuration is managed by an external system. When set to true, the controller will ignore the management of this configuration.",
											MarkdownDescription: "ExternalManaged indicates whether the configuration is managed by an external system. When set to true, the controller will ignore the management of this configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the template.",
											MarkdownDescription: "Specifies the name of the template.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"namespace": schema.StringAttribute{
											Description:         "Specifies the namespace of the referenced template ConfigMap object.",
											MarkdownDescription: "Specifies the namespace of the referenced template ConfigMap object.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"restart_on_file_change": schema.BoolAttribute{
											Description:         "Specifies whether to restart the pod when the file changes.",
											MarkdownDescription: "Specifies whether to restart the pod when the file changes.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template": schema.StringAttribute{
											Description:         "Specifies the name of the referenced template ConfigMap object.",
											MarkdownDescription: "Specifies the name of the referenced template ConfigMap object.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"volume_name": schema.StringAttribute{
											Description:         "Refers to the volume name of PodTemplate. The file produced through the template will be mounted to the corresponding volume. Must be a DNS_LABEL name. The volume name must be defined in podSpec.containers[*].volumeMounts.",
											MarkdownDescription: "Refers to the volume name of PodTemplate. The file produced through the template will be mounted to the corresponding volume. Must be a DNS_LABEL name. The volume name must be defined in podSpec.containers[*].volumeMounts.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Defines the unique identifier of the configuration template. It must be a string of maximum 63 characters, and can only include lowercase alphanumeric characters, hyphens, and periods. The name must start and end with an alphanumeric character.",
									MarkdownDescription: "Defines the unique identifier of the configuration template. It must be a string of maximum 63 characters, and can only include lowercase alphanumeric characters, hyphens, and periods. The name must start and end with an alphanumeric character.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"payload": schema.MapAttribute{
									Description:         "External controllers can trigger a configuration rerender by modifying this field. Note: Currently, the 'payload' field is opaque and its content is not interpreted by the system. Modifying this field will cause a rerender, regardless of the specific content of this field.",
									MarkdownDescription: "External controllers can trigger a configuration rerender by modifying this field. Note: Currently, the 'payload' field is opaque and its content is not interpreted by the system. Modifying this field will cause a rerender, regardless of the specific content of this field.",
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

func (r *ParametersKubeblocksIoComponentParameterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_parameters_kubeblocks_io_component_parameter_v1alpha1_manifest")

	var model ParametersKubeblocksIoComponentParameterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("parameters.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("ComponentParameter")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
