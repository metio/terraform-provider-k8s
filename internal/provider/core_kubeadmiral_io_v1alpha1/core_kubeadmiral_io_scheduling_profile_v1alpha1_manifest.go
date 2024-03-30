/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package core_kubeadmiral_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &CoreKubeadmiralIoSchedulingProfileV1Alpha1Manifest{}
)

func NewCoreKubeadmiralIoSchedulingProfileV1Alpha1Manifest() datasource.DataSource {
	return &CoreKubeadmiralIoSchedulingProfileV1Alpha1Manifest{}
}

type CoreKubeadmiralIoSchedulingProfileV1Alpha1Manifest struct{}

type CoreKubeadmiralIoSchedulingProfileV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		PluginConfig *[]struct {
			Args *map[string]string `tfsdk:"args" json:"args,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"plugin_config" json:"pluginConfig,omitempty"`
		Plugins *struct {
			Filter *struct {
				Disabled *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
					Wait *int64  `tfsdk:"wait" json:"wait,omitempty"`
				} `tfsdk:"disabled" json:"disabled,omitempty"`
				Enabled *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
					Wait *int64  `tfsdk:"wait" json:"wait,omitempty"`
				} `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"filter" json:"filter,omitempty"`
			Score *struct {
				Disabled *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
					Wait *int64  `tfsdk:"wait" json:"wait,omitempty"`
				} `tfsdk:"disabled" json:"disabled,omitempty"`
				Enabled *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
					Wait *int64  `tfsdk:"wait" json:"wait,omitempty"`
				} `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"score" json:"score,omitempty"`
			Select *struct {
				Disabled *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
					Wait *int64  `tfsdk:"wait" json:"wait,omitempty"`
				} `tfsdk:"disabled" json:"disabled,omitempty"`
				Enabled *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
					Wait *int64  `tfsdk:"wait" json:"wait,omitempty"`
				} `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"select" json:"select,omitempty"`
		} `tfsdk:"plugins" json:"plugins,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CoreKubeadmiralIoSchedulingProfileV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_core_kubeadmiral_io_scheduling_profile_v1alpha1_manifest"
}

func (r *CoreKubeadmiralIoSchedulingProfileV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SchedulingProfile configures the plugins to use when scheduling a resource",
		MarkdownDescription: "SchedulingProfile configures the plugins to use when scheduling a resource",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"plugin_config": schema.ListNestedAttribute{
						Description:         "PluginConfig is an optional set of custom plugin arguments for each plugin. Omitting config args for a plugin is equivalent to using the default config for that plugin.",
						MarkdownDescription: "PluginConfig is an optional set of custom plugin arguments for each plugin. Omitting config args for a plugin is equivalent to using the default config for that plugin.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.MapAttribute{
									Description:         "Args defines the arguments passed to the plugins at the time of initialization. Args can have arbitrary structure.",
									MarkdownDescription: "Args defines the arguments passed to the plugins at the time of initialization. Args can have arbitrary structure.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name defines the name of plugin being configured.",
									MarkdownDescription: "Name defines the name of plugin being configured.",
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

					"plugins": schema.SingleNestedAttribute{
						Description:         "Plugins specify the set of plugins that should be enabled or disabled. Enabled plugins are the ones that should be enabled in addition to the default plugins. Disabled plugins are any of the default plugins that should be disabled. When no enabled or disabled plugin is specified for an extension point, default plugins for that extension point will be used if there is any.",
						MarkdownDescription: "Plugins specify the set of plugins that should be enabled or disabled. Enabled plugins are the ones that should be enabled in addition to the default plugins. Disabled plugins are any of the default plugins that should be disabled. When no enabled or disabled plugin is specified for an extension point, default plugins for that extension point will be used if there is any.",
						Attributes: map[string]schema.Attribute{
							"filter": schema.SingleNestedAttribute{
								Description:         "Filter is the list of plugins that should be invoked during the filter phase.",
								MarkdownDescription: "Filter is the list of plugins that should be invoked during the filter phase.",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.ListNestedAttribute{
										Description:         "Disabled specifies default plugins that should be disabled.",
										MarkdownDescription: "Disabled specifies default plugins that should be disabled.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name defines the name of the plugin.",
													MarkdownDescription: "Name defines the name of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													MarkdownDescription: "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Webhook"),
													},
												},

												"wait": schema.Int64Attribute{
													Description:         "Weight defines the weight of the plugin.",
													MarkdownDescription: "Weight defines the weight of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": schema.ListNestedAttribute{
										Description:         "Enabled specifies plugins that should be enabled in addition to the default plugins. Enabled plugins are called in the order specified here, after default plugins. If they need to be invoked before default plugins, default plugins must be disabled and re-enabled here in desired order.",
										MarkdownDescription: "Enabled specifies plugins that should be enabled in addition to the default plugins. Enabled plugins are called in the order specified here, after default plugins. If they need to be invoked before default plugins, default plugins must be disabled and re-enabled here in desired order.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name defines the name of the plugin.",
													MarkdownDescription: "Name defines the name of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													MarkdownDescription: "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Webhook"),
													},
												},

												"wait": schema.Int64Attribute{
													Description:         "Weight defines the weight of the plugin.",
													MarkdownDescription: "Weight defines the weight of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
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

							"score": schema.SingleNestedAttribute{
								Description:         "Score is the list of plugins that should be invoked during the score phase.",
								MarkdownDescription: "Score is the list of plugins that should be invoked during the score phase.",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.ListNestedAttribute{
										Description:         "Disabled specifies default plugins that should be disabled.",
										MarkdownDescription: "Disabled specifies default plugins that should be disabled.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name defines the name of the plugin.",
													MarkdownDescription: "Name defines the name of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													MarkdownDescription: "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Webhook"),
													},
												},

												"wait": schema.Int64Attribute{
													Description:         "Weight defines the weight of the plugin.",
													MarkdownDescription: "Weight defines the weight of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": schema.ListNestedAttribute{
										Description:         "Enabled specifies plugins that should be enabled in addition to the default plugins. Enabled plugins are called in the order specified here, after default plugins. If they need to be invoked before default plugins, default plugins must be disabled and re-enabled here in desired order.",
										MarkdownDescription: "Enabled specifies plugins that should be enabled in addition to the default plugins. Enabled plugins are called in the order specified here, after default plugins. If they need to be invoked before default plugins, default plugins must be disabled and re-enabled here in desired order.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name defines the name of the plugin.",
													MarkdownDescription: "Name defines the name of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													MarkdownDescription: "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Webhook"),
													},
												},

												"wait": schema.Int64Attribute{
													Description:         "Weight defines the weight of the plugin.",
													MarkdownDescription: "Weight defines the weight of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
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

							"select": schema.SingleNestedAttribute{
								Description:         "Select is the list of plugins that should be invoked during the select phase.",
								MarkdownDescription: "Select is the list of plugins that should be invoked during the select phase.",
								Attributes: map[string]schema.Attribute{
									"disabled": schema.ListNestedAttribute{
										Description:         "Disabled specifies default plugins that should be disabled.",
										MarkdownDescription: "Disabled specifies default plugins that should be disabled.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name defines the name of the plugin.",
													MarkdownDescription: "Name defines the name of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													MarkdownDescription: "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Webhook"),
													},
												},

												"wait": schema.Int64Attribute{
													Description:         "Weight defines the weight of the plugin.",
													MarkdownDescription: "Weight defines the weight of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": schema.ListNestedAttribute{
										Description:         "Enabled specifies plugins that should be enabled in addition to the default plugins. Enabled plugins are called in the order specified here, after default plugins. If they need to be invoked before default plugins, default plugins must be disabled and re-enabled here in desired order.",
										MarkdownDescription: "Enabled specifies plugins that should be enabled in addition to the default plugins. Enabled plugins are called in the order specified here, after default plugins. If they need to be invoked before default plugins, default plugins must be disabled and re-enabled here in desired order.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name defines the name of the plugin.",
													MarkdownDescription: "Name defines the name of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													MarkdownDescription: "Type defines the type of the plugin. Type should be omitted when referencing in-tree plugins.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Webhook"),
													},
												},

												"wait": schema.Int64Attribute{
													Description:         "Weight defines the weight of the plugin.",
													MarkdownDescription: "Weight defines the weight of the plugin.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
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
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CoreKubeadmiralIoSchedulingProfileV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_core_kubeadmiral_io_scheduling_profile_v1alpha1_manifest")

	var model CoreKubeadmiralIoSchedulingProfileV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("core.kubeadmiral.io/v1alpha1")
	model.Kind = pointer.String("SchedulingProfile")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
