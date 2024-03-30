/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package chaos_mesh_org_v1alpha1

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
	_ datasource.DataSource = &ChaosMeshOrgStressChaosV1Alpha1Manifest{}
)

func NewChaosMeshOrgStressChaosV1Alpha1Manifest() datasource.DataSource {
	return &ChaosMeshOrgStressChaosV1Alpha1Manifest{}
}

type ChaosMeshOrgStressChaosV1Alpha1Manifest struct{}

type ChaosMeshOrgStressChaosV1Alpha1ManifestData struct {
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
		ContainerNames *[]string `tfsdk:"container_names" json:"containerNames,omitempty"`
		Duration       *string   `tfsdk:"duration" json:"duration,omitempty"`
		Mode           *string   `tfsdk:"mode" json:"mode,omitempty"`
		RemoteCluster  *string   `tfsdk:"remote_cluster" json:"remoteCluster,omitempty"`
		Selector       *struct {
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
		StressngStressors *string `tfsdk:"stressng_stressors" json:"stressngStressors,omitempty"`
		Stressors         *struct {
			Cpu *struct {
				Load    *int64    `tfsdk:"load" json:"load,omitempty"`
				Options *[]string `tfsdk:"options" json:"options,omitempty"`
				Workers *int64    `tfsdk:"workers" json:"workers,omitempty"`
			} `tfsdk:"cpu" json:"cpu,omitempty"`
			Memory *struct {
				OomScoreAdj *int64    `tfsdk:"oom_score_adj" json:"oomScoreAdj,omitempty"`
				Options     *[]string `tfsdk:"options" json:"options,omitempty"`
				Size        *string   `tfsdk:"size" json:"size,omitempty"`
				Workers     *int64    `tfsdk:"workers" json:"workers,omitempty"`
			} `tfsdk:"memory" json:"memory,omitempty"`
		} `tfsdk:"stressors" json:"stressors,omitempty"`
		Value *string `tfsdk:"value" json:"value,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ChaosMeshOrgStressChaosV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_chaos_mesh_org_stress_chaos_v1alpha1_manifest"
}

func (r *ChaosMeshOrgStressChaosV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StressChaos is the Schema for the stresschaos API",
		MarkdownDescription: "StressChaos is the Schema for the stresschaos API",
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
				Description:         "Spec defines the behavior of a time chaos experiment",
				MarkdownDescription: "Spec defines the behavior of a time chaos experiment",
				Attributes: map[string]schema.Attribute{
					"container_names": schema.ListAttribute{
						Description:         "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
						MarkdownDescription: "ContainerNames indicates list of the name of affected container. If not set, the first container will be injected",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"duration": schema.StringAttribute{
						Description:         "Duration represents the duration of the chaos action",
						MarkdownDescription: "Duration represents the duration of the chaos action",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						MarkdownDescription: "Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("one", "all", "fixed", "fixed-percent", "random-max-percent"),
						},
					},

					"remote_cluster": schema.StringAttribute{
						Description:         "RemoteCluster represents the remote cluster where the chaos will be deployed",
						MarkdownDescription: "RemoteCluster represents the remote cluster where the chaos will be deployed",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
								Optional:            true,
								Computed:            false,
							},

							"expression_selectors": schema.ListNestedAttribute{
								Description:         "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								MarkdownDescription: "a slice of label selector expressions that can be used to select objects. A list of selectors based on set-based label expressions.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"field_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on fields.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on fields.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select objects. A selector based on labels.",
								MarkdownDescription: "Map of string keys and values that can be used to select objects. A selector based on labels.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespaces": schema.ListAttribute{
								Description:         "Namespaces is a set of namespace to which objects belong.",
								MarkdownDescription: "Namespaces is a set of namespace to which objects belong.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selectors": schema.MapAttribute{
								Description:         "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								MarkdownDescription: "Map of string keys and values that can be used to select nodes. Selector which must match a node's labels, and objects must belong to these selected nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"nodes": schema.ListAttribute{
								Description:         "Nodes is a set of node name and objects must belong to these nodes.",
								MarkdownDescription: "Nodes is a set of node name and objects must belong to these nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_phase_selectors": schema.ListAttribute{
								Description:         "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								MarkdownDescription: "PodPhaseSelectors is a set of condition of a pod at the current time. supported value: Pending / Running / Succeeded / Failed / Unknown",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pods": schema.MapAttribute{
								Description:         "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								MarkdownDescription: "Pods is a map of string keys and a set values that used to select pods. The key defines the namespace which pods belong, and the each values is a set of pod names.",
								ElementType:         types.ListType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"stressng_stressors": schema.StringAttribute{
						Description:         "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
						MarkdownDescription: "StressngStressors defines plenty of stressors just like 'Stressors' except that it's an experimental feature and more powerful. You can define stressors in 'stress-ng' (see also 'man stress-ng') dialect, however not all of the supported stressors are well tested. It maybe retired in later releases. You should always use 'Stressors' to define the stressors and use this only when you want more stressors unsupported by 'Stressors'. When both 'StressngStressors' and 'Stressors' are defined, 'StressngStressors' wins.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stressors": schema.SingleNestedAttribute{
						Description:         "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
						MarkdownDescription: "Stressors defines plenty of stressors supported to stress system components out. You can use one or more of them to make up various kinds of stresses. At least one of the stressors should be specified.",
						Attributes: map[string]schema.Attribute{
							"cpu": schema.SingleNestedAttribute{
								Description:         "CPUStressor stresses CPU out",
								MarkdownDescription: "CPUStressor stresses CPU out",
								Attributes: map[string]schema.Attribute{
									"load": schema.Int64Attribute{
										Description:         "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
										MarkdownDescription: "Load specifies P percent loading per CPU worker. 0 is effectively a sleep (no load) and 100 is full loading.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(100),
										},
									},

									"options": schema.ListAttribute{
										Description:         "extend stress-ng options",
										MarkdownDescription: "extend stress-ng options",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"workers": schema.Int64Attribute{
										Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
										MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtMost(8192),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"memory": schema.SingleNestedAttribute{
								Description:         "MemoryStressor stresses virtual memory out",
								MarkdownDescription: "MemoryStressor stresses virtual memory out",
								Attributes: map[string]schema.Attribute{
									"oom_score_adj": schema.Int64Attribute{
										Description:         "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
										MarkdownDescription: "OOMScoreAdj sets the oom_score_adj of the stress process. See 'man 5 proc' to know more about this option.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(-1000),
											int64validator.AtMost(1000),
										},
									},

									"options": schema.ListAttribute{
										Description:         "extend stress-ng options",
										MarkdownDescription: "extend stress-ng options",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"size": schema.StringAttribute{
										Description:         "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
										MarkdownDescription: "Size specifies N bytes consumed per vm worker, default is the total available memory. One can specify the size as % of total available memory or in units of B, KB/KiB, MB/MiB, GB/GiB, TB/TiB.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"workers": schema.Int64Attribute{
										Description:         "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
										MarkdownDescription: "Workers specifies N workers to apply the stressor. Maximum 8192 workers can run by stress-ng",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtMost(8192),
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

					"value": schema.StringAttribute{
						Description:         "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						MarkdownDescription: "Value is required when the mode is set to 'FixedMode' / 'FixedPercentMode' / 'RandomMaxPercentMode'. If 'FixedMode', provide an integer of pods to do chaos action. If 'FixedPercentMode', provide a number from 0-100 to specify the percent of pods the server can do chaos action. IF 'RandomMaxPercentMode',  provide a number from 0-100 to specify the max percent of pods to do chaos action",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ChaosMeshOrgStressChaosV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_chaos_mesh_org_stress_chaos_v1alpha1_manifest")

	var model ChaosMeshOrgStressChaosV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("chaos-mesh.org/v1alpha1")
	model.Kind = pointer.String("StressChaos")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
