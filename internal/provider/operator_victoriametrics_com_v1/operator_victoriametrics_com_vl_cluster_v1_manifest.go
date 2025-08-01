/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1

import (
	"context"
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
	_ datasource.DataSource = &OperatorVictoriametricsComVlclusterV1Manifest{}
)

func NewOperatorVictoriametricsComVlclusterV1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVlclusterV1Manifest{}
}

type OperatorVictoriametricsComVlclusterV1Manifest struct{}

type OperatorVictoriametricsComVlclusterV1ManifestData struct {
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
		ClusterDomainName *string `tfsdk:"cluster_domain_name" json:"clusterDomainName,omitempty"`
		ClusterVersion    *string `tfsdk:"cluster_version" json:"clusterVersion,omitempty"`
		ImagePullSecrets  *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		ManagedMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"managed_metadata" json:"managedMetadata,omitempty"`
		Paused               *bool `tfsdk:"paused" json:"paused,omitempty"`
		RequestsLoadBalancer *struct {
			DisableInsertBalancing *bool              `tfsdk:"disable_insert_balancing" json:"disableInsertBalancing,omitempty"`
			DisableSelectBalancing *bool              `tfsdk:"disable_select_balancing" json:"disableSelectBalancing,omitempty"`
			Enabled                *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Spec                   *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"requests_load_balancer" json:"requestsLoadBalancer,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		UseStrictSecurity  *bool   `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
		Vlinsert           *struct {
			Affinity                            *map[string]string   `tfsdk:"affinity" json:"affinity,omitempty"`
			ConfigMaps                          *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
			Containers                          *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
			DisableAutomountServiceAccountToken *bool                `tfsdk:"disable_automount_service_account_token" json:"disableAutomountServiceAccountToken,omitempty"`
			DisableSelfServiceScrape            *bool                `tfsdk:"disable_self_service_scrape" json:"disableSelfServiceScrape,omitempty"`
			DnsConfig                           *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy     *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			ExtraArgs     *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
			ExtraEnvs     *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
			ExtraEnvsFrom *[]struct {
				ConfigMapRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
				Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
				SecretRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"extra_envs_from" json:"extraEnvsFrom,omitempty"`
			HostAliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
			HostNetwork  *bool `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			Host_aliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"host_aliases,omitempty"`
			Hpa   *map[string]string `tfsdk:"hpa" json:"hpa,omitempty"`
			Image *struct {
				PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			InitContainers      *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
			LivenessProbe       *map[string]string   `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			LogFormat           *string              `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogLevel            *string              `tfsdk:"log_level" json:"logLevel,omitempty"`
			MinReadySeconds     *int64               `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
			NodeSelector        *map[string]string   `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Paused              *bool                `tfsdk:"paused" json:"paused,omitempty"`
			PodDisruptionBudget *struct {
				MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
				SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
			} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			PodMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Port              *string `tfsdk:"port" json:"port,omitempty"`
			PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ReadinessGates    *[]struct {
				ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
			ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			ReplicaCount   *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RevisionHistoryLimitCount *int64 `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
			RollingUpdate             *struct {
				MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
			RuntimeClassName  *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName     *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			Secrets           *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext   *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
			ServiceScrapeSpec *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
			ServiceSpec       *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			StartupProbe *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
			SyslogSpec   *struct {
				TcpListeners *[]struct {
					CompressMethod   *string `tfsdk:"compress_method" json:"compressMethod,omitempty"`
					DecolorizeFields *string `tfsdk:"decolorize_fields" json:"decolorizeFields,omitempty"`
					IgnoreFields     *string `tfsdk:"ignore_fields" json:"ignoreFields,omitempty"`
					ListenPort       *int64  `tfsdk:"listen_port" json:"listenPort,omitempty"`
					StreamFields     *string `tfsdk:"stream_fields" json:"streamFields,omitempty"`
					TenantID         *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
					TlsConfig        *struct {
						CertFile   *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						CertSecret *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"cert_secret" json:"certSecret,omitempty"`
						KeyFile   *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
					} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				} `tfsdk:"tcp_listeners" json:"tcpListeners,omitempty"`
				UdpListeners *[]struct {
					CompressMethod   *string `tfsdk:"compress_method" json:"compressMethod,omitempty"`
					DecolorizeFields *string `tfsdk:"decolorize_fields" json:"decolorizeFields,omitempty"`
					IgnoreFields     *string `tfsdk:"ignore_fields" json:"ignoreFields,omitempty"`
					ListenPort       *int64  `tfsdk:"listen_port" json:"listenPort,omitempty"`
					StreamFields     *string `tfsdk:"stream_fields" json:"streamFields,omitempty"`
					TenantID         *string `tfsdk:"tenant_id" json:"tenantID,omitempty"`
				} `tfsdk:"udp_listeners" json:"udpListeners,omitempty"`
			} `tfsdk:"syslog_spec" json:"syslogSpec,omitempty"`
			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			Tolerations                   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
			UpdateStrategy            *string              `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
			UseDefaultResources       *bool                `tfsdk:"use_default_resources" json:"useDefaultResources,omitempty"`
			UseStrictSecurity         *bool                `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
			VolumeMounts              *[]struct {
				MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
				SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"vlinsert" json:"vlinsert,omitempty"`
		Vlselect *struct {
			Affinity                            *map[string]string   `tfsdk:"affinity" json:"affinity,omitempty"`
			ConfigMaps                          *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
			Containers                          *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
			DisableAutomountServiceAccountToken *bool                `tfsdk:"disable_automount_service_account_token" json:"disableAutomountServiceAccountToken,omitempty"`
			DisableSelfServiceScrape            *bool                `tfsdk:"disable_self_service_scrape" json:"disableSelfServiceScrape,omitempty"`
			DnsConfig                           *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy     *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			ExtraArgs     *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
			ExtraEnvs     *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
			ExtraEnvsFrom *[]struct {
				ConfigMapRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
				Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
				SecretRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"extra_envs_from" json:"extraEnvsFrom,omitempty"`
			HostAliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
			HostNetwork  *bool `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			Host_aliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"host_aliases,omitempty"`
			Hpa   *map[string]string `tfsdk:"hpa" json:"hpa,omitempty"`
			Image *struct {
				PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			InitContainers      *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
			LivenessProbe       *map[string]string   `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			LogFormat           *string              `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogLevel            *string              `tfsdk:"log_level" json:"logLevel,omitempty"`
			MinReadySeconds     *int64               `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
			NodeSelector        *map[string]string   `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Paused              *bool                `tfsdk:"paused" json:"paused,omitempty"`
			PodDisruptionBudget *struct {
				MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
				SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
			} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			PodMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Port              *string `tfsdk:"port" json:"port,omitempty"`
			PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ReadinessGates    *[]struct {
				ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
			ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			ReplicaCount   *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RevisionHistoryLimitCount *int64 `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
			RollingUpdate             *struct {
				MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
			RuntimeClassName  *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName     *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			Secrets           *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext   *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
			ServiceScrapeSpec *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
			ServiceSpec       *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			StartupProbe                  *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
			TerminationGracePeriodSeconds *int64             `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			Tolerations                   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
			UpdateStrategy            *string              `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
			UseDefaultResources       *bool                `tfsdk:"use_default_resources" json:"useDefaultResources,omitempty"`
			UseStrictSecurity         *bool                `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
			VolumeMounts              *[]struct {
				MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
				SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"vlselect" json:"vlselect,omitempty"`
		Vlstorage *struct {
			Affinity                            *map[string]string   `tfsdk:"affinity" json:"affinity,omitempty"`
			ClaimTemplates                      *[]map[string]string `tfsdk:"claim_templates" json:"claimTemplates,omitempty"`
			ConfigMaps                          *[]string            `tfsdk:"config_maps" json:"configMaps,omitempty"`
			Containers                          *[]map[string]string `tfsdk:"containers" json:"containers,omitempty"`
			DisableAutomountServiceAccountToken *bool                `tfsdk:"disable_automount_service_account_token" json:"disableAutomountServiceAccountToken,omitempty"`
			DisableSelfServiceScrape            *bool                `tfsdk:"disable_self_service_scrape" json:"disableSelfServiceScrape,omitempty"`
			DnsConfig                           *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy     *string              `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			ExtraArgs     *map[string]string   `tfsdk:"extra_args" json:"extraArgs,omitempty"`
			ExtraEnvs     *[]map[string]string `tfsdk:"extra_envs" json:"extraEnvs,omitempty"`
			ExtraEnvsFrom *[]struct {
				ConfigMapRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
				Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
				SecretRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"extra_envs_from" json:"extraEnvsFrom,omitempty"`
			FutureRetention *string `tfsdk:"future_retention" json:"futureRetention,omitempty"`
			HostAliases     *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
			HostNetwork  *bool `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			Host_aliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"host_aliases,omitempty"`
			Image *struct {
				PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
				Repository *string `tfsdk:"repository" json:"repository,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"image" json:"image,omitempty"`
			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			InitContainers                       *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
			LivenessProbe                        *map[string]string   `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			LogFormat                            *string              `tfsdk:"log_format" json:"logFormat,omitempty"`
			LogIngestedRows                      *bool                `tfsdk:"log_ingested_rows" json:"logIngestedRows,omitempty"`
			LogLevel                             *string              `tfsdk:"log_level" json:"logLevel,omitempty"`
			LogNewStreams                        *bool                `tfsdk:"log_new_streams" json:"logNewStreams,omitempty"`
			MaintenanceInsertNodeIDs             *[]string            `tfsdk:"maintenance_insert_node_i_ds" json:"maintenanceInsertNodeIDs,omitempty"`
			MaintenanceSelectNodeIDs             *[]string            `tfsdk:"maintenance_select_node_i_ds" json:"maintenanceSelectNodeIDs,omitempty"`
			MinReadySeconds                      *int64               `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
			NodeSelector                         *map[string]string   `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Paused                               *bool                `tfsdk:"paused" json:"paused,omitempty"`
			PersistentVolumeClaimRetentionPolicy *struct {
				WhenDeleted *string `tfsdk:"when_deleted" json:"whenDeleted,omitempty"`
				WhenScaled  *string `tfsdk:"when_scaled" json:"whenScaled,omitempty"`
			} `tfsdk:"persistent_volume_claim_retention_policy" json:"persistentVolumeClaimRetentionPolicy,omitempty"`
			PodDisruptionBudget *struct {
				MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
				SelectorLabels *map[string]string `tfsdk:"selector_labels" json:"selectorLabels,omitempty"`
			} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			PodMetadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Port              *string `tfsdk:"port" json:"port,omitempty"`
			PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ReadinessGates    *[]struct {
				ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
			ReadinessProbe *map[string]string `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			ReplicaCount   *int64             `tfsdk:"replica_count" json:"replicaCount,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Request *string `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RetentionMaxDiskSpaceUsageBytes *string `tfsdk:"retention_max_disk_space_usage_bytes" json:"retentionMaxDiskSpaceUsageBytes,omitempty"`
			RetentionPeriod                 *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
			RevisionHistoryLimitCount       *int64  `tfsdk:"revision_history_limit_count" json:"revisionHistoryLimitCount,omitempty"`
			RollingUpdateStrategy           *string `tfsdk:"rolling_update_strategy" json:"rollingUpdateStrategy,omitempty"`
			RollingUpdateStrategyBehavior   *struct {
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			} `tfsdk:"rolling_update_strategy_behavior" json:"rollingUpdateStrategyBehavior,omitempty"`
			RuntimeClassName  *string            `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName     *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			Secrets           *[]string          `tfsdk:"secrets" json:"secrets,omitempty"`
			SecurityContext   *map[string]string `tfsdk:"security_context" json:"securityContext,omitempty"`
			ServiceScrapeSpec *map[string]string `tfsdk:"service_scrape_spec" json:"serviceScrapeSpec,omitempty"`
			ServiceSpec       *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				UseAsDefault *bool              `tfsdk:"use_as_default" json:"useAsDefault,omitempty"`
			} `tfsdk:"service_spec" json:"serviceSpec,omitempty"`
			StartupProbe *map[string]string `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
			Storage      *struct {
				DisableMountSubPath *bool `tfsdk:"disable_mount_sub_path" json:"disableMountSubPath,omitempty"`
				EmptyDir            *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				VolumeClaimTemplate *map[string]string `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			StorageDataPath               *string `tfsdk:"storage_data_path" json:"storageDataPath,omitempty"`
			TerminationGracePeriodSeconds *int64  `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			Tolerations                   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologySpreadConstraints *[]map[string]string `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
			UseDefaultResources       *bool                `tfsdk:"use_default_resources" json:"useDefaultResources,omitempty"`
			UseStrictSecurity         *bool                `tfsdk:"use_strict_security" json:"useStrictSecurity,omitempty"`
			VolumeMounts              *[]struct {
				MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
				SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			Volumes *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"vlstorage" json:"vlstorage,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVlclusterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vl_cluster_v1_manifest"
}

func (r *OperatorVictoriametricsComVlclusterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VLCluster is fast, cost-effective and scalable logs database.",
		MarkdownDescription: "VLCluster is fast, cost-effective and scalable logs database.",
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
				Description:         "VLClusterSpec defines the desired state of VLCluster",
				MarkdownDescription: "VLClusterSpec defines the desired state of VLCluster",
				Attributes: map[string]schema.Attribute{
					"cluster_domain_name": schema.StringAttribute{
						Description:         "ClusterDomainName defines domain name suffix for in-cluster dns addresses aka .cluster.local used by vlinsert and vlselect to build vlstorage address",
						MarkdownDescription: "ClusterDomainName defines domain name suffix for in-cluster dns addresses aka .cluster.local used by vlinsert and vlselect to build vlstorage address",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_version": schema.StringAttribute{
						Description:         "ClusterVersion defines default images tag for all components. it can be overwritten with component specific image.tag value.",
						MarkdownDescription: "ClusterVersion defines default images tag for all components. it can be overwritten with component specific image.tag value.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
									MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

					"managed_metadata": schema.SingleNestedAttribute{
						Description:         "ManagedMetadata defines metadata that will be added to the all objects created by operator for the given CustomResource",
						MarkdownDescription: "ManagedMetadata defines metadata that will be added to the all objects created by operator for the given CustomResource",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
								MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
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

					"paused": schema.BoolAttribute{
						Description:         "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
						MarkdownDescription: "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"requests_load_balancer": schema.SingleNestedAttribute{
						Description:         "RequestsLoadBalancer configures load-balancing for vlinsert and vlselect requests. It helps to evenly spread load across pods. Usually it's not possible with Kubernetes TCP-based services.",
						MarkdownDescription: "RequestsLoadBalancer configures load-balancing for vlinsert and vlselect requests. It helps to evenly spread load across pods. Usually it's not possible with Kubernetes TCP-based services.",
						Attributes: map[string]schema.Attribute{
							"disable_insert_balancing": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_select_balancing": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spec": schema.MapAttribute{
								Description:         "VMAuthLoadBalancerSpec defines configuration spec for VMAuth used as load-balancer for VMCluster component",
								MarkdownDescription: "VMAuthLoadBalancerSpec defines configuration spec for VMAuth used as load-balancer for VMCluster component",
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

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to use to run the VLSelect, VLInsert and VLStorage Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run the VLSelect, VLInsert and VLStorage Pods.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"use_strict_security": schema.BoolAttribute{
						Description:         "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
						MarkdownDescription: "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vlinsert": schema.SingleNestedAttribute{
						Description:         "VLInsert defines vlinsert component configuration at victoria-logs cluster",
						MarkdownDescription: "VLInsert defines vlinsert component configuration at victoria-logs cluster",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.MapAttribute{
								Description:         "Affinity If specified, the pod's scheduling constraints.",
								MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_maps": schema.ListAttribute{
								Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
								MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListAttribute{
								Description:         "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
								MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_automount_service_account_token": schema.BoolAttribute{
								Description:         "DisableAutomountServiceAccountToken whether to disable serviceAccount auto mount by Kubernetes (available from v0.54.0). Operator will conditionally create volumes and volumeMounts for containers if it requires k8s API access. For example, vmagent and vm-config-reloader requires k8s API access. Operator creates volumes with name: 'kube-api-access', which can be used as volumeMount for extraContainers if needed. And also adds VolumeMounts at /var/run/secrets/kubernetes.io/serviceaccount.",
								MarkdownDescription: "DisableAutomountServiceAccountToken whether to disable serviceAccount auto mount by Kubernetes (available from v0.54.0). Operator will conditionally create volumes and volumeMounts for containers if it requires k8s API access. For example, vmagent and vm-config-reloader requires k8s API access. Operator creates volumes with name: 'kube-api-access', which can be used as volumeMount for extraContainers if needed. And also adds VolumeMounts at /var/run/secrets/kubernetes.io/serviceaccount.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_self_service_scrape": schema.BoolAttribute{
								Description:         "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
								MarkdownDescription: "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_config": schema.SingleNestedAttribute{
								Description:         "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
								MarkdownDescription: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
										MarkdownDescription: "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
										Description:         "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
										MarkdownDescription: "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is this DNS resolver option's name. Required.",
													MarkdownDescription: "Name is this DNS resolver option's name. Required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is this DNS resolver option's value.",
													MarkdownDescription: "Value is this DNS resolver option's value.",
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

									"searches": schema.ListAttribute{
										Description:         "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
										MarkdownDescription: "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
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

							"dns_policy": schema.StringAttribute{
								Description:         "DNSPolicy sets DNS policy for the pod",
								MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_args": schema.MapAttribute{
								Description:         "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
								MarkdownDescription: "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs": schema.ListAttribute{
								Description:         "ExtraEnvs that will be passed to the application container",
								MarkdownDescription: "ExtraEnvs that will be passed to the application container",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs_from": schema.ListNestedAttribute{
								Description:         "ExtraEnvsFrom defines source of env variables for the application container could either be secret or configmap",
								MarkdownDescription: "ExtraEnvsFrom defines source of env variables for the application container could either be secret or configmap",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config_map_ref": schema.SingleNestedAttribute{
											Description:         "The ConfigMap to select from",
											MarkdownDescription: "The ConfigMap to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap must be defined",
													MarkdownDescription: "Specify whether the ConfigMap must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"prefix": schema.StringAttribute{
											Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "The Secret to select from",
											MarkdownDescription: "The Secret to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret must be defined",
													MarkdownDescription: "Specify whether the Secret must be defined",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_aliases": schema.ListNestedAttribute{
								Description:         "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
								MarkdownDescription: "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostnames": schema.ListAttribute{
											Description:         "Hostnames for the above IP address.",
											MarkdownDescription: "Hostnames for the above IP address.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "IP address of the host file entry.",
											MarkdownDescription: "IP address of the host file entry.",
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

							"host_network": schema.BoolAttribute{
								Description:         "HostNetwork controls whether the pod may use the node network namespace",
								MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_aliases": schema.ListNestedAttribute{
								Description:         "HostAliasesUnderScore provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork. Has Priority over hostAliases field",
								MarkdownDescription: "HostAliasesUnderScore provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork. Has Priority over hostAliases field",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostnames": schema.ListAttribute{
											Description:         "Hostnames for the above IP address.",
											MarkdownDescription: "Hostnames for the above IP address.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "IP address of the host file entry.",
											MarkdownDescription: "IP address of the host file entry.",
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

							"hpa": schema.MapAttribute{
								Description:         "Configures horizontal pod autoscaling.",
								MarkdownDescription: "Configures horizontal pod autoscaling.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "Image - docker image settings if no specified operator uses default version from operator config",
								MarkdownDescription: "Image - docker image settings if no specified operator uses default version from operator config",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes how to pull docker image",
										MarkdownDescription: "PullPolicy describes how to pull docker image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository contains name of docker image + it's repository if needed",
										MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag contains desired docker image version",
										MarkdownDescription: "Tag contains desired docker image version",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
								Description:         "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
								MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"init_containers": schema.ListAttribute{
								Description:         "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
								MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.MapAttribute{
								Description:         "LivenessProbe that will be added CRD pod",
								MarkdownDescription: "LivenessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_format": schema.StringAttribute{
								Description:         "LogFormat for VLSelect to be configured with. default or json",
								MarkdownDescription: "LogFormat for VLSelect to be configured with. default or json",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("default", "json"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel for VLSelect to be configured with.",
								MarkdownDescription: "LogLevel for VLSelect to be configured with.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
								},
							},

							"min_ready_seconds": schema.Int64Attribute{
								Description:         "MinReadySeconds defines a minimum number of seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
								MarkdownDescription: "MinReadySeconds defines a minimum number of seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
								MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"paused": schema.BoolAttribute{
								Description:         "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
								MarkdownDescription: "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "PodDisruptionBudget created by operator",
								MarkdownDescription: "PodDisruptionBudget created by operator",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_available": schema.StringAttribute{
										Description:         "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
										MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector_labels": schema.MapAttribute{
										Description:         "replaces default labels selector generated by operator it's useful when you need to create custom budget",
										MarkdownDescription: "replaces default labels selector generated by operator it's useful when you need to create custom budget",
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

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata configures Labels and Annotations which are propagated to the VLSelect pods.",
								MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VLSelect pods.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.StringAttribute{
								Description:         "Port listen address",
								MarkdownDescription: "Port listen address",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "PriorityClassName class assigned to the Pods",
								MarkdownDescription: "PriorityClassName class assigned to the Pods",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_gates": schema.ListNestedAttribute{
								Description:         "ReadinessGates defines pod readiness gates",
								MarkdownDescription: "ReadinessGates defines pod readiness gates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition_type": schema.StringAttribute{
											Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
											MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
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

							"readiness_probe": schema.MapAttribute{
								Description:         "ReadinessProbe that will be added CRD pod",
								MarkdownDescription: "ReadinessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_count": schema.Int64Attribute{
								Description:         "ReplicaCount is the expected size of the Application.",
								MarkdownDescription: "ReplicaCount is the expected size of the Application.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
								MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
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

							"revision_history_limit_count": schema.Int64Attribute{
								Description:         "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
								MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update": schema.SingleNestedAttribute{
								Description:         "RollingUpdate - overrides deployment update params.",
								MarkdownDescription: "RollingUpdate - overrides deployment update params.",
								Attributes: map[string]schema.Attribute{
									"max_surge": schema.StringAttribute{
										Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_unavailable": schema.StringAttribute{
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"runtime_class_name": schema.StringAttribute{
								Description:         "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
								MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "SchedulerName - defines kubernetes scheduler name",
								MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListAttribute{
								Description:         "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
								MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.MapAttribute{
								Description:         "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_scrape_spec": schema.MapAttribute{
								Description:         "ServiceScrapeSpec that will be added to vlselect VMServiceScrape spec",
								MarkdownDescription: "ServiceScrapeSpec that will be added to vlselect VMServiceScrape spec",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "ServiceSpec that will be added to vlselect service spec",
								MarkdownDescription: "ServiceSpec that will be added to vlselect service spec",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
										MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.MapAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"use_as_default": schema.BoolAttribute{
										Description:         "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
										MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"startup_probe": schema.MapAttribute{
								Description:         "StartupProbe that will be added to CRD pod",
								MarkdownDescription: "StartupProbe that will be added to CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syslog_spec": schema.SingleNestedAttribute{
								Description:         "SyslogSpec defines syslog listener configuration",
								MarkdownDescription: "SyslogSpec defines syslog listener configuration",
								Attributes: map[string]schema.Attribute{
									"tcp_listeners": schema.ListNestedAttribute{
										Description:         "TCPListeners defines syslog server TCP listener configuration",
										MarkdownDescription: "TCPListeners defines syslog server TCP listener configuration",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"compress_method": schema.StringAttribute{
													Description:         "CompressMethod for syslog messages see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#compression",
													MarkdownDescription: "CompressMethod for syslog messages see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#compression",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(none|zstd|gzip|deflate)$`), ""),
													},
												},

												"decolorize_fields": schema.StringAttribute{
													Description:         "DecolorizeFields to remove ANSI color codes across logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#decolorizing-fields",
													MarkdownDescription: "DecolorizeFields to remove ANSI color codes across logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#decolorizing-fields",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ignore_fields": schema.StringAttribute{
													Description:         "IgnoreFields to ignore at logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#dropping-fields",
													MarkdownDescription: "IgnoreFields to ignore at logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#dropping-fields",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"listen_port": schema.Int64Attribute{
													Description:         "ListenPort defines listen port",
													MarkdownDescription: "ListenPort defines listen port",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"stream_fields": schema.StringAttribute{
													Description:         "StreamFields to use as log stream labels see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#stream-fields",
													MarkdownDescription: "StreamFields to use as log stream labels see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#stream-fields",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tenant_id": schema.StringAttribute{
													Description:         "TenantID for logs ingested in form of accountID:projectID see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#multiple-configs",
													MarkdownDescription: "TenantID for logs ingested in form of accountID:projectID see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#multiple-configs",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tls_config": schema.SingleNestedAttribute{
													Description:         "TLSServerConfig defines VictoriaMetrics TLS configuration for the application's server",
													MarkdownDescription: "TLSServerConfig defines VictoriaMetrics TLS configuration for the application's server",
													Attributes: map[string]schema.Attribute{
														"cert_file": schema.StringAttribute{
															Description:         "CertFile defines path to the pre-mounted file with certificate mutually exclusive with CertSecret",
															MarkdownDescription: "CertFile defines path to the pre-mounted file with certificate mutually exclusive with CertSecret",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_secret": schema.SingleNestedAttribute{
															Description:         "CertSecretRef defines reference for secret with certificate content under given key mutually exclusive with CertFile",
															MarkdownDescription: "CertSecretRef defines reference for secret with certificate content under given key mutually exclusive with CertFile",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from. Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the Secret or its key must be defined",
																	MarkdownDescription: "Specify whether the Secret or its key must be defined",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"key_file": schema.StringAttribute{
															Description:         "KeyFile defines path to the pre-mounted file with certificate key mutually exclusive with KeySecretRef",
															MarkdownDescription: "KeyFile defines path to the pre-mounted file with certificate key mutually exclusive with KeySecretRef",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key_secret": schema.SingleNestedAttribute{
															Description:         "Key defines reference for secret with certificate key content under given key mutually exclusive with KeyFile",
															MarkdownDescription: "Key defines reference for secret with certificate key content under given key mutually exclusive with KeyFile",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from. Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the Secret or its key must be defined",
																	MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"udp_listeners": schema.ListNestedAttribute{
										Description:         "UDPListeners defines syslog server UDP listener configuration",
										MarkdownDescription: "UDPListeners defines syslog server UDP listener configuration",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"compress_method": schema.StringAttribute{
													Description:         "CompressMethod for syslog messages see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#compression",
													MarkdownDescription: "CompressMethod for syslog messages see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#compression",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(none|zstd|gzip|deflate)$`), ""),
													},
												},

												"decolorize_fields": schema.StringAttribute{
													Description:         "DecolorizeFields to remove ANSI color codes across logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#decolorizing-fields",
													MarkdownDescription: "DecolorizeFields to remove ANSI color codes across logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#decolorizing-fields",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ignore_fields": schema.StringAttribute{
													Description:         "IgnoreFields to ignore at logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#dropping-fields",
													MarkdownDescription: "IgnoreFields to ignore at logs see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#dropping-fields",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"listen_port": schema.Int64Attribute{
													Description:         "ListenPort defines listen port",
													MarkdownDescription: "ListenPort defines listen port",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"stream_fields": schema.StringAttribute{
													Description:         "StreamFields to use as log stream labels see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#stream-fields",
													MarkdownDescription: "StreamFields to use as log stream labels see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#stream-fields",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tenant_id": schema.StringAttribute{
													Description:         "TenantID for logs ingested in form of accountID:projectID see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#multiple-configs",
													MarkdownDescription: "TenantID for logs ingested in form of accountID:projectID see https://docs.victoriametrics.com/victorialogs/data-ingestion/syslog/#multiple-configs",
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

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "TerminationGracePeriodSeconds period for container graceful termination",
								MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations If specified, the pod's tolerations.",
								MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
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

							"topology_spread_constraints": schema.ListAttribute{
								Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"update_strategy": schema.StringAttribute{
								Description:         "UpdateStrategy - overrides default update strategy.",
								MarkdownDescription: "UpdateStrategy - overrides default update strategy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Recreate", "RollingUpdate"),
								},
							},

							"use_default_resources": schema.BoolAttribute{
								Description:         "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
								MarkdownDescription: "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_strict_security": schema.BoolAttribute{
								Description:         "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
								MarkdownDescription: "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted. Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted. Must not contain ':'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"recursive_read_only": schema.StringAttribute{
											Description:         "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
											MarkdownDescription: "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

							"volumes": schema.ListAttribute{
								Description:         "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
								MarkdownDescription: "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vlselect": schema.SingleNestedAttribute{
						Description:         "VLSelect defines vlselect component configuration at victoria-logs cluster",
						MarkdownDescription: "VLSelect defines vlselect component configuration at victoria-logs cluster",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.MapAttribute{
								Description:         "Affinity If specified, the pod's scheduling constraints.",
								MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_maps": schema.ListAttribute{
								Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
								MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListAttribute{
								Description:         "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
								MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_automount_service_account_token": schema.BoolAttribute{
								Description:         "DisableAutomountServiceAccountToken whether to disable serviceAccount auto mount by Kubernetes (available from v0.54.0). Operator will conditionally create volumes and volumeMounts for containers if it requires k8s API access. For example, vmagent and vm-config-reloader requires k8s API access. Operator creates volumes with name: 'kube-api-access', which can be used as volumeMount for extraContainers if needed. And also adds VolumeMounts at /var/run/secrets/kubernetes.io/serviceaccount.",
								MarkdownDescription: "DisableAutomountServiceAccountToken whether to disable serviceAccount auto mount by Kubernetes (available from v0.54.0). Operator will conditionally create volumes and volumeMounts for containers if it requires k8s API access. For example, vmagent and vm-config-reloader requires k8s API access. Operator creates volumes with name: 'kube-api-access', which can be used as volumeMount for extraContainers if needed. And also adds VolumeMounts at /var/run/secrets/kubernetes.io/serviceaccount.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_self_service_scrape": schema.BoolAttribute{
								Description:         "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
								MarkdownDescription: "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_config": schema.SingleNestedAttribute{
								Description:         "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
								MarkdownDescription: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
										MarkdownDescription: "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
										Description:         "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
										MarkdownDescription: "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is this DNS resolver option's name. Required.",
													MarkdownDescription: "Name is this DNS resolver option's name. Required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is this DNS resolver option's value.",
													MarkdownDescription: "Value is this DNS resolver option's value.",
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

									"searches": schema.ListAttribute{
										Description:         "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
										MarkdownDescription: "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
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

							"dns_policy": schema.StringAttribute{
								Description:         "DNSPolicy sets DNS policy for the pod",
								MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_args": schema.MapAttribute{
								Description:         "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
								MarkdownDescription: "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs": schema.ListAttribute{
								Description:         "ExtraEnvs that will be passed to the application container",
								MarkdownDescription: "ExtraEnvs that will be passed to the application container",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs_from": schema.ListNestedAttribute{
								Description:         "ExtraEnvsFrom defines source of env variables for the application container could either be secret or configmap",
								MarkdownDescription: "ExtraEnvsFrom defines source of env variables for the application container could either be secret or configmap",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config_map_ref": schema.SingleNestedAttribute{
											Description:         "The ConfigMap to select from",
											MarkdownDescription: "The ConfigMap to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap must be defined",
													MarkdownDescription: "Specify whether the ConfigMap must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"prefix": schema.StringAttribute{
											Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "The Secret to select from",
											MarkdownDescription: "The Secret to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret must be defined",
													MarkdownDescription: "Specify whether the Secret must be defined",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_aliases": schema.ListNestedAttribute{
								Description:         "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
								MarkdownDescription: "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostnames": schema.ListAttribute{
											Description:         "Hostnames for the above IP address.",
											MarkdownDescription: "Hostnames for the above IP address.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "IP address of the host file entry.",
											MarkdownDescription: "IP address of the host file entry.",
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

							"host_network": schema.BoolAttribute{
								Description:         "HostNetwork controls whether the pod may use the node network namespace",
								MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_aliases": schema.ListNestedAttribute{
								Description:         "HostAliasesUnderScore provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork. Has Priority over hostAliases field",
								MarkdownDescription: "HostAliasesUnderScore provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork. Has Priority over hostAliases field",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostnames": schema.ListAttribute{
											Description:         "Hostnames for the above IP address.",
											MarkdownDescription: "Hostnames for the above IP address.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "IP address of the host file entry.",
											MarkdownDescription: "IP address of the host file entry.",
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

							"hpa": schema.MapAttribute{
								Description:         "Configures horizontal pod autoscaling.",
								MarkdownDescription: "Configures horizontal pod autoscaling.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.SingleNestedAttribute{
								Description:         "Image - docker image settings if no specified operator uses default version from operator config",
								MarkdownDescription: "Image - docker image settings if no specified operator uses default version from operator config",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes how to pull docker image",
										MarkdownDescription: "PullPolicy describes how to pull docker image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository contains name of docker image + it's repository if needed",
										MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag contains desired docker image version",
										MarkdownDescription: "Tag contains desired docker image version",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
								Description:         "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
								MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"init_containers": schema.ListAttribute{
								Description:         "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
								MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.MapAttribute{
								Description:         "LivenessProbe that will be added CRD pod",
								MarkdownDescription: "LivenessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_format": schema.StringAttribute{
								Description:         "LogFormat for VLSelect to be configured with. default or json",
								MarkdownDescription: "LogFormat for VLSelect to be configured with. default or json",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("default", "json"),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel for VLSelect to be configured with.",
								MarkdownDescription: "LogLevel for VLSelect to be configured with.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
								},
							},

							"min_ready_seconds": schema.Int64Attribute{
								Description:         "MinReadySeconds defines a minimum number of seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
								MarkdownDescription: "MinReadySeconds defines a minimum number of seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
								MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"paused": schema.BoolAttribute{
								Description:         "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
								MarkdownDescription: "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "PodDisruptionBudget created by operator",
								MarkdownDescription: "PodDisruptionBudget created by operator",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_available": schema.StringAttribute{
										Description:         "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
										MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector_labels": schema.MapAttribute{
										Description:         "replaces default labels selector generated by operator it's useful when you need to create custom budget",
										MarkdownDescription: "replaces default labels selector generated by operator it's useful when you need to create custom budget",
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

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata configures Labels and Annotations which are propagated to the VLSelect pods.",
								MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VLSelect pods.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.StringAttribute{
								Description:         "Port listen address",
								MarkdownDescription: "Port listen address",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "PriorityClassName class assigned to the Pods",
								MarkdownDescription: "PriorityClassName class assigned to the Pods",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_gates": schema.ListNestedAttribute{
								Description:         "ReadinessGates defines pod readiness gates",
								MarkdownDescription: "ReadinessGates defines pod readiness gates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition_type": schema.StringAttribute{
											Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
											MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
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

							"readiness_probe": schema.MapAttribute{
								Description:         "ReadinessProbe that will be added CRD pod",
								MarkdownDescription: "ReadinessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_count": schema.Int64Attribute{
								Description:         "ReplicaCount is the expected size of the Application.",
								MarkdownDescription: "ReplicaCount is the expected size of the Application.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
								MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
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

							"revision_history_limit_count": schema.Int64Attribute{
								Description:         "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
								MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update": schema.SingleNestedAttribute{
								Description:         "RollingUpdate - overrides deployment update params.",
								MarkdownDescription: "RollingUpdate - overrides deployment update params.",
								Attributes: map[string]schema.Attribute{
									"max_surge": schema.StringAttribute{
										Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_unavailable": schema.StringAttribute{
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"runtime_class_name": schema.StringAttribute{
								Description:         "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
								MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "SchedulerName - defines kubernetes scheduler name",
								MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListAttribute{
								Description:         "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
								MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.MapAttribute{
								Description:         "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_scrape_spec": schema.MapAttribute{
								Description:         "ServiceScrapeSpec that will be added to vlselect VMServiceScrape spec",
								MarkdownDescription: "ServiceScrapeSpec that will be added to vlselect VMServiceScrape spec",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "ServiceSpec that will be added to vlselect service spec",
								MarkdownDescription: "ServiceSpec that will be added to vlselect service spec",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
										MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.MapAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"use_as_default": schema.BoolAttribute{
										Description:         "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
										MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"startup_probe": schema.MapAttribute{
								Description:         "StartupProbe that will be added to CRD pod",
								MarkdownDescription: "StartupProbe that will be added to CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "TerminationGracePeriodSeconds period for container graceful termination",
								MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations If specified, the pod's tolerations.",
								MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
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

							"topology_spread_constraints": schema.ListAttribute{
								Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"update_strategy": schema.StringAttribute{
								Description:         "UpdateStrategy - overrides default update strategy.",
								MarkdownDescription: "UpdateStrategy - overrides default update strategy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Recreate", "RollingUpdate"),
								},
							},

							"use_default_resources": schema.BoolAttribute{
								Description:         "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
								MarkdownDescription: "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_strict_security": schema.BoolAttribute{
								Description:         "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
								MarkdownDescription: "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted. Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted. Must not contain ':'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"recursive_read_only": schema.StringAttribute{
											Description:         "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
											MarkdownDescription: "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

							"volumes": schema.ListAttribute{
								Description:         "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
								MarkdownDescription: "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vlstorage": schema.SingleNestedAttribute{
						Description:         "VLStorage defines vlstorage component configuration at victoria-logs cluster",
						MarkdownDescription: "VLStorage defines vlstorage component configuration at victoria-logs cluster",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.MapAttribute{
								Description:         "Affinity If specified, the pod's scheduling constraints.",
								MarkdownDescription: "Affinity If specified, the pod's scheduling constraints.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"claim_templates": schema.ListAttribute{
								Description:         "ClaimTemplates allows adding additional VolumeClaimTemplates for StatefulSet",
								MarkdownDescription: "ClaimTemplates allows adding additional VolumeClaimTemplates for StatefulSet",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config_maps": schema.ListAttribute{
								Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
								MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/configs/CONFIGMAP_NAME folder",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"containers": schema.ListAttribute{
								Description:         "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
								MarkdownDescription: "Containers property allows to inject additions sidecars or to patch existing containers. It can be useful for proxies, backup, etc.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_automount_service_account_token": schema.BoolAttribute{
								Description:         "DisableAutomountServiceAccountToken whether to disable serviceAccount auto mount by Kubernetes (available from v0.54.0). Operator will conditionally create volumes and volumeMounts for containers if it requires k8s API access. For example, vmagent and vm-config-reloader requires k8s API access. Operator creates volumes with name: 'kube-api-access', which can be used as volumeMount for extraContainers if needed. And also adds VolumeMounts at /var/run/secrets/kubernetes.io/serviceaccount.",
								MarkdownDescription: "DisableAutomountServiceAccountToken whether to disable serviceAccount auto mount by Kubernetes (available from v0.54.0). Operator will conditionally create volumes and volumeMounts for containers if it requires k8s API access. For example, vmagent and vm-config-reloader requires k8s API access. Operator creates volumes with name: 'kube-api-access', which can be used as volumeMount for extraContainers if needed. And also adds VolumeMounts at /var/run/secrets/kubernetes.io/serviceaccount.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_self_service_scrape": schema.BoolAttribute{
								Description:         "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
								MarkdownDescription: "DisableSelfServiceScrape controls creation of VMServiceScrape by operator for the application. Has priority over 'VM_DISABLESELFSERVICESCRAPECREATION' operator env variable",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_config": schema.SingleNestedAttribute{
								Description:         "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
								MarkdownDescription: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
										MarkdownDescription: "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
										Description:         "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
										MarkdownDescription: "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is this DNS resolver option's name. Required.",
													MarkdownDescription: "Name is this DNS resolver option's name. Required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is this DNS resolver option's value.",
													MarkdownDescription: "Value is this DNS resolver option's value.",
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

									"searches": schema.ListAttribute{
										Description:         "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
										MarkdownDescription: "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
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

							"dns_policy": schema.StringAttribute{
								Description:         "DNSPolicy sets DNS policy for the pod",
								MarkdownDescription: "DNSPolicy sets DNS policy for the pod",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_args": schema.MapAttribute{
								Description:         "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
								MarkdownDescription: "ExtraArgs that will be passed to the application container for example remoteWrite.tmpDataPath: /tmp",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs": schema.ListAttribute{
								Description:         "ExtraEnvs that will be passed to the application container",
								MarkdownDescription: "ExtraEnvs that will be passed to the application container",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_envs_from": schema.ListNestedAttribute{
								Description:         "ExtraEnvsFrom defines source of env variables for the application container could either be secret or configmap",
								MarkdownDescription: "ExtraEnvsFrom defines source of env variables for the application container could either be secret or configmap",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config_map_ref": schema.SingleNestedAttribute{
											Description:         "The ConfigMap to select from",
											MarkdownDescription: "The ConfigMap to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap must be defined",
													MarkdownDescription: "Specify whether the ConfigMap must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"prefix": schema.StringAttribute{
											Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "The Secret to select from",
											MarkdownDescription: "The Secret to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret must be defined",
													MarkdownDescription: "Specify whether the Secret must be defined",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"future_retention": schema.StringAttribute{
								Description:         "FutureRetention for the stored logs Log entries with timestamps bigger than now+futureRetention are rejected during data ingestion; see https://docs.victoriametrics.com/victorialogs/#retention",
								MarkdownDescription: "FutureRetention for the stored logs Log entries with timestamps bigger than now+futureRetention are rejected during data ingestion; see https://docs.victoriametrics.com/victorialogs/#retention",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+(h|d|w|y)?$`), ""),
								},
							},

							"host_aliases": schema.ListNestedAttribute{
								Description:         "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
								MarkdownDescription: "HostAliases provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostnames": schema.ListAttribute{
											Description:         "Hostnames for the above IP address.",
											MarkdownDescription: "Hostnames for the above IP address.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "IP address of the host file entry.",
											MarkdownDescription: "IP address of the host file entry.",
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

							"host_network": schema.BoolAttribute{
								Description:         "HostNetwork controls whether the pod may use the node network namespace",
								MarkdownDescription: "HostNetwork controls whether the pod may use the node network namespace",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_aliases": schema.ListNestedAttribute{
								Description:         "HostAliasesUnderScore provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork. Has Priority over hostAliases field",
								MarkdownDescription: "HostAliasesUnderScore provides mapping for ip and hostname, that would be propagated to pod, cannot be used with HostNetwork. Has Priority over hostAliases field",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostnames": schema.ListAttribute{
											Description:         "Hostnames for the above IP address.",
											MarkdownDescription: "Hostnames for the above IP address.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ip": schema.StringAttribute{
											Description:         "IP address of the host file entry.",
											MarkdownDescription: "IP address of the host file entry.",
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

							"image": schema.SingleNestedAttribute{
								Description:         "Image - docker image settings if no specified operator uses default version from operator config",
								MarkdownDescription: "Image - docker image settings if no specified operator uses default version from operator config",
								Attributes: map[string]schema.Attribute{
									"pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes how to pull docker image",
										MarkdownDescription: "PullPolicy describes how to pull docker image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"repository": schema.StringAttribute{
										Description:         "Repository contains name of docker image + it's repository if needed",
										MarkdownDescription: "Repository contains name of docker image + it's repository if needed",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "Tag contains desired docker image version",
										MarkdownDescription: "Tag contains desired docker image version",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
								Description:         "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
								MarkdownDescription: "ImagePullSecrets An optional list of references to secrets in the same namespace to use for pulling images from registries see https://kubernetes.io/docs/concepts/containers/images/#referring-to-an-imagepullsecrets-on-a-pod",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"init_containers": schema.ListAttribute{
								Description:         "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
								MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.MapAttribute{
								Description:         "LivenessProbe that will be added CRD pod",
								MarkdownDescription: "LivenessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_format": schema.StringAttribute{
								Description:         "LogFormat for VLStorage to be configured with. default or json",
								MarkdownDescription: "LogFormat for VLStorage to be configured with. default or json",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("default", "json"),
								},
							},

							"log_ingested_rows": schema.BoolAttribute{
								Description:         "Whether to log all the ingested log entries; this can be useful for debugging of data ingestion; see https://docs.victoriametrics.com/victorialogs/data-ingestion/",
								MarkdownDescription: "Whether to log all the ingested log entries; this can be useful for debugging of data ingestion; see https://docs.victoriametrics.com/victorialogs/data-ingestion/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel for VLStorage to be configured with.",
								MarkdownDescription: "LogLevel for VLStorage to be configured with.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("INFO", "WARN", "ERROR", "FATAL", "PANIC"),
								},
							},

							"log_new_streams": schema.BoolAttribute{
								Description:         "LogNewStreams Whether to log creation of new streams; this can be useful for debugging of high cardinality issues with log streams; see https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields",
								MarkdownDescription: "LogNewStreams Whether to log creation of new streams; this can be useful for debugging of high cardinality issues with log streams; see https://docs.victoriametrics.com/victorialogs/keyconcepts/#stream-fields",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maintenance_insert_node_i_ds": schema.ListAttribute{
								Description:         "MaintenanceInsertNodeIDs - excludes given node ids from insert requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc. lets say, you have pod-0, pod-1, pod-2, pod-3. to exclude pod-0 and pod-3 from insert routing, define nodeIDs: [0,3]. Useful at storage expanding, when you want to rebalance some data at cluster.",
								MarkdownDescription: "MaintenanceInsertNodeIDs - excludes given node ids from insert requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc. lets say, you have pod-0, pod-1, pod-2, pod-3. to exclude pod-0 and pod-3 from insert routing, define nodeIDs: [0,3]. Useful at storage expanding, when you want to rebalance some data at cluster.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maintenance_select_node_i_ds": schema.ListAttribute{
								Description:         "MaintenanceInsertNodeIDs - excludes given node ids from select requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc.",
								MarkdownDescription: "MaintenanceInsertNodeIDs - excludes given node ids from select requests routing, must contain pod suffixes - for pod-0, id will be 0 and etc.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_ready_seconds": schema.Int64Attribute{
								Description:         "MinReadySeconds defines a minimum number of seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
								MarkdownDescription: "MinReadySeconds defines a minimum number of seconds to wait before starting update next pod if previous in healthy state Has no effect for VLogs and VMSingle",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector Define which Nodes the Pods are scheduled on.",
								MarkdownDescription: "NodeSelector Define which Nodes the Pods are scheduled on.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"paused": schema.BoolAttribute{
								Description:         "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
								MarkdownDescription: "Paused If set to true all actions on the underlying managed objects are not going to be performed, except for delete actions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"persistent_volume_claim_retention_policy": schema.SingleNestedAttribute{
								Description:         "PersistentVolumeClaimRetentionPolicy allows configuration of PVC retention policy",
								MarkdownDescription: "PersistentVolumeClaimRetentionPolicy allows configuration of PVC retention policy",
								Attributes: map[string]schema.Attribute{
									"when_deleted": schema.StringAttribute{
										Description:         "WhenDeleted specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is deleted. The default policy of 'Retain' causes PVCs to not be affected by StatefulSet deletion. The 'Delete' policy causes those PVCs to be deleted.",
										MarkdownDescription: "WhenDeleted specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is deleted. The default policy of 'Retain' causes PVCs to not be affected by StatefulSet deletion. The 'Delete' policy causes those PVCs to be deleted.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"when_scaled": schema.StringAttribute{
										Description:         "WhenScaled specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is scaled down. The default policy of 'Retain' causes PVCs to not be affected by a scaledown. The 'Delete' policy causes the associated PVCs for any excess pods above the replica count to be deleted.",
										MarkdownDescription: "WhenScaled specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is scaled down. The default policy of 'Retain' causes PVCs to not be affected by a scaledown. The 'Delete' policy causes the associated PVCs for any excess pods above the replica count to be deleted.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "PodDisruptionBudget created by operator",
								MarkdownDescription: "PodDisruptionBudget created by operator",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_available": schema.StringAttribute{
										Description:         "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
										MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod. So for example you can prevent all voluntary evictions by specifying '100%'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector_labels": schema.MapAttribute{
										Description:         "replaces default labels selector generated by operator it's useful when you need to create custom budget",
										MarkdownDescription: "replaces default labels selector generated by operator it's useful when you need to create custom budget",
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

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata configures Labels and Annotations which are propagated to the VLStorage pods.",
								MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the VLStorage pods.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": schema.StringAttribute{
								Description:         "Port listen address",
								MarkdownDescription: "Port listen address",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "PriorityClassName class assigned to the Pods",
								MarkdownDescription: "PriorityClassName class assigned to the Pods",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_gates": schema.ListNestedAttribute{
								Description:         "ReadinessGates defines pod readiness gates",
								MarkdownDescription: "ReadinessGates defines pod readiness gates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition_type": schema.StringAttribute{
											Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
											MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",
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

							"readiness_probe": schema.MapAttribute{
								Description:         "ReadinessProbe that will be added CRD pod",
								MarkdownDescription: "ReadinessProbe that will be added CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_count": schema.Int64Attribute{
								Description:         "ReplicaCount is the expected size of the Application.",
								MarkdownDescription: "ReplicaCount is the expected size of the Application.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
								MarkdownDescription: "Resources container resource request and limits, https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/ if not defined default resources from operator config will be used",
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

							"retention_max_disk_space_usage_bytes": schema.StringAttribute{
								Description:         "RetentionMaxDiskSpaceUsageBytes for the stored logs VictoriaLogs keeps at least two last days of data in order to guarantee that the logs for the last day can be returned in queries. This means that the total disk space usage may exceed the -retention.maxDiskSpaceUsageBytes, if the size of the last two days of data exceeds the -retention.maxDiskSpaceUsageBytes. https://docs.victoriametrics.com/victorialogs/#retention-by-disk-space-usage",
								MarkdownDescription: "RetentionMaxDiskSpaceUsageBytes for the stored logs VictoriaLogs keeps at least two last days of data in order to guarantee that the logs for the last day can be returned in queries. This means that the total disk space usage may exceed the -retention.maxDiskSpaceUsageBytes, if the size of the last two days of data exceeds the -retention.maxDiskSpaceUsageBytes. https://docs.victoriametrics.com/victorialogs/#retention-by-disk-space-usage",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retention_period": schema.StringAttribute{
								Description:         "RetentionPeriod for the stored logs https://docs.victoriametrics.com/victorialogs/#retention",
								MarkdownDescription: "RetentionPeriod for the stored logs https://docs.victoriametrics.com/victorialogs/#retention",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+(h|d|w|y)?$`), ""),
								},
							},

							"revision_history_limit_count": schema.Int64Attribute{
								Description:         "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
								MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback in deployment or maximum number of revisions that will be maintained in the Deployment revision history. Has no effect at StatefulSets Defaults to 10.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update_strategy": schema.StringAttribute{
								Description:         "RollingUpdateStrategy defines strategy for application updates Default is OnDelete, in this case operator handles update process Can be changed for RollingUpdate",
								MarkdownDescription: "RollingUpdateStrategy defines strategy for application updates Default is OnDelete, in this case operator handles update process Can be changed for RollingUpdate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update_strategy_behavior": schema.SingleNestedAttribute{
								Description:         "RollingUpdateStrategyBehavior defines customized behavior for rolling updates. It applies if the RollingUpdateStrategy is set to OnDelete, which is the default.",
								MarkdownDescription: "RollingUpdateStrategyBehavior defines customized behavior for rolling updates. It applies if the RollingUpdateStrategy is set to OnDelete, which is the default.",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "MaxUnavailable defines the maximum number of pods that can be unavailable during the update. It can be specified as an absolute number (e.g. 2) or a percentage of the total pods (e.g. '50%'). For example, if set to 100%, all pods will be upgraded at once, minimizing downtime when needed.",
										MarkdownDescription: "MaxUnavailable defines the maximum number of pods that can be unavailable during the update. It can be specified as an absolute number (e.g. 2) or a percentage of the total pods (e.g. '50%'). For example, if set to 100%, all pods will be upgraded at once, minimizing downtime when needed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"runtime_class_name": schema.StringAttribute{
								Description:         "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
								MarkdownDescription: "RuntimeClassName - defines runtime class for kubernetes pod. https://kubernetes.io/docs/concepts/containers/runtime-class/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "SchedulerName - defines kubernetes scheduler name",
								MarkdownDescription: "SchedulerName - defines kubernetes scheduler name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secrets": schema.ListAttribute{
								Description:         "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
								MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the Application object, which shall be mounted into the Application container at /etc/vm/secrets/SECRET_NAME folder",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.MapAttribute{
								Description:         "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_scrape_spec": schema.MapAttribute{
								Description:         "ServiceScrapeSpec that will be added to vlselect VMServiceScrape spec",
								MarkdownDescription: "ServiceScrapeSpec that will be added to vlselect VMServiceScrape spec",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_spec": schema.SingleNestedAttribute{
								Description:         "ServiceSpec that will be added to vlselect service spec",
								MarkdownDescription: "ServiceSpec that will be added to vlselect service spec",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "EmbeddedObjectMetadata defines objectMeta for additional service.",
										MarkdownDescription: "EmbeddedObjectMetadata defines objectMeta for additional service.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												MarkdownDescription: "Labels Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": schema.MapAttribute{
										Description:         "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"use_as_default": schema.BoolAttribute{
										Description:         "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
										MarkdownDescription: "UseAsDefault applies changes from given service definition to the main object Service Changing from headless service to clusterIP or loadbalancer may break cross-component communication",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"startup_probe": schema.MapAttribute{
								Description:         "StartupProbe that will be added to CRD pod",
								MarkdownDescription: "StartupProbe that will be added to CRD pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage": schema.SingleNestedAttribute{
								Description:         "Storage configures persistent volume for VLStorage",
								MarkdownDescription: "Storage configures persistent volume for VLStorage",
								Attributes: map[string]schema.Attribute{
									"disable_mount_sub_path": schema.BoolAttribute{
										Description:         "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary. DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										MarkdownDescription: "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary. DisableMountSubPath allows to remove any subPath usage in volume mounts.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"empty_dir": schema.SingleNestedAttribute{
										Description:         "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
										MarkdownDescription: "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
										Attributes: map[string]schema.Attribute{
											"medium": schema.StringAttribute{
												Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"size_limit": schema.StringAttribute{
												Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_claim_template": schema.MapAttribute{
										Description:         "A PVC spec to be used by the StatefulSets/Deployments.",
										MarkdownDescription: "A PVC spec to be used by the StatefulSets/Deployments.",
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

							"storage_data_path": schema.StringAttribute{
								Description:         "StorageDataPath - path to storage data",
								MarkdownDescription: "StorageDataPath - path to storage data",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "TerminationGracePeriodSeconds period for container graceful termination",
								MarkdownDescription: "TerminationGracePeriodSeconds period for container graceful termination",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations If specified, the pod's tolerations.",
								MarkdownDescription: "Tolerations If specified, the pod's tolerations.",
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

							"topology_spread_constraints": schema.ListAttribute{
								Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined topology domains https://kubernetes.io/docs/concepts/workloads/pods/pod-topology-spread-constraints/",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_default_resources": schema.BoolAttribute{
								Description:         "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
								MarkdownDescription: "UseDefaultResources controls resource settings By default, operator sets built-in resource requirements",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"use_strict_security": schema.BoolAttribute{
								Description:         "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
								MarkdownDescription: "UseStrictSecurity enables strict security mode for component it restricts disk writes access uses non-root user out of the box drops not needed security permissions",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output Deployment/StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the Application container",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted. Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted. Must not contain ':'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified (which defaults to None).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"recursive_read_only": schema.StringAttribute{
											Description:         "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
											MarkdownDescription: "RecursiveReadOnly specifies whether read-only mounts should be handled recursively. If ReadOnly is false, this field has no meaning and must be unspecified. If ReadOnly is true, and this field is set to Disabled, the mount is not made recursively read-only. If this field is set to IfPossible, the mount is made recursively read-only, if it is supported by the container runtime. If this field is set to Enabled, the mount is made recursively read-only if it is supported by the container runtime, otherwise the pod will not be started and an error will be generated to indicate the reason. If this field is set to IfPossible or Enabled, MountPropagation must be set to None (or be unspecified, which defaults to None). If this field is not specified, it is treated as an equivalent of Disabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

							"volumes": schema.ListAttribute{
								Description:         "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
								MarkdownDescription: "Volumes allows configuration of additional volumes on the output Deployment/StatefulSet definition. Volumes specified will be appended to other volumes that are generated. / +optional",
								ElementType:         types.MapType{ElemType: types.StringType},
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

func (r *OperatorVictoriametricsComVlclusterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vl_cluster_v1_manifest")

	var model OperatorVictoriametricsComVlclusterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1")
	model.Kind = pointer.String("VLCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
