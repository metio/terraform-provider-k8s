/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1beta1

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &ExternalSecretsIoClusterExternalSecretV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &ExternalSecretsIoClusterExternalSecretV1Beta1DataSource{}
)

func NewExternalSecretsIoClusterExternalSecretV1Beta1DataSource() datasource.DataSource {
	return &ExternalSecretsIoClusterExternalSecretV1Beta1DataSource{}
}

type ExternalSecretsIoClusterExternalSecretV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ExternalSecretsIoClusterExternalSecretV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ExternalSecretMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"external_secret_metadata" json:"externalSecretMetadata,omitempty"`
		ExternalSecretName *string `tfsdk:"external_secret_name" json:"externalSecretName,omitempty"`
		ExternalSecretSpec *struct {
			Data *[]struct {
				RemoteRef *struct {
					ConversionStrategy *string `tfsdk:"conversion_strategy" json:"conversionStrategy,omitempty"`
					DecodingStrategy   *string `tfsdk:"decoding_strategy" json:"decodingStrategy,omitempty"`
					Key                *string `tfsdk:"key" json:"key,omitempty"`
					MetadataPolicy     *string `tfsdk:"metadata_policy" json:"metadataPolicy,omitempty"`
					Property           *string `tfsdk:"property" json:"property,omitempty"`
					Version            *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"remote_ref" json:"remoteRef,omitempty"`
				SecretKey *string `tfsdk:"secret_key" json:"secretKey,omitempty"`
				SourceRef *struct {
					GeneratorRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"generator_ref" json:"generatorRef,omitempty"`
					StoreRef *struct {
						Kind *string `tfsdk:"kind" json:"kind,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"store_ref" json:"storeRef,omitempty"`
				} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
			} `tfsdk:"data" json:"data,omitempty"`
			DataFrom *[]struct {
				Extract *struct {
					ConversionStrategy *string `tfsdk:"conversion_strategy" json:"conversionStrategy,omitempty"`
					DecodingStrategy   *string `tfsdk:"decoding_strategy" json:"decodingStrategy,omitempty"`
					Key                *string `tfsdk:"key" json:"key,omitempty"`
					MetadataPolicy     *string `tfsdk:"metadata_policy" json:"metadataPolicy,omitempty"`
					Property           *string `tfsdk:"property" json:"property,omitempty"`
					Version            *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"extract" json:"extract,omitempty"`
				Find *struct {
					ConversionStrategy *string `tfsdk:"conversion_strategy" json:"conversionStrategy,omitempty"`
					DecodingStrategy   *string `tfsdk:"decoding_strategy" json:"decodingStrategy,omitempty"`
					Name               *struct {
						Regexp *string `tfsdk:"regexp" json:"regexp,omitempty"`
					} `tfsdk:"name" json:"name,omitempty"`
					Path *string            `tfsdk:"path" json:"path,omitempty"`
					Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
				} `tfsdk:"find" json:"find,omitempty"`
				Rewrite *[]struct {
					Regexp *struct {
						Source *string `tfsdk:"source" json:"source,omitempty"`
						Target *string `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"regexp" json:"regexp,omitempty"`
				} `tfsdk:"rewrite" json:"rewrite,omitempty"`
				SourceRef *struct {
					GeneratorRef *struct {
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"generator_ref" json:"generatorRef,omitempty"`
					StoreRef *struct {
						Kind *string `tfsdk:"kind" json:"kind,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"store_ref" json:"storeRef,omitempty"`
				} `tfsdk:"source_ref" json:"sourceRef,omitempty"`
			} `tfsdk:"data_from" json:"dataFrom,omitempty"`
			RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
			SecretStoreRef  *struct {
				Kind *string `tfsdk:"kind" json:"kind,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_store_ref" json:"secretStoreRef,omitempty"`
			Target *struct {
				CreationPolicy *string `tfsdk:"creation_policy" json:"creationPolicy,omitempty"`
				DeletionPolicy *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
				Immutable      *bool   `tfsdk:"immutable" json:"immutable,omitempty"`
				Name           *string `tfsdk:"name" json:"name,omitempty"`
				Template       *struct {
					Data          *map[string]string `tfsdk:"data" json:"data,omitempty"`
					EngineVersion *string            `tfsdk:"engine_version" json:"engineVersion,omitempty"`
					MergePolicy   *string            `tfsdk:"merge_policy" json:"mergePolicy,omitempty"`
					Metadata      *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					TemplateFrom *[]struct {
						ConfigMap *struct {
							Items *[]struct {
								Key        *string `tfsdk:"key" json:"key,omitempty"`
								TemplateAs *string `tfsdk:"template_as" json:"templateAs,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						Literal *string `tfsdk:"literal" json:"literal,omitempty"`
						Secret  *struct {
							Items *[]struct {
								Key        *string `tfsdk:"key" json:"key,omitempty"`
								TemplateAs *string `tfsdk:"template_as" json:"templateAs,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						Target *string `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"template_from" json:"templateFrom,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"target" json:"target,omitempty"`
		} `tfsdk:"external_secret_spec" json:"externalSecretSpec,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		RefreshTime *string `tfsdk:"refresh_time" json:"refreshTime,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_external_secrets_io_cluster_external_secret_v1beta1"
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterExternalSecret is the Schema for the clusterexternalsecrets API.",
		MarkdownDescription: "ClusterExternalSecret is the Schema for the clusterexternalsecrets API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "ClusterExternalSecretSpec defines the desired state of ClusterExternalSecret.",
				MarkdownDescription: "ClusterExternalSecretSpec defines the desired state of ClusterExternalSecret.",
				Attributes: map[string]schema.Attribute{
					"external_secret_metadata": schema.SingleNestedAttribute{
						Description:         "The metadata of the external secrets to be created",
						MarkdownDescription: "The metadata of the external secrets to be created",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"external_secret_name": schema.StringAttribute{
						Description:         "The name of the external secrets to be created defaults to the name of the ClusterExternalSecret",
						MarkdownDescription: "The name of the external secrets to be created defaults to the name of the ClusterExternalSecret",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"external_secret_spec": schema.SingleNestedAttribute{
						Description:         "The spec for the ExternalSecrets to be created",
						MarkdownDescription: "The spec for the ExternalSecrets to be created",
						Attributes: map[string]schema.Attribute{
							"data": schema.ListNestedAttribute{
								Description:         "Data defines the connection between the Kubernetes Secret keys and the Provider data",
								MarkdownDescription: "Data defines the connection between the Kubernetes Secret keys and the Provider data",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"remote_ref": schema.SingleNestedAttribute{
											Description:         "RemoteRef points to the remote secret and defines which secret (version/property/..) to fetch.",
											MarkdownDescription: "RemoteRef points to the remote secret and defines which secret (version/property/..) to fetch.",
											Attributes: map[string]schema.Attribute{
												"conversion_strategy": schema.StringAttribute{
													Description:         "Used to define a conversion Strategy",
													MarkdownDescription: "Used to define a conversion Strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"decoding_strategy": schema.StringAttribute{
													Description:         "Used to define a decoding Strategy",
													MarkdownDescription: "Used to define a decoding Strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the key used in the Provider, mandatory",
													MarkdownDescription: "Key is the key used in the Provider, mandatory",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"metadata_policy": schema.StringAttribute{
													Description:         "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
													MarkdownDescription: "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"property": schema.StringAttribute{
													Description:         "Used to select a specific property of the Provider value (if a map), if supported",
													MarkdownDescription: "Used to select a specific property of the Provider value (if a map), if supported",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"version": schema.StringAttribute{
													Description:         "Used to select a specific version of the Provider value, if supported",
													MarkdownDescription: "Used to select a specific version of the Provider value, if supported",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"secret_key": schema.StringAttribute{
											Description:         "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",
											MarkdownDescription: "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"source_ref": schema.SingleNestedAttribute{
											Description:         "SourceRef allows you to override the source from which the value will pulled from.",
											MarkdownDescription: "SourceRef allows you to override the source from which the value will pulled from.",
											Attributes: map[string]schema.Attribute{
												"generator_ref": schema.SingleNestedAttribute{
													Description:         "GeneratorRef points to a generator custom resource in",
													MarkdownDescription: "GeneratorRef points to a generator custom resource in",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "Specify the apiVersion of the generator resource",
															MarkdownDescription: "Specify the apiVersion of the generator resource",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"kind": schema.StringAttribute{
															Description:         "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
															MarkdownDescription: "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Specify the name of the generator resource",
															MarkdownDescription: "Specify the name of the generator resource",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"store_ref": schema.SingleNestedAttribute{
													Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
													MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
													Attributes: map[string]schema.Attribute{
														"kind": schema.StringAttribute{
															Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
															MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the SecretStore resource",
															MarkdownDescription: "Name of the SecretStore resource",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"data_from": schema.ListNestedAttribute{
								Description:         "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
								MarkdownDescription: "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"extract": schema.SingleNestedAttribute{
											Description:         "Used to extract multiple key/value pairs from one secret Note: Extract does not support sourceRef.Generator or sourceRef.GeneratorRef.",
											MarkdownDescription: "Used to extract multiple key/value pairs from one secret Note: Extract does not support sourceRef.Generator or sourceRef.GeneratorRef.",
											Attributes: map[string]schema.Attribute{
												"conversion_strategy": schema.StringAttribute{
													Description:         "Used to define a conversion Strategy",
													MarkdownDescription: "Used to define a conversion Strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"decoding_strategy": schema.StringAttribute{
													Description:         "Used to define a decoding Strategy",
													MarkdownDescription: "Used to define a decoding Strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the key used in the Provider, mandatory",
													MarkdownDescription: "Key is the key used in the Provider, mandatory",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"metadata_policy": schema.StringAttribute{
													Description:         "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
													MarkdownDescription: "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"property": schema.StringAttribute{
													Description:         "Used to select a specific property of the Provider value (if a map), if supported",
													MarkdownDescription: "Used to select a specific property of the Provider value (if a map), if supported",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"version": schema.StringAttribute{
													Description:         "Used to select a specific version of the Provider value, if supported",
													MarkdownDescription: "Used to select a specific version of the Provider value, if supported",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"find": schema.SingleNestedAttribute{
											Description:         "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",
											MarkdownDescription: "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",
											Attributes: map[string]schema.Attribute{
												"conversion_strategy": schema.StringAttribute{
													Description:         "Used to define a conversion Strategy",
													MarkdownDescription: "Used to define a conversion Strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"decoding_strategy": schema.StringAttribute{
													Description:         "Used to define a decoding Strategy",
													MarkdownDescription: "Used to define a decoding Strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.SingleNestedAttribute{
													Description:         "Finds secrets based on the name.",
													MarkdownDescription: "Finds secrets based on the name.",
													Attributes: map[string]schema.Attribute{
														"regexp": schema.StringAttribute{
															Description:         "Finds secrets base",
															MarkdownDescription: "Finds secrets base",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"path": schema.StringAttribute{
													Description:         "A root path to start the find operations.",
													MarkdownDescription: "A root path to start the find operations.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tags": schema.MapAttribute{
													Description:         "Find secrets based on tags.",
													MarkdownDescription: "Find secrets based on tags.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"rewrite": schema.ListNestedAttribute{
											Description:         "Used to rewrite secret Keys after getting them from the secret Provider Multiple Rewrite operations can be provided. They are applied in a layered order (first to last)",
											MarkdownDescription: "Used to rewrite secret Keys after getting them from the secret Provider Multiple Rewrite operations can be provided. They are applied in a layered order (first to last)",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"regexp": schema.SingleNestedAttribute{
														Description:         "Used to rewrite with regular expressions. The resulting key will be the output of a regexp.ReplaceAll operation.",
														MarkdownDescription: "Used to rewrite with regular expressions. The resulting key will be the output of a regexp.ReplaceAll operation.",
														Attributes: map[string]schema.Attribute{
															"source": schema.StringAttribute{
																Description:         "Used to define the regular expression of a re.Compiler.",
																MarkdownDescription: "Used to define the regular expression of a re.Compiler.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"target": schema.StringAttribute{
																Description:         "Used to define the target pattern of a ReplaceAll operation.",
																MarkdownDescription: "Used to define the target pattern of a ReplaceAll operation.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"source_ref": schema.SingleNestedAttribute{
											Description:         "SourceRef points to a store or generator which contains secret values ready to use. Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values",
											MarkdownDescription: "SourceRef points to a store or generator which contains secret values ready to use. Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values",
											Attributes: map[string]schema.Attribute{
												"generator_ref": schema.SingleNestedAttribute{
													Description:         "GeneratorRef points to a generator custom resource in",
													MarkdownDescription: "GeneratorRef points to a generator custom resource in",
													Attributes: map[string]schema.Attribute{
														"api_version": schema.StringAttribute{
															Description:         "Specify the apiVersion of the generator resource",
															MarkdownDescription: "Specify the apiVersion of the generator resource",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"kind": schema.StringAttribute{
															Description:         "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
															MarkdownDescription: "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Specify the name of the generator resource",
															MarkdownDescription: "Specify the name of the generator resource",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"store_ref": schema.SingleNestedAttribute{
													Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
													MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
													Attributes: map[string]schema.Attribute{
														"kind": schema.StringAttribute{
															Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
															MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the SecretStore resource",
															MarkdownDescription: "Name of the SecretStore resource",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"refresh_interval": schema.StringAttribute{
								Description:         "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' May be set to zero to fetch and create it once. Defaults to 1h.",
								MarkdownDescription: "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' May be set to zero to fetch and create it once. Defaults to 1h.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"secret_store_ref": schema.SingleNestedAttribute{
								Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
								MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
								Attributes: map[string]schema.Attribute{
									"kind": schema.StringAttribute{
										Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
										MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the SecretStore resource",
										MarkdownDescription: "Name of the SecretStore resource",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"target": schema.SingleNestedAttribute{
								Description:         "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
								MarkdownDescription: "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
								Attributes: map[string]schema.Attribute{
									"creation_policy": schema.StringAttribute{
										Description:         "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
										MarkdownDescription: "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"deletion_policy": schema.StringAttribute{
										Description:         "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",
										MarkdownDescription: "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"immutable": schema.BoolAttribute{
										Description:         "Immutable defines if the final secret will be immutable",
										MarkdownDescription: "Immutable defines if the final secret will be immutable",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
										MarkdownDescription: "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"engine_version": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"merge_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"labels": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
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
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"template_as": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"literal": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
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
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"template_as": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"target": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "The labels to select by to find the Namespaces to create the ExternalSecrets in.",
						MarkdownDescription: "The labels to select by to find the Namespaces to create the ExternalSecrets in.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"refresh_time": schema.StringAttribute{
						Description:         "The time in which the controller should reconcile it's objects and recheck namespaces for labels.",
						MarkdownDescription: "The time in which the controller should reconcile it's objects and recheck namespaces for labels.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_external_secrets_io_cluster_external_secret_v1beta1")

	var data ExternalSecretsIoClusterExternalSecretV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "external-secrets.io", Version: "v1beta1", Resource: "clusterexternalsecrets"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ExternalSecretsIoClusterExternalSecretV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("external-secrets.io/v1beta1")
	data.Kind = pointer.String("ClusterExternalSecret")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
