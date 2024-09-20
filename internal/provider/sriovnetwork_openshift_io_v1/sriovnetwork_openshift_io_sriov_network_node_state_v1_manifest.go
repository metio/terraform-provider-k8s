/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sriovnetwork_openshift_io_v1

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
	_ datasource.DataSource = &SriovnetworkOpenshiftIoSriovNetworkNodeStateV1Manifest{}
)

func NewSriovnetworkOpenshiftIoSriovNetworkNodeStateV1Manifest() datasource.DataSource {
	return &SriovnetworkOpenshiftIoSriovNetworkNodeStateV1Manifest{}
}

type SriovnetworkOpenshiftIoSriovNetworkNodeStateV1Manifest struct{}

type SriovnetworkOpenshiftIoSriovNetworkNodeStateV1ManifestData struct {
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
		Bridges *struct {
			Ovs *[]struct {
				Bridge *struct {
					DatapathType *string            `tfsdk:"datapath_type" json:"datapathType,omitempty"`
					ExternalIDs  *map[string]string `tfsdk:"external_i_ds" json:"externalIDs,omitempty"`
					OtherConfig  *map[string]string `tfsdk:"other_config" json:"otherConfig,omitempty"`
				} `tfsdk:"bridge" json:"bridge,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Uplinks *[]struct {
					Interface *struct {
						ExternalIDs *map[string]string `tfsdk:"external_i_ds" json:"externalIDs,omitempty"`
						Options     *map[string]string `tfsdk:"options" json:"options,omitempty"`
						OtherConfig *map[string]string `tfsdk:"other_config" json:"otherConfig,omitempty"`
						Type        *string            `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"interface" json:"interface,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					PciAddress *string `tfsdk:"pci_address" json:"pciAddress,omitempty"`
				} `tfsdk:"uplinks" json:"uplinks,omitempty"`
			} `tfsdk:"ovs" json:"ovs,omitempty"`
		} `tfsdk:"bridges" json:"bridges,omitempty"`
		Interfaces *[]struct {
			ESwitchMode       *string `tfsdk:"e_switch_mode" json:"eSwitchMode,omitempty"`
			ExternallyManaged *bool   `tfsdk:"externally_managed" json:"externallyManaged,omitempty"`
			LinkType          *string `tfsdk:"link_type" json:"linkType,omitempty"`
			Mtu               *int64  `tfsdk:"mtu" json:"mtu,omitempty"`
			Name              *string `tfsdk:"name" json:"name,omitempty"`
			NumVfs            *int64  `tfsdk:"num_vfs" json:"numVfs,omitempty"`
			PciAddress        *string `tfsdk:"pci_address" json:"pciAddress,omitempty"`
			VfGroups          *[]struct {
				DeviceType   *string `tfsdk:"device_type" json:"deviceType,omitempty"`
				IsRdma       *bool   `tfsdk:"is_rdma" json:"isRdma,omitempty"`
				Mtu          *int64  `tfsdk:"mtu" json:"mtu,omitempty"`
				PolicyName   *string `tfsdk:"policy_name" json:"policyName,omitempty"`
				ResourceName *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
				VdpaType     *string `tfsdk:"vdpa_type" json:"vdpaType,omitempty"`
				VfRange      *string `tfsdk:"vf_range" json:"vfRange,omitempty"`
			} `tfsdk:"vf_groups" json:"vfGroups,omitempty"`
		} `tfsdk:"interfaces" json:"interfaces,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SriovnetworkOpenshiftIoSriovNetworkNodeStateV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sriovnetwork_openshift_io_sriov_network_node_state_v1_manifest"
}

func (r *SriovnetworkOpenshiftIoSriovNetworkNodeStateV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SriovNetworkNodeState is the Schema for the sriovnetworknodestates API",
		MarkdownDescription: "SriovNetworkNodeState is the Schema for the sriovnetworknodestates API",
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
				Description:         "SriovNetworkNodeStateSpec defines the desired state of SriovNetworkNodeState",
				MarkdownDescription: "SriovNetworkNodeStateSpec defines the desired state of SriovNetworkNodeState",
				Attributes: map[string]schema.Attribute{
					"bridges": schema.SingleNestedAttribute{
						Description:         "Bridges contains list of bridges",
						MarkdownDescription: "Bridges contains list of bridges",
						Attributes: map[string]schema.Attribute{
							"ovs": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"bridge": schema.SingleNestedAttribute{
											Description:         "bridge-level configuration for the bridge",
											MarkdownDescription: "bridge-level configuration for the bridge",
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

										"name": schema.StringAttribute{
											Description:         "name of the bridge",
											MarkdownDescription: "name of the bridge",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"uplinks": schema.ListNestedAttribute{
											Description:         "uplink-level bridge configuration for each uplink(PF). currently must contain only one element",
											MarkdownDescription: "uplink-level bridge configuration for each uplink(PF). currently must contain only one element",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"interface": schema.SingleNestedAttribute{
														Description:         "configuration from the Interface OVS table for the PF",
														MarkdownDescription: "configuration from the Interface OVS table for the PF",
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

													"name": schema.StringAttribute{
														Description:         "name of the PF interface",
														MarkdownDescription: "name of the PF interface",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pci_address": schema.StringAttribute{
														Description:         "pci address of the PF",
														MarkdownDescription: "pci address of the PF",
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

					"interfaces": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"e_switch_mode": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"externally_managed": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"link_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"mtu": schema.Int64Attribute{
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

								"num_vfs": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"pci_address": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"vf_groups": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"device_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"is_rdma": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mtu": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"policy_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"vdpa_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"vf_range": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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
	}
}

func (r *SriovnetworkOpenshiftIoSriovNetworkNodeStateV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sriovnetwork_openshift_io_sriov_network_node_state_v1_manifest")

	var model SriovnetworkOpenshiftIoSriovNetworkNodeStateV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sriovnetwork.openshift.io/v1")
	model.Kind = pointer.String("SriovNetworkNodeState")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
