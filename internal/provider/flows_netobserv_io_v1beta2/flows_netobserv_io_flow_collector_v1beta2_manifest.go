/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flows_netobserv_io_v1beta2

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
	_ datasource.DataSource = &FlowsNetobservIoFlowCollectorV1Beta2Manifest{}
)

func NewFlowsNetobservIoFlowCollectorV1Beta2Manifest() datasource.DataSource {
	return &FlowsNetobservIoFlowCollectorV1Beta2Manifest{}
}

type FlowsNetobservIoFlowCollectorV1Beta2Manifest struct{}

type FlowsNetobservIoFlowCollectorV1Beta2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Agent *struct {
			Ebpf *struct {
				Advanced *struct {
					Env        *map[string]string `tfsdk:"env" json:"env,omitempty"`
					Scheduling *struct {
						Affinity *struct {
							NodeAffinity *struct {
								PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
									Preference *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchFields *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_fields" json:"matchFields,omitempty"`
									} `tfsdk:"preference" json:"preference,omitempty"`
									Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
								} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
								RequiredDuringSchedulingIgnoredDuringExecution *struct {
									NodeSelectorTerms *[]struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchFields *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_fields" json:"matchFields,omitempty"`
									} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
								} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
							} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
							PodAffinity *struct {
								PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
									PodAffinityTerm *struct {
										LabelSelector *struct {
											MatchExpressions *[]struct {
												Key      *string   `tfsdk:"key" json:"key,omitempty"`
												Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
												Values   *[]string `tfsdk:"values" json:"values,omitempty"`
											} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
											MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
										} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
										MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
										MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
										NamespaceSelector *struct {
											MatchExpressions *[]struct {
												Key      *string   `tfsdk:"key" json:"key,omitempty"`
												Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
												Values   *[]string `tfsdk:"values" json:"values,omitempty"`
											} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
											MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
										} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
										Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
										TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
									} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
									Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
								} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
								RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
									Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
									TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
								} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
							} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
							PodAntiAffinity *struct {
								PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
									PodAffinityTerm *struct {
										LabelSelector *struct {
											MatchExpressions *[]struct {
												Key      *string   `tfsdk:"key" json:"key,omitempty"`
												Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
												Values   *[]string `tfsdk:"values" json:"values,omitempty"`
											} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
											MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
										} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
										MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
										MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
										NamespaceSelector *struct {
											MatchExpressions *[]struct {
												Key      *string   `tfsdk:"key" json:"key,omitempty"`
												Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
												Values   *[]string `tfsdk:"values" json:"values,omitempty"`
											} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
											MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
										} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
										Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
										TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
									} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
									Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
								} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
								RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
									Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
									TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
								} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
							} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
						} `tfsdk:"affinity" json:"affinity,omitempty"`
						NodeSelector      *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
						Tolerations       *[]struct {
							Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
							Key               *string `tfsdk:"key" json:"key,omitempty"`
							Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
							TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
							Value             *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					} `tfsdk:"scheduling" json:"scheduling,omitempty"`
				} `tfsdk:"advanced" json:"advanced,omitempty"`
				CacheActiveTimeout *string   `tfsdk:"cache_active_timeout" json:"cacheActiveTimeout,omitempty"`
				CacheMaxFlows      *int64    `tfsdk:"cache_max_flows" json:"cacheMaxFlows,omitempty"`
				ExcludeInterfaces  *[]string `tfsdk:"exclude_interfaces" json:"excludeInterfaces,omitempty"`
				Features           *[]string `tfsdk:"features" json:"features,omitempty"`
				FlowFilter         *struct {
					Action      *string `tfsdk:"action" json:"action,omitempty"`
					Cidr        *string `tfsdk:"cidr" json:"cidr,omitempty"`
					DestPorts   *string `tfsdk:"dest_ports" json:"destPorts,omitempty"`
					Direction   *string `tfsdk:"direction" json:"direction,omitempty"`
					Enable      *bool   `tfsdk:"enable" json:"enable,omitempty"`
					IcmpCode    *int64  `tfsdk:"icmp_code" json:"icmpCode,omitempty"`
					IcmpType    *int64  `tfsdk:"icmp_type" json:"icmpType,omitempty"`
					PeerIP      *string `tfsdk:"peer_ip" json:"peerIP,omitempty"`
					PktDrops    *bool   `tfsdk:"pkt_drops" json:"pktDrops,omitempty"`
					Ports       *string `tfsdk:"ports" json:"ports,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
					SourcePorts *string `tfsdk:"source_ports" json:"sourcePorts,omitempty"`
					TcpFlags    *string `tfsdk:"tcp_flags" json:"tcpFlags,omitempty"`
				} `tfsdk:"flow_filter" json:"flowFilter,omitempty"`
				ImagePullPolicy *string   `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Interfaces      *[]string `tfsdk:"interfaces" json:"interfaces,omitempty"`
				KafkaBatchSize  *int64    `tfsdk:"kafka_batch_size" json:"kafkaBatchSize,omitempty"`
				LogLevel        *string   `tfsdk:"log_level" json:"logLevel,omitempty"`
				Metrics         *struct {
					DisableAlerts *[]string `tfsdk:"disable_alerts" json:"disableAlerts,omitempty"`
					Enable        *bool     `tfsdk:"enable" json:"enable,omitempty"`
					Server        *struct {
						Port *int64 `tfsdk:"port" json:"port,omitempty"`
						Tls  *struct {
							InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
							Provided           *struct {
								CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
								CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								Type      *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"provided" json:"provided,omitempty"`
							ProvidedCaFile *struct {
								File      *string `tfsdk:"file" json:"file,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								Type      *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"provided_ca_file" json:"providedCaFile,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"tls" json:"tls,omitempty"`
					} `tfsdk:"server" json:"server,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				Privileged *bool `tfsdk:"privileged" json:"privileged,omitempty"`
				Resources  *struct {
					Claims *[]struct {
						Name    *string `tfsdk:"name" json:"name,omitempty"`
						Request *string `tfsdk:"request" json:"request,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Sampling *int64 `tfsdk:"sampling" json:"sampling,omitempty"`
			} `tfsdk:"ebpf" json:"ebpf,omitempty"`
			Ipfix *struct {
				CacheActiveTimeout     *string `tfsdk:"cache_active_timeout" json:"cacheActiveTimeout,omitempty"`
				CacheMaxFlows          *int64  `tfsdk:"cache_max_flows" json:"cacheMaxFlows,omitempty"`
				ClusterNetworkOperator *struct {
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"cluster_network_operator" json:"clusterNetworkOperator,omitempty"`
				ForceSampleAll *bool `tfsdk:"force_sample_all" json:"forceSampleAll,omitempty"`
				OvnKubernetes  *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					DaemonSetName *string `tfsdk:"daemon_set_name" json:"daemonSetName,omitempty"`
					Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ovn_kubernetes" json:"ovnKubernetes,omitempty"`
				Sampling *int64 `tfsdk:"sampling" json:"sampling,omitempty"`
			} `tfsdk:"ipfix" json:"ipfix,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"agent" json:"agent,omitempty"`
		ConsolePlugin *struct {
			Advanced *struct {
				Args       *[]string          `tfsdk:"args" json:"args,omitempty"`
				Env        *map[string]string `tfsdk:"env" json:"env,omitempty"`
				Port       *int64             `tfsdk:"port" json:"port,omitempty"`
				Register   *bool              `tfsdk:"register" json:"register,omitempty"`
				Scheduling *struct {
					Affinity *struct {
						NodeAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								Preference *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchFields *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_fields" json:"matchFields,omitempty"`
								} `tfsdk:"preference" json:"preference,omitempty"`
								Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
							RequiredDuringSchedulingIgnoredDuringExecution *struct {
								NodeSelectorTerms *[]struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchFields *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_fields" json:"matchFields,omitempty"`
								} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
						PodAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								PodAffinityTerm *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
									Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
									TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
								} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
								Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
							RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
								MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
						PodAntiAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								PodAffinityTerm *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
									Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
									TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
								} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
								Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
							RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
								MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" json:"affinity,omitempty"`
					NodeSelector      *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
					Tolerations       *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"scheduling" json:"scheduling,omitempty"`
			} `tfsdk:"advanced" json:"advanced,omitempty"`
			Autoscaler *struct {
				MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
				Metrics     *[]struct {
					ContainerResource *struct {
						Container *string `tfsdk:"container" json:"container,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Target    *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"container_resource" json:"containerResource,omitempty"`
					External *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"external" json:"external,omitempty"`
					Object *struct {
						DescribedObject *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
							Name       *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"described_object" json:"describedObject,omitempty"`
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"object" json:"object,omitempty"`
					Pods *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"pods" json:"pods,omitempty"`
					Resource *struct {
						Name   *string `tfsdk:"name" json:"name,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"resource" json:"resource,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
				Status      *string `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"autoscaler" json:"autoscaler,omitempty"`
			Enable          *bool   `tfsdk:"enable" json:"enable,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			LogLevel        *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			PortNaming      *struct {
				Enable    *bool              `tfsdk:"enable" json:"enable,omitempty"`
				PortNames *map[string]string `tfsdk:"port_names" json:"portNames,omitempty"`
			} `tfsdk:"port_naming" json:"portNaming,omitempty"`
			QuickFilters *[]struct {
				Default *bool              `tfsdk:"default" json:"default,omitempty"`
				Filter  *map[string]string `tfsdk:"filter" json:"filter,omitempty"`
				Name    *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"quick_filters" json:"quickFilters,omitempty"`
			Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"console_plugin" json:"consolePlugin,omitempty"`
		DeploymentModel *string `tfsdk:"deployment_model" json:"deploymentModel,omitempty"`
		Exporters       *[]struct {
			Ipfix *struct {
				TargetHost *string `tfsdk:"target_host" json:"targetHost,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
				Transport  *string `tfsdk:"transport" json:"transport,omitempty"`
			} `tfsdk:"ipfix" json:"ipfix,omitempty"`
			Kafka *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Sasl    *struct {
					ClientIDReference *struct {
						File      *string `tfsdk:"file" json:"file,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"client_id_reference" json:"clientIDReference,omitempty"`
					ClientSecretReference *struct {
						File      *string `tfsdk:"file" json:"file,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"client_secret_reference" json:"clientSecretReference,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"sasl" json:"sasl,omitempty"`
				Tls *struct {
					CaCert *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_cert" json:"caCert,omitempty"`
					Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					UserCert           *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"user_cert" json:"userCert,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Topic *string `tfsdk:"topic" json:"topic,omitempty"`
			} `tfsdk:"kafka" json:"kafka,omitempty"`
			OpenTelemetry *struct {
				FieldsMapping *[]struct {
					Input      *string `tfsdk:"input" json:"input,omitempty"`
					Multiplier *int64  `tfsdk:"multiplier" json:"multiplier,omitempty"`
					Output     *string `tfsdk:"output" json:"output,omitempty"`
				} `tfsdk:"fields_mapping" json:"fieldsMapping,omitempty"`
				Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
				Logs    *struct {
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"logs" json:"logs,omitempty"`
				Metrics *struct {
					Enable           *bool   `tfsdk:"enable" json:"enable,omitempty"`
					PushTimeInterval *string `tfsdk:"push_time_interval" json:"pushTimeInterval,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetHost *string `tfsdk:"target_host" json:"targetHost,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
				Tls        *struct {
					CaCert *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_cert" json:"caCert,omitempty"`
					Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					UserCert           *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"user_cert" json:"userCert,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"open_telemetry" json:"openTelemetry,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"exporters" json:"exporters,omitempty"`
		Kafka *struct {
			Address *string `tfsdk:"address" json:"address,omitempty"`
			Sasl    *struct {
				ClientIDReference *struct {
					File      *string `tfsdk:"file" json:"file,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"client_id_reference" json:"clientIDReference,omitempty"`
				ClientSecretReference *struct {
					File      *string `tfsdk:"file" json:"file,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"client_secret_reference" json:"clientSecretReference,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"sasl" json:"sasl,omitempty"`
			Tls *struct {
				CaCert *struct {
					CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_cert" json:"caCert,omitempty"`
				Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				UserCert           *struct {
					CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Type      *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"user_cert" json:"userCert,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Topic *string `tfsdk:"topic" json:"topic,omitempty"`
		} `tfsdk:"kafka" json:"kafka,omitempty"`
		Loki *struct {
			Advanced *struct {
				StaticLabels    *map[string]string `tfsdk:"static_labels" json:"staticLabels,omitempty"`
				WriteMaxBackoff *string            `tfsdk:"write_max_backoff" json:"writeMaxBackoff,omitempty"`
				WriteMaxRetries *int64             `tfsdk:"write_max_retries" json:"writeMaxRetries,omitempty"`
				WriteMinBackoff *string            `tfsdk:"write_min_backoff" json:"writeMinBackoff,omitempty"`
			} `tfsdk:"advanced" json:"advanced,omitempty"`
			Enable    *bool `tfsdk:"enable" json:"enable,omitempty"`
			LokiStack *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"loki_stack" json:"lokiStack,omitempty"`
			Manual *struct {
				AuthToken   *string `tfsdk:"auth_token" json:"authToken,omitempty"`
				IngesterUrl *string `tfsdk:"ingester_url" json:"ingesterUrl,omitempty"`
				QuerierUrl  *string `tfsdk:"querier_url" json:"querierUrl,omitempty"`
				StatusTls   *struct {
					CaCert *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_cert" json:"caCert,omitempty"`
					Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					UserCert           *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"user_cert" json:"userCert,omitempty"`
				} `tfsdk:"status_tls" json:"statusTls,omitempty"`
				StatusUrl *string `tfsdk:"status_url" json:"statusUrl,omitempty"`
				TenantID  *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				Tls       *struct {
					CaCert *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_cert" json:"caCert,omitempty"`
					Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					UserCert           *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"user_cert" json:"userCert,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"manual" json:"manual,omitempty"`
			Microservices *struct {
				IngesterUrl *string `tfsdk:"ingester_url" json:"ingesterUrl,omitempty"`
				QuerierUrl  *string `tfsdk:"querier_url" json:"querierUrl,omitempty"`
				TenantID    *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				Tls         *struct {
					CaCert *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_cert" json:"caCert,omitempty"`
					Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					UserCert           *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"user_cert" json:"userCert,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"microservices" json:"microservices,omitempty"`
			Mode       *string `tfsdk:"mode" json:"mode,omitempty"`
			Monolithic *struct {
				TenantID *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				Tls      *struct {
					CaCert *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"ca_cert" json:"caCert,omitempty"`
					Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					UserCert           *struct {
						CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Type      *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"user_cert" json:"userCert,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"monolithic" json:"monolithic,omitempty"`
			ReadTimeout    *string `tfsdk:"read_timeout" json:"readTimeout,omitempty"`
			WriteBatchSize *int64  `tfsdk:"write_batch_size" json:"writeBatchSize,omitempty"`
			WriteBatchWait *string `tfsdk:"write_batch_wait" json:"writeBatchWait,omitempty"`
			WriteTimeout   *string `tfsdk:"write_timeout" json:"writeTimeout,omitempty"`
		} `tfsdk:"loki" json:"loki,omitempty"`
		Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
		NetworkPolicy *struct {
			AdditionalNamespaces *[]string `tfsdk:"additional_namespaces" json:"additionalNamespaces,omitempty"`
			Enable               *bool     `tfsdk:"enable" json:"enable,omitempty"`
		} `tfsdk:"network_policy" json:"networkPolicy,omitempty"`
		Processor *struct {
			AddZone  *bool `tfsdk:"add_zone" json:"addZone,omitempty"`
			Advanced *struct {
				ConversationEndTimeout         *string            `tfsdk:"conversation_end_timeout" json:"conversationEndTimeout,omitempty"`
				ConversationHeartbeatInterval  *string            `tfsdk:"conversation_heartbeat_interval" json:"conversationHeartbeatInterval,omitempty"`
				ConversationTerminatingTimeout *string            `tfsdk:"conversation_terminating_timeout" json:"conversationTerminatingTimeout,omitempty"`
				DropUnusedFields               *bool              `tfsdk:"drop_unused_fields" json:"dropUnusedFields,omitempty"`
				EnableKubeProbes               *bool              `tfsdk:"enable_kube_probes" json:"enableKubeProbes,omitempty"`
				Env                            *map[string]string `tfsdk:"env" json:"env,omitempty"`
				HealthPort                     *int64             `tfsdk:"health_port" json:"healthPort,omitempty"`
				Port                           *int64             `tfsdk:"port" json:"port,omitempty"`
				ProfilePort                    *int64             `tfsdk:"profile_port" json:"profilePort,omitempty"`
				Scheduling                     *struct {
					Affinity *struct {
						NodeAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								Preference *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchFields *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_fields" json:"matchFields,omitempty"`
								} `tfsdk:"preference" json:"preference,omitempty"`
								Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
							RequiredDuringSchedulingIgnoredDuringExecution *struct {
								NodeSelectorTerms *[]struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchFields *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_fields" json:"matchFields,omitempty"`
								} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
						PodAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								PodAffinityTerm *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
									Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
									TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
								} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
								Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
							RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
								MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
						PodAntiAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								PodAffinityTerm *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
									MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
									MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key      *string   `tfsdk:"key" json:"key,omitempty"`
											Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
											Values   *[]string `tfsdk:"values" json:"values,omitempty"`
										} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
										MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
									Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
									TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
								} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
								Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
							RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
								MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
								MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
								Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
								TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" json:"affinity,omitempty"`
					NodeSelector      *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
					Tolerations       *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"scheduling" json:"scheduling,omitempty"`
				SecondaryNetworks *[]struct {
					Index *[]string `tfsdk:"index" json:"index,omitempty"`
					Name  *string   `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secondary_networks" json:"secondaryNetworks,omitempty"`
			} `tfsdk:"advanced" json:"advanced,omitempty"`
			ClusterName             *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
			ImagePullPolicy         *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			KafkaConsumerAutoscaler *struct {
				MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
				Metrics     *[]struct {
					ContainerResource *struct {
						Container *string `tfsdk:"container" json:"container,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Target    *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"container_resource" json:"containerResource,omitempty"`
					External *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"external" json:"external,omitempty"`
					Object *struct {
						DescribedObject *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
							Name       *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"described_object" json:"describedObject,omitempty"`
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"object" json:"object,omitempty"`
					Pods *struct {
						Metric *struct {
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Selector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selector" json:"selector,omitempty"`
						} `tfsdk:"metric" json:"metric,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"pods" json:"pods,omitempty"`
					Resource *struct {
						Name   *string `tfsdk:"name" json:"name,omitempty"`
						Target *struct {
							AverageUtilization *int64  `tfsdk:"average_utilization" json:"averageUtilization,omitempty"`
							AverageValue       *string `tfsdk:"average_value" json:"averageValue,omitempty"`
							Type               *string `tfsdk:"type" json:"type,omitempty"`
							Value              *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"target" json:"target,omitempty"`
					} `tfsdk:"resource" json:"resource,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				MinReplicas *int64  `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
				Status      *string `tfsdk:"status" json:"status,omitempty"`
			} `tfsdk:"kafka_consumer_autoscaler" json:"kafkaConsumerAutoscaler,omitempty"`
			KafkaConsumerBatchSize     *int64  `tfsdk:"kafka_consumer_batch_size" json:"kafkaConsumerBatchSize,omitempty"`
			KafkaConsumerQueueCapacity *int64  `tfsdk:"kafka_consumer_queue_capacity" json:"kafkaConsumerQueueCapacity,omitempty"`
			KafkaConsumerReplicas      *int64  `tfsdk:"kafka_consumer_replicas" json:"kafkaConsumerReplicas,omitempty"`
			LogLevel                   *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			LogTypes                   *string `tfsdk:"log_types" json:"logTypes,omitempty"`
			Metrics                    *struct {
				DisableAlerts *[]string `tfsdk:"disable_alerts" json:"disableAlerts,omitempty"`
				IncludeList   *[]string `tfsdk:"include_list" json:"includeList,omitempty"`
				Server        *struct {
					Port *int64 `tfsdk:"port" json:"port,omitempty"`
					Tls  *struct {
						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Provided           *struct {
							CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
							CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Type      *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"provided" json:"provided,omitempty"`
						ProvidedCaFile *struct {
							File      *string `tfsdk:"file" json:"file,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Type      *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"provided_ca_file" json:"providedCaFile,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"server" json:"server,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			MultiClusterDeployment *bool `tfsdk:"multi_cluster_deployment" json:"multiClusterDeployment,omitempty"`
			Resources              *struct {
				Claims *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			SubnetLabels *struct {
				CustomLabels *[]struct {
					Cidrs *[]string `tfsdk:"cidrs" json:"cidrs,omitempty"`
					Name  *string   `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"custom_labels" json:"customLabels,omitempty"`
				OpenShiftAutoDetect *bool `tfsdk:"open_shift_auto_detect" json:"openShiftAutoDetect,omitempty"`
			} `tfsdk:"subnet_labels" json:"subnetLabels,omitempty"`
		} `tfsdk:"processor" json:"processor,omitempty"`
		Prometheus *struct {
			Querier *struct {
				Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				Manual *struct {
					ForwardUserToken *bool `tfsdk:"forward_user_token" json:"forwardUserToken,omitempty"`
					Tls              *struct {
						CaCert *struct {
							CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
							CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Type      *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"ca_cert" json:"caCert,omitempty"`
						Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						UserCert           *struct {
							CertFile  *string `tfsdk:"cert_file" json:"certFile,omitempty"`
							CertKey   *string `tfsdk:"cert_key" json:"certKey,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							Type      *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"user_cert" json:"userCert,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"manual" json:"manual,omitempty"`
				Mode    *string `tfsdk:"mode" json:"mode,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"querier" json:"querier,omitempty"`
		} `tfsdk:"prometheus" json:"prometheus,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlowsNetobservIoFlowCollectorV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flows_netobserv_io_flow_collector_v1beta2_manifest"
}

func (r *FlowsNetobservIoFlowCollectorV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "'FlowCollector' is the schema for the network flows collection API, which pilots and configures the underlying deployments.",
		MarkdownDescription: "'FlowCollector' is the schema for the network flows collection API, which pilots and configures the underlying deployments.",
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
				Description:         "Defines the desired state of the FlowCollector resource. <br><br> *: the mention of 'unsupported' or 'deprecated' for a feature throughout this document means that this feature is not officially supported by Red Hat. It might have been, for example, contributed by the community and accepted without a formal agreement for maintenance. The product maintainers might provide some support for these features as a best effort only.",
				MarkdownDescription: "Defines the desired state of the FlowCollector resource. <br><br> *: the mention of 'unsupported' or 'deprecated' for a feature throughout this document means that this feature is not officially supported by Red Hat. It might have been, for example, contributed by the community and accepted without a formal agreement for maintenance. The product maintainers might provide some support for these features as a best effort only.",
				Attributes: map[string]schema.Attribute{
					"agent": schema.SingleNestedAttribute{
						Description:         "Agent configuration for flows extraction.",
						MarkdownDescription: "Agent configuration for flows extraction.",
						Attributes: map[string]schema.Attribute{
							"ebpf": schema.SingleNestedAttribute{
								Description:         "'ebpf' describes the settings related to the eBPF-based flow reporter when 'spec.agent.type' is set to 'eBPF'.",
								MarkdownDescription: "'ebpf' describes the settings related to the eBPF-based flow reporter when 'spec.agent.type' is set to 'eBPF'.",
								Attributes: map[string]schema.Attribute{
									"advanced": schema.SingleNestedAttribute{
										Description:         "'advanced' allows setting some aspects of the internal configuration of the eBPF agent. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
										MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the eBPF agent. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
										Attributes: map[string]schema.Attribute{
											"env": schema.MapAttribute{
												Description:         "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
												MarkdownDescription: "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"scheduling": schema.SingleNestedAttribute{
												Description:         "scheduling controls how the pods are scheduled on nodes.",
												MarkdownDescription: "scheduling controls how the pods are scheduled on nodes.",
												Attributes: map[string]schema.Attribute{
													"affinity": schema.SingleNestedAttribute{
														Description:         "If specified, the pod's scheduling constraints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
														MarkdownDescription: "If specified, the pod's scheduling constraints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
														Attributes: map[string]schema.Attribute{
															"node_affinity": schema.SingleNestedAttribute{
																Description:         "Describes node affinity scheduling rules for the pod.",
																MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
																Attributes: map[string]schema.Attribute{
																	"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																		MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"preference": schema.SingleNestedAttribute{
																					Description:         "A node selector term, associated with the corresponding weight.",
																					MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "A list of node selector requirements by node's labels.",
																							MarkdownDescription: "A list of node selector requirements by node's labels.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "The label key that the selector applies to.",
																										MarkdownDescription: "The label key that the selector applies to.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																										MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																						"match_fields": schema.ListNestedAttribute{
																							Description:         "A list of node selector requirements by node's fields.",
																							MarkdownDescription: "A list of node selector requirements by node's fields.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "The label key that the selector applies to.",
																										MarkdownDescription: "The label key that the selector applies to.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																										MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																					},
																					Required: true,
																					Optional: false,
																					Computed: false,
																				},

																				"weight": schema.Int64Attribute{
																					Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																					MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

																	"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
																		Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																		MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																		Attributes: map[string]schema.Attribute{
																			"node_selector_terms": schema.ListNestedAttribute{
																				Description:         "Required. A list of node selector terms. The terms are ORed.",
																				MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "A list of node selector requirements by node's labels.",
																							MarkdownDescription: "A list of node selector requirements by node's labels.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "The label key that the selector applies to.",
																										MarkdownDescription: "The label key that the selector applies to.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																										MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																						"match_fields": schema.ListNestedAttribute{
																							Description:         "A list of node selector requirements by node's fields.",
																							MarkdownDescription: "A list of node selector requirements by node's fields.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "The label key that the selector applies to.",
																										MarkdownDescription: "The label key that the selector applies to.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																										MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																					},
																				},
																				Required: true,
																				Optional: false,
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

															"pod_affinity": schema.SingleNestedAttribute{
																Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
																Attributes: map[string]schema.Attribute{
																	"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																		MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"pod_affinity_term": schema.SingleNestedAttribute{
																					Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																					MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																					Attributes: map[string]schema.Attribute{
																						"label_selector": schema.SingleNestedAttribute{
																							Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																							MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																							Attributes: map[string]schema.Attribute{
																								"match_expressions": schema.ListNestedAttribute{
																									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																								"match_labels": schema.MapAttribute{
																									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"match_label_keys": schema.ListAttribute{
																							Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"mismatch_label_keys": schema.ListAttribute{
																							Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"namespace_selector": schema.SingleNestedAttribute{
																							Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																							MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																							Attributes: map[string]schema.Attribute{
																								"match_expressions": schema.ListNestedAttribute{
																									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																								"match_labels": schema.MapAttribute{
																									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"namespaces": schema.ListAttribute{
																							Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"topology_key": schema.StringAttribute{
																							Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																							MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: true,
																					Optional: false,
																					Computed: false,
																				},

																				"weight": schema.Int64Attribute{
																					Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																					MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

																	"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																		MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"match_label_keys": schema.ListAttribute{
																					Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"namespaces": schema.ListAttribute{
																					Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"topology_key": schema.StringAttribute{
																					Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"pod_anti_affinity": schema.SingleNestedAttribute{
																Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
																Attributes: map[string]schema.Attribute{
																	"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																		MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"pod_affinity_term": schema.SingleNestedAttribute{
																					Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																					MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																					Attributes: map[string]schema.Attribute{
																						"label_selector": schema.SingleNestedAttribute{
																							Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																							MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																							Attributes: map[string]schema.Attribute{
																								"match_expressions": schema.ListNestedAttribute{
																									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																								"match_labels": schema.MapAttribute{
																									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"match_label_keys": schema.ListAttribute{
																							Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"mismatch_label_keys": schema.ListAttribute{
																							Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"namespace_selector": schema.SingleNestedAttribute{
																							Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																							MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																							Attributes: map[string]schema.Attribute{
																								"match_expressions": schema.ListNestedAttribute{
																									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																								"match_labels": schema.MapAttribute{
																									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},
																							},
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"namespaces": schema.ListAttribute{
																							Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"topology_key": schema.StringAttribute{
																							Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																							MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},
																					},
																					Required: true,
																					Optional: false,
																					Computed: false,
																				},

																				"weight": schema.Int64Attribute{
																					Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																					MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

																	"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																		MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"match_label_keys": schema.ListAttribute{
																					Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"namespaces": schema.ListAttribute{
																					Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"topology_key": schema.StringAttribute{
																					Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

													"node_selector": schema.MapAttribute{
														Description:         "'nodeSelector' allows scheduling of pods only onto nodes that have each of the specified labels. For documentation, refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														MarkdownDescription: "'nodeSelector' allows scheduling of pods only onto nodes that have each of the specified labels. For documentation, refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"priority_class_name": schema.StringAttribute{
														Description:         "If specified, indicates the pod's priority. For documentation, refer to https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#how-to-use-priority-and-preemption. If not specified, default priority is used, or zero if there is no default.",
														MarkdownDescription: "If specified, indicates the pod's priority. For documentation, refer to https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#how-to-use-priority-and-preemption. If not specified, default priority is used, or zero if there is no default.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tolerations": schema.ListNestedAttribute{
														Description:         "'tolerations' is a list of tolerations that allow the pod to schedule onto nodes with matching taints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
														MarkdownDescription: "'tolerations' is a list of tolerations that allow the pod to schedule onto nodes with matching taints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"effect": schema.StringAttribute{
																	Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																	MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"key": schema.StringAttribute{
																	Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																	MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																	MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"toleration_seconds": schema.Int64Attribute{
																	Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																	MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
																	MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

									"cache_active_timeout": schema.StringAttribute{
										Description:         "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										MarkdownDescription: "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(ns|ms|s|m)?$`), ""),
										},
									},

									"cache_max_flows": schema.Int64Attribute{
										Description:         "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										MarkdownDescription: "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows. Increasing 'cacheMaxFlows' and 'cacheActiveTimeout' can decrease the network traffic overhead and the CPU load, however you can expect higher memory consumption and an increased latency in the flow collection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"exclude_interfaces": schema.ListAttribute{
										Description:         "'excludeInterfaces' contains the interface names that are excluded from flow tracing. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										MarkdownDescription: "'excludeInterfaces' contains the interface names that are excluded from flow tracing. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"features": schema.ListAttribute{
										Description:         "List of additional features to enable. They are all disabled by default. Enabling additional features might have performance impacts. Possible values are:<br> - 'PacketDrop': enable the packets drop flows logging feature. This feature requires mounting the kernel debug filesystem, so the eBPF pod has to run as privileged. If the 'spec.agent.ebpf.privileged' parameter is not set, an error is reported.<br> - 'DNSTracking': enable the DNS tracking feature.<br> - 'FlowRTT': enable flow latency (sRTT) extraction in the eBPF agent from TCP traffic.<br> - 'NetworkEvents': enable the Network events monitoring feature. This feature requires mounting the kernel debug filesystem, so the eBPF pod has to run as privileged. It requires using the OVN-Kubernetes network plugin with the Observability feature. IMPORTANT: this feature is available as a Developer Preview.<br>",
										MarkdownDescription: "List of additional features to enable. They are all disabled by default. Enabling additional features might have performance impacts. Possible values are:<br> - 'PacketDrop': enable the packets drop flows logging feature. This feature requires mounting the kernel debug filesystem, so the eBPF pod has to run as privileged. If the 'spec.agent.ebpf.privileged' parameter is not set, an error is reported.<br> - 'DNSTracking': enable the DNS tracking feature.<br> - 'FlowRTT': enable flow latency (sRTT) extraction in the eBPF agent from TCP traffic.<br> - 'NetworkEvents': enable the Network events monitoring feature. This feature requires mounting the kernel debug filesystem, so the eBPF pod has to run as privileged. It requires using the OVN-Kubernetes network plugin with the Observability feature. IMPORTANT: this feature is available as a Developer Preview.<br>",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"flow_filter": schema.SingleNestedAttribute{
										Description:         "'flowFilter' defines the eBPF agent configuration regarding flow filtering.",
										MarkdownDescription: "'flowFilter' defines the eBPF agent configuration regarding flow filtering.",
										Attributes: map[string]schema.Attribute{
											"action": schema.StringAttribute{
												Description:         "'action' defines the action to perform on the flows that match the filter.",
												MarkdownDescription: "'action' defines the action to perform on the flows that match the filter.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Accept", "Reject"),
												},
											},

											"cidr": schema.StringAttribute{
												Description:         "'cidr' defines the IP CIDR to filter flows by. Examples: '10.10.10.0/24' or '100:100:100:100::/64'",
												MarkdownDescription: "'cidr' defines the IP CIDR to filter flows by. Examples: '10.10.10.0/24' or '100:100:100:100::/64'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"dest_ports": schema.StringAttribute{
												Description:         "'destPorts' defines the destination ports to filter flows by. To filter a single port, set a single port as an integer value. For example, 'destPorts: 80'. To filter a range of ports, use a 'start-end' range in string format. For example, 'destPorts: '80-100''. To filter two ports, use a 'port1,port2' in string format. For example, 'ports: '80,100''.",
												MarkdownDescription: "'destPorts' defines the destination ports to filter flows by. To filter a single port, set a single port as an integer value. For example, 'destPorts: 80'. To filter a range of ports, use a 'start-end' range in string format. For example, 'destPorts: '80-100''. To filter two ports, use a 'port1,port2' in string format. For example, 'ports: '80,100''.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"direction": schema.StringAttribute{
												Description:         "'direction' defines the direction to filter flows by.",
												MarkdownDescription: "'direction' defines the direction to filter flows by.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Ingress", "Egress"),
												},
											},

											"enable": schema.BoolAttribute{
												Description:         "Set 'enable' to 'true' to enable the eBPF flow filtering feature.",
												MarkdownDescription: "Set 'enable' to 'true' to enable the eBPF flow filtering feature.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"icmp_code": schema.Int64Attribute{
												Description:         "'icmpCode', for Internet Control Message Protocol (ICMP) traffic, defines the ICMP code to filter flows by.",
												MarkdownDescription: "'icmpCode', for Internet Control Message Protocol (ICMP) traffic, defines the ICMP code to filter flows by.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"icmp_type": schema.Int64Attribute{
												Description:         "'icmpType', for ICMP traffic, defines the ICMP type to filter flows by.",
												MarkdownDescription: "'icmpType', for ICMP traffic, defines the ICMP type to filter flows by.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"peer_ip": schema.StringAttribute{
												Description:         "'peerIP' defines the IP address to filter flows by. Example: '10.10.10.10'.",
												MarkdownDescription: "'peerIP' defines the IP address to filter flows by. Example: '10.10.10.10'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pkt_drops": schema.BoolAttribute{
												Description:         "'pktDrops', to filter flows with packet drops",
												MarkdownDescription: "'pktDrops', to filter flows with packet drops",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ports": schema.StringAttribute{
												Description:         "'ports' defines the ports to filter flows by. It is used both for source and destination ports. To filter a single port, set a single port as an integer value. For example, 'ports: 80'. To filter a range of ports, use a 'start-end' range in string format. For example, 'ports: '80-100''. To filter two ports, use a 'port1,port2' in string format. For example, 'ports: '80,100''.",
												MarkdownDescription: "'ports' defines the ports to filter flows by. It is used both for source and destination ports. To filter a single port, set a single port as an integer value. For example, 'ports: 80'. To filter a range of ports, use a 'start-end' range in string format. For example, 'ports: '80-100''. To filter two ports, use a 'port1,port2' in string format. For example, 'ports: '80,100''.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"protocol": schema.StringAttribute{
												Description:         "'protocol' defines the protocol to filter flows by.",
												MarkdownDescription: "'protocol' defines the protocol to filter flows by.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("TCP", "UDP", "ICMP", "ICMPv6", "SCTP"),
												},
											},

											"source_ports": schema.StringAttribute{
												Description:         "'sourcePorts' defines the source ports to filter flows by. To filter a single port, set a single port as an integer value. For example, 'sourcePorts: 80'. To filter a range of ports, use a 'start-end' range in string format. For example, 'sourcePorts: '80-100''. To filter two ports, use a 'port1,port2' in string format. For example, 'ports: '80,100''.",
												MarkdownDescription: "'sourcePorts' defines the source ports to filter flows by. To filter a single port, set a single port as an integer value. For example, 'sourcePorts: 80'. To filter a range of ports, use a 'start-end' range in string format. For example, 'sourcePorts: '80-100''. To filter two ports, use a 'port1,port2' in string format. For example, 'ports: '80,100''.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tcp_flags": schema.StringAttribute{
												Description:         "'tcpFlags' defines the TCP flags to filter flows by.",
												MarkdownDescription: "'tcpFlags' defines the TCP flags to filter flows by.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("SYN", "SYN-ACK", "ACK", "FIN", "RST", "URG", "ECE", "CWR", "FIN-ACK", "RST-ACK"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
										MarkdownDescription: "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
										},
									},

									"interfaces": schema.ListAttribute{
										Description:         "'interfaces' contains the interface names from where flows are collected. If empty, the agent fetches all the interfaces in the system, excepting the ones listed in 'excludeInterfaces'. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										MarkdownDescription: "'interfaces' contains the interface names from where flows are collected. If empty, the agent fetches all the interfaces in the system, excepting the ones listed in 'excludeInterfaces'. An entry enclosed by slashes, such as '/br-/', is matched as a regular expression. Otherwise it is matched as a case-sensitive string.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kafka_batch_size": schema.Int64Attribute{
										Description:         "'kafkaBatchSize' limits the maximum size of a request in bytes before being sent to a partition. Ignored when not using Kafka. Default: 1MB.",
										MarkdownDescription: "'kafkaBatchSize' limits the maximum size of a request in bytes before being sent to a partition. Ignored when not using Kafka. Default: 1MB.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_level": schema.StringAttribute{
										Description:         "'logLevel' defines the log level for the NetObserv eBPF Agent",
										MarkdownDescription: "'logLevel' defines the log level for the NetObserv eBPF Agent",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
										},
									},

									"metrics": schema.SingleNestedAttribute{
										Description:         "'metrics' defines the eBPF agent configuration regarding metrics.",
										MarkdownDescription: "'metrics' defines the eBPF agent configuration regarding metrics.",
										Attributes: map[string]schema.Attribute{
											"disable_alerts": schema.ListAttribute{
												Description:         "'disableAlerts' is a list of alerts that should be disabled. Possible values are:<br> 'NetObservDroppedFlows', which is triggered when the eBPF agent is missing packets or flows, such as when the BPF hashmap is busy or full, or the capacity limiter is being triggered.<br>",
												MarkdownDescription: "'disableAlerts' is a list of alerts that should be disabled. Possible values are:<br> 'NetObservDroppedFlows', which is triggered when the eBPF agent is missing packets or flows, such as when the BPF hashmap is busy or full, or the capacity limiter is being triggered.<br>",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Set 'enable' to 'false' to disable eBPF agent metrics collection. It is enabled by default.",
												MarkdownDescription: "Set 'enable' to 'false' to disable eBPF agent metrics collection. It is enabled by default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"server": schema.SingleNestedAttribute{
												Description:         "Metrics server endpoint configuration for the Prometheus scraper.",
												MarkdownDescription: "Metrics server endpoint configuration for the Prometheus scraper.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "The metrics server HTTP port.",
														MarkdownDescription: "The metrics server HTTP port.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"tls": schema.SingleNestedAttribute{
														Description:         "TLS configuration.",
														MarkdownDescription: "TLS configuration.",
														Attributes: map[string]schema.Attribute{
															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
																MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"provided": schema.SingleNestedAttribute{
																Description:         "TLS configuration when 'type' is set to 'Provided'.",
																MarkdownDescription: "TLS configuration when 'type' is set to 'Provided'.",
																Attributes: map[string]schema.Attribute{
																	"cert_file": schema.StringAttribute{
																		Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
																		MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"cert_key": schema.StringAttribute{
																		Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																		MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the config map or secret containing certificates.",
																		MarkdownDescription: "Name of the config map or secret containing certificates.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																		MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
																		MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("configmap", "secret"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"provided_ca_file": schema.SingleNestedAttribute{
																Description:         "Reference to the CA file when 'type' is set to 'Provided'.",
																MarkdownDescription: "Reference to the CA file when 'type' is set to 'Provided'.",
																Attributes: map[string]schema.Attribute{
																	"file": schema.StringAttribute{
																		Description:         "File name within the config map or secret.",
																		MarkdownDescription: "File name within the config map or secret.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the config map or secret containing the file.",
																		MarkdownDescription: "Name of the config map or secret containing the file.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																		MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "Type for the file reference: 'configmap' or 'secret'.",
																		MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("configmap", "secret"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": schema.StringAttribute{
																Description:         "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. [Unsupported (*)]. - 'Auto' to use OpenShift auto generated certificate using annotations.",
																MarkdownDescription: "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. [Unsupported (*)]. - 'Auto' to use OpenShift auto generated certificate using annotations.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Disabled", "Provided", "Auto"),
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

									"privileged": schema.BoolAttribute{
										Description:         "Privileged mode for the eBPF Agent container. When ignored or set to 'false', the operator sets granular capabilities (BPF, PERFMON, NET_ADMIN, SYS_RESOURCE) to the container. If for some reason these capabilities cannot be set, such as if an old kernel version not knowing CAP_BPF is in use, then you can turn on this mode for more global privileges. Some agent features require the privileged mode, such as packet drops tracking (see 'features') and SR-IOV support.",
										MarkdownDescription: "Privileged mode for the eBPF Agent container. When ignored or set to 'false', the operator sets granular capabilities (BPF, PERFMON, NET_ADMIN, SYS_RESOURCE) to the container. If for some reason these capabilities cannot be set, such as if an old kernel version not knowing CAP_BPF is in use, then you can turn on this mode for more global privileges. Some agent features require the privileged mode, such as packet drops tracking (see 'features') and SR-IOV support.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "'resources' are the compute resources required by this container. For more information, see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "'resources' are the compute resources required by this container. For more information, see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"request": schema.StringAttribute{
															Description:         "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
															MarkdownDescription: "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
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

											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"sampling": schema.Int64Attribute{
										Description:         "Sampling rate of the flow reporter. 100 means one flow on 100 is sent. 0 or 1 means all flows are sampled.",
										MarkdownDescription: "Sampling rate of the flow reporter. 100 means one flow on 100 is sent. 0 or 1 means all flows are sampled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipfix": schema.SingleNestedAttribute{
								Description:         "'ipfix' [deprecated (*)] - describes the settings related to the IPFIX-based flow reporter when 'spec.agent.type' is set to 'IPFIX'.",
								MarkdownDescription: "'ipfix' [deprecated (*)] - describes the settings related to the IPFIX-based flow reporter when 'spec.agent.type' is set to 'IPFIX'.",
								Attributes: map[string]schema.Attribute{
									"cache_active_timeout": schema.StringAttribute{
										Description:         "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending.",
										MarkdownDescription: "'cacheActiveTimeout' is the max period during which the reporter aggregates flows before sending.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^\d+(ns|ms|s|m)?$`), ""),
										},
									},

									"cache_max_flows": schema.Int64Attribute{
										Description:         "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows.",
										MarkdownDescription: "'cacheMaxFlows' is the max number of flows in an aggregate; when reached, the reporter sends the flows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"cluster_network_operator": schema.SingleNestedAttribute{
										Description:         "'clusterNetworkOperator' defines the settings related to the OpenShift Cluster Network Operator, when available.",
										MarkdownDescription: "'clusterNetworkOperator' defines the settings related to the OpenShift Cluster Network Operator, when available.",
										Attributes: map[string]schema.Attribute{
											"namespace": schema.StringAttribute{
												Description:         "Namespace where the config map is going to be deployed.",
												MarkdownDescription: "Namespace where the config map is going to be deployed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"force_sample_all": schema.BoolAttribute{
										Description:         "'forceSampleAll' allows disabling sampling in the IPFIX-based flow reporter. It is not recommended to sample all the traffic with IPFIX, as it might generate cluster instability. If you REALLY want to do that, set this flag to 'true'. Use at your own risk. When it is set to 'true', the value of 'sampling' is ignored.",
										MarkdownDescription: "'forceSampleAll' allows disabling sampling in the IPFIX-based flow reporter. It is not recommended to sample all the traffic with IPFIX, as it might generate cluster instability. If you REALLY want to do that, set this flag to 'true'. Use at your own risk. When it is set to 'true', the value of 'sampling' is ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ovn_kubernetes": schema.SingleNestedAttribute{
										Description:         "'ovnKubernetes' defines the settings of the OVN-Kubernetes network plugin, when available. This configuration is used when using OVN's IPFIX exports, without OpenShift. When using OpenShift, refer to the 'clusterNetworkOperator' property instead.",
										MarkdownDescription: "'ovnKubernetes' defines the settings of the OVN-Kubernetes network plugin, when available. This configuration is used when using OVN's IPFIX exports, without OpenShift. When using OpenShift, refer to the 'clusterNetworkOperator' property instead.",
										Attributes: map[string]schema.Attribute{
											"container_name": schema.StringAttribute{
												Description:         "'containerName' defines the name of the container to configure for IPFIX.",
												MarkdownDescription: "'containerName' defines the name of the container to configure for IPFIX.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"daemon_set_name": schema.StringAttribute{
												Description:         "'daemonSetName' defines the name of the DaemonSet controlling the OVN-Kubernetes pods.",
												MarkdownDescription: "'daemonSetName' defines the name of the DaemonSet controlling the OVN-Kubernetes pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace where OVN-Kubernetes pods are deployed.",
												MarkdownDescription: "Namespace where OVN-Kubernetes pods are deployed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"sampling": schema.Int64Attribute{
										Description:         "'sampling' is the sampling rate on the reporter. 100 means one flow on 100 is sent. To ensure cluster stability, it is not possible to set a value below 2. If you really want to sample every packet, which might impact the cluster stability, refer to 'forceSampleAll'. Alternatively, you can use the eBPF Agent instead of IPFIX.",
										MarkdownDescription: "'sampling' is the sampling rate on the reporter. 100 means one flow on 100 is sent. To ensure cluster stability, it is not possible to set a value below 2. If you really want to sample every packet, which might impact the cluster stability, refer to 'forceSampleAll'. Alternatively, you can use the eBPF Agent instead of IPFIX.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(2),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "'type' [deprecated (*)] selects the flows tracing agent. Previously, this field allowed to select between 'eBPF' or 'IPFIX'. Only 'eBPF' is allowed now, so this field is deprecated and is planned for removal in a future version of the API.",
								MarkdownDescription: "'type' [deprecated (*)] selects the flows tracing agent. Previously, this field allowed to select between 'eBPF' or 'IPFIX'. Only 'eBPF' is allowed now, so this field is deprecated and is planned for removal in a future version of the API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("eBPF", "IPFIX"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"console_plugin": schema.SingleNestedAttribute{
						Description:         "'consolePlugin' defines the settings related to the OpenShift Console plugin, when available.",
						MarkdownDescription: "'consolePlugin' defines the settings related to the OpenShift Console plugin, when available.",
						Attributes: map[string]schema.Attribute{
							"advanced": schema.SingleNestedAttribute{
								Description:         "'advanced' allows setting some aspects of the internal configuration of the console plugin. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the console plugin. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								Attributes: map[string]schema.Attribute{
									"args": schema.ListAttribute{
										Description:         "'args' allows passing custom arguments to underlying components. Useful for overriding some parameters, such as a URL or a configuration path, that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										MarkdownDescription: "'args' allows passing custom arguments to underlying components. Useful for overriding some parameters, such as a URL or a configuration path, that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env": schema.MapAttribute{
										Description:         "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										MarkdownDescription: "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "'port' is the plugin service port. Do not use 9002, which is reserved for metrics.",
										MarkdownDescription: "'port' is the plugin service port. Do not use 9002, which is reserved for metrics.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},

									"register": schema.BoolAttribute{
										Description:         "'register' allows, when set to 'true', to automatically register the provided console plugin with the OpenShift Console operator. When set to 'false', you can still register it manually by editing console.operator.openshift.io/cluster with the following command: 'oc patch console.operator.openshift.io cluster --type='json' -p '[{'op': 'add', 'path': '/spec/plugins/-', 'value': 'netobserv-plugin'}]''",
										MarkdownDescription: "'register' allows, when set to 'true', to automatically register the provided console plugin with the OpenShift Console operator. When set to 'false', you can still register it manually by editing console.operator.openshift.io/cluster with the following command: 'oc patch console.operator.openshift.io cluster --type='json' -p '[{'op': 'add', 'path': '/spec/plugins/-', 'value': 'netobserv-plugin'}]''",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheduling": schema.SingleNestedAttribute{
										Description:         "'scheduling' controls how the pods are scheduled on nodes.",
										MarkdownDescription: "'scheduling' controls how the pods are scheduled on nodes.",
										Attributes: map[string]schema.Attribute{
											"affinity": schema.SingleNestedAttribute{
												Description:         "If specified, the pod's scheduling constraints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												MarkdownDescription: "If specified, the pod's scheduling constraints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												Attributes: map[string]schema.Attribute{
													"node_affinity": schema.SingleNestedAttribute{
														Description:         "Describes node affinity scheduling rules for the pod.",
														MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
														Attributes: map[string]schema.Attribute{
															"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"preference": schema.SingleNestedAttribute{
																			Description:         "A node selector term, associated with the corresponding weight.",
																			MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's labels.",
																					MarkdownDescription: "A list of node selector requirements by node's labels.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																				"match_fields": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's fields.",
																					MarkdownDescription: "A list of node selector requirements by node's fields.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"weight": schema.Int64Attribute{
																			Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																			MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

															"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
																Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																Attributes: map[string]schema.Attribute{
																	"node_selector_terms": schema.ListNestedAttribute{
																		Description:         "Required. A list of node selector terms. The terms are ORed.",
																		MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's labels.",
																					MarkdownDescription: "A list of node selector requirements by node's labels.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																				"match_fields": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's fields.",
																					MarkdownDescription: "A list of node selector requirements by node's fields.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																			},
																		},
																		Required: true,
																		Optional: false,
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

													"pod_affinity": schema.SingleNestedAttribute{
														Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
														MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
														Attributes: map[string]schema.Attribute{
															"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"pod_affinity_term": schema.SingleNestedAttribute{
																			Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																			MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"match_label_keys": schema.ListAttribute{
																					Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"namespaces": schema.ListAttribute{
																					Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"topology_key": schema.StringAttribute{
																					Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"weight": schema.Int64Attribute{
																			Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																			MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

															"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"match_label_keys": schema.ListAttribute{
																			Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"mismatch_label_keys": schema.ListAttribute{
																			Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_anti_affinity": schema.SingleNestedAttribute{
														Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
														MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
														Attributes: map[string]schema.Attribute{
															"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"pod_affinity_term": schema.SingleNestedAttribute{
																			Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																			MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"match_label_keys": schema.ListAttribute{
																					Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"namespaces": schema.ListAttribute{
																					Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"topology_key": schema.StringAttribute{
																					Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"weight": schema.Int64Attribute{
																			Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																			MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

															"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"match_label_keys": schema.ListAttribute{
																			Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"mismatch_label_keys": schema.ListAttribute{
																			Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

											"node_selector": schema.MapAttribute{
												Description:         "'nodeSelector' allows scheduling of pods only onto nodes that have each of the specified labels. For documentation, refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
												MarkdownDescription: "'nodeSelector' allows scheduling of pods only onto nodes that have each of the specified labels. For documentation, refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_class_name": schema.StringAttribute{
												Description:         "If specified, indicates the pod's priority. For documentation, refer to https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#how-to-use-priority-and-preemption. If not specified, default priority is used, or zero if there is no default.",
												MarkdownDescription: "If specified, indicates the pod's priority. For documentation, refer to https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#how-to-use-priority-and-preemption. If not specified, default priority is used, or zero if there is no default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tolerations": schema.ListNestedAttribute{
												Description:         "'tolerations' is a list of tolerations that allow the pod to schedule onto nodes with matching taints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												MarkdownDescription: "'tolerations' is a list of tolerations that allow the pod to schedule onto nodes with matching taints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"effect": schema.StringAttribute{
															Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"toleration_seconds": schema.Int64Attribute{
															Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
															MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"autoscaler": schema.SingleNestedAttribute{
								Description:         "'autoscaler' spec of a horizontal pod autoscaler to set up for the plugin Deployment.",
								MarkdownDescription: "'autoscaler' spec of a horizontal pod autoscaler to set up for the plugin Deployment.",
								Attributes: map[string]schema.Attribute{
									"max_replicas": schema.Int64Attribute{
										Description:         "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										MarkdownDescription: "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.ListNestedAttribute{
										Description:         "Metrics used by the pod autoscaler. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/horizontal-pod-autoscaler-v2/",
										MarkdownDescription: "Metrics used by the pod autoscaler. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/horizontal-pod-autoscaler-v2/",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"container_resource": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"container": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"external": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_labels": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"object": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"described_object": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"metric": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_labels": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"pods": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_labels": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
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

									"min_replicas": schema.Int64Attribute{
										Description:         "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										MarkdownDescription: "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status": schema.StringAttribute{
										Description:         "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										MarkdownDescription: "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Enabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Enables the console plugin deployment.",
								MarkdownDescription: "Enables the console plugin deployment.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								MarkdownDescription: "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "'logLevel' for the console plugin backend",
								MarkdownDescription: "'logLevel' for the console plugin backend",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
								},
							},

							"port_naming": schema.SingleNestedAttribute{
								Description:         "'portNaming' defines the configuration of the port-to-service name translation",
								MarkdownDescription: "'portNaming' defines the configuration of the port-to-service name translation",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "Enable the console plugin port-to-service name translation",
										MarkdownDescription: "Enable the console plugin port-to-service name translation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_names": schema.MapAttribute{
										Description:         "'portNames' defines additional port names to use in the console, for example, 'portNames: {'3100': 'loki'}'.",
										MarkdownDescription: "'portNames' defines additional port names to use in the console, for example, 'portNames: {'3100': 'loki'}'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"quick_filters": schema.ListNestedAttribute{
								Description:         "'quickFilters' configures quick filter presets for the Console plugin",
								MarkdownDescription: "'quickFilters' configures quick filter presets for the Console plugin",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"default": schema.BoolAttribute{
											Description:         "'default' defines whether this filter should be active by default or not",
											MarkdownDescription: "'default' defines whether this filter should be active by default or not",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"filter": schema.MapAttribute{
											Description:         "'filter' is a set of keys and values to be set when this filter is selected. Each key can relate to a list of values using a coma-separated string, for example, 'filter: {'src_namespace': 'namespace1,namespace2'}'.",
											MarkdownDescription: "'filter' is a set of keys and values to be set when this filter is selected. Each key can relate to a list of values using a coma-separated string, for example, 'filter: {'src_namespace': 'namespace1,namespace2'}'.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the filter, that is displayed in the Console",
											MarkdownDescription: "Name of the filter, that is displayed in the Console",
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

							"replicas": schema.Int64Attribute{
								Description:         "'replicas' defines the number of replicas (pods) to start.",
								MarkdownDescription: "'replicas' defines the number of replicas (pods) to start.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "'resources', in terms of compute resources, required by this container. For more information, see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "'resources', in terms of compute resources, required by this container. For more information, see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"request": schema.StringAttribute{
													Description:         "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
													MarkdownDescription: "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
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

					"deployment_model": schema.StringAttribute{
						Description:         "'deploymentModel' defines the desired type of deployment for flow processing. Possible values are:<br> - 'Direct' (default) to make the flow processor listen directly from the agents.<br> - 'Kafka' to make flows sent to a Kafka pipeline before consumption by the processor.<br> Kafka can provide better scalability, resiliency, and high availability (for more details, see https://www.redhat.com/en/topics/integration/what-is-apache-kafka).",
						MarkdownDescription: "'deploymentModel' defines the desired type of deployment for flow processing. Possible values are:<br> - 'Direct' (default) to make the flow processor listen directly from the agents.<br> - 'Kafka' to make flows sent to a Kafka pipeline before consumption by the processor.<br> Kafka can provide better scalability, resiliency, and high availability (for more details, see https://www.redhat.com/en/topics/integration/what-is-apache-kafka).",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Direct", "Kafka"),
						},
					},

					"exporters": schema.ListNestedAttribute{
						Description:         "'exporters' define additional optional exporters for custom consumption or storage.",
						MarkdownDescription: "'exporters' define additional optional exporters for custom consumption or storage.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ipfix": schema.SingleNestedAttribute{
									Description:         "IPFIX configuration, such as the IP address and port to send enriched IPFIX flows to.",
									MarkdownDescription: "IPFIX configuration, such as the IP address and port to send enriched IPFIX flows to.",
									Attributes: map[string]schema.Attribute{
										"target_host": schema.StringAttribute{
											Description:         "Address of the IPFIX external receiver",
											MarkdownDescription: "Address of the IPFIX external receiver",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target_port": schema.Int64Attribute{
											Description:         "Port for the IPFIX external receiver",
											MarkdownDescription: "Port for the IPFIX external receiver",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"transport": schema.StringAttribute{
											Description:         "Transport protocol ('TCP' or 'UDP') to be used for the IPFIX connection, defaults to 'TCP'.",
											MarkdownDescription: "Transport protocol ('TCP' or 'UDP') to be used for the IPFIX connection, defaults to 'TCP'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("TCP", "UDP"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"kafka": schema.SingleNestedAttribute{
									Description:         "Kafka configuration, such as the address and topic, to send enriched flows to.",
									MarkdownDescription: "Kafka configuration, such as the address and topic, to send enriched flows to.",
									Attributes: map[string]schema.Attribute{
										"address": schema.StringAttribute{
											Description:         "Address of the Kafka server",
											MarkdownDescription: "Address of the Kafka server",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"sasl": schema.SingleNestedAttribute{
											Description:         "SASL authentication configuration. [Unsupported (*)].",
											MarkdownDescription: "SASL authentication configuration. [Unsupported (*)].",
											Attributes: map[string]schema.Attribute{
												"client_id_reference": schema.SingleNestedAttribute{
													Description:         "Reference to the secret or config map containing the client ID",
													MarkdownDescription: "Reference to the secret or config map containing the client ID",
													Attributes: map[string]schema.Attribute{
														"file": schema.StringAttribute{
															Description:         "File name within the config map or secret.",
															MarkdownDescription: "File name within the config map or secret.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing the file.",
															MarkdownDescription: "Name of the config map or secret containing the file.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the file reference: 'configmap' or 'secret'.",
															MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"client_secret_reference": schema.SingleNestedAttribute{
													Description:         "Reference to the secret or config map containing the client secret",
													MarkdownDescription: "Reference to the secret or config map containing the client secret",
													Attributes: map[string]schema.Attribute{
														"file": schema.StringAttribute{
															Description:         "File name within the config map or secret.",
															MarkdownDescription: "File name within the config map or secret.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing the file.",
															MarkdownDescription: "Name of the config map or secret containing the file.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the file reference: 'configmap' or 'secret'.",
															MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
													Description:         "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
													MarkdownDescription: "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Disabled", "Plain", "ScramSHA512"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
											MarkdownDescription: "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
											Attributes: map[string]schema.Attribute{
												"ca_cert": schema.SingleNestedAttribute{
													Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
													MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
															MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing certificates.",
															MarkdownDescription: "Name of the config map or secret containing certificates.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
															MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"enable": schema.BoolAttribute{
													Description:         "Enable TLS",
													MarkdownDescription: "Enable TLS",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
													MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user_cert": schema.SingleNestedAttribute{
													Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
													MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
															MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing certificates.",
															MarkdownDescription: "Name of the config map or secret containing certificates.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
															MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
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

										"topic": schema.StringAttribute{
											Description:         "Kafka topic to use. It must exist. NetObserv does not create it.",
											MarkdownDescription: "Kafka topic to use. It must exist. NetObserv does not create it.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"open_telemetry": schema.SingleNestedAttribute{
									Description:         "Open telemetry configuration, such as the IP address and port to send enriched logs, metrics and or traces to.",
									MarkdownDescription: "Open telemetry configuration, such as the IP address and port to send enriched logs, metrics and or traces to.",
									Attributes: map[string]schema.Attribute{
										"fields_mapping": schema.ListNestedAttribute{
											Description:         "Custom fields mapping to an OpenTelemetry conformant format. By default, NetObserv format proposal is used: https://github.com/rhobs/observability-data-model/blob/main/network-observability.md#format-proposal . As there is currently no accepted otlp standard for L3/4 network logs, you can freely override it with your own.",
											MarkdownDescription: "Custom fields mapping to an OpenTelemetry conformant format. By default, NetObserv format proposal is used: https://github.com/rhobs/observability-data-model/blob/main/network-observability.md#format-proposal . As there is currently no accepted otlp standard for L3/4 network logs, you can freely override it with your own.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"input": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"multiplier": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"output": schema.StringAttribute{
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

										"headers": schema.MapAttribute{
											Description:         "Headers to add to messages (optional)",
											MarkdownDescription: "Headers to add to messages (optional)",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"logs": schema.SingleNestedAttribute{
											Description:         "Open telemetry configuration for logs.",
											MarkdownDescription: "Open telemetry configuration for logs.",
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													Description:         "Set 'enable' to 'true' to send logs to Open Telemetry receiver.",
													MarkdownDescription: "Set 'enable' to 'true' to send logs to Open Telemetry receiver.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"metrics": schema.SingleNestedAttribute{
											Description:         "Open telemetry configuration for metrics.",
											MarkdownDescription: "Open telemetry configuration for metrics.",
											Attributes: map[string]schema.Attribute{
												"enable": schema.BoolAttribute{
													Description:         "Set 'enable' to 'true' to send metrics to Open Telemetry receiver.",
													MarkdownDescription: "Set 'enable' to 'true' to send metrics to Open Telemetry receiver.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"push_time_interval": schema.StringAttribute{
													Description:         "How often should metrics be sent to collector",
													MarkdownDescription: "How often should metrics be sent to collector",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol of Open Telemetry connection. The available options are 'http' and 'grpc'.",
											MarkdownDescription: "Protocol of Open Telemetry connection. The available options are 'http' and 'grpc'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("http", "grpc"),
											},
										},

										"target_host": schema.StringAttribute{
											Description:         "Address of the Open Telemetry receiver",
											MarkdownDescription: "Address of the Open Telemetry receiver",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target_port": schema.Int64Attribute{
											Description:         "Port for the Open Telemetry receiver",
											MarkdownDescription: "Port for the Open Telemetry receiver",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "TLS client configuration.",
											MarkdownDescription: "TLS client configuration.",
											Attributes: map[string]schema.Attribute{
												"ca_cert": schema.SingleNestedAttribute{
													Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
													MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
															MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing certificates.",
															MarkdownDescription: "Name of the config map or secret containing certificates.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
															MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"enable": schema.BoolAttribute{
													Description:         "Enable TLS",
													MarkdownDescription: "Enable TLS",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
													MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user_cert": schema.SingleNestedAttribute{
													Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
													MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
															MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_key": schema.StringAttribute{
															Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the config map or secret containing certificates.",
															MarkdownDescription: "Name of the config map or secret containing certificates.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
															MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("configmap", "secret"),
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

								"type": schema.StringAttribute{
									Description:         "'type' selects the type of exporters. The available options are 'Kafka' and 'IPFIX'.",
									MarkdownDescription: "'type' selects the type of exporters. The available options are 'Kafka' and 'IPFIX'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Kafka", "IPFIX", "OpenTelemetry"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka": schema.SingleNestedAttribute{
						Description:         "Kafka configuration, allowing to use Kafka as a broker as part of the flow collection pipeline. Available when the 'spec.deploymentModel' is 'Kafka'.",
						MarkdownDescription: "Kafka configuration, allowing to use Kafka as a broker as part of the flow collection pipeline. Available when the 'spec.deploymentModel' is 'Kafka'.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Address of the Kafka server",
								MarkdownDescription: "Address of the Kafka server",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"sasl": schema.SingleNestedAttribute{
								Description:         "SASL authentication configuration. [Unsupported (*)].",
								MarkdownDescription: "SASL authentication configuration. [Unsupported (*)].",
								Attributes: map[string]schema.Attribute{
									"client_id_reference": schema.SingleNestedAttribute{
										Description:         "Reference to the secret or config map containing the client ID",
										MarkdownDescription: "Reference to the secret or config map containing the client ID",
										Attributes: map[string]schema.Attribute{
											"file": schema.StringAttribute{
												Description:         "File name within the config map or secret.",
												MarkdownDescription: "File name within the config map or secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing the file.",
												MarkdownDescription: "Name of the config map or secret containing the file.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the file reference: 'configmap' or 'secret'.",
												MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_secret_reference": schema.SingleNestedAttribute{
										Description:         "Reference to the secret or config map containing the client secret",
										MarkdownDescription: "Reference to the secret or config map containing the client secret",
										Attributes: map[string]schema.Attribute{
											"file": schema.StringAttribute{
												Description:         "File name within the config map or secret.",
												MarkdownDescription: "File name within the config map or secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing the file.",
												MarkdownDescription: "Name of the config map or secret containing the file.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the file reference: 'configmap' or 'secret'.",
												MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
										MarkdownDescription: "Type of SASL authentication to use, or 'Disabled' if SASL is not used",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Plain", "ScramSHA512"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
								MarkdownDescription: "TLS client configuration. When using TLS, verify that the address matches the Kafka port used for TLS, generally 9093.",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.SingleNestedAttribute{
										Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
										MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
												MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing certificates.",
												MarkdownDescription: "Name of the config map or secret containing certificates.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
												MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable": schema.BoolAttribute{
										Description:         "Enable TLS",
										MarkdownDescription: "Enable TLS",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
										MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_cert": schema.SingleNestedAttribute{
										Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
										MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
										Attributes: map[string]schema.Attribute{
											"cert_file": schema.StringAttribute{
												Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
												MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_key": schema.StringAttribute{
												Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the config map or secret containing certificates.",
												MarkdownDescription: "Name of the config map or secret containing certificates.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
												MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("configmap", "secret"),
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

							"topic": schema.StringAttribute{
								Description:         "Kafka topic to use. It must exist. NetObserv does not create it.",
								MarkdownDescription: "Kafka topic to use. It must exist. NetObserv does not create it.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"loki": schema.SingleNestedAttribute{
						Description:         "'loki', the flow store, client settings.",
						MarkdownDescription: "'loki', the flow store, client settings.",
						Attributes: map[string]schema.Attribute{
							"advanced": schema.SingleNestedAttribute{
								Description:         "'advanced' allows setting some aspects of the internal configuration of the Loki clients. This section is aimed mostly for debugging and fine-grained performance optimizations.",
								MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the Loki clients. This section is aimed mostly for debugging and fine-grained performance optimizations.",
								Attributes: map[string]schema.Attribute{
									"static_labels": schema.MapAttribute{
										Description:         "'staticLabels' is a map of common labels to set on each flow in Loki storage.",
										MarkdownDescription: "'staticLabels' is a map of common labels to set on each flow in Loki storage.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"write_max_backoff": schema.StringAttribute{
										Description:         "'writeMaxBackoff' is the maximum backoff time for Loki client connection between retries.",
										MarkdownDescription: "'writeMaxBackoff' is the maximum backoff time for Loki client connection between retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"write_max_retries": schema.Int64Attribute{
										Description:         "'writeMaxRetries' is the maximum number of retries for Loki client connections.",
										MarkdownDescription: "'writeMaxRetries' is the maximum number of retries for Loki client connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"write_min_backoff": schema.StringAttribute{
										Description:         "'writeMinBackoff' is the initial backoff time for Loki client connection between retries.",
										MarkdownDescription: "'writeMinBackoff' is the initial backoff time for Loki client connection between retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Set 'enable' to 'true' to store flows in Loki. The Console plugin can use either Loki or Prometheus as a data source for metrics (see also 'spec.prometheus.querier'), or both. Not all queries are transposable from Loki to Prometheus. Hence, if Loki is disabled, some features of the plugin are disabled as well, such as getting per-pod information or viewing raw flows. If both Prometheus and Loki are enabled, Prometheus takes precedence and Loki is used as a fallback for queries that Prometheus cannot handle. If they are both disabled, the Console plugin is not deployed.",
								MarkdownDescription: "Set 'enable' to 'true' to store flows in Loki. The Console plugin can use either Loki or Prometheus as a data source for metrics (see also 'spec.prometheus.querier'), or both. Not all queries are transposable from Loki to Prometheus. Hence, if Loki is disabled, some features of the plugin are disabled as well, such as getting per-pod information or viewing raw flows. If both Prometheus and Loki are enabled, Prometheus takes precedence and Loki is used as a fallback for queries that Prometheus cannot handle. If they are both disabled, the Console plugin is not deployed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"loki_stack": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'LokiStack' mode. This is useful for an easy Loki Operator configuration. It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'LokiStack' mode. This is useful for an easy Loki Operator configuration. It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of an existing LokiStack resource to use.",
										MarkdownDescription: "Name of an existing LokiStack resource to use.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namespace where this 'LokiStack' resource is located. If omitted, it is assumed to be the same as 'spec.namespace'.",
										MarkdownDescription: "Namespace where this 'LokiStack' resource is located. If omitted, it is assumed to be the same as 'spec.namespace'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"manual": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'Manual' mode. This is the most flexible configuration. It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'Manual' mode. This is the most flexible configuration. It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"auth_token": schema.StringAttribute{
										Description:         "'authToken' describes the way to get a token to authenticate to Loki.<br> - 'Disabled' does not send any token with the request.<br> - 'Forward' forwards the user token for authorization.<br> - 'Host' [deprecated (*)] - uses the local pod service account to authenticate to Loki.<br> When using the Loki Operator, this must be set to 'Forward'.",
										MarkdownDescription: "'authToken' describes the way to get a token to authenticate to Loki.<br> - 'Disabled' does not send any token with the request.<br> - 'Forward' forwards the user token for authorization.<br> - 'Host' [deprecated (*)] - uses the local pod service account to authenticate to Loki.<br> When using the Loki Operator, this must be set to 'Forward'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Host", "Forward"),
										},
									},

									"ingester_url": schema.StringAttribute{
										Description:         "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										MarkdownDescription: "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"querier_url": schema.StringAttribute{
										Description:         "'querierUrl' specifies the address of the Loki querier service. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										MarkdownDescription: "'querierUrl' specifies the address of the Loki querier service. When using the Loki Operator, set it to the Loki gateway service with the 'network' tenant set in path, for example https://loki-gateway-http.netobserv.svc:8080/api/logs/v1/network.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status_tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki status URL.",
										MarkdownDescription: "TLS client configuration for Loki status URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
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

									"status_url": schema.StringAttribute{
										Description:         "'statusUrl' specifies the address of the Loki '/ready', '/metrics' and '/config' endpoints, in case it is different from the Loki querier URL. If empty, the 'querierUrl' value is used. This is useful to show error messages and some context in the frontend. When using the Loki Operator, set it to the Loki HTTP query frontend service, for example https://loki-query-frontend-http.netobserv.svc:3100/. 'statusTLS' configuration is used when 'statusUrl' is set.",
										MarkdownDescription: "'statusUrl' specifies the address of the Loki '/ready', '/metrics' and '/config' endpoints, in case it is different from the Loki querier URL. If empty, the 'querierUrl' value is used. This is useful to show error messages and some context in the frontend. When using the Loki Operator, set it to the Loki HTTP query frontend service, for example https://loki-query-frontend-http.netobserv.svc:3100/. 'statusTLS' configuration is used when 'statusUrl' is set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tenant_id": schema.StringAttribute{
										Description:         "'tenantID' is the Loki 'X-Scope-OrgID' that identifies the tenant for each request. When using the Loki Operator, set it to 'network', which corresponds to a special tenant mode.",
										MarkdownDescription: "'tenantID' is the Loki 'X-Scope-OrgID' that identifies the tenant for each request. When using the Loki Operator, set it to 'network', which corresponds to a special tenant mode.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki URL.",
										MarkdownDescription: "TLS client configuration for Loki URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
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

							"microservices": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'Microservices' mode. Use this option when Loki is installed using the microservices deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#microservices-mode). It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'Microservices' mode. Use this option when Loki is installed using the microservices deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#microservices-mode). It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"ingester_url": schema.StringAttribute{
										Description:         "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to.",
										MarkdownDescription: "'ingesterUrl' is the address of an existing Loki ingester service to push the flows to.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"querier_url": schema.StringAttribute{
										Description:         "'querierURL' specifies the address of the Loki querier service.",
										MarkdownDescription: "'querierURL' specifies the address of the Loki querier service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tenant_id": schema.StringAttribute{
										Description:         "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										MarkdownDescription: "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki URL.",
										MarkdownDescription: "TLS client configuration for Loki URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
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

							"mode": schema.StringAttribute{
								Description:         "'mode' must be set according to the installation mode of Loki:<br> - Use 'LokiStack' when Loki is managed using the Loki Operator<br> - Use 'Monolithic' when Loki is installed as a monolithic workload<br> - Use 'Microservices' when Loki is installed as microservices, but without Loki Operator<br> - Use 'Manual' if none of the options above match your setup<br>",
								MarkdownDescription: "'mode' must be set according to the installation mode of Loki:<br> - Use 'LokiStack' when Loki is managed using the Loki Operator<br> - Use 'Monolithic' when Loki is installed as a monolithic workload<br> - Use 'Microservices' when Loki is installed as microservices, but without Loki Operator<br> - Use 'Manual' if none of the options above match your setup<br>",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Manual", "LokiStack", "Monolithic", "Microservices"),
								},
							},

							"monolithic": schema.SingleNestedAttribute{
								Description:         "Loki configuration for 'Monolithic' mode. Use this option when Loki is installed using the monolithic deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#monolithic-mode). It is ignored for other modes.",
								MarkdownDescription: "Loki configuration for 'Monolithic' mode. Use this option when Loki is installed using the monolithic deployment mode (https://grafana.com/docs/loki/latest/fundamentals/architecture/deployment-modes/#monolithic-mode). It is ignored for other modes.",
								Attributes: map[string]schema.Attribute{
									"tenant_id": schema.StringAttribute{
										Description:         "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										MarkdownDescription: "'tenantID' is the Loki 'X-Scope-OrgID' header that identifies the tenant for each request.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS client configuration for Loki URL.",
										MarkdownDescription: "TLS client configuration for Loki URL.",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.SingleNestedAttribute{
												Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
												MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable": schema.BoolAttribute{
												Description:         "Enable TLS",
												MarkdownDescription: "Enable TLS",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_cert": schema.SingleNestedAttribute{
												Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
												Attributes: map[string]schema.Attribute{
													"cert_file": schema.StringAttribute{
														Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
														MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_key": schema.StringAttribute{
														Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the config map or secret containing certificates.",
														MarkdownDescription: "Name of the config map or secret containing certificates.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
														MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("configmap", "secret"),
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

									"url": schema.StringAttribute{
										Description:         "'url' is the unique address of an existing Loki service that points to both the ingester and the querier.",
										MarkdownDescription: "'url' is the unique address of an existing Loki service that points to both the ingester and the querier.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_timeout": schema.StringAttribute{
								Description:         "'readTimeout' is the maximum console plugin loki query total time limit. A timeout of zero means no timeout.",
								MarkdownDescription: "'readTimeout' is the maximum console plugin loki query total time limit. A timeout of zero means no timeout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_batch_size": schema.Int64Attribute{
								Description:         "'writeBatchSize' is the maximum batch size (in bytes) of Loki logs to accumulate before sending.",
								MarkdownDescription: "'writeBatchSize' is the maximum batch size (in bytes) of Loki logs to accumulate before sending.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"write_batch_wait": schema.StringAttribute{
								Description:         "'writeBatchWait' is the maximum time to wait before sending a Loki batch.",
								MarkdownDescription: "'writeBatchWait' is the maximum time to wait before sending a Loki batch.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"write_timeout": schema.StringAttribute{
								Description:         "'writeTimeout' is the maximum Loki time connection / request limit. A timeout of zero means no timeout.",
								MarkdownDescription: "'writeTimeout' is the maximum Loki time connection / request limit. A timeout of zero means no timeout.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace where NetObserv pods are deployed.",
						MarkdownDescription: "Namespace where NetObserv pods are deployed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"network_policy": schema.SingleNestedAttribute{
						Description:         "'networkPolicy' defines ingress network policy settings for NetObserv components isolation.",
						MarkdownDescription: "'networkPolicy' defines ingress network policy settings for NetObserv components isolation.",
						Attributes: map[string]schema.Attribute{
							"additional_namespaces": schema.ListAttribute{
								Description:         "'additionalNamespaces' contains additional namespaces allowed to connect to the NetObserv namespace. It gives some flexibility in the network policy configuration, however should you need a more specific configuration, you can disable it and install your own instead.",
								MarkdownDescription: "'additionalNamespaces' contains additional namespaces allowed to connect to the NetObserv namespace. It gives some flexibility in the network policy configuration, however should you need a more specific configuration, you can disable it and install your own instead.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Set 'enable' to 'true' to deploy network policies on the namespaces used by NetObserv (main and privileged). It is disabled by default. These network policies better isolate the NetObserv components to prevent undesired connections to them. We recommend you either enable it, or create your own network policy for NetObserv.",
								MarkdownDescription: "Set 'enable' to 'true' to deploy network policies on the namespaces used by NetObserv (main and privileged). It is disabled by default. These network policies better isolate the NetObserv components to prevent undesired connections to them. We recommend you either enable it, or create your own network policy for NetObserv.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"processor": schema.SingleNestedAttribute{
						Description:         "'processor' defines the settings of the component that receives the flows from the agent, enriches them, generates metrics, and forwards them to the Loki persistence layer and/or any available exporter.",
						MarkdownDescription: "'processor' defines the settings of the component that receives the flows from the agent, enriches them, generates metrics, and forwards them to the Loki persistence layer and/or any available exporter.",
						Attributes: map[string]schema.Attribute{
							"add_zone": schema.BoolAttribute{
								Description:         "'addZone' allows availability zone awareness by labelling flows with their source and destination zones. This feature requires the 'topology.kubernetes.io/zone' label to be set on nodes.",
								MarkdownDescription: "'addZone' allows availability zone awareness by labelling flows with their source and destination zones. This feature requires the 'topology.kubernetes.io/zone' label to be set on nodes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"advanced": schema.SingleNestedAttribute{
								Description:         "'advanced' allows setting some aspects of the internal configuration of the flow processor. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								MarkdownDescription: "'advanced' allows setting some aspects of the internal configuration of the flow processor. This section is aimed mostly for debugging and fine-grained performance optimizations, such as 'GOGC' and 'GOMAXPROCS' env vars. Set these values at your own risk.",
								Attributes: map[string]schema.Attribute{
									"conversation_end_timeout": schema.StringAttribute{
										Description:         "'conversationEndTimeout' is the time to wait after a network flow is received, to consider the conversation ended. This delay is ignored when a FIN packet is collected for TCP flows (see 'conversationTerminatingTimeout' instead).",
										MarkdownDescription: "'conversationEndTimeout' is the time to wait after a network flow is received, to consider the conversation ended. This delay is ignored when a FIN packet is collected for TCP flows (see 'conversationTerminatingTimeout' instead).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conversation_heartbeat_interval": schema.StringAttribute{
										Description:         "'conversationHeartbeatInterval' is the time to wait between 'tick' events of a conversation",
										MarkdownDescription: "'conversationHeartbeatInterval' is the time to wait between 'tick' events of a conversation",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"conversation_terminating_timeout": schema.StringAttribute{
										Description:         "'conversationTerminatingTimeout' is the time to wait from detected FIN flag to end a conversation. Only relevant for TCP flows.",
										MarkdownDescription: "'conversationTerminatingTimeout' is the time to wait from detected FIN flag to end a conversation. Only relevant for TCP flows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"drop_unused_fields": schema.BoolAttribute{
										Description:         "'dropUnusedFields' [deprecated (*)] this setting is not used anymore.",
										MarkdownDescription: "'dropUnusedFields' [deprecated (*)] this setting is not used anymore.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_kube_probes": schema.BoolAttribute{
										Description:         "'enableKubeProbes' is a flag to enable or disable Kubernetes liveness and readiness probes",
										MarkdownDescription: "'enableKubeProbes' is a flag to enable or disable Kubernetes liveness and readiness probes",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env": schema.MapAttribute{
										Description:         "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										MarkdownDescription: "'env' allows passing custom environment variables to underlying components. Useful for passing some very concrete performance-tuning options, such as 'GOGC' and 'GOMAXPROCS', that should not be publicly exposed as part of the FlowCollector descriptor, as they are only useful in edge debug or support scenarios.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"health_port": schema.Int64Attribute{
										Description:         "'healthPort' is a collector HTTP port in the Pod that exposes the health check API",
										MarkdownDescription: "'healthPort' is a collector HTTP port in the Pod that exposes the health check API",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Port of the flow collector (host port). By convention, some values are forbidden. It must be greater than 1024 and different from 4500, 4789 and 6081.",
										MarkdownDescription: "Port of the flow collector (host port). By convention, some values are forbidden. It must be greater than 1024 and different from 4500, 4789 and 6081.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1025),
											int64validator.AtMost(65535),
										},
									},

									"profile_port": schema.Int64Attribute{
										Description:         "'profilePort' allows setting up a Go pprof profiler listening to this port",
										MarkdownDescription: "'profilePort' allows setting up a Go pprof profiler listening to this port",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(65535),
										},
									},

									"scheduling": schema.SingleNestedAttribute{
										Description:         "scheduling controls how the pods are scheduled on nodes.",
										MarkdownDescription: "scheduling controls how the pods are scheduled on nodes.",
										Attributes: map[string]schema.Attribute{
											"affinity": schema.SingleNestedAttribute{
												Description:         "If specified, the pod's scheduling constraints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												MarkdownDescription: "If specified, the pod's scheduling constraints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												Attributes: map[string]schema.Attribute{
													"node_affinity": schema.SingleNestedAttribute{
														Description:         "Describes node affinity scheduling rules for the pod.",
														MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
														Attributes: map[string]schema.Attribute{
															"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"preference": schema.SingleNestedAttribute{
																			Description:         "A node selector term, associated with the corresponding weight.",
																			MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's labels.",
																					MarkdownDescription: "A list of node selector requirements by node's labels.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																				"match_fields": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's fields.",
																					MarkdownDescription: "A list of node selector requirements by node's fields.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"weight": schema.Int64Attribute{
																			Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																			MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

															"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
																Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																Attributes: map[string]schema.Attribute{
																	"node_selector_terms": schema.ListNestedAttribute{
																		Description:         "Required. A list of node selector terms. The terms are ORed.",
																		MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's labels.",
																					MarkdownDescription: "A list of node selector requirements by node's labels.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

																				"match_fields": schema.ListNestedAttribute{
																					Description:         "A list of node selector requirements by node's fields.",
																					MarkdownDescription: "A list of node selector requirements by node's fields.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The label key that the selector applies to.",
																								MarkdownDescription: "The label key that the selector applies to.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"values": schema.ListAttribute{
																								Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																			},
																		},
																		Required: true,
																		Optional: false,
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

													"pod_affinity": schema.SingleNestedAttribute{
														Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
														MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
														Attributes: map[string]schema.Attribute{
															"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"pod_affinity_term": schema.SingleNestedAttribute{
																			Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																			MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"match_label_keys": schema.ListAttribute{
																					Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"namespaces": schema.ListAttribute{
																					Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"topology_key": schema.StringAttribute{
																					Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"weight": schema.Int64Attribute{
																			Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																			MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

															"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"match_label_keys": schema.ListAttribute{
																			Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"mismatch_label_keys": schema.ListAttribute{
																			Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_anti_affinity": schema.SingleNestedAttribute{
														Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
														MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
														Attributes: map[string]schema.Attribute{
															"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"pod_affinity_term": schema.SingleNestedAttribute{
																			Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																			MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"match_label_keys": schema.ListAttribute{
																					Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
																					Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																					Attributes: map[string]schema.Attribute{
																						"match_expressions": schema.ListNestedAttribute{
																							Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																							MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																						"match_labels": schema.MapAttribute{
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"namespaces": schema.ListAttribute{
																					Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"topology_key": schema.StringAttribute{
																					Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"weight": schema.Int64Attribute{
																			Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																			MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

															"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"match_label_keys": schema.ListAttribute{
																			Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"mismatch_label_keys": schema.ListAttribute{
																			Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set. This is a beta field and requires enabling MatchLabelKeysInPodAffinity feature gate (enabled by default).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"namespace_selector": schema.SingleNestedAttribute{
																			Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"namespaces": schema.ListAttribute{
																			Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																			MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

											"node_selector": schema.MapAttribute{
												Description:         "'nodeSelector' allows scheduling of pods only onto nodes that have each of the specified labels. For documentation, refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
												MarkdownDescription: "'nodeSelector' allows scheduling of pods only onto nodes that have each of the specified labels. For documentation, refer to https://kubernetes.io/docs/concepts/configuration/assign-pod-node/.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_class_name": schema.StringAttribute{
												Description:         "If specified, indicates the pod's priority. For documentation, refer to https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#how-to-use-priority-and-preemption. If not specified, default priority is used, or zero if there is no default.",
												MarkdownDescription: "If specified, indicates the pod's priority. For documentation, refer to https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#how-to-use-priority-and-preemption. If not specified, default priority is used, or zero if there is no default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tolerations": schema.ListNestedAttribute{
												Description:         "'tolerations' is a list of tolerations that allow the pod to schedule onto nodes with matching taints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												MarkdownDescription: "'tolerations' is a list of tolerations that allow the pod to schedule onto nodes with matching taints. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#scheduling.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"effect": schema.StringAttribute{
															Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"toleration_seconds": schema.Int64Attribute{
															Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
															MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secondary_networks": schema.ListNestedAttribute{
										Description:         "Define secondary networks to be checked for resources identification. In order to guarantee a correct identification, it is important that the indexed values form an unique identifier across the cluster. If there are collisions in the indexes (same index used by several resources), those resources might be wrongly labelled.",
										MarkdownDescription: "Define secondary networks to be checked for resources identification. In order to guarantee a correct identification, it is important that the indexed values form an unique identifier across the cluster. If there are collisions in the indexes (same index used by several resources), those resources might be wrongly labelled.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"index": schema.ListAttribute{
													Description:         "'index' is a list of fields to use for indexing the pods. They should form a unique Pod identifier across the cluster. Can be any of: MAC, IP, Interface",
													MarkdownDescription: "'index' is a list of fields to use for indexing the pods. They should form a unique Pod identifier across the cluster. Can be any of: MAC, IP, Interface",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "'name' should match the network name as visible in the pods annotation 'k8s.v1.cni.cncf.io/network-status'.",
													MarkdownDescription: "'name' should match the network name as visible in the pods annotation 'k8s.v1.cni.cncf.io/network-status'.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_name": schema.StringAttribute{
								Description:         "'clusterName' is the name of the cluster to appear in the flows data. This is useful in a multi-cluster context. When using OpenShift, leave empty to make it automatically determined.",
								MarkdownDescription: "'clusterName' is the name of the cluster to appear in the flows data. This is useful in a multi-cluster context. When using OpenShift, leave empty to make it automatically determined.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								MarkdownDescription: "'imagePullPolicy' is the Kubernetes pull policy for the image defined above",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"kafka_consumer_autoscaler": schema.SingleNestedAttribute{
								Description:         "'kafkaConsumerAutoscaler' is the spec of a horizontal pod autoscaler to set up for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								MarkdownDescription: "'kafkaConsumerAutoscaler' is the spec of a horizontal pod autoscaler to set up for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								Attributes: map[string]schema.Attribute{
									"max_replicas": schema.Int64Attribute{
										Description:         "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										MarkdownDescription: "'maxReplicas' is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics": schema.ListNestedAttribute{
										Description:         "Metrics used by the pod autoscaler. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/horizontal-pod-autoscaler-v2/",
										MarkdownDescription: "Metrics used by the pod autoscaler. For documentation, refer to https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/horizontal-pod-autoscaler-v2/",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"container_resource": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"container": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"external": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_labels": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"object": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"described_object": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"metric": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_labels": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"pods": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"metric": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_labels": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
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
															Required: true,
															Optional: false,
															Computed: false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resource": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"target": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"average_utilization": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"average_value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"type": schema.StringAttribute{
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

									"min_replicas": schema.Int64Attribute{
										Description:         "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										MarkdownDescription: "'minReplicas' is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured. Scaling is active as long as at least one metric value is available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status": schema.StringAttribute{
										Description:         "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										MarkdownDescription: "'status' describes the desired status regarding deploying an horizontal pod autoscaler.<br> - 'Disabled' does not deploy an horizontal pod autoscaler.<br> - 'Enabled' deploys an horizontal pod autoscaler.<br>",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Disabled", "Enabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kafka_consumer_batch_size": schema.Int64Attribute{
								Description:         "'kafkaConsumerBatchSize' indicates to the broker the maximum batch size, in bytes, that the consumer accepts. Ignored when not using Kafka. Default: 10MB.",
								MarkdownDescription: "'kafkaConsumerBatchSize' indicates to the broker the maximum batch size, in bytes, that the consumer accepts. Ignored when not using Kafka. Default: 10MB.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_consumer_queue_capacity": schema.Int64Attribute{
								Description:         "'kafkaConsumerQueueCapacity' defines the capacity of the internal message queue used in the Kafka consumer client. Ignored when not using Kafka.",
								MarkdownDescription: "'kafkaConsumerQueueCapacity' defines the capacity of the internal message queue used in the Kafka consumer client. Ignored when not using Kafka.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_consumer_replicas": schema.Int64Attribute{
								Description:         "'kafkaConsumerReplicas' defines the number of replicas (pods) to start for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								MarkdownDescription: "'kafkaConsumerReplicas' defines the number of replicas (pods) to start for 'flowlogs-pipeline-transformer', which consumes Kafka messages. This setting is ignored when Kafka is disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "'logLevel' of the processor runtime",
								MarkdownDescription: "'logLevel' of the processor runtime",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal", "panic"),
								},
							},

							"log_types": schema.StringAttribute{
								Description:         "'logTypes' defines the desired record types to generate. Possible values are:<br> - 'Flows' (default) to export regular network flows<br> - 'Conversations' to generate events for started conversations, ended conversations as well as periodic 'tick' updates<br> - 'EndedConversations' to generate only ended conversations events<br> - 'All' to generate both network flows and all conversations events<br>",
								MarkdownDescription: "'logTypes' defines the desired record types to generate. Possible values are:<br> - 'Flows' (default) to export regular network flows<br> - 'Conversations' to generate events for started conversations, ended conversations as well as periodic 'tick' updates<br> - 'EndedConversations' to generate only ended conversations events<br> - 'All' to generate both network flows and all conversations events<br>",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Flows", "Conversations", "EndedConversations", "All"),
								},
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "'Metrics' define the processor configuration regarding metrics",
								MarkdownDescription: "'Metrics' define the processor configuration regarding metrics",
								Attributes: map[string]schema.Attribute{
									"disable_alerts": schema.ListAttribute{
										Description:         "'disableAlerts' is a list of alerts that should be disabled. Possible values are:<br> 'NetObservNoFlows', which is triggered when no flows are being observed for a certain period.<br> 'NetObservLokiError', which is triggered when flows are being dropped due to Loki errors.<br>",
										MarkdownDescription: "'disableAlerts' is a list of alerts that should be disabled. Possible values are:<br> 'NetObservNoFlows', which is triggered when no flows are being observed for a certain period.<br> 'NetObservLokiError', which is triggered when flows are being dropped due to Loki errors.<br>",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"include_list": schema.ListAttribute{
										Description:         "'includeList' is a list of metric names to specify which ones to generate. The names correspond to the names in Prometheus without the prefix. For example, 'namespace_egress_packets_total' shows up as 'netobserv_namespace_egress_packets_total' in Prometheus. Note that the more metrics you add, the bigger is the impact on Prometheus workload resources. Metrics enabled by default are: 'namespace_flows_total', 'node_ingress_bytes_total', 'node_egress_bytes_total', 'workload_ingress_bytes_total', 'workload_egress_bytes_total', 'namespace_drop_packets_total' (when 'PacketDrop' feature is enabled), 'namespace_rtt_seconds' (when 'FlowRTT' feature is enabled), 'namespace_dns_latency_seconds' (when 'DNSTracking' feature is enabled). More information, with full list of available metrics: https://github.com/netobserv/network-observability-operator/blob/main/docs/Metrics.md",
										MarkdownDescription: "'includeList' is a list of metric names to specify which ones to generate. The names correspond to the names in Prometheus without the prefix. For example, 'namespace_egress_packets_total' shows up as 'netobserv_namespace_egress_packets_total' in Prometheus. Note that the more metrics you add, the bigger is the impact on Prometheus workload resources. Metrics enabled by default are: 'namespace_flows_total', 'node_ingress_bytes_total', 'node_egress_bytes_total', 'workload_ingress_bytes_total', 'workload_egress_bytes_total', 'namespace_drop_packets_total' (when 'PacketDrop' feature is enabled), 'namespace_rtt_seconds' (when 'FlowRTT' feature is enabled), 'namespace_dns_latency_seconds' (when 'DNSTracking' feature is enabled). More information, with full list of available metrics: https://github.com/netobserv/network-observability-operator/blob/main/docs/Metrics.md",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server": schema.SingleNestedAttribute{
										Description:         "Metrics server endpoint configuration for Prometheus scraper",
										MarkdownDescription: "Metrics server endpoint configuration for Prometheus scraper",
										Attributes: map[string]schema.Attribute{
											"port": schema.Int64Attribute{
												Description:         "The metrics server HTTP port.",
												MarkdownDescription: "The metrics server HTTP port.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
													int64validator.AtMost(65535),
												},
											},

											"tls": schema.SingleNestedAttribute{
												Description:         "TLS configuration.",
												MarkdownDescription: "TLS configuration.",
												Attributes: map[string]schema.Attribute{
													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
														MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the provided certificate. If set to 'true', the 'providedCaFile' field is ignored.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"provided": schema.SingleNestedAttribute{
														Description:         "TLS configuration when 'type' is set to 'Provided'.",
														MarkdownDescription: "TLS configuration when 'type' is set to 'Provided'.",
														Attributes: map[string]schema.Attribute{
															"cert_file": schema.StringAttribute{
																Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
																MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert_key": schema.StringAttribute{
																Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the config map or secret containing certificates.",
																MarkdownDescription: "Name of the config map or secret containing certificates.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
																MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("configmap", "secret"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"provided_ca_file": schema.SingleNestedAttribute{
														Description:         "Reference to the CA file when 'type' is set to 'Provided'.",
														MarkdownDescription: "Reference to the CA file when 'type' is set to 'Provided'.",
														Attributes: map[string]schema.Attribute{
															"file": schema.StringAttribute{
																Description:         "File name within the config map or secret.",
																MarkdownDescription: "File name within the config map or secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the config map or secret containing the file.",
																MarkdownDescription: "Name of the config map or secret containing the file.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																MarkdownDescription: "Namespace of the config map or secret containing the file. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type for the file reference: 'configmap' or 'secret'.",
																MarkdownDescription: "Type for the file reference: 'configmap' or 'secret'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("configmap", "secret"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": schema.StringAttribute{
														Description:         "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. [Unsupported (*)]. - 'Auto' to use OpenShift auto generated certificate using annotations.",
														MarkdownDescription: "Select the type of TLS configuration:<br> - 'Disabled' (default) to not configure TLS for the endpoint. - 'Provided' to manually provide cert file and a key file. [Unsupported (*)]. - 'Auto' to use OpenShift auto generated certificate using annotations.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Disabled", "Provided", "Auto"),
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

							"multi_cluster_deployment": schema.BoolAttribute{
								Description:         "Set 'multiClusterDeployment' to 'true' to enable multi clusters feature. This adds 'clusterName' label to flows data",
								MarkdownDescription: "Set 'multiClusterDeployment' to 'true' to enable multi clusters feature. This adds 'clusterName' label to flows data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "'resources' are the compute resources required by this container. For more information, see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "'resources' are the compute resources required by this container. For more information, see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container. This is an alpha field and requires enabling the DynamicResourceAllocation feature gate. This field is immutable. It can only be set for containers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"request": schema.StringAttribute{
													Description:         "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
													MarkdownDescription: "Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.",
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

									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"subnet_labels": schema.SingleNestedAttribute{
								Description:         "'subnetLabels' allows to define custom labels on subnets and IPs or to enable automatic labelling of recognized subnets in OpenShift, which is used to identify cluster external traffic. When a subnet matches the source or destination IP of a flow, a corresponding field is added: 'SrcSubnetLabel' or 'DstSubnetLabel'.",
								MarkdownDescription: "'subnetLabels' allows to define custom labels on subnets and IPs or to enable automatic labelling of recognized subnets in OpenShift, which is used to identify cluster external traffic. When a subnet matches the source or destination IP of a flow, a corresponding field is added: 'SrcSubnetLabel' or 'DstSubnetLabel'.",
								Attributes: map[string]schema.Attribute{
									"custom_labels": schema.ListNestedAttribute{
										Description:         "'customLabels' allows to customize subnets and IPs labelling, such as to identify cluster-external workloads or web services. If you enable 'openShiftAutoDetect', 'customLabels' can override the detected subnets in case they overlap.",
										MarkdownDescription: "'customLabels' allows to customize subnets and IPs labelling, such as to identify cluster-external workloads or web services. If you enable 'openShiftAutoDetect', 'customLabels' can override the detected subnets in case they overlap.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidrs": schema.ListAttribute{
													Description:         "List of CIDRs, such as '['1.2.3.4/32']'.",
													MarkdownDescription: "List of CIDRs, such as '['1.2.3.4/32']'.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Label name, used to flag matching flows.",
													MarkdownDescription: "Label name, used to flag matching flows.",
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

									"open_shift_auto_detect": schema.BoolAttribute{
										Description:         "'openShiftAutoDetect' allows, when set to 'true', to detect automatically the machines, pods and services subnets based on the OpenShift install configuration and the Cluster Network Operator configuration. Indirectly, this is a way to accurately detect external traffic: flows that are not labeled for those subnets are external to the cluster. Enabled by default on OpenShift.",
										MarkdownDescription: "'openShiftAutoDetect' allows, when set to 'true', to detect automatically the machines, pods and services subnets based on the OpenShift install configuration and the Cluster Network Operator configuration. Indirectly, this is a way to accurately detect external traffic: flows that are not labeled for those subnets are external to the cluster. Enabled by default on OpenShift.",
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

					"prometheus": schema.SingleNestedAttribute{
						Description:         "'prometheus' defines Prometheus settings, such as querier configuration used to fetch metrics from the Console plugin.",
						MarkdownDescription: "'prometheus' defines Prometheus settings, such as querier configuration used to fetch metrics from the Console plugin.",
						Attributes: map[string]schema.Attribute{
							"querier": schema.SingleNestedAttribute{
								Description:         "Prometheus querying configuration, such as client settings, used in the Console plugin.",
								MarkdownDescription: "Prometheus querying configuration, such as client settings, used in the Console plugin.",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "When 'enable' is 'true', the Console plugin queries flow metrics from Prometheus instead of Loki whenever possible. It is enbaled by default: set it to 'false' to disable this feature. The Console plugin can use either Loki or Prometheus as a data source for metrics (see also 'spec.loki'), or both. Not all queries are transposable from Loki to Prometheus. Hence, if Loki is disabled, some features of the plugin are disabled as well, such as getting per-pod information or viewing raw flows. If both Prometheus and Loki are enabled, Prometheus takes precedence and Loki is used as a fallback for queries that Prometheus cannot handle. If they are both disabled, the Console plugin is not deployed.",
										MarkdownDescription: "When 'enable' is 'true', the Console plugin queries flow metrics from Prometheus instead of Loki whenever possible. It is enbaled by default: set it to 'false' to disable this feature. The Console plugin can use either Loki or Prometheus as a data source for metrics (see also 'spec.loki'), or both. Not all queries are transposable from Loki to Prometheus. Hence, if Loki is disabled, some features of the plugin are disabled as well, such as getting per-pod information or viewing raw flows. If both Prometheus and Loki are enabled, Prometheus takes precedence and Loki is used as a fallback for queries that Prometheus cannot handle. If they are both disabled, the Console plugin is not deployed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"manual": schema.SingleNestedAttribute{
										Description:         "Prometheus configuration for 'Manual' mode.",
										MarkdownDescription: "Prometheus configuration for 'Manual' mode.",
										Attributes: map[string]schema.Attribute{
											"forward_user_token": schema.BoolAttribute{
												Description:         "Set 'true' to forward logged in user token in queries to Prometheus",
												MarkdownDescription: "Set 'true' to forward logged in user token in queries to Prometheus",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls": schema.SingleNestedAttribute{
												Description:         "TLS client configuration for Prometheus URL.",
												MarkdownDescription: "TLS client configuration for Prometheus URL.",
												Attributes: map[string]schema.Attribute{
													"ca_cert": schema.SingleNestedAttribute{
														Description:         "'caCert' defines the reference of the certificate for the Certificate Authority",
														MarkdownDescription: "'caCert' defines the reference of the certificate for the Certificate Authority",
														Attributes: map[string]schema.Attribute{
															"cert_file": schema.StringAttribute{
																Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
																MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert_key": schema.StringAttribute{
																Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the config map or secret containing certificates.",
																MarkdownDescription: "Name of the config map or secret containing certificates.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
																MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("configmap", "secret"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"enable": schema.BoolAttribute{
														Description:         "Enable TLS",
														MarkdownDescription: "Enable TLS",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
														MarkdownDescription: "'insecureSkipVerify' allows skipping client-side verification of the server certificate. If set to 'true', the 'caCert' field is ignored.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_cert": schema.SingleNestedAttribute{
														Description:         "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
														MarkdownDescription: "'userCert' defines the user certificate reference and is used for mTLS (you can ignore it when using one-way TLS)",
														Attributes: map[string]schema.Attribute{
															"cert_file": schema.StringAttribute{
																Description:         "'certFile' defines the path to the certificate file name within the config map or secret.",
																MarkdownDescription: "'certFile' defines the path to the certificate file name within the config map or secret.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert_key": schema.StringAttribute{
																Description:         "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																MarkdownDescription: "'certKey' defines the path to the certificate private key file name within the config map or secret. Omit when the key is not necessary.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the config map or secret containing certificates.",
																MarkdownDescription: "Name of the config map or secret containing certificates.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																MarkdownDescription: "Namespace of the config map or secret containing certificates. If omitted, the default is to use the same namespace as where NetObserv is deployed. If the namespace is different, the config map or the secret is copied so that it can be mounted as required.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type for the certificate reference: 'configmap' or 'secret'.",
																MarkdownDescription: "Type for the certificate reference: 'configmap' or 'secret'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("configmap", "secret"),
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

											"url": schema.StringAttribute{
												Description:         "'url' is the address of an existing Prometheus service to use for querying metrics.",
												MarkdownDescription: "'url' is the address of an existing Prometheus service to use for querying metrics.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": schema.StringAttribute{
										Description:         "'mode' must be set according to the type of Prometheus installation that stores NetObserv metrics:<br> - Use 'Auto' to try configuring automatically. In OpenShift, it uses the Thanos querier from OpenShift Cluster Monitoring<br> - Use 'Manual' for a manual setup<br>",
										MarkdownDescription: "'mode' must be set according to the type of Prometheus installation that stores NetObserv metrics:<br> - Use 'Auto' to try configuring automatically. In OpenShift, it uses the Thanos querier from OpenShift Cluster Monitoring<br> - Use 'Manual' for a manual setup<br>",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Manual", "Auto"),
										},
									},

									"timeout": schema.StringAttribute{
										Description:         "'timeout' is the read timeout for console plugin queries to Prometheus. A timeout of zero means no timeout.",
										MarkdownDescription: "'timeout' is the read timeout for console plugin queries to Prometheus. A timeout of zero means no timeout.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *FlowsNetobservIoFlowCollectorV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flows_netobserv_io_flow_collector_v1beta2_manifest")

	var model FlowsNetobservIoFlowCollectorV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("flows.netobserv.io/v1beta2")
	model.Kind = pointer.String("FlowCollector")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
