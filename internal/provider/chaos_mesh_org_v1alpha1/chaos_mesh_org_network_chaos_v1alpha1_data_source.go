/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource              = &ChaosMeshOrgNetworkChaosV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &ChaosMeshOrgNetworkChaosV1Alpha1DataSource{}
)

func NewChaosMeshOrgNetworkChaosV1Alpha1DataSource() datasource.DataSource {
	return &ChaosMeshOrgNetworkChaosV1Alpha1DataSource{}
}

type ChaosMeshOrgNetworkChaosV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ChaosMeshOrgNetworkChaosV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Action    *string `tfsdk:"action" json:"action,omitempty"`
		Bandwidth *struct {
			Buffer   *int64  `tfsdk:"buffer" json:"buffer,omitempty"`
			Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
			Minburst *int64  `tfsdk:"minburst" json:"minburst,omitempty"`
			Peakrate *int64  `tfsdk:"peakrate" json:"peakrate,omitempty"`
			Rate     *string `tfsdk:"rate" json:"rate,omitempty"`
		} `tfsdk:"bandwidth" json:"bandwidth,omitempty"`
		Corrupt *struct {
			Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Corrupt     *string `tfsdk:"corrupt" json:"corrupt,omitempty"`
		} `tfsdk:"corrupt" json:"corrupt,omitempty"`
		Delay *struct {
			Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Jitter      *string `tfsdk:"jitter" json:"jitter,omitempty"`
			Latency     *string `tfsdk:"latency" json:"latency,omitempty"`
			Reorder     *struct {
				Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
				Gap         *int64  `tfsdk:"gap" json:"gap,omitempty"`
				Reorder     *string `tfsdk:"reorder" json:"reorder,omitempty"`
			} `tfsdk:"reorder" json:"reorder,omitempty"`
		} `tfsdk:"delay" json:"delay,omitempty"`
		Device    *string `tfsdk:"device" json:"device,omitempty"`
		Direction *string `tfsdk:"direction" json:"direction,omitempty"`
		Duplicate *struct {
			Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Duplicate   *string `tfsdk:"duplicate" json:"duplicate,omitempty"`
		} `tfsdk:"duplicate" json:"duplicate,omitempty"`
		Duration        *string   `tfsdk:"duration" json:"duration,omitempty"`
		ExternalTargets *[]string `tfsdk:"external_targets" json:"externalTargets,omitempty"`
		Loss            *struct {
			Correlation *string `tfsdk:"correlation" json:"correlation,omitempty"`
			Loss        *string `tfsdk:"loss" json:"loss,omitempty"`
		} `tfsdk:"loss" json:"loss,omitempty"`
		Mode          *string `tfsdk:"mode" json:"mode,omitempty"`
		RemoteCluster *string `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
		Selector      *struct {
			AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
			ExpressionSelectors *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
			FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
			LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
			Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
			NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
			Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
			PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
			Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		Target *struct {
			Mode     *string `tfsdk:"mode" json:"mode,omitempty"`
			Selector *struct {
				AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" json:"annotationSelectors,omitempty"`
				ExpressionSelectors *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"expression_selectors" json:"expressionSelectors,omitempty"`
				FieldSelectors    *map[string]string   `tfsdk:"field_selectors" json:"fieldSelectors,omitempty"`
				LabelSelectors    *map[string]string   `tfsdk:"label_selectors" json:"labelSelectors,omitempty"`
				Namespaces        *[]string            `tfsdk:"namespaces" json:"namespaces,omitempty"`
				NodeSelectors     *map[string]string   `tfsdk:"node_selectors" json:"nodeSelectors,omitempty"`
				Nodes             *[]string            `tfsdk:"nodes" json:"nodes,omitempty"`
				PodPhaseSelectors *[]string            `tfsdk:"pod_phase_selectors" json:"podPhaseSelectors,omitempty"`
				Pods              *map[string][]string `tfsdk:"pods" json:"pods,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"target" json:"target,omitempty"`
		TargetDevice *string `tfsdk:"target_device" json:"targetDevice,omitempty"`
		Value        *string `tfsdk:"value" json:"value,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_network_chaos_v1alpha1"
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NetworkChaos is the Schema for the networkchaos API",
		MarkdownDescription: "NetworkChaos is the Schema for the networkchaos API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "Spec defines the behavior of a pod chaos experiment",
				MarkdownDescription: "Spec defines the behavior of a pod chaos experiment",
				Attributes: map[string]schema.Attribute{
					"action": schema.StringAttribute{
						Description:         "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
						MarkdownDescription: "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"bandwidth": schema.SingleNestedAttribute{
						Description:         "Bandwidth represents the detail about bandwidth control action",
						MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",
						Attributes: map[string]schema.Attribute{
							"buffer": schema.Int64Attribute{
								Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
								MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"limit": schema.Int64Attribute{
								Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
								MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"minburst": schema.Int64Attribute{
								Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
								MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"peakrate": schema.Int64Attribute{
								Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
								MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"rate": schema.StringAttribute{
								Description:         "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
								MarkdownDescription: "Rate is the speed knob. Allows bit, kbit, mbit, gbit, tbit, bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"corrupt": schema.SingleNestedAttribute{
						Description:         "Corrupt represents the detail about corrupt action",
						MarkdownDescription: "Corrupt represents the detail about corrupt action",
						Attributes: map[string]schema.Attribute{
							"correlation": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"corrupt": schema.StringAttribute{
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

					"delay": schema.SingleNestedAttribute{
						Description:         "Delay represents the detail about delay action",
						MarkdownDescription: "Delay represents the detail about delay action",
						Attributes: map[string]schema.Attribute{
							"correlation": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jitter": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"latency": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"reorder": schema.SingleNestedAttribute{
								Description:         "ReorderSpec defines details of packet reorder.",
								MarkdownDescription: "ReorderSpec defines details of packet reorder.",
								Attributes: map[string]schema.Attribute{
									"correlation": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"gap": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"reorder": schema.StringAttribute{
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
						Required: false,
						Optional: false,
						Computed: true,
					},

					"device": schema.StringAttribute{
						Description:         "Device represents the network device to be affected.",
						MarkdownDescription: "Device represents the network device to be affected.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"direction": schema.StringAttribute{
						Description:         "Direction represents the direction, this applies on netem and network partition action",
						MarkdownDescription: "Direction represents the direction, this applies on netem and network partition action",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"duplicate": schema.SingleNestedAttribute{
						Description:         "DuplicateSpec represents the detail about loss action",
						MarkdownDescription: "DuplicateSpec represents the detail about loss action",
						Attributes: map[string]schema.Attribute{
							"correlation": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"duplicate": schema.StringAttribute{
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

					"duration": schema.StringAttribute{
						Description:         "Duration represents the duration of the chaos action",
						MarkdownDescription: "Duration represents the duration of the chaos action",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"external_targets": schema.ListAttribute{
						Description:         "ExternalTargets represents network targets outside k8s",
						MarkdownDescription: "ExternalTargets represents network targets outside k8s",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"loss": schema.SingleNestedAttribute{
						Description:         "Loss represents the detail about loss action",
						MarkdownDescription: "Loss represents the detail about loss action",
						Attributes: map[string]schema.Attribute{
							"correlation": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"loss": schema.StringAttribute{
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

					"mode": schema.StringAttribute{
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"remote_cluster": schema.StringAttribute{
						Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
						MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"selector": schema.SingleNestedAttribute{
						Description:         "Selector is used to select pods that are used to inject chaos action.",
						MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
						Attributes: map[string]schema.Attribute{
							"annotation_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"expression_selectors": schema.ListNestedAttribute{
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
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

							"field_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"label_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespaces": schema.ListAttribute{
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"nodes": schema.ListAttribute{
								Description:         "Nodes is a set of node name and objects must belong to these nodes.",
								MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_phase_selectors": schema.ListAttribute{
								Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pods": schema.MapAttribute{
								Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"target": schema.SingleNestedAttribute{
						Description:         "Target represents network target, this applies on netem and network partition action",
						MarkdownDescription: "Target represents network target, this applies on netem and network partition action",
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
								MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Selector is used to select pods that are used to inject chaos action.",
								MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",
								Attributes: map[string]schema.Attribute{
									"annotation_selectors": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
										MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"expression_selectors": schema.ListNestedAttribute{
										Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
										MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
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

									"field_selectors": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
										MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"label_selectors": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
										MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespaces": schema.ListAttribute{
										Description:         "Namespaces is a set of namespace to which objects belong.",
										MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_selectors": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
										MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"nodes": schema.ListAttribute{
										Description:         "Nodes is a set of node name and objects must belong to these nodes.",
										MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_phase_selectors": schema.ListAttribute{
										Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
										MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pods": schema.MapAttribute{
										Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
										MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"value": schema.StringAttribute{
								Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
								MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"target_device": schema.StringAttribute{
						Description:         "TargetDevice represents the network device to be affected in target scope.",
						MarkdownDescription: "TargetDevice represents the network device to be affected in target scope.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"value": schema.StringAttribute{
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
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

func (r *ChaosMeshOrgNetworkChaosV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_chaos_mesh_org_network_chaos_v1alpha1")

	var data ChaosMeshOrgNetworkChaosV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "chaos-mesh.org", Version: "v1alpha1", Resource: "networkchaos"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse ChaosMeshOrgNetworkChaosV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	data.Kind = pointer.String("NetworkChaos")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
