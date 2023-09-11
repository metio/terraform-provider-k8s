/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package topology_node_k8s_io_v1alpha1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource{}
)

func NewTopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource() datasource.DataSource {
	return &TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource{}
}

type TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_topology_node_k8s_io_node_resource_topology_v1alpha1"
}

func (r *TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"topology_policies": schema.ListAttribute{
				Description:         "",
				MarkdownDescription: "",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            false,
				Computed:            true,
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
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"value": schema.StringAttribute{
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

						"costs": schema.ListNestedAttribute{
							Description:         "CostList contains an array of CostInfo objects.",
							MarkdownDescription: "CostList contains an array of CostInfo objects.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"value": schema.Int64Attribute{
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

						"name": schema.StringAttribute{
							Description:         "",
							MarkdownDescription: "",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"parent": schema.StringAttribute{
							Description:         "",
							MarkdownDescription: "",
							Required:            false,
							Optional:            false,
							Computed:            true,
						},

						"resources": schema.ListNestedAttribute{
							Description:         "ResourceInfoList contains an array of ResourceInfo objects.",
							MarkdownDescription: "ResourceInfoList contains an array of ResourceInfo objects.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"allocatable": schema.StringAttribute{
										Description:         "Allocatable quantity of the resource, corresponding to allocatable in node status, i.e. total amount of this resource available to be used by pods.",
										MarkdownDescription: "Allocatable quantity of the resource, corresponding to allocatable in node status, i.e. total amount of this resource available to be used by pods.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"available": schema.StringAttribute{
										Description:         "Available is the amount of this resource currently available for new (to be scheduled) pods, i.e. Allocatable minus the resources reserved by currently running pods.",
										MarkdownDescription: "Available is the amount of this resource currently available for new (to be scheduled) pods, i.e. Allocatable minus the resources reserved by currently running pods.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"capacity": schema.StringAttribute{
										Description:         "Capacity of the resource, corresponding to capacity in node status, i.e. total amount of this resource that the node has.",
										MarkdownDescription: "Capacity of the resource, corresponding to capacity in node status, i.e. total amount of this resource that the node has.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the resource.",
										MarkdownDescription: "Name of the resource.",
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

						"type": schema.StringAttribute{
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
	}
}

func (r *TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_topology_node_k8s_io_node_resource_topology_v1alpha1")

	var data TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "topology.node.k8s.io", Version: "v1alpha1", Resource: "noderesourcetopologies"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name configured.\n\n"+
						"Name: %s", data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse TopologyNodeK8SIoNodeResourceTopologyV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("topology.node.k8s.io/v1alpha1")
	data.Kind = pointer.String("NodeResourceTopology")
	data.Metadata = readResponse.Metadata
	data.TopologyPolicies = readResponse.TopologyPolicies
	data.Zones = readResponse.Zones

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
