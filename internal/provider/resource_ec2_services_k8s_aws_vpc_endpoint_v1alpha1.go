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

type Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource)(nil)
)

type Ec2ServicesK8SAwsVPCEndpointV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type Ec2ServicesK8SAwsVPCEndpointV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		DnsOptions *struct {
			DnsRecordIPType *string `tfsdk:"dns_record_ip_type" yaml:"dnsRecordIPType,omitempty"`
		} `tfsdk:"dns_options" yaml:"dnsOptions,omitempty"`

		IpAddressType *string `tfsdk:"ip_address_type" yaml:"ipAddressType,omitempty"`

		PolicyDocument *string `tfsdk:"policy_document" yaml:"policyDocument,omitempty"`

		PrivateDNSEnabled *bool `tfsdk:"private_dns_enabled" yaml:"privateDNSEnabled,omitempty"`

		RouteTableIDs *[]string `tfsdk:"route_table_i_ds" yaml:"routeTableIDs,omitempty"`

		RouteTableRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"route_table_refs" yaml:"routeTableRefs,omitempty"`

		SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

		SecurityGroupRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"security_group_refs" yaml:"securityGroupRefs,omitempty"`

		ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`

		SubnetIDs *[]string `tfsdk:"subnet_i_ds" yaml:"subnetIDs,omitempty"`

		SubnetRefs *[]struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"subnet_refs" yaml:"subnetRefs,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		VpcEndpointType *string `tfsdk:"vpc_endpoint_type" yaml:"vpcEndpointType,omitempty"`

		VpcID *string `tfsdk:"vpc_id" yaml:"vpcID,omitempty"`

		VpcRef *struct {
			From *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"from" yaml:"from,omitempty"`
		} `tfsdk:"vpc_ref" yaml:"vpcRef,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewEc2ServicesK8SAwsVPCEndpointV1Alpha1Resource() resource.Resource {
	return &Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource{}
}

func (r *Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ec2_services_k8s_aws_vpc_endpoint_v1alpha1"
}

func (r *Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "VPCEndpoint is the Schema for the VPCEndpoints API",
		MarkdownDescription: "VPCEndpoint is the Schema for the VPCEndpoints API",
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
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "VpcEndpointSpec defines the desired state of VpcEndpoint.  Describes a VPC endpoint.",
				MarkdownDescription: "VpcEndpointSpec defines the desired state of VpcEndpoint.  Describes a VPC endpoint.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"dns_options": {
						Description:         "The DNS options for the endpoint.",
						MarkdownDescription: "The DNS options for the endpoint.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dns_record_ip_type": {
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

					"ip_address_type": {
						Description:         "The IP address type for the endpoint.",
						MarkdownDescription: "The IP address type for the endpoint.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"policy_document": {
						Description:         "(Interface and gateway endpoints) A policy to attach to the endpoint that controls access to the service. The policy must be in valid JSON format. If this parameter is not specified, we attach a default policy that allows full access to the service.",
						MarkdownDescription: "(Interface and gateway endpoints) A policy to attach to the endpoint that controls access to the service. The policy must be in valid JSON format. If this parameter is not specified, we attach a default policy that allows full access to the service.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"private_dns_enabled": {
						Description:         "(Interface endpoint) Indicates whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, kinesis.us-east-1.amazonaws.com), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.  To use a private hosted zone, you must set the following VPC attributes to true: enableDnsHostnames and enableDnsSupport. Use ModifyVpcAttribute to set the VPC attributes.  Default: true",
						MarkdownDescription: "(Interface endpoint) Indicates whether to associate a private hosted zone with the specified VPC. The private hosted zone contains a record set for the default public DNS name for the service for the Region (for example, kinesis.us-east-1.amazonaws.com), which resolves to the private IP addresses of the endpoint network interfaces in the VPC. This enables you to make requests to the default public DNS name for the service instead of the public DNS names that are automatically generated by the VPC endpoint service.  To use a private hosted zone, you must set the following VPC attributes to true: enableDnsHostnames and enableDnsSupport. Use ModifyVpcAttribute to set the VPC attributes.  Default: true",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_table_i_ds": {
						Description:         "(Gateway endpoint) One or more route table IDs.",
						MarkdownDescription: "(Gateway endpoint) One or more route table IDs.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_table_refs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"security_group_i_ds": {
						Description:         "(Interface endpoint) The ID of one or more security groups to associate with the endpoint network interface.",
						MarkdownDescription: "(Interface endpoint) The ID of one or more security groups to associate with the endpoint network interface.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"security_group_refs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"service_name": {
						Description:         "The service name. To get a list of available services, use the DescribeVpcEndpointServices request, or get the name from the service provider.",
						MarkdownDescription: "The service name. To get a list of available services, use the DescribeVpcEndpointServices request, or get the name from the service provider.",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"subnet_i_ds": {
						Description:         "(Interface and Gateway Load Balancer endpoints) The ID of one or more subnets in which to create an endpoint network interface. For a Gateway Load Balancer endpoint, you can specify one subnet only.",
						MarkdownDescription: "(Interface and Gateway Load Balancer endpoints) The ID of one or more subnets in which to create an endpoint network interface. For a Gateway Load Balancer endpoint, you can specify one subnet only.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"subnet_refs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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

					"tags": {
						Description:         "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
						MarkdownDescription: "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
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

					"vpc_endpoint_type": {
						Description:         "The type of endpoint.  Default: Gateway",
						MarkdownDescription: "The type of endpoint.  Default: Gateway",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vpc_id": {
						Description:         "The ID of the VPC in which the endpoint will be used.",
						MarkdownDescription: "The ID of the VPC in which the endpoint will be used.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vpc_ref": {
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef: from: name: my-api",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"from": {
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
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
		},
	}, nil
}

func (r *Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1")

	var state Ec2ServicesK8SAwsVPCEndpointV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Ec2ServicesK8SAwsVPCEndpointV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ec2.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("VPCEndpoint")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1")

	var state Ec2ServicesK8SAwsVPCEndpointV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Ec2ServicesK8SAwsVPCEndpointV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ec2.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("VPCEndpoint")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *Ec2ServicesK8SAwsVPCEndpointV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ec2_services_k8s_aws_vpc_endpoint_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
