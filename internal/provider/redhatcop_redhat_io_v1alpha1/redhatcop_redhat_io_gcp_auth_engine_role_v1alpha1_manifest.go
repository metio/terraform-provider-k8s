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
	_ datasource.DataSource = &RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoGcpauthEngineRoleV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1Manifest{}
}

type RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1Manifest struct{}

type RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1ManifestData struct {
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
		AddGroupAliases   *bool `tfsdk:"add_group_aliases" json:"addGroupAliases,omitempty"`
		AllowGCEInference *bool `tfsdk:"allow_gce_inference" json:"allowGCEInference,omitempty"`
		Authentication    *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		BoundInstanceGroups  *[]string `tfsdk:"bound_instance_groups" json:"boundInstanceGroups,omitempty"`
		BoundLabels          *[]string `tfsdk:"bound_labels" json:"boundLabels,omitempty"`
		BoundProjects        *[]string `tfsdk:"bound_projects" json:"boundProjects,omitempty"`
		BoundRegions         *[]string `tfsdk:"bound_regions" json:"boundRegions,omitempty"`
		BoundServiceAccounts *[]string `tfsdk:"bound_service_accounts" json:"boundServiceAccounts,omitempty"`
		BoundZones           *[]string `tfsdk:"bound_zones" json:"boundZones,omitempty"`
		Connection           *struct {
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
		MaxJWTExp            *string   `tfsdk:"max_jwt_exp" json:"maxJWTExp,omitempty"`
		Name                 *string   `tfsdk:"name" json:"name,omitempty"`
		Path                 *string   `tfsdk:"path" json:"path,omitempty"`
		Policies             *[]string `tfsdk:"policies" json:"policies,omitempty"`
		TokenBoundCIDRs      *[]string `tfsdk:"token_bound_cidrs" json:"tokenBoundCIDRs,omitempty"`
		TokenExplicitMaxTTL  *string   `tfsdk:"token_explicit_max_ttl" json:"tokenExplicitMaxTTL,omitempty"`
		TokenMaxTTL          *string   `tfsdk:"token_max_ttl" json:"tokenMaxTTL,omitempty"`
		TokenNoDefaultPolicy *bool     `tfsdk:"token_no_default_policy" json:"tokenNoDefaultPolicy,omitempty"`
		TokenNumUses         *int64    `tfsdk:"token_num_uses" json:"tokenNumUses,omitempty"`
		TokenPeriod          *int64    `tfsdk:"token_period" json:"tokenPeriod,omitempty"`
		TokenPolicies        *[]string `tfsdk:"token_policies" json:"tokenPolicies,omitempty"`
		TokenTTL             *string   `tfsdk:"token_ttl" json:"tokenTTL,omitempty"`
		TokenType            *string   `tfsdk:"token_type" json:"tokenType,omitempty"`
		Type                 *string   `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_gcp_auth_engine_role_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GCPAuthEngineRole is the Schema for the gcpauthengineroles API",
		MarkdownDescription: "GCPAuthEngineRole is the Schema for the gcpauthengineroles API",
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
				Description:         "GCPAuthEngineRoleSpec defines the desired state of GCPAuthEngineRole",
				MarkdownDescription: "GCPAuthEngineRoleSpec defines the desired state of GCPAuthEngineRole",
				Attributes: map[string]schema.Attribute{
					"add_group_aliases": schema.BoolAttribute{
						Description:         "If true, any auth token generated under this token will have associated group aliases, namely project-$PROJECT_ID, folder-$PROJECT_ID, and organization-$ORG_ID for the entities project and all its folder or organization ancestors. This requires Vault to have IAM permission resourcemanager.projects.get.",
						MarkdownDescription: "If true, any auth token generated under this token will have associated group aliases, namely project-$PROJECT_ID, folder-$PROJECT_ID, and organization-$ORG_ID for the entities project and all its folder or organization ancestors. This requires Vault to have IAM permission resourcemanager.projects.get.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_gce_inference": schema.BoolAttribute{
						Description:         "A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.",
						MarkdownDescription: "A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.",
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

					"bound_instance_groups": schema.ListAttribute{
						Description:         "The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either bound_zones or bound_regions must be set too. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either bound_zones or bound_regions must be set too. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_labels": schema.ListAttribute{
						Description:         "A comma-separated list of GCP labels formatted as 'key:value' strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "A comma-separated list of GCP labels formatted as 'key:value' strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_projects": schema.ListAttribute{
						Description:         "An array of GCP project IDs. Only entities belonging to this project can authenticate under the role.",
						MarkdownDescription: "An array of GCP project IDs. Only entities belonging to this project can authenticate under the role.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_regions": schema.ListAttribute{
						Description:         "The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_service_accounts": schema.ListAttribute{
						Description:         "An array of service account emails or IDs that login is restricted to, either directly or through an associated instance. If set to *, all service accounts are allowed (you can bind this further using bound_projects.)",
						MarkdownDescription: "An array of service account emails or IDs that login is restricted to, either directly or through an associated instance. If set to *, all service accounts are allowed (you can bind this further using bound_projects.)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bound_zones": schema.ListAttribute{
						Description:         "The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
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

					"max_jwt_exp": schema.StringAttribute{
						Description:         "The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter exp. The GCE metadata tokens currently do not allow the exp claim to be customized. The following parameter is only valid when the role is of type 'iam'.",
						MarkdownDescription: "The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter exp. The GCE metadata tokens currently do not allow the exp claim to be customized. The following parameter is only valid when the role is of type 'iam'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the role.",
						MarkdownDescription: "Name of the role.",
						Required:            true,
						Optional:            false,
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

					"policies": schema.ListAttribute{
						Description:         "DEPRECATED: Please use the token_policies parameter instead. List of token policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "DEPRECATED: Please use the token_policies parameter instead. List of token policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values. kubebuilder:validation:UniqueItems=true",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_bound_cidrs": schema.ListAttribute{
						Description:         "List of CIDR blocks. If set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "List of CIDR blocks. If set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well. kubebuilder:validation:UniqueItems=true",
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
						Description:         "The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						MarkdownDescription: "The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_no_default_policy": schema.BoolAttribute{
						Description:         "If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.",
						MarkdownDescription: "If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.",
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

					"token_period": schema.Int64Attribute{
						Description:         "The maximum allowed period value when a periodic token is requested from this role.",
						MarkdownDescription: "The maximum allowed period value when a periodic token is requested from this role.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"token_policies": schema.ListAttribute{
						Description:         "List of token policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "List of token policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values. kubebuilder:validation:UniqueItems=true",
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

					"type": schema.StringAttribute{
						Description:         "The type of this role. Certain fields correspond to specific roles and will be rejected otherwise. Please see below for more information.",
						MarkdownDescription: "The type of this role. Certain fields correspond to specific roles and will be rejected otherwise. Please see below for more information.",
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
	}
}

func (r *RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_gcp_auth_engine_role_v1alpha1_manifest")

	var model RedhatcopRedhatIoGcpauthEngineRoleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("GCPAuthEngineRole")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
