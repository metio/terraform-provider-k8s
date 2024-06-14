/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1alpha4

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoVsphereVmV1Alpha4Manifest{}
)

func NewInfrastructureClusterXK8SIoVsphereVmV1Alpha4Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoVsphereVmV1Alpha4Manifest{}
}

type InfrastructureClusterXK8SIoVsphereVmV1Alpha4Manifest struct{}

type InfrastructureClusterXK8SIoVsphereVmV1Alpha4ManifestData struct {
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
		BiosUUID     *string `tfsdk:"bios_uuid" json:"biosUUID,omitempty"`
		BootstrapRef *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"bootstrap_ref" json:"bootstrapRef,omitempty"`
		CloneMode     *string            `tfsdk:"clone_mode" json:"cloneMode,omitempty"`
		CustomVMXKeys *map[string]string `tfsdk:"custom_vmx_keys" json:"customVMXKeys,omitempty"`
		Datacenter    *string            `tfsdk:"datacenter" json:"datacenter,omitempty"`
		Datastore     *string            `tfsdk:"datastore" json:"datastore,omitempty"`
		DiskGiB       *int64             `tfsdk:"disk_gi_b" json:"diskGiB,omitempty"`
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
		ResourcePool      *string `tfsdk:"resource_pool" json:"resourcePool,omitempty"`
		Server            *string `tfsdk:"server" json:"server,omitempty"`
		Snapshot          *string `tfsdk:"snapshot" json:"snapshot,omitempty"`
		StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
		Template          *string `tfsdk:"template" json:"template,omitempty"`
		Thumbprint        *string `tfsdk:"thumbprint" json:"thumbprint,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoVsphereVmV1Alpha4Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha4_manifest"
}

