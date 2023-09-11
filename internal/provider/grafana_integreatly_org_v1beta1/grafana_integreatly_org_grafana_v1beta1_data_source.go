/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package grafana_integreatly_org_v1beta1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource{}
)

func NewGrafanaIntegreatlyOrgGrafanaV1Beta1DataSource() datasource.DataSource {
	return &GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource{}
}

type GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GrafanaIntegreatlyOrgGrafanaV1Beta1DataSourceData struct {
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
		Client *struct {
			PreferIngress *bool  `tfsdk:"prefer_ingress" json:"preferIngress,omitempty"`
			Timeout       *int64 `tfsdk:"timeout" json:"timeout,omitempty"`
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
							ResizePolicy *[]struct {
								ResourceName  *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
								RestartPolicy *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							} `tfsdk:"resize_policy" json:"resizePolicy,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
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
							ResizePolicy *[]struct {
								ResourceName  *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
								RestartPolicy *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							} `tfsdk:"resize_policy" json:"resizePolicy,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
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
							TargetContainerName      *string `tfsdk:"target_container_name" json:"targetContainerName,omitempty"`
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
							ResizePolicy *[]struct {
								ResourceName  *string `tfsdk:"resource_name" json:"resourceName,omitempty"`
								RestartPolicy *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
							} `tfsdk:"resize_policy" json:"resizePolicy,omitempty"`
							Resources *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
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
							Sysctls            *[]struct {
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
											Claims *[]struct {
												Name *string `tfsdk:"name" json:"name,omitempty"`
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
						} `tfsdk:"volumes" json:"volumes,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"deployment" json:"deployment,omitempty"`
		External *struct {
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
						Name *string `tfsdk:"name" json:"name,omitempty"`
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
				Tls *struct {
					CaCertificate                 *string `tfsdk:"ca_certificate" json:"caCertificate,omitempty"`
					Certificate                   *string `tfsdk:"certificate" json:"certificate,omitempty"`
					DestinationCACertificate      *string `tfsdk:"destination_ca_certificate" json:"destinationCACertificate,omitempty"`
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
				Type *string `tfsdk:"type" json:"type,omitempty"`
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
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_grafana_integreatly_org_grafana_v1beta1"
}

func (r *GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"client": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"prefer_ingress": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout": schema.Int64Attribute{
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

					"config": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"deployment": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"min_ready_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"paused": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"progress_deadline_seconds": schema.Int64Attribute{
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

									"revision_history_limit": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"operator": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"values": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
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
														Optional:            false,
														Computed:            true,
													},

													"max_unavailable": schema.StringAttribute{
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
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"spec": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"active_deadline_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
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
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																									Description:         "",
																									MarkdownDescription: "",
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
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																									Description:         "",
																									MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"topology_key": schema.StringAttribute{
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

																				"weight": schema.Int64Attribute{
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
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
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
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
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
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"topology_key": schema.StringAttribute{
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
																},
																Required: false,
																Optional: false,
																Computed: true,
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
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																									Description:         "",
																									MarkdownDescription: "",
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
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																									Description:         "",
																									MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"topology_key": schema.StringAttribute{
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

																				"weight": schema.Int64Attribute{
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
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
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
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"operator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"values": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
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
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"topology_key": schema.StringAttribute{
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

													"automount_service_account_token": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
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
																	Optional:            false,
																	Computed:            true,
																},

																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"env": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
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

																							"optional": schema.BoolAttribute{
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

																					"field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"api_version": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"field_path": schema.StringAttribute{
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

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"resource": schema.StringAttribute{
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

																					"secret_key_ref": schema.SingleNestedAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"optional": schema.BoolAttribute{
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
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"optional": schema.BoolAttribute{
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

																			"prefix": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"secret_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"optional": schema.BoolAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
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
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"value": schema.StringAttribute{
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

																						"path": schema.StringAttribute{
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

																						"scheme": schema.StringAttribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"value": schema.StringAttribute{
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

																						"path": schema.StringAttribute{
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

																						"scheme": schema.StringAttribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"ports": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"container_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"host_ip": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"host_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"protocol": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"resize_policy": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"resource_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"restart_policy": schema.StringAttribute{
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

																		"limits": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																"restart_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"security_context": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"allow_privilege_escalation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
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
																					Optional:            false,
																					Computed:            true,
																				},

																				"drop": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"proc_mount": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only_root_filesystem": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_group": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_non_root": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_user": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"se_linux_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"level": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"role": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"user": schema.StringAttribute{
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

																		"seccomp_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
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

																		"windows_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"gmsa_credential_spec": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"gmsa_credential_spec_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"host_process": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"run_as_user_name": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"stdin": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"stdin_once": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"termination_message_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"termination_message_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tty": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"volume_devices": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"device_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"volume_mounts": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mount_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"mount_propagation": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"read_only": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"sub_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"sub_path_expr": schema.StringAttribute{
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

																"working_dir": schema.StringAttribute{
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

													"dns_config": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"nameservers": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
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
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.StringAttribute{
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

															"searches": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
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

													"dns_policy": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"enable_service_links": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
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
																	Optional:            false,
																	Computed:            true,
																},

																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"env": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
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

																							"optional": schema.BoolAttribute{
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

																					"field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"api_version": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"field_path": schema.StringAttribute{
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

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"resource": schema.StringAttribute{
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

																					"secret_key_ref": schema.SingleNestedAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"optional": schema.BoolAttribute{
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
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"optional": schema.BoolAttribute{
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

																			"prefix": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"secret_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"optional": schema.BoolAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
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
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"value": schema.StringAttribute{
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

																						"path": schema.StringAttribute{
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

																						"scheme": schema.StringAttribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"value": schema.StringAttribute{
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

																						"path": schema.StringAttribute{
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

																						"scheme": schema.StringAttribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"ports": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"container_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"host_ip": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"host_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"protocol": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"resize_policy": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"resource_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"restart_policy": schema.StringAttribute{
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

																		"limits": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																"restart_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"security_context": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"allow_privilege_escalation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
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
																					Optional:            false,
																					Computed:            true,
																				},

																				"drop": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"proc_mount": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only_root_filesystem": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_group": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_non_root": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_user": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"se_linux_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"level": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"role": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"user": schema.StringAttribute{
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

																		"seccomp_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
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

																		"windows_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"gmsa_credential_spec": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"gmsa_credential_spec_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"host_process": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"run_as_user_name": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"stdin": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"stdin_once": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"target_container_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"termination_message_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"termination_message_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tty": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"volume_devices": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"device_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"volume_mounts": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mount_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"mount_propagation": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"read_only": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"sub_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"sub_path_expr": schema.StringAttribute{
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

																"working_dir": schema.StringAttribute{
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
																	Optional:            false,
																	Computed:            true,
																},

																"ip": schema.StringAttribute{
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

													"host_ipc": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_network": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_pid": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_users": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"hostname": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
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
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
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
																	Optional:            false,
																	Computed:            true,
																},

																"command": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"env": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
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

																							"optional": schema.BoolAttribute{
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

																					"field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"api_version": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"field_path": schema.StringAttribute{
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

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"resource": schema.StringAttribute{
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

																					"secret_key_ref": schema.SingleNestedAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"optional": schema.BoolAttribute{
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
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"optional": schema.BoolAttribute{
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

																			"prefix": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"secret_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"optional": schema.BoolAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
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
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"value": schema.StringAttribute{
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

																						"path": schema.StringAttribute{
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

																						"scheme": schema.StringAttribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"http_get": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"http_headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"value": schema.StringAttribute{
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

																						"path": schema.StringAttribute{
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

																						"scheme": schema.StringAttribute{
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

																				"tcp_socket": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"host": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"ports": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"container_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"host_ip": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"host_port": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"protocol": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"resize_policy": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"resource_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"restart_policy": schema.StringAttribute{
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

																		"limits": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

																"restart_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"security_context": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"allow_privilege_escalation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
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
																					Optional:            false,
																					Computed:            true,
																				},

																				"drop": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"proc_mount": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only_root_filesystem": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_group": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_non_root": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"run_as_user": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"se_linux_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"level": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"role": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"user": schema.StringAttribute{
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

																		"seccomp_profile": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"localhost_profile": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
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

																		"windows_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"gmsa_credential_spec": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"gmsa_credential_spec_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"host_process": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"run_as_user_name": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"failure_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"grpc": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"port": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"service": schema.StringAttribute{
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

																		"http_get": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.StringAttribute{
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

																				"path": schema.StringAttribute{
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

																				"scheme": schema.StringAttribute{
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

																		"initial_delay_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"success_threshold": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tcp_socket": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
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
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"termination_grace_period_seconds": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"timeout_seconds": schema.Int64Attribute{
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

																"stdin": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"stdin_once": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"termination_message_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"termination_message_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tty": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"volume_devices": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"device_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"volume_mounts": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mount_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"mount_propagation": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																			"read_only": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"sub_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"sub_path_expr": schema.StringAttribute{
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

																"working_dir": schema.StringAttribute{
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

													"node_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"node_selector": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"os": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
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

													"overhead": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"preemption_policy": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"priority": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"priority_class_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"readiness_gates": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"condition_type": schema.StringAttribute{
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

													"restart_policy": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"runtime_class_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"scheduler_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"security_context": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"fs_group": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"fs_group_change_policy": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"run_as_group": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"run_as_non_root": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"run_as_user": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"se_linux_options": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"role": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"user": schema.StringAttribute{
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

															"seccomp_profile": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"localhost_profile": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
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

															"supplemental_groups": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"sysctls": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.StringAttribute{
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

															"windows_options": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"gmsa_credential_spec": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"gmsa_credential_spec_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"host_process": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"run_as_user_name": schema.StringAttribute{
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

													"service_account": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"service_account_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"set_hostname_as_fqdn": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"share_process_namespace": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"subdomain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"termination_grace_period_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
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
																	Optional:            false,
																	Computed:            true,
																},

																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"operator": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"toleration_seconds": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
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
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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

																"match_label_keys": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_skew": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"min_domains": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"node_affinity_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"node_taints_policy": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"when_unsatisfiable": schema.StringAttribute{
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
																			Optional:            false,
																			Computed:            true,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_id": schema.StringAttribute{
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

																"azure_disk": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"caching_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"disk_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"disk_uri": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
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

																"azure_file": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"read_only": schema.BoolAttribute{
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

																		"share_name": schema.StringAttribute{
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

																"cephfs": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"monitors": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_file": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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

																		"user": schema.StringAttribute{
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

																"cinder": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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

																		"volume_id": schema.StringAttribute{
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

																"config_map": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.StringAttribute{
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

																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"optional": schema.BoolAttribute{
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

																"csi": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"node_publish_secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_attributes": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
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
																								Optional:            false,
																								Computed:            true,
																							},

																							"field_path": schema.StringAttribute{
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

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"resource": schema.StringAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"medium": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"size_limit": schema.StringAttribute{
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
																					Optional:            false,
																					Computed:            true,
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
																							Optional:            false,
																							Computed:            true,
																						},

																						"data_source": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"api_group": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"kind": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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

																						"data_source_ref": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"api_group": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"kind": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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

																								"namespace": schema.StringAttribute{
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

																								"limits": schema.MapAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"requests": schema.MapAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"operator": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"values": schema.ListAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																									Description:         "",
																									MarkdownDescription: "",
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
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"volume_mode": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"volume_name": schema.StringAttribute{
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
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"fc": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"target_ww_ns": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"wwids": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"options": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"flocker": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"dataset_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"dataset_uuid": schema.StringAttribute{
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

																"gce_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"pd_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
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

																"git_repo": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"directory": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"repository": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"revision": schema.StringAttribute{
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

																"glusterfs": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"endpoints": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
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

																"host_path": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
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

																"iscsi": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"chap_auth_discovery": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"chap_auth_session": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"initiator_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"iqn": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"iscsi_interface": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"portals": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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

																		"target_portal": schema.StringAttribute{
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

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"nfs": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"server": schema.StringAttribute{
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

																"persistent_volume_claim": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claim_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
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

																"photon_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"pd_id": schema.StringAttribute{
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

																"portworx_volume": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_id": schema.StringAttribute{
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

																"projected": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"sources": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
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
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"path": schema.StringAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"optional": schema.BoolAttribute{
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
																													Optional:            false,
																													Computed:            true,
																												},

																												"field_path": schema.StringAttribute{
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

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"path": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"resource_field_ref": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"container_name": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"divisor": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"resource": schema.StringAttribute{
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
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"mode": schema.Int64Attribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"path": schema.StringAttribute{
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

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"optional": schema.BoolAttribute{
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

																					"service_account_token": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"audience": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"expiration_seconds": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"path": schema.StringAttribute{
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
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"registry": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"tenant": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"user": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume": schema.StringAttribute{
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

																"rbd": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"image": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"keyring": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"monitors": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"pool": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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

																		"user": schema.StringAttribute{
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

																"scale_io": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"gateway": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"protection_domain": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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

																		"ssl_enabled": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"storage_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"storage_pool": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"system": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_name": schema.StringAttribute{
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

																"secret": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"mode": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.StringAttribute{
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

																		"optional": schema.BoolAttribute{
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

																"storageos": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
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

																		"volume_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_namespace": schema.StringAttribute{
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

																"vsphere_volume": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"storage_policy_id": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"storage_policy_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_path": schema.StringAttribute{
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

					"external": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"admin_password": schema.SingleNestedAttribute{
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

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"optional": schema.BoolAttribute{
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

							"admin_user": schema.SingleNestedAttribute{
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

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"optional": schema.BoolAttribute{
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

							"api_key": schema.SingleNestedAttribute{
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

									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"optional": schema.BoolAttribute{
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

							"url": schema.StringAttribute{
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

					"ingress": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"default_backend": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"resource": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"kind": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

											"service": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"number": schema.Int64Attribute{
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ingress_class_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"rules": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"http": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"paths": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"backend": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"resource": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"api_group": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"kind": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																			"service": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"port": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"number": schema.Int64Attribute{
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
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"path": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"path_type": schema.StringAttribute{
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

									"tls": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hosts": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
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

					"jsonnet": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"library_label_selector": schema.SingleNestedAttribute{
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"values": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
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

					"persistent_volume_claim": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"data_source": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"data_source_ref": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"kind": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

											"limits": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
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
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"operator": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"values": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volume_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volume_name": schema.StringAttribute{
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

					"route": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
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
													Description:         "",
													MarkdownDescription: "",
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

												"weight": schema.Int64Attribute{
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

									"host": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"target_port": schema.StringAttribute{
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

									"tls": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"ca_certificate": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"certificate": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"destination_ca_certificate": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"insecure_edge_termination_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"termination": schema.StringAttribute{
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

									"to": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"kind": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

											"weight": schema.Int64Attribute{
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

									"wildcard_policy": schema.StringAttribute{
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

					"service": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"allocate_load_balancer_node_ports": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"cluster_ip": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"cluster_i_ps": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"external_i_ps": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"external_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"external_traffic_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"health_check_node_port": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"internal_traffic_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ip_families": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ip_family_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"load_balancer_class": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"load_balancer_ip": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"load_balancer_source_ranges": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ports": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"app_protocol": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"protocol": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"target_port": schema.StringAttribute{
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

									"publish_not_ready_addresses": schema.BoolAttribute{
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

									"session_affinity": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"session_affinity_config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"client_ip": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"timeout_seconds": schema.Int64Attribute{
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

					"service_account": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"automount_service_account_token": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"field_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"resource_version": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"uid": schema.StringAttribute{
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

func (r *GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *GrafanaIntegreatlyOrgGrafanaV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_grafana_integreatly_org_grafana_v1beta1")

	var data GrafanaIntegreatlyOrgGrafanaV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "grafana.integreatly.org", Version: "v1beta1", Resource: "grafanas"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
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

	var readResponse GrafanaIntegreatlyOrgGrafanaV1Beta1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("grafana.integreatly.org/v1beta1")
	data.Kind = pointer.String("Grafana")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
