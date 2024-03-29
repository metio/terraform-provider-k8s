/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package network_openshift_io_v1

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
	_ datasource.DataSource = &NetworkOpenshiftIoClusterNetworkV1Manifest{}
)

func NewNetworkOpenshiftIoClusterNetworkV1Manifest() datasource.DataSource {
	return &NetworkOpenshiftIoClusterNetworkV1Manifest{}
}

type NetworkOpenshiftIoClusterNetworkV1Manifest struct{}

type NetworkOpenshiftIoClusterNetworkV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	ClusterNetworks *[]struct {
		CIDR             *string `tfsdk:"cidr" json:"CIDR,omitempty"`
		HostSubnetLength *int64  `tfsdk:"host_subnet_length" json:"hostSubnetLength,omitempty"`
	} `tfsdk:"cluster_networks" json:"clusterNetworks,omitempty"`
	Hostsubnetlength *int64  `tfsdk:"hostsubnetlength" json:"hostsubnetlength,omitempty"`
	Mtu              *int64  `tfsdk:"mtu" json:"mtu,omitempty"`
	Network          *string `tfsdk:"network" json:"network,omitempty"`
	PluginName       *string `tfsdk:"plugin_name" json:"pluginName,omitempty"`
	ServiceNetwork   *string `tfsdk:"service_network" json:"serviceNetwork,omitempty"`
	VxlanPort        *int64  `tfsdk:"vxlan_port" json:"vxlanPort,omitempty"`
}

func (r *NetworkOpenshiftIoClusterNetworkV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_network_openshift_io_cluster_network_v1_manifest"
}

func (r *NetworkOpenshiftIoClusterNetworkV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterNetwork describes the cluster network. There is normally only one object of this type, named 'default', which is created by the SDN network plugin based on the master configuration when the cluster is brought up for the first time.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ClusterNetwork describes the cluster network. There is normally only one object of this type, named 'default', which is created by the SDN network plugin based on the master configuration when the cluster is brought up for the first time.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"cluster_networks": schema.ListNestedAttribute{
				Description:         "ClusterNetworks is a list of ClusterNetwork objects that defines the global overlay network's L3 space by specifying a set of CIDR and netmasks that the SDN can allocate addresses from.",
				MarkdownDescription: "ClusterNetworks is a list of ClusterNetwork objects that defines the global overlay network's L3 space by specifying a set of CIDR and netmasks that the SDN can allocate addresses from.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cidr": schema.StringAttribute{
							Description:         "CIDR defines the total range of a cluster networks address space.",
							MarkdownDescription: "CIDR defines the total range of a cluster networks address space.",
							Required:            true,
							Optional:            false,
							Computed:            false,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/([0-9]|[12][0-9]|3[0-2])$`), ""),
							},
						},

						"host_subnet_length": schema.Int64Attribute{
							Description:         "HostSubnetLength is the number of bits of the accompanying CIDR address to allocate to each node. eg, 8 would mean that each node would have a /24 slice of the overlay network for its pods.",
							MarkdownDescription: "HostSubnetLength is the number of bits of the accompanying CIDR address to allocate to each node. eg, 8 would mean that each node would have a /24 slice of the overlay network for its pods.",
							Required:            true,
							Optional:            false,
							Computed:            false,
							Validators: []validator.Int64{
								int64validator.AtLeast(2),
								int64validator.AtMost(30),
							},
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},

			"hostsubnetlength": schema.Int64Attribute{
				Description:         "HostSubnetLength is the number of bits of network to allocate to each node. eg, 8 would mean that each node would have a /24 slice of the overlay network for its pods",
				MarkdownDescription: "HostSubnetLength is the number of bits of network to allocate to each node. eg, 8 would mean that each node would have a /24 slice of the overlay network for its pods",
				Required:            false,
				Optional:            true,
				Computed:            false,
				Validators: []validator.Int64{
					int64validator.AtLeast(2),
					int64validator.AtMost(30),
				},
			},

			"mtu": schema.Int64Attribute{
				Description:         "MTU is the MTU for the overlay network. This should be 50 less than the MTU of the network connecting the nodes. It is normally autodetected by the cluster network operator.",
				MarkdownDescription: "MTU is the MTU for the overlay network. This should be 50 less than the MTU of the network connecting the nodes. It is normally autodetected by the cluster network operator.",
				Required:            false,
				Optional:            true,
				Computed:            false,
				Validators: []validator.Int64{
					int64validator.AtLeast(576),
					int64validator.AtMost(65536),
				},
			},

			"network": schema.StringAttribute{
				Description:         "Network is a CIDR string specifying the global overlay network's L3 space",
				MarkdownDescription: "Network is a CIDR string specifying the global overlay network's L3 space",
				Required:            false,
				Optional:            true,
				Computed:            false,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/([0-9]|[12][0-9]|3[0-2])$`), ""),
				},
			},

			"plugin_name": schema.StringAttribute{
				Description:         "PluginName is the name of the network plugin being used",
				MarkdownDescription: "PluginName is the name of the network plugin being used",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"service_network": schema.StringAttribute{
				Description:         "ServiceNetwork is the CIDR range that Service IP addresses are allocated from",
				MarkdownDescription: "ServiceNetwork is the CIDR range that Service IP addresses are allocated from",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[0-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/([0-9]|[12][0-9]|3[0-2])$`), ""),
				},
			},

			"vxlan_port": schema.Int64Attribute{
				Description:         "VXLANPort sets the VXLAN destination port used by the cluster. It is set by the master configuration file on startup and cannot be edited manually. Valid values for VXLANPort are integers 1-65535 inclusive and if unset defaults to 4789. Changing VXLANPort allows users to resolve issues between openshift SDN and other software trying to use the same VXLAN destination port.",
				MarkdownDescription: "VXLANPort sets the VXLAN destination port used by the cluster. It is set by the master configuration file on startup and cannot be edited manually. Valid values for VXLANPort are integers 1-65535 inclusive and if unset defaults to 4789. Changing VXLANPort allows users to resolve issues between openshift SDN and other software trying to use the same VXLAN destination port.",
				Required:            false,
				Optional:            true,
				Computed:            false,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
					int64validator.AtMost(65535),
				},
			},
		},
	}
}

func (r *NetworkOpenshiftIoClusterNetworkV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_network_openshift_io_cluster_network_v1_manifest")

	var model NetworkOpenshiftIoClusterNetworkV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("network.openshift.io/v1")
	model.Kind = pointer.String("ClusterNetwork")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
