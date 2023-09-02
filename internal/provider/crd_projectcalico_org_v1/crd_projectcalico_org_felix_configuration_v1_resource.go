/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"regexp"
)

var (
	_ resource.Resource                = &CrdProjectcalicoOrgFelixConfigurationV1Resource{}
	_ resource.ResourceWithConfigure   = &CrdProjectcalicoOrgFelixConfigurationV1Resource{}
	_ resource.ResourceWithImportState = &CrdProjectcalicoOrgFelixConfigurationV1Resource{}
)

func NewCrdProjectcalicoOrgFelixConfigurationV1Resource() resource.Resource {
	return &CrdProjectcalicoOrgFelixConfigurationV1Resource{}
}

type CrdProjectcalicoOrgFelixConfigurationV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CrdProjectcalicoOrgFelixConfigurationV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AllowIPIPPacketsFromWorkloads      *bool              `tfsdk:"allow_ipip_packets_from_workloads" json:"allowIPIPPacketsFromWorkloads,omitempty"`
		AllowVXLANPacketsFromWorkloads     *bool              `tfsdk:"allow_vxlan_packets_from_workloads" json:"allowVXLANPacketsFromWorkloads,omitempty"`
		AwsSrcDstCheck                     *string            `tfsdk:"aws_src_dst_check" json:"awsSrcDstCheck,omitempty"`
		BpfCTLBLogFilter                   *string            `tfsdk:"bpf_ctlb_log_filter" json:"bpfCTLBLogFilter,omitempty"`
		BpfConnectTimeLoadBalancingEnabled *bool              `tfsdk:"bpf_connect_time_load_balancing_enabled" json:"bpfConnectTimeLoadBalancingEnabled,omitempty"`
		BpfDSROptoutCIDRs                  *[]string          `tfsdk:"bpf_dsr_optout_cid_rs" json:"bpfDSROptoutCIDRs,omitempty"`
		BpfDataIfacePattern                *string            `tfsdk:"bpf_data_iface_pattern" json:"bpfDataIfacePattern,omitempty"`
		BpfDisableGROForIfaces             *string            `tfsdk:"bpf_disable_gro_for_ifaces" json:"bpfDisableGROForIfaces,omitempty"`
		BpfDisableUnprivileged             *bool              `tfsdk:"bpf_disable_unprivileged" json:"bpfDisableUnprivileged,omitempty"`
		BpfEnabled                         *bool              `tfsdk:"bpf_enabled" json:"bpfEnabled,omitempty"`
		BpfEnforceRPF                      *string            `tfsdk:"bpf_enforce_rpf" json:"bpfEnforceRPF,omitempty"`
		BpfExtToServiceConnmark            *int64             `tfsdk:"bpf_ext_to_service_connmark" json:"bpfExtToServiceConnmark,omitempty"`
		BpfExternalServiceMode             *string            `tfsdk:"bpf_external_service_mode" json:"bpfExternalServiceMode,omitempty"`
		BpfForceTrackPacketsFromIfaces     *[]string          `tfsdk:"bpf_force_track_packets_from_ifaces" json:"bpfForceTrackPacketsFromIfaces,omitempty"`
		BpfHostConntrackBypass             *bool              `tfsdk:"bpf_host_conntrack_bypass" json:"bpfHostConntrackBypass,omitempty"`
		BpfKubeProxyEndpointSlicesEnabled  *bool              `tfsdk:"bpf_kube_proxy_endpoint_slices_enabled" json:"bpfKubeProxyEndpointSlicesEnabled,omitempty"`
		BpfKubeProxyIptablesCleanupEnabled *bool              `tfsdk:"bpf_kube_proxy_iptables_cleanup_enabled" json:"bpfKubeProxyIptablesCleanupEnabled,omitempty"`
		BpfKubeProxyMinSyncPeriod          *string            `tfsdk:"bpf_kube_proxy_min_sync_period" json:"bpfKubeProxyMinSyncPeriod,omitempty"`
		BpfL3IfacePattern                  *string            `tfsdk:"bpf_l3_iface_pattern" json:"bpfL3IfacePattern,omitempty"`
		BpfLogFilters                      *map[string]string `tfsdk:"bpf_log_filters" json:"bpfLogFilters,omitempty"`
		BpfLogLevel                        *string            `tfsdk:"bpf_log_level" json:"bpfLogLevel,omitempty"`
		BpfMapSizeConntrack                *int64             `tfsdk:"bpf_map_size_conntrack" json:"bpfMapSizeConntrack,omitempty"`
		BpfMapSizeIPSets                   *int64             `tfsdk:"bpf_map_size_ip_sets" json:"bpfMapSizeIPSets,omitempty"`
		BpfMapSizeIfState                  *int64             `tfsdk:"bpf_map_size_if_state" json:"bpfMapSizeIfState,omitempty"`
		BpfMapSizeNATAffinity              *int64             `tfsdk:"bpf_map_size_nat_affinity" json:"bpfMapSizeNATAffinity,omitempty"`
		BpfMapSizeNATBackend               *int64             `tfsdk:"bpf_map_size_nat_backend" json:"bpfMapSizeNATBackend,omitempty"`
		BpfMapSizeNATFrontend              *int64             `tfsdk:"bpf_map_size_nat_frontend" json:"bpfMapSizeNATFrontend,omitempty"`
		BpfMapSizeRoute                    *int64             `tfsdk:"bpf_map_size_route" json:"bpfMapSizeRoute,omitempty"`
		BpfPSNATPorts                      *string            `tfsdk:"bpf_psnat_ports" json:"bpfPSNATPorts,omitempty"`
		BpfPolicyDebugEnabled              *bool              `tfsdk:"bpf_policy_debug_enabled" json:"bpfPolicyDebugEnabled,omitempty"`
		ChainInsertMode                    *string            `tfsdk:"chain_insert_mode" json:"chainInsertMode,omitempty"`
		DataplaneDriver                    *string            `tfsdk:"dataplane_driver" json:"dataplaneDriver,omitempty"`
		DataplaneWatchdogTimeout           *string            `tfsdk:"dataplane_watchdog_timeout" json:"dataplaneWatchdogTimeout,omitempty"`
		DebugDisableLogDropping            *bool              `tfsdk:"debug_disable_log_dropping" json:"debugDisableLogDropping,omitempty"`
		DebugMemoryProfilePath             *string            `tfsdk:"debug_memory_profile_path" json:"debugMemoryProfilePath,omitempty"`
		DebugSimulateCalcGraphHangAfter    *string            `tfsdk:"debug_simulate_calc_graph_hang_after" json:"debugSimulateCalcGraphHangAfter,omitempty"`
		DebugSimulateDataplaneHangAfter    *string            `tfsdk:"debug_simulate_dataplane_hang_after" json:"debugSimulateDataplaneHangAfter,omitempty"`
		DefaultEndpointToHostAction        *string            `tfsdk:"default_endpoint_to_host_action" json:"defaultEndpointToHostAction,omitempty"`
		DeviceRouteProtocol                *int64             `tfsdk:"device_route_protocol" json:"deviceRouteProtocol,omitempty"`
		DeviceRouteSourceAddress           *string            `tfsdk:"device_route_source_address" json:"deviceRouteSourceAddress,omitempty"`
		DeviceRouteSourceAddressIPv6       *string            `tfsdk:"device_route_source_address_i_pv6" json:"deviceRouteSourceAddressIPv6,omitempty"`
		DisableConntrackInvalidCheck       *bool              `tfsdk:"disable_conntrack_invalid_check" json:"disableConntrackInvalidCheck,omitempty"`
		EndpointReportingDelay             *string            `tfsdk:"endpoint_reporting_delay" json:"endpointReportingDelay,omitempty"`
		EndpointReportingEnabled           *bool              `tfsdk:"endpoint_reporting_enabled" json:"endpointReportingEnabled,omitempty"`
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
		FeatureDetectOverride  *string `tfsdk:"feature_detect_override" json:"featureDetectOverride,omitempty"`
		FeatureGates           *string `tfsdk:"feature_gates" json:"featureGates,omitempty"`
		FloatingIPs            *string `tfsdk:"floating_i_ps" json:"floatingIPs,omitempty"`
		GenericXDPEnabled      *bool   `tfsdk:"generic_xdp_enabled" json:"genericXDPEnabled,omitempty"`
		HealthEnabled          *bool   `tfsdk:"health_enabled" json:"healthEnabled,omitempty"`
		HealthHost             *string `tfsdk:"health_host" json:"healthHost,omitempty"`
		HealthPort             *int64  `tfsdk:"health_port" json:"healthPort,omitempty"`
		HealthTimeoutOverrides *[]struct {
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
		} `tfsdk:"health_timeout_overrides" json:"healthTimeoutOverrides,omitempty"`
		InterfaceExclude                   *string   `tfsdk:"interface_exclude" json:"interfaceExclude,omitempty"`
		InterfacePrefix                    *string   `tfsdk:"interface_prefix" json:"interfacePrefix,omitempty"`
		InterfaceRefreshInterval           *string   `tfsdk:"interface_refresh_interval" json:"interfaceRefreshInterval,omitempty"`
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
		NatPortRange                       *string   `tfsdk:"nat_port_range" json:"natPortRange,omitempty"`
		NetlinkTimeout                     *string   `tfsdk:"netlink_timeout" json:"netlinkTimeout,omitempty"`
		OpenstackRegion                    *string   `tfsdk:"openstack_region" json:"openstackRegion,omitempty"`
		PolicySyncPathPrefix               *string   `tfsdk:"policy_sync_path_prefix" json:"policySyncPathPrefix,omitempty"`
		PrometheusGoMetricsEnabled         *bool     `tfsdk:"prometheus_go_metrics_enabled" json:"prometheusGoMetricsEnabled,omitempty"`
		PrometheusMetricsEnabled           *bool     `tfsdk:"prometheus_metrics_enabled" json:"prometheusMetricsEnabled,omitempty"`
		PrometheusMetricsHost              *string   `tfsdk:"prometheus_metrics_host" json:"prometheusMetricsHost,omitempty"`
		PrometheusMetricsPort              *int64    `tfsdk:"prometheus_metrics_port" json:"prometheusMetricsPort,omitempty"`
		PrometheusProcessMetricsEnabled    *bool     `tfsdk:"prometheus_process_metrics_enabled" json:"prometheusProcessMetricsEnabled,omitempty"`
		PrometheusWireGuardMetricsEnabled  *bool     `tfsdk:"prometheus_wire_guard_metrics_enabled" json:"prometheusWireGuardMetricsEnabled,omitempty"`
		RemoveExternalRoutes               *bool     `tfsdk:"remove_external_routes" json:"removeExternalRoutes,omitempty"`
		ReportingInterval                  *string   `tfsdk:"reporting_interval" json:"reportingInterval,omitempty"`
		ReportingTTL                       *string   `tfsdk:"reporting_ttl" json:"reportingTTL,omitempty"`
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
		WorkloadSourceSpoofing         *string `tfsdk:"workload_source_spoofing" json:"workloadSourceSpoofing,omitempty"`
		XdpEnabled                     *bool   `tfsdk:"xdp_enabled" json:"xdpEnabled,omitempty"`
		XdpRefreshInterval             *string `tfsdk:"xdp_refresh_interval" json:"xdpRefreshInterval,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_felix_configuration_v1"
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Felix Configuration contains the configuration for Felix.",
		MarkdownDescription: "Felix Configuration contains the configuration for Felix.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
						Description:         "AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic from workloads [Default: false]",
						MarkdownDescription: "AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic from workloads [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allow_vxlan_packets_from_workloads": schema.BoolAttribute{
						Description:         "AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic from workloads [Default: false]",
						MarkdownDescription: "AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic from workloads [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"aws_src_dst_check": schema.StringAttribute{
						Description:         "Set source-destination-check on AWS EC2 instances. Accepted value must be one of 'DoNothing', 'Enable' or 'Disable'. [Default: DoNothing]",
						MarkdownDescription: "Set source-destination-check on AWS EC2 instances. Accepted value must be one of 'DoNothing', 'Enable' or 'Disable'. [Default: DoNothing]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("DoNothing", "Enable", "Disable"),
						},
					},

					"bpf_ctlb_log_filter": schema.StringAttribute{
						Description:         "BPFCTLBLogFilter specifies, what is logged by connect time load balancer when BPFLogLevel is debug. Currently has to be specified as 'all' when BPFLogFilters is set to see CTLB logs. [Default: unset - means logs are emitted when BPFLogLevel id debug and BPFLogFilters not set.]",
						MarkdownDescription: "BPFCTLBLogFilter specifies, what is logged by connect time load balancer when BPFLogLevel is debug. Currently has to be specified as 'all' when BPFLogFilters is set to see CTLB logs. [Default: unset - means logs are emitted when BPFLogLevel id debug and BPFLogFilters not set.]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_connect_time_load_balancing_enabled": schema.BoolAttribute{
						Description:         "BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load balancer.  The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections.  The only reason to disable it is for debugging purposes.  [Default: true]",
						MarkdownDescription: "BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load balancer.  The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections.  The only reason to disable it is for debugging purposes.  [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_dsr_optout_cid_rs": schema.ListAttribute{
						Description:         "BPFDSROptoutCIDRs is a list of CIDRs which are excluded from DSR. That is, clients in those CIDRs will accesses nodeports as if BPFExternalServiceMode was set to Tunnel.",
						MarkdownDescription: "BPFDSROptoutCIDRs is a list of CIDRs which are excluded from DSR. That is, clients in those CIDRs will accesses nodeports as if BPFExternalServiceMode was set to Tunnel.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_data_iface_pattern": schema.StringAttribute{
						Description:         "BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to in order to catch traffic to/from the network.  This needs to match the interfaces that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.  It should not match the workload interfaces (usually named cali...).",
						MarkdownDescription: "BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to in order to catch traffic to/from the network.  This needs to match the interfaces that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.  It should not match the workload interfaces (usually named cali...).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_disable_gro_for_ifaces": schema.StringAttribute{
						Description:         "BPFDisableGROForIfaces is a regular expression that controls which interfaces Felix should disable the Generic Receive Offload [GRO] option.  It should not match the workload interfaces (usually named cali...).",
						MarkdownDescription: "BPFDisableGROForIfaces is a regular expression that controls which interfaces Felix should disable the Generic Receive Offload [GRO] option.  It should not match the workload interfaces (usually named cali...).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_disable_unprivileged": schema.BoolAttribute{
						Description:         "BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable unprivileged use of BPF.  This ensures that unprivileged users cannot access Calico's BPF maps and cannot insert their own BPF programs to interfere with Calico's. [Default: true]",
						MarkdownDescription: "BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable unprivileged use of BPF.  This ensures that unprivileged users cannot access Calico's BPF maps and cannot insert their own BPF programs to interfere with Calico's. [Default: true]",
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

					"bpf_ext_to_service_connmark": schema.Int64Attribute{
						Description:         "BPFExtToServiceConnmark in BPF mode, control a 32bit mark that is set on connections from an external client to a local service. This mark allows us to control how packets of that connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]",
						MarkdownDescription: "BPFExtToServiceConnmark in BPF mode, control a 32bit mark that is set on connections from an external client to a local service. This mark allows us to control how packets of that connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_external_service_mode": schema.StringAttribute{
						Description:         "BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports and cluster IPs) are forwarded to remote workloads.  If set to 'Tunnel' then both request and response traffic is tunneled to the remote node.  If set to 'DSR', the request traffic is tunneled but the response traffic is sent directly from the remote node.  In 'DSR' mode, the remote node appears to use the IP of the ingress node; this requires a permissive L2 network.  [Default: Tunnel]",
						MarkdownDescription: "BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports and cluster IPs) are forwarded to remote workloads.  If set to 'Tunnel' then both request and response traffic is tunneled to the remote node.  If set to 'DSR', the request traffic is tunneled but the response traffic is sent directly from the remote node.  In 'DSR' mode, the remote node appears to use the IP of the ingress node; this requires a permissive L2 network.  [Default: Tunnel]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Tunnel|DSR)?$`), ""),
						},
					},

					"bpf_force_track_packets_from_ifaces": schema.ListAttribute{
						Description:         "BPFForceTrackPacketsFromIfaces in BPF mode, forces traffic from these interfaces to skip Calico's iptables NOTRACK rule, allowing traffic from those interfaces to be tracked by Linux conntrack.  Should only be used for interfaces that are not used for the Calico fabric.  For example, a docker bridge device for non-Calico-networked containers. [Default: docker+]",
						MarkdownDescription: "BPFForceTrackPacketsFromIfaces in BPF mode, forces traffic from these interfaces to skip Calico's iptables NOTRACK rule, allowing traffic from those interfaces to be tracked by Linux conntrack.  Should only be used for interfaces that are not used for the Calico fabric.  For example, a docker bridge device for non-Calico-networked containers. [Default: docker+]",
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

					"bpf_kube_proxy_endpoint_slices_enabled": schema.BoolAttribute{
						Description:         "BPFKubeProxyEndpointSlicesEnabled in BPF mode, controls whether Felix's embedded kube-proxy accepts EndpointSlices or not.",
						MarkdownDescription: "BPFKubeProxyEndpointSlicesEnabled in BPF mode, controls whether Felix's embedded kube-proxy accepts EndpointSlices or not.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_kube_proxy_iptables_cleanup_enabled": schema.BoolAttribute{
						Description:         "BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream Kubernetes kube-proxy's iptables chains.  Should only be enabled if kube-proxy is not running.  [Default: true]",
						MarkdownDescription: "BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream Kubernetes kube-proxy's iptables chains.  Should only be enabled if kube-proxy is not running.  [Default: true]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_kube_proxy_min_sync_period": schema.StringAttribute{
						Description:         "BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's embedded kube-proxy.  Lower values give reduced set-up latency.  Higher values reduce Felix CPU usage by batching up more work.  [Default: 1s]",
						MarkdownDescription: "BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's embedded kube-proxy.  Lower values give reduced set-up latency.  Higher values reduce Felix CPU usage by batching up more work.  [Default: 1s]",
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
						Description:         "BPFLogFilters is a map of key=values where the value is a pcap filter expression and the key is an interface name with 'all' denoting all interfaces, 'weps' all workload endpoints and 'heps' all host endpoints.  When specified as an env var, it accepts a comma-separated list of key=values. [Default: unset - means all debug logs are emitted]",
						MarkdownDescription: "BPFLogFilters is a map of key=values where the value is a pcap filter expression and the key is an interface name with 'all' denoting all interfaces, 'weps' all workload endpoints and 'heps' all host endpoints.  When specified as an env var, it accepts a comma-separated list of key=values. [Default: unset - means all debug logs are emitted]",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_log_level": schema.StringAttribute{
						Description:         "BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode.  One of 'Off', 'Info', or 'Debug'.  The logs are emitted to the BPF trace pipe, accessible with the command 'tc exec bpf debug'. [Default: Off].",
						MarkdownDescription: "BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode.  One of 'Off', 'Info', or 'Debug'.  The logs are emitted to the BPF trace pipe, accessible with the command 'tc exec bpf debug'. [Default: Off].",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Off|Info|Debug)?$`), ""),
						},
					},

					"bpf_map_size_conntrack": schema.Int64Attribute{
						Description:         "BPFMapSizeConntrack sets the size for the conntrack map.  This map must be large enough to hold an entry for each active connection.  Warning: changing the size of the conntrack map can cause disruption.",
						MarkdownDescription: "BPFMapSizeConntrack sets the size for the conntrack map.  This map must be large enough to hold an entry for each active connection.  Warning: changing the size of the conntrack map can cause disruption.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_ip_sets": schema.Int64Attribute{
						Description:         "BPFMapSizeIPSets sets the size for ipsets map.  The IP sets map must be large enough to hold an entry for each endpoint matched by every selector in the source/destination matches in network policy.  Selectors such as 'all()' can result in large numbers of entries (one entry per endpoint in that case).",
						MarkdownDescription: "BPFMapSizeIPSets sets the size for ipsets map.  The IP sets map must be large enough to hold an entry for each endpoint matched by every selector in the source/destination matches in network policy.  Selectors such as 'all()' can result in large numbers of entries (one entry per endpoint in that case).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_if_state": schema.Int64Attribute{
						Description:         "BPFMapSizeIfState sets the size for ifstate map.  The ifstate map must be large enough to hold an entry for each device (host + workloads) on a host.",
						MarkdownDescription: "BPFMapSizeIfState sets the size for ifstate map.  The ifstate map must be large enough to hold an entry for each device (host + workloads) on a host.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_nat_affinity": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_nat_backend": schema.Int64Attribute{
						Description:         "BPFMapSizeNATBackend sets the size for nat back end map. This is the total number of endpoints. This is mostly more than the size of the number of services.",
						MarkdownDescription: "BPFMapSizeNATBackend sets the size for nat back end map. This is the total number of endpoints. This is mostly more than the size of the number of services.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_nat_frontend": schema.Int64Attribute{
						Description:         "BPFMapSizeNATFrontend sets the size for nat front end map. FrontendMap should be large enough to hold an entry for each nodeport, external IP and each port in each service.",
						MarkdownDescription: "BPFMapSizeNATFrontend sets the size for nat front end map. FrontendMap should be large enough to hold an entry for each nodeport, external IP and each port in each service.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_map_size_route": schema.Int64Attribute{
						Description:         "BPFMapSizeRoute sets the size for the routes map.  The routes map should be large enough to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and tunnel IPs).",
						MarkdownDescription: "BPFMapSizeRoute sets the size for the routes map.  The routes map should be large enough to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and tunnel IPs).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bpf_psnat_ports": schema.StringAttribute{
						Description:         "BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision. This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the  ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]",
						MarkdownDescription: "BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision. This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the  ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]",
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

					"chain_insert_mode": schema.StringAttribute{
						Description:         "ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule at the top of the chain or by appending a rule at the bottom. insert is the safe default since it prevents Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed. [Default: insert]",
						MarkdownDescription: "ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule at the top of the chain or by appending a rule at the bottom. insert is the safe default since it prevents Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed. [Default: insert]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(insert|append)?$`), ""),
						},
					},

					"dataplane_driver": schema.StringAttribute{
						Description:         "DataplaneDriver filename of the external dataplane driver to use.  Only used if UseInternalDataplaneDriver is set to false.",
						MarkdownDescription: "DataplaneDriver filename of the external dataplane driver to use.  Only used if UseInternalDataplaneDriver is set to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"dataplane_watchdog_timeout": schema.StringAttribute{
						Description:         "DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver. Increase this value if you experience spurious non-ready or non-live events when Felix is under heavy load. Decrease the value to get felix to report non-live or non-ready more quickly. [Default: 90s]  Deprecated: replaced by the generic HealthTimeoutOverrides.",
						MarkdownDescription: "DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver. Increase this value if you experience spurious non-ready or non-live events when Felix is under heavy load. Decrease the value to get felix to report non-live or non-ready more quickly. [Default: 90s]  Deprecated: replaced by the generic HealthTimeoutOverrides.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_disable_log_dropping": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_memory_profile_path": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug_simulate_calc_graph_hang_after": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"debug_simulate_dataplane_hang_after": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"default_endpoint_to_host_action": schema.StringAttribute{
						Description:         "DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host itself (after the traffic hits the endpoint egress policy). By default Calico blocks traffic from workload endpoints to the host itself with an iptables 'DROP' action. If you want to allow some or all traffic from endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables 'INPUT' chain; Calico will insert its rules at the top of that chain, then 'RETURN' packets to the 'INPUT' chain once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets from workloads after processing workload endpoint egress policy. [Default: Drop]",
						MarkdownDescription: "DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host itself (after the traffic hits the endpoint egress policy). By default Calico blocks traffic from workload endpoints to the host itself with an iptables 'DROP' action. If you want to allow some or all traffic from endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables 'INPUT' chain; Calico will insert its rules at the top of that chain, then 'RETURN' packets to the 'INPUT' chain once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets from workloads after processing workload endpoint egress policy. [Default: Drop]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Drop|Accept|Return)?$`), ""),
						},
					},

					"device_route_protocol": schema.Int64Attribute{
						Description:         "This defines the route protocol added to programmed device routes, by default this will be RTPROT_BOOT when left blank.",
						MarkdownDescription: "This defines the route protocol added to programmed device routes, by default this will be RTPROT_BOOT when left blank.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"device_route_source_address": schema.StringAttribute{
						Description:         "This is the IPv4 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",
						MarkdownDescription: "This is the IPv4 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"device_route_source_address_i_pv6": schema.StringAttribute{
						Description:         "This is the IPv6 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",
						MarkdownDescription: "This is the IPv6 source address to use on programmed device routes. By default the source address is left blank, leaving the kernel to choose the source address used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_conntrack_invalid_check": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"endpoint_reporting_delay": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"endpoint_reporting_enabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"external_nodes_list": schema.ListAttribute{
						Description:         "ExternalNodesCIDRList is a list of CIDR's of external-non-calico-nodes which may source tunnel traffic and have the tunneled traffic be accepted at calico nodes.",
						MarkdownDescription: "ExternalNodesCIDRList is a list of CIDR's of external-non-calico-nodes which may source tunnel traffic and have the tunneled traffic be accepted at calico nodes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"failsafe_inbound_host_ports": schema.ListNestedAttribute{
						Description:         "FailsafeInboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow incoming traffic to host endpoints on irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports, use the value none. The default value allows ssh access and DHCP. [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667]",
						MarkdownDescription: "FailsafeInboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow incoming traffic to host endpoints on irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports, use the value none. The default value allows ssh access and DHCP. [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667]",
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

					"failsafe_outbound_host_ports": schema.ListNestedAttribute{
						Description:         "FailsafeOutboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow outgoing traffic from host endpoints to irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports, use the value none. The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd as well as allowing DHCP and DNS. [Default: tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667, udp:53, udp:67]",
						MarkdownDescription: "FailsafeOutboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow outgoing traffic from host endpoints to irrespective of the security policy. This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to 'tcp'. If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports, use the value none. The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd as well as allowing DHCP and DNS. [Default: tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667, udp:53, udp:67]",
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

					"feature_detect_override": schema.StringAttribute{
						Description:         "FeatureDetectOverride is used to override feature detection based on auto-detected platform capabilities.  Values are specified in a comma separated list with no spaces, example; 'SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock='.  'true' or 'false' will force the feature, empty or omitted values are auto-detected.",
						MarkdownDescription: "FeatureDetectOverride is used to override feature detection based on auto-detected platform capabilities.  Values are specified in a comma separated list with no spaces, example; 'SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock='.  'true' or 'false' will force the feature, empty or omitted values are auto-detected.",
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
						Description:         "FloatingIPs configures whether or not Felix will program non-OpenStack floating IP addresses.  (OpenStack-derived floating IPs are always programmed, regardless of this setting.)",
						MarkdownDescription: "FloatingIPs configures whether or not Felix will program non-OpenStack floating IP addresses.  (OpenStack-derived floating IPs are always programmed, regardless of this setting.)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"generic_xdp_enabled": schema.BoolAttribute{
						Description:         "GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver modes can use XDP. This is not recommended since it doesn't provide better performance than iptables. [Default: false]",
						MarkdownDescription: "GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver modes can use XDP. This is not recommended since it doesn't provide better performance than iptables. [Default: false]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_enabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_port": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_timeout_overrides": schema.ListNestedAttribute{
						Description:         "HealthTimeoutOverrides allows the internal watchdog timeouts of individual subcomponents to be overridden.  This is useful for working around 'false positive' liveness timeouts that can occur in particularly stressful workloads or if CPU is constrained.  For a list of active subcomponents, see Felix's logs.",
						MarkdownDescription: "HealthTimeoutOverrides allows the internal watchdog timeouts of individual subcomponents to be overridden.  This is useful for working around 'false positive' liveness timeouts that can occur in particularly stressful workloads or if CPU is constrained.  For a list of active subcomponents, see Felix's logs.",
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
						Description:         "InterfaceExclude is a comma-separated list of interfaces that Felix should exclude when monitoring for host endpoints. The default value ensures that Felix ignores Kubernetes' IPVS dummy interface, which is used internally by kube-proxy. If you want to exclude multiple interface names using a single value, the list supports regular expressions. For regular expressions you must wrap the value with '/'. For example having values '/^kube/,veth1' will exclude all interfaces that begin with 'kube' and also the interface 'veth1'. [Default: kube-ipvs0]",
						MarkdownDescription: "InterfaceExclude is a comma-separated list of interfaces that Felix should exclude when monitoring for host endpoints. The default value ensures that Felix ignores Kubernetes' IPVS dummy interface, which is used internally by kube-proxy. If you want to exclude multiple interface names using a single value, the list supports regular expressions. For regular expressions you must wrap the value with '/'. For example having values '/^kube/,veth1' will exclude all interfaces that begin with 'kube' and also the interface 'veth1'. [Default: kube-ipvs0]",
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

					"ipip_enabled": schema.BoolAttribute{
						Description:         "IPIPEnabled overrides whether Felix should configure an IPIP interface on the host. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						MarkdownDescription: "IPIPEnabled overrides whether Felix should configure an IPIP interface on the host. Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipip_mtu": schema.Int64Attribute{
						Description:         "IPIPMTU is the MTU to set on the tunnel device. See Configuring MTU [Default: 1440]",
						MarkdownDescription: "IPIPMTU is the MTU to set on the tunnel device. See Configuring MTU [Default: 1440]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipsets_refresh_interval": schema.StringAttribute{
						Description:         "IpsetsRefreshInterval is the period at which Felix re-checks all iptables state to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable iptables refresh. [Default: 90s]",
						MarkdownDescription: "IpsetsRefreshInterval is the period at which Felix re-checks all iptables state to ensure that no other process has accidentally broken Calico's rules. Set to 0 to disable iptables refresh. [Default: 90s]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"iptables_backend": schema.StringAttribute{
						Description:         "IptablesBackend specifies which backend of iptables will be used. The default is Auto.",
						MarkdownDescription: "IptablesBackend specifies which backend of iptables will be used. The default is Auto.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Auto|FelixConfiguration|FelixConfigurationList|Legacy|NFT)?$`), ""),
						},
					},

					"iptables_filter_allow_action": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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
						Description:         "IptablesLockProbeInterval is the time that Felix will wait between attempts to acquire the iptables lock if it is not available. Lower values make Felix more responsive when the lock is contended, but use more CPU. [Default: 50ms]",
						MarkdownDescription: "IptablesLockProbeInterval is the time that Felix will wait between attempts to acquire the iptables lock if it is not available. Lower values make Felix more responsive when the lock is contended, but use more CPU. [Default: 50ms]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"iptables_lock_timeout": schema.StringAttribute{
						Description:         "IptablesLockTimeout is the time that Felix will wait for the iptables lock, or 0, to disable. To use this feature, Felix must share the iptables lock file with all other processes that also take the lock. When running Felix inside a container, this requires the /run directory of the host to be mounted into the calico/node or calico/felix container. [Default: 0s disabled]",
						MarkdownDescription: "IptablesLockTimeout is the time that Felix will wait for the iptables lock, or 0, to disable. To use this feature, Felix must share the iptables lock file with all other processes that also take the lock. When running Felix inside a container, this requires the /run directory of the host to be mounted into the calico/node or calico/felix container. [Default: 0s disabled]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
					},

					"iptables_mangle_allow_action": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Accept|Return)?$`), ""),
						},
					},

					"iptables_mark_mask": schema.Int64Attribute{
						Description:         "IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xff000000]",
						MarkdownDescription: "IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from. Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xff000000]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"iptables_nat_outgoing_interface_filter": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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
						Description:         "LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs. Only logs from files with names that match the given regular expression are included.  The filter only applies to Debug level logs.",
						MarkdownDescription: "LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs. Only logs from files with names that match the given regular expression are included.  The filter only applies to Debug level logs.",
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
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Debug|Info|Warning|Error|Fatal)?$`), ""),
						},
					},

					"log_severity_screen": schema.StringAttribute{
						Description:         "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						MarkdownDescription: "LogSeverityScreen is the log severity above which logs are sent to the stdout. [Default: Info]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Debug|Info|Warning|Error|Fatal)?$`), ""),
						},
					},

					"log_severity_sys": schema.StringAttribute{
						Description:         "LogSeveritySys is the log severity above which logs are sent to the syslog. Set to None for no logging to syslog. [Default: Info]",
						MarkdownDescription: "LogSeveritySys is the log severity above which logs are sent to the syslog. Set to None for no logging to syslog. [Default: Info]",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(?i)(Debug|Info|Warning|Error|Fatal)?$`), ""),
						},
					},

					"max_ipset_size": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata_addr": schema.StringAttribute{
						Description:         "MetadataAddr is the IP address or domain name of the server that can answer VM queries for cloud-init metadata. In OpenStack, this corresponds to the machine running nova-api (or in Ubuntu, nova-api-metadata). A value of none (case insensitive) means that Felix should not set up any NAT rule for the metadata path. [Default: 127.0.0.1]",
						MarkdownDescription: "MetadataAddr is the IP address or domain name of the server that can answer VM queries for cloud-init metadata. In OpenStack, this corresponds to the machine running nova-api (or in Ubuntu, nova-api-metadata). A value of none (case insensitive) means that Felix should not set up any NAT rule for the metadata path. [Default: 127.0.0.1]",
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
						Description:         "NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that is leaving the network. By default the address used is an address on the interface the traffic is leaving on (ie it uses the iptables MASQUERADE target)",
						MarkdownDescription: "NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that is leaving the network. By default the address used is an address on the interface the traffic is leaving on (ie it uses the iptables MASQUERADE target)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nat_port_range": schema.StringAttribute{
						Description:         "NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT. When unset the default behavior of the network stack is used.",
						MarkdownDescription: "NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT. When unset the default behavior of the network stack is used.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"netlink_timeout": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\\.[0-9]+)?(ms|s|m|h))*$`), ""),
						},
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
						Description:         "Whether or not to remove device routes that have not been programmed by Felix. Disabling this will allow external applications to also add device routes. This is enabled by default which means we will remove externally added routes.",
						MarkdownDescription: "Whether or not to remove device routes that have not been programmed by Felix. Disabling this will allow external applications to also add device routes. This is enabled by default which means we will remove externally added routes.",
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
						Description:         "UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic.  If false, it will launch an external dataplane driver and communicate with it over protobuf.",
						MarkdownDescription: "UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic.  If false, it will launch an external dataplane driver and communicate with it over protobuf.",
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
						Description:         "VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device. See Configuring MTU [Default: 1410]",
						MarkdownDescription: "VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device. See Configuring MTU [Default: 1410]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_mtuv6": schema.Int64Attribute{
						Description:         "VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device. See Configuring MTU [Default: 1390]",
						MarkdownDescription: "VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device. See Configuring MTU [Default: 1390]",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_port": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vxlan_vni": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
						Description:         "WireguardKeepAlive controls Wireguard PersistentKeepalive option. Set 0 to disable. [Default: 0]",
						MarkdownDescription: "WireguardKeepAlive controls Wireguard PersistentKeepalive option. Set 0 to disable. [Default: 0]",
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

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_crd_projectcalico_org_felix_configuration_v1")

	var model CrdProjectcalicoOrgFelixConfigurationV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("FelixConfiguration")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "FelixConfiguration"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CrdProjectcalicoOrgFelixConfigurationV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_crd_projectcalico_org_felix_configuration_v1")

	var data CrdProjectcalicoOrgFelixConfigurationV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "FelixConfiguration"}).
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

	var readResponse CrdProjectcalicoOrgFelixConfigurationV1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_crd_projectcalico_org_felix_configuration_v1")

	var model CrdProjectcalicoOrgFelixConfigurationV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	model.Kind = pointer.String("FelixConfiguration")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "FelixConfiguration"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse CrdProjectcalicoOrgFelixConfigurationV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_crd_projectcalico_org_felix_configuration_v1")

	var data CrdProjectcalicoOrgFelixConfigurationV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "FelixConfiguration"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *CrdProjectcalicoOrgFelixConfigurationV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
