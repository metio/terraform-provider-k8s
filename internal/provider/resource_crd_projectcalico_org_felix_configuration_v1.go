/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type CrdProjectcalicoOrgFelixConfigurationV1Resource struct{}

var (
	_ resource.Resource = (*CrdProjectcalicoOrgFelixConfigurationV1Resource)(nil)
)

type CrdProjectcalicoOrgFelixConfigurationV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CrdProjectcalicoOrgFelixConfigurationV1GoModel struct {
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
		AllowIPIPPacketsFromWorkloads *bool `tfsdk:"allow_ipip_packets_from_workloads" yaml:"allowIPIPPacketsFromWorkloads,omitempty"`

		AllowVXLANPacketsFromWorkloads *bool `tfsdk:"allow_vxlan_packets_from_workloads" yaml:"allowVXLANPacketsFromWorkloads,omitempty"`

		AwsSrcDstCheck *string `tfsdk:"aws_src_dst_check" yaml:"awsSrcDstCheck,omitempty"`

		BpfConnectTimeLoadBalancingEnabled *bool `tfsdk:"bpf_connect_time_load_balancing_enabled" yaml:"bpfConnectTimeLoadBalancingEnabled,omitempty"`

		BpfDataIfacePattern *string `tfsdk:"bpf_data_iface_pattern" yaml:"bpfDataIfacePattern,omitempty"`

		BpfDisableUnprivileged *bool `tfsdk:"bpf_disable_unprivileged" yaml:"bpfDisableUnprivileged,omitempty"`

		BpfEnabled *bool `tfsdk:"bpf_enabled" yaml:"bpfEnabled,omitempty"`

		BpfEnforceRPF *string `tfsdk:"bpf_enforce_rpf" yaml:"bpfEnforceRPF,omitempty"`

		BpfExtToServiceConnmark *int64 `tfsdk:"bpf_ext_to_service_connmark" yaml:"bpfExtToServiceConnmark,omitempty"`

		BpfExternalServiceMode *string `tfsdk:"bpf_external_service_mode" yaml:"bpfExternalServiceMode,omitempty"`

		BpfHostConntrackBypass *bool `tfsdk:"bpf_host_conntrack_bypass" yaml:"bpfHostConntrackBypass,omitempty"`

		BpfKubeProxyEndpointSlicesEnabled *bool `tfsdk:"bpf_kube_proxy_endpoint_slices_enabled" yaml:"bpfKubeProxyEndpointSlicesEnabled,omitempty"`

		BpfKubeProxyIptablesCleanupEnabled *bool `tfsdk:"bpf_kube_proxy_iptables_cleanup_enabled" yaml:"bpfKubeProxyIptablesCleanupEnabled,omitempty"`

		BpfKubeProxyMinSyncPeriod *string `tfsdk:"bpf_kube_proxy_min_sync_period" yaml:"bpfKubeProxyMinSyncPeriod,omitempty"`

		BpfL3IfacePattern *string `tfsdk:"bpf_l3_iface_pattern" yaml:"bpfL3IfacePattern,omitempty"`

		BpfLogLevel *string `tfsdk:"bpf_log_level" yaml:"bpfLogLevel,omitempty"`

		BpfMapSizeConntrack *int64 `tfsdk:"bpf_map_size_conntrack" yaml:"bpfMapSizeConntrack,omitempty"`

		BpfMapSizeIPSets *int64 `tfsdk:"bpf_map_size_ip_sets" yaml:"bpfMapSizeIPSets,omitempty"`

		BpfMapSizeIfState *int64 `tfsdk:"bpf_map_size_if_state" yaml:"bpfMapSizeIfState,omitempty"`

		BpfMapSizeNATAffinity *int64 `tfsdk:"bpf_map_size_nat_affinity" yaml:"bpfMapSizeNATAffinity,omitempty"`

		BpfMapSizeNATBackend *int64 `tfsdk:"bpf_map_size_nat_backend" yaml:"bpfMapSizeNATBackend,omitempty"`

		BpfMapSizeNATFrontend *int64 `tfsdk:"bpf_map_size_nat_frontend" yaml:"bpfMapSizeNATFrontend,omitempty"`

		BpfMapSizeRoute *int64 `tfsdk:"bpf_map_size_route" yaml:"bpfMapSizeRoute,omitempty"`

		BpfPSNATPorts utilities.IntOrString `tfsdk:"bpf_psnat_ports" yaml:"bpfPSNATPorts,omitempty"`

		BpfPolicyDebugEnabled *bool `tfsdk:"bpf_policy_debug_enabled" yaml:"bpfPolicyDebugEnabled,omitempty"`

		ChainInsertMode *string `tfsdk:"chain_insert_mode" yaml:"chainInsertMode,omitempty"`

		DataplaneDriver *string `tfsdk:"dataplane_driver" yaml:"dataplaneDriver,omitempty"`

		DataplaneWatchdogTimeout *string `tfsdk:"dataplane_watchdog_timeout" yaml:"dataplaneWatchdogTimeout,omitempty"`

		DebugDisableLogDropping *bool `tfsdk:"debug_disable_log_dropping" yaml:"debugDisableLogDropping,omitempty"`

		DebugMemoryProfilePath *string `tfsdk:"debug_memory_profile_path" yaml:"debugMemoryProfilePath,omitempty"`

		DebugSimulateCalcGraphHangAfter *string `tfsdk:"debug_simulate_calc_graph_hang_after" yaml:"debugSimulateCalcGraphHangAfter,omitempty"`

		DebugSimulateDataplaneHangAfter *string `tfsdk:"debug_simulate_dataplane_hang_after" yaml:"debugSimulateDataplaneHangAfter,omitempty"`

		DefaultEndpointToHostAction *string `tfsdk:"default_endpoint_to_host_action" yaml:"defaultEndpointToHostAction,omitempty"`

		DeviceRouteProtocol *int64 `tfsdk:"device_route_protocol" yaml:"deviceRouteProtocol,omitempty"`

		DeviceRouteSourceAddress *string `tfsdk:"device_route_source_address" yaml:"deviceRouteSourceAddress,omitempty"`

		DeviceRouteSourceAddressIPv6 *string `tfsdk:"device_route_source_address_i_pv6" yaml:"deviceRouteSourceAddressIPv6,omitempty"`

		DisableConntrackInvalidCheck *bool `tfsdk:"disable_conntrack_invalid_check" yaml:"disableConntrackInvalidCheck,omitempty"`

		EndpointReportingDelay *string `tfsdk:"endpoint_reporting_delay" yaml:"endpointReportingDelay,omitempty"`

		EndpointReportingEnabled *bool `tfsdk:"endpoint_reporting_enabled" yaml:"endpointReportingEnabled,omitempty"`

		ExternalNodesList *[]string `tfsdk:"external_nodes_list" yaml:"externalNodesList,omitempty"`

		FailsafeInboundHostPorts *[]struct {
			Net *string `tfsdk:"net" yaml:"net,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
		} `tfsdk:"failsafe_inbound_host_ports" yaml:"failsafeInboundHostPorts,omitempty"`

		FailsafeOutboundHostPorts *[]struct {
			Net *string `tfsdk:"net" yaml:"net,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
		} `tfsdk:"failsafe_outbound_host_ports" yaml:"failsafeOutboundHostPorts,omitempty"`

		FeatureDetectOverride *string `tfsdk:"feature_detect_override" yaml:"featureDetectOverride,omitempty"`

		FeatureGates *string `tfsdk:"feature_gates" yaml:"featureGates,omitempty"`

		FloatingIPs *string `tfsdk:"floating_i_ps" yaml:"floatingIPs,omitempty"`

		GenericXDPEnabled *bool `tfsdk:"generic_xdp_enabled" yaml:"genericXDPEnabled,omitempty"`

		HealthEnabled *bool `tfsdk:"health_enabled" yaml:"healthEnabled,omitempty"`

		HealthHost *string `tfsdk:"health_host" yaml:"healthHost,omitempty"`

		HealthPort *int64 `tfsdk:"health_port" yaml:"healthPort,omitempty"`

		InterfaceExclude *string `tfsdk:"interface_exclude" yaml:"interfaceExclude,omitempty"`

		InterfacePrefix *string `tfsdk:"interface_prefix" yaml:"interfacePrefix,omitempty"`

		InterfaceRefreshInterval *string `tfsdk:"interface_refresh_interval" yaml:"interfaceRefreshInterval,omitempty"`

		IpipEnabled *bool `tfsdk:"ipip_enabled" yaml:"ipipEnabled,omitempty"`

		IpipMTU *int64 `tfsdk:"ipip_mtu" yaml:"ipipMTU,omitempty"`

		IpsetsRefreshInterval *string `tfsdk:"ipsets_refresh_interval" yaml:"ipsetsRefreshInterval,omitempty"`

		IptablesBackend *string `tfsdk:"iptables_backend" yaml:"iptablesBackend,omitempty"`

		IptablesFilterAllowAction *string `tfsdk:"iptables_filter_allow_action" yaml:"iptablesFilterAllowAction,omitempty"`

		IptablesLockFilePath *string `tfsdk:"iptables_lock_file_path" yaml:"iptablesLockFilePath,omitempty"`

		IptablesLockProbeInterval *string `tfsdk:"iptables_lock_probe_interval" yaml:"iptablesLockProbeInterval,omitempty"`

		IptablesLockTimeout *string `tfsdk:"iptables_lock_timeout" yaml:"iptablesLockTimeout,omitempty"`

		IptablesMangleAllowAction *string `tfsdk:"iptables_mangle_allow_action" yaml:"iptablesMangleAllowAction,omitempty"`

		IptablesMarkMask *int64 `tfsdk:"iptables_mark_mask" yaml:"iptablesMarkMask,omitempty"`

		IptablesNATOutgoingInterfaceFilter *string `tfsdk:"iptables_nat_outgoing_interface_filter" yaml:"iptablesNATOutgoingInterfaceFilter,omitempty"`

		IptablesPostWriteCheckInterval *string `tfsdk:"iptables_post_write_check_interval" yaml:"iptablesPostWriteCheckInterval,omitempty"`

		IptablesRefreshInterval *string `tfsdk:"iptables_refresh_interval" yaml:"iptablesRefreshInterval,omitempty"`

		Ipv6Support *bool `tfsdk:"ipv6_support" yaml:"ipv6Support,omitempty"`

		KubeNodePortRanges *[]string `tfsdk:"kube_node_port_ranges" yaml:"kubeNodePortRanges,omitempty"`

		LogDebugFilenameRegex *string `tfsdk:"log_debug_filename_regex" yaml:"logDebugFilenameRegex,omitempty"`

		LogFilePath *string `tfsdk:"log_file_path" yaml:"logFilePath,omitempty"`

		LogPrefix *string `tfsdk:"log_prefix" yaml:"logPrefix,omitempty"`

		LogSeverityFile *string `tfsdk:"log_severity_file" yaml:"logSeverityFile,omitempty"`

		LogSeverityScreen *string `tfsdk:"log_severity_screen" yaml:"logSeverityScreen,omitempty"`

		LogSeveritySys *string `tfsdk:"log_severity_sys" yaml:"logSeveritySys,omitempty"`

		MaxIpsetSize *int64 `tfsdk:"max_ipset_size" yaml:"maxIpsetSize,omitempty"`

		MetadataAddr *string `tfsdk:"metadata_addr" yaml:"metadataAddr,omitempty"`

		MetadataPort *int64 `tfsdk:"metadata_port" yaml:"metadataPort,omitempty"`

		MtuIfacePattern *string `tfsdk:"mtu_iface_pattern" yaml:"mtuIfacePattern,omitempty"`

		NatOutgoingAddress *string `tfsdk:"nat_outgoing_address" yaml:"natOutgoingAddress,omitempty"`

		NatPortRange utilities.IntOrString `tfsdk:"nat_port_range" yaml:"natPortRange,omitempty"`

		NetlinkTimeout *string `tfsdk:"netlink_timeout" yaml:"netlinkTimeout,omitempty"`

		OpenstackRegion *string `tfsdk:"openstack_region" yaml:"openstackRegion,omitempty"`

		PolicySyncPathPrefix *string `tfsdk:"policy_sync_path_prefix" yaml:"policySyncPathPrefix,omitempty"`

		PrometheusGoMetricsEnabled *bool `tfsdk:"prometheus_go_metrics_enabled" yaml:"prometheusGoMetricsEnabled,omitempty"`

		PrometheusMetricsEnabled *bool `tfsdk:"prometheus_metrics_enabled" yaml:"prometheusMetricsEnabled,omitempty"`

		PrometheusMetricsHost *string `tfsdk:"prometheus_metrics_host" yaml:"prometheusMetricsHost,omitempty"`

		PrometheusMetricsPort *int64 `tfsdk:"prometheus_metrics_port" yaml:"prometheusMetricsPort,omitempty"`

		PrometheusProcessMetricsEnabled *bool `tfsdk:"prometheus_process_metrics_enabled" yaml:"prometheusProcessMetricsEnabled,omitempty"`

		PrometheusWireGuardMetricsEnabled *bool `tfsdk:"prometheus_wire_guard_metrics_enabled" yaml:"prometheusWireGuardMetricsEnabled,omitempty"`

		RemoveExternalRoutes *bool `tfsdk:"remove_external_routes" yaml:"removeExternalRoutes,omitempty"`

		ReportingInterval *string `tfsdk:"reporting_interval" yaml:"reportingInterval,omitempty"`

		ReportingTTL *string `tfsdk:"reporting_ttl" yaml:"reportingTTL,omitempty"`

		RouteRefreshInterval *string `tfsdk:"route_refresh_interval" yaml:"routeRefreshInterval,omitempty"`

		RouteSource *string `tfsdk:"route_source" yaml:"routeSource,omitempty"`

		RouteSyncDisabled *bool `tfsdk:"route_sync_disabled" yaml:"routeSyncDisabled,omitempty"`

		RouteTableRange *struct {
			Max *int64 `tfsdk:"max" yaml:"max,omitempty"`

			Min *int64 `tfsdk:"min" yaml:"min,omitempty"`
		} `tfsdk:"route_table_range" yaml:"routeTableRange,omitempty"`

		RouteTableRanges *[]struct {
			Max *int64 `tfsdk:"max" yaml:"max,omitempty"`

			Min *int64 `tfsdk:"min" yaml:"min,omitempty"`
		} `tfsdk:"route_table_ranges" yaml:"routeTableRanges,omitempty"`

		ServiceLoopPrevention *string `tfsdk:"service_loop_prevention" yaml:"serviceLoopPrevention,omitempty"`

		SidecarAccelerationEnabled *bool `tfsdk:"sidecar_acceleration_enabled" yaml:"sidecarAccelerationEnabled,omitempty"`

		UsageReportingEnabled *bool `tfsdk:"usage_reporting_enabled" yaml:"usageReportingEnabled,omitempty"`

		UsageReportingInitialDelay *string `tfsdk:"usage_reporting_initial_delay" yaml:"usageReportingInitialDelay,omitempty"`

		UsageReportingInterval *string `tfsdk:"usage_reporting_interval" yaml:"usageReportingInterval,omitempty"`

		UseInternalDataplaneDriver *bool `tfsdk:"use_internal_dataplane_driver" yaml:"useInternalDataplaneDriver,omitempty"`

		VxlanEnabled *bool `tfsdk:"vxlan_enabled" yaml:"vxlanEnabled,omitempty"`

		VxlanMTU *int64 `tfsdk:"vxlan_mtu" yaml:"vxlanMTU,omitempty"`

		VxlanMTUV6 *int64 `tfsdk:"vxlan_mtuv6" yaml:"vxlanMTUV6,omitempty"`

		VxlanPort *int64 `tfsdk:"vxlan_port" yaml:"vxlanPort,omitempty"`

		VxlanVNI *int64 `tfsdk:"vxlan_vni" yaml:"vxlanVNI,omitempty"`

		WireguardEnabled *bool `tfsdk:"wireguard_enabled" yaml:"wireguardEnabled,omitempty"`

		WireguardEnabledV6 *bool `tfsdk:"wireguard_enabled_v6" yaml:"wireguardEnabledV6,omitempty"`

		WireguardHostEncryptionEnabled *bool `tfsdk:"wireguard_host_encryption_enabled" yaml:"wireguardHostEncryptionEnabled,omitempty"`

		WireguardInterfaceName *string `tfsdk:"wireguard_interface_name" yaml:"wireguardInterfaceName,omitempty"`

		WireguardInterfaceNameV6 *string `tfsdk:"wireguard_interface_name_v6" yaml:"wireguardInterfaceNameV6,omitempty"`

		WireguardKeepAlive *string `tfsdk:"wireguard_keep_alive" yaml:"wireguardKeepAlive,omitempty"`

		WireguardListeningPort *int64 `tfsdk:"wireguard_listening_port" yaml:"wireguardListeningPort,omitempty"`

		WireguardListeningPortV6 *int64 `tfsdk:"wireguard_listening_port_v6" yaml:"wireguardListeningPortV6,omitempty"`

		WireguardMTU *int64 `tfsdk:"wireguard_mtu" yaml:"wireguardMTU,omitempty"`

		WireguardMTUV6 *int64 `tfsdk:"wireguard_mtuv6" yaml:"wireguardMTUV6,omitempty"`

		WireguardRoutingRulePriority *int64 `tfsdk:"wireguard_routing_rule_priority" yaml:"wireguardRoutingRulePriority,omitempty"`

		WorkloadSourceSpoofing *string `tfsdk:"workload_source_spoofing" yaml:"workloadSourceSpoofing,omitempty"`

		XdpEnabled *bool `tfsdk:"xdp_enabled" yaml:"xdpEnabled,omitempty"`

		XdpRefreshInterval *string `tfsdk:"xdp_refresh_interval" yaml:"xdpRefreshInterval,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCrdProjectcalicoOrgFelixConfigurationV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgFelixConfigurationV1Resource{}
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crd_projectcalico_org_felix_configuration_v1"
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Felix Configuration contains the configuration for Felix.",
		MarkdownDescription: "Felix Configuration contains the configuration for Felix.",
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
				Description:         "FelixConfigurationSpec contains the values of the Felix configuration.",
				MarkdownDescription: "FelixConfigurationSpec contains the values of the Felix configuration.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"allow_ipip_packets_from_workloads": {
						Description:         "AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic from workloads [Default: false]",
						MarkdownDescription: "AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic from workloads [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"allow_vxlan_packets_from_workloads": {
						Description:         "AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic from workloads [Default: false]",
						MarkdownDescription: "AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic from workloads [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"aws_src_dst_check": {
						Description:         "Set source-destination-check on AWS EC2 instances. Accepted value must be one of 'DoNothing', 'Enable' or 'Disable'. [Default: DoNothing]",
						MarkdownDescription: "Set source-destination-check on AWS EC2 instances. Accepted value must be one of 'DoNothing', 'Enable' or 'Disable'. [Default: DoNothing]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("DoNothing", "Enable", "Disable"),
						},
					},

					"bpf_connect_time_load_balancing_enabled": {
						Description:         "BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load balancer.  The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections.  The only reason to disable it is for debugging purposes.  [Default: true]",
						MarkdownDescription: "BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load balancer.  The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections.  The only reason to disable it is for debugging purposes.  [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_data_iface_pattern": {
						Description:         "BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to in order to catch traffic to/from the network.  This needs to match the interfaces that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.  It should not match the workload interfaces (usually named cali...).",
						MarkdownDescription: "BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to in order to catch traffic to/from the network.  This needs to match the interfaces that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.  It should not match the workload interfaces (usually named cali...).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_disable_unprivileged": {
						Description:         "BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable unprivileged use of BPF.  This ensures that unprivileged users cannot access Calico's BPF maps and cannot insert their own BPF programs to interfere with Calico's. [Default: true]",
						MarkdownDescription: "BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable unprivileged use of BPF.  This ensures that unprivileged users cannot access Calico's BPF maps and cannot insert their own BPF programs to interfere with Calico's. [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_enabled": {
						Description:         "BPFEnabled, if enabled Felix will use the BPF dataplane. [Default: false]",
						MarkdownDescription: "BPFEnabled, if enabled Felix will use the BPF dataplane. [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_enforce_rpf": {
						Description:         "BPFEnforceRPF enforce strict RPF on all interfaces with BPF programs regardless of what is the per-interfaces or global setting. Possible values are Disabled or Strict. [Default: Strict]",
						MarkdownDescription: "BPFEnforceRPF enforce strict RPF on all interfaces with BPF programs regardless of what is the per-interfaces or global setting. Possible values are Disabled or Strict. [Default: Strict]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_ext_to_service_connmark": {
						Description:         "BPFExtToServiceConnmark in BPF mode, control a 32bit mark that is set on connections from an external client to a local service. This mark allows us to control how packets of that connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]",
						MarkdownDescription: "BPFExtToServiceConnmark in BPF mode, control a 32bit mark that is set on connections from an external client to a local service. This mark allows us to control how packets of that connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_external_service_mode": {
						Description:         "BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports and cluster IPs) are forwarded to remote workloads.  If set to 'Tunnel' then both request and response traffic is tunneled to the remote node.  If set to 'DSR', the request traffic is tunneled but the response traffic is sent directly from the remote node.  In 'DSR' mode, the remote node appears to use the IP of the ingress node; this requires a permissive L2 network.  [Default: Tunnel]",
						MarkdownDescription: "BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports and cluster IPs) are forwarded to remote workloads.  If set to 'Tunnel' then both request and response traffic is tunneled to the remote node.  If set to 'DSR', the request traffic is tunneled but the response traffic is sent directly from the remote node.  In 'DSR' mode, the remote node appears to use the IP of the ingress node; this requires a permissive L2 network.  [Default: Tunnel]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_host_conntrack_bypass": {
						Description:         "BPFHostConntrackBypass Controls whether to bypass Linux conntrack in BPF mode for workloads and services. [Default: true - bypass Linux conntrack]",
						MarkdownDescription: "BPFHostConntrackBypass Controls whether to bypass Linux conntrack in BPF mode for workloads and services. [Default: true - bypass Linux conntrack]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_kube_proxy_endpoint_slices_enabled": {
						Description:         "BPFKubeProxyEndpointSlicesEnabled in BPF mode, controls whether Felix's embedded kube-proxy accepts EndpointSlices or not.",
						MarkdownDescription: "BPFKubeProxyEndpointSlicesEnabled in BPF mode, controls whether Felix's embedded kube-proxy accepts EndpointSlices or not.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_kube_proxy_iptables_cleanup_enabled": {
						Description:         "BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream Kubernetes kube-proxy's iptables chains.  Should only be enabled if kube-proxy is not running.  [Default: true]",
						MarkdownDescription: "BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream Kubernetes kube-proxy's iptables chains.  Should only be enabled if kube-proxy is not running.  [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_kube_proxy_min_sync_period": {
						Description:         "BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's embedded kube-proxy.  Lower values give reduced set-up latency.  Higher values reduce Felix CPU usage by batching up more work.  [Default: 1s]",
						MarkdownDescription: "BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's embedded kube-proxy.  Lower values give reduced set-up latency.  Higher values reduce Felix CPU usage by batching up more work.  [Default: 1s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_l3_iface_pattern": {
						Description:         "BPFL3IfacePattern is a regular expression that allows to list tunnel devices like wireguard or vxlan (i.e., L3 devices) in addition to BPFDataIfacePattern. That is, tunnel interfaces not created by Calico, that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.",
						MarkdownDescription: "BPFL3IfacePattern is a regular expression that allows to list tunnel devices like wireguard or vxlan (i.e., L3 devices) in addition to BPFDataIfacePattern. That is, tunnel interfaces not created by Calico, that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_log_level": {
						Description:         "BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode.  One of 'Off', 'Info', or 'Debug'.  The logs are emitted to the BPF trace pipe, accessible with the command 'tc exec bpf debug'. [Default: Off].",
						MarkdownDescription: "BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode.  One of 'Off', 'Info', or 'Debug'.  The logs are emitted to the BPF trace pipe, accessible with the command 'tc exec bpf debug'. [Default: Off].",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_map_size_conntrack": {
						Description:         "BPFMapSizeConntrack sets the size for the conntrack map.  This map must be large enough to hold an entry for each active connection.  Warning: changing the size of the conntrack map can cause disruption.",
						MarkdownDescription: "BPFMapSizeConntrack sets the size for the conntrack map.  This map must be large enough to hold an entry for each active connection.  Warning: changing the size of the conntrack map can cause disruption.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_map_size_ip_sets": {
						Description:         "BPFMapSizeIPSets sets the size for ipsets map.  The IP sets map must be large enough to hold an entry for each endpoint matched by every selector in the source/destination matches in network policy.  Selectors such as 'all()' can result in large numbers of entries (one entry per endpoint in that case).",
						MarkdownDescription: "BPFMapSizeIPSets sets the size for ipsets map.  The IP sets map must be large enough to hold an entry for each endpoint matched by every selector in the source/destination matches in network policy.  Selectors such as 'all()' can result in large numbers of entries (one entry per endpoint in that case).",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_map_size_if_state": {
						Description:         "BPFMapSizeIfState sets the size for ifstate map.  The ifstate map must be large enough to hold an entry for each device (host + workloads) on a host.",
						MarkdownDescription: "BPFMapSizeIfState sets the size for ifstate map.  The ifstate map must be large enough to hold an entry for each device (host + workloads) on a host.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_map_size_nat_affinity": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_map_size_nat_backend": {
						Description:         "BPFMapSizeNATBackend sets the size for nat back end map. This is the total number of endpoints. This is mostly more than the size of the number of services.",
						MarkdownDescription: "BPFMapSizeNATBackend sets the size for nat back end map. This is the total number of endpoints. This is mostly more than the size of the number of services.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_map_size_nat_frontend": {
						Description:         "BPFMapSizeNATFrontend sets the size for nat front end map. FrontendMap should be large enough to hold an entry for each nodeport, external IP and each port in each service.",
						MarkdownDescription: "BPFMapSizeNATFrontend sets the size for nat front end map. FrontendMap should be large enough to hold an entry for each nodeport, external IP and each port in each service.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_map_size_route": {
						Description:         "BPFMapSizeRoute sets the size for the routes map.  The routes map should be large enough to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and tunnel IPs).",
						MarkdownDescription: "BPFMapSizeRoute sets the size for the routes map.  The routes map should be large enough to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and tunnel IPs).",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_psnat_ports": {
						Description:         "BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision. This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the  ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]",
						MarkdownDescription: "BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision. This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the  ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]",

						Type: utilities.IntOrStringType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"bpf_policy_debug_enabled": {
						Description:         "BPFPolicyDebugEnabled when true, Felix records detailed information about the BPF policy programs, which can be examined with the calico-bpf command-line tool.",
						MarkdownDescription: "BPFPolicyDebugEnabled when true, Felix records detailed information about the BPF policy programs, which can be examined with the calico-bpf command-line tool.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"chain_insert_mode": {
						Description:         "ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule at the top of the chain or by appending a rule at the bottom. insert is the safe default since it prevents Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed. [Default: insert]",
						MarkdownDescription: "ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule at the top of the chain or by appending a rule at the bottom. insert is the safe default since it prevents Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed. [Default: insert]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dataplane_driver": {
						Description:         "DataplaneDriver filename of the external dataplane driver to use.  Only used if UseInternalDataplaneDriver is set to false.",
						MarkdownDescription: "DataplaneDriver filename of the external dataplane driver to use.  Only used if UseInternalDataplaneDriver is set to false.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dataplane_watchdog_timeout": {
						Description:         "DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver. Increase this value if you experience spurious non-ready or non-live events when Felix is under heavy load. Decrease the value to get felix to report non-live or non-ready more quickly. [Default: 90s]",
						MarkdownDescription: "DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver. Increase this value if you experience spurious non-ready or non-live events when Felix is under heavy load. Decrease the value to get felix to report non-live or non-ready more quickly. [Default: 90s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"debug_disable_log_dropping": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"debug_memory_profile_path": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"debug_simulate_calc_graph_hang_after": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"debug_simulate_dataplane_hang_after": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_endpoint_to_host_action": {
						Description:         "DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host itself (after the traffic hits the endpoint egress policy). By default Calico blocks traffic from workload endpoints to the host itself with an iptables 'DROP' action. If you want to allow some or all traffic from endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables 'INPUT' chain; Calico will insert its rules at the top of that chain, then 'RETURN' packets to the 'INPUT' chain once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets from workloads after processing workload endpoint egress policy. [Default: Drop]",
						MarkdownDescription: "DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host itself (after the traffic hits the endpoint egress policy). By default Calico blocks traffic from workload endpoints to the host itself with an iptables 'DROP' action. If you want to allow some or all traffic from endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables 'INPUT' chain; Calico will insert its rules at the top of that chain, then 'RETURN' packets to the 'INPUT' chain once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets from workloads after processing workload endpoint egress policy. [Default: Drop]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"device_route_protocol": {
						Description:         "This defines the route protocol added to programmed device routes, by default this will be RTPROT_BOOT when left blank.",
						MarkdownDescription: "This defines the route protocol added to programmed device routes, by default this will be RTPROT_BOOT when left blank.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"device_route_source_address": {
						Description:         "This is the IPv4 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",
						MarkdownDescription: "This is the IPv4 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"device_route_source_address_i_pv6": {
						Description:         "This is the IPv6 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",
						MarkdownDescription: "This is the IPv6 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_conntrack_invalid_check": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoint_reporting_delay": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoint_reporting_enabled": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_nodes_list": {
						Description:         "ExternalNodesCIDRList is a list of CIDR's of external-non-calico-nodes which may source tunnel traffic and have the tunneled traffic be accepted at calico nodes.",
						MarkdownDescription: "ExternalNodesCIDRList is a list of CIDR's of external-non-calico-nodes which may source tunnel traffic and have the tunneled traffic be accepted at calico nodes.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"failsafe_inbound_host_ports": {
						Description:         "FailsafeInboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow incoming traffic to host endpoints on irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports, use the value none. The default value allows ssh access and DHCP. [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667]",
						MarkdownDescription: "FailsafeInboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow incoming traffic to host endpoints on irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports, use the value none. The default value allows ssh access and DHCP. [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667]",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"net": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"protocol": {
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

					"failsafe_outbound_host_ports": {
						Description:         "FailsafeOutboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow outgoing traffic from host endpoints to irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports, use the value none. The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd as well as allowing DHCP and DNS. [Default: tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667, udp:53, udp:67]",
						MarkdownDescription: "FailsafeOutboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow outgoing traffic from host endpoints to irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports, use the value none. The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd as well as allowing DHCP and DNS. [Default: tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667, udp:53, udp:67]",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"net": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"protocol": {
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

					"feature_detect_override": {
						Description:         "FeatureDetectOverride is used to override feature detection based on auto-detected platform capabilities.  Values are specified in a comma separated list with no spaces, example; 'SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock='.  'true' or 'false' will force the feature, empty or omitted values are auto-detected.",
						MarkdownDescription: "FeatureDetectOverride is used to override feature detection based on auto-detected platform capabilities.  Values are specified in a comma separated list with no spaces, example; 'SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock='.  'true' or 'false' will force the feature, empty or omitted values are auto-detected.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"feature_gates": {
						Description:         "FeatureGates is used to enable or disable tech-preview Calico features. Values are specified in a comma separated list with no spaces, example; 'BPFConnectTimeLoadBalancingWorkaround=enabled,XyZ=false'. This is used to enable features that are not fully production ready.",
						MarkdownDescription: "FeatureGates is used to enable or disable tech-preview Calico features. Values are specified in a comma separated list with no spaces, example; 'BPFConnectTimeLoadBalancingWorkaround=enabled,XyZ=false'. This is used to enable features that are not fully production ready.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"floating_i_ps": {
						Description:         "FloatingIPs configures whether or not Felix will program non-OpenStack floating IP addresses.  (OpenStack-derived floating IPs are always programmed, regardless of this setting.)",
						MarkdownDescription: "FloatingIPs configures whether or not Felix will program non-OpenStack floating IP addresses.  (OpenStack-derived floating IPs are always programmed, regardless of this setting.)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"generic_xdp_enabled": {
						Description:         "GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver modes can use XDP. This is not recommended since it doesn't provide better performance than iptables. [Default: false]",
						MarkdownDescription: "GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver modes can use XDP. This is not recommended since it doesn't provide better performance than iptables. [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"health_enabled": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"health_host": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"health_port": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interface_exclude": {
						Description:         "InterfaceExclude is a comma-separated list of interfaces that Felix should exclude when monitoring for host endpoints. The default value ensures that Felix ignores Kubernetes' IPVS dummy interface, which is used internally by kube-proxy. If you want to exclude multiple interface names using a single value, the list supports regular expressions. For regular expressions you must wrap the value with '/'. For example having values '/^kube/,veth1' will exclude all interfaces that begin with 'kube' and also the interface 'veth1'. [Default: kube-ipvs0]",
						MarkdownDescription: "InterfaceExclude is a comma-separated list of interfaces that Felix should exclude when monitoring for host endpoints. The default value ensures that Felix ignores Kubernetes' IPVS dummy interface, which is used internally by kube-proxy. If you want to exclude multiple interface names using a single value, the list supports regular expressions. For regular expressions you must wrap the value with '/'. For example having values '/^kube/,veth1' will exclude all interfaces that begin with 'kube' and also the interface 'veth1'. [Default: kube-ipvs0]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interface_prefix": {
						Description:         "InterfacePrefix is the interface name prefix that identifies workload endpoints and so distinguishes them from host endpoint interfaces. Note: in environments other than bare metal, the orchestrators configure this appropriately. For example our Kubernetes and Docker integrations set the 'cali' value, and our OpenStack integration sets the 'tap' value. [Default: cali]",
						MarkdownDescription: "InterfacePrefix is the interface name prefix that identifies workload endpoints and so distinguishes them from host endpoint interfaces. Note: in environments other than bare metal, the orchestrators configure this appropriately. For example our Kubernetes and Docker integrations set the 'cali' value, and our OpenStack integration sets the 'tap' value. [Default: cali]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"interface_refresh_interval": {
						Description:         "InterfaceRefreshInterval is the period at which Felix rescans local interfaces to verify their state. The rescan can be disabled by setting the interval to 0.",
						MarkdownDescription: "InterfaceRefreshInterval is the period at which Felix rescans local interfaces to verify their state. The rescan can be disabled by setting the interval to 0.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipip_enabled": {
						Description:         "IPIPEnabled overrides whether Felix should configure an IPIP interface on the host. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						MarkdownDescription: "IPIPEnabled overrides whether Felix should configure an IPIP interface on the host. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipip_mtu": {
						Description:         "IPIPMTU is the MTU to set on the tunnel device. See Configuring MTU [Default: 1440]",
						MarkdownDescription: "IPIPMTU is the MTU to set on the tunnel device. See Configuring MTU [Default: 1440]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipsets_refresh_interval": {
						Description:         "IpsetsRefreshInterval is the period at which Felix re-checks all iptables state to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable iptables refresh. [Default: 90s]",
						MarkdownDescription: "IpsetsRefreshInterval is the period at which Felix re-checks all iptables state to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable iptables refresh. [Default: 90s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_backend": {
						Description:         "IptablesBackend specifies which backend of iptables will be used. The default is legacy.",
						MarkdownDescription: "IptablesBackend specifies which backend of iptables will be used. The default is legacy.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_filter_allow_action": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_lock_file_path": {
						Description:         "IptablesLockFilePath is the location of the iptables lock file. You may need to change this if the lock file is not in its standard location (for example if you have mapped it into Felix's container at a different path). [Default: /run/xtables.lock]",
						MarkdownDescription: "IptablesLockFilePath is the location of the iptables lock file. You may need to change this if the lock file is not in its standard location (for example if you have mapped it into Felix's container at a different path). [Default: /run/xtables.lock]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_lock_probe_interval": {
						Description:         "IptablesLockProbeInterval is the time that Felix will wait between attempts to acquire the iptables lock if it is not available. Lower values make Felix more responsive when the lock is contended, but use more CPU. [Default: 50ms]",
						MarkdownDescription: "IptablesLockProbeInterval is the time that Felix will wait between attempts to acquire the iptables lock if it is not available. Lower values make Felix more responsive when the lock is contended, but use more CPU. [Default: 50ms]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_lock_timeout": {
						Description:         "IptablesLockTimeout is the time that Felix will wait for the iptables lock, or 0, to disable. To use this feature, Felix must share the iptables lock file with all other processes that also take the lock. When running Felix inside a container, this requires the /run directory of the host to be mounted into the calico/node or calico/felix container. [Default: 0s disabled]",
						MarkdownDescription: "IptablesLockTimeout is the time that Felix will wait for the iptables lock, or 0, to disable. To use this feature, Felix must share the iptables lock file with all other processes that also take the lock. When running Felix inside a container, this requires the /run directory of the host to be mounted into the calico/node or calico/felix container. [Default: 0s disabled]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_mangle_allow_action": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_mark_mask": {
						Description:         "IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xff000000]",
						MarkdownDescription: "IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xff000000]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_nat_outgoing_interface_filter": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_post_write_check_interval": {
						Description:         "IptablesPostWriteCheckInterval is the period after Felix has done a write to the dataplane that it schedules an extra read back in order to check the write was not clobbered by another process. This should only occur if another application on the system doesn't respect the iptables lock. [Default: 1s]",
						MarkdownDescription: "IptablesPostWriteCheckInterval is the period after Felix has done a write to the dataplane that it schedules an extra read back in order to check the write was not clobbered by another process. This should only occur if another application on the system doesn't respect the iptables lock. [Default: 1s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iptables_refresh_interval": {
						Description:         "IptablesRefreshInterval is the period at which Felix re-checks the IP sets in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable IP sets refresh. Note: the default for this value is lower than the other refresh intervals as a workaround for a Linux kernel bug that was fixed in kernel version 4.11. If you are using v4.11 or greater you may want to set this to, a higher value to reduce Felix CPU usage. [Default: 10s]",
						MarkdownDescription: "IptablesRefreshInterval is the period at which Felix re-checks the IP sets in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable IP sets refresh. Note: the default for this value is lower than the other refresh intervals as a workaround for a Linux kernel bug that was fixed in kernel version 4.11. If you are using v4.11 or greater you may want to set this to, a higher value to reduce Felix CPU usage. [Default: 10s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipv6_support": {
						Description:         "IPv6Support controls whether Felix enables support for IPv6 (if supported by the in-use dataplane).",
						MarkdownDescription: "IPv6Support controls whether Felix enables support for IPv6 (if supported by the in-use dataplane).",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kube_node_port_ranges": {
						Description:         "KubeNodePortRanges holds list of port ranges used for service node ports. Only used if felix detects kube-proxy running in ipvs mode. Felix uses these ranges to separate host and workload traffic. [Default: 30000:32767].",
						MarkdownDescription: "KubeNodePortRanges holds list of port ranges used for service node ports. Only used if felix detects kube-proxy running in ipvs mode. Felix uses these ranges to separate host and workload traffic. [Default: 30000:32767].",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_debug_filename_regex": {
						Description:         "LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs. Only logs from files with names that match the given regular expression are included.  The filter only applies to Debug level logs.",
						MarkdownDescription: "LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs. Only logs from files with names that match the given regular expression are included.  The filter only applies to Debug level logs.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_file_path": {
						Description:         "LogFilePath is the full path to the Felix log. Set to none to disable file logging. [Default: /var/log/calico/felix.log]",
						MarkdownDescription: "LogFilePath is the full path to the Felix log. Set to none to disable file logging. [Default: /var/log/calico/felix.log]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_prefix": {
						Description:         "LogPrefix is the log prefix that Felix uses when rendering LOG rules. [Default: calico-packet]",
						MarkdownDescription: "LogPrefix is the log prefix that Felix uses when rendering LOG rules. [Default: calico-packet]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_severity_file": {
						Description:         "LogSeverityFile is the log severity above which logs are sent to the log file. [Default: Info]",
						MarkdownDescription: "LogSeverityFile is the log severity above which logs are sent to the log file. [Default: Info]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_severity_screen": {
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_severity_sys": {
						Description:         "LogSeveritySys is the log severity above which logs are sent to the syslog. Set to None for no logging to syslog. [Default: Info]",
						MarkdownDescription: "LogSeveritySys is the log severity above which logs are sent to the syslog. Set to None for no logging to syslog. [Default: Info]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_ipset_size": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metadata_addr": {
						Description:         "MetadataAddr is the IP address or domain name of the server that can answer VM queries for cloud-init metadata. In OpenStack, this corresponds to the machine running nova-api (or in Ubuntu, nova-api-metadata). A value of none (case insensitive) means that Felix should not set up any NAT rule for the metadata path. [Default: 127.0.0.1]",
						MarkdownDescription: "MetadataAddr is the IP address or domain name of the server that can answer VM queries for cloud-init metadata. In OpenStack, this corresponds to the machine running nova-api (or in Ubuntu, nova-api-metadata). A value of none (case insensitive) means that Felix should not set up any NAT rule for the metadata path. [Default: 127.0.0.1]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metadata_port": {
						Description:         "MetadataPort is the port of the metadata server. This, combined with global.MetadataAddr (if not 'None'), is used to set up a NAT rule, from 169.254.169.254:80 to MetadataAddr:MetadataPort. In most cases this should not need to be changed [Default: 8775].",
						MarkdownDescription: "MetadataPort is the port of the metadata server. This, combined with global.MetadataAddr (if not 'None'), is used to set up a NAT rule, from 169.254.169.254:80 to MetadataAddr:MetadataPort. In most cases this should not need to be changed [Default: 8775].",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mtu_iface_pattern": {
						Description:         "MTUIfacePattern is a regular expression that controls which interfaces Felix should scan in order to calculate the host's MTU. This should not match workload interfaces (usually named cali...).",
						MarkdownDescription: "MTUIfacePattern is a regular expression that controls which interfaces Felix should scan in order to calculate the host's MTU. This should not match workload interfaces (usually named cali...).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nat_outgoing_address": {
						Description:         "NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that is leaving the network. By default the address used is an address on the interface the traffic is leaving on (ie it uses the iptables MASQUERADE target)",
						MarkdownDescription: "NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that is leaving the network. By default the address used is an address on the interface the traffic is leaving on (ie it uses the iptables MASQUERADE target)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nat_port_range": {
						Description:         "NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT. When unset the default behavior of the network stack is used.",
						MarkdownDescription: "NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT. When unset the default behavior of the network stack is used.",

						Type: utilities.IntOrStringType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"netlink_timeout": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"openstack_region": {
						Description:         "OpenstackRegion is the name of the region that a particular Felix belongs to. In a multi-region Calico/OpenStack deployment, this must be configured somehow for each Felix (here in the datamodel, or in felix.cfg or the environment on each compute node), and must match the [calico] openstack_region value configured in neutron.conf on each node. [Default: Empty]",
						MarkdownDescription: "OpenstackRegion is the name of the region that a particular Felix belongs to. In a multi-region Calico/OpenStack deployment, this must be configured somehow for each Felix (here in the datamodel, or in felix.cfg or the environment on each compute node), and must match the [calico] openstack_region value configured in neutron.conf on each node. [Default: Empty]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"policy_sync_path_prefix": {
						Description:         "PolicySyncPathPrefix is used to by Felix to communicate policy changes to external services, like Application layer policy. [Default: Empty]",
						MarkdownDescription: "PolicySyncPathPrefix is used to by Felix to communicate policy changes to external services, like Application layer policy. [Default: Empty]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_go_metrics_enabled": {
						Description:         "PrometheusGoMetricsEnabled disables Go runtime metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						MarkdownDescription: "PrometheusGoMetricsEnabled disables Go runtime metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_metrics_enabled": {
						Description:         "PrometheusMetricsEnabled enables the Prometheus metrics server in Felix if set to true. [Default: false]",
						MarkdownDescription: "PrometheusMetricsEnabled enables the Prometheus metrics server in Felix if set to true. [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_metrics_host": {
						Description:         "PrometheusMetricsHost is the host that the Prometheus metrics server should bind to. [Default: empty]",
						MarkdownDescription: "PrometheusMetricsHost is the host that the Prometheus metrics server should bind to. [Default: empty]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_metrics_port": {
						Description:         "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. [Default: 9091]",
						MarkdownDescription: "PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to. [Default: 9091]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_process_metrics_enabled": {
						Description:         "PrometheusProcessMetricsEnabled disables process metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						MarkdownDescription: "PrometheusProcessMetricsEnabled disables process metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_wire_guard_metrics_enabled": {
						Description:         "PrometheusWireGuardMetricsEnabled disables wireguard metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",
						MarkdownDescription: "PrometheusWireGuardMetricsEnabled disables wireguard metrics collection, which the Prometheus client does by default, when set to false. This reduces the number of metrics reported, reducing Prometheus load. [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"remove_external_routes": {
						Description:         "Whether or not to remove device routes that have not been programmed by Felix. Disabling this will allow external applications to also add device routes. This is enabled by default which means we will remove externally added routes.",
						MarkdownDescription: "Whether or not to remove device routes that have not been programmed by Felix. Disabling this will allow external applications to also add device routes. This is enabled by default which means we will remove externally added routes.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"reporting_interval": {
						Description:         "ReportingInterval is the interval at which Felix reports its status into the datastore or 0 to disable. Must be non-zero in OpenStack deployments. [Default: 30s]",
						MarkdownDescription: "ReportingInterval is the interval at which Felix reports its status into the datastore or 0 to disable. Must be non-zero in OpenStack deployments. [Default: 30s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"reporting_ttl": {
						Description:         "ReportingTTL is the time-to-live setting for process-wide status reports. [Default: 90s]",
						MarkdownDescription: "ReportingTTL is the time-to-live setting for process-wide status reports. [Default: 90s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_refresh_interval": {
						Description:         "RouteRefreshInterval is the period at which Felix re-checks the routes in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable route refresh. [Default: 90s]",
						MarkdownDescription: "RouteRefreshInterval is the period at which Felix re-checks the routes in the dataplane to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable route refresh. [Default: 90s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_source": {
						Description:         "RouteSource configures where Felix gets its routing information. - WorkloadIPs: use workload endpoints to construct routes. - CalicoIPAM: the default - use IPAM data to construct routes.",
						MarkdownDescription: "RouteSource configures where Felix gets its routing information. - WorkloadIPs: use workload endpoints to construct routes. - CalicoIPAM: the default - use IPAM data to construct routes.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_sync_disabled": {
						Description:         "RouteSyncDisabled will disable all operations performed on the route table. Set to true to run in network-policy mode only.",
						MarkdownDescription: "RouteSyncDisabled will disable all operations performed on the route table. Set to true to run in network-policy mode only.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_table_range": {
						Description:         "Deprecated in favor of RouteTableRanges. Calico programs additional Linux route tables for various purposes. RouteTableRange specifies the indices of the route tables that Calico should use.",
						MarkdownDescription: "Deprecated in favor of RouteTableRanges. Calico programs additional Linux route tables for various purposes. RouteTableRange specifies the indices of the route tables that Calico should use.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"min": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_table_ranges": {
						Description:         "Calico programs additional Linux route tables for various purposes. RouteTableRanges specifies a set of table index ranges that Calico should use. Deprecates'RouteTableRange', overrides 'RouteTableRange'.",
						MarkdownDescription: "Calico programs additional Linux route tables for various purposes. RouteTableRanges specifies a set of table index ranges that Calico should use. Deprecates'RouteTableRange', overrides 'RouteTableRange'.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"max": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"min": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_loop_prevention": {
						Description:         "When service IP advertisement is enabled, prevent routing loops to service IPs that are not in use, by dropping or rejecting packets that do not get DNAT'd by kube-proxy. Unless set to 'Disabled', in which case such routing loops continue to be allowed. [Default: Drop]",
						MarkdownDescription: "When service IP advertisement is enabled, prevent routing loops to service IPs that are not in use, by dropping or rejecting packets that do not get DNAT'd by kube-proxy. Unless set to 'Disabled', in which case such routing loops continue to be allowed. [Default: Drop]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sidecar_acceleration_enabled": {
						Description:         "SidecarAccelerationEnabled enables experimental sidecar acceleration [Default: false]",
						MarkdownDescription: "SidecarAccelerationEnabled enables experimental sidecar acceleration [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"usage_reporting_enabled": {
						Description:         "UsageReportingEnabled reports anonymous Calico version number and cluster size to projectcalico.org. Logs warnings returned by the usage server. For example, if a significant security vulnerability has been discovered in the version of Calico being used. [Default: true]",
						MarkdownDescription: "UsageReportingEnabled reports anonymous Calico version number and cluster size to projectcalico.org. Logs warnings returned by the usage server. For example, if a significant security vulnerability has been discovered in the version of Calico being used. [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"usage_reporting_initial_delay": {
						Description:         "UsageReportingInitialDelay controls the minimum delay before Felix makes a report. [Default: 300s]",
						MarkdownDescription: "UsageReportingInitialDelay controls the minimum delay before Felix makes a report. [Default: 300s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"usage_reporting_interval": {
						Description:         "UsageReportingInterval controls the interval at which Felix makes reports. [Default: 86400s]",
						MarkdownDescription: "UsageReportingInterval controls the interval at which Felix makes reports. [Default: 86400s]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"use_internal_dataplane_driver": {
						Description:         "UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic.  If false, it will launch an external dataplane driver and communicate with it over protobuf.",
						MarkdownDescription: "UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic.  If false, it will launch an external dataplane driver and communicate with it over protobuf.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vxlan_enabled": {
						Description:         "VXLANEnabled overrides whether Felix should create the VXLAN tunnel device for IPv4 VXLAN networking. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						MarkdownDescription: "VXLANEnabled overrides whether Felix should create the VXLAN tunnel device for IPv4 VXLAN networking. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vxlan_mtu": {
						Description:         "VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device. See Configuring MTU [Default: 1410]",
						MarkdownDescription: "VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device. See Configuring MTU [Default: 1410]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vxlan_mtuv6": {
						Description:         "VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device. See Configuring MTU [Default: 1390]",
						MarkdownDescription: "VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device. See Configuring MTU [Default: 1390]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vxlan_port": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vxlan_vni": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_enabled": {
						Description:         "WireguardEnabled controls whether Wireguard is enabled for IPv4 (encapsulating IPv4 traffic over an IPv4 underlay network). [Default: false]",
						MarkdownDescription: "WireguardEnabled controls whether Wireguard is enabled for IPv4 (encapsulating IPv4 traffic over an IPv4 underlay network). [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_enabled_v6": {
						Description:         "WireguardEnabledV6 controls whether Wireguard is enabled for IPv6 (encapsulating IPv6 traffic over an IPv6 underlay network). [Default: false]",
						MarkdownDescription: "WireguardEnabledV6 controls whether Wireguard is enabled for IPv6 (encapsulating IPv6 traffic over an IPv6 underlay network). [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_host_encryption_enabled": {
						Description:         "WireguardHostEncryptionEnabled controls whether Wireguard host-to-host encryption is enabled. [Default: false]",
						MarkdownDescription: "WireguardHostEncryptionEnabled controls whether Wireguard host-to-host encryption is enabled. [Default: false]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_interface_name": {
						Description:         "WireguardInterfaceName specifies the name to use for the IPv4 Wireguard interface. [Default: wireguard.cali]",
						MarkdownDescription: "WireguardInterfaceName specifies the name to use for the IPv4 Wireguard interface. [Default: wireguard.cali]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_interface_name_v6": {
						Description:         "WireguardInterfaceNameV6 specifies the name to use for the IPv6 Wireguard interface. [Default: wg-v6.cali]",
						MarkdownDescription: "WireguardInterfaceNameV6 specifies the name to use for the IPv6 Wireguard interface. [Default: wg-v6.cali]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_keep_alive": {
						Description:         "WireguardKeepAlive controls Wireguard PersistentKeepalive option. Set 0 to disable. [Default: 0]",
						MarkdownDescription: "WireguardKeepAlive controls Wireguard PersistentKeepalive option. Set 0 to disable. [Default: 0]",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_listening_port": {
						Description:         "WireguardListeningPort controls the listening port used by IPv4 Wireguard. [Default: 51820]",
						MarkdownDescription: "WireguardListeningPort controls the listening port used by IPv4 Wireguard. [Default: 51820]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_listening_port_v6": {
						Description:         "WireguardListeningPortV6 controls the listening port used by IPv6 Wireguard. [Default: 51821]",
						MarkdownDescription: "WireguardListeningPortV6 controls the listening port used by IPv6 Wireguard. [Default: 51821]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_mtu": {
						Description:         "WireguardMTU controls the MTU on the IPv4 Wireguard interface. See Configuring MTU [Default: 1440]",
						MarkdownDescription: "WireguardMTU controls the MTU on the IPv4 Wireguard interface. See Configuring MTU [Default: 1440]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_mtuv6": {
						Description:         "WireguardMTUV6 controls the MTU on the IPv6 Wireguard interface. See Configuring MTU [Default: 1420]",
						MarkdownDescription: "WireguardMTUV6 controls the MTU on the IPv6 Wireguard interface. See Configuring MTU [Default: 1420]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"wireguard_routing_rule_priority": {
						Description:         "WireguardRoutingRulePriority controls the priority value to use for the Wireguard routing rule. [Default: 99]",
						MarkdownDescription: "WireguardRoutingRulePriority controls the priority value to use for the Wireguard routing rule. [Default: 99]",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"workload_source_spoofing": {
						Description:         "WorkloadSourceSpoofing controls whether pods can use the allowedSourcePrefixes annotation to send traffic with a source IP address that is not theirs. This is disabled by default. When set to 'Any', pods can request any prefix.",
						MarkdownDescription: "WorkloadSourceSpoofing controls whether pods can use the allowedSourcePrefixes annotation to send traffic with a source IP address that is not theirs. This is disabled by default. When set to 'Any', pods can request any prefix.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"xdp_enabled": {
						Description:         "XDPEnabled enables XDP acceleration for suitable untracked incoming deny rules. [Default: true]",
						MarkdownDescription: "XDPEnabled enables XDP acceleration for suitable untracked incoming deny rules. [Default: true]",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"xdp_refresh_interval": {
						Description:         "XDPRefreshInterval is the period at which Felix re-checks all XDP state to ensure that no other process has accidentally broken Calico's BPF maps or attached programs. Set to 0 to disable XDP refresh. [Default: 90s]",
						MarkdownDescription: "XDPRefreshInterval is the period at which Felix re-checks all XDP state to ensure that no other process has accidentally broken Calico's BPF maps or attached programs. Set to 0 to disable XDP refresh. [Default: 90s]",

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
		},
	}, nil
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_felix_configuration_v1")

	var state CrdProjectcalicoOrgFelixConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgFelixConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("FelixConfiguration")

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

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_felix_configuration_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_felix_configuration_v1")

	var state CrdProjectcalicoOrgFelixConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CrdProjectcalicoOrgFelixConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("crd.projectcalico.org/v1")
	goModel.Kind = utilities.Ptr("FelixConfiguration")

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

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_felix_configuration_v1")
	// NO-OP: Terraform removes the state automatically for us
}
