/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package policy_networking_k8s_io_v1alpha1

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
	_ datasource.DataSource              = &PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource{}
)

func NewPolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource() datasource.DataSource {
	return &PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource{}
}

type PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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
				NamedPort  *string `tfsdk:"named_port" json:"namedPort,omitempty"`
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
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
					SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
				} `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Pods *struct {
					Namespaces *struct {
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
						SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
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
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
					SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
				} `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Pods *struct {
					Namespaces *struct {
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
						SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
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
				NamedPort  *string `tfsdk:"named_port" json:"namedPort,omitempty"`
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

func (r *PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_policy_networking_k8s_io_admin_network_policy_v1alpha1"
}

func (r *PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AdminNetworkPolicy is  a cluster level resource that is part of the AdminNetworkPolicy API.",
		MarkdownDescription: "AdminNetworkPolicy is  a cluster level resource that is part of the AdminNetworkPolicy API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "Specification of the desired behavior of AdminNetworkPolicy.",
				MarkdownDescription: "Specification of the desired behavior of AdminNetworkPolicy.",
				Attributes: map[string]schema.Attribute{
					"egress": schema.ListNestedAttribute{
						Description:         "Egress is the list of Egress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of egress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the egress rules would take the highest precedence. ANPs with no egress rules do not affect egress traffic.",
						MarkdownDescription: "Egress is the list of Egress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of egress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the egress rules would take the highest precedence. ANPs with no egress rules do not affect egress traffic.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod.",
									MarkdownDescription: "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies.",
									MarkdownDescription: "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports allows for matching traffic based on port and protocols. This field is a list of destination ports for the outging egress traffic. If Ports is not set then the rule does not filter traffic via port.",
									MarkdownDescription: "Ports allows for matching traffic based on port and protocols. This field is a list of destination ports for the outging egress traffic. If Ports is not set then the rule does not filter traffic via port.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"named_port": schema.StringAttribute{
												Description:         "NamedPort selects a port on a pod(s) based on name.",
												MarkdownDescription: "NamedPort selects a port on a pod(s) based on name.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"port_number": schema.SingleNestedAttribute{
												Description:         "Port selects a port on a pod(s) based on number.",
												MarkdownDescription: "Port selects a port on a pod(s) based on number.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Number defines a network port value.",
														MarkdownDescription: "Number defines a network port value.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"port_range": schema.SingleNestedAttribute{
												Description:         "PortRange selects a port range on a pod(s) based on provided start and end values.",
												MarkdownDescription: "PortRange selects a port range on a pod(s) based on provided start and end values.",
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														MarkdownDescription: "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"start": schema.Int64Attribute{
														Description:         "Start defines a network port that is the start of a port range, the Start value must be less than End.",
														MarkdownDescription: "Start defines a network port that is the start of a port range, the Start value must be less than End.",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"to": schema.ListNestedAttribute{
									Description:         "To is the List of destinations whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the destination of outgoing traffic then the specified action is applied. This field must be defined and contain at least one item.",
									MarkdownDescription: "To is the List of destinations whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the destination of outgoing traffic then the specified action is applied. This field must be defined and contain at least one item.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"namespaces": schema.SingleNestedAttribute{
												Description:         "Namespaces defines a way to select a set of Namespaces.",
												MarkdownDescription: "Namespaces defines a way to select a set of Namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
														MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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

													"not_same_labels": schema.ListAttribute{
														Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"same_labels": schema.ListAttribute{
														Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
														MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
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

											"pods": schema.SingleNestedAttribute{
												Description:         "Pods defines a way to select a set of pods in in a set of namespaces.",
												MarkdownDescription: "Pods defines a way to select a set of pods in in a set of namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespaces": schema.SingleNestedAttribute{
														Description:         "Namespaces is used to select a set of Namespaces.",
														MarkdownDescription: "Namespaces is used to select a set of Namespaces.",
														Attributes: map[string]schema.Attribute{
															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
																MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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

															"not_same_labels": schema.ListAttribute{
																Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"same_labels": schema.ListAttribute{
																Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
																MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
														MarkdownDescription: "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
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
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ingress": schema.ListNestedAttribute{
						Description:         "Ingress is the list of Ingress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of ingress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the ingress rules would take the highest precedence. ANPs with no ingress rules do not affect ingress traffic.",
						MarkdownDescription: "Ingress is the list of Ingress rules to be applied to the selected pods. A total of 100 rules will be allowed in each ANP instance. The relative precedence of ingress rules within a single ANP object (all of which share the priority) will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the ingress rules would take the highest precedence. ANPs with no ingress rules do not affect ingress traffic.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod.",
									MarkdownDescription: "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic (even if it would otherwise have been denied by NetworkPolicy) Deny: denies the selected traffic Pass: instructs the selected traffic to skip any remaining ANP rules, and then pass execution to any NetworkPolicies that select the pod. If the pod is not selected by any NetworkPolicies then execution is passed to any BaselineAdminNetworkPolicies that select the pod.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"from": schema.ListNestedAttribute{
									Description:         "From is the list of sources whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the source of incoming traffic then the specified action is applied. This field must be defined and contain at least one item.",
									MarkdownDescription: "From is the list of sources whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the source of incoming traffic then the specified action is applied. This field must be defined and contain at least one item.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"namespaces": schema.SingleNestedAttribute{
												Description:         "Namespaces defines a way to select a set of Namespaces.",
												MarkdownDescription: "Namespaces defines a way to select a set of Namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
														MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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

													"not_same_labels": schema.ListAttribute{
														Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"same_labels": schema.ListAttribute{
														Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
														MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
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

											"pods": schema.SingleNestedAttribute{
												Description:         "Pods defines a way to select a set of pods in in a set of namespaces.",
												MarkdownDescription: "Pods defines a way to select a set of pods in in a set of namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespaces": schema.SingleNestedAttribute{
														Description:         "Namespaces is used to select a set of Namespaces.",
														MarkdownDescription: "Namespaces is used to select a set of Namespaces.",
														Attributes: map[string]schema.Attribute{
															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
																MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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

															"not_same_labels": schema.ListAttribute{
																Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"same_labels": schema.ListAttribute{
																Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
																MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
														MarkdownDescription: "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
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
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies.",
									MarkdownDescription: "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied AdminNetworkPolicies.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports allows for matching traffic based on port and protocols. This field is a list of ports which should be matched on the pods selected for this policy i.e the subject of the policy. So it matches on the destination port for the ingress traffic. If Ports is not set then the rule does not filter traffic via port.",
									MarkdownDescription: "Ports allows for matching traffic based on port and protocols. This field is a list of ports which should be matched on the pods selected for this policy i.e the subject of the policy. So it matches on the destination port for the ingress traffic. If Ports is not set then the rule does not filter traffic via port.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"named_port": schema.StringAttribute{
												Description:         "NamedPort selects a port on a pod(s) based on name.",
												MarkdownDescription: "NamedPort selects a port on a pod(s) based on name.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"port_number": schema.SingleNestedAttribute{
												Description:         "Port selects a port on a pod(s) based on number.",
												MarkdownDescription: "Port selects a port on a pod(s) based on number.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Number defines a network port value.",
														MarkdownDescription: "Number defines a network port value.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"port_range": schema.SingleNestedAttribute{
												Description:         "PortRange selects a port range on a pod(s) based on provided start and end values.",
												MarkdownDescription: "PortRange selects a port range on a pod(s) based on provided start and end values.",
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														MarkdownDescription: "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"start": schema.Int64Attribute{
														Description:         "Start defines a network port that is the start of a port range, the Start value must be less than End.",
														MarkdownDescription: "Start defines a network port that is the start of a port range, the Start value must be less than End.",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority is a value from 0 to 1000. Rules with lower priority values have higher precedence, and are checked before rules with higher priority values. All AdminNetworkPolicy rules have higher precedence than NetworkPolicy or BaselineAdminNetworkPolicy rules The behavior is undefined if two ANP objects have same priority.",
						MarkdownDescription: "Priority is a value from 0 to 1000. Rules with lower priority values have higher precedence, and are checked before rules with higher priority values. All AdminNetworkPolicy rules have higher precedence than NetworkPolicy or BaselineAdminNetworkPolicy rules The behavior is undefined if two ANP objects have same priority.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"subject": schema.SingleNestedAttribute{
						Description:         "Subject defines the pods to which this AdminNetworkPolicy applies.",
						MarkdownDescription: "Subject defines the pods to which this AdminNetworkPolicy applies.",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_policy_networking_k8s_io_admin_network_policy_v1alpha1")

	var data PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "policy.networking.k8s.io", Version: "v1alpha1", Resource: "adminnetworkpolicies"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse PolicyNetworkingK8SIoAdminNetworkPolicyV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("policy.networking.k8s.io/v1alpha1")
	data.Kind = pointer.String("AdminNetworkPolicy")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
