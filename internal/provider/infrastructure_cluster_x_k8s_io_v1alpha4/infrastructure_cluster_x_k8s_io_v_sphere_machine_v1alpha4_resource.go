/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha4

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource{}
	_ resource.ResourceWithConfigure   = &InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource{}
	_ resource.ResourceWithImportState = &InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource{}
)

func NewInfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource() resource.Resource {
	return &InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource{}
}

type InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CloneMode     *string            `tfsdk:"clone_mode" json:"cloneMode,omitempty"`
		CustomVMXKeys *map[string]string `tfsdk:"custom_vmx_keys" json:"customVMXKeys,omitempty"`
		Datacenter    *string            `tfsdk:"datacenter" json:"datacenter,omitempty"`
		Datastore     *string            `tfsdk:"datastore" json:"datastore,omitempty"`
		DiskGiB       *int64             `tfsdk:"disk_gi_b" json:"diskGiB,omitempty"`
		FailureDomain *string            `tfsdk:"failure_domain" json:"failureDomain,omitempty"`
		Folder        *string            `tfsdk:"folder" json:"folder,omitempty"`
		MemoryMiB     *int64             `tfsdk:"memory_mi_b" json:"memoryMiB,omitempty"`
		Network       *struct {
			Devices *[]struct {
				DeviceName  *string   `tfsdk:"device_name" json:"deviceName,omitempty"`
				Dhcp4       *bool     `tfsdk:"dhcp4" json:"dhcp4,omitempty"`
				Dhcp6       *bool     `tfsdk:"dhcp6" json:"dhcp6,omitempty"`
				Gateway4    *string   `tfsdk:"gateway4" json:"gateway4,omitempty"`
				Gateway6    *string   `tfsdk:"gateway6" json:"gateway6,omitempty"`
				IpAddrs     *[]string `tfsdk:"ip_addrs" json:"ipAddrs,omitempty"`
				MacAddr     *string   `tfsdk:"mac_addr" json:"macAddr,omitempty"`
				Mtu         *int64    `tfsdk:"mtu" json:"mtu,omitempty"`
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				NetworkName *string   `tfsdk:"network_name" json:"networkName,omitempty"`
				Routes      *[]struct {
					Metric *int64  `tfsdk:"metric" json:"metric,omitempty"`
					To     *string `tfsdk:"to" json:"to,omitempty"`
					Via    *string `tfsdk:"via" json:"via,omitempty"`
				} `tfsdk:"routes" json:"routes,omitempty"`
				SearchDomains *[]string `tfsdk:"search_domains" json:"searchDomains,omitempty"`
			} `tfsdk:"devices" json:"devices,omitempty"`
			PreferredAPIServerCidr *string `tfsdk:"preferred_api_server_cidr" json:"preferredAPIServerCidr,omitempty"`
			Routes                 *[]struct {
				Metric *int64  `tfsdk:"metric" json:"metric,omitempty"`
				To     *string `tfsdk:"to" json:"to,omitempty"`
				Via    *string `tfsdk:"via" json:"via,omitempty"`
			} `tfsdk:"routes" json:"routes,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		NumCPUs           *int64  `tfsdk:"num_cp_us" json:"numCPUs,omitempty"`
		NumCoresPerSocket *int64  `tfsdk:"num_cores_per_socket" json:"numCoresPerSocket,omitempty"`
		ProviderID        *string `tfsdk:"provider_id" json:"providerID,omitempty"`
		ResourcePool      *string `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
		Server            *string `tfsdk:"server" json:"server,omitempty"`
		Snapshot          *string `tfsdk:"snapshot" json:"snapshot,omitempty"`
		StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
		Template          *string `tfsdk:"template" json:"template,omitempty"`
		Thumbprint        *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4"
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereMachine is the Schema for the vspheremachines API  Deprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "VSphereMachine is the Schema for the vspheremachines API  Deprecated: This type will be removed in one of the next releases.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "VSphereMachineSpec defines the desired state of VSphereMachine",
				MarkdownDescription: "VSphereMachineSpec defines the desired state of VSphereMachine",
				Attributes: map[string]schema.Attribute{
					"clone_mode": schema.StringAttribute{
						Description:         "CloneMode specifies the type of clone operation. The LinkedClone mode is only support for templates that have at least one snapshot. If the template has no snapshots, then CloneMode defaults to FullClone. When LinkedClone mode is enabled the DiskGiB field is ignored as it is not possible to expand disks of linked clones. Defaults to LinkedClone, but fails gracefully to FullClone if the source of the clone operation has no snapshots.",
						MarkdownDescription: "CloneMode specifies the type of clone operation. The LinkedClone mode is only support for templates that have at least one snapshot. If the template has no snapshots, then CloneMode defaults to FullClone. When LinkedClone mode is enabled the DiskGiB field is ignored as it is not possible to expand disks of linked clones. Defaults to LinkedClone, but fails gracefully to FullClone if the source of the clone operation has no snapshots.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_vmx_keys": schema.MapAttribute{
						Description:         "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VM Defaults to empty map",
						MarkdownDescription: "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VM Defaults to empty map",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datacenter": schema.StringAttribute{
						Description:         "Datacenter is the name or inventory path of the datacenter in which the virtual machine is created/located.",
						MarkdownDescription: "Datacenter is the name or inventory path of the datacenter in which the virtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datastore": schema.StringAttribute{
						Description:         "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
						MarkdownDescription: "Datastore is the name or inventory path of the datastore in which the virtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disk_gi_b": schema.Int64Attribute{
						Description:         "DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						MarkdownDescription: "DiskGiB is the size of a virtual machine's disk, in GiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failure_domain": schema.StringAttribute{
						Description:         "FailureDomain is the failure domain unique identifier this Machine should be attached to, as defined in Cluster API. For this infrastructure provider, the name is equivalent to the name of the VSphereDeploymentZone.",
						MarkdownDescription: "FailureDomain is the failure domain unique identifier this Machine should be attached to, as defined in Cluster API. For this infrastructure provider, the name is equivalent to the name of the VSphereDeploymentZone.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"folder": schema.StringAttribute{
						Description:         "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
						MarkdownDescription: "Folder is the name or inventory path of the folder in which the virtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"memory_mi_b": schema.Int64Attribute{
						Description:         "MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						MarkdownDescription: "MemoryMiB is the size of a virtual machine's memory, in MiB. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "Network is the network configuration for this machine's VM.",
						MarkdownDescription: "Network is the network configuration for this machine's VM.",
						Attributes: map[string]schema.Attribute{
							"devices": schema.ListNestedAttribute{
								Description:         "Devices is the list of network devices used by the virtual machine. TODO(akutz) Make sure at least one network matches the ClusterSpec.CloudProviderConfiguration.Network.Name",
								MarkdownDescription: "Devices is the list of network devices used by the virtual machine. TODO(akutz) Make sure at least one network matches the ClusterSpec.CloudProviderConfiguration.Network.Name",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"device_name": schema.StringAttribute{
											Description:         "DeviceName may be used to explicitly assign a name to the network device as it exists in the guest operating system.",
											MarkdownDescription: "DeviceName may be used to explicitly assign a name to the network device as it exists in the guest operating system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dhcp4": schema.BoolAttribute{
											Description:         "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4 on this device. If true then IPAddrs should not contain any IPv4 addresses.",
											MarkdownDescription: "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4 on this device. If true then IPAddrs should not contain any IPv4 addresses.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dhcp6": schema.BoolAttribute{
											Description:         "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6 on this device. If true then IPAddrs should not contain any IPv6 addresses.",
											MarkdownDescription: "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6 on this device. If true then IPAddrs should not contain any IPv6 addresses.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"gateway4": schema.StringAttribute{
											Description:         "Gateway4 is the IPv4 gateway used by this device. Required when DHCP4 is false.",
											MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device. Required when DHCP4 is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"gateway6": schema.StringAttribute{
											Description:         "Gateway4 is the IPv4 gateway used by this device. Required when DHCP6 is false.",
											MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device. Required when DHCP6 is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip_addrs": schema.ListAttribute{
											Description:         "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign to this device. IP addresses must also specify the segment length in CIDR notation. Required when DHCP4 and DHCP6 are both false.",
											MarkdownDescription: "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assign to this device. IP addresses must also specify the segment length in CIDR notation. Required when DHCP4 and DHCP6 are both false.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mac_addr": schema.StringAttribute{
											Description:         "MACAddr is the MAC address used by this device. It is generally a good idea to omit this field and allow a MAC address to be generated. Please note that this value must use the VMware OUI to work with the in-tree vSphere cloud provider.",
											MarkdownDescription: "MACAddr is the MAC address used by this device. It is generally a good idea to omit this field and allow a MAC address to be generated. Please note that this value must use the VMware OUI to work with the in-tree vSphere cloud provider.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mtu": schema.Int64Attribute{
											Description:         "MTU is the device’s Maximum Transmission Unit size in bytes.",
											MarkdownDescription: "MTU is the device’s Maximum Transmission Unit size in bytes.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"nameservers": schema.ListAttribute{
											Description:         "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS nameservers. Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
											MarkdownDescription: "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNS nameservers. Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"network_name": schema.StringAttribute{
											Description:         "NetworkName is the name of the vSphere network to which the device will be connected.",
											MarkdownDescription: "NetworkName is the name of the vSphere network to which the device will be connected.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"routes": schema.ListNestedAttribute{
											Description:         "Routes is a list of optional, static routes applied to the device.",
											MarkdownDescription: "Routes is a list of optional, static routes applied to the device.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"metric": schema.Int64Attribute{
														Description:         "Metric is the weight/priority of the route.",
														MarkdownDescription: "Metric is the weight/priority of the route.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"to": schema.StringAttribute{
														Description:         "To is an IPv4 or IPv6 address.",
														MarkdownDescription: "To is an IPv4 or IPv6 address.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"via": schema.StringAttribute{
														Description:         "Via is an IPv4 or IPv6 address.",
														MarkdownDescription: "Via is an IPv4 or IPv6 address.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"search_domains": schema.ListAttribute{
											Description:         "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
											MarkdownDescription: "SearchDomains is a list of search domains used when resolving IP addresses with DNS.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"preferred_api_server_cidr": schema.StringAttribute{
								Description:         "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API server endpoint on this machine",
								MarkdownDescription: "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes API server endpoint on this machine",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"routes": schema.ListNestedAttribute{
								Description:         "Routes is a list of optional, static routes applied to the virtual machine.",
								MarkdownDescription: "Routes is a list of optional, static routes applied to the virtual machine.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metric": schema.Int64Attribute{
											Description:         "Metric is the weight/priority of the route.",
											MarkdownDescription: "Metric is the weight/priority of the route.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"to": schema.StringAttribute{
											Description:         "To is an IPv4 or IPv6 address.",
											MarkdownDescription: "To is an IPv4 or IPv6 address.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"via": schema.StringAttribute{
											Description:         "Via is an IPv4 or IPv6 address.",
											MarkdownDescription: "Via is an IPv4 or IPv6 address.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"num_cp_us": schema.Int64Attribute{
						Description:         "NumCPUs is the number of virtual processors in a virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						MarkdownDescription: "NumCPUs is the number of virtual processors in a virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"num_cores_per_socket": schema.Int64Attribute{
						Description:         "NumCPUs is the number of cores among which to distribute CPUs in this virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						MarkdownDescription: "NumCPUs is the number of cores among which to distribute CPUs in this virtual machine. Defaults to the eponymous property value in the template from which the virtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"provider_id": schema.StringAttribute{
						Description:         "ProviderID is the virtual machine's BIOS UUID formated as vsphere://12345678-1234-1234-1234-123456789abc",
						MarkdownDescription: "ProviderID is the virtual machine's BIOS UUID formated as vsphere://12345678-1234-1234-1234-123456789abc",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_pool": schema.StringAttribute{
						Description:         "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
						MarkdownDescription: "ResourcePool is the name or inventory path of the resource pool in which the virtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"server": schema.StringAttribute{
						Description:         "Server is the IP address or FQDN of the vSphere server on which the virtual machine is created/located.",
						MarkdownDescription: "Server is the IP address or FQDN of the vSphere server on which the virtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot": schema.StringAttribute{
						Description:         "Snapshot is the name of the snapshot from which to create a linked clone. This field is ignored if LinkedClone is not enabled. Defaults to the source's current snapshot.",
						MarkdownDescription: "Snapshot is the name of the snapshot from which to create a linked clone. This field is ignored if LinkedClone is not enabled. Defaults to the source's current snapshot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_policy_name": schema.StringAttribute{
						Description:         "StoragePolicyName of the storage policy to use with this Virtual Machine",
						MarkdownDescription: "StoragePolicyName of the storage policy to use with this Virtual Machine",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.StringAttribute{
						Description:         "Template is the name or inventory path of the template used to clone the virtual machine.",
						MarkdownDescription: "Template is the name or inventory path of the template used to clone the virtual machine.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"thumbprint": schema.StringAttribute{
						Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When this is set to empty, this VirtualMachine would be created without TLS certificate validation of the communication between Cluster API Provider vSphere and the VMware vCenter server.",
						MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificate When this is set to empty, this VirtualMachine would be created without TLS certificate validation of the communication between Cluster API Provider vSphere and the VMware vCenter server.",
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
	}
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4")

	var model InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha4")
	model.Kind = pointer.String("VSphereMachine")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha4", Resource: "vspheremachines"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4")

	var data InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha4", Resource: "vspheremachines"}).
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

	var readResponse InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4")

	var model InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha4")
	model.Kind = pointer.String("VSphereMachine")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha4", Resource: "vspheremachines"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_machine_v1alpha4")

	var data InfrastructureClusterXK8SIoVsphereMachineV1Alpha4ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha4", Resource: "vspheremachines"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1alpha4", Resource: "vspheremachines"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *InfrastructureClusterXK8SIoVsphereMachineV1Alpha4Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
