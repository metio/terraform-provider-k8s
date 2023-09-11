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
	_ datasource.DataSource              = &Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource{}
)

func NewEc2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource() datasource.DataSource {
	return &Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource{}
}

type Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSourceData struct {
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
		Address               *string `tfsdk:"address" json:"address,omitempty"`
		CustomerOwnedIPv4Pool *string `tfsdk:"customer_owned_i_pv4_pool" json:"customerOwnedIPv4Pool,omitempty"`
		NetworkBorderGroup    *string `tfsdk:"network_border_group" json:"networkBorderGroup,omitempty"`
		PublicIPv4Pool        *string `tfsdk:"public_i_pv4_pool" json:"publicIPv4Pool,omitempty"`
		Tags                  *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_elastic_ip_address_v1alpha1"
}

func (r *Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ElasticIPAddress is the Schema for the ElasticIPAddresses API",
		MarkdownDescription: "ElasticIPAddress is the Schema for the ElasticIPAddresses API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "ElasticIPAddressSpec defines the desired state of ElasticIPAddress.",
				MarkdownDescription: "ElasticIPAddressSpec defines the desired state of ElasticIPAddress.",
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description:         "[EC2-VPC] The Elastic IP address to recover or an IPv4 address from an address pool.",
						MarkdownDescription: "[EC2-VPC] The Elastic IP address to recover or an IPv4 address from an address pool.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"customer_owned_i_pv4_pool": schema.StringAttribute{
						Description:         "The ID of a customer-owned address pool. Use this parameter to let Amazon EC2 select an address from the address pool. Alternatively, specify a specific address from the address pool.",
						MarkdownDescription: "The ID of a customer-owned address pool. Use this parameter to let Amazon EC2 select an address from the address pool. Alternatively, specify a specific address from the address pool.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"network_border_group": schema.StringAttribute{
						Description:         "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses. Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.  Use DescribeAvailabilityZones (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.  You cannot use a network border group with EC2 Classic. If you attempt this operation on EC2 Classic, you receive an InvalidParameterCombination error.",
						MarkdownDescription: "A unique set of Availability Zones, Local Zones, or Wavelength Zones from which Amazon Web Services advertises IP addresses. Use this parameter to limit the IP address to this location. IP addresses cannot move between network border groups.  Use DescribeAvailabilityZones (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html) to view the network border groups.  You cannot use a network border group with EC2 Classic. If you attempt this operation on EC2 Classic, you receive an InvalidParameterCombination error.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"public_i_pv4_pool": schema.StringAttribute{
						Description:         "The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool. To specify a specific address from the address pool, use the Address parameter instead.",
						MarkdownDescription: "The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an address from the address pool. To specify a specific address from the address pool, use the Address parameter instead.",
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

func (r *Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_ec2_services_k8s_aws_elastic_ip_address_v1alpha1")

	var data Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "elasticipaddresses"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse Ec2ServicesK8SAwsElasticIpaddressV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("ElasticIPAddress")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
