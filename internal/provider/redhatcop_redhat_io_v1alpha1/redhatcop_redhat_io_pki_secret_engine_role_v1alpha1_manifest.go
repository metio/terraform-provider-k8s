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
	_ datasource.DataSource = &RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoPkisecretEngineRoleV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1Manifest{}
}

type RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1Manifest struct{}

type RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1ManifestData struct {
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
		TTL                    *string   `tfsdk:"ttl" json:"TTL,omitempty"`
		AllowAnyName           *bool     `tfsdk:"allow_any_name" json:"allowAnyName,omitempty"`
		AllowBareDomains       *bool     `tfsdk:"allow_bare_domains" json:"allowBareDomains,omitempty"`
		AllowGlobDomains       *bool     `tfsdk:"allow_glob_domains" json:"allowGlobDomains,omitempty"`
		AllowIPSans            *bool     `tfsdk:"allow_ip_sans" json:"allowIPSans,omitempty"`
		AllowLocalhost         *bool     `tfsdk:"allow_localhost" json:"allowLocalhost,omitempty"`
		AllowSubdomains        *bool     `tfsdk:"allow_subdomains" json:"allowSubdomains,omitempty"`
		AllowedDomains         *[]string `tfsdk:"allowed_domains" json:"allowedDomains,omitempty"`
		AllowedDomainsTemplate *bool     `tfsdk:"allowed_domains_template" json:"allowedDomainsTemplate,omitempty"`
		AllowedOtherSans       *string   `tfsdk:"allowed_other_sans" json:"allowedOtherSans,omitempty"`
		AllowedURISans         *[]string `tfsdk:"allowed_uri_sans" json:"allowedURISans,omitempty"`
		Authentication         *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		BasicConstraintsValidForNonCa *bool `tfsdk:"basic_constraints_valid_for_non_ca" json:"basicConstraintsValidForNonCa,omitempty"`
		ClientFlag                    *bool `tfsdk:"client_flag" json:"clientFlag,omitempty"`
		CodeSigningFlag               *bool `tfsdk:"code_signing_flag" json:"codeSigningFlag,omitempty"`
		Connection                    *struct {
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
		Country             *string   `tfsdk:"country" json:"country,omitempty"`
		EmailProtectionFlag *bool     `tfsdk:"email_protection_flag" json:"emailProtectionFlag,omitempty"`
		EnforceHostnames    *bool     `tfsdk:"enforce_hostnames" json:"enforceHostnames,omitempty"`
		ExtKeyUsage         *[]string `tfsdk:"ext_key_usage" json:"extKeyUsage,omitempty"`
		ExtKeyUsageOids     *[]string `tfsdk:"ext_key_usage_oids" json:"extKeyUsageOids,omitempty"`
		GenerateLease       *bool     `tfsdk:"generate_lease" json:"generateLease,omitempty"`
		KeyBits             *int64    `tfsdk:"key_bits" json:"keyBits,omitempty"`
		KeyType             *string   `tfsdk:"key_type" json:"keyType,omitempty"`
		KeyUsage            *[]string `tfsdk:"key_usage" json:"keyUsage,omitempty"`
		Locality            *string   `tfsdk:"locality" json:"locality,omitempty"`
		MaxTTL              *string   `tfsdk:"max_ttl" json:"maxTTL,omitempty"`
		Name                *string   `tfsdk:"name" json:"name,omitempty"`
		NoStore             *bool     `tfsdk:"no_store" json:"noStore,omitempty"`
		NotBeforeDuration   *string   `tfsdk:"not_before_duration" json:"notBeforeDuration,omitempty"`
		Organization        *string   `tfsdk:"organization" json:"organization,omitempty"`
		Ou                  *string   `tfsdk:"ou" json:"ou,omitempty"`
		Path                *string   `tfsdk:"path" json:"path,omitempty"`
		PolicyIdentifiers   *[]string `tfsdk:"policy_identifiers" json:"policyIdentifiers,omitempty"`
		PostalCode          *string   `tfsdk:"postal_code" json:"postalCode,omitempty"`
		Province            *string   `tfsdk:"province" json:"province,omitempty"`
		RequireCn           *bool     `tfsdk:"require_cn" json:"requireCn,omitempty"`
		SerialNumber        *string   `tfsdk:"serial_number" json:"serialNumber,omitempty"`
		ServerFlag          *bool     `tfsdk:"server_flag" json:"serverFlag,omitempty"`
		StreetAddress       *string   `tfsdk:"street_address" json:"streetAddress,omitempty"`
		UseCSRCommonName    *bool     `tfsdk:"use_csr_common_name" json:"useCSRCommonName,omitempty"`
		UseCSRSans          *bool     `tfsdk:"use_csr_sans" json:"useCSRSans,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_pki_secret_engine_role_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PKISecretEngineRole is the Schema for the pkisecretengineroles API",
		MarkdownDescription: "PKISecretEngineRole is the Schema for the pkisecretengineroles API",
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
				Description:         "PKISecretEngineRoleSpec defines the desired state of PKISecretEngineRole",
				MarkdownDescription: "PKISecretEngineRoleSpec defines the desired state of PKISecretEngineRole",
				Attributes: map[string]schema.Attribute{
					"ttl": schema.StringAttribute{
						Description:         "Specifies the Time To Live value provided as a string duration with time suffix. Hour is the largest suffix. If not set, uses the system default value or the value of max_ttl, whichever is shorter.",
						MarkdownDescription: "Specifies the Time To Live value provided as a string duration with time suffix. Hour is the largest suffix. If not set, uses the system default value or the value of max_ttl, whichever is shorter.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_any_name": schema.BoolAttribute{
						Description:         "Specifies if clients can request any CN. Useful in some circumstances, but make sure you understand whether it is appropriate for your installation before enabling it.",
						MarkdownDescription: "Specifies if clients can request any CN. Useful in some circumstances, but make sure you understand whether it is appropriate for your installation before enabling it.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_bare_domains": schema.BoolAttribute{
						Description:         "Specifies if clients can request certificates matching the value of the actual domains themselves; e.g. if a configured domain set with allowed_domains is example.com, this allows clients to actually request a certificate containing the name example.com as one of the DNS values on the final certificate. In some scenarios, this can be considered a security risk.",
						MarkdownDescription: "Specifies if clients can request certificates matching the value of the actual domains themselves; e.g. if a configured domain set with allowed_domains is example.com, this allows clients to actually request a certificate containing the name example.com as one of the DNS values on the final certificate. In some scenarios, this can be considered a security risk.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_glob_domains": schema.BoolAttribute{
						Description:         "Allows names specified in allowed_domains to contain glob patterns (e.g. ftp*.example.com). Clients will be allowed to request certificates with names matching the glob patterns.",
						MarkdownDescription: "Allows names specified in allowed_domains to contain glob patterns (e.g. ftp*.example.com). Clients will be allowed to request certificates with names matching the glob patterns.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_ip_sans": schema.BoolAttribute{
						Description:         "Specifies if clients can request IP Subject Alternative Names. No authorization checking is performed except to verify that the given values are valid IP addresses.",
						MarkdownDescription: "Specifies if clients can request IP Subject Alternative Names. No authorization checking is performed except to verify that the given values are valid IP addresses.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_localhost": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_subdomains": schema.BoolAttribute{
						Description:         "Specifies if clients can request certificates with CNs that are subdomains of the CNs allowed by the other role options. This includes wildcard subdomains. For example, an allowed_domains value of example.com with this option set to true will allow foo.example.com and bar.example.com as well as *.example.com. This is redundant when using the allow_any_name option.",
						MarkdownDescription: "Specifies if clients can request certificates with CNs that are subdomains of the CNs allowed by the other role options. This includes wildcard subdomains. For example, an allowed_domains value of example.com with this option set to true will allow foo.example.com and bar.example.com as well as *.example.com. This is redundant when using the allow_any_name option.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_domains": schema.ListAttribute{
						Description:         "Specifies the domains of the role. This is used with the allow_bare_domains and allow_subdomains options. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "Specifies the domains of the role. This is used with the allow_bare_domains and allow_subdomains options. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_domains_template": schema.BoolAttribute{
						Description:         "When set, allowed_domains may contain templates, as with ACL Path Templating.",
						MarkdownDescription: "When set, allowed_domains may contain templates, as with ACL Path Templating.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_other_sans": schema.StringAttribute{
						Description:         "Defines allowed custom OID/UTF8-string SANs. This can be a comma-delimited list or a JSON string slice, where each element has the same format as OpenSSL: <oid>;<type>:<value>, but the only valid type is UTF8 or UTF-8. The value part of an element may be a * to allow any value with that OID. Alternatively, specifying a single * will allow any other_sans input.",
						MarkdownDescription: "Defines allowed custom OID/UTF8-string SANs. This can be a comma-delimited list or a JSON string slice, where each element has the same format as OpenSSL: <oid>;<type>:<value>, but the only valid type is UTF8 or UTF-8. The value part of an element may be a * to allow any value with that OID. Alternatively, specifying a single * will allow any other_sans input.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_uri_sans": schema.ListAttribute{
						Description:         "Defines allowed URI Subject Alternative Names. No authorization checking is performed except to verify that the given values are valid URIs. This can be a comma-delimited list or a JSON string slice. Values can contain glob patterns (e.g. spiffe://hostname/*). kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "Defines allowed URI Subject Alternative Names. No authorization checking is performed except to verify that the given values are valid URIs. This can be a comma-delimited list or a JSON string slice. Values can contain glob patterns (e.g. spiffe://hostname/*). kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
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

					"basic_constraints_valid_for_non_ca": schema.BoolAttribute{
						Description:         "Mark Basic Constraints valid when issuing non-CA certificates.",
						MarkdownDescription: "Mark Basic Constraints valid when issuing non-CA certificates.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client_flag": schema.BoolAttribute{
						Description:         "Specifies if certificates are flagged for client use.",
						MarkdownDescription: "Specifies if certificates are flagged for client use.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"code_signing_flag": schema.BoolAttribute{
						Description:         "Specifies if certificates are flagged for code signing use.",
						MarkdownDescription: "Specifies if certificates are flagged for code signing use.",
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

					"email_protection_flag": schema.BoolAttribute{
						Description:         "Specifies if certificates are flagged for email protection use.",
						MarkdownDescription: "Specifies if certificates are flagged for email protection use.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enforce_hostnames": schema.BoolAttribute{
						Description:         "Specifies if only valid host names are allowed for CNs, DNS SANs, and the host part of email addresses.",
						MarkdownDescription: "Specifies if only valid host names are allowed for CNs, DNS SANs, and the host part of email addresses.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ext_key_usage": schema.ListAttribute{
						Description:         "Specifies the allowed extended key usage constraint on issued certificates. Valid values can be found at https://golang.org/pkg/crypto/x509/#ExtKeyUsage - simply drop the ExtKeyUsage part of the value. Values are not case-sensitive. To specify no key usage constraints, set this to an empty list. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "Specifies the allowed extended key usage constraint on issued certificates. Valid values can be found at https://golang.org/pkg/crypto/x509/#ExtKeyUsage - simply drop the ExtKeyUsage part of the value. Values are not case-sensitive. To specify no key usage constraints, set this to an empty list. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ext_key_usage_oids": schema.ListAttribute{
						Description:         "A comma-separated string or list of extended key usage oids. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "A comma-separated string or list of extended key usage oids. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"generate_lease": schema.BoolAttribute{
						Description:         "Specifies if certificates issued/signed against this role will have Vault leases attached to them. Certificates can be added to the CRL by vault revoke <lease_id> when certificates are associated with leases. It can also be done using the pki/revoke endpoint. However, when lease generation is disabled, invoking pki/revoke would be the only way to add the certificates to the CRL.",
						MarkdownDescription: "Specifies if certificates issued/signed against this role will have Vault leases attached to them. Certificates can be added to the CRL by vault revoke <lease_id> when certificates are associated with leases. It can also be done using the pki/revoke endpoint. However, when lease generation is disabled, invoking pki/revoke would be the only way to add the certificates to the CRL.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_bits": schema.Int64Attribute{
						Description:         "Specifies the number of bits to use for the generated keys. This will need to be changed for ec keys, e.g., 224, 256, 384 or 521.",
						MarkdownDescription: "Specifies the number of bits to use for the generated keys. This will need to be changed for ec keys, e.g., 224, 256, 384 or 521.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_type": schema.StringAttribute{
						Description:         "Specifies the type of key to generate for generated private keys and the type of key expected for submitted CSRs. Currently, rsa and ec are supported, or when signing CSRs any can be specified to allow keys of either type and with any bit size (subject to > 1024 bits for RSA keys).",
						MarkdownDescription: "Specifies the type of key to generate for generated private keys and the type of key expected for submitted CSRs. Currently, rsa and ec are supported, or when signing CSRs any can be specified to allow keys of either type and with any bit size (subject to > 1024 bits for RSA keys).",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("rsa", "ec"),
						},
					},

					"key_usage": schema.ListAttribute{
						Description:         "Specifies the allowed key usage constraint on issued certificates. Valid values can be found at https://golang.org/pkg/crypto/x509/#KeyUsage - simply drop the KeyUsage part of the value. Values are not case-sensitive. To specify no key usage constraints, set this to an empty list. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "Specifies the allowed key usage constraint on issued certificates. Valid values can be found at https://golang.org/pkg/crypto/x509/#KeyUsage - simply drop the KeyUsage part of the value. Values are not case-sensitive. To specify no key usage constraints, set this to an empty list. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"locality": schema.StringAttribute{
						Description:         "Specifies the L (Locality) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the L (Locality) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_ttl": schema.StringAttribute{
						Description:         "Specifies the maximum Time To Live provided as a string duration with time suffix. Hour is the largest suffix. If not set, defaults to the system maximum lease TTL.",
						MarkdownDescription: "Specifies the maximum Time To Live provided as a string duration with time suffix. Hour is the largest suffix. If not set, defaults to the system maximum lease TTL.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the obejct created in Vault. If this is specified it takes precedence over {metatada.name}",
						MarkdownDescription: "The name of the obejct created in Vault. If this is specified it takes precedence over {metatada.name}",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
						},
					},

					"no_store": schema.BoolAttribute{
						Description:         "If set, certificates issued/signed against this role will not be stored in the storage backend. This can improve performance when issuing large numbers of certificates. However, certificates issued in this way cannot be enumerated or revoked, so this option is recommended only for certificates that are non-sensitive, or extremely short-lived. This option implies a value of false for generate_lease.",
						MarkdownDescription: "If set, certificates issued/signed against this role will not be stored in the storage backend. This can improve performance when issuing large numbers of certificates. However, certificates issued in this way cannot be enumerated or revoked, so this option is recommended only for certificates that are non-sensitive, or extremely short-lived. This option implies a value of false for generate_lease.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"not_before_duration": schema.StringAttribute{
						Description:         "Specifies the duration by which to backdate the NotBefore property.",
						MarkdownDescription: "Specifies the duration by which to backdate the NotBefore property.",
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

					"ou": schema.StringAttribute{
						Description:         "Specifies the OU (OrganizationalUnit) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the OU (OrganizationalUnit) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path at which to create the role. The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/roles/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						MarkdownDescription: "Path at which to create the role. The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/roles/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
						},
					},

					"policy_identifiers": schema.ListAttribute{
						Description:         "A comma-separated string or list of policy OIDs. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "A comma-separated string or list of policy OIDs. kubebuilder:validation:UniqueItems=true",
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

					"province": schema.StringAttribute{
						Description:         "Specifies the ST (Province) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						MarkdownDescription: "Specifies the ST (Province) values in the subject field of issued certificates. This is a comma-separated string or JSON array.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"require_cn": schema.BoolAttribute{
						Description:         "If set to false, makes the common_name field optional while generating a certificate.",
						MarkdownDescription: "If set to false, makes the common_name field optional while generating a certificate.",
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

					"server_flag": schema.BoolAttribute{
						Description:         "Specifies if certificates are flagged for server use.",
						MarkdownDescription: "Specifies if certificates are flagged for server use.",
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

					"use_csr_common_name": schema.BoolAttribute{
						Description:         "When used with the CSR signing endpoint, the common name in the CSR will be used instead of taken from the JSON data. This does not include any requested SANs in the CSR; use use_csr_sans for that.",
						MarkdownDescription: "When used with the CSR signing endpoint, the common name in the CSR will be used instead of taken from the JSON data. This does not include any requested SANs in the CSR; use use_csr_sans for that.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_csr_sans": schema.BoolAttribute{
						Description:         "When used with the CSR signing endpoint, the subject alternate names in the CSR will be used instead of taken from the JSON data. This does not include the common name in the CSR; use use_csr_common_name for that.",
						MarkdownDescription: "When used with the CSR signing endpoint, the subject alternate names in the CSR will be used instead of taken from the JSON data. This does not include the common name in the CSR; use use_csr_common_name for that.",
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
	}
}

func (r *RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_pki_secret_engine_role_v1alpha1_manifest")

	var model RedhatcopRedhatIoPkisecretEngineRoleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("PKISecretEngineRole")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