func (r *InfrastructureClusterXK8SIoVsphereVmV1Alpha4Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VSphereVM is the Schema for the vspherevms APIDeprecated: This type will be removed in one of the next releases.",
		MarkdownDescription: "VSphereVM is the Schema for the vspherevms APIDeprecated: This type will be removed in one of the next releases.",
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
				Description:         "VSphereVMSpec defines the desired state of VSphereVM.",
				MarkdownDescription: "VSphereVMSpec defines the desired state of VSphereVM.",
				Attributes: map[string]schema.Attribute{
					"bios_uuid": schema.StringAttribute{
						Description:         "BiosUUID is the VM's BIOS UUID that is assigned at runtime afterthe VM has been created.This field is required at runtime for other controllers that readthis CRD as unstructured data.",
						MarkdownDescription: "BiosUUID is the VM's BIOS UUID that is assigned at runtime afterthe VM has been created.This field is required at runtime for other controllers that readthis CRD as unstructured data.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bootstrap_ref": schema.SingleNestedAttribute{
						Description:         "BootstrapRef is a reference to a bootstrap provider-specific resourcethat holds configuration details.This field is optional in case no bootstrap data is required to createa VM.",
						MarkdownDescription: "BootstrapRef is a reference to a bootstrap provider-specific resourcethat holds configuration details.This field is optional in case no bootstrap data is required to createa VM.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"clone_mode": schema.StringAttribute{
						Description:         "CloneMode specifies the type of clone operation.The LinkedClone mode is only support for templates that have at leastone snapshot. If the template has no snapshots, then CloneMode defaultsto FullClone.When LinkedClone mode is enabled the DiskGiB field is ignored as it isnot possible to expand disks of linked clones.Defaults to LinkedClone, but fails gracefully to FullClone if the sourceof the clone operation has no snapshots.",
						MarkdownDescription: "CloneMode specifies the type of clone operation.The LinkedClone mode is only support for templates that have at leastone snapshot. If the template has no snapshots, then CloneMode defaultsto FullClone.When LinkedClone mode is enabled the DiskGiB field is ignored as it isnot possible to expand disks of linked clones.Defaults to LinkedClone, but fails gracefully to FullClone if the sourceof the clone operation has no snapshots.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"custom_vmx_keys": schema.MapAttribute{
						Description:         "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VMDefaults to empty map",
						MarkdownDescription: "CustomVMXKeys is a dictionary of advanced VMX options that can be set on VMDefaults to empty map",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datacenter": schema.StringAttribute{
						Description:         "Datacenter is the name or inventory path of the datacenter in which thevirtual machine is created/located.",
						MarkdownDescription: "Datacenter is the name or inventory path of the datacenter in which thevirtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"datastore": schema.StringAttribute{
						Description:         "Datastore is the name or inventory path of the datastore in which thevirtual machine is created/located.",
						MarkdownDescription: "Datastore is the name or inventory path of the datastore in which thevirtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disk_gi_b": schema.Int64Attribute{
						Description:         "DiskGiB is the size of a virtual machine's disk, in GiB.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						MarkdownDescription: "DiskGiB is the size of a virtual machine's disk, in GiB.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"folder": schema.StringAttribute{
						Description:         "Folder is the name or inventory path of the folder in which thevirtual machine is created/located.",
						MarkdownDescription: "Folder is the name or inventory path of the folder in which thevirtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"memory_mi_b": schema.Int64Attribute{
						Description:         "MemoryMiB is the size of a virtual machine's memory, in MiB.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						MarkdownDescription: "MemoryMiB is the size of a virtual machine's memory, in MiB.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "Network is the network configuration for this machine's VM.",
						MarkdownDescription: "Network is the network configuration for this machine's VM.",
						Attributes: map[string]schema.Attribute{
							"devices": schema.ListNestedAttribute{
								Description:         "Devices is the list of network devices used by the virtual machine.TODO(akutz) Make sure at least one network matches the            ClusterSpec.CloudProviderConfiguration.Network.Name",
								MarkdownDescription: "Devices is the list of network devices used by the virtual machine.TODO(akutz) Make sure at least one network matches the            ClusterSpec.CloudProviderConfiguration.Network.Name",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"device_name": schema.StringAttribute{
											Description:         "DeviceName may be used to explicitly assign a name to the network deviceas it exists in the guest operating system.",
											MarkdownDescription: "DeviceName may be used to explicitly assign a name to the network deviceas it exists in the guest operating system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dhcp4": schema.BoolAttribute{
											Description:         "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4on this device.If true then IPAddrs should not contain any IPv4 addresses.",
											MarkdownDescription: "DHCP4 is a flag that indicates whether or not to use DHCP for IPv4on this device.If true then IPAddrs should not contain any IPv4 addresses.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dhcp6": schema.BoolAttribute{
											Description:         "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6on this device.If true then IPAddrs should not contain any IPv6 addresses.",
											MarkdownDescription: "DHCP6 is a flag that indicates whether or not to use DHCP for IPv6on this device.If true then IPAddrs should not contain any IPv6 addresses.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"gateway4": schema.StringAttribute{
											Description:         "Gateway4 is the IPv4 gateway used by this device.Required when DHCP4 is false.",
											MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device.Required when DHCP4 is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"gateway6": schema.StringAttribute{
											Description:         "Gateway4 is the IPv4 gateway used by this device.Required when DHCP6 is false.",
											MarkdownDescription: "Gateway4 is the IPv4 gateway used by this device.Required when DHCP6 is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip_addrs": schema.ListAttribute{
											Description:         "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assignto this device. IP addresses must also specify the segment length inCIDR notation.Required when DHCP4 and DHCP6 are both false.",
											MarkdownDescription: "IPAddrs is a list of one or more IPv4 and/or IPv6 addresses to assignto this device. IP addresses must also specify the segment length inCIDR notation.Required when DHCP4 and DHCP6 are both false.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mac_addr": schema.StringAttribute{
											Description:         "MACAddr is the MAC address used by this device.It is generally a good idea to omit this field and allow a MAC addressto be generated.Please note that this value must use the VMware OUI to work with thein-tree vSphere cloud provider.",
											MarkdownDescription: "MACAddr is the MAC address used by this device.It is generally a good idea to omit this field and allow a MAC addressto be generated.Please note that this value must use the VMware OUI to work with thein-tree vSphere cloud provider.",
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
											Description:         "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNSnameservers.Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
											MarkdownDescription: "Nameservers is a list of IPv4 and/or IPv6 addresses used as DNSnameservers.Please note that Linux allows only three nameservers (https://linux.die.net/man/5/resolv.conf).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"network_name": schema.StringAttribute{
											Description:         "NetworkName is the name of the vSphere network to which the devicewill be connected.",
											MarkdownDescription: "NetworkName is the name of the vSphere network to which the devicewill be connected.",
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
											Description:         "SearchDomains is a list of search domains used when resolving IPaddresses with DNS.",
											MarkdownDescription: "SearchDomains is a list of search domains used when resolving IPaddresses with DNS.",
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
								Description:         "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes APIserver endpoint on this machine",
								MarkdownDescription: "PreferredAPIServeCIDR is the preferred CIDR for the Kubernetes APIserver endpoint on this machine",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"routes": schema.ListNestedAttribute{
								Description:         "Routes is a list of optional, static routes applied to the virtualmachine.",
								MarkdownDescription: "Routes is a list of optional, static routes applied to the virtualmachine.",
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
						Description:         "NumCPUs is the number of virtual processors in a virtual machine.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						MarkdownDescription: "NumCPUs is the number of virtual processors in a virtual machine.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"num_cores_per_socket": schema.Int64Attribute{
						Description:         "NumCPUs is the number of cores among which to distribute CPUs in thisvirtual machine.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						MarkdownDescription: "NumCPUs is the number of cores among which to distribute CPUs in thisvirtual machine.Defaults to the eponymous property value in the template from which thevirtual machine is cloned.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_pool": schema.StringAttribute{
						Description:         "ResourcePool is the name or inventory path of the resource pool in whichthe virtual machine is created/located.",
						MarkdownDescription: "ResourcePool is the name or inventory path of the resource pool in whichthe virtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"server": schema.StringAttribute{
						Description:         "Server is the IP address or FQDN of the vSphere server on whichthe virtual machine is created/located.",
						MarkdownDescription: "Server is the IP address or FQDN of the vSphere server on whichthe virtual machine is created/located.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"snapshot": schema.StringAttribute{
						Description:         "Snapshot is the name of the snapshot from which to create a linked clone.This field is ignored if LinkedClone is not enabled.Defaults to the source's current snapshot.",
						MarkdownDescription: "Snapshot is the name of the snapshot from which to create a linked clone.This field is ignored if LinkedClone is not enabled.Defaults to the source's current snapshot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_policy_name": schema.StringAttribute{
						Description:         "StoragePolicyName of the storage policy to use with thisVirtual Machine",
						MarkdownDescription: "StoragePolicyName of the storage policy to use with thisVirtual Machine",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.StringAttribute{
						Description:         "Template is the name or inventory path of the template used to clonethe virtual machine.",
						MarkdownDescription: "Template is the name or inventory path of the template used to clonethe virtual machine.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"thumbprint": schema.StringAttribute{
						Description:         "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificateWhen this is set to empty, this VirtualMachine would be createdwithout TLS certificate validation of the communication between Cluster API Provider vSphereand the VMware vCenter server.",
						MarkdownDescription: "Thumbprint is the colon-separated SHA-1 checksum of the given vCenter server's host certificateWhen this is set to empty, this VirtualMachine would be createdwithout TLS certificate validation of the communication between Cluster API Provider vSphereand the VMware vCenter server.",
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

func (r *InfrastructureClusterXK8SIoVsphereVmV1Alpha4Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_v_sphere_vm_v1alpha4_manifest")

	var model InfrastructureClusterXK8SIoVsphereVmV1Alpha4ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1alpha4")
	model.Kind = pointer.String("VSphereVM")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
