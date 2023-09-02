/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scheduling_koordinator_sh_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SchedulingKoordinatorShDeviceV1Alpha1Manifest{}
)

func NewSchedulingKoordinatorShDeviceV1Alpha1Manifest() datasource.DataSource {
	return &SchedulingKoordinatorShDeviceV1Alpha1Manifest{}
}

type SchedulingKoordinatorShDeviceV1Alpha1Manifest struct{}

type SchedulingKoordinatorShDeviceV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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

func (r *SchedulingKoordinatorShDeviceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scheduling_koordinator_sh_device_v1alpha1_manifest"
}

func (r *SchedulingKoordinatorShDeviceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
									Optional:            true,
									Computed:            false,
								},

								"id": schema.StringAttribute{
									Description:         "UUID represents the UUID of device",
									MarkdownDescription: "UUID represents the UUID of device",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels represents the device properties that can be used to organize and categorize (scope and select) objects",
									MarkdownDescription: "Labels represents the device properties that can be used to organize and categorize (scope and select) objects",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"minor": schema.Int64Attribute{
									Description:         "Minor represents the Minor number of Device, starting from 0",
									MarkdownDescription: "Minor represents the Minor number of Device, starting from 0",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"module_id": schema.Int64Attribute{
									Description:         "ModuleID represents the physical id of Device",
									MarkdownDescription: "ModuleID represents the physical id of Device",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resources": schema.MapAttribute{
									Description:         "Resources is a set of (resource name, quantity) pairs",
									MarkdownDescription: "Resources is a set of (resource name, quantity) pairs",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"topology": schema.SingleNestedAttribute{
									Description:         "Topology represents the topology information about the device",
									MarkdownDescription: "Topology represents the topology information about the device",
									Attributes: map[string]schema.Attribute{
										"bus_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_id": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"pcie_id": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"socket_id": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"type": schema.StringAttribute{
									Description:         "Type represents the type of device",
									MarkdownDescription: "Type represents the type of device",
									Required:            false,
									Optional:            true,
									Computed:            false,
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
												Optional:            true,
												Computed:            false,
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
															Optional:            true,
															Computed:            false,
														},

														"minor": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
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

func (r *SchedulingKoordinatorShDeviceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_scheduling_koordinator_sh_device_v1alpha1_manifest")

	var model SchedulingKoordinatorShDeviceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("scheduling.koordinator.sh/v1alpha1")
	model.Kind = pointer.String("Device")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
