/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package admissionregistration_k8s_io_v1

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
	_ datasource.DataSource              = &AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource{}
	_ datasource.DataSourceWithConfigure = &AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource{}
)

func NewAdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource() datasource.DataSource {
	return &AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource{}
}

type AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Webhooks *[]struct {
		AdmissionReviewVersions *[]string `tfsdk:"admission_review_versions" json:"admissionReviewVersions,omitempty"`
		ClientConfig            *struct {
			CaBundle *string `tfsdk:"ca_bundle" json:"caBundle,omitempty"`
			Service  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Path      *string `tfsdk:"path" json:"path,omitempty"`
				Port      *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"service" json:"service,omitempty"`
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"client_config" json:"clientConfig,omitempty"`
		FailurePolicy   *string `tfsdk:"failure_policy" json:"failurePolicy,omitempty"`
		MatchConditions *[]struct {
			Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"match_conditions" json:"matchConditions,omitempty"`
		MatchPolicy       *string `tfsdk:"match_policy" json:"matchPolicy,omitempty"`
		Name              *string `tfsdk:"name" json:"name,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		ObjectSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"object_selector" json:"objectSelector,omitempty"`
		Rules *[]struct {
			ApiGroups   *[]string `tfsdk:"api_groups" json:"apiGroups,omitempty"`
			ApiVersions *[]string `tfsdk:"api_versions" json:"apiVersions,omitempty"`
			Operations  *[]string `tfsdk:"operations" json:"operations,omitempty"`
			Resources   *[]string `tfsdk:"resources" json:"resources,omitempty"`
			Scope       *string   `tfsdk:"scope" json:"scope,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		SideEffects    *string `tfsdk:"side_effects" json:"sideEffects,omitempty"`
		TimeoutSeconds *int64  `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
	} `tfsdk:"webhooks" json:"webhooks,omitempty"`
}

func (r *AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_admissionregistration_k8s_io_validating_webhook_configuration_v1"
}

func (r *AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.",
		MarkdownDescription: "ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"webhooks": schema.ListNestedAttribute{
				Description:         "Webhooks is a list of webhooks and the affected resources and operations.",
				MarkdownDescription: "Webhooks is a list of webhooks and the affected resources and operations.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"admission_review_versions": schema.ListAttribute{
							Description:         "AdmissionReviewVersions is an ordered list of preferred 'AdmissionReview' versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy.",
							MarkdownDescription: "AdmissionReviewVersions is an ordered list of preferred 'AdmissionReview' versions the Webhook expects. API server will try to use first version in the list which it supports. If none of the versions specified in this list supported by API server, validation will fail for this object. If a persisted webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail and be subject to the failure policy.",
							ElementType:         types.StringType,
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"client_config": schema.SingleNestedAttribute{
							Description:         "WebhookClientConfig contains the information to make a TLS connection with the webhook",
							MarkdownDescription: "WebhookClientConfig contains the information to make a TLS connection with the webhook",
							Attributes: map[string]schema.Attribute{
								"ca_bundle": schema.StringAttribute{
									Description:         "'caBundle' is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
									MarkdownDescription: "'caBundle' is a PEM encoded CA bundle which will be used to validate the webhook's server certificate. If unspecified, system trust roots on the apiserver are used.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"service": schema.SingleNestedAttribute{
									Description:         "ServiceReference holds a reference to Service.legacy.k8s.io",
									MarkdownDescription: "ServiceReference holds a reference to Service.legacy.k8s.io",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "'name' is the name of the service. Required",
											MarkdownDescription: "'name' is the name of the service. Required",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"namespace": schema.StringAttribute{
											Description:         "'namespace' is the namespace of the service. Required",
											MarkdownDescription: "'namespace' is the namespace of the service. Required",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "'path' is an optional URL path which will be sent in any request to this service.",
											MarkdownDescription: "'path' is an optional URL path which will be sent in any request to this service.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"port": schema.Int64Attribute{
											Description:         "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
											MarkdownDescription: "If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. 'port' should be a valid port number (1-65535, inclusive).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"url": schema.StringAttribute{
									Description:         "'url' gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.The scheme must be 'https'; the URL must begin with 'https://'.A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
									MarkdownDescription: "'url' gives the location of the webhook, in standard URL form ('scheme://host:port/path'). Exactly one of 'url' or 'service' must be specified.The 'host' should not refer to a service running in the cluster; use the 'service' field instead. The host might be resolved via external DNS in some apiservers (e.g., 'kube-apiserver' cannot resolve in-cluster DNS as that would be a layering violation). 'host' may also be an IP address.Please note that using 'localhost' or '127.0.0.1' as a 'host' is risky unless you take great care to run this webhook on all hosts which run an apiserver which might need to make calls to this webhook. Such installs are likely to be non-portable, i.e., not easy to turn up in a new cluster.The scheme must be 'https'; the URL must begin with 'https://'.A path is optional, and if present may be any string permissible in a URL. You may use the path to pass an arbitrary string to the webhook, for example, a cluster identifier.Attempting to use a user or basic auth e.g. 'user:password@' is not allowed. Fragments ('#...') and query parameters ('?...') are not allowed, either.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
							Required: false,
							Optional: false,
							Computed: true,
						},

						"failure_policy": schema.StringAttribute{
							Description:         "FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Fail.",
							MarkdownDescription: "FailurePolicy defines how unrecognized errors from the admission endpoint are handled - allowed values are Ignore or Fail. Defaults to Fail.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"match_conditions": schema.ListNestedAttribute{
							Description:         "MatchConditions is a list of conditions that must be met for a request to be sent to this webhook. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.The exact matching logic is (in order):  1. If ANY matchCondition evaluates to FALSE, the webhook is skipped.  2. If ALL matchConditions evaluate to TRUE, the webhook is called.  3. If any matchCondition evaluates to an error (but none are FALSE):     - If failurePolicy=Fail, reject the request     - If failurePolicy=Ignore, the error is ignored and the webhook is skippedThis is a beta feature and managed by the AdmissionWebhookMatchConditions feature gate.",
							MarkdownDescription: "MatchConditions is a list of conditions that must be met for a request to be sent to this webhook. Match conditions filter requests that have already been matched by the rules, namespaceSelector, and objectSelector. An empty list of matchConditions matches all requests. There are a maximum of 64 match conditions allowed.The exact matching logic is (in order):  1. If ANY matchCondition evaluates to FALSE, the webhook is skipped.  2. If ALL matchConditions evaluate to TRUE, the webhook is called.  3. If any matchCondition evaluates to an error (but none are FALSE):     - If failurePolicy=Fail, reject the request     - If failurePolicy=Ignore, the error is ignored and the webhook is skippedThis is a beta feature and managed by the AdmissionWebhookMatchConditions feature gate.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"expression": schema.StringAttribute{
										Description:         "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/Required.",
										MarkdownDescription: "Expression represents the expression which will be evaluated by CEL. Must evaluate to bool. CEL expressions have access to the contents of the AdmissionRequest and Authorizer, organized into CEL variables:'object' - The object from the incoming request. The value is null for DELETE requests. 'oldObject' - The existing object. The value is null for CREATE requests. 'request' - Attributes of the admission request(/pkg/apis/admission/types.go#AdmissionRequest). 'authorizer' - A CEL Authorizer. May be used to perform authorization checks for the principal (user or service account) of the request.  See https://pkg.go.dev/k8s.io/apiserver/pkg/cel/library#Authz'authorizer.requestResource' - A CEL ResourceCheck constructed from the 'authorizer' and configured with the  request resource.Documentation on CEL: https://kubernetes.io/docs/reference/using-api/cel/Required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character (e.g. 'MyName',  or 'my.name',  or '123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with an optional DNS subdomain prefix and '/' (e.g. 'example.com/MyName')Required.",
										MarkdownDescription: "Name is an identifier for this match condition, used for strategic merging of MatchConditions, as well as providing an identifier for logging purposes. A good name should be descriptive of the associated expression. Name must be a qualified name consisting of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character (e.g. 'MyName',  or 'my.name',  or '123-abc', regex used for validation is '([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]') with an optional DNS subdomain prefix and '/' (e.g. 'example.com/MyName')Required.",
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

						"match_policy": schema.StringAttribute{
							Description:         "matchPolicy defines how the 'rules' list is used to match incoming requests. Allowed values are 'Exact' or 'Equivalent'.- Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.- Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.Defaults to 'Equivalent'",
							MarkdownDescription: "matchPolicy defines how the 'rules' list is used to match incoming requests. Allowed values are 'Exact' or 'Equivalent'.- Exact: match a request only if it exactly matches a specified rule. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, but 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would not be sent to the webhook.- Equivalent: match a request if modifies a resource listed in rules, even via another API group or version. For example, if deployments can be modified via apps/v1, apps/v1beta1, and extensions/v1beta1, and 'rules' only included 'apiGroups:['apps'], apiVersions:['v1'], resources: ['deployments']', a request to apps/v1beta1 or extensions/v1beta1 would be converted to apps/v1 and sent to the webhook.Defaults to 'Equivalent'",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"name": schema.StringAttribute{
							Description:         "The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where 'imagepolicy' is the name of the webhook, and kubernetes.io is the name of the organization. Required.",
							MarkdownDescription: "The name of the admission webhook. Name should be fully qualified, e.g., imagepolicy.kubernetes.io, where 'imagepolicy' is the name of the webhook, and kubernetes.io is the name of the organization. Required.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"namespace_selector": schema.SingleNestedAttribute{
							Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
							MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

						"object_selector": schema.SingleNestedAttribute{
							Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
							MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

						"rules": schema.ListNestedAttribute{
							Description:         "Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.",
							MarkdownDescription: "Rules describes what operations on what resources/subresources the webhook cares about. The webhook cares about an operation if it matches _any_ Rule. However, in order to prevent ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks from putting the cluster in a state which cannot be recovered from without completely disabling the plugin, ValidatingAdmissionWebhooks and MutatingAdmissionWebhooks are never called on admission requests for ValidatingWebhookConfiguration and MutatingWebhookConfiguration objects.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"api_groups": schema.ListAttribute{
										Description:         "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
										MarkdownDescription: "APIGroups is the API groups the resources belong to. '*' is all groups. If '*' is present, the length of the slice must be one. Required.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"api_versions": schema.ListAttribute{
										Description:         "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
										MarkdownDescription: "APIVersions is the API versions the resources belong to. '*' is all versions. If '*' is present, the length of the slice must be one. Required.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"operations": schema.ListAttribute{
										Description:         "Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.",
										MarkdownDescription: "Operations is the operations the admission hook cares about - CREATE, UPDATE, DELETE, CONNECT or * for all of those operations and any future admission operations that are added. If '*' is present, the length of the slice must be one. Required.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"resources": schema.ListAttribute{
										Description:         "Resources is a list of resources this rule applies to.For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources.If wildcard is present, the validation rule will ensure resources do not overlap with each other.Depending on the enclosing object, subresources might not be allowed. Required.",
										MarkdownDescription: "Resources is a list of resources this rule applies to.For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources.If wildcard is present, the validation rule will ensure resources do not overlap with each other.Depending on the enclosing object, subresources might not be allowed. Required.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"scope": schema.StringAttribute{
										Description:         "scope specifies the scope of this rule. Valid values are 'Cluster', 'Namespaced', and '*' 'Cluster' means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. 'Namespaced' means that only namespaced resources will match this rule. '*' means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is '*'.",
										MarkdownDescription: "scope specifies the scope of this rule. Valid values are 'Cluster', 'Namespaced', and '*' 'Cluster' means that only cluster-scoped resources will match this rule. Namespace API objects are cluster-scoped. 'Namespaced' means that only namespaced resources will match this rule. '*' means that there are no scope restrictions. Subresources match the scope of their parent resource. Default is '*'.",
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

						"side_effects": schema.StringAttribute{
							Description:         "SideEffects states whether this webhook has side effects. Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission chain and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some.",
							MarkdownDescription: "SideEffects states whether this webhook has side effects. Acceptable values are: None, NoneOnDryRun (webhooks created via v1beta1 may also specify Some or Unknown). Webhooks with side effects MUST implement a reconciliation system, since a request may be rejected by a future step in the admission chain and the side effects therefore need to be undone. Requests with the dryRun attribute will be auto-rejected if they match a webhook with sideEffects == Unknown or Some.",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"timeout_seconds": schema.Int64Attribute{
							Description:         "TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.",
							MarkdownDescription: "TimeoutSeconds specifies the timeout for this webhook. After the timeout passes, the webhook call will be ignored or the API call will fail based on the failure policy. The timeout value must be between 1 and 30 seconds. Default to 10 seconds.",
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
		},
	}
}

func (r *AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_admissionregistration_k8s_io_validating_webhook_configuration_v1")

	var data AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "admissionregistration.k8s.io", Version: "v1", Resource: "ValidatingWebhookConfiguration"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse AdmissionregistrationK8SIoValidatingWebhookConfigurationV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("admissionregistration.k8s.io/v1")
	data.Kind = pointer.String("ValidatingWebhookConfiguration")
	data.Metadata = readResponse.Metadata
	data.Webhooks = readResponse.Webhooks

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
