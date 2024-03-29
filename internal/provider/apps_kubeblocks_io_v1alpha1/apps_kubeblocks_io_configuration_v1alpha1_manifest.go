/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
		Description:         "Configuration is the Schema for the configurations API",
		MarkdownDescription: "Configuration is the Schema for the configurations API",
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
				Description:         "ConfigurationSpec defines the desired state of a Configuration resource.",
				MarkdownDescription: "ConfigurationSpec defines the desired state of a Configuration resource.",
				Attributes: map[string]schema.Attribute{
					"cluster_ref": schema.StringAttribute{
						Description:         "Specifies the name of the cluster that this configuration is associated with.",
						MarkdownDescription: "Specifies the name of the cluster that this configuration is associated with.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"component_name": schema.StringAttribute{
						Description:         "Represents the name of the cluster component that this configuration pertains to.",
						MarkdownDescription: "Represents the name of the cluster component that this configuration pertains to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"config_item_details": schema.ListNestedAttribute{
						Description:         "An array of ConfigurationItemDetail objects that describe user-defined configuration templates.",
						MarkdownDescription: "An array of ConfigurationItemDetail objects that describe user-defined configuration templates.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_file_params": schema.SingleNestedAttribute{
									Description:         "Used to set the parameters to be updated. It is optional.",
									MarkdownDescription: "Used to set the parameters to be updated. It is optional.",
									Attributes: map[string]schema.Attribute{
										"content": schema.StringAttribute{
											Description:         "Holds the configuration keys and values. This field is a workaround for issues found in kubebuilder and code-generator. Refer to https://github.com/kubernetes-sigs/kubebuilder/issues/528 and https://github.com/kubernetes/code-generator/issues/50 for more details.  Represents the content of the configuration file.",
											MarkdownDescription: "Holds the configuration keys and values. This field is a workaround for issues found in kubebuilder and code-generator. Refer to https://github.com/kubernetes-sigs/kubebuilder/issues/528 and https://github.com/kubernetes/code-generator/issues/50 for more details.  Represents the content of the configuration file.",
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
									Description:         "Used to set the configuration template. It is optional.",
									MarkdownDescription: "Used to set the configuration template. It is optional.",
									Attributes: map[string]schema.Attribute{
										"as_env_from": schema.ListAttribute{
											Description:         "An optional field where the list of containers will be injected into EnvFrom.",
											MarkdownDescription: "An optional field where the list of containers will be injected into EnvFrom.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"constraint_ref": schema.StringAttribute{
											Description:         "An optional field that defines the name of the referenced configuration constraints object.",
											MarkdownDescription: "An optional field that defines the name of the referenced configuration constraints object.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"default_mode": schema.Int64Attribute{
											Description:         "Refers to the mode bits used to set permissions on created files by default.  Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.  Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "Refers to the mode bits used to set permissions on created files by default.  Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644.  Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"keys": schema.ListAttribute{
											Description:         "Defines a list of keys. If left empty, ConfigConstraint applies to all keys in the configmap.",
											MarkdownDescription: "Defines a list of keys. If left empty, ConfigConstraint applies to all keys in the configmap.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"legacy_rendered_config_spec": schema.SingleNestedAttribute{
											Description:         "An optional field that defines the secondary rendered config spec.",
											MarkdownDescription: "An optional field that defines the secondary rendered config spec.",
											Attributes: map[string]schema.Attribute{
												"namespace": schema.StringAttribute{
													Description:         "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
													MarkdownDescription: "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(63),
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
														stringvalidator.LengthAtMost(63),
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
											Description:         "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
											MarkdownDescription: "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"re_render_resource_types": schema.ListAttribute{
											Description:         "An optional field defines which resources change trigger re-render config.",
											MarkdownDescription: "An optional field defines which resources change trigger re-render config.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"template_ref": schema.StringAttribute{
											Description:         "Specifies the name of the referenced configuration template ConfigMap object.",
											MarkdownDescription: "Specifies the name of the referenced configuration template ConfigMap object.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"volume_name": schema.StringAttribute{
											Description:         "Refers to the volume name of PodTemplate. The configuration file produced through the configuration template will be mounted to the corresponding volume. Must be a DNS_LABEL name. The volume name must be defined in podSpec.containers[*].volumeMounts.",
											MarkdownDescription: "Refers to the volume name of PodTemplate. The configuration file produced through the configuration template will be mounted to the corresponding volume. Must be a DNS_LABEL name. The volume name must be defined in podSpec.containers[*].volumeMounts.",
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
									Description:         "Specifies the configuration template. It is optional.",
									MarkdownDescription: "Specifies the configuration template. It is optional.",
									Attributes: map[string]schema.Attribute{
										"namespace": schema.StringAttribute{
											Description:         "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
											MarkdownDescription: "Specifies the namespace of the referenced configuration template ConfigMap object. An empty namespace is equivalent to the 'default' namespace.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
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
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
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
									Description:         "Holds the configuration-related rerender. Preserves unknown fields and is optional.",
									MarkdownDescription: "Holds the configuration-related rerender. Preserves unknown fields and is optional.",
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

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
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
