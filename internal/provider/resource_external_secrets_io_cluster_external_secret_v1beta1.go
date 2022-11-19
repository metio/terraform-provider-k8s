/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type ExternalSecretsIoClusterExternalSecretV1Beta1Resource struct{}

var (
	_ resource.Resource = (*ExternalSecretsIoClusterExternalSecretV1Beta1Resource)(nil)
)

type ExternalSecretsIoClusterExternalSecretV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ExternalSecretsIoClusterExternalSecretV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ExternalSecretName *string `tfsdk:"external_secret_name" yaml:"externalSecretName,omitempty"`

		ExternalSecretSpec *struct {
			Data *[]struct {
				RemoteRef *struct {
					ConversionStrategy *string `tfsdk:"conversion_strategy" yaml:"conversionStrategy,omitempty"`

					DecodingStrategy *string `tfsdk:"decoding_strategy" yaml:"decodingStrategy,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					MetadataPolicy *string `tfsdk:"metadata_policy" yaml:"metadataPolicy,omitempty"`

					Property *string `tfsdk:"property" yaml:"property,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"remote_ref" yaml:"remoteRef,omitempty"`

				SecretKey *string `tfsdk:"secret_key" yaml:"secretKey,omitempty"`

				SourceRef *struct {
					GeneratorRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"generator_ref" yaml:"generatorRef,omitempty"`

					StoreRef *struct {
						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"store_ref" yaml:"storeRef,omitempty"`
				} `tfsdk:"source_ref" yaml:"sourceRef,omitempty"`
			} `tfsdk:"data" yaml:"data,omitempty"`

			DataFrom *[]struct {
				Extract *struct {
					ConversionStrategy *string `tfsdk:"conversion_strategy" yaml:"conversionStrategy,omitempty"`

					DecodingStrategy *string `tfsdk:"decoding_strategy" yaml:"decodingStrategy,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					MetadataPolicy *string `tfsdk:"metadata_policy" yaml:"metadataPolicy,omitempty"`

					Property *string `tfsdk:"property" yaml:"property,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"extract" yaml:"extract,omitempty"`

				Find *struct {
					ConversionStrategy *string `tfsdk:"conversion_strategy" yaml:"conversionStrategy,omitempty"`

					DecodingStrategy *string `tfsdk:"decoding_strategy" yaml:"decodingStrategy,omitempty"`

					Name *struct {
						Regexp *string `tfsdk:"regexp" yaml:"regexp,omitempty"`
					} `tfsdk:"name" yaml:"name,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Tags *map[string]string `tfsdk:"tags" yaml:"tags,omitempty"`
				} `tfsdk:"find" yaml:"find,omitempty"`

				Rewrite *[]struct {
					Regexp *struct {
						Source *string `tfsdk:"source" yaml:"source,omitempty"`

						Target *string `tfsdk:"target" yaml:"target,omitempty"`
					} `tfsdk:"regexp" yaml:"regexp,omitempty"`
				} `tfsdk:"rewrite" yaml:"rewrite,omitempty"`

				SourceRef *struct {
					GeneratorRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"generator_ref" yaml:"generatorRef,omitempty"`

					StoreRef *struct {
						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"store_ref" yaml:"storeRef,omitempty"`
				} `tfsdk:"source_ref" yaml:"sourceRef,omitempty"`
			} `tfsdk:"data_from" yaml:"dataFrom,omitempty"`

			RefreshInterval *string `tfsdk:"refresh_interval" yaml:"refreshInterval,omitempty"`

			SecretStoreRef *struct {
				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"secret_store_ref" yaml:"secretStoreRef,omitempty"`

			Target *struct {
				CreationPolicy *string `tfsdk:"creation_policy" yaml:"creationPolicy,omitempty"`

				DeletionPolicy *string `tfsdk:"deletion_policy" yaml:"deletionPolicy,omitempty"`

				Immutable *bool `tfsdk:"immutable" yaml:"immutable,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Template *struct {
					Data *map[string]string `tfsdk:"data" yaml:"data,omitempty"`

					EngineVersion *string `tfsdk:"engine_version" yaml:"engineVersion,omitempty"`

					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					TemplateFrom *[]struct {
						ConfigMap *struct {
							Items *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`
							} `tfsdk:"items" yaml:"items,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"config_map" yaml:"configMap,omitempty"`

						Secret *struct {
							Items *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`
							} `tfsdk:"items" yaml:"items,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret" yaml:"secret,omitempty"`
					} `tfsdk:"template_from" yaml:"templateFrom,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"target" yaml:"target,omitempty"`
		} `tfsdk:"external_secret_spec" yaml:"externalSecretSpec,omitempty"`

		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

		RefreshTime *string `tfsdk:"refresh_time" yaml:"refreshTime,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewExternalSecretsIoClusterExternalSecretV1Beta1Resource() resource.Resource {
	return &ExternalSecretsIoClusterExternalSecretV1Beta1Resource{}
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_external_secrets_io_cluster_external_secret_v1beta1"
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterExternalSecret is the Schema for the clusterexternalsecrets API.",
		MarkdownDescription: "ClusterExternalSecret is the Schema for the clusterexternalsecrets API.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "ClusterExternalSecretSpec defines the desired state of ClusterExternalSecret.",
				MarkdownDescription: "ClusterExternalSecretSpec defines the desired state of ClusterExternalSecret.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"external_secret_name": {
						Description:         "The name of the external secrets to be created defaults to the name of the ClusterExternalSecret",
						MarkdownDescription: "The name of the external secrets to be created defaults to the name of the ClusterExternalSecret",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_secret_spec": {
						Description:         "The spec for the ExternalSecrets to be created",
						MarkdownDescription: "The spec for the ExternalSecrets to be created",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"data": {
								Description:         "Data defines the connection between the Kubernetes Secret keys and the Provider data",
								MarkdownDescription: "Data defines the connection between the Kubernetes Secret keys and the Provider data",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"remote_ref": {
										Description:         "RemoteRef points to the remote secret and defines which secret (version/property/..) to fetch.",
										MarkdownDescription: "RemoteRef points to the remote secret and defines which secret (version/property/..) to fetch.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"conversion_strategy": {
												Description:         "Used to define a conversion Strategy",
												MarkdownDescription: "Used to define a conversion Strategy",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"decoding_strategy": {
												Description:         "Used to define a decoding Strategy",
												MarkdownDescription: "Used to define a decoding Strategy",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the key used in the Provider, mandatory",
												MarkdownDescription: "Key is the key used in the Provider, mandatory",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"metadata_policy": {
												Description:         "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
												MarkdownDescription: "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"property": {
												Description:         "Used to select a specific property of the Provider value (if a map), if supported",
												MarkdownDescription: "Used to select a specific property of the Provider value (if a map), if supported",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"version": {
												Description:         "Used to select a specific version of the Provider value, if supported",
												MarkdownDescription: "Used to select a specific version of the Provider value, if supported",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_key": {
										Description:         "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",
										MarkdownDescription: "SecretKey defines the key in which the controller stores the value. This is the key in the Kind=Secret",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"source_ref": {
										Description:         "SourceRef allows you to override the source from which the value will pulled from.",
										MarkdownDescription: "SourceRef allows you to override the source from which the value will pulled from.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"generator_ref": {
												Description:         "GeneratorRef points to a generator custom resource in",
												MarkdownDescription: "GeneratorRef points to a generator custom resource in",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Specify the apiVersion of the generator resource",
														MarkdownDescription: "Specify the apiVersion of the generator resource",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
														MarkdownDescription: "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Specify the name of the generator resource",
														MarkdownDescription: "Specify the name of the generator resource",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"store_ref": {
												Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
												MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"kind": {
														Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
														MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name of the SecretStore resource",
														MarkdownDescription: "Name of the SecretStore resource",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_from": {
								Description:         "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
								MarkdownDescription: "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"extract": {
										Description:         "Used to extract multiple key/value pairs from one secret Note: Extract does not support sourceRef.Generator or sourceRef.GeneratorRef.",
										MarkdownDescription: "Used to extract multiple key/value pairs from one secret Note: Extract does not support sourceRef.Generator or sourceRef.GeneratorRef.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"conversion_strategy": {
												Description:         "Used to define a conversion Strategy",
												MarkdownDescription: "Used to define a conversion Strategy",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"decoding_strategy": {
												Description:         "Used to define a decoding Strategy",
												MarkdownDescription: "Used to define a decoding Strategy",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the key used in the Provider, mandatory",
												MarkdownDescription: "Key is the key used in the Provider, mandatory",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"metadata_policy": {
												Description:         "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",
												MarkdownDescription: "Policy for fetching tags/labels from provider secrets, possible options are Fetch, None. Defaults to None",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"property": {
												Description:         "Used to select a specific property of the Provider value (if a map), if supported",
												MarkdownDescription: "Used to select a specific property of the Provider value (if a map), if supported",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"version": {
												Description:         "Used to select a specific version of the Provider value, if supported",
												MarkdownDescription: "Used to select a specific version of the Provider value, if supported",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"find": {
										Description:         "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",
										MarkdownDescription: "Used to find secrets based on tags or regular expressions Note: Find does not support sourceRef.Generator or sourceRef.GeneratorRef.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"conversion_strategy": {
												Description:         "Used to define a conversion Strategy",
												MarkdownDescription: "Used to define a conversion Strategy",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"decoding_strategy": {
												Description:         "Used to define a decoding Strategy",
												MarkdownDescription: "Used to define a decoding Strategy",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Finds secrets based on the name.",
												MarkdownDescription: "Finds secrets based on the name.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"regexp": {
														Description:         "Finds secrets base",
														MarkdownDescription: "Finds secrets base",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "A root path to start the find operations.",
												MarkdownDescription: "A root path to start the find operations.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tags": {
												Description:         "Find secrets based on tags.",
												MarkdownDescription: "Find secrets based on tags.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rewrite": {
										Description:         "Used to rewrite secret Keys after getting them from the secret Provider Multiple Rewrite operations can be provided. They are applied in a layered order (first to last)",
										MarkdownDescription: "Used to rewrite secret Keys after getting them from the secret Provider Multiple Rewrite operations can be provided. They are applied in a layered order (first to last)",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"regexp": {
												Description:         "Used to rewrite with regular expressions. The resulting key will be the output of a regexp.ReplaceAll operation.",
												MarkdownDescription: "Used to rewrite with regular expressions. The resulting key will be the output of a regexp.ReplaceAll operation.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"source": {
														Description:         "Used to define the regular expression of a re.Compiler.",
														MarkdownDescription: "Used to define the regular expression of a re.Compiler.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"target": {
														Description:         "Used to define the target pattern of a ReplaceAll operation.",
														MarkdownDescription: "Used to define the target pattern of a ReplaceAll operation.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_ref": {
										Description:         "SourceRef points to a store or generator which contains secret values ready to use. Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values",
										MarkdownDescription: "SourceRef points to a store or generator which contains secret values ready to use. Use this in combination with Extract or Find pull values out of a specific SecretStore. When sourceRef points to a generator Extract or Find is not supported. The generator returns a static map of values",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"generator_ref": {
												Description:         "GeneratorRef points to a generator custom resource in",
												MarkdownDescription: "GeneratorRef points to a generator custom resource in",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "Specify the apiVersion of the generator resource",
														MarkdownDescription: "Specify the apiVersion of the generator resource",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",
														MarkdownDescription: "Specify the Kind of the resource, e.g. Password, ACRAccessToken etc.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Specify the name of the generator resource",
														MarkdownDescription: "Specify the name of the generator resource",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"store_ref": {
												Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
												MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"kind": {
														Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
														MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "Name of the SecretStore resource",
														MarkdownDescription: "Name of the SecretStore resource",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"refresh_interval": {
								Description:         "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' May be set to zero to fetch and create it once. Defaults to 1h.",
								MarkdownDescription: "RefreshInterval is the amount of time before the values are read again from the SecretStore provider Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h' May be set to zero to fetch and create it once. Defaults to 1h.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_store_ref": {
								Description:         "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",
								MarkdownDescription: "SecretStoreRef defines which SecretStore to fetch the ExternalSecret data.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"kind": {
										Description:         "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",
										MarkdownDescription: "Kind of the SecretStore resource (SecretStore or ClusterSecretStore) Defaults to 'SecretStore'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the SecretStore resource",
										MarkdownDescription: "Name of the SecretStore resource",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target": {
								Description:         "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",
								MarkdownDescription: "ExternalSecretTarget defines the Kubernetes Secret to be created There can be only one target per ExternalSecret.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"creation_policy": {
										Description:         "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",
										MarkdownDescription: "CreationPolicy defines rules on how to create the resulting Secret Defaults to 'Owner'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Owner", "Orphan", "Merge", "None"),
										},
									},

									"deletion_policy": {
										Description:         "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",
										MarkdownDescription: "DeletionPolicy defines rules on how to delete the resulting Secret Defaults to 'Retain'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Delete", "Merge", "Retain"),
										},
									},

									"immutable": {
										Description:         "Immutable defines if the final secret will be immutable",
										MarkdownDescription: "Immutable defines if the final secret will be immutable",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",
										MarkdownDescription: "Name defines the name of the Secret resource to be managed This field is immutable Defaults to the .metadata.name of the ExternalSecret resource",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"template": {
										Description:         "Template defines a blueprint for the created Secret resource.",
										MarkdownDescription: "Template defines a blueprint for the created Secret resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"data": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"engine_version": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"metadata": {
												Description:         "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",
												MarkdownDescription: "ExternalSecretTemplateMetadata defines metadata fields for the Secret blueprint.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"template_from": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"items": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("secret")),
														},
													},

													"secret": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"items": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															schemavalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("config_map")),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"namespace_selector": {
						Description:         "The labels to select by to find the Namespaces to create the ExternalSecrets in.",
						MarkdownDescription: "The labels to select by to find the Namespaces to create the ExternalSecrets in.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"refresh_time": {
						Description:         "The time in which the controller should reconcile it's objects and recheck namespaces for labels.",
						MarkdownDescription: "The time in which the controller should reconcile it's objects and recheck namespaces for labels.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_external_secrets_io_cluster_external_secret_v1beta1")

	var state ExternalSecretsIoClusterExternalSecretV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExternalSecretsIoClusterExternalSecretV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("external-secrets.io/v1beta1")
	goModel.Kind = utilities.Ptr("ClusterExternalSecret")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_external_secrets_io_cluster_external_secret_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_external_secrets_io_cluster_external_secret_v1beta1")

	var state ExternalSecretsIoClusterExternalSecretV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExternalSecretsIoClusterExternalSecretV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("external-secrets.io/v1beta1")
	goModel.Kind = utilities.Ptr("ClusterExternalSecret")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ExternalSecretsIoClusterExternalSecretV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_external_secrets_io_cluster_external_secret_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
