/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"
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

type CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource struct{}

var (
	_ resource.Resource = (*CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource)(nil)
)

type CrdProjectcalicoOrgGlobalNetworkPolicyV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CrdProjectcalicoOrgGlobalNetworkPolicyV1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		DoNotTrack *bool `tfsdk:"do_not_track" yaml:"doNotTrack,omitempty"`

		NamespaceSelector *string `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

		Order *float64 `tfsdk:"order" yaml:"order,omitempty"`

		Types *[]string `tfsdk:"types" yaml:"types,omitempty"`

		ApplyOnForward *bool `tfsdk:"apply_on_forward" yaml:"applyOnForward,omitempty"`

		Egress *[]struct {
			Http *struct {
				Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

				Paths *[]struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
				} `tfsdk:"paths" yaml:"paths,omitempty"`
			} `tfsdk:"http" yaml:"http,omitempty"`

			Icmp *struct {
				Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

				Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"icmp" yaml:"icmp,omitempty"`

			IpVersion *int64 `tfsdk:"ip_version" yaml:"ipVersion,omitempty"`

			NotProtocol *string `tfsdk:"not_protocol" yaml:"notProtocol,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

			Source *struct {
				Services *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"services" yaml:"services,omitempty"`

				NotNets *[]string `tfsdk:"not_nets" yaml:"notNets,omitempty"`

				NotPorts *[]string `tfsdk:"not_ports" yaml:"notPorts,omitempty"`

				Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`

				ServiceAccounts *struct {
					Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

					Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"service_accounts" yaml:"serviceAccounts,omitempty"`

				NamespaceSelector *string `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

				Nets *[]string `tfsdk:"nets" yaml:"nets,omitempty"`

				NotSelector *string `tfsdk:"not_selector" yaml:"notSelector,omitempty"`

				Ports *[]string `tfsdk:"ports" yaml:"ports,omitempty"`
			} `tfsdk:"source" yaml:"source,omitempty"`

			Action *string `tfsdk:"action" yaml:"action,omitempty"`

			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`
			} `tfsdk:"metadata" yaml:"metadata,omitempty"`

			NotICMP *struct {
				Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

				Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"not_icmp" yaml:"notICMP,omitempty"`

			Destination *struct {
				NotSelector *string `tfsdk:"not_selector" yaml:"notSelector,omitempty"`

				ServiceAccounts *struct {
					Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

					Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"service_accounts" yaml:"serviceAccounts,omitempty"`

				Nets *[]string `tfsdk:"nets" yaml:"nets,omitempty"`

				NotPorts *[]string `tfsdk:"not_ports" yaml:"notPorts,omitempty"`

				Ports *[]string `tfsdk:"ports" yaml:"ports,omitempty"`

				Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`

				Services *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"services" yaml:"services,omitempty"`

				NamespaceSelector *string `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

				NotNets *[]string `tfsdk:"not_nets" yaml:"notNets,omitempty"`
			} `tfsdk:"destination" yaml:"destination,omitempty"`
		} `tfsdk:"egress" yaml:"egress,omitempty"`

		Ingress *[]struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`
			} `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Source *struct {
				NotNets *[]string `tfsdk:"not_nets" yaml:"notNets,omitempty"`

				Ports *[]string `tfsdk:"ports" yaml:"ports,omitempty"`

				ServiceAccounts *struct {
					Names *[]string `tfsdk:"names" yaml:"names,omitempty"`

					Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"service_accounts" yaml:"serviceAccounts,omitempty"`

				Services *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"services" yaml:"services,omitempty"`

				NamespaceSelector *string `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

				Nets *[]string `tfsdk:"nets" yaml:"nets,omitempty"`

				Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`

				NotPorts *[]string `tfsdk:"not_ports" yaml:"notPorts,omitempty"`

				NotSelector *string `tfsdk:"not_selector" yaml:"notSelector,omitempty"`
			} `tfsdk:"source" yaml:"source,omitempty"`

			Destination *struct {
				NotPorts *[]string `tfsdk:"not_ports" yaml:"notPorts,omitempty"`

				Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`

				ServiceAccounts *struct {
					Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`

					Names *[]string `tfsdk:"names" yaml:"names,omitempty"`
				} `tfsdk:"service_accounts" yaml:"serviceAccounts,omitempty"`

				Services *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"services" yaml:"services,omitempty"`

				NamespaceSelector *string `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

				Nets *[]string `tfsdk:"nets" yaml:"nets,omitempty"`

				Ports *[]string `tfsdk:"ports" yaml:"ports,omitempty"`

				NotNets *[]string `tfsdk:"not_nets" yaml:"notNets,omitempty"`

				NotSelector *string `tfsdk:"not_selector" yaml:"notSelector,omitempty"`
			} `tfsdk:"destination" yaml:"destination,omitempty"`

			Icmp *struct {
				Type *int64 `tfsdk:"type" yaml:"type,omitempty"`

				Code *int64 `tfsdk:"code" yaml:"code,omitempty"`
			} `tfsdk:"icmp" yaml:"icmp,omitempty"`

			IpVersion *int64 `tfsdk:"ip_version" yaml:"ipVersion,omitempty"`

			NotICMP *struct {
				Code *int64 `tfsdk:"code" yaml:"code,omitempty"`

				Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"not_icmp" yaml:"notICMP,omitempty"`

			NotProtocol *string `tfsdk:"not_protocol" yaml:"notProtocol,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

			Action *string `tfsdk:"action" yaml:"action,omitempty"`

			Http *struct {
				Methods *[]string `tfsdk:"methods" yaml:"methods,omitempty"`

				Paths *[]struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
				} `tfsdk:"paths" yaml:"paths,omitempty"`
			} `tfsdk:"http" yaml:"http,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		PreDNAT *bool `tfsdk:"pre_dnat" yaml:"preDNAT,omitempty"`

		Selector *string `tfsdk:"selector" yaml:"selector,omitempty"`

		ServiceAccountSelector *string `tfsdk:"service_account_selector" yaml:"serviceAccountSelector,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCrdProjectcalicoOrgGlobalNetworkPolicyV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource{}
}

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crd_projectcalico_org_global_network_policy_v1"
}

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"do_not_track": {
						Description:         "DoNotTrack indicates whether packets matched by the rules in this policy should go through the data plane's connection tracking, such as Linux conntrack.  If True, the rules in this policy are applied before any data plane connection tracking, and packets allowed by this policy are marked as not to be tracked.",
						MarkdownDescription: "DoNotTrack indicates whether packets matched by the rules in this policy should go through the data plane's connection tracking, such as Linux conntrack.  If True, the rules in this policy are applied before any data plane connection tracking, and packets allowed by this policy are marked as not to be tracked.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace_selector": {
						Description:         "NamespaceSelector is an optional field for an expression used to select a pod based on namespaces.",
						MarkdownDescription: "NamespaceSelector is an optional field for an expression used to select a pod based on namespaces.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"order": {
						Description:         "Order is an optional field that specifies the order in which the policy is applied. Policies with higher 'order' are applied after those with lower order.  If the order is omitted, it may be considered to be 'infinite' - i.e. the policy will be applied last.  Policies with identical order will be applied in alphanumerical order based on the Policy 'Name'.",
						MarkdownDescription: "Order is an optional field that specifies the order in which the policy is applied. Policies with higher 'order' are applied after those with lower order.  If the order is omitted, it may be considered to be 'infinite' - i.e. the policy will be applied last.  Policies with identical order will be applied in alphanumerical order based on the Policy 'Name'.",

						Type: types.NumberType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"types": {
						Description:         "Types indicates whether this policy applies to ingress, or to egress, or to both.  When not explicitly specified (and so the value on creation is empty or nil), Calico defaults Types according to what Ingress and Egress rules are present in the policy.  The default is:  - [ PolicyTypeIngress ], if there are no Egress rules (including the case where there are   also no Ingress rules)  - [ PolicyTypeEgress ], if there are Egress rules but no Ingress rules  - [ PolicyTypeIngress, PolicyTypeEgress ], if there are both Ingress and Egress rules.  When the policy is read back again, Types will always be one of these values, never empty or nil.",
						MarkdownDescription: "Types indicates whether this policy applies to ingress, or to egress, or to both.  When not explicitly specified (and so the value on creation is empty or nil), Calico defaults Types according to what Ingress and Egress rules are present in the policy.  The default is:  - [ PolicyTypeIngress ], if there are no Egress rules (including the case where there are   also no Ingress rules)  - [ PolicyTypeEgress ], if there are Egress rules but no Ingress rules  - [ PolicyTypeIngress, PolicyTypeEgress ], if there are both Ingress and Egress rules.  When the policy is read back again, Types will always be one of these values, never empty or nil.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"apply_on_forward": {
						Description:         "ApplyOnForward indicates to apply the rules in this policy on forward traffic.",
						MarkdownDescription: "ApplyOnForward indicates to apply the rules in this policy on forward traffic.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress": {
						Description:         "The ordered set of egress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",
						MarkdownDescription: "The ordered set of egress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"http": {
								Description:         "HTTP contains match criteria that apply to HTTP requests.",
								MarkdownDescription: "HTTP contains match criteria that apply to HTTP requests.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"methods": {
										Description:         "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",
										MarkdownDescription: "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"paths": {
										Description:         "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",
										MarkdownDescription: "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"icmp": {
								Description:         "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",
								MarkdownDescription: "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"code": {
										Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
										MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
										MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",

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

							"ip_version": {
								Description:         "IPVersion is an optional field that restricts the rule to only match a specific IP version.",
								MarkdownDescription: "IPVersion is an optional field that restricts the rule to only match a specific IP version.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"not_protocol": {
								Description:         "NotProtocol is the negated version of the Protocol field.",
								MarkdownDescription: "NotProtocol is the negated version of the Protocol field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"protocol": {
								Description:         "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",
								MarkdownDescription: "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source": {
								Description:         "Source contains the match criteria that apply to source entity.",
								MarkdownDescription: "Source contains the match criteria that apply to source entity.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"services": {
										Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
										MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name specifies the name of a Kubernetes Service to match.",
												MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
												MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",

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

									"not_nets": {
										Description:         "NotNets is the negated version of the Nets field.",
										MarkdownDescription: "NotNets is the negated version of the Nets field.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_ports": {
										Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
										MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_accounts": {
										Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
										MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"names": {
												Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
												MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
												MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",

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

									"namespace_selector": {
										Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
										MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"nets": {
										Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
										MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_selector": {
										Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
										MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

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

							"action": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"metadata": {
								Description:         "Metadata contains additional information for this rule",
								MarkdownDescription: "Metadata contains additional information for this rule",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is a set of key value pairs that give extra information about the rule",
										MarkdownDescription: "Annotations is a set of key value pairs that give extra information about the rule",

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

							"not_icmp": {
								Description:         "NotICMP is the negated version of the ICMP field.",
								MarkdownDescription: "NotICMP is the negated version of the ICMP field.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"code": {
										Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
										MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
										MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",

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

							"destination": {
								Description:         "Destination contains the match criteria that apply to destination entity.",
								MarkdownDescription: "Destination contains the match criteria that apply to destination entity.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"not_selector": {
										Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
										MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_accounts": {
										Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
										MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"names": {
												Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
												MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
												MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",

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

									"nets": {
										Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
										MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_ports": {
										Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
										MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"services": {
										Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
										MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name specifies the name of a Kubernetes Service to match.",
												MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
												MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",

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

									"namespace_selector": {
										Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
										MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_nets": {
										Description:         "NotNets is the negated version of the Nets field.",
										MarkdownDescription: "NotNets is the negated version of the Nets field.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": {
						Description:         "The ordered set of ingress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",
						MarkdownDescription: "The ordered set of ingress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"metadata": {
								Description:         "Metadata contains additional information for this rule",
								MarkdownDescription: "Metadata contains additional information for this rule",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is a set of key value pairs that give extra information about the rule",
										MarkdownDescription: "Annotations is a set of key value pairs that give extra information about the rule",

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

							"source": {
								Description:         "Source contains the match criteria that apply to source entity.",
								MarkdownDescription: "Source contains the match criteria that apply to source entity.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"not_nets": {
										Description:         "NotNets is the negated version of the Nets field.",
										MarkdownDescription: "NotNets is the negated version of the Nets field.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_accounts": {
										Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
										MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"names": {
												Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
												MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
												MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",

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

									"services": {
										Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
										MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
												MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name specifies the name of a Kubernetes Service to match.",
												MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",

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

									"namespace_selector": {
										Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
										MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"nets": {
										Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
										MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
										MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_ports": {
										Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_selector": {
										Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
										MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",

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

							"destination": {
								Description:         "Destination contains the match criteria that apply to destination entity.",
								MarkdownDescription: "Destination contains the match criteria that apply to destination entity.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"not_ports": {
										Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selector": {
										Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
										MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_accounts": {
										Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
										MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"selector": {
												Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
												MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"names": {
												Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
												MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",

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

									"services": {
										Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
										MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name specifies the name of a Kubernetes Service to match.",
												MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
												MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",

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

									"namespace_selector": {
										Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
										MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"nets": {
										Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
										MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
										MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_nets": {
										Description:         "NotNets is the negated version of the Nets field.",
										MarkdownDescription: "NotNets is the negated version of the Nets field.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"not_selector": {
										Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
										MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",

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

							"icmp": {
								Description:         "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",
								MarkdownDescription: "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"type": {
										Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
										MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"code": {
										Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
										MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",

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

							"ip_version": {
								Description:         "IPVersion is an optional field that restricts the rule to only match a specific IP version.",
								MarkdownDescription: "IPVersion is an optional field that restricts the rule to only match a specific IP version.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"not_icmp": {
								Description:         "NotICMP is the negated version of the ICMP field.",
								MarkdownDescription: "NotICMP is the negated version of the ICMP field.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"code": {
										Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
										MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
										MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",

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

							"not_protocol": {
								Description:         "NotProtocol is the negated version of the Protocol field.",
								MarkdownDescription: "NotProtocol is the negated version of the Protocol field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"protocol": {
								Description:         "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",
								MarkdownDescription: "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"action": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"http": {
								Description:         "HTTP contains match criteria that apply to HTTP requests.",
								MarkdownDescription: "HTTP contains match criteria that apply to HTTP requests.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"methods": {
										Description:         "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",
										MarkdownDescription: "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"paths": {
										Description:         "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",
										MarkdownDescription: "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
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

					"pre_dnat": {
						Description:         "PreDNAT indicates to apply the rules in this policy before any DNAT.",
						MarkdownDescription: "PreDNAT indicates to apply the rules in this policy before any DNAT.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "The selector is an expression used to pick pick out the endpoints that the policy should be applied to.  Selector expressions follow this syntax:  	label == 'string_literal'  ->  comparison, e.g. my_label == 'foo bar' 	label != 'string_literal'   ->  not equal; also matches if label is not present 	label in { 'a', 'b', 'c', ... }  ->  true if the value of label X is one of 'a', 'b', 'c' 	label not in { 'a', 'b', 'c', ... }  ->  true if the value of label X is not one of 'a', 'b', 'c' 	has(label_name)  -> True if that label is present 	! expr -> negation of expr 	expr && expr  -> Short-circuit and 	expr || expr  -> Short-circuit or 	( expr ) -> parens for grouping 	all() or the empty selector -> matches all endpoints.  Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive but they do not support escape characters.  Examples (with made-up labels):  	type == 'webserver' && deployment == 'prod' 	type in {'frontend', 'backend'} 	deployment != 'dev' 	! has(label_name)",
						MarkdownDescription: "The selector is an expression used to pick pick out the endpoints that the policy should be applied to.  Selector expressions follow this syntax:  	label == 'string_literal'  ->  comparison, e.g. my_label == 'foo bar' 	label != 'string_literal'   ->  not equal; also matches if label is not present 	label in { 'a', 'b', 'c', ... }  ->  true if the value of label X is one of 'a', 'b', 'c' 	label not in { 'a', 'b', 'c', ... }  ->  true if the value of label X is not one of 'a', 'b', 'c' 	has(label_name)  -> True if that label is present 	! expr -> negation of expr 	expr && expr  -> Short-circuit and 	expr || expr  -> Short-circuit or 	( expr ) -> parens for grouping 	all() or the empty selector -> matches all endpoints.  Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive but they do not support escape characters.  Examples (with made-up labels):  	type == 'webserver' && deployment == 'prod' 	type in {'frontend', 'backend'} 	deployment != 'dev' 	! has(label_name)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_selector": {
						Description:         "ServiceAccountSelector is an optional field for an expression used to select a pod based on service accounts.",
						MarkdownDescription: "ServiceAccountSelector is an optional field for an expression used to select a pod based on service accounts.",

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
		},
	}, nil
}

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_global_network_policy_v1")

	var state CrdProjectcalicoOrgGlobalNetworkPolicyV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgGlobalNetworkPolicyV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("GlobalNetworkPolicy")

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

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_global_network_policy_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_global_network_policy_v1")

	var state CrdProjectcalicoOrgGlobalNetworkPolicyV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgGlobalNetworkPolicyV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("GlobalNetworkPolicy")

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

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_global_network_policy_v1")
	// NO-OP: Terraform removes the state automatically for us
}
