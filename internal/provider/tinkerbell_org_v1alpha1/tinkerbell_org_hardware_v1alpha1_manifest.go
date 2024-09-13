/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tinkerbell_org_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &TinkerbellOrgHardwareV1Alpha1Manifest{}
)

func NewTinkerbellOrgHardwareV1Alpha1Manifest() datasource.DataSource {
	return &TinkerbellOrgHardwareV1Alpha1Manifest{}
}

type TinkerbellOrgHardwareV1Alpha1Manifest struct{}

type TinkerbellOrgHardwareV1Alpha1ManifestData struct {
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
		BmcRef *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"bmc_ref" json:"bmcRef,omitempty"`
		Disks *[]struct {
			Device *string `tfsdk:"device" json:"device,omitempty"`
		} `tfsdk:"disks" json:"disks,omitempty"`
		Interfaces *[]struct {
			Dhcp *struct {
				Arch       *string `tfsdk:"arch" json:"arch,omitempty"`
				Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
				Iface_name *string `tfsdk:"iface_name" json:"iface_name,omitempty"`
				Ip         *struct {
					Address *string `tfsdk:"address" json:"address,omitempty"`
					Family  *int64  `tfsdk:"family" json:"family,omitempty"`
					Gateway *string `tfsdk:"gateway" json:"gateway,omitempty"`
					Netmask *string `tfsdk:"netmask" json:"netmask,omitempty"`
				} `tfsdk:"ip" json:"ip,omitempty"`
				Lease_time   *int64    `tfsdk:"lease_time" json:"lease_time,omitempty"`
				Mac          *string   `tfsdk:"mac" json:"mac,omitempty"`
				Name_servers *[]string `tfsdk:"name_servers" json:"name_servers,omitempty"`
				Time_servers *[]string `tfsdk:"time_servers" json:"time_servers,omitempty"`
				Uefi         *bool     `tfsdk:"uefi" json:"uefi,omitempty"`
				Vlan_id      *string   `tfsdk:"vlan_id" json:"vlan_id,omitempty"`
			} `tfsdk:"dhcp" json:"dhcp,omitempty"`
			Netboot *struct {
				AllowPXE      *bool `tfsdk:"allow_pxe" json:"allowPXE,omitempty"`
				AllowWorkflow *bool `tfsdk:"allow_workflow" json:"allowWorkflow,omitempty"`
				Ipxe          *struct {
					Contents *string `tfsdk:"contents" json:"contents,omitempty"`
					Url      *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"ipxe" json:"ipxe,omitempty"`
				Osie *struct {
					BaseURL *string `tfsdk:"base_url" json:"baseURL,omitempty"`
					Initrd  *string `tfsdk:"initrd" json:"initrd,omitempty"`
					Kernel  *string `tfsdk:"kernel" json:"kernel,omitempty"`
				} `tfsdk:"osie" json:"osie,omitempty"`
			} `tfsdk:"netboot" json:"netboot,omitempty"`
		} `tfsdk:"interfaces" json:"interfaces,omitempty"`
		Metadata *struct {
			Bonding_mode *int64 `tfsdk:"bonding_mode" json:"bonding_mode,omitempty"`
			Custom       *struct {
				Preinstalled_operating_system_version *struct {
					Distro    *string `tfsdk:"distro" json:"distro,omitempty"`
					Image_tag *string `tfsdk:"image_tag" json:"image_tag,omitempty"`
					Os_slug   *string `tfsdk:"os_slug" json:"os_slug,omitempty"`
					Slug      *string `tfsdk:"slug" json:"slug,omitempty"`
					Version   *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"preinstalled_operating_system_version" json:"preinstalled_operating_system_version,omitempty"`
				Private_subnets *[]string `tfsdk:"private_subnets" json:"private_subnets,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Facility *struct {
				Facility_code     *string `tfsdk:"facility_code" json:"facility_code,omitempty"`
				Plan_slug         *string `tfsdk:"plan_slug" json:"plan_slug,omitempty"`
				Plan_version_slug *string `tfsdk:"plan_version_slug" json:"plan_version_slug,omitempty"`
			} `tfsdk:"facility" json:"facility,omitempty"`
			Instance *struct {
				Allow_pxe             *bool   `tfsdk:"allow_pxe" json:"allow_pxe,omitempty"`
				Always_pxe            *bool   `tfsdk:"always_pxe" json:"always_pxe,omitempty"`
				Crypted_root_password *string `tfsdk:"crypted_root_password" json:"crypted_root_password,omitempty"`
				Hostname              *string `tfsdk:"hostname" json:"hostname,omitempty"`
				Id                    *string `tfsdk:"id" json:"id,omitempty"`
				Ips                   *[]struct {
					Address    *string `tfsdk:"address" json:"address,omitempty"`
					Family     *int64  `tfsdk:"family" json:"family,omitempty"`
					Gateway    *string `tfsdk:"gateway" json:"gateway,omitempty"`
					Management *bool   `tfsdk:"management" json:"management,omitempty"`
					Netmask    *string `tfsdk:"netmask" json:"netmask,omitempty"`
					Public     *bool   `tfsdk:"public" json:"public,omitempty"`
				} `tfsdk:"ips" json:"ips,omitempty"`
				Ipxe_script_url  *string `tfsdk:"ipxe_script_url" json:"ipxe_script_url,omitempty"`
				Network_ready    *bool   `tfsdk:"network_ready" json:"network_ready,omitempty"`
				Operating_system *struct {
					Distro    *string `tfsdk:"distro" json:"distro,omitempty"`
					Image_tag *string `tfsdk:"image_tag" json:"image_tag,omitempty"`
					Os_slug   *string `tfsdk:"os_slug" json:"os_slug,omitempty"`
					Slug      *string `tfsdk:"slug" json:"slug,omitempty"`
					Version   *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"operating_system" json:"operating_system,omitempty"`
				Rescue   *bool     `tfsdk:"rescue" json:"rescue,omitempty"`
				Ssh_keys *[]string `tfsdk:"ssh_keys" json:"ssh_keys,omitempty"`
				State    *string   `tfsdk:"state" json:"state,omitempty"`
				Storage  *struct {
					Disks *[]struct {
						Device     *string `tfsdk:"device" json:"device,omitempty"`
						Partitions *[]struct {
							Label     *string `tfsdk:"label" json:"label,omitempty"`
							Number    *int64  `tfsdk:"number" json:"number,omitempty"`
							Size      *int64  `tfsdk:"size" json:"size,omitempty"`
							Start     *int64  `tfsdk:"start" json:"start,omitempty"`
							Type_guid *string `tfsdk:"type_guid" json:"type_guid,omitempty"`
						} `tfsdk:"partitions" json:"partitions,omitempty"`
						Wipe_table *bool `tfsdk:"wipe_table" json:"wipe_table,omitempty"`
					} `tfsdk:"disks" json:"disks,omitempty"`
					Filesystems *[]struct {
						Mount *struct {
							Create *struct {
								Force   *bool     `tfsdk:"force" json:"force,omitempty"`
								Options *[]string `tfsdk:"options" json:"options,omitempty"`
							} `tfsdk:"create" json:"create,omitempty"`
							Device *string `tfsdk:"device" json:"device,omitempty"`
							Files  *[]struct {
								Contents *string `tfsdk:"contents" json:"contents,omitempty"`
								Gid      *int64  `tfsdk:"gid" json:"gid,omitempty"`
								Mode     *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path     *string `tfsdk:"path" json:"path,omitempty"`
								Uid      *int64  `tfsdk:"uid" json:"uid,omitempty"`
							} `tfsdk:"files" json:"files,omitempty"`
							Format *string `tfsdk:"format" json:"format,omitempty"`
							Point  *string `tfsdk:"point" json:"point,omitempty"`
						} `tfsdk:"mount" json:"mount,omitempty"`
					} `tfsdk:"filesystems" json:"filesystems,omitempty"`
					Raid *[]struct {
						Devices *[]string `tfsdk:"devices" json:"devices,omitempty"`
						Level   *string   `tfsdk:"level" json:"level,omitempty"`
						Name    *string   `tfsdk:"name" json:"name,omitempty"`
						Spare   *int64    `tfsdk:"spare" json:"spare,omitempty"`
					} `tfsdk:"raid" json:"raid,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
				Tags     *[]string `tfsdk:"tags" json:"tags,omitempty"`
				Userdata *string   `tfsdk:"userdata" json:"userdata,omitempty"`
			} `tfsdk:"instance" json:"instance,omitempty"`
			Manufacturer *struct {
				Id   *string `tfsdk:"id" json:"id,omitempty"`
				Slug *string `tfsdk:"slug" json:"slug,omitempty"`
			} `tfsdk:"manufacturer" json:"manufacturer,omitempty"`
			State *string `tfsdk:"state" json:"state,omitempty"`
		} `tfsdk:"metadata" json:"metadata,omitempty"`
		Resources   *map[string]string `tfsdk:"resources" json:"resources,omitempty"`
		TinkVersion *int64             `tfsdk:"tink_version" json:"tinkVersion,omitempty"`
		UserData    *string            `tfsdk:"user_data" json:"userData,omitempty"`
		VendorData  *string            `tfsdk:"vendor_data" json:"vendorData,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TinkerbellOrgHardwareV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tinkerbell_org_hardware_v1alpha1_manifest"
}

func (r *TinkerbellOrgHardwareV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Hardware is the Schema for the Hardware API.",
		MarkdownDescription: "Hardware is the Schema for the Hardware API.",
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
				Description:         "HardwareSpec defines the desired state of Hardware.",
				MarkdownDescription: "HardwareSpec defines the desired state of Hardware.",
				Attributes: map[string]schema.Attribute{
					"bmc_ref": schema.SingleNestedAttribute{
						Description:         "BMCRef contains a relation to a BMC state management type in the same namespace as the Hardware. This may be used for BMC management by orchestrators.",
						MarkdownDescription: "BMCRef contains a relation to a BMC state management type in the same namespace as the Hardware. This may be used for BMC management by orchestrators.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
								MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is the type of resource being referenced",
								MarkdownDescription: "Kind is the type of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of resource being referenced",
								MarkdownDescription: "Name is the name of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disks": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"device": schema.StringAttribute{
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

					"interfaces": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"dhcp": schema.SingleNestedAttribute{
									Description:         "DHCP configuration.",
									MarkdownDescription: "DHCP configuration.",
									Attributes: map[string]schema.Attribute{
										"arch": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hostname": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"iface_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.SingleNestedAttribute{
											Description:         "IP configuration.",
											MarkdownDescription: "IP configuration.",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"family": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"gateway": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"netmask": schema.StringAttribute{
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

										"lease_time": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mac": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`([0-9a-f]{2}[:]){5}([0-9a-f]{2})`), ""),
											},
										},

										"name_servers": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"time_servers": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uefi": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"vlan_id": schema.StringAttribute{
											Description:         "validation pattern for VLANDID is a string number between 0-4096",
											MarkdownDescription: "validation pattern for VLANDID is a string number between 0-4096",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9][0-9]{0,2}|[1-3][0-9][0-9][0-9]|40([0-8][0-9]|9[0-6]))(,[1-9][0-9]{0,2}|[1-3][0-9][0-9][0-9]|40([0-8][0-9]|9[0-6]))*)$`), ""),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"netboot": schema.SingleNestedAttribute{
									Description:         "Netboot configuration.",
									MarkdownDescription: "Netboot configuration.",
									Attributes: map[string]schema.Attribute{
										"allow_pxe": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"allow_workflow": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ipxe": schema.SingleNestedAttribute{
											Description:         "IPXE configuration.",
											MarkdownDescription: "IPXE configuration.",
											Attributes: map[string]schema.Attribute{
												"contents": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"url": schema.StringAttribute{
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

										"osie": schema.SingleNestedAttribute{
											Description:         "OSIE configuration.",
											MarkdownDescription: "OSIE configuration.",
											Attributes: map[string]schema.Attribute{
												"base_url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"initrd": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kernel": schema.StringAttribute{
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

					"metadata": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"bonding_mode": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"preinstalled_operating_system_version": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"distro": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_tag": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"os_slug": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"slug": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"version": schema.StringAttribute{
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

									"private_subnets": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
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

							"facility": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"facility_code": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"plan_slug": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"plan_version_slug": schema.StringAttribute{
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

							"instance": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allow_pxe": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"always_pxe": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"crypted_root_password": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hostname": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ips": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"family": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"gateway": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"management": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"netmask": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"public": schema.BoolAttribute{
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

									"ipxe_script_url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"network_ready": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"operating_system": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"distro": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_tag": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"os_slug": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"slug": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"version": schema.StringAttribute{
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

									"rescue": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ssh_keys": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"state": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"disks": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"device": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"partitions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"label": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"number": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"size": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"start": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type_guid": schema.StringAttribute{
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

														"wipe_table": schema.BoolAttribute{
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

											"filesystems": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"mount": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"create": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"force": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"options": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																"device": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"files": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"contents": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"gid": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"mode": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"uid": schema.Int64Attribute{
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

																"format": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"point": schema.StringAttribute{
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
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"raid": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"devices": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"level": schema.StringAttribute{
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

														"spare": schema.Int64Attribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tags": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"userdata": schema.StringAttribute{
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

							"manufacturer": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"slug": schema.StringAttribute{
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

							"state": schema.StringAttribute{
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

					"resources": schema.MapAttribute{
						Description:         "Resources represents known resources that are available on a machine. Resources may be used for scheduling by orchestrators.",
						MarkdownDescription: "Resources represents known resources that are available on a machine. Resources may be used for scheduling by orchestrators.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tink_version": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_data": schema.StringAttribute{
						Description:         "UserData is the user data to configure in the hardware's metadata",
						MarkdownDescription: "UserData is the user data to configure in the hardware's metadata",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vendor_data": schema.StringAttribute{
						Description:         "VendorData is the vendor data to configure in the hardware's metadata",
						MarkdownDescription: "VendorData is the vendor data to configure in the hardware's metadata",
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

func (r *TinkerbellOrgHardwareV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tinkerbell_org_hardware_v1alpha1_manifest")

	var model TinkerbellOrgHardwareV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tinkerbell.org/v1alpha1")
	model.Kind = pointer.String("Hardware")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
