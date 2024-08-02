/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package tinkerbell_org_v1alpha2

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
	_ datasource.DataSource = &TinkerbellOrgHardwareV1Alpha2Manifest{}
)

func NewTinkerbellOrgHardwareV1Alpha2Manifest() datasource.DataSource {
	return &TinkerbellOrgHardwareV1Alpha2Manifest{}
}

type TinkerbellOrgHardwareV1Alpha2Manifest struct{}

type TinkerbellOrgHardwareV1Alpha2ManifestData struct {
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
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"bmc_ref" json:"bmcRef,omitempty"`
		Instance *struct {
			Userdata   *string `tfsdk:"userdata" json:"userdata,omitempty"`
			Vendordata *string `tfsdk:"vendordata" json:"vendordata,omitempty"`
		} `tfsdk:"instance" json:"instance,omitempty"`
		Ipxe *struct {
			Inline *string `tfsdk:"inline" json:"inline,omitempty"`
			Url    *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"ipxe" json:"ipxe,omitempty"`
		KernelParams      *[]string `tfsdk:"kernel_params" json:"kernelParams,omitempty"`
		NetworkInterfaces *struct {
			Dhcp *struct {
				Gateway          *string   `tfsdk:"gateway" json:"gateway,omitempty"`
				Hostname         *string   `tfsdk:"hostname" json:"hostname,omitempty"`
				Ip               *string   `tfsdk:"ip" json:"ip,omitempty"`
				LeaseTimeSeconds *int64    `tfsdk:"lease_time_seconds" json:"leaseTimeSeconds,omitempty"`
				Nameservers      *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Netmask          *string   `tfsdk:"netmask" json:"netmask,omitempty"`
				Timeservers      *[]string `tfsdk:"timeservers" json:"timeservers,omitempty"`
				VlanId           *string   `tfsdk:"vlan_id" json:"vlanId,omitempty"`
			} `tfsdk:"dhcp" json:"dhcp,omitempty"`
			DisableDhcp    *bool `tfsdk:"disable_dhcp" json:"disableDhcp,omitempty"`
			DisableNetboot *bool `tfsdk:"disable_netboot" json:"disableNetboot,omitempty"`
		} `tfsdk:"network_interfaces" json:"networkInterfaces,omitempty"`
		Osie *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"osie" json:"osie,omitempty"`
		StorageDevices *[]string `tfsdk:"storage_devices" json:"storageDevices,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TinkerbellOrgHardwareV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_tinkerbell_org_hardware_v1alpha2_manifest"
}

