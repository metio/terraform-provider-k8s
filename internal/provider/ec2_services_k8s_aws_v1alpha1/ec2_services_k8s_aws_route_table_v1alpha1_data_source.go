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
	_ datasource.DataSource              = &Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource{}
)

func NewEc2ServicesK8SAwsRouteTableV1Alpha1DataSource() datasource.DataSource {
	return &Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource{}
}

type Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Ec2ServicesK8SAwsRouteTableV1Alpha1DataSourceData struct {
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
		Routes *[]struct {
			CarrierGatewayID            *string `tfsdk:"carrier_gateway_id" json:"carrierGatewayID,omitempty"`
			CoreNetworkARN              *string `tfsdk:"core_network_arn" json:"coreNetworkARN,omitempty"`
			DestinationCIDRBlock        *string `tfsdk:"destination_cidr_block" json:"destinationCIDRBlock,omitempty"`
			DestinationIPv6CIDRBlock    *string `tfsdk:"destination_i_pv6_cidr_block" json:"destinationIPv6CIDRBlock,omitempty"`
			DestinationPrefixListID     *string `tfsdk:"destination_prefix_list_id" json:"destinationPrefixListID,omitempty"`
			EgressOnlyInternetGatewayID *string `tfsdk:"egress_only_internet_gateway_id" json:"egressOnlyInternetGatewayID,omitempty"`
			GatewayID                   *string `tfsdk:"gateway_id" json:"gatewayID,omitempty"`
			GatewayRef                  *struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"gateway_ref" json:"gatewayRef,omitempty"`
			InstanceID     *string `tfsdk:"instance_id" json:"instanceID,omitempty"`
			LocalGatewayID *string `tfsdk:"local_gateway_id" json:"localGatewayID,omitempty"`
			NatGatewayID   *string `tfsdk:"nat_gateway_id" json:"natGatewayID,omitempty"`
			NatGatewayRef  *struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"nat_gateway_ref" json:"natGatewayRef,omitempty"`
			NetworkInterfaceID *string `tfsdk:"network_interface_id" json:"networkInterfaceID,omitempty"`
			TransitGatewayID   *string `tfsdk:"transit_gateway_id" json:"transitGatewayID,omitempty"`
			TransitGatewayRef  *struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"transit_gateway_ref" json:"transitGatewayRef,omitempty"`
			VpcEndpointID  *string `tfsdk:"vpc_endpoint_id" json:"vpcEndpointID,omitempty"`
			VpcEndpointRef *struct {
				From *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
			} `tfsdk:"vpc_endpoint_ref" json:"vpcEndpointRef,omitempty"`
			VpcPeeringConnectionID *string `tfsdk:"vpc_peering_connection_id" json:"vpcPeeringConnectionID,omitempty"`
		} `tfsdk:"routes" json:"routes,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcID  *string `tfsdk:"vpc_id" json:"vpcID,omitempty"`
		VpcRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"vpc_ref" json:"vpcRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_route_table_v1alpha1"
}

func (r *Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RouteTable is the Schema for the RouteTables API",
		MarkdownDescription: "RouteTable is the Schema for the RouteTables API",
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
				Description:         "RouteTableSpec defines the desired state of RouteTable.  Describes a route table.",
				MarkdownDescription: "RouteTableSpec defines the desired state of RouteTable.  Describes a route table.",
				Attributes: map[string]schema.Attribute{
					"routes": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"carrier_gateway_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"core_network_arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"destination_cidr_block": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"destination_i_pv6_cidr_block": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"destination_prefix_list_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"egress_only_internet_gateway_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"gateway_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"gateway_ref": schema.SingleNestedAttribute{
									Description:         "Reference field for GatewayID",
									MarkdownDescription: "Reference field for GatewayID",
									Attributes: map[string]schema.Attribute{
										"from": schema.SingleNestedAttribute{
											Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"instance_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"local_gateway_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"nat_gateway_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"nat_gateway_ref": schema.SingleNestedAttribute{
									Description:         "Reference field for NATGatewayID",
									MarkdownDescription: "Reference field for NATGatewayID",
									Attributes: map[string]schema.Attribute{
										"from": schema.SingleNestedAttribute{
											Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"network_interface_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"transit_gateway_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"transit_gateway_ref": schema.SingleNestedAttribute{
									Description:         "Reference field for TransitGatewayID",
									MarkdownDescription: "Reference field for TransitGatewayID",
									Attributes: map[string]schema.Attribute{
										"from": schema.SingleNestedAttribute{
											Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"vpc_endpoint_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"vpc_endpoint_ref": schema.SingleNestedAttribute{
									Description:         "Reference field for VPCEndpointID",
									MarkdownDescription: "Reference field for VPCEndpointID",
									Attributes: map[string]schema.Attribute{
										"from": schema.SingleNestedAttribute{
											Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

								"vpc_peering_connection_id": schema.StringAttribute{
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

					"vpc_id": schema.StringAttribute{
						Description:         "The ID of the VPC.",
						MarkdownDescription: "The ID of the VPC.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"vpc_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReference type to provide more user friendly syntax for references using 'from' field Ex: APIIDRef:  from: name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
	}
}

func (r *Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *Ec2ServicesK8SAwsRouteTableV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_ec2_services_k8s_aws_route_table_v1alpha1")

	var data Ec2ServicesK8SAwsRouteTableV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "routetables"}).
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

	var readResponse Ec2ServicesK8SAwsRouteTableV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("RouteTable")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
