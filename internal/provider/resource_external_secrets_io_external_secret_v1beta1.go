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

type ExternalSecretsIoExternalSecretV1Beta1Resource struct{}

var (
	_ resource.Resource = (*ExternalSecretsIoExternalSecretV1Beta1Resource)(nil)
)

type ExternalSecretsIoExternalSecretV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ExternalSecretsIoExternalSecretV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
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
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewExternalSecretsIoExternalSecretV1Beta1Resource() resource.Resource {
	return &ExternalSecretsIoExternalSecretV1Beta1Resource{}
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_external_secrets_io_external_secret_v1beta1"
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ExternalSecret is the Schema for the external-secrets API.",
		MarkdownDescription: "ExternalSecret is the Schema for the external-secrets API.",
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

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "ExternalSecretSpec defines the desired state of ExternalSecret.",
				MarkdownDescription: "ExternalSecretSpec defines the desired state of ExternalSecret.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"data": {
						Description:         "Data defines the connection between the Kubernetes Secret keys and the Provider data",
						MarkdownDescription: "Data defines the connection between the Kubernetes Secret keys and the Provider data",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"remote_ref": {
								Description:         "ExternalSecretDataRemoteRef defines Provider data location.",
								MarkdownDescription: "ExternalSecretDataRemoteRef defines Provider data location.",

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
					},

					"data_from": {
						Description:         "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
						MarkdownDescription: "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"extract": {
								Description:         "Used to extract multiple key/value pairs from one secret",
								MarkdownDescription: "Used to extract multiple key/value pairs from one secret",

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
								Description:         "Used to find secrets based on tags or regular expressions",
								MarkdownDescription: "Used to find secrets based on tags or regular expressions",

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

						Required: true,
						Optional: false,
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

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_external_secrets_io_external_secret_v1beta1")

	var state ExternalSecretsIoExternalSecretV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExternalSecretsIoExternalSecretV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("external-secrets.io/v1beta1")
	goModel.Kind = utilities.Ptr("ExternalSecret")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_external_secrets_io_external_secret_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_external_secrets_io_external_secret_v1beta1")

	var state ExternalSecretsIoExternalSecretV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ExternalSecretsIoExternalSecretV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("external-secrets.io/v1beta1")
	goModel.Kind = utilities.Ptr("ExternalSecret")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ExternalSecretsIoExternalSecretV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_external_secrets_io_external_secret_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
