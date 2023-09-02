/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package topology_node_k8s_io_v1alpha1

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
	_ datasource.DataSource = &TopologyNodeK8SIoNodeResourceTopologyV1Alpha1Manifest{}
)

func NewTopologyNodeK8SIoNodeResourceTopologyV1Alpha1Manifest() datasource.DataSource {
	return &TopologyNodeK8SIoNodeResourceTopologyV1Alpha1Manifest{}
}

type TopologyNodeK8SIoNodeResourceTopologyV1Alpha1Manifest struct{}

type TopologyNodeK8SIoNodeResourceTopologyV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	TopologyPolicies *[]string `tfsdk:"topology_policies" json:"topologyPolicies,omitempty"`
	Zones            *[]struct {
		Attributes *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"attributes" json:"attributes,omitempty"`
		Costs *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *int64  `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"costs" json:"costs,omitempty"`
		Name      *string `tfsdk:"name" json:"name,omitempty"`
		Parent    *string `tfsdk:"parent" json:"parent,omitempty"`
		Resources *[]struct {
			Allocatable *string `tfsdk:"allocatable" json:"allocatable,omitempty"`
			Available   *string `tfsdk:"available" json:"available,omitempty"`
			Capacity    *string `tfsdk:"capacity" json:"capacity,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Type *string `tfsdk:"type" json:"type,omitempty"`
	} `tfsdk:"zones" json:"zones,omitempty"`
}

func (r *TopologyNodeK8SIoNodeResourceTopologyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_topology_node_k8s_io_node_resource_topology_v1alpha1_manifest"
}

func (r *TopologyNodeK8SIoNodeResourceTopologyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeResourceTopology describes node resources and their topology.",
		MarkdownDescription: "NodeResourceTopology describes node resources and their topology.",
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

			"topology_policies": schema.ListAttribute{
				Description:         "",
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Required:            true,
				Optional:            false,
				Computed:            false,
			},

			"zones": schema.ListNestedAttribute{
				Description:         "ZoneList contains an array of Zone objects.",
				MarkdownDescription: "ZoneList contains an array of Zone objects.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"attributes": schema.ListNestedAttribute{
							Description:         "AttributeList contains an array of AttributeInfo objects.",
							MarkdownDescription: "AttributeList contains an array of AttributeInfo objects.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"value": schema.StringAttribute{
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

						"costs": schema.ListNestedAttribute{
							Description:         "CostList contains an array of CostInfo objects.",
							MarkdownDescription: "CostList contains an array of CostInfo objects.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"value": schema.Int64Attribute{
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

						"name": schema.StringAttribute{
							Description:         "",
							MarkdownDescription: "",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},

						"parent": schema.StringAttribute{
							Description:         "",
							MarkdownDescription: "",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"resources": schema.ListNestedAttribute{
							Description:         "ResourceInfoList contains an array of ResourceInfo objects.",
							MarkdownDescription: "ResourceInfoList contains an array of ResourceInfo objects.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"allocatable": schema.StringAttribute{
										Description:         "Allocatable quantity of the resource, corresponding to allocatable in node status, i.e. total amount of this resource available to be used by pods.",
										MarkdownDescription: "Allocatable quantity of the resource, corresponding to allocatable in node status, i.e. total amount of this resource available to be used by pods.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"available": schema.StringAttribute{
										Description:         "Available is the amount of this resource currently available for new (to be scheduled) pods, i.e. Allocatable minus the resources reserved by currently running pods.",
										MarkdownDescription: "Available is the amount of this resource currently available for new (to be scheduled) pods, i.e. Allocatable minus the resources reserved by currently running pods.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"capacity": schema.StringAttribute{
										Description:         "Capacity of the resource, corresponding to capacity in node status, i.e. total amount of this resource that the node has.",
										MarkdownDescription: "Capacity of the resource, corresponding to capacity in node status, i.e. total amount of this resource that the node has.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the resource.",
										MarkdownDescription: "Name of the resource.",
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

						"type": schema.StringAttribute{
							Description:         "",
							MarkdownDescription: "",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *TopologyNodeK8SIoNodeResourceTopologyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_topology_node_k8s_io_node_resource_topology_v1alpha1_manifest")

	var model TopologyNodeK8SIoNodeResourceTopologyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("topology.node.k8s.io/v1alpha1")
	model.Kind = pointer.String("NodeResourceTopology")

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
