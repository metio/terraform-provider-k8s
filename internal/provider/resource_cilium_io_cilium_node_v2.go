/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type CiliumIoCiliumNodeV2Resource struct{}

var (
	_ resource.Resource = (*CiliumIoCiliumNodeV2Resource)(nil)
)

type CiliumIoCiliumNodeV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CiliumIoCiliumNodeV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Addresses *[]struct {
			Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"addresses" yaml:"addresses,omitempty"`

		Alibaba_cloud *struct {
			Availability_zone *string `tfsdk:"availability__zone" yaml:"availability-zone,omitempty"`

			Cidr_block *string `tfsdk:"cidr__block" yaml:"cidr-block,omitempty"`

			Instance_type *string `tfsdk:"instance__type" yaml:"instance-type,omitempty"`

			Security_group_tags *map[string]string `tfsdk:"security__group__tags" yaml:"security-group-tags,omitempty"`

			Security_groups *[]string `tfsdk:"security__groups" yaml:"security-groups,omitempty"`

			Vpc_id *string `tfsdk:"vpc__id" yaml:"vpc-id,omitempty"`

			Vswitch_tags *map[string]string `tfsdk:"vswitch__tags" yaml:"vswitch-tags,omitempty"`

			Vswitches *[]string `tfsdk:"vswitches" yaml:"vswitches,omitempty"`
		} `tfsdk:"alibaba__cloud" yaml:"alibaba-cloud,omitempty"`

		Azure *struct {
			Interface_name *string `tfsdk:"interface__name" yaml:"interface-name,omitempty"`
		} `tfsdk:"azure" yaml:"azure,omitempty"`

		Encryption *struct {
			Key *int64 `tfsdk:"key" yaml:"key,omitempty"`
		} `tfsdk:"encryption" yaml:"encryption,omitempty"`

		Eni *struct {
			Availability_zone *string `tfsdk:"availability__zone" yaml:"availability-zone,omitempty"`

			Delete_on_termination *bool `tfsdk:"delete__on__termination" yaml:"delete-on-termination,omitempty"`

			Disable_prefix_delegation *bool `tfsdk:"disable__prefix__delegation" yaml:"disable-prefix-delegation,omitempty"`

			Exclude_interface_tags *map[string]string `tfsdk:"exclude__interface__tags" yaml:"exclude-interface-tags,omitempty"`

			First_interface_index *int64 `tfsdk:"first__interface__index" yaml:"first-interface-index,omitempty"`

			Instance_id *string `tfsdk:"instance__id" yaml:"instance-id,omitempty"`

			Instance_type *string `tfsdk:"instance__type" yaml:"instance-type,omitempty"`

			Max_above_watermark *int64 `tfsdk:"max__above__watermark" yaml:"max-above-watermark,omitempty"`

			Min_allocate *int64 `tfsdk:"min__allocate" yaml:"min-allocate,omitempty"`

			Pre_allocate *int64 `tfsdk:"pre__allocate" yaml:"pre-allocate,omitempty"`

			Security_group_tags *map[string]string `tfsdk:"security__group__tags" yaml:"security-group-tags,omitempty"`

			Security_groups *[]string `tfsdk:"security__groups" yaml:"security-groups,omitempty"`

			Subnet_ids *[]string `tfsdk:"subnet__ids" yaml:"subnet-ids,omitempty"`

			Subnet_tags *map[string]string `tfsdk:"subnet__tags" yaml:"subnet-tags,omitempty"`

			Use_primary_address *bool `tfsdk:"use__primary__address" yaml:"use-primary-address,omitempty"`

			Vpc_id *string `tfsdk:"vpc__id" yaml:"vpc-id,omitempty"`
		} `tfsdk:"eni" yaml:"eni,omitempty"`

		Health *struct {
			Ipv4 *string `tfsdk:"ipv4" yaml:"ipv4,omitempty"`

			Ipv6 *string `tfsdk:"ipv6" yaml:"ipv6,omitempty"`
		} `tfsdk:"health" yaml:"health,omitempty"`

		Ingress *struct {
			Ipv4 *string `tfsdk:"ipv4" yaml:"ipv4,omitempty"`

			Ipv6 *string `tfsdk:"ipv6" yaml:"ipv6,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		Instance_id *string `tfsdk:"instance__id" yaml:"instance-id,omitempty"`

		Ipam *struct {
			Max_above_watermark *int64 `tfsdk:"max__above__watermark" yaml:"max-above-watermark,omitempty"`

			Max_allocate *int64 `tfsdk:"max__allocate" yaml:"max-allocate,omitempty"`

			Min_allocate *int64 `tfsdk:"min__allocate" yaml:"min-allocate,omitempty"`

			Pod_cidr_allocation_threshold *int64 `tfsdk:"pod__cidr__allocation__threshold" yaml:"pod-cidr-allocation-threshold,omitempty"`

			Pod_cidr_release_threshold *int64 `tfsdk:"pod__cidr__release__threshold" yaml:"pod-cidr-release-threshold,omitempty"`

			PodCIDRs *[]string `tfsdk:"pod_cid_rs" yaml:"podCIDRs,omitempty"`

			Pool *struct {
				Owner *string `tfsdk:"owner" yaml:"owner,omitempty"`

				Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
			} `tfsdk:"pool" yaml:"pool,omitempty"`

			Pre_allocate *int64 `tfsdk:"pre__allocate" yaml:"pre-allocate,omitempty"`
		} `tfsdk:"ipam" yaml:"ipam,omitempty"`

		Nodeidentity *int64 `tfsdk:"nodeidentity" yaml:"nodeidentity,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCiliumIoCiliumNodeV2Resource() resource.Resource {
	return &CiliumIoCiliumNodeV2Resource{}
}

func (r *CiliumIoCiliumNodeV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cilium_io_cilium_node_v2"
}

func (r *CiliumIoCiliumNodeV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CiliumNode represents a node managed by Cilium. It contains a specification to control various node specific configuration aspects and a status section to represent the status of the node.",
		MarkdownDescription: "CiliumNode represents a node managed by Cilium. It contains a specification to control various node specific configuration aspects and a status section to represent the status of the node.",
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
				Description:         "Spec defines the desired specification/configuration of the node.",
				MarkdownDescription: "Spec defines the desired specification/configuration of the node.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"addresses": {
						Description:         "Addresses is the list of all node addresses.",
						MarkdownDescription: "Addresses is the list of all node addresses.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"ip": {
								Description:         "IP is an IP of a node",
								MarkdownDescription: "IP is an IP of a node",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "Type is the type of the node address",
								MarkdownDescription: "Type is the type of the node address",

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

					"alibaba__cloud": {
						Description:         "AlibabaCloud is the AlibabaCloud IPAM specific configuration.",
						MarkdownDescription: "AlibabaCloud is the AlibabaCloud IPAM specific configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"availability__zone": {
								Description:         "AvailabilityZone is the availability zone to use when allocating ENIs.",
								MarkdownDescription: "AvailabilityZone is the availability zone to use when allocating ENIs.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cidr__block": {
								Description:         "CIDRBlock is vpc ipv4 CIDR",
								MarkdownDescription: "CIDRBlock is vpc ipv4 CIDR",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance__type": {
								Description:         "InstanceType is the ECS instance type, e.g. 'ecs.g6.2xlarge'",
								MarkdownDescription: "InstanceType is the ECS instance type, e.g. 'ecs.g6.2xlarge'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security__group__tags": {
								Description:         "SecurityGroupTags is the list of tags to use when evaluating which security groups to use for the ENI.",
								MarkdownDescription: "SecurityGroupTags is the list of tags to use when evaluating which security groups to use for the ENI.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security__groups": {
								Description:         "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",
								MarkdownDescription: "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vpc__id": {
								Description:         "VPCID is the VPC ID to use when allocating ENIs.",
								MarkdownDescription: "VPCID is the VPC ID to use when allocating ENIs.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vswitch__tags": {
								Description:         "VSwitchTags is the list of tags to use when evaluating which vSwitch to use for the ENI.",
								MarkdownDescription: "VSwitchTags is the list of tags to use when evaluating which vSwitch to use for the ENI.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vswitches": {
								Description:         "VSwitches is the ID of vSwitch available for ENI",
								MarkdownDescription: "VSwitches is the ID of vSwitch available for ENI",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"azure": {
						Description:         "Azure is the Azure IPAM specific configuration.",
						MarkdownDescription: "Azure is the Azure IPAM specific configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"interface__name": {
								Description:         "InterfaceName is the name of the interface the cilium-operator will use to allocate all the IPs on",
								MarkdownDescription: "InterfaceName is the name of the interface the cilium-operator will use to allocate all the IPs on",

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

					"encryption": {
						Description:         "Encryption is the encryption configuration of the node.",
						MarkdownDescription: "Encryption is the encryption configuration of the node.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "Key is the index to the key to use for encryption or 0 if encryption is disabled.",
								MarkdownDescription: "Key is the index to the key to use for encryption or 0 if encryption is disabled.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"eni": {
						Description:         "ENI is the AWS ENI specific configuration.",
						MarkdownDescription: "ENI is the AWS ENI specific configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"availability__zone": {
								Description:         "AvailabilityZone is the availability zone to use when allocating ENIs.",
								MarkdownDescription: "AvailabilityZone is the availability zone to use when allocating ENIs.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"delete__on__termination": {
								Description:         "DeleteOnTermination defines that the ENI should be deleted when the associated instance is terminated. If the parameter is not set the default behavior is to delete the ENI on instance termination.",
								MarkdownDescription: "DeleteOnTermination defines that the ENI should be deleted when the associated instance is terminated. If the parameter is not set the default behavior is to delete the ENI on instance termination.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable__prefix__delegation": {
								Description:         "DisablePrefixDelegation determines whether ENI prefix delegation should be disabled on this node.",
								MarkdownDescription: "DisablePrefixDelegation determines whether ENI prefix delegation should be disabled on this node.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exclude__interface__tags": {
								Description:         "ExcludeInterfaceTags is the list of tags to use when excluding ENIs for Cilium IP allocation. Any interface matching this set of tags will not be managed by Cilium.",
								MarkdownDescription: "ExcludeInterfaceTags is the list of tags to use when excluding ENIs for Cilium IP allocation. Any interface matching this set of tags will not be managed by Cilium.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"first__interface__index": {
								Description:         "FirstInterfaceIndex is the index of the first ENI to use for IP allocation, e.g. if the node has eth0, eth1, eth2 and FirstInterfaceIndex is set to 1, then only eth1 and eth2 will be used for IP allocation, eth0 will be ignored for PodIP allocation.",
								MarkdownDescription: "FirstInterfaceIndex is the index of the first ENI to use for IP allocation, e.g. if the node has eth0, eth1, eth2 and FirstInterfaceIndex is set to 1, then only eth1 and eth2 will be used for IP allocation, eth0 will be ignored for PodIP allocation.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"instance__id": {
								Description:         "InstanceID is the AWS InstanceId of the node. The InstanceID is used to retrieve AWS metadata for the node.  OBSOLETE: This field is obsolete, please use Spec.InstanceID",
								MarkdownDescription: "InstanceID is the AWS InstanceId of the node. The InstanceID is used to retrieve AWS metadata for the node.  OBSOLETE: This field is obsolete, please use Spec.InstanceID",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance__type": {
								Description:         "InstanceType is the AWS EC2 instance type, e.g. 'm5.large'",
								MarkdownDescription: "InstanceType is the AWS EC2 instance type, e.g. 'm5.large'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max__above__watermark": {
								Description:         "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs.  OBSOLETE: This field is obsolete, please use Spec.IPAM.MaxAboveWatermark",
								MarkdownDescription: "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs.  OBSOLETE: This field is obsolete, please use Spec.IPAM.MaxAboveWatermark",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"min__allocate": {
								Description:         "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs.  OBSOLETE: This field is obsolete, please use Spec.IPAM.MinAllocate",
								MarkdownDescription: "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs.  OBSOLETE: This field is obsolete, please use Spec.IPAM.MinAllocate",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"pre__allocate": {
								Description:         "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved.  OBSOLETE: This field is obsolete, please use Spec.IPAM.PreAllocate",
								MarkdownDescription: "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved.  OBSOLETE: This field is obsolete, please use Spec.IPAM.PreAllocate",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"security__group__tags": {
								Description:         "SecurityGroupTags is the list of tags to use when evaliating what AWS security groups to use for the ENI.",
								MarkdownDescription: "SecurityGroupTags is the list of tags to use when evaliating what AWS security groups to use for the ENI.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security__groups": {
								Description:         "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",
								MarkdownDescription: "SecurityGroups is the list of security groups to attach to any ENI that is created and attached to the instance.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subnet__ids": {
								Description:         "SubnetIDs is the list of subnet ids to use when evaluating what AWS subnets to use for ENI and IP allocation.",
								MarkdownDescription: "SubnetIDs is the list of subnet ids to use when evaluating what AWS subnets to use for ENI and IP allocation.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subnet__tags": {
								Description:         "SubnetTags is the list of tags to use when evaluating what AWS subnets to use for ENI and IP allocation.",
								MarkdownDescription: "SubnetTags is the list of tags to use when evaluating what AWS subnets to use for ENI and IP allocation.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"use__primary__address": {
								Description:         "UsePrimaryAddress determines whether an ENI's primary address should be available for allocations on the node",
								MarkdownDescription: "UsePrimaryAddress determines whether an ENI's primary address should be available for allocations on the node",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"vpc__id": {
								Description:         "VpcID is the VPC ID to use when allocating ENIs.",
								MarkdownDescription: "VpcID is the VPC ID to use when allocating ENIs.",

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

					"health": {
						Description:         "HealthAddressing is the addressing information for health connectivity checking.",
						MarkdownDescription: "HealthAddressing is the addressing information for health connectivity checking.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ipv4": {
								Description:         "IPv4 is the IPv4 address of the IPv4 health endpoint.",
								MarkdownDescription: "IPv4 is the IPv4 address of the IPv4 health endpoint.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv6": {
								Description:         "IPv6 is the IPv6 address of the IPv4 health endpoint.",
								MarkdownDescription: "IPv6 is the IPv6 address of the IPv4 health endpoint.",

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

					"ingress": {
						Description:         "IngressAddressing is the addressing information for Ingress listener.",
						MarkdownDescription: "IngressAddressing is the addressing information for Ingress listener.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ipv4": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv6": {
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

					"instance__id": {
						Description:         "InstanceID is the identifier of the node. This is different from the node name which is typically the FQDN of the node. The InstanceID typically refers to the identifier used by the cloud provider or some other means of identification.",
						MarkdownDescription: "InstanceID is the identifier of the node. This is different from the node name which is typically the FQDN of the node. The InstanceID typically refers to the identifier used by the cloud provider or some other means of identification.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipam": {
						Description:         "IPAM is the address management specification. This section can be populated by a user or it can be automatically populated by an IPAM operator.",
						MarkdownDescription: "IPAM is the address management specification. This section can be populated by a user or it can be automatically populated by an IPAM operator.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max__above__watermark": {
								Description:         "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs.",
								MarkdownDescription: "MaxAboveWatermark is the maximum number of addresses to allocate beyond the addresses needed to reach the PreAllocate watermark. Going above the watermark can help reduce the number of API calls to allocate IPs, e.g. when a new ENI is allocated, as many secondary IPs as possible are allocated. Limiting the amount can help reduce waste of IPs.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"max__allocate": {
								Description:         "MaxAllocate is the maximum number of IPs that can be allocated to the node. When the current amount of allocated IPs will approach this value, the considered value for PreAllocate will decrease down to 0 in order to not attempt to allocate more addresses than defined.",
								MarkdownDescription: "MaxAllocate is the maximum number of IPs that can be allocated to the node. When the current amount of allocated IPs will approach this value, the considered value for PreAllocate will decrease down to 0 in order to not attempt to allocate more addresses than defined.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"min__allocate": {
								Description:         "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs.",
								MarkdownDescription: "MinAllocate is the minimum number of IPs that must be allocated when the node is first bootstrapped. It defines the minimum base socket of addresses that must be available. After reaching this watermark, the PreAllocate and MaxAboveWatermark logic takes over to continue allocating IPs.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"pod__cidr__allocation__threshold": {
								Description:         "PodCIDRAllocationThreshold defines the minimum number of free IPs which must be available to this node via its pod CIDR pool. If the total number of IP addresses in the pod CIDR pool is less than this value, the pod CIDRs currently in-use by this node will be marked as depleted and cilium-operator will allocate a new pod CIDR to this node. This value effectively defines the buffer of IP addresses available immediately without requiring cilium-operator to get involved.",
								MarkdownDescription: "PodCIDRAllocationThreshold defines the minimum number of free IPs which must be available to this node via its pod CIDR pool. If the total number of IP addresses in the pod CIDR pool is less than this value, the pod CIDRs currently in-use by this node will be marked as depleted and cilium-operator will allocate a new pod CIDR to this node. This value effectively defines the buffer of IP addresses available immediately without requiring cilium-operator to get involved.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"pod__cidr__release__threshold": {
								Description:         "PodCIDRReleaseThreshold defines the maximum number of free IPs which may be available to this node via its pod CIDR pool. While the total number of free IP addresses in the pod CIDR pool is larger than this value, cilium-agent will attempt to release currently unused pod CIDRs.",
								MarkdownDescription: "PodCIDRReleaseThreshold defines the maximum number of free IPs which may be available to this node via its pod CIDR pool. While the total number of free IP addresses in the pod CIDR pool is larger than this value, cilium-agent will attempt to release currently unused pod CIDRs.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"pod_cid_rs": {
								Description:         "PodCIDRs is the list of CIDRs available to the node for allocation. When an IP is used, the IP will be added to Status.IPAM.Used",
								MarkdownDescription: "PodCIDRs is the list of CIDRs available to the node for allocation. When an IP is used, the IP will be added to Status.IPAM.Used",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pool": {
								Description:         "Pool is the list of IPs available to the node for allocation. When an IP is used, the IP will remain on this list but will be added to Status.IPAM.Used",
								MarkdownDescription: "Pool is the list of IPs available to the node for allocation. When an IP is used, the IP will remain on this list but will be added to Status.IPAM.Used",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"owner": {
										Description:         "Owner is the owner of the IP. This field is set if the IP has been allocated. It will be set to the pod name or another identifier representing the usage of the IP  The owner field is left blank for an entry in Spec.IPAM.Pool and filled out as the IP is used and also added to Status.IPAM.Used.",
										MarkdownDescription: "Owner is the owner of the IP. This field is set if the IP has been allocated. It will be set to the pod name or another identifier representing the usage of the IP  The owner field is left blank for an entry in Spec.IPAM.Pool and filled out as the IP is used and also added to Status.IPAM.Used.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource": {
										Description:         "Resource is set for both available and allocated IPs, it represents what resource the IP is associated with, e.g. in combination with AWS ENI, this will refer to the ID of the ENI",
										MarkdownDescription: "Resource is set for both available and allocated IPs, it represents what resource the IP is associated with, e.g. in combination with AWS ENI, this will refer to the ID of the ENI",

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

							"pre__allocate": {
								Description:         "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved.",
								MarkdownDescription: "PreAllocate defines the number of IP addresses that must be available for allocation in the IPAMspec. It defines the buffer of addresses available immediately without requiring cilium-operator to get involved.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nodeidentity": {
						Description:         "NodeIdentity is the Cilium numeric identity allocated for the node, if any.",
						MarkdownDescription: "NodeIdentity is the Cilium numeric identity allocated for the node, if any.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CiliumIoCiliumNodeV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cilium_io_cilium_node_v2")

	var state CiliumIoCiliumNodeV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumNodeV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2")
	goModel.Kind = utilities.Ptr("CiliumNode")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CiliumIoCiliumNodeV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_node_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CiliumIoCiliumNodeV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cilium_io_cilium_node_v2")

	var state CiliumIoCiliumNodeV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumNodeV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2")
	goModel.Kind = utilities.Ptr("CiliumNode")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CiliumIoCiliumNodeV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cilium_io_cilium_node_v2")
	// NO-OP: Terraform removes the state automatically for us
}
