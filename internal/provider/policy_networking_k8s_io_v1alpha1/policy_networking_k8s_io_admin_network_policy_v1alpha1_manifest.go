/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package policy_networking_k8s_io_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1Manifest{}
)

func NewPolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1Manifest() datasource.DataSource {
	return &PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1Manifest{}
}

type PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1Manifest struct{}

type PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Egress *[]struct {
			Action *string `tfsdk:"action" json:"action,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Ports  *[]struct {
				PortNumber *struct {
					Port     *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"port_number" json:"portNumber,omitempty"`
				PortRange *struct {
					End      *int64  `tfsdk:"end" json:"end,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Start    *int64  `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"port_range" json:"portRange,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
			To *[]struct {
				Namespaces *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Pods *struct {
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				} `tfsdk:"pods" json:"pods,omitempty"`
			} `tfsdk:"to" json:"to,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		Ingress *[]struct {
			Action *string `tfsdk:"action" json:"action,omitempty"`
			From   *[]struct {
				Namespaces *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Pods *struct {
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				} `tfsdk:"pods" json:"pods,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Ports *[]struct {
				PortNumber *struct {
					Port     *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"port_number" json:"portNumber,omitempty"`
				PortRange *struct {
					End      *int64  `tfsdk:"end" json:"end,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Start    *int64  `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"port_range" json:"portRange,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Priority *int64 `tfsdk:"priority" json:"priority,omitempty"`
		Subject  *struct {
			Namespaces *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespaces" json:"namespaces,omitempty"`
			Pods *struct {
				NamespaceSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				PodSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
			} `tfsdk:"pods" json:"pods,omitempty"`
		} `tfsdk:"subject" json:"subject,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest"
}

func (r *PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AdminNetworkPolicy is a cluster level resource that is part of the AdminNetworkPolicy API.",
		MarkdownDescription: "AdminNetworkPolicy is a cluster level resource that is part of the AdminNetworkPolicy API.",
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
				Description:         "Specification of the desired behavior of AdminNetworkPolicy.",
				MarkdownDescription: "Specification of the desired behavior of AdminNetworkPolicy.",
				Attributes: map[string]schema.Attribute{
					"egress": schema.ListNestedAttribute{
						Description:         "Egress is the list of Egress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of egress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the egress rules would take the highest precedence. ANPs with no egress rules do not affect egress traffic. Support: Core",
						MarkdownDescription: "Egress is the list of Egress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of egress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the egress rules would take the highest precedence. ANPs with no egress rules do not affect egress traffic. Support: Core",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod. Support: Core",
									MarkdownDescription: "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod. Support: Core",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Allow", "Deny", "Pass"),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies. Support: Core",
									MarkdownDescription: "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies. Support: Core",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(100),
									},
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports allows for matching traffic based on port and protocols. This field is a list of destination ports for the outgoing egress traffic. If Ports is not set then the rule does not filter traffic via port. Support: Core",
									MarkdownDescription: "Ports allows for matching traffic based on port and protocols. This field is a list of destination ports for the outgoing egress traffic. If Ports is not set then the rule does not filter traffic via port. Support: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"port_number": schema.SingleNestedAttribute{
												Description:         "Port selects a port on a pod(s) based on number. Support: Core",
												MarkdownDescription: "Port selects a port on a pod(s) based on number. Support: Core",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Number defines a network port value. Support: Core",
														MarkdownDescription: "Number defines a network port value. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("port_range")),
												},
											},

											"port_range": schema.SingleNestedAttribute{
												Description:         "PortRange selects a port range on a pod(s) based on provided start and end values. Support: Core",
												MarkdownDescription: "PortRange selects a port range on a pod(s) based on provided start and end values. Support: Core",
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "End defines a network port that is the end of a port range, the End value must be greater than Start. Support: Core",
														MarkdownDescription: "End defines a network port that is the end of a port range, the End value must be greater than Start. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"start": schema.Int64Attribute{
														Description:         "Start defines a network port that is the start of a port range, the Start value must be less than End. Support: Core",
														MarkdownDescription: "Start defines a network port that is the start of a port range, the Start value must be less than End. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("port_number")),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to": schema.ListNestedAttribute{
									Description:         "To is the List of destinations whose traffic this rule applies to. If any AdminNetworkPolicyEgressPeer matches the destination of outgoing traffic then the specified action is applied. This field must be defined and contain at least one item. Support: Core",
									MarkdownDescription: "To is the List of destinations whose traffic this rule applies to. If any AdminNetworkPolicyEgressPeer matches the destination of outgoing traffic then the specified action is applied. This field must be defined and contain at least one item. Support: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"namespaces": schema.SingleNestedAttribute{
												Description:         "Namespaces defines a way to select all pods within a set of Namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
												MarkdownDescription: "Namespaces defines a way to select all pods within a set of Namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
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
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("pods")),
												},
											},

											"pods": schema.SingleNestedAttribute{
												Description:         "Pods defines a way to select a set of pods in a set of namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
												MarkdownDescription: "Pods defines a way to select a set of pods in a set of namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
												Attributes: map[string]schema.Attribute{
													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
														MarkdownDescription: "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
														MarkdownDescription: "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespaces")),
												},
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": schema.ListNestedAttribute{
						Description:         "Ingress is the list of Ingress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of ingress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the ingress rules would take the highest precedence. ANPs with no ingress rules do not affect ingress traffic. Support: Core",
						MarkdownDescription: "Ingress is the list of Ingress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of ingress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the ingress rules would take the highest precedence. ANPs with no ingress rules do not affect ingress traffic. Support: Core",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod. Support: Core",
									MarkdownDescription: "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod. Support: Core",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Allow", "Deny", "Pass"),
									},
								},

								"from": schema.ListNestedAttribute{
									Description:         "From is the list of sources whose traffic this rule applies to. If any AdminNetworkPolicyIngressPeer matches the source of incoming traffic then the specified action is applied. This field must be defined and contain at least one item. Support: Core",
									MarkdownDescription: "From is the list of sources whose traffic this rule applies to. If any AdminNetworkPolicyIngressPeer matches the source of incoming traffic then the specified action is applied. This field must be defined and contain at least one item. Support: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"namespaces": schema.SingleNestedAttribute{
												Description:         "Namespaces defines a way to select all pods within a set of Namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
												MarkdownDescription: "Namespaces defines a way to select all pods within a set of Namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
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
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("pods")),
												},
											},

											"pods": schema.SingleNestedAttribute{
												Description:         "Pods defines a way to select a set of pods in a set of namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
												MarkdownDescription: "Pods defines a way to select a set of pods in a set of namespaces. Note that host-networked pods are not included in this type of peer. Support: Core",
												Attributes: map[string]schema.Attribute{
													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
														MarkdownDescription: "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
														MarkdownDescription: "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespaces")),
												},
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies. Support: Core",
									MarkdownDescription: "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies. Support: Core",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(100),
									},
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports allows for matching traffic based on port and protocols. This field is a list of ports which should be matched on the pods selected for this policy i.e the subject of the policy. So it matches on the destination port for the ingress traffic. If Ports is not set then the rule does not filter traffic via port. Support: Core",
									MarkdownDescription: "Ports allows for matching traffic based on port and protocols. This field is a list of ports which should be matched on the pods selected for this policy i.e the subject of the policy. So it matches on the destination port for the ingress traffic. If Ports is not set then the rule does not filter traffic via port. Support: Core",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"port_number": schema.SingleNestedAttribute{
												Description:         "Port selects a port on a pod(s) based on number. Support: Core",
												MarkdownDescription: "Port selects a port on a pod(s) based on number. Support: Core",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Number defines a network port value. Support: Core",
														MarkdownDescription: "Number defines a network port value. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("port_range")),
												},
											},

											"port_range": schema.SingleNestedAttribute{
												Description:         "PortRange selects a port range on a pod(s) based on provided start and end values. Support: Core",
												MarkdownDescription: "PortRange selects a port range on a pod(s) based on provided start and end values. Support: Core",
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "End defines a network port that is the end of a port range, the End value must be greater than Start. Support: Core",
														MarkdownDescription: "End defines a network port that is the end of a port range, the End value must be greater than Start. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP. Support: Core",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"start": schema.Int64Attribute{
														Description:         "Start defines a network port that is the start of a port range, the Start value must be less than End. Support: Core",
														MarkdownDescription: "Start defines a network port that is the start of a port range, the Start value must be less than End. Support: Core",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("port_number")),
												},
											},
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

					"priority": schema.Int64Attribute{
						Description:         "Priority is a value from 0 to 1000. Policies with lower priority values have higher precedence, and are checked before policies with higher priority values. All AdminNetworkPolicy rules have higher precedence than NetworkPolicy or BaselineAdminNetworkPolicy rules Every AdminNetworkPolicy should have a unique priority value; if two (or more) policies with the same priority could both match a connection, then the implementation can apply any of the matching policies to the connection, and there is no way for the user to reliably determine which one it will choose. Support: Core",
						MarkdownDescription: "Priority is a value from 0 to 1000. Policies with lower priority values have higher precedence, and are checked before policies with higher priority values. All AdminNetworkPolicy rules have higher precedence than NetworkPolicy or BaselineAdminNetworkPolicy rules Every AdminNetworkPolicy should have a unique priority value; if two (or more) policies with the same priority could both match a connection, then the implementation can apply any of the matching policies to the connection, and there is no way for the user to reliably determine which one it will choose. Support: Core",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(1000),
						},
					},

					"subject": schema.SingleNestedAttribute{
						Description:         "Subject defines the pods to which this AdminNetworkPolicy applies. Note that host-networked pods are not included in subject selection. Support: Core",
						MarkdownDescription: "Subject defines the pods to which this AdminNetworkPolicy applies. Note that host-networked pods are not included in subject selection. Support: Core",
						Attributes: map[string]schema.Attribute{
							"namespaces": schema.SingleNestedAttribute{
								Description:         "Namespaces is used to select pods via namespace selectors.",
								MarkdownDescription: "Namespaces is used to select pods via namespace selectors.",
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
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("pods")),
								},
							},

							"pods": schema.SingleNestedAttribute{
								Description:         "Pods is used to select pods via namespace AND pod selectors.",
								MarkdownDescription: "Pods is used to select pods via namespace AND pod selectors.",
								Attributes: map[string]schema.Attribute{
									"namespace_selector": schema.SingleNestedAttribute{
										Description:         "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
										MarkdownDescription: "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"pod_selector": schema.SingleNestedAttribute{
										Description:         "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
										MarkdownDescription: "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespaces")),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_policy_networking_k8s_io_admin_network_policy_v1alpha1_manifest")

	var model PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("policy.networking.k8s.io/v1alpha1")
	model.Kind = pointer.String("AdminNetworkPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
