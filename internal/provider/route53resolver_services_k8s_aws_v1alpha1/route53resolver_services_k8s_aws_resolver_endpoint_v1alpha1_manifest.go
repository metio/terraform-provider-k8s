/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package route53resolver_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1Manifest{}
)

func NewRoute53ResolverServicesK8SAwsResolverEndpointV1Alpha1Manifest() datasource.DataSource {
	return &Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1Manifest{}
}

type Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1Manifest struct{}

type Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1ManifestData struct {
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
		Direction   *string `tfsdk:"direction" json:"direction,omitempty"`
		IpAddresses *[]struct {
			Ip        *string `tfsdk:"ip" json:"ip,omitempty"`
			Ipv6      *string `tfsdk:"ipv6" json:"ipv6,omitempty"`
			SubnetID  *string `tfsdk:"subnet_id" json:"subnetID,omitempty"`
			SubnetRef *struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"subnet_ref" json:"subnetRef,omitempty"`
		} `tfsdk:"ip_addresses" json:"ipAddresses,omitempty"`
		Name                 *string   `tfsdk:"name" json:"name,omitempty"`
		ResolverEndpointType *string   `tfsdk:"resolver_endpoint_type" json:"resolverEndpointType,omitempty"`
		SecurityGroupIDs     *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
		SecurityGroupRefs    *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"security_group_refs" json:"securityGroupRefs,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest"
}

func (r *Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResolverEndpoint is the Schema for the ResolverEndpoints API",
		MarkdownDescription: "ResolverEndpoint is the Schema for the ResolverEndpoints API",
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
				Description:         "ResolverEndpointSpec defines the desired state of ResolverEndpoint.In the response to a CreateResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html),DeleteResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverEndpoint.html),GetResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html),Updates the name, or ResolverEndpointType for an endpoint, or UpdateResolverEndpoint(https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverEndpoint.html)request, a complex type that contains settings for an existing inbound oroutbound Resolver endpoint.",
				MarkdownDescription: "ResolverEndpointSpec defines the desired state of ResolverEndpoint.In the response to a CreateResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverEndpoint.html),DeleteResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DeleteResolverEndpoint.html),GetResolverEndpoint (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_GetResolverEndpoint.html),Updates the name, or ResolverEndpointType for an endpoint, or UpdateResolverEndpoint(https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_UpdateResolverEndpoint.html)request, a complex type that contains settings for an existing inbound oroutbound Resolver endpoint.",
				Attributes: map[string]schema.Attribute{
					"direction": schema.StringAttribute{
						Description:         "Specify the applicable value:   * INBOUND: Resolver forwards DNS queries to the DNS service for a VPC   from your network   * OUTBOUND: Resolver forwards DNS queries from the DNS service for a VPC   to your network",
						MarkdownDescription: "Specify the applicable value:   * INBOUND: Resolver forwards DNS queries to the DNS service for a VPC   from your network   * OUTBOUND: Resolver forwards DNS queries from the DNS service for a VPC   to your network",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"ip_addresses": schema.ListNestedAttribute{
						Description:         "The subnets and IP addresses in your VPC that DNS queries originate from(for outbound endpoints) or that you forward DNS queries to (for inboundendpoints). The subnet ID uniquely identifies a VPC.",
						MarkdownDescription: "The subnets and IP addresses in your VPC that DNS queries originate from(for outbound endpoints) or that you forward DNS queries to (for inboundendpoints). The subnet ID uniquely identifies a VPC.",
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

								"subnet_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subnet_ref": schema.SingleNestedAttribute{
									Description:         "Reference field for SubnetID",
									MarkdownDescription: "Reference field for SubnetID",
									Attributes: map[string]schema.Attribute{
										"from": schema.SingleNestedAttribute{
											Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
											MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "A friendly name that lets you easily find a configuration in the Resolverdashboard in the Route 53 console.",
						MarkdownDescription: "A friendly name that lets you easily find a configuration in the Resolverdashboard in the Route 53 console.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resolver_endpoint_type": schema.StringAttribute{
						Description:         "For the endpoint type you can choose either IPv4, IPv6. or dual-stack. Adual-stack endpoint means that it will resolve via both IPv4 and IPv6. Thisendpoint type is applied to all IP addresses.",
						MarkdownDescription: "For the endpoint type you can choose either IPv4, IPv6. or dual-stack. Adual-stack endpoint means that it will resolve via both IPv4 and IPv6. Thisendpoint type is applied to all IP addresses.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_group_i_ds": schema.ListAttribute{
						Description:         "The ID of one or more security groups that you want to use to control accessto this VPC. The security group that you specify must include one or moreinbound rules (for inbound Resolver endpoints) or outbound rules (for outboundResolver endpoints). Inbound and outbound rules must allow TCP and UDP access.For inbound access, open port 53. For outbound access, open the port thatyou're using for DNS queries on your network.",
						MarkdownDescription: "The ID of one or more security groups that you want to use to control accessto this VPC. The security group that you specify must include one or moreinbound rules (for inbound Resolver endpoints) or outbound rules (for outboundResolver endpoints). Inbound and outbound rules must allow TCP and UDP access.For inbound access, open port 53. For outbound access, open the port thatyou're using for DNS queries on your network.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_group_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.SingleNestedAttribute{
									Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_route53resolver_services_k8s_aws_resolver_endpoint_v1alpha1_manifest")

	var model Route53ResolverServicesK8SAwsResolverEndpointV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("route53resolver.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("ResolverEndpoint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
