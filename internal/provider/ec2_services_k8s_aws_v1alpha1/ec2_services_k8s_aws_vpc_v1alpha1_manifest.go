/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ec2_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &Ec2ServicesK8SAwsVpcV1Alpha1Manifest{}
)

func NewEc2ServicesK8SAwsVpcV1Alpha1Manifest() datasource.DataSource {
	return &Ec2ServicesK8SAwsVpcV1Alpha1Manifest{}
}

type Ec2ServicesK8SAwsVpcV1Alpha1Manifest struct{}

type Ec2ServicesK8SAwsVpcV1Alpha1ManifestData struct {
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
		AmazonProvidedIPv6CIDRBlock     *bool     `tfsdk:"amazon_provided_i_pv6_cidr_block" json:"amazonProvidedIPv6CIDRBlock,omitempty"`
		CidrBlocks                      *[]string `tfsdk:"cidr_blocks" json:"cidrBlocks,omitempty"`
		EnableDNSHostnames              *bool     `tfsdk:"enable_dns_hostnames" json:"enableDNSHostnames,omitempty"`
		EnableDNSSupport                *bool     `tfsdk:"enable_dns_support" json:"enableDNSSupport,omitempty"`
		InstanceTenancy                 *string   `tfsdk:"instance_tenancy" json:"instanceTenancy,omitempty"`
		Ipv4IPAMPoolID                  *string   `tfsdk:"ipv4_ipam_pool_id" json:"ipv4IPAMPoolID,omitempty"`
		Ipv4NetmaskLength               *int64    `tfsdk:"ipv4_netmask_length" json:"ipv4NetmaskLength,omitempty"`
		Ipv6CIDRBlock                   *string   `tfsdk:"ipv6_cidr_block" json:"ipv6CIDRBlock,omitempty"`
		Ipv6CIDRBlockNetworkBorderGroup *string   `tfsdk:"ipv6_cidr_block_network_border_group" json:"ipv6CIDRBlockNetworkBorderGroup,omitempty"`
		Ipv6IPAMPoolID                  *string   `tfsdk:"ipv6_ipam_pool_id" json:"ipv6IPAMPoolID,omitempty"`
		Ipv6NetmaskLength               *int64    `tfsdk:"ipv6_netmask_length" json:"ipv6NetmaskLength,omitempty"`
		Ipv6Pool                        *string   `tfsdk:"ipv6_pool" json:"ipv6Pool,omitempty"`
		Tags                            *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_vpc_v1alpha1_manifest"
}

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VPC is the Schema for the VPCS API",
		MarkdownDescription: "VPC is the Schema for the VPCS API",
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
				Description:         "VpcSpec defines the desired state of Vpc.  Describes a VPC.",
				MarkdownDescription: "VpcSpec defines the desired state of Vpc.  Describes a VPC.",
				Attributes: map[string]schema.Attribute{
					"amazon_provided_i_pv6_cidr_block": schema.BoolAttribute{
						Description:         "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block.",
						MarkdownDescription: "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cidr_blocks": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"enable_dns_hostnames": schema.BoolAttribute{
						Description:         "The attribute value. The valid values are true or false.",
						MarkdownDescription: "The attribute value. The valid values are true or false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_dns_support": schema.BoolAttribute{
						Description:         "The attribute value. The valid values are true or false.",
						MarkdownDescription: "The attribute value. The valid values are true or false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_tenancy": schema.StringAttribute{
						Description:         "The tenancy options for instances launched into the VPC. For default, instances are launched with shared tenancy by default. You can launch instances with any tenancy into a shared tenancy VPC. For dedicated, instances are launched as dedicated tenancy instances by default. You can only launch instances with a tenancy of dedicated or host into a dedicated tenancy VPC.  Important: The host value cannot be used with this parameter. Use the default or dedicated values only.  Default: default",
						MarkdownDescription: "The tenancy options for instances launched into the VPC. For default, instances are launched with shared tenancy by default. You can launch instances with any tenancy into a shared tenancy VPC. For dedicated, instances are launched as dedicated tenancy instances by default. You can only launch instances with a tenancy of dedicated or host into a dedicated tenancy VPC.  Important: The host value cannot be used with this parameter. Use the default or dedicated values only.  Default: default",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv4_ipam_pool_id": schema.StringAttribute{
						Description:         "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv4_netmask_length": schema.Int64Attribute{
						Description:         "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_cidr_block": schema.StringAttribute{
						Description:         "The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool in the request.  To let Amazon choose the IPv6 CIDR block for you, omit this parameter.",
						MarkdownDescription: "The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool in the request.  To let Amazon choose the IPv6 CIDR block for you, omit this parameter.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_cidr_block_network_border_group": schema.StringAttribute{
						Description:         "The name of the location from which we advertise the IPV6 CIDR block. Use this parameter to limit the address to this location.  You must set AmazonProvidedIpv6CidrBlock to true to use this parameter.",
						MarkdownDescription: "The name of the location from which we advertise the IPV6 CIDR block. Use this parameter to limit the address to this location.  You must set AmazonProvidedIpv6CidrBlock to true to use this parameter.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_ipam_pool_id": schema.StringAttribute{
						Description:         "The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6 CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across Amazon Web Services Regions and accounts throughout your Amazon Web Services Organization. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6 CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across Amazon Web Services Regions and accounts throughout your Amazon Web Services Organization. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_netmask_length": schema.Int64Attribute{
						Description:         "The netmask length of the IPv6 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The netmask length of the IPv6 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_pool": schema.StringAttribute{
						Description:         "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.",
						MarkdownDescription: "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
						MarkdownDescription: "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
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

func (r *Ec2ServicesK8SAwsVpcV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ec2_services_k8s_aws_vpc_v1alpha1_manifest")

	var model Ec2ServicesK8SAwsVpcV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("VPC")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
