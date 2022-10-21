/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type GetambassadorIoHostV3Alpha1Resource struct{}

var (
	_ resource.Resource = (*GetambassadorIoHostV3Alpha1Resource)(nil)
)

type GetambassadorIoHostV3Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GetambassadorIoHostV3Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AcmeProvider *struct {
			Authority *string `tfsdk:"authority" yaml:"authority,omitempty"`

			Email *string `tfsdk:"email" yaml:"email,omitempty"`

			PrivateKeySecret *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"private_key_secret" yaml:"privateKeySecret,omitempty"`

			Registration *string `tfsdk:"registration" yaml:"registration,omitempty"`
		} `tfsdk:"acme_provider" yaml:"acmeProvider,omitempty"`

		Ambassador_id *[]string `tfsdk:"ambassador_id" yaml:"ambassador_id,omitempty"`

		Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

		MappingSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"mapping_selector" yaml:"mappingSelector,omitempty"`

		PreviewUrl *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"preview_url" yaml:"previewUrl,omitempty"`

		RequestPolicy *struct {
			Insecure *struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				AdditionalPort *int64 `tfsdk:"additional_port" yaml:"additionalPort,omitempty"`
			} `tfsdk:"insecure" yaml:"insecure,omitempty"`
		} `tfsdk:"request_policy" yaml:"requestPolicy,omitempty"`

		Selector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`

		Tls *struct {
			Alpn_protocols *string `tfsdk:"alpn_protocols" yaml:"alpn_protocols,omitempty"`

			Ca_secret *string `tfsdk:"ca_secret" yaml:"ca_secret,omitempty"`

			Cacert_chain_file *string `tfsdk:"cacert_chain_file" yaml:"cacert_chain_file,omitempty"`

			Cert_chain_file *string `tfsdk:"cert_chain_file" yaml:"cert_chain_file,omitempty"`

			Cert_required *bool `tfsdk:"cert_required" yaml:"cert_required,omitempty"`

			Cipher_suites *[]string `tfsdk:"cipher_suites" yaml:"cipher_suites,omitempty"`

			Crl_secret *string `tfsdk:"crl_secret" yaml:"crl_secret,omitempty"`

			Ecdh_curves *[]string `tfsdk:"ecdh_curves" yaml:"ecdh_curves,omitempty"`

			Max_tls_version *string `tfsdk:"max_tls_version" yaml:"max_tls_version,omitempty"`

			Min_tls_version *string `tfsdk:"min_tls_version" yaml:"min_tls_version,omitempty"`

			Private_key_file *string `tfsdk:"private_key_file" yaml:"private_key_file,omitempty"`

			Redirect_cleartext_from *int64 `tfsdk:"redirect_cleartext_from" yaml:"redirect_cleartext_from,omitempty"`

			Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`
		} `tfsdk:"tls" yaml:"tls,omitempty"`

		TlsContext *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"tls_context" yaml:"tlsContext,omitempty"`

		TlsSecret *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"tls_secret" yaml:"tlsSecret,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGetambassadorIoHostV3Alpha1Resource() resource.Resource {
	return &GetambassadorIoHostV3Alpha1Resource{}
}

func (r *GetambassadorIoHostV3Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_getambassador_io_host_v3alpha1"
}

func (r *GetambassadorIoHostV3Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Host is the Schema for the hosts API",
		MarkdownDescription: "Host is the Schema for the hosts API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "HostSpec defines the desired state of Host",
				MarkdownDescription: "HostSpec defines the desired state of Host",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"acme_provider": {
						Description:         "Specifies whether/who to talk ACME with to automatically manage the $tlsSecret.",
						MarkdownDescription: "Specifies whether/who to talk ACME with to automatically manage the $tlsSecret.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"authority": {
								Description:         "Specifies who to talk ACME with to get certs. Defaults to Let's Encrypt; if 'none' (case-insensitive), do not try to do ACME for this Host.",
								MarkdownDescription: "Specifies who to talk ACME with to get certs. Defaults to Let's Encrypt; if 'none' (case-insensitive), do not try to do ACME for this Host.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"email": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"private_key_secret": {
								Description:         "Specifies the Kubernetes Secret to use to store the private key of the ACME account (essentially, where to store the auto-generated password for the auto-created ACME account).  You should not normally need to set this--the default value is based on a combination of the ACME authority being registered wit and the email address associated with the account.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
								MarkdownDescription: "Specifies the Kubernetes Secret to use to store the private key of the ACME account (essentially, where to store the auto-generated password for the auto-created ACME account).  You should not normally need to set this--the default value is based on a combination of the ACME authority being registered wit and the email address associated with the account.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"registration": {
								Description:         "This is normally set automatically",
								MarkdownDescription: "This is normally set automatically",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ambassador_id": {
						Description:         "Common to all Ambassador objects (and optional).",
						MarkdownDescription: "Common to all Ambassador objects (and optional).",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hostname": {
						Description:         "Hostname by which the Ambassador can be reached.",
						MarkdownDescription: "Hostname by which the Ambassador can be reached.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mapping_selector": {
						Description:         "Selector for Mappings we'll associate with this Host. At the moment, Selector and MappingSelector are synonyms, but that will change soon.",
						MarkdownDescription: "Selector for Mappings we'll associate with this Host. At the moment, Selector and MappingSelector are synonyms, but that will change soon.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"preview_url": {
						Description:         "Configuration for the Preview URL feature of Service Preview. Defaults to preview URLs not enabled.",
						MarkdownDescription: "Configuration for the Preview URL feature of Service Preview. Defaults to preview URLs not enabled.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Is the Preview URL feature enabled?",
								MarkdownDescription: "Is the Preview URL feature enabled?",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "What type of Preview URL is allowed?",
								MarkdownDescription: "What type of Preview URL is allowed?",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Path"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"request_policy": {
						Description:         "Request policy definition.",
						MarkdownDescription: "Request policy definition.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"insecure": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Redirect", "Reject", "Route"),
										},
									},

									"additional_port": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "DEPRECATED: Selector by which we can find further configuration. Use MappingSelector instead.  TODO(lukeshu): In v3alpha2, figure out how to get rid of HostSpec.DeprecatedSelector.",
						MarkdownDescription: "DEPRECATED: Selector by which we can find further configuration. Use MappingSelector instead.  TODO(lukeshu): In v3alpha2, figure out how to get rid of HostSpec.DeprecatedSelector.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": {
						Description:         "TLS configuration.  It is not valid to specify both 'tlsContext' and 'tls'.",
						MarkdownDescription: "TLS configuration.  It is not valid to specify both 'tlsContext' and 'tls'.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"alpn_protocols": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ca_secret": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cacert_chain_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cert_chain_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cert_required": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cipher_suites": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"crl_secret": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ecdh_curves": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_tls_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"min_tls_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"private_key_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"redirect_cleartext_from": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sni": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls_context": {
						Description:         "Name of the TLSContext the Host resource is linked with. It is not valid to specify both 'tlsContext' and 'tls'.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",
						MarkdownDescription: "Name of the TLSContext the Host resource is linked with. It is not valid to specify both 'tlsContext' and 'tls'.  Note that this is a native-Kubernetes-style core.v1.LocalObjectReference, not an Ambassador-style '{name}.{namespace}' string.  Because we're opinionated, it does not support referencing a Secret in another namespace (because most native Kubernetes resources don't support that), but if we ever abandon that opinion and decide to support non-local references it, it would be by adding a 'namespace:' field by changing it from a core.v1.LocalObjectReference to a core.v1.SecretReference, not by adopting the '{name}.{namespace}' notation.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls_secret": {
						Description:         "Name of the Kubernetes secret into which to save generated certificates.  If ACME is enabled (see $acmeProvider), then the default is $hostname; otherwise the default is ''.  If the value is '', then we do not do TLS for this Host.",
						MarkdownDescription: "Name of the Kubernetes secret into which to save generated certificates.  If ACME is enabled (see $acmeProvider), then the default is $hostname; otherwise the default is ''.  If the value is '', then we do not do TLS for this Host.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "Name is unique within a namespace to reference a secret resource.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "Namespace defines the space within which the secret name must be unique.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *GetambassadorIoHostV3Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_getambassador_io_host_v3alpha1")

	var state GetambassadorIoHostV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoHostV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("Host")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoHostV3Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_getambassador_io_host_v3alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *GetambassadorIoHostV3Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_getambassador_io_host_v3alpha1")

	var state GetambassadorIoHostV3Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GetambassadorIoHostV3Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("getambassador.io/v3alpha1")
	goModel.Kind = utilities.Ptr("Host")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GetambassadorIoHostV3Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_getambassador_io_host_v3alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
