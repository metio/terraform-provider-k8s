/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package secrets_stackable_tech_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &SecretsStackableTechSecretClassV1Alpha1Manifest{}
)

func NewSecretsStackableTechSecretClassV1Alpha1Manifest() datasource.DataSource {
	return &SecretsStackableTechSecretClassV1Alpha1Manifest{}
}

type SecretsStackableTechSecretClassV1Alpha1Manifest struct{}

type SecretsStackableTechSecretClassV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Backend *struct {
			AutoTls *struct {
				Ca *struct {
					AutoGenerate          *bool   `tfsdk:"auto_generate" json:"autoGenerate,omitempty"`
					CaCertificateLifetime *string `tfsdk:"ca_certificate_lifetime" json:"caCertificateLifetime,omitempty"`
					Secret                *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"ca" json:"ca,omitempty"`
				MaxCertificateLifetime *string `tfsdk:"max_certificate_lifetime" json:"maxCertificateLifetime,omitempty"`
			} `tfsdk:"auto_tls" json:"autoTls,omitempty"`
			ExperimentalCertManager *struct {
				DefaultCertificateLifetime *string `tfsdk:"default_certificate_lifetime" json:"defaultCertificateLifetime,omitempty"`
				Issuer                     *struct {
					Kind *string `tfsdk:"kind" json:"kind,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"issuer" json:"issuer,omitempty"`
			} `tfsdk:"experimental_cert_manager" json:"experimentalCertManager,omitempty"`
			K8sSearch *struct {
				SearchNamespace *struct {
					Name *string            `tfsdk:"name" json:"name,omitempty"`
					Pod  *map[string]string `tfsdk:"pod" json:"pod,omitempty"`
				} `tfsdk:"search_namespace" json:"searchNamespace,omitempty"`
			} `tfsdk:"k8s_search" json:"k8sSearch,omitempty"`
			KerberosKeytab *struct {
				Admin *struct {
					ActiveDirectory *struct {
						ExperimentalGenerateSamAccountName *struct {
							Prefix      *string `tfsdk:"prefix" json:"prefix,omitempty"`
							TotalLength *int64  `tfsdk:"total_length" json:"totalLength,omitempty"`
						} `tfsdk:"experimental_generate_sam_account_name" json:"experimentalGenerateSamAccountName,omitempty"`
						LdapServer      *string `tfsdk:"ldap_server" json:"ldapServer,omitempty"`
						LdapTlsCaSecret *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"ldap_tls_ca_secret" json:"ldapTlsCaSecret,omitempty"`
						PasswordCacheSecret *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"password_cache_secret" json:"passwordCacheSecret,omitempty"`
						SchemaDistinguishedName *string `tfsdk:"schema_distinguished_name" json:"schemaDistinguishedName,omitempty"`
						UserDistinguishedName   *string `tfsdk:"user_distinguished_name" json:"userDistinguishedName,omitempty"`
					} `tfsdk:"active_directory" json:"activeDirectory,omitempty"`
					Mit *struct {
						KadminServer *string `tfsdk:"kadmin_server" json:"kadminServer,omitempty"`
					} `tfsdk:"mit" json:"mit,omitempty"`
				} `tfsdk:"admin" json:"admin,omitempty"`
				AdminKeytabSecret *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"admin_keytab_secret" json:"adminKeytabSecret,omitempty"`
				AdminPrincipal *string `tfsdk:"admin_principal" json:"adminPrincipal,omitempty"`
				Kdc            *string `tfsdk:"kdc" json:"kdc,omitempty"`
				RealmName      *string `tfsdk:"realm_name" json:"realmName,omitempty"`
			} `tfsdk:"kerberos_keytab" json:"kerberosKeytab,omitempty"`
		} `tfsdk:"backend" json:"backend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecretsStackableTechSecretClassV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_secrets_stackable_tech_secret_class_v1alpha1_manifest"
}

func (r *SecretsStackableTechSecretClassV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for SecretClassSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for SecretClassSpec via 'CustomResource'",
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
				Description:         "A [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) is a cluster-global Kubernetes resource that defines a category of secrets that the Secret Operator knows how to provision.",
				MarkdownDescription: "A [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) is a cluster-global Kubernetes resource that defines a category of secrets that the Secret Operator knows how to provision.",
				Attributes: map[string]schema.Attribute{
					"backend": schema.SingleNestedAttribute{
						Description:         "Each SecretClass is associated with a single [backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend), which dictates the mechanism for issuing that kind of Secret.",
						MarkdownDescription: "Each SecretClass is associated with a single [backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend), which dictates the mechanism for issuing that kind of Secret.",
						Attributes: map[string]schema.Attribute{
							"auto_tls": schema.SingleNestedAttribute{
								Description:         "The ['autoTls' backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-autotls) issues a TLS certificate signed by the Secret Operator. The certificate authority can be provided by the administrator, or managed automatically by the Secret Operator. A new certificate and keypair will be generated and signed for each Pod, keys or certificates are never reused.",
								MarkdownDescription: "The ['autoTls' backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-autotls) issues a TLS certificate signed by the Secret Operator. The certificate authority can be provided by the administrator, or managed automatically by the Secret Operator. A new certificate and keypair will be generated and signed for each Pod, keys or certificates are never reused.",
								Attributes: map[string]schema.Attribute{
									"ca": schema.SingleNestedAttribute{
										Description:         "Configures the certificate authority used to issue Pod certificates.",
										MarkdownDescription: "Configures the certificate authority used to issue Pod certificates.",
										Attributes: map[string]schema.Attribute{
											"auto_generate": schema.BoolAttribute{
												Description:         "Whether the certificate authority should be managed by Secret Operator, including being generated if it does not already exist.",
												MarkdownDescription: "Whether the certificate authority should be managed by Secret Operator, including being generated if it does not already exist.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ca_certificate_lifetime": schema.StringAttribute{
												Description:         "The lifetime of each generated certificate authority. Should always be more than double 'maxCertificateLifetime'. If 'autoGenerate: true' then the Secret Operator will prepare a new CA certificate the old CA approaches expiration. If 'autoGenerate: false' then the Secret Operator will log a warning instead.",
												MarkdownDescription: "The lifetime of each generated certificate authority. Should always be more than double 'maxCertificateLifetime'. If 'autoGenerate: true' then the Secret Operator will prepare a new CA certificate the old CA approaches expiration. If 'autoGenerate: false' then the Secret Operator will log a warning instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret": schema.SingleNestedAttribute{
												Description:         "Reference (name and namespace) to a Kubernetes Secret object where the CA certificate and key is stored in the keys 'ca.crt' and 'ca.key' respectively.",
												MarkdownDescription: "Reference (name and namespace) to a Kubernetes Secret object where the CA certificate and key is stored in the keys 'ca.crt' and 'ca.key' respectively.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the Secret being referred to.",
														MarkdownDescription: "Name of the Secret being referred to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the Secret being referred to.",
														MarkdownDescription: "Namespace of the Secret being referred to.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"max_certificate_lifetime": schema.StringAttribute{
										Description:         "Maximum lifetime the created certificates are allowed to have. In case consumers request a longer lifetime than allowed by this setting, the lifetime will be the minimum of both, so this setting takes precedence. The default value is 15 days.",
										MarkdownDescription: "Maximum lifetime the created certificates are allowed to have. In case consumers request a longer lifetime than allowed by this setting, the lifetime will be the minimum of both, so this setting takes precedence. The default value is 15 days.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"experimental_cert_manager": schema.SingleNestedAttribute{
								Description:         "The ['experimentalCertManager' backend][1] injects a TLS certificate issued by [cert-manager](https://cert-manager.io/). A new certificate will be requested the first time it is used by a Pod, it will be reused after that (subject to cert-manager renewal rules). [1]: https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-certmanager",
								MarkdownDescription: "The ['experimentalCertManager' backend][1] injects a TLS certificate issued by [cert-manager](https://cert-manager.io/). A new certificate will be requested the first time it is used by a Pod, it will be reused after that (subject to cert-manager renewal rules). [1]: https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-certmanager",
								Attributes: map[string]schema.Attribute{
									"default_certificate_lifetime": schema.StringAttribute{
										Description:         "The default lifetime of certificates. Defaults to 1 day. This may need to be increased for external issuers that impose rate limits (such as Let's Encrypt).",
										MarkdownDescription: "The default lifetime of certificates. Defaults to 1 day. This may need to be increased for external issuers that impose rate limits (such as Let's Encrypt).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"issuer": schema.SingleNestedAttribute{
										Description:         "A reference to the cert-manager issuer that the certificates should be requested from.",
										MarkdownDescription: "A reference to the cert-manager issuer that the certificates should be requested from.",
										Attributes: map[string]schema.Attribute{
											"kind": schema.StringAttribute{
												Description:         "The kind of the issuer, Issuer or ClusterIssuer. If Issuer then it must be in the same namespace as the Pods using it.",
												MarkdownDescription: "The kind of the issuer, Issuer or ClusterIssuer. If Issuer then it must be in the same namespace as the Pods using it.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Issuer", "ClusterIssuer"),
												},
											},

											"name": schema.StringAttribute{
												Description:         "The name of the issuer.",
												MarkdownDescription: "The name of the issuer.",
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

							"k8s_search": schema.SingleNestedAttribute{
								Description:         "The ['k8sSearch' backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-k8ssearch) can be used to mount Secrets across namespaces into Pods.",
								MarkdownDescription: "The ['k8sSearch' backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-k8ssearch) can be used to mount Secrets across namespaces into Pods.",
								Attributes: map[string]schema.Attribute{
									"search_namespace": schema.SingleNestedAttribute{
										Description:         "Configures the namespace searched for Secret objects.",
										MarkdownDescription: "Configures the namespace searched for Secret objects.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "The Secret objects are located in a single global namespace. Should be used for secrets that are provisioned by the cluster administrator.",
												MarkdownDescription: "The Secret objects are located in a single global namespace. Should be used for secrets that are provisioned by the cluster administrator.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod": schema.MapAttribute{
												Description:         "The Secret objects are located in the same namespace as the Pod object. Should be used for Secrets that are provisioned by the application administrator.",
												MarkdownDescription: "The Secret objects are located in the same namespace as the Pod object. Should be used for Secrets that are provisioned by the application administrator.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
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

							"kerberos_keytab": schema.SingleNestedAttribute{
								Description:         "The ['kerberosKeytab' backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-kerberoskeytab) creates a Kerberos keytab file for a selected realm. The Kerberos KDC and administrator credentials must be provided by the administrator.",
								MarkdownDescription: "The ['kerberosKeytab' backend](https://docs.stackable.tech/home/nightly/secret-operator/secretclass#backend-kerberoskeytab) creates a Kerberos keytab file for a selected realm. The Kerberos KDC and administrator credentials must be provided by the administrator.",
								Attributes: map[string]schema.Attribute{
									"admin": schema.SingleNestedAttribute{
										Description:         "Kerberos admin configuration settings.",
										MarkdownDescription: "Kerberos admin configuration settings.",
										Attributes: map[string]schema.Attribute{
											"active_directory": schema.SingleNestedAttribute{
												Description:         "Credentials should be provisioned in a Microsoft Active Directory domain.",
												MarkdownDescription: "Credentials should be provisioned in a Microsoft Active Directory domain.",
												Attributes: map[string]schema.Attribute{
													"experimental_generate_sam_account_name": schema.SingleNestedAttribute{
														Description:         "Allows samAccountName generation for new accounts to be customized. Note that setting this field (even if empty) makes the Secret Operator take over the generation duty from the domain controller.",
														MarkdownDescription: "Allows samAccountName generation for new accounts to be customized. Note that setting this field (even if empty) makes the Secret Operator take over the generation duty from the domain controller.",
														Attributes: map[string]schema.Attribute{
															"prefix": schema.StringAttribute{
																Description:         "A prefix to be prepended to generated samAccountNames.",
																MarkdownDescription: "A prefix to be prepended to generated samAccountNames.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"total_length": schema.Int64Attribute{
																Description:         "The total length of generated samAccountNames, _including_ 'prefix'. Must be larger than the length of 'prefix', but at most '20'. Note that this should be as large as possible, to minimize the risk of collisions.",
																MarkdownDescription: "The total length of generated samAccountNames, _including_ 'prefix'. Must be larger than the length of 'prefix', but at most '20'. Note that this should be as large as possible, to minimize the risk of collisions.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ldap_server": schema.StringAttribute{
														Description:         "An AD LDAP server, such as the AD Domain Controller. This must match the server’s FQDN, or GSSAPI authentication will fail.",
														MarkdownDescription: "An AD LDAP server, such as the AD Domain Controller. This must match the server’s FQDN, or GSSAPI authentication will fail.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"ldap_tls_ca_secret": schema.SingleNestedAttribute{
														Description:         "Reference (name and namespace) to a Kubernetes Secret object containing the TLS CA (in 'ca.crt') that the LDAP server’s certificate should be authenticated against.",
														MarkdownDescription: "Reference (name and namespace) to a Kubernetes Secret object containing the TLS CA (in 'ca.crt') that the LDAP server’s certificate should be authenticated against.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the Secret being referred to.",
																MarkdownDescription: "Name of the Secret being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the Secret being referred to.",
																MarkdownDescription: "Namespace of the Secret being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"password_cache_secret": schema.SingleNestedAttribute{
														Description:         "Reference (name and namespace) to a Kubernetes Secret object where workload passwords will be stored. This must not be accessible to end users.",
														MarkdownDescription: "Reference (name and namespace) to a Kubernetes Secret object where workload passwords will be stored. This must not be accessible to end users.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the Secret being referred to.",
																MarkdownDescription: "Name of the Secret being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the Secret being referred to.",
																MarkdownDescription: "Namespace of the Secret being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"schema_distinguished_name": schema.StringAttribute{
														Description:         "The root Distinguished Name (DN) for AD-managed schemas, typically 'CN=Schema,CN=Configuration,{domain_dn}'.",
														MarkdownDescription: "The root Distinguished Name (DN) for AD-managed schemas, typically 'CN=Schema,CN=Configuration,{domain_dn}'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"user_distinguished_name": schema.StringAttribute{
														Description:         "The root Distinguished Name (DN) where service accounts should be provisioned, typically 'CN=Users,{domain_dn}'.",
														MarkdownDescription: "The root Distinguished Name (DN) where service accounts should be provisioned, typically 'CN=Users,{domain_dn}'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"mit": schema.SingleNestedAttribute{
												Description:         "Credentials should be provisioned in a MIT Kerberos Admin Server.",
												MarkdownDescription: "Credentials should be provisioned in a MIT Kerberos Admin Server.",
												Attributes: map[string]schema.Attribute{
													"kadmin_server": schema.StringAttribute{
														Description:         "The hostname of the Kerberos Admin Server. This should be provided by the Kerberos administrator.",
														MarkdownDescription: "The hostname of the Kerberos Admin Server. This should be provided by the Kerberos administrator.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"admin_keytab_secret": schema.SingleNestedAttribute{
										Description:         "Reference ('name' and 'namespace') to a K8s Secret object where a keytab with administrative privileges is stored in the key 'keytab'.",
										MarkdownDescription: "Reference ('name' and 'namespace') to a K8s Secret object where a keytab with administrative privileges is stored in the key 'keytab'.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the Secret being referred to.",
												MarkdownDescription: "Name of the Secret being referred to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the Secret being referred to.",
												MarkdownDescription: "Namespace of the Secret being referred to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"admin_principal": schema.StringAttribute{
										Description:         "The admin principal.",
										MarkdownDescription: "The admin principal.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"kdc": schema.StringAttribute{
										Description:         "The hostname of the Kerberos Key Distribution Center (KDC). This should be provided by the Kerberos administrator.",
										MarkdownDescription: "The hostname of the Kerberos Key Distribution Center (KDC). This should be provided by the Kerberos administrator.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"realm_name": schema.StringAttribute{
										Description:         "The name of the Kerberos realm. This should be provided by the Kerberos administrator.",
										MarkdownDescription: "The name of the Kerberos realm. This should be provided by the Kerberos administrator.",
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
	}
}

func (r *SecretsStackableTechSecretClassV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_secrets_stackable_tech_secret_class_v1alpha1_manifest")

	var model SecretsStackableTechSecretClassV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("secrets.stackable.tech/v1alpha1")
	model.Kind = pointer.String("SecretClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
