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
	_ datasource.DataSource = &NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1Manifest{}
)

func NewNetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1Manifest() datasource.DataSource {
	return &NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1Manifest{}
}

type NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1Manifest struct{}

type NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1ManifestData struct {
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
		Description             *string `tfsdk:"description" json:"description,omitempty"`
		EncryptionConfiguration *struct {
			KeyID *string `tfsdk:"key_id" json:"keyID,omitempty"`
			Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
		} `tfsdk:"encryption_configuration" json:"encryptionConfiguration,omitempty"`
		FirewallPolicy *struct {
			PolicyVariables *struct {
				RuleVariables *struct {
					Definition *[]string `tfsdk:"definition" json:"definition,omitempty"`
				} `tfsdk:"rule_variables" json:"ruleVariables,omitempty"`
			} `tfsdk:"policy_variables" json:"policyVariables,omitempty"`
			StatefulDefaultActions *[]string `tfsdk:"stateful_default_actions" json:"statefulDefaultActions,omitempty"`
			StatefulEngineOptions  *struct {
				RuleOrder             *string `tfsdk:"rule_order" json:"ruleOrder,omitempty"`
				StreamExceptionPolicy *string `tfsdk:"stream_exception_policy" json:"streamExceptionPolicy,omitempty"`
			} `tfsdk:"stateful_engine_options" json:"statefulEngineOptions,omitempty"`
			StatefulRuleGroupReferences *[]struct {
				Override *struct {
					Action *string `tfsdk:"action" json:"action,omitempty"`
				} `tfsdk:"override" json:"override,omitempty"`
				Priority    *int64  `tfsdk:"priority" json:"priority,omitempty"`
				ResourceARN *string `tfsdk:"resource_arn" json:"resourceARN,omitempty"`
			} `tfsdk:"stateful_rule_group_references" json:"statefulRuleGroupReferences,omitempty"`
			StatelessCustomActions *[]struct {
				ActionDefinition *struct {
					PublishMetricAction *struct {
						Dimensions *[]struct {
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"dimensions" json:"dimensions,omitempty"`
					} `tfsdk:"publish_metric_action" json:"publishMetricAction,omitempty"`
				} `tfsdk:"action_definition" json:"actionDefinition,omitempty"`
				ActionName *string `tfsdk:"action_name" json:"actionName,omitempty"`
			} `tfsdk:"stateless_custom_actions" json:"statelessCustomActions,omitempty"`
			StatelessDefaultActions         *[]string `tfsdk:"stateless_default_actions" json:"statelessDefaultActions,omitempty"`
			StatelessFragmentDefaultActions *[]string `tfsdk:"stateless_fragment_default_actions" json:"statelessFragmentDefaultActions,omitempty"`
			StatelessRuleGroupReferences    *[]struct {
				Priority    *int64  `tfsdk:"priority" json:"priority,omitempty"`
				ResourceARN *string `tfsdk:"resource_arn" json:"resourceARN,omitempty"`
			} `tfsdk:"stateless_rule_group_references" json:"statelessRuleGroupReferences,omitempty"`
			TlsInspectionConfigurationARN *string `tfsdk:"tls_inspection_configuration_arn" json:"tlsInspectionConfigurationARN,omitempty"`
		} `tfsdk:"firewall_policy" json:"firewallPolicy,omitempty"`
		FirewallPolicyName *string `tfsdk:"firewall_policy_name" json:"firewallPolicyName,omitempty"`
		Tags               *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest"
}

func (r *NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "FirewallPolicy is the Schema for the FirewallPolicies API",
		MarkdownDescription: "FirewallPolicy is the Schema for the FirewallPolicies API",
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
				Description:         "FirewallPolicySpec defines the desired state of FirewallPolicy.The firewall policy defines the behavior of a firewall using a collectionof stateless and stateful rule groups and other settings. You can use onefirewall policy for multiple firewalls.This, along with FirewallPolicyResponse, define the policy. You can retrieveall objects for a firewall policy by calling DescribeFirewallPolicy.",
				MarkdownDescription: "FirewallPolicySpec defines the desired state of FirewallPolicy.The firewall policy defines the behavior of a firewall using a collectionof stateless and stateful rule groups and other settings. You can use onefirewall policy for multiple firewalls.This, along with FirewallPolicyResponse, define the policy. You can retrieveall objects for a firewall policy by calling DescribeFirewallPolicy.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "A description of the firewall policy.",
						MarkdownDescription: "A description of the firewall policy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encryption_configuration": schema.SingleNestedAttribute{
						Description:         "A complex type that contains settings for encryption of your firewall policyresources.",
						MarkdownDescription: "A complex type that contains settings for encryption of your firewall policyresources.",
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

					"firewall_policy": schema.SingleNestedAttribute{
						Description:         "The rule groups and policy actions to use in the firewall policy.",
						MarkdownDescription: "The rule groups and policy actions to use in the firewall policy.",
						Attributes: map[string]schema.Attribute{
							"policy_variables": schema.SingleNestedAttribute{
								Description:         "Contains variables that you can use to override default Suricata settingsin your firewall policy.",
								MarkdownDescription: "Contains variables that you can use to override default Suricata settingsin your firewall policy.",
								Attributes: map[string]schema.Attribute{
									"rule_variables": schema.SingleNestedAttribute{
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

							"stateful_default_actions": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stateful_engine_options": schema.SingleNestedAttribute{
								Description:         "Configuration settings for the handling of the stateful rule groups in afirewall policy.",
								MarkdownDescription: "Configuration settings for the handling of the stateful rule groups in afirewall policy.",
								Attributes: map[string]schema.Attribute{
									"rule_order": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"stream_exception_policy": schema.StringAttribute{
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

							"stateful_rule_group_references": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"override": schema.SingleNestedAttribute{
											Description:         "The setting that allows the policy owner to change the behavior of the rulegroup within a policy.",
											MarkdownDescription: "The setting that allows the policy owner to change the behavior of the rulegroup within a policy.",
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
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

										"priority": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resource_arn": schema.StringAttribute{
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

							"stateless_custom_actions": schema.ListNestedAttribute{
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

							"stateless_default_actions": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stateless_fragment_default_actions": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stateless_rule_group_references": schema.ListNestedAttribute{
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

										"resource_arn": schema.StringAttribute{
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

							"tls_inspection_configuration_arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"firewall_policy_name": schema.StringAttribute{
						Description:         "The descriptive name of the firewall policy. You can't change the name ofa firewall policy after you create it.",
						MarkdownDescription: "The descriptive name of the firewall policy. You can't change the name ofa firewall policy after you create it.",
						Required:            true,
						Optional:            false,
						Computed:            false,
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networkfirewall_services_k8s_aws_firewall_policy_v1alpha1_manifest")

	var model NetworkfirewallServicesK8SAwsFirewallPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("networkfirewall.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("FirewallPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
