/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package route53resolver_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Route53ResolverServicesK8SAwsResolverRuleV1Alpha1Manifest{}
)

func NewRoute53ResolverServicesK8SAwsResolverRuleV1Alpha1Manifest() datasource.DataSource {
	return &Route53ResolverServicesK8SAwsResolverRuleV1Alpha1Manifest{}
}

type Route53ResolverServicesK8SAwsResolverRuleV1Alpha1Manifest struct{}

type Route53ResolverServicesK8SAwsResolverRuleV1Alpha1ManifestData struct {
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
		Associations *[]struct {
			Id             *string `tfsdk:"id" json:"id,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			ResolverRuleID *string `tfsdk:"resolver_rule_id" json:"resolverRuleID,omitempty"`
			Status         *string `tfsdk:"status" json:"status,omitempty"`
			StatusMessage  *string `tfsdk:"status_message" json:"statusMessage,omitempty"`
			VpcID          *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
		} `tfsdk:"associations" json:"associations,omitempty"`
		DomainName         *string `tfsdk:"domain_name" json:"domainName,omitempty"`
		Name               *string `tfsdk:"name" json:"name,omitempty"`
		ResolverEndpointID *string `tfsdk:"resolver_endpoint_id" json:"resolverEndpointID,omitempty"`
		RuleType           *string `tfsdk:"rule_type" json:"ruleType,omitempty"`
		Tags               *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		TargetIPs *[]struct {
			Ip   *string `tfsdk:"ip" json:"ip,omitempty"`
			Ipv6 *string `tfsdk:"ipv6" json:"ipv6,omitempty"`
			Port *int64  `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"target_i_ps" json:"targetIPs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Route53ResolverServicesK8SAwsResolverRuleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest"
}

func (r *Route53ResolverServicesK8SAwsResolverRuleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResolverRule is the Schema for the ResolverRules API",
		MarkdownDescription: "ResolverRule is the Schema for the ResolverRules API",
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
				Description:         "ResolverRuleSpec defines the desired state of ResolverRule.For queries that originate in your VPC, detailed information about a Resolverrule, which specifies how to route DNS queries out of the VPC. The ResolverRuleparameter appears in the response to a CreateResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html),DeleteResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverRule.html),GetResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverRule.html),ListResolverRules (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRules.html),or UpdateResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverRule.html)request.",
				MarkdownDescription: "ResolverRuleSpec defines the desired state of ResolverRule.For queries that originate in your VPC, detailed information about a Resolverrule, which specifies how to route DNS queries out of the VPC. The ResolverRuleparameter appears in the response to a CreateResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html),DeleteResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverRule.html),GetResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverRule.html),ListResolverRules (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ListResolverRules.html),or UpdateResolverRule (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverRule.html)request.",
				Attributes: map[string]schema.Attribute{
					"associations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resolver_rule_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"status": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"status_message": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"vpc_id": schema.StringAttribute{
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

					"domain_name": schema.StringAttribute{
						Description:         "DNS queries for this domain name are forwarded to the IP addresses that youspecify in TargetIps. If a query matches multiple Resolver rules (example.comand www.example.com), outbound DNS queries are routed using the Resolverrule that contains the most specific domain name (www.example.com).",
						MarkdownDescription: "DNS queries for this domain name are forwarded to the IP addresses that youspecify in TargetIps. If a query matches multiple Resolver rules (example.comand www.example.com), outbound DNS queries are routed using the Resolverrule that contains the most specific domain name (www.example.com).",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "A friendly name that lets you easily find a rule in the Resolver dashboardin the Route 53 console.",
						MarkdownDescription: "A friendly name that lets you easily find a rule in the Resolver dashboardin the Route 53 console.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resolver_endpoint_id": schema.StringAttribute{
						Description:         "The ID of the outbound Resolver endpoint that you want to use to route DNSqueries to the IP addresses that you specify in TargetIps.",
						MarkdownDescription: "The ID of the outbound Resolver endpoint that you want to use to route DNSqueries to the IP addresses that you specify in TargetIps.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rule_type": schema.StringAttribute{
						Description:         "When you want to forward DNS queries for specified domain name to resolverson your network, specify FORWARD.When you have a forwarding rule to forward DNS queries for a domain to yournetwork and you want Resolver to process queries for a subdomain of thatdomain, specify SYSTEM.For example, to forward DNS queries for example.com to resolvers on yournetwork, you create a rule and specify FORWARD for RuleType. To then haveResolver process queries for apex.example.com, you create a rule and specifySYSTEM for RuleType.Currently, only Resolver can create rules that have a value of RECURSIVEfor RuleType.",
						MarkdownDescription: "When you want to forward DNS queries for specified domain name to resolverson your network, specify FORWARD.When you have a forwarding rule to forward DNS queries for a domain to yournetwork and you want Resolver to process queries for a subdomain of thatdomain, specify SYSTEM.For example, to forward DNS queries for example.com to resolvers on yournetwork, you create a rule and specify FORWARD for RuleType. To then haveResolver process queries for apex.example.com, you create a rule and specifySYSTEM for RuleType.Currently, only Resolver can create rules that have a value of RECURSIVEfor RuleType.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "A list of the tag keys and values that you want to associate with the endpoint.",
						MarkdownDescription: "A list of the tag keys and values that you want to associate with the endpoint.",
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

					"target_i_ps": schema.ListNestedAttribute{
						Description:         "The IPs that you want Resolver to forward DNS queries to. You can specifyonly IPv4 addresses. Separate IP addresses with a space.TargetIps is available only when the value of Rule type is FORWARD.",
						MarkdownDescription: "The IPs that you want Resolver to forward DNS queries to. You can specifyonly IPv4 addresses. Separate IP addresses with a space.TargetIps is available only when the value of Rule type is FORWARD.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ipv6": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
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

func (r *Route53ResolverServicesK8SAwsResolverRuleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest")

	var model Route53ResolverServicesK8SAwsResolverRuleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("route53resolver.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("ResolverRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
