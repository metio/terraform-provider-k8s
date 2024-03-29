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
	_ datasource.DataSource = &AppsKubeblocksIoClusterVersionV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoClusterVersionV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoClusterVersionV1Alpha1Manifest{}
}

type AppsKubeblocksIoClusterVersionV1Alpha1Manifest struct{}

type AppsKubeblocksIoClusterVersionV1Alpha1ManifestData struct {
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
		ClusterDefinitionRef *string `tfsdk:"cluster_definition_ref" json:"clusterDefinitionRef,omitempty"`
		ComponentVersions    *[]struct {
			ComponentDefRef *string `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
			ConfigSpecs     *[]struct {
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
			} `tfsdk:"config_specs" json:"configSpecs,omitempty"`
			SwitchoverSpec *struct {
				CmdExecutorConfig *struct {
					Env   *map[string]string `tfsdk:"env" json:"env,omitempty"`
					Image *string            `tfsdk:"image" json:"image,omitempty"`
				} `tfsdk:"cmd_executor_config" json:"cmdExecutorConfig,omitempty"`
			} `tfsdk:"switchover_spec" json:"switchoverSpec,omitempty"`
			SystemAccountSpec *struct {
				CmdExecutorConfig *struct {
					Env   *map[string]string `tfsdk:"env" json:"env,omitempty"`
					Image *string            `tfsdk:"image" json:"image,omitempty"`
				} `tfsdk:"cmd_executor_config" json:"cmdExecutorConfig,omitempty"`
			} `tfsdk:"system_account_spec" json:"systemAccountSpec,omitempty"`
			VersionsContext *struct {
				Containers     *map[string]string `tfsdk:"containers" json:"containers,omitempty"`
				InitContainers *map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
			} `tfsdk:"versions_context" json:"versionsContext,omitempty"`
		} `tfsdk:"component_versions" json:"componentVersions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoClusterVersionV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_cluster_version_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoClusterVersionV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterVersion is the Schema for the ClusterVersions API",
		MarkdownDescription: "ClusterVersion is the Schema for the ClusterVersions API",
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
				Description:         "ClusterVersionSpec defines the desired state of ClusterVersion",
				MarkdownDescription: "ClusterVersionSpec defines the desired state of ClusterVersion",
				Attributes: map[string]schema.Attribute{
					"cluster_definition_ref": schema.StringAttribute{
						Description:         "Specifies a reference to the ClusterDefinition.",
						MarkdownDescription: "Specifies a reference to the ClusterDefinition.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"component_versions": schema.ListNestedAttribute{
						Description:         "Contains a list of versioning contexts for the components' containers.",
						MarkdownDescription: "Contains a list of versioning contexts for the components' containers.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_def_ref": schema.StringAttribute{
									Description:         "Specifies a reference to one of the cluster component definition names in the ClusterDefinition API (spec.componentDefs.name).",
									MarkdownDescription: "Specifies a reference to one of the cluster component definition names in the ClusterDefinition API (spec.componentDefs.name).",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"config_specs": schema.ListNestedAttribute{
									Description:         "Defines a configuration extension mechanism to handle configuration differences between versions. The configTemplateRefs field, in conjunction with the configTemplateRefs in the ClusterDefinition, determines the final configuration file.",
									MarkdownDescription: "Defines a configuration extension mechanism to handle configuration differences between versions. The configTemplateRefs field, in conjunction with the configTemplateRefs in the ClusterDefinition, determines the final configuration file.",
									NestedObject: schema.NestedAttributeObject{
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"switchover_spec": schema.SingleNestedAttribute{
									Description:         "Defines the images for the component to perform a switchover. This overrides the image and env attributes defined in clusterDefinition.spec.componentDefs.SwitchoverSpec.CommandExecutorEnvItem.",
									MarkdownDescription: "Defines the images for the component to perform a switchover. This overrides the image and env attributes defined in clusterDefinition.spec.componentDefs.SwitchoverSpec.CommandExecutorEnvItem.",
									Attributes: map[string]schema.Attribute{
										"cmd_executor_config": schema.SingleNestedAttribute{
											Description:         "Represents the configuration for the command executor.",
											MarkdownDescription: "Represents the configuration for the command executor.",
											Attributes: map[string]schema.Attribute{
												"env": schema.MapAttribute{
													Description:         "A list of environment variables that will be injected into the command execution context.",
													MarkdownDescription: "A list of environment variables that will be injected into the command execution context.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image": schema.StringAttribute{
													Description:         "Specifies the image used to execute the command.",
													MarkdownDescription: "Specifies the image used to execute the command.",
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

								"system_account_spec": schema.SingleNestedAttribute{
									Description:         "Defines the image for the component to connect to databases or engines. This overrides the 'image' and 'env' attributes defined in clusterDefinition.spec.componentDefs.systemAccountSpec.cmdExecutorConfig. To clear default environment settings, set systemAccountSpec.cmdExecutorConfig.env to an empty list.",
									MarkdownDescription: "Defines the image for the component to connect to databases or engines. This overrides the 'image' and 'env' attributes defined in clusterDefinition.spec.componentDefs.systemAccountSpec.cmdExecutorConfig. To clear default environment settings, set systemAccountSpec.cmdExecutorConfig.env to an empty list.",
									Attributes: map[string]schema.Attribute{
										"cmd_executor_config": schema.SingleNestedAttribute{
											Description:         "Configures the method for obtaining the client SDK and executing statements.",
											MarkdownDescription: "Configures the method for obtaining the client SDK and executing statements.",
											Attributes: map[string]schema.Attribute{
												"env": schema.MapAttribute{
													Description:         "A list of environment variables that will be injected into the command execution context.",
													MarkdownDescription: "A list of environment variables that will be injected into the command execution context.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image": schema.StringAttribute{
													Description:         "Specifies the image used to execute the command.",
													MarkdownDescription: "Specifies the image used to execute the command.",
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

								"versions_context": schema.SingleNestedAttribute{
									Description:         "Defines the context for container images for component versions. This value replaces the values in clusterDefinition.spec.componentDefs.podSpec.[initContainers | containers].",
									MarkdownDescription: "Defines the context for container images for component versions. This value replaces the values in clusterDefinition.spec.componentDefs.podSpec.[initContainers | containers].",
									Attributes: map[string]schema.Attribute{
										"containers": schema.MapAttribute{
											Description:         "Provides override values for ClusterDefinition.spec.componentDefs.podSpec.containers. Typically used in scenarios such as updating application container images.",
											MarkdownDescription: "Provides override values for ClusterDefinition.spec.componentDefs.podSpec.containers. Typically used in scenarios such as updating application container images.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"init_containers": schema.MapAttribute{
											Description:         "Provides override values for ClusterDefinition.spec.componentDefs.podSpec.initContainers. Typically used in scenarios such as updating application container images.",
											MarkdownDescription: "Provides override values for ClusterDefinition.spec.componentDefs.podSpec.initContainers. Typically used in scenarios such as updating application container images.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
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

func (r *AppsKubeblocksIoClusterVersionV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_cluster_version_v1alpha1_manifest")

	var model AppsKubeblocksIoClusterVersionV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("ClusterVersion")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
