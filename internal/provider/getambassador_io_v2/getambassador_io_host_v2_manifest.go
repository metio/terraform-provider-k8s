/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v2

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
	_ datasource.DataSource = &GetambassadorIoHostV2Manifest{}
)

func NewGetambassadorIoHostV2Manifest() datasource.DataSource {
	return &GetambassadorIoHostV2Manifest{}
}

type GetambassadorIoHostV2Manifest struct{}

type GetambassadorIoHostV2ManifestData struct {
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
		AcmeProvider *struct {
			Authority        *string `tfsdk:"authority" json:"authority,omitempty"`
			Email            *string `tfsdk:"email" json:"email,omitempty"`
			PrivateKeySecret *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"private_key_secret" json:"privateKeySecret,omitempty"`
			Registration *string `tfsdk:"registration" json:"registration,omitempty"`
		} `tfsdk:"acme_provider" json:"acmeProvider,omitempty"`
		Ambassador_id *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Hostname      *string   `tfsdk:"hostname" json:"hostname,omitempty"`
		PreviewUrl    *struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Type    *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"preview_url" json:"previewUrl,omitempty"`
		RequestPolicy *struct {
			Insecure *struct {
				Action         *string `tfsdk:"action" json:"action,omitempty"`
				AdditionalPort *int64  `tfsdk:"additional_port" json:"additionalPort,omitempty"`
			} `tfsdk:"insecure" json:"insecure,omitempty"`
		} `tfsdk:"request_policy" json:"requestPolicy,omitempty"`
		Selector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Tls *struct {
			Alpn_protocols          *string   `tfsdk:"alpn_protocols" json:"alpn_protocols,omitempty"`
			Ca_secret               *string   `tfsdk:"ca_secret" json:"ca_secret,omitempty"`
			Cacert_chain_file       *string   `tfsdk:"cacert_chain_file" json:"cacert_chain_file,omitempty"`
			Cert_chain_file         *string   `tfsdk:"cert_chain_file" json:"cert_chain_file,omitempty"`
			Cert_required           *bool     `tfsdk:"cert_required" json:"cert_required,omitempty"`
			Cipher_suites           *[]string `tfsdk:"cipher_suites" json:"cipher_suites,omitempty"`
			Ecdh_curves             *[]string `tfsdk:"ecdh_curves" json:"ecdh_curves,omitempty"`
			Max_tls_version         *string   `tfsdk:"max_tls_version" json:"max_tls_version,omitempty"`
			Min_tls_version         *string   `tfsdk:"min_tls_version" json:"min_tls_version,omitempty"`
			Private_key_file        *string   `tfsdk:"private_key_file" json:"private_key_file,omitempty"`
			Redirect_cleartext_from *int64    `tfsdk:"redirect_cleartext_from" json:"redirect_cleartext_from,omitempty"`
			Sni                     *string   `tfsdk:"sni" json:"sni,omitempty"`
			V3CRLSecret             *string   `tfsdk:"v3_crl_secret" json:"v3CRLSecret,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		TlsContext *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"tls_context" json:"tlsContext,omitempty"`
		TlsSecret *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"tls_secret" json:"tlsSecret,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GetambassadorIoHostV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_host_v2_manifest"
}

func (r *GetambassadorIoHostV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Host is the Schema for the hosts API",
		MarkdownDescription: "Host is the Schema for the hosts API",
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
				Description:         "HostSpec defines the desired state of Host",
				MarkdownDescription: "HostSpec defines the desired state of Host",
				Attributes: map[string]schema.Attribute{
					"acme_provider": schema.SingleNestedAttribute{
						Description:         "Specifies whether/who to talk ACME with to automatically manage the $tlsSecret.",
						MarkdownDescription: "Specifies whether/who to talk ACME with to automatically manage the $tlsSecret.",
						Attributes: map[string]schema.Attribute{
							"authority": schema.StringAttribute{
								Description:         "Specifies who to talk ACME with to get certs. Defaults to Let's Encrypt; if 'none' (case-insensitive), do not try to do ACME for this Host.",
								MarkdownDescription: "Specifies who to talk ACME with to get certs. Defaults to Let's Encrypt; if 'none' (case-insensitive), do not try to do ACME for this Host.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"email": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_key_secret": schema.SingleNestedAttribute{
								Description:         "Specifies the Kubernetes Secret to use to store the private key of the ACME account (essentially, where to store the auto-generated password for the auto-created ACME account).  You should not normally need to set this--the default value is based on a combination of the ACME authority being registered wit and the email address associated with the account.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
								MarkdownDescription: "Specifies the Kubernetes Secret to use to store the private key of the ACME account (essentially, where to store the auto-generated password for the auto-created ACME account).  You should not normally need to set this--the default value is based on a combination of the ACME authority being registered wit and the email address associated with the account.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"registration": schema.StringAttribute{
								Description:         "This is normally set automatically",
								MarkdownDescription: "This is normally set automatically",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ambassador_id": schema.ListAttribute{
						Description:         "Common to all Ambassador objects (and optional).",
						MarkdownDescription: "Common to all Ambassador objects (and optional).",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"hostname": schema.StringAttribute{
						Description:         "Hostname by which the Ambassador can be reached.",
						MarkdownDescription: "Hostname by which the Ambassador can be reached.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"preview_url": schema.SingleNestedAttribute{
						Description:         "Configuration for the Preview URL feature of Service Preview. Defaults to preview URLs not enabled.",
						MarkdownDescription: "Configuration for the Preview URL feature of Service Preview. Defaults to preview URLs not enabled.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Is the Preview URL feature enabled?",
								MarkdownDescription: "Is the Preview URL feature enabled?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "What type of Preview URL is allowed?",
								MarkdownDescription: "What type of Preview URL is allowed?",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Path"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"request_policy": schema.SingleNestedAttribute{
						Description:         "Request policy definition.",
						MarkdownDescription: "Request policy definition.",
						Attributes: map[string]schema.Attribute{
							"insecure": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Redirect", "Reject", "Route"),
										},
									},

									"additional_port": schema.Int64Attribute{
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

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector by which we can find further configuration. Defaults to hostname=$hostname",
						MarkdownDescription: "Selector by which we can find further configuration. Defaults to hostname=$hostname",
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

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS configuration.  It is not valid to specify both 'tlsContext' and 'tls'.",
						MarkdownDescription: "TLS configuration.  It is not valid to specify both 'tlsContext' and 'tls'.",
						Attributes: map[string]schema.Attribute{
							"alpn_protocols": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cacert_chain_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cert_chain_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cert_required": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cipher_suites": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ecdh_curves": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_tls_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_tls_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_key_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"redirect_cleartext_from": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sni": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"v3_crl_secret": schema.StringAttribute{
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

					"tls_context": schema.SingleNestedAttribute{
						Description:         "Name of the TLSContext the Host resource is linked with. It is not valid to specify both 'tlsContext' and 'tls'.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
						MarkdownDescription: "Name of the TLSContext the Host resource is linked with. It is not valid to specify both 'tlsContext' and 'tls'.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. TODO: Add other useful fields. apiVersion, kind, uid? More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls_secret": schema.SingleNestedAttribute{
						Description:         "Name of the Kubernetes secret into which to save generated certificates.  If ACME is enabled (see $acmeProvider), then the default is $hostname; otherwise the default is ''.  If the value is '', then we do not do TLS for this Host.",
						MarkdownDescription: "Name of the Kubernetes secret into which to save generated certificates.  If ACME is enabled (see $acmeProvider), then the default is $hostname; otherwise the default is ''.  If the value is '', then we do not do TLS for this Host.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

func (r *GetambassadorIoHostV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_host_v2_manifest")

	var model GetambassadorIoHostV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("getambassador.io/v2")
	model.Kind = pointer.String("Host")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
