/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package caching_ibm_com_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &CachingIbmComVarnishClusterV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &CachingIbmComVarnishClusterV1Alpha1DataSource{}
)

func NewCachingIbmComVarnishClusterV1Alpha1DataSource() datasource.DataSource {
	return &CachingIbmComVarnishClusterV1Alpha1DataSource{}
}

type CachingIbmComVarnishClusterV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CachingIbmComVarnishClusterV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
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
		Backend *struct {
			Namespaces    *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
			OnlyReady     *bool              `tfsdk:"only_ready" json:"onlyReady,omitempty"`
			Port          *string            `tfsdk:"port" json:"port,omitempty"`
			Selector      *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
			ZoneBalancing *struct {
				Thresholds *[]struct {
					Local     *int64 `tfsdk:"local" json:"local,omitempty"`
					Remote    *int64 `tfsdk:"remote" json:"remote,omitempty"`
					Threshold *int64 `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"thresholds" json:"thresholds,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"zone_balancing" json:"zoneBalancing,omitempty"`
		} `tfsdk:"backend" json:"backend,omitempty"`
		LogFormat  *string `tfsdk:"log_format" json:"logFormat,omitempty"`
		LogLevel   *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		Monitoring *struct {
			GrafanaDashboard *struct {
				DatasourceName *string            `tfsdk:"datasource_name" json:"datasourceName,omitempty"`
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Namespace      *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				Title          *string            `tfsdk:"title" json:"title,omitempty"`
			} `tfsdk:"grafana_dashboard" json:"grafanaDashboard,omitempty"`
			PrometheusServiceMonitor *struct {
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Namespace      *string            `tfsdk:"namespace" json:"namespace,omitempty"`
				ScrapeInterval *string            `tfsdk:"scrape_interval" json:"scrapeInterval,omitempty"`
			} `tfsdk:"prometheus_service_monitor" json:"prometheusServiceMonitor,omitempty"`
		} `tfsdk:"monitoring" json:"monitoring,omitempty"`
		NodeSelector        *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		PodDisruptionBudget *struct {
			MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			MinAvailable   *string `tfsdk:"min_available" json:"minAvailable,omitempty"`
			Selector       *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		Replicas          *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		Service           *struct {
			Annotations               *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			ControllerMetricsNodePort *int64             `tfsdk:"controller_metrics_node_port" json:"controllerMetricsNodePort,omitempty"`
			MetricsNodePort           *int64             `tfsdk:"metrics_node_port" json:"metricsNodePort,omitempty"`
			MetricsPort               *int64             `tfsdk:"metrics_port" json:"metricsPort,omitempty"`
			NodePort                  *int64             `tfsdk:"node_port" json:"nodePort,omitempty"`
			Port                      *int64             `tfsdk:"port" json:"port,omitempty"`
			Type                      *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		Tolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		UpdateStrategy *struct {
			DelayedRollingUpdate *struct {
				DelaySeconds *int64 `tfsdk:"delay_seconds" json:"delaySeconds,omitempty"`
			} `tfsdk:"delayed_rolling_update" json:"delayedRollingUpdate,omitempty"`
			RollingUpdate *struct {
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				Partition      *int64  `tfsdk:"partition" json:"partition,omitempty"`
			} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		Varnish *struct {
			AdmAuth *struct {
				Key        *string `tfsdk:"key" json:"key,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"adm_auth" json:"admAuth,omitempty"`
			Args       *[]string `tfsdk:"args" json:"args,omitempty"`
			Controller *struct {
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Resources       *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"controller" json:"controller,omitempty"`
			EnvFrom *[]struct {
				ConfigMapRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
				Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
				SecretRef *struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"env_from" json:"envFrom,omitempty"`
			ExtraInitContainers *[]struct {
				Args    *[]string `tfsdk:"args" json:"args,omitempty"`
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Env     *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
							Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				EnvFrom *[]struct {
					ConfigMapRef *struct {
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
					Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
					SecretRef *struct {
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"env_from" json:"envFrom,omitempty"`
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Lifecycle       *struct {
					PostStart *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
						HttpGet *struct {
							Host        *string `tfsdk:"host" json:"host,omitempty"`
							HttpHeaders *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
							Path   *string `tfsdk:"path" json:"path,omitempty"`
							Port   *string `tfsdk:"port" json:"port,omitempty"`
							Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
						} `tfsdk:"http_get" json:"httpGet,omitempty"`
						TcpSocket *struct {
							Host *string `tfsdk:"host" json:"host,omitempty"`
							Port *string `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					} `tfsdk:"post_start" json:"postStart,omitempty"`
					PreStop *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" json:"command,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
						HttpGet *struct {
							Host        *string `tfsdk:"host" json:"host,omitempty"`
							HttpHeaders *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
							Path   *string `tfsdk:"path" json:"path,omitempty"`
							Port   *string `tfsdk:"port" json:"port,omitempty"`
							Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
						} `tfsdk:"http_get" json:"httpGet,omitempty"`
						TcpSocket *struct {
							Host *string `tfsdk:"host" json:"host,omitempty"`
							Port *string `tfsdk:"port" json:"port,omitempty"`
						} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					} `tfsdk:"pre_stop" json:"preStop,omitempty"`
				} `tfsdk:"lifecycle" json:"lifecycle,omitempty"`
				LivenessProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					Grpc             *struct {
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
						Service *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"grpc" json:"grpc,omitempty"`
					HttpGet *struct {
						Host        *string `tfsdk:"host" json:"host,omitempty"`
						HttpHeaders *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TcpSocket           *struct {
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *string `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Ports *[]struct {
					ContainerPort *int64  `tfsdk:"container_port" json:"containerPort,omitempty"`
					HostIP        *string `tfsdk:"host_ip" json:"hostIP,omitempty"`
					HostPort      *int64  `tfsdk:"host_port" json:"hostPort,omitempty"`
					Name          *string `tfsdk:"name" json:"name,omitempty"`
					Protocol      *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				ReadinessProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					Grpc             *struct {
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
						Service *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"grpc" json:"grpc,omitempty"`
					HttpGet *struct {
						Host        *string `tfsdk:"host" json:"host,omitempty"`
						HttpHeaders *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TcpSocket           *struct {
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *string `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				SecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					Capabilities             *struct {
						Add  *[]string `tfsdk:"add" json:"add,omitempty"`
						Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
					} `tfsdk:"capabilities" json:"capabilities,omitempty"`
					Privileged             *bool   `tfsdk:"privileged" json:"privileged,omitempty"`
					ProcMount              *string `tfsdk:"proc_mount" json:"procMount,omitempty"`
					ReadOnlyRootFilesystem *bool   `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
					RunAsGroup             *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
					RunAsNonRoot           *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser              *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					SeLinuxOptions         *struct {
						Level *string `tfsdk:"level" json:"level,omitempty"`
						Role  *string `tfsdk:"role" json:"role,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						User  *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
					SeccompProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
					WindowsOptions *struct {
						GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
						HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
						RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
				} `tfsdk:"security_context" json:"securityContext,omitempty"`
				StartupProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					Grpc             *struct {
						Port    *int64  `tfsdk:"port" json:"port,omitempty"`
						Service *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"grpc" json:"grpc,omitempty"`
					HttpGet *struct {
						Host        *string `tfsdk:"host" json:"host,omitempty"`
						HttpHeaders *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TcpSocket           *struct {
						Host *string `tfsdk:"host" json:"host,omitempty"`
						Port *string `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
				Stdin                    *bool   `tfsdk:"stdin" json:"stdin,omitempty"`
				StdinOnce                *bool   `tfsdk:"stdin_once" json:"stdinOnce,omitempty"`
				TerminationMessagePath   *string `tfsdk:"termination_message_path" json:"terminationMessagePath,omitempty"`
				TerminationMessagePolicy *string `tfsdk:"termination_message_policy" json:"terminationMessagePolicy,omitempty"`
				Tty                      *bool   `tfsdk:"tty" json:"tty,omitempty"`
				VolumeDevices            *[]struct {
					DevicePath *string `tfsdk:"device_path" json:"devicePath,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"volume_devices" json:"volumeDevices,omitempty"`
				VolumeMounts *[]struct {
					MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
					Name             *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
				WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
			} `tfsdk:"extra_init_containers" json:"extraInitContainers,omitempty"`
			ExtraVolumeClaimTemplates *[]struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					DataSource  *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"data_source" json:"dataSource,omitempty"`
					DataSourceRef *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
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
					VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"extra_volume_claim_templates" json:"extraVolumeClaimTemplates,omitempty"`
			ExtraVolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"extra_volume_mounts" json:"extraVolumeMounts,omitempty"`
			ExtraVolumes *[]struct {
				AwsElasticBlockStore *struct {
					FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					VolumeID  *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
				} `tfsdk:"aws_elastic_block_store" json:"awsElasticBlockStore,omitempty"`
				AzureDisk *struct {
					CachingMode *string `tfsdk:"caching_mode" json:"cachingMode,omitempty"`
					DiskName    *string `tfsdk:"disk_name" json:"diskName,omitempty"`
					DiskURI     *string `tfsdk:"disk_uri" json:"diskURI,omitempty"`
					FsType      *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
					ReadOnly    *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				} `tfsdk:"azure_disk" json:"azureDisk,omitempty"`
				AzureFile *struct {
					ReadOnly   *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					ShareName  *string `tfsdk:"share_name" json:"shareName,omitempty"`
				} `tfsdk:"azure_file" json:"azureFile,omitempty"`
				Cephfs *struct {
					Monitors   *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
					Path       *string   `tfsdk:"path" json:"path,omitempty"`
					ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretFile *string   `tfsdk:"secret_file" json:"secretFile,omitempty"`
					SecretRef  *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					User *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"cephfs" json:"cephfs,omitempty"`
				Cinder *struct {
					FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
				} `tfsdk:"cinder" json:"cinder,omitempty"`
				ConfigMap *struct {
					DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
					Items       *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Csi *struct {
					Driver               *string `tfsdk:"driver" json:"driver,omitempty"`
					FsType               *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					NodePublishSecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
					ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
					VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
				} `tfsdk:"csi" json:"csi,omitempty"`
				DownwardAPI *struct {
					DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
					Items       *[]struct {
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
						Path             *string `tfsdk:"path" json:"path,omitempty"`
						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
							Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
				EmptyDir *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				Ephemeral *struct {
					VolumeClaimTemplate *struct {
						Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
						Spec     *struct {
							AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
							DataSource  *struct {
								ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
								Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"data_source" json:"dataSource,omitempty"`
							DataSourceRef *struct {
								ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
								Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
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
							VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
							VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
				} `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
				Fc *struct {
					FsType     *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
					Lun        *int64    `tfsdk:"lun" json:"lun,omitempty"`
					ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
					TargetWWNs *[]string `tfsdk:"target_ww_ns" json:"targetWWNs,omitempty"`
					Wwids      *[]string `tfsdk:"wwids" json:"wwids,omitempty"`
				} `tfsdk:"fc" json:"fc,omitempty"`
				FlexVolume *struct {
					Driver    *string            `tfsdk:"driver" json:"driver,omitempty"`
					FsType    *string            `tfsdk:"fs_type" json:"fsType,omitempty"`
					Options   *map[string]string `tfsdk:"options" json:"options,omitempty"`
					ReadOnly  *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"flex_volume" json:"flexVolume,omitempty"`
				Flocker *struct {
					DatasetName *string `tfsdk:"dataset_name" json:"datasetName,omitempty"`
					DatasetUUID *string `tfsdk:"dataset_uuid" json:"datasetUUID,omitempty"`
				} `tfsdk:"flocker" json:"flocker,omitempty"`
				GcePersistentDisk *struct {
					FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
					PdName    *string `tfsdk:"pd_name" json:"pdName,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				} `tfsdk:"gce_persistent_disk" json:"gcePersistentDisk,omitempty"`
				GitRepo *struct {
					Directory  *string `tfsdk:"directory" json:"directory,omitempty"`
					Repository *string `tfsdk:"repository" json:"repository,omitempty"`
					Revision   *string `tfsdk:"revision" json:"revision,omitempty"`
				} `tfsdk:"git_repo" json:"gitRepo,omitempty"`
				Glusterfs *struct {
					Endpoints *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
					Path      *string `tfsdk:"path" json:"path,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				} `tfsdk:"glusterfs" json:"glusterfs,omitempty"`
				HostPath *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"host_path" json:"hostPath,omitempty"`
				Iscsi *struct {
					ChapAuthDiscovery *bool     `tfsdk:"chap_auth_discovery" json:"chapAuthDiscovery,omitempty"`
					ChapAuthSession   *bool     `tfsdk:"chap_auth_session" json:"chapAuthSession,omitempty"`
					FsType            *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
					InitiatorName     *string   `tfsdk:"initiator_name" json:"initiatorName,omitempty"`
					Iqn               *string   `tfsdk:"iqn" json:"iqn,omitempty"`
					IscsiInterface    *string   `tfsdk:"iscsi_interface" json:"iscsiInterface,omitempty"`
					Lun               *int64    `tfsdk:"lun" json:"lun,omitempty"`
					Portals           *[]string `tfsdk:"portals" json:"portals,omitempty"`
					ReadOnly          *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef         *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					TargetPortal *string `tfsdk:"target_portal" json:"targetPortal,omitempty"`
				} `tfsdk:"iscsi" json:"iscsi,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Nfs  *struct {
					Path     *string `tfsdk:"path" json:"path,omitempty"`
					ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					Server   *string `tfsdk:"server" json:"server,omitempty"`
				} `tfsdk:"nfs" json:"nfs,omitempty"`
				PersistentVolumeClaim *struct {
					ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
				PhotonPersistentDisk *struct {
					FsType *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					PdID   *string `tfsdk:"pd_id" json:"pdID,omitempty"`
				} `tfsdk:"photon_persistent_disk" json:"photonPersistentDisk,omitempty"`
				PortworxVolume *struct {
					FsType   *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
				} `tfsdk:"portworx_volume" json:"portworxVolume,omitempty"`
				Projected *struct {
					DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
					Sources     *[]struct {
						ConfigMap *struct {
							Items *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						DownwardAPI *struct {
							Items *[]struct {
								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
									FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
								Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path             *string `tfsdk:"path" json:"path,omitempty"`
								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
									Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
									Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
						Secret *struct {
							Items *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						ServiceAccountToken *struct {
							Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
							ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
							Path              *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
					} `tfsdk:"sources" json:"sources,omitempty"`
				} `tfsdk:"projected" json:"projected,omitempty"`
				Quobyte *struct {
					Group    *string `tfsdk:"group" json:"group,omitempty"`
					ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					Registry *string `tfsdk:"registry" json:"registry,omitempty"`
					Tenant   *string `tfsdk:"tenant" json:"tenant,omitempty"`
					User     *string `tfsdk:"user" json:"user,omitempty"`
					Volume   *string `tfsdk:"volume" json:"volume,omitempty"`
				} `tfsdk:"quobyte" json:"quobyte,omitempty"`
				Rbd *struct {
					FsType    *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
					Image     *string   `tfsdk:"image" json:"image,omitempty"`
					Keyring   *string   `tfsdk:"keyring" json:"keyring,omitempty"`
					Monitors  *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
					Pool      *string   `tfsdk:"pool" json:"pool,omitempty"`
					ReadOnly  *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					User *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"rbd" json:"rbd,omitempty"`
				ScaleIO *struct {
					FsType           *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					Gateway          *string `tfsdk:"gateway" json:"gateway,omitempty"`
					ProtectionDomain *string `tfsdk:"protection_domain" json:"protectionDomain,omitempty"`
					ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef        *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					SslEnabled  *bool   `tfsdk:"ssl_enabled" json:"sslEnabled,omitempty"`
					StorageMode *string `tfsdk:"storage_mode" json:"storageMode,omitempty"`
					StoragePool *string `tfsdk:"storage_pool" json:"storagePool,omitempty"`
					System      *string `tfsdk:"system" json:"system,omitempty"`
					VolumeName  *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"scale_io" json:"scaleIO,omitempty"`
				Secret *struct {
					DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
					Items       *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
					Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
					SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Storageos *struct {
					FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					VolumeName      *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
					VolumeNamespace *string `tfsdk:"volume_namespace" json:"volumeNamespace,omitempty"`
				} `tfsdk:"storageos" json:"storageos,omitempty"`
				VsphereVolume *struct {
					FsType            *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					StoragePolicyID   *string `tfsdk:"storage_policy_id" json:"storagePolicyID,omitempty"`
					StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
					VolumePath        *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
				} `tfsdk:"vsphere_volume" json:"vsphereVolume,omitempty"`
			} `tfsdk:"extra_volumes" json:"extraVolumes,omitempty"`
			Image           *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ImagePullSecret *string `tfsdk:"image_pull_secret" json:"imagePullSecret,omitempty"`
			MetricsExporter *struct {
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Resources       *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"metrics_exporter" json:"metricsExporter,omitempty"`
			Resources *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"varnish" json:"varnish,omitempty"`
		Vcl *struct {
			ConfigMapName      *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
			EntrypointFileName *string `tfsdk:"entrypoint_file_name" json:"entrypointFileName,omitempty"`
		} `tfsdk:"vcl" json:"vcl,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CachingIbmComVarnishClusterV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_caching_ibm_com_varnish_cluster_v1alpha1"
}

func (r *CachingIbmComVarnishClusterV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VarnishCluster is the Schema for the varnishclusters API",
		MarkdownDescription: "VarnishCluster is the Schema for the varnishclusters API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "VarnishClusterSpec defines the desired state of VarnishCluster",
				MarkdownDescription: "VarnishClusterSpec defines the desired state of VarnishCluster",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Affinity is a group of affinity scheduling rules.",
						MarkdownDescription: "Affinity is a group of affinity scheduling rules.",
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"weight": schema.Int64Attribute{
													Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
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
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
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
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
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
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"backend": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"namespaces": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"only_ready": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"selector": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"zone_balancing": schema.SingleNestedAttribute{
								Description:         "Defines the type and parameters for backend traffic distribution in multi-zone clusters",
								MarkdownDescription: "Defines the type and parameters for backend traffic distribution in multi-zone clusters",
								Attributes: map[string]schema.Attribute{
									"thresholds": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"local": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"remote": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"threshold": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"log_format": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"log_level": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"monitoring": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"grafana_dashboard": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"datasource_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"title": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"prometheus_service_monitor": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"scrape_interval": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"node_selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pod_disruption_budget": schema.SingleNestedAttribute{
						Description:         "PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.",
						MarkdownDescription: "PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.",
						Attributes: map[string]schema.Attribute{
							"max_unavailable": schema.StringAttribute{
								Description:         "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
								MarkdownDescription: "An eviction is allowed if at most 'maxUnavailable' pods selected by 'selector' are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_available": schema.StringAttribute{
								Description:         "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying '100%'.",
								MarkdownDescription: "An eviction is allowed if at least 'minAvailable' pods selected by 'selector' will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying '100%'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"selector": schema.SingleNestedAttribute{
								Description:         "Label query over pods whose evictions are managed by the disruption budget. A null selector will match no pods, while an empty ({}) selector will select all pods within the namespace.",
								MarkdownDescription: "Label query over pods whose evictions are managed by the disruption budget. A null selector will match no pods, while an empty ({}) selector will select all pods within the namespace.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"priority_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replicas": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"controller_metrics_node_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metrics_node_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metrics_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Service Type string describes ingress methods for a service",
								MarkdownDescription: "Service Type string describes ingress methods for a service",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"update_strategy": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"delayed_rolling_update": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"delay_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"rolling_update": schema.SingleNestedAttribute{
								Description:         "RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.",
								MarkdownDescription: "RollingUpdateStatefulSetStrategy is used to communicate parameter for RollingUpdateStatefulSetStrategyType.",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"partition": schema.Int64Attribute{
										Description:         "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
										MarkdownDescription: "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"varnish": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"adm_auth": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"args": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"controller": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes a policy for if/when to pull a container image",
										MarkdownDescription: "PullPolicy describes a policy for if/when to pull a container image",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "ResourceRequirements describes the compute resource requirements.",
										MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"env_from": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config_map_ref": schema.SingleNestedAttribute{
											Description:         "The ConfigMap to select from",
											MarkdownDescription: "The ConfigMap to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap must be defined",
													MarkdownDescription: "Specify whether the ConfigMap must be defined",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"prefix": schema.StringAttribute{
											Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "The Secret to select from",
											MarkdownDescription: "The Secret to select from",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret must be defined",
													MarkdownDescription: "Specify whether the Secret must be defined",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"extra_init_containers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"args": schema.ListAttribute{
											Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"command": schema.ListAttribute{
											Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"env": schema.ListNestedAttribute{
											Description:         "List of environment variables to set in the container. Cannot be updated.",
											MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
														MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value": schema.StringAttribute{
														Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"value_from": schema.SingleNestedAttribute{
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
														Attributes: map[string]schema.Attribute{
															"config_map_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"field_ref": schema.SingleNestedAttribute{
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																Attributes: map[string]schema.Attribute{
																	"api_version": schema.StringAttribute{
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"field_path": schema.StringAttribute{
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"resource_field_ref": schema.SingleNestedAttribute{
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"secret_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"env_from": schema.ListNestedAttribute{
											Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
											MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"prefix": schema.StringAttribute{
														Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "The Secret to select from",
														MarkdownDescription: "The Secret to select from",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"image": schema.StringAttribute{
											Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
											MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image_pull_policy": schema.StringAttribute{
											Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
											MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"lifecycle": schema.SingleNestedAttribute{
											Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
											MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
											Attributes: map[string]schema.Attribute{
												"post_start": schema.SingleNestedAttribute{
													Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"http_get": schema.SingleNestedAttribute{
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"path": schema.StringAttribute{
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
															MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"pre_stop": schema.SingleNestedAttribute{
													Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"http_get": schema.SingleNestedAttribute{
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"path": schema.StringAttribute{
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
															MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"liveness_probe": schema.SingleNestedAttribute{
											Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"failure_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"http_get": schema.SingleNestedAttribute{
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"path": schema.StringAttribute{
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"initial_delay_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
											MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
											MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_port": schema.Int64Attribute{
														Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_ip": schema.StringAttribute{
														Description:         "What host IP to bind the external port to.",
														MarkdownDescription: "What host IP to bind the external port to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_port": schema.Int64Attribute{
														Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
														MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
														MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"readiness_probe": schema.SingleNestedAttribute{
											Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"failure_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"http_get": schema.SingleNestedAttribute{
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"path": schema.StringAttribute{
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"initial_delay_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											Attributes: map[string]schema.Attribute{
												"limits": schema.MapAttribute{
													Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"security_context": schema.SingleNestedAttribute{
											Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
											MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
											Attributes: map[string]schema.Attribute{
												"allow_privilege_escalation": schema.BoolAttribute{
													Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"capabilities": schema.SingleNestedAttribute{
													Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
													Attributes: map[string]schema.Attribute{
														"add": schema.ListAttribute{
															Description:         "Added capabilities",
															MarkdownDescription: "Added capabilities",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"drop": schema.ListAttribute{
															Description:         "Removed capabilities",
															MarkdownDescription: "Removed capabilities",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"privileged": schema.BoolAttribute{
													Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"proc_mount": schema.StringAttribute{
													Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only_root_filesystem": schema.BoolAttribute{
													Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"run_as_group": schema.Int64Attribute{
													Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"run_as_non_root": schema.BoolAttribute{
													Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
													MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"run_as_user": schema.Int64Attribute{
													Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"se_linux_options": schema.SingleNestedAttribute{
													Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
													Attributes: map[string]schema.Attribute{
														"level": schema.StringAttribute{
															Description:         "Level is SELinux level label that applies to the container.",
															MarkdownDescription: "Level is SELinux level label that applies to the container.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"role": schema.StringAttribute{
															Description:         "Role is a SELinux role label that applies to the container.",
															MarkdownDescription: "Role is a SELinux role label that applies to the container.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"type": schema.StringAttribute{
															Description:         "Type is a SELinux type label that applies to the container.",
															MarkdownDescription: "Type is a SELinux type label that applies to the container.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"user": schema.StringAttribute{
															Description:         "User is a SELinux user label that applies to the container.",
															MarkdownDescription: "User is a SELinux user label that applies to the container.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"seccomp_profile": schema.SingleNestedAttribute{
													Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
													Attributes: map[string]schema.Attribute{
														"localhost_profile": schema.StringAttribute{
															Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
															MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"type": schema.StringAttribute{
															Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
															MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"windows_options": schema.SingleNestedAttribute{
													Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
													MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
													Attributes: map[string]schema.Attribute{
														"gmsa_credential_spec": schema.StringAttribute{
															Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
															MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"gmsa_credential_spec_name": schema.StringAttribute{
															Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
															MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"host_process": schema.BoolAttribute{
															Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
															MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"run_as_user_name": schema.StringAttribute{
															Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
															MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"startup_probe": schema.SingleNestedAttribute{
											Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"failure_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"http_get": schema.SingleNestedAttribute{
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"path": schema.StringAttribute{
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"initial_delay_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"stdin": schema.BoolAttribute{
											Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
											MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"stdin_once": schema.BoolAttribute{
											Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
											MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"termination_message_path": schema.StringAttribute{
											Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
											MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"termination_message_policy": schema.StringAttribute{
											Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
											MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tty": schema.BoolAttribute{
											Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
											MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_devices": schema.ListNestedAttribute{
											Description:         "volumeDevices is the list of block devices to be used by the container.",
											MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"device_path": schema.StringAttribute{
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"volume_mounts": schema.ListNestedAttribute{
											Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
											MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"mount_propagation": schema.StringAttribute{
														Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "This must match the Name of a Volume.",
														MarkdownDescription: "This must match the Name of a Volume.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"read_only": schema.BoolAttribute{
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"sub_path": schema.StringAttribute{
														Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"sub_path_expr": schema.StringAttribute{
														Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"working_dir": schema.StringAttribute{
											Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
											MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"extra_volume_claim_templates": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"labels": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes",
											MarkdownDescription: "PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes",
											Attributes: map[string]schema.Attribute{
												"access_modes": schema.ListAttribute{
													Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"data_source": schema.SingleNestedAttribute{
													Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
													MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
													Attributes: map[string]schema.Attribute{
														"api_group": schema.StringAttribute{
															Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
															MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"kind": schema.StringAttribute{
															Description:         "Kind is the type of resource being referenced",
															MarkdownDescription: "Kind is the type of resource being referenced",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of resource being referenced",
															MarkdownDescription: "Name is the name of resource being referenced",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"data_source_ref": schema.SingleNestedAttribute{
													Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
													MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
													Attributes: map[string]schema.Attribute{
														"api_group": schema.StringAttribute{
															Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
															MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"kind": schema.StringAttribute{
															Description:         "Kind is the type of resource being referenced",
															MarkdownDescription: "Kind is the type of resource being referenced",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of resource being referenced",
															MarkdownDescription: "Name is the name of resource being referenced",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"resources": schema.SingleNestedAttribute{
													Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
													MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
													Attributes: map[string]schema.Attribute{
														"limits": schema.MapAttribute{
															Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"requests": schema.MapAttribute{
															Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "selector is a label query over volumes to consider for binding.",
													MarkdownDescription: "selector is a label query over volumes to consider for binding.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"storage_class_name": schema.StringAttribute{
													Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
													MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_mode": schema.StringAttribute{
													Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
													MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_name": schema.StringAttribute{
													Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
													MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"extra_volume_mounts": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"extra_volumes": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"aws_elastic_block_store": schema.SingleNestedAttribute{
											Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"partition": schema.Int64Attribute{
													Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
													MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_id": schema.StringAttribute{
													Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"azure_disk": schema.SingleNestedAttribute{
											Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
											MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
											Attributes: map[string]schema.Attribute{
												"caching_mode": schema.StringAttribute{
													Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
													MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"disk_name": schema.StringAttribute{
													Description:         "diskName is the Name of the data disk in the blob storage",
													MarkdownDescription: "diskName is the Name of the data disk in the blob storage",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"disk_uri": schema.StringAttribute{
													Description:         "diskURI is the URI of data disk in the blob storage",
													MarkdownDescription: "diskURI is the URI of data disk in the blob storage",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"fs_type": schema.StringAttribute{
													Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"kind": schema.StringAttribute{
													Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
													MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"azure_file": schema.SingleNestedAttribute{
											Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
											MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
											Attributes: map[string]schema.Attribute{
												"read_only": schema.BoolAttribute{
													Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_name": schema.StringAttribute{
													Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
													MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"share_name": schema.StringAttribute{
													Description:         "shareName is the azure share Name",
													MarkdownDescription: "shareName is the azure share Name",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"cephfs": schema.SingleNestedAttribute{
											Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
											MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
											Attributes: map[string]schema.Attribute{
												"monitors": schema.ListAttribute{
													Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"path": schema.StringAttribute{
													Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
													MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_file": schema.StringAttribute{
													Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"user": schema.StringAttribute{
													Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"cinder": schema.SingleNestedAttribute{
											Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
													MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"volume_id": schema.StringAttribute{
													Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"config_map": schema.SingleNestedAttribute{
											Description:         "configMap represents a configMap that should populate this volume",
											MarkdownDescription: "configMap represents a configMap that should populate this volume",
											Attributes: map[string]schema.Attribute{
												"default_mode": schema.Int64Attribute{
													Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"items": schema.ListNestedAttribute{
													Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
													MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"mode": schema.Int64Attribute{
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"optional": schema.BoolAttribute{
													Description:         "optional specify whether the ConfigMap or its keys must be defined",
													MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"csi": schema.SingleNestedAttribute{
											Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
											MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
											Attributes: map[string]schema.Attribute{
												"driver": schema.StringAttribute{
													Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
													MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"fs_type": schema.StringAttribute{
													Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
													MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_publish_secret_ref": schema.SingleNestedAttribute{
													Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
													MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
													MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_attributes": schema.MapAttribute{
													Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
													MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"downward_api": schema.SingleNestedAttribute{
											Description:         "downwardAPI represents downward API about the pod that should populate this volume",
											MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",
											Attributes: map[string]schema.Attribute{
												"default_mode": schema.Int64Attribute{
													Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"items": schema.ListNestedAttribute{
													Description:         "Items is a list of downward API volume file",
													MarkdownDescription: "Items is a list of downward API volume file",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"field_ref": schema.SingleNestedAttribute{
																Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																Attributes: map[string]schema.Attribute{
																	"api_version": schema.StringAttribute{
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"field_path": schema.StringAttribute{
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"mode": schema.Int64Attribute{
																Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"resource_field_ref": schema.SingleNestedAttribute{
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"empty_dir": schema.SingleNestedAttribute{
											Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											Attributes: map[string]schema.Attribute{
												"medium": schema.StringAttribute{
													Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"size_limit": schema.StringAttribute{
													Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
													MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"ephemeral": schema.SingleNestedAttribute{
											Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
											MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
											Attributes: map[string]schema.Attribute{
												"volume_claim_template": schema.SingleNestedAttribute{
													Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
													MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
													Attributes: map[string]schema.Attribute{
														"metadata": schema.MapAttribute{
															Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
															MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"spec": schema.SingleNestedAttribute{
															Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
															MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
															Attributes: map[string]schema.Attribute{
																"access_modes": schema.ListAttribute{
																	Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																	MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"data_source": schema.SingleNestedAttribute{
																	Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																	MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																	Attributes: map[string]schema.Attribute{
																		"api_group": schema.StringAttribute{
																			Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																			MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "Kind is the type of resource being referenced",
																			MarkdownDescription: "Kind is the type of resource being referenced",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name is the name of resource being referenced",
																			MarkdownDescription: "Name is the name of resource being referenced",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"data_source_ref": schema.SingleNestedAttribute{
																	Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																	MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																	Attributes: map[string]schema.Attribute{
																		"api_group": schema.StringAttribute{
																			Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																			MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "Kind is the type of resource being referenced",
																			MarkdownDescription: "Kind is the type of resource being referenced",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name is the name of resource being referenced",
																			MarkdownDescription: "Name is the name of resource being referenced",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																	MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																	Attributes: map[string]schema.Attribute{
																		"limits": schema.MapAttribute{
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"selector": schema.SingleNestedAttribute{
																	Description:         "selector is a label query over volumes to consider for binding.",
																	MarkdownDescription: "selector is a label query over volumes to consider for binding.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"storage_class_name": schema.StringAttribute{
																	Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																	MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"volume_mode": schema.StringAttribute{
																	Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																	MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"volume_name": schema.StringAttribute{
																	Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
																	MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"fc": schema.SingleNestedAttribute{
											Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
											MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"lun": schema.Int64Attribute{
													Description:         "lun is Optional: FC target lun number",
													MarkdownDescription: "lun is Optional: FC target lun number",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"target_ww_ns": schema.ListAttribute{
													Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
													MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"wwids": schema.ListAttribute{
													Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
													MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"flex_volume": schema.SingleNestedAttribute{
											Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
											MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
											Attributes: map[string]schema.Attribute{
												"driver": schema.StringAttribute{
													Description:         "driver is the name of the driver to use for this volume.",
													MarkdownDescription: "driver is the name of the driver to use for this volume.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
													MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"options": schema.MapAttribute{
													Description:         "options is Optional: this field holds extra command options if any.",
													MarkdownDescription: "options is Optional: this field holds extra command options if any.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
													MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"flocker": schema.SingleNestedAttribute{
											Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
											MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
											Attributes: map[string]schema.Attribute{
												"dataset_name": schema.StringAttribute{
													Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
													MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"dataset_uuid": schema.StringAttribute{
													Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
													MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"gce_persistent_disk": schema.SingleNestedAttribute{
											Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"partition": schema.Int64Attribute{
													Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"pd_name": schema.StringAttribute{
													Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"git_repo": schema.SingleNestedAttribute{
											Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
											MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
											Attributes: map[string]schema.Attribute{
												"directory": schema.StringAttribute{
													Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
													MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"repository": schema.StringAttribute{
													Description:         "repository is the URL",
													MarkdownDescription: "repository is the URL",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"revision": schema.StringAttribute{
													Description:         "revision is the commit hash for the specified revision.",
													MarkdownDescription: "revision is the commit hash for the specified revision.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"glusterfs": schema.SingleNestedAttribute{
											Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
											MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
											Attributes: map[string]schema.Attribute{
												"endpoints": schema.StringAttribute{
													Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"path": schema.StringAttribute{
													Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"host_path": schema.SingleNestedAttribute{
											Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
											MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
													MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"type": schema.StringAttribute{
													Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
													MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"iscsi": schema.SingleNestedAttribute{
											Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
											MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
											Attributes: map[string]schema.Attribute{
												"chap_auth_discovery": schema.BoolAttribute{
													Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
													MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"chap_auth_session": schema.BoolAttribute{
													Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
													MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"initiator_name": schema.StringAttribute{
													Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
													MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"iqn": schema.StringAttribute{
													Description:         "iqn is the target iSCSI Qualified Name.",
													MarkdownDescription: "iqn is the target iSCSI Qualified Name.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"iscsi_interface": schema.StringAttribute{
													Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
													MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"lun": schema.Int64Attribute{
													Description:         "lun represents iSCSI Target Lun number.",
													MarkdownDescription: "lun represents iSCSI Target Lun number.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"portals": schema.ListAttribute{
													Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
													MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
													MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
													MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"target_portal": schema.StringAttribute{
													Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
													MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"nfs": schema.SingleNestedAttribute{
											Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"server": schema.StringAttribute{
													Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"persistent_volume_claim": schema.SingleNestedAttribute{
											Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											Attributes: map[string]schema.Attribute{
												"claim_name": schema.StringAttribute{
													Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
													MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
													MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"photon_persistent_disk": schema.SingleNestedAttribute{
											Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
											MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"pd_id": schema.StringAttribute{
													Description:         "pdID is the ID that identifies Photon Controller persistent disk",
													MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"portworx_volume": schema.SingleNestedAttribute{
											Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
											MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_id": schema.StringAttribute{
													Description:         "volumeID uniquely identifies a Portworx volume",
													MarkdownDescription: "volumeID uniquely identifies a Portworx volume",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"projected": schema.SingleNestedAttribute{
											Description:         "projected items for all in one resources secrets, configmaps, and downward API",
											MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",
											Attributes: map[string]schema.Attribute{
												"default_mode": schema.Int64Attribute{
													Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"sources": schema.ListNestedAttribute{
													Description:         "sources is the list of volume projections",
													MarkdownDescription: "sources is the list of volume projections",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config_map": schema.SingleNestedAttribute{
																Description:         "configMap information about the configMap data to project",
																MarkdownDescription: "configMap information about the configMap data to project",
																Attributes: map[string]schema.Attribute{
																	"items": schema.ListNestedAttribute{
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the key to project.",
																					MarkdownDescription: "key is the key to project.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"mode": schema.Int64Attribute{
																					Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"path": schema.StringAttribute{
																					Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "optional specify whether the ConfigMap or its keys must be defined",
																		MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"downward_api": schema.SingleNestedAttribute{
																Description:         "downwardAPI information about the downwardAPI data to project",
																MarkdownDescription: "downwardAPI information about the downwardAPI data to project",
																Attributes: map[string]schema.Attribute{
																	"items": schema.ListNestedAttribute{
																		Description:         "Items is a list of DownwardAPIVolume file",
																		MarkdownDescription: "Items is a list of DownwardAPIVolume file",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"field_ref": schema.SingleNestedAttribute{
																					Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																					MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																					Attributes: map[string]schema.Attribute{
																						"api_version": schema.StringAttribute{
																							Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																							MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"field_path": schema.StringAttribute{
																							Description:         "Path of the field to select in the specified API version.",
																							MarkdownDescription: "Path of the field to select in the specified API version.",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"mode": schema.Int64Attribute{
																					Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"path": schema.StringAttribute{
																					Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																					MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"resource_field_ref": schema.SingleNestedAttribute{
																					Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																					MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																					Attributes: map[string]schema.Attribute{
																						"container_name": schema.StringAttribute{
																							Description:         "Container name: required for volumes, optional for env vars",
																							MarkdownDescription: "Container name: required for volumes, optional for env vars",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"divisor": schema.StringAttribute{
																							Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																							MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"resource": schema.StringAttribute{
																							Description:         "Required: resource to select",
																							MarkdownDescription: "Required: resource to select",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"secret": schema.SingleNestedAttribute{
																Description:         "secret information about the secret data to project",
																MarkdownDescription: "secret information about the secret data to project",
																Attributes: map[string]schema.Attribute{
																	"items": schema.ListNestedAttribute{
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the key to project.",
																					MarkdownDescription: "key is the key to project.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"mode": schema.Int64Attribute{
																					Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"path": schema.StringAttribute{
																					Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "optional field specify whether the Secret or its key must be defined",
																		MarkdownDescription: "optional field specify whether the Secret or its key must be defined",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"service_account_token": schema.SingleNestedAttribute{
																Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
																MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",
																Attributes: map[string]schema.Attribute{
																	"audience": schema.StringAttribute{
																		Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																		MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"expiration_seconds": schema.Int64Attribute{
																		Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																		MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"path": schema.StringAttribute{
																		Description:         "path is the path relative to the mount point of the file to project the token into.",
																		MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"quobyte": schema.SingleNestedAttribute{
											Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
											MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
											Attributes: map[string]schema.Attribute{
												"group": schema.StringAttribute{
													Description:         "group to map volume access to Default is no group",
													MarkdownDescription: "group to map volume access to Default is no group",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
													MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"registry": schema.StringAttribute{
													Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
													MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tenant": schema.StringAttribute{
													Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
													MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"user": schema.StringAttribute{
													Description:         "user to map volume access to Defaults to serivceaccount user",
													MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume": schema.StringAttribute{
													Description:         "volume is a string that references an already created Quobyte volume by name.",
													MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"rbd": schema.SingleNestedAttribute{
											Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
											MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"image": schema.StringAttribute{
													Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"keyring": schema.StringAttribute{
													Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"monitors": schema.ListAttribute{
													Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"pool": schema.StringAttribute{
													Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"user": schema.StringAttribute{
													Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"scale_io": schema.SingleNestedAttribute{
											Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
											MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
													MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"gateway": schema.StringAttribute{
													Description:         "gateway is the host address of the ScaleIO API Gateway.",
													MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"protection_domain": schema.StringAttribute{
													Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
													MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
													MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"ssl_enabled": schema.BoolAttribute{
													Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
													MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"storage_mode": schema.StringAttribute{
													Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
													MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"storage_pool": schema.StringAttribute{
													Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
													MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"system": schema.StringAttribute{
													Description:         "system is the name of the storage system as configured in ScaleIO.",
													MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_name": schema.StringAttribute{
													Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
													MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"secret": schema.SingleNestedAttribute{
											Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
											MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
											Attributes: map[string]schema.Attribute{
												"default_mode": schema.Int64Attribute{
													Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"items": schema.ListNestedAttribute{
													Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
													MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"mode": schema.Int64Attribute{
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"optional": schema.BoolAttribute{
													Description:         "optional field specify whether the Secret or its keys must be defined",
													MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_name": schema.StringAttribute{
													Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
													MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"storageos": schema.SingleNestedAttribute{
											Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
											MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
													MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"volume_name": schema.StringAttribute{
													Description:         "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
													MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_namespace": schema.StringAttribute{
													Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
													MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"vsphere_volume": schema.SingleNestedAttribute{
											Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
											MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"storage_policy_id": schema.StringAttribute{
													Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
													MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"storage_policy_name": schema.StringAttribute{
													Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
													MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"volume_path": schema.StringAttribute{
													Description:         "volumePath is the path that identifies vSphere volume vmdk",
													MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "PullPolicy describes a policy for if/when to pull a container image",
								MarkdownDescription: "PullPolicy describes a policy for if/when to pull a container image",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_pull_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metrics_exporter": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "PullPolicy describes a policy for if/when to pull a container image",
										MarkdownDescription: "PullPolicy describes a policy for if/when to pull a container image",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "ResourceRequirements describes the compute resource requirements.",
										MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "ResourceRequirements describes the compute resource requirements.",
								MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"vcl": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"config_map_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"entrypoint_file_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CachingIbmComVarnishClusterV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *CachingIbmComVarnishClusterV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_caching_ibm_com_varnish_cluster_v1alpha1")

	var data CachingIbmComVarnishClusterV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "caching.ibm.com", Version: "v1alpha1", Resource: "varnishclusters"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CachingIbmComVarnishClusterV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("caching.ibm.com/v1alpha1")
	data.Kind = pointer.String("VarnishCluster")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
