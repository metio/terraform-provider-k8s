/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &CiliumIoCiliumNodeV2Manifest{}
)

func NewCiliumIoCiliumNodeV2Manifest() datasource.DataSource {
	return &CiliumIoCiliumNodeV2Manifest{}
}

type CiliumIoCiliumNodeV2Manifest struct{}

type CiliumIoCiliumNodeV2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Addresses *[]struct {
			Ip   *string `tfsdk:"ip" json:"ip,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"addresses" json:"addresses,omitempty"`
		Alibaba_cloud *struct {
			Availability_zone   *string            `tfsdk:"availability_zone" json:"availability-zone,omitempty"`
			Cidr_block          *string            `tfsdk:"cidr_block" json:"cidr-block,omitempty"`
			Instance_type       *string            `tfsdk:"instance_type" json:"instance-type,omitempty"`
			Security_group_tags *map[string]string `tfsdk:"security_group_tags" json:"security-group-tags,omitempty"`
			Security_groups     *[]string          `tfsdk:"security_groups" json:"security-groups,omitempty"`
			Vpc_id              *string            `tfsdk:"vpc_id" json:"vpc-id,omitempty"`
			Vswitch_tags        *map[string]string `tfsdk:"vswitch_tags" json:"vswitch-tags,omitempty"`
			Vswitches           *[]string          `tfsdk:"vswitches" json:"vswitches,omitempty"`
		} `tfsdk:"alibaba_cloud" json:"alibaba-cloud,omitempty"`
		Azure *struct {
			Interface_name *string `tfsdk:"interface_name" json:"interface-name,omitempty"`
		} `tfsdk:"azure" json:"azure,omitempty"`
		Bootid     *string `tfsdk:"bootid" json:"bootid,omitempty"`
		Encryption *struct {
			Key *int64 `tfsdk:"key" json:"key,omitempty"`
		} `tfsdk:"encryption" json:"encryption,omitempty"`
		Eni *struct {
			Availability_zone         *string            `tfsdk:"availability_zone" json:"availability-zone,omitempty"`
			Delete_on_termination     *bool              `tfsdk:"delete_on_termination" json:"delete-on-termination,omitempty"`
			Disable_prefix_delegation *bool              `tfsdk:"disable_prefix_delegation" json:"disable-prefix-delegation,omitempty"`
			Exclude_interface_tags    *map[string]string `tfsdk:"exclude_interface_tags" json:"exclude-interface-tags,omitempty"`
			First_interface_index     *int64             `tfsdk:"first_interface_index" json:"first-interface-index,omitempty"`
			Instance_id               *string            `tfsdk:"instance_id" json:"instance-id,omitempty"`
			Instance_type             *string            `tfsdk:"instance_type" json:"instance-type,omitempty"`
			Max_above_watermark       *int64             `tfsdk:"max_above_watermark" json:"max-above-watermark,omitempty"`
			Min_allocate              *int64             `tfsdk:"min_allocate" json:"min-allocate,omitempty"`
			Node_subnet_id            *string            `tfsdk:"node_subnet_id" json:"node-subnet-id,omitempty"`
			Pre_allocate              *int64             `tfsdk:"pre_allocate" json:"pre-allocate,omitempty"`
			Security_group_tags       *map[string]string `tfsdk:"security_group_tags" json:"security-group-tags,omitempty"`
			Security_groups           *[]string          `tfsdk:"security_groups" json:"security-groups,omitempty"`
			Subnet_ids                *[]string          `tfsdk:"subnet_ids" json:"subnet-ids,omitempty"`
			Subnet_tags               *map[string]string `tfsdk:"subnet_tags" json:"subnet-tags,omitempty"`
			Use_primary_address       *bool              `tfsdk:"use_primary_address" json:"use-primary-address,omitempty"`
			Vpc_id                    *string            `tfsdk:"vpc_id" json:"vpc-id,omitempty"`
		} `tfsdk:"eni" json:"eni,omitempty"`
		Health *struct {
			Ipv4 *string `tfsdk:"ipv4" json:"ipv4,omitempty"`
			Ipv6 *string `tfsdk:"ipv6" json:"ipv6,omitempty"`
		} `tfsdk:"health" json:"health,omitempty"`
		Ingress *struct {
			Ipv4 *string `tfsdk:"ipv4" json:"ipv4,omitempty"`
			Ipv6 *string `tfsdk:"ipv6" json:"ipv6,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Instance_id *string `tfsdk:"instance_id" json:"instance-id,omitempty"`
		Ipam        *struct {
			Ipv6_pool *struct {
				Owner    *string `tfsdk:"owner" json:"owner,omitempty"`
				Resource *string `tfsdk:"resource" json:"resource,omitempty"`
			} `tfsdk:"ipv6_pool" json:"ipv6-pool,omitempty"`
			Max_above_watermark *int64    `tfsdk:"max_above_watermark" json:"max-above-watermark,omitempty"`
			Max_allocate        *int64    `tfsdk:"max_allocate" json:"max-allocate,omitempty"`
			Min_allocate        *int64    `tfsdk:"min_allocate" json:"min-allocate,omitempty"`
			PodCIDRs            *[]string `tfsdk:"pod_cidrs" json:"podCIDRs,omitempty"`
			Pool                *struct {
				Owner    *string `tfsdk:"owner" json:"owner,omitempty"`
				Resource *string `tfsdk:"resource" json:"resource,omitempty"`
			} `tfsdk:"pool" json:"pool,omitempty"`
			Pools *struct {
				Allocated *[]struct {
					Cidrs *[]string `tfsdk:"cidrs" json:"cidrs,omitempty"`
					Pool  *string   `tfsdk:"pool" json:"pool,omitempty"`
				} `tfsdk:"allocated" json:"allocated,omitempty"`
				Requested *[]struct {
					Needed *struct {
						Ipv4_addrs *int64 `tfsdk:"ipv4_addrs" json:"ipv4-addrs,omitempty"`
						Ipv6_addrs *int64 `tfsdk:"ipv6_addrs" json:"ipv6-addrs,omitempty"`
					} `tfsdk:"needed" json:"needed,omitempty"`
					Pool *string `tfsdk:"pool" json:"pool,omitempty"`
				} `tfsdk:"requested" json:"requested,omitempty"`
			} `tfsdk:"pools" json:"pools,omitempty"`
			Pre_allocate   *int64             `tfsdk:"pre_allocate" json:"pre-allocate,omitempty"`
			Static_ip_tags *map[string]string `tfsdk:"static_ip_tags" json:"static-ip-tags,omitempty"`
		} `tfsdk:"ipam" json:"ipam,omitempty"`
		Nodeidentity *int64 `tfsdk:"nodeidentity" json:"nodeidentity,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CiliumIoCiliumNodeV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_node_v2_manifest"
}

func (r *CiliumIoCiliumNodeV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumNode represents a node managed by Cilium. It contains a specification to control various node specific configuration aspects and a status section to represent the status of the node.",
		MarkdownDescription: "CiliumNode represents a node managed by Cilium. It contains a specification to control various node specific configuration aspects and a status section to represent the status of the node.",
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
				Description:         "Spec defines the desired specification/configuration of the node.",
				MarkdownDescription: "Spec defines the desired specification/configuration of the node.",
				Attributes: map[string]schema.Attribute{
					"addresses": schema.ListNestedAttribute{
						Description:         "Addresses is the list of all node addresses.",
						MarkdownDescription: "Addresses is the list of all node addresses.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ip": schema.StringAttribute{
									Description:         "IP is an IP of a node",
									MarkdownDescription: "IP is an IP of a node",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type is the type of the node address",
									MarkdownDescription: "Type is the type of the node address",
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

					"alibaba_cloud": schema.SingleNestedAttribute{
						Description:         "AlibabaCloud is the AlibabaCloud IPAM specific configuration.",
						MarkdownDescription: "AlibabaCloud is the AlibabaCloud IPAM specific configuration.",
						Attributes: map[string]schema.Attribute{
							"availability_zone": schema.StringAttribute{
								Description:         "AvailabilityZone is the availability zone to use when allocating ENIs.",
								MarkdownDescription: "AvailabilityZone is the availability zone to use when allocating ENIs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cidr_block": schema.StringAttribute{
								Description:         "CIDRBlock is vpc ipv4 CIDR",
								MarkdownDescription: "CIDRBlock is vpc ipv4 CIDR",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance_type": schema.StringAttribute{
								Description:         "InstanceType is the ECS instance type, e.g. 'ecs.g6.2xlarge'",
								MarkdownDescription: "InstanceType is the ECS instance type, e.g. 'ecs.g6.2xlarge'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_group_tags": schema.MapAttribute{
								Description:         "SecurityGroupTags is the list of tags to use when evaluating which security groups to use for the ENI.",
								MarkdownDescription: "SecurityGroupTags is the list of tags to use when evaluating which security groups to use for the ENI.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_groups": schema.ListAttribute{
								Description:         "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",
								MarkdownDescription: "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vpc_id": schema.StringAttribute{
								Description:         "VPCID is the VPC ID to use when allocating ENIs.",
								MarkdownDescription: "VPCID is the VPC ID to use when allocating ENIs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vswitch_tags": schema.MapAttribute{
								Description:         "VSwitchTags is the list of tags to use when evaluating which vSwitch to use for the ENI.",
								MarkdownDescription: "VSwitchTags is the list of tags to use when evaluating which vSwitch to use for the ENI.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vswitches": schema.ListAttribute{
								Description:         "VSwitches is the ID of vSwitch available for ENI",
								MarkdownDescription: "VSwitches is the ID of vSwitch available for ENI",
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

					"azure": schema.SingleNestedAttribute{
						Description:         "Azure is the Azure IPAM specific configuration.",
						MarkdownDescription: "Azure is the Azure IPAM specific configuration.",
						Attributes: map[string]schema.Attribute{
							"interface_name": schema.StringAttribute{
								Description:         "InterfaceName is the name of the interface the cilium-operator will use to allocate all the IPs on",
								MarkdownDescription: "InterfaceName is the name of the interface the cilium-operator will use to allocate all the IPs on",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"bootid": schema.StringAttribute{
						Description:         "BootID is a unique node identifier generated on boot",
						MarkdownDescription: "BootID is a unique node identifier generated on boot",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encryption": schema.SingleNestedAttribute{
						Description:         "Encryption is the encryption configuration of the node.",
						MarkdownDescription: "Encryption is the encryption configuration of the node.",
						Attributes: map[string]schema.Attribute{
							"key": schema.Int64Attribute{
								Description:         "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
								MarkdownDescription: "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"eni": schema.SingleNestedAttribute{
						Description:         "ENI is the AWS ENI specific configuration.",
						MarkdownDescription: "ENI is the AWS ENI specific configuration.",
						Attributes: map[string]schema.Attribute{
							"availability_zone": schema.StringAttribute{
								Description:         "AvailabilityZone is the availability zone to use when allocating ENIs.",
								MarkdownDescription: "AvailabilityZone is the availability zone to use when allocating ENIs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delete_on_termination": schema.BoolAttribute{
								Description:         "DeleteOnTermination defines that the ENI should be deleted when the associated instance is terminated. If the parameter is not set the default behavior is to delete the ENI on instance termination.",
								MarkdownDescription: "DeleteOnTermination defines that the ENI should be deleted when the associated instance is terminated. If the parameter is not set the default behavior is to delete the ENI on instance termination.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_prefix_delegation": schema.BoolAttribute{
								Description:         "DisablePrefixDelegation determines whether ENI prefix delegation should be disabled on this node.",
								MarkdownDescription: "DisablePrefixDelegation determines whether ENI prefix delegation should be disabled on this node.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exclude_interface_tags": schema.MapAttribute{
								Description:         "ExcludeInterfaceTags is the list of tags to use when excluding ENIs for Cilium IP allocation. Any interface matching this set of tags will not be managed by Cilium.",
								MarkdownDescription: "ExcludeInterfaceTags is the list of tags to use when excluding ENIs for Cilium IP allocation. Any interface matching this set of tags will not be managed by Cilium.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"first_interface_index": schema.Int64Attribute{
								Description:         "FirstInterfaceIndex is the index of the first ENI to use for IP allocation, e.g. if the node has eth0, eth1, eth2 and FirstInterfaceIndex is set to 1, then only eth1 and eth2 will be used for IP allocation, eth0 will be ignored for PodIP allocation.",
								MarkdownDescription: "FirstInterfaceIndex is the index of the first ENI to use for IP allocation, e.g. if the node has eth0, eth1, eth2 and FirstInterfaceIndex is set to 1, then only eth1 and eth2 will be used for IP allocation, eth0 will be ignored for PodIP allocation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"instance_id": schema.StringAttribute{
								Description:         "InstanceID is the AWS InstanceId of the node. The InstanceID is used to retrieve AWS metadata for the node. OBSOLETE: This field is obsolete, please use Spec.InstanceID",
								MarkdownDescription: "InstanceID is the AWS InstanceId of the node. The InstanceID is used to retrieve AWS metadata for the node. OBSOLETE: This field is obsolete, please use Spec.InstanceID",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance_type": schema.StringAttribute{
								Description:         "InstanceType is the AWS EC2 instance type, e.g. 'm5.large'",
								MarkdownDescription: "InstanceType is the AWS EC2 instance type, e.g. 'm5.large'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_above_watermark": schema.Int64Attribute{
								Description:         "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs. OBSOLETE: This field is obsolete, please use Spec.IPAM.MaxAboveWatermark",
								MarkdownDescription: "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs. OBSOLETE: This field is obsolete, please use Spec.IPAM.MaxAboveWatermark",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"min_allocate": schema.Int64Attribute{
								Description:         "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs. OBSOLETE: This field is obsolete, please use Spec.IPAM.MinAllocate",
								MarkdownDescription: "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs. OBSOLETE: This field is obsolete, please use Spec.IPAM.MinAllocate",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"node_subnet_id": schema.StringAttribute{
								Description:         "NodeSubnetID is the subnet of the primary ENI the instance was brought up with. It is used as a sensible default subnet to create ENIs in.",
								MarkdownDescription: "NodeSubnetID is the subnet of the primary ENI the instance was brought up with. It is used as a sensible default subnet to create ENIs in.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pre_allocate": schema.Int64Attribute{
								Description:         "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved. OBSOLETE: This field is obsolete, please use Spec.IPAM.PreAllocate",
								MarkdownDescription: "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved. OBSOLETE: This field is obsolete, please use Spec.IPAM.PreAllocate",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"security_group_tags": schema.MapAttribute{
								Description:         "SecurityGroupTags is the list of tags to use when evaliating what AWS security groups to use for the ENI.",
								MarkdownDescription: "SecurityGroupTags is the list of tags to use when evaliating what AWS security groups to use for the ENI.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_groups": schema.ListAttribute{
								Description:         "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",
								MarkdownDescription: "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subnet_ids": schema.ListAttribute{
								Description:         "SubnetIDs is the list of subnet ids to use when evaluating what AWS subnets to use for ENI and IP allocation.",
								MarkdownDescription: "SubnetIDs is the list of subnet ids to use when evaluating what AWS subnets to use for ENI and IP allocation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subnet_tags": schema.MapAttribute{
								Description:         "SubnetTags is the list of tags to use when evaluating what AWS subnets to use for ENI and IP allocation.",
								MarkdownDescription: "SubnetTags is the list of tags to use when evaluating what AWS subnets to use for ENI and IP allocation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_primary_address": schema.BoolAttribute{
								Description:         "UsePrimaryAddress determines whether an ENI's primary address should be available for allocations on the node",
								MarkdownDescription: "UsePrimaryAddress determines whether an ENI's primary address should be available for allocations on the node",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vpc_id": schema.StringAttribute{
								Description:         "VpcID is the VPC ID to use when allocating ENIs.",
								MarkdownDescription: "VpcID is the VPC ID to use when allocating ENIs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"health": schema.SingleNestedAttribute{
						Description:         "HealthAddressing is the addressing information for health connectivity checking.",
						MarkdownDescription: "HealthAddressing is the addressing information for health connectivity checking.",
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.StringAttribute{
								Description:         "IPv4 is the IPv4 address of the IPv4 health endpoint.",
								MarkdownDescription: "IPv4 is the IPv4 address of the IPv4 health endpoint.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ipv6": schema.StringAttribute{
								Description:         "IPv6 is the IPv6 address of the IPv4 health endpoint.",
								MarkdownDescription: "IPv6 is the IPv6 address of the IPv4 health endpoint.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": schema.SingleNestedAttribute{
						Description:         "IngressAddressing is the addressing information for Ingress listener.",
						MarkdownDescription: "IngressAddressing is the addressing information for Ingress listener.",
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"instance_id": schema.StringAttribute{
						Description:         "InstanceID is the identifier of the node. This is different from the node name which is typically the FQDN of the node. The InstanceID typically refers to the identifier used by the cloud provider or some other means of identification.",
						MarkdownDescription: "InstanceID is the identifier of the node. This is different from the node name which is typically the FQDN of the node. The InstanceID typically refers to the identifier used by the cloud provider or some other means of identification.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipam": schema.SingleNestedAttribute{
						Description:         "IPAM is the address management specification. This section can be populated by a user or it can be automatically populated by an IPAM operator.",
						MarkdownDescription: "IPAM is the address management specification. This section can be populated by a user or it can be automatically populated by an IPAM operator.",
						Attributes: map[string]schema.Attribute{
							"ipv6_pool": schema.SingleNestedAttribute{
								Description:         "IPv6Pool is the list of IPv6 addresses available to the node for allocation. When an IPv6 address is used, it will remain on this list but will be added to Status.IPAM.IPv6Used",
								MarkdownDescription: "IPv6Pool is the list of IPv6 addresses available to the node for allocation. When an IPv6 address is used, it will remain on this list but will be added to Status.IPAM.IPv6Used",
								Attributes: map[string]schema.Attribute{
									"owner": schema.StringAttribute{
										Description:         "Owner is the owner of the IP. This field is set if the IP has been allocated. It will be set to the pod name or another identifier representing the usage of the IP The owner field is left blank for an entry in Spec.IPAM.Pool and filled out as the IP is used and also added to Status.IPAM.Used.",
										MarkdownDescription: "Owner is the owner of the IP. This field is set if the IP has been allocated. It will be set to the pod name or another identifier representing the usage of the IP The owner field is left blank for an entry in Spec.IPAM.Pool and filled out as the IP is used and also added to Status.IPAM.Used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resource": schema.StringAttribute{
										Description:         "Resource is set for both available and allocated IPs, it represents what resource the IP is associated with, e.g. in combination with AWS ENI, this will refer to the ID of the ENI",
										MarkdownDescription: "Resource is set for both available and allocated IPs, it represents what resource the IP is associated with, e.g. in combination with AWS ENI, this will refer to the ID of the ENI",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_above_watermark": schema.Int64Attribute{
								Description:         "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs.",
								MarkdownDescription: "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"max_allocate": schema.Int64Attribute{
								Description:         "MaxAllocate is the maximum number of IPs that can be allocated to the node. When the current amount of allocated IPs will approach this value, the considered value for PreAllocate will decrease down to 0 in order to not attempt to allocate more addresses than defined.",
								MarkdownDescription: "MaxAllocate is the maximum number of IPs that can be allocated to the node. When the current amount of allocated IPs will approach this value, the considered value for PreAllocate will decrease down to 0 in order to not attempt to allocate more addresses than defined.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"min_allocate": schema.Int64Attribute{
								Description:         "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs.",
								MarkdownDescription: "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"pod_cidrs": schema.ListAttribute{
								Description:         "PodCIDRs is the list of CIDRs available to the node for allocation. When an IP is used, the IP will be added to Status.IPAM.Used",
								MarkdownDescription: "PodCIDRs is the list of CIDRs available to the node for allocation. When an IP is used, the IP will be added to Status.IPAM.Used",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pool": schema.SingleNestedAttribute{
								Description:         "Pool is the list of IPv4 addresses available to the node for allocation. When an IPv4 address is used, it will remain on this list but will be added to Status.IPAM.Used",
								MarkdownDescription: "Pool is the list of IPv4 addresses available to the node for allocation. When an IPv4 address is used, it will remain on this list but will be added to Status.IPAM.Used",
								Attributes: map[string]schema.Attribute{
									"owner": schema.StringAttribute{
										Description:         "Owner is the owner of the IP. This field is set if the IP has been allocated. It will be set to the pod name or another identifier representing the usage of the IP The owner field is left blank for an entry in Spec.IPAM.Pool and filled out as the IP is used and also added to Status.IPAM.Used.",
										MarkdownDescription: "Owner is the owner of the IP. This field is set if the IP has been allocated. It will be set to the pod name or another identifier representing the usage of the IP The owner field is left blank for an entry in Spec.IPAM.Pool and filled out as the IP is used and also added to Status.IPAM.Used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resource": schema.StringAttribute{
										Description:         "Resource is set for both available and allocated IPs, it represents what resource the IP is associated with, e.g. in combination with AWS ENI, this will refer to the ID of the ENI",
										MarkdownDescription: "Resource is set for both available and allocated IPs, it represents what resource the IP is associated with, e.g. in combination with AWS ENI, this will refer to the ID of the ENI",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pools": schema.SingleNestedAttribute{
								Description:         "Pools contains the list of assigned IPAM pools for this node.",
								MarkdownDescription: "Pools contains the list of assigned IPAM pools for this node.",
								Attributes: map[string]schema.Attribute{
									"allocated": schema.ListNestedAttribute{
										Description:         "Allocated contains the list of pooled CIDR assigned to this node. The operator will add new pod CIDRs to this field, whereas the agent will remove CIDRs it has released.",
										MarkdownDescription: "Allocated contains the list of pooled CIDR assigned to this node. The operator will add new pod CIDRs to this field, whereas the agent will remove CIDRs it has released.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidrs": schema.ListAttribute{
													Description:         "CIDRs contains a list of pod CIDRs currently allocated from this pool",
													MarkdownDescription: "CIDRs contains a list of pod CIDRs currently allocated from this pool",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pool": schema.StringAttribute{
													Description:         "Pool is the name of the IPAM pool backing this allocation",
													MarkdownDescription: "Pool is the name of the IPAM pool backing this allocation",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"requested": schema.ListNestedAttribute{
										Description:         "Requested contains a list of IPAM pool requests, i.e. indicates how many addresses this node requests out of each pool listed here. This field is owned and written to by cilium-agent and read by the operator.",
										MarkdownDescription: "Requested contains a list of IPAM pool requests, i.e. indicates how many addresses this node requests out of each pool listed here. This field is owned and written to by cilium-agent and read by the operator.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"needed": schema.SingleNestedAttribute{
													Description:         "Needed indicates how many IPs out of the above Pool this node requests from the operator. The operator runs a reconciliation loop to ensure each node always has enough PodCIDRs allocated in each pool to fulfill the requested number of IPs here.",
													MarkdownDescription: "Needed indicates how many IPs out of the above Pool this node requests from the operator. The operator runs a reconciliation loop to ensure each node always has enough PodCIDRs allocated in each pool to fulfill the requested number of IPs here.",
													Attributes: map[string]schema.Attribute{
														"ipv4_addrs": schema.Int64Attribute{
															Description:         "IPv4Addrs contains the number of requested IPv4 addresses out of a given pool",
															MarkdownDescription: "IPv4Addrs contains the number of requested IPv4 addresses out of a given pool",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ipv6_addrs": schema.Int64Attribute{
															Description:         "IPv6Addrs contains the number of requested IPv6 addresses out of a given pool",
															MarkdownDescription: "IPv6Addrs contains the number of requested IPv6 addresses out of a given pool",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"pool": schema.StringAttribute{
													Description:         "Pool is the name of the IPAM pool backing this request",
													MarkdownDescription: "Pool is the name of the IPAM pool backing this request",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
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

							"pre_allocate": schema.Int64Attribute{
								Description:         "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved.",
								MarkdownDescription: "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"static_ip_tags": schema.MapAttribute{
								Description:         "StaticIPTags are used to determine the pool of IPs from which to attribute a static IP to the node. For example in AWS this is used to filter Elastic IP Addresses.",
								MarkdownDescription: "StaticIPTags are used to determine the pool of IPs from which to attribute a static IP to the node. For example in AWS this is used to filter Elastic IP Addresses.",
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

					"nodeidentity": schema.Int64Attribute{
						Description:         "NodeIdentity is the Cilium numeric identity allocated for the node, if any.",
						MarkdownDescription: "NodeIdentity is the Cilium numeric identity allocated for the node, if any.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumNodeV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_node_v2_manifest")

	var model CiliumIoCiliumNodeV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2")
	model.Kind = pointer.String("CiliumNode")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
