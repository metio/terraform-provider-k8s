/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networkfirewall_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1Manifest{}
)

func NewNetworkfirewallServicesK8SAwsRuleGroupV1Alpha1Manifest() datasource.DataSource {
	return &NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1Manifest{}
}

type NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1Manifest struct{}

type NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		AnalyzeRuleGroup        *bool   `tfsdk:"analyze_rule_group" json:"analyzeRuleGroup,omitempty"`
		Capacity                *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
		Description             *string `tfsdk:"description" json:"description,omitempty"`
		DryRun                  *bool   `tfsdk:"dry_run" json:"dryRun,omitempty"`
		EncryptionConfiguration *struct {
			KeyID *string `tfsdk:"key_id" json:"keyID,omitempty"`
			Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
		} `tfsdk:"encryption_configuration" json:"encryptionConfiguration,omitempty"`
		RuleGroup *struct {
			ReferenceSets *struct {
				IPSetReferences *struct {
					ReferenceARN *string `tfsdk:"reference_arn" json:"referenceARN,omitempty"`
				} `tfsdk:"i_p_set_references" json:"iPSetReferences,omitempty"`
			} `tfsdk:"reference_sets" json:"referenceSets,omitempty"`
			RuleVariables *struct {
				IPSets *struct {
					Definition *[]string `tfsdk:"definition" json:"definition,omitempty"`
				} `tfsdk:"i_p_sets" json:"iPSets,omitempty"`
				PortSets *struct {
					Definition *[]string `tfsdk:"definition" json:"definition,omitempty"`
				} `tfsdk:"port_sets" json:"portSets,omitempty"`
			} `tfsdk:"rule_variables" json:"ruleVariables,omitempty"`
			RulesSource *struct {
				RulesSourceList *struct {
					GeneratedRulesType *string   `tfsdk:"generated_rules_type" json:"generatedRulesType,omitempty"`
					TargetTypes        *[]string `tfsdk:"target_types" json:"targetTypes,omitempty"`
					Targets            *[]string `tfsdk:"targets" json:"targets,omitempty"`
				} `tfsdk:"rules_source_list" json:"rulesSourceList,omitempty"`
				RulesString   *string `tfsdk:"rules_string" json:"rulesString,omitempty"`
				StatefulRules *[]struct {
					Action *string `tfsdk:"action" json:"action,omitempty"`
					Header *struct {
						Destination     *string `tfsdk:"destination" json:"destination,omitempty"`
						DestinationPort *string `tfsdk:"destination_port" json:"destinationPort,omitempty"`
						Direction       *string `tfsdk:"direction" json:"direction,omitempty"`
						Protocol        *string `tfsdk:"protocol" json:"protocol,omitempty"`
						Source          *string `tfsdk:"source" json:"source,omitempty"`
						SourcePort      *string `tfsdk:"source_port" json:"sourcePort,omitempty"`
					} `tfsdk:"header" json:"header,omitempty"`
					RuleOptions *[]struct {
						Keyword  *string   `tfsdk:"keyword" json:"keyword,omitempty"`
						Settings *[]string `tfsdk:"settings" json:"settings,omitempty"`
					} `tfsdk:"rule_options" json:"ruleOptions,omitempty"`
				} `tfsdk:"stateful_rules" json:"statefulRules,omitempty"`
				StatelessRulesAndCustomActions *struct {
					CustomActions *[]struct {
						ActionDefinition *struct {
							PublishMetricAction *struct {
								Dimensions *[]struct {
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dimensions" json:"dimensions,omitempty"`
							} `tfsdk:"publish_metric_action" json:"publishMetricAction,omitempty"`
						} `tfsdk:"action_definition" json:"actionDefinition,omitempty"`
						ActionName *string `tfsdk:"action_name" json:"actionName,omitempty"`
					} `tfsdk:"custom_actions" json:"customActions,omitempty"`
					StatelessRules *[]struct {
						Priority       *int64 `tfsdk:"priority" json:"priority,omitempty"`
						RuleDefinition *struct {
							Actions         *[]string `tfsdk:"actions" json:"actions,omitempty"`
							MatchAttributes *struct {
								DestinationPorts *[]struct {
									FromPort *int64 `tfsdk:"from_port" json:"fromPort,omitempty"`
									ToPort   *int64 `tfsdk:"to_port" json:"toPort,omitempty"`
								} `tfsdk:"destination_ports" json:"destinationPorts,omitempty"`
								Destinations *[]struct {
									AddressDefinition *string `tfsdk:"address_definition" json:"addressDefinition,omitempty"`
								} `tfsdk:"destinations" json:"destinations,omitempty"`
								Protocols   *[]string `tfsdk:"protocols" json:"protocols,omitempty"`
								SourcePorts *[]struct {
									FromPort *int64 `tfsdk:"from_port" json:"fromPort,omitempty"`
									ToPort   *int64 `tfsdk:"to_port" json:"toPort,omitempty"`
								} `tfsdk:"source_ports" json:"sourcePorts,omitempty"`
								Sources *[]struct {
									AddressDefinition *string `tfsdk:"address_definition" json:"addressDefinition,omitempty"`
								} `tfsdk:"sources" json:"sources,omitempty"`
								TcpFlags *[]struct {
									Flags *[]string `tfsdk:"flags" json:"flags,omitempty"`
									Masks *[]string `tfsdk:"masks" json:"masks,omitempty"`
								} `tfsdk:"tcp_flags" json:"tcpFlags,omitempty"`
							} `tfsdk:"match_attributes" json:"matchAttributes,omitempty"`
						} `tfsdk:"rule_definition" json:"ruleDefinition,omitempty"`
					} `tfsdk:"stateless_rules" json:"statelessRules,omitempty"`
				} `tfsdk:"stateless_rules_and_custom_actions" json:"statelessRulesAndCustomActions,omitempty"`
			} `tfsdk:"rules_source" json:"rulesSource,omitempty"`
			StatefulRuleOptions *struct {
				RuleOrder *string `tfsdk:"rule_order" json:"ruleOrder,omitempty"`
			} `tfsdk:"stateful_rule_options" json:"statefulRuleOptions,omitempty"`
		} `tfsdk:"rule_group" json:"ruleGroup,omitempty"`
		RuleGroupName  *string `tfsdk:"rule_group_name" json:"ruleGroupName,omitempty"`
		Rules          *string `tfsdk:"rules" json:"rules,omitempty"`
		SourceMetadata *struct {
			SourceARN         *string `tfsdk:"source_arn" json:"sourceARN,omitempty"`
			SourceUpdateToken *string `tfsdk:"source_update_token" json:"sourceUpdateToken,omitempty"`
		} `tfsdk:"source_metadata" json:"sourceMetadata,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networkfirewall_services_k8s_aws_rule_group_v1alpha1_manifest"
}

func (r *NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RuleGroup is the Schema for the RuleGroups API",
		MarkdownDescription: "RuleGroup is the Schema for the RuleGroups API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "RuleGroupSpec defines the desired state of RuleGroup.The object that defines the rules in a rule group. This, along with RuleGroupResponse,define the rule group. You can retrieve all objects for a rule group by callingDescribeRuleGroup.Network Firewall uses a rule group to inspect and control network traffic.You define stateless rule groups to inspect individual packets and you definestateful rule groups to inspect packets in the context of their traffic flow.To use a rule group, you include it by reference in an Network Firewall firewallpolicy, then you use the policy in a firewall. You can reference a rule groupfrom more than one firewall policy, and you can use a firewall policy inmore than one firewall.",
				MarkdownDescription: "RuleGroupSpec defines the desired state of RuleGroup.The object that defines the rules in a rule group. This, along with RuleGroupResponse,define the rule group. You can retrieve all objects for a rule group by callingDescribeRuleGroup.Network Firewall uses a rule group to inspect and control network traffic.You define stateless rule groups to inspect individual packets and you definestateful rule groups to inspect packets in the context of their traffic flow.To use a rule group, you include it by reference in an Network Firewall firewallpolicy, then you use the policy in a firewall. You can reference a rule groupfrom more than one firewall policy, and you can use a firewall policy inmore than one firewall.",
				Attributes: map[string]schema.Attribute{
					"analyze_rule_group": schema.BoolAttribute{
						Description:         "Indicates whether you want Network Firewall to analyze the stateless rulesin the rule group for rule behavior such as asymmetric routing. If set toTRUE, Network Firewall runs the analysis and then creates the rule groupfor you. To run the stateless rule group analyzer without creating the rulegroup, set DryRun to TRUE.",
						MarkdownDescription: "Indicates whether you want Network Firewall to analyze the stateless rulesin the rule group for rule behavior such as asymmetric routing. If set toTRUE, Network Firewall runs the analysis and then creates the rule groupfor you. To run the stateless rule group analyzer without creating the rulegroup, set DryRun to TRUE.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"capacity": schema.Int64Attribute{
						Description:         "The maximum operating resources that this rule group can use. Rule groupcapacity is fixed at creation. When you update a rule group, you are limitedto this capacity. When you reference a rule group from a firewall policy,Network Firewall reserves this capacity for the rule group.You can retrieve the capacity that would be required for a rule group beforeyou create the rule group by calling CreateRuleGroup with DryRun set to TRUE.You can't change or exceed this capacity when you update the rule group,so leave room for your rule group to grow.Capacity for a stateless rule groupFor a stateless rule group, the capacity required is the sum of the capacityrequirements of the individual rules that you expect to have in the rulegroup.To calculate the capacity requirement of a single rule, multiply the capacityrequirement values of each of the rule's match settings:   * A match setting with no criteria specified has a value of 1.   * A match setting with Any specified has a value of 1.   * All other match settings have a value equal to the number of elements   provided in the setting. For example, a protocol setting ['UDP'] and a   source setting ['10.0.0.0/24'] each have a value of 1. A protocol setting   ['UDP','TCP'] has a value of 2. A source setting ['10.0.0.0/24','10.0.0.1/24','10.0.0.2/24']   has a value of 3.A rule with no criteria specified in any of its match settings has a capacityrequirement of 1. A rule with protocol setting ['UDP','TCP'], source setting['10.0.0.0/24','10.0.0.1/24','10.0.0.2/24'], and a single specification orno specification for each of the other match settings has a capacity requirementof 6.Capacity for a stateful rule groupFor a stateful rule group, the minimum capacity required is the number ofindividual rules that you expect to have in the rule group.",
						MarkdownDescription: "The maximum operating resources that this rule group can use. Rule groupcapacity is fixed at creation. When you update a rule group, you are limitedto this capacity. When you reference a rule group from a firewall policy,Network Firewall reserves this capacity for the rule group.You can retrieve the capacity that would be required for a rule group beforeyou create the rule group by calling CreateRuleGroup with DryRun set to TRUE.You can't change or exceed this capacity when you update the rule group,so leave room for your rule group to grow.Capacity for a stateless rule groupFor a stateless rule group, the capacity required is the sum of the capacityrequirements of the individual rules that you expect to have in the rulegroup.To calculate the capacity requirement of a single rule, multiply the capacityrequirement values of each of the rule's match settings:   * A match setting with no criteria specified has a value of 1.   * A match setting with Any specified has a value of 1.   * All other match settings have a value equal to the number of elements   provided in the setting. For example, a protocol setting ['UDP'] and a   source setting ['10.0.0.0/24'] each have a value of 1. A protocol setting   ['UDP','TCP'] has a value of 2. A source setting ['10.0.0.0/24','10.0.0.1/24','10.0.0.2/24']   has a value of 3.A rule with no criteria specified in any of its match settings has a capacityrequirement of 1. A rule with protocol setting ['UDP','TCP'], source setting['10.0.0.0/24','10.0.0.1/24','10.0.0.2/24'], and a single specification orno specification for each of the other match settings has a capacity requirementof 6.Capacity for a stateful rule groupFor a stateful rule group, the minimum capacity required is the number ofindividual rules that you expect to have in the rule group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"description": schema.StringAttribute{
						Description:         "A description of the rule group.",
						MarkdownDescription: "A description of the rule group.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dry_run": schema.BoolAttribute{
						Description:         "Indicates whether you want Network Firewall to just check the validity ofthe request, rather than run the request.If set to TRUE, Network Firewall checks whether the request can run successfully,but doesn't actually make the requested changes. The call returns the valuethat the request would return if you ran it with dry run set to FALSE, butdoesn't make additions or changes to your resources. This option allows youto make sure that you have the required permissions to run the request andthat your request parameters are valid.If set to FALSE, Network Firewall makes the requested changes to your resources.",
						MarkdownDescription: "Indicates whether you want Network Firewall to just check the validity ofthe request, rather than run the request.If set to TRUE, Network Firewall checks whether the request can run successfully,but doesn't actually make the requested changes. The call returns the valuethat the request would return if you ran it with dry run set to FALSE, butdoesn't make additions or changes to your resources. This option allows youto make sure that you have the required permissions to run the request andthat your request parameters are valid.If set to FALSE, Network Firewall makes the requested changes to your resources.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encryption_configuration": schema.SingleNestedAttribute{
						Description:         "A complex type that contains settings for encryption of your rule group resources.",
						MarkdownDescription: "A complex type that contains settings for encryption of your rule group resources.",
						Attributes: map[string]schema.Attribute{
							"key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type_": schema.StringAttribute{
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

					"rule_group": schema.SingleNestedAttribute{
						Description:         "An object that defines the rule group rules.You must provide either this rule group setting or a Rules setting, but notboth.",
						MarkdownDescription: "An object that defines the rule group rules.You must provide either this rule group setting or a Rules setting, but notboth.",
						Attributes: map[string]schema.Attribute{
							"reference_sets": schema.SingleNestedAttribute{
								Description:         "Contains a set of IP set references.",
								MarkdownDescription: "Contains a set of IP set references.",
								Attributes: map[string]schema.Attribute{
									"i_p_set_references": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"reference_arn": schema.StringAttribute{
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

							"rule_variables": schema.SingleNestedAttribute{
								Description:         "Settings that are available for use in the rules in the RuleGroup where thisis defined.",
								MarkdownDescription: "Settings that are available for use in the rules in the RuleGroup where thisis defined.",
								Attributes: map[string]schema.Attribute{
									"i_p_sets": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"definition": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"port_sets": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"definition": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rules_source": schema.SingleNestedAttribute{
								Description:         "The stateless or stateful rules definitions for use in a single rule group.Each rule group requires a single RulesSource. You can use an instance ofthis for either stateless rules or stateful rules.",
								MarkdownDescription: "The stateless or stateful rules definitions for use in a single rule group.Each rule group requires a single RulesSource. You can use an instance ofthis for either stateless rules or stateful rules.",
								Attributes: map[string]schema.Attribute{
									"rules_source_list": schema.SingleNestedAttribute{
										Description:         "Stateful inspection criteria for a domain list rule group.For HTTPS traffic, domain filtering is SNI-based. It uses the server nameindicator extension of the TLS handshake.By default, Network Firewall domain list inspection only includes trafficcoming from the VPC where you deploy the firewall. To inspect traffic fromIP addresses outside of the deployment VPC, you set the HOME_NET rule variableto include the CIDR range of the deployment VPC plus the other CIDR ranges.For more information, see RuleVariables in this guide and Stateful domainlist rule groups in Network Firewall (https://docs.aws.amazon.com/network-firewall/latest/developerguide/stateful-rule-groups-domain-names.html)in the Network Firewall Developer Guide.",
										MarkdownDescription: "Stateful inspection criteria for a domain list rule group.For HTTPS traffic, domain filtering is SNI-based. It uses the server nameindicator extension of the TLS handshake.By default, Network Firewall domain list inspection only includes trafficcoming from the VPC where you deploy the firewall. To inspect traffic fromIP addresses outside of the deployment VPC, you set the HOME_NET rule variableto include the CIDR range of the deployment VPC plus the other CIDR ranges.For more information, see RuleVariables in this guide and Stateful domainlist rule groups in Network Firewall (https://docs.aws.amazon.com/network-firewall/latest/developerguide/stateful-rule-groups-domain-names.html)in the Network Firewall Developer Guide.",
										Attributes: map[string]schema.Attribute{
											"generated_rules_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_types": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"targets": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"rules_string": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"stateful_rules": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"header": schema.SingleNestedAttribute{
													Description:         "The basic rule criteria for Network Firewall to use to inspect packet headersin stateful traffic flow inspection. Traffic flows that match the criteriaare a match for the corresponding StatefulRule.",
													MarkdownDescription: "The basic rule criteria for Network Firewall to use to inspect packet headersin stateful traffic flow inspection. Traffic flows that match the criteriaare a match for the corresponding StatefulRule.",
													Attributes: map[string]schema.Attribute{
														"destination": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"destination_port": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"direction": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"protocol": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_port": schema.StringAttribute{
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

												"rule_options": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"keyword": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"settings": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"stateless_rules_and_custom_actions": schema.SingleNestedAttribute{
										Description:         "Stateless inspection criteria. Each stateless rule group uses exactly oneof these data types to define its stateless rules.",
										MarkdownDescription: "Stateless inspection criteria. Each stateless rule group uses exactly oneof these data types to define its stateless rules.",
										Attributes: map[string]schema.Attribute{
											"custom_actions": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action_definition": schema.SingleNestedAttribute{
															Description:         "A custom action to use in stateless rule actions settings. This is used inCustomAction.",
															MarkdownDescription: "A custom action to use in stateless rule actions settings. This is used inCustomAction.",
															Attributes: map[string]schema.Attribute{
																"publish_metric_action": schema.SingleNestedAttribute{
																	Description:         "Stateless inspection criteria that publishes the specified metrics to AmazonCloudWatch for the matching packet. This setting defines a CloudWatch dimensionvalue to be published.",
																	MarkdownDescription: "Stateless inspection criteria that publishes the specified metrics to AmazonCloudWatch for the matching packet. This setting defines a CloudWatch dimensionvalue to be published.",
																	Attributes: map[string]schema.Attribute{
																		"dimensions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"value": schema.StringAttribute{
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"action_name": schema.StringAttribute{
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

											"stateless_rules": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"priority": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"rule_definition": schema.SingleNestedAttribute{
															Description:         "The inspection criteria and action for a single stateless rule. Network Firewallinspects each packet for the specified matching criteria. When a packet matchesthe criteria, Network Firewall performs the rule's actions on the packet.",
															MarkdownDescription: "The inspection criteria and action for a single stateless rule. Network Firewallinspects each packet for the specified matching criteria. When a packet matchesthe criteria, Network Firewall performs the rule's actions on the packet.",
															Attributes: map[string]schema.Attribute{
																"actions": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"match_attributes": schema.SingleNestedAttribute{
																	Description:         "Criteria for Network Firewall to use to inspect an individual packet in statelessrule inspection. Each match attributes set can include one or more itemssuch as IP address, CIDR range, port number, protocol, and TCP flags.",
																	MarkdownDescription: "Criteria for Network Firewall to use to inspect an individual packet in statelessrule inspection. Each match attributes set can include one or more itemssuch as IP address, CIDR range, port number, protocol, and TCP flags.",
																	Attributes: map[string]schema.Attribute{
																		"destination_ports": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"from_port": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"to_port": schema.Int64Attribute{
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

																		"destinations": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"address_definition": schema.StringAttribute{
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

																		"protocols": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"source_ports": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"from_port": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"to_port": schema.Int64Attribute{
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

																		"sources": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"address_definition": schema.StringAttribute{
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

																		"tcp_flags": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"flags": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"masks": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

							"stateful_rule_options": schema.SingleNestedAttribute{
								Description:         "Additional options governing how Network Firewall handles the rule group.You can only use these for stateful rule groups.",
								MarkdownDescription: "Additional options governing how Network Firewall handles the rule group.You can only use these for stateful rule groups.",
								Attributes: map[string]schema.Attribute{
									"rule_order": schema.StringAttribute{
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

					"rule_group_name": schema.StringAttribute{
						Description:         "The descriptive name of the rule group. You can't change the name of a rulegroup after you create it.",
						MarkdownDescription: "The descriptive name of the rule group. You can't change the name of a rulegroup after you create it.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"rules": schema.StringAttribute{
						Description:         "A string containing stateful rule group rules specifications in Suricataflat format, with one rule per line. Use this to import your existing Suricatacompatible rule groups.You must provide either this rules setting or a populated RuleGroup setting,but not both.You can provide your rule group specification in Suricata flat format throughthis setting when you create or update your rule group. The call responsereturns a RuleGroup object that Network Firewall has populated from yourstring.",
						MarkdownDescription: "A string containing stateful rule group rules specifications in Suricataflat format, with one rule per line. Use this to import your existing Suricatacompatible rule groups.You must provide either this rules setting or a populated RuleGroup setting,but not both.You can provide your rule group specification in Suricata flat format throughthis setting when you create or update your rule group. The call responsereturns a RuleGroup object that Network Firewall has populated from yourstring.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_metadata": schema.SingleNestedAttribute{
						Description:         "A complex type that contains metadata about the rule group that your ownrule group is copied from. You can use the metadata to keep track of updatesmade to the originating rule group.",
						MarkdownDescription: "A complex type that contains metadata about the rule group that your ownrule group is copied from. You can use the metadata to keep track of updatesmade to the originating rule group.",
						Attributes: map[string]schema.Attribute{
							"source_arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source_update_token": schema.StringAttribute{
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

					"tags": schema.ListNestedAttribute{
						Description:         "The key:value pairs to associate with the resource.",
						MarkdownDescription: "The key:value pairs to associate with the resource.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
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

					"type_": schema.StringAttribute{
						Description:         "Indicates whether the rule group is stateless or stateful. If the rule groupis stateless, it contains stateless rules. If it is stateful, it containsstateful rules.",
						MarkdownDescription: "Indicates whether the rule group is stateless or stateful. If the rule groupis stateless, it contains stateless rules. If it is stateful, it containsstateful rules.",
						Required:            true,
						Optional:            false,
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

func (r *NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networkfirewall_services_k8s_aws_rule_group_v1alpha1_manifest")

	var model NetworkfirewallServicesK8SAwsRuleGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("networkfirewall.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("RuleGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
