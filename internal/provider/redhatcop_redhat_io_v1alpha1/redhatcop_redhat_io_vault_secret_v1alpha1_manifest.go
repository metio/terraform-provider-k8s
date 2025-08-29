/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package redhatcop_redhat_io_v1alpha1

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
	_ datasource.DataSource = &RedhatcopRedhatIoVaultSecretV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoVaultSecretV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoVaultSecretV1Alpha1Manifest{}
}

type RedhatcopRedhatIoVaultSecretV1Alpha1Manifest struct{}

type RedhatcopRedhatIoVaultSecretV1Alpha1ManifestData struct {
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
		Output *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			StringData  *map[string]string `tfsdk:"string_data" json:"stringData,omitempty"`
			Type        *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"output" json:"output,omitempty"`
		RefreshPeriod          *string `tfsdk:"refresh_period" json:"refreshPeriod,omitempty"`
		RefreshThreshold       *int64  `tfsdk:"refresh_threshold" json:"refreshThreshold,omitempty"`
		VaultSecretDefinitions *[]struct {
			Authentication *struct {
				Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Path           *string `tfsdk:"path" json:"path,omitempty"`
				Role           *string `tfsdk:"role" json:"role,omitempty"`
				ServiceAccount *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			Connection *struct {
				Address    *string `tfsdk:"address" json:"address,omitempty"`
				MaxRetries *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
				TLSConfig  *struct {
					Cacert     *string `tfsdk:"cacert" json:"cacert,omitempty"`
					SkipVerify *bool   `tfsdk:"skip_verify" json:"skipVerify,omitempty"`
					TlsSecret  *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
					TlsServerName *string `tfsdk:"tls_server_name" json:"tlsServerName,omitempty"`
				} `tfsdk:"t_ls_config" json:"tLSConfig,omitempty"`
				TimeOut *string `tfsdk:"time_out" json:"timeOut,omitempty"`
			} `tfsdk:"connection" json:"connection,omitempty"`
			Name           *string            `tfsdk:"name" json:"name,omitempty"`
			Path           *string            `tfsdk:"path" json:"path,omitempty"`
			RequestPayload *map[string]string `tfsdk:"request_payload" json:"requestPayload,omitempty"`
			RequestType    *string            `tfsdk:"request_type" json:"requestType,omitempty"`
		} `tfsdk:"vault_secret_definitions" json:"vaultSecretDefinitions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoVaultSecretV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_vault_secret_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoVaultSecretV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VaultSecret is the Schema for the vaultsecrets API",
		MarkdownDescription: "VaultSecret is the Schema for the vaultsecrets API",
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
				Description:         "VaultSecretSpec defines the desired state of VaultSecret",
				MarkdownDescription: "VaultSecretSpec defines the desired state of VaultSecret",
				Attributes: map[string]schema.Attribute{
					"output": schema.SingleNestedAttribute{
						Description:         "TemplatizedK8sSecret is the formatted K8s Secret created by templating from the Vault KV secrets.",
						MarkdownDescription: "TemplatizedK8sSecret is the formatted K8s Secret created by templating from the Vault KV secrets.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations are annotations to add to the final K8s Secret.",
								MarkdownDescription: "Annotations are annotations to add to the final K8s Secret.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels are labels to add to the final K8s Secret.",
								MarkdownDescription: "Labels are labels to add to the final K8s Secret.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the K8s Secret name to output to.",
								MarkdownDescription: "Name is the K8s Secret name to output to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"string_data": schema.MapAttribute{
								Description:         "StringData is the K8s Secret stringData and allows specifying non-binary secret data in string form with go templating support to transform the Vault KV secrets into a formatted K8s Secret. The Sprig template library and Helm functions (like toYaml) are supported.",
								MarkdownDescription: "StringData is the K8s Secret stringData and allows specifying non-binary secret data in string form with go templating support to transform the Vault KV secrets into a formatted K8s Secret. The Sprig template library and Helm functions (like toYaml) are supported.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the K8s Secret type to output to.",
								MarkdownDescription: "Type is the K8s Secret type to output to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"refresh_period": schema.StringAttribute{
						Description:         "RefreshPeriod if specified, the operator will refresh the secret with the given frequency. This takes precedence over any vault secret lease duration and can be used to force a refresh.",
						MarkdownDescription: "RefreshPeriod if specified, the operator will refresh the secret with the given frequency. This takes precedence over any vault secret lease duration and can be used to force a refresh.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"refresh_threshold": schema.Int64Attribute{
						Description:         "RefreshThreshold if specified, will instruct the operator to refresh when a percentage of the lease duration is met when there is no RefreshPeriod specified. This is particularly useful for controlling when dynamic secrets should be refreshed before the lease duration is exceeded. The default is 90, meaning the secret would refresh after 90% of the time has passed from the vault secret's lease duration.",
						MarkdownDescription: "RefreshThreshold if specified, will instruct the operator to refresh when a percentage of the lease duration is met when there is no RefreshPeriod specified. This is particularly useful for controlling when dynamic secrets should be refreshed before the lease duration is exceeded. The default is 90, meaning the secret would refresh after 90% of the time has passed from the vault secret's lease duration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vault_secret_definitions": schema.ListNestedAttribute{
						Description:         "VaultSecretDefinitions are the secrets in Vault.",
						MarkdownDescription: "VaultSecretDefinitions are the secrets in Vault.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"authentication": schema.SingleNestedAttribute{
									Description:         "Authentication is the kube auth configuraiton to be used to execute this request",
									MarkdownDescription: "Authentication is the kube auth configuraiton to be used to execute this request",
									Attributes: map[string]schema.Attribute{
										"namespace": schema.StringAttribute{
											Description:         "Namespace is the Vault namespace to be used in all the operations withing this connection/authentication. Only available in Vault Enterprise.",
											MarkdownDescription: "Namespace is the Vault namespace to be used in all the operations withing this connection/authentication. Only available in Vault Enterprise.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path is the path of the role used for this kube auth authentication. The operator will try to authenticate at {[namespace/]}auth/{spec.path}",
											MarkdownDescription: "Path is the path of the role used for this kube auth authentication. The operator will try to authenticate at {[namespace/]}auth/{spec.path}",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
											},
										},

										"role": schema.StringAttribute{
											Description:         "Role the role to be used during authentication",
											MarkdownDescription: "Role the role to be used during authentication",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_account": schema.SingleNestedAttribute{
											Description:         "ServiceAccount is the service account used for the kube auth authentication",
											MarkdownDescription: "ServiceAccount is the service account used for the kube auth authentication",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

								"connection": schema.SingleNestedAttribute{
									Description:         "Connection represents the information needed to connect to Vault. This operator uses the standard Vault environment variables to connect to Vault. If you need to override those settings and for example connect to a different Vault instance, you can do with this section of the CR.",
									MarkdownDescription: "Connection represents the information needed to connect to Vault. This operator uses the standard Vault environment variables to connect to Vault. If you need to override those settings and for example connect to a different Vault instance, you can do with this section of the CR.",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "Address Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/",
											MarkdownDescription: "Address Address of the Vault server expressed as a URL and port, for example: https://127.0.0.1:8200/",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_retries": schema.Int64Attribute{
											Description:         "MaxRetries Maximum number of retries when certain error codes are encountered. The default is 2, for three total attempts. Set this to 0 or less to disable retrying. Error codes that are retried are 412 (client consistency requirement not satisfied) and all 5xx except for 501 (not implemented).",
											MarkdownDescription: "MaxRetries Maximum number of retries when certain error codes are encountered. The default is 2, for three total attempts. Set this to 0 or less to disable retrying. Error codes that are retried are 412 (client consistency requirement not satisfied) and all 5xx except for 501 (not implemented).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"t_ls_config": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"cacert": schema.StringAttribute{
													Description:         "Cacert Path to a PEM-encoded CA certificate file on the local disk. This file is used to verify the Vault server's SSL certificate. This environment variable takes precedence over a cert passed via the secret.",
													MarkdownDescription: "Cacert Path to a PEM-encoded CA certificate file on the local disk. This file is used to verify the Vault server's SSL certificate. This environment variable takes precedence over a cert passed via the secret.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"skip_verify": schema.BoolAttribute{
													Description:         "SkipVerify Do not verify Vault's presented certificate before communicating with it. Setting this variable is not recommended and voids Vault's security model.",
													MarkdownDescription: "SkipVerify Do not verify Vault's presented certificate before communicating with it. Setting this variable is not recommended and voids Vault's security model.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls_secret": schema.SingleNestedAttribute{
													Description:         "TLSSecret namespace-local secret containing the tls material for the connection. the expected keys for the secret are: ca bundle -> 'ca.crt', certificate -> 'tls.crt', key -> 'tls.key'",
													MarkdownDescription: "TLSSecret namespace-local secret containing the tls material for the connection. the expected keys for the secret are: ca bundle -> 'ca.crt', certificate -> 'tls.crt', key -> 'tls.key'",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"tls_server_name": schema.StringAttribute{
													Description:         "TLSServerName Name to use as the SNI host when connecting via TLS.",
													MarkdownDescription: "TLSServerName Name to use as the SNI host when connecting via TLS.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"time_out": schema.StringAttribute{
											Description:         "Timeout Timeout variable. The default value is 60s.",
											MarkdownDescription: "Timeout Timeout variable. The default value is 60s.",
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
									Description:         "Name is an arbitrary, but unique, name for this KV Vault secret and referenced when templating.",
									MarkdownDescription: "Name is an arbitrary, but unique, name for this KV Vault secret and referenced when templating.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"path": schema.StringAttribute{
									Description:         "Path is the path of the secret.",
									MarkdownDescription: "Path is the path of the secret.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
									},
								},

								"request_payload": schema.MapAttribute{
									Description:         "RequestPayload for POST type of requests, this field contains the payload of the request. Not used for GET requests.",
									MarkdownDescription: "RequestPayload for POST type of requests, this field contains the payload of the request. Not used for GET requests.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"request_type": schema.StringAttribute{
									Description:         "RequestType the type of request needed to retrieve a secret. Normally a GET, but some secret engnes require a POST.",
									MarkdownDescription: "RequestType the type of request needed to retrieve a secret. Normally a GET, but some secret engnes require a POST.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("GET", "POST"),
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
	}
}

func (r *RedhatcopRedhatIoVaultSecretV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_vault_secret_v1alpha1_manifest")

	var model RedhatcopRedhatIoVaultSecretV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("VaultSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
