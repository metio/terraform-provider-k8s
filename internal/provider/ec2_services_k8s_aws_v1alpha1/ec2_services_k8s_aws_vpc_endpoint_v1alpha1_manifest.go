/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ec2_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Ec2ServicesK8SAwsVpcendpointV1Alpha1Manifest{}
)

func NewEc2ServicesK8SAwsVpcendpointV1Alpha1Manifest() datasource.DataSource {
	return &Ec2ServicesK8SAwsVpcendpointV1Alpha1Manifest{}
}

type Ec2ServicesK8SAwsVpcendpointV1Alpha1Manifest struct{}

type Ec2ServicesK8SAwsVpcendpointV1Alpha1ManifestData struct {
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
		DnsOptions *struct {
			DnsRecordIPType *string `tfsdk:"dns_record_ip_type" json:"dnsRecordIPType,omitempty"`
		} `tfsdk:"dns_options" json:"dnsOptions,omitempty"`
		IpAddressType     *string   `tfsdk:"ip_address_type" json:"ipAddressType,omitempty"`
		PolicyDocument    *string   `tfsdk:"policy_document" json:"policyDocument,omitempty"`
		PrivateDNSEnabled *bool     `tfsdk:"private_dns_enabled" json:"privateDNSEnabled,omitempty"`
		RouteTableIDs     *[]string `tfsdk:"route_table_i_ds" json:"routeTableIDs,omitempty"`
		RouteTableRefs    *[]struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"route_table_refs" json:"routeTableRefs,omitempty"`
		SecurityGroupIDs  *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
		SecurityGroupRefs *[]struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"security_group_refs" json:"securityGroupRefs,omitempty"`
		ServiceName *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
		SubnetIDs   *[]string `tfsdk:"subnet_i_ds" json:"subnetIDs,omitempty"`
		SubnetRefs  *[]struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"subnet_refs" json:"subnetRefs,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcEndpointType *string `tfsdk:"vpc_endpoint_type" json:"vpcEndpointType,omitempty"`
		VpcID           *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
		VpcRef          *struct {
			From *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"vpc_ref" json:"vpcRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Ec2ServicesK8SAwsVpcendpointV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest"
}

func (r *Ec2ServicesK8SAwsVpcendpointV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VPCEndpoint is the Schema for the VPCEndpoints API",
		MarkdownDescription: "VPCEndpoint is the Schema for the VPCEndpoints API",
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
				Description:         "VpcEndpointSpec defines the desired state of VpcEndpoint.Describes a VPC endpoint.",
				MarkdownDescription: "VpcEndpointSpec defines the desired state of VpcEndpoint.Describes a VPC endpoint.",
				Attributes: map[string]schema.Attribute{
					"dns_options": schema.SingleNestedAttribute{
						Description:         "The DNS options for the endpoint.",
						MarkdownDescription: "The DNS options for the endpoint.",
						Attributes: map[string]schema.Attribute{
							"dns_record_ip_type": schema.StringAttribute{
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

					"ip_address_type": schema.StringAttribute{
						Description:         "The IP address type for the endpoint.",
						MarkdownDescription: "The IP address type for the endpoint.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy_document": schema.StringAttribute{
						Description:         "(Interface and gateway endpoints) A policy to attach to the endpoint thatcontrols access to the service. The policy must be in valid JSON format.If this parameter is not specified, we attach a default policy that allowsfull access to the service.",
						MarkdownDescription: "(Interface and gateway endpoints) A policy to attach to the endpoint thatcontrols access to the service. The policy must be in valid JSON format.If this parameter is not specified, we attach a default policy that allowsfull access to the service.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"private_dns_enabled": schema.BoolAttribute{
						Description:         "(Interface endpoint) Indicates whether to associate a private hosted zonewith the specified VPC. The private hosted zone contains a record set forthe default public DNS name for the service for the Region (for example,kinesis.us-east-1.amazonaws.com), which resolves to the private IP addressesof the endpoint network interfaces in the VPC. This enables you to make requeststo the default public DNS name for the service instead of the public DNSnames that are automatically generated by the VPC endpoint service.To use a private hosted zone, you must set the following VPC attributes totrue: enableDnsHostnames and enableDnsSupport. Use ModifyVpcAttribute toset the VPC attributes.Default: true",
						MarkdownDescription: "(Interface endpoint) Indicates whether to associate a private hosted zonewith the specified VPC. The private hosted zone contains a record set forthe default public DNS name for the service for the Region (for example,kinesis.us-east-1.amazonaws.com), which resolves to the private IP addressesof the endpoint network interfaces in the VPC. This enables you to make requeststo the default public DNS name for the service instead of the public DNSnames that are automatically generated by the VPC endpoint service.To use a private hosted zone, you must set the following VPC attributes totrue: enableDnsHostnames and enableDnsSupport. Use ModifyVpcAttribute toset the VPC attributes.Default: true",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_table_i_ds": schema.ListAttribute{
						Description:         "(Gateway endpoint) One or more route table IDs.",
						MarkdownDescription: "(Gateway endpoint) One or more route table IDs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_table_refs": schema.ListNestedAttribute{
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

										"namespace": schema.StringAttribute{
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

					"security_group_i_ds": schema.ListAttribute{
						Description:         "(Interface endpoint) The ID of one or more security groups to associate withthe endpoint network interface.",
						MarkdownDescription: "(Interface endpoint) The ID of one or more security groups to associate withthe endpoint network interface.",
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

										"namespace": schema.StringAttribute{
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

					"service_name": schema.StringAttribute{
						Description:         "The service name. To get a list of available services, use the DescribeVpcEndpointServicesrequest, or get the name from the service provider.",
						MarkdownDescription: "The service name. To get a list of available services, use the DescribeVpcEndpointServicesrequest, or get the name from the service provider.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"subnet_i_ds": schema.ListAttribute{
						Description:         "(Interface and Gateway Load Balancer endpoints) The ID of one or more subnetsin which to create an endpoint network interface. For a Gateway Load Balancerendpoint, you can specify one subnet only.",
						MarkdownDescription: "(Interface and Gateway Load Balancer endpoints) The ID of one or more subnetsin which to create an endpoint network interface. For a Gateway Load Balancerendpoint, you can specify one subnet only.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subnet_refs": schema.ListNestedAttribute{
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

										"namespace": schema.StringAttribute{
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
						Description:         "The tags. The value parameter is required, but if you don't want the tagto have a value, specify the parameter with no value, and we set the valueto an empty string.",
						MarkdownDescription: "The tags. The value parameter is required, but if you don't want the tagto have a value, specify the parameter with no value, and we set the valueto an empty string.",
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

					"vpc_endpoint_type": schema.StringAttribute{
						Description:         "The type of endpoint.Default: Gateway",
						MarkdownDescription: "The type of endpoint.Default: Gateway",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vpc_id": schema.StringAttribute{
						Description:         "The ID of the VPC in which the endpoint will be used.",
						MarkdownDescription: "The ID of the VPC in which the endpoint will be used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vpc_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
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

									"namespace": schema.StringAttribute{
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Ec2ServicesK8SAwsVpcendpointV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1_manifest")

	var model Ec2ServicesK8SAwsVpcendpointV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("VPCEndpoint")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
