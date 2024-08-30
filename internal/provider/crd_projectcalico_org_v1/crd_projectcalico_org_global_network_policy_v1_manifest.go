/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	_ datasource.DataSource = &CrdProjectcalicoOrgGlobalNetworkPolicyV1Manifest{}
)

func NewCrdProjectcalicoOrgGlobalNetworkPolicyV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgGlobalNetworkPolicyV1Manifest{}
}

type CrdProjectcalicoOrgGlobalNetworkPolicyV1Manifest struct{}

type CrdProjectcalicoOrgGlobalNetworkPolicyV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ApplyOnForward *bool `tfsdk:"apply_on_forward" json:"applyOnForward,omitempty"`
		DoNotTrack     *bool `tfsdk:"do_not_track" json:"doNotTrack,omitempty"`
		Egress         *[]struct {
			Action      *string `tfsdk:"action" json:"action,omitempty"`
			Destination *struct {
				NamespaceSelector *string   `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				Nets              *[]string `tfsdk:"nets" json:"nets,omitempty"`
				NotNets           *[]string `tfsdk:"not_nets" json:"notNets,omitempty"`
				NotPorts          *[]string `tfsdk:"not_ports" json:"notPorts,omitempty"`
				NotSelector       *string   `tfsdk:"not_selector" json:"notSelector,omitempty"`
				Ports             *[]string `tfsdk:"ports" json:"ports,omitempty"`
				Selector          *string   `tfsdk:"selector" json:"selector,omitempty"`
				ServiceAccounts   *struct {
					Names    *[]string `tfsdk:"names" json:"names,omitempty"`
					Selector *string   `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"service_accounts" json:"serviceAccounts,omitempty"`
				Services *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"services" json:"services,omitempty"`
			} `tfsdk:"destination" json:"destination,omitempty"`
			Http *struct {
				Methods *[]string `tfsdk:"methods" json:"methods,omitempty"`
				Paths   *[]struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
				} `tfsdk:"paths" json:"paths,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			Icmp *struct {
				Code *int64 `tfsdk:"code" json:"code,omitempty"`
				Type *int64 `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"icmp" json:"icmp,omitempty"`
			IpVersion *int64 `tfsdk:"ip_version" json:"ipVersion,omitempty"`
			Metadata  *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			NotICMP *struct {
				Code *int64 `tfsdk:"code" json:"code,omitempty"`
				Type *int64 `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"not_icmp" json:"notICMP,omitempty"`
			NotProtocol *string `tfsdk:"not_protocol" json:"notProtocol,omitempty"`
			Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
			Source      *struct {
				NamespaceSelector *string   `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				Nets              *[]string `tfsdk:"nets" json:"nets,omitempty"`
				NotNets           *[]string `tfsdk:"not_nets" json:"notNets,omitempty"`
				NotPorts          *[]string `tfsdk:"not_ports" json:"notPorts,omitempty"`
				NotSelector       *string   `tfsdk:"not_selector" json:"notSelector,omitempty"`
				Ports             *[]string `tfsdk:"ports" json:"ports,omitempty"`
				Selector          *string   `tfsdk:"selector" json:"selector,omitempty"`
				ServiceAccounts   *struct {
					Names    *[]string `tfsdk:"names" json:"names,omitempty"`
					Selector *string   `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"service_accounts" json:"serviceAccounts,omitempty"`
				Services *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"services" json:"services,omitempty"`
			} `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		Ingress *[]struct {
			Action      *string `tfsdk:"action" json:"action,omitempty"`
			Destination *struct {
				NamespaceSelector *string   `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				Nets              *[]string `tfsdk:"nets" json:"nets,omitempty"`
				NotNets           *[]string `tfsdk:"not_nets" json:"notNets,omitempty"`
				NotPorts          *[]string `tfsdk:"not_ports" json:"notPorts,omitempty"`
				NotSelector       *string   `tfsdk:"not_selector" json:"notSelector,omitempty"`
				Ports             *[]string `tfsdk:"ports" json:"ports,omitempty"`
				Selector          *string   `tfsdk:"selector" json:"selector,omitempty"`
				ServiceAccounts   *struct {
					Names    *[]string `tfsdk:"names" json:"names,omitempty"`
					Selector *string   `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"service_accounts" json:"serviceAccounts,omitempty"`
				Services *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"services" json:"services,omitempty"`
			} `tfsdk:"destination" json:"destination,omitempty"`
			Http *struct {
				Methods *[]string `tfsdk:"methods" json:"methods,omitempty"`
				Paths   *[]struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
				} `tfsdk:"paths" json:"paths,omitempty"`
			} `tfsdk:"http" json:"http,omitempty"`
			Icmp *struct {
				Code *int64 `tfsdk:"code" json:"code,omitempty"`
				Type *int64 `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"icmp" json:"icmp,omitempty"`
			IpVersion *int64 `tfsdk:"ip_version" json:"ipVersion,omitempty"`
			Metadata  *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			NotICMP *struct {
				Code *int64 `tfsdk:"code" json:"code,omitempty"`
				Type *int64 `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"not_icmp" json:"notICMP,omitempty"`
			NotProtocol *string `tfsdk:"not_protocol" json:"notProtocol,omitempty"`
			Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
			Source      *struct {
				NamespaceSelector *string   `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				Nets              *[]string `tfsdk:"nets" json:"nets,omitempty"`
				NotNets           *[]string `tfsdk:"not_nets" json:"notNets,omitempty"`
				NotPorts          *[]string `tfsdk:"not_ports" json:"notPorts,omitempty"`
				NotSelector       *string   `tfsdk:"not_selector" json:"notSelector,omitempty"`
				Ports             *[]string `tfsdk:"ports" json:"ports,omitempty"`
				Selector          *string   `tfsdk:"selector" json:"selector,omitempty"`
				ServiceAccounts   *struct {
					Names    *[]string `tfsdk:"names" json:"names,omitempty"`
					Selector *string   `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"service_accounts" json:"serviceAccounts,omitempty"`
				Services *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"services" json:"services,omitempty"`
			} `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		NamespaceSelector      *string   `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		Order                  *float64  `tfsdk:"order" json:"order,omitempty"`
		PerformanceHints       *[]string `tfsdk:"performance_hints" json:"performanceHints,omitempty"`
		PreDNAT                *bool     `tfsdk:"pre_dnat" json:"preDNAT,omitempty"`
		Selector               *string   `tfsdk:"selector" json:"selector,omitempty"`
		ServiceAccountSelector *string   `tfsdk:"service_account_selector" json:"serviceAccountSelector,omitempty"`
		Tier                   *string   `tfsdk:"tier" json:"tier,omitempty"`
		Types                  *[]string `tfsdk:"types" json:"types,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_global_network_policy_v1_manifest"
}

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"apply_on_forward": schema.BoolAttribute{
						Description:         "ApplyOnForward indicates to apply the rules in this policy on forward traffic.",
						MarkdownDescription: "ApplyOnForward indicates to apply the rules in this policy on forward traffic.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"do_not_track": schema.BoolAttribute{
						Description:         "DoNotTrack indicates whether packets matched by the rules in this policy should go through the data plane's connection tracking, such as Linux conntrack.  If True, the rules in this policy are applied before any data plane connection tracking, and packets allowed by this policy are marked as not to be tracked.",
						MarkdownDescription: "DoNotTrack indicates whether packets matched by the rules in this policy should go through the data plane's connection tracking, such as Linux conntrack.  If True, the rules in this policy are applied before any data plane connection tracking, and packets allowed by this policy are marked as not to be tracked.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"egress": schema.ListNestedAttribute{
						Description:         "The ordered set of egress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",
						MarkdownDescription: "The ordered set of egress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"destination": schema.SingleNestedAttribute{
									Description:         "Destination contains the match criteria that apply to destination entity.",
									MarkdownDescription: "Destination contains the match criteria that apply to destination entity.",
									Attributes: map[string]schema.Attribute{
										"namespace_selector": schema.StringAttribute{
											Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"nets": schema.ListAttribute{
											Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_nets": schema.ListAttribute{
											Description:         "NotNets is the negated version of the Nets field.",
											MarkdownDescription: "NotNets is the negated version of the Nets field.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_ports": schema.ListAttribute{
											Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_selector": schema.StringAttribute{
											Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListAttribute{
											Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.StringAttribute{
											Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts": schema.SingleNestedAttribute{
											Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											Attributes: map[string]schema.Attribute{
												"names": schema.ListAttribute{
													Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"services": schema.SingleNestedAttribute{
											Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name specifies the name of a Kubernetes Service to match.",
													MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
													MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
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

								"http": schema.SingleNestedAttribute{
									Description:         "HTTP contains match criteria that apply to HTTP requests.",
									MarkdownDescription: "HTTP contains match criteria that apply to HTTP requests.",
									Attributes: map[string]schema.Attribute{
										"methods": schema.ListAttribute{
											Description:         "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",
											MarkdownDescription: "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"paths": schema.ListNestedAttribute{
											Description:         "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",
											MarkdownDescription: "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"prefix": schema.StringAttribute{
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

								"icmp": schema.SingleNestedAttribute{
									Description:         "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",
									MarkdownDescription: "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",
									Attributes: map[string]schema.Attribute{
										"code": schema.Int64Attribute{
											Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.Int64Attribute{
											Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ip_version": schema.Int64Attribute{
									Description:         "IPVersion is an optional field that restricts the rule to only match a specific IP version.",
									MarkdownDescription: "IPVersion is an optional field that restricts the rule to only match a specific IP version.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metadata": schema.SingleNestedAttribute{
									Description:         "Metadata contains additional information for this rule",
									MarkdownDescription: "Metadata contains additional information for this rule",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "Annotations is a set of key value pairs that give extra information about the rule",
											MarkdownDescription: "Annotations is a set of key value pairs that give extra information about the rule",
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

								"not_icmp": schema.SingleNestedAttribute{
									Description:         "NotICMP is the negated version of the ICMP field.",
									MarkdownDescription: "NotICMP is the negated version of the ICMP field.",
									Attributes: map[string]schema.Attribute{
										"code": schema.Int64Attribute{
											Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.Int64Attribute{
											Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"not_protocol": schema.StringAttribute{
									Description:         "NotProtocol is the negated version of the Protocol field.",
									MarkdownDescription: "NotProtocol is the negated version of the Protocol field.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
									Description:         "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",
									MarkdownDescription: "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source": schema.SingleNestedAttribute{
									Description:         "Source contains the match criteria that apply to source entity.",
									MarkdownDescription: "Source contains the match criteria that apply to source entity.",
									Attributes: map[string]schema.Attribute{
										"namespace_selector": schema.StringAttribute{
											Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"nets": schema.ListAttribute{
											Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_nets": schema.ListAttribute{
											Description:         "NotNets is the negated version of the Nets field.",
											MarkdownDescription: "NotNets is the negated version of the Nets field.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_ports": schema.ListAttribute{
											Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_selector": schema.StringAttribute{
											Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListAttribute{
											Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.StringAttribute{
											Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts": schema.SingleNestedAttribute{
											Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											Attributes: map[string]schema.Attribute{
												"names": schema.ListAttribute{
													Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"services": schema.SingleNestedAttribute{
											Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name specifies the name of a Kubernetes Service to match.",
													MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
													MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": schema.ListNestedAttribute{
						Description:         "The ordered set of ingress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",
						MarkdownDescription: "The ordered set of ingress rules.  Each rule contains a set of packet match criteria and a corresponding action to apply.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"destination": schema.SingleNestedAttribute{
									Description:         "Destination contains the match criteria that apply to destination entity.",
									MarkdownDescription: "Destination contains the match criteria that apply to destination entity.",
									Attributes: map[string]schema.Attribute{
										"namespace_selector": schema.StringAttribute{
											Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"nets": schema.ListAttribute{
											Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_nets": schema.ListAttribute{
											Description:         "NotNets is the negated version of the Nets field.",
											MarkdownDescription: "NotNets is the negated version of the Nets field.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_ports": schema.ListAttribute{
											Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_selector": schema.StringAttribute{
											Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListAttribute{
											Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.StringAttribute{
											Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts": schema.SingleNestedAttribute{
											Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											Attributes: map[string]schema.Attribute{
												"names": schema.ListAttribute{
													Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"services": schema.SingleNestedAttribute{
											Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name specifies the name of a Kubernetes Service to match.",
													MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
													MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
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

								"http": schema.SingleNestedAttribute{
									Description:         "HTTP contains match criteria that apply to HTTP requests.",
									MarkdownDescription: "HTTP contains match criteria that apply to HTTP requests.",
									Attributes: map[string]schema.Attribute{
										"methods": schema.ListAttribute{
											Description:         "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",
											MarkdownDescription: "Methods is an optional field that restricts the rule to apply only to HTTP requests that use one of the listed HTTP Methods (e.g. GET, PUT, etc.) Multiple methods are OR'd together.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"paths": schema.ListNestedAttribute{
											Description:         "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",
											MarkdownDescription: "Paths is an optional field that restricts the rule to apply to HTTP requests that use one of the listed HTTP Paths. Multiple paths are OR'd together. e.g: - exact: /foo - prefix: /bar NOTE: Each entry may ONLY specify either a 'exact' or a 'prefix' match. The validator will check for it.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"exact": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"prefix": schema.StringAttribute{
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

								"icmp": schema.SingleNestedAttribute{
									Description:         "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",
									MarkdownDescription: "ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.  This should only be specified if the Protocol field is set to 'ICMP' or 'ICMPv6'.",
									Attributes: map[string]schema.Attribute{
										"code": schema.Int64Attribute{
											Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.Int64Attribute{
											Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ip_version": schema.Int64Attribute{
									Description:         "IPVersion is an optional field that restricts the rule to only match a specific IP version.",
									MarkdownDescription: "IPVersion is an optional field that restricts the rule to only match a specific IP version.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"metadata": schema.SingleNestedAttribute{
									Description:         "Metadata contains additional information for this rule",
									MarkdownDescription: "Metadata contains additional information for this rule",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "Annotations is a set of key value pairs that give extra information about the rule",
											MarkdownDescription: "Annotations is a set of key value pairs that give extra information about the rule",
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

								"not_icmp": schema.SingleNestedAttribute{
									Description:         "NotICMP is the negated version of the ICMP field.",
									MarkdownDescription: "NotICMP is the negated version of the ICMP field.",
									Attributes: map[string]schema.Attribute{
										"code": schema.Int64Attribute{
											Description:         "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											MarkdownDescription: "Match on a specific ICMP code.  If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.Int64Attribute{
											Description:         "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											MarkdownDescription: "Match on a specific ICMP type.  For example a value of 8 refers to ICMP Echo Request (i.e. pings).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"not_protocol": schema.StringAttribute{
									Description:         "NotProtocol is the negated version of the Protocol field.",
									MarkdownDescription: "NotProtocol is the negated version of the Protocol field.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
									Description:         "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",
									MarkdownDescription: "Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol. Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).  Must be one of these string values: 'TCP', 'UDP', 'ICMP', 'ICMPv6', 'SCTP', 'UDPLite' or an integer in the range 1-255.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source": schema.SingleNestedAttribute{
									Description:         "Source contains the match criteria that apply to source entity.",
									MarkdownDescription: "Source contains the match criteria that apply to source entity.",
									Attributes: map[string]schema.Attribute{
										"namespace_selector": schema.StringAttribute{
											Description:         "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											MarkdownDescription: "NamespaceSelector is an optional field that contains a selector expression. Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.  For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.  For NetworkPolicy, 'global()' NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.  For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"nets": schema.ListAttribute{
											Description:         "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											MarkdownDescription: "Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_nets": schema.ListAttribute{
											Description:         "NotNets is the negated version of the Nets field.",
											MarkdownDescription: "NotNets is the negated version of the Nets field.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_ports": schema.ListAttribute{
											Description:         "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "NotPorts is the negated version of the Ports field. Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"not_selector": schema.StringAttribute{
											Description:         "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											MarkdownDescription: "NotSelector is the negated version of the Selector field.  See Selector field for subtleties with negated selectors.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListAttribute{
											Description:         "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											MarkdownDescription: "Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values. This value is a list of integers or strings that represent ranges of ports.  Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to 'TCP' or 'UDP'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"selector": schema.StringAttribute{
											Description:         "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											MarkdownDescription: "Selector is an optional field that contains a selector expression (see Policy for sample syntax).  Only traffic that originates from (terminates at) endpoints matching the selector will be matched.  Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:  	Selector = '!has(my_label)' matches packets that are from other Calico-controlled 	endpoints that do not have the label 'my_label'.  	NotSelector = 'has(my_label)' matches packets that are not from Calico-controlled 	endpoints that do have the label 'my_label'.  The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_accounts": schema.SingleNestedAttribute{
											Description:         "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											MarkdownDescription: "ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.",
											Attributes: map[string]schema.Attribute{
												"names": schema.ListAttribute{
													Description:         "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													MarkdownDescription: "Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													MarkdownDescription: "Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector. If both Names and Selector are specified then they are AND'ed.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"services": schema.SingleNestedAttribute{
											Description:         "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											MarkdownDescription: "Services is an optional field that contains options for matching Kubernetes Services. If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.  Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.  Ports and NotPorts can only be specified with Services on ingress rules.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name specifies the name of a Kubernetes Service to match.",
													MarkdownDescription: "Name specifies the name of a Kubernetes Service to match.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
													MarkdownDescription: "Namespace specifies the namespace of the given Service. If left empty, the rule will match within this policy's namespace.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace_selector": schema.StringAttribute{
						Description:         "NamespaceSelector is an optional field for an expression used to select a pod based on namespaces.",
						MarkdownDescription: "NamespaceSelector is an optional field for an expression used to select a pod based on namespaces.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"order": schema.Float64Attribute{
						Description:         "Order is an optional field that specifies the order in which the policy is applied. Policies with higher 'order' are applied after those with lower order within the same tier.  If the order is omitted, it may be considered to be 'infinite' - i.e. the policy will be applied last.  Policies with identical order will be applied in alphanumerical order based on the Policy 'Name' within the tier.",
						MarkdownDescription: "Order is an optional field that specifies the order in which the policy is applied. Policies with higher 'order' are applied after those with lower order within the same tier.  If the order is omitted, it may be considered to be 'infinite' - i.e. the policy will be applied last.  Policies with identical order will be applied in alphanumerical order based on the Policy 'Name' within the tier.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_hints": schema.ListAttribute{
						Description:         "PerformanceHints contains a list of hints to Calico's policy engine to help process the policy more efficiently.  Hints never change the enforcement behaviour of the policy.  Currently, the only available hint is 'AssumeNeededOnEveryNode'.  When that hint is set on a policy, Felix will act as if the policy matches a local endpoint even if it does not. This is useful for 'preloading' any large static policies that are known to be used on every node. If the policy is _not_ used on a particular node then the work done to preload the policy (and to maintain it) is wasted.",
						MarkdownDescription: "PerformanceHints contains a list of hints to Calico's policy engine to help process the policy more efficiently.  Hints never change the enforcement behaviour of the policy.  Currently, the only available hint is 'AssumeNeededOnEveryNode'.  When that hint is set on a policy, Felix will act as if the policy matches a local endpoint even if it does not. This is useful for 'preloading' any large static policies that are known to be used on every node. If the policy is _not_ used on a particular node then the work done to preload the policy (and to maintain it) is wasted.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pre_dnat": schema.BoolAttribute{
						Description:         "PreDNAT indicates to apply the rules in this policy before any DNAT.",
						MarkdownDescription: "PreDNAT indicates to apply the rules in this policy before any DNAT.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.StringAttribute{
						Description:         "The selector is an expression used to pick out the endpoints that the policy should be applied to.  Selector expressions follow this syntax:  	label == 'string_literal'  ->  comparison, e.g. my_label == 'foo bar' 	label != 'string_literal'   ->  not equal; also matches if label is not present 	label in { 'a', 'b', 'c', ... }  ->  true if the value of label X is one of 'a', 'b', 'c' 	label not in { 'a', 'b', 'c', ... }  ->  true if the value of label X is not one of 'a', 'b', 'c' 	has(label_name)  -> True if that label is present 	! expr -> negation of expr 	expr && expr  -> Short-circuit and 	expr || expr  -> Short-circuit or 	( expr ) -> parens for grouping 	all() or the empty selector -> matches all endpoints.  Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive but they do not support escape characters.  Examples (with made-up labels):  	type == 'webserver' && deployment == 'prod' 	type in {'frontend', 'backend'} 	deployment != 'dev' 	! has(label_name)",
						MarkdownDescription: "The selector is an expression used to pick out the endpoints that the policy should be applied to.  Selector expressions follow this syntax:  	label == 'string_literal'  ->  comparison, e.g. my_label == 'foo bar' 	label != 'string_literal'   ->  not equal; also matches if label is not present 	label in { 'a', 'b', 'c', ... }  ->  true if the value of label X is one of 'a', 'b', 'c' 	label not in { 'a', 'b', 'c', ... }  ->  true if the value of label X is not one of 'a', 'b', 'c' 	has(label_name)  -> True if that label is present 	! expr -> negation of expr 	expr && expr  -> Short-circuit and 	expr || expr  -> Short-circuit or 	( expr ) -> parens for grouping 	all() or the empty selector -> matches all endpoints.  Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive but they do not support escape characters.  Examples (with made-up labels):  	type == 'webserver' && deployment == 'prod' 	type in {'frontend', 'backend'} 	deployment != 'dev' 	! has(label_name)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_selector": schema.StringAttribute{
						Description:         "ServiceAccountSelector is an optional field for an expression used to select a pod based on service accounts.",
						MarkdownDescription: "ServiceAccountSelector is an optional field for an expression used to select a pod based on service accounts.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tier": schema.StringAttribute{
						Description:         "The name of the tier that this policy belongs to.  If this is omitted, the default tier (name is 'default') is assumed.  The specified tier must exist in order to create security policies within the tier, the 'default' tier is created automatically if it does not exist, this means for deployments requiring only a single Tier, the tier name may be omitted on all policy management requests.",
						MarkdownDescription: "The name of the tier that this policy belongs to.  If this is omitted, the default tier (name is 'default') is assumed.  The specified tier must exist in order to create security policies within the tier, the 'default' tier is created automatically if it does not exist, this means for deployments requiring only a single Tier, the tier name may be omitted on all policy management requests.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"types": schema.ListAttribute{
						Description:         "Types indicates whether this policy applies to ingress, or to egress, or to both.  When not explicitly specified (and so the value on creation is empty or nil), Calico defaults Types according to what Ingress and Egress rules are present in the policy.  The default is:  - [ PolicyTypeIngress ], if there are no Egress rules (including the case where there are   also no Ingress rules)  - [ PolicyTypeEgress ], if there are Egress rules but no Ingress rules  - [ PolicyTypeIngress, PolicyTypeEgress ], if there are both Ingress and Egress rules.  When the policy is read back again, Types will always be one of these values, never empty or nil.",
						MarkdownDescription: "Types indicates whether this policy applies to ingress, or to egress, or to both.  When not explicitly specified (and so the value on creation is empty or nil), Calico defaults Types according to what Ingress and Egress rules are present in the policy.  The default is:  - [ PolicyTypeIngress ], if there are no Egress rules (including the case where there are   also no Ingress rules)  - [ PolicyTypeEgress ], if there are Egress rules but no Ingress rules  - [ PolicyTypeIngress, PolicyTypeEgress ], if there are both Ingress and Egress rules.  When the policy is read back again, Types will always be one of these values, never empty or nil.",
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
	}
}

func (r *CrdProjectcalicoOrgGlobalNetworkPolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_global_network_policy_v1_manifest")

	var model CrdProjectcalicoOrgGlobalNetworkPolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("GlobalNetworkPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
