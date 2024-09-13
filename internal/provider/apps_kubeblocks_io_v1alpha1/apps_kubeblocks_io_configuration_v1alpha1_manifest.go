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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AppsKubeblocksIoConfigurationV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoConfigurationV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoConfigurationV1Alpha1Manifest{}
}

type AppsKubeblocksIoConfigurationV1Alpha1Manifest struct{}

type AppsKubeblocksIoConfigurationV1Alpha1ManifestData struct {
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
		ClusterRef        *string `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		ComponentName     *string `tfsdk:"component_name" json:"componentName,omitempty"`
		ConfigItemDetails *[]struct {
			ConfigFileParams *struct {
				Content    *string            `tfsdk:"content" json:"content,omitempty"`
				Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			} `tfsdk:"config_file_params" json:"configFileParams,omitempty"`
			ConfigSpec *struct {
				AsEnvFrom                *[]string `tfsdk:"as_env_from" json:"asEnvFrom,omitempty"`
				ConstraintRef            *string   `tfsdk:"constraint_ref" json:"constraintRef,omitempty"`
				DefaultMode              *int64    `tfsdk:"default_mode" json:"defaultMode,omitempty"`
				InjectEnvTo              *[]string `tfsdk:"inject_env_to" json:"injectEnvTo,omitempty"`
				Keys                     *[]string `tfsdk:"keys" json:"keys,omitempty"`
				LegacyRenderedConfigSpec *struct {
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Policy      *string `tfsdk:"policy" json:"policy,omitempty"`
					TemplateRef *string `tfsdk:"template_ref" json:"templateRef,omitempty"`
				} `tfsdk:"legacy_rendered_config_spec" json:"legacyRenderedConfigSpec,omitempty"`
				Name                  *string   `tfsdk:"name" json:"name,omitempty"`
				Namespace             *string   `tfsdk:"namespace" json:"namespace,omitempty"`
				ReRenderResourceTypes *[]string `tfsdk:"re_render_resource_types" json:"reRenderResourceTypes,omitempty"`
				TemplateRef           *string   `tfsdk:"template_ref" json:"templateRef,omitempty"`
				VolumeName            *string   `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"config_spec" json:"configSpec,omitempty"`
			ImportTemplateRef *struct {
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Policy      *string `tfsdk:"policy" json:"policy,omitempty"`
				TemplateRef *string `tfsdk:"template_ref" json:"templateRef,omitempty"`
			} `tfsdk:"import_template_ref" json:"importTemplateRef,omitempty"`
			Name    *string            `tfsdk:"name" json:"name,omitempty"`
			Payload *map[string]string `tfsdk:"payload" json:"payload,omitempty"`
			Version *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"config_item_details" json:"configItemDetails,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoConfigurationV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_configuration_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoConfigurationV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Configuration represents the complete set of configurations for a specific Component of a Cluster.This includes templates for each configuration file, their corresponding ConfigConstraints, volume mounts,and other relevant details.",
		MarkdownDescription: "Configuration represents the complete set of configurations for a specific Component of a Cluster.This includes templates for each configuration file, their corresponding ConfigConstraints, volume mounts,and other relevant details.",
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
				Description:         "ConfigurationSpec defines the desired state of a Configuration resource.",
				MarkdownDescription: "ConfigurationSpec defines the desired state of a Configuration resource.",
				Attributes: map[string]schema.Attribute{
					"cluster_ref": schema.StringAttribute{
						Description:         "Specifies the name of the Cluster that this configuration is associated with.",
						MarkdownDescription: "Specifies the name of the Cluster that this configuration is associated with.",
						Required:            true,
						Optional:            false,
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
						Description:         "ConfigItemDetails is an array of ConfigurationItemDetail objects.Each ConfigurationItemDetail corresponds to a configuration template,which is a ConfigMap that contains multiple configuration files.Each configuration file is stored as a key-value pair within the ConfigMap.The ConfigurationItemDetail includes information such as:- The configuration template (a ConfigMap)- The corresponding ConfigConstraint (constraints and validation rules for the configuration)- Volume mounts (for mounting the configuration files)",
						MarkdownDescription: "ConfigItemDetails is an array of ConfigurationItemDetail objects.Each ConfigurationItemDetail corresponds to a configuration template,which is a ConfigMap that contains multiple configuration files.Each configuration file is stored as a key-value pair within the ConfigMap.The ConfigurationItemDetail includes information such as:- The configuration template (a ConfigMap)- The corresponding ConfigConstraint (constraints and validation rules for the configuration)- Volume mounts (for mounting the configuration files)",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_file_params": schema.SingleNestedAttribute{
									Description:         "Specifies the user-defined configuration parameters.When provided, the parameter values in 'configFileParams' override the default configuration parameters.This allows users to override the default configuration according to their specific needs.",
									MarkdownDescription: "Specifies the user-defined configuration parameters.When provided, the parameter values in 'configFileParams' override the default configuration parameters.This allows users to override the default configuration according to their specific needs.",
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Description:         "Holds the configuration keys and values. This field is a workaround for issues found in kubebuilder and code-generator.Refer to https://github.com/kubernetes-sigs/kubebuilder/issues/528 and https://github.com/kubernetes/code-generator/issues/50 for more details.Represents the content of the configuration file.",
											MarkdownDescription: "Holds the configuration keys and values. This field is a workaround for issues found in kubebuilder and code-generator.Refer to https://github.com/kubernetes-sigs/kubebuilder/issues/528 and https://github.com/kubernetes/code-generator/issues/50 for more details.Represents the content of the configuration file.",
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
									Description:         "Specifies the name of the configuration template (a ConfigMap), ConfigConstraint, and other miscellaneous options.The configuration template is a ConfigMap that contains multiple configuration files.Each configuration file is stored as a key-value pair within the ConfigMap.ConfigConstraint allows defining constraints and validation rules for configuration parameters.It ensures that the configuration adheres to certain requirements and limitations.",
									MarkdownDescription: "Specifies the name of the configuration template (a ConfigMap), ConfigConstraint, and other miscellaneous options.The configuration template is a ConfigMap that contains multiple configuration files.Each configuration file is stored as a key-value pair within the ConfigMap.ConfigConstraint allows defining constraints and validation rules for configuration parameters.It ensures that the configuration adheres to certain requirements and limitations.",
									Attributes: map[string]schema.Attribute{
										"as_env_from": schema.ListAttribute{
											Description:         "Specifies the containers to inject the ConfigMap parameters as environment variables.This is useful when application images accept parameters through environment variables andgenerate the final configuration file in the startup script based on these variables.This field allows users to specify a list of container names, and KubeBlocks will inject the environmentvariables converted from the ConfigMap into these designated containers. This provides a flexible way topass the configuration items from the ConfigMap to the container without modifying the image.Deprecated: 'asEnvFrom' has been deprecated since 0.9.0 and will be removed in 0.10.0.Use 'injectEnvTo' instead.",
											MarkdownDescription: "Specifies the containers to inject the ConfigMap parameters as environment variables.This is useful when application images accept parameters through environment variables andgenerate the final configuration file in the startup script based on these variables.This field allows users to specify a list of container names, and KubeBlocks will inject the environmentvariables converted from the ConfigMap into these designated containers. This provides a flexible way topass the configuration items from the ConfigMap to the container without modifying the image.Deprecated: 'asEnvFrom' has been deprecated since 0.9.0 and will be removed in 0.10.0.Use 'injectEnvTo' instead.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"constraint_ref": schema.StringAttribute{
											Description:         "Specifies the name of the referenced configuration constraints object.",
											MarkdownDescription: "Specifies the name of the referenced configuration constraints object.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"default_mode": schema.Int64Attribute{
											Description:         "The operator attempts to set default file permissions for scripts (0555) and configurations (0444).However, certain database engines may require different file permissions.You can specify the desired file permissions here.Must be specified as an octal value between 0000 and 0777 (inclusive),or as a decimal value between 0 and 511 (inclusive).YAML supports both octal and decimal values for file permissions.Please note that this setting only affects the permissions of the files themselves.Directories within the specified path are not impacted by this setting.It's important to be aware that this setting might conflict with other optionsthat influence the file mode, such as fsGroup.In such cases, the resulting file mode may have additional bits set.Refers to documents of k8s.ConfigMapVolumeSource.defaultMode for more information.",
											MarkdownDescription: "The operator attempts to set default file permissions for scripts (0555) and configurations (0444).However, certain database engines may require different file permissions.You can specify the desired file permissions here.Must be specified as an octal value between 0000 and 0777 (inclusive),or as a decimal value between 0 and 511 (inclusive).YAML supports both octal and decimal values for file permissions.Please note that this setting only affects the permissions of the files themselves.Directories within the specified path are not impacted by this setting.It's important to be aware that this setting might conflict with other optionsthat influence the file mode, such as fsGroup.In such cases, the resulting file mode may have additional bits set.Refers to documents of k8s.ConfigMapVolumeSource.defaultMode for more information.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"inject_env_to": schema.ListAttribute{
											Description:         "Specifies the containers to inject the ConfigMap parameters as environment variables.This is useful when application images accept parameters through environment variables andgenerate the final configuration file in the startup script based on these variables.This field allows users to specify a list of container names, and KubeBlocks will inject the environmentvariables converted from the ConfigMap into these designated containers. This provides a flexible way topass the configuration items from the ConfigMap to the container without modifying the image.",
											MarkdownDescription: "Specifies the containers to inject the ConfigMap parameters as environment variables.This is useful when application images accept parameters through environment variables andgenerate the final configuration file in the startup script based on these variables.This field allows users to specify a list of container names, and KubeBlocks will inject the environmentvariables converted from the ConfigMap into these designated containers. This provides a flexible way topass the configuration items from the ConfigMap to the container without modifying the image.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keys": schema.ListAttribute{
											Description:         "Specifies the configuration files within the ConfigMap that support dynamic updates.A configuration template (provided in the form of a ConfigMap) may contain templates for multipleconfiguration files.Each configuration file corresponds to a key in the ConfigMap.Some of these configuration files may support dynamic modification and reloading without requiringa pod restart.If empty or omitted, all configuration files in the ConfigMap are assumed to support dynamic updates,and ConfigConstraint applies to all keys.",
											MarkdownDescription: "Specifies the configuration files within the ConfigMap that support dynamic updates.A configuration template (provided in the form of a ConfigMap) may contain templates for multipleconfiguration files.Each configuration file corresponds to a key in the ConfigMap.Some of these configuration files may support dynamic modification and reloading without requiringa pod restart.If empty or omitted, all configuration files in the ConfigMap are assumed to support dynamic updates,and ConfigConstraint applies to all keys.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"legacy_rendered_config_spec": schema.SingleNestedAttribute{
											Description:         "Specifies the secondary rendered config spec for pod-specific customization.The template is rendered inside the pod (by the 'config-manager' sidecar container) and merged with the maintemplate's render result to generate the final configuration file.This field is intended to handle scenarios where different pods within the same Component havevarying configurations. It allows for pod-specific customization of the configuration.Note: This field will be deprecated in future versions, and the functionality will be moved to'cluster.spec.componentSpecs[*].instances[*]'.",
											MarkdownDescription: "Specifies the secondary rendered config spec for pod-specific customization.The template is rendered inside the pod (by the 'config-manager' sidecar container) and merged with the maintemplate's render result to generate the final configuration file.This field is intended to handle scenarios where different pods within the same Component havevarying configurations. It allows for pod-specific customization of the configuration.Note: This field will be deprecated in future versions, and the functionality will be moved to'cluster.spec.componentSpecs[*].instances[*]'.",
											Attributes: map[string]schema.Attribute{
												"namespace": schema.StringAttribute{
													Description:         "Specifies the namespace of the referenced configuration template ConfigMap object.An empty namespace is equivalent to the 'default' namespace.",
													MarkdownDescription: "Specifies the namespace of the referenced configuration template ConfigMap object.An empty namespace is equivalent to the 'default' namespace.",
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

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the configuration template.",
											MarkdownDescription: "Specifies the name of the configuration template.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"namespace": schema.StringAttribute{
											Description:         "Specifies the namespace of the referenced configuration template ConfigMap object.An empty namespace is equivalent to the 'default' namespace.",
											MarkdownDescription: "Specifies the namespace of the referenced configuration template ConfigMap object.An empty namespace is equivalent to the 'default' namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"re_render_resource_types": schema.ListAttribute{
											Description:         "Specifies whether the configuration needs to be re-rendered after v-scale or h-scale operations to reflect changes.In some scenarios, the configuration may need to be updated to reflect the changes in resource allocationor cluster topology. Examples:- Redis: adjust maxmemory after v-scale operation.- MySQL: increase max connections after v-scale operation.- Zookeeper: update zoo.cfg with new node addresses after h-scale operation.",
											MarkdownDescription: "Specifies whether the configuration needs to be re-rendered after v-scale or h-scale operations to reflect changes.In some scenarios, the configuration may need to be updated to reflect the changes in resource allocationor cluster topology. Examples:- Redis: adjust maxmemory after v-scale operation.- MySQL: increase max connections after v-scale operation.- Zookeeper: update zoo.cfg with new node addresses after h-scale operation.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template_ref": schema.StringAttribute{
											Description:         "Specifies the name of the referenced configuration template ConfigMap object.",
											MarkdownDescription: "Specifies the name of the referenced configuration template ConfigMap object.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"volume_name": schema.StringAttribute{
											Description:         "Refers to the volume name of PodTemplate. The configuration file produced through the configurationtemplate will be mounted to the corresponding volume. Must be a DNS_LABEL name.The volume name must be defined in podSpec.containers[*].volumeMounts.",
											MarkdownDescription: "Refers to the volume name of PodTemplate. The configuration file produced through the configurationtemplate will be mounted to the corresponding volume. Must be a DNS_LABEL name.The volume name must be defined in podSpec.containers[*].volumeMounts.",
											Required:            true,
											Optional:            false,
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

								"import_template_ref": schema.SingleNestedAttribute{
									Description:         "Specifies the user-defined configuration template.When provided, the 'importTemplateRef' overrides the default configuration templatespecified in 'configSpec.templateRef'.This allows users to customize the configuration template according to their specific requirements.",
									MarkdownDescription: "Specifies the user-defined configuration template.When provided, the 'importTemplateRef' overrides the default configuration templatespecified in 'configSpec.templateRef'.This allows users to customize the configuration template according to their specific requirements.",
									Attributes: map[string]schema.Attribute{
										"namespace": schema.StringAttribute{
											Description:         "Specifies the namespace of the referenced configuration template ConfigMap object.An empty namespace is equivalent to the 'default' namespace.",
											MarkdownDescription: "Specifies the namespace of the referenced configuration template ConfigMap object.An empty namespace is equivalent to the 'default' namespace.",
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

								"name": schema.StringAttribute{
									Description:         "Defines the unique identifier of the configuration template.It must be a string of maximum 63 characters, and can only include lowercase alphanumeric characters,hyphens, and periods.The name must start and end with an alphanumeric character.",
									MarkdownDescription: "Defines the unique identifier of the configuration template.It must be a string of maximum 63 characters, and can only include lowercase alphanumeric characters,hyphens, and periods.The name must start and end with an alphanumeric character.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"payload": schema.MapAttribute{
									Description:         "External controllers can trigger a configuration rerender by modifying this field.Note: Currently, the 'payload' field is opaque and its content is not interpreted by the system.Modifying this field will cause a rerender, regardless of the specific content of this field.",
									MarkdownDescription: "External controllers can trigger a configuration rerender by modifying this field.Note: Currently, the 'payload' field is opaque and its content is not interpreted by the system.Modifying this field will cause a rerender, regardless of the specific content of this field.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"version": schema.StringAttribute{
									Description:         "Deprecated: No longer used. Please use 'Payload' instead. Previously represented the version of the configuration template.",
									MarkdownDescription: "Deprecated: No longer used. Please use 'Payload' instead. Previously represented the version of the configuration template.",
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

func (r *AppsKubeblocksIoConfigurationV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_configuration_v1alpha1_manifest")

	var model AppsKubeblocksIoConfigurationV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Configuration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
