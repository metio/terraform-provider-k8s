/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_nginx_org_v1

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
	_ datasource.DataSource = &K8SNginxOrgPolicyV1Manifest{}
)

func NewK8SNginxOrgPolicyV1Manifest() datasource.DataSource {
	return &K8SNginxOrgPolicyV1Manifest{}
}

type K8SNginxOrgPolicyV1Manifest struct{}

type K8SNginxOrgPolicyV1ManifestData struct {
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
		AccessControl *struct {
			Allow *[]string `tfsdk:"allow" json:"allow,omitempty"`
			Deny  *[]string `tfsdk:"deny" json:"deny,omitempty"`
		} `tfsdk:"access_control" json:"accessControl,omitempty"`
		ApiKey *struct {
			ClientSecret *string `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			SuppliedIn   *struct {
				Header *[]string `tfsdk:"header" json:"header,omitempty"`
				Query  *[]string `tfsdk:"query" json:"query,omitempty"`
			} `tfsdk:"supplied_in" json:"suppliedIn,omitempty"`
		} `tfsdk:"api_key" json:"apiKey,omitempty"`
		BasicAuth *struct {
			Realm  *string `tfsdk:"realm" json:"realm,omitempty"`
			Secret *string `tfsdk:"secret" json:"secret,omitempty"`
		} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
		EgressMTLS *struct {
			Ciphers           *string `tfsdk:"ciphers" json:"ciphers,omitempty"`
			Protocols         *string `tfsdk:"protocols" json:"protocols,omitempty"`
			ServerName        *bool   `tfsdk:"server_name" json:"serverName,omitempty"`
			SessionReuse      *bool   `tfsdk:"session_reuse" json:"sessionReuse,omitempty"`
			SslName           *string `tfsdk:"ssl_name" json:"sslName,omitempty"`
			TlsSecret         *string `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
			TrustedCertSecret *string `tfsdk:"trusted_cert_secret" json:"trustedCertSecret,omitempty"`
			VerifyDepth       *int64  `tfsdk:"verify_depth" json:"verifyDepth,omitempty"`
			VerifyServer      *bool   `tfsdk:"verify_server" json:"verifyServer,omitempty"`
		} `tfsdk:"egress_mtls" json:"egressMTLS,omitempty"`
		IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
		IngressMTLS      *struct {
			ClientCertSecret *string `tfsdk:"client_cert_secret" json:"clientCertSecret,omitempty"`
			CrlFileName      *string `tfsdk:"crl_file_name" json:"crlFileName,omitempty"`
			VerifyClient     *string `tfsdk:"verify_client" json:"verifyClient,omitempty"`
			VerifyDepth      *int64  `tfsdk:"verify_depth" json:"verifyDepth,omitempty"`
		} `tfsdk:"ingress_mtls" json:"ingressMTLS,omitempty"`
		Jwt *struct {
			JwksURI  *string `tfsdk:"jwks_uri" json:"jwksURI,omitempty"`
			KeyCache *string `tfsdk:"key_cache" json:"keyCache,omitempty"`
			Realm    *string `tfsdk:"realm" json:"realm,omitempty"`
			Secret   *string `tfsdk:"secret" json:"secret,omitempty"`
			Token    *string `tfsdk:"token" json:"token,omitempty"`
		} `tfsdk:"jwt" json:"jwt,omitempty"`
		Oidc *struct {
			AccessTokenEnable     *bool     `tfsdk:"access_token_enable" json:"accessTokenEnable,omitempty"`
			AuthEndpoint          *string   `tfsdk:"auth_endpoint" json:"authEndpoint,omitempty"`
			AuthExtraArgs         *[]string `tfsdk:"auth_extra_args" json:"authExtraArgs,omitempty"`
			ClientID              *string   `tfsdk:"client_id" json:"clientID,omitempty"`
			ClientSecret          *string   `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			EndSessionEndpoint    *string   `tfsdk:"end_session_endpoint" json:"endSessionEndpoint,omitempty"`
			JwksURI               *string   `tfsdk:"jwks_uri" json:"jwksURI,omitempty"`
			PostLogoutRedirectURI *string   `tfsdk:"post_logout_redirect_uri" json:"postLogoutRedirectURI,omitempty"`
			RedirectURI           *string   `tfsdk:"redirect_uri" json:"redirectURI,omitempty"`
			Scope                 *string   `tfsdk:"scope" json:"scope,omitempty"`
			TokenEndpoint         *string   `tfsdk:"token_endpoint" json:"tokenEndpoint,omitempty"`
			ZoneSyncLeeway        *int64    `tfsdk:"zone_sync_leeway" json:"zoneSyncLeeway,omitempty"`
		} `tfsdk:"oidc" json:"oidc,omitempty"`
		RateLimit *struct {
			Burst      *int64  `tfsdk:"burst" json:"burst,omitempty"`
			Delay      *int64  `tfsdk:"delay" json:"delay,omitempty"`
			DryRun     *bool   `tfsdk:"dry_run" json:"dryRun,omitempty"`
			Key        *string `tfsdk:"key" json:"key,omitempty"`
			LogLevel   *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			NoDelay    *bool   `tfsdk:"no_delay" json:"noDelay,omitempty"`
			Rate       *string `tfsdk:"rate" json:"rate,omitempty"`
			RejectCode *int64  `tfsdk:"reject_code" json:"rejectCode,omitempty"`
			Scale      *bool   `tfsdk:"scale" json:"scale,omitempty"`
			ZoneSize   *string `tfsdk:"zone_size" json:"zoneSize,omitempty"`
		} `tfsdk:"rate_limit" json:"rateLimit,omitempty"`
		Waf *struct {
			ApBundle    *string `tfsdk:"ap_bundle" json:"apBundle,omitempty"`
			ApPolicy    *string `tfsdk:"ap_policy" json:"apPolicy,omitempty"`
			Enable      *bool   `tfsdk:"enable" json:"enable,omitempty"`
			SecurityLog *struct {
				ApLogBundle *string `tfsdk:"ap_log_bundle" json:"apLogBundle,omitempty"`
				ApLogConf   *string `tfsdk:"ap_log_conf" json:"apLogConf,omitempty"`
				Enable      *bool   `tfsdk:"enable" json:"enable,omitempty"`
				LogDest     *string `tfsdk:"log_dest" json:"logDest,omitempty"`
			} `tfsdk:"security_log" json:"securityLog,omitempty"`
			SecurityLogs *[]struct {
				ApLogBundle *string `tfsdk:"ap_log_bundle" json:"apLogBundle,omitempty"`
				ApLogConf   *string `tfsdk:"ap_log_conf" json:"apLogConf,omitempty"`
				Enable      *bool   `tfsdk:"enable" json:"enable,omitempty"`
				LogDest     *string `tfsdk:"log_dest" json:"logDest,omitempty"`
			} `tfsdk:"security_logs" json:"securityLogs,omitempty"`
		} `tfsdk:"waf" json:"waf,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SNginxOrgPolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_nginx_org_policy_v1_manifest"
}

func (r *K8SNginxOrgPolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Policy defines a Policy for VirtualServer and VirtualServerRoute resources.",
		MarkdownDescription: "Policy defines a Policy for VirtualServer and VirtualServerRoute resources.",
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
				Description:         "PolicySpec is the spec of the Policy resource. The spec includes multiple fields, where each field represents a different policy. Only one policy (field) is allowed.",
				MarkdownDescription: "PolicySpec is the spec of the Policy resource. The spec includes multiple fields, where each field represents a different policy. Only one policy (field) is allowed.",
				Attributes: map[string]schema.Attribute{
					"access_control": schema.SingleNestedAttribute{
						Description:         "AccessControl defines an access policy based on the source IP of a request.",
						MarkdownDescription: "AccessControl defines an access policy based on the source IP of a request.",
						Attributes: map[string]schema.Attribute{
							"allow": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"deny": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"api_key": schema.SingleNestedAttribute{
						Description:         "APIKey defines an API Key policy.",
						MarkdownDescription: "APIKey defines an API Key policy.",
						Attributes: map[string]schema.Attribute{
							"client_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supplied_in": schema.SingleNestedAttribute{
								Description:         "SuppliedIn defines the locations API Key should be supplied in.",
								MarkdownDescription: "SuppliedIn defines the locations API Key should be supplied in.",
								Attributes: map[string]schema.Attribute{
									"header": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
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

					"basic_auth": schema.SingleNestedAttribute{
						Description:         "BasicAuth holds HTTP Basic authentication configuration",
						MarkdownDescription: "BasicAuth holds HTTP Basic authentication configuration",
						Attributes: map[string]schema.Attribute{
							"realm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret": schema.StringAttribute{
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

					"egress_mtls": schema.SingleNestedAttribute{
						Description:         "EgressMTLS defines an Egress MTLS policy.",
						MarkdownDescription: "EgressMTLS defines an Egress MTLS policy.",
						Attributes: map[string]schema.Attribute{
							"ciphers": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocols": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_name": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"session_reuse": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ssl_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"trusted_cert_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_depth": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_server": schema.BoolAttribute{
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

					"ingress_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_mtls": schema.SingleNestedAttribute{
						Description:         "IngressMTLS defines an Ingress MTLS policy.",
						MarkdownDescription: "IngressMTLS defines an Ingress MTLS policy.",
						Attributes: map[string]schema.Attribute{
							"client_cert_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"crl_file_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_client": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"verify_depth": schema.Int64Attribute{
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

					"jwt": schema.SingleNestedAttribute{
						Description:         "JWTAuth holds JWT authentication configuration.",
						MarkdownDescription: "JWTAuth holds JWT authentication configuration.",
						Attributes: map[string]schema.Attribute{
							"jwks_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key_cache": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"realm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token": schema.StringAttribute{
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

					"oidc": schema.SingleNestedAttribute{
						Description:         "OIDC defines an Open ID Connect policy.",
						MarkdownDescription: "OIDC defines an Open ID Connect policy.",
						Attributes: map[string]schema.Attribute{
							"access_token_enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auth_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auth_extra_args": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"end_session_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jwks_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"post_logout_redirect_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"redirect_uri": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scope": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"token_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zone_sync_leeway": schema.Int64Attribute{
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

					"rate_limit": schema.SingleNestedAttribute{
						Description:         "RateLimit defines a rate limit policy.",
						MarkdownDescription: "RateLimit defines a rate limit policy.",
						Attributes: map[string]schema.Attribute{
							"burst": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delay": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dry_run": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"no_delay": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rate": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reject_code": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scale": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zone_size": schema.StringAttribute{
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

					"waf": schema.SingleNestedAttribute{
						Description:         "WAF defines an WAF policy.",
						MarkdownDescription: "WAF defines an WAF policy.",
						Attributes: map[string]schema.Attribute{
							"ap_bundle": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ap_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_log": schema.SingleNestedAttribute{
								Description:         "SecurityLog defines the security log of a WAF policy.",
								MarkdownDescription: "SecurityLog defines the security log of a WAF policy.",
								Attributes: map[string]schema.Attribute{
									"ap_log_bundle": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ap_log_conf": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_dest": schema.StringAttribute{
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

							"security_logs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ap_log_bundle": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ap_log_conf": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enable": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"log_dest": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
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
		},
	}
}

func (r *K8SNginxOrgPolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_nginx_org_policy_v1_manifest")

	var model K8SNginxOrgPolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.nginx.org/v1")
	model.Kind = pointer.String("Policy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
