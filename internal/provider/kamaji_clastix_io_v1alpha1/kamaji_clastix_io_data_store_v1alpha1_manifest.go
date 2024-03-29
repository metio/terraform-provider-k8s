/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kamaji_clastix_io_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KamajiClastixIoDataStoreV1Alpha1Manifest{}
)

func NewKamajiClastixIoDataStoreV1Alpha1Manifest() datasource.DataSource {
	return &KamajiClastixIoDataStoreV1Alpha1Manifest{}
}

type KamajiClastixIoDataStoreV1Alpha1Manifest struct{}

type KamajiClastixIoDataStoreV1Alpha1ManifestData struct {
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
		BasicAuth *struct {
			Password *struct {
				Content         *string `tfsdk:"content" json:"content,omitempty"`
				SecretReference *struct {
					KeyPath   *string `tfsdk:"key_path" json:"keyPath,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_reference" json:"secretReference,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Username *struct {
				Content         *string `tfsdk:"content" json:"content,omitempty"`
				SecretReference *struct {
					KeyPath   *string `tfsdk:"key_path" json:"keyPath,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_reference" json:"secretReference,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
		Driver    *string   `tfsdk:"driver" json:"driver,omitempty"`
		Endpoints *[]string `tfsdk:"endpoints" json:"endpoints,omitempty"`
		TlsConfig *struct {
			CertificateAuthority *struct {
				Certificate *struct {
					Content         *string `tfsdk:"content" json:"content,omitempty"`
					SecretReference *struct {
						KeyPath   *string `tfsdk:"key_path" json:"keyPath,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_reference" json:"secretReference,omitempty"`
				} `tfsdk:"certificate" json:"certificate,omitempty"`
				PrivateKey *struct {
					Content         *string `tfsdk:"content" json:"content,omitempty"`
					SecretReference *struct {
						KeyPath   *string `tfsdk:"key_path" json:"keyPath,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_reference" json:"secretReference,omitempty"`
				} `tfsdk:"private_key" json:"privateKey,omitempty"`
			} `tfsdk:"certificate_authority" json:"certificateAuthority,omitempty"`
			ClientCertificate *struct {
				Certificate *struct {
					Content         *string `tfsdk:"content" json:"content,omitempty"`
					SecretReference *struct {
						KeyPath   *string `tfsdk:"key_path" json:"keyPath,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_reference" json:"secretReference,omitempty"`
				} `tfsdk:"certificate" json:"certificate,omitempty"`
				PrivateKey *struct {
					Content         *string `tfsdk:"content" json:"content,omitempty"`
					SecretReference *struct {
						KeyPath   *string `tfsdk:"key_path" json:"keyPath,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_reference" json:"secretReference,omitempty"`
				} `tfsdk:"private_key" json:"privateKey,omitempty"`
			} `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
		} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KamajiClastixIoDataStoreV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kamaji_clastix_io_data_store_v1alpha1_manifest"
}

func (r *KamajiClastixIoDataStoreV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DataStore is the Schema for the datastores API.",
		MarkdownDescription: "DataStore is the Schema for the datastores API.",
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
				Description:         "DataStoreSpec defines the desired state of DataStore.",
				MarkdownDescription: "DataStoreSpec defines the desired state of DataStore.",
				Attributes: map[string]schema.Attribute{
					"basic_auth": schema.SingleNestedAttribute{
						Description:         "In case of authentication enabled for the given data store, specifies the username and password pair. This value is optional.",
						MarkdownDescription: "In case of authentication enabled for the given data store, specifies the username and password pair. This value is optional.",
						Attributes: map[string]schema.Attribute{
							"password": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"content": schema.StringAttribute{
										Description:         "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
										MarkdownDescription: "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.Base64Validator(),
										},
									},

									"secret_reference": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"key_path": schema.StringAttribute{
												Description:         "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
												MarkdownDescription: "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

							"username": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"content": schema.StringAttribute{
										Description:         "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
										MarkdownDescription: "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.Base64Validator(),
										},
									},

									"secret_reference": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"key_path": schema.StringAttribute{
												Description:         "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
												MarkdownDescription: "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

					"driver": schema.StringAttribute{
						Description:         "The driver to use to connect to the shared datastore.",
						MarkdownDescription: "The driver to use to connect to the shared datastore.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("etcd", "MySQL", "PostgreSQL"),
						},
					},

					"endpoints": schema.ListAttribute{
						Description:         "List of the endpoints to connect to the shared datastore. No need for protocol, just bare IP/FQDN and port.",
						MarkdownDescription: "List of the endpoints to connect to the shared datastore. No need for protocol, just bare IP/FQDN and port.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tls_config": schema.SingleNestedAttribute{
						Description:         "Defines the TLS/SSL configuration required to connect to the data store in a secure way.",
						MarkdownDescription: "Defines the TLS/SSL configuration required to connect to the data store in a secure way.",
						Attributes: map[string]schema.Attribute{
							"certificate_authority": schema.SingleNestedAttribute{
								Description:         "Retrieve the Certificate Authority certificate and private key, such as bare content of the file, or a SecretReference. The key reference is required since etcd authentication is based on certificates, and Kamaji is responsible in creating this.",
								MarkdownDescription: "Retrieve the Certificate Authority certificate and private key, such as bare content of the file, or a SecretReference. The key reference is required since etcd authentication is based on certificates, and Kamaji is responsible in creating this.",
								Attributes: map[string]schema.Attribute{
									"certificate": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"content": schema.StringAttribute{
												Description:         "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												MarkdownDescription: "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"secret_reference": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"key_path": schema.StringAttribute{
														Description:         "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														MarkdownDescription: "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

									"private_key": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"content": schema.StringAttribute{
												Description:         "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												MarkdownDescription: "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"secret_reference": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"key_path": schema.StringAttribute{
														Description:         "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														MarkdownDescription: "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"client_certificate": schema.SingleNestedAttribute{
								Description:         "Specifies the SSL/TLS key and private key pair used to connect to the data store.",
								MarkdownDescription: "Specifies the SSL/TLS key and private key pair used to connect to the data store.",
								Attributes: map[string]schema.Attribute{
									"certificate": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"content": schema.StringAttribute{
												Description:         "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												MarkdownDescription: "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"secret_reference": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"key_path": schema.StringAttribute{
														Description:         "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														MarkdownDescription: "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

									"private_key": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"content": schema.StringAttribute{
												Description:         "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												MarkdownDescription: "Bare content of the file, base64 encoded. It has precedence over the SecretReference value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													validators.Base64Validator(),
												},
											},

											"secret_reference": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"key_path": schema.StringAttribute{
														Description:         "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														MarkdownDescription: "Name of the key for the given Secret reference where the content is stored. This value is mandatory.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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
								Required: true,
								Optional: false,
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

func (r *KamajiClastixIoDataStoreV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kamaji_clastix_io_data_store_v1alpha1_manifest")

	var model KamajiClastixIoDataStoreV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("kamaji.clastix.io/v1alpha1")
	model.Kind = pointer.String("DataStore")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
