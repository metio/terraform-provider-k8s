/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package perses_dev_v1alpha1

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
	_ datasource.DataSource = &PersesDevPersesDatasourceV1Alpha1Manifest{}
)

func NewPersesDevPersesDatasourceV1Alpha1Manifest() datasource.DataSource {
	return &PersesDevPersesDatasourceV1Alpha1Manifest{}
}

type PersesDevPersesDatasourceV1Alpha1Manifest struct{}

type PersesDevPersesDatasourceV1Alpha1ManifestData struct {
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
		Client *struct {
			BasicAuth *struct {
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Password_path *string `tfsdk:"password_path" json:"password_path,omitempty"`
				Type          *string `tfsdk:"type" json:"type,omitempty"`
				Username      *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			KubernetesAuth *struct {
				Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
			} `tfsdk:"kubernetes_auth" json:"kubernetesAuth,omitempty"`
			Oauth *struct {
				AuthStyle        *int64               `tfsdk:"auth_style" json:"authStyle,omitempty"`
				ClientIDPath     *string              `tfsdk:"client_id_path" json:"clientIDPath,omitempty"`
				ClientSecretPath *string              `tfsdk:"client_secret_path" json:"clientSecretPath,omitempty"`
				EndpointParams   *map[string][]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
				Name             *string              `tfsdk:"name" json:"name,omitempty"`
				Namespace        *string              `tfsdk:"namespace" json:"namespace,omitempty"`
				Scopes           *[]string            `tfsdk:"scopes" json:"scopes,omitempty"`
				TokenURL         *string              `tfsdk:"token_url" json:"tokenURL,omitempty"`
				Type             *string              `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"oauth" json:"oauth,omitempty"`
			Tls *struct {
				CaCert *struct {
					CertPath       *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
					PrivateKeyPath *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
					Type           *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_cert" json:"caCert,omitempty"`
				Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				UserCert           *struct {
					CertPath       *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
					PrivateKeyPath *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
					Type           *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"user_cert" json:"userCert,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"client" json:"client,omitempty"`
		Config *struct {
			Default *bool `tfsdk:"default" json:"default,omitempty"`
			Display *struct {
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"display" json:"display,omitempty"`
			Plugin *struct {
				Kind *string            `tfsdk:"kind" json:"kind,omitempty"`
				Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"plugin" json:"plugin,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PersesDevPersesDatasourceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_perses_dev_perses_datasource_v1alpha1_manifest"
}

func (r *PersesDevPersesDatasourceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PersesDatasource is the Schema for the PersesDatasources API",
		MarkdownDescription: "PersesDatasource is the Schema for the PersesDatasources API",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"client": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"basic_auth": schema.SingleNestedAttribute{
								Description:         "BasicAuth basic auth config for perses client",
								MarkdownDescription: "BasicAuth basic auth config for perses client",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password_path": schema.StringAttribute{
										Description:         "Path to password",
										MarkdownDescription: "Path to password",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type source type of secret",
										MarkdownDescription: "Type source type of secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("secret", "configmap", "file"),
										},
									},

									"username": schema.StringAttribute{
										Description:         "Username for basic auth",
										MarkdownDescription: "Username for basic auth",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_auth": schema.SingleNestedAttribute{
								Description:         "KubernetesAuth configuration for perses client",
								MarkdownDescription: "KubernetesAuth configuration for perses client",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "Enable kubernetes auth for perses client",
										MarkdownDescription: "Enable kubernetes auth for perses client",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"oauth": schema.SingleNestedAttribute{
								Description:         "OAuth configuration for perses client",
								MarkdownDescription: "OAuth configuration for perses client",
								Attributes: map[string]schema.Attribute{
									"auth_style": schema.Int64Attribute{
										Description:         "AuthStyle optionally specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect.",
										MarkdownDescription: "AuthStyle optionally specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_id_path": schema.StringAttribute{
										Description:         "Path to client id",
										MarkdownDescription: "Path to client id",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_secret_path": schema.StringAttribute{
										Description:         "Path to client secret",
										MarkdownDescription: "Path to client secret",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"endpoint_params": schema.MapAttribute{
										Description:         "EndpointParams specifies additional parameters for requests to the token endpoint.",
										MarkdownDescription: "EndpointParams specifies additional parameters for requests to the token endpoint.",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scopes": schema.ListAttribute{
										Description:         "Scope specifies optional requested permissions.",
										MarkdownDescription: "Scope specifies optional requested permissions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"token_url": schema.StringAttribute{
										Description:         "TokenURL is the resource server's token endpoint URL. This is a constant specific to each server.",
										MarkdownDescription: "TokenURL is the resource server's token endpoint URL. This is a constant specific to each server.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type source type of secret",
										MarkdownDescription: "Type source type of secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("secret", "configmap", "file"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS the equivalent to the tls_config for perses client",
								MarkdownDescription: "TLS the equivalent to the tls_config for perses client",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.SingleNestedAttribute{
										Description:         "CaCert to verify the perses certificate",
										MarkdownDescription: "CaCert to verify the perses certificate",
										Attributes: map[string]schema.Attribute{
											"cert_path": schema.StringAttribute{
												Description:         "Path to Certificate",
												MarkdownDescription: "Path to Certificate",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"private_key_path": schema.StringAttribute{
												Description:         "Path to Private key certificate",
												MarkdownDescription: "Path to Private key certificate",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type source type of secret",
												MarkdownDescription: "Type source type of secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("secret", "configmap", "file"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable": schema.BoolAttribute{
										Description:         "Enable TLS connection to perses",
										MarkdownDescription: "Enable TLS connection to perses",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "InsecureSkipVerify skip verify of perses certificate",
										MarkdownDescription: "InsecureSkipVerify skip verify of perses certificate",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_cert": schema.SingleNestedAttribute{
										Description:         "UserCert client cert/key for mTLS",
										MarkdownDescription: "UserCert client cert/key for mTLS",
										Attributes: map[string]schema.Attribute{
											"cert_path": schema.StringAttribute{
												Description:         "Path to Certificate",
												MarkdownDescription: "Path to Certificate",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"private_key_path": schema.StringAttribute{
												Description:         "Path to Private key certificate",
												MarkdownDescription: "Path to Private key certificate",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type source type of secret",
												MarkdownDescription: "Type source type of secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("secret", "configmap", "file"),
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

					"config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"default": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"display": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"description": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"plugin": schema.SingleNestedAttribute{
								Description:         "Plugin will contain the datasource configuration. The data typed is available in Cue.",
								MarkdownDescription: "Plugin will contain the datasource configuration. The data typed is available in Cue.",
								Attributes: map[string]schema.Attribute{
									"kind": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"spec": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *PersesDevPersesDatasourceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_perses_dev_perses_datasource_v1alpha1_manifest")

	var model PersesDevPersesDatasourceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("perses.dev/v1alpha1")
	model.Kind = pointer.String("PersesDatasource")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
