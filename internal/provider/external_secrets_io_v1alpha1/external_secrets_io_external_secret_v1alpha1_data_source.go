/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package external_secrets_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
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
	_ datasource.DataSource              = &ExternalSecretsIoExternalSecretV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ExternalSecretsIoExternalSecretV1Alpha1DataSource{}
)

func NewExternalSecretsIoExternalSecretV1Alpha1DataSource() datasource.DataSource {
	return &ExternalSecretsIoExternalSecretV1Alpha1DataSource{}
}

type ExternalSecretsIoExternalSecretV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ExternalSecretsIoExternalSecretV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *ExternalSecretsIoExternalSecretV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_external_secrets_io_external_secret_v1alpha1"
}

func (r *ExternalSecretsIoExternalSecretV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

					"data_from": schema.ListNestedAttribute{
						Description:         "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
						MarkdownDescription: "DataFrom is used to fetch all properties from a specific Provider data If multiple entries are specified, the Secret keys are merged in the specified order",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"conversion_strategy": schema.StringAttribute{
									Description:         "Used to define a conversion Strategy",
									MarkdownDescription: "Used to define a conversion Strategy",
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
										Description:         "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
										MarkdownDescription: "EngineVersion specifies the template engine version that should be used to compile/execute the template specified in .data and .templateFrom[].",
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
		},
	}
}

func (r *ExternalSecretsIoExternalSecretV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ExternalSecretsIoExternalSecretV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_external_secrets_io_external_secret_v1alpha1")

	var data ExternalSecretsIoExternalSecretV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "external-secrets.io", Version: "v1alpha1", Resource: "ExternalSecret"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ExternalSecretsIoExternalSecretV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("external-secrets.io/v1alpha1")
	data.Kind = pointer.String("ExternalSecret")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
