/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package grafana_integreatly_org_v1beta1

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
	_ datasource.DataSource = &GrafanaIntegreatlyOrgGrafanaV1Beta1Manifest{}
)

func NewGrafanaIntegreatlyOrgGrafanaV1Beta1Manifest() datasource.DataSource {
	return &GrafanaIntegreatlyOrgGrafanaV1Beta1Manifest{}
}

type GrafanaIntegreatlyOrgGrafanaV1Beta1Manifest struct{}

type GrafanaIntegreatlyOrgGrafanaV1Beta1ManifestData struct {
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
		Client *struct {
			Headers       *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
			PreferIngress *bool              `tfsdk:"prefer_ingress" json:"preferIngress,omitempty"`
			Timeout       *int64             `tfsdk:"timeout" json:"timeout,omitempty"`
			Tls           *struct {
				CertSecretRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"client" json:"client,omitempty"`
		Config     *map[string]string `tfsdk:"config" json:"config,omitempty"`
		Deployment *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				MinReadySeconds         *int64 `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
				Paused                  *bool  `tfsdk:"paused" json:"paused,omitempty"`
				ProgressDeadlineSeconds *int64 `tfsdk:"progress_deadline_seconds" json:"progressDeadlineSeconds,omitempty"`
				Replicas                *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				RevisionHistoryLimit    *int64 `tfsdk:"revision_history_limit" json:"revisionHistoryLimit,omitempty"`
				Selector                *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Strategy *struct {
					RollingUpdate *struct {
						MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
						MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"strategy" json:"strategy,omitempty"`
				Template *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec *struct {
						ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
						Affinity              *struct {
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
						AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" json:"automountServiceAccountToken,omitempty"`
						Containers                   *[]struct {
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
									Sleep *struct {
										Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
									} `tfsdk:"sleep" json:"sleep,omitempty"`
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
									Sleep *struct {
										Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
									} `tfsdk:"sleep" json:"sleep,omitempty"`
									TcpSocket *struct {
										Host *string `tfsdk:"host" json:"host,omitempty"`
										Port *string `tfsdk:"port" json:"port,omitempty"`
									} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
								} `tfsdk:"pre_stop" json:"preStop,omitempty"`
								StopSignal *string `tfsdk:"stop_signal" json:"stopSignal,omitempty"`
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
							ResizePolicy *[]struct {
								ResourceName  *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
								RestartPolicy *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							} `tfsdk:"resize_policy" json:"resizePolicy,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name    *string `tfsdk:"name" json:"name,omitempty"`
									Request *string `tfsdk:"request" json:"request,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							SecurityContext *struct {
								AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
								AppArmorProfile          *struct {
									LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
									Type             *string `tfsdk:"type" json:"type,omitempty"`
								} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
								Capabilities *struct {
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
								MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
								MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								Name              *string `tfsdk:"name" json:"name,omitempty"`
								ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
								SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
							WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						DnsConfig *struct {
							Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
							Options     *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"options" json:"options,omitempty"`
							Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
						} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
						DnsPolicy           *string `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
						EnableServiceLinks  *bool   `tfsdk:"enable_service_links" json:"enableServiceLinks,omitempty"`
						EphemeralContainers *[]struct {
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
									Sleep *struct {
										Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
									} `tfsdk:"sleep" json:"sleep,omitempty"`
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
									Sleep *struct {
										Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
									} `tfsdk:"sleep" json:"sleep,omitempty"`
									TcpSocket *struct {
										Host *string `tfsdk:"host" json:"host,omitempty"`
										Port *string `tfsdk:"port" json:"port,omitempty"`
									} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
								} `tfsdk:"pre_stop" json:"preStop,omitempty"`
								StopSignal *string `tfsdk:"stop_signal" json:"stopSignal,omitempty"`
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
							ResizePolicy *[]struct {
								ResourceName  *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
								RestartPolicy *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							} `tfsdk:"resize_policy" json:"resizePolicy,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name    *string `tfsdk:"name" json:"name,omitempty"`
									Request *string `tfsdk:"request" json:"request,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							SecurityContext *struct {
								AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
								AppArmorProfile          *struct {
									LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
									Type             *string `tfsdk:"type" json:"type,omitempty"`
								} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
								Capabilities *struct {
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
							TargetContainerName      *string `tfsdk:"target_container_name" json:"targetContainerName,omitempty"`
							TerminationMessagePath   *string `tfsdk:"termination_message_path" json:"terminationMessagePath,omitempty"`
							TerminationMessagePolicy *string `tfsdk:"termination_message_policy" json:"terminationMessagePolicy,omitempty"`
							Tty                      *bool   `tfsdk:"tty" json:"tty,omitempty"`
							VolumeDevices            *[]struct {
								DevicePath *string `tfsdk:"device_path" json:"devicePath,omitempty"`
								Name       *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"volume_devices" json:"volumeDevices,omitempty"`
							VolumeMounts *[]struct {
								MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
								MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								Name              *string `tfsdk:"name" json:"name,omitempty"`
								ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
								SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
							WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
						} `tfsdk:"ephemeral_containers" json:"ephemeralContainers,omitempty"`
						HostAliases *[]struct {
							Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
							Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
						} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
						HostIPC          *bool   `tfsdk:"host_ipc" json:"hostIPC,omitempty"`
						HostNetwork      *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
						HostPID          *bool   `tfsdk:"host_pid" json:"hostPID,omitempty"`
						HostUsers        *bool   `tfsdk:"host_users" json:"hostUsers,omitempty"`
						Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
						ImagePullSecrets *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
						InitContainers *[]struct {
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
									Sleep *struct {
										Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
									} `tfsdk:"sleep" json:"sleep,omitempty"`
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
									Sleep *struct {
										Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
									} `tfsdk:"sleep" json:"sleep,omitempty"`
									TcpSocket *struct {
										Host *string `tfsdk:"host" json:"host,omitempty"`
										Port *string `tfsdk:"port" json:"port,omitempty"`
									} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
								} `tfsdk:"pre_stop" json:"preStop,omitempty"`
								StopSignal *string `tfsdk:"stop_signal" json:"stopSignal,omitempty"`
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
							ResizePolicy *[]struct {
								ResourceName  *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
								RestartPolicy *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							} `tfsdk:"resize_policy" json:"resizePolicy,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name    *string `tfsdk:"name" json:"name,omitempty"`
									Request *string `tfsdk:"request" json:"request,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							SecurityContext *struct {
								AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
								AppArmorProfile          *struct {
									LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
									Type             *string `tfsdk:"type" json:"type,omitempty"`
								} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
								Capabilities *struct {
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
								MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
								MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								Name              *string `tfsdk:"name" json:"name,omitempty"`
								ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
								SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
							WorkingDir *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
						} `tfsdk:"init_containers" json:"initContainers,omitempty"`
						NodeName     *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
						NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						Os           *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"os" json:"os,omitempty"`
						Overhead          *map[string]string `tfsdk:"overhead" json:"overhead,omitempty"`
						PreemptionPolicy  *string            `tfsdk:"preemption_policy" json:"preemptionPolicy,omitempty"`
						Priority          *int64             `tfsdk:"priority" json:"priority,omitempty"`
						PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
						ReadinessGates    *[]struct {
							ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
						} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
						RestartPolicy    *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
						RuntimeClassName *string `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
						SchedulerName    *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
						SecurityContext  *struct {
							AppArmorProfile *struct {
								LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
								Type             *string `tfsdk:"type" json:"type,omitempty"`
							} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
							FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
							FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
							RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
							RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
							RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
							SeLinuxChangePolicy *string `tfsdk:"se_linux_change_policy" json:"seLinuxChangePolicy,omitempty"`
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
							SupplementalGroups       *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
							SupplementalGroupsPolicy *string   `tfsdk:"supplemental_groups_policy" json:"supplementalGroupsPolicy,omitempty"`
							Sysctls                  *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"sysctls" json:"sysctls,omitempty"`
							WindowsOptions *struct {
								GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
								GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
								HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
								RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
							} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
						} `tfsdk:"security_context" json:"securityContext,omitempty"`
						ServiceAccount                *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
						ServiceAccountName            *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
						SetHostnameAsFQDN             *bool   `tfsdk:"set_hostname_as_fqdn" json:"setHostnameAsFQDN,omitempty"`
						ShareProcessNamespace         *bool   `tfsdk:"share_process_namespace" json:"shareProcessNamespace,omitempty"`
						Subdomain                     *string `tfsdk:"subdomain" json:"subdomain,omitempty"`
						TerminationGracePeriodSeconds *int64  `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
						Tolerations                   *[]struct {
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
						Volumes *[]struct {
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
											ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
											Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
											Name      *string `tfsdk:"name" json:"name,omitempty"`
											Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
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
										StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
										VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
										VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
										VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
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
							Image *struct {
								PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
								Reference  *string `tfsdk:"reference" json:"reference,omitempty"`
							} `tfsdk:"image" json:"image,omitempty"`
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
									ClusterTrustBundle *struct {
										LabelSelector *struct {
											MatchExpressions *[]struct {
												Key      *string   `tfsdk:"key" json:"key,omitempty"`
												Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
												Values   *[]string `tfsdk:"values" json:"values,omitempty"`
											} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
											MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
										} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
										Name       *string `tfsdk:"name" json:"name,omitempty"`
										Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
										Path       *string `tfsdk:"path" json:"path,omitempty"`
										SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
									} `tfsdk:"cluster_trust_bundle" json:"clusterTrustBundle,omitempty"`
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
						} `tfsdk:"volumes" json:"volumes,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"deployment" json:"deployment,omitempty"`
		DisableDefaultAdminSecret     *bool   `tfsdk:"disable_default_admin_secret" json:"disableDefaultAdminSecret,omitempty"`
		DisableDefaultSecurityContext *string `tfsdk:"disable_default_security_context" json:"disableDefaultSecurityContext,omitempty"`
		External                      *struct {
			AdminPassword *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"admin_password" json:"adminPassword,omitempty"`
			AdminUser *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"admin_user" json:"adminUser,omitempty"`
			ApiKey *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"api_key" json:"apiKey,omitempty"`
			Tls *struct {
				CertSecretRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"cert_secret_ref" json:"certSecretRef,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"external" json:"external,omitempty"`
		Ingress *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				DefaultBackend *struct {
					Resource *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"resource" json:"resource,omitempty"`
					Service *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Port *struct {
							Name   *string `tfsdk:"name" json:"name,omitempty"`
							Number *int64  `tfsdk:"number" json:"number,omitempty"`
						} `tfsdk:"port" json:"port,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"default_backend" json:"defaultBackend,omitempty"`
				IngressClassName *string `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
				Rules            *[]struct {
					Host *string `tfsdk:"host" json:"host,omitempty"`
					Http *struct {
						Paths *[]struct {
							Backend *struct {
								Resource *struct {
									ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
									Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"resource" json:"resource,omitempty"`
								Service *struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
									Port *struct {
										Name   *string `tfsdk:"name" json:"name,omitempty"`
										Number *int64  `tfsdk:"number" json:"number,omitempty"`
									} `tfsdk:"port" json:"port,omitempty"`
								} `tfsdk:"service" json:"service,omitempty"`
							} `tfsdk:"backend" json:"backend,omitempty"`
							Path     *string `tfsdk:"path" json:"path,omitempty"`
							PathType *string `tfsdk:"path_type" json:"pathType,omitempty"`
						} `tfsdk:"paths" json:"paths,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
				Tls *[]struct {
					Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
					SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Jsonnet *struct {
			LibraryLabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"library_label_selector" json:"libraryLabelSelector,omitempty"`
		} `tfsdk:"jsonnet" json:"jsonnet,omitempty"`
		PersistentVolumeClaim *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
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
					Claims *[]struct {
						Name    *string `tfsdk:"name" json:"name,omitempty"`
						Request *string `tfsdk:"request" json:"request,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
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
		} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
		Preferences *struct {
			HomeDashboardUid *string `tfsdk:"home_dashboard_uid" json:"homeDashboardUid,omitempty"`
		} `tfsdk:"preferences" json:"preferences,omitempty"`
		Route *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				AlternateBackends *[]struct {
					Kind   *string `tfsdk:"kind" json:"kind,omitempty"`
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Weight *int64  `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"alternate_backends" json:"alternateBackends,omitempty"`
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Path *string `tfsdk:"path" json:"path,omitempty"`
				Port *struct {
					TargetPort *string `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Subdomain *string `tfsdk:"subdomain" json:"subdomain,omitempty"`
				Tls       *struct {
					CaCertificate            *string `tfsdk:"ca_certificate" json:"caCertificate,omitempty"`
					Certificate              *string `tfsdk:"certificate" json:"certificate,omitempty"`
					DestinationCACertificate *string `tfsdk:"destination_ca_certificate" json:"destinationCACertificate,omitempty"`
					ExternalCertificate      *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"external_certificate" json:"externalCertificate,omitempty"`
					InsecureEdgeTerminationPolicy *string `tfsdk:"insecure_edge_termination_policy" json:"insecureEdgeTerminationPolicy,omitempty"`
					Key                           *string `tfsdk:"key" json:"key,omitempty"`
					Termination                   *string `tfsdk:"termination" json:"termination,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				To *struct {
					Kind   *string `tfsdk:"kind" json:"kind,omitempty"`
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Weight *int64  `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"to" json:"to,omitempty"`
				WildcardPolicy *string `tfsdk:"wildcard_policy" json:"wildcardPolicy,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
		Service *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
				ClusterIP                     *string   `tfsdk:"cluster_ip" json:"clusterIP,omitempty"`
				ClusterIPs                    *[]string `tfsdk:"cluster_i_ps" json:"clusterIPs,omitempty"`
				ExternalIPs                   *[]string `tfsdk:"external_i_ps" json:"externalIPs,omitempty"`
				ExternalName                  *string   `tfsdk:"external_name" json:"externalName,omitempty"`
				ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
				HealthCheckNodePort           *int64    `tfsdk:"health_check_node_port" json:"healthCheckNodePort,omitempty"`
				InternalTrafficPolicy         *string   `tfsdk:"internal_traffic_policy" json:"internalTrafficPolicy,omitempty"`
				IpFamilies                    *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
				IpFamilyPolicy                *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
				LoadBalancerClass             *string   `tfsdk:"load_balancer_class" json:"loadBalancerClass,omitempty"`
				LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
				LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
				Ports                         *[]struct {
					AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
					TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				PublishNotReadyAddresses *bool              `tfsdk:"publish_not_ready_addresses" json:"publishNotReadyAddresses,omitempty"`
				Selector                 *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
				SessionAffinity          *string            `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
				SessionAffinityConfig    *struct {
					ClientIP *struct {
						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
					} `tfsdk:"client_ip" json:"clientIP,omitempty"`
				} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
				TrafficDistribution *string `tfsdk:"traffic_distribution" json:"trafficDistribution,omitempty"`
				Type                *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		ServiceAccount *struct {
			AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" json:"automountServiceAccountToken,omitempty"`
			ImagePullSecrets             *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Secrets *[]struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"secrets" json:"secrets,omitempty"`
		} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		Suspend *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GrafanaIntegreatlyOrgGrafanaV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_grafana_integreatly_org_grafana_v1beta1_manifest"
}

func (r *GrafanaIntegreatlyOrgGrafanaV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Grafana is the Schema for the grafanas API",
		MarkdownDescription: "Grafana is the Schema for the grafanas API",
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
				Description:         "GrafanaSpec defines the desired state of Grafana",
				MarkdownDescription: "GrafanaSpec defines the desired state of Grafana",
				Attributes: map[string]schema.Attribute{
					"client": schema.SingleNestedAttribute{
						Description:         "Client defines how the grafana-operator talks to the grafana instance.",
						MarkdownDescription: "Client defines how the grafana-operator talks to the grafana instance.",
						Attributes: map[string]schema.Attribute{
							"headers": schema.MapAttribute{
								Description:         "Custom HTTP headers to use when interacting with this Grafana.",
								MarkdownDescription: "Custom HTTP headers to use when interacting with this Grafana.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"prefer_ingress": schema.BoolAttribute{
								Description:         "If the operator should send it's request through the grafana instances ingress object instead of through the service.",
								MarkdownDescription: "If the operator should send it's request through the grafana instances ingress object instead of through the service.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS Configuration used to talk with the grafana instance.",
								MarkdownDescription: "TLS Configuration used to talk with the grafana instance.",
								Attributes: map[string]schema.Attribute{
									"cert_secret_ref": schema.SingleNestedAttribute{
										Description:         "Use a secret as a reference to give TLS Certificate information",
										MarkdownDescription: "Use a secret as a reference to give TLS Certificate information",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "Disable the CA check of the server",
										MarkdownDescription: "Disable the CA check of the server",
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

					"config": schema.MapAttribute{
						Description:         "Config defines how your grafana ini file should looks like.",
						MarkdownDescription: "Config defines how your grafana ini file should looks like.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment": schema.SingleNestedAttribute{
						Description:         "Deployment sets how the deployment object should look like with your grafana instance, contains a number of defaults.",
						MarkdownDescription: "Deployment sets how the deployment object should look like with your grafana instance, contains a number of defaults.",
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
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"min_ready_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"paused": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"progress_deadline_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"revision_history_limit": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
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

									"strategy": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"rolling_update": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max_surge": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_unavailable": schema.StringAttribute{
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

											"type": schema.StringAttribute{
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

									"template": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
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
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
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

											"spec": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"active_deadline_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"affinity": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"node_affinity": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"preference": schema.SingleNestedAttribute{
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
																		Description:         "",
																		MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"pod_affinity_term": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"label_selector": schema.SingleNestedAttribute{
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

																						"match_label_keys": schema.ListAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"mismatch_label_keys": schema.ListAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"namespace_selector": schema.SingleNestedAttribute{
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

																						"namespaces": schema.ListAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
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

																				"match_label_keys": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
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

																				"namespaces": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
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
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"pod_affinity_term": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"label_selector": schema.SingleNestedAttribute{
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

																						"match_label_keys": schema.ListAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"mismatch_label_keys": schema.ListAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"namespace_selector": schema.SingleNestedAttribute{
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

																						"namespaces": schema.ListAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
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

																				"match_label_keys": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"mismatch_label_keys": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace_selector": schema.SingleNestedAttribute{
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

																				"namespaces": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
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

													"automount_service_account_token": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"args": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"env": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
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

																			"value_from": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"config_map_key_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																							"optional": schema.BoolAttribute{
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

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"resource": schema.StringAttribute{
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
																						Description:         "",
																						MarkdownDescription: "",
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

																							"optional": schema.BoolAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"config_map_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_pull_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"lifecycle": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"post_start": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"exec": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"value": schema.StringAttribute{
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

																				"sleep": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"seconds": schema.Int64Attribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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

																		"pre_stop": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"exec": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"value": schema.StringAttribute{
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

																				"sleep": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"seconds": schema.Int64Attribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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

																		"stop_signal": schema.StringAttribute{
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

																"liveness_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"ports": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"container_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"host_ip": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"host_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
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

																"readiness_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"resize_policy": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"resource_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"restart_policy": schema.StringAttribute{
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

																"resources": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"request": schema.StringAttribute{
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

																		"limits": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
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

																"restart_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"security_context": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"allow_privilege_escalation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"app_armor_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"capabilities": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"drop": schema.ListAttribute{
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

																		"privileged": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"proc_mount": schema.StringAttribute{
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

																		"se_linux_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"level": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"role": schema.StringAttribute{
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

																				"user": schema.StringAttribute{
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

																		"seccomp_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"windows_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"gmsa_credential_spec": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"gmsa_credential_spec_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"host_process": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"run_as_user_name": schema.StringAttribute{
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

																"startup_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"stdin": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"stdin_once": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"termination_message_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"termination_message_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tty": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_devices": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"device_path": schema.StringAttribute{
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"volume_mounts": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mount_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"mount_propagation": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"recursive_read_only": schema.StringAttribute{
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

																			"sub_path_expr": schema.StringAttribute{
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

																"working_dir": schema.StringAttribute{
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

													"dns_config": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"nameservers": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"options": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"searches": schema.ListAttribute{
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

													"dns_policy": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_service_links": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ephemeral_containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"args": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"env": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
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

																			"value_from": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"config_map_key_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																							"optional": schema.BoolAttribute{
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

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"resource": schema.StringAttribute{
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
																						Description:         "",
																						MarkdownDescription: "",
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

																							"optional": schema.BoolAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"config_map_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_pull_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"lifecycle": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"post_start": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"exec": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"value": schema.StringAttribute{
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

																				"sleep": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"seconds": schema.Int64Attribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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

																		"pre_stop": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"exec": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"value": schema.StringAttribute{
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

																				"sleep": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"seconds": schema.Int64Attribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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

																		"stop_signal": schema.StringAttribute{
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

																"liveness_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"ports": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"container_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"host_ip": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"host_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
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

																"readiness_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"resize_policy": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"resource_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"restart_policy": schema.StringAttribute{
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

																"resources": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"request": schema.StringAttribute{
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

																		"limits": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
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

																"restart_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"security_context": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"allow_privilege_escalation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"app_armor_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"capabilities": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"drop": schema.ListAttribute{
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

																		"privileged": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"proc_mount": schema.StringAttribute{
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

																		"se_linux_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"level": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"role": schema.StringAttribute{
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

																				"user": schema.StringAttribute{
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

																		"seccomp_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"windows_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"gmsa_credential_spec": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"gmsa_credential_spec_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"host_process": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"run_as_user_name": schema.StringAttribute{
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

																"startup_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"stdin": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"stdin_once": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"target_container_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"termination_message_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"termination_message_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tty": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_devices": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"device_path": schema.StringAttribute{
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"volume_mounts": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mount_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"mount_propagation": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"recursive_read_only": schema.StringAttribute{
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

																			"sub_path_expr": schema.StringAttribute{
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

																"working_dir": schema.StringAttribute{
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

													"host_aliases": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"hostnames": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"ip": schema.StringAttribute{
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

													"host_ipc": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_network": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_pid": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_users": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hostname": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"image_pull_secrets": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"init_containers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"args": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"env": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
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

																			"value_from": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"config_map_key_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																							"optional": schema.BoolAttribute{
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

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"resource": schema.StringAttribute{
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
																						Description:         "",
																						MarkdownDescription: "",
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

																							"optional": schema.BoolAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"config_map_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
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

																"image": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_pull_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"lifecycle": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"post_start": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"exec": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"value": schema.StringAttribute{
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

																				"sleep": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"seconds": schema.Int64Attribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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

																		"pre_stop": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"exec": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"value": schema.StringAttribute{
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

																				"sleep": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"seconds": schema.Int64Attribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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

																		"stop_signal": schema.StringAttribute{
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

																"liveness_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"ports": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"container_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"host_ip": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"host_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
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

																"readiness_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"resize_policy": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"resource_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"restart_policy": schema.StringAttribute{
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

																"resources": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"request": schema.StringAttribute{
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

																		"limits": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
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

																"restart_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"security_context": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"allow_privilege_escalation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"app_armor_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"capabilities": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"drop": schema.ListAttribute{
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

																		"privileged": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"proc_mount": schema.StringAttribute{
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

																		"se_linux_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"level": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"role": schema.StringAttribute{
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

																				"user": schema.StringAttribute{
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

																		"seccomp_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"windows_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"gmsa_credential_spec": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"gmsa_credential_spec_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"host_process": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"run_as_user_name": schema.StringAttribute{
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

																"startup_probe": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
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

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
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

																"stdin": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"stdin_once": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"termination_message_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"termination_message_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tty": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_devices": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"device_path": schema.StringAttribute{
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"volume_mounts": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mount_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"mount_propagation": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"recursive_read_only": schema.StringAttribute{
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

																			"sub_path_expr": schema.StringAttribute{
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

																"working_dir": schema.StringAttribute{
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

													"node_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_selector": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"os": schema.SingleNestedAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"overhead": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preemption_policy": schema.StringAttribute{
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

													"priority_class_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"readiness_gates": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"condition_type": schema.StringAttribute{
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

													"restart_policy": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"runtime_class_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scheduler_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_context": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"app_armor_profile": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"localhost_profile": schema.StringAttribute{
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

															"se_linux_change_policy": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"se_linux_options": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"role": schema.StringAttribute{
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

																	"user": schema.StringAttribute{
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

															"seccomp_profile": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"localhost_profile": schema.StringAttribute{
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

															"supplemental_groups_policy": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sysctls": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"value": schema.StringAttribute{
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

															"windows_options": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"gmsa_credential_spec": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"gmsa_credential_spec_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host_process": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"run_as_user_name": schema.StringAttribute{
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

													"service_account": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_account_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"set_hostname_as_fqdn": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"share_process_namespace": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"subdomain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"termination_grace_period_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tolerations": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"effect": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"toleration_seconds": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_spread_constraints": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"node_taints_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"volumes": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"aws_elastic_block_store": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_id": schema.StringAttribute{
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

																"azure_disk": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"caching_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"disk_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"disk_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
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

																"azure_file": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"share_name": schema.StringAttribute{
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

																"cephfs": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"monitors": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_file": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"user": schema.StringAttribute{
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

																"cinder": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"volume_id": schema.StringAttribute{
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

																"config_map": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
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

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
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

																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"optional": schema.BoolAttribute{
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

																"csi": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"node_publish_secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_attributes": schema.MapAttribute{
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

																"downward_api": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"field_ref": schema.SingleNestedAttribute{
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

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"resource": schema.StringAttribute{
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

																"empty_dir": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"medium": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"size_limit": schema.StringAttribute{
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

																"ephemeral": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"volume_claim_template": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"metadata": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"spec": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"access_modes": schema.ListAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"data_source": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"api_group": schema.StringAttribute{
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
																							Required: false,
																							Optional: true,
																							Computed: false,
																						},

																						"data_source_ref": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"api_group": schema.StringAttribute{
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

																								"namespace": schema.StringAttribute{
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

																						"resources": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"limits": schema.MapAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"requests": schema.MapAttribute{
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

																						"storage_class_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_attributes_class_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_mode": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_name": schema.StringAttribute{
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"fc": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"target_ww_ns": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"wwids": schema.ListAttribute{
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

																"flex_volume": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"options": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"flocker": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"dataset_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"dataset_uuid": schema.StringAttribute{
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

																"gce_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pd_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"git_repo": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"directory": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"repository": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"revision": schema.StringAttribute{
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

																"glusterfs": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"endpoints": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"host_path": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
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
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"image": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"pull_policy": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"reference": schema.StringAttribute{
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

																"iscsi": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"chap_auth_discovery": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"chap_auth_session": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"initiator_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"iqn": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"iscsi_interface": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"portals": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"target_portal": schema.StringAttribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"nfs": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"server": schema.StringAttribute{
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

																"persistent_volume_claim": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claim_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"photon_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pd_id": schema.StringAttribute{
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

																"portworx_volume": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_id": schema.StringAttribute{
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

																"projected": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sources": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"cluster_trust_bundle": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"label_selector": schema.SingleNestedAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"signer_name": schema.StringAttribute{
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

																					"config_map": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
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

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
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

																					"downward_api": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"field_ref": schema.SingleNestedAttribute{
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

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"resource_field_ref": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"container_name": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            true,
																													Computed:            false,
																												},

																												"divisor": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            true,
																													Computed:            false,
																												},

																												"resource": schema.StringAttribute{
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

																					"secret": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
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

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
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

																					"service_account_token": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"audience": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"expiration_seconds": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
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

																"quobyte": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"registry": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"tenant": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"user": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume": schema.StringAttribute{
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

																"rbd": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"image": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"keyring": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"monitors": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"pool": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"user": schema.StringAttribute{
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

																"scale_io": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"gateway": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"protection_domain": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
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

																		"ssl_enabled": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_pool": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"system": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"volume_name": schema.StringAttribute{
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

																"secret": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
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

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
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

																		"optional": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_name": schema.StringAttribute{
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

																"storageos": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																		"volume_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_namespace": schema.StringAttribute{
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

																"vsphere_volume": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_policy_id": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_policy_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_path": schema.StringAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_default_admin_secret": schema.BoolAttribute{
						Description:         "DisableDefaultAdminSecret prevents operator from creating default admin-credentials secret",
						MarkdownDescription: "DisableDefaultAdminSecret prevents operator from creating default admin-credentials secret",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_default_security_context": schema.StringAttribute{
						Description:         "DisableDefaultSecurityContext prevents the operator from populating securityContext on deployments",
						MarkdownDescription: "DisableDefaultSecurityContext prevents the operator from populating securityContext on deployments",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Pod", "Container", "All"),
						},
					},

					"external": schema.SingleNestedAttribute{
						Description:         "External enables you to configure external grafana instances that is not managed by the operator.",
						MarkdownDescription: "External enables you to configure external grafana instances that is not managed by the operator.",
						Attributes: map[string]schema.Attribute{
							"admin_password": schema.SingleNestedAttribute{
								Description:         "AdminPassword key to talk to the external grafana instance.",
								MarkdownDescription: "AdminPassword key to talk to the external grafana instance.",
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

							"admin_user": schema.SingleNestedAttribute{
								Description:         "AdminUser key to talk to the external grafana instance.",
								MarkdownDescription: "AdminUser key to talk to the external grafana instance.",
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

							"api_key": schema.SingleNestedAttribute{
								Description:         "The API key to talk to the external grafana instance, you need to define ether apiKey or adminUser/adminPassword.",
								MarkdownDescription: "The API key to talk to the external grafana instance, you need to define ether apiKey or adminUser/adminPassword.",
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

							"tls": schema.SingleNestedAttribute{
								Description:         "DEPRECATED, use top level 'tls' instead.",
								MarkdownDescription: "DEPRECATED, use top level 'tls' instead.",
								Attributes: map[string]schema.Attribute{
									"cert_secret_ref": schema.SingleNestedAttribute{
										Description:         "Use a secret as a reference to give TLS Certificate information",
										MarkdownDescription: "Use a secret as a reference to give TLS Certificate information",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is unique within a namespace to reference a secret resource.",
												MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "namespace defines the space within which the secret name must be unique.",
												MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "Disable the CA check of the server",
										MarkdownDescription: "Disable the CA check of the server",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": schema.StringAttribute{
								Description:         "URL of the external grafana instance you want to manage.",
								MarkdownDescription: "URL of the external grafana instance you want to manage.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": schema.SingleNestedAttribute{
						Description:         "Ingress sets how the ingress object should look like with your grafana instance.",
						MarkdownDescription: "Ingress sets how the ingress object should look like with your grafana instance.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								MarkdownDescription: "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"spec": schema.SingleNestedAttribute{
								Description:         "IngressSpec describes the Ingress the user wishes to exist.",
								MarkdownDescription: "IngressSpec describes the Ingress the user wishes to exist.",
								Attributes: map[string]schema.Attribute{
									"default_backend": schema.SingleNestedAttribute{
										Description:         "defaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified, DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the rules will be up to the Ingress controller.",
										MarkdownDescription: "defaultBackend is the backend that should handle requests that don't match any rule. If Rules are not specified, DefaultBackend must be specified. If DefaultBackend is not set, the handling of requests that do not match any of the rules will be up to the Ingress controller.",
										Attributes: map[string]schema.Attribute{
											"resource": schema.SingleNestedAttribute{
												Description:         "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
												MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
														MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the type of resource being referenced",
														MarkdownDescription: "Kind is the type of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of resource being referenced",
														MarkdownDescription: "Name is the name of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"service": schema.SingleNestedAttribute{
												Description:         "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
												MarkdownDescription: "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
														MarkdownDescription: "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"port": schema.SingleNestedAttribute{
														Description:         "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
														MarkdownDescription: "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																MarkdownDescription: "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"number": schema.Int64Attribute{
																Description:         "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
																MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
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

									"ingress_class_name": schema.StringAttribute{
										Description:         "ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the 'kubernetes.io/ingress.class' annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",
										MarkdownDescription: "ingressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource, by a transitive connection (controller -> IngressClass -> Ingress resource). Although the 'kubernetes.io/ingress.class' annotation (simple constant name) was never formally defined, it was widely supported by Ingress controllers to create a direct binding between Ingress controller and Ingress resources. Newly created Ingress resources should prefer using the field. However, even though the annotation is officially deprecated, for backwards compatibility reasons, ingress controllers should still honor that annotation if present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"rules": schema.ListNestedAttribute{
										Description:         "rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
										MarkdownDescription: "rules is a list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The ':' delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue. host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If host is precise, the request matches this rule if the http host header is equal to Host. 2. If host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
													MarkdownDescription: "host is the fully qualified domain name of a network host, as defined by RFC 3986. Note the following deviations from the 'host' part of the URI as defined in RFC 3986: 1. IPs are not allowed. Currently an IngressRuleValue can only apply to the IP in the Spec of the parent Ingress. 2. The ':' delimiter is not respected because ports are not allowed. Currently the port of an Ingress is implicitly :80 for http and :443 for https. Both these may change in the future. Incoming requests are matched against the host before the IngressRuleValue. If the host is unspecified, the Ingress routes all traffic based on the specified IngressRuleValue. host can be 'precise' which is a domain name without the terminating dot of a network host (e.g. 'foo.bar.com') or 'wildcard', which is a domain name prefixed with a single wildcard label (e.g. '*.foo.com'). The wildcard character '*' must appear by itself as the first DNS label and matches only a single label. You cannot have a wildcard label by itself (e.g. Host == '*'). Requests will be matched against the Host field in the following way: 1. If host is precise, the request matches this rule if the http host header is equal to Host. 2. If host is a wildcard, then the request matches this rule if the http host header is to equal to the suffix (removing the first label) of the wildcard rule.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http": schema.SingleNestedAttribute{
													Description:         "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
													MarkdownDescription: "HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.",
													Attributes: map[string]schema.Attribute{
														"paths": schema.ListNestedAttribute{
															Description:         "paths is a collection of paths that map requests to backends.",
															MarkdownDescription: "paths is a collection of paths that map requests to backends.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"backend": schema.SingleNestedAttribute{
																		Description:         "backend defines the referenced service endpoint to which the traffic will be forwarded to.",
																		MarkdownDescription: "backend defines the referenced service endpoint to which the traffic will be forwarded to.",
																		Attributes: map[string]schema.Attribute{
																			"resource": schema.SingleNestedAttribute{
																				Description:         "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
																				MarkdownDescription: "resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with 'Service'.",
																				Attributes: map[string]schema.Attribute{
																					"api_group": schema.StringAttribute{
																						Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																						MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"kind": schema.StringAttribute{
																						Description:         "Kind is the type of resource being referenced",
																						MarkdownDescription: "Kind is the type of resource being referenced",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name is the name of resource being referenced",
																						MarkdownDescription: "Name is the name of resource being referenced",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"service": schema.SingleNestedAttribute{
																				Description:         "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
																				MarkdownDescription: "service references a service as a backend. This is a mutually exclusive setting with 'Resource'.",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
																						MarkdownDescription: "name is the referenced service. The service must exist in the same namespace as the Ingress object.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"port": schema.SingleNestedAttribute{
																						Description:         "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
																						MarkdownDescription: "port of the referenced service. A port name or port number is required for a IngressServiceBackend.",
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																								MarkdownDescription: "name is the name of the port on the Service. This is a mutually exclusive setting with 'Number'.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"number": schema.Int64Attribute{
																								Description:         "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
																								MarkdownDescription: "number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with 'Name'.",
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
																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"path": schema.StringAttribute{
																		Description:         "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",
																		MarkdownDescription: "path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986. Paths must begin with a '/' and must be present when using PathType with value 'Exact' or 'Prefix'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"path_type": schema.StringAttribute{
																		Description:         "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.",
																		MarkdownDescription: "pathType determines the interpretation of the path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by '/'. Matching is done on a path element by element basis. A path element refers is the list of labels in the path split by the '/' separator. A request is a match for path p if every p is an element-wise prefix of p of the request path. Note that if the last element of the path is a substring of the last element in request path, it is not a match (e.g. /foo/bar matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to the IngressClass. Implementations can treat this as a separate PathType or treat it identically to Prefix or Exact path types. Implementations are required to support all path types.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": schema.ListNestedAttribute{
										Description:         "tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
										MarkdownDescription: "tls represents the TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hosts": schema.ListAttribute{
													Description:         "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
													MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
													MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the 'Host' header is used for routing.",
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

					"jsonnet": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"library_label_selector": schema.SingleNestedAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"persistent_volume_claim": schema.SingleNestedAttribute{
						Description:         "PersistentVolumeClaim creates a PVC if you need to attach one to your grafana instance.",
						MarkdownDescription: "PersistentVolumeClaim creates a PVC if you need to attach one to your grafana instance.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								MarkdownDescription: "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"data_source": schema.SingleNestedAttribute{
										Description:         "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",
										MarkdownDescription: "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"data_source_ref": schema.SingleNestedAttribute{
										Description:         "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",
										MarkdownDescription: "TypedLocalObjectReference contains enough information to let you locate the typed referenced object inside the same namespace.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "ResourceRequirements describes the compute resource requirements.",
										MarkdownDescription: "ResourceRequirements describes the compute resource requirements.",
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

									"volume_mode": schema.StringAttribute{
										Description:         "PersistentVolumeMode describes how a volume is intended to be consumed, either Block or Filesystem.",
										MarkdownDescription: "PersistentVolumeMode describes how a volume is intended to be consumed, either Block or Filesystem.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_name": schema.StringAttribute{
										Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
										MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",
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

					"preferences": schema.SingleNestedAttribute{
						Description:         "Preferences holds the Grafana Preferences settings",
						MarkdownDescription: "Preferences holds the Grafana Preferences settings",
						Attributes: map[string]schema.Attribute{
							"home_dashboard_uid": schema.StringAttribute{
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

					"route": schema.SingleNestedAttribute{
						Description:         "Route sets how the ingress object should look like with your grafana instance, this only works in Openshift.",
						MarkdownDescription: "Route sets how the ingress object should look like with your grafana instance, this only works in Openshift.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								MarkdownDescription: "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"alternate_backends": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"kind": schema.StringAttribute{
													Description:         "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
													MarkdownDescription: "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Service", ""),
													},
												},

												"name": schema.StringAttribute{
													Description:         "name of the service/target that is being referred to. e.g. name of the service",
													MarkdownDescription: "name of the service/target that is being referred to. e.g. name of the service",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},

												"weight": schema.Int64Attribute{
													Description:         "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
													MarkdownDescription: "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
														int64validator.AtMost(256),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

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

									"port": schema.SingleNestedAttribute{
										Description:         "RoutePort defines a port mapping from a router to an endpoint in the service endpoints.",
										MarkdownDescription: "RoutePort defines a port mapping from a router to an endpoint in the service endpoints.",
										Attributes: map[string]schema.Attribute{
											"target_port": schema.StringAttribute{
												Description:         "The target port on pods selected by the service this route points to. If this is a string, it will be looked up as a named port in the target endpoints port list. Required",
												MarkdownDescription: "The target port on pods selected by the service this route points to. If this is a string, it will be looked up as a named port in the target endpoints port list. Required",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"subdomain": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLSConfig defines config used to secure a route and provide termination",
										MarkdownDescription: "TLSConfig defines config used to secure a route and provide termination",
										Attributes: map[string]schema.Attribute{
											"ca_certificate": schema.StringAttribute{
												Description:         "caCertificate provides the cert authority certificate contents",
												MarkdownDescription: "caCertificate provides the cert authority certificate contents",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"certificate": schema.StringAttribute{
												Description:         "certificate provides certificate contents. This should be a single serving certificate, not a certificate chain. Do not include a CA certificate.",
												MarkdownDescription: "certificate provides certificate contents. This should be a single serving certificate, not a certificate chain. Do not include a CA certificate.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"destination_ca_certificate": schema.StringAttribute{
												Description:         "destinationCACertificate provides the contents of the ca certificate of the final destination. When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
												MarkdownDescription: "destinationCACertificate provides the contents of the ca certificate of the final destination. When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_certificate": schema.SingleNestedAttribute{
												Description:         "externalCertificate provides certificate contents as a secret reference. This should be a single serving certificate, not a certificate chain. Do not include a CA certificate. The secret referenced should be present in the same namespace as that of the Route. Forbidden when 'certificate' is set. The router service account needs to be granted with read-only access to this secret, please refer to openshift docs for additional details.",
												MarkdownDescription: "externalCertificate provides certificate contents as a secret reference. This should be a single serving certificate, not a certificate chain. Do not include a CA certificate. The secret referenced should be present in the same namespace as that of the Route. Forbidden when 'certificate' is set. The router service account needs to be granted with read-only access to this secret, please refer to openshift docs for additional details.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_edge_termination_policy": schema.StringAttribute{
												Description:         "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80. If a route does not specify insecureEdgeTerminationPolicy, then the default behavior is 'None'. * Allow - traffic is sent to the server on the insecure port (edge/reencrypt terminations only). * None - no traffic is allowed on the insecure port (default). * Redirect - clients are redirected to the secure port.",
												MarkdownDescription: "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80. If a route does not specify insecureEdgeTerminationPolicy, then the default behavior is 'None'. * Allow - traffic is sent to the server on the insecure port (edge/reencrypt terminations only). * None - no traffic is allowed on the insecure port (default). * Redirect - clients are redirected to the secure port.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Allow", "None", "Redirect", ""),
												},
											},

											"key": schema.StringAttribute{
												Description:         "key provides key file contents",
												MarkdownDescription: "key provides key file contents",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"termination": schema.StringAttribute{
												Description:         "termination indicates termination type. * edge - TLS termination is done by the router and http is used to communicate with the backend (default) * passthrough - Traffic is sent straight to the destination without the router providing TLS termination * reencrypt - TLS termination is done by the router and https is used to communicate with the backend Note: passthrough termination is incompatible with httpHeader actions",
												MarkdownDescription: "termination indicates termination type. * edge - TLS termination is done by the router and http is used to communicate with the backend (default) * passthrough - Traffic is sent straight to the destination without the router providing TLS termination * reencrypt - TLS termination is done by the router and https is used to communicate with the backend Note: passthrough termination is incompatible with httpHeader actions",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("edge", "reencrypt", "passthrough"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to": schema.SingleNestedAttribute{
										Description:         "RouteTargetReference specifies the target that resolve into endpoints. Only the 'Service' kind is allowed. Use 'weight' field to emphasize one over others.",
										MarkdownDescription: "RouteTargetReference specifies the target that resolve into endpoints. Only the 'Service' kind is allowed. Use 'weight' field to emphasize one over others.",
										Attributes: map[string]schema.Attribute{
											"kind": schema.StringAttribute{
												Description:         "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
												MarkdownDescription: "The kind of target that the route is referring to. Currently, only 'Service' is allowed",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Service", ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "name of the service/target that is being referred to. e.g. name of the service",
												MarkdownDescription: "name of the service/target that is being referred to. e.g. name of the service",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"weight": schema.Int64Attribute{
												Description:         "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
												MarkdownDescription: "weight as an integer between 0 and 256, default 100, that specifies the target's relative weight against other target reference objects. 0 suppresses requests to this backend.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(256),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"wildcard_policy": schema.StringAttribute{
										Description:         "WildcardPolicyType indicates the type of wildcard support needed by routes.",
										MarkdownDescription: "WildcardPolicyType indicates the type of wildcard support needed by routes.",
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

					"service": schema.SingleNestedAttribute{
						Description:         "Service sets how the service object should look like with your grafana instance, contains a number of defaults.",
						MarkdownDescription: "Service sets how the service object should look like with your grafana instance, contains a number of defaults.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								MarkdownDescription: "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"spec": schema.SingleNestedAttribute{
								Description:         "ServiceSpec describes the attributes that a user creates on a service.",
								MarkdownDescription: "ServiceSpec describes the attributes that a user creates on a service.",
								Attributes: map[string]schema.Attribute{
									"allocate_load_balancer_node_ports": schema.BoolAttribute{
										Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
										MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automatically allocated for services with type LoadBalancer. Default is 'true'. It may be set to 'false' if the cluster load-balancer does not rely on NodePorts. If the caller requests specific NodePorts (by specifying a value), those requests will be respected, regardless of this field. This field may only be set for services with type LoadBalancer and will be cleared if the type is changed to any other type.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_ip": schema.StringAttribute{
										Description:         "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										MarkdownDescription: "clusterIP is the IP address of the service and is usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be blank) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_i_ps": schema.ListAttribute{
										Description:         "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value. This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and are usually assigned randomly. If an address is specified manually, is in-range (as per system configuration), and is not in use, it will be allocated to the service; otherwise creation of the service will fail. This field may not be changed through updates unless the type field is also being changed to ExternalName (which requires this field to be empty) or the type field is being changed from ExternalName (in which case this field may optionally be specified, as describe above). Valid values are 'None', empty string (''), or a valid IP address. Setting this to 'None' makes a 'headless service' (no virtual IP), which is useful when direct endpoint connections are preferred and proxying is not required. Only applies to types ClusterIP, NodePort, and LoadBalancer. If this field is specified when creating a Service of type ExternalName, creation will fail. This field will be wiped when updating a Service to type ExternalName. If this field is not specified, it will be initialized from the clusterIP field. If this field is specified, clients must ensure that clusterIPs[0] and clusterIP have the same value. This field may hold a maximum of two entries (dual-stack IPs, in either order). These IPs must correspond to the values of the ipFamilies field. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_i_ps": schema.ListAttribute{
										Description:         "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.",
										MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the cluster will also accept traffic for this service. These IPs are not managed by Kubernetes. The user is responsible for ensuring that traffic arrives at a node with this IP. A common example is external load-balancers that are not part of the Kubernetes system.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_name": schema.StringAttribute{
										Description:         "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
										MarkdownDescription: "externalName is the external reference that discovery mechanisms will return as an alias for this service (e.g. a DNS CNAME record). No proxying will be involved. Must be a lowercase RFC-1123 hostname (https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_traffic_policy": schema.StringAttribute{
										Description:         "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
										MarkdownDescription: "externalTrafficPolicy describes how nodes distribute service traffic they receive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configure the service in a way that assumes that external load balancers will take care of balancing the service traffic between nodes, and so each node will deliver traffic only to the node-local endpoints of the service, without masquerading the client source IP. (Traffic mistakenly sent to a node with no endpoints will be dropped.) The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features). Note that traffic sent to an External IP or LoadBalancer IP from within the cluster will always get 'Cluster' semantics, but clients sending to a NodePort from within the cluster may need to take traffic policy into account when picking a node.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"health_check_node_port": schema.Int64Attribute{
										Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
										MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service. This only applies when type is set to LoadBalancer and externalTrafficPolicy is set to Local. If a value is specified, is in-range, and is not in use, it will be used. If not specified, a value will be automatically allocated. External systems (e.g. load-balancers) can use this port to determine if a given node holds endpoints for this service or not. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type). This field cannot be updated once set.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"internal_traffic_policy": schema.StringAttribute{
										Description:         "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
										MarkdownDescription: "InternalTrafficPolicy describes how nodes distribute service traffic they receive on the ClusterIP. If set to 'Local', the proxy will assume that pods only want to talk to endpoints of the service on the same node as the pod, dropping the traffic if there are no local endpoints. The default value, 'Cluster', uses the standard behavior of routing to all endpoints evenly (possibly modified by topology and other features).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ip_families": schema.ListAttribute{
										Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName. This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
										MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to this service. This field is usually assigned automatically based on cluster configuration and the ipFamilyPolicy field. If this field is specified manually, the requested family is available in the cluster, and ipFamilyPolicy allows it, it will be used; otherwise creation of the service will fail. This field is conditionally mutable: it allows for adding or removing a secondary IP family, but it does not allow changing the primary IP family of the Service. Valid values are 'IPv4' and 'IPv6'. This field only applies to Services of types ClusterIP, NodePort, and LoadBalancer, and does apply to 'headless' services. This field will be wiped when updating a Service to type ExternalName. This field may hold a maximum of two entries (dual-stack families, in either order). These families must correspond to the values of the clusterIPs field, if specified. Both clusterIPs and ipFamilies are governed by the ipFamilyPolicy field.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ip_family_policy": schema.StringAttribute{
										Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
										MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required by this Service. If there is no value provided, then this field will be set to SingleStack. Services can be 'SingleStack' (a single IP family), 'PreferDualStack' (two IP families on dual-stack configured clusters or a single IP family on single-stack clusters), or 'RequireDualStack' (two IP families on dual-stack configured clusters, otherwise fail). The ipFamilies and clusterIPs fields depend on the value of this field. This field will be wiped when updating a service to type ExternalName.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"load_balancer_class": schema.StringAttribute{
										Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
										MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to. If specified, the value of this field must be a label-style identifier, with an optional prefix, e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users. This field can only be set when the Service type is 'LoadBalancer'. If not set, the default load balancer implementation is used, today this is typically done through the cloud provider integration, but should apply for any default implementation. If set, it is assumed that a load balancer implementation is watching for Services with a matching class. Any default load balancer implementation (e.g. cloud providers) should ignore Services that set this field. This field can only be set when creating or updating a Service to type 'LoadBalancer'. Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"load_balancer_ip": schema.StringAttribute{
										Description:         "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations. Using it is non-portable and it may not support dual-stack. Users are encouraged to use implementation-specific annotations when available.",
										MarkdownDescription: "Only applies to Service Type: LoadBalancer. This feature depends on whether the underlying cloud-provider supports specifying the loadBalancerIP when a load balancer is created. This field will be ignored if the cloud-provider does not support the feature. Deprecated: This field was under-specified and its meaning varies across implementations. Using it is non-portable and it may not support dual-stack. Users are encouraged to use implementation-specific annotations when available.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"load_balancer_source_ranges": schema.ListAttribute{
										Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
										MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-provider load-balancer will be restricted to the specified client IPs. This field will be ignored if the cloud-provider does not support the feature.' More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ports": schema.ListNestedAttribute{
										Description:         "The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										MarkdownDescription: "The list of ports that are exposed by this service. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"app_protocol": schema.StringAttribute{
													Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
													MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either: * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior- * 'kubernetes.io/ws' - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455 * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
													MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_port": schema.Int64Attribute{
													Description:         "The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this Service requires one. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
													MarkdownDescription: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer. Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail. If not specified, a port will be allocated if this Service requires one. If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The port that will be exposed by this service.",
													MarkdownDescription: "The port that will be exposed by this service.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
													MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target_port": schema.StringAttribute{
													Description:         "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
													MarkdownDescription: "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
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

									"publish_not_ready_addresses": schema.BoolAttribute{
										Description:         "publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.",
										MarkdownDescription: "publishNotReadyAddresses indicates that any agent which deals with endpoints for this Service should disregard any indications of ready/not-ready. The primary use case for setting this field is for a StatefulSet's Headless Service to propagate SRV DNS records for its Pods for the purpose of peer discovery. The Kubernetes controllers that generate Endpoints and EndpointSlice resources for Services interpret this to mean that all endpoints are considered 'ready' even if the Pods themselves are not. Agents which consume only Kubernetes generated endpoints through the Endpoints or EndpointSlice resources can safely assume this behavior.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"selector": schema.MapAttribute{
										Description:         "Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										MarkdownDescription: "Route service traffic to pods with label keys and values matching this selector. If empty or not present, the service is assumed to have an external process managing its endpoints, which Kubernetes will not modify. Only applies to types ClusterIP, NodePort, and LoadBalancer. Ignored if type is ExternalName. More info: https://kubernetes.io/docs/concepts/services-networking/service/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"session_affinity": schema.StringAttribute{
										Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"session_affinity_config": schema.SingleNestedAttribute{
										Description:         "sessionAffinityConfig contains the configurations of session affinity.",
										MarkdownDescription: "sessionAffinityConfig contains the configurations of session affinity.",
										Attributes: map[string]schema.Attribute{
											"client_ip": schema.SingleNestedAttribute{
												Description:         "clientIP contains the configurations of Client IP based session affinity.",
												MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",
												Attributes: map[string]schema.Attribute{
													"timeout_seconds": schema.Int64Attribute{
														Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
														MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
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

									"traffic_distribution": schema.StringAttribute{
										Description:         "TrafficDistribution offers a way to express preferences for how traffic is distributed to Service endpoints. Implementations can use this field as a hint, but are not required to guarantee strict adherence. If the field is not set, the implementation will apply its default routing strategy. If set to 'PreferClose', implementations should prioritize endpoints that are in the same zone.",
										MarkdownDescription: "TrafficDistribution offers a way to express preferences for how traffic is distributed to Service endpoints. Implementations can use this field as a hint, but are not required to guarantee strict adherence. If the field is not set, the implementation will apply its default routing strategy. If set to 'PreferClose', implementations should prioritize endpoints that are in the same zone.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
										MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. 'ExternalName' aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
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

					"service_account": schema.SingleNestedAttribute{
						Description:         "ServiceAccount sets how the ServiceAccount object should look like with your grafana instance, contains a number of defaults.",
						MarkdownDescription: "ServiceAccount sets how the ServiceAccount object should look like with your grafana instance, contains a number of defaults.",
						Attributes: map[string]schema.Attribute{
							"automount_service_account_token": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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

							"metadata": schema.SingleNestedAttribute{
								Description:         "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								MarkdownDescription: "ObjectMeta contains only a [subset of the fields included in k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta).",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
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

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"field_path": schema.StringAttribute{
											Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
											MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
											Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
											MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"resource_version": schema.StringAttribute{
											Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
											MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uid": schema.StringAttribute{
											Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
											MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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

					"suspend": schema.BoolAttribute{
						Description:         "Suspend pauses reconciliation of owned resources like deployments, Services, Etc. upon changes",
						MarkdownDescription: "Suspend pauses reconciliation of owned resources like deployments, Services, Etc. upon changes",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"version": schema.StringAttribute{
						Description:         "Version specifies the version of Grafana to use for this deployment. It follows the same format as the docker.io/grafana/grafana tags",
						MarkdownDescription: "Version specifies the version of Grafana to use for this deployment. It follows the same format as the docker.io/grafana/grafana tags",
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

func (r *GrafanaIntegreatlyOrgGrafanaV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_grafana_integreatly_org_grafana_v1beta1_manifest")

	var model GrafanaIntegreatlyOrgGrafanaV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("grafana.integreatly.org/v1beta1")
	model.Kind = pointer.String("Grafana")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
