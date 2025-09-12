/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CrdProjectcalicoOrgFelixConfigurationV1Manifest{}
)

func NewCrdProjectcalicoOrgFelixConfigurationV1Manifest() datasource.DataSource {
	return &CrdProjectcalicoOrgFelixConfigurationV1Manifest{}
}

type CrdProjectcalicoOrgFelixConfigurationV1Manifest struct{}

type CrdProjectcalicoOrgFelixConfigurationV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AllowIPIPPacketsFromWorkloads      *bool   `tfsdk:"allow_ipip_packets_from_workloads" json:"allowIPIPPacketsFromWorkloads,omitempty"`
		AllowVXLANPacketsFromWorkloads     *bool   `tfsdk:"allow_vxlan_packets_from_workloads" json:"allowVXLANPacketsFromWorkloads,omitempty"`
		AwsSrcDstCheck                     *string `tfsdk:"aws_src_dst_check" json:"awsSrcDstCheck,omitempty"`
		BpfAttachType                      *string `tfsdk:"bpf_attach_type" json:"bpfAttachType,omitempty"`
		BpfCTLBLogFilter                   *string `tfsdk:"bpf_ctlb_log_filter" json:"bpfCTLBLogFilter,omitempty"`
		BpfConnectTimeLoadBalancing        *string `tfsdk:"bpf_connect_time_load_balancing" json:"bpfConnectTimeLoadBalancing,omitempty"`
		BpfConnectTimeLoadBalancingEnabled *bool   `tfsdk:"bpf_connect_time_load_balancing_enabled" json:"bpfConnectTimeLoadBalancingEnabled,omitempty"`
		BpfConntrackLogLevel               *string `tfsdk:"bpf_conntrack_log_level" json:"bpfConntrackLogLevel,omitempty"`
		BpfConntrackMode                   *string `tfsdk:"bpf_conntrack_mode" json:"bpfConntrackMode,omitempty"`
		BpfConntrackTimeouts               *struct {
			CreationGracePeriod *string `tfsdk:"creation_grace_period" json:"creationGracePeriod,omitempty"`
			GenericTimeout      *string `tfsdk:"generic_timeout" json:"genericTimeout,omitempty"`
			IcmpTimeout         *string `tfsdk:"icmp_timeout" json:"icmpTimeout,omitempty"`
			TcpEstablished      *string `tfsdk:"tcp_established" json:"tcpEstablished,omitempty"`
			TcpFinsSeen         *string `tfsdk:"tcp_fins_seen" json:"tcpFinsSeen,omitempty"`
			TcpResetSeen        *string `tfsdk:"tcp_reset_seen" json:"tcpResetSeen,omitempty"`
			TcpSynSent          *string `tfsdk:"tcp_syn_sent" json:"tcpSynSent,omitempty"`
			UdpTimeout          *string `tfsdk:"udp_timeout" json:"udpTimeout,omitempty"`
		} `tfsdk:"bpf_conntrack_timeouts" json:"bpfConntrackTimeouts,omitempty"`
		BpfDSROptoutCIDRs                  *[]string          `tfsdk:"bpf_dsr_optout_cidrs" json:"bpfDSROptoutCIDRs,omitempty"`
		BpfDataIfacePattern                *string            `tfsdk:"bpf_data_iface_pattern" json:"bpfDataIfacePattern,omitempty"`
		BpfDisableGROForIfaces             *string            `tfsdk:"bpf_disable_gro_for_ifaces" json:"bpfDisableGROForIfaces,omitempty"`
		BpfDisableUnprivileged             *bool              `tfsdk:"bpf_disable_unprivileged" json:"bpfDisableUnprivileged,omitempty"`
		BpfEnabled                         *bool              `tfsdk:"bpf_enabled" json:"bpfEnabled,omitempty"`
		BpfEnforceRPF                      *string            `tfsdk:"bpf_enforce_rpf" json:"bpfEnforceRPF,omitempty"`
		BpfExcludeCIDRsFromNAT             *[]string          `tfsdk:"bpf_exclude_cidrs_from_nat" json:"bpfExcludeCIDRsFromNAT,omitempty"`
		BpfExportBufferSizeMB              *int64             `tfsdk:"bpf_export_buffer_size_mb" json:"bpfExportBufferSizeMB,omitempty"`
		BpfExtToServiceConnmark            *int64             `tfsdk:"bpf_ext_to_service_connmark" json:"bpfExtToServiceConnmark,omitempty"`
		BpfExternalServiceMode             *string            `tfsdk:"bpf_external_service_mode" json:"bpfExternalServiceMode,omitempty"`
		BpfForceTrackPacketsFromIfaces     *[]string          `tfsdk:"bpf_force_track_packets_from_ifaces" json:"bpfForceTrackPacketsFromIfaces,omitempty"`
		BpfHostConntrackBypass             *bool              `tfsdk:"bpf_host_conntrack_bypass" json:"bpfHostConntrackBypass,omitempty"`
		BpfHostNetworkedNATWithoutCTLB     *string            `tfsdk:"bpf_host_networked_nat_without_ctlb" json:"bpfHostNetworkedNATWithoutCTLB,omitempty"`
		BpfKubeProxyEndpointSlicesEnabled  *bool              `tfsdk:"bpf_kube_proxy_endpoint_slices_enabled" json:"bpfKubeProxyEndpointSlicesEnabled,omitempty"`
		BpfKubeProxyIptablesCleanupEnabled *bool              `tfsdk:"bpf_kube_proxy_iptables_cleanup_enabled" json:"bpfKubeProxyIptablesCleanupEnabled,omitempty"`
		BpfKubeProxyMinSyncPeriod          *string            `tfsdk:"bpf_kube_proxy_min_sync_period" json:"bpfKubeProxyMinSyncPeriod,omitempty"`
		BpfL3IfacePattern                  *string            `tfsdk:"bpf_l3_iface_pattern" json:"bpfL3IfacePattern,omitempty"`
		BpfLogFilters                      *map[string]string `tfsdk:"bpf_log_filters" json:"bpfLogFilters,omitempty"`
		BpfLogLevel                        *string            `tfsdk:"bpf_log_level" json:"bpfLogLevel,omitempty"`
		BpfMapSizeConntrack                *int64             `tfsdk:"bpf_map_size_conntrack" json:"bpfMapSizeConntrack,omitempty"`
		BpfMapSizeConntrackCleanupQueue    *int64             `tfsdk:"bpf_map_size_conntrack_cleanup_queue" json:"bpfMapSizeConntrackCleanupQueue,omitempty"`
		BpfMapSizeConntrackScaling         *string            `tfsdk:"bpf_map_size_conntrack_scaling" json:"bpfMapSizeConntrackScaling,omitempty"`
		BpfMapSizeIPSets                   *int64             `tfsdk:"bpf_map_size_ip_sets" json:"bpfMapSizeIPSets,omitempty"`
		BpfMapSizeIfState                  *int64             `tfsdk:"bpf_map_size_if_state" json:"bpfMapSizeIfState,omitempty"`
		BpfMapSizeNATAffinity              *int64             `tfsdk:"bpf_map_size_nat_affinity" json:"bpfMapSizeNATAffinity,omitempty"`
		BpfMapSizeNATBackend               *int64             `tfsdk:"bpf_map_size_nat_backend" json:"bpfMapSizeNATBackend,omitempty"`
		BpfMapSizeNATFrontend              *int64             `tfsdk:"bpf_map_size_nat_frontend" json:"bpfMapSizeNATFrontend,omitempty"`
		BpfMapSizePerCpuConntrack          *int64             `tfsdk:"bpf_map_size_per_cpu_conntrack" json:"bpfMapSizePerCpuConntrack,omitempty"`
		BpfMapSizeRoute                    *int64             `tfsdk:"bpf_map_size_route" json:"bpfMapSizeRoute,omitempty"`
		BpfPSNATPorts                      *string            `tfsdk:"bpf_psnat_ports" json:"bpfPSNATPorts,omitempty"`
		BpfPolicyDebugEnabled              *bool              `tfsdk:"bpf_policy_debug_enabled" json:"bpfPolicyDebugEnabled,omitempty"`
		BpfProfiling                       *string            `tfsdk:"bpf_profiling" json:"bpfProfiling,omitempty"`
		BpfRedirectToPeer                  *string            `tfsdk:"bpf_redirect_to_peer" json:"bpfRedirectToPeer,omitempty"`
		CgroupV2Path                       *string            `tfsdk:"cgroup_v2_path" json:"cgroupV2Path,omitempty"`
		ChainInsertMode                    *string            `tfsdk:"chain_insert_mode" json:"chainInsertMode,omitempty"`
		DataplaneDriver                    *string            `tfsdk:"dataplane_driver" json:"dataplaneDriver,omitempty"`
		DataplaneWatchdogTimeout           *string            `tfsdk:"dataplane_watchdog_timeout" json:"dataplaneWatchdogTimeout,omitempty"`
		DebugDisableLogDropping            *bool              `tfsdk:"debug_disable_log_dropping" json:"debugDisableLogDropping,omitempty"`
		DebugHost                          *string            `tfsdk:"debug_host" json:"debugHost,omitempty"`
		DebugMemoryProfilePath             *string            `tfsdk:"debug_memory_profile_path" json:"debugMemoryProfilePath,omitempty"`
		DebugPort                          *int64             `tfsdk:"debug_port" json:"debugPort,omitempty"`
		DebugSimulateCalcGraphHangAfter    *string            `tfsdk:"debug_simulate_calc_graph_hang_after" json:"debugSimulateCalcGraphHangAfter,omitempty"`
		DebugSimulateDataplaneApplyDelay   *string            `tfsdk:"debug_simulate_dataplane_apply_delay" json:"debugSimulateDataplaneApplyDelay,omitempty"`
		DebugSimulateDataplaneHangAfter    *string            `tfsdk:"debug_simulate_dataplane_hang_after" json:"debugSimulateDataplaneHangAfter,omitempty"`
		DefaultEndpointToHostAction        *string            `tfsdk:"default_endpoint_to_host_action" json:"defaultEndpointToHostAction,omitempty"`
		DeviceRouteProtocol                *int64             `tfsdk:"device_route_protocol" json:"deviceRouteProtocol,omitempty"`
		DeviceRouteSourceAddress           *string            `tfsdk:"device_route_source_address" json:"deviceRouteSourceAddress,omitempty"`
		DeviceRouteSourceAddressIPv6       *string            `tfsdk:"device_route_source_address_i_pv6" json:"deviceRouteSourceAddressIPv6,omitempty"`
		DisableConntrackInvalidCheck       *bool              `tfsdk:"disable_conntrack_invalid_check" json:"disableConntrackInvalidCheck,omitempty"`
		EndpointReportingDelay             *string            `tfsdk:"endpoint_reporting_delay" json:"endpointReportingDelay,omitempty"`
		EndpointReportingEnabled           *bool              `tfsdk:"endpoint_reporting_enabled" json:"endpointReportingEnabled,omitempty"`
		EndpointStatusPathPrefix           *string            `tfsdk:"endpoint_status_path_prefix" json:"endpointStatusPathPrefix,omitempty"`
		ExternalNodesList                  *[]string          `tfsdk:"external_nodes_list" json:"externalNodesList,omitempty"`
		FailsafeInboundHostPorts           *[]struct {
			Net      *string `tfsdk:"net" json:"net,omitempty"`
			Port     *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"failsafe_inbound_host_ports" json:"failsafeInboundHostPorts,omitempty"`
		FailsafeOutboundHostPorts *[]struct {
			Net      *string `tfsdk:"net" json:"net,omitempty"`
			Port     *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"failsafe_outbound_host_ports" json:"failsafeOutboundHostPorts,omitempty"`
		FeatureDetectOverride        *string `tfsdk:"feature_detect_override" json:"featureDetectOverride,omitempty"`
		FeatureGates                 *string `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		FloatingIPs                  *string `tfsdk:"floating_i_ps" json:"floatingIPs,omitempty"`
		FlowLogsCollectorDebugTrace  *bool   `tfsdk:"flow_logs_collector_debug_trace" json:"flowLogsCollectorDebugTrace,omitempty"`
		FlowLogsFlushInterval        *string `tfsdk:"flow_logs_flush_interval" json:"flowLogsFlushInterval,omitempty"`
		FlowLogsGoldmaneServer       *string `tfsdk:"flow_logs_goldmane_server" json:"flowLogsGoldmaneServer,omitempty"`
		FlowLogsLocalReporter        *string `tfsdk:"flow_logs_local_reporter" json:"flowLogsLocalReporter,omitempty"`
		FlowLogsPolicyEvaluationMode *string `tfsdk:"flow_logs_policy_evaluation_mode" json:"flowLogsPolicyEvaluationMode,omitempty"`
		GenericXDPEnabled            *bool   `tfsdk:"generic_xdp_enabled" json:"genericXDPEnabled,omitempty"`
		GoGCThreshold                *int64  `tfsdk:"go_gc_threshold" json:"goGCThreshold,omitempty"`
		GoMaxProcs                   *int64  `tfsdk:"go_max_procs" json:"goMaxProcs,omitempty"`
		GoMemoryLimitMB              *int64  `tfsdk:"go_memory_limit_mb" json:"goMemoryLimitMB,omitempty"`
		HealthEnabled                *bool   `tfsdk:"health_enabled" json:"healthEnabled,omitempty"`
		HealthHost                   *string `tfsdk:"health_host" json:"healthHost,omitempty"`
		HealthPort                   *int64  `tfsdk:"health_port" json:"healthPort,omitempty"`
		HealthTimeoutOverrides       *[]struct {
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"health_timeout_overrides" json:"healthTimeoutOverrides,omitempty"`
		InterfaceExclude                   *string   `tfsdk:"interface_exclude" json:"interfaceExclude,omitempty"`
		InterfacePrefix                    *string   `tfsdk:"interface_prefix" json:"interfacePrefix,omitempty"`
		InterfaceRefreshInterval           *string   `tfsdk:"interface_refresh_interval" json:"interfaceRefreshInterval,omitempty"`
		IpForwarding                       *string   `tfsdk:"ip_forwarding" json:"ipForwarding,omitempty"`
		IpipEnabled                        *bool     `tfsdk:"ipip_enabled" json:"ipipEnabled,omitempty"`
		IpipMTU                            *int64    `tfsdk:"ipip_mtu" json:"ipipMTU,omitempty"`
		IpsetsRefreshInterval              *string   `tfsdk:"ipsets_refresh_interval" json:"ipsetsRefreshInterval,omitempty"`
		IptablesBackend                    *string   `tfsdk:"iptables_backend" json:"iptablesBackend,omitempty"`
		IptablesFilterAllowAction          *string   `tfsdk:"iptables_filter_allow_action" json:"iptablesFilterAllowAction,omitempty"`
		IptablesFilterDenyAction           *string   `tfsdk:"iptables_filter_deny_action" json:"iptablesFilterDenyAction,omitempty"`
		IptablesLockFilePath               *string   `tfsdk:"iptables_lock_file_path" json:"iptablesLockFilePath,omitempty"`
		IptablesLockProbeInterval          *string   `tfsdk:"iptables_lock_probe_interval" json:"iptablesLockProbeInterval,omitempty"`
		IptablesLockTimeout                *string   `tfsdk:"iptables_lock_timeout" json:"iptablesLockTimeout,omitempty"`
		IptablesMangleAllowAction          *string   `tfsdk:"iptables_mangle_allow_action" json:"iptablesMangleAllowAction,omitempty"`
		IptablesMarkMask                   *int64    `tfsdk:"iptables_mark_mask" json:"iptablesMarkMask,omitempty"`
		IptablesNATOutgoingInterfaceFilter *string   `tfsdk:"iptables_nat_outgoing_interface_filter" json:"iptablesNATOutgoingInterfaceFilter,omitempty"`
		IptablesPostWriteCheckInterval     *string   `tfsdk:"iptables_post_write_check_interval" json:"iptablesPostWriteCheckInterval,omitempty"`
		IptablesRefreshInterval            *string   `tfsdk:"iptables_refresh_interval" json:"iptablesRefreshInterval,omitempty"`
		Ipv6Support                        *bool     `tfsdk:"ipv6_support" json:"ipv6Support,omitempty"`
		KubeNodePortRanges                 *[]string `tfsdk:"kube_node_port_ranges" json:"kubeNodePortRanges,omitempty"`
		LogDebugFilenameRegex              *string   `tfsdk:"log_debug_filename_regex" json:"logDebugFilenameRegex,omitempty"`
		LogFilePath                        *string   `tfsdk:"log_file_path" json:"logFilePath,omitempty"`
		LogPrefix                          *string   `tfsdk:"log_prefix" json:"logPrefix,omitempty"`
		LogSeverityFile                    *string   `tfsdk:"log_severity_file" json:"logSeverityFile,omitempty"`
		LogSeverityScreen                  *string   `tfsdk:"log_severity_screen" json:"logSeverityScreen,omitempty"`
		LogSeveritySys                     *string   `tfsdk:"log_severity_sys" json:"logSeveritySys,omitempty"`
		MaxIpsetSize                       *int64    `tfsdk:"max_ipset_size" json:"maxIpsetSize,omitempty"`
		MetadataAddr                       *string   `tfsdk:"metadata_addr" json:"metadataAddr,omitempty"`
		MetadataPort                       *int64    `tfsdk:"metadata_port" json:"metadataPort,omitempty"`
		MtuIfacePattern                    *string   `tfsdk:"mtu_iface_pattern" json:"mtuIfacePattern,omitempty"`
		NatOutgoingAddress                 *string   `tfsdk:"nat_outgoing_address" json:"natOutgoingAddress,omitempty"`
		NatOutgoingExclusions              *string   `tfsdk:"nat_outgoing_exclusions" json:"natOutgoingExclusions,omitempty"`
		NatPortRange                       *string   `tfsdk:"nat_port_range" json:"natPortRange,omitempty"`
		NetlinkTimeout                     *string   `tfsdk:"netlink_timeout" json:"netlinkTimeout,omitempty"`
		NftablesFilterAllowAction          *string   `tfsdk:"nftables_filter_allow_action" json:"nftablesFilterAllowAction,omitempty"`
		NftablesFilterDenyAction           *string   `tfsdk:"nftables_filter_deny_action" json:"nftablesFilterDenyAction,omitempty"`
		NftablesMangleAllowAction          *string   `tfsdk:"nftables_mangle_allow_action" json:"nftablesMangleAllowAction,omitempty"`
		NftablesMarkMask                   *int64    `tfsdk:"nftables_mark_mask" json:"nftablesMarkMask,omitempty"`
		NftablesMode                       *string   `tfsdk:"nftables_mode" json:"nftablesMode,omitempty"`
		NftablesRefreshInterval            *string   `tfsdk:"nftables_refresh_interval" json:"nftablesRefreshInterval,omitempty"`
		OpenstackRegion                    *string   `tfsdk:"openstack_region" json:"openstackRegion,omitempty"`
		PolicySyncPathPrefix               *string   `tfsdk:"policy_sync_path_prefix" json:"policySyncPathPrefix,omitempty"`
		ProgramClusterRoutes               *string   `tfsdk:"program_cluster_routes" json:"programClusterRoutes,omitempty"`
		PrometheusGoMetricsEnabled         *bool     `tfsdk:"prometheus_go_metrics_enabled" json:"prometheusGoMetricsEnabled,omitempty"`
		PrometheusMetricsEnabled           *bool     `tfsdk:"prometheus_metrics_enabled" json:"prometheusMetricsEnabled,omitempty"`
		PrometheusMetricsHost              *string   `tfsdk:"prometheus_metrics_host" json:"prometheusMetricsHost,omitempty"`
		PrometheusMetricsPort              *int64    `tfsdk:"prometheus_metrics_port" json:"prometheusMetricsPort,omitempty"`
		PrometheusProcessMetricsEnabled    *bool     `tfsdk:"prometheus_process_metrics_enabled" json:"prometheusProcessMetricsEnabled,omitempty"`
		PrometheusWireGuardMetricsEnabled  *bool     `tfsdk:"prometheus_wire_guard_metrics_enabled" json:"prometheusWireGuardMetricsEnabled,omitempty"`
		RemoveExternalRoutes               *bool     `tfsdk:"remove_external_routes" json:"removeExternalRoutes,omitempty"`
		ReportingInterval                  *string   `tfsdk:"reporting_interval" json:"reportingInterval,omitempty"`
		ReportingTTL                       *string   `tfsdk:"reporting_ttl" json:"reportingTTL,omitempty"`
		RequireMTUFile                     *bool     `tfsdk:"require_mtu_file" json:"requireMTUFile,omitempty"`
		RouteRefreshInterval               *string   `tfsdk:"route_refresh_interval" json:"routeRefreshInterval,omitempty"`
		RouteSource                        *string   `tfsdk:"route_source" json:"routeSource,omitempty"`
		RouteSyncDisabled                  *bool     `tfsdk:"route_sync_disabled" json:"routeSyncDisabled,omitempty"`
		RouteTableRange                    *struct {
			Max *int64 `tfsdk:"max" json:"max,omitempty"`
			Min *int64 `tfsdk:"min" json:"min,omitempty"`
		} `tfsdk:"route_table_range" json:"routeTableRange,omitempty"`
		RouteTableRanges *[]struct {
			Max *int64 `tfsdk:"max" json:"max,omitempty"`
			Min *int64 `tfsdk:"min" json:"min,omitempty"`
		} `tfsdk:"route_table_ranges" json:"routeTableRanges,omitempty"`
		ServiceLoopPrevention          *string `tfsdk:"service_loop_prevention" json:"serviceLoopPrevention,omitempty"`
		SidecarAccelerationEnabled     *bool   `tfsdk:"sidecar_acceleration_enabled" json:"sidecarAccelerationEnabled,omitempty"`
		UsageReportingEnabled          *bool   `tfsdk:"usage_reporting_enabled" json:"usageReportingEnabled,omitempty"`
		UsageReportingInitialDelay     *string `tfsdk:"usage_reporting_initial_delay" json:"usageReportingInitialDelay,omitempty"`
		UsageReportingInterval         *string `tfsdk:"usage_reporting_interval" json:"usageReportingInterval,omitempty"`
		UseInternalDataplaneDriver     *bool   `tfsdk:"use_internal_dataplane_driver" json:"useInternalDataplaneDriver,omitempty"`
		VxlanEnabled                   *bool   `tfsdk:"vxlan_enabled" json:"vxlanEnabled,omitempty"`
		VxlanMTU                       *int64  `tfsdk:"vxlan_mtu" json:"vxlanMTU,omitempty"`
		VxlanMTUV6                     *int64  `tfsdk:"vxlan_mtuv6" json:"vxlanMTUV6,omitempty"`
		VxlanPort                      *int64  `tfsdk:"vxlan_port" json:"vxlanPort,omitempty"`
		VxlanVNI                       *int64  `tfsdk:"vxlan_vni" json:"vxlanVNI,omitempty"`
		WindowsManageFirewallRules     *string `tfsdk:"windows_manage_firewall_rules" json:"windowsManageFirewallRules,omitempty"`
		WireguardEnabled               *bool   `tfsdk:"wireguard_enabled" json:"wireguardEnabled,omitempty"`
		WireguardEnabledV6             *bool   `tfsdk:"wireguard_enabled_v6" json:"wireguardEnabledV6,omitempty"`
		WireguardHostEncryptionEnabled *bool   `tfsdk:"wireguard_host_encryption_enabled" json:"wireguardHostEncryptionEnabled,omitempty"`
		WireguardInterfaceName         *string `tfsdk:"wireguard_interface_name" json:"wireguardInterfaceName,omitempty"`
		WireguardInterfaceNameV6       *string `tfsdk:"wireguard_interface_name_v6" json:"wireguardInterfaceNameV6,omitempty"`
		WireguardKeepAlive             *string `tfsdk:"wireguard_keep_alive" json:"wireguardKeepAlive,omitempty"`
		WireguardListeningPort         *int64  `tfsdk:"wireguard_listening_port" json:"wireguardListeningPort,omitempty"`
		WireguardListeningPortV6       *int64  `tfsdk:"wireguard_listening_port_v6" json:"wireguardListeningPortV6,omitempty"`
		WireguardMTU                   *int64  `tfsdk:"wireguard_mtu" json:"wireguardMTU,omitempty"`
		WireguardMTUV6                 *int64  `tfsdk:"wireguard_mtuv6" json:"wireguardMTUV6,omitempty"`
		WireguardRoutingRulePriority   *int64  `tfsdk:"wireguard_routing_rule_priority" json:"wireguardRoutingRulePriority,omitempty"`
		WireguardThreadingEnabled      *bool   `tfsdk:"wireguard_threading_enabled" json:"wireguardThreadingEnabled,omitempty"`
		WorkloadSourceSpoofing         *string `tfsdk:"workload_source_spoofing" json:"workloadSourceSpoofing,omitempty"`
		XdpEnabled                     *bool   `tfsdk:"xdp_enabled" json:"xdpEnabled,omitempty"`
		XdpRefreshInterval             *string `tfsdk:"xdp_refresh_interval" json:"xdpRefreshInterval,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_felix_configuration_v1_manifest"
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Felix Configuration contains the configuration for Felix.",
		MarkdownDescription: "Felix Configuration contains the configuration for Felix.",
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
				Description:         "FelixConfigurationSpec contains the values of the Felix configuration.",
				MarkdownDescription: "FelixConfigurationSpec contains the values of the Felix configuration.",
				Attributes: map[string]schema.Attribute{
					"allow_ipip_packets_from_workloads": schema.BoolAttribute{
						Description:         "AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic from workloads. [Default: false]",
						MarkdownDescription: "AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic from workloads. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_vxlan_packets_from_workloads": schema.BoolAttribute{
						Description:         "AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic from workloads. [Default: false]",
						MarkdownDescription: "AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic from workloads. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"aws_src_dst_check": schema.StringAttribute{
						Description:         "AWSSrcDstCheck controls whether Felix will try to change the 'source/dest check' setting on the EC2 instance on which it is running. A value of 'Disable' will try to disable the source/dest check. Disabling the check allows for sending workload traffic without encapsulation within the same AWS subnet. [Default: DoNothing]",
						MarkdownDescription: "AWSSrcDstCheck controls whether Felix will try to change the 'source/dest check' setting on the EC2 instance on which it is running. A value of 'Disable' will try to disable the source/dest check. Disabling the check allows for sending workload traffic without encapsulation within the same AWS subnet. [Default: DoNothing]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("DoNothing", "Enable", "Disable"),
						},
					},

					"bpf_attach_type": schema.StringAttribute{
						Description:         "BPFAttachType controls how are the BPF programs at the network interfaces attached. By default 'TCX' is used where available to enable easier coexistence with 3rd party programs. 'TC' can force the legacy method of attaching via a qdisc. 'TCX' falls back to 'TC' if 'TCX' is not available. [Default: TCX]",
						MarkdownDescription: "BPFAttachType controls how are the BPF programs at the network interfaces attached. By default 'TCX' is used where available to enable easier coexistence with 3rd party programs. 'TC' can force the legacy method of attaching via a qdisc. 'TCX' falls back to 'TC' if 'TCX' is not available. [Default: TCX]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("TC", "TCX"),
						},
					},

					"bpf_ctlb_log_filter": schema.StringAttribute{
						Description:         "BPFCTLBLogFilter specifies, what is logged by connect time load balancer when BPFLogLevel is debug. Currently has to be specified as 'all' when BPFLogFilters is set to see CTLB logs. [Default: unset - means logs are emitted when BPFLogLevel id debug and BPFLogFilters not set.]",
						MarkdownDescription: "BPFCTLBLogFilter specifies, what is logged by connect time load balancer when BPFLogLevel is debug. Currently has to be specified as 'all' when BPFLogFilters is set to see CTLB logs. [Default: unset - means logs are emitted when BPFLogLevel id debug and BPFLogFilters not set.]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_connect_time_load_balancing": schema.StringAttribute{
						Description:         "BPFConnectTimeLoadBalancing when in BPF mode, controls whether Felix installs the connect-time load balancer. The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections.When set to TCP, connect time load balancing is available only for services with TCP ports. [Default: TCP]",
						MarkdownDescription: "BPFConnectTimeLoadBalancing when in BPF mode, controls whether Felix installs the connect-time load balancer. The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections.When set to TCP, connect time load balancing is available only for services with TCP ports. [Default: TCP]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("TCP", "Enabled", "Disabled"),
						},
					},

					"bpf_connect_time_load_balancing_enabled": schema.BoolAttribute{
						Description:         "BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load balancer. The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections. The only reason to disable it is for debugging purposes. Deprecated: Use BPFConnectTimeLoadBalancing [Default: true]",
						MarkdownDescription: "BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load balancer. The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections. The only reason to disable it is for debugging purposes. Deprecated: Use BPFConnectTimeLoadBalancing [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_conntrack_log_level": schema.StringAttribute{
						Description:         "BPFConntrackLogLevel controls the log level of the BPF conntrack cleanup program, which runs periodically to clean up expired BPF conntrack entries. [Default: Off].",
						MarkdownDescription: "BPFConntrackLogLevel controls the log level of the BPF conntrack cleanup program, which runs periodically to clean up expired BPF conntrack entries. [Default: Off].",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Off", "Debug"),
						},
					},

					"bpf_conntrack_mode": schema.StringAttribute{
						Description:         "BPFConntrackCleanupMode controls how BPF conntrack entries are cleaned up. 'Auto' will use a BPF program if supported, falling back to userspace if not. 'Userspace' will always use the userspace cleanup code. 'BPFProgram' will always use the BPF program (failing if not supported). /To be deprecated in future versions as conntrack map type changed to lru_hash and userspace cleanup is the only mode that is supported. [Default: Userspace]",
						MarkdownDescription: "BPFConntrackCleanupMode controls how BPF conntrack entries are cleaned up. 'Auto' will use a BPF program if supported, falling back to userspace if not. 'Userspace' will always use the userspace cleanup code. 'BPFProgram' will always use the BPF program (failing if not supported). /To be deprecated in future versions as conntrack map type changed to lru_hash and userspace cleanup is the only mode that is supported. [Default: Userspace]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Auto", "Userspace", "BPFProgram"),
						},
					},

					"bpf_conntrack_timeouts": schema.SingleNestedAttribute{
						Description:         "BPFConntrackTimers overrides the default values for the specified conntrack timer if set. Each value can be either a duration or 'Auto' to pick the value from a Linux conntrack timeout. Configurable timers are: CreationGracePeriod, TCPSynSent, TCPEstablished, TCPFinsSeen, TCPResetSeen, UDPTimeout, GenericTimeout, ICMPTimeout. Unset values are replaced by the default values with a warning log for incorrect values.",
						MarkdownDescription: "BPFConntrackTimers overrides the default values for the specified conntrack timer if set. Each value can be either a duration or 'Auto' to pick the value from a Linux conntrack timeout. Configurable timers are: CreationGracePeriod, TCPSynSent, TCPEstablished, TCPFinsSeen, TCPResetSeen, UDPTimeout, GenericTimeout, ICMPTimeout. Unset values are replaced by the default values with a warning log for incorrect values.",
						Attributes: map[string]schema.Attribute{
							"creation_grace_period": schema.StringAttribute{
								Description:         "CreationGracePeriod gives a generic grace period to new connections before they are considered for cleanup [Default: 10s].",
								MarkdownDescription: "CreationGracePeriod gives a generic grace period to new connections before they are considered for cleanup [Default: 10s].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},

							"generic_timeout": schema.StringAttribute{
								Description:         "GenericTimeout controls how long it takes before considering this entry for cleanup after the connection became idle. If set to 'Auto', the value from nf_conntrack_generic_timeout is used. If nil, Calico uses its own default value. [Default: 10m].",
								MarkdownDescription: "GenericTimeout controls how long it takes before considering this entry for cleanup after the connection became idle. If set to 'Auto', the value from nf_conntrack_generic_timeout is used. If nil, Calico uses its own default value. [Default: 10m].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},

							"icmp_timeout": schema.StringAttribute{
								Description:         "ICMPTimeout controls how long it takes before considering this entry for cleanup after the connection became idle. If set to 'Auto', the value from nf_conntrack_icmp_timeout is used. If nil, Calico uses its own default value. [Default: 5s].",
								MarkdownDescription: "ICMPTimeout controls how long it takes before considering this entry for cleanup after the connection became idle. If set to 'Auto', the value from nf_conntrack_icmp_timeout is used. If nil, Calico uses its own default value. [Default: 5s].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},

							"tcp_established": schema.StringAttribute{
								Description:         "TCPEstablished controls how long it takes before considering this entry for cleanup after the connection became idle. If set to 'Auto', the value from nf_conntrack_tcp_timeout_established is used. If nil, Calico uses its own default value. [Default: 1h].",
								MarkdownDescription: "TCPEstablished controls how long it takes before considering this entry for cleanup after the connection became idle. If set to 'Auto', the value from nf_conntrack_tcp_timeout_established is used. If nil, Calico uses its own default value. [Default: 1h].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},

							"tcp_fins_seen": schema.StringAttribute{
								Description:         "TCPFinsSeen controls how long it takes before considering this entry for cleanup after the connection was closed gracefully. If set to 'Auto', the value from nf_conntrack_tcp_timeout_time_wait is used. If nil, Calico uses its own default value. [Default: Auto].",
								MarkdownDescription: "TCPFinsSeen controls how long it takes before considering this entry for cleanup after the connection was closed gracefully. If set to 'Auto', the value from nf_conntrack_tcp_timeout_time_wait is used. If nil, Calico uses its own default value. [Default: Auto].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},

							"tcp_reset_seen": schema.StringAttribute{
								Description:         "TCPResetSeen controls how long it takes before considering this entry for cleanup after the connection was aborted. If nil, Calico uses its own default value. [Default: 40s].",
								MarkdownDescription: "TCPResetSeen controls how long it takes before considering this entry for cleanup after the connection was aborted. If nil, Calico uses its own default value. [Default: 40s].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},

							"tcp_syn_sent": schema.StringAttribute{
								Description:         "TCPSynSent controls how long it takes before considering this entry for cleanup after the last SYN without a response. If set to 'Auto', the value from nf_conntrack_tcp_timeout_syn_sent is used. If nil, Calico uses its own default value. [Default: 20s].",
								MarkdownDescription: "TCPSynSent controls how long it takes before considering this entry for cleanup after the last SYN without a response. If set to 'Auto', the value from nf_conntrack_tcp_timeout_syn_sent is used. If nil, Calico uses its own default value. [Default: 20s].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},

							"udp_timeout": schema.StringAttribute{
								Description:         "UDPTimeout controls how long it takes before considering this entry for cleanup after the connection became idle. If nil, Calico uses its own default value. [Default: 60s].",
								MarkdownDescription: "UDPTimeout controls how long it takes before considering this entry for cleanup after the connection became idle. If nil, Calico uses its own default value. [Default: 60s].",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]*(\.[0-9]*)?(ms|s|h|m|us)+)+|Auto)$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_dsr_optout_cidrs": schema.ListAttribute{
						Description:         "BPFDSROptoutCIDRs is a list of CIDRs which are excluded from DSR. That is, clients in those CIDRs will access service node ports as if BPFExternalServiceMode was set to Tunnel.",
						MarkdownDescription: "BPFDSROptoutCIDRs is a list of CIDRs which are excluded from DSR. That is, clients in those CIDRs will access service node ports as if BPFExternalServiceMode was set to Tunnel.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_data_iface_pattern": schema.StringAttribute{
						Description:         "BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to in order to catch traffic to/from the network. This needs to match the interfaces that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster. It should not match the workload interfaces (usually named cali...) or any other special device managed by Calico itself (e.g., tunnels).",
						MarkdownDescription: "BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to in order to catch traffic to/from the network. This needs to match the interfaces that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster. It should not match the workload interfaces (usually named cali...) or any other special device managed by Calico itself (e.g., tunnels).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_disable_gro_for_ifaces": schema.StringAttribute{
						Description:         "BPFDisableGROForIfaces is a regular expression that controls which interfaces Felix should disable the Generic Receive Offload [GRO] option. It should not match the workload interfaces (usually named cali...).",
						MarkdownDescription: "BPFDisableGROForIfaces is a regular expression that controls which interfaces Felix should disable the Generic Receive Offload [GRO] option. It should not match the workload interfaces (usually named cali...).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_disable_unprivileged": schema.BoolAttribute{
						Description:         "BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable unprivileged use of BPF. This ensures that unprivileged users cannot access Calico's BPF maps and cannot insert their own BPF programs to interfere with Calico's. [Default: true]",
						MarkdownDescription: "BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable unprivileged use of BPF. This ensures that unprivileged users cannot access Calico's BPF maps and cannot insert their own BPF programs to interfere with Calico's. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_enabled": schema.BoolAttribute{
						Description:         "BPFEnabled, if enabled Felix will use the BPF dataplane. [Default: false]",
						MarkdownDescription: "BPFEnabled, if enabled Felix will use the BPF dataplane. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_enforce_rpf": schema.StringAttribute{
						Description:         "BPFEnforceRPF enforce strict RPF on all host interfaces with BPF programs regardless of what is the per-interfaces or global setting. Possible values are Disabled, Strict or Loose. [Default: Loose]",
						MarkdownDescription: "BPFEnforceRPF enforce strict RPF on all host interfaces with BPF programs regardless of what is the per-interfaces or global setting. Possible values are Disabled, Strict or Loose. [Default: Loose]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Disabled|Strict|Loose)?$`), ""),
						},
					},

					"bpf_exclude_cidrs_from_nat": schema.ListAttribute{
						Description:         "BPFExcludeCIDRsFromNAT is a list of CIDRs that are to be excluded from NAT resolution so that host can handle them. A typical usecase is node local DNS cache.",
						MarkdownDescription: "BPFExcludeCIDRsFromNAT is a list of CIDRs that are to be excluded from NAT resolution so that host can handle them. A typical usecase is node local DNS cache.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_export_buffer_size_mb": schema.Int64Attribute{
						Description:         "BPFExportBufferSizeMB in BPF mode, controls the buffer size used for sending BPF events to felix. [Default: 1]",
						MarkdownDescription: "BPFExportBufferSizeMB in BPF mode, controls the buffer size used for sending BPF events to felix. [Default: 1]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_ext_to_service_connmark": schema.Int64Attribute{
						Description:         "BPFExtToServiceConnmark in BPF mode, controls a 32bit mark that is set on connections from an external client to a local service. This mark allows us to control how packets of that connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]",
						MarkdownDescription: "BPFExtToServiceConnmark in BPF mode, controls a 32bit mark that is set on connections from an external client to a local service. This mark allows us to control how packets of that connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_external_service_mode": schema.StringAttribute{
						Description:         "BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports and cluster IPs) are forwarded to remote workloads. If set to 'Tunnel' then both request and response traffic is tunneled to the remote node. If set to 'DSR', the request traffic is tunneled but the response traffic is sent directly from the remote node. In 'DSR' mode, the remote node appears to use the IP of the ingress node; this requires a permissive L2 network. [Default: Tunnel]",
						MarkdownDescription: "BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports and cluster IPs) are forwarded to remote workloads. If set to 'Tunnel' then both request and response traffic is tunneled to the remote node. If set to 'DSR', the request traffic is tunneled but the response traffic is sent directly from the remote node. In 'DSR' mode, the remote node appears to use the IP of the ingress node; this requires a permissive L2 network. [Default: Tunnel]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Tunnel|DSR)?$`), ""),
						},
					},

					"bpf_force_track_packets_from_ifaces": schema.ListAttribute{
						Description:         "BPFForceTrackPacketsFromIfaces in BPF mode, forces traffic from these interfaces to skip Calico's iptables NOTRACK rule, allowing traffic from those interfaces to be tracked by Linux conntrack. Should only be used for interfaces that are not used for the Calico fabric. For example, a docker bridge device for non-Calico-networked containers. [Default: docker+]",
						MarkdownDescription: "BPFForceTrackPacketsFromIfaces in BPF mode, forces traffic from these interfaces to skip Calico's iptables NOTRACK rule, allowing traffic from those interfaces to be tracked by Linux conntrack. Should only be used for interfaces that are not used for the Calico fabric. For example, a docker bridge device for non-Calico-networked containers. [Default: docker+]",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_host_conntrack_bypass": schema.BoolAttribute{
						Description:         "BPFHostConntrackBypass Controls whether to bypass Linux conntrack in BPF mode for workloads and services. [Default: true - bypass Linux conntrack]",
						MarkdownDescription: "BPFHostConntrackBypass Controls whether to bypass Linux conntrack in BPF mode for workloads and services. [Default: true - bypass Linux conntrack]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_host_networked_nat_without_ctlb": schema.StringAttribute{
						Description:         "BPFHostNetworkedNATWithoutCTLB when in BPF mode, controls whether Felix does a NAT without CTLB. This along with BPFConnectTimeLoadBalancing determines the CTLB behavior. [Default: Enabled]",
						MarkdownDescription: "BPFHostNetworkedNATWithoutCTLB when in BPF mode, controls whether Felix does a NAT without CTLB. This along with BPFConnectTimeLoadBalancing determines the CTLB behavior. [Default: Enabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"bpf_kube_proxy_endpoint_slices_enabled": schema.BoolAttribute{
						Description:         "BPFKubeProxyEndpointSlicesEnabled is deprecated and has no effect. BPF kube-proxy always accepts endpoint slices. This option will be removed in the next release.",
						MarkdownDescription: "BPFKubeProxyEndpointSlicesEnabled is deprecated and has no effect. BPF kube-proxy always accepts endpoint slices. This option will be removed in the next release.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_kube_proxy_iptables_cleanup_enabled": schema.BoolAttribute{
						Description:         "BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream Kubernetes kube-proxy's iptables chains. Should only be enabled if kube-proxy is not running. [Default: true]",
						MarkdownDescription: "BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream Kubernetes kube-proxy's iptables chains. Should only be enabled if kube-proxy is not running. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_kube_proxy_min_sync_period": schema.StringAttribute{
						Description:         "BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's embedded kube-proxy. Lower values give reduced set-up latency. Higher values reduce Felix CPU usage by batching up more work. [Default: 1s]",
						MarkdownDescription: "BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's embedded kube-proxy. Lower values give reduced set-up latency. Higher values reduce Felix CPU usage by batching up more work. [Default: 1s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"bpf_l3_iface_pattern": schema.StringAttribute{
						Description:         "BPFL3IfacePattern is a regular expression that allows to list tunnel devices like wireguard or vxlan (i.e., L3 devices) in addition to BPFDataIfacePattern. That is, tunnel interfaces not created by Calico, that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.",
						MarkdownDescription: "BPFL3IfacePattern is a regular expression that allows to list tunnel devices like wireguard or vxlan (i.e., L3 devices) in addition to BPFDataIfacePattern. That is, tunnel interfaces not created by Calico, that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_log_filters": schema.MapAttribute{
						Description:         "BPFLogFilters is a map of key=values where the value is a pcap filter expression and the key is an interface name with 'all' denoting all interfaces, 'weps' all workload endpoints and 'heps' all host endpoints. When specified as an env var, it accepts a comma-separated list of key=values. [Default: unset - means all debug logs are emitted]",
						MarkdownDescription: "BPFLogFilters is a map of key=values where the value is a pcap filter expression and the key is an interface name with 'all' denoting all interfaces, 'weps' all workload endpoints and 'heps' all host endpoints. When specified as an env var, it accepts a comma-separated list of key=values. [Default: unset - means all debug logs are emitted]",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_log_level": schema.StringAttribute{
						Description:         "BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode. One of 'Off', 'Info', or 'Debug'. The logs are emitted to the BPF trace pipe, accessible with the command 'tc exec bpf debug'. [Default: Off].",
						MarkdownDescription: "BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode. One of 'Off', 'Info', or 'Debug'. The logs are emitted to the BPF trace pipe, accessible with the command 'tc exec bpf debug'. [Default: Off].",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Off|Info|Debug)?$`), ""),
						},
					},

					"bpf_map_size_conntrack": schema.Int64Attribute{
						Description:         "BPFMapSizeConntrack sets the size for the conntrack map. This map must be large enough to hold an entry for each active connection. Warning: changing the size of the conntrack map can cause disruption.",
						MarkdownDescription: "BPFMapSizeConntrack sets the size for the conntrack map. This map must be large enough to hold an entry for each active connection. Warning: changing the size of the conntrack map can cause disruption.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_conntrack_cleanup_queue": schema.Int64Attribute{
						Description:         "BPFMapSizeConntrackCleanupQueue sets the size for the map used to hold NAT conntrack entries that are queued for cleanup. This should be big enough to hold all the NAT entries that expire within one cleanup interval.",
						MarkdownDescription: "BPFMapSizeConntrackCleanupQueue sets the size for the map used to hold NAT conntrack entries that are queued for cleanup. This should be big enough to hold all the NAT entries that expire within one cleanup interval.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"bpf_map_size_conntrack_scaling": schema.StringAttribute{
						Description:         "BPFMapSizeConntrackScaling controls whether and how we scale the conntrack map size depending on its usage. 'Disabled' make the size stay at the default or whatever is set by BPFMapSizeConntrack*. 'DoubleIfFull' doubles the size when the map is pretty much full even after cleanups. [Default: DoubleIfFull]",
						MarkdownDescription: "BPFMapSizeConntrackScaling controls whether and how we scale the conntrack map size depending on its usage. 'Disabled' make the size stay at the default or whatever is set by BPFMapSizeConntrack*. 'DoubleIfFull' doubles the size when the map is pretty much full even after cleanups. [Default: DoubleIfFull]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Disabled|DoubleIfFull)?$`), ""),
						},
					},

					"bpf_map_size_ip_sets": schema.Int64Attribute{
						Description:         "BPFMapSizeIPSets sets the size for ipsets map. The IP sets map must be large enough to hold an entry for each endpoint matched by every selector in the source/destination matches in network policy. Selectors such as 'all()' can result in large numbers of entries (one entry per endpoint in that case).",
						MarkdownDescription: "BPFMapSizeIPSets sets the size for ipsets map. The IP sets map must be large enough to hold an entry for each endpoint matched by every selector in the source/destination matches in network policy. Selectors such as 'all()' can result in large numbers of entries (one entry per endpoint in that case).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_if_state": schema.Int64Attribute{
						Description:         "BPFMapSizeIfState sets the size for ifstate map. The ifstate map must be large enough to hold an entry for each device (host + workloads) on a host.",
						MarkdownDescription: "BPFMapSizeIfState sets the size for ifstate map. The ifstate map must be large enough to hold an entry for each device (host + workloads) on a host.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_nat_affinity": schema.Int64Attribute{
						Description:         "BPFMapSizeNATAffinity sets the size of the BPF map that stores the affinity of a connection (for services that enable that feature.",
						MarkdownDescription: "BPFMapSizeNATAffinity sets the size of the BPF map that stores the affinity of a connection (for services that enable that feature.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_nat_backend": schema.Int64Attribute{
						Description:         "BPFMapSizeNATBackend sets the size for NAT back end map. This is the total number of endpoints. This is mostly more than the size of the number of services.",
						MarkdownDescription: "BPFMapSizeNATBackend sets the size for NAT back end map. This is the total number of endpoints. This is mostly more than the size of the number of services.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_nat_frontend": schema.Int64Attribute{
						Description:         "BPFMapSizeNATFrontend sets the size for NAT front end map. FrontendMap should be large enough to hold an entry for each nodeport, external IP and each port in each service.",
						MarkdownDescription: "BPFMapSizeNATFrontend sets the size for NAT front end map. FrontendMap should be large enough to hold an entry for each nodeport, external IP and each port in each service.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_per_cpu_conntrack": schema.Int64Attribute{
						Description:         "BPFMapSizePerCPUConntrack determines the size of conntrack map based on the number of CPUs. If set to a non-zero value, overrides BPFMapSizeConntrack with 'BPFMapSizePerCPUConntrack * (Number of CPUs)'. This map must be large enough to hold an entry for each active connection. Warning: changing the size of the conntrack map can cause disruption.",
						MarkdownDescription: "BPFMapSizePerCPUConntrack determines the size of conntrack map based on the number of CPUs. If set to a non-zero value, overrides BPFMapSizeConntrack with 'BPFMapSizePerCPUConntrack * (Number of CPUs)'. This map must be large enough to hold an entry for each active connection. Warning: changing the size of the conntrack map can cause disruption.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_route": schema.Int64Attribute{
						Description:         "BPFMapSizeRoute sets the size for the routes map. The routes map should be large enough to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and tunnel IPs).",
						MarkdownDescription: "BPFMapSizeRoute sets the size for the routes map. The routes map should be large enough to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and tunnel IPs).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_psnat_ports": schema.StringAttribute{
						Description:         "BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision. This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]",
						MarkdownDescription: "BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision. This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_policy_debug_enabled": schema.BoolAttribute{
						Description:         "BPFPolicyDebugEnabled when true, Felix records detailed information about the BPF policy programs, which can be examined with the calico-bpf command-line tool.",
						MarkdownDescription: "BPFPolicyDebugEnabled when true, Felix records detailed information about the BPF policy programs, which can be examined with the calico-bpf command-line tool.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_profiling": schema.StringAttribute{
						Description:         "BPFProfiling controls profiling of BPF programs. At the monent, it can be Disabled or Enabled. [Default: Disabled]",
						MarkdownDescription: "BPFProfiling controls profiling of BPF programs. At the monent, it can be Disabled or Enabled. [Default: Disabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"bpf_redirect_to_peer": schema.StringAttribute{
						Description:         "BPFRedirectToPeer controls which whether it is allowed to forward straight to the peer side of the workload devices. It is allowed for any host L2 devices by default (L2Only), but it breaks TCP dump on the host side of workload device as it bypasses it on ingress. Value of Enabled also allows redirection from L3 host devices like IPIP tunnel or Wireguard directly to the peer side of the workload's device. This makes redirection faster, however, it breaks tools like tcpdump on the peer side. Use Enabled with caution. [Default: L2Only]",
						MarkdownDescription: "BPFRedirectToPeer controls which whether it is allowed to forward straight to the peer side of the workload devices. It is allowed for any host L2 devices by default (L2Only), but it breaks TCP dump on the host side of workload device as it bypasses it on ingress. Value of Enabled also allows redirection from L3 host devices like IPIP tunnel or Wireguard directly to the peer side of the workload's device. This makes redirection faster, however, it breaks tools like tcpdump on the peer side. Use Enabled with caution. [Default: L2Only]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled", "L2Only"),
						},
					},

					"cgroup_v2_path": schema.StringAttribute{
						Description:         "CgroupV2Path overrides the default location where to find the cgroup hierarchy.",
						MarkdownDescription: "CgroupV2Path overrides the default location where to find the cgroup hierarchy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"chain_insert_mode": schema.StringAttribute{
						Description:         "ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule at the top of the chain or by appending a rule at the bottom. insert is the safe default since it prevents Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed. [Default: insert]",
						MarkdownDescription: "ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule at the top of the chain or by appending a rule at the bottom. insert is the safe default since it prevents Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed. [Default: insert]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Insert|Append)?$`), ""),
						},
					},

					"dataplane_driver": schema.StringAttribute{
						Description:         "DataplaneDriver filename of the external dataplane driver to use. Only used if UseInternalDataplaneDriver is set to false.",
						MarkdownDescription: "DataplaneDriver filename of the external dataplane driver to use. Only used if UseInternalDataplaneDriver is set to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dataplane_watchdog_timeout": schema.StringAttribute{
						Description:         "DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver. Deprecated: replaced by the generic HealthTimeoutOverrides.",
						MarkdownDescription: "DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver. Deprecated: replaced by the generic HealthTimeoutOverrides.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_disable_log_dropping": schema.BoolAttribute{
						Description:         "DebugDisableLogDropping disables the dropping of log messages when the log buffer is full. This can significantly impact performance if log write-out is a bottleneck. [Default: false]",
						MarkdownDescription: "DebugDisableLogDropping disables the dropping of log messages when the log buffer is full. This can significantly impact performance if log write-out is a bottleneck. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_host": schema.StringAttribute{
						Description:         "DebugHost is the host IP or hostname to bind the debug port to. Only used if DebugPort is set. [Default:localhost]",
						MarkdownDescription: "DebugHost is the host IP or hostname to bind the debug port to. Only used if DebugPort is set. [Default:localhost]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_memory_profile_path": schema.StringAttribute{
						Description:         "DebugMemoryProfilePath is the path to write the memory profile to when triggered by signal.",
						MarkdownDescription: "DebugMemoryProfilePath is the path to write the memory profile to when triggered by signal.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_port": schema.Int64Attribute{
						Description:         "DebugPort if set, enables Felix's debug HTTP port, which allows memory and CPU profiles to be retrieved. The debug port is not secure, it should not be exposed to the internet.",
						MarkdownDescription: "DebugPort if set, enables Felix's debug HTTP port, which allows memory and CPU profiles to be retrieved. The debug port is not secure, it should not be exposed to the internet.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_simulate_calc_graph_hang_after": schema.StringAttribute{
						Description:         "DebugSimulateCalcGraphHangAfter is used to simulate a hang in the calculation graph after the specified duration. This is useful in tests of the watchdog system only!",
						MarkdownDescription: "DebugSimulateCalcGraphHangAfter is used to simulate a hang in the calculation graph after the specified duration. This is useful in tests of the watchdog system only!",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"debug_simulate_dataplane_apply_delay": schema.StringAttribute{
						Description:         "DebugSimulateDataplaneApplyDelay adds an artificial delay to every dataplane operation. This is useful for simulating a heavily loaded system for test purposes only.",
						MarkdownDescription: "DebugSimulateDataplaneApplyDelay adds an artificial delay to every dataplane operation. This is useful for simulating a heavily loaded system for test purposes only.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"debug_simulate_dataplane_hang_after": schema.StringAttribute{
						Description:         "DebugSimulateDataplaneHangAfter is used to simulate a hang in the dataplane after the specified duration. This is useful in tests of the watchdog system only!",
						MarkdownDescription: "DebugSimulateDataplaneHangAfter is used to simulate a hang in the dataplane after the specified duration. This is useful in tests of the watchdog system only!",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"default_endpoint_to_host_action": schema.StringAttribute{
						Description:         "DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host itself (after the endpoint's egress policy is applied). By default, Calico blocks traffic from workload endpoints to the host itself with an iptables 'DROP' action. If you want to allow some or all traffic from endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables 'INPUT' chain; Calico will insert its rules at the top of that chain, then 'RETURN' packets to the 'INPUT' chain once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets from workloads after processing workload endpoint egress policy. [Default: Drop]",
						MarkdownDescription: "DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host itself (after the endpoint's egress policy is applied). By default, Calico blocks traffic from workload endpoints to the host itself with an iptables 'DROP' action. If you want to allow some or all traffic from endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables 'INPUT' chain; Calico will insert its rules at the top of that chain, then 'RETURN' packets to the 'INPUT' chain once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets from workloads after processing workload endpoint egress policy. [Default: Drop]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Drop|Accept|Return)?$`), ""),
						},
					},

					"device_route_protocol": schema.Int64Attribute{
						Description:         "DeviceRouteProtocol controls the protocol to set on routes programmed by Felix. The protocol is an 8-bit label used to identify the owner of the route.",
						MarkdownDescription: "DeviceRouteProtocol controls the protocol to set on routes programmed by Felix. The protocol is an 8-bit label used to identify the owner of the route.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"device_route_source_address": schema.StringAttribute{
						Description:         "DeviceRouteSourceAddress IPv4 address to set as the source hint for routes programmed by Felix. When not set the source address for local traffic from host to workload will be determined by the kernel.",
						MarkdownDescription: "DeviceRouteSourceAddress IPv4 address to set as the source hint for routes programmed by Felix. When not set the source address for local traffic from host to workload will be determined by the kernel.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"device_route_source_address_i_pv6": schema.StringAttribute{
						Description:         "DeviceRouteSourceAddressIPv6 IPv6 address to set as the source hint for routes programmed by Felix. When not set the source address for local traffic from host to workload will be determined by the kernel.",
						MarkdownDescription: "DeviceRouteSourceAddressIPv6 IPv6 address to set as the source hint for routes programmed by Felix. When not set the source address for local traffic from host to workload will be determined by the kernel.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_conntrack_invalid_check": schema.BoolAttribute{
						Description:         "DisableConntrackInvalidCheck disables the check for invalid connections in conntrack. While the conntrack invalid check helps to detect malicious traffic, it can also cause issues with certain multi-NIC scenarios.",
						MarkdownDescription: "DisableConntrackInvalidCheck disables the check for invalid connections in conntrack. While the conntrack invalid check helps to detect malicious traffic, it can also cause issues with certain multi-NIC scenarios.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoint_reporting_delay": schema.StringAttribute{
						Description:         "EndpointReportingDelay is the delay before Felix reports endpoint status to the datastore. This is only used by the OpenStack integration. [Default: 1s]",
						MarkdownDescription: "EndpointReportingDelay is the delay before Felix reports endpoint status to the datastore. This is only used by the OpenStack integration. [Default: 1s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"endpoint_reporting_enabled": schema.BoolAttribute{
						Description:         "EndpointReportingEnabled controls whether Felix reports endpoint status to the datastore. This is only used by the OpenStack integration. [Default: false]",
						MarkdownDescription: "EndpointReportingEnabled controls whether Felix reports endpoint status to the datastore. This is only used by the OpenStack integration. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoint_status_path_prefix": schema.StringAttribute{
						Description:         "EndpointStatusPathPrefix is the path to the directory where endpoint status will be written. Endpoint status file reporting is disabled if field is left empty. Chosen directory should match the directory used by the CNI plugin for PodStartupDelay. [Default: /var/run/calico]",
						MarkdownDescription: "EndpointStatusPathPrefix is the path to the directory where endpoint status will be written. Endpoint status file reporting is disabled if field is left empty. Chosen directory should match the directory used by the CNI plugin for PodStartupDelay. [Default: /var/run/calico]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"external_nodes_list": schema.ListAttribute{
						Description:         "ExternalNodesCIDRList is a list of CIDR's of external, non-Calico nodes from which VXLAN/IPIP overlay traffic will be allowed. By default, external tunneled traffic is blocked to reduce attack surface.",
						MarkdownDescription: "ExternalNodesCIDRList is a list of CIDR's of external, non-Calico nodes from which VXLAN/IPIP overlay traffic will be allowed. By default, external tunneled traffic is blocked to reduce attack surface.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failsafe_inbound_host_ports": schema.ListNestedAttribute{
						Description:         "FailsafeInboundHostPorts is a list of ProtoPort struct objects including UDP/TCP/SCTP ports and CIDRs that Felix will allow incoming traffic to host endpoints on irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For backwards compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports, use the value '[]'. The default value allows ssh access, DHCP, BGP, etcd and the Kubernetes API. [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:5473, tcp:6443, tcp:6666, tcp:6667 ]",
						MarkdownDescription: "FailsafeInboundHostPorts is a list of ProtoPort struct objects including UDP/TCP/SCTP ports and CIDRs that Felix will allow incoming traffic to host endpoints on irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For backwards compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports, use the value '[]'. The default value allows ssh access, DHCP, BGP, etcd and the Kubernetes API. [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:5473, tcp:6443, tcp:6666, tcp:6667 ]",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"net": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
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

					"failsafe_outbound_host_ports": schema.ListNestedAttribute{
						Description:         "FailsafeOutboundHostPorts is a list of PortProto struct objects including UDP/TCP/SCTP ports and CIDRs that Felix will allow outgoing traffic from host endpoints to irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For backwards compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports, use the value '[]'. The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd as well as allowing DHCP, DNS, BGP and the Kubernetes API. [Default: udp:53, udp:67, tcp:179, tcp:2379, tcp:2380, tcp:5473, tcp:6443, tcp:6666, tcp:6667 ]",
						MarkdownDescription: "FailsafeOutboundHostPorts is a list of PortProto struct objects including UDP/TCP/SCTP ports and CIDRs that Felix will allow outgoing traffic from host endpoints to irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For backwards compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports, use the value '[]'. The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd as well as allowing DHCP, DNS, BGP and the Kubernetes API. [Default: udp:53, udp:67, tcp:179, tcp:2379, tcp:2380, tcp:5473, tcp:6443, tcp:6666, tcp:6667 ]",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"net": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
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

					"feature_detect_override": schema.StringAttribute{
						Description:         "FeatureDetectOverride is used to override feature detection based on auto-detected platform capabilities. Values are specified in a comma separated list with no spaces, example; 'SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock='. A value of 'true' or 'false' will force enable/disable feature, empty or omitted values fall back to auto-detection.",
						MarkdownDescription: "FeatureDetectOverride is used to override feature detection based on auto-detected platform capabilities. Values are specified in a comma separated list with no spaces, example; 'SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock='. A value of 'true' or 'false' will force enable/disable feature, empty or omitted values fall back to auto-detection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9-_]+=(true|false|),)*([a-zA-Z0-9-_]+=(true|false|))?$`), ""),
						},
					},

					"feature_gates": schema.StringAttribute{
						Description:         "FeatureGates is used to enable or disable tech-preview Calico features. Values are specified in a comma separated list with no spaces, example; 'BPFConnectTimeLoadBalancingWorkaround=enabled,XyZ=false'. This is used to enable features that are not fully production ready.",
						MarkdownDescription: "FeatureGates is used to enable or disable tech-preview Calico features. Values are specified in a comma separated list with no spaces, example; 'BPFConnectTimeLoadBalancingWorkaround=enabled,XyZ=false'. This is used to enable features that are not fully production ready.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([a-zA-Z0-9-_]+=([^=]+),)*([a-zA-Z0-9-_]+=([^=]+))?$`), ""),
						},
					},

					"floating_i_ps": schema.StringAttribute{
						Description:         "FloatingIPs configures whether or not Felix will program non-OpenStack floating IP addresses. (OpenStack-derived floating IPs are always programmed, regardless of this setting.)",
						MarkdownDescription: "FloatingIPs configures whether or not Felix will program non-OpenStack floating IP addresses. (OpenStack-derived floating IPs are always programmed, regardless of this setting.)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"flow_logs_collector_debug_trace": schema.BoolAttribute{
						Description:         "When FlowLogsCollectorDebugTrace is set to true, enables the logs in the collector to be printed in their entirety.",
						MarkdownDescription: "When FlowLogsCollectorDebugTrace is set to true, enables the logs in the collector to be printed in their entirety.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"flow_logs_flush_interval": schema.StringAttribute{
						Description:         "FlowLogsFlushInterval configures the interval at which Felix exports flow logs.",
						MarkdownDescription: "FlowLogsFlushInterval configures the interval at which Felix exports flow logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"flow_logs_goldmane_server": schema.StringAttribute{
						Description:         "FlowLogGoldmaneServer is the flow server endpoint to which flow data should be published.",
						MarkdownDescription: "FlowLogGoldmaneServer is the flow server endpoint to which flow data should be published.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"flow_logs_local_reporter": schema.StringAttribute{
						Description:         "FlowLogsLocalReporter configures local unix socket for reporting flow data from each node. [Default: Disabled]",
						MarkdownDescription: "FlowLogsLocalReporter configures local unix socket for reporting flow data from each node. [Default: Disabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Disabled", "Enabled"),
						},
					},

					"flow_logs_policy_evaluation_mode": schema.StringAttribute{
						Description:         "Continuous - Felix evaluates active flows on a regular basis to determine the rule traces in the flow logs. Any policy updates that impact a flow will be reflected in the pending_policies field, offering a near-real-time view of policy changes across flows. None - Felix stops evaluating pending traces. [Default: Continuous]",
						MarkdownDescription: "Continuous - Felix evaluates active flows on a regular basis to determine the rule traces in the flow logs. Any policy updates that impact a flow will be reflected in the pending_policies field, offering a near-real-time view of policy changes across flows. None - Felix stops evaluating pending traces. [Default: Continuous]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("None", "Continuous"),
						},
					},

					"generic_xdp_enabled": schema.BoolAttribute{
						Description:         "GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver modes can use XDP. This is not recommended since it doesn't provide better performance than iptables. [Default: false]",
						MarkdownDescription: "GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver modes can use XDP. This is not recommended since it doesn't provide better performance than iptables. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"go_gc_threshold": schema.Int64Attribute{
						Description:         "GoGCThreshold Sets the Go runtime's garbage collection threshold. I.e. the percentage that the heap is allowed to grow before garbage collection is triggered. In general, doubling the value halves the CPU time spent doing GC, but it also doubles peak GC memory overhead. A special value of -1 can be used to disable GC entirely; this should only be used in conjunction with the GoMemoryLimitMB setting. This setting is overridden by the GOGC environment variable. [Default: 40]",
						MarkdownDescription: "GoGCThreshold Sets the Go runtime's garbage collection threshold. I.e. the percentage that the heap is allowed to grow before garbage collection is triggered. In general, doubling the value halves the CPU time spent doing GC, but it also doubles peak GC memory overhead. A special value of -1 can be used to disable GC entirely; this should only be used in conjunction with the GoMemoryLimitMB setting. This setting is overridden by the GOGC environment variable. [Default: 40]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"go_max_procs": schema.Int64Attribute{
						Description:         "GoMaxProcs sets the maximum number of CPUs that the Go runtime will use concurrently. A value of -1 means 'use the system default'; typically the number of real CPUs on the system. this setting is overridden by the GOMAXPROCS environment variable. [Default: -1]",
						MarkdownDescription: "GoMaxProcs sets the maximum number of CPUs that the Go runtime will use concurrently. A value of -1 means 'use the system default'; typically the number of real CPUs on the system. this setting is overridden by the GOMAXPROCS environment variable. [Default: -1]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"go_memory_limit_mb": schema.Int64Attribute{
						Description:         "GoMemoryLimitMB sets a (soft) memory limit for the Go runtime in MB. The Go runtime will try to keep its memory usage under the limit by triggering GC as needed. To avoid thrashing, it will exceed the limit if GC starts to take more than 50% of the process's CPU time. A value of -1 disables the memory limit. Note that the memory limit, if used, must be considerably less than any hard resource limit set at the container or pod level. This is because felix is not the only process that must run in the container or pod. This setting is overridden by the GOMEMLIMIT environment variable. [Default: -1]",
						MarkdownDescription: "GoMemoryLimitMB sets a (soft) memory limit for the Go runtime in MB. The Go runtime will try to keep its memory usage under the limit by triggering GC as needed. To avoid thrashing, it will exceed the limit if GC starts to take more than 50% of the process's CPU time. A value of -1 disables the memory limit. Note that the memory limit, if used, must be considerably less than any hard resource limit set at the container or pod level. This is because felix is not the only process that must run in the container or pod. This setting is overridden by the GOMEMLIMIT environment variable. [Default: -1]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_enabled": schema.BoolAttribute{
						Description:         "HealthEnabled if set to true, enables Felix's health port, which provides readiness and liveness endpoints. [Default: false]",
						MarkdownDescription: "HealthEnabled if set to true, enables Felix's health port, which provides readiness and liveness endpoints. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_host": schema.StringAttribute{
						Description:         "HealthHost is the host that the health server should bind to. [Default: localhost]",
						MarkdownDescription: "HealthHost is the host that the health server should bind to. [Default: localhost]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_port": schema.Int64Attribute{
						Description:         "HealthPort is the TCP port that the health server should bind to. [Default: 9099]",
						MarkdownDescription: "HealthPort is the TCP port that the health server should bind to. [Default: 9099]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_timeout_overrides": schema.ListNestedAttribute{
						Description:         "HealthTimeoutOverrides allows the internal watchdog timeouts of individual subcomponents to be overridden. This is useful for working around 'false positive' liveness timeouts that can occur in particularly stressful workloads or if CPU is constrained. For a list of active subcomponents, see Felix's logs.",
						MarkdownDescription: "HealthTimeoutOverrides allows the internal watchdog timeouts of individual subcomponents to be overridden. This is useful for working around 'false positive' liveness timeouts that can occur in particularly stressful workloads or if CPU is constrained. For a list of active subcomponents, see Felix's logs.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"timeout": schema.StringAttribute{
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

					"interface_exclude": schema.StringAttribute{
						Description:         "InterfaceExclude A comma-separated list of interface names that should be excluded when Felix is resolving host endpoints. The default value ensures that Felix ignores Kubernetes' internal 'kube-ipvs0' device. If you want to exclude multiple interface names using a single value, the list supports regular expressions. For regular expressions you must wrap the value with '/'. For example having values '/^kube/,veth1' will exclude all interfaces that begin with 'kube' and also the interface 'veth1'. [Default: kube-ipvs0]",
						MarkdownDescription: "InterfaceExclude A comma-separated list of interface names that should be excluded when Felix is resolving host endpoints. The default value ensures that Felix ignores Kubernetes' internal 'kube-ipvs0' device. If you want to exclude multiple interface names using a single value, the list supports regular expressions. For regular expressions you must wrap the value with '/'. For example having values '/^kube/,veth1' will exclude all interfaces that begin with 'kube' and also the interface 'veth1'. [Default: kube-ipvs0]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interface_prefix": schema.StringAttribute{
						Description:         "InterfacePrefix is the interface name prefix that identifies workload endpoints and so distinguishes them from host endpoint interfaces. Note: in environments other than bare metal, the orchestrators configure this appropriately. For example our Kubernetes and Docker integrations set the 'cali' value, and our OpenStack integration sets the 'tap' value. [Default: cali]",
						MarkdownDescription: "InterfacePrefix is the interface name prefix that identifies workload endpoints and so distinguishes them from host endpoint interfaces. Note: in environments other than bare metal, the orchestrators configure this appropriately. For example our Kubernetes and Docker integrations set the 'cali' value, and our OpenStack integration sets the 'tap' value. [Default: cali]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"interface_refresh_interval": schema.StringAttribute{
						Description:         "InterfaceRefreshInterval is the period at which Felix rescans local interfaces to verify their state. The rescan can be disabled by setting the interval to 0.",
						MarkdownDescription: "InterfaceRefreshInterval is the period at which Felix rescans local interfaces to verify their state. The rescan can be disabled by setting the interval to 0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"ip_forwarding": schema.StringAttribute{
						Description:         "IPForwarding controls whether Felix sets the host sysctls to enable IP forwarding. IP forwarding is required when using Calico for workload networking. This should be disabled only on hosts where Calico is used solely for host protection. In BPF mode, due to a kernel interaction, either IPForwarding must be enabled or BPFEnforceRPF must be disabled. [Default: Enabled]",
						MarkdownDescription: "IPForwarding controls whether Felix sets the host sysctls to enable IP forwarding. IP forwarding is required when using Calico for workload networking. This should be disabled only on hosts where Calico is used solely for host protection. In BPF mode, due to a kernel interaction, either IPForwarding must be enabled or BPFEnforceRPF must be disabled. [Default: Enabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"ipip_enabled": schema.BoolAttribute{
						Description:         "IPIPEnabled overrides whether Felix should configure an IPIP interface on the host. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						MarkdownDescription: "IPIPEnabled overrides whether Felix should configure an IPIP interface on the host. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipip_mtu": schema.Int64Attribute{
						Description:         "IPIPMTU controls the MTU to set on the IPIP tunnel device. Optional as Felix auto-detects the MTU based on the MTU of the host's interfaces. [Default: 0 (auto-detect)]",
						MarkdownDescription: "IPIPMTU controls the MTU to set on the IPIP tunnel device. Optional as Felix auto-detects the MTU based on the MTU of the host's interfaces. [Default: 0 (auto-detect)]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipsets_refresh_interval": schema.StringAttribute{
						Description:         "IpsetsRefreshInterval controls the period at which Felix re-checks all IP sets to look for discrepancies. Set to 0 to disable the periodic refresh. [Default: 90s]",
						MarkdownDescription: "IpsetsRefreshInterval controls the period at which Felix re-checks all IP sets to look for discrepancies. Set to 0 to disable the periodic refresh. [Default: 90s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"iptables_backend": schema.StringAttribute{
						Description:         "IptablesBackend controls which backend of iptables will be used. The default is 'Auto'. Warning: changing this on a running system can leave 'orphaned' rules in the 'other' backend. These should be cleaned up to avoid confusing interactions.",
						MarkdownDescription: "IptablesBackend controls which backend of iptables will be used. The default is 'Auto'. Warning: changing this on a running system can leave 'orphaned' rules in the 'other' backend. These should be cleaned up to avoid confusing interactions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Auto|Legacy|NFT)?$`), ""),
						},
					},

					"iptables_filter_allow_action": schema.StringAttribute{
						Description:         "IptablesFilterAllowAction controls what happens to traffic that is accepted by a Felix policy chain in the iptables filter table (which is used for 'normal' policy). The default will immediately 'Accept' the traffic. Use 'Return' to send the traffic back up to the system chains for further processing.",
						MarkdownDescription: "IptablesFilterAllowAction controls what happens to traffic that is accepted by a Felix policy chain in the iptables filter table (which is used for 'normal' policy). The default will immediately 'Accept' the traffic. Use 'Return' to send the traffic back up to the system chains for further processing.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Accept|Return)?$`), ""),
						},
					},

					"iptables_filter_deny_action": schema.StringAttribute{
						Description:         "IptablesFilterDenyAction controls what happens to traffic that is denied by network policy. By default Calico blocks traffic with an iptables 'DROP' action. If you want to use 'REJECT' action instead you can configure it in here.",
						MarkdownDescription: "IptablesFilterDenyAction controls what happens to traffic that is denied by network policy. By default Calico blocks traffic with an iptables 'DROP' action. If you want to use 'REJECT' action instead you can configure it in here.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Drop|Reject)?$`), ""),
						},
					},

					"iptables_lock_file_path": schema.StringAttribute{
						Description:         "IptablesLockFilePath is the location of the iptables lock file. You may need to change this if the lock file is not in its standard location (for example if you have mapped it into Felix's container at a different path). [Default: /run/xtables.lock]",
						MarkdownDescription: "IptablesLockFilePath is the location of the iptables lock file. You may need to change this if the lock file is not in its standard location (for example if you have mapped it into Felix's container at a different path). [Default: /run/xtables.lock]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"iptables_lock_probe_interval": schema.StringAttribute{
						Description:         "IptablesLockProbeInterval when IptablesLockTimeout is enabled: the time that Felix will wait between attempts to acquire the iptables lock if it is not available. Lower values make Felix more responsive when the lock is contended, but use more CPU. [Default: 50ms]",
						MarkdownDescription: "IptablesLockProbeInterval when IptablesLockTimeout is enabled: the time that Felix will wait between attempts to acquire the iptables lock if it is not available. Lower values make Felix more responsive when the lock is contended, but use more CPU. [Default: 50ms]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"iptables_lock_timeout": schema.StringAttribute{
						Description:         "IptablesLockTimeout is the time that Felix itself will wait for the iptables lock (rather than delegating the lock handling to the 'iptables' command). Deprecated: 'iptables-restore' v1.8+ always takes the lock, so enabling this feature results in deadlock. [Default: 0s disabled]",
						MarkdownDescription: "IptablesLockTimeout is the time that Felix itself will wait for the iptables lock (rather than delegating the lock handling to the 'iptables' command). Deprecated: 'iptables-restore' v1.8+ always takes the lock, so enabling this feature results in deadlock. [Default: 0s disabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"iptables_mangle_allow_action": schema.StringAttribute{
						Description:         "IptablesMangleAllowAction controls what happens to traffic that is accepted by a Felix policy chain in the iptables mangle table (which is used for 'pre-DNAT' policy). The default will immediately 'Accept' the traffic. Use 'Return' to send the traffic back up to the system chains for further processing.",
						MarkdownDescription: "IptablesMangleAllowAction controls what happens to traffic that is accepted by a Felix policy chain in the iptables mangle table (which is used for 'pre-DNAT' policy). The default will immediately 'Accept' the traffic. Use 'Return' to send the traffic back up to the system chains for further processing.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Accept|Return)?$`), ""),
						},
					},

					"iptables_mark_mask": schema.Int64Attribute{
						Description:         "IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xffff0000]",
						MarkdownDescription: "IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xffff0000]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"iptables_nat_outgoing_interface_filter": schema.StringAttribute{
						Description:         "This parameter can be used to limit the host interfaces on which Calico will apply SNAT to traffic leaving a Calico IPAM pool with 'NAT outgoing' enabled. This can be useful if you have a main data interface, where traffic should be SNATted and a secondary device (such as the docker bridge) which is local to the host and doesn't require SNAT. This parameter uses the iptables interface matching syntax, which allows + as a wildcard. Most users will not need to set this. Example: if your data interfaces are eth0 and eth1 and you want to exclude the docker bridge, you could set this to eth+",
						MarkdownDescription: "This parameter can be used to limit the host interfaces on which Calico will apply SNAT to traffic leaving a Calico IPAM pool with 'NAT outgoing' enabled. This can be useful if you have a main data interface, where traffic should be SNATted and a secondary device (such as the docker bridge) which is local to the host and doesn't require SNAT. This parameter uses the iptables interface matching syntax, which allows + as a wildcard. Most users will not need to set this. Example: if your data interfaces are eth0 and eth1 and you want to exclude the docker bridge, you could set this to eth+",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"iptables_post_write_check_interval": schema.StringAttribute{
						Description:         "IptablesPostWriteCheckInterval is the period after Felix has done a write to the dataplane that it schedules an extra read back in order to check the write was not clobbered by another process. This should only occur if another application on the system doesn't respect the iptables lock. [Default: 1s]",
						MarkdownDescription: "IptablesPostWriteCheckInterval is the period after Felix has done a write to the dataplane that it schedules an extra read back in order to check the write was not clobbered by another process. This should only occur if another application on the system doesn't respect the iptables lock. [Default: 1s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"iptables_refresh_interval": schema.StringAttribute{
						Description:         "IptablesRefreshInterval is the period at which Felix re-checks the IP sets in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable IP sets refresh. Note: the default for this value is lower than the other refresh intervals as a workaround for a Linux kernel bug that was fixed in kernel version 4.11. If you are using v4.11 or greater you may want to set this to, a higher value to reduce Felix CPU usage. [Default: 10s]",
						MarkdownDescription: "IptablesRefreshInterval is the period at which Felix re-checks the IP sets in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable IP sets refresh. Note: the default for this value is lower than the other refresh intervals as a workaround for a Linux kernel bug that was fixed in kernel version 4.11. If you are using v4.11 or greater you may want to set this to, a higher value to reduce Felix CPU usage. [Default: 10s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"ipv6_support": schema.BoolAttribute{
						Description:         "IPv6Support controls whether Felix enables support for IPv6 (if supported by the in-use dataplane).",
						MarkdownDescription: "IPv6Support controls whether Felix enables support for IPv6 (if supported by the in-use dataplane).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kube_node_port_ranges": schema.ListAttribute{
						Description:         "KubeNodePortRanges holds list of port ranges used for service node ports. Only used if felix detects kube-proxy running in ipvs mode. Felix uses these ranges to separate host and workload traffic. [Default: 30000:32767].",
						MarkdownDescription: "KubeNodePortRanges holds list of port ranges used for service node ports. Only used if felix detects kube-proxy running in ipvs mode. Felix uses these ranges to separate host and workload traffic. [Default: 30000:32767].",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_debug_filename_regex": schema.StringAttribute{
						Description:         "LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs. Only logs from files with names that match the given regular expression are included. The filter only applies to Debug level logs.",
						MarkdownDescription: "LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs. Only logs from files with names that match the given regular expression are included. The filter only applies to Debug level logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_file_path": schema.StringAttribute{
						Description:         "LogFilePath is the full path to the Felix log. Set to none to disable file logging. [Default: /var/log/calico/felix.log]",
						MarkdownDescription: "LogFilePath is the full path to the Felix log. Set to none to disable file logging. [Default: /var/log/calico/felix.log]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_prefix": schema.StringAttribute{
						Description:         "LogPrefix is the log prefix that Felix uses when rendering LOG rules. [Default: calico-packet]",
						MarkdownDescription: "LogPrefix is the log prefix that Felix uses when rendering LOG rules. [Default: calico-packet]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_severity_file": schema.StringAttribute{
						Description:         "LogSeverityFile is the log severity above which logs are sent to the log file. [Default: Info]",
						MarkdownDescription: "LogSeverityFile is the log severity above which logs are sent to the log file. [Default: Info]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Trace|Debug|Info|Warning|Error|Fatal)?$`), ""),
						},
					},

					"log_severity_screen": schema.StringAttribute{
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Trace|Debug|Info|Warning|Error|Fatal)?$`), ""),
						},
					},

					"log_severity_sys": schema.StringAttribute{
						Description:         "LogSeveritySys is the log severity above which logs are sent to the syslog. Set to None for no logging to syslog. [Default: Info]",
						MarkdownDescription: "LogSeveritySys is the log severity above which logs are sent to the syslog. Set to None for no logging to syslog. [Default: Info]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Trace|Debug|Info|Warning|Error|Fatal)?$`), ""),
						},
					},

					"max_ipset_size": schema.Int64Attribute{
						Description:         "MaxIpsetSize is the maximum number of IP addresses that can be stored in an IP set. Not applicable if using the nftables backend.",
						MarkdownDescription: "MaxIpsetSize is the maximum number of IP addresses that can be stored in an IP set. Not applicable if using the nftables backend.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata_addr": schema.StringAttribute{
						Description:         "MetadataAddr is the IP address or domain name of the server that can answer VM queries for cloud-init metadata. In OpenStack, this corresponds to the machine running nova-api (or in Ubuntu, nova-api-metadata). A value of none (case-insensitive) means that Felix should not set up any NAT rule for the metadata path. [Default: 127.0.0.1]",
						MarkdownDescription: "MetadataAddr is the IP address or domain name of the server that can answer VM queries for cloud-init metadata. In OpenStack, this corresponds to the machine running nova-api (or in Ubuntu, nova-api-metadata). A value of none (case-insensitive) means that Felix should not set up any NAT rule for the metadata path. [Default: 127.0.0.1]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata_port": schema.Int64Attribute{
						Description:         "MetadataPort is the port of the metadata server. This, combined with global.MetadataAddr (if not 'None'), is used to set up a NAT rule, from 169.254.169.254:80 to MetadataAddr:MetadataPort. In most cases this should not need to be changed [Default: 8775].",
						MarkdownDescription: "MetadataPort is the port of the metadata server. This, combined with global.MetadataAddr (if not 'None'), is used to set up a NAT rule, from 169.254.169.254:80 to MetadataAddr:MetadataPort. In most cases this should not need to be changed [Default: 8775].",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mtu_iface_pattern": schema.StringAttribute{
						Description:         "MTUIfacePattern is a regular expression that controls which interfaces Felix should scan in order to calculate the host's MTU. This should not match workload interfaces (usually named cali...).",
						MarkdownDescription: "MTUIfacePattern is a regular expression that controls which interfaces Felix should scan in order to calculate the host's MTU. This should not match workload interfaces (usually named cali...).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nat_outgoing_address": schema.StringAttribute{
						Description:         "NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that is leaving the network. By default the address used is an address on the interface the traffic is leaving on (i.e. it uses the iptables MASQUERADE target).",
						MarkdownDescription: "NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that is leaving the network. By default the address used is an address on the interface the traffic is leaving on (i.e. it uses the iptables MASQUERADE target).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nat_outgoing_exclusions": schema.StringAttribute{
						Description:         "When a IP pool setting 'natOutgoing' is true, packets sent from Calico networked containers in this IP pool to destinations will be masqueraded. Configure which type of destinations is excluded from being masqueraded. - IPPoolsOnly: destinations outside of this IP pool will be masqueraded. - IPPoolsAndHostIPs: destinations outside of this IP pool and all hosts will be masqueraded. [Default: IPPoolsOnly]",
						MarkdownDescription: "When a IP pool setting 'natOutgoing' is true, packets sent from Calico networked containers in this IP pool to destinations will be masqueraded. Configure which type of destinations is excluded from being masqueraded. - IPPoolsOnly: destinations outside of this IP pool will be masqueraded. - IPPoolsAndHostIPs: destinations outside of this IP pool and all hosts will be masqueraded. [Default: IPPoolsOnly]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("IPPoolsOnly", "IPPoolsAndHostIPs"),
						},
					},

					"nat_port_range": schema.StringAttribute{
						Description:         "NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT. When unset the default behavior of the network stack is used.",
						MarkdownDescription: "NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT. When unset the default behavior of the network stack is used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"netlink_timeout": schema.StringAttribute{
						Description:         "NetlinkTimeout is the timeout when talking to the kernel over the netlink protocol, used for programming routes, rules, and other kernel objects. [Default: 10s]",
						MarkdownDescription: "NetlinkTimeout is the timeout when talking to the kernel over the netlink protocol, used for programming routes, rules, and other kernel objects. [Default: 10s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"nftables_filter_allow_action": schema.StringAttribute{
						Description:         "NftablesFilterAllowAction controls the nftables action that Felix uses to represent the 'allow' policy verdict in the filter table. The default is to 'ACCEPT' the traffic, which is a terminal action. Alternatively, 'RETURN' can be used to return the traffic back to the top-level chain for further processing by your rules.",
						MarkdownDescription: "NftablesFilterAllowAction controls the nftables action that Felix uses to represent the 'allow' policy verdict in the filter table. The default is to 'ACCEPT' the traffic, which is a terminal action. Alternatively, 'RETURN' can be used to return the traffic back to the top-level chain for further processing by your rules.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Accept|Return)?$`), ""),
						},
					},

					"nftables_filter_deny_action": schema.StringAttribute{
						Description:         "NftablesFilterDenyAction controls what happens to traffic that is denied by network policy. By default, Calico blocks traffic with a 'drop' action. If you want to use a 'reject' action instead you can configure it here.",
						MarkdownDescription: "NftablesFilterDenyAction controls what happens to traffic that is denied by network policy. By default, Calico blocks traffic with a 'drop' action. If you want to use a 'reject' action instead you can configure it here.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Drop|Reject)?$`), ""),
						},
					},

					"nftables_mangle_allow_action": schema.StringAttribute{
						Description:         "NftablesMangleAllowAction controls the nftables action that Felix uses to represent the 'allow' policy verdict in the mangle table. The default is to 'ACCEPT' the traffic, which is a terminal action. Alternatively, 'RETURN' can be used to return the traffic back to the top-level chain for further processing by your rules.",
						MarkdownDescription: "NftablesMangleAllowAction controls the nftables action that Felix uses to represent the 'allow' policy verdict in the mangle table. The default is to 'ACCEPT' the traffic, which is a terminal action. Alternatively, 'RETURN' can be used to return the traffic back to the top-level chain for further processing by your rules.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Accept|Return)?$`), ""),
						},
					},

					"nftables_mark_mask": schema.Int64Attribute{
						Description:         "NftablesMarkMask is the mask that Felix selects its nftables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xffff0000]",
						MarkdownDescription: "NftablesMarkMask is the mask that Felix selects its nftables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xffff0000]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nftables_mode": schema.StringAttribute{
						Description:         "NFTablesMode configures nftables support in Felix. [Default: Disabled]",
						MarkdownDescription: "NFTablesMode configures nftables support in Felix. [Default: Disabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Disabled", "Enabled", "Auto"),
						},
					},

					"nftables_refresh_interval": schema.StringAttribute{
						Description:         "NftablesRefreshInterval controls the interval at which Felix periodically refreshes the nftables rules. [Default: 90s]",
						MarkdownDescription: "NftablesRefreshInterval controls the interval at which Felix periodically refreshes the nftables rules. [Default: 90s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"openstack_region": schema.StringAttribute{
						Description:         "OpenstackRegion is the name of the region that a particular Felix belongs to. In a multi-region Calico/OpenStack deployment, this must be configured somehow for each Felix (here in the datamodel, or in felix.cfg or the environment on each compute node), and must match the [calico] openstack_region value configured in neutron.conf on each node. [Default: Empty]",
						MarkdownDescription: "OpenstackRegion is the name of the region that a particular Felix belongs to. In a multi-region Calico/OpenStack deployment, this must be configured somehow for each Felix (here in the datamodel, or in felix.cfg or the environment on each compute node), and must match the [calico] openstack_region value configured in neutron.conf on each node. [Default: Empty]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"policy_sync_path_prefix": schema.StringAttribute{
						Description:         "PolicySyncPathPrefix is used to by Felix to communicate policy changes to external services, like Application layer policy. [Default: Empty]",
						MarkdownDescription: "PolicySyncPathPrefix is used to by Felix to communicate policy changes to external services, like Application layer policy. [Default: Empty]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"program_cluster_routes": schema.StringAttribute{
						Description:         "ProgramClusterRoutes specifies whether Felix should program IPIP routes instead of BIRD. Felix always programs VXLAN routes. [Default: Disabled]",
						MarkdownDescription: "ProgramClusterRoutes specifies whether Felix should program IPIP routes instead of BIRD. Felix always programs VXLAN routes. [Default: Disabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"prometheus_go_metrics_enabled": schema.BoolAttribute{
						Description:         "PrometheusGoMetricsEnabled disables Go runtime metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						MarkdownDescription: "PrometheusGoMetricsEnabled disables Go runtime metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_metrics_enabled": schema.BoolAttribute{
						Description:         "PrometheusMetricsEnabled enables the Prometheus metrics server in Felix if set to true. [Default: false]",
						MarkdownDescription: "PrometheusMetricsEnabled enables the Prometheus metrics server in Felix if set to true. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_metrics_host": schema.StringAttribute{
						Description:         "PrometheusMetricsHost is the host that the Prometheus metrics server should bind to. [Default: empty]",
						MarkdownDescription: "PrometheusMetricsHost is the host that the Prometheus metrics server should bind to. [Default: empty]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_metrics_port": schema.Int64Attribute{
						Description:         "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. [Default: 9091]",
						MarkdownDescription: "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. [Default: 9091]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_process_metrics_enabled": schema.BoolAttribute{
						Description:         "PrometheusProcessMetricsEnabled disables process metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						MarkdownDescription: "PrometheusProcessMetricsEnabled disables process metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"prometheus_wire_guard_metrics_enabled": schema.BoolAttribute{
						Description:         "PrometheusWireGuardMetricsEnabled disables wireguard metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						MarkdownDescription: "PrometheusWireGuardMetricsEnabled disables wireguard metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"remove_external_routes": schema.BoolAttribute{
						Description:         "RemoveExternalRoutes Controls whether Felix will remove unexpected routes to workload interfaces. Felix will always clean up expected routes that use the configured DeviceRouteProtocol. To add your own routes, you must use a distinct protocol (in addition to setting this field to false).",
						MarkdownDescription: "RemoveExternalRoutes Controls whether Felix will remove unexpected routes to workload interfaces. Felix will always clean up expected routes that use the configured DeviceRouteProtocol. To add your own routes, you must use a distinct protocol (in addition to setting this field to false).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"reporting_interval": schema.StringAttribute{
						Description:         "ReportingInterval is the interval at which Felix reports its status into the datastore or 0 to disable. Must be non-zero in OpenStack deployments. [Default: 30s]",
						MarkdownDescription: "ReportingInterval is the interval at which Felix reports its status into the datastore or 0 to disable. Must be non-zero in OpenStack deployments. [Default: 30s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"reporting_ttl": schema.StringAttribute{
						Description:         "ReportingTTL is the time-to-live setting for process-wide status reports. [Default: 90s]",
						MarkdownDescription: "ReportingTTL is the time-to-live setting for process-wide status reports. [Default: 90s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"require_mtu_file": schema.BoolAttribute{
						Description:         "RequireMTUFile specifies whether mtu file is required to start the felix. Optional as to keep the same as previous behavior. [Default: false]",
						MarkdownDescription: "RequireMTUFile specifies whether mtu file is required to start the felix. Optional as to keep the same as previous behavior. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_refresh_interval": schema.StringAttribute{
						Description:         "RouteRefreshInterval is the period at which Felix re-checks the routes in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable route refresh. [Default: 90s]",
						MarkdownDescription: "RouteRefreshInterval is the period at which Felix re-checks the routes in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable route refresh. [Default: 90s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"route_source": schema.StringAttribute{
						Description:         "RouteSource configures where Felix gets its routing information. - WorkloadIPs: use workload endpoints to construct routes. - CalicoIPAM: the default - use IPAM data to construct routes.",
						MarkdownDescription: "RouteSource configures where Felix gets its routing information. - WorkloadIPs: use workload endpoints to construct routes. - CalicoIPAM: the default - use IPAM data to construct routes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(WorkloadIPs|CalicoIPAM)?$`), ""),
						},
					},

					"route_sync_disabled": schema.BoolAttribute{
						Description:         "RouteSyncDisabled will disable all operations performed on the route table. Set to true to run in network-policy mode only.",
						MarkdownDescription: "RouteSyncDisabled will disable all operations performed on the route table. Set to true to run in network-policy mode only.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_table_range": schema.SingleNestedAttribute{
						Description:         "Deprecated in favor of RouteTableRanges. Calico programs additional Linux route tables for various purposes. RouteTableRange specifies the indices of the route tables that Calico should use.",
						MarkdownDescription: "Deprecated in favor of RouteTableRanges. Calico programs additional Linux route tables for various purposes. RouteTableRange specifies the indices of the route tables that Calico should use.",
						Attributes: map[string]schema.Attribute{
							"max": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"min": schema.Int64Attribute{
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

					"route_table_ranges": schema.ListNestedAttribute{
						Description:         "Calico programs additional Linux route tables for various purposes. RouteTableRanges specifies a set of table index ranges that Calico should use. Deprecates'RouteTableRange', overrides 'RouteTableRange'.",
						MarkdownDescription: "Calico programs additional Linux route tables for various purposes. RouteTableRanges specifies a set of table index ranges that Calico should use. Deprecates'RouteTableRange', overrides 'RouteTableRange'.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"max": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"min": schema.Int64Attribute{
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

					"service_loop_prevention": schema.StringAttribute{
						Description:         "When service IP advertisement is enabled, prevent routing loops to service IPs that are not in use, by dropping or rejecting packets that do not get DNAT'd by kube-proxy. Unless set to 'Disabled', in which case such routing loops continue to be allowed. [Default: Drop]",
						MarkdownDescription: "When service IP advertisement is enabled, prevent routing loops to service IPs that are not in use, by dropping or rejecting packets that do not get DNAT'd by kube-proxy. Unless set to 'Disabled', in which case such routing loops continue to be allowed. [Default: Drop]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Drop|Reject|Disabled)?$`), ""),
						},
					},

					"sidecar_acceleration_enabled": schema.BoolAttribute{
						Description:         "SidecarAccelerationEnabled enables experimental sidecar acceleration [Default: false]",
						MarkdownDescription: "SidecarAccelerationEnabled enables experimental sidecar acceleration [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"usage_reporting_enabled": schema.BoolAttribute{
						Description:         "UsageReportingEnabled reports anonymous Calico version number and cluster size to projectcalico.org. Logs warnings returned by the usage server. For example, if a significant security vulnerability has been discovered in the version of Calico being used. [Default: true]",
						MarkdownDescription: "UsageReportingEnabled reports anonymous Calico version number and cluster size to projectcalico.org. Logs warnings returned by the usage server. For example, if a significant security vulnerability has been discovered in the version of Calico being used. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"usage_reporting_initial_delay": schema.StringAttribute{
						Description:         "UsageReportingInitialDelay controls the minimum delay before Felix makes a report. [Default: 300s]",
						MarkdownDescription: "UsageReportingInitialDelay controls the minimum delay before Felix makes a report. [Default: 300s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"usage_reporting_interval": schema.StringAttribute{
						Description:         "UsageReportingInterval controls the interval at which Felix makes reports. [Default: 86400s]",
						MarkdownDescription: "UsageReportingInterval controls the interval at which Felix makes reports. [Default: 86400s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"use_internal_dataplane_driver": schema.BoolAttribute{
						Description:         "UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic. If false, it will launch an external dataplane driver and communicate with it over protobuf.",
						MarkdownDescription: "UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic. If false, it will launch an external dataplane driver and communicate with it over protobuf.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_enabled": schema.BoolAttribute{
						Description:         "VXLANEnabled overrides whether Felix should create the VXLAN tunnel device for IPv4 VXLAN networking. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						MarkdownDescription: "VXLANEnabled overrides whether Felix should create the VXLAN tunnel device for IPv4 VXLAN networking. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_mtu": schema.Int64Attribute{
						Description:         "VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device. Optional as Felix auto-detects the MTU based on the MTU of the host's interfaces. [Default: 0 (auto-detect)]",
						MarkdownDescription: "VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device. Optional as Felix auto-detects the MTU based on the MTU of the host's interfaces. [Default: 0 (auto-detect)]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_mtuv6": schema.Int64Attribute{
						Description:         "VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device. Optional as Felix auto-detects the MTU based on the MTU of the host's interfaces. [Default: 0 (auto-detect)]",
						MarkdownDescription: "VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device. Optional as Felix auto-detects the MTU based on the MTU of the host's interfaces. [Default: 0 (auto-detect)]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_port": schema.Int64Attribute{
						Description:         "VXLANPort is the UDP port number to use for VXLAN traffic. [Default: 4789]",
						MarkdownDescription: "VXLANPort is the UDP port number to use for VXLAN traffic. [Default: 4789]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_vni": schema.Int64Attribute{
						Description:         "VXLANVNI is the VXLAN VNI to use for VXLAN traffic. You may need to change this if the default value is in use on your system. [Default: 4096]",
						MarkdownDescription: "VXLANVNI is the VXLAN VNI to use for VXLAN traffic. You may need to change this if the default value is in use on your system. [Default: 4096]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"windows_manage_firewall_rules": schema.StringAttribute{
						Description:         "WindowsManageFirewallRules configures whether or not Felix will program Windows Firewall rules (to allow inbound access to its own metrics ports). [Default: Disabled]",
						MarkdownDescription: "WindowsManageFirewallRules configures whether or not Felix will program Windows Firewall rules (to allow inbound access to its own metrics ports). [Default: Disabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"wireguard_enabled": schema.BoolAttribute{
						Description:         "WireguardEnabled controls whether Wireguard is enabled for IPv4 (encapsulating IPv4 traffic over an IPv4 underlay network). [Default: false]",
						MarkdownDescription: "WireguardEnabled controls whether Wireguard is enabled for IPv4 (encapsulating IPv4 traffic over an IPv4 underlay network). [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_enabled_v6": schema.BoolAttribute{
						Description:         "WireguardEnabledV6 controls whether Wireguard is enabled for IPv6 (encapsulating IPv6 traffic over an IPv6 underlay network). [Default: false]",
						MarkdownDescription: "WireguardEnabledV6 controls whether Wireguard is enabled for IPv6 (encapsulating IPv6 traffic over an IPv6 underlay network). [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_host_encryption_enabled": schema.BoolAttribute{
						Description:         "WireguardHostEncryptionEnabled controls whether Wireguard host-to-host encryption is enabled. [Default: false]",
						MarkdownDescription: "WireguardHostEncryptionEnabled controls whether Wireguard host-to-host encryption is enabled. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_interface_name": schema.StringAttribute{
						Description:         "WireguardInterfaceName specifies the name to use for the IPv4 Wireguard interface. [Default: wireguard.cali]",
						MarkdownDescription: "WireguardInterfaceName specifies the name to use for the IPv4 Wireguard interface. [Default: wireguard.cali]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_interface_name_v6": schema.StringAttribute{
						Description:         "WireguardInterfaceNameV6 specifies the name to use for the IPv6 Wireguard interface. [Default: wg-v6.cali]",
						MarkdownDescription: "WireguardInterfaceNameV6 specifies the name to use for the IPv6 Wireguard interface. [Default: wg-v6.cali]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_keep_alive": schema.StringAttribute{
						Description:         "WireguardPersistentKeepAlive controls Wireguard PersistentKeepalive option. Set 0 to disable. [Default: 0]",
						MarkdownDescription: "WireguardPersistentKeepAlive controls Wireguard PersistentKeepalive option. Set 0 to disable. [Default: 0]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"wireguard_listening_port": schema.Int64Attribute{
						Description:         "WireguardListeningPort controls the listening port used by IPv4 Wireguard. [Default: 51820]",
						MarkdownDescription: "WireguardListeningPort controls the listening port used by IPv4 Wireguard. [Default: 51820]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_listening_port_v6": schema.Int64Attribute{
						Description:         "WireguardListeningPortV6 controls the listening port used by IPv6 Wireguard. [Default: 51821]",
						MarkdownDescription: "WireguardListeningPortV6 controls the listening port used by IPv6 Wireguard. [Default: 51821]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_mtu": schema.Int64Attribute{
						Description:         "WireguardMTU controls the MTU on the IPv4 Wireguard interface. See Configuring MTU [Default: 1440]",
						MarkdownDescription: "WireguardMTU controls the MTU on the IPv4 Wireguard interface. See Configuring MTU [Default: 1440]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_mtuv6": schema.Int64Attribute{
						Description:         "WireguardMTUV6 controls the MTU on the IPv6 Wireguard interface. See Configuring MTU [Default: 1420]",
						MarkdownDescription: "WireguardMTUV6 controls the MTU on the IPv6 Wireguard interface. See Configuring MTU [Default: 1420]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_routing_rule_priority": schema.Int64Attribute{
						Description:         "WireguardRoutingRulePriority controls the priority value to use for the Wireguard routing rule. [Default: 99]",
						MarkdownDescription: "WireguardRoutingRulePriority controls the priority value to use for the Wireguard routing rule. [Default: 99]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wireguard_threading_enabled": schema.BoolAttribute{
						Description:         "WireguardThreadingEnabled controls whether Wireguard has Threaded NAPI enabled. [Default: false] This increases the maximum number of packets a Wireguard interface can process. Consider threaded NAPI only if you have high packets per second workloads that are causing dropping packets due to a saturated 'softirq' CPU core. There is a [known issue](https://lore.kernel.org/netdev/CALrw=nEoT2emQ0OAYCjM1d_6Xe_kNLSZ6dhjb5FxrLFYh4kozA@mail.gmail.com/T/) with this setting that may cause NAPI to get stuck holding the global 'rtnl_mutex' when a peer is removed. Workaround: Make sure your Linux kernel [includes this patch](https://github.com/torvalds/linux/commit/56364c910691f6d10ba88c964c9041b9ab777bd6) to unwedge NAPI.",
						MarkdownDescription: "WireguardThreadingEnabled controls whether Wireguard has Threaded NAPI enabled. [Default: false] This increases the maximum number of packets a Wireguard interface can process. Consider threaded NAPI only if you have high packets per second workloads that are causing dropping packets due to a saturated 'softirq' CPU core. There is a [known issue](https://lore.kernel.org/netdev/CALrw=nEoT2emQ0OAYCjM1d_6Xe_kNLSZ6dhjb5FxrLFYh4kozA@mail.gmail.com/T/) with this setting that may cause NAPI to get stuck holding the global 'rtnl_mutex' when a peer is removed. Workaround: Make sure your Linux kernel [includes this patch](https://github.com/torvalds/linux/commit/56364c910691f6d10ba88c964c9041b9ab777bd6) to unwedge NAPI.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"workload_source_spoofing": schema.StringAttribute{
						Description:         "WorkloadSourceSpoofing controls whether pods can use the allowedSourcePrefixes annotation to send traffic with a source IP address that is not theirs. This is disabled by default. When set to 'Any', pods can request any prefix.",
						MarkdownDescription: "WorkloadSourceSpoofing controls whether pods can use the allowedSourcePrefixes annotation to send traffic with a source IP address that is not theirs. This is disabled by default. When set to 'Any', pods can request any prefix.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Disabled|Any)?$`), ""),
						},
					},

					"xdp_enabled": schema.BoolAttribute{
						Description:         "XDPEnabled enables XDP acceleration for suitable untracked incoming deny rules. [Default: true]",
						MarkdownDescription: "XDPEnabled enables XDP acceleration for suitable untracked incoming deny rules. [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"xdp_refresh_interval": schema.StringAttribute{
						Description:         "XDPRefreshInterval is the period at which Felix re-checks all XDP state to ensure that no other process has accidentally broken Calico's BPF maps or attached programs. Set to 0 to disable XDP refresh. [Default: 90s]",
						MarkdownDescription: "XDPRefreshInterval is the period at which Felix re-checks all XDP state to ensure that no other process has accidentally broken Calico's BPF maps or attached programs. Set to 0 to disable XDP refresh. [Default: 90s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_felix_configuration_v1_manifest")

	var model CrdProjectcalicoOrgFelixConfigurationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("FelixConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
