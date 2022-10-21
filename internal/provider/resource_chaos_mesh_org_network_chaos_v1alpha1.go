/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type ChaosMeshOrgNetworkChaosV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ChaosMeshOrgNetworkChaosV1Alpha1Resource)(nil)
)

type ChaosMeshOrgNetworkChaosV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ChaosMeshOrgNetworkChaosV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Action *string `tfsdk:"action" yaml:"action,omitempty"`

		Bandwidth *struct {
			Buffer *int64 `tfsdk:"buffer" yaml:"buffer,omitempty"`

			Limit *int64 `tfsdk:"limit" yaml:"limit,omitempty"`

			Minburst *int64 `tfsdk:"minburst" yaml:"minburst,omitempty"`

			Peakrate *int64 `tfsdk:"peakrate" yaml:"peakrate,omitempty"`

			Rate *string `tfsdk:"rate" yaml:"rate,omitempty"`
		} `tfsdk:"bandwidth" yaml:"bandwidth,omitempty"`

		Corrupt *struct {
			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Corrupt *string `tfsdk:"corrupt" yaml:"corrupt,omitempty"`
		} `tfsdk:"corrupt" yaml:"corrupt,omitempty"`

		Delay *struct {
			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Jitter *string `tfsdk:"jitter" yaml:"jitter,omitempty"`

			Latency *string `tfsdk:"latency" yaml:"latency,omitempty"`

			Reorder *struct {
				Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

				Gap *int64 `tfsdk:"gap" yaml:"gap,omitempty"`

				Reorder *string `tfsdk:"reorder" yaml:"reorder,omitempty"`
			} `tfsdk:"reorder" yaml:"reorder,omitempty"`
		} `tfsdk:"delay" yaml:"delay,omitempty"`

		Device *string `tfsdk:"device" yaml:"device,omitempty"`

		Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

		Duplicate *struct {
			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Duplicate *string `tfsdk:"duplicate" yaml:"duplicate,omitempty"`
		} `tfsdk:"duplicate" yaml:"duplicate,omitempty"`

		Duration *string `tfsdk:"duration" yaml:"duration,omitempty"`

		ExternalTargets *[]string `tfsdk:"external_targets" yaml:"externalTargets,omitempty"`

		Loss *struct {
			Correlation *string `tfsdk:"correlation" yaml:"correlation,omitempty"`

			Loss *string `tfsdk:"loss" yaml:"loss,omitempty"`
		} `tfsdk:"loss" yaml:"loss,omitempty"`

		Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

		Selector *struct {
			AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

			ExpressionSelectors *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

			FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

			LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

			Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

			NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

			Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

			PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

			Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`

		Target *struct {
			Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

			Selector *struct {
				AnnotationSelectors *map[string]string `tfsdk:"annotation_selectors" yaml:"annotationSelectors,omitempty"`

				ExpressionSelectors *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"expression_selectors" yaml:"expressionSelectors,omitempty"`

				FieldSelectors *map[string]string `tfsdk:"field_selectors" yaml:"fieldSelectors,omitempty"`

				LabelSelectors *map[string]string `tfsdk:"label_selectors" yaml:"labelSelectors,omitempty"`

				Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

				NodeSelectors *map[string]string `tfsdk:"node_selectors" yaml:"nodeSelectors,omitempty"`

				Nodes *[]string `tfsdk:"nodes" yaml:"nodes,omitempty"`

				PodPhaseSelectors *[]string `tfsdk:"pod_phase_selectors" yaml:"podPhaseSelectors,omitempty"`

				Pods *map[string][]string `tfsdk:"pods" yaml:"pods,omitempty"`
			} `tfsdk:"selector" yaml:"selector,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"target" yaml:"target,omitempty"`

		TargetDevice *string `tfsdk:"target_device" yaml:"targetDevice,omitempty"`

		Value *string `tfsdk:"value" yaml:"value,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewChaosMeshOrgNetworkChaosV1Alpha1Resource() resource.Resource {
	return &ChaosMeshOrgNetworkChaosV1Alpha1Resource{}
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chaos_mesh_org_network_chaos_v1alpha1"
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "NetworkChaos is the Schema for the networkchaos API",
		MarkdownDescription: "NetworkChaos is the Schema for the networkchaos API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Spec defines the behavior of a pod chaos experiment",
				MarkdownDescription: "Spec defines the behavior of a pod chaos experiment",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"action": {
						Description:         "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",
						MarkdownDescription: "Action defines the specific network chaos action. Supported action: partition, netem, delay, loss, duplicate, corrupt Default action: delay",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("netem", "delay", "loss", "duplicate", "corrupt", "partition", "bandwidth"),
						},
					},

					"bandwidth": {
						Description:         "Bandwidth represents the detail about bandwidth control action",
						MarkdownDescription: "Bandwidth represents the detail about bandwidth control action",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"buffer": {
								Description:         "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",
								MarkdownDescription: "Buffer is the maximum amount of bytes that tokens can be available for instantaneously.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"limit": {
								Description:         "Limit is the number of bytes that can be queued waiting for tokens to become available.",
								MarkdownDescription: "Limit is the number of bytes that can be queued waiting for tokens to become available.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"minburst": {
								Description:         "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",
								MarkdownDescription: "Minburst specifies the size of the peakrate bucket. For perfect accuracy, should be set to the MTU of the interface.  If a peakrate is needed, but some burstiness is acceptable, this size can be raised. A 3000 byte minburst allows around 3mbit/s of peakrate, given 1000 byte packets.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"peakrate": {
								Description:         "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",
								MarkdownDescription: "Peakrate is the maximum depletion rate of the bucket. The peakrate does not need to be set, it is only necessary if perfect millisecond timescale shaping is required.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"rate": {
								Description:         "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",
								MarkdownDescription: "Rate is the speed knob. Allows bps, kbps, mbps, gbps, tbps unit. bps means bytes per second.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"corrupt": {
						Description:         "Corrupt represents the detail about corrupt action",
						MarkdownDescription: "Corrupt represents the detail about corrupt action",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"correlation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"corrupt": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"delay": {
						Description:         "Delay represents the detail about delay action",
						MarkdownDescription: "Delay represents the detail about delay action",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"correlation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jitter": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"reorder": {
								Description:         "ReorderSpec defines details of packet reorder.",
								MarkdownDescription: "ReorderSpec defines details of packet reorder.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"gap": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"reorder": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"device": {
						Description:         "Device represents the network device to be affected.",
						MarkdownDescription: "Device represents the network device to be affected.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"direction": {
						Description:         "Direction represents the direction, this applies on netem and network partition action",
						MarkdownDescription: "Direction represents the direction, this applies on netem and network partition action",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("to", "from", "both"),
						},
					},

					"duplicate": {
						Description:         "DuplicateSpec represents the detail about loss action",
						MarkdownDescription: "DuplicateSpec represents the detail about loss action",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"correlation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplicate": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"duration": {
						Description:         "Duration represents the duration of the chaos action",
						MarkdownDescription: "Duration represents the duration of the chaos action",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_targets": {
						Description:         "ExternalTargets represents network targets outside k8s",
						MarkdownDescription: "ExternalTargets represents network targets outside k8s",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"loss": {
						Description:         "Loss represents the detail about loss action",
						MarkdownDescription: "Loss represents the detail about loss action",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"correlation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mode": {
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
						},
					},

					"selector": {
						Description:         "Selector is used to select pods that are used to inject chaos action.",
						MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotation_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"expression_selectors": {
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"field_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"label_selectors": {
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespaces": {
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_selectors": {
								Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"nodes": {
								Description:         "Nodes is a set of node name and objects must belong to these nodes.",
								MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_phase_selectors": {
								Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pods": {
								Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

								Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"target": {
						Description:         "Target represents network target, this applies on netem and network partition action",
						MarkdownDescription: "Target represents network target, this applies on netem and network partition action",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"mode": {
								Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
								MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
								},
							},

							"selector": {
								Description:         "Selector is used to select pods that are used to inject chaos action.",
								MarkdownDescription: "Selector is used to select pods that are used to inject chaos action.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotation_selectors": {
										Description:         "Map of string keys and values that can be used to select objects. A selector based on annotations.",
										MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on annotations.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expression_selectors": {
										Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
										MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"field_selectors": {
										Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
										MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"label_selectors": {
										Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
										MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespaces": {
										Description:         "Namespaces is a set of namespace to which objects belong.",
										MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_selectors": {
										Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
										MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"nodes": {
										Description:         "Nodes is a set of node name and objects must belong to these nodes.",
										MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_phase_selectors": {
										Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
										MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pods": {
										Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
										MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",

										Type: types.MapType{ElemType: types.ListType{ElemType: types.StringType}},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
								MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_device": {
						Description:         "TargetDevice represents the network device to be affected in target scope.",
						MarkdownDescription: "TargetDevice represents the network device to be affected in target scope.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"value": {
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_chaos_mesh_org_network_chaos_v1alpha1")

	var state ChaosMeshOrgNetworkChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgNetworkChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("NetworkChaos")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_network_chaos_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_chaos_mesh_org_network_chaos_v1alpha1")

	var state ChaosMeshOrgNetworkChaosV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ChaosMeshOrgNetworkChaosV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("chaos-mesh.org/v1alpha1")
	goModel.Kind = utilities.Ptr("NetworkChaos")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ChaosMeshOrgNetworkChaosV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_chaos_mesh_org_network_chaos_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
