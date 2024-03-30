/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ExternalSecretsIoExternalSecretV1Alpha1Manifest{}
)

func NewExternalSecretsIoExternalSecretV1Alpha1Manifest() datasource.DataSource {
	return &ExternalSecretsIoExternalSecretV1Alpha1Manifest{}
}

type ExternalSecretsIoExternalSecretV1Alpha1Manifest struct{}

type ExternalSecretsIoExternalSecretV1Alpha1ManifestData struct {
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
		Data *[]struct {
			RemoteRef *struct {
				ConversionStrategy *string `tfsdk:"conversion_strategy" json:"conversionStrategy,omitempty"`
				Key                *string `tfsdk:"key" json:"key,omitempty"`
				Property           *string `tfsdk:"property" json:"property,omitempty"`
				Version            *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"remote_ref" json:"remoteRef,omitempty"`
			SecretKey *string `tfsdk:"secret_key" json:"secretKey,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		DataFrom *[]struct {
			ConversionStrategy *string `tfsdk:"conversion_strategy" json:"conversionStrategy,omitempty"`
			Key                *string `tfsdk:"key" json:"key,omitempty"`
			Property           *string `tfsdk:"property" json:"property,omitempty"`
			Version            *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"data_from" json:"dataFrom,omitempty"`
		RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
		SecretStoreRef  *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"secret_store_ref" json:"secretStoreRef,omitempty"`
		Target *struct {
			CreationPolicy *string `tfsdk:"creation_policy" json:"creationPolicy,omitempty"`
			Immutable      *bool   `tfsdk:"immutable" json:"immutable,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Template       *struct {
				Data          *map[string]string `tfsdk:"data" json:"data,omitempty"`
				EngineVersion *string            `tfsdk:"engine_version" json:"engineVersion,omitempty"`
				Metadata      *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				TemplateFrom *[]struct {
					ConfigMap *struct {
						Items *[]struct {
							Key *string `tfsdk:"key" json:"key,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					Secret *struct {
						Items *[]struct {
							Key *string `tfsdk:"key" json:"key,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"template_from" json:"templateFrom,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"target" json:"target,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExternalSecretsIoExternalSecretV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_external_secrets_io_external_secret_v1alpha1_manifest"
}

func (r *ExternalSecretsIoExternalSecretV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExternalSecret is the Schema for the external-secrets API.",
		MarkdownDescription: "ExternalSecret is the Schema for the external-secrets API.",
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
				Description:         "ExternalSecretSpec defines the desired state of ExternalSecret.",
				MarkdownDescription: "ExternalSecretSpec defines the desired state of ExternalSecret.",
				Attributes: map[string]schema.Attribute{
					"data": schema.ListNestedAttribute{
						Description:         "Data defines the connection between the Kubernetes Secret keys and the Provider data",
						MarkdownDescription: "Data defines the connection between the Kubernetes Secret keys and the Provider data",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"remote_ref": schema.SingleNestedAttribute{
									Description:         "ExternalSecretDataRemoteRef defines Provider data location.",
									MarkdownDescription: "ExternalSecretDataRemoteRef defines Provider data location.",
									Attributes: map[string]schema.Attribute{
										"conversion_strategy": schema.StringAttribute{
											Description:         "Used to define a conversion Strategy",
											MarkdownDescription: "Used to define a conversion Strategy",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Default", "Unicode"),
											},
										},

										"key": schema.StringAttribute{
											Description:         "Key is the key used in the Provider, mandatory",
											MarkdownDescription: "Key is the key used in the Provider, mandatory",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"property": schema.StringAttribute{
											Description:         "Used to select a specific property of the Provider value (if a map), if supported",
											MarkdownDescription: "Used to select a specific property of the Provider value (if a map), if supported",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "Used to select a specific version of the Provider value, if supported",
											MarkdownDescription: "Used to select a specific version of the Provider value, if supported",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"secret_key": schema.StringAttribute{
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

					"data_from": schema.ListNestedAttribute{
						Description:         "DataFrom is used to fetch all properties from a specific Provider dataIf multiple entries are specified, the Secret keys are merged in the specified order",
						MarkdownDescription: "DataFrom is used to fetch all properties from a specific Provider dataIf multiple entries are specified, the Secret keys are merged in the specified order",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conversion_strategy": schema.StringAttribute{
									Description:         "Used to define a conversion Strategy",
									MarkdownDescription: "Used to define a conversion Strategy",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Default", "Unicode"),
									},
								},

								"key": schema.StringAttribute{
									Description:         "Key is the key used in the Provider, mandatory",
									MarkdownDescription: "Key is the key used in the Provider, mandatory",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"property": schema.StringAttribute{
									Description:         "Used to select a specific property of the Provider value (if a map), if supported",
									MarkdownDescription: "Used to select a specific property of the Provider value (if a map), if supported",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"version": schema.StringAttribute{
									Description:         "Used to select a specific version of the Provider value, if supported",
									MarkdownDescription: "Used to select a specific version of the Provider value, if supported",
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

					"refresh_interval": schema.StringAttribute{
						Description:         "RefreshInterval is the amount of time before the values are read again from the SecretStore providerValid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'May be set to zero to fetch and create it once. Defaults to 1h.",
						MarkdownDescription: "RefreshInterval is the amount of time before the values are read again from the SecretStore providerValid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'May be set to zero to fetch and create it once. Defaults to 1h.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_store_ref": schema.SingleNestedAttribute{
						Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
						MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore)Defaults to 'SecretStore'",
								MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore)Defaults to 'SecretStore'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the SecretStore resource",
								MarkdownDescription: "Name of the SecretStore resource",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"target": schema.SingleNestedAttribute{
						Description:         "ExternalSecretTarget defines the Kubernetes Secret to be createdThere can be only one target per ExternalSecret.",
						MarkdownDescription: "ExternalSecretTarget defines the Kubernetes Secret to be createdThere can be only one target per ExternalSecret.",
						Attributes: map[string]schema.Attribute{
							"creation_policy": schema.StringAttribute{
								Description:         "CreationPolicy defines rules on how to create the resulting SecretDefaults to 'Owner'",
								MarkdownDescription: "CreationPolicy defines rules on how to create the resulting SecretDefaults to 'Owner'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Owner", "Merge", "None"),
								},
							},

							"immutable": schema.BoolAttribute{
								Description:         "Immutable defines if the final secret will be immutable",
								MarkdownDescription: "Immutable defines if the final secret will be immutable",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name defines the name of the Secret resource to be managedThis field is immutableDefaults to the .metadata.name of the ExternalSecret resource",
								MarkdownDescription: "Name defines the name of the Secret resource to be managedThis field is immutableDefaults to the .metadata.name of the ExternalSecret resource",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"template": schema.SingleNestedAttribute{
								Description:         "Template defines a blueprint for the created Secret resource.",
								MarkdownDescription: "Template defines a blueprint for the created Secret resource.",
								Attributes: map[string]schema.Attribute{
									"data": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"engine_version": schema.StringAttribute{
										Description:         "EngineVersion specifies the template engine versionthat should be used to compile/execute thetemplate specified in .data and .templateFrom[].",
										MarkdownDescription: "EngineVersion specifies the template engine versionthat should be used to compile/execute thetemplate specified in .data and .templateFrom[].",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("v1", "v2"),
										},
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
										MarkdownDescription: "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"template_from": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"items": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("secret")),
													},
												},

												"secret": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"items": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
													Validators: []validator.Object{
														objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("config_map")),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

func (r *ExternalSecretsIoExternalSecretV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_external_secrets_io_external_secret_v1alpha1_manifest")

	var model ExternalSecretsIoExternalSecretV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("external-secrets.io/v1alpha1")
	model.Kind = pointer.String("ExternalSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
