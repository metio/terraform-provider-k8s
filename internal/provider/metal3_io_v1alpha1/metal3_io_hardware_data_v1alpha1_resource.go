/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	"regexp"
	"strings"
)

var (
	_ resource.Resource                = &Metal3IoHardwareDataV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &Metal3IoHardwareDataV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &Metal3IoHardwareDataV1Alpha1Resource{}
)

func NewMetal3IoHardwareDataV1Alpha1Resource() resource.Resource {
	return &Metal3IoHardwareDataV1Alpha1Resource{}
}

type Metal3IoHardwareDataV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type Metal3IoHardwareDataV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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
				Hctl               *string `tfsdk:"hctl" json:"hctl,omitempty"`
				Model              *string `tfsdk:"model" json:"model,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				Rotational         *bool   `tfsdk:"rotational" json:"rotational,omitempty"`
				SerialNumber       *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
				SizeBytes          *int64  `tfsdk:"size_bytes" json:"sizeBytes,omitempty"`
				Type               *string `tfsdk:"type" json:"type,omitempty"`
				Vendor             *string `tfsdk:"vendor" json:"vendor,omitempty"`
				Wwn                *string `tfsdk:"wwn" json:"wwn,omitempty"`
				WwnVendorExtension *string `tfsdk:"wwn_vendor_extension" json:"wwnVendorExtension,omitempty"`
				WwnWithExtension   *string `tfsdk:"wwn_with_extension" json:"wwnWithExtension,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			SystemVendor *struct {
				Manufacturer *string `tfsdk:"manufacturer" json:"manufacturer,omitempty"`
				ProductName  *string `tfsdk:"product_name" json:"productName,omitempty"`
				SerialNumber *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
			} `tfsdk:"system_vendor" json:"systemVendor,omitempty"`
		} `tfsdk:"hardware" json:"hardware,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Metal3IoHardwareDataV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metal3_io_hardware_data_v1alpha1"
}

func (r *Metal3IoHardwareDataV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "HardwareData is the Schema for the hardwaredata API",
		MarkdownDescription: "HardwareData is the Schema for the hardwaredata API",
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

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
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
				Description:         "HardwareDataSpec defines the desired state of HardwareData",
				MarkdownDescription: "HardwareDataSpec defines the desired state of HardwareData",
				Attributes: map[string]schema.Attribute{
					"hardware": schema.SingleNestedAttribute{
						Description:         "The hardware discovered on the host during its inspection.",
						MarkdownDescription: "The hardware discovered on the host during its inspection.",
						Attributes: map[string]schema.Attribute{
							"cpu": schema.SingleNestedAttribute{
								Description:         "CPU describes one processor on the host.",
								MarkdownDescription: "CPU describes one processor on the host.",
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
								Description:         "Firmware describes the firmware on the host.",
								MarkdownDescription: "Firmware describes the firmware on the host.",
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
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description:         "The IP address of the interface. This will be an IPv4 or IPv6 address if one is present.  If both IPv4 and IPv6 addresses are present in a dual-stack environment, two nics will be output, one with each IP.",
											MarkdownDescription: "The IP address of the interface. This will be an IPv4 or IPv6 address if one is present.  If both IPv4 and IPv6 addresses are present in a dual-stack environment, two nics will be output, one with each IP.",
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
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
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
											Description:         "The Linux device name of the disk, e.g. '/dev/sda'. Note that this may not be stable across reboots.",
											MarkdownDescription: "The Linux device name of the disk, e.g. '/dev/sda'. Note that this may not be stable across reboots.",
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
								Description:         "HardwareSystemVendor stores details about the whole hardware system.",
								MarkdownDescription: "HardwareSystemVendor stores details about the whole hardware system.",
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

func (r *Metal3IoHardwareDataV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *Metal3IoHardwareDataV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_metal3_io_hardware_data_v1alpha1")

	var model Metal3IoHardwareDataV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("metal3.io/v1alpha1")
	model.Kind = pointer.String("HardwareData")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "metal3.io", Version: "v1alpha1", Resource: "hardwaredata"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse Metal3IoHardwareDataV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *Metal3IoHardwareDataV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_metal3_io_hardware_data_v1alpha1")

	var data Metal3IoHardwareDataV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "metal3.io", Version: "v1alpha1", Resource: "hardwaredata"}).
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

	var readResponse Metal3IoHardwareDataV1Alpha1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *Metal3IoHardwareDataV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_metal3_io_hardware_data_v1alpha1")

	var model Metal3IoHardwareDataV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("metal3.io/v1alpha1")
	model.Kind = pointer.String("HardwareData")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "metal3.io", Version: "v1alpha1", Resource: "hardwaredata"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse Metal3IoHardwareDataV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *Metal3IoHardwareDataV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_metal3_io_hardware_data_v1alpha1")

	var data Metal3IoHardwareDataV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "metal3.io", Version: "v1alpha1", Resource: "hardwaredata"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *Metal3IoHardwareDataV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
