/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scheduling_koordinator_sh_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &SchedulingKoordinatorShDeviceV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &SchedulingKoordinatorShDeviceV1Alpha1DataSource{}
)

func NewSchedulingKoordinatorShDeviceV1Alpha1DataSource() datasource.DataSource {
	return &SchedulingKoordinatorShDeviceV1Alpha1DataSource{}
}

type SchedulingKoordinatorShDeviceV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SchedulingKoordinatorShDeviceV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Devices *[]struct {
			Health    *bool              `tfsdk:"health" json:"health,omitempty"`
			Id        *string            `tfsdk:"id" json:"id,omitempty"`
			Labels    *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Minor     *int64             `tfsdk:"minor" json:"minor,omitempty"`
			ModuleID  *int64             `tfsdk:"module_id" json:"moduleID,omitempty"`
			Resources *map[string]string `tfsdk:"resources" json:"resources,omitempty"`
			Topology  *struct {
				BusID    *string `tfsdk:"bus_id" json:"busID,omitempty"`
				NodeID   *int64  `tfsdk:"node_id" json:"nodeID,omitempty"`
				PcieID   *int64  `tfsdk:"pcie_id" json:"pcieID,omitempty"`
				SocketID *int64  `tfsdk:"socket_id" json:"socketID,omitempty"`
			} `tfsdk:"topology" json:"topology,omitempty"`
			Type     *string `tfsdk:"type" json:"type,omitempty"`
			VfGroups *[]struct {
				Labels *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Vfs    *[]struct {
					BusID *string `tfsdk:"bus_id" json:"busID,omitempty"`
					Minor *int64  `tfsdk:"minor" json:"minor,omitempty"`
				} `tfsdk:"vfs" json:"vfs,omitempty"`
			} `tfsdk:"vf_groups" json:"vfGroups,omitempty"`
		} `tfsdk:"devices" json:"devices,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SchedulingKoordinatorShDeviceV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scheduling_koordinator_sh_device_v1alpha1"
}

func (r *SchedulingKoordinatorShDeviceV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"devices": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"health": schema.BoolAttribute{
									Description:         "Health indicates whether the device is normal",
									MarkdownDescription: "Health indicates whether the device is normal",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"id": schema.StringAttribute{
									Description:         "UUID represents the UUID of device",
									MarkdownDescription: "UUID represents the UUID of device",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels represents the device properties that can be used to organize and categorize (scope and select) objects",
									MarkdownDescription: "Labels represents the device properties that can be used to organize and categorize (scope and select) objects",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"minor": schema.Int64Attribute{
									Description:         "Minor represents the Minor number of Device, starting from 0",
									MarkdownDescription: "Minor represents the Minor number of Device, starting from 0",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"module_id": schema.Int64Attribute{
									Description:         "ModuleID represents the physical id of Device",
									MarkdownDescription: "ModuleID represents the physical id of Device",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"resources": schema.MapAttribute{
									Description:         "Resources is a set of (resource name, quantity) pairs",
									MarkdownDescription: "Resources is a set of (resource name, quantity) pairs",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"topology": schema.SingleNestedAttribute{
									Description:         "Topology represents the topology information about the device",
									MarkdownDescription: "Topology represents the topology information about the device",
									Attributes: map[string]schema.Attribute{
										"bus_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"node_id": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"pcie_id": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"socket_id": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"type": schema.StringAttribute{
									Description:         "Type represents the type of device",
									MarkdownDescription: "Type represents the type of device",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"vf_groups": schema.ListNestedAttribute{
									Description:         "VFGroups represents the virtual function devices",
									MarkdownDescription: "VFGroups represents the virtual function devices",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"labels": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"vfs": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"bus_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"minor": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *SchedulingKoordinatorShDeviceV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *SchedulingKoordinatorShDeviceV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_scheduling_koordinator_sh_device_v1alpha1")

	var data SchedulingKoordinatorShDeviceV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "scheduling.koordinator.sh", Version: "v1alpha1", Resource: "Device"}).
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

	var readResponse SchedulingKoordinatorShDeviceV1Alpha1DataSourceData
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

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("scheduling.koordinator.sh/v1alpha1")
	data.Kind = pointer.String("Device")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
