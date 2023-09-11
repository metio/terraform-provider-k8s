/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1beta1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ExternalSecretsIoExternalSecretV1Beta1Manifest{}
)

func NewExternalSecretsIoExternalSecretV1Beta1Manifest() datasource.DataSource {
	return &ExternalSecretsIoExternalSecretV1Beta1Manifest{}
}

type ExternalSecretsIoExternalSecretV1Beta1Manifest struct{}

type ExternalSecretsIoExternalSecretV1Beta1ManifestData struct {
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
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_external_secrets_io_external_secret_v1beta1_manifest"
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ExternalSecret is the Schema for the external-secrets API.",
		MarkdownDescription: "ExternalSecret is the Schema for the external-secrets API.",
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
				Description:         "ExternalSecretSpec defines the desired state of ExternalSecret.",
				MarkdownDescription: "ExternalSecretSpec defines the desired state of ExternalSecret.",
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
											Optional:            true,
											Computed:            false,
										},

										"decoding_strategy": schema.StringAttribute{
											Description:         "Used to define a decoding Strategy",
											MarkdownDescription: "Used to define a decoding Strategy",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the key used in the Provider, mandatory",
											MarkdownDescription: "Key is the key used in the Provider, mandatory",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"metadata_policy": schema.StringAttribute{
											Description:         "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
											MarkdownDescription: "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
											Required:            false,
											Optional:            true,
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
									Description:         "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",
									MarkdownDescription: "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",
									Required:            true,
									Optional:            false,
									Computed:            false,
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
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
													MarkdownDescription: "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Specify the name of the generator resource",
													MarkdownDescription: "Specify the name of the generator resource",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"store_ref": schema.SingleNestedAttribute{
											Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
											MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
											Attributes: map[string]schema.Attribute{
												"kind": schema.StringAttribute{
													Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
													MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
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
											Optional:            true,
											Computed:            false,
										},

										"decoding_strategy": schema.StringAttribute{
											Description:         "Used to define a decoding Strategy",
											MarkdownDescription: "Used to define a decoding Strategy",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the key used in the Provider, mandatory",
											MarkdownDescription: "Key is the key used in the Provider, mandatory",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"metadata_policy": schema.StringAttribute{
											Description:         "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
											MarkdownDescription: "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
											Required:            false,
											Optional:            true,
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"find": schema.SingleNestedAttribute{
									Description:         "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",
									MarkdownDescription: "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",
									Attributes: map[string]schema.Attribute{
										"conversion_strategy": schema.StringAttribute{
											Description:         "Used to define a conversion Strategy",
											MarkdownDescription: "Used to define a conversion Strategy",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"decoding_strategy": schema.StringAttribute{
											Description:         "Used to define a decoding Strategy",
											MarkdownDescription: "Used to define a decoding Strategy",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.SingleNestedAttribute{
											Description:         "Finds secrets based on the name.",
											MarkdownDescription: "Finds secrets based on the name.",
											Attributes: map[string]schema.Attribute{
												"regexp": schema.StringAttribute{
													Description:         "Finds secrets base",
													MarkdownDescription: "Finds secrets base",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"path": schema.StringAttribute{
											Description:         "A root path to start the find operations.",
											MarkdownDescription: "A root path to start the find operations.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Find secrets based on tags.",
											MarkdownDescription: "Find secrets based on tags.",
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
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"target": schema.StringAttribute{
														Description:         "Used to define the target pattern of a ReplaceAll operation.",
														MarkdownDescription: "Used to define the target pattern of a ReplaceAll operation.",
														Required:            true,
														Optional:            false,
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
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
													Description:         "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
													MarkdownDescription: "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Specify the name of the generator resource",
													MarkdownDescription: "Specify the name of the generator resource",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"store_ref": schema.SingleNestedAttribute{
											Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
											MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
											Attributes: map[string]schema.Attribute{
												"kind": schema.StringAttribute{
													Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
													MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"refresh_interval": schema.StringAttribute{
						Description:         "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' May be set to zero to fetch and create it once. Defaults to 1h.",
						MarkdownDescription: "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' May be set to zero to fetch and create it once. Defaults to 1h.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_store_ref": schema.SingleNestedAttribute{
						Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
						MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
								MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target": schema.SingleNestedAttribute{
						Description:         "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
						MarkdownDescription: "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
						Attributes: map[string]schema.Attribute{
							"creation_policy": schema.StringAttribute{
								Description:         "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
								MarkdownDescription: "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Owner", "Orphan", "Merge", "None"),
								},
							},

							"deletion_policy": schema.StringAttribute{
								Description:         "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",
								MarkdownDescription: "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Delete", "Merge", "Retain"),
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
								Description:         "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
								MarkdownDescription: "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
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
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"merge_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
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

																	"template_as": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
												},

												"literal": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
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

																	"template_as": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
												},

												"target": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

func (r *ExternalSecretsIoExternalSecretV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_external_secrets_io_external_secret_v1beta1_manifest")

	var model ExternalSecretsIoExternalSecretV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("external-secrets.io/v1beta1")
	model.Kind = pointer.String("ExternalSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
