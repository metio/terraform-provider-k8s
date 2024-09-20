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
	_ datasource.DataSource = &SriovnetworkOpenshiftIoSriovNetworkV1Manifest{}
)

func NewSriovnetworkOpenshiftIoSriovNetworkV1Manifest() datasource.DataSource {
	return &SriovnetworkOpenshiftIoSriovNetworkV1Manifest{}
}

type SriovnetworkOpenshiftIoSriovNetworkV1Manifest struct{}

type SriovnetworkOpenshiftIoSriovNetworkV1ManifestData struct {
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
		Capabilities     *string `tfsdk:"capabilities" json:"capabilities,omitempty"`
		Ipam             *string `tfsdk:"ipam" json:"ipam,omitempty"`
		LinkState        *string `tfsdk:"link_state" json:"linkState,omitempty"`
		LogFile          *string `tfsdk:"log_file" json:"logFile,omitempty"`
		LogLevel         *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		MaxTxRate        *int64  `tfsdk:"max_tx_rate" json:"maxTxRate,omitempty"`
		MetaPlugins      *string `tfsdk:"meta_plugins" json:"metaPlugins,omitempty"`
		MinTxRate        *int64  `tfsdk:"min_tx_rate" json:"minTxRate,omitempty"`
		NetworkNamespace *string `tfsdk:"network_namespace" json:"networkNamespace,omitempty"`
		ResourceName     *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
		SpoofChk         *string `tfsdk:"spoof_chk" json:"spoofChk,omitempty"`
		Trust            *string `tfsdk:"trust" json:"trust,omitempty"`
		Vlan             *int64  `tfsdk:"vlan" json:"vlan,omitempty"`
		VlanProto        *string `tfsdk:"vlan_proto" json:"vlanProto,omitempty"`
		VlanQoS          *int64  `tfsdk:"vlan_qo_s" json:"vlanQoS,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SriovnetworkOpenshiftIoSriovNetworkV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sriovnetwork_openshift_io_sriov_network_v1_manifest"
}

func (r *SriovnetworkOpenshiftIoSriovNetworkV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SriovNetwork is the Schema for the sriovnetworks API",
		MarkdownDescription: "SriovNetwork is the Schema for the sriovnetworks API",
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
				Description:         "SriovNetworkSpec defines the desired state of SriovNetwork",
				MarkdownDescription: "SriovNetworkSpec defines the desired state of SriovNetwork",
				Attributes: map[string]schema.Attribute{
					"capabilities": schema.StringAttribute{
						Description:         "Capabilities to be configured for this network. Capabilities supported: (mac|ips), e.g. '{'mac': true}'",
						MarkdownDescription: "Capabilities to be configured for this network. Capabilities supported: (mac|ips), e.g. '{'mac': true}'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipam": schema.StringAttribute{
						Description:         "IPAM configuration to be used for this network.",
						MarkdownDescription: "IPAM configuration to be used for this network.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"link_state": schema.StringAttribute{
						Description:         "VF link state (enable|disable|auto)",
						MarkdownDescription: "VF link state (enable|disable|auto)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("auto", "enable", "disable"),
						},
					},

					"log_file": schema.StringAttribute{
						Description:         "LogFile sets the log file of the SRIOV CNI plugin logs. If unset (default), this will log to stderr and thus to multus and container runtime logs.",
						MarkdownDescription: "LogFile sets the log file of the SRIOV CNI plugin logs. If unset (default), this will log to stderr and thus to multus and container runtime logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.StringAttribute{
						Description:         "LogLevel sets the log level of the SRIOV CNI plugin - either of panic, error, warning, info, debug. Defaults to info if left blank.",
						MarkdownDescription: "LogLevel sets the log level of the SRIOV CNI plugin - either of panic, error, warning, info, debug. Defaults to info if left blank.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("panic", "error", "warning", "info", "debug", ""),
						},
					},

					"max_tx_rate": schema.Int64Attribute{
						Description:         "Maximum tx rate, in Mbps, for the VF. Defaults to 0 (no rate limiting)",
						MarkdownDescription: "Maximum tx rate, in Mbps, for the VF. Defaults to 0 (no rate limiting)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"meta_plugins": schema.StringAttribute{
						Description:         "MetaPluginsConfig configuration to be used in order to chain metaplugins to the sriov interface returned by the operator.",
						MarkdownDescription: "MetaPluginsConfig configuration to be used in order to chain metaplugins to the sriov interface returned by the operator.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"min_tx_rate": schema.Int64Attribute{
						Description:         "Minimum tx rate, in Mbps, for the VF. Defaults to 0 (no rate limiting). min_tx_rate should be <= max_tx_rate.",
						MarkdownDescription: "Minimum tx rate, in Mbps, for the VF. Defaults to 0 (no rate limiting). min_tx_rate should be <= max_tx_rate.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"network_namespace": schema.StringAttribute{
						Description:         "Namespace of the NetworkAttachmentDefinition custom resource",
						MarkdownDescription: "Namespace of the NetworkAttachmentDefinition custom resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_name": schema.StringAttribute{
						Description:         "SRIOV Network device plugin endpoint resource name",
						MarkdownDescription: "SRIOV Network device plugin endpoint resource name",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"spoof_chk": schema.StringAttribute{
						Description:         "VF spoof check, (on|off)",
						MarkdownDescription: "VF spoof check, (on|off)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("on", "off"),
						},
					},

					"trust": schema.StringAttribute{
						Description:         "VF trust mode (on|off)",
						MarkdownDescription: "VF trust mode (on|off)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("on", "off"),
						},
					},

					"vlan": schema.Int64Attribute{
						Description:         "VLAN ID to assign for the VF. Defaults to 0.",
						MarkdownDescription: "VLAN ID to assign for the VF. Defaults to 0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4096),
						},
					},

					"vlan_proto": schema.StringAttribute{
						Description:         "VLAN proto to assign for the VF. Defaults to 802.1q.",
						MarkdownDescription: "VLAN proto to assign for the VF. Defaults to 802.1q.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("802.1q", "802.1Q", "802.1ad", "802.1AD"),
						},
					},

					"vlan_qo_s": schema.Int64Attribute{
						Description:         "VLAN QoS ID to assign for the VF. Defaults to 0.",
						MarkdownDescription: "VLAN QoS ID to assign for the VF. Defaults to 0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(7),
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

func (r *SriovnetworkOpenshiftIoSriovNetworkV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sriovnetwork_openshift_io_sriov_network_v1_manifest")

	var model SriovnetworkOpenshiftIoSriovNetworkV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sriovnetwork.openshift.io/v1")
	model.Kind = pointer.String("SriovNetwork")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
