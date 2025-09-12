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
	_ datasource.DataSource = &RedhatcopRedhatIoSecretEngineMountV1Alpha1Manifest{}
)

func NewRedhatcopRedhatIoSecretEngineMountV1Alpha1Manifest() datasource.DataSource {
	return &RedhatcopRedhatIoSecretEngineMountV1Alpha1Manifest{}
}

type RedhatcopRedhatIoSecretEngineMountV1Alpha1Manifest struct{}

type RedhatcopRedhatIoSecretEngineMountV1Alpha1ManifestData struct {
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
		Authentication *struct {
			Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Path           *string `tfsdk:"path" json:"path,omitempty"`
			Role           *string `tfsdk:"role" json:"role,omitempty"`
			ServiceAccount *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		Config *struct {
			AllowedResponseHeaders    *[]string `tfsdk:"allowed_response_headers" json:"allowedResponseHeaders,omitempty"`
			AuditNonHMACRequestKeys   *[]string `tfsdk:"audit_non_hmac_request_keys" json:"auditNonHMACRequestKeys,omitempty"`
			AuditNonHMACResponseKeys  *[]string `tfsdk:"audit_non_hmac_response_keys" json:"auditNonHMACResponseKeys,omitempty"`
			DefaultLeaseTTL           *string   `tfsdk:"default_lease_ttl" json:"defaultLeaseTTL,omitempty"`
			ForceNoCache              *bool     `tfsdk:"force_no_cache" json:"forceNoCache,omitempty"`
			ListingVisibility         *string   `tfsdk:"listing_visibility" json:"listingVisibility,omitempty"`
			MaxLeaseTTL               *string   `tfsdk:"max_lease_ttl" json:"maxLeaseTTL,omitempty"`
			PassthroughRequestHeaders *[]string `tfsdk:"passthrough_request_headers" json:"passthroughRequestHeaders,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
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
		Description           *string            `tfsdk:"description" json:"description,omitempty"`
		ExternalEntropyAccess *bool              `tfsdk:"external_entropy_access" json:"externalEntropyAccess,omitempty"`
		Local                 *bool              `tfsdk:"local" json:"local,omitempty"`
		Name                  *string            `tfsdk:"name" json:"name,omitempty"`
		Options               *map[string]string `tfsdk:"options" json:"options,omitempty"`
		Path                  *string            `tfsdk:"path" json:"path,omitempty"`
		SealWrap              *bool              `tfsdk:"seal_wrap" json:"sealWrap,omitempty"`
		Type                  *string            `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RedhatcopRedhatIoSecretEngineMountV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_redhatcop_redhat_io_secret_engine_mount_v1alpha1_manifest"
}

func (r *RedhatcopRedhatIoSecretEngineMountV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SecretEngineMount is the Schema for the secretenginemounts API",
		MarkdownDescription: "SecretEngineMount is the Schema for the secretenginemounts API",
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
				Description:         "SecretEngineMountSpec defines the desired state of SecretEngineMount",
				MarkdownDescription: "SecretEngineMountSpec defines the desired state of SecretEngineMount",
				Attributes: map[string]schema.Attribute{
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

					"config": schema.SingleNestedAttribute{
						Description:         "Specifies configuration options for this mount; if set on a specific mount, values will override any global defaults (e.g. the system TTL/Max TTL)",
						MarkdownDescription: "Specifies configuration options for this mount; if set on a specific mount, values will override any global defaults (e.g. the system TTL/Max TTL)",
						Attributes: map[string]schema.Attribute{
							"allowed_response_headers": schema.ListAttribute{
								Description:         "AllowedResponseHeaders list of headers to whitelist, allowing a plugin to include them in the response. kubebuilder:validation:UniqueItems=true",
								MarkdownDescription: "AllowedResponseHeaders list of headers to whitelist, allowing a plugin to include them in the response. kubebuilder:validation:UniqueItems=true",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"audit_non_hmac_request_keys": schema.ListAttribute{
								Description:         "AuditNonHMACRequestKeys list of keys that will not be HMAC'd by audit devices in the request data object. kubebuilder:validation:UniqueItems=true",
								MarkdownDescription: "AuditNonHMACRequestKeys list of keys that will not be HMAC'd by audit devices in the request data object. kubebuilder:validation:UniqueItems=true",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"audit_non_hmac_response_keys": schema.ListAttribute{
								Description:         "AuditNonHMACResponseKeys list of keys that will not be HMAC'd by audit devices in the response data object. kubebuilder:validation:UniqueItems=true",
								MarkdownDescription: "AuditNonHMACResponseKeys list of keys that will not be HMAC'd by audit devices in the response data object. kubebuilder:validation:UniqueItems=true",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_lease_ttl": schema.StringAttribute{
								Description:         "DefaultLeaseTTL The default lease duration, specified as a string duration like '5s' or '30m'.",
								MarkdownDescription: "DefaultLeaseTTL The default lease duration, specified as a string duration like '5s' or '30m'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"force_no_cache": schema.BoolAttribute{
								Description:         "ForceNoCache Disable caching.",
								MarkdownDescription: "ForceNoCache Disable caching.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"listing_visibility": schema.StringAttribute{
								Description:         "ListingVisibility Specifies whether to show this mount in the UI-specific listing endpoint. Valid values are 'unauth' or 'hidden'. If not set, behaves like 'hidden'",
								MarkdownDescription: "ListingVisibility Specifies whether to show this mount in the UI-specific listing endpoint. Valid values are 'unauth' or 'hidden'. If not set, behaves like 'hidden'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("unauth", "hidden"),
								},
							},

							"max_lease_ttl": schema.StringAttribute{
								Description:         "MaxLeaseTTL The maximum lease duration, specified as a string duration like '5s' or '30m'.",
								MarkdownDescription: "MaxLeaseTTL The maximum lease duration, specified as a string duration like '5s' or '30m'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"passthrough_request_headers": schema.ListAttribute{
								Description:         "PassthroughRequestHeaders list of headers to whitelist and pass from the request to the plugin. kubebuilder:validation:UniqueItems=true",
								MarkdownDescription: "PassthroughRequestHeaders list of headers to whitelist and pass from the request to the plugin. kubebuilder:validation:UniqueItems=true",
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

					"description": schema.StringAttribute{
						Description:         "Description Specifies the human-friendly description of the mount.",
						MarkdownDescription: "Description Specifies the human-friendly description of the mount.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"external_entropy_access": schema.BoolAttribute{
						Description:         "ExternalEntropyAccess Enable the secrets engine to access Vault's external entropy source.",
						MarkdownDescription: "ExternalEntropyAccess Enable the secrets engine to access Vault's external entropy source.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"local": schema.BoolAttribute{
						Description:         "Local Specifies if the secrets engine is a local mount only. Local mounts are not replicated nor (if a secondary) removed by replication.",
						MarkdownDescription: "Local Specifies if the secrets engine is a local mount only. Local mounts are not replicated nor (if a secondary) removed by replication.",
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

					"options": schema.MapAttribute{
						Description:         "Options Specifies mount type specific options that are passed to the backend.",
						MarkdownDescription: "Options Specifies mount type specific options that are passed to the backend.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"path": schema.StringAttribute{
						Description:         "Path at which this secret engine will be available The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path /sys/mounts/{[spec.authentication.namespace]}/{spec.path}/{metadata.name}.",
						MarkdownDescription: "Path at which this secret engine will be available The final path in Vault will be {[spec.authentication.namespace]}/{spec.path}/{metadata.name}. The authentication role must have the following capabilities = [ 'create', 'read', 'update', 'delete'] on that path /sys/mounts/{[spec.authentication.namespace]}/{spec.path}/{metadata.name}.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?:/?[\w;:@&=\$-\.\+]*)+/?`), ""),
						},
					},

					"seal_wrap": schema.BoolAttribute{
						Description:         "SealWrap Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.",
						MarkdownDescription: "SealWrap Enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type Specifies the type of the backend, such as 'aws'.",
						MarkdownDescription: "Type Specifies the type of the backend, such as 'aws'.",
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

func (r *RedhatcopRedhatIoSecretEngineMountV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_redhatcop_redhat_io_secret_engine_mount_v1alpha1_manifest")

	var model RedhatcopRedhatIoSecretEngineMountV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("redhatcop.redhat.io/v1alpha1")
	model.Kind = pointer.String("SecretEngineMount")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
