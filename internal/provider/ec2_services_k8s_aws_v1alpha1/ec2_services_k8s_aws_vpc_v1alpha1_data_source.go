/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ec2_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource              = &Ec2ServicesK8SAwsVPCV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &Ec2ServicesK8SAwsVPCV1Alpha1DataSource{}
)

func NewEc2ServicesK8SAwsVPCV1Alpha1DataSource() datasource.DataSource {
	return &Ec2ServicesK8SAwsVPCV1Alpha1DataSource{}
}

type Ec2ServicesK8SAwsVPCV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Ec2ServicesK8SAwsVPCV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *Ec2ServicesK8SAwsVPCV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_vpc_v1alpha1"
}

func (r *Ec2ServicesK8SAwsVPCV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "VpcSpec defines the desired state of Vpc.  Describes a VPC.",
				MarkdownDescription: "VpcSpec defines the desired state of Vpc.  Describes a VPC.",
				Attributes: map[string]schema.Attribute{
					"amazon_provided_i_pv6_cidr_block": schema.BoolAttribute{
						Description:         "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block.",
						MarkdownDescription: "Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC. You cannot specify the range of IP addresses, or the size of the CIDR block.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cidr_blocks": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_dns_hostnames": schema.BoolAttribute{
						Description:         "The attribute value. The valid values are true or false.",
						MarkdownDescription: "The attribute value. The valid values are true or false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_dns_support": schema.BoolAttribute{
						Description:         "The attribute value. The valid values are true or false.",
						MarkdownDescription: "The attribute value. The valid values are true or false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"instance_tenancy": schema.StringAttribute{
						Description:         "The tenancy options for instances launched into the VPC. For default, instances are launched with shared tenancy by default. You can launch instances with any tenancy into a shared tenancy VPC. For dedicated, instances are launched as dedicated tenancy instances by default. You can only launch instances with a tenancy of dedicated or host into a dedicated tenancy VPC.  Important: The host value cannot be used with this parameter. Use the default or dedicated values only.  Default: default",
						MarkdownDescription: "The tenancy options for instances launched into the VPC. For default, instances are launched with shared tenancy by default. You can launch instances with any tenancy into a shared tenancy VPC. For dedicated, instances are launched as dedicated tenancy instances by default. You can only launch instances with a tenancy of dedicated or host into a dedicated tenancy VPC.  Important: The host value cannot be used with this parameter. Use the default or dedicated values only.  Default: default",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv4_ipam_pool_id": schema.StringAttribute{
						Description:         "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv4_netmask_length": schema.Int64Attribute{
						Description:         "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv6_cidr_block": schema.StringAttribute{
						Description:         "The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool in the request.  To let Amazon choose the IPv6 CIDR block for you, omit this parameter.",
						MarkdownDescription: "The IPv6 CIDR block from the IPv6 address pool. You must also specify Ipv6Pool in the request.  To let Amazon choose the IPv6 CIDR block for you, omit this parameter.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv6_cidr_block_network_border_group": schema.StringAttribute{
						Description:         "The name of the location from which we advertise the IPV6 CIDR block. Use this parameter to limit the address to this location.  You must set AmazonProvidedIpv6CidrBlock to true to use this parameter.",
						MarkdownDescription: "The name of the location from which we advertise the IPV6 CIDR block. Use this parameter to limit the address to this location.  You must set AmazonProvidedIpv6CidrBlock to true to use this parameter.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv6_ipam_pool_id": schema.StringAttribute{
						Description:         "The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6 CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across Amazon Web Services Regions and accounts throughout your Amazon Web Services Organization. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The ID of an IPv6 IPAM pool which will be used to allocate this VPC an IPv6 CIDR. IPAM is a VPC feature that you can use to automate your IP address management workflows including assigning, tracking, troubleshooting, and auditing IP addresses across Amazon Web Services Regions and accounts throughout your Amazon Web Services Organization. For more information, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv6_netmask_length": schema.Int64Attribute{
						Description:         "The netmask length of the IPv6 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						MarkdownDescription: "The netmask length of the IPv6 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool. For more information about IPAM, see What is IPAM? (https://docs.aws.amazon.com/vpc/latest/ipam/what-is-it-ipam.html) in the Amazon VPC IPAM User Guide.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv6_pool": schema.StringAttribute{
						Description:         "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.",
						MarkdownDescription: "The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *Ec2ServicesK8SAwsVPCV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *Ec2ServicesK8SAwsVPCV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_ec2_services_k8s_aws_vpc_v1alpha1")

	var data Ec2ServicesK8SAwsVPCV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "VPC"}).
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

	var readResponse Ec2ServicesK8SAwsVPCV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("VPC")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
