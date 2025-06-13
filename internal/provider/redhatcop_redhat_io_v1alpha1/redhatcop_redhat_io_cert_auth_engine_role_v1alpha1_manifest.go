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
	_ datasource.DataSource = &RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoCertAuthEngineRoleV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1Manifest{}
}

type RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1Manifest struct{}

type RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1ManifestData struct {
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
		AllowedCommonNames         *[]string `tfsdk:"allowed_common_names" json:"allowedCommonNames,omitempty"`
		AllowedDNSSANs             *[]string `tfsdk:"allowed_dnssa_ns" json:"allowedDNSSANs,omitempty"`
		AllowedEmailSANs           *[]string `tfsdk:"allowed_email_sa_ns" json:"allowedEmailSANs,omitempty"`
		AllowedMetadataExtensions  *[]string `tfsdk:"allowed_metadata_extensions" json:"allowedMetadataExtensions,omitempty"`
		AllowedOrganizationalUnits *[]string `tfsdk:"allowed_organizational_units" json:"allowedOrganizationalUnits,omitempty"`
		AllowedURISANs             *[]string `tfsdk:"allowed_urisa_ns" json:"allowedURISANs,omitempty"`
		Authentication             *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
		Connection  *struct {
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
		DisplayName          *string   `tfsdk:"display_name" json:"displayName,omitempty"`
		Name                 *string   `tfsdk:"name" json:"name,omitempty"`
		OcspCACertificates   *string   `tfsdk:"ocsp_ca_certificates" json:"ocspCACertificates,omitempty"`
		OcspEnabled          *bool     `tfsdk:"ocsp_enabled" json:"ocspEnabled,omitempty"`
		OcspFailOpen         *bool     `tfsdk:"ocsp_fail_open" json:"ocspFailOpen,omitempty"`
		OcspMaxRetries       *int64    `tfsdk:"ocsp_max_retries" json:"ocspMaxRetries,omitempty"`
		OcspQueryAllServers  *bool     `tfsdk:"ocsp_query_all_servers" json:"ocspQueryAllServers,omitempty"`
		OcspServersOverride  *[]string `tfsdk:"ocsp_servers_override" json:"ocspServersOverride,omitempty"`
		OcspThisUpdateMaxAge *string   `tfsdk:"ocsp_this_update_max_age" json:"ocspThisUpdateMaxAge,omitempty"`
		Path                 *string   `tfsdk:"path" json:"path,omitempty"`
		RequiredExtensions   *[]string `tfsdk:"required_extensions" json:"requiredExtensions,omitempty"`
		TokenBoundCIDRs      *[]string `tfsdk:"token_bound_cidrs" json:"tokenBoundCIDRs,omitempty"`
		TokenExplicitMaxTTL  *string   `tfsdk:"token_explicit_max_ttl" json:"tokenExplicitMaxTTL,omitempty"`
		TokenMaxTTL          *string   `tfsdk:"token_max_ttl" json:"tokenMaxTTL,omitempty"`
		TokenNoDefaultPolicy *bool     `tfsdk:"token_no_default_policy" json:"tokenNoDefaultPolicy,omitempty"`
		TokenNumUses         *int64    `tfsdk:"token_num_uses" json:"tokenNumUses,omitempty"`
		TokenPeriod          *string   `tfsdk:"token_period" json:"tokenPeriod,omitempty"`
		TokenPolicies        *[]string `tfsdk:"token_policies" json:"tokenPolicies,omitempty"`
		TokenTTL             *string   `tfsdk:"token_ttl" json:"tokenTTL,omitempty"`
		TokenType            *string   `tfsdk:"token_type" json:"tokenType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_cert_auth_engine_role_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CertAuthEngineRole is the Schema for the certauthengineroles API",
		MarkdownDescription: "CertAuthEngineRole is the Schema for the certauthengineroles API",
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
				Description:         "CertAuthEngineRoleSpec defines the desired state of CertAuthEngineRole",
				MarkdownDescription: "CertAuthEngineRoleSpec defines the desired state of CertAuthEngineRole",
				Attributes: map[string]schema.Attribute{
					"allowed_common_names": schema.ListAttribute{
						Description:         "Constrain the Common Names in the client certificate with a globbed pattern. Value is a comma-separated list of patterns. Authentication requires at least one Name matching at least one pattern. If not set, defaults to allowing all names.",
						MarkdownDescription: "Constrain the Common Names in the client certificate with a globbed pattern. Value is a comma-separated list of patterns. Authentication requires at least one Name matching at least one pattern. If not set, defaults to allowing all names.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_dnssa_ns": schema.ListAttribute{
						Description:         "Constrain the Alternative Names in the client certificate with a globbed pattern. Value is a comma-separated list of patterns. Authentication requires at least one DNS matching at least one pattern. If not set, defaults to allowing all dns.",
						MarkdownDescription: "Constrain the Alternative Names in the client certificate with a globbed pattern. Value is a comma-separated list of patterns. Authentication requires at least one DNS matching at least one pattern. If not set, defaults to allowing all dns.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_email_sa_ns": schema.ListAttribute{
						Description:         "Constrain the Alternative Names in the client certificate with a globbed pattern. Value is a comma-separated list of patterns. Authentication requires at least one Email matching at least one pattern. If not set, defaults to allowing all emails.",
						MarkdownDescription: "Constrain the Alternative Names in the client certificate with a globbed pattern. Value is a comma-separated list of patterns. Authentication requires at least one Email matching at least one pattern. If not set, defaults to allowing all emails.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_metadata_extensions": schema.ListAttribute{
						Description:         "A comma separated string or array of oid extensions. Upon successful authentication, these extensions will be added as metadata if they are present in the certificate. The metadata key will be the string consisting of the oid numbers separated by a dash (-) instead of a dot (.) to allow usage in ACL templates.",
						MarkdownDescription: "A comma separated string or array of oid extensions. Upon successful authentication, these extensions will be added as metadata if they are present in the certificate. The metadata key will be the string consisting of the oid numbers separated by a dash (-) instead of a dot (.) to allow usage in ACL templates.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_organizational_units": schema.ListAttribute{
						Description:         "Constrain the Organizational Units (OU) in the client certificate with a globbed pattern. Value is a comma-separated list of OU patterns. Authentication requires at least one OU matching at least one pattern. If not set, defaults to allowing all OUs.",
						MarkdownDescription: "Constrain the Organizational Units (OU) in the client certificate with a globbed pattern. Value is a comma-separated list of OU patterns. Authentication requires at least one OU matching at least one pattern. If not set, defaults to allowing all OUs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_urisa_ns": schema.ListAttribute{
						Description:         "Constrain the Alternative Names in the client certificate with a globbed pattern. Value is a comma-separated list of URI patterns. Authentication requires at least one URI matching at least one pattern. If not set, defaults to allowing all URIs.",
						MarkdownDescription: "Constrain the Alternative Names in the client certificate with a globbed pattern. Value is a comma-separated list of URI patterns. Authentication requires at least one URI matching at least one pattern. If not set, defaults to allowing all URIs.",
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

					"certificate": schema.StringAttribute{
						Description:         "The PEM-format CA certificate.",
						MarkdownDescription: "The PEM-format CA certificate.",
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

					"display_name": schema.StringAttribute{
						Description:         "The display_name to set on tokens issued when authenticating against this CA certificate. If not set, defaults to the name of the role.",
						MarkdownDescription: "The display_name to set on tokens issued when authenticating against this CA certificate. If not set, defaults to the name of the role.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the object created in Vault. If this is specified it takes precedence over {metatada.name}",
						MarkdownDescription: "The name of the object created in Vault. If this is specified it takes precedence over {metatada.name}",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[a-z0-9]([-a-z0-9]*[a-z0-9])?`), ""),
						},
					},

					"ocsp_ca_certificates": schema.StringAttribute{
						Description:         "Any additional OCSP responder certificates needed to verify OCSP responses. Provided as base64 encoded PEM data.",
						MarkdownDescription: "Any additional OCSP responder certificates needed to verify OCSP responses. Provided as base64 encoded PEM data.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_enabled": schema.BoolAttribute{
						Description:         "If enabled, validate certificates' revocation status using OCSP.",
						MarkdownDescription: "If enabled, validate certificates' revocation status using OCSP.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_fail_open": schema.BoolAttribute{
						Description:         "If true and an OCSP response cannot be fetched or is of an unknown status, the login will proceed as if the certificate has not been revoked.",
						MarkdownDescription: "If true and an OCSP response cannot be fetched or is of an unknown status, the login will proceed as if the certificate has not been revoked.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_max_retries": schema.Int64Attribute{
						Description:         "The number of retries attempted before giving up on an OCSP request. 0 will disable retries.",
						MarkdownDescription: "The number of retries attempted before giving up on an OCSP request. 0 will disable retries.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_query_all_servers": schema.BoolAttribute{
						Description:         "If set to true, rather than accepting the first successful OCSP response, query all servers and consider the certificate valid only if all servers agree.",
						MarkdownDescription: "If set to true, rather than accepting the first successful OCSP response, query all servers and consider the certificate valid only if all servers agree.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_servers_override": schema.ListAttribute{
						Description:         "A comma-separated list of OCSP server addresses. If unset, the OCSP server is determined from the AuthorityInformationAccess extension on the certificate being inspected.",
						MarkdownDescription: "A comma-separated list of OCSP server addresses. If unset, the OCSP server is determined from the AuthorityInformationAccess extension on the certificate being inspected.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_this_update_max_age": schema.StringAttribute{
						Description:         "If greater than 0, specifies the maximum age of an OCSP thisUpdate field. This avoids accepting old responses without a nextUpdate field.",
						MarkdownDescription: "If greater than 0, specifies the maximum age of an OCSP thisUpdate field. This avoids accepting old responses without a nextUpdate field.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/auth/{spec.path}/certs/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						MarkdownDescription: "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/auth/{spec.path}/certs/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
						},
					},

					"required_extensions": schema.ListAttribute{
						Description:         "Require specific Custom Extension OIDs to exist and match the pattern. Value is a comma separated string or array of oid:value. Expects the extension value to be some type of ASN1 encoded string. All conditions must be met. To match on the hex-encoded value of the extension, including non-string extensions, use the format hex:<oid>:<value>. Supports globbing on value.",
						MarkdownDescription: "Require specific Custom Extension OIDs to exist and match the pattern. Value is a comma separated string or array of oid:value. Expects the extension value to be some type of ASN1 encoded string. All conditions must be met. To match on the hex-encoded value of the extension, including non-string extensions, use the format hex:<oid>:<value>. Supports globbing on value.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_bound_cidrs": schema.ListAttribute{
						Description:         "List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.",
						MarkdownDescription: "List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_explicit_max_ttl": schema.StringAttribute{
						Description:         "If set, will encode an explicit max TTL onto the token. This is a hard cap even if tokenTTL and tokenMaxTTL would otherwise allow a renewal.",
						MarkdownDescription: "If set, will encode an explicit max TTL onto the token. This is a hard cap even if tokenTTL and tokenMaxTTL would otherwise allow a renewal.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_max_ttl": schema.StringAttribute{
						Description:         "The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						MarkdownDescription: "The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_no_default_policy": schema.BoolAttribute{
						Description:         "If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in tokenPolicies.",
						MarkdownDescription: "If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in tokenPolicies.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_num_uses": schema.Int64Attribute{
						Description:         "The maximum number of times a generated token may be used (within its lifetime); 0 means unlimited. If you require the token to have the ability to create child tokens, you will need to set this value to 0.",
						MarkdownDescription: "The maximum number of times a generated token may be used (within its lifetime); 0 means unlimited. If you require the token to have the ability to create child tokens, you will need to set this value to 0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_period": schema.StringAttribute{
						Description:         "The maximum allowed period value when a periodic token is requested from this role.",
						MarkdownDescription: "The maximum allowed period value when a periodic token is requested from this role.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_policies": schema.ListAttribute{
						Description:         "List of token policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.",
						MarkdownDescription: "List of token policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_ttl": schema.StringAttribute{
						Description:         "The incremental lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						MarkdownDescription: "The incremental lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_type": schema.StringAttribute{
						Description:         "The type of token that should be generated. Can be service, batch, or default to use the mount's tuned default (which unless changed will be service tokens). For token store roles, there are two additional possibilities: default-service and default-batch which specify the type to return unless the client requests a different type at generation time. For machine based authentication cases, you should use batch type tokens.",
						MarkdownDescription: "The type of token that should be generated. Can be service, batch, or default to use the mount's tuned default (which unless changed will be service tokens). For token store roles, there are two additional possibilities: default-service and default-batch which specify the type to return unless the client requests a different type at generation time. For machine based authentication cases, you should use batch type tokens.",
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

func (r *RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_cert_auth_engine_role_v1alpha1_manifest")

	var model RedhatcopRedhatIoCertAuthEngineRoleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("CertAuthEngineRole")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
