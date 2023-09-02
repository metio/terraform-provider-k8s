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
	_ datasource.DataSource              = &Ec2ServicesK8SAwsSubnetV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &Ec2ServicesK8SAwsSubnetV1Alpha1DataSource{}
)

func NewEc2ServicesK8SAwsSubnetV1Alpha1DataSource() datasource.DataSource {
	return &Ec2ServicesK8SAwsSubnetV1Alpha1DataSource{}
}

type Ec2ServicesK8SAwsSubnetV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type Ec2ServicesK8SAwsSubnetV1Alpha1DataSourceData struct {
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
		AssignIPv6AddressOnCreation     *bool   `tfsdk:"assign_i_pv6_address_on_creation" json:"assignIPv6AddressOnCreation,omitempty"`
		AvailabilityZone                *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
		AvailabilityZoneID              *string `tfsdk:"availability_zone_id" json:"availabilityZoneID,omitempty"`
		CidrBlock                       *string `tfsdk:"cidr_block" json:"cidrBlock,omitempty"`
		CustomerOwnedIPv4Pool           *string `tfsdk:"customer_owned_i_pv4_pool" json:"customerOwnedIPv4Pool,omitempty"`
		EnableDNS64                     *bool   `tfsdk:"enable_dns64" json:"enableDNS64,omitempty"`
		EnableResourceNameDNSAAAARecord *bool   `tfsdk:"enable_resource_name_dnsaaaa_record" json:"enableResourceNameDNSAAAARecord,omitempty"`
		EnableResourceNameDNSARecord    *bool   `tfsdk:"enable_resource_name_dnsa_record" json:"enableResourceNameDNSARecord,omitempty"`
		HostnameType                    *string `tfsdk:"hostname_type" json:"hostnameType,omitempty"`
		Ipv6CIDRBlock                   *string `tfsdk:"ipv6_cidr_block" json:"ipv6CIDRBlock,omitempty"`
		Ipv6Native                      *bool   `tfsdk:"ipv6_native" json:"ipv6Native,omitempty"`
		MapPublicIPOnLaunch             *bool   `tfsdk:"map_public_ip_on_launch" json:"mapPublicIPOnLaunch,omitempty"`
		OutpostARN                      *string `tfsdk:"outpost_arn" json:"outpostARN,omitempty"`
		RouteTableRefs                  *[]struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"route_table_refs" json:"routeTableRefs,omitempty"`
		RouteTables *[]string `tfsdk:"route_tables" json:"routeTables,omitempty"`
		Tags        *[]struct {
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

func (r *Ec2ServicesK8SAwsSubnetV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_subnet_v1alpha1"
}

func (r *Ec2ServicesK8SAwsSubnetV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Subnet is the Schema for the Subnets API",
		MarkdownDescription: "Subnet is the Schema for the Subnets API",
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
				Description:         "SubnetSpec defines the desired state of Subnet.  Describes a subnet.",
				MarkdownDescription: "SubnetSpec defines the desired state of Subnet.  Describes a subnet.",
				Attributes: map[string]schema.Attribute{
					"assign_i_pv6_address_on_creation": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"availability_zone": schema.StringAttribute{
						Description:         "The Availability Zone or Local Zone for the subnet.  Default: Amazon Web Services selects one for you. If you create more than one subnet in your VPC, we do not necessarily select a different zone for each subnet.  To create a subnet in a Local Zone, set this value to the Local Zone ID, for example us-west-2-lax-1a. For information about the Regions that support Local Zones, see Available Regions (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions) in the Amazon Elastic Compute Cloud User Guide.  To create a subnet in an Outpost, set this value to the Availability Zone for the Outpost and specify the Outpost ARN.",
						MarkdownDescription: "The Availability Zone or Local Zone for the subnet.  Default: Amazon Web Services selects one for you. If you create more than one subnet in your VPC, we do not necessarily select a different zone for each subnet.  To create a subnet in a Local Zone, set this value to the Local Zone ID, for example us-west-2-lax-1a. For information about the Regions that support Local Zones, see Available Regions (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#concepts-available-regions) in the Amazon Elastic Compute Cloud User Guide.  To create a subnet in an Outpost, set this value to the Availability Zone for the Outpost and specify the Outpost ARN.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"availability_zone_id": schema.StringAttribute{
						Description:         "The AZ ID or the Local Zone ID of the subnet.",
						MarkdownDescription: "The AZ ID or the Local Zone ID of the subnet.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cidr_block": schema.StringAttribute{
						Description:         "The IPv4 network range for the subnet, in CIDR notation. For example, 10.0.0.0/24. We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18.  This parameter is not supported for an IPv6 only subnet.",
						MarkdownDescription: "The IPv4 network range for the subnet, in CIDR notation. For example, 10.0.0.0/24. We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18.  This parameter is not supported for an IPv6 only subnet.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"customer_owned_i_pv4_pool": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_dns64": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_resource_name_dnsaaaa_record": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"enable_resource_name_dnsa_record": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"hostname_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv6_cidr_block": schema.StringAttribute{
						Description:         "The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length.  This parameter is required for an IPv6 only subnet.",
						MarkdownDescription: "The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length.  This parameter is required for an IPv6 only subnet.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ipv6_native": schema.BoolAttribute{
						Description:         "Indicates whether to create an IPv6 only subnet.",
						MarkdownDescription: "Indicates whether to create an IPv6 only subnet.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"map_public_ip_on_launch": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"outpost_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the Outpost. If you specify an Outpost ARN, you must also specify the Availability Zone of the Outpost subnet.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the Outpost. If you specify an Outpost ARN, you must also specify the Availability Zone of the Outpost subnet.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"route_table_refs": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"route_tables": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
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

func (r *Ec2ServicesK8SAwsSubnetV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *Ec2ServicesK8SAwsSubnetV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_ec2_services_k8s_aws_subnet_v1alpha1")

	var data Ec2ServicesK8SAwsSubnetV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "ec2.services.k8s.aws", Version: "v1alpha1", Resource: "Subnet"}).
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

	var readResponse Ec2ServicesK8SAwsSubnetV1Alpha1DataSourceData
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
	data.Kind = pointer.String("Subnet")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
