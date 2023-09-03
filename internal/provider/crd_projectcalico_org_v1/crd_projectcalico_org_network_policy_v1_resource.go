/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
)

var (
	_ resource.Resource                = &CrdProjectcalicoOrgNetworkPolicyV1Resource{}
	_ resource.ResourceWithConfigure   = &CrdProjectcalicoOrgNetworkPolicyV1Resource{}
	_ resource.ResourceWithImportState = &CrdProjectcalicoOrgNetworkPolicyV1Resource{}
)

func NewCrdProjectcalicoOrgNetworkPolicyV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgNetworkPolicyV1Resource{}
}

type CrdProjectcalicoOrgNetworkPolicyV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CrdProjectcalicoOrgNetworkPolicyV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Egress *[]struct {
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
		Order                  *float64  `tfsdk:"order" json:"order,omitempty"`
		Selector               *string   `tfsdk:"selector" json:"selector,omitempty"`
		ServiceAccountSelector *string   `tfsdk:"service_account_selector" json:"serviceAccountSelector,omitempty"`
		Types                  *[]string `tfsdk:"types" json:"types,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_network_policy_v1"
}

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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

					"order": schema.Float64Attribute{
						Description:         "Order is an optional field that specifies the order in which the policy is applied. Policies with higher 'order' are applied after those with lower order.  If the order is omitted, it may be considered to be 'infinite' - i.e. the policy will be applied last.  Policies with identical order will be applied in alphanumerical order based on the Policy 'Name'.",
						MarkdownDescription: "Order is an optional field that specifies the order in which the policy is applied. Policies with higher 'order' are applied after those with lower order.  If the order is omitted, it may be considered to be 'infinite' - i.e. the policy will be applied last.  Policies with identical order will be applied in alphanumerical order based on the Policy 'Name'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selector": schema.StringAttribute{
						Description:         "The selector is an expression used to pick pick out the endpoints that the policy should be applied to.  Selector expressions follow this syntax:  	label == 'string_literal'  ->  comparison, e.g. my_label == 'foo bar' 	label != 'string_literal'   ->  not equal; also matches if label is not present 	label in { 'a', 'b', 'c', ... }  ->  true if the value of label X is one of 'a', 'b', 'c' 	label not in { 'a', 'b', 'c', ... }  ->  true if the value of label X is not one of 'a', 'b', 'c' 	has(label_name)  -> True if that label is present 	! expr -> negation of expr 	expr && expr  -> Short-circuit and 	expr || expr  -> Short-circuit or 	( expr ) -> parens for grouping 	all() or the empty selector -> matches all endpoints.  Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive but they do not support escape characters.  Examples (with made-up labels):  	type == 'webserver' && deployment == 'prod' 	type in {'frontend', 'backend'} 	deployment != 'dev' 	! has(label_name)",
						MarkdownDescription: "The selector is an expression used to pick pick out the endpoints that the policy should be applied to.  Selector expressions follow this syntax:  	label == 'string_literal'  ->  comparison, e.g. my_label == 'foo bar' 	label != 'string_literal'   ->  not equal; also matches if label is not present 	label in { 'a', 'b', 'c', ... }  ->  true if the value of label X is one of 'a', 'b', 'c' 	label not in { 'a', 'b', 'c', ... }  ->  true if the value of label X is not one of 'a', 'b', 'c' 	has(label_name)  -> True if that label is present 	! expr -> negation of expr 	expr && expr  -> Short-circuit and 	expr || expr  -> Short-circuit or 	( expr ) -> parens for grouping 	all() or the empty selector -> matches all endpoints.  Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive but they do not support escape characters.  Examples (with made-up labels):  	type == 'webserver' && deployment == 'prod' 	type in {'frontend', 'backend'} 	deployment != 'dev' 	! has(label_name)",
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

					"types": schema.ListAttribute{
						Description:         "Types indicates whether this policy applies to ingress, or to egress, or to both.  When not explicitly specified (and so the value on creation is empty or nil), Calico defaults Types according to what Ingress and Egress are present in the policy.  The default is:  - [ PolicyTypeIngress ], if there are no Egress rules (including the case where there are   also no Ingress rules)  - [ PolicyTypeEgress ], if there are Egress rules but no Ingress rules  - [ PolicyTypeIngress, PolicyTypeEgress ], if there are both Ingress and Egress rules.  When the policy is read back again, Types will always be one of these values, never empty or nil.",
						MarkdownDescription: "Types indicates whether this policy applies to ingress, or to egress, or to both.  When not explicitly specified (and so the value on creation is empty or nil), Calico defaults Types according to what Ingress and Egress are present in the policy.  The default is:  - [ PolicyTypeIngress ], if there are no Egress rules (including the case where there are   also no Ingress rules)  - [ PolicyTypeEgress ], if there are Egress rules but no Ingress rules  - [ PolicyTypeIngress, PolicyTypeEgress ], if there are both Ingress and Egress rules.  When the policy is read back again, Types will always be one of these values, never empty or nil.",
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

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_network_policy_v1")

	var model CrdProjectcalicoOrgNetworkPolicyV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("NetworkPolicy")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "NetworkPolicy"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CrdProjectcalicoOrgNetworkPolicyV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_network_policy_v1")

	var data CrdProjectcalicoOrgNetworkPolicyV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "NetworkPolicy"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse CrdProjectcalicoOrgNetworkPolicyV1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_network_policy_v1")

	var model CrdProjectcalicoOrgNetworkPolicyV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("NetworkPolicy")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "NetworkPolicy"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CrdProjectcalicoOrgNetworkPolicyV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_network_policy_v1")

	var data CrdProjectcalicoOrgNetworkPolicyV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "NetworkPolicy"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *CrdProjectcalicoOrgNetworkPolicyV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
