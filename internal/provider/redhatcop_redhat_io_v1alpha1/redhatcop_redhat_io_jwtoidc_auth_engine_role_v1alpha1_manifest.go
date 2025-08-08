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
	_ datasource.DataSource = &RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1Manifest{}
}

type RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1Manifest struct{}

type RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1ManifestData struct {
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
		OIDCScopes          *[]string `tfsdk:"oidc_scopes" json:"OIDCScopes,omitempty"`
		AllowedRedirectURIs *[]string `tfsdk:"allowed_redirect_ur_is" json:"allowedRedirectURIs,omitempty"`
		Authentication      *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		BoundAudiences  *[]string          `tfsdk:"bound_audiences" json:"boundAudiences,omitempty"`
		BoundClaims     *map[string]string `tfsdk:"bound_claims" json:"boundClaims,omitempty"`
		BoundClaimsType *string            `tfsdk:"bound_claims_type" json:"boundClaimsType,omitempty"`
		BoundSubject    *string            `tfsdk:"bound_subject" json:"boundSubject,omitempty"`
		ClaimMappings   *map[string]string `tfsdk:"claim_mappings" json:"claimMappings,omitempty"`
		ClockSkewLeeway *int64             `tfsdk:"clock_skew_leeway" json:"clockSkewLeeway,omitempty"`
		Connection      *struct {
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
		ExpirationLeeway     *int64    `tfsdk:"expiration_leeway" json:"expirationLeeway,omitempty"`
		GroupsClaim          *string   `tfsdk:"groups_claim" json:"groupsClaim,omitempty"`
		Maxage               *int64    `tfsdk:"maxage" json:"maxage,omitempty"`
		Name                 *string   `tfsdk:"name" json:"name,omitempty"`
		NotBeforeLeeway      *int64    `tfsdk:"not_before_leeway" json:"notBeforeLeeway,omitempty"`
		Path                 *string   `tfsdk:"path" json:"path,omitempty"`
		RoleType             *string   `tfsdk:"role_type" json:"roleType,omitempty"`
		TokenBoundCIDRs      *[]string `tfsdk:"token_bound_cidrs" json:"tokenBoundCIDRs,omitempty"`
		TokenExplicitMaxTTL  *string   `tfsdk:"token_explicit_max_ttl" json:"tokenExplicitMaxTTL,omitempty"`
		TokenMaxTTL          *string   `tfsdk:"token_max_ttl" json:"tokenMaxTTL,omitempty"`
		TokenNoDefaultPolicy *bool     `tfsdk:"token_no_default_policy" json:"tokenNoDefaultPolicy,omitempty"`
		TokenNumUses         *int64    `tfsdk:"token_num_uses" json:"tokenNumUses,omitempty"`
		TokenPeriod          *int64    `tfsdk:"token_period" json:"tokenPeriod,omitempty"`
		TokenPolicies        *[]string `tfsdk:"token_policies" json:"tokenPolicies,omitempty"`
		TokenTTL             *string   `tfsdk:"token_ttl" json:"tokenTTL,omitempty"`
		TokenType            *string   `tfsdk:"token_type" json:"tokenType,omitempty"`
		UserClaim            *string   `tfsdk:"user_claim" json:"userClaim,omitempty"`
		UserClaimJSONPointer *bool     `tfsdk:"user_claim_json_pointer" json:"userClaimJSONPointer,omitempty"`
		VerboseOIDCLogging   *bool     `tfsdk:"verbose_oidc_logging" json:"verboseOIDCLogging,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_jwtoidc_auth_engine_role_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "JWTOIDCAuthEngineRole is the Schema for the jwtoidcauthengineroles API",
		MarkdownDescription: "JWTOIDCAuthEngineRole is the Schema for the jwtoidcauthengineroles API",
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
				Description:         "JWTOIDCAuthEngineRoleSpec defines the desired state of JWTOIDCAuthEngineRole",
				MarkdownDescription: "JWTOIDCAuthEngineRoleSpec defines the desired state of JWTOIDCAuthEngineRole",
				Attributes: map[string]schema.Attribute{
					"oidc_scopes": schema.ListAttribute{
						Description:         "If set, a list of OIDC scopes to be used with an OIDC role The standard scope 'openid' is automatically included and need not be specified kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "If set, a list of OIDC scopes to be used with an OIDC role The standard scope 'openid' is automatically included and need not be specified kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_redirect_ur_is": schema.ListAttribute{
						Description:         "The list of allowed values for redirect_uri during OIDC logins kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "The list of allowed values for redirect_uri during OIDC logins kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

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

					"bound_audiences": schema.ListAttribute{
						Description:         "List of aud claims to match against. Any match is sufficient. Required for 'jwt' roles, optional for 'oidc' roles kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "List of aud claims to match against. Any match is sufficient. Required for 'jwt' roles, optional for 'oidc' roles kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_claims": schema.MapAttribute{
						Description:         "If set, a map of claims (keys) to match against respective claim values (values) The expected value may be a single string or a list of strings The interpretation of the bound claim values is configured with bound_claims_type Keys support JSON pointer syntax for referencing claims",
						MarkdownDescription: "If set, a map of claims (keys) to match against respective claim values (values) The expected value may be a single string or a list of strings The interpretation of the bound claim values is configured with bound_claims_type Keys support JSON pointer syntax for referencing claims",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_claims_type": schema.StringAttribute{
						Description:         "Configures the interpretation of the bound_claims values. If 'string' (the default), the values will treated as string literals and must match exactly. If set to 'glob', the values will be interpreted as globs, with * matching any number of characters",
						MarkdownDescription: "Configures the interpretation of the bound_claims values. If 'string' (the default), the values will treated as string literals and must match exactly. If set to 'glob', the values will be interpreted as globs, with * matching any number of characters",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_subject": schema.StringAttribute{
						Description:         "If set, requires that the sub claim matches this value.",
						MarkdownDescription: "If set, requires that the sub claim matches this value.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"claim_mappings": schema.MapAttribute{
						Description:         "If set, a map of claims (keys) to be copied to specified metadata fields (values) Keys support JSON pointer syntax for referencing claims",
						MarkdownDescription: "If set, a map of claims (keys) to be copied to specified metadata fields (values) Keys support JSON pointer syntax for referencing claims",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clock_skew_leeway": schema.Int64Attribute{
						Description:         "The amount of leeway to add to all claims to account for clock skew, in seconds. Defaults to 60 seconds if set to 0 and can be disabled if set to -1. Accepts an integer number of seconds, or a Go duration format string. Only applicable with 'jwt' roles",
						MarkdownDescription: "The amount of leeway to add to all claims to account for clock skew, in seconds. Defaults to 60 seconds if set to 0 and can be disabled if set to -1. Accepts an integer number of seconds, or a Go duration format string. Only applicable with 'jwt' roles",
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

					"expiration_leeway": schema.Int64Attribute{
						Description:         "The amount of leeway to add to expiration (exp) claims to account for clock skew, in seconds. Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Accepts an integer number of seconds, or a Go duration format string. Only applicable with 'jwt' roles.",
						MarkdownDescription: "The amount of leeway to add to expiration (exp) claims to account for clock skew, in seconds. Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Accepts an integer number of seconds, or a Go duration format string. Only applicable with 'jwt' roles.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups_claim": schema.StringAttribute{
						Description:         "The claim to use to uniquely identify the set of groups to which the user belongs; this will be used as the names for the Identity group aliases created due to a successful login. The claim value must be a list of strings. Supports JSON pointer syntax for referencing claims",
						MarkdownDescription: "The claim to use to uniquely identify the set of groups to which the user belongs; this will be used as the names for the Identity group aliases created due to a successful login. The claim value must be a list of strings. Supports JSON pointer syntax for referencing claims",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"maxage": schema.Int64Attribute{
						Description:         "Specifies the allowable elapsed time in seconds since the last time the user was actively authenticated with the OIDC provider If set, the max_age request parameter will be included in the authentication request See AuthRequest for additional details Accepts an integer number of seconds, or a Go duration format string",
						MarkdownDescription: "Specifies the allowable elapsed time in seconds since the last time the user was actively authenticated with the OIDC provider If set, the max_age request parameter will be included in the authentication request See AuthRequest for additional details Accepts an integer number of seconds, or a Go duration format string",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the role",
						MarkdownDescription: "Name of the role",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"not_before_leeway": schema.Int64Attribute{
						Description:         "he amount of leeway to add to not before (nbf) claims to account for clock skew, in seconds Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Accepts an integer number of seconds, or a Go duration format string. Only applicable with 'jwt' roles",
						MarkdownDescription: "he amount of leeway to add to not before (nbf) claims to account for clock skew, in seconds Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Accepts an integer number of seconds, or a Go duration format string. Only applicable with 'jwt' roles",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/auth/{spec.path}/groups/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						MarkdownDescription: "Path at which to make the configuration. The final path in Vault will be {[spec.authentication.namespace]}/auth/{spec.path}/groups/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
						},
					},

					"role_type": schema.StringAttribute{
						Description:         "Type of role, either 'oidc' (default) or 'jwt'",
						MarkdownDescription: "Type of role, either 'oidc' (default) or 'jwt'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_bound_cidrs": schema.ListAttribute{
						Description:         "List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_explicit_max_ttl": schema.StringAttribute{
						Description:         "If set, will encode an explicit max TTL onto the token. This is a hard cap even if token_ttl and token_max_ttl would otherwise allow a renewal.",
						MarkdownDescription: "If set, will encode an explicit max TTL onto the token. This is a hard cap even if token_ttl and token_max_ttl would otherwise allow a renewal.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_max_ttl": schema.StringAttribute{
						Description:         "The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time",
						MarkdownDescription: "The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_no_default_policy": schema.BoolAttribute{
						Description:         "If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies",
						MarkdownDescription: "If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_num_uses": schema.Int64Attribute{
						Description:         "The maximum number of times a generated token may be used (within its lifetime); 0 means unlimited. If you require the token to have the ability to create child tokens, you will need to set this value to 0",
						MarkdownDescription: "The maximum number of times a generated token may be used (within its lifetime); 0 means unlimited. If you require the token to have the ability to create child tokens, you will need to set this value to 0",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_period": schema.Int64Attribute{
						Description:         "The period, if any, to set on the token",
						MarkdownDescription: "The period, if any, to set on the token",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_policies": schema.ListAttribute{
						Description:         "List of policies to encode onto generated tokens Depending on the auth method, this list may be supplemented by user/group/other values kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "List of policies to encode onto generated tokens Depending on the auth method, this list may be supplemented by user/group/other values kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_ttl": schema.StringAttribute{
						Description:         "The incremental lifetime for generated tokens This current value of this will be referenced at renewal time",
						MarkdownDescription: "The incremental lifetime for generated tokens This current value of this will be referenced at renewal time",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_type": schema.StringAttribute{
						Description:         "The type of token that should be generated. Can be service, batch, or default to use the mount's tuned default (which unless changed will be service tokens). For token store roles, there are two additional possibilities: default-service and default-batch which specify the type to return unless the client requests a different type at generation time.",
						MarkdownDescription: "The type of token that should be generated. Can be service, batch, or default to use the mount's tuned default (which unless changed will be service tokens). For token store roles, there are two additional possibilities: default-service and default-batch which specify the type to return unless the client requests a different type at generation time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_claim": schema.StringAttribute{
						Description:         "The claim to use to uniquely identify the user; this will be used as the name for the Identity entity alias created due to a successful login. The claim value must be a string",
						MarkdownDescription: "The claim to use to uniquely identify the user; this will be used as the name for the Identity entity alias created due to a successful login. The claim value must be a string",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"user_claim_json_pointer": schema.BoolAttribute{
						Description:         "Specifies if the user_claim value uses JSON pointer syntax for referencing claims. By default, the user_claim value will not use JSON pointer.",
						MarkdownDescription: "Specifies if the user_claim value uses JSON pointer syntax for referencing claims. By default, the user_claim value will not use JSON pointer.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"verbose_oidc_logging": schema.BoolAttribute{
						Description:         "Log received OIDC tokens and claims when debug-level logging is active Not recommended in production since sensitive information may be present in OIDC responses",
						MarkdownDescription: "Log received OIDC tokens and claims when debug-level logging is active Not recommended in production since sensitive information may be present in OIDC responses",
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

func (r *RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_jwtoidc_auth_engine_role_v1alpha1_manifest")

	var model RedhatcopRedhatIoJwtoidcauthEngineRoleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("JWTOIDCAuthEngineRole")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
