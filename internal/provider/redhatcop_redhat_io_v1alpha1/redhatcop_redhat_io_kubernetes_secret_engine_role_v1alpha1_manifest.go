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
	_ datasource.DataSource = &RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1Manifest{}
}

type RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1Manifest struct{}

type RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1ManifestData struct {
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
		AllowedKubernetesNamespaceSelector *string   `tfsdk:"allowed_kubernetes_namespace_selector" json:"allowedKubernetesNamespaceSelector,omitempty"`
		AllowedKubernetesNamespaces        *[]string `tfsdk:"allowed_kubernetes_namespaces" json:"allowedKubernetesNamespaces,omitempty"`
		Authentication                     *struct {
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
		DefaultAudiences   *string            `tfsdk:"default_audiences" json:"defaultAudiences,omitempty"`
		DefaultTTL         *string            `tfsdk:"default_ttl" json:"defaultTTL,omitempty"`
		ExtraAnnotations   *map[string]string `tfsdk:"extra_annotations" json:"extraAnnotations,omitempty"`
		ExtraLabels        *map[string]string `tfsdk:"extra_labels" json:"extraLabels,omitempty"`
		GenerateRoleRules  *string            `tfsdk:"generate_role_rules" json:"generateRoleRules,omitempty"`
		KubernetesRoleName *string            `tfsdk:"kubernetes_role_name" json:"kubernetesRoleName,omitempty"`
		KubernetesRoleType *string            `tfsdk:"kubernetes_role_type" json:"kubernetesRoleType,omitempty"`
		MaxTTL             *string            `tfsdk:"max_ttl" json:"maxTTL,omitempty"`
		Name               *string            `tfsdk:"name" json:"name,omitempty"`
		NameTemplate       *string            `tfsdk:"name_template" json:"nameTemplate,omitempty"`
		Path               *string            `tfsdk:"path" json:"path,omitempty"`
		ServiceAccountName *string            `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		TargetNamespaces   *struct {
			TargetNamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"target_namespace_selector" json:"targetNamespaceSelector,omitempty"`
			TargetNamespaces *[]string `tfsdk:"target_namespaces" json:"targetNamespaces,omitempty"`
		} `tfsdk:"target_namespaces" json:"targetNamespaces,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_kubernetes_secret_engine_role_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KubernetesSecretEngineRole is the Schema for the kubernetessecretengineroles API",
		MarkdownDescription: "KubernetesSecretEngineRole is the Schema for the kubernetessecretengineroles API",
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
				Description:         "KubernetesSecretEngineRoleSpec defines the desired state of KubernetesSecretEngineRole",
				MarkdownDescription: "KubernetesSecretEngineRoleSpec defines the desired state of KubernetesSecretEngineRole",
				Attributes: map[string]schema.Attribute{
					"allowed_kubernetes_namespace_selector": schema.StringAttribute{
						Description:         "A label selector for Kubernetes namespaces in which credentials can be generated. Accepts either a JSON or YAML object. The value should be of type LabelSelector as illustrated: ''{'matchLabels':{'stage':'prod','sa-generator':'vault'}}'. If set with allowed_kubernetes_namespaces, the conditions are ORed.",
						MarkdownDescription: "A label selector for Kubernetes namespaces in which credentials can be generated. Accepts either a JSON or YAML object. The value should be of type LabelSelector as illustrated: ''{'matchLabels':{'stage':'prod','sa-generator':'vault'}}'. If set with allowed_kubernetes_namespaces, the conditions are ORed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_kubernetes_namespaces": schema.ListAttribute{
						Description:         "AllowedKubernetesNamespaces The list of Kubernetes namespaces this role can generate credentials for. If set to '*' all namespaces are allowed. kubebuilder:validation:UniqueItems=true",
						MarkdownDescription: "AllowedKubernetesNamespaces The list of Kubernetes namespaces this role can generate credentials for. If set to '*' all namespaces are allowed. kubebuilder:validation:UniqueItems=true",
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

					"default_audiences": schema.StringAttribute{
						Description:         "DefaultAudiences The default intended audiences for generated Kubernetes tokens, specified by a comma separated string. e.g 'custom-audience-0,custom-audience-1'. If not set or set to '', the Kubernetes cluster default for audiences of service account tokens will be used.",
						MarkdownDescription: "DefaultAudiences The default intended audiences for generated Kubernetes tokens, specified by a comma separated string. e.g 'custom-audience-0,custom-audience-1'. If not set or set to '', the Kubernetes cluster default for audiences of service account tokens will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"default_ttl": schema.StringAttribute{
						Description:         "DeafulTTL Specifies the TTL for the leases associated with this role. Accepts time suffixed strings ('1h') or an integer number of seconds. Defaults to system/engine default TTL time.",
						MarkdownDescription: "DeafulTTL Specifies the TTL for the leases associated with this role. Accepts time suffixed strings ('1h') or an integer number of seconds. Defaults to system/engine default TTL time.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_annotations": schema.MapAttribute{
						Description:         "ExtraAnnotations Additional annotations to apply to all generated Kubernetes objects. See the Kubernetes annotations documentation for more details on annotations.",
						MarkdownDescription: "ExtraAnnotations Additional annotations to apply to all generated Kubernetes objects. See the Kubernetes annotations documentation for more details on annotations.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_labels": schema.MapAttribute{
						Description:         "ExtraLabels Additional labels to apply to all generated Kubernetes objects. See the Kubernetes labels documentation for more details on labels.",
						MarkdownDescription: "ExtraLabels Additional labels to apply to all generated Kubernetes objects. See the Kubernetes labels documentation for more details on labels.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"generate_role_rules": schema.StringAttribute{
						Description:         "GenerateRoleRules The Role or ClusterRole rules to use when generating a role. Accepts either JSON or YAML formatted rules. If set, the entire chain of Kubernetes objects will be generated when credentials are requested. The value should be a rules key with an array of PolicyRule objects, as illustrated in the Kubernetes RBAC documentation and Sample Payload 3 below.",
						MarkdownDescription: "GenerateRoleRules The Role or ClusterRole rules to use when generating a role. Accepts either JSON or YAML formatted rules. If set, the entire chain of Kubernetes objects will be generated when credentials are requested. The value should be a rules key with an array of PolicyRule objects, as illustrated in the Kubernetes RBAC documentation and Sample Payload 3 below.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kubernetes_role_name": schema.StringAttribute{
						Description:         "KubernetesRoleName The pre-existing Role or ClusterRole to bind a generated service account to. If set, Kubernetes token, service account, and role binding objects will be created when credentials are requested. See the Kubernetes roles documentation for more details on Kubernetes roles.",
						MarkdownDescription: "KubernetesRoleName The pre-existing Role or ClusterRole to bind a generated service account to. If set, Kubernetes token, service account, and role binding objects will be created when credentials are requested. See the Kubernetes roles documentation for more details on Kubernetes roles.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kubernetes_role_type": schema.StringAttribute{
						Description:         "KubernetesRoleType Specifies whether the Kubernetes role is a Role or ClusterRole",
						MarkdownDescription: "KubernetesRoleType Specifies whether the Kubernetes role is a Role or ClusterRole",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Role", "ClusterRole"),
						},
					},

					"max_ttl": schema.StringAttribute{
						Description:         "MaxTTL Specifies the maximum TTL for the leases associated with this role. Accepts time suffixed strings ('1h') or an integer number of seconds. Defaults to system/mount default TTL time; this value is allowed to be less than the mount max TTL (or, if not set, the system max TTL), but it is not allowed to be longer. See also The TTL General Case.",
						MarkdownDescription: "MaxTTL Specifies the maximum TTL for the leases associated with this role. Accepts time suffixed strings ('1h') or an integer number of seconds. Defaults to system/mount default TTL time; this value is allowed to be less than the mount max TTL (or, if not set, the system max TTL), but it is not allowed to be longer. See also The TTL General Case.",
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

					"name_template": schema.StringAttribute{
						Description:         "NameTemplate The name template to use when generating service accounts, roles and role bindings. If unset, a default template is used. See username templating for details on how to write a custom template.",
						MarkdownDescription: "NameTemplate The name template to use when generating service accounts, roles and role bindings. If unset, a default template is used. See username templating for details on how to write a custom template.",
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

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName The pre-existing service account to generate tokens for. Mutually exclusive with all role parameters. If set, only a Kubernetes token will be created when credentials are requested. See the Kubernetes service account documentation for more details on service accounts.",
						MarkdownDescription: "ServiceAccountName The pre-existing service account to generate tokens for. Mutually exclusive with all role parameters. If set, only a Kubernetes token will be created when credentials are requested. See the Kubernetes service account documentation for more details on service accounts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_namespaces": schema.SingleNestedAttribute{
						Description:         "TargetNamespaces specifies how to retrieve the list of Kubernetes namespaces this role can generate credentials for.",
						MarkdownDescription: "TargetNamespaces specifies how to retrieve the list of Kubernetes namespaces this role can generate credentials for.",
						Attributes: map[string]schema.Attribute{
							"target_namespace_selector": schema.SingleNestedAttribute{
								Description:         "TargetNamespaceSelector is a selector of namespaces from which service accounts will receove this role. Either TargetNamespaceSelector or TargetNamespaces can be specified",
								MarkdownDescription: "TargetNamespaceSelector is a selector of namespaces from which service accounts will receove this role. Either TargetNamespaceSelector or TargetNamespaces can be specified",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

							"target_namespaces": schema.ListAttribute{
								Description:         "TargetNamespaces is a list of namespace from which service accounts will receive this role. Either TargetNamespaceSelector or TargetNamespaces can be specified. kubebuilder:validation:UniqueItems=true",
								MarkdownDescription: "TargetNamespaces is a list of namespace from which service accounts will receive this role. Either TargetNamespaceSelector or TargetNamespaces can be specified. kubebuilder:validation:UniqueItems=true",
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
		},
	}
}

func (r *RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_kubernetes_secret_engine_role_v1alpha1_manifest")

	var model RedhatcopRedhatIoKubernetesSecretEngineRoleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("KubernetesSecretEngineRole")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
