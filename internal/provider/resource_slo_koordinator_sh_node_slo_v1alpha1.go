/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type SloKoordinatorShNodeSLOV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SloKoordinatorShNodeSLOV1Alpha1Resource)(nil)
)

type SloKoordinatorShNodeSLOV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SloKoordinatorShNodeSLOV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		CpuBurstStrategy *struct {
			CfsQuotaBurstPercent *int64 `tfsdk:"cfs_quota_burst_percent" yaml:"cfsQuotaBurstPercent,omitempty"`

			CfsQuotaBurstPeriodSeconds *int64 `tfsdk:"cfs_quota_burst_period_seconds" yaml:"cfsQuotaBurstPeriodSeconds,omitempty"`

			CpuBurstPercent *int64 `tfsdk:"cpu_burst_percent" yaml:"cpuBurstPercent,omitempty"`

			Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`

			SharePoolThresholdPercent *int64 `tfsdk:"share_pool_threshold_percent" yaml:"sharePoolThresholdPercent,omitempty"`
		} `tfsdk:"cpu_burst_strategy" yaml:"cpuBurstStrategy,omitempty"`

		Extensions utilities.Dynamic `tfsdk:"extensions" yaml:"extensions,omitempty"`

		ResourceQOSStrategy *struct {
			BeClass *struct {
				CpuQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					GroupIdentity *int64 `tfsdk:"group_identity" yaml:"groupIdentity,omitempty"`
				} `tfsdk:"cpu_qos" yaml:"cpuQOS,omitempty"`

				MemoryQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					LowLimitPercent *int64 `tfsdk:"low_limit_percent" yaml:"lowLimitPercent,omitempty"`

					MinLimitPercent *int64 `tfsdk:"min_limit_percent" yaml:"minLimitPercent,omitempty"`

					OomKillGroup *int64 `tfsdk:"oom_kill_group" yaml:"oomKillGroup,omitempty"`

					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					PriorityEnable *int64 `tfsdk:"priority_enable" yaml:"priorityEnable,omitempty"`

					ThrottlingPercent *int64 `tfsdk:"throttling_percent" yaml:"throttlingPercent,omitempty"`

					WmarkMinAdj *int64 `tfsdk:"wmark_min_adj" yaml:"wmarkMinAdj,omitempty"`

					WmarkRatio *int64 `tfsdk:"wmark_ratio" yaml:"wmarkRatio,omitempty"`

					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" yaml:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" yaml:"memoryQOS,omitempty"`

				ResctrlQOS *struct {
					CatRangeEndPercent *int64 `tfsdk:"cat_range_end_percent" yaml:"catRangeEndPercent,omitempty"`

					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" yaml:"catRangeStartPercent,omitempty"`

					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					MbaPercent *int64 `tfsdk:"mba_percent" yaml:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" yaml:"resctrlQOS,omitempty"`
			} `tfsdk:"be_class" yaml:"beClass,omitempty"`

			CgroupRoot *struct {
				CpuQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					GroupIdentity *int64 `tfsdk:"group_identity" yaml:"groupIdentity,omitempty"`
				} `tfsdk:"cpu_qos" yaml:"cpuQOS,omitempty"`

				MemoryQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					LowLimitPercent *int64 `tfsdk:"low_limit_percent" yaml:"lowLimitPercent,omitempty"`

					MinLimitPercent *int64 `tfsdk:"min_limit_percent" yaml:"minLimitPercent,omitempty"`

					OomKillGroup *int64 `tfsdk:"oom_kill_group" yaml:"oomKillGroup,omitempty"`

					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					PriorityEnable *int64 `tfsdk:"priority_enable" yaml:"priorityEnable,omitempty"`

					ThrottlingPercent *int64 `tfsdk:"throttling_percent" yaml:"throttlingPercent,omitempty"`

					WmarkMinAdj *int64 `tfsdk:"wmark_min_adj" yaml:"wmarkMinAdj,omitempty"`

					WmarkRatio *int64 `tfsdk:"wmark_ratio" yaml:"wmarkRatio,omitempty"`

					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" yaml:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" yaml:"memoryQOS,omitempty"`

				ResctrlQOS *struct {
					CatRangeEndPercent *int64 `tfsdk:"cat_range_end_percent" yaml:"catRangeEndPercent,omitempty"`

					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" yaml:"catRangeStartPercent,omitempty"`

					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					MbaPercent *int64 `tfsdk:"mba_percent" yaml:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" yaml:"resctrlQOS,omitempty"`
			} `tfsdk:"cgroup_root" yaml:"cgroupRoot,omitempty"`

			LsClass *struct {
				CpuQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					GroupIdentity *int64 `tfsdk:"group_identity" yaml:"groupIdentity,omitempty"`
				} `tfsdk:"cpu_qos" yaml:"cpuQOS,omitempty"`

				MemoryQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					LowLimitPercent *int64 `tfsdk:"low_limit_percent" yaml:"lowLimitPercent,omitempty"`

					MinLimitPercent *int64 `tfsdk:"min_limit_percent" yaml:"minLimitPercent,omitempty"`

					OomKillGroup *int64 `tfsdk:"oom_kill_group" yaml:"oomKillGroup,omitempty"`

					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					PriorityEnable *int64 `tfsdk:"priority_enable" yaml:"priorityEnable,omitempty"`

					ThrottlingPercent *int64 `tfsdk:"throttling_percent" yaml:"throttlingPercent,omitempty"`

					WmarkMinAdj *int64 `tfsdk:"wmark_min_adj" yaml:"wmarkMinAdj,omitempty"`

					WmarkRatio *int64 `tfsdk:"wmark_ratio" yaml:"wmarkRatio,omitempty"`

					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" yaml:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" yaml:"memoryQOS,omitempty"`

				ResctrlQOS *struct {
					CatRangeEndPercent *int64 `tfsdk:"cat_range_end_percent" yaml:"catRangeEndPercent,omitempty"`

					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" yaml:"catRangeStartPercent,omitempty"`

					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					MbaPercent *int64 `tfsdk:"mba_percent" yaml:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" yaml:"resctrlQOS,omitempty"`
			} `tfsdk:"ls_class" yaml:"lsClass,omitempty"`

			LsrClass *struct {
				CpuQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					GroupIdentity *int64 `tfsdk:"group_identity" yaml:"groupIdentity,omitempty"`
				} `tfsdk:"cpu_qos" yaml:"cpuQOS,omitempty"`

				MemoryQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					LowLimitPercent *int64 `tfsdk:"low_limit_percent" yaml:"lowLimitPercent,omitempty"`

					MinLimitPercent *int64 `tfsdk:"min_limit_percent" yaml:"minLimitPercent,omitempty"`

					OomKillGroup *int64 `tfsdk:"oom_kill_group" yaml:"oomKillGroup,omitempty"`

					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					PriorityEnable *int64 `tfsdk:"priority_enable" yaml:"priorityEnable,omitempty"`

					ThrottlingPercent *int64 `tfsdk:"throttling_percent" yaml:"throttlingPercent,omitempty"`

					WmarkMinAdj *int64 `tfsdk:"wmark_min_adj" yaml:"wmarkMinAdj,omitempty"`

					WmarkRatio *int64 `tfsdk:"wmark_ratio" yaml:"wmarkRatio,omitempty"`

					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" yaml:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" yaml:"memoryQOS,omitempty"`

				ResctrlQOS *struct {
					CatRangeEndPercent *int64 `tfsdk:"cat_range_end_percent" yaml:"catRangeEndPercent,omitempty"`

					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" yaml:"catRangeStartPercent,omitempty"`

					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					MbaPercent *int64 `tfsdk:"mba_percent" yaml:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" yaml:"resctrlQOS,omitempty"`
			} `tfsdk:"lsr_class" yaml:"lsrClass,omitempty"`

			SystemClass *struct {
				CpuQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					GroupIdentity *int64 `tfsdk:"group_identity" yaml:"groupIdentity,omitempty"`
				} `tfsdk:"cpu_qos" yaml:"cpuQOS,omitempty"`

				MemoryQOS *struct {
					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					LowLimitPercent *int64 `tfsdk:"low_limit_percent" yaml:"lowLimitPercent,omitempty"`

					MinLimitPercent *int64 `tfsdk:"min_limit_percent" yaml:"minLimitPercent,omitempty"`

					OomKillGroup *int64 `tfsdk:"oom_kill_group" yaml:"oomKillGroup,omitempty"`

					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					PriorityEnable *int64 `tfsdk:"priority_enable" yaml:"priorityEnable,omitempty"`

					ThrottlingPercent *int64 `tfsdk:"throttling_percent" yaml:"throttlingPercent,omitempty"`

					WmarkMinAdj *int64 `tfsdk:"wmark_min_adj" yaml:"wmarkMinAdj,omitempty"`

					WmarkRatio *int64 `tfsdk:"wmark_ratio" yaml:"wmarkRatio,omitempty"`

					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" yaml:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" yaml:"memoryQOS,omitempty"`

				ResctrlQOS *struct {
					CatRangeEndPercent *int64 `tfsdk:"cat_range_end_percent" yaml:"catRangeEndPercent,omitempty"`

					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" yaml:"catRangeStartPercent,omitempty"`

					Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

					MbaPercent *int64 `tfsdk:"mba_percent" yaml:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" yaml:"resctrlQOS,omitempty"`
			} `tfsdk:"system_class" yaml:"systemClass,omitempty"`
		} `tfsdk:"resource_qos_strategy" yaml:"resourceQOSStrategy,omitempty"`

		ResourceUsedThresholdWithBE *struct {
			CpuEvictBESatisfactionLowerPercent *int64 `tfsdk:"cpu_evict_be_satisfaction_lower_percent" yaml:"cpuEvictBESatisfactionLowerPercent,omitempty"`

			CpuEvictBESatisfactionUpperPercent *int64 `tfsdk:"cpu_evict_be_satisfaction_upper_percent" yaml:"cpuEvictBESatisfactionUpperPercent,omitempty"`

			CpuEvictTimeWindowSeconds *int64 `tfsdk:"cpu_evict_time_window_seconds" yaml:"cpuEvictTimeWindowSeconds,omitempty"`

			CpuSuppressPolicy *string `tfsdk:"cpu_suppress_policy" yaml:"cpuSuppressPolicy,omitempty"`

			CpuSuppressThresholdPercent *int64 `tfsdk:"cpu_suppress_threshold_percent" yaml:"cpuSuppressThresholdPercent,omitempty"`

			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

			MemoryEvictLowerPercent *int64 `tfsdk:"memory_evict_lower_percent" yaml:"memoryEvictLowerPercent,omitempty"`

			MemoryEvictThresholdPercent *int64 `tfsdk:"memory_evict_threshold_percent" yaml:"memoryEvictThresholdPercent,omitempty"`
		} `tfsdk:"resource_used_threshold_with_be" yaml:"resourceUsedThresholdWithBE,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSloKoordinatorShNodeSLOV1Alpha1Resource() resource.Resource {
	return &SloKoordinatorShNodeSLOV1Alpha1Resource{}
}

func (r *SloKoordinatorShNodeSLOV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_slo_koordinator_sh_node_slo_v1alpha1"
}

func (r *SloKoordinatorShNodeSLOV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "NodeSLO is the Schema for the nodeslos API",
		MarkdownDescription: "NodeSLO is the Schema for the nodeslos API",
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
				Description:         "NodeSLOSpec defines the desired state of NodeSLO",
				MarkdownDescription: "NodeSLOSpec defines the desired state of NodeSLO",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"cpu_burst_strategy": {
						Description:         "CPU Burst Strategy",
						MarkdownDescription: "CPU Burst Strategy",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cfs_quota_burst_percent": {
								Description:         "pod cfs quota scale up ceil percentage, default = 300 (300%)",
								MarkdownDescription: "pod cfs quota scale up ceil percentage, default = 300 (300%)",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cfs_quota_burst_period_seconds": {
								Description:         "specifies a period of time for pod can use at burst, default = -1 (unlimited)",
								MarkdownDescription: "specifies a period of time for pod can use at burst, default = -1 (unlimited)",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cpu_burst_percent": {
								Description:         "cpu burst percentage for setting cpu.cfs_burst_us, legal range: [0, 10000], default as 1000 (1000%)",
								MarkdownDescription: "cpu burst percentage for setting cpu.cfs_burst_us, legal range: [0, 10000], default as 1000 (1000%)",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(10000),
								},
							},

							"policy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"share_pool_threshold_percent": {
								Description:         "scale down cfs quota if node cpu overload, default = 50",
								MarkdownDescription: "scale down cfs quota if node cpu overload, default = 50",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"extensions": {
						Description:         "Third party extensions for NodeSLO",
						MarkdownDescription: "Third party extensions for NodeSLO",

						Type: utilities.DynamicType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_qos_strategy": {
						Description:         "QoS config strategy for pods of different qos-class",
						MarkdownDescription: "QoS config strategy for pods of different qos-class",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"be_class": {
								Description:         "ResourceQOS for BE pods.",
								MarkdownDescription: "ResourceQOS for BE pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu_qos": {
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"group_identity": {
												Description:         "group identity value for pods, default = 0",
												MarkdownDescription: "group identity value for pods, default = 0",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": {
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"low_limit_percent": {
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": {
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority_enable": {
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"throttling_percent": {
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": {
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(-25),

													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": {
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": {
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),

													int64validator.AtMost(1000),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": {
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cat_range_end_percent": {
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": {
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"enable": {
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mba_percent": {
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
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

							"cgroup_root": {
								Description:         "ResourceQOS for root cgroup.",
								MarkdownDescription: "ResourceQOS for root cgroup.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu_qos": {
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"group_identity": {
												Description:         "group identity value for pods, default = 0",
												MarkdownDescription: "group identity value for pods, default = 0",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": {
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"low_limit_percent": {
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": {
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority_enable": {
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"throttling_percent": {
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": {
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(-25),

													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": {
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": {
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),

													int64validator.AtMost(1000),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": {
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cat_range_end_percent": {
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": {
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"enable": {
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mba_percent": {
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
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

							"ls_class": {
								Description:         "ResourceQOS for LS pods.",
								MarkdownDescription: "ResourceQOS for LS pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu_qos": {
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"group_identity": {
												Description:         "group identity value for pods, default = 0",
												MarkdownDescription: "group identity value for pods, default = 0",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": {
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"low_limit_percent": {
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": {
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority_enable": {
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"throttling_percent": {
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": {
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(-25),

													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": {
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": {
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),

													int64validator.AtMost(1000),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": {
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cat_range_end_percent": {
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": {
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"enable": {
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mba_percent": {
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
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

							"lsr_class": {
								Description:         "ResourceQOS for LSR pods.",
								MarkdownDescription: "ResourceQOS for LSR pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu_qos": {
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"group_identity": {
												Description:         "group identity value for pods, default = 0",
												MarkdownDescription: "group identity value for pods, default = 0",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": {
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"low_limit_percent": {
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": {
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority_enable": {
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"throttling_percent": {
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": {
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(-25),

													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": {
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": {
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),

													int64validator.AtMost(1000),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": {
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cat_range_end_percent": {
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": {
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"enable": {
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mba_percent": {
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
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

							"system_class": {
								Description:         "ResourceQOS for system pods",
								MarkdownDescription: "ResourceQOS for system pods",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu_qos": {
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"group_identity": {
												Description:         "group identity value for pods, default = 0",
												MarkdownDescription: "group identity value for pods, default = 0",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": {
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enable": {
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"low_limit_percent": {
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": {
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := memory.limit_in_bytes * throttlingFactor / 100 (use 'max' if memory.high <= memory.min) MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority_enable": {
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"throttling_percent": {
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": {
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(-25),

													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": {
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": {
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),

													int64validator.AtMost(1000),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": {
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cat_range_end_percent": {
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": {
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
											},

											"enable": {
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mba_percent": {
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(100),
												},
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_used_threshold_with_be": {
						Description:         "BE pods will be limited if node resource usage overload",
						MarkdownDescription: "BE pods will be limited if node resource usage overload",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cpu_evict_be_satisfaction_lower_percent": {
								Description:         "if be CPU (RealLimit/allocatedLimit < CPUEvictBESatisfactionLowerPercent and usage nearly 100%) continue CPUEvictTimeWindowSeconds,then start evict",
								MarkdownDescription: "if be CPU (RealLimit/allocatedLimit < CPUEvictBESatisfactionLowerPercent and usage nearly 100%) continue CPUEvictTimeWindowSeconds,then start evict",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cpu_evict_be_satisfaction_upper_percent": {
								Description:         "if be CPU RealLimit/allocatedLimit > CPUEvictBESatisfactionUpperPercent, then stop evict BE pods",
								MarkdownDescription: "if be CPU RealLimit/allocatedLimit > CPUEvictBESatisfactionUpperPercent, then stop evict BE pods",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cpu_evict_time_window_seconds": {
								Description:         "cpu evict start after continue avg(cpuusage) > CPUEvictThresholdPercent in seconds",
								MarkdownDescription: "cpu evict start after continue avg(cpuusage) > CPUEvictThresholdPercent in seconds",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cpu_suppress_policy": {
								Description:         "CPUSuppressPolicy",
								MarkdownDescription: "CPUSuppressPolicy",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cpu_suppress_threshold_percent": {
								Description:         "cpu suppress threshold percentage (0,100), default = 65",
								MarkdownDescription: "cpu suppress threshold percentage (0,100), default = 65",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(100),
								},
							},

							"enable": {
								Description:         "whether the strategy is enabled, default = false",
								MarkdownDescription: "whether the strategy is enabled, default = false",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"memory_evict_lower_percent": {
								Description:         "lower: memory release util usage under MemoryEvictLowerPercent, default = MemoryEvictThresholdPercent - 2",
								MarkdownDescription: "lower: memory release util usage under MemoryEvictLowerPercent, default = MemoryEvictThresholdPercent - 2",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(100),
								},
							},

							"memory_evict_threshold_percent": {
								Description:         "upper: memory evict threshold percentage (0,100), default = 70",
								MarkdownDescription: "upper: memory evict threshold percentage (0,100), default = 70",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(100),
								},
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
		},
	}, nil
}

func (r *SloKoordinatorShNodeSLOV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_slo_koordinator_sh_node_slo_v1alpha1")

	var state SloKoordinatorShNodeSLOV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SloKoordinatorShNodeSLOV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("slo.koordinator.sh/v1alpha1")
	goModel.Kind = utilities.Ptr("NodeSLO")

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

func (r *SloKoordinatorShNodeSLOV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_slo_koordinator_sh_node_slo_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SloKoordinatorShNodeSLOV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_slo_koordinator_sh_node_slo_v1alpha1")

	var state SloKoordinatorShNodeSLOV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SloKoordinatorShNodeSLOV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("slo.koordinator.sh/v1alpha1")
	goModel.Kind = utilities.Ptr("NodeSLO")

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

func (r *SloKoordinatorShNodeSLOV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_slo_koordinator_sh_node_slo_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
