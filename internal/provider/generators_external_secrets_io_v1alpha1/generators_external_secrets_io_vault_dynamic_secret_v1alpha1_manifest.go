/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package generators_external_secrets_io_v1alpha1

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
	_ datasource.DataSource = &GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1Manifest{}
)

func NewGeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1Manifest() datasource.DataSource {
	return &GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1Manifest{}
}

type GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1Manifest struct{}

type GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1ManifestData struct {
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
		AllowEmptyResponse *bool              `tfsdk:"allow_empty_response" json:"allowEmptyResponse,omitempty"`
		Controller         *string            `tfsdk:"controller" json:"controller,omitempty"`
		Method             *string            `tfsdk:"method" json:"method,omitempty"`
		Parameters         *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		Path               *string            `tfsdk:"path" json:"path,omitempty"`
		Provider           *struct {
			Auth *struct {
				AppRole *struct {
					Path    *string `tfsdk:"path" json:"path,omitempty"`
					RoleId  *string `tfsdk:"role_id" json:"roleId,omitempty"`
					RoleRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"role_ref" json:"roleRef,omitempty"`
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"app_role" json:"appRole,omitempty"`
				Cert *struct {
					ClientCert *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"client_cert" json:"clientCert,omitempty"`
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"cert" json:"cert,omitempty"`
				Iam *struct {
					ExternalID *string `tfsdk:"external_id" json:"externalID,omitempty"`
					Jwt        *struct {
						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							Name      *string   `tfsdk:"name" json:"name,omitempty"`
							Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
					} `tfsdk:"jwt" json:"jwt,omitempty"`
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					Region    *string `tfsdk:"region" json:"region,omitempty"`
					Role      *string `tfsdk:"role" json:"role,omitempty"`
					SecretRef *struct {
						AccessKeyIDSecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"access_key_id_secret_ref" json:"accessKeyIDSecretRef,omitempty"`
						SecretAccessKeySecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"secret_access_key_secret_ref" json:"secretAccessKeySecretRef,omitempty"`
						SessionTokenSecretRef *struct {
							Key       *string `tfsdk:"key" json:"key,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"session_token_secret_ref" json:"sessionTokenSecretRef,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					VaultAwsIamServerID *string `tfsdk:"vault_aws_iam_server_id" json:"vaultAwsIamServerID,omitempty"`
					VaultRole           *string `tfsdk:"vault_role" json:"vaultRole,omitempty"`
				} `tfsdk:"iam" json:"iam,omitempty"`
				Jwt *struct {
					KubernetesServiceAccountToken *struct {
						Audiences         *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
						ExpirationSeconds *int64    `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
						ServiceAccountRef *struct {
							Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							Name      *string   `tfsdk:"name" json:"name,omitempty"`
							Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
					} `tfsdk:"kubernetes_service_account_token" json:"kubernetesServiceAccountToken,omitempty"`
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					Role      *string `tfsdk:"role" json:"role,omitempty"`
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"jwt" json:"jwt,omitempty"`
				Kubernetes *struct {
					MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					Role      *string `tfsdk:"role" json:"role,omitempty"`
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					ServiceAccountRef *struct {
						Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
						Name      *string   `tfsdk:"name" json:"name,omitempty"`
						Namespace *string   `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service_account_ref" json:"serviceAccountRef,omitempty"`
				} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
				Ldap *struct {
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"ldap" json:"ldap,omitempty"`
				Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
				TokenSecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"token_secret_ref" json:"tokenSecretRef,omitempty"`
				UserPass *struct {
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					SecretRef *struct {
						Key       *string `tfsdk:"key" json:"key,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"user_pass" json:"userPass,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			CaBundle   *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			CaProvider *struct {
				Key       *string `tfsdk:"key" json:"key,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Type      *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"ca_provider" json:"caProvider,omitempty"`
			CheckAndSet *struct {
				Required *bool `tfsdk:"required" json:"required,omitempty"`
			} `tfsdk:"check_and_set" json:"checkAndSet,omitempty"`
			ForwardInconsistent *bool              `tfsdk:"forward_inconsistent" json:"forwardInconsistent,omitempty"`
			Headers             *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
			Namespace           *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			Path                *string            `tfsdk:"path" json:"path,omitempty"`
			ReadYourWrites      *bool              `tfsdk:"read_your_writes" json:"readYourWrites,omitempty"`
			Server              *string            `tfsdk:"server" json:"server,omitempty"`
			Tls                 *struct {
				CertSecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
				KeySecretRef *struct {
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"key_secret_ref" json:"keySecretRef,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"provider" json:"provider,omitempty"`
		ResultType    *string `tfsdk:"result_type" json:"resultType,omitempty"`
		RetrySettings *struct {
			MaxRetries    *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
			RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
		} `tfsdk:"retry_settings" json:"retrySettings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_generators_external_secrets_io_vault_dynamic_secret_v1alpha1_manifest"
}

func (r *GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
					"allow_empty_response": schema.BoolAttribute{
						Description:         "Do not fail if no secrets are found. Useful for requests where no data is expected.",
						MarkdownDescription: "Do not fail if no secrets are found. Useful for requests where no data is expected.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"controller": schema.StringAttribute{
						Description:         "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters VDS based on this property",
						MarkdownDescription: "Used to select the correct ESO controller (think: ingress.ingressClassName) The ESO controller is instantiated with a specific controller name and filters VDS based on this property",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"method": schema.StringAttribute{
						Description:         "Vault API method to use (GET/POST/other)",
						MarkdownDescription: "Vault API method to use (GET/POST/other)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters": schema.MapAttribute{
						Description:         "Parameters to pass to Vault write (for non-GET methods)",
						MarkdownDescription: "Parameters to pass to Vault write (for non-GET methods)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Vault path to obtain the dynamic secret from",
						MarkdownDescription: "Vault path to obtain the dynamic secret from",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"provider": schema.SingleNestedAttribute{
						Description:         "Vault provider common spec",
						MarkdownDescription: "Vault provider common spec",
						Attributes: map[string]schema.Attribute{
							"auth": schema.SingleNestedAttribute{
								Description:         "Auth configures how secret-manager authenticates with the Vault server.",
								MarkdownDescription: "Auth configures how secret-manager authenticates with the Vault server.",
								Attributes: map[string]schema.Attribute{
									"app_role": schema.SingleNestedAttribute{
										Description:         "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
										MarkdownDescription: "AppRole authenticates with Vault using the App Role auth mechanism, with the role and secret stored in a Kubernetes Secret resource.",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",
												MarkdownDescription: "Path where the App Role authentication backend is mounted in Vault, e.g: 'approle'",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"role_id": schema.StringAttribute{
												Description:         "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
												MarkdownDescription: "RoleID configured in the App Role authentication backend when setting up the authentication backend in Vault.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"role_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a key in a Secret that contains the App Role ID used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role id.",
												MarkdownDescription: "Reference to a key in a Secret that contains the App Role ID used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role id.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
												MarkdownDescription: "Reference to a key in a Secret that contains the App Role secret used to authenticate with Vault. The 'key' field must be specified and denotes which entry within the Secret resource is used as the app role secret.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
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

									"cert": schema.SingleNestedAttribute{
										Description:         "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
										MarkdownDescription: "Cert authenticates with TLS Certificates by passing client certificate, private key and ca certificate Cert authentication method",
										Attributes: map[string]schema.Attribute{
											"client_cert": schema.SingleNestedAttribute{
												Description:         "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
												MarkdownDescription: "ClientCert is a certificate to authenticate using the Cert Vault authentication method",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
												MarkdownDescription: "SecretRef to a key in a Secret resource containing client private key to authenticate with Vault using the Cert authentication method",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

									"iam": schema.SingleNestedAttribute{
										Description:         "Iam authenticates with vault by passing a special AWS request signed with AWS IAM credentials AWS IAM authentication method",
										MarkdownDescription: "Iam authenticates with vault by passing a special AWS request signed with AWS IAM credentials AWS IAM authentication method",
										Attributes: map[string]schema.Attribute{
											"external_id": schema.StringAttribute{
												Description:         "AWS External ID set on assumed IAM roles",
												MarkdownDescription: "AWS External ID set on assumed IAM roles",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"jwt": schema.SingleNestedAttribute{
												Description:         "Specify a service account with IRSA enabled",
												MarkdownDescription: "Specify a service account with IRSA enabled",
												Attributes: map[string]schema.Attribute{
													"service_account_ref": schema.SingleNestedAttribute{
														Description:         "A reference to a ServiceAccount resource.",
														MarkdownDescription: "A reference to a ServiceAccount resource.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

											"path": schema.StringAttribute{
												Description:         "Path where the AWS auth method is enabled in Vault, e.g: 'aws'",
												MarkdownDescription: "Path where the AWS auth method is enabled in Vault, e.g: 'aws'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"region": schema.StringAttribute{
												Description:         "AWS region",
												MarkdownDescription: "AWS region",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"role": schema.StringAttribute{
												Description:         "This is the AWS role to be assumed before talking to vault",
												MarkdownDescription: "This is the AWS role to be assumed before talking to vault",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "Specify credentials in a Secret object",
												MarkdownDescription: "Specify credentials in a Secret object",
												Attributes: map[string]schema.Attribute{
													"access_key_id_secret_ref": schema.SingleNestedAttribute{
														Description:         "The AccessKeyID is used for authentication",
														MarkdownDescription: "The AccessKeyID is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"namespace": schema.StringAttribute{
																Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_access_key_secret_ref": schema.SingleNestedAttribute{
														Description:         "The SecretAccessKey is used for authentication",
														MarkdownDescription: "The SecretAccessKey is used for authentication",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"namespace": schema.StringAttribute{
																Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"session_token_secret_ref": schema.SingleNestedAttribute{
														Description:         "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
														MarkdownDescription: "The SessionToken used for authentication This must be defined if AccessKeyID and SecretAccessKey are temporary credentials see: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
																MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the Secret resource being referred to.",
																MarkdownDescription: "The name of the Secret resource being referred to.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"namespace": schema.StringAttribute{
																Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

											"vault_aws_iam_server_id": schema.StringAttribute{
												Description:         "X-Vault-AWS-IAM-Server-ID is an additional header used by Vault IAM auth method to mitigate against different types of replay attacks. More details here: https://developer.hashicorp.com/vault/docs/auth/aws",
												MarkdownDescription: "X-Vault-AWS-IAM-Server-ID is an additional header used by Vault IAM auth method to mitigate against different types of replay attacks. More details here: https://developer.hashicorp.com/vault/docs/auth/aws",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"vault_role": schema.StringAttribute{
												Description:         "Vault Role. In vault, a role describes an identity with a set of permissions, groups, or policies you want to attach a user of the secrets engine",
												MarkdownDescription: "Vault Role. In vault, a role describes an identity with a set of permissions, groups, or policies you want to attach a user of the secrets engine",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"jwt": schema.SingleNestedAttribute{
										Description:         "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
										MarkdownDescription: "Jwt authenticates with Vault by passing role and JWT token using the JWT/OIDC authentication method",
										Attributes: map[string]schema.Attribute{
											"kubernetes_service_account_token": schema.SingleNestedAttribute{
												Description:         "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the 'TokenRequest' API.",
												MarkdownDescription: "Optional ServiceAccountToken specifies the Kubernetes service account for which to request a token for with the 'TokenRequest' API.",
												Attributes: map[string]schema.Attribute{
													"audiences": schema.ListAttribute{
														Description:         "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified. Deprecated: use serviceAccountRef.Audiences instead",
														MarkdownDescription: "Optional audiences field that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Defaults to a single audience 'vault' it not specified. Deprecated: use serviceAccountRef.Audiences instead",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"expiration_seconds": schema.Int64Attribute{
														Description:         "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Deprecated: this will be removed in the future. Defaults to 10 minutes.",
														MarkdownDescription: "Optional expiration time in seconds that will be used to request a temporary Kubernetes service account token for the service account referenced by 'serviceAccountRef'. Deprecated: this will be removed in the future. Defaults to 10 minutes.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_account_ref": schema.SingleNestedAttribute{
														Description:         "Service account field containing the name of a kubernetes ServiceAccount.",
														MarkdownDescription: "Service account field containing the name of a kubernetes ServiceAccount.",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the ServiceAccount resource being referred to.",
																MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(253),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
																},
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
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

											"path": schema.StringAttribute{
												Description:         "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",
												MarkdownDescription: "Path where the JWT authentication backend is mounted in Vault, e.g: 'jwt'",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"role": schema.StringAttribute{
												Description:         "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
												MarkdownDescription: "Role is a JWT role to authenticate using the JWT/OIDC Vault authentication method",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
												MarkdownDescription: "Optional SecretRef that refers to a key in a Secret resource containing JWT token to authenticate with Vault using the JWT/OIDC authentication method.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

									"kubernetes": schema.SingleNestedAttribute{
										Description:         "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
										MarkdownDescription: "Kubernetes authenticates with Vault by passing the ServiceAccount token stored in the named Secret resource to the Vault server.",
										Attributes: map[string]schema.Attribute{
											"mount_path": schema.StringAttribute{
												Description:         "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",
												MarkdownDescription: "Path where the Kubernetes authentication backend is mounted in Vault, e.g: 'kubernetes'",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"role": schema.StringAttribute{
												Description:         "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
												MarkdownDescription: "A required field containing the Vault Role to assume. A Role binds a Kubernetes ServiceAccount with a set of Vault policies.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
												MarkdownDescription: "Optional secret field containing a Kubernetes ServiceAccount JWT used for authenticating with Vault. If a name is specified without a key, 'token' is the default. If one is not specified, the one bound to the controller will be used.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_account_ref": schema.SingleNestedAttribute{
												Description:         "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
												MarkdownDescription: "Optional service account field containing the name of a kubernetes ServiceAccount. If the service account is specified, the service account secret token JWT will be used for authenticating with Vault. If the service account selector is not supplied, the secretRef will be used instead.",
												Attributes: map[string]schema.Attribute{
													"audiences": schema.ListAttribute{
														Description:         "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
														MarkdownDescription: "Audience specifies the 'aud' claim for the service account token If the service account uses a well-known annotation for e.g. IRSA or GCP Workload Identity then this audiences will be appended to the list",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the ServiceAccount resource being referred to.",
														MarkdownDescription: "The name of the ServiceAccount resource being referred to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "Namespace of the resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

									"ldap": schema.SingleNestedAttribute{
										Description:         "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
										MarkdownDescription: "Ldap authenticates with Vault by passing username/password pair using the LDAP authentication method",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",
												MarkdownDescription: "Path where the LDAP authentication backend is mounted in Vault, e.g: 'ldap'",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
												MarkdownDescription: "SecretRef to a key in a Secret resource containing password for the LDAP user used to authenticate with Vault using the LDAP authentication method",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": schema.StringAttribute{
												Description:         "Username is an LDAP username used to authenticate using the LDAP Vault authentication method",
												MarkdownDescription: "Username is an LDAP username used to authenticate using the LDAP Vault authentication method",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Name of the vault namespace to authenticate to. This can be different than the namespace your secret is in. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces This will default to Vault.Namespace field if set, or empty otherwise",
										MarkdownDescription: "Name of the vault namespace to authenticate to. This can be different than the namespace your secret is in. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces This will default to Vault.Namespace field if set, or empty otherwise",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"token_secret_ref": schema.SingleNestedAttribute{
										Description:         "TokenSecretRef authenticates with Vault by presenting a token.",
										MarkdownDescription: "TokenSecretRef authenticates with Vault by presenting a token.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "The name of the Secret resource being referred to.",
												MarkdownDescription: "The name of the Secret resource being referred to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_pass": schema.SingleNestedAttribute{
										Description:         "UserPass authenticates with Vault by passing username/password pair",
										MarkdownDescription: "UserPass authenticates with Vault by passing username/password pair",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path where the UserPassword authentication backend is mounted in Vault, e.g: 'userpass'",
												MarkdownDescription: "Path where the UserPassword authentication backend is mounted in Vault, e.g: 'userpass'",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretRef to a key in a Secret resource containing password for the user used to authenticate with Vault using the UserPass authentication method",
												MarkdownDescription: "SecretRef to a key in a Secret resource containing password for the user used to authenticate with Vault using the UserPass authentication method",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the Secret resource being referred to.",
														MarkdownDescription: "The name of the Secret resource being referred to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(253),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
														},
													},

													"namespace": schema.StringAttribute{
														Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": schema.StringAttribute{
												Description:         "Username is a username used to authenticate using the UserPass Vault authentication method",
												MarkdownDescription: "Username is a username used to authenticate using the UserPass Vault authentication method",
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

							"ca_bundle": schema.StringAttribute{
								Description:         "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
								MarkdownDescription: "PEM encoded CA bundle used to validate Vault server certificate. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. If not set the system root certificates are used to validate the TLS connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},

							"ca_provider": schema.SingleNestedAttribute{
								Description:         "The provider for the CA bundle to use to validate Vault server certificate.",
								MarkdownDescription: "The provider for the CA bundle to use to validate Vault server certificate.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key where the CA certificate can be found in the Secret or ConfigMap.",
										MarkdownDescription: "The key where the CA certificate can be found in the Secret or ConfigMap.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
										},
									},

									"name": schema.StringAttribute{
										Description:         "The name of the object located at the provider type.",
										MarkdownDescription: "The name of the object located at the provider type.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
										},
									},

									"namespace": schema.StringAttribute{
										Description:         "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
										MarkdownDescription: "The namespace the Provider type is in. Can only be defined when used in a ClusterSecretStore.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(63),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
										},
									},

									"type": schema.StringAttribute{
										Description:         "The type of provider to use such as 'Secret', or 'ConfigMap'.",
										MarkdownDescription: "The type of provider to use such as 'Secret', or 'ConfigMap'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Secret", "ConfigMap"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"check_and_set": schema.SingleNestedAttribute{
								Description:         "CheckAndSet defines the Check-And-Set (CAS) settings for PushSecret operations. Only applies to Vault KV v2 stores. When enabled, write operations must include the current version of the secret to prevent unintentional overwrites.",
								MarkdownDescription: "CheckAndSet defines the Check-And-Set (CAS) settings for PushSecret operations. Only applies to Vault KV v2 stores. When enabled, write operations must include the current version of the secret to prevent unintentional overwrites.",
								Attributes: map[string]schema.Attribute{
									"required": schema.BoolAttribute{
										Description:         "Required when true, all write operations must include a check-and-set parameter. This helps prevent unintentional overwrites of secrets.",
										MarkdownDescription: "Required when true, all write operations must include a check-and-set parameter. This helps prevent unintentional overwrites of secrets.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"forward_inconsistent": schema.BoolAttribute{
								Description:         "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
								MarkdownDescription: "ForwardInconsistent tells Vault to forward read-after-write requests to the Vault leader instead of simply retrying within a loop. This can increase performance if the option is enabled serverside. https://www.vaultproject.io/docs/configuration/replication#allow_forwarding_via_header",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"headers": schema.MapAttribute{
								Description:         "Headers to be added in Vault request",
								MarkdownDescription: "Headers to be added in Vault request",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
								MarkdownDescription: "Name of the vault namespace. Namespaces is a set of features within Vault Enterprise that allows Vault environments to support Secure Multi-tenancy. e.g: 'ns1'. More about namespaces can be found here https://www.vaultproject.io/docs/enterprise/namespaces",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"path": schema.StringAttribute{
								Description:         "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
								MarkdownDescription: "Path is the mount path of the Vault KV backend endpoint, e.g: 'secret'. The v2 KV secret engine version specific '/data' path suffix for fetching secrets from Vault is optional and will be appended if not present in specified path.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read_your_writes": schema.BoolAttribute{
								Description:         "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
								MarkdownDescription: "ReadYourWrites ensures isolated read-after-write semantics by providing discovered cluster replication states in each request. More information about eventual consistency in Vault can be found here https://www.vaultproject.io/docs/enterprise/consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server": schema.StringAttribute{
								Description:         "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
								MarkdownDescription: "Server is the connection address for the Vault server, e.g: 'https://vault.example.com:8200'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "The configuration used for client side related TLS communication, when the Vault server requires mutual authentication. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. It's worth noting this configuration is different from the 'TLS certificates auth method', which is available under the 'auth.cert' section.",
								MarkdownDescription: "The configuration used for client side related TLS communication, when the Vault server requires mutual authentication. Only used if the Server URL is using HTTPS protocol. This parameter is ignored for plain HTTP protocol connection. It's worth noting this configuration is different from the 'TLS certificates auth method', which is available under the 'auth.cert' section.",
								Attributes: map[string]schema.Attribute{
									"cert_secret_ref": schema.SingleNestedAttribute{
										Description:         "CertSecretRef is a certificate added to the transport layer when communicating with the Vault server. If no key for the Secret is specified, external-secret will default to 'tls.crt'.",
										MarkdownDescription: "CertSecretRef is a certificate added to the transport layer when communicating with the Vault server. If no key for the Secret is specified, external-secret will default to 'tls.crt'.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "The name of the Secret resource being referred to.",
												MarkdownDescription: "The name of the Secret resource being referred to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_secret_ref": schema.SingleNestedAttribute{
										Description:         "KeySecretRef to a key in a Secret resource containing client private key added to the transport layer when communicating with the Vault server. If no key for the Secret is specified, external-secret will default to 'tls.key'.",
										MarkdownDescription: "KeySecretRef to a key in a Secret resource containing client private key added to the transport layer when communicating with the Vault server. If no key for the Secret is specified, external-secret will default to 'tls.key'.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												MarkdownDescription: "A key in the referenced Secret. Some instances of this field may be defaulted, in others it may be required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[-._a-zA-Z0-9]+$`), ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "The name of the Secret resource being referred to.",
												MarkdownDescription: "The name of the Secret resource being referred to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												MarkdownDescription: "The namespace of the Secret resource being referred to. Ignored if referent is not cluster-scoped, otherwise defaults to the namespace of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
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

							"version": schema.StringAttribute{
								Description:         "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",
								MarkdownDescription: "Version is the Vault KV secret engine version. This can be either 'v1' or 'v2'. Version defaults to 'v2'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("v1", "v2"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"result_type": schema.StringAttribute{
						Description:         "Result type defines which data is returned from the generator. By default it is the 'data' section of the Vault API response. When using e.g. /auth/token/create the 'data' section is empty but the 'auth' section contains the generated token. Please refer to the vault docs regarding the result data structure. Additionally, accessing the raw response is possibly by using 'Raw' result type.",
						MarkdownDescription: "Result type defines which data is returned from the generator. By default it is the 'data' section of the Vault API response. When using e.g. /auth/token/create the 'data' section is empty but the 'auth' section contains the generated token. Please refer to the vault docs regarding the result data structure. Additionally, accessing the raw response is possibly by using 'Raw' result type.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Data", "Auth", "Raw"),
						},
					},

					"retry_settings": schema.SingleNestedAttribute{
						Description:         "Used to configure http retries if failed",
						MarkdownDescription: "Used to configure http retries if failed",
						Attributes: map[string]schema.Attribute{
							"max_retries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retry_interval": schema.StringAttribute{
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
	}
}

func (r *GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_generators_external_secrets_io_vault_dynamic_secret_v1alpha1_manifest")

	var model GeneratorsExternalSecretsIoVaultDynamicSecretV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("generators.external-secrets.io/v1alpha1")
	model.Kind = pointer.String("VaultDynamicSecret")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
