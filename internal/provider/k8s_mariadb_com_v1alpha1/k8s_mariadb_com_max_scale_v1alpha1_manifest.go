/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_mariadb_com_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &K8SMariadbComMaxScaleV1Alpha1Manifest{}
)

func NewK8SMariadbComMaxScaleV1Alpha1Manifest() datasource.DataSource {
	return &K8SMariadbComMaxScaleV1Alpha1Manifest{}
}

type K8SMariadbComMaxScaleV1Alpha1Manifest struct{}

type K8SMariadbComMaxScaleV1Alpha1ManifestData struct {
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
		Admin *struct {
			GuiEnabled *bool  `tfsdk:"gui_enabled" json:"guiEnabled,omitempty"`
			Port       *int64 `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"admin" json:"admin,omitempty"`
		Affinity *struct {
			AntiAffinityEnabled *bool `tfsdk:"anti_affinity_enabled" json:"antiAffinityEnabled,omitempty"`
			NodeAffinity        *struct {
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
						TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
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
					TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		Args *[]string `tfsdk:"args" json:"args,omitempty"`
		Auth *struct {
			AdminPasswordSecretKeyRef *struct {
				Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"admin_password_secret_key_ref" json:"adminPasswordSecretKeyRef,omitempty"`
			AdminUsername              *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
			ClientMaxConnections       *int64  `tfsdk:"client_max_connections" json:"clientMaxConnections,omitempty"`
			ClientPasswordSecretKeyRef *struct {
				Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"client_password_secret_key_ref" json:"clientPasswordSecretKeyRef,omitempty"`
			ClientUsername              *string `tfsdk:"client_username" json:"clientUsername,omitempty"`
			DeleteDefaultAdmin          *bool   `tfsdk:"delete_default_admin" json:"deleteDefaultAdmin,omitempty"`
			Generate                    *bool   `tfsdk:"generate" json:"generate,omitempty"`
			MetricsPasswordSecretKeyRef *struct {
				Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metrics_password_secret_key_ref" json:"metricsPasswordSecretKeyRef,omitempty"`
			MetricsUsername             *string `tfsdk:"metrics_username" json:"metricsUsername,omitempty"`
			MonitorMaxConnections       *int64  `tfsdk:"monitor_max_connections" json:"monitorMaxConnections,omitempty"`
			MonitorPasswordSecretKeyRef *struct {
				Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"monitor_password_secret_key_ref" json:"monitorPasswordSecretKeyRef,omitempty"`
			MonitorUsername            *string `tfsdk:"monitor_username" json:"monitorUsername,omitempty"`
			ServerMaxConnections       *int64  `tfsdk:"server_max_connections" json:"serverMaxConnections,omitempty"`
			ServerPasswordSecretKeyRef *struct {
				Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"server_password_secret_key_ref" json:"serverPasswordSecretKeyRef,omitempty"`
			ServerUsername           *string `tfsdk:"server_username" json:"serverUsername,omitempty"`
			SyncMaxConnections       *int64  `tfsdk:"sync_max_connections" json:"syncMaxConnections,omitempty"`
			SyncPasswordSecretKeyRef *struct {
				Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"sync_password_secret_key_ref" json:"syncPasswordSecretKeyRef,omitempty"`
			SyncUsername *string `tfsdk:"sync_username" json:"syncUsername,omitempty"`
		} `tfsdk:"auth" json:"auth,omitempty"`
		Command *[]string `tfsdk:"command" json:"command,omitempty"`
		Config  *struct {
			Params *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Sync   *struct {
				Database *string `tfsdk:"database" json:"database,omitempty"`
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
				Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"sync" json:"sync,omitempty"`
			VolumeClaimTemplate *struct {
				AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				Metadata    *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
		Connection *struct {
			HealthCheck *struct {
				Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
				RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			Params         *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Port           *int64             `tfsdk:"port" json:"port,omitempty"`
			SecretName     *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			SecretTemplate *struct {
				DatabaseKey *string `tfsdk:"database_key" json:"databaseKey,omitempty"`
				Format      *string `tfsdk:"format" json:"format,omitempty"`
				HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
				Key         *string `tfsdk:"key" json:"key,omitempty"`
				Metadata    *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"secret_template" json:"secretTemplate,omitempty"`
			ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		} `tfsdk:"connection" json:"connection,omitempty"`
		Env *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
				SecretKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		EnvFrom *[]struct {
			ConfigMapRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
			Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"env_from" json:"envFrom,omitempty"`
		GuiKubernetesService *struct {
			AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
			ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
			LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
			Metadata                      *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			SessionAffinity *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
			Type            *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"gui_kubernetes_service" json:"guiKubernetesService,omitempty"`
		Image            *string `tfsdk:"image" json:"image,omitempty"`
		ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		InheritMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"inherit_metadata" json:"inheritMetadata,omitempty"`
		KubernetesService *struct {
			AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
			ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
			LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
			Metadata                      *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			SessionAffinity *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
			Type            *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"kubernetes_service" json:"kubernetesService,omitempty"`
		LivenessProbe *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
			} `tfsdk:"exec" json:"exec,omitempty"`
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			HttpGet          *struct {
				Host   *string `tfsdk:"host" json:"host,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *string `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"http_get" json:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		MariaDbRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			WaitForIt *bool   `tfsdk:"wait_for_it" json:"waitForIt,omitempty"`
		} `tfsdk:"maria_db_ref" json:"mariaDbRef,omitempty"`
		Metrics *struct {
			Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Exporter *struct {
				Affinity *struct {
					AntiAffinityEnabled *bool `tfsdk:"anti_affinity_enabled" json:"antiAffinityEnabled,omitempty"`
					NodeAffinity        *struct {
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
								TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
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
							TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Image            *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				PodMetadata  *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
				PodSecurityContext *struct {
					AppArmorProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
					FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
					FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
					RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
					RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					SeLinuxOptions      *struct {
						Level *string `tfsdk:"level" json:"level,omitempty"`
						Role  *string `tfsdk:"role" json:"role,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						User  *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
					SeccompProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
					SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
				} `tfsdk:"pod_security_context" json:"podSecurityContext,omitempty"`
				Port              *int64  `tfsdk:"port" json:"port,omitempty"`
				PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
				Resources         *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				SecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					Capabilities             *struct {
						Add  *[]string `tfsdk:"add" json:"add,omitempty"`
						Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
					} `tfsdk:"capabilities" json:"capabilities,omitempty"`
					Privileged             *bool  `tfsdk:"privileged" json:"privileged,omitempty"`
					ReadOnlyRootFilesystem *bool  `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
					RunAsGroup             *int64 `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
					RunAsNonRoot           *bool  `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser              *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
				} `tfsdk:"security_context" json:"securityContext,omitempty"`
				Tolerations *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"exporter" json:"exporter,omitempty"`
			ServiceMonitor *struct {
				Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
				JobLabel          *string `tfsdk:"job_label" json:"jobLabel,omitempty"`
				PrometheusRelease *string `tfsdk:"prometheus_release" json:"prometheusRelease,omitempty"`
				ScrapeTimeout     *string `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
			} `tfsdk:"service_monitor" json:"serviceMonitor,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Monitor *struct {
			CooperativeMonitoring *string            `tfsdk:"cooperative_monitoring" json:"cooperativeMonitoring,omitempty"`
			Interval              *string            `tfsdk:"interval" json:"interval,omitempty"`
			Module                *string            `tfsdk:"module" json:"module,omitempty"`
			Name                  *string            `tfsdk:"name" json:"name,omitempty"`
			Params                *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Suspend               *bool              `tfsdk:"suspend" json:"suspend,omitempty"`
		} `tfsdk:"monitor" json:"monitor,omitempty"`
		NodeSelector        *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		PodDisruptionBudget *struct {
			MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			MinAvailable   *string `tfsdk:"min_available" json:"minAvailable,omitempty"`
		} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
		PodMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
		PodSecurityContext *struct {
			AppArmorProfile *struct {
				LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
				Type             *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
			FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
			FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
			RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
			RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
			RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
			SeLinuxOptions      *struct {
				Level *string `tfsdk:"level" json:"level,omitempty"`
				Role  *string `tfsdk:"role" json:"role,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				User  *string `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
			SeccompProfile *struct {
				LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
				Type             *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
			SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
		} `tfsdk:"pod_security_context" json:"podSecurityContext,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		ReadinessProbe    *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
			} `tfsdk:"exec" json:"exec,omitempty"`
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			HttpGet          *struct {
				Host   *string `tfsdk:"host" json:"host,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *string `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"http_get" json:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		Replicas        *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		RequeueInterval *string `tfsdk:"requeue_interval" json:"requeueInterval,omitempty"`
		Resources       *struct {
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		SecurityContext *struct {
			AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
			Capabilities             *struct {
				Add  *[]string `tfsdk:"add" json:"add,omitempty"`
				Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
			} `tfsdk:"capabilities" json:"capabilities,omitempty"`
			Privileged             *bool  `tfsdk:"privileged" json:"privileged,omitempty"`
			ReadOnlyRootFilesystem *bool  `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
			RunAsGroup             *int64 `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
			RunAsNonRoot           *bool  `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
			RunAsUser              *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
		} `tfsdk:"security_context" json:"securityContext,omitempty"`
		Servers *[]struct {
			Address     *string            `tfsdk:"address" json:"address,omitempty"`
			Maintenance *bool              `tfsdk:"maintenance" json:"maintenance,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			Params      *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Port        *int64             `tfsdk:"port" json:"port,omitempty"`
			Protocol    *string            `tfsdk:"protocol" json:"protocol,omitempty"`
		} `tfsdk:"servers" json:"servers,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		Services           *[]struct {
			Listener *struct {
				Name     *string            `tfsdk:"name" json:"name,omitempty"`
				Params   *map[string]string `tfsdk:"params" json:"params,omitempty"`
				Port     *int64             `tfsdk:"port" json:"port,omitempty"`
				Protocol *string            `tfsdk:"protocol" json:"protocol,omitempty"`
				Suspend  *bool              `tfsdk:"suspend" json:"suspend,omitempty"`
			} `tfsdk:"listener" json:"listener,omitempty"`
			Name    *string            `tfsdk:"name" json:"name,omitempty"`
			Params  *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Router  *string            `tfsdk:"router" json:"router,omitempty"`
			Suspend *bool              `tfsdk:"suspend" json:"suspend,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		Suspend     *bool `tfsdk:"suspend" json:"suspend,omitempty"`
		Tolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		TopologySpreadConstraints *[]struct {
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			MatchLabelKeys     *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
			MaxSkew            *int64    `tfsdk:"max_skew" json:"maxSkew,omitempty"`
			MinDomains         *int64    `tfsdk:"min_domains" json:"minDomains,omitempty"`
			NodeAffinityPolicy *string   `tfsdk:"node_affinity_policy" json:"nodeAffinityPolicy,omitempty"`
			NodeTaintsPolicy   *string   `tfsdk:"node_taints_policy" json:"nodeTaintsPolicy,omitempty"`
			TopologyKey        *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
			WhenUnsatisfiable  *string   `tfsdk:"when_unsatisfiable" json:"whenUnsatisfiable,omitempty"`
		} `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
		UpdateStrategy *struct {
			RollingUpdate *struct {
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				Partition      *int64  `tfsdk:"partition" json:"partition,omitempty"`
			} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		VolumeMounts *[]struct {
			MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
		} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SMariadbComMaxScaleV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_mariadb_com_max_scale_v1alpha1_manifest"
}

func (r *K8SMariadbComMaxScaleV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MaxScale is the Schema for the maxscales API. It is used to define MaxScale clusters.",
		MarkdownDescription: "MaxScale is the Schema for the maxscales API. It is used to define MaxScale clusters.",
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
				Description:         "MaxScaleSpec defines the desired state of MaxScale.",
				MarkdownDescription: "MaxScaleSpec defines the desired state of MaxScale.",
				Attributes: map[string]schema.Attribute{
					"admin": schema.SingleNestedAttribute{
						Description:         "Admin configures the admin REST API and GUI.",
						MarkdownDescription: "Admin configures the admin REST API and GUI.",
						Attributes: map[string]schema.Attribute{
							"gui_enabled": schema.BoolAttribute{
								Description:         "GuiEnabled indicates whether the admin GUI should be enabled.",
								MarkdownDescription: "GuiEnabled indicates whether the admin GUI should be enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port where the admin REST API and GUI will be exposed.",
								MarkdownDescription: "Port where the admin REST API and GUI will be exposed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"affinity": schema.SingleNestedAttribute{
						Description:         "Affinity to be used in the Pod.",
						MarkdownDescription: "Affinity to be used in the Pod.",
						Attributes: map[string]schema.Attribute{
							"anti_affinity_enabled": schema.BoolAttribute{
								Description:         "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
								MarkdownDescription: "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_affinity": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeaffinity-v1-core",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeaffinity-v1-core",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"preference": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselectorterm-v1-core",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselectorterm-v1-core",
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
																		Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																		MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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

														"match_fields": schema.ListNestedAttribute{
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
																		Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																		MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
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

									"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
										Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselector-v1-core",
										MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselector-v1-core",
										Attributes: map[string]schema.Attribute{
											"node_selector_terms": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
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
																		Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																		MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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

														"match_fields": schema.ListNestedAttribute{
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
																		Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																		MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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

							"pod_anti_affinity": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
															MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
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
																				Description:         "A label selector operator is the set of operators that can be used in a selector requirement.",
																				MarkdownDescription: "A label selector operator is the set of operators that can be used in a selector requirement.",
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

														"topology_key": schema.StringAttribute{
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

												"weight": schema.Int64Attribute{
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

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
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
																		Description:         "A label selector operator is the set of operators that can be used in a selector requirement.",
																		MarkdownDescription: "A label selector operator is the set of operators that can be used in a selector requirement.",
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

												"topology_key": schema.StringAttribute{
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

					"args": schema.ListAttribute{
						Description:         "Args to be used in the Container.",
						MarkdownDescription: "Args to be used in the Container.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auth": schema.SingleNestedAttribute{
						Description:         "Auth defines the credentials required for MaxScale to connect to MariaDB.",
						MarkdownDescription: "Auth defines the credentials required for MaxScale to connect to MariaDB.",
						Attributes: map[string]schema.Attribute{
							"admin_password_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "AdminPasswordSecretKeyRef is Secret key reference to the admin password to call the admin REST API. It is defaulted if not provided.",
								MarkdownDescription: "AdminPasswordSecretKeyRef is Secret key reference to the admin password to call the admin REST API. It is defaulted if not provided.",
								Attributes: map[string]schema.Attribute{
									"generate": schema.BoolAttribute{
										Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"admin_username": schema.StringAttribute{
								Description:         "AdminUsername is an admin username to call the admin REST API. It is defaulted if not provided.",
								MarkdownDescription: "AdminUsername is an admin username to call the admin REST API. It is defaulted if not provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_max_connections": schema.Int64Attribute{
								Description:         "ClientMaxConnections defines the maximum number of connections that the client can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								MarkdownDescription: "ClientMaxConnections defines the maximum number of connections that the client can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"client_password_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "ClientPasswordSecretKeyRef is Secret key reference to the password to connect to MaxScale. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								MarkdownDescription: "ClientPasswordSecretKeyRef is Secret key reference to the password to connect to MaxScale. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								Attributes: map[string]schema.Attribute{
									"generate": schema.BoolAttribute{
										Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"client_username": schema.StringAttribute{
								Description:         "ClientUsername is the user to connect to MaxScale. It is defaulted if not provided.",
								MarkdownDescription: "ClientUsername is the user to connect to MaxScale. It is defaulted if not provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delete_default_admin": schema.BoolAttribute{
								Description:         "DeleteDefaultAdmin determines whether the default admin user should be deleted after the initial configuration. If not provided, it defaults to true.",
								MarkdownDescription: "DeleteDefaultAdmin determines whether the default admin user should be deleted after the initial configuration. If not provided, it defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"generate": schema.BoolAttribute{
								Description:         "Generate defies whether the operator should generate users and grants for MaxScale to work. It only supports MariaDBs specified via spec.mariaDbRef.",
								MarkdownDescription: "Generate defies whether the operator should generate users and grants for MaxScale to work. It only supports MariaDBs specified via spec.mariaDbRef.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_password_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "MetricsPasswordSecretKeyRef is Secret key reference to the metrics password to call the admib REST API. It is defaulted if metrics are enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								MarkdownDescription: "MetricsPasswordSecretKeyRef is Secret key reference to the metrics password to call the admib REST API. It is defaulted if metrics are enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								Attributes: map[string]schema.Attribute{
									"generate": schema.BoolAttribute{
										Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"metrics_username": schema.StringAttribute{
								Description:         "MetricsUsername is an metrics username to call the REST API. It is defaulted if metrics are enabled.",
								MarkdownDescription: "MetricsUsername is an metrics username to call the REST API. It is defaulted if metrics are enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"monitor_max_connections": schema.Int64Attribute{
								Description:         "MonitorMaxConnections defines the maximum number of connections that the monitor can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								MarkdownDescription: "MonitorMaxConnections defines the maximum number of connections that the monitor can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"monitor_password_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "MonitorPasswordSecretKeyRef is Secret key reference to the password used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								MarkdownDescription: "MonitorPasswordSecretKeyRef is Secret key reference to the password used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								Attributes: map[string]schema.Attribute{
									"generate": schema.BoolAttribute{
										Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"monitor_username": schema.StringAttribute{
								Description:         "MonitorUsername is the user used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided.",
								MarkdownDescription: "MonitorUsername is the user used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_max_connections": schema.Int64Attribute{
								Description:         "ServerMaxConnections defines the maximum number of connections that the server can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								MarkdownDescription: "ServerMaxConnections defines the maximum number of connections that the server can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_password_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "ServerPasswordSecretKeyRef is Secret key reference to the password used by MaxScale to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								MarkdownDescription: "ServerPasswordSecretKeyRef is Secret key reference to the password used by MaxScale to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								Attributes: map[string]schema.Attribute{
									"generate": schema.BoolAttribute{
										Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"server_username": schema.StringAttribute{
								Description:         "ServerUsername is the user used by MaxScale to connect to MariaDB server. It is defaulted if not provided.",
								MarkdownDescription: "ServerUsername is the user used by MaxScale to connect to MariaDB server. It is defaulted if not provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sync_max_connections": schema.Int64Attribute{
								Description:         "SyncMaxConnections defines the maximum number of connections that the sync can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								MarkdownDescription: "SyncMaxConnections defines the maximum number of connections that the sync can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sync_password_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "SyncPasswordSecretKeyRef is Secret key reference to the password used by MaxScale config to connect to MariaDB server. It is defaulted when HA is enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								MarkdownDescription: "SyncPasswordSecretKeyRef is Secret key reference to the password used by MaxScale config to connect to MariaDB server. It is defaulted when HA is enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								Attributes: map[string]schema.Attribute{
									"generate": schema.BoolAttribute{
										Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"sync_username": schema.StringAttribute{
								Description:         "MonitoSyncUsernamerUsername is the user used by MaxScale config sync to connect to MariaDB server. It is defaulted when HA is enabled.",
								MarkdownDescription: "MonitoSyncUsernamerUsername is the user used by MaxScale config sync to connect to MariaDB server. It is defaulted when HA is enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"command": schema.ListAttribute{
						Description:         "Command to be used in the Container.",
						MarkdownDescription: "Command to be used in the Container.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"config": schema.SingleNestedAttribute{
						Description:         "Config defines the MaxScale configuration.",
						MarkdownDescription: "Config defines the MaxScale configuration.",
						Attributes: map[string]schema.Attribute{
							"params": schema.MapAttribute{
								Description:         "Params is a key value pair of parameters to be used in the MaxScale static configuration file. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#global-settings.",
								MarkdownDescription: "Params is a key value pair of parameters to be used in the MaxScale static configuration file. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#global-settings.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sync": schema.SingleNestedAttribute{
								Description:         "Sync defines how to replicate configuration across MaxScale replicas. It is defaulted when HA is enabled.",
								MarkdownDescription: "Sync defines how to replicate configuration across MaxScale replicas. It is defaulted when HA is enabled.",
								Attributes: map[string]schema.Attribute{
									"database": schema.StringAttribute{
										Description:         "Database is the MariaDB logical database where the 'maxscale_config' table will be created in order to persist and synchronize config changes. If not provided, it defaults to 'mysql'.",
										MarkdownDescription: "Database is the MariaDB logical database where the 'maxscale_config' table will be created in order to persist and synchronize config changes. If not provided, it defaults to 'mysql'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "Interval defines the config synchronization interval. It is defaulted if not provided.",
										MarkdownDescription: "Interval defines the config synchronization interval. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"timeout": schema.StringAttribute{
										Description:         "Interval defines the config synchronization timeout. It is defaulted if not provided.",
										MarkdownDescription: "Interval defines the config synchronization timeout. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_claim_template": schema.SingleNestedAttribute{
								Description:         "VolumeClaimTemplate provides a template to define the PVCs for storing MaxScale runtime configuration files. It is defaulted if not provided.",
								MarkdownDescription: "VolumeClaimTemplate provides a template to define the PVCs for storing MaxScale runtime configuration files. It is defaulted if not provided.",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the PVC metadata.",
										MarkdownDescription: "Metadata to be added to the PVC metadata.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"resources": schema.SingleNestedAttribute{
										Description:         "VolumeResourceRequirements describes the storage resource requirements for a volume.",
										MarkdownDescription: "VolumeResourceRequirements describes the storage resource requirements for a volume.",
										Attributes: map[string]schema.Attribute{
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

									"selector": schema.SingleNestedAttribute{
										Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
										MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

									"storage_class_name": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"connection": schema.SingleNestedAttribute{
						Description:         "Connection provides a template to define the Connection for MaxScale.",
						MarkdownDescription: "Connection provides a template to define the Connection for MaxScale.",
						Attributes: map[string]schema.Attribute{
							"health_check": schema.SingleNestedAttribute{
								Description:         "HealthCheck to be used in the Connection.",
								MarkdownDescription: "HealthCheck to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "Interval used to perform health checks.",
										MarkdownDescription: "Interval used to perform health checks.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_interval": schema.StringAttribute{
										Description:         "RetryInterval is the interval used to perform health check retries.",
										MarkdownDescription: "RetryInterval is the interval used to perform health check retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"params": schema.MapAttribute{
								Description:         "Params to be used in the Connection.",
								MarkdownDescription: "Params to be used in the Connection.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								MarkdownDescription: "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName to be used in the Connection.",
								MarkdownDescription: "SecretName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_template": schema.SingleNestedAttribute{
								Description:         "SecretTemplate to be used in the Connection.",
								MarkdownDescription: "SecretTemplate to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"database_key": schema.StringAttribute{
										Description:         "DatabaseKey to be used in the Secret.",
										MarkdownDescription: "DatabaseKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"format": schema.StringAttribute{
										Description:         "Format to be used in the Secret.",
										MarkdownDescription: "Format to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_key": schema.StringAttribute{
										Description:         "HostKey to be used in the Secret.",
										MarkdownDescription: "HostKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "Key to be used in the Secret.",
										MarkdownDescription: "Key to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the Secret object.",
										MarkdownDescription: "Metadata to be added to the Secret object.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"password_key": schema.StringAttribute{
										Description:         "PasswordKey to be used in the Secret.",
										MarkdownDescription: "PasswordKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_key": schema.StringAttribute{
										Description:         "PortKey to be used in the Secret.",
										MarkdownDescription: "PortKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_key": schema.StringAttribute{
										Description:         "UsernameKey to be used in the Secret.",
										MarkdownDescription: "UsernameKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName to be used in the Connection.",
								MarkdownDescription: "ServiceName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"env": schema.ListNestedAttribute{
						Description:         "Env represents the environment variables to be injected in a container.",
						MarkdownDescription: "Env represents the environment variables to be injected in a container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
									MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
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

								"value_from": schema.SingleNestedAttribute{
									Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
									MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
											MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
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

										"field_ref": schema.SingleNestedAttribute{
											Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
											MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
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

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
											MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
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

					"env_from": schema.ListNestedAttribute{
						Description:         "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
						MarkdownDescription: "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_map_ref": schema.SingleNestedAttribute{
									Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

								"prefix": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"gui_kubernetes_service": schema.SingleNestedAttribute{
						Description:         "GuiKubernetesService defines a template for a Kubernetes Service object to connect to MaxScale's GUI.",
						MarkdownDescription: "GuiKubernetesService defines a template for a Kubernetes Service object to connect to MaxScale's GUI.",
						Attributes: map[string]schema.Attribute{
							"allocate_load_balancer_node_ports": schema.BoolAttribute{
								Description:         "AllocateLoadBalancerNodePorts Service field.",
								MarkdownDescription: "AllocateLoadBalancerNodePorts Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_traffic_policy": schema.StringAttribute{
								Description:         "ExternalTrafficPolicy Service field.",
								MarkdownDescription: "ExternalTrafficPolicy Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_ip": schema.StringAttribute{
								Description:         "LoadBalancerIP Service field.",
								MarkdownDescription: "LoadBalancerIP Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges": schema.ListAttribute{
								Description:         "LoadBalancerSourceRanges Service field.",
								MarkdownDescription: "LoadBalancerSourceRanges Service field.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata to be added to the Service metadata.",
								MarkdownDescription: "Metadata to be added to the Service metadata.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations to be added to children resources.",
										MarkdownDescription: "Annotations to be added to children resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels to be added to children resources.",
										MarkdownDescription: "Labels to be added to children resources.",
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

							"session_affinity": schema.StringAttribute{
								Description:         "SessionAffinity Service field.",
								MarkdownDescription: "SessionAffinity Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
								MarkdownDescription: "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Image name to be used by the MaxScale instances. The supported format is '<image>:<tag>'. Only MaxScale official images are supported.",
						MarkdownDescription: "Image name to be used by the MaxScale instances. The supported format is '<image>:<tag>'. Only MaxScale official images are supported.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
						MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
						},
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
						MarkdownDescription: "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
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

					"inherit_metadata": schema.SingleNestedAttribute{
						Description:         "InheritMetadata defines the metadata to be inherited by children resources.",
						MarkdownDescription: "InheritMetadata defines the metadata to be inherited by children resources.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to be added to children resources.",
								MarkdownDescription: "Annotations to be added to children resources.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to be added to children resources.",
								MarkdownDescription: "Labels to be added to children resources.",
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

					"kubernetes_service": schema.SingleNestedAttribute{
						Description:         "KubernetesService defines a template for a Kubernetes Service object to connect to MaxScale.",
						MarkdownDescription: "KubernetesService defines a template for a Kubernetes Service object to connect to MaxScale.",
						Attributes: map[string]schema.Attribute{
							"allocate_load_balancer_node_ports": schema.BoolAttribute{
								Description:         "AllocateLoadBalancerNodePorts Service field.",
								MarkdownDescription: "AllocateLoadBalancerNodePorts Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_traffic_policy": schema.StringAttribute{
								Description:         "ExternalTrafficPolicy Service field.",
								MarkdownDescription: "ExternalTrafficPolicy Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_ip": schema.StringAttribute{
								Description:         "LoadBalancerIP Service field.",
								MarkdownDescription: "LoadBalancerIP Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges": schema.ListAttribute{
								Description:         "LoadBalancerSourceRanges Service field.",
								MarkdownDescription: "LoadBalancerSourceRanges Service field.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata to be added to the Service metadata.",
								MarkdownDescription: "Metadata to be added to the Service metadata.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations to be added to children resources.",
										MarkdownDescription: "Annotations to be added to children resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels to be added to children resources.",
										MarkdownDescription: "Labels to be added to children resources.",
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

							"session_affinity": schema.StringAttribute{
								Description:         "SessionAffinity Service field.",
								MarkdownDescription: "SessionAffinity Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
								MarkdownDescription: "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"liveness_probe": schema.SingleNestedAttribute{
						Description:         "LivenessProbe to be used in the Container.",
						MarkdownDescription: "LivenessProbe to be used in the Container.",
						Attributes: map[string]schema.Attribute{
							"exec": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
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

							"failure_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_get": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
										MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"period_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"success_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_seconds": schema.Int64Attribute{
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

					"maria_db_ref": schema.SingleNestedAttribute{
						Description:         "MariaDBRef is a reference to the MariaDB that MaxScale points to. It is used to initialize the servers field.",
						MarkdownDescription: "MariaDBRef is a reference to the MariaDB that MaxScale points to. It is used to initialize the servers field.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_for_it": schema.BoolAttribute{
								Description:         "WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.",
								MarkdownDescription: "WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.",
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
						Description:         "Metrics configures metrics and how to scrape them.",
						MarkdownDescription: "Metrics configures metrics and how to scrape them.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled is a flag to enable Metrics",
								MarkdownDescription: "Enabled is a flag to enable Metrics",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exporter": schema.SingleNestedAttribute{
								Description:         "Exporter defines the metrics exporter container.",
								MarkdownDescription: "Exporter defines the metrics exporter container.",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "Affinity to be used in the Pod.",
										MarkdownDescription: "Affinity to be used in the Pod.",
										Attributes: map[string]schema.Attribute{
											"anti_affinity_enabled": schema.BoolAttribute{
												Description:         "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
												MarkdownDescription: "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_affinity": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeaffinity-v1-core",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeaffinity-v1-core",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"preference": schema.SingleNestedAttribute{
																	Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselectorterm-v1-core",
																	MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselectorterm-v1-core",
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
																						Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																						MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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

																		"match_fields": schema.ListNestedAttribute{
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
																						Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																						MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
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

													"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
														Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselector-v1-core",
														MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#nodeselector-v1-core",
														Attributes: map[string]schema.Attribute{
															"node_selector_terms": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
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
																						Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																						MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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

																		"match_fields": schema.ListNestedAttribute{
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
																						Description:         "A node selector operator is the set of operators that can be used in a node selector requirement.",
																						MarkdownDescription: "A node selector operator is the set of operators that can be used in a node selector requirement.",
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

											"pod_anti_affinity": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"pod_affinity_term": schema.SingleNestedAttribute{
																	Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																	MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
																			MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
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
																								Description:         "A label selector operator is the set of operators that can be used in a selector requirement.",
																								MarkdownDescription: "A label selector operator is the set of operators that can be used in a selector requirement.",
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

																		"topology_key": schema.StringAttribute{
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

																"weight": schema.Int64Attribute{
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

													"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
																	MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#labelselector-v1-meta",
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
																						Description:         "A label selector operator is the set of operators that can be used in a selector requirement.",
																						MarkdownDescription: "A label selector operator is the set of operators that can be used in a selector requirement.",
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

																"topology_key": schema.StringAttribute{
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

									"image": schema.StringAttribute{
										Description:         "Image name to be used as metrics exporter. The supported format is '<image>:<tag>'. Only mysqld-exporter >= v0.15.0 is supported: https://github.com/prometheus/mysqld_exporter",
										MarkdownDescription: "Image name to be used as metrics exporter. The supported format is '<image>:<tag>'. Only mysqld-exporter >= v0.15.0 is supported: https://github.com/prometheus/mysqld_exporter",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"image_pull_secrets": schema.ListNestedAttribute{
										Description:         "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
										MarkdownDescription: "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector to be used in the Pod.",
										MarkdownDescription: "NodeSelector to be used in the Pod.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_metadata": schema.SingleNestedAttribute{
										Description:         "PodMetadata defines extra metadata for the Pod.",
										MarkdownDescription: "PodMetadata defines extra metadata for the Pod.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"pod_security_context": schema.SingleNestedAttribute{
										Description:         "SecurityContext holds pod-level security attributes and common container settings.",
										MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.",
										Attributes: map[string]schema.Attribute{
											"app_armor_profile": schema.SingleNestedAttribute{
												Description:         "AppArmorProfile defines a pod or container's AppArmor settings.",
												MarkdownDescription: "AppArmorProfile defines a pod or container's AppArmor settings.",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
														MarkdownDescription: "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
														MarkdownDescription: "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"fs_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"fs_group_change_policy": schema.StringAttribute{
												Description:         "PodFSGroupChangePolicy holds policies that will be used for applying fsGroup to a volume when volume is mounted.",
												MarkdownDescription: "PodFSGroupChangePolicy holds policies that will be used for applying fsGroup to a volume when volume is mounted.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "SELinuxOptions are the labels to be applied to the container",
												MarkdownDescription: "SELinuxOptions are the labels to be applied to the container",
												Attributes: map[string]schema.Attribute{
													"level": schema.StringAttribute{
														Description:         "Level is SELinux level label that applies to the container.",
														MarkdownDescription: "Level is SELinux level label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"role": schema.StringAttribute{
														Description:         "Role is a SELinux role label that applies to the container.",
														MarkdownDescription: "Role is a SELinux role label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type is a SELinux type label that applies to the container.",
														MarkdownDescription: "Type is a SELinux type label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user": schema.StringAttribute{
														Description:         "User is a SELinux user label that applies to the container.",
														MarkdownDescription: "User is a SELinux user label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"seccomp_profile": schema.SingleNestedAttribute{
												Description:         "SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.",
												MarkdownDescription: "SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type indicates which kind of seccomp profile will be applied. Valid options are: Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are: Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"supplemental_groups": schema.ListAttribute{
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

									"port": schema.Int64Attribute{
										Description:         "Port where the exporter will be listening for connections.",
										MarkdownDescription: "Port where the exporter will be listening for connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"priority_class_name": schema.StringAttribute{
										Description:         "PriorityClassName to be used in the Pod.",
										MarkdownDescription: "PriorityClassName to be used in the Pod.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resouces describes the compute resource requirements.",
										MarkdownDescription: "Resouces describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "SecurityContext holds container-level security attributes.",
										MarkdownDescription: "SecurityContext holds container-level security attributes.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities": schema.SingleNestedAttribute{
												Description:         "Adds and removes POSIX capabilities from running containers.",
												MarkdownDescription: "Adds and removes POSIX capabilities from running containers.",
												Attributes: map[string]schema.Attribute{
													"add": schema.ListAttribute{
														Description:         "Added capabilities",
														MarkdownDescription: "Added capabilities",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"drop": schema.ListAttribute{
														Description:         "Removed capabilities",
														MarkdownDescription: "Removed capabilities",
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

											"privileged": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_only_root_filesystem": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations to be used in the Pod.",
										MarkdownDescription: "Tolerations to be used in the Pod.",
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

							"service_monitor": schema.SingleNestedAttribute{
								Description:         "ServiceMonitor defines the ServiceMonior object.",
								MarkdownDescription: "ServiceMonitor defines the ServiceMonior object.",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "Interval for scraping metrics.",
										MarkdownDescription: "Interval for scraping metrics.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"job_label": schema.StringAttribute{
										Description:         "JobLabel to add to the ServiceMonitor object.",
										MarkdownDescription: "JobLabel to add to the ServiceMonitor object.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"prometheus_release": schema.StringAttribute{
										Description:         "PrometheusRelease is the release label to add to the ServiceMonitor object.",
										MarkdownDescription: "PrometheusRelease is the release label to add to the ServiceMonitor object.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scrape_timeout": schema.StringAttribute{
										Description:         "ScrapeTimeout defines the timeout for scraping metrics.",
										MarkdownDescription: "ScrapeTimeout defines the timeout for scraping metrics.",
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

					"monitor": schema.SingleNestedAttribute{
						Description:         "Monitor monitors MariaDB server instances. It is required if 'spec.mariaDbRef' is not provided.",
						MarkdownDescription: "Monitor monitors MariaDB server instances. It is required if 'spec.mariaDbRef' is not provided.",
						Attributes: map[string]schema.Attribute{
							"cooperative_monitoring": schema.StringAttribute{
								Description:         "CooperativeMonitoring enables coordination between multiple MaxScale instances running monitors. It is defaulted when HA is enabled.",
								MarkdownDescription: "CooperativeMonitoring enables coordination between multiple MaxScale instances running monitors. It is defaulted when HA is enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("majority_of_all", "majority_of_running"),
								},
							},

							"interval": schema.StringAttribute{
								Description:         "Interval used to monitor MariaDB servers. It is defaulted if not provided.",
								MarkdownDescription: "Interval used to monitor MariaDB servers. It is defaulted if not provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"module": schema.StringAttribute{
								Description:         "Module is the module to use to monitor MariaDB servers. It is mandatory when no MariaDB reference is provided.",
								MarkdownDescription: "Module is the module to use to monitor MariaDB servers. It is mandatory when no MariaDB reference is provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the identifier of the monitor. It is defaulted if not provided.",
								MarkdownDescription: "Name is the identifier of the monitor. It is defaulted if not provided.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"params": schema.MapAttribute{
								Description:         "Params defines extra parameters to pass to the monitor. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-common-monitor-parameters/. Monitor specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-galera-monitor/#galera-monitor-optional-parameters. https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-monitor/#configuration.",
								MarkdownDescription: "Params defines extra parameters to pass to the monitor. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-common-monitor-parameters/. Monitor specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-galera-monitor/#galera-monitor-optional-parameters. https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-monitor/#configuration.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"suspend": schema.BoolAttribute{
								Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
								MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "NodeSelector to be used in the Pod.",
						MarkdownDescription: "NodeSelector to be used in the Pod.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_disruption_budget": schema.SingleNestedAttribute{
						Description:         "PodDisruptionBudget defines the budget for replica availability.",
						MarkdownDescription: "PodDisruptionBudget defines the budget for replica availability.",
						Attributes: map[string]schema.Attribute{
							"max_unavailable": schema.StringAttribute{
								Description:         "MaxUnavailable defines the number of maximum unavailable Pods.",
								MarkdownDescription: "MaxUnavailable defines the number of maximum unavailable Pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_available": schema.StringAttribute{
								Description:         "MinAvailable defines the number of minimum available Pods.",
								MarkdownDescription: "MinAvailable defines the number of minimum available Pods.",
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
						Description:         "PodMetadata defines extra metadata for the Pod.",
						MarkdownDescription: "PodMetadata defines extra metadata for the Pod.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to be added to children resources.",
								MarkdownDescription: "Annotations to be added to children resources.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to be added to children resources.",
								MarkdownDescription: "Labels to be added to children resources.",
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

					"pod_security_context": schema.SingleNestedAttribute{
						Description:         "SecurityContext holds pod-level security attributes and common container settings.",
						MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.",
						Attributes: map[string]schema.Attribute{
							"app_armor_profile": schema.SingleNestedAttribute{
								Description:         "AppArmorProfile defines a pod or container's AppArmor settings.",
								MarkdownDescription: "AppArmorProfile defines a pod or container's AppArmor settings.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
										MarkdownDescription: "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
										MarkdownDescription: "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"fs_group": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fs_group_change_policy": schema.StringAttribute{
								Description:         "PodFSGroupChangePolicy holds policies that will be used for applying fsGroup to a volume when volume is mounted.",
								MarkdownDescription: "PodFSGroupChangePolicy holds policies that will be used for applying fsGroup to a volume when volume is mounted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_group": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_non_root": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_user": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"se_linux_options": schema.SingleNestedAttribute{
								Description:         "SELinuxOptions are the labels to be applied to the container",
								MarkdownDescription: "SELinuxOptions are the labels to be applied to the container",
								Attributes: map[string]schema.Attribute{
									"level": schema.StringAttribute{
										Description:         "Level is SELinux level label that applies to the container.",
										MarkdownDescription: "Level is SELinux level label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"role": schema.StringAttribute{
										Description:         "Role is a SELinux role label that applies to the container.",
										MarkdownDescription: "Role is a SELinux role label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is a SELinux type label that applies to the container.",
										MarkdownDescription: "Type is a SELinux type label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user": schema.StringAttribute{
										Description:         "User is a SELinux user label that applies to the container.",
										MarkdownDescription: "User is a SELinux user label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"seccomp_profile": schema.SingleNestedAttribute{
								Description:         "SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.",
								MarkdownDescription: "SeccompProfile defines a pod/container's seccomp profile settings. Only one profile source may be set.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
										MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must be set if type is 'Localhost'. Must NOT be set for any other type.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of seccomp profile will be applied. Valid options are: Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
										MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are: Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"supplemental_groups": schema.ListAttribute{
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

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName to be used in the Pod.",
						MarkdownDescription: "PriorityClassName to be used in the Pod.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"readiness_probe": schema.SingleNestedAttribute{
						Description:         "ReadinessProbe to be used in the Container.",
						MarkdownDescription: "ReadinessProbe to be used in the Container.",
						Attributes: map[string]schema.Attribute{
							"exec": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
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

							"failure_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_get": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
										MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"period_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"success_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_seconds": schema.Int64Attribute{
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

					"replicas": schema.Int64Attribute{
						Description:         "Replicas indicates the number of desired instances.",
						MarkdownDescription: "Replicas indicates the number of desired instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"requeue_interval": schema.StringAttribute{
						Description:         "RequeueInterval is used to perform requeue reconciliations. If not defined, it defaults to 10s.",
						MarkdownDescription: "RequeueInterval is used to perform requeue reconciliations. If not defined, it defaults to 10s.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resouces describes the compute resource requirements.",
						MarkdownDescription: "Resouces describes the compute resource requirements.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.MapAttribute{
								Description:         "ResourceList is a set of (resource name, quantity) pairs.",
								MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "ResourceList is a set of (resource name, quantity) pairs.",
								MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

					"security_context": schema.SingleNestedAttribute{
						Description:         "SecurityContext holds security configuration that will be applied to a container.",
						MarkdownDescription: "SecurityContext holds security configuration that will be applied to a container.",
						Attributes: map[string]schema.Attribute{
							"allow_privilege_escalation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capabilities": schema.SingleNestedAttribute{
								Description:         "Adds and removes POSIX capabilities from running containers.",
								MarkdownDescription: "Adds and removes POSIX capabilities from running containers.",
								Attributes: map[string]schema.Attribute{
									"add": schema.ListAttribute{
										Description:         "Added capabilities",
										MarkdownDescription: "Added capabilities",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"drop": schema.ListAttribute{
										Description:         "Removed capabilities",
										MarkdownDescription: "Removed capabilities",
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

							"privileged": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read_only_root_filesystem": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_group": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_non_root": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_user": schema.Int64Attribute{
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

					"servers": schema.ListNestedAttribute{
						Description:         "Servers are the MariaDB servers to forward traffic to. It is required if 'spec.mariaDbRef' is not provided.",
						MarkdownDescription: "Servers are the MariaDB servers to forward traffic to. It is required if 'spec.mariaDbRef' is not provided.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"address": schema.StringAttribute{
									Description:         "Address is the network address of the MariaDB server.",
									MarkdownDescription: "Address is the network address of the MariaDB server.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"maintenance": schema.BoolAttribute{
									Description:         "Maintenance indicates whether the server is in maintenance mode.",
									MarkdownDescription: "Maintenance indicates whether the server is in maintenance mode.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the identifier of the MariaDB server.",
									MarkdownDescription: "Name is the identifier of the MariaDB server.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"params": schema.MapAttribute{
									Description:         "Params defines extra parameters to pass to the server. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#server_1.",
									MarkdownDescription: "Params defines extra parameters to pass to the server. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#server_1.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"port": schema.Int64Attribute{
									Description:         "Port is the network port of the MariaDB server. If not provided, it defaults to 3306.",
									MarkdownDescription: "Port is the network port of the MariaDB server. If not provided, it defaults to 3306.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"protocol": schema.StringAttribute{
									Description:         "Protocol is the MaxScale protocol to use when communicating with this MariaDB server. If not provided, it defaults to MariaDBBackend.",
									MarkdownDescription: "Protocol is the MaxScale protocol to use when communicating with this MariaDB server. If not provided, it defaults to MariaDBBackend.",
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

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to be used by the Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to be used by the Pods.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"services": schema.ListNestedAttribute{
						Description:         "Services define how the traffic is forwarded to the MariaDB servers. It is defaulted if not provided.",
						MarkdownDescription: "Services define how the traffic is forwarded to the MariaDB servers. It is defaulted if not provided.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"listener": schema.SingleNestedAttribute{
									Description:         "MaxScaleListener defines how the MaxScale server will listen for connections.",
									MarkdownDescription: "MaxScaleListener defines how the MaxScale server will listen for connections.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name is the identifier of the listener. It is defaulted if not provided",
											MarkdownDescription: "Name is the identifier of the listener. It is defaulted if not provided",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"params": schema.MapAttribute{
											Description:         "Params defines extra parameters to pass to the listener. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#listener_1.",
											MarkdownDescription: "Params defines extra parameters to pass to the listener. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#listener_1.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port is the network port where the MaxScale server will listen.",
											MarkdownDescription: "Port is the network port where the MaxScale server will listen.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"protocol": schema.StringAttribute{
											Description:         "Protocol is the MaxScale protocol to use when communicating with the client. If not provided, it defaults to MariaDBProtocol.",
											MarkdownDescription: "Protocol is the MaxScale protocol to use when communicating with the client. If not provided, it defaults to MariaDBProtocol.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"suspend": schema.BoolAttribute{
											Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
											MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the identifier of the MaxScale service.",
									MarkdownDescription: "Name is the identifier of the MaxScale service.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"params": schema.MapAttribute{
									Description:         "Params defines extra parameters to pass to the service. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#service_1. Router specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-readwritesplit/#configuration. https://mariadb.com/kb/en/mariadb-maxscale-2308-readconnroute/#configuration.",
									MarkdownDescription: "Params defines extra parameters to pass to the service. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#service_1. Router specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-readwritesplit/#configuration. https://mariadb.com/kb/en/mariadb-maxscale-2308-readconnroute/#configuration.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"router": schema.StringAttribute{
									Description:         "Router is the type of router to use.",
									MarkdownDescription: "Router is the type of router to use.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("readwritesplit", "readconnroute"),
									},
								},

								"suspend": schema.BoolAttribute{
									Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
									MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
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

					"suspend": schema.BoolAttribute{
						Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
						MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Tolerations to be used in the Pod.",
						MarkdownDescription: "Tolerations to be used in the Pod.",
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

					"topology_spread_constraints": schema.ListNestedAttribute{
						Description:         "TopologySpreadConstraints to be used in the Pod.",
						MarkdownDescription: "TopologySpreadConstraints to be used in the Pod.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"label_selector": schema.SingleNestedAttribute{
									Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
									MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_skew": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"min_domains": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"node_affinity_policy": schema.StringAttribute{
									Description:         "NodeInclusionPolicy defines the type of node inclusion policy",
									MarkdownDescription: "NodeInclusionPolicy defines the type of node inclusion policy",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"node_taints_policy": schema.StringAttribute{
									Description:         "NodeInclusionPolicy defines the type of node inclusion policy",
									MarkdownDescription: "NodeInclusionPolicy defines the type of node inclusion policy",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"topology_key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"when_unsatisfiable": schema.StringAttribute{
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

					"update_strategy": schema.SingleNestedAttribute{
						Description:         "UpdateStrategy defines the update strategy for the StatefulSet object.",
						MarkdownDescription: "UpdateStrategy defines the update strategy for the StatefulSet object.",
						Attributes: map[string]schema.Attribute{
							"rolling_update": schema.SingleNestedAttribute{
								Description:         "RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.",
								MarkdownDescription: "RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"partition": schema.Int64Attribute{
										Description:         "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
										MarkdownDescription: "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.",
								MarkdownDescription: "Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts to be used in the Container.",
						MarkdownDescription: "VolumeMounts to be used in the Container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
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
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"sub_path": schema.StringAttribute{
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *K8SMariadbComMaxScaleV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_mariadb_com_max_scale_v1alpha1_manifest")

	var model K8SMariadbComMaxScaleV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.mariadb.com/v1alpha1")
	model.Kind = pointer.String("MaxScale")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
