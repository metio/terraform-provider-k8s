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
	_ datasource.DataSource = &SriovnetworkOpenshiftIoOvsnetworkV1Manifest{}
)

func NewSriovnetworkOpenshiftIoOvsnetworkV1Manifest() datasource.DataSource {
	return &SriovnetworkOpenshiftIoOvsnetworkV1Manifest{}
}

type SriovnetworkOpenshiftIoOvsnetworkV1Manifest struct{}

type SriovnetworkOpenshiftIoOvsnetworkV1ManifestData struct {
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
		Bridge           *string `tfsdk:"bridge" json:"bridge,omitempty"`
		Capabilities     *string `tfsdk:"capabilities" json:"capabilities,omitempty"`
		InterfaceType    *string `tfsdk:"interface_type" json:"interfaceType,omitempty"`
		Ipam             *string `tfsdk:"ipam" json:"ipam,omitempty"`
		MetaPlugins      *string `tfsdk:"meta_plugins" json:"metaPlugins,omitempty"`
		Mtu              *int64  `tfsdk:"mtu" json:"mtu,omitempty"`
		NetworkNamespace *string `tfsdk:"network_namespace" json:"networkNamespace,omitempty"`
		ResourceName     *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
		Trunk            *[]struct {
			Id    *int64 `tfsdk:"id" json:"id,omitempty"`
			MaxID *int64 `tfsdk:"max_id" json:"maxID,omitempty"`
			MinID *int64 `tfsdk:"min_id" json:"minID,omitempty"`
		} `tfsdk:"trunk" json:"trunk,omitempty"`
		Vlan *int64 `tfsdk:"vlan" json:"vlan,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SriovnetworkOpenshiftIoOvsnetworkV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sriovnetwork_openshift_io_ovs_network_v1_manifest"
}

func (r *SriovnetworkOpenshiftIoOvsnetworkV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OVSNetwork is the Schema for the ovsnetworks API",
		MarkdownDescription: "OVSNetwork is the Schema for the ovsnetworks API",
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
				Description:         "OVSNetworkSpec defines the desired state of OVSNetwork",
				MarkdownDescription: "OVSNetworkSpec defines the desired state of OVSNetwork",
				Attributes: map[string]schema.Attribute{
					"bridge": schema.StringAttribute{
						Description:         "name of the OVS bridge, if not set OVS will automatically select bridge based on VF PCI address",
						MarkdownDescription: "name of the OVS bridge, if not set OVS will automatically select bridge based on VF PCI address",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"capabilities": schema.StringAttribute{
						Description:         "Capabilities to be configured for this network. Capabilities supported: (mac|ips), e.g. '{'mac': true}'",
						MarkdownDescription: "Capabilities to be configured for this network. Capabilities supported: (mac|ips), e.g. '{'mac': true}'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interface_type": schema.StringAttribute{
						Description:         "The type of interface on ovs.",
						MarkdownDescription: "The type of interface on ovs.",
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

					"meta_plugins": schema.StringAttribute{
						Description:         "MetaPluginsConfig configuration to be used in order to chain metaplugins",
						MarkdownDescription: "MetaPluginsConfig configuration to be used in order to chain metaplugins",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mtu": schema.Int64Attribute{
						Description:         "Mtu for the OVS port",
						MarkdownDescription: "Mtu for the OVS port",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network_namespace": schema.StringAttribute{
						Description:         "Namespace of the NetworkAttachmentDefinition custom resource",
						MarkdownDescription: "Namespace of the NetworkAttachmentDefinition custom resource",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resource_name": schema.StringAttribute{
						Description:         "OVS Network device plugin endpoint resource name",
						MarkdownDescription: "OVS Network device plugin endpoint resource name",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"trunk": schema.ListNestedAttribute{
						Description:         "Trunk configuration for the OVS port",
						MarkdownDescription: "Trunk configuration for the OVS port",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(4095),
									},
								},

								"max_id": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(4095),
									},
								},

								"min_id": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
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

					"vlan": schema.Int64Attribute{
						Description:         "Vlan to assign for the OVS port",
						MarkdownDescription: "Vlan to assign for the OVS port",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
							int64validator.AtMost(4095),
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

func (r *SriovnetworkOpenshiftIoOvsnetworkV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sriovnetwork_openshift_io_ovs_network_v1_manifest")

	var model SriovnetworkOpenshiftIoOvsnetworkV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sriovnetwork.openshift.io/v1")
	model.Kind = pointer.String("OVSNetwork")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
