/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package slo_koordinator_sh_v1alpha1

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
	_ datasource.DataSource = &SloKoordinatorShNodeSloV1Alpha1Manifest{}
)

func NewSloKoordinatorShNodeSloV1Alpha1Manifest() datasource.DataSource {
	return &SloKoordinatorShNodeSloV1Alpha1Manifest{}
}

type SloKoordinatorShNodeSloV1Alpha1Manifest struct{}

type SloKoordinatorShNodeSloV1Alpha1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CpuBurstStrategy *struct {
			CfsQuotaBurstPercent       *int64  `tfsdk:"cfs_quota_burst_percent" json:"cfsQuotaBurstPercent,omitempty"`
			CfsQuotaBurstPeriodSeconds *int64  `tfsdk:"cfs_quota_burst_period_seconds" json:"cfsQuotaBurstPeriodSeconds,omitempty"`
			CpuBurstPercent            *int64  `tfsdk:"cpu_burst_percent" json:"cpuBurstPercent,omitempty"`
			Policy                     *string `tfsdk:"policy" json:"policy,omitempty"`
			SharePoolThresholdPercent  *int64  `tfsdk:"share_pool_threshold_percent" json:"sharePoolThresholdPercent,omitempty"`
		} `tfsdk:"cpu_burst_strategy" json:"cpuBurstStrategy,omitempty"`
		Extensions       *map[string]string `tfsdk:"extensions" json:"extensions,omitempty"`
		HostApplications *[]struct {
			CgroupPath *struct {
				Base         *string `tfsdk:"base" json:"base,omitempty"`
				ParentDir    *string `tfsdk:"parent_dir" json:"parentDir,omitempty"`
				RelativePath *string `tfsdk:"relative_path" json:"relativePath,omitempty"`
			} `tfsdk:"cgroup_path" json:"cgroupPath,omitempty"`
			Name     *string            `tfsdk:"name" json:"name,omitempty"`
			Priority *string            `tfsdk:"priority" json:"priority,omitempty"`
			Qos      *string            `tfsdk:"qos" json:"qos,omitempty"`
			Strategy *map[string]string `tfsdk:"strategy" json:"strategy,omitempty"`
		} `tfsdk:"host_applications" json:"hostApplications,omitempty"`
		ResourceQOSStrategy *struct {
			BeClass *struct {
				BlkioQOS *struct {
					Blocks *[]struct {
						IoCfg *struct {
							EnableUserModel     *bool  `tfsdk:"enable_user_model" json:"enableUserModel,omitempty"`
							IoWeightPercent     *int64 `tfsdk:"io_weight_percent" json:"ioWeightPercent,omitempty"`
							ModelReadBPS        *int64 `tfsdk:"model_read_bps" json:"modelReadBPS,omitempty"`
							ModelReadRandIOPS   *int64 `tfsdk:"model_read_rand_iops" json:"modelReadRandIOPS,omitempty"`
							ModelReadSeqIOPS    *int64 `tfsdk:"model_read_seq_iops" json:"modelReadSeqIOPS,omitempty"`
							ModelWriteBPS       *int64 `tfsdk:"model_write_bps" json:"modelWriteBPS,omitempty"`
							ModelWriteRandIOPS  *int64 `tfsdk:"model_write_rand_iops" json:"modelWriteRandIOPS,omitempty"`
							ModelWriteSeqIOPS   *int64 `tfsdk:"model_write_seq_iops" json:"modelWriteSeqIOPS,omitempty"`
							ReadBPS             *int64 `tfsdk:"read_bps" json:"readBPS,omitempty"`
							ReadIOPS            *int64 `tfsdk:"read_iops" json:"readIOPS,omitempty"`
							ReadLatency         *int64 `tfsdk:"read_latency" json:"readLatency,omitempty"`
							ReadLatencyPercent  *int64 `tfsdk:"read_latency_percent" json:"readLatencyPercent,omitempty"`
							WriteBPS            *int64 `tfsdk:"write_bps" json:"writeBPS,omitempty"`
							WriteIOPS           *int64 `tfsdk:"write_iops" json:"writeIOPS,omitempty"`
							WriteLatency        *int64 `tfsdk:"write_latency" json:"writeLatency,omitempty"`
							WriteLatencyPercent *int64 `tfsdk:"write_latency_percent" json:"writeLatencyPercent,omitempty"`
						} `tfsdk:"io_cfg" json:"ioCfg,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"blocks" json:"blocks,omitempty"`
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"blkio_qos" json:"blkioQOS,omitempty"`
				CpuQOS *struct {
					CoreExpeller  *bool  `tfsdk:"core_expeller" json:"coreExpeller,omitempty"`
					Enable        *bool  `tfsdk:"enable" json:"enable,omitempty"`
					GroupIdentity *int64 `tfsdk:"group_identity" json:"groupIdentity,omitempty"`
					SchedIdle     *int64 `tfsdk:"sched_idle" json:"schedIdle,omitempty"`
				} `tfsdk:"cpu_qos" json:"cpuQOS,omitempty"`
				MemoryQOS *struct {
					Enable            *bool  `tfsdk:"enable" json:"enable,omitempty"`
					LowLimitPercent   *int64 `tfsdk:"low_limit_percent" json:"lowLimitPercent,omitempty"`
					MinLimitPercent   *int64 `tfsdk:"min_limit_percent" json:"minLimitPercent,omitempty"`
					OomKillGroup      *int64 `tfsdk:"oom_kill_group" json:"oomKillGroup,omitempty"`
					Priority          *int64 `tfsdk:"priority" json:"priority,omitempty"`
					PriorityEnable    *int64 `tfsdk:"priority_enable" json:"priorityEnable,omitempty"`
					ThrottlingPercent *int64 `tfsdk:"throttling_percent" json:"throttlingPercent,omitempty"`
					WmarkMinAdj       *int64 `tfsdk:"wmark_min_adj" json:"wmarkMinAdj,omitempty"`
					WmarkRatio        *int64 `tfsdk:"wmark_ratio" json:"wmarkRatio,omitempty"`
					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" json:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" json:"memoryQOS,omitempty"`
				NetworkQOS *struct {
					EgressLimit    *string `tfsdk:"egress_limit" json:"egressLimit,omitempty"`
					EgressRequest  *string `tfsdk:"egress_request" json:"egressRequest,omitempty"`
					Enable         *bool   `tfsdk:"enable" json:"enable,omitempty"`
					IngressLimit   *string `tfsdk:"ingress_limit" json:"ingressLimit,omitempty"`
					IngressRequest *string `tfsdk:"ingress_request" json:"ingressRequest,omitempty"`
				} `tfsdk:"network_qos" json:"networkQOS,omitempty"`
				ResctrlQOS *struct {
					CatRangeEndPercent   *int64 `tfsdk:"cat_range_end_percent" json:"catRangeEndPercent,omitempty"`
					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" json:"catRangeStartPercent,omitempty"`
					Enable               *bool  `tfsdk:"enable" json:"enable,omitempty"`
					MbaPercent           *int64 `tfsdk:"mba_percent" json:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" json:"resctrlQOS,omitempty"`
			} `tfsdk:"be_class" json:"beClass,omitempty"`
			CgroupRoot *struct {
				BlkioQOS *struct {
					Blocks *[]struct {
						IoCfg *struct {
							EnableUserModel     *bool  `tfsdk:"enable_user_model" json:"enableUserModel,omitempty"`
							IoWeightPercent     *int64 `tfsdk:"io_weight_percent" json:"ioWeightPercent,omitempty"`
							ModelReadBPS        *int64 `tfsdk:"model_read_bps" json:"modelReadBPS,omitempty"`
							ModelReadRandIOPS   *int64 `tfsdk:"model_read_rand_iops" json:"modelReadRandIOPS,omitempty"`
							ModelReadSeqIOPS    *int64 `tfsdk:"model_read_seq_iops" json:"modelReadSeqIOPS,omitempty"`
							ModelWriteBPS       *int64 `tfsdk:"model_write_bps" json:"modelWriteBPS,omitempty"`
							ModelWriteRandIOPS  *int64 `tfsdk:"model_write_rand_iops" json:"modelWriteRandIOPS,omitempty"`
							ModelWriteSeqIOPS   *int64 `tfsdk:"model_write_seq_iops" json:"modelWriteSeqIOPS,omitempty"`
							ReadBPS             *int64 `tfsdk:"read_bps" json:"readBPS,omitempty"`
							ReadIOPS            *int64 `tfsdk:"read_iops" json:"readIOPS,omitempty"`
							ReadLatency         *int64 `tfsdk:"read_latency" json:"readLatency,omitempty"`
							ReadLatencyPercent  *int64 `tfsdk:"read_latency_percent" json:"readLatencyPercent,omitempty"`
							WriteBPS            *int64 `tfsdk:"write_bps" json:"writeBPS,omitempty"`
							WriteIOPS           *int64 `tfsdk:"write_iops" json:"writeIOPS,omitempty"`
							WriteLatency        *int64 `tfsdk:"write_latency" json:"writeLatency,omitempty"`
							WriteLatencyPercent *int64 `tfsdk:"write_latency_percent" json:"writeLatencyPercent,omitempty"`
						} `tfsdk:"io_cfg" json:"ioCfg,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"blocks" json:"blocks,omitempty"`
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"blkio_qos" json:"blkioQOS,omitempty"`
				CpuQOS *struct {
					CoreExpeller  *bool  `tfsdk:"core_expeller" json:"coreExpeller,omitempty"`
					Enable        *bool  `tfsdk:"enable" json:"enable,omitempty"`
					GroupIdentity *int64 `tfsdk:"group_identity" json:"groupIdentity,omitempty"`
					SchedIdle     *int64 `tfsdk:"sched_idle" json:"schedIdle,omitempty"`
				} `tfsdk:"cpu_qos" json:"cpuQOS,omitempty"`
				MemoryQOS *struct {
					Enable            *bool  `tfsdk:"enable" json:"enable,omitempty"`
					LowLimitPercent   *int64 `tfsdk:"low_limit_percent" json:"lowLimitPercent,omitempty"`
					MinLimitPercent   *int64 `tfsdk:"min_limit_percent" json:"minLimitPercent,omitempty"`
					OomKillGroup      *int64 `tfsdk:"oom_kill_group" json:"oomKillGroup,omitempty"`
					Priority          *int64 `tfsdk:"priority" json:"priority,omitempty"`
					PriorityEnable    *int64 `tfsdk:"priority_enable" json:"priorityEnable,omitempty"`
					ThrottlingPercent *int64 `tfsdk:"throttling_percent" json:"throttlingPercent,omitempty"`
					WmarkMinAdj       *int64 `tfsdk:"wmark_min_adj" json:"wmarkMinAdj,omitempty"`
					WmarkRatio        *int64 `tfsdk:"wmark_ratio" json:"wmarkRatio,omitempty"`
					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" json:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" json:"memoryQOS,omitempty"`
				NetworkQOS *struct {
					EgressLimit    *string `tfsdk:"egress_limit" json:"egressLimit,omitempty"`
					EgressRequest  *string `tfsdk:"egress_request" json:"egressRequest,omitempty"`
					Enable         *bool   `tfsdk:"enable" json:"enable,omitempty"`
					IngressLimit   *string `tfsdk:"ingress_limit" json:"ingressLimit,omitempty"`
					IngressRequest *string `tfsdk:"ingress_request" json:"ingressRequest,omitempty"`
				} `tfsdk:"network_qos" json:"networkQOS,omitempty"`
				ResctrlQOS *struct {
					CatRangeEndPercent   *int64 `tfsdk:"cat_range_end_percent" json:"catRangeEndPercent,omitempty"`
					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" json:"catRangeStartPercent,omitempty"`
					Enable               *bool  `tfsdk:"enable" json:"enable,omitempty"`
					MbaPercent           *int64 `tfsdk:"mba_percent" json:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" json:"resctrlQOS,omitempty"`
			} `tfsdk:"cgroup_root" json:"cgroupRoot,omitempty"`
			LsClass *struct {
				BlkioQOS *struct {
					Blocks *[]struct {
						IoCfg *struct {
							EnableUserModel     *bool  `tfsdk:"enable_user_model" json:"enableUserModel,omitempty"`
							IoWeightPercent     *int64 `tfsdk:"io_weight_percent" json:"ioWeightPercent,omitempty"`
							ModelReadBPS        *int64 `tfsdk:"model_read_bps" json:"modelReadBPS,omitempty"`
							ModelReadRandIOPS   *int64 `tfsdk:"model_read_rand_iops" json:"modelReadRandIOPS,omitempty"`
							ModelReadSeqIOPS    *int64 `tfsdk:"model_read_seq_iops" json:"modelReadSeqIOPS,omitempty"`
							ModelWriteBPS       *int64 `tfsdk:"model_write_bps" json:"modelWriteBPS,omitempty"`
							ModelWriteRandIOPS  *int64 `tfsdk:"model_write_rand_iops" json:"modelWriteRandIOPS,omitempty"`
							ModelWriteSeqIOPS   *int64 `tfsdk:"model_write_seq_iops" json:"modelWriteSeqIOPS,omitempty"`
							ReadBPS             *int64 `tfsdk:"read_bps" json:"readBPS,omitempty"`
							ReadIOPS            *int64 `tfsdk:"read_iops" json:"readIOPS,omitempty"`
							ReadLatency         *int64 `tfsdk:"read_latency" json:"readLatency,omitempty"`
							ReadLatencyPercent  *int64 `tfsdk:"read_latency_percent" json:"readLatencyPercent,omitempty"`
							WriteBPS            *int64 `tfsdk:"write_bps" json:"writeBPS,omitempty"`
							WriteIOPS           *int64 `tfsdk:"write_iops" json:"writeIOPS,omitempty"`
							WriteLatency        *int64 `tfsdk:"write_latency" json:"writeLatency,omitempty"`
							WriteLatencyPercent *int64 `tfsdk:"write_latency_percent" json:"writeLatencyPercent,omitempty"`
						} `tfsdk:"io_cfg" json:"ioCfg,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"blocks" json:"blocks,omitempty"`
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"blkio_qos" json:"blkioQOS,omitempty"`
				CpuQOS *struct {
					CoreExpeller  *bool  `tfsdk:"core_expeller" json:"coreExpeller,omitempty"`
					Enable        *bool  `tfsdk:"enable" json:"enable,omitempty"`
					GroupIdentity *int64 `tfsdk:"group_identity" json:"groupIdentity,omitempty"`
					SchedIdle     *int64 `tfsdk:"sched_idle" json:"schedIdle,omitempty"`
				} `tfsdk:"cpu_qos" json:"cpuQOS,omitempty"`
				MemoryQOS *struct {
					Enable            *bool  `tfsdk:"enable" json:"enable,omitempty"`
					LowLimitPercent   *int64 `tfsdk:"low_limit_percent" json:"lowLimitPercent,omitempty"`
					MinLimitPercent   *int64 `tfsdk:"min_limit_percent" json:"minLimitPercent,omitempty"`
					OomKillGroup      *int64 `tfsdk:"oom_kill_group" json:"oomKillGroup,omitempty"`
					Priority          *int64 `tfsdk:"priority" json:"priority,omitempty"`
					PriorityEnable    *int64 `tfsdk:"priority_enable" json:"priorityEnable,omitempty"`
					ThrottlingPercent *int64 `tfsdk:"throttling_percent" json:"throttlingPercent,omitempty"`
					WmarkMinAdj       *int64 `tfsdk:"wmark_min_adj" json:"wmarkMinAdj,omitempty"`
					WmarkRatio        *int64 `tfsdk:"wmark_ratio" json:"wmarkRatio,omitempty"`
					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" json:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" json:"memoryQOS,omitempty"`
				NetworkQOS *struct {
					EgressLimit    *string `tfsdk:"egress_limit" json:"egressLimit,omitempty"`
					EgressRequest  *string `tfsdk:"egress_request" json:"egressRequest,omitempty"`
					Enable         *bool   `tfsdk:"enable" json:"enable,omitempty"`
					IngressLimit   *string `tfsdk:"ingress_limit" json:"ingressLimit,omitempty"`
					IngressRequest *string `tfsdk:"ingress_request" json:"ingressRequest,omitempty"`
				} `tfsdk:"network_qos" json:"networkQOS,omitempty"`
				ResctrlQOS *struct {
					CatRangeEndPercent   *int64 `tfsdk:"cat_range_end_percent" json:"catRangeEndPercent,omitempty"`
					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" json:"catRangeStartPercent,omitempty"`
					Enable               *bool  `tfsdk:"enable" json:"enable,omitempty"`
					MbaPercent           *int64 `tfsdk:"mba_percent" json:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" json:"resctrlQOS,omitempty"`
			} `tfsdk:"ls_class" json:"lsClass,omitempty"`
			LsrClass *struct {
				BlkioQOS *struct {
					Blocks *[]struct {
						IoCfg *struct {
							EnableUserModel     *bool  `tfsdk:"enable_user_model" json:"enableUserModel,omitempty"`
							IoWeightPercent     *int64 `tfsdk:"io_weight_percent" json:"ioWeightPercent,omitempty"`
							ModelReadBPS        *int64 `tfsdk:"model_read_bps" json:"modelReadBPS,omitempty"`
							ModelReadRandIOPS   *int64 `tfsdk:"model_read_rand_iops" json:"modelReadRandIOPS,omitempty"`
							ModelReadSeqIOPS    *int64 `tfsdk:"model_read_seq_iops" json:"modelReadSeqIOPS,omitempty"`
							ModelWriteBPS       *int64 `tfsdk:"model_write_bps" json:"modelWriteBPS,omitempty"`
							ModelWriteRandIOPS  *int64 `tfsdk:"model_write_rand_iops" json:"modelWriteRandIOPS,omitempty"`
							ModelWriteSeqIOPS   *int64 `tfsdk:"model_write_seq_iops" json:"modelWriteSeqIOPS,omitempty"`
							ReadBPS             *int64 `tfsdk:"read_bps" json:"readBPS,omitempty"`
							ReadIOPS            *int64 `tfsdk:"read_iops" json:"readIOPS,omitempty"`
							ReadLatency         *int64 `tfsdk:"read_latency" json:"readLatency,omitempty"`
							ReadLatencyPercent  *int64 `tfsdk:"read_latency_percent" json:"readLatencyPercent,omitempty"`
							WriteBPS            *int64 `tfsdk:"write_bps" json:"writeBPS,omitempty"`
							WriteIOPS           *int64 `tfsdk:"write_iops" json:"writeIOPS,omitempty"`
							WriteLatency        *int64 `tfsdk:"write_latency" json:"writeLatency,omitempty"`
							WriteLatencyPercent *int64 `tfsdk:"write_latency_percent" json:"writeLatencyPercent,omitempty"`
						} `tfsdk:"io_cfg" json:"ioCfg,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"blocks" json:"blocks,omitempty"`
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"blkio_qos" json:"blkioQOS,omitempty"`
				CpuQOS *struct {
					CoreExpeller  *bool  `tfsdk:"core_expeller" json:"coreExpeller,omitempty"`
					Enable        *bool  `tfsdk:"enable" json:"enable,omitempty"`
					GroupIdentity *int64 `tfsdk:"group_identity" json:"groupIdentity,omitempty"`
					SchedIdle     *int64 `tfsdk:"sched_idle" json:"schedIdle,omitempty"`
				} `tfsdk:"cpu_qos" json:"cpuQOS,omitempty"`
				MemoryQOS *struct {
					Enable            *bool  `tfsdk:"enable" json:"enable,omitempty"`
					LowLimitPercent   *int64 `tfsdk:"low_limit_percent" json:"lowLimitPercent,omitempty"`
					MinLimitPercent   *int64 `tfsdk:"min_limit_percent" json:"minLimitPercent,omitempty"`
					OomKillGroup      *int64 `tfsdk:"oom_kill_group" json:"oomKillGroup,omitempty"`
					Priority          *int64 `tfsdk:"priority" json:"priority,omitempty"`
					PriorityEnable    *int64 `tfsdk:"priority_enable" json:"priorityEnable,omitempty"`
					ThrottlingPercent *int64 `tfsdk:"throttling_percent" json:"throttlingPercent,omitempty"`
					WmarkMinAdj       *int64 `tfsdk:"wmark_min_adj" json:"wmarkMinAdj,omitempty"`
					WmarkRatio        *int64 `tfsdk:"wmark_ratio" json:"wmarkRatio,omitempty"`
					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" json:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" json:"memoryQOS,omitempty"`
				NetworkQOS *struct {
					EgressLimit    *string `tfsdk:"egress_limit" json:"egressLimit,omitempty"`
					EgressRequest  *string `tfsdk:"egress_request" json:"egressRequest,omitempty"`
					Enable         *bool   `tfsdk:"enable" json:"enable,omitempty"`
					IngressLimit   *string `tfsdk:"ingress_limit" json:"ingressLimit,omitempty"`
					IngressRequest *string `tfsdk:"ingress_request" json:"ingressRequest,omitempty"`
				} `tfsdk:"network_qos" json:"networkQOS,omitempty"`
				ResctrlQOS *struct {
					CatRangeEndPercent   *int64 `tfsdk:"cat_range_end_percent" json:"catRangeEndPercent,omitempty"`
					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" json:"catRangeStartPercent,omitempty"`
					Enable               *bool  `tfsdk:"enable" json:"enable,omitempty"`
					MbaPercent           *int64 `tfsdk:"mba_percent" json:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" json:"resctrlQOS,omitempty"`
			} `tfsdk:"lsr_class" json:"lsrClass,omitempty"`
			Policies *struct {
				CpuPolicy    *string `tfsdk:"cpu_policy" json:"cpuPolicy,omitempty"`
				NetQOSPolicy *string `tfsdk:"net_qos_policy" json:"netQOSPolicy,omitempty"`
			} `tfsdk:"policies" json:"policies,omitempty"`
			SystemClass *struct {
				BlkioQOS *struct {
					Blocks *[]struct {
						IoCfg *struct {
							EnableUserModel     *bool  `tfsdk:"enable_user_model" json:"enableUserModel,omitempty"`
							IoWeightPercent     *int64 `tfsdk:"io_weight_percent" json:"ioWeightPercent,omitempty"`
							ModelReadBPS        *int64 `tfsdk:"model_read_bps" json:"modelReadBPS,omitempty"`
							ModelReadRandIOPS   *int64 `tfsdk:"model_read_rand_iops" json:"modelReadRandIOPS,omitempty"`
							ModelReadSeqIOPS    *int64 `tfsdk:"model_read_seq_iops" json:"modelReadSeqIOPS,omitempty"`
							ModelWriteBPS       *int64 `tfsdk:"model_write_bps" json:"modelWriteBPS,omitempty"`
							ModelWriteRandIOPS  *int64 `tfsdk:"model_write_rand_iops" json:"modelWriteRandIOPS,omitempty"`
							ModelWriteSeqIOPS   *int64 `tfsdk:"model_write_seq_iops" json:"modelWriteSeqIOPS,omitempty"`
							ReadBPS             *int64 `tfsdk:"read_bps" json:"readBPS,omitempty"`
							ReadIOPS            *int64 `tfsdk:"read_iops" json:"readIOPS,omitempty"`
							ReadLatency         *int64 `tfsdk:"read_latency" json:"readLatency,omitempty"`
							ReadLatencyPercent  *int64 `tfsdk:"read_latency_percent" json:"readLatencyPercent,omitempty"`
							WriteBPS            *int64 `tfsdk:"write_bps" json:"writeBPS,omitempty"`
							WriteIOPS           *int64 `tfsdk:"write_iops" json:"writeIOPS,omitempty"`
							WriteLatency        *int64 `tfsdk:"write_latency" json:"writeLatency,omitempty"`
							WriteLatencyPercent *int64 `tfsdk:"write_latency_percent" json:"writeLatencyPercent,omitempty"`
						} `tfsdk:"io_cfg" json:"ioCfg,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"blocks" json:"blocks,omitempty"`
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"blkio_qos" json:"blkioQOS,omitempty"`
				CpuQOS *struct {
					CoreExpeller  *bool  `tfsdk:"core_expeller" json:"coreExpeller,omitempty"`
					Enable        *bool  `tfsdk:"enable" json:"enable,omitempty"`
					GroupIdentity *int64 `tfsdk:"group_identity" json:"groupIdentity,omitempty"`
					SchedIdle     *int64 `tfsdk:"sched_idle" json:"schedIdle,omitempty"`
				} `tfsdk:"cpu_qos" json:"cpuQOS,omitempty"`
				MemoryQOS *struct {
					Enable            *bool  `tfsdk:"enable" json:"enable,omitempty"`
					LowLimitPercent   *int64 `tfsdk:"low_limit_percent" json:"lowLimitPercent,omitempty"`
					MinLimitPercent   *int64 `tfsdk:"min_limit_percent" json:"minLimitPercent,omitempty"`
					OomKillGroup      *int64 `tfsdk:"oom_kill_group" json:"oomKillGroup,omitempty"`
					Priority          *int64 `tfsdk:"priority" json:"priority,omitempty"`
					PriorityEnable    *int64 `tfsdk:"priority_enable" json:"priorityEnable,omitempty"`
					ThrottlingPercent *int64 `tfsdk:"throttling_percent" json:"throttlingPercent,omitempty"`
					WmarkMinAdj       *int64 `tfsdk:"wmark_min_adj" json:"wmarkMinAdj,omitempty"`
					WmarkRatio        *int64 `tfsdk:"wmark_ratio" json:"wmarkRatio,omitempty"`
					WmarkScalePermill *int64 `tfsdk:"wmark_scale_permill" json:"wmarkScalePermill,omitempty"`
				} `tfsdk:"memory_qos" json:"memoryQOS,omitempty"`
				NetworkQOS *struct {
					EgressLimit    *string `tfsdk:"egress_limit" json:"egressLimit,omitempty"`
					EgressRequest  *string `tfsdk:"egress_request" json:"egressRequest,omitempty"`
					Enable         *bool   `tfsdk:"enable" json:"enable,omitempty"`
					IngressLimit   *string `tfsdk:"ingress_limit" json:"ingressLimit,omitempty"`
					IngressRequest *string `tfsdk:"ingress_request" json:"ingressRequest,omitempty"`
				} `tfsdk:"network_qos" json:"networkQOS,omitempty"`
				ResctrlQOS *struct {
					CatRangeEndPercent   *int64 `tfsdk:"cat_range_end_percent" json:"catRangeEndPercent,omitempty"`
					CatRangeStartPercent *int64 `tfsdk:"cat_range_start_percent" json:"catRangeStartPercent,omitempty"`
					Enable               *bool  `tfsdk:"enable" json:"enable,omitempty"`
					MbaPercent           *int64 `tfsdk:"mba_percent" json:"mbaPercent,omitempty"`
				} `tfsdk:"resctrl_qos" json:"resctrlQOS,omitempty"`
			} `tfsdk:"system_class" json:"systemClass,omitempty"`
		} `tfsdk:"resource_qos_strategy" json:"resourceQOSStrategy,omitempty"`
		ResourceUsedThresholdWithBE *struct {
			CpuEvictBESatisfactionLowerPercent *int64  `tfsdk:"cpu_evict_be_satisfaction_lower_percent" json:"cpuEvictBESatisfactionLowerPercent,omitempty"`
			CpuEvictBESatisfactionUpperPercent *int64  `tfsdk:"cpu_evict_be_satisfaction_upper_percent" json:"cpuEvictBESatisfactionUpperPercent,omitempty"`
			CpuEvictBEUsageThresholdPercent    *int64  `tfsdk:"cpu_evict_be_usage_threshold_percent" json:"cpuEvictBEUsageThresholdPercent,omitempty"`
			CpuEvictPolicy                     *string `tfsdk:"cpu_evict_policy" json:"cpuEvictPolicy,omitempty"`
			CpuEvictTimeWindowSeconds          *int64  `tfsdk:"cpu_evict_time_window_seconds" json:"cpuEvictTimeWindowSeconds,omitempty"`
			CpuSuppressPolicy                  *string `tfsdk:"cpu_suppress_policy" json:"cpuSuppressPolicy,omitempty"`
			CpuSuppressThresholdPercent        *int64  `tfsdk:"cpu_suppress_threshold_percent" json:"cpuSuppressThresholdPercent,omitempty"`
			Enable                             *bool   `tfsdk:"enable" json:"enable,omitempty"`
			MemoryEvictLowerPercent            *int64  `tfsdk:"memory_evict_lower_percent" json:"memoryEvictLowerPercent,omitempty"`
			MemoryEvictThresholdPercent        *int64  `tfsdk:"memory_evict_threshold_percent" json:"memoryEvictThresholdPercent,omitempty"`
		} `tfsdk:"resource_used_threshold_with_be" json:"resourceUsedThresholdWithBE,omitempty"`
		SystemStrategy *struct {
			MemcgReapBackGround   *int64  `tfsdk:"memcg_reap_back_ground" json:"memcgReapBackGround,omitempty"`
			MinFreeKbytesFactor   *int64  `tfsdk:"min_free_kbytes_factor" json:"minFreeKbytesFactor,omitempty"`
			TotalNetworkBandwidth *string `tfsdk:"total_network_bandwidth" json:"totalNetworkBandwidth,omitempty"`
			WatermarkScaleFactor  *int64  `tfsdk:"watermark_scale_factor" json:"watermarkScaleFactor,omitempty"`
		} `tfsdk:"system_strategy" json:"systemStrategy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SloKoordinatorShNodeSloV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_slo_koordinator_sh_node_slo_v1alpha1_manifest"
}

func (r *SloKoordinatorShNodeSloV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodeSLO is the Schema for the nodeslos API",
		MarkdownDescription: "NodeSLO is the Schema for the nodeslos API",
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
				Description:         "NodeSLOSpec defines the desired state of NodeSLO",
				MarkdownDescription: "NodeSLOSpec defines the desired state of NodeSLO",
				Attributes: map[string]schema.Attribute{
					"cpu_burst_strategy": schema.SingleNestedAttribute{
						Description:         "CPU Burst Strategy",
						MarkdownDescription: "CPU Burst Strategy",
						Attributes: map[string]schema.Attribute{
							"cfs_quota_burst_percent": schema.Int64Attribute{
								Description:         "pod cfs quota scale up ceil percentage, default = 300 (300%)",
								MarkdownDescription: "pod cfs quota scale up ceil percentage, default = 300 (300%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cfs_quota_burst_period_seconds": schema.Int64Attribute{
								Description:         "specifies a period of time for pod can use at burst, default = -1 (unlimited)",
								MarkdownDescription: "specifies a period of time for pod can use at burst, default = -1 (unlimited)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_burst_percent": schema.Int64Attribute{
								Description:         "cpu burst percentage for setting cpu.cfs_burst_us, legal range: [0, 10000], default as 1000 (1000%)",
								MarkdownDescription: "cpu burst percentage for setting cpu.cfs_burst_us, legal range: [0, 10000], default as 1000 (1000%)",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(10000),
								},
							},

							"policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"share_pool_threshold_percent": schema.Int64Attribute{
								Description:         "scale down cfs quota if node cpu overload, default = 50",
								MarkdownDescription: "scale down cfs quota if node cpu overload, default = 50",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"extensions": schema.MapAttribute{
						Description:         "Third party extensions for NodeSLO",
						MarkdownDescription: "Third party extensions for NodeSLO",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_applications": schema.ListNestedAttribute{
						Description:         "QoS management for out-of-band applications",
						MarkdownDescription: "QoS management for out-of-band applications",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"cgroup_path": schema.SingleNestedAttribute{
									Description:         "Optional, defines the host cgroup configuration, use default if not specified according to priority and qos",
									MarkdownDescription: "Optional, defines the host cgroup configuration, use default if not specified according to priority and qos",
									Attributes: map[string]schema.Attribute{
										"base": schema.StringAttribute{
											Description:         "cgroup base dir, the format is various across cgroup drivers",
											MarkdownDescription: "cgroup base dir, the format is various across cgroup drivers",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"parent_dir": schema.StringAttribute{
											Description:         "cgroup parent path under base dir",
											MarkdownDescription: "cgroup parent path under base dir",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"relative_path": schema.StringAttribute{
											Description:         "cgroup relative path under parent dir",
											MarkdownDescription: "cgroup relative path under parent dir",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"priority": schema.StringAttribute{
									Description:         "Priority class of the application",
									MarkdownDescription: "Priority class of the application",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"qos": schema.StringAttribute{
									Description:         "QoS class of the application",
									MarkdownDescription: "QoS class of the application",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"strategy": schema.MapAttribute{
									Description:         "QoS Strategy of host application",
									MarkdownDescription: "QoS Strategy of host application",
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

					"resource_qos_strategy": schema.SingleNestedAttribute{
						Description:         "QoS config strategy for pods of different qos-class",
						MarkdownDescription: "QoS config strategy for pods of different qos-class",
						Attributes: map[string]schema.Attribute{
							"be_class": schema.SingleNestedAttribute{
								Description:         "ResourceQOS for BE pods.",
								MarkdownDescription: "ResourceQOS for BE pods.",
								Attributes: map[string]schema.Attribute{
									"blkio_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"blocks": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"io_cfg": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"enable_user_model": schema.BoolAttribute{
																	Description:         "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	MarkdownDescription: "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"io_weight_percent": schema.Int64Attribute{
																	Description:         "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	MarkdownDescription: "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																		int64validator.AtMost(100),
																	},
																},

																"model_read_bps": schema.Int64Attribute{
																	Description:         "the read BPS of user model",
																	MarkdownDescription: "the read BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_rand_iops": schema.Int64Attribute{
																	Description:         "the random read iops of user model",
																	MarkdownDescription: "the random read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential read iops of user model",
																	MarkdownDescription: "the sequential read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_bps": schema.Int64Attribute{
																	Description:         "the write BPS of user model",
																	MarkdownDescription: "the write BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_rand_iops": schema.Int64Attribute{
																	Description:         "the random write iops of user model",
																	MarkdownDescription: "the random write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential write iops of user model",
																	MarkdownDescription: "the sequential write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"read_bps": schema.Int64Attribute{
																	Description:         "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_iops": schema.Int64Attribute{
																	Description:         "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_latency": schema.Int64Attribute{
																	Description:         "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	MarkdownDescription: "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_latency_percent": schema.Int64Attribute{
																	Description:         "the read latency percentile",
																	MarkdownDescription: "the read latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
																	},
																},

																"write_bps": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_iops": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_latency": schema.Int64Attribute{
																	Description:         "the write latency threshold. Unit: microseconds.",
																	MarkdownDescription: "the write latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"write_latency_percent": schema.Int64Attribute{
																	Description:         "the write latency percentile",
																	MarkdownDescription: "the write latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cpu_qos": schema.SingleNestedAttribute{
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",
										Attributes: map[string]schema.Attribute{
											"core_expeller": schema.BoolAttribute{
												Description:         "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"group_identity": schema.Int64Attribute{
												Description:         "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												MarkdownDescription: "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sched_idle": schema.Int64Attribute{
												Description:         "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": schema.SingleNestedAttribute{
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"low_limit_percent": schema.Int64Attribute{
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": schema.Int64Attribute{
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_enable": schema.Int64Attribute{
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"throttling_percent": schema.Int64Attribute{
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": schema.Int64Attribute{
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(-25),
													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": schema.Int64Attribute{
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": schema.Int64Attribute{
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(1000),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"egress_limit": schema.StringAttribute{
												Description:         "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"egress_request": schema.StringAttribute{
												Description:         "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_limit": schema.StringAttribute{
												Description:         "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_request": schema.StringAttribute{
												Description:         "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": schema.SingleNestedAttribute{
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",
										Attributes: map[string]schema.Attribute{
											"cat_range_end_percent": schema.Int64Attribute{
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": schema.Int64Attribute{
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mba_percent": schema.Int64Attribute{
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
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

							"cgroup_root": schema.SingleNestedAttribute{
								Description:         "ResourceQOS for root cgroup.",
								MarkdownDescription: "ResourceQOS for root cgroup.",
								Attributes: map[string]schema.Attribute{
									"blkio_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"blocks": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"io_cfg": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"enable_user_model": schema.BoolAttribute{
																	Description:         "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	MarkdownDescription: "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"io_weight_percent": schema.Int64Attribute{
																	Description:         "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	MarkdownDescription: "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																		int64validator.AtMost(100),
																	},
																},

																"model_read_bps": schema.Int64Attribute{
																	Description:         "the read BPS of user model",
																	MarkdownDescription: "the read BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_rand_iops": schema.Int64Attribute{
																	Description:         "the random read iops of user model",
																	MarkdownDescription: "the random read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential read iops of user model",
																	MarkdownDescription: "the sequential read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_bps": schema.Int64Attribute{
																	Description:         "the write BPS of user model",
																	MarkdownDescription: "the write BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_rand_iops": schema.Int64Attribute{
																	Description:         "the random write iops of user model",
																	MarkdownDescription: "the random write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential write iops of user model",
																	MarkdownDescription: "the sequential write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"read_bps": schema.Int64Attribute{
																	Description:         "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_iops": schema.Int64Attribute{
																	Description:         "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_latency": schema.Int64Attribute{
																	Description:         "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	MarkdownDescription: "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_latency_percent": schema.Int64Attribute{
																	Description:         "the read latency percentile",
																	MarkdownDescription: "the read latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
																	},
																},

																"write_bps": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_iops": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_latency": schema.Int64Attribute{
																	Description:         "the write latency threshold. Unit: microseconds.",
																	MarkdownDescription: "the write latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"write_latency_percent": schema.Int64Attribute{
																	Description:         "the write latency percentile",
																	MarkdownDescription: "the write latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cpu_qos": schema.SingleNestedAttribute{
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",
										Attributes: map[string]schema.Attribute{
											"core_expeller": schema.BoolAttribute{
												Description:         "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"group_identity": schema.Int64Attribute{
												Description:         "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												MarkdownDescription: "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sched_idle": schema.Int64Attribute{
												Description:         "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": schema.SingleNestedAttribute{
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"low_limit_percent": schema.Int64Attribute{
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": schema.Int64Attribute{
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_enable": schema.Int64Attribute{
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"throttling_percent": schema.Int64Attribute{
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": schema.Int64Attribute{
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(-25),
													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": schema.Int64Attribute{
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": schema.Int64Attribute{
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(1000),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"egress_limit": schema.StringAttribute{
												Description:         "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"egress_request": schema.StringAttribute{
												Description:         "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_limit": schema.StringAttribute{
												Description:         "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_request": schema.StringAttribute{
												Description:         "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": schema.SingleNestedAttribute{
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",
										Attributes: map[string]schema.Attribute{
											"cat_range_end_percent": schema.Int64Attribute{
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": schema.Int64Attribute{
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mba_percent": schema.Int64Attribute{
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
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

							"ls_class": schema.SingleNestedAttribute{
								Description:         "ResourceQOS for LS pods.",
								MarkdownDescription: "ResourceQOS for LS pods.",
								Attributes: map[string]schema.Attribute{
									"blkio_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"blocks": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"io_cfg": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"enable_user_model": schema.BoolAttribute{
																	Description:         "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	MarkdownDescription: "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"io_weight_percent": schema.Int64Attribute{
																	Description:         "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	MarkdownDescription: "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																		int64validator.AtMost(100),
																	},
																},

																"model_read_bps": schema.Int64Attribute{
																	Description:         "the read BPS of user model",
																	MarkdownDescription: "the read BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_rand_iops": schema.Int64Attribute{
																	Description:         "the random read iops of user model",
																	MarkdownDescription: "the random read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential read iops of user model",
																	MarkdownDescription: "the sequential read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_bps": schema.Int64Attribute{
																	Description:         "the write BPS of user model",
																	MarkdownDescription: "the write BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_rand_iops": schema.Int64Attribute{
																	Description:         "the random write iops of user model",
																	MarkdownDescription: "the random write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential write iops of user model",
																	MarkdownDescription: "the sequential write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"read_bps": schema.Int64Attribute{
																	Description:         "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_iops": schema.Int64Attribute{
																	Description:         "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_latency": schema.Int64Attribute{
																	Description:         "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	MarkdownDescription: "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_latency_percent": schema.Int64Attribute{
																	Description:         "the read latency percentile",
																	MarkdownDescription: "the read latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
																	},
																},

																"write_bps": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_iops": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_latency": schema.Int64Attribute{
																	Description:         "the write latency threshold. Unit: microseconds.",
																	MarkdownDescription: "the write latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"write_latency_percent": schema.Int64Attribute{
																	Description:         "the write latency percentile",
																	MarkdownDescription: "the write latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cpu_qos": schema.SingleNestedAttribute{
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",
										Attributes: map[string]schema.Attribute{
											"core_expeller": schema.BoolAttribute{
												Description:         "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"group_identity": schema.Int64Attribute{
												Description:         "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												MarkdownDescription: "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sched_idle": schema.Int64Attribute{
												Description:         "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": schema.SingleNestedAttribute{
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"low_limit_percent": schema.Int64Attribute{
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": schema.Int64Attribute{
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_enable": schema.Int64Attribute{
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"throttling_percent": schema.Int64Attribute{
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": schema.Int64Attribute{
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(-25),
													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": schema.Int64Attribute{
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": schema.Int64Attribute{
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(1000),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"egress_limit": schema.StringAttribute{
												Description:         "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"egress_request": schema.StringAttribute{
												Description:         "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_limit": schema.StringAttribute{
												Description:         "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_request": schema.StringAttribute{
												Description:         "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": schema.SingleNestedAttribute{
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",
										Attributes: map[string]schema.Attribute{
											"cat_range_end_percent": schema.Int64Attribute{
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": schema.Int64Attribute{
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mba_percent": schema.Int64Attribute{
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
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

							"lsr_class": schema.SingleNestedAttribute{
								Description:         "ResourceQOS for LSR pods.",
								MarkdownDescription: "ResourceQOS for LSR pods.",
								Attributes: map[string]schema.Attribute{
									"blkio_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"blocks": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"io_cfg": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"enable_user_model": schema.BoolAttribute{
																	Description:         "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	MarkdownDescription: "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"io_weight_percent": schema.Int64Attribute{
																	Description:         "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	MarkdownDescription: "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																		int64validator.AtMost(100),
																	},
																},

																"model_read_bps": schema.Int64Attribute{
																	Description:         "the read BPS of user model",
																	MarkdownDescription: "the read BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_rand_iops": schema.Int64Attribute{
																	Description:         "the random read iops of user model",
																	MarkdownDescription: "the random read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential read iops of user model",
																	MarkdownDescription: "the sequential read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_bps": schema.Int64Attribute{
																	Description:         "the write BPS of user model",
																	MarkdownDescription: "the write BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_rand_iops": schema.Int64Attribute{
																	Description:         "the random write iops of user model",
																	MarkdownDescription: "the random write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential write iops of user model",
																	MarkdownDescription: "the sequential write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"read_bps": schema.Int64Attribute{
																	Description:         "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_iops": schema.Int64Attribute{
																	Description:         "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_latency": schema.Int64Attribute{
																	Description:         "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	MarkdownDescription: "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_latency_percent": schema.Int64Attribute{
																	Description:         "the read latency percentile",
																	MarkdownDescription: "the read latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
																	},
																},

																"write_bps": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_iops": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_latency": schema.Int64Attribute{
																	Description:         "the write latency threshold. Unit: microseconds.",
																	MarkdownDescription: "the write latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"write_latency_percent": schema.Int64Attribute{
																	Description:         "the write latency percentile",
																	MarkdownDescription: "the write latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cpu_qos": schema.SingleNestedAttribute{
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",
										Attributes: map[string]schema.Attribute{
											"core_expeller": schema.BoolAttribute{
												Description:         "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"group_identity": schema.Int64Attribute{
												Description:         "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												MarkdownDescription: "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sched_idle": schema.Int64Attribute{
												Description:         "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": schema.SingleNestedAttribute{
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"low_limit_percent": schema.Int64Attribute{
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": schema.Int64Attribute{
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_enable": schema.Int64Attribute{
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"throttling_percent": schema.Int64Attribute{
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": schema.Int64Attribute{
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(-25),
													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": schema.Int64Attribute{
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": schema.Int64Attribute{
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(1000),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"egress_limit": schema.StringAttribute{
												Description:         "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"egress_request": schema.StringAttribute{
												Description:         "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_limit": schema.StringAttribute{
												Description:         "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_request": schema.StringAttribute{
												Description:         "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": schema.SingleNestedAttribute{
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",
										Attributes: map[string]schema.Attribute{
											"cat_range_end_percent": schema.Int64Attribute{
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": schema.Int64Attribute{
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mba_percent": schema.Int64Attribute{
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
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

							"policies": schema.SingleNestedAttribute{
								Description:         "Policies of pod QoS.",
								MarkdownDescription: "Policies of pod QoS.",
								Attributes: map[string]schema.Attribute{
									"cpu_policy": schema.StringAttribute{
										Description:         "applied policy for the CPU QoS, default = 'groupIdentity'",
										MarkdownDescription: "applied policy for the CPU QoS, default = 'groupIdentity'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"net_qos_policy": schema.StringAttribute{
										Description:         "applied policy for the Net QoS, default = 'tc'",
										MarkdownDescription: "applied policy for the Net QoS, default = 'tc'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"system_class": schema.SingleNestedAttribute{
								Description:         "ResourceQOS for system pods",
								MarkdownDescription: "ResourceQOS for system pods",
								Attributes: map[string]schema.Attribute{
									"blkio_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"blocks": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"io_cfg": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"enable_user_model": schema.BoolAttribute{
																	Description:         "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	MarkdownDescription: "configure the cost model of blkio-cost manually whether the user model is enabled. Default value: false",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"io_weight_percent": schema.Int64Attribute{
																	Description:         "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	MarkdownDescription: "This field is used to set the weight of a sub-group. Default value: 100. Valid values: 1 to 100.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																		int64validator.AtMost(100),
																	},
																},

																"model_read_bps": schema.Int64Attribute{
																	Description:         "the read BPS of user model",
																	MarkdownDescription: "the read BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_rand_iops": schema.Int64Attribute{
																	Description:         "the random read iops of user model",
																	MarkdownDescription: "the random read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_read_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential read iops of user model",
																	MarkdownDescription: "the sequential read iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_bps": schema.Int64Attribute{
																	Description:         "the write BPS of user model",
																	MarkdownDescription: "the write BPS of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_rand_iops": schema.Int64Attribute{
																	Description:         "the random write iops of user model",
																	MarkdownDescription: "the random write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"model_write_seq_iops": schema.Int64Attribute{
																	Description:         "the sequential write iops of user model",
																	MarkdownDescription: "the sequential write iops of user model",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(1),
																	},
																},

																"read_bps": schema.Int64Attribute{
																	Description:         "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of throughput The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_iops": schema.Int64Attribute{
																	Description:         "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	MarkdownDescription: "Throttling of IOPS The value is set to 0, which indicates that the feature is disabled.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"read_latency": schema.Int64Attribute{
																	Description:         "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	MarkdownDescription: "Configure the weight-based throttling feature of blk-iocost Only used for RootClass After blk-iocost is enabled, the kernel calculates the proportion of requests that exceed the read or write latency threshold out of all requests. When the proportion is greater than the read or write latency percentile (95%), the kernel considers the disk to be saturated and reduces the rate at which requests are sent to the disk. the read latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_latency_percent": schema.Int64Attribute{
																	Description:         "the read latency percentile",
																	MarkdownDescription: "the read latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
																	},
																},

																"write_bps": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_iops": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																	},
																},

																"write_latency": schema.Int64Attribute{
																	Description:         "the write latency threshold. Unit: microseconds.",
																	MarkdownDescription: "the write latency threshold. Unit: microseconds.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"write_latency_percent": schema.Int64Attribute{
																	Description:         "the write latency percentile",
																	MarkdownDescription: "the write latency percentile",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.Int64{
																		int64validator.AtLeast(0),
																		int64validator.AtMost(100),
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
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cpu_qos": schema.SingleNestedAttribute{
										Description:         "CPUQOSCfg stores node-level config of cpu qos",
										MarkdownDescription: "CPUQOSCfg stores node-level config of cpu qos",
										Attributes: map[string]schema.Attribute{
											"core_expeller": schema.BoolAttribute{
												Description:         "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "whether pods of the QoS class can expel the cgroup idle pods at the SMT-level. default = false If set to true, pods of this QoS will use a dedicated core sched group for noise clean with the SchedIdle pods. NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the cpu qos is enabled.",
												MarkdownDescription: "Enable indicates whether the cpu qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"group_identity": schema.Int64Attribute{
												Description:         "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												MarkdownDescription: "group identity value for pods, default = 0 NOTE: It takes effect if cpuPolicy = 'groupIdentity'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sched_idle": schema.Int64Attribute{
												Description:         "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												MarkdownDescription: "cpu.idle value for pods, default = 0. '1' means using SCHED_IDLE. CGroup Idle (introduced since mainline Linux 5.15): https://lore.kernel.org/lkml/162971078674.25758.15464079371945307825.tip-bot2@tip-bot2/#r NOTE: It takes effect if cpuPolicy = 'coreSched'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"memory_qos": schema.SingleNestedAttribute{
										Description:         "MemoryQOSCfg stores node-level config of memory qos",
										MarkdownDescription: "MemoryQOSCfg stores node-level config of memory qos",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												MarkdownDescription: "Enable indicates whether the memory qos is enabled (default: false). This field is used for node-level control, while pod-level configuration is done with MemoryQOS and 'Policy' instead of an 'Enable' option. Please view the differences between MemoryQOSCfg and PodMemoryQOSConfig structs.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"low_limit_percent": schema.Int64Attribute{
												Description:         "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												MarkdownDescription: "LowLimitPercent specifies the lowLimitFactor percentage to calculate 'memory.low', which TRIES BEST protecting memory from global reclamation when memory usage does not exceed the low limit unless no unprotected memcg can be reclaimed. NOTE: 'memory.low' should be larger than 'memory.min'. If spec.requests.memory == spec.limits.memory, pod 'memory.low' and 'memory.high' become invalid, while 'memory.wmark_ratio' is still in effect. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"min_limit_percent": schema.Int64Attribute{
												Description:         "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												MarkdownDescription: "memcg qos If enabled, memcg qos will be set by the agent, where some fields are implicitly calculated from pod spec. 1. 'memory.min' := spec.requests.memory * minLimitFactor / 100 (use 0 if requests.memory is not set) 2. 'memory.low' := spec.requests.memory * lowLimitFactor / 100 (use 0 if requests.memory is not set) 3. 'memory.limit_in_bytes' := spec.limits.memory (set $node.allocatable.memory if limits.memory is not set) 4. 'memory.high' := floor[(spec.requests.memory + throttlingFactor / 100 * (memory.limit_in_bytes or node allocatable memory - spec.requests.memory))/pageSize] * pageSize MinLimitPercent specifies the minLimitFactor percentage to calculate 'memory.min', which protects memory from global reclamation when memory usage does not exceed the min limit. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"oom_kill_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_enable": schema.Int64Attribute{
												Description:         "TODO: enhance the usages of oom priority and oom kill group",
												MarkdownDescription: "TODO: enhance the usages of oom priority and oom kill group",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"throttling_percent": schema.Int64Attribute{
												Description:         "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												MarkdownDescription: "ThrottlingPercent specifies the throttlingFactor percentage to calculate 'memory.high' with pod memory.limits or node allocatable memory, which triggers memcg direct reclamation when memory usage exceeds. Lower the factor brings more heavier reclaim pressure. Close: 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"wmark_min_adj": schema.Int64Attribute{
												Description:         "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												MarkdownDescription: "wmark_min_adj (Anolis OS required) WmarkMinAdj specifies 'memory.wmark_min_adj' which adjusts per-memcg threshold for global memory reclamation. Lower the factor brings later reclamation. The adjustment uses different formula for different value range. [-25, 0)：global_wmark_min' = global_wmark_min + (global_wmark_min - 0) * wmarkMinAdj (0, 50]：global_wmark_min' = global_wmark_min + (global_wmark_low - global_wmark_min) * wmarkMinAdj Close: [LSR:0, LS:0, BE:0]. Recommended: [LSR:-25, LS:-25, BE:50].",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(-25),
													int64validator.AtMost(50),
												},
											},

											"wmark_ratio": schema.Int64Attribute{
												Description:         "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												MarkdownDescription: "wmark_ratio (Anolis OS required) Async memory reclamation is triggered when cgroup memory usage exceeds 'memory.wmark_high' and the reclamation stops when usage is below 'memory.wmark_low'. Basically, 'memory.wmark_high' := min(memory.high, memory.limit_in_bytes) * memory.memory.wmark_ratio 'memory.wmark_low' := min(memory.high, memory.limit_in_bytes) * (memory.wmark_ratio - memory.wmark_scale_factor) WmarkRatio specifies 'memory.wmark_ratio' that help calculate 'memory.wmark_high', which triggers async memory reclamation when memory usage exceeds. Close: 0. Recommended: 95.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"wmark_scale_permill": schema.Int64Attribute{
												Description:         "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												MarkdownDescription: "WmarkScalePermill specifies 'memory.wmark_scale_factor' that helps calculate 'memory.wmark_low', which stops async memory reclamation when memory usage belows. Close: 50. Recommended: 20.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(1000),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"network_qos": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"egress_limit": schema.StringAttribute{
												Description:         "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressLimit describes the maximum network bandwidth can be used in the egress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"egress_request": schema.StringAttribute{
												Description:         "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "EgressRequest describes the minimum network bandwidth guaranteed in the egress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_limit": schema.StringAttribute{
												Description:         "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressLimit describes the maximum network bandwidth can be used in the ingress direction, unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_request": schema.StringAttribute{
												Description:         "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												MarkdownDescription: "IngressRequest describes the minimum network bandwidth guaranteed in the ingress direction. unit: bps(bytes per second), two expressions are supported，int and string, int: percentage based on total bandwidth，valid in 0-100 string: a specific network bandwidth value, eg: 50M.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resctrl_qos": schema.SingleNestedAttribute{
										Description:         "ResctrlQOSCfg stores node-level config of resctrl qos",
										MarkdownDescription: "ResctrlQOSCfg stores node-level config of resctrl qos",
										Attributes: map[string]schema.Attribute{
											"cat_range_end_percent": schema.Int64Attribute{
												Description:         "LLC available range end for pods by percentage",
												MarkdownDescription: "LLC available range end for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"cat_range_start_percent": schema.Int64Attribute{
												Description:         "LLC available range start for pods by percentage",
												MarkdownDescription: "LLC available range start for pods by percentage",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable indicates whether the resctrl qos is enabled.",
												MarkdownDescription: "Enable indicates whether the resctrl qos is enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mba_percent": schema.Int64Attribute{
												Description:         "MBA percent",
												MarkdownDescription: "MBA percent",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_used_threshold_with_be": schema.SingleNestedAttribute{
						Description:         "BE pods will be limited if node resource usage overload",
						MarkdownDescription: "BE pods will be limited if node resource usage overload",
						Attributes: map[string]schema.Attribute{
							"cpu_evict_be_satisfaction_lower_percent": schema.Int64Attribute{
								Description:         "be.satisfactionRate = be.CPURealLimit/be.CPURequest; be.cpuUsage = be.CPUUsed/be.CPURealLimit if be.satisfactionRate < CPUEvictBESatisfactionLowerPercent/100 && be.usage >= CPUEvictBEUsageThresholdPercent/100, then start to evict pod, and will evict to ${CPUEvictBESatisfactionUpperPercent}",
								MarkdownDescription: "be.satisfactionRate = be.CPURealLimit/be.CPURequest; be.cpuUsage = be.CPUUsed/be.CPURealLimit if be.satisfactionRate < CPUEvictBESatisfactionLowerPercent/100 && be.usage >= CPUEvictBEUsageThresholdPercent/100, then start to evict pod, and will evict to ${CPUEvictBESatisfactionUpperPercent}",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_evict_be_satisfaction_upper_percent": schema.Int64Attribute{
								Description:         "be.satisfactionRate = be.CPURealLimit/be.CPURequest if be.satisfactionRate > CPUEvictBESatisfactionUpperPercent/100, then stop to evict.",
								MarkdownDescription: "be.satisfactionRate = be.CPURealLimit/be.CPURequest if be.satisfactionRate > CPUEvictBESatisfactionUpperPercent/100, then stop to evict.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_evict_be_usage_threshold_percent": schema.Int64Attribute{
								Description:         "if be.cpuUsage >= CPUEvictBEUsageThresholdPercent/100, then start to calculate the resources need to be released.",
								MarkdownDescription: "if be.cpuUsage >= CPUEvictBEUsageThresholdPercent/100, then start to calculate the resources need to be released.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_evict_policy": schema.StringAttribute{
								Description:         "CPUEvictPolicy defines the policy for the BECPUEvict feature. Default: 'evictByRealLimit'.",
								MarkdownDescription: "CPUEvictPolicy defines the policy for the BECPUEvict feature. Default: 'evictByRealLimit'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_evict_time_window_seconds": schema.Int64Attribute{
								Description:         "when avg(cpuusage) > CPUEvictThresholdPercent, will start to evict pod by cpu, and avg(cpuusage) is calculated based on the most recent CPUEvictTimeWindowSeconds data",
								MarkdownDescription: "when avg(cpuusage) > CPUEvictThresholdPercent, will start to evict pod by cpu, and avg(cpuusage) is calculated based on the most recent CPUEvictTimeWindowSeconds data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_suppress_policy": schema.StringAttribute{
								Description:         "CPUSuppressPolicy",
								MarkdownDescription: "CPUSuppressPolicy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cpu_suppress_threshold_percent": schema.Int64Attribute{
								Description:         "cpu suppress threshold percentage (0,100), default = 65",
								MarkdownDescription: "cpu suppress threshold percentage (0,100), default = 65",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(100),
								},
							},

							"enable": schema.BoolAttribute{
								Description:         "whether the strategy is enabled, default = false",
								MarkdownDescription: "whether the strategy is enabled, default = false",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"memory_evict_lower_percent": schema.Int64Attribute{
								Description:         "lower: memory release util usage under MemoryEvictLowerPercent, default = MemoryEvictThresholdPercent - 2",
								MarkdownDescription: "lower: memory release util usage under MemoryEvictLowerPercent, default = MemoryEvictThresholdPercent - 2",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(100),
								},
							},

							"memory_evict_threshold_percent": schema.Int64Attribute{
								Description:         "upper: memory evict threshold percentage (0,100), default = 70",
								MarkdownDescription: "upper: memory evict threshold percentage (0,100), default = 70",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(100),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"system_strategy": schema.SingleNestedAttribute{
						Description:         "node global system config",
						MarkdownDescription: "node global system config",
						Attributes: map[string]schema.Attribute{
							"memcg_reap_back_ground": schema.Int64Attribute{
								Description:         "/sys/kernel/mm/memcg_reaper/reap_background",
								MarkdownDescription: "/sys/kernel/mm/memcg_reaper/reap_background",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_free_kbytes_factor": schema.Int64Attribute{
								Description:         "for /proc/sys/vm/min_free_kbytes, min_free_kbytes = minFreeKbytesFactor * nodeTotalMemory /10000",
								MarkdownDescription: "for /proc/sys/vm/min_free_kbytes, min_free_kbytes = minFreeKbytesFactor * nodeTotalMemory /10000",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"total_network_bandwidth": schema.StringAttribute{
								Description:         "TotalNetworkBandwidth indicates the overall network bandwidth, cluster manager can set this field, and default value taken from /sys/class/net/${NIC_NAME}/speed, unit: Mbps",
								MarkdownDescription: "TotalNetworkBandwidth indicates the overall network bandwidth, cluster manager can set this field, and default value taken from /sys/class/net/${NIC_NAME}/speed, unit: Mbps",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"watermark_scale_factor": schema.Int64Attribute{
								Description:         "/proc/sys/vm/watermark_scale_factor",
								MarkdownDescription: "/proc/sys/vm/watermark_scale_factor",
								Required:            false,
								Optional:            true,
								Computed:            false,
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

func (r *SloKoordinatorShNodeSloV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_slo_koordinator_sh_node_slo_v1alpha1_manifest")

	var model SloKoordinatorShNodeSloV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("slo.koordinator.sh/v1alpha1")
	model.Kind = pointer.String("NodeSLO")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
