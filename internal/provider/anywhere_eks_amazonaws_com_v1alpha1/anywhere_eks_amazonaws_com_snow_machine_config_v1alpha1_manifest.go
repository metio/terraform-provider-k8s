/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComSnowMachineConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1ManifestData struct {
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
		AmiID            *string `tfsdk:"ami_id" json:"amiID,omitempty"`
		ContainersVolume *struct {
			DeviceName *string `tfsdk:"device_name" json:"deviceName,omitempty"`
			Size       *int64  `tfsdk:"size" json:"size,omitempty"`
			Type       *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"containers_volume" json:"containersVolume,omitempty"`
		Devices             *[]string `tfsdk:"devices" json:"devices,omitempty"`
		HostOSConfiguration *struct {
			BottlerocketConfiguration *struct {
				Boot *struct {
					BootKernelParameters *map[string][]string `tfsdk:"boot_kernel_parameters" json:"bootKernelParameters,omitempty"`
				} `tfsdk:"boot" json:"boot,omitempty"`
				Kernel *struct {
					SysctlSettings *map[string]string `tfsdk:"sysctl_settings" json:"sysctlSettings,omitempty"`
				} `tfsdk:"kernel" json:"kernel,omitempty"`
				Kubernetes *struct {
					AllowedUnsafeSysctls *[]string `tfsdk:"allowed_unsafe_sysctls" json:"allowedUnsafeSysctls,omitempty"`
					ClusterDNSIPs        *[]string `tfsdk:"cluster_dnsi_ps" json:"clusterDNSIPs,omitempty"`
					MaxPods              *int64    `tfsdk:"max_pods" json:"maxPods,omitempty"`
				} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
			} `tfsdk:"bottlerocket_configuration" json:"bottlerocketConfiguration,omitempty"`
			CertBundles *[]struct {
				Data *string `tfsdk:"data" json:"data,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"cert_bundles" json:"certBundles,omitempty"`
			NtpConfiguration *struct {
				Servers *[]string `tfsdk:"servers" json:"servers,omitempty"`
			} `tfsdk:"ntp_configuration" json:"ntpConfiguration,omitempty"`
		} `tfsdk:"host_os_configuration" json:"hostOSConfiguration,omitempty"`
		InstanceType *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
		Network      *struct {
			DirectNetworkInterfaces *[]struct {
				Dhcp      *bool  `tfsdk:"dhcp" json:"dhcp,omitempty"`
				Index     *int64 `tfsdk:"index" json:"index,omitempty"`
				IpPoolRef *struct {
					Kind *string `tfsdk:"kind" json:"kind,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ip_pool_ref" json:"ipPoolRef,omitempty"`
				Primary *bool  `tfsdk:"primary" json:"primary,omitempty"`
				VlanID  *int64 `tfsdk:"vlan_id" json:"vlanID,omitempty"`
			} `tfsdk:"direct_network_interfaces" json:"directNetworkInterfaces,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		NonRootVolumes *[]struct {
			DeviceName *string `tfsdk:"device_name" json:"deviceName,omitempty"`
			Size       *int64  `tfsdk:"size" json:"size,omitempty"`
			Type       *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"non_root_volumes" json:"nonRootVolumes,omitempty"`
		OsFamily                 *string `tfsdk:"os_family" json:"osFamily,omitempty"`
		PhysicalNetworkConnector *string `tfsdk:"physical_network_connector" json:"physicalNetworkConnector,omitempty"`
		SshKeyName               *string `tfsdk:"ssh_key_name" json:"sshKeyName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SnowMachineConfig is the Schema for the SnowMachineConfigs API.",
		MarkdownDescription: "SnowMachineConfig is the Schema for the SnowMachineConfigs API.",
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
				Description:         "SnowMachineConfigSpec defines the desired state of SnowMachineConfigSpec.",
				MarkdownDescription: "SnowMachineConfigSpec defines the desired state of SnowMachineConfigSpec.",
				Attributes: map[string]schema.Attribute{
					"ami_id": schema.StringAttribute{
						Description:         "The AMI ID from which to create the machine instance.",
						MarkdownDescription: "The AMI ID from which to create the machine instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"containers_volume": schema.SingleNestedAttribute{
						Description:         "ContainersVolume provides the configuration options for the containers data storage volume.",
						MarkdownDescription: "ContainersVolume provides the configuration options for the containers data storage volume.",
						Attributes: map[string]schema.Attribute{
							"device_name": schema.StringAttribute{
								Description:         "Device name",
								MarkdownDescription: "Device name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.Int64Attribute{
								Description:         "Size specifies size (in Gi) of the storage device. Must be greater than the image snapshot size or 8 (whichever is greater).",
								MarkdownDescription: "Size specifies size (in Gi) of the storage device. Must be greater than the image snapshot size or 8 (whichever is greater).",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(8),
								},
							},

							"type": schema.StringAttribute{
								Description:         "Type is the type of the volume (sbp1 for capacity-optimized HDD, sbg1 performance-optimized SSD, default is sbp1)",
								MarkdownDescription: "Type is the type of the volume (sbp1 for capacity-optimized HDD, sbg1 performance-optimized SSD, default is sbp1)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("sbp1", "sbg1"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"devices": schema.ListAttribute{
						Description:         "Devices contains a device ip list assigned by the user to provision machines.",
						MarkdownDescription: "Devices contains a device ip list assigned by the user to provision machines.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_os_configuration": schema.SingleNestedAttribute{
						Description:         "HostOSConfiguration provides OS specific configurations for the machine",
						MarkdownDescription: "HostOSConfiguration provides OS specific configurations for the machine",
						Attributes: map[string]schema.Attribute{
							"bottlerocket_configuration": schema.SingleNestedAttribute{
								Description:         "BottlerocketConfiguration defines the Bottlerocket configuration on the host OS. These settings only take effect when the 'osFamily' is bottlerocket.",
								MarkdownDescription: "BottlerocketConfiguration defines the Bottlerocket configuration on the host OS. These settings only take effect when the 'osFamily' is bottlerocket.",
								Attributes: map[string]schema.Attribute{
									"boot": schema.SingleNestedAttribute{
										Description:         "Boot defines the boot settings for bottlerocket.",
										MarkdownDescription: "Boot defines the boot settings for bottlerocket.",
										Attributes: map[string]schema.Attribute{
											"boot_kernel_parameters": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.ListType{ElemType: types.StringType},
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kernel": schema.SingleNestedAttribute{
										Description:         "Kernel defines the kernel settings for bottlerocket.",
										MarkdownDescription: "Kernel defines the kernel settings for bottlerocket.",
										Attributes: map[string]schema.Attribute{
											"sysctl_settings": schema.MapAttribute{
												Description:         "SysctlSettings defines the kernel sysctl settings to set for bottlerocket nodes.",
												MarkdownDescription: "SysctlSettings defines the kernel sysctl settings to set for bottlerocket nodes.",
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

									"kubernetes": schema.SingleNestedAttribute{
										Description:         "Kubernetes defines the Kubernetes settings on the host OS.",
										MarkdownDescription: "Kubernetes defines the Kubernetes settings on the host OS.",
										Attributes: map[string]schema.Attribute{
											"allowed_unsafe_sysctls": schema.ListAttribute{
												Description:         "AllowedUnsafeSysctls defines the list of unsafe sysctls that can be set on a node.",
												MarkdownDescription: "AllowedUnsafeSysctls defines the list of unsafe sysctls that can be set on a node.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_dnsi_ps": schema.ListAttribute{
												Description:         "ClusterDNSIPs defines IP addresses of the DNS servers.",
												MarkdownDescription: "ClusterDNSIPs defines IP addresses of the DNS servers.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_pods": schema.Int64Attribute{
												Description:         "MaxPods defines the maximum number of pods that can run on a node.",
												MarkdownDescription: "MaxPods defines the maximum number of pods that can run on a node.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cert_bundles": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"data": schema.StringAttribute{
											Description:         "Data defines the cert bundle data.",
											MarkdownDescription: "Data defines the cert bundle data.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name defines the cert bundle name.",
											MarkdownDescription: "Name defines the cert bundle name.",
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

							"ntp_configuration": schema.SingleNestedAttribute{
								Description:         "NTPConfiguration defines the NTP configuration on the host OS.",
								MarkdownDescription: "NTPConfiguration defines the NTP configuration on the host OS.",
								Attributes: map[string]schema.Attribute{
									"servers": schema.ListAttribute{
										Description:         "Servers defines a list of NTP servers to be configured on the host OS.",
										MarkdownDescription: "Servers defines a list of NTP servers to be configured on the host OS.",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
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

					"instance_type": schema.StringAttribute{
						Description:         "InstanceType is the type of instance to create.",
						MarkdownDescription: "InstanceType is the type of instance to create.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "Network provides the custom network setting for the machine.",
						MarkdownDescription: "Network provides the custom network setting for the machine.",
						Attributes: map[string]schema.Attribute{
							"direct_network_interfaces": schema.ListNestedAttribute{
								Description:         "DirectNetworkInterfaces contains a list of direct network interface (DNI) configuration.",
								MarkdownDescription: "DirectNetworkInterfaces contains a list of direct network interface (DNI) configuration.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dhcp": schema.BoolAttribute{
											Description:         "DHCP defines whether DHCP is used to assign ip for the DNI.",
											MarkdownDescription: "DHCP defines whether DHCP is used to assign ip for the DNI.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"index": schema.Int64Attribute{
											Description:         "Index is the index number of DNI used to clarify the position in the list. Usually starts with 1.",
											MarkdownDescription: "Index is the index number of DNI used to clarify the position in the list. Usually starts with 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
												int64validator.AtMost(8),
											},
										},

										"ip_pool_ref": schema.SingleNestedAttribute{
											Description:         "IPPool contains a reference to a snow ip pool which provides a range of ip addresses. When specified, an ip address selected from the pool is allocated to this DNI.",
											MarkdownDescription: "IPPool contains a reference to a snow ip pool which provides a range of ip addresses. When specified, an ip address selected from the pool is allocated to this DNI.",
											Attributes: map[string]schema.Attribute{
												"kind": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
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

										"primary": schema.BoolAttribute{
											Description:         "Primary indicates whether the DNI is primary or not.",
											MarkdownDescription: "Primary indicates whether the DNI is primary or not.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"vlan_id": schema.Int64Attribute{
											Description:         "VlanID is the vlan id assigned by the user for the DNI.",
											MarkdownDescription: "VlanID is the vlan id assigned by the user for the DNI.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(4095),
											},
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

					"non_root_volumes": schema.ListNestedAttribute{
						Description:         "NonRootVolumes provides the configuration options for the non root storage volumes.",
						MarkdownDescription: "NonRootVolumes provides the configuration options for the non root storage volumes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"device_name": schema.StringAttribute{
									Description:         "Device name",
									MarkdownDescription: "Device name",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"size": schema.Int64Attribute{
									Description:         "Size specifies size (in Gi) of the storage device. Must be greater than the image snapshot size or 8 (whichever is greater).",
									MarkdownDescription: "Size specifies size (in Gi) of the storage device. Must be greater than the image snapshot size or 8 (whichever is greater).",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(8),
									},
								},

								"type": schema.StringAttribute{
									Description:         "Type is the type of the volume (sbp1 for capacity-optimized HDD, sbg1 performance-optimized SSD, default is sbp1)",
									MarkdownDescription: "Type is the type of the volume (sbp1 for capacity-optimized HDD, sbg1 performance-optimized SSD, default is sbp1)",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("sbp1", "sbg1"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"os_family": schema.StringAttribute{
						Description:         "OSFamily is the node instance OS. Valid values: 'bottlerocket' and 'ubuntu'.",
						MarkdownDescription: "OSFamily is the node instance OS. Valid values: 'bottlerocket' and 'ubuntu'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"physical_network_connector": schema.StringAttribute{
						Description:         "PhysicalNetworkConnector is the physical network connector type to use for creating direct network interfaces (DNI). Valid values: 'SFP_PLUS' (default), 'QSFP' and 'RJ45'.",
						MarkdownDescription: "PhysicalNetworkConnector is the physical network connector type to use for creating direct network interfaces (DNI). Valid values: 'SFP_PLUS' (default), 'QSFP' and 'RJ45'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ssh_key_name": schema.StringAttribute{
						Description:         "SSHKeyName is the name of the ssh key defined in the aws snow key pairs, to attach to the instance.",
						MarkdownDescription: "SSHKeyName is the name of the ssh key defined in the aws snow key pairs, to attach to the instance.",
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

func (r *AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_snow_machine_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComSnowMachineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("SnowMachineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
