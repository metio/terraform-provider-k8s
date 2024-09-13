/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &Metal3IoHardwareDataV1Alpha1Manifest{}
)

func NewMetal3IoHardwareDataV1Alpha1Manifest() datasource.DataSource {
	return &Metal3IoHardwareDataV1Alpha1Manifest{}
}

type Metal3IoHardwareDataV1Alpha1Manifest struct{}

type Metal3IoHardwareDataV1Alpha1ManifestData struct {
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
		Hardware *struct {
			Cpu *struct {
				Arch           *string   `tfsdk:"arch" json:"arch,omitempty"`
				ClockMegahertz *float64  `tfsdk:"clock_megahertz" json:"clockMegahertz,omitempty"`
				Count          *int64    `tfsdk:"count" json:"count,omitempty"`
				Flags          *[]string `tfsdk:"flags" json:"flags,omitempty"`
				Model          *string   `tfsdk:"model" json:"model,omitempty"`
			} `tfsdk:"cpu" json:"cpu,omitempty"`
			Firmware *struct {
				Bios *struct {
					Date    *string `tfsdk:"date" json:"date,omitempty"`
					Vendor  *string `tfsdk:"vendor" json:"vendor,omitempty"`
					Version *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"bios" json:"bios,omitempty"`
			} `tfsdk:"firmware" json:"firmware,omitempty"`
			Hostname *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Nics     *[]struct {
				Ip        *string `tfsdk:"ip" json:"ip,omitempty"`
				Mac       *string `tfsdk:"mac" json:"mac,omitempty"`
				Model     *string `tfsdk:"model" json:"model,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Pxe       *bool   `tfsdk:"pxe" json:"pxe,omitempty"`
				SpeedGbps *int64  `tfsdk:"speed_gbps" json:"speedGbps,omitempty"`
				VlanId    *int64  `tfsdk:"vlan_id" json:"vlanId,omitempty"`
				Vlans     *[]struct {
					Id   *int64  `tfsdk:"id" json:"id,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"vlans" json:"vlans,omitempty"`
			} `tfsdk:"nics" json:"nics,omitempty"`
			RamMebibytes *int64 `tfsdk:"ram_mebibytes" json:"ramMebibytes,omitempty"`
			Storage      *[]struct {
				AlternateNames     *[]string `tfsdk:"alternate_names" json:"alternateNames,omitempty"`
				Hctl               *string   `tfsdk:"hctl" json:"hctl,omitempty"`
				Model              *string   `tfsdk:"model" json:"model,omitempty"`
				Name               *string   `tfsdk:"name" json:"name,omitempty"`
				Rotational         *bool     `tfsdk:"rotational" json:"rotational,omitempty"`
				SerialNumber       *string   `tfsdk:"serial_number" json:"serialNumber,omitempty"`
				SizeBytes          *int64    `tfsdk:"size_bytes" json:"sizeBytes,omitempty"`
				Type               *string   `tfsdk:"type" json:"type,omitempty"`
				Vendor             *string   `tfsdk:"vendor" json:"vendor,omitempty"`
				Wwn                *string   `tfsdk:"wwn" json:"wwn,omitempty"`
				WwnVendorExtension *string   `tfsdk:"wwn_vendor_extension" json:"wwnVendorExtension,omitempty"`
				WwnWithExtension   *string   `tfsdk:"wwn_with_extension" json:"wwnWithExtension,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			SystemVendor *struct {
				Manufacturer *string `tfsdk:"manufacturer" json:"manufacturer,omitempty"`
				ProductName  *string `tfsdk:"product_name" json:"productName,omitempty"`
				SerialNumber *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
			} `tfsdk:"system_vendor" json:"systemVendor,omitempty"`
		} `tfsdk:"hardware" json:"hardware,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Metal3IoHardwareDataV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metal3_io_hardware_data_v1alpha1_manifest"
}

func (r *Metal3IoHardwareDataV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HardwareData is the Schema for the hardwaredata API.",
		MarkdownDescription: "HardwareData is the Schema for the hardwaredata API.",
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
				Description:         "HardwareDataSpec defines the desired state of HardwareData.",
				MarkdownDescription: "HardwareDataSpec defines the desired state of HardwareData.",
				Attributes: map[string]schema.Attribute{
					"hardware": schema.SingleNestedAttribute{
						Description:         "The hardware discovered on the host during its inspection.",
						MarkdownDescription: "The hardware discovered on the host during its inspection.",
						Attributes: map[string]schema.Attribute{
							"cpu": schema.SingleNestedAttribute{
								Description:         "Details of the CPU(s) in the system.",
								MarkdownDescription: "Details of the CPU(s) in the system.",
								Attributes: map[string]schema.Attribute{
									"arch": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"clock_megahertz": schema.Float64Attribute{
										Description:         "ClockSpeed is a clock speed in MHz",
										MarkdownDescription: "ClockSpeed is a clock speed in MHz",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"count": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"flags": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"model": schema.StringAttribute{
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

							"firmware": schema.SingleNestedAttribute{
								Description:         "System firmware information.",
								MarkdownDescription: "System firmware information.",
								Attributes: map[string]schema.Attribute{
									"bios": schema.SingleNestedAttribute{
										Description:         "The BIOS for this firmware",
										MarkdownDescription: "The BIOS for this firmware",
										Attributes: map[string]schema.Attribute{
											"date": schema.StringAttribute{
												Description:         "The release/build date for this BIOS",
												MarkdownDescription: "The release/build date for this BIOS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"vendor": schema.StringAttribute{
												Description:         "The vendor name for this BIOS",
												MarkdownDescription: "The vendor name for this BIOS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"version": schema.StringAttribute{
												Description:         "The version of the BIOS",
												MarkdownDescription: "The version of the BIOS",
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

							"hostname": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"nics": schema.ListNestedAttribute{
								Description:         "List of network interfaces for the host.",
								MarkdownDescription: "List of network interfaces for the host.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description:         "The IP address of the interface. This will be an IPv4 or IPv6 address if one is present. If both IPv4 and IPv6 addresses are present in a dual-stack environment, two nics will be output, one with each IP.",
											MarkdownDescription: "The IP address of the interface. This will be an IPv4 or IPv6 address if one is present. If both IPv4 and IPv6 addresses are present in a dual-stack environment, two nics will be output, one with each IP.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mac": schema.StringAttribute{
											Description:         "The device MAC address",
											MarkdownDescription: "The device MAC address",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`), ""),
											},
										},

										"model": schema.StringAttribute{
											Description:         "The vendor and product IDs of the NIC, e.g. '0x8086 0x1572'",
											MarkdownDescription: "The vendor and product IDs of the NIC, e.g. '0x8086 0x1572'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "The name of the network interface, e.g. 'en0'",
											MarkdownDescription: "The name of the network interface, e.g. 'en0'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pxe": schema.BoolAttribute{
											Description:         "Whether the NIC is PXE Bootable",
											MarkdownDescription: "Whether the NIC is PXE Bootable",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"speed_gbps": schema.Int64Attribute{
											Description:         "The speed of the device in Gigabits per second",
											MarkdownDescription: "The speed of the device in Gigabits per second",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"vlan_id": schema.Int64Attribute{
											Description:         "The untagged VLAN ID",
											MarkdownDescription: "The untagged VLAN ID",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(4094),
											},
										},

										"vlans": schema.ListNestedAttribute{
											Description:         "The VLANs available",
											MarkdownDescription: "The VLANs available",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"id": schema.Int64Attribute{
														Description:         "VLANID is a 12-bit 802.1Q VLAN identifier",
														MarkdownDescription: "VLANID is a 12-bit 802.1Q VLAN identifier",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(4094),
														},
													},

													"name": schema.StringAttribute{
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

							"ram_mebibytes": schema.Int64Attribute{
								Description:         "The host's amount of memory in Mebibytes.",
								MarkdownDescription: "The host's amount of memory in Mebibytes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage": schema.ListNestedAttribute{
								Description:         "List of storage (disk, SSD, etc.) available to the host.",
								MarkdownDescription: "List of storage (disk, SSD, etc.) available to the host.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"alternate_names": schema.ListAttribute{
											Description:         "A list of alternate Linux device names of the disk, e.g. '/dev/sda'. Note that this list is not exhaustive, and names may not be stable across reboots.",
											MarkdownDescription: "A list of alternate Linux device names of the disk, e.g. '/dev/sda'. Note that this list is not exhaustive, and names may not be stable across reboots.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"hctl": schema.StringAttribute{
											Description:         "The SCSI location of the device",
											MarkdownDescription: "The SCSI location of the device",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"model": schema.StringAttribute{
											Description:         "Hardware model",
											MarkdownDescription: "Hardware model",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "A Linux device name of the disk, e.g. '/dev/disk/by-path/pci-0000:01:00.0-scsi-0:2:0:0'. This will be a name that is stable across reboots if one is available.",
											MarkdownDescription: "A Linux device name of the disk, e.g. '/dev/disk/by-path/pci-0000:01:00.0-scsi-0:2:0:0'. This will be a name that is stable across reboots if one is available.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rotational": schema.BoolAttribute{
											Description:         "Whether this disk represents rotational storage. This field is not recommended for usage, please prefer using 'Type' field instead, this field will be deprecated eventually.",
											MarkdownDescription: "Whether this disk represents rotational storage. This field is not recommended for usage, please prefer using 'Type' field instead, this field will be deprecated eventually.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"serial_number": schema.StringAttribute{
											Description:         "The serial number of the device",
											MarkdownDescription: "The serial number of the device",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size_bytes": schema.Int64Attribute{
											Description:         "The size of the disk in Bytes",
											MarkdownDescription: "The size of the disk in Bytes",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Device type, one of: HDD, SSD, NVME.",
											MarkdownDescription: "Device type, one of: HDD, SSD, NVME.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("HDD", "SSD", "NVME"),
											},
										},

										"vendor": schema.StringAttribute{
											Description:         "The name of the vendor of the device",
											MarkdownDescription: "The name of the vendor of the device",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wwn": schema.StringAttribute{
											Description:         "The WWN of the device",
											MarkdownDescription: "The WWN of the device",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wwn_vendor_extension": schema.StringAttribute{
											Description:         "The WWN Vendor extension of the device",
											MarkdownDescription: "The WWN Vendor extension of the device",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wwn_with_extension": schema.StringAttribute{
											Description:         "The WWN with the extension",
											MarkdownDescription: "The WWN with the extension",
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

							"system_vendor": schema.SingleNestedAttribute{
								Description:         "System vendor information.",
								MarkdownDescription: "System vendor information.",
								Attributes: map[string]schema.Attribute{
									"manufacturer": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"product_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"serial_number": schema.StringAttribute{
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *Metal3IoHardwareDataV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_metal3_io_hardware_data_v1alpha1_manifest")

	var model Metal3IoHardwareDataV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("metal3.io/v1alpha1")
	model.Kind = pointer.String("HardwareData")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
