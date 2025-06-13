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
	_ datasource.DataSource = &RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoPkisecretEngineConfigV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1Manifest{}
}

type RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1Manifest struct{}

type RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1ManifestData struct {
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
		CRLDisable            *bool     `tfsdk:"crl_disable" json:"CRLDisable,omitempty"`
		CRLDistributionPoints *[]string `tfsdk:"crl_distribution_points" json:"CRLDistributionPoints,omitempty"`
		CRLExpiry             *string   `tfsdk:"crl_expiry" json:"CRLExpiry,omitempty"`
		IPSans                *string   `tfsdk:"ip_sans" json:"IPSans,omitempty"`
		TTL                   *string   `tfsdk:"ttl" json:"TTL,omitempty"`
		URISans               *string   `tfsdk:"uri_sans" json:"URISans,omitempty"`
		AltNames              *string   `tfsdk:"alt_names" json:"altNames,omitempty"`
		Authentication        *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		CertificateKey *string `tfsdk:"certificate_key" json:"certificateKey,omitempty"`
		CommonName     *string `tfsdk:"common_name" json:"commonName,omitempty"`
		Connection     *struct {
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
		Country            *string `tfsdk:"country" json:"country,omitempty"`
		ExcludeCnFromSans  *bool   `tfsdk:"exclude_cn_from_sans" json:"excludeCnFromSans,omitempty"`
		ExternalSignSecret *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"external_sign_secret" json:"externalSignSecret,omitempty"`
		Format       *string `tfsdk:"format" json:"format,omitempty"`
		InternalSign *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"internal_sign" json:"internalSign,omitempty"`
		IssuingCertificates *[]string `tfsdk:"issuing_certificates" json:"issuingCertificates,omitempty"`
		KeyBits             *int64    `tfsdk:"key_bits" json:"keyBits,omitempty"`
		KeyType             *string   `tfsdk:"key_type" json:"keyType,omitempty"`
		Locality            *string   `tfsdk:"locality" json:"locality,omitempty"`
		MaxPathLength       *int64    `tfsdk:"max_path_length" json:"maxPathLength,omitempty"`
		OcspServers         *[]string `tfsdk:"ocsp_servers" json:"ocspServers,omitempty"`
		Organization        *string   `tfsdk:"organization" json:"organization,omitempty"`
		OtherSans           *string   `tfsdk:"other_sans" json:"otherSans,omitempty"`
		Ou                  *string   `tfsdk:"ou" json:"ou,omitempty"`
		Path                *string   `tfsdk:"path" json:"path,omitempty"`
		PermittedDnsDomains *[]string `tfsdk:"permitted_dns_domains" json:"permittedDnsDomains,omitempty"`
		PostalCode          *string   `tfsdk:"postal_code" json:"postalCode,omitempty"`
		PrivateKeyFormat    *string   `tfsdk:"private_key_format" json:"privateKeyFormat,omitempty"`
		PrivateKeyType      *string   `tfsdk:"private_key_type" json:"privateKeyType,omitempty"`
		Province            *string   `tfsdk:"province" json:"province,omitempty"`
		SerialNumber        *string   `tfsdk:"serial_number" json:"serialNumber,omitempty"`
		StreetAddress       *string   `tfsdk:"street_address" json:"streetAddress,omitempty"`
		Type                *string   `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_pki_secret_engine_config_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PKISecretEngineConfig is the Schema for the pkisecretengineconfigs API",
		MarkdownDescription: "PKISecretEngineConfig is the Schema for the pkisecretengineconfigs API",
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
				Description:         "PKISecretEngineConfigSpec defines the desired state of PKISecretEngineConfig",
				MarkdownDescription: "PKISecretEngineConfigSpec defines the desired state of PKISecretEngineConfig",
				Attributes: map[string]schema.Attribute{
					"crl_disable": schema.BoolAttribute{
						Description:         "Disables or enables CRL building.",
						MarkdownDescription: "Disables or enables CRL building.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"crl_distribution_points": schema.ListAttribute{
						Description:         "Specifies the URL values for the CRL Distribution Points field. This can be an array or a comma-separated string list. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "Specifies the URL values for the CRL Distribution Points field. This can be an array or a comma-separated string list. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"crl_expiry": schema.StringAttribute{
						Description:         "Specifies the time until expiration.",
						MarkdownDescription: "Specifies the time until expiration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ip_sans": schema.StringAttribute{
						Description:         "Specifies the requested IP Subject Alternative Names, in a comma-delimited list.",
						MarkdownDescription: "Specifies the requested IP Subject Alternative Names, in a comma-delimited list.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ttl": schema.StringAttribute{
						Description:         "Specifies the requested Time To Live (after which the certificate will be expired). This cannot be larger than the engine's max (or, if not set, the system max).",
						MarkdownDescription: "Specifies the requested Time To Live (after which the certificate will be expired). This cannot be larger than the engine's max (or, if not set, the system max).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uri_sans": schema.StringAttribute{
						Description:         "Specifies the requested URI Subject Alternative Names, in a comma-delimited list.",
						MarkdownDescription: "Specifies the requested URI Subject Alternative Names, in a comma-delimited list.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"alt_names": schema.StringAttribute{
						Description:         "Specifies the requested Subject Alternative Names, in a comma-delimited list. These can be host names or email addresses; they will be parsed into their respective fields.",
						MarkdownDescription: "Specifies the requested Subject Alternative Names, in a comma-delimited list. These can be host names or email addresses; they will be parsed into their respective fields.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"authentication": schema.SingleNestedAttribute{
						Description:         "Authentication is the kube auth configuration to be used to execute this request",
						MarkdownDescription: "Authentication is the kube auth configuration to be used to execute this request",
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

					"certificate_key": schema.StringAttribute{
						Description:         "CertificateKey key to be used when retrieving the signed certificate",
						MarkdownDescription: "CertificateKey key to be used when retrieving the signed certificate",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"common_name": schema.StringAttribute{
						Description:         "Specifies the requested CN for the certificate.",
						MarkdownDescription: "Specifies the requested CN for the certificate.",
						Required:            false,
						Optional:            true,
						Computed:            false,
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

					"country": schema.StringAttribute{
						Description:         "Specifies the C (Country) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the C (Country) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"exclude_cn_from_sans": schema.BoolAttribute{
						Description:         "If set, the given common_name will not be included in DNS or Email Subject Alternate Names (as appropriate). Useful if the CN is not a hostname or email address, but is instead some human-readable identifier.",
						MarkdownDescription: "If set, the given common_name will not be included in DNS or Email Subject Alternate Names (as appropriate). Useful if the CN is not a hostname or email address, but is instead some human-readable identifier.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"external_sign_secret": schema.SingleNestedAttribute{
						Description:         "ExternalSignSecret retrieves the signed intermediate certificate from a Kubernetes secret. Allows submitting the signed CA certificate corresponding to a private key generated.",
						MarkdownDescription: "ExternalSignSecret retrieves the signed intermediate certificate from a Kubernetes secret. Allows submitting the signed CA certificate corresponding to a private key generated.",
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

					"format": schema.StringAttribute{
						Description:         "Specifies the format for returned data. Can be pem, der, or pem_bundle. If der, the output is base64 encoded. If pem_bundle, the certificate field will contain the private key (if exported) and certificate, concatenated; if the issuing CA is not a Vault-derived self-signed root, this will be included as well.",
						MarkdownDescription: "Specifies the format for returned data. Can be pem, der, or pem_bundle. If der, the output is base64 encoded. If pem_bundle, the certificate field will contain the private key (if exported) and certificate, concatenated; if the issuing CA is not a Vault-derived self-signed root, this will be included as well.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("pem", "pem_bundle", "der"),
						},
					},

					"internal_sign": schema.SingleNestedAttribute{
						Description:         "Use the configured refered Vault PKISecretEngineConfig to issue a certificate with appropriate values for acting as an intermediate CA.",
						MarkdownDescription: "Use the configured refered Vault PKISecretEngineConfig to issue a certificate with appropriate values for acting as an intermediate CA.",
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

					"issuing_certificates": schema.ListAttribute{
						Description:         "Specifies the URL values for the Issuing Certificate field. This can be an array or a comma-separated string list. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "Specifies the URL values for the Issuing Certificate field. This can be an array or a comma-separated string list. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_bits": schema.Int64Attribute{
						Description:         "Specifies the number of bits to use. This must be changed to a valid value if the key_type is ec, e.g., 224, 256, 384 or 521.",
						MarkdownDescription: "Specifies the number of bits to use. This must be changed to a valid value if the key_type is ec, e.g., 224, 256, 384 or 521.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_type": schema.StringAttribute{
						Description:         "Specifies the desired key type; must be rsa or ec.",
						MarkdownDescription: "Specifies the desired key type; must be rsa or ec.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("rsa", "ec"),
						},
					},

					"locality": schema.StringAttribute{
						Description:         "Specifies the L (Locality) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the L (Locality) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_path_length": schema.Int64Attribute{
						Description:         "Specifies the maximum path length to encode in the generated certificate. -1 means no limit. Unless the signing certificate has a maximum path length set, in which case the path length is set to one less than that of the signing certificate. A limit of 0 means a literal path length of zero.",
						MarkdownDescription: "Specifies the maximum path length to encode in the generated certificate. -1 means no limit. Unless the signing certificate has a maximum path length set, in which case the path length is set to one less than that of the signing certificate. A limit of 0 means a literal path length of zero.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_servers": schema.ListAttribute{
						Description:         "Specifies the URL values for the OCSP Servers field. This can be an array or a comma-separated string list. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "Specifies the URL values for the OCSP Servers field. This can be an array or a comma-separated string list. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"organization": schema.StringAttribute{
						Description:         "Specifies the O (Organization) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the O (Organization) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"other_sans": schema.StringAttribute{
						Description:         "Specifies custom OID/UTF8-string SANs. These must match values specified on the role in allowed_other_sans (see role creation for allowed_other_sans globbing rules). The format is the same as OpenSSL: <oid>;<type>:<value> where the only current valid type is UTF8. This can be a comma-delimited list or a JSON string slice.",
						MarkdownDescription: "Specifies custom OID/UTF8-string SANs. These must match values specified on the role in allowed_other_sans (see role creation for allowed_other_sans globbing rules). The format is the same as OpenSSL: <oid>;<type>:<value> where the only current valid type is UTF8. This can be a comma-delimited list or a JSON string slice.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ou": schema.StringAttribute{
						Description:         "Specifies the OU (OrganizationalUnit) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the OU (OrganizationalUnit) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path at which to create the role. The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/config/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						MarkdownDescription: "Path at which to create the role. The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/config/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
						},
					},

					"permitted_dns_domains": schema.ListAttribute{
						Description:         "A comma separated string (or, string array) containing DNS domains for which certificates are allowed to be issued or signed by this CA certificate. Note that subdomains are allowed, as per RFC. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "A comma separated string (or, string array) containing DNS domains for which certificates are allowed to be issued or signed by this CA certificate. Note that subdomains are allowed, as per RFC. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postal_code": schema.StringAttribute{
						Description:         "Specifies the Postal Code values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the Postal Code values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_key_format": schema.StringAttribute{
						Description:         "Specifies the format for marshaling the private key. Defaults to der which will return either base64-encoded DER or PEM-encoded DER, depending on the value of format. The other option is pkcs8 which will return the key marshalled as PEM-encoded PKCS8.",
						MarkdownDescription: "Specifies the format for marshaling the private key. Defaults to der which will return either base64-encoded DER or PEM-encoded DER, depending on the value of format. The other option is pkcs8 which will return the key marshalled as PEM-encoded PKCS8.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_key_type": schema.StringAttribute{
						Description:         "Specifies the type of the root to create. If exported, the private key will be returned in the response; if internal the private key will not be returned and cannot be retrieved later. This is part of the request URL.",
						MarkdownDescription: "Specifies the type of the root to create. If exported, the private key will be returned in the response; if internal the private key will not be returned and cannot be retrieved later. This is part of the request URL.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("internal", "exported"),
						},
					},

					"province": schema.StringAttribute{
						Description:         "Specifies the ST (Province) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the ST (Province) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"serial_number": schema.StringAttribute{
						Description:         "Specifies the Serial Number, if any. Otherwise Vault will generate a random serial for you. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5.",
						MarkdownDescription: "Specifies the Serial Number, if any. Otherwise Vault will generate a random serial for you. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"street_address": schema.StringAttribute{
						Description:         "Specifies the Street Address values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the Street Address values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Specifies the type of certificate authority. Root CA or Intermediate CA. This is part of the request URL.",
						MarkdownDescription: "Specifies the type of certificate authority. Root CA or Intermediate CA. This is part of the request URL.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("root", "intermediate"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_pki_secret_engine_config_v1alpha1_manifest")

	var model RedhatcopRedhatIoPkisecretEngineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("PKISecretEngineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
