/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scheduling_volcano_sh_v1beta1

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
	_ datasource.DataSource              = &SchedulingVolcanoShQueueV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &SchedulingVolcanoShQueueV1Beta1DataSource{}
)

func NewSchedulingVolcanoShQueueV1Beta1DataSource() datasource.DataSource {
	return &SchedulingVolcanoShQueueV1Beta1DataSource{}
}

type SchedulingVolcanoShQueueV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SchedulingVolcanoShQueueV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Affinity *struct {
			NodeGroupAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]string `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution  *[]string `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_group_affinity" json:"nodeGroupAffinity,omitempty"`
			NodeGroupAntiAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]string `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution  *[]string `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_group_anti_affinity" json:"nodeGroupAntiAffinity,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		Capability     *map[string]string `tfsdk:"capability" json:"capability,omitempty"`
		ExtendClusters *[]struct {
			Capacity *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
			Name     *string            `tfsdk:"name" json:"name,omitempty"`
			Weight   *int64             `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"extend_clusters" json:"extendClusters,omitempty"`
		Guarantee *struct {
			Resource *map[string]string `tfsdk:"resource" json:"resource,omitempty"`
		} `tfsdk:"guarantee" json:"guarantee,omitempty"`
		Reclaimable *bool   `tfsdk:"reclaimable" json:"reclaimable,omitempty"`
		Type        *string `tfsdk:"type" json:"type,omitempty"`
		Weight      *int64  `tfsdk:"weight" json:"weight,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SchedulingVolcanoShQueueV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scheduling_volcano_sh_queue_v1beta1"
}

func (r *SchedulingVolcanoShQueueV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Queue is a queue of PodGroup.",
		MarkdownDescription: "Queue is a queue of PodGroup.",
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
				Description:         "Specification of the desired behavior of the queue. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired behavior of the queue. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "If specified, the pod owned by the queue will be scheduled with constraint",
						MarkdownDescription: "If specified, the pod owned by the queue will be scheduled with constraint",
						Attributes: map[string]schema.Attribute{
							"node_group_affinity": schema.SingleNestedAttribute{
								Description:         "Describes nodegroup affinity scheduling rules for the queue(e.g. putting pods of the queue in the nodes of the nodegroup)",
								MarkdownDescription: "Describes nodegroup affinity scheduling rules for the queue(e.g. putting pods of the queue in the nodes of the nodegroup)",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"node_group_anti_affinity": schema.SingleNestedAttribute{
								Description:         "Describes nodegroup anti-affinity scheduling rules for the queue(e.g. avoid putting pods of the queue in the nodes of the nodegroup).",
								MarkdownDescription: "Describes nodegroup anti-affinity scheduling rules for the queue(e.g. avoid putting pods of the queue in the nodes of the nodegroup).",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
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

					"capability": schema.MapAttribute{
						Description:         "ResourceList is a set of (resource name, quantity) pairs.",
						MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"extend_clusters": schema.ListNestedAttribute{
						Description:         "extendCluster indicate the jobs in this Queue will be dispatched to these clusters.",
						MarkdownDescription: "extendCluster indicate the jobs in this Queue will be dispatched to these clusters.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"capacity": schema.MapAttribute{
									Description:         "ResourceList is a set of (resource name, quantity) pairs.",
									MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"weight": schema.Int64Attribute{
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

					"guarantee": schema.SingleNestedAttribute{
						Description:         "Guarantee indicate configuration about resource reservation",
						MarkdownDescription: "Guarantee indicate configuration about resource reservation",
						Attributes: map[string]schema.Attribute{
							"resource": schema.MapAttribute{
								Description:         "The amount of cluster resource reserved for queue. Just set either 'percentage' or 'resource'",
								MarkdownDescription: "The amount of cluster resource reserved for queue. Just set either 'percentage' or 'resource'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"reclaimable": schema.BoolAttribute{
						Description:         "Reclaimable indicate whether the queue can be reclaimed by other queue",
						MarkdownDescription: "Reclaimable indicate whether the queue can be reclaimed by other queue",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"type": schema.StringAttribute{
						Description:         "Type define the type of queue",
						MarkdownDescription: "Type define the type of queue",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"weight": schema.Int64Attribute{
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
		},
	}
}

func (r *SchedulingVolcanoShQueueV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *SchedulingVolcanoShQueueV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_scheduling_volcano_sh_queue_v1beta1")

	var data SchedulingVolcanoShQueueV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "scheduling.volcano.sh", Version: "v1beta1", Resource: "Queue"}).
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

	var readResponse SchedulingVolcanoShQueueV1Beta1DataSourceData
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
	data.ApiVersion = pointer.String("scheduling.volcano.sh/v1beta1")
	data.Kind = pointer.String("Queue")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