func (r *TinkerbellOrgHardwareV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Hardware is a logical representation of a machine that can execute Workflows.",
		MarkdownDescription: "Hardware is a logical representation of a machine that can execute Workflows.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"bmc_ref": schema.SingleNestedAttribute{
						Description:         "BMCRef references a Rufio Machine object.",
						MarkdownDescription: "BMCRef references a Rufio Machine object.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
						Description:         "Instance describes instance specific data that is generally unused by Tinkerbell core.",
						MarkdownDescription: "Instance describes instance specific data that is generally unused by Tinkerbell core.",
						Attributes: map[string]schema.Attribute{
							"userdata": schema.StringAttribute{
								Description:         "Userdata is data with a structure understood by the producer and consumer of the data.",
								MarkdownDescription: "Userdata is data with a structure understood by the producer and consumer of the data.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vendordata": schema.StringAttribute{
								Description:         "Vendordata is data with a structure understood by the producer and consumer of the data.",
								MarkdownDescription: "Vendordata is data with a structure understood by the producer and consumer of the data.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipxe": schema.SingleNestedAttribute{
						Description:         "IPXE provides iPXE script override fields. This is useful for debugging or netbootcustomization.",
						MarkdownDescription: "IPXE provides iPXE script override fields. This is useful for debugging or netbootcustomization.",
						Attributes: map[string]schema.Attribute{
							"inline": schema.StringAttribute{
								Description:         "Content is an inline iPXE script.",
								MarkdownDescription: "Content is an inline iPXE script.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "URL is a URL to a hosted iPXE script.",
								MarkdownDescription: "URL is a URL to a hosted iPXE script.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kernel_params": schema.ListAttribute{
						Description:         "KernelParams passed to the kernel when launching the OSIE. Parameters are joined with aspace.",
						MarkdownDescription: "KernelParams passed to the kernel when launching the OSIE. Parameters are joined with aspace.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network_interfaces": schema.SingleNestedAttribute{
						Description:         "NetworkInterfaces defines the desired DHCP and netboot configuration for a network interface.",
						MarkdownDescription: "NetworkInterfaces defines the desired DHCP and netboot configuration for a network interface.",
						Attributes: map[string]schema.Attribute{
							"dhcp": schema.SingleNestedAttribute{
								Description:         "DHCP is the basic network information for serving DHCP requests. Required when DisbaleDHCPis false.",
								MarkdownDescription: "DHCP is the basic network information for serving DHCP requests. Required when DisbaleDHCPis false.",
								Attributes: map[string]schema.Attribute{
									"gateway": schema.StringAttribute{
										Description:         "Gateway is the default gateway address to serve.",
										MarkdownDescription: "Gateway is the default gateway address to serve.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`), ""),
										},
									},

									"hostname": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9]"[A-Za-z0-9\-]*[A-Za-z0-9])$`), ""),
										},
									},

									"ip": schema.StringAttribute{
										Description:         "IP is an IPv4 address to serve.",
										MarkdownDescription: "IP is an IPv4 address to serve.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`), ""),
										},
									},

									"lease_time_seconds": schema.Int64Attribute{
										Description:         "LeaseTimeSeconds to serve. 24h default. Maximum equates to max uint32 as defined by RFC 2132§ 9.2 (https://www.rfc-editor.org/rfc/rfc2132.html#section-9.2).",
										MarkdownDescription: "LeaseTimeSeconds to serve. 24h default. Maximum equates to max uint32 as defined by RFC 2132§ 9.2 (https://www.rfc-editor.org/rfc/rfc2132.html#section-9.2).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(4.294967295e+09),
										},
									},

									"nameservers": schema.ListAttribute{
										Description:         "Nameservers to serve.",
										MarkdownDescription: "Nameservers to serve.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"netmask": schema.StringAttribute{
										Description:         "Netmask is an IPv4 netmask to serve.",
										MarkdownDescription: "Netmask is an IPv4 netmask to serve.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeservers": schema.ListAttribute{
										Description:         "Timeservers to serve.",
										MarkdownDescription: "Timeservers to serve.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vlan_id": schema.StringAttribute{
										Description:         "VLANID is a VLAN ID between 0 and 4096.",
										MarkdownDescription: "VLANID is a VLAN ID between 0 and 4096.",
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

							"disable_dhcp": schema.BoolAttribute{
								Description:         "DisableDHCP disables DHCP for this interface. Implies DisableNetboot.",
								MarkdownDescription: "DisableDHCP disables DHCP for this interface. Implies DisableNetboot.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_netboot": schema.BoolAttribute{
								Description:         "DisableNetboot disables netbooting for this interface. The interface will still receivenetwork information specified by DHCP.",
								MarkdownDescription: "DisableNetboot disables netbooting for this interface. The interface will still receivenetwork information specified by DHCP.",
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
						Description:         "OSIE describes the Operating System Installation Environment to be netbooted.",
						MarkdownDescription: "OSIE describes the Operating System Installation Environment to be netbooted.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage_devices": schema.ListAttribute{
						Description:         "StorageDevices is a list of storage devices that will be available in the OSIE.",
						MarkdownDescription: "StorageDevices is a list of storage devices that will be available in the OSIE.",
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
		},
	}
}

func (r *TinkerbellOrgHardwareV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_tinkerbell_org_hardware_v1alpha2_manifest")

	var model TinkerbellOrgHardwareV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("tinkerbell.org/v1alpha2")
	model.Kind = pointer.String("Hardware")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
