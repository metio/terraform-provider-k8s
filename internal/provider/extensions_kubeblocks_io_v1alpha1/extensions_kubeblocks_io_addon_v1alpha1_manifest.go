/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package extensions_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &ExtensionsKubeblocksIoAddonV1Alpha1Manifest{}
)

func NewExtensionsKubeblocksIoAddonV1Alpha1Manifest() datasource.DataSource {
	return &ExtensionsKubeblocksIoAddonV1Alpha1Manifest{}
}

type ExtensionsKubeblocksIoAddonV1Alpha1Manifest struct{}

type ExtensionsKubeblocksIoAddonV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CliPlugins *[]struct {
			Description     *string `tfsdk:"description" json:"description,omitempty"`
			IndexRepository *string `tfsdk:"index_repository" json:"indexRepository,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"cli_plugins" json:"cliPlugins,omitempty"`
		DefaultInstallValues *[]struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Extras  *[]struct {
				Name                    *string `tfsdk:"name" json:"name,omitempty"`
				PersistentVolumeEnabled *bool   `tfsdk:"persistent_volume_enabled" json:"persistentVolumeEnabled,omitempty"`
				Replicas                *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources               *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				Tolerations  *string `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"extras" json:"extras,omitempty"`
			PersistentVolumeEnabled *bool  `tfsdk:"persistent_volume_enabled" json:"persistentVolumeEnabled,omitempty"`
			Replicas                *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources               *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Selectors *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"selectors" json:"selectors,omitempty"`
			StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
			Tolerations  *string `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"default_install_values" json:"defaultInstallValues,omitempty"`
		Description *string `tfsdk:"description" json:"description,omitempty"`
		Helm        *struct {
			ChartLocationURL  *string            `tfsdk:"chart_location_url" json:"chartLocationURL,omitempty"`
			ChartsImage       *string            `tfsdk:"charts_image" json:"chartsImage,omitempty"`
			ChartsPathInImage *string            `tfsdk:"charts_path_in_image" json:"chartsPathInImage,omitempty"`
			InstallOptions    *map[string]string `tfsdk:"install_options" json:"installOptions,omitempty"`
			InstallValues     *struct {
				ConfigMapRefs *[]struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
				SecretRefs *[]struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
				SetJSONValues *[]string `tfsdk:"set_json_values" json:"setJSONValues,omitempty"`
				SetValues     *[]string `tfsdk:"set_values" json:"setValues,omitempty"`
				Urls          *[]string `tfsdk:"urls" json:"urls,omitempty"`
			} `tfsdk:"install_values" json:"installValues,omitempty"`
			ValuesMapping *struct {
				Extras *[]struct {
					JsonMap *struct {
						Tolerations *string `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"json_map" json:"jsonMap,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Resources *struct {
						Cpu *struct {
							Limits   *string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							Limits   *string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
						Storage *string `tfsdk:"storage" json:"storage,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					ValueMap *struct {
						PersistentVolumeEnabled *string `tfsdk:"persistent_volume_enabled" json:"persistentVolumeEnabled,omitempty"`
						ReplicaCount            *string `tfsdk:"replica_count" json:"replicaCount,omitempty"`
						StorageClass            *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
					} `tfsdk:"value_map" json:"valueMap,omitempty"`
				} `tfsdk:"extras" json:"extras,omitempty"`
				JsonMap *struct {
					Tolerations *string `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"json_map" json:"jsonMap,omitempty"`
				Resources *struct {
					Cpu *struct {
						Limits   *string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limits   *string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				ValueMap *struct {
					PersistentVolumeEnabled *string `tfsdk:"persistent_volume_enabled" json:"persistentVolumeEnabled,omitempty"`
					ReplicaCount            *string `tfsdk:"replica_count" json:"replicaCount,omitempty"`
					StorageClass            *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				} `tfsdk:"value_map" json:"valueMap,omitempty"`
			} `tfsdk:"values_mapping" json:"valuesMapping,omitempty"`
		} `tfsdk:"helm" json:"helm,omitempty"`
		Install *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Extras  *[]struct {
				Name                    *string `tfsdk:"name" json:"name,omitempty"`
				PersistentVolumeEnabled *bool   `tfsdk:"persistent_volume_enabled" json:"persistentVolumeEnabled,omitempty"`
				Replicas                *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources               *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				Tolerations  *string `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"extras" json:"extras,omitempty"`
			PersistentVolumeEnabled *bool  `tfsdk:"persistent_volume_enabled" json:"persistentVolumeEnabled,omitempty"`
			Replicas                *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources               *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
			Tolerations  *string `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"install" json:"install,omitempty"`
		Installable *struct {
			AutoInstall *bool `tfsdk:"auto_install" json:"autoInstall,omitempty"`
			Selectors   *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"selectors" json:"selectors,omitempty"`
		} `tfsdk:"installable" json:"installable,omitempty"`
		Provider *string `tfsdk:"provider" json:"provider,omitempty"`
		Type     *string `tfsdk:"type" json:"type,omitempty"`
		Version  *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExtensionsKubeblocksIoAddonV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_extensions_kubeblocks_io_addon_v1alpha1_manifest"
}

func (r *ExtensionsKubeblocksIoAddonV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Addon is the Schema for the add-ons API.",
		MarkdownDescription: "Addon is the Schema for the add-ons API.",
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
				Description:         "AddonSpec defines the desired state of an add-on.",
				MarkdownDescription: "AddonSpec defines the desired state of an add-on.",
				Attributes: map[string]schema.Attribute{
					"cli_plugins": schema.ListNestedAttribute{
						Description:         "Specifies the CLI plugin installation specifications.",
						MarkdownDescription: "Specifies the CLI plugin installation specifications.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"description": schema.StringAttribute{
									Description:         "Provides a brief description of the plugin.",
									MarkdownDescription: "Provides a brief description of the plugin.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"index_repository": schema.StringAttribute{
									Description:         "Defines the index repository of the plugin.",
									MarkdownDescription: "Defines the index repository of the plugin.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Specifies the name of the plugin.",
									MarkdownDescription: "Specifies the name of the plugin.",
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

					"default_install_values": schema.ListNestedAttribute{
						Description:         "Specifies the default installation parameters.",
						MarkdownDescription: "Specifies the default installation parameters.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description:         "Can be set to true if there are no specific installation attributes to be set.",
									MarkdownDescription: "Can be set to true if there are no specific installation attributes to be set.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"extras": schema.ListNestedAttribute{
									Description:         "Specifies the installation specifications for extra items.",
									MarkdownDescription: "Specifies the installation specifications for extra items.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Specifies the name of the item.",
												MarkdownDescription: "Specifies the name of the item.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"persistent_volume_enabled": schema.BoolAttribute{
												Description:         "Indicates whether the Persistent Volume is enabled or not.",
												MarkdownDescription: "Indicates whether the Persistent Volume is enabled or not.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replicas": schema.Int64Attribute{
												Description:         "Specifies the number of replicas.",
												MarkdownDescription: "Specifies the number of replicas.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Specifies the resource requirements.",
												MarkdownDescription: "Specifies the resource requirements.",
												Attributes: map[string]schema.Attribute{
													"limits": schema.MapAttribute{
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
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

											"storage_class": schema.StringAttribute{
												Description:         "Specifies the name of the storage class.",
												MarkdownDescription: "Specifies the name of the storage class.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tolerations": schema.StringAttribute{
												Description:         "Specifies the tolerations in a JSON array string format.",
												MarkdownDescription: "Specifies the tolerations in a JSON array string format.",
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

								"persistent_volume_enabled": schema.BoolAttribute{
									Description:         "Indicates whether the Persistent Volume is enabled or not.",
									MarkdownDescription: "Indicates whether the Persistent Volume is enabled or not.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Specifies the number of replicas.",
									MarkdownDescription: "Specifies the number of replicas.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Specifies the resource requirements.",
									MarkdownDescription: "Specifies the resource requirements.",
									Attributes: map[string]schema.Attribute{
										"limits": schema.MapAttribute{
											Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
											MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
											MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
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

								"selectors": schema.ListNestedAttribute{
									Description:         "Indicates the default selectors for add-on installations. If multiple selectors are provided, all selectors must evaluate to true.",
									MarkdownDescription: "Indicates the default selectors for add-on installations. If multiple selectors are provided, all selectors must evaluate to true.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The selector key. Valid values are KubeVersion, KubeGitVersion and KubeProvider. - 'KubeVersion' the semver expression of Kubernetes versions, i.e., v1.24. - 'KubeGitVersion' may contain distro. info., i.e., v1.24.4+eks. - 'KubeProvider' the Kubernetes provider, i.e., aws, gcp, azure, huaweiCloud, tencentCloud etc.",
												MarkdownDescription: "The selector key. Valid values are KubeVersion, KubeGitVersion and KubeProvider. - 'KubeVersion' the semver expression of Kubernetes versions, i.e., v1.24. - 'KubeGitVersion' may contain distro. info., i.e., v1.24.4+eks. - 'KubeProvider' the Kubernetes provider, i.e., aws, gcp, azure, huaweiCloud, tencentCloud etc.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("KubeGitVersion", "KubeVersion", "KubeProvider"),
												},
											},

											"operator": schema.StringAttribute{
												Description:         "Represents a key's relationship to a set of values. Valid operators are Contains, NotIn, DoesNotContain, MatchRegex, and DoesNoteMatchRegex. Possible enum values: - 'Contains' line contains a string. - 'DoesNotContain' line does not contain a string. - 'MatchRegex' line contains a match to the regular expression. - 'DoesNotMatchRegex' line does not contain a match to the regular expression.",
												MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are Contains, NotIn, DoesNotContain, MatchRegex, and DoesNoteMatchRegex. Possible enum values: - 'Contains' line contains a string. - 'DoesNotContain' line does not contain a string. - 'MatchRegex' line contains a match to the regular expression. - 'DoesNotMatchRegex' line does not contain a match to the regular expression.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Contains", "DoesNotContain", "MatchRegex", "DoesNotMatchRegex"),
												},
											},

											"values": schema.ListAttribute{
												Description:         "Represents an array of string values. This serves as an 'OR' expression to the operator.",
												MarkdownDescription: "Represents an array of string values. This serves as an 'OR' expression to the operator.",
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

								"storage_class": schema.StringAttribute{
									Description:         "Specifies the name of the storage class.",
									MarkdownDescription: "Specifies the name of the storage class.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tolerations": schema.StringAttribute{
									Description:         "Specifies the tolerations in a JSON array string format.",
									MarkdownDescription: "Specifies the tolerations in a JSON array string format.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "Specifies the description of the add-on.",
						MarkdownDescription: "Specifies the description of the add-on.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"helm": schema.SingleNestedAttribute{
						Description:         "Represents the Helm installation specifications. This is only processed when the type is set to 'helm'.",
						MarkdownDescription: "Represents the Helm installation specifications. This is only processed when the type is set to 'helm'.",
						Attributes: map[string]schema.Attribute{
							"chart_location_url": schema.StringAttribute{
								Description:         "Specifies the URL location of the Helm Chart.",
								MarkdownDescription: "Specifies the URL location of the Helm Chart.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"charts_image": schema.StringAttribute{
								Description:         "Defines the image of Helm charts.",
								MarkdownDescription: "Defines the image of Helm charts.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"charts_path_in_image": schema.StringAttribute{
								Description:         "Defines the path of Helm charts in the image. This path is used to copy Helm charts from the image to the shared volume. The default path is '/charts'.",
								MarkdownDescription: "Defines the path of Helm charts in the image. This path is used to copy Helm charts from the image to the shared volume. The default path is '/charts'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"install_options": schema.MapAttribute{
								Description:         "Defines the options for Helm release installation.",
								MarkdownDescription: "Defines the options for Helm release installation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"install_values": schema.SingleNestedAttribute{
								Description:         "Defines the set values for Helm release installation.",
								MarkdownDescription: "Defines the set values for Helm release installation.",
								Attributes: map[string]schema.Attribute{
									"config_map_refs": schema.ListNestedAttribute{
										Description:         "Selects a key from a ConfigMap item list. The value can be a JSON or YAML string content. Use a key name with '.json', '.yaml', or '.yml' extension to specify a content type.",
										MarkdownDescription: "Selects a key from a ConfigMap item list. The value can be a JSON or YAML string content. Use a key name with '.json', '.yaml', or '.yml' extension to specify a content type.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "Specifies the key to be selected.",
													MarkdownDescription: "Specifies the key to be selected.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Defines the name of the object being referred to.",
													MarkdownDescription: "Defines the name of the object being referred to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_refs": schema.ListNestedAttribute{
										Description:         "Selects a key from a Secrets item list. The value can be a JSON or YAML string content. Use a key name with '.json', '.yaml', or '.yml' extension to specify a content type.",
										MarkdownDescription: "Selects a key from a Secrets item list. The value can be a JSON or YAML string content. Use a key name with '.json', '.yaml', or '.yml' extension to specify a content type.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "Specifies the key to be selected.",
													MarkdownDescription: "Specifies the key to be selected.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Defines the name of the object being referred to.",
													MarkdownDescription: "Defines the name of the object being referred to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"set_json_values": schema.ListAttribute{
										Description:         "JSON values set during Helm installation. Multiple or separate values can be specified with commas (key1=jsonval1,key2=jsonval2).",
										MarkdownDescription: "JSON values set during Helm installation. Multiple or separate values can be specified with commas (key1=jsonval1,key2=jsonval2).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"set_values": schema.ListAttribute{
										Description:         "Values set during Helm installation. Multiple or separate values can be specified with commas (key1=val1,key2=val2).",
										MarkdownDescription: "Values set during Helm installation. Multiple or separate values can be specified with commas (key1=val1,key2=val2).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"urls": schema.ListAttribute{
										Description:         "Specifies the URL location of the values file.",
										MarkdownDescription: "Specifies the URL location of the values file.",
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

							"values_mapping": schema.SingleNestedAttribute{
								Description:         "Defines the mapping of add-on normalized resources parameters to Helm values' keys.",
								MarkdownDescription: "Defines the mapping of add-on normalized resources parameters to Helm values' keys.",
								Attributes: map[string]schema.Attribute{
									"extras": schema.ListNestedAttribute{
										Description:         "Helm value mapping items for extra items.",
										MarkdownDescription: "Helm value mapping items for extra items.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"json_map": schema.SingleNestedAttribute{
													Description:         "Defines the 'key' mapping values. The valid key is tolerations. Enum values explained: - 'tolerations' sets the toleration mapping key.",
													MarkdownDescription: "Defines the 'key' mapping values. The valid key is tolerations. Enum values explained: - 'tolerations' sets the toleration mapping key.",
													Attributes: map[string]schema.Attribute{
														"tolerations": schema.StringAttribute{
															Description:         "Specifies the toleration mapping key.",
															MarkdownDescription: "Specifies the toleration mapping key.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the item.",
													MarkdownDescription: "Name of the item.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"resources": schema.SingleNestedAttribute{
													Description:         "Sets resources related mapping keys.",
													MarkdownDescription: "Sets resources related mapping keys.",
													Attributes: map[string]schema.Attribute{
														"cpu": schema.SingleNestedAttribute{
															Description:         "Specifies the key used for mapping both CPU requests and limits.",
															MarkdownDescription: "Specifies the key used for mapping both CPU requests and limits.",
															Attributes: map[string]schema.Attribute{
																"limits": schema.StringAttribute{
																	Description:         "Specifies the mapping key for the limit value.",
																	MarkdownDescription: "Specifies the mapping key for the limit value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"requests": schema.StringAttribute{
																	Description:         "Specifies the mapping key for the request value.",
																	MarkdownDescription: "Specifies the mapping key for the request value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"memory": schema.SingleNestedAttribute{
															Description:         "Specifies the key used for mapping both Memory requests and limits.",
															MarkdownDescription: "Specifies the key used for mapping both Memory requests and limits.",
															Attributes: map[string]schema.Attribute{
																"limits": schema.StringAttribute{
																	Description:         "Specifies the mapping key for the limit value.",
																	MarkdownDescription: "Specifies the mapping key for the limit value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"requests": schema.StringAttribute{
																	Description:         "Specifies the mapping key for the request value.",
																	MarkdownDescription: "Specifies the mapping key for the request value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"storage": schema.StringAttribute{
															Description:         "Specifies the key used for mapping the storage size value.",
															MarkdownDescription: "Specifies the key used for mapping the storage size value.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"value_map": schema.SingleNestedAttribute{
													Description:         "Defines the 'key' mapping values. Valid keys include 'replicaCount', 'persistentVolumeEnabled', and 'storageClass'. Enum values explained: - 'replicaCount' sets the replicaCount value mapping key. - 'persistentVolumeEnabled' sets the persistent volume enabled mapping key. - 'storageClass' sets the storageClass mapping key.",
													MarkdownDescription: "Defines the 'key' mapping values. Valid keys include 'replicaCount', 'persistentVolumeEnabled', and 'storageClass'. Enum values explained: - 'replicaCount' sets the replicaCount value mapping key. - 'persistentVolumeEnabled' sets the persistent volume enabled mapping key. - 'storageClass' sets the storageClass mapping key.",
													Attributes: map[string]schema.Attribute{
														"persistent_volume_enabled": schema.StringAttribute{
															Description:         "Indicates whether the persistent volume is enabled in the Helm values map.",
															MarkdownDescription: "Indicates whether the persistent volume is enabled in the Helm values map.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"replica_count": schema.StringAttribute{
															Description:         "Defines the key for setting the replica count in the Helm values map.",
															MarkdownDescription: "Defines the key for setting the replica count in the Helm values map.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"storage_class": schema.StringAttribute{
															Description:         "Specifies the key for setting the storage class in the Helm values map.",
															MarkdownDescription: "Specifies the key for setting the storage class in the Helm values map.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_map": schema.SingleNestedAttribute{
										Description:         "Defines the 'key' mapping values. The valid key is tolerations. Enum values explained: - 'tolerations' sets the toleration mapping key.",
										MarkdownDescription: "Defines the 'key' mapping values. The valid key is tolerations. Enum values explained: - 'tolerations' sets the toleration mapping key.",
										Attributes: map[string]schema.Attribute{
											"tolerations": schema.StringAttribute{
												Description:         "Specifies the toleration mapping key.",
												MarkdownDescription: "Specifies the toleration mapping key.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Sets resources related mapping keys.",
										MarkdownDescription: "Sets resources related mapping keys.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.SingleNestedAttribute{
												Description:         "Specifies the key used for mapping both CPU requests and limits.",
												MarkdownDescription: "Specifies the key used for mapping both CPU requests and limits.",
												Attributes: map[string]schema.Attribute{
													"limits": schema.StringAttribute{
														Description:         "Specifies the mapping key for the limit value.",
														MarkdownDescription: "Specifies the mapping key for the limit value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.StringAttribute{
														Description:         "Specifies the mapping key for the request value.",
														MarkdownDescription: "Specifies the mapping key for the request value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": schema.SingleNestedAttribute{
												Description:         "Specifies the key used for mapping both Memory requests and limits.",
												MarkdownDescription: "Specifies the key used for mapping both Memory requests and limits.",
												Attributes: map[string]schema.Attribute{
													"limits": schema.StringAttribute{
														Description:         "Specifies the mapping key for the limit value.",
														MarkdownDescription: "Specifies the mapping key for the limit value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.StringAttribute{
														Description:         "Specifies the mapping key for the request value.",
														MarkdownDescription: "Specifies the mapping key for the request value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage": schema.StringAttribute{
												Description:         "Specifies the key used for mapping the storage size value.",
												MarkdownDescription: "Specifies the key used for mapping the storage size value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value_map": schema.SingleNestedAttribute{
										Description:         "Defines the 'key' mapping values. Valid keys include 'replicaCount', 'persistentVolumeEnabled', and 'storageClass'. Enum values explained: - 'replicaCount' sets the replicaCount value mapping key. - 'persistentVolumeEnabled' sets the persistent volume enabled mapping key. - 'storageClass' sets the storageClass mapping key.",
										MarkdownDescription: "Defines the 'key' mapping values. Valid keys include 'replicaCount', 'persistentVolumeEnabled', and 'storageClass'. Enum values explained: - 'replicaCount' sets the replicaCount value mapping key. - 'persistentVolumeEnabled' sets the persistent volume enabled mapping key. - 'storageClass' sets the storageClass mapping key.",
										Attributes: map[string]schema.Attribute{
											"persistent_volume_enabled": schema.StringAttribute{
												Description:         "Indicates whether the persistent volume is enabled in the Helm values map.",
												MarkdownDescription: "Indicates whether the persistent volume is enabled in the Helm values map.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replica_count": schema.StringAttribute{
												Description:         "Defines the key for setting the replica count in the Helm values map.",
												MarkdownDescription: "Defines the key for setting the replica count in the Helm values map.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"storage_class": schema.StringAttribute{
												Description:         "Specifies the key for setting the storage class in the Helm values map.",
												MarkdownDescription: "Specifies the key for setting the storage class in the Helm values map.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"install": schema.SingleNestedAttribute{
						Description:         "Defines the installation parameters.",
						MarkdownDescription: "Defines the installation parameters.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Can be set to true if there are no specific installation attributes to be set.",
								MarkdownDescription: "Can be set to true if there are no specific installation attributes to be set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extras": schema.ListNestedAttribute{
								Description:         "Specifies the installation specifications for extra items.",
								MarkdownDescription: "Specifies the installation specifications for extra items.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Specifies the name of the item.",
											MarkdownDescription: "Specifies the name of the item.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"persistent_volume_enabled": schema.BoolAttribute{
											Description:         "Indicates whether the Persistent Volume is enabled or not.",
											MarkdownDescription: "Indicates whether the Persistent Volume is enabled or not.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replicas": schema.Int64Attribute{
											Description:         "Specifies the number of replicas.",
											MarkdownDescription: "Specifies the number of replicas.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "Specifies the resource requirements.",
											MarkdownDescription: "Specifies the resource requirements.",
											Attributes: map[string]schema.Attribute{
												"limits": schema.MapAttribute{
													Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
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

										"storage_class": schema.StringAttribute{
											Description:         "Specifies the name of the storage class.",
											MarkdownDescription: "Specifies the name of the storage class.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tolerations": schema.StringAttribute{
											Description:         "Specifies the tolerations in a JSON array string format.",
											MarkdownDescription: "Specifies the tolerations in a JSON array string format.",
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

							"persistent_volume_enabled": schema.BoolAttribute{
								Description:         "Indicates whether the Persistent Volume is enabled or not.",
								MarkdownDescription: "Indicates whether the Persistent Volume is enabled or not.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Specifies the number of replicas.",
								MarkdownDescription: "Specifies the number of replicas.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Specifies the resource requirements.",
								MarkdownDescription: "Specifies the resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified; otherwise, it defaults to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/.",
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

							"storage_class": schema.StringAttribute{
								Description:         "Specifies the name of the storage class.",
								MarkdownDescription: "Specifies the name of the storage class.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.StringAttribute{
								Description:         "Specifies the tolerations in a JSON array string format.",
								MarkdownDescription: "Specifies the tolerations in a JSON array string format.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"installable": schema.SingleNestedAttribute{
						Description:         "Represents the installable specifications of the add-on. This includes the selector and auto-install settings.",
						MarkdownDescription: "Represents the installable specifications of the add-on. This includes the selector and auto-install settings.",
						Attributes: map[string]schema.Attribute{
							"auto_install": schema.BoolAttribute{
								Description:         "Indicates whether an add-on should be installed automatically.",
								MarkdownDescription: "Indicates whether an add-on should be installed automatically.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"selectors": schema.ListNestedAttribute{
								Description:         "Specifies the selectors for add-on installation. If multiple selectors are provided, they must all evaluate to true for the add-on to be installed.",
								MarkdownDescription: "Specifies the selectors for add-on installation. If multiple selectors are provided, they must all evaluate to true for the add-on to be installed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The selector key. Valid values are KubeVersion, KubeGitVersion and KubeProvider. - 'KubeVersion' the semver expression of Kubernetes versions, i.e., v1.24. - 'KubeGitVersion' may contain distro. info., i.e., v1.24.4+eks. - 'KubeProvider' the Kubernetes provider, i.e., aws, gcp, azure, huaweiCloud, tencentCloud etc.",
											MarkdownDescription: "The selector key. Valid values are KubeVersion, KubeGitVersion and KubeProvider. - 'KubeVersion' the semver expression of Kubernetes versions, i.e., v1.24. - 'KubeGitVersion' may contain distro. info., i.e., v1.24.4+eks. - 'KubeProvider' the Kubernetes provider, i.e., aws, gcp, azure, huaweiCloud, tencentCloud etc.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("KubeGitVersion", "KubeVersion", "KubeProvider"),
											},
										},

										"operator": schema.StringAttribute{
											Description:         "Represents a key's relationship to a set of values. Valid operators are Contains, NotIn, DoesNotContain, MatchRegex, and DoesNoteMatchRegex. Possible enum values: - 'Contains' line contains a string. - 'DoesNotContain' line does not contain a string. - 'MatchRegex' line contains a match to the regular expression. - 'DoesNotMatchRegex' line does not contain a match to the regular expression.",
											MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are Contains, NotIn, DoesNotContain, MatchRegex, and DoesNoteMatchRegex. Possible enum values: - 'Contains' line contains a string. - 'DoesNotContain' line does not contain a string. - 'MatchRegex' line contains a match to the regular expression. - 'DoesNotMatchRegex' line does not contain a match to the regular expression.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Contains", "DoesNotContain", "MatchRegex", "DoesNotMatchRegex"),
											},
										},

										"values": schema.ListAttribute{
											Description:         "Represents an array of string values. This serves as an 'OR' expression to the operator.",
											MarkdownDescription: "Represents an array of string values. This serves as an 'OR' expression to the operator.",
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

					"provider": schema.StringAttribute{
						Description:         "Specifies the provider of the add-on.",
						MarkdownDescription: "Specifies the provider of the add-on.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Defines the type of the add-on. The only valid value is 'helm'.",
						MarkdownDescription: "Defines the type of the add-on. The only valid value is 'helm'.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Helm"),
						},
					},

					"version": schema.StringAttribute{
						Description:         "Indicates the version of the add-on.",
						MarkdownDescription: "Indicates the version of the add-on.",
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

func (r *ExtensionsKubeblocksIoAddonV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_extensions_kubeblocks_io_addon_v1alpha1_manifest")

	var model ExtensionsKubeblocksIoAddonV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("extensions.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Addon")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
