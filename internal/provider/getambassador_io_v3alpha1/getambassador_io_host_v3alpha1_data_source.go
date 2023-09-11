/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package getambassador_io_v3alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &GetambassadorIoHostV3Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &GetambassadorIoHostV3Alpha1DataSource{}
)

func NewGetambassadorIoHostV3Alpha1DataSource() datasource.DataSource {
	return &GetambassadorIoHostV3Alpha1DataSource{}
}

type GetambassadorIoHostV3Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GetambassadorIoHostV3Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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
		Ambassador_id   *[]string `tfsdk:"ambassador_id" json:"ambassador_id,omitempty"`
		Hostname        *string   `tfsdk:"hostname" json:"hostname,omitempty"`
		MappingSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"mapping_selector" json:"mappingSelector,omitempty"`
		PreviewUrl *struct {
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
			Crl_secret              *string   `tfsdk:"crl_secret" json:"crl_secret,omitempty"`
			Ecdh_curves             *[]string `tfsdk:"ecdh_curves" json:"ecdh_curves,omitempty"`
			Max_tls_version         *string   `tfsdk:"max_tls_version" json:"max_tls_version,omitempty"`
			Min_tls_version         *string   `tfsdk:"min_tls_version" json:"min_tls_version,omitempty"`
			Private_key_file        *string   `tfsdk:"private_key_file" json:"private_key_file,omitempty"`
			Redirect_cleartext_from *int64    `tfsdk:"redirect_cleartext_from" json:"redirect_cleartext_from,omitempty"`
			Sni                     *string   `tfsdk:"sni" json:"sni,omitempty"`
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

func (r *GetambassadorIoHostV3Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_getambassador_io_host_v3alpha1"
}

func (r *GetambassadorIoHostV3Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Host is the Schema for the hosts API",
		MarkdownDescription: "Host is the Schema for the hosts API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},

							"email": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"private_key_secret": schema.SingleNestedAttribute{
								Description:         "Specifies the Kubernetes Secret to use to store the private key of the ACME account (essentially, where to store the auto-generated password for the auto-created ACME account).  You should not normally need to set this--the default value is based on a combination of the ACME authority being registered wit and the email address associated with the account.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
								MarkdownDescription: "Specifies the Kubernetes Secret to use to store the private key of the ACME account (essentially, where to store the auto-generated password for the auto-created ACME account).  You should not normally need to set this--the default value is based on a combination of the ACME authority being registered wit and the email address associated with the account.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"registration": schema.StringAttribute{
								Description:         "This is normally set automatically",
								MarkdownDescription: "This is normally set automatically",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ambassador_id": schema.ListAttribute{
						Description:         "Common to all Ambassador objects (and optional).",
						MarkdownDescription: "Common to all Ambassador objects (and optional).",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"hostname": schema.StringAttribute{
						Description:         "Hostname by which the Ambassador can be reached.",
						MarkdownDescription: "Hostname by which the Ambassador can be reached.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mapping_selector": schema.SingleNestedAttribute{
						Description:         "Selector for Mappings we'll associate with this Host. At the moment, Selector and MappingSelector are synonyms, but that will change soon.",
						MarkdownDescription: "Selector for Mappings we'll associate with this Host. At the moment, Selector and MappingSelector are synonyms, but that will change soon.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"preview_url": schema.SingleNestedAttribute{
						Description:         "Configuration for the Preview URL feature of Service Preview. Defaults to preview URLs not enabled.",
						MarkdownDescription: "Configuration for the Preview URL feature of Service Preview. Defaults to preview URLs not enabled.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Is the Preview URL feature enabled?",
								MarkdownDescription: "Is the Preview URL feature enabled?",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "What type of Preview URL is allowed?",
								MarkdownDescription: "What type of Preview URL is allowed?",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"additional_port": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "DEPRECATED: Selector by which we can find further configuration. Use MappingSelector instead.",
						MarkdownDescription: "DEPRECATED: Selector by which we can find further configuration. Use MappingSelector instead.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS configuration.  It is not valid to specify both 'tlsContext' and 'tls'.",
						MarkdownDescription: "TLS configuration.  It is not valid to specify both 'tlsContext' and 'tls'.",
						Attributes: map[string]schema.Attribute{
							"alpn_protocols": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ca_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cacert_chain_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cert_chain_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cert_required": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cipher_suites": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"crl_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ecdh_curves": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_tls_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_tls_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"private_key_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"redirect_cleartext_from": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sni": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tls_context": schema.SingleNestedAttribute{
						Description:         "Name of the TLSContext the Host resource is linked with. It is not valid to specify both 'tlsContext' and 'tls'.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
						MarkdownDescription: "Name of the TLSContext the Host resource is linked with. It is not valid to specify both 'tlsContext' and 'tls'.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tls_secret": schema.SingleNestedAttribute{
						Description:         "Name of the Kubernetes secret into which to save generated certificates.  If ACME is enabled (see $acmeProvider), then the default is $hostname; otherwise the default is ''.  If the value is '', then we do not do TLS for this Host.",
						MarkdownDescription: "Name of the Kubernetes secret into which to save generated certificates.  If ACME is enabled (see $acmeProvider), then the default is $hostname; otherwise the default is ''.  If the value is '', then we do not do TLS for this Host.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *GetambassadorIoHostV3Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *GetambassadorIoHostV3Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_getambassador_io_host_v3alpha1")

	var data GetambassadorIoHostV3Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "getambassador.io", Version: "v3alpha1", Resource: "hosts"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse GetambassadorIoHostV3Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("getambassador.io/v3alpha1")
	data.Kind = pointer.String("Host")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
