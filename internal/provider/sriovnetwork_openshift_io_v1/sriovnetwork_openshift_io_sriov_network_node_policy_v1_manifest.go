/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sriovnetwork_openshift_io_v1

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
	_ datasource.DataSource = &SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1Manifest{}
)

func NewSriovnetworkOpenshiftIoSriovNetworkNodePolicyV1Manifest() datasource.DataSource {
	return &SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1Manifest{}
}

type SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1Manifest struct{}

type SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1ManifestData struct {
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
		Bridge *struct {
			Ovs *struct {
				Bridge *struct {
					DatapathType *string            `tfsdk:"datapath_type" json:"datapathType,omitempty"`
					ExternalIDs  *map[string]string `tfsdk:"external_i_ds" json:"externalIDs,omitempty"`
					OtherConfig  *map[string]string `tfsdk:"other_config" json:"otherConfig,omitempty"`
				} `tfsdk:"bridge" json:"bridge,omitempty"`
				Uplink *struct {
					Interface *struct {
						ExternalIDs *map[string]string `tfsdk:"external_i_ds" json:"externalIDs,omitempty"`
						Options     *map[string]string `tfsdk:"options" json:"options,omitempty"`
						OtherConfig *map[string]string `tfsdk:"other_config" json:"otherConfig,omitempty"`
						Type        *string            `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"interface" json:"interface,omitempty"`
				} `tfsdk:"uplink" json:"uplink,omitempty"`
			} `tfsdk:"ovs" json:"ovs,omitempty"`
		} `tfsdk:"bridge" json:"bridge,omitempty"`
		DeviceType        *string `tfsdk:"device_type" json:"deviceType,omitempty"`
		ESwitchMode       *string `tfsdk:"e_switch_mode" json:"eSwitchMode,omitempty"`
		ExcludeTopology   *bool   `tfsdk:"exclude_topology" json:"excludeTopology,omitempty"`
		ExternallyManaged *bool   `tfsdk:"externally_managed" json:"externallyManaged,omitempty"`
		IsRdma            *bool   `tfsdk:"is_rdma" json:"isRdma,omitempty"`
		LinkType          *string `tfsdk:"link_type" json:"linkType,omitempty"`
		Mtu               *int64  `tfsdk:"mtu" json:"mtu,omitempty"`
		NeedVhostNet      *bool   `tfsdk:"need_vhost_net" json:"needVhostNet,omitempty"`
		NicSelector       *struct {
			DeviceID    *string   `tfsdk:"device_id" json:"deviceID,omitempty"`
			NetFilter   *string   `tfsdk:"net_filter" json:"netFilter,omitempty"`
			PfNames     *[]string `tfsdk:"pf_names" json:"pfNames,omitempty"`
			RootDevices *[]string `tfsdk:"root_devices" json:"rootDevices,omitempty"`
			Vendor      *string   `tfsdk:"vendor" json:"vendor,omitempty"`
		} `tfsdk:"nic_selector" json:"nicSelector,omitempty"`
		NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		NumVfs       *int64             `tfsdk:"num_vfs" json:"numVfs,omitempty"`
		Priority     *int64             `tfsdk:"priority" json:"priority,omitempty"`
		ResourceName *string            `tfsdk:"resource_name" json:"resourceName,omitempty"`
		VdpaType     *string            `tfsdk:"vdpa_type" json:"vdpaType,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sriovnetwork_openshift_io_sriov_network_node_policy_v1_manifest"
}

func (r *SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SriovNetworkNodePolicy is the Schema for the sriovnetworknodepolicies API",
		MarkdownDescription: "SriovNetworkNodePolicy is the Schema for the sriovnetworknodepolicies API",
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
				Description:         "SriovNetworkNodePolicySpec defines the desired state of SriovNetworkNodePolicy",
				MarkdownDescription: "SriovNetworkNodePolicySpec defines the desired state of SriovNetworkNodePolicy",
				Attributes: map[string]schema.Attribute{
					"bridge": schema.SingleNestedAttribute{
						Description:         "contains bridge configuration for matching PFs, valid only for eSwitchMode==switchdev",
						MarkdownDescription: "contains bridge configuration for matching PFs, valid only for eSwitchMode==switchdev",
						Attributes: map[string]schema.Attribute{
							"ovs": schema.SingleNestedAttribute{
								Description:         "contains configuration for the OVS bridge,",
								MarkdownDescription: "contains configuration for the OVS bridge,",
								Attributes: map[string]schema.Attribute{
									"bridge": schema.SingleNestedAttribute{
										Description:         "contains bridge level settings",
										MarkdownDescription: "contains bridge level settings",
										Attributes: map[string]schema.Attribute{
											"datapath_type": schema.StringAttribute{
												Description:         "configure datapath_type field in the Bridge table in OVSDB",
												MarkdownDescription: "configure datapath_type field in the Bridge table in OVSDB",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_i_ds": schema.MapAttribute{
												Description:         "IDs to inject to external_ids field in the Bridge table in OVSDB",
												MarkdownDescription: "IDs to inject to external_ids field in the Bridge table in OVSDB",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"other_config": schema.MapAttribute{
												Description:         "additional options to inject to other_config field in the bridge table in OVSDB",
												MarkdownDescription: "additional options to inject to other_config field in the bridge table in OVSDB",
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

									"uplink": schema.SingleNestedAttribute{
										Description:         "contains settings for uplink (PF)",
										MarkdownDescription: "contains settings for uplink (PF)",
										Attributes: map[string]schema.Attribute{
											"interface": schema.SingleNestedAttribute{
												Description:         "contains settings for PF interface in the OVS bridge",
												MarkdownDescription: "contains settings for PF interface in the OVS bridge",
												Attributes: map[string]schema.Attribute{
													"external_i_ds": schema.MapAttribute{
														Description:         "external_ids field in the Interface table in OVSDB",
														MarkdownDescription: "external_ids field in the Interface table in OVSDB",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"options": schema.MapAttribute{
														Description:         "options field in the Interface table in OVSDB",
														MarkdownDescription: "options field in the Interface table in OVSDB",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"other_config": schema.MapAttribute{
														Description:         "other_config field in the Interface table in OVSDB",
														MarkdownDescription: "other_config field in the Interface table in OVSDB",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type field in the Interface table in OVSDB",
														MarkdownDescription: "type field in the Interface table in OVSDB",
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

					"device_type": schema.StringAttribute{
						Description:         "The driver type for configured VFs. Allowed value 'netdevice', 'vfio-pci'. Defaults to netdevice.",
						MarkdownDescription: "The driver type for configured VFs. Allowed value 'netdevice', 'vfio-pci'. Defaults to netdevice.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("netdevice", "vfio-pci"),
						},
					},

					"e_switch_mode": schema.StringAttribute{
						Description:         "NIC Device Mode. Allowed value 'legacy','switchdev'.",
						MarkdownDescription: "NIC Device Mode. Allowed value 'legacy','switchdev'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("legacy", "switchdev"),
						},
					},

					"exclude_topology": schema.BoolAttribute{
						Description:         "Exclude device's NUMA node when advertising this resource by SRIOV network device plugin. Default to false.",
						MarkdownDescription: "Exclude device's NUMA node when advertising this resource by SRIOV network device plugin. Default to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"externally_managed": schema.BoolAttribute{
						Description:         "don't create the virtual function only allocated them to the device plugin. Defaults to false.",
						MarkdownDescription: "don't create the virtual function only allocated them to the device plugin. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"is_rdma": schema.BoolAttribute{
						Description:         "RDMA mode. Defaults to false.",
						MarkdownDescription: "RDMA mode. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"link_type": schema.StringAttribute{
						Description:         "NIC Link Type. Allowed value 'eth', 'ETH', 'ib', and 'IB'.",
						MarkdownDescription: "NIC Link Type. Allowed value 'eth', 'ETH', 'ib', and 'IB'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("eth", "ETH", "ib", "IB"),
						},
					},

					"mtu": schema.Int64Attribute{
						Description:         "MTU of VF",
						MarkdownDescription: "MTU of VF",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"need_vhost_net": schema.BoolAttribute{
						Description:         "mount vhost-net device. Defaults to false.",
						MarkdownDescription: "mount vhost-net device. Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nic_selector": schema.SingleNestedAttribute{
						Description:         "NicSelector selects the NICs to be configured",
						MarkdownDescription: "NicSelector selects the NICs to be configured",
						Attributes: map[string]schema.Attribute{
							"device_id": schema.StringAttribute{
								Description:         "The device hex code of SR-IoV device. Allowed value '0d58', '1572', '158b', '1013', '1015', '1017', '101b'.",
								MarkdownDescription: "The device hex code of SR-IoV device. Allowed value '0d58', '1572', '158b', '1013', '1015', '1017', '101b'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"net_filter": schema.StringAttribute{
								Description:         "Infrastructure Networking selection filter. Allowed value 'openstack/NetworkID:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'",
								MarkdownDescription: "Infrastructure Networking selection filter. Allowed value 'openstack/NetworkID:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pf_names": schema.ListAttribute{
								Description:         "Name of SR-IoV PF.",
								MarkdownDescription: "Name of SR-IoV PF.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"root_devices": schema.ListAttribute{
								Description:         "PCI address of SR-IoV PF.",
								MarkdownDescription: "PCI address of SR-IoV PF.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vendor": schema.StringAttribute{
								Description:         "The vendor hex code of SR-IoV device. Allowed value '8086', '15b3'.",
								MarkdownDescription: "The vendor hex code of SR-IoV device. Allowed value '8086', '15b3'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "NodeSelector selects the nodes to be configured",
						MarkdownDescription: "NodeSelector selects the nodes to be configured",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"num_vfs": schema.Int64Attribute{
						Description:         "Number of VFs for each PF",
						MarkdownDescription: "Number of VFs for each PF",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"priority": schema.Int64Attribute{
						Description:         "Priority of the policy, higher priority policies can override lower ones.",
						MarkdownDescription: "Priority of the policy, higher priority policies can override lower ones.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(99),
						},
					},

					"resource_name": schema.StringAttribute{
						Description:         "SRIOV Network device plugin endpoint resource name",
						MarkdownDescription: "SRIOV Network device plugin endpoint resource name",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"vdpa_type": schema.StringAttribute{
						Description:         "VDPA device type. Allowed value 'virtio', 'vhost'",
						MarkdownDescription: "VDPA device type. Allowed value 'virtio', 'vhost'",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("virtio", "vhost"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sriovnetwork_openshift_io_sriov_network_node_policy_v1_manifest")

	var model SriovnetworkOpenshiftIoSriovNetworkNodePolicyV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sriovnetwork.openshift.io/v1")
	model.Kind = pointer.String("SriovNetworkNodePolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
