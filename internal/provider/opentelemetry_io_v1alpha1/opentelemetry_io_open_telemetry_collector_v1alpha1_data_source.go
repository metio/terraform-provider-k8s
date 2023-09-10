/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package opentelemetry_io_v1alpha1

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
	_ datasource.DataSource              = &OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource{}
)

func NewOpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource() datasource.DataSource {
	return &OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource{}
}

type OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSourceData struct {
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
		AdditionalContainers *[]struct {
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
		} `tfsdk:"additional_containers" json:"additionalContainers,omitempty"`
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
		Args       *map[string]string `tfsdk:"args" json:"args,omitempty"`
		Autoscaler *struct {
			Behavior *struct {
				ScaleDown *struct {
					Policies *[]struct {
						PeriodSeconds *int64  `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
						Type          *string `tfsdk:"type" json:"type,omitempty"`
						Value         *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"policies" json:"policies,omitempty"`
					SelectPolicy               *string `tfsdk:"select_policy" json:"selectPolicy,omitempty"`
					StabilizationWindowSeconds *int64  `tfsdk:"stabilization_window_seconds" json:"stabilizationWindowSeconds,omitempty"`
				} `tfsdk:"scale_down" json:"scaleDown,omitempty"`
				ScaleUp *struct {
					Policies *[]struct {
						PeriodSeconds *int64  `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
						Type          *string `tfsdk:"type" json:"type,omitempty"`
						Value         *int64  `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"policies" json:"policies,omitempty"`
					SelectPolicy               *string `tfsdk:"select_policy" json:"selectPolicy,omitempty"`
					StabilizationWindowSeconds *int64  `tfsdk:"stabilization_window_seconds" json:"stabilizationWindowSeconds,omitempty"`
				} `tfsdk:"scale_up" json:"scaleUp,omitempty"`
			} `tfsdk:"behavior" json:"behavior,omitempty"`
			MaxReplicas *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
			Metrics     *[]struct {
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
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			MinReplicas             *int64 `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
			TargetCPUUtilization    *int64 `tfsdk:"target_cpu_utilization" json:"targetCPUUtilization,omitempty"`
			TargetMemoryUtilization *int64 `tfsdk:"target_memory_utilization" json:"targetMemoryUtilization,omitempty"`
		} `tfsdk:"autoscaler" json:"autoscaler,omitempty"`
		Config *string `tfsdk:"config" json:"config,omitempty"`
		Env    *[]struct {
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
		HostNetwork     *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
		Image           *string `tfsdk:"image" json:"image,omitempty"`
		ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		Ingress         *struct {
			Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Hostname         *string            `tfsdk:"hostname" json:"hostname,omitempty"`
			IngressClassName *string            `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
			Route            *struct {
				Termination *string `tfsdk:"termination" json:"termination,omitempty"`
			} `tfsdk:"route" json:"route,omitempty"`
			RuleType *string `tfsdk:"rule_type" json:"ruleType,omitempty"`
			Tls      *[]struct {
				Hosts      *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
				SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
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
		Lifecycle *struct {
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
			FailureThreshold              *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			InitialDelaySeconds           *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds                 *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold              *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		ManagementState *string            `tfsdk:"management_state" json:"managementState,omitempty"`
		MaxReplicas     *int64             `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
		MinReplicas     *int64             `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		Mode            *string            `tfsdk:"mode" json:"mode,omitempty"`
		NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Observability   *struct {
			Metrics *struct {
				EnableMetrics *bool `tfsdk:"enable_metrics" json:"enableMetrics,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
		} `tfsdk:"observability" json:"observability,omitempty"`
		PodAnnotations     *map[string]string `tfsdk:"pod_annotations" json:"podAnnotations,omitempty"`
		PodSecurityContext *struct {
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
		} `tfsdk:"pod_security_context" json:"podSecurityContext,omitempty"`
		Ports *[]struct {
			AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
			Port        *int64  `tfsdk:"port" json:"port,omitempty"`
			Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
			TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
		} `tfsdk:"ports" json:"ports,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		Replicas          *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources         *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
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
		ServiceAccount  *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
		TargetAllocator *struct {
			AllocationStrategy *string `tfsdk:"allocation_strategy" json:"allocationStrategy,omitempty"`
			Enabled            *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Env                *[]struct {
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
			FilterStrategy *string            `tfsdk:"filter_strategy" json:"filterStrategy,omitempty"`
			Image          *string            `tfsdk:"image" json:"image,omitempty"`
			NodeSelector   *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PrometheusCR   *struct {
				Enabled                *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				PodMonitorSelector     *map[string]string `tfsdk:"pod_monitor_selector" json:"podMonitorSelector,omitempty"`
				ScrapeInterval         *string            `tfsdk:"scrape_interval" json:"scrapeInterval,omitempty"`
				ServiceMonitorSelector *map[string]string `tfsdk:"service_monitor_selector" json:"serviceMonitorSelector,omitempty"`
			} `tfsdk:"prometheus_cr" json:"prometheusCR,omitempty"`
			Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			ServiceAccount            *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
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
		} `tfsdk:"target_allocator" json:"targetAllocator,omitempty"`
		TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
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
		UpgradeStrategy      *string `tfsdk:"upgrade_strategy" json:"upgradeStrategy,omitempty"`
		VolumeClaimTemplates *[]struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Metadata   *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
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
			Status *struct {
				AccessModes               *[]string          `tfsdk:"access_modes" json:"accessModes,omitempty"`
				AllocatedResourceStatuses *map[string]string `tfsdk:"allocated_resource_statuses" json:"allocatedResourceStatuses,omitempty"`
				AllocatedResources        *map[string]string `tfsdk:"allocated_resources" json:"allocatedResources,omitempty"`
				Capacity                  *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
				Conditions                *[]struct {
					LastProbeTime      *string `tfsdk:"last_probe_time" json:"lastProbeTime,omitempty"`
					LastTransitionTime *string `tfsdk:"last_transition_time" json:"lastTransitionTime,omitempty"`
					Message            *string `tfsdk:"message" json:"message,omitempty"`
					Reason             *string `tfsdk:"reason" json:"reason,omitempty"`
					Status             *string `tfsdk:"status" json:"status,omitempty"`
					Type               *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"conditions" json:"conditions,omitempty"`
				Phase *string `tfsdk:"phase" json:"phase,omitempty"`
			} `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		VolumeMounts *[]struct {
			MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
		} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
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
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Name        *string            `tfsdk:"name" json:"name,omitempty"`
						Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec *struct {
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
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_opentelemetry_io_open_telemetry_collector_v1alpha1"
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OpenTelemetryCollector is the Schema for the opentelemetrycollectors API.",
		MarkdownDescription: "OpenTelemetryCollector is the Schema for the opentelemetrycollectors API.",
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
				Description:         "OpenTelemetryCollectorSpec defines the desired state of OpenTelemetryCollector.",
				MarkdownDescription: "OpenTelemetryCollectorSpec defines the desired state of OpenTelemetryCollector.",
				Attributes: map[string]schema.Attribute{
					"additional_containers": schema.ListNestedAttribute{
						Description:         "AdditionalContainers allows injecting additional containers into the Collector's pod definition.",
						MarkdownDescription: "AdditionalContainers allows injecting additional containers into the Collector's pod definition.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.ListAttribute{
									Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
									MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"command": schema.ListAttribute{
									Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
									MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
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
												Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
												MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
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
														Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
														MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
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
														Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
														MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
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
									Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER.",
									MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER.",
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
									Description:         "Container image name. More info: https://kubernetes.",
									MarkdownDescription: "Container image name. More info: https://kubernetes.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"image_pull_policy": schema.StringAttribute{
									Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.",
									MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"lifecycle": schema.SingleNestedAttribute{
									Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
									MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
									Attributes: map[string]schema.Attribute{
										"post_start": schema.SingleNestedAttribute{
											Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy.",
											MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy.",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
																		Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
													Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
													MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
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
											Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc.",
											MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc.",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
																		Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
													Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
													MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
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
													Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
													MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
											Description:         "GRPC specifies an action involving a GRPC port.",
											MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"service": schema.StringAttribute{
													Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
													MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
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
																Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
											Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
											MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
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
									Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.",
									MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.",
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
									Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.",
									MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.",
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "Exec specifies the action to take.",
											MarkdownDescription: "Exec specifies the action to take.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
													MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
											Description:         "GRPC specifies an action involving a GRPC port.",
											MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"service": schema.StringAttribute{
													Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
													MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
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
																Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
											Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
											MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
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

								"resize_policy": schema.ListNestedAttribute{
									Description:         "Resources resize policy for the container.",
									MarkdownDescription: "Resources resize policy for the container.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"resource_name": schema.StringAttribute{
												Description:         "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
												MarkdownDescription: "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"restart_policy": schema.StringAttribute{
												Description:         "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
												MarkdownDescription: "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
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
									Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
									MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
									Attributes: map[string]schema.Attribute{
										"claims": schema.ListNestedAttribute{
											Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
											MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
														MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
											Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"requests": schema.MapAttribute{
											Description:         "Requests describes the minimum amount of compute resources required.",
											MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
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
									Description:         "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'.",
									MarkdownDescription: "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"security_context": schema.SingleNestedAttribute{
									Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
									MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
									Attributes: map[string]schema.Attribute{
										"allow_privilege_escalation": schema.BoolAttribute{
											Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process.",
											MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process.",
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
											Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths.",
											MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths.",
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
											Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.",
											MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"run_as_non_root": schema.BoolAttribute{
											Description:         "Indicates that the container must run as a non-root user.",
											MarkdownDescription: "Indicates that the container must run as a non-root user.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"run_as_user": schema.Int64Attribute{
											Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
											MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"se_linux_options": schema.SingleNestedAttribute{
											Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.",
											MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.",
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
											Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options.",
											MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options.",
											Attributes: map[string]schema.Attribute{
												"localhost_profile": schema.StringAttribute{
													Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
													MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"type": schema.StringAttribute{
													Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
													MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
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
											Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used.",
											MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used.",
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
													Description:         "HostProcess determines if a container should be run as a 'Host Process' container.",
													MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"run_as_user_name": schema.StringAttribute{
													Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
													MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
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
									Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully.",
									MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully.",
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "Exec specifies the action to take.",
											MarkdownDescription: "Exec specifies the action to take.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
													MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
											Description:         "GRPC specifies an action involving a GRPC port.",
											MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"service": schema.StringAttribute{
													Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
													MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
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
																Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
											Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
											MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
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
									Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions.",
									MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"termination_message_path": schema.StringAttribute{
									Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.",
									MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"termination_message_policy": schema.StringAttribute{
									Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure.",
									MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure.",
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
												Description:         "Expanded path within the volume from which the container's volume should be mounted.",
												MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.",
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

					"affinity": schema.SingleNestedAttribute{
						Description:         "If specified, indicates the pod's scheduling constraints",
						MarkdownDescription: "If specified, indicates the pod's scheduling constraints",
						Attributes: map[string]schema.Attribute{
							"node_affinity": schema.SingleNestedAttribute{
								Description:         "Describes node affinity scheduling rules for the pod.",
								MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions.",
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
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node.",
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
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions.",
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
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																	Description:         "matchLabels is a map of {key,value} pairs.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
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
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																	Description:         "matchLabels is a map of {key,value} pairs.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
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
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node.",
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
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
															Description:         "matchLabels is a map of {key,value} pairs.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
													Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
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
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
															Description:         "matchLabels is a map of {key,value} pairs.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
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
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions.",
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
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																	Description:         "matchLabels is a map of {key,value} pairs.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
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
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																	Description:         "matchLabels is a map of {key,value} pairs.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
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
										Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node.",
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
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
															Description:         "matchLabels is a map of {key,value} pairs.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
													Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field.",
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
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
															Description:         "matchLabels is a map of {key,value} pairs.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose ",
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

					"args": schema.MapAttribute{
						Description:         "Args is the set of arguments to pass to the OpenTelemetry Collector binary",
						MarkdownDescription: "Args is the set of arguments to pass to the OpenTelemetry Collector binary",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"autoscaler": schema.SingleNestedAttribute{
						Description:         "Autoscaler specifies the pod autoscaling configuration to use for the OpenTelemetryCollector workload.",
						MarkdownDescription: "Autoscaler specifies the pod autoscaling configuration to use for the OpenTelemetryCollector workload.",
						Attributes: map[string]schema.Attribute{
							"behavior": schema.SingleNestedAttribute{
								Description:         "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
								MarkdownDescription: "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
								Attributes: map[string]schema.Attribute{
									"scale_down": schema.SingleNestedAttribute{
										Description:         "scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e.",
										MarkdownDescription: "scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e.",
										Attributes: map[string]schema.Attribute{
											"policies": schema.ListNestedAttribute{
												Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
												MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"period_seconds": schema.Int64Attribute{
															Description:         "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
															MarkdownDescription: "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"type": schema.StringAttribute{
															Description:         "type is used to specify the scaling policy.",
															MarkdownDescription: "type is used to specify the scaling policy.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"value": schema.Int64Attribute{
															Description:         "value contains the amount of change which is permitted by the policy. It must be greater than zero",
															MarkdownDescription: "value contains the amount of change which is permitted by the policy. It must be greater than zero",
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

											"select_policy": schema.StringAttribute{
												Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
												MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"stabilization_window_seconds": schema.Int64Attribute{
												Description:         "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down.",
												MarkdownDescription: "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"scale_up": schema.SingleNestedAttribute{
										Description:         "scaleUp is scaling policy for scaling Up.",
										MarkdownDescription: "scaleUp is scaling policy for scaling Up.",
										Attributes: map[string]schema.Attribute{
											"policies": schema.ListNestedAttribute{
												Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
												MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"period_seconds": schema.Int64Attribute{
															Description:         "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
															MarkdownDescription: "periodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"type": schema.StringAttribute{
															Description:         "type is used to specify the scaling policy.",
															MarkdownDescription: "type is used to specify the scaling policy.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"value": schema.Int64Attribute{
															Description:         "value contains the amount of change which is permitted by the policy. It must be greater than zero",
															MarkdownDescription: "value contains the amount of change which is permitted by the policy. It must be greater than zero",
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

											"select_policy": schema.StringAttribute{
												Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
												MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"stabilization_window_seconds": schema.Int64Attribute{
												Description:         "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down.",
												MarkdownDescription: "stabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down.",
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

							"max_replicas": schema.Int64Attribute{
								Description:         "MaxReplicas sets an upper bound to the autoscaling feature. If MaxReplicas is set autoscaling is enabled.",
								MarkdownDescription: "MaxReplicas sets an upper bound to the autoscaling feature. If MaxReplicas is set autoscaling is enabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metrics": schema.ListNestedAttribute{
								Description:         "Metrics is meant to provide a customizable way to configure HPA metrics. currently the only supported custom metrics is type=Pod.",
								MarkdownDescription: "Metrics is meant to provide a customizable way to configure HPA metrics. currently the only supported custom metrics is type=Pod.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"pods": schema.SingleNestedAttribute{
											Description:         "PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second).",
											MarkdownDescription: "PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second).",
											Attributes: map[string]schema.Attribute{
												"metric": schema.SingleNestedAttribute{
													Description:         "metric identifies the target metric by name and selector",
													MarkdownDescription: "metric identifies the target metric by name and selector",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "name is the name of the given metric",
															MarkdownDescription: "name is the name of the given metric",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"selector": schema.SingleNestedAttribute{
															Description:         "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scopi",
															MarkdownDescription: "selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scopi",
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
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																	Description:         "matchLabels is a map of {key,value} pairs.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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

												"target": schema.SingleNestedAttribute{
													Description:         "target specifies the target value for the given metric",
													MarkdownDescription: "target specifies the target value for the given metric",
													Attributes: map[string]schema.Attribute{
														"average_utilization": schema.Int64Attribute{
															Description:         "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.",
															MarkdownDescription: "averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"average_value": schema.StringAttribute{
															Description:         "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
															MarkdownDescription: "averageValue is the target value of the average of the metric across all relevant pods (as a quantity)",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"type": schema.StringAttribute{
															Description:         "type represents whether the metric type is Utilization, Value, or AverageValue",
															MarkdownDescription: "type represents whether the metric type is Utilization, Value, or AverageValue",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"value": schema.StringAttribute{
															Description:         "value is the target value of the metric (as a quantity).",
															MarkdownDescription: "value is the target value of the metric (as a quantity).",
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
											Description:         "MetricSourceType indicates the type of metric.",
											MarkdownDescription: "MetricSourceType indicates the type of metric.",
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

							"min_replicas": schema.Int64Attribute{
								Description:         "MinReplicas sets a lower bound to the autoscaling feature.  Set this if your are using autoscaling. It must be at least 1",
								MarkdownDescription: "MinReplicas sets a lower bound to the autoscaling feature.  Set this if your are using autoscaling. It must be at least 1",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"target_cpu_utilization": schema.Int64Attribute{
								Description:         "TargetCPUUtilization sets the target average CPU used across all replicas. If average CPU exceeds this value, the HPA will scale up. Defaults to 90 percent.",
								MarkdownDescription: "TargetCPUUtilization sets the target average CPU used across all replicas. If average CPU exceeds this value, the HPA will scale up. Defaults to 90 percent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"target_memory_utilization": schema.Int64Attribute{
								Description:         "TargetMemoryUtilization sets the target average memory utilization across all replicas",
								MarkdownDescription: "TargetMemoryUtilization sets the target average memory utilization across all replicas",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"config": schema.StringAttribute{
						Description:         "Config is the raw JSON to be used as the collector's configuration. Refer to the OpenTelemetry Collector documentation for details.",
						MarkdownDescription: "Config is the raw JSON to be used as the collector's configuration. Refer to the OpenTelemetry Collector documentation for details.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"env": schema.ListNestedAttribute{
						Description:         "ENV vars to set on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",
						MarkdownDescription: "ENV vars to set on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",
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
									Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
									MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
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
											Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
											MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
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
											Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
											MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
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
						Description:         "List of sources to populate environment variables on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",
						MarkdownDescription: "List of sources to populate environment variables on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",
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

					"host_network": schema.BoolAttribute{
						Description:         "HostNetwork indicates if the pod should run in the host networking namespace.",
						MarkdownDescription: "HostNetwork indicates if the pod should run in the host networking namespace.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image": schema.StringAttribute{
						Description:         "Image indicates the container image to use for the OpenTelemetry Collector.",
						MarkdownDescription: "Image indicates the container image to use for the OpenTelemetry Collector.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "ImagePullPolicy indicates the pull policy to be used for retrieving the container image (Always, Never, IfNotPresent)",
						MarkdownDescription: "ImagePullPolicy indicates the pull policy to be used for retrieving the container image (Always, Never, IfNotPresent)",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ingress": schema.SingleNestedAttribute{
						Description:         "Ingress is used to specify how OpenTelemetry Collector is exposed. This functionality is only available if one of the valid modes is set. Valid modes are: deployment, daemonset and statefulset.",
						MarkdownDescription: "Ingress is used to specify how OpenTelemetry Collector is exposed. This functionality is only available if one of the valid modes is set. Valid modes are: deployment, daemonset and statefulset.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to add to ingress. e.g. 'cert-manager.io/cluster-issuer: 'letsencrypt''",
								MarkdownDescription: "Annotations to add to ingress. e.g. 'cert-manager.io/cluster-issuer: 'letsencrypt''",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"hostname": schema.StringAttribute{
								Description:         "Hostname by which the ingress proxy can be reached.",
								MarkdownDescription: "Hostname by which the ingress proxy can be reached.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ingress_class_name": schema.StringAttribute{
								Description:         "IngressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource.",
								MarkdownDescription: "IngressClassName is the name of an IngressClass cluster resource. Ingress controller implementations use this field to know whether they should be serving this Ingress resource.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"route": schema.SingleNestedAttribute{
								Description:         "Route is an OpenShift specific section that is only considered when type 'route' is used.",
								MarkdownDescription: "Route is an OpenShift specific section that is only considered when type 'route' is used.",
								Attributes: map[string]schema.Attribute{
									"termination": schema.StringAttribute{
										Description:         "Termination indicates termination type. By default 'edge' is used.",
										MarkdownDescription: "Termination indicates termination type. By default 'edge' is used.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"rule_type": schema.StringAttribute{
								Description:         "RuleType defines how Ingress exposes collector receivers. IngressRuleTypePath ('path') exposes each receiver port on a unique path on single domain defined in Hostname.",
								MarkdownDescription: "RuleType defines how Ingress exposes collector receivers. IngressRuleTypePath ('path') exposes each receiver port on a unique path on single domain defined in Hostname.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tls": schema.ListNestedAttribute{
								Description:         "TLS configuration.",
								MarkdownDescription: "TLS configuration.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hosts": schema.ListAttribute{
											Description:         "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret.",
											MarkdownDescription: "hosts is a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_name": schema.StringAttribute{
											Description:         "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone.",
											MarkdownDescription: "secretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone.",
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
								Description:         "Type default value is: '' Supported types are: ingress, route",
								MarkdownDescription: "Type default value is: '' Supported types are: ingress, route",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"init_containers": schema.ListNestedAttribute{
						Description:         "InitContainers allows injecting initContainers to the Collector's pod definition.",
						MarkdownDescription: "InitContainers allows injecting initContainers to the Collector's pod definition.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.ListAttribute{
									Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
									MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"command": schema.ListAttribute{
									Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
									MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.",
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
												Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
												MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
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
														Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
														MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
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
														Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
														MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
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
									Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER.",
									MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER.",
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
									Description:         "Container image name. More info: https://kubernetes.",
									MarkdownDescription: "Container image name. More info: https://kubernetes.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"image_pull_policy": schema.StringAttribute{
									Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.",
									MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"lifecycle": schema.SingleNestedAttribute{
									Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
									MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
									Attributes: map[string]schema.Attribute{
										"post_start": schema.SingleNestedAttribute{
											Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy.",
											MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy.",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
																		Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
													Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
													MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
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
											Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc.",
											MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc.",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
																		Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
													Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
													MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
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
													Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
													MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
											Description:         "GRPC specifies an action involving a GRPC port.",
											MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"service": schema.StringAttribute{
													Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
													MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
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
																Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
											Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
											MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
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
									Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.",
									MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.",
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
									Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.",
									MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.",
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "Exec specifies the action to take.",
											MarkdownDescription: "Exec specifies the action to take.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
													MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
											Description:         "GRPC specifies an action involving a GRPC port.",
											MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"service": schema.StringAttribute{
													Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
													MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
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
																Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
											Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
											MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
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

								"resize_policy": schema.ListNestedAttribute{
									Description:         "Resources resize policy for the container.",
									MarkdownDescription: "Resources resize policy for the container.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"resource_name": schema.StringAttribute{
												Description:         "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
												MarkdownDescription: "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"restart_policy": schema.StringAttribute{
												Description:         "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
												MarkdownDescription: "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
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
									Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
									MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
									Attributes: map[string]schema.Attribute{
										"claims": schema.ListNestedAttribute{
											Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
											MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
														MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
											Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"requests": schema.MapAttribute{
											Description:         "Requests describes the minimum amount of compute resources required.",
											MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
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
									Description:         "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'.",
									MarkdownDescription: "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"security_context": schema.SingleNestedAttribute{
									Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
									MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
									Attributes: map[string]schema.Attribute{
										"allow_privilege_escalation": schema.BoolAttribute{
											Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process.",
											MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process.",
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
											Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths.",
											MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths.",
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
											Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.",
											MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"run_as_non_root": schema.BoolAttribute{
											Description:         "Indicates that the container must run as a non-root user.",
											MarkdownDescription: "Indicates that the container must run as a non-root user.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"run_as_user": schema.Int64Attribute{
											Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
											MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"se_linux_options": schema.SingleNestedAttribute{
											Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.",
											MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.",
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
											Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options.",
											MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options.",
											Attributes: map[string]schema.Attribute{
												"localhost_profile": schema.StringAttribute{
													Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
													MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"type": schema.StringAttribute{
													Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
													MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
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
											Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used.",
											MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used.",
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
													Description:         "HostProcess determines if a container should be run as a 'Host Process' container.",
													MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"run_as_user_name": schema.StringAttribute{
													Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
													MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
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
									Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully.",
									MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully.",
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "Exec specifies the action to take.",
											MarkdownDescription: "Exec specifies the action to take.",
											Attributes: map[string]schema.Attribute{
												"command": schema.ListAttribute{
													Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
													MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
											Description:         "GRPC specifies an action involving a GRPC port.",
											MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"service": schema.StringAttribute{
													Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
													MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).",
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
																Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
											Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
											MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
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
									Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions.",
									MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"termination_message_path": schema.StringAttribute{
									Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.",
									MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"termination_message_policy": schema.StringAttribute{
									Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure.",
									MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure.",
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
												Description:         "Expanded path within the volume from which the container's volume should be mounted.",
												MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.",
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

					"lifecycle": schema.SingleNestedAttribute{
						Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
						MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
						Attributes: map[string]schema.Attribute{
							"post_start": schema.SingleNestedAttribute{
								Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy.",
								MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy.",
								Attributes: map[string]schema.Attribute{
									"exec": schema.SingleNestedAttribute{
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",
										Attributes: map[string]schema.Attribute{
											"command": schema.ListAttribute{
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
															Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
															MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
										Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
										MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
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
								Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc.",
								MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc.",
								Attributes: map[string]schema.Attribute{
									"exec": schema.SingleNestedAttribute{
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",
										Attributes: map[string]schema.Attribute{
											"command": schema.ListAttribute{
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem.",
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
															Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
															MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
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
										Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
										MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility.",
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
						Description:         "Liveness config for the OpenTelemetry Collector except the probe handler which is auto generated from the health extension of the collector.",
						MarkdownDescription: "Liveness config for the OpenTelemetry Collector except the probe handler which is auto generated from the health extension of the collector.",
						Attributes: map[string]schema.Attribute{
							"failure_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "Number of seconds after the container has started before liveness probes are initiated. Defaults to 0 seconds. Minimum value is 0. More info: https://kubernetes.",
								MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. Defaults to 0 seconds. Minimum value is 0. More info: https://kubernetes.",
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

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
								MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.",
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

					"management_state": schema.StringAttribute{
						Description:         "ManagementState defines if the CR should be managed by the operator or not. Default is managed.",
						MarkdownDescription: "ManagementState defines if the CR should be managed by the operator or not. Default is managed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"max_replicas": schema.Int64Attribute{
						Description:         "MaxReplicas sets an upper bound to the autoscaling feature. If MaxReplicas is set autoscaling is enabled. Deprecated: use 'OpenTelemetryCollector.Spec.Autoscaler.MaxReplicas' instead.",
						MarkdownDescription: "MaxReplicas sets an upper bound to the autoscaling feature. If MaxReplicas is set autoscaling is enabled. Deprecated: use 'OpenTelemetryCollector.Spec.Autoscaler.MaxReplicas' instead.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"min_replicas": schema.Int64Attribute{
						Description:         "MinReplicas sets a lower bound to the autoscaling feature.  Set this if you are using autoscaling. It must be at least 1 Deprecated: use 'OpenTelemetryCollector.Spec.Autoscaler.MinReplicas' instead.",
						MarkdownDescription: "MinReplicas sets a lower bound to the autoscaling feature.  Set this if you are using autoscaling. It must be at least 1 Deprecated: use 'OpenTelemetryCollector.Spec.Autoscaler.MinReplicas' instead.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"mode": schema.StringAttribute{
						Description:         "Mode represents how the collector should be deployed (deployment, daemonset, statefulset or sidecar)",
						MarkdownDescription: "Mode represents how the collector should be deployed (deployment, daemonset, statefulset or sidecar)",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"node_selector": schema.MapAttribute{
						Description:         "NodeSelector to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",
						MarkdownDescription: "NodeSelector to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"observability": schema.SingleNestedAttribute{
						Description:         "ObservabilitySpec defines how telemetry data gets handled.",
						MarkdownDescription: "ObservabilitySpec defines how telemetry data gets handled.",
						Attributes: map[string]schema.Attribute{
							"metrics": schema.SingleNestedAttribute{
								Description:         "Metrics defines the metrics configuration for operands.",
								MarkdownDescription: "Metrics defines the metrics configuration for operands.",
								Attributes: map[string]schema.Attribute{
									"enable_metrics": schema.BoolAttribute{
										Description:         "EnableMetrics specifies if ServiceMonitor should be created for the OpenTelemetry Collector and Prometheus Exporters. The operator.observability.",
										MarkdownDescription: "EnableMetrics specifies if ServiceMonitor should be created for the OpenTelemetry Collector and Prometheus Exporters. The operator.observability.",
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

					"pod_annotations": schema.MapAttribute{
						Description:         "PodAnnotations is the set of annotations that will be attached to Collector and Target Allocator pods.",
						MarkdownDescription: "PodAnnotations is the set of annotations that will be attached to Collector and Target Allocator pods.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"pod_security_context": schema.SingleNestedAttribute{
						Description:         "PodSecurityContext configures the pod security context for the opentelemetry-collector pod, when running as a deployment, daemonset, or statefulset.",
						MarkdownDescription: "PodSecurityContext configures the pod security context for the opentelemetry-collector pod, when running as a deployment, daemonset, or statefulset.",
						Attributes: map[string]schema.Attribute{
							"fs_group": schema.Int64Attribute{
								Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1.",
								MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"fs_group_change_policy": schema.StringAttribute{
								Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod.",
								MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"run_as_group": schema.Int64Attribute{
								Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.",
								MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"run_as_non_root": schema.BoolAttribute{
								Description:         "Indicates that the container must run as a non-root user.",
								MarkdownDescription: "Indicates that the container must run as a non-root user.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"run_as_user": schema.Int64Attribute{
								Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.",
								MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"se_linux_options": schema.SingleNestedAttribute{
								Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.",
								MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.",
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
								Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
										MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
										MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
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
								Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for th",
								MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for th",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sysctls": schema.ListNestedAttribute{
								Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.",
								MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of a property to set",
											MarkdownDescription: "Name of a property to set",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "Value of a property to set",
											MarkdownDescription: "Value of a property to set",
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
								Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used.",
								MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used.",
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
										Description:         "HostProcess determines if a container should be run as a 'Host Process' container.",
										MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"run_as_user_name": schema.StringAttribute{
										Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
										MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
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

					"ports": schema.ListNestedAttribute{
						Description:         "Ports allows a set of ports to be exposed by the underlying v1.Service. By default, the operator will attempt to infer the required ports by parsing the .Spec.",
						MarkdownDescription: "Ports allows a set of ports to be exposed by the underlying v1.Service. By default, the operator will attempt to infer the required ports by parsing the .Spec.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"app_protocol": schema.StringAttribute{
									Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax.",
									MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names.",
									MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"node_port": schema.Int64Attribute{
									Description:         "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system.",
									MarkdownDescription: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"port": schema.Int64Attribute{
									Description:         "The port that will be exposed by this service.",
									MarkdownDescription: "The port that will be exposed by this service.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"protocol": schema.StringAttribute{
									Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
									MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target_port": schema.StringAttribute{
									Description:         "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
									MarkdownDescription: "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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

					"priority_class_name": schema.StringAttribute{
						Description:         "If specified, indicates the pod's priority. If not specified, the pod priority will be default or zero if there is no default.",
						MarkdownDescription: "If specified, indicates the pod's priority. If not specified, the pod priority will be default or zero if there is no default.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Replicas is the number of pod instances for the underlying OpenTelemetry Collector. Set this if your are not using autoscaling",
						MarkdownDescription: "Replicas is the number of pod instances for the underlying OpenTelemetry Collector. Set this if your are not using autoscaling",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources to set on the OpenTelemetry Collector pods.",
						MarkdownDescription: "Resources to set on the OpenTelemetry Collector pods.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
								Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required.",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
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
						Description:         "SecurityContext configures the container security context for the opentelemetry-collector container.",
						MarkdownDescription: "SecurityContext configures the container security context for the opentelemetry-collector container.",
						Attributes: map[string]schema.Attribute{
							"allow_privilege_escalation": schema.BoolAttribute{
								Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process.",
								MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process.",
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
								Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths.",
								MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths.",
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
								Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.",
								MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"run_as_non_root": schema.BoolAttribute{
								Description:         "Indicates that the container must run as a non-root user.",
								MarkdownDescription: "Indicates that the container must run as a non-root user.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"run_as_user": schema.Int64Attribute{
								Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
								MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"se_linux_options": schema.SingleNestedAttribute{
								Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.",
								MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.",
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
								Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options.",
								MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
										MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
										MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used.",
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
								Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used.",
								MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used.",
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
										Description:         "HostProcess determines if a container should be run as a 'Host Process' container.",
										MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"run_as_user_name": schema.StringAttribute{
										Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
										MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext.",
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
						Description:         "ServiceAccount indicates the name of an existing service account to use with this instance. When set, the operator will not automatically create a ServiceAccount for the collector.",
						MarkdownDescription: "ServiceAccount indicates the name of an existing service account to use with this instance. When set, the operator will not automatically create a ServiceAccount for the collector.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"target_allocator": schema.SingleNestedAttribute{
						Description:         "TargetAllocator indicates a value which determines whether to spawn a target allocation resource or not.",
						MarkdownDescription: "TargetAllocator indicates a value which determines whether to spawn a target allocation resource or not.",
						Attributes: map[string]schema.Attribute{
							"allocation_strategy": schema.StringAttribute{
								Description:         "AllocationStrategy determines which strategy the target allocator should use for allocation. The current options are least-weighted and consistent-hashing. The default option is least-weighted",
								MarkdownDescription: "AllocationStrategy determines which strategy the target allocator should use for allocation. The current options are least-weighted and consistent-hashing. The default option is least-weighted",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enabled indicates whether to use a target allocation mechanism for Prometheus targets or not.",
								MarkdownDescription: "Enabled indicates whether to use a target allocation mechanism for Prometheus targets or not.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env": schema.ListNestedAttribute{
								Description:         "ENV vars to set on the OpenTelemetry TargetAllocator's Pods. These can then in certain cases be consumed in the config file for the TargetAllocator.",
								MarkdownDescription: "ENV vars to set on the OpenTelemetry TargetAllocator's Pods. These can then in certain cases be consumed in the config file for the TargetAllocator.",
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
											Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
											MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables.",
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
													Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
													MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.",
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
													Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
													MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.",
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

							"filter_strategy": schema.StringAttribute{
								Description:         "FilterStrategy determines how to filter targets before allocating them among the collectors. The only current option is relabel-config (drops targets based on prom relabel_config).",
								MarkdownDescription: "FilterStrategy determines how to filter targets before allocating them among the collectors. The only current option is relabel-config (drops targets based on prom relabel_config).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image": schema.StringAttribute{
								Description:         "Image indicates the container image to use for the OpenTelemetry TargetAllocator.",
								MarkdownDescription: "Image indicates the container image to use for the OpenTelemetry TargetAllocator.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector to schedule OpenTelemetry TargetAllocator pods.",
								MarkdownDescription: "NodeSelector to schedule OpenTelemetry TargetAllocator pods.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"prometheus_cr": schema.SingleNestedAttribute{
								Description:         "PrometheusCR defines the configuration for the retrieval of PrometheusOperator CRDs ( servicemonitor.monitoring.coreos.com/v1 and podmonitor.monitoring.coreos.com/v1 )  retrieval.",
								MarkdownDescription: "PrometheusCR defines the configuration for the retrieval of PrometheusOperator CRDs ( servicemonitor.monitoring.coreos.com/v1 and podmonitor.monitoring.coreos.com/v1 )  retrieval.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled indicates whether to use a PrometheusOperator custom resources as targets or not.",
										MarkdownDescription: "Enabled indicates whether to use a PrometheusOperator custom resources as targets or not.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_monitor_selector": schema.MapAttribute{
										Description:         "PodMonitors to be selected for target discovery. This is a map of {key,value} pairs. Each {key,value} in the map is going to exactly match a label in a PodMonitor's meta labels.",
										MarkdownDescription: "PodMonitors to be selected for target discovery. This is a map of {key,value} pairs. Each {key,value} in the map is going to exactly match a label in a PodMonitor's meta labels.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"scrape_interval": schema.StringAttribute{
										Description:         "Interval between consecutive scrapes. Equivalent to the same setting on the Prometheus CRD.  Default: '30s'",
										MarkdownDescription: "Interval between consecutive scrapes. Equivalent to the same setting on the Prometheus CRD.  Default: '30s'",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"service_monitor_selector": schema.MapAttribute{
										Description:         "ServiceMonitors to be selected for target discovery. This is a map of {key,value} pairs. Each {key,value} in the map is going to exactly match a label in a ServiceMonitor's meta labels.",
										MarkdownDescription: "ServiceMonitors to be selected for target discovery. This is a map of {key,value} pairs. Each {key,value} in the map is going to exactly match a label in a ServiceMonitor's meta labels.",
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

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the number of pod instances for the underlying TargetAllocator. This should only be set to a value other than 1 if a strategy that allows for high availability is chosen.",
								MarkdownDescription: "Replicas is the number of pod instances for the underlying TargetAllocator. This should only be set to a value other than 1 if a strategy that allows for high availability is chosen.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources to set on the OpenTelemetryTargetAllocator containers.",
								MarkdownDescription: "Resources to set on the OpenTelemetryTargetAllocator containers.",
								Attributes: map[string]schema.Attribute{
									"claims": schema.ListNestedAttribute{
										Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
										MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
													MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required.",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
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

							"service_account": schema.StringAttribute{
								Description:         "ServiceAccount indicates the name of an existing service account to use with this instance. When set, the operator will not automatically create a ServiceAccount for the TargetAllocator.",
								MarkdownDescription: "ServiceAccount indicates the name of an existing service account to use with this instance. When set, the operator will not automatically create a ServiceAccount for the TargetAllocator.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"topology_spread_constraints": schema.ListNestedAttribute{
								Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined top",
								MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined top",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"label_selector": schema.SingleNestedAttribute{
											Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
											MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
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
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
													Description:         "matchLabels is a map of {key,value} pairs.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
											Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated.",
											MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"max_skew": schema.Int64Attribute{
											Description:         "MaxSkew describes the degree to which pods may be unevenly distributed.",
											MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"min_domains": schema.Int64Attribute{
											Description:         "MinDomains indicates a minimum number of eligible domains.",
											MarkdownDescription: "MinDomains indicates a minimum number of eligible domains.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"node_affinity_policy": schema.StringAttribute{
											Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew.",
											MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"node_taints_policy": schema.StringAttribute{
											Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew.",
											MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"topology_key": schema.StringAttribute{
											Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology.",
											MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"when_unsatisfiable": schema.StringAttribute{
											Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it.",
											MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it.",
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

					"termination_grace_period_seconds": schema.Int64Attribute{
						Description:         "Duration in seconds the pod needs to terminate gracefully upon probe failure.",
						MarkdownDescription: "Duration in seconds the pod needs to terminate gracefully upon probe failure.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Toleration to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",
						MarkdownDescription: "Toleration to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",
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
									Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal.",
									MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint.",
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

					"topology_spread_constraints": schema.ListNestedAttribute{
						Description:         "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined top",
						MarkdownDescription: "TopologySpreadConstraints embedded kubernetes pod configuration option, controls how pods are spread across your cluster among failure-domains such as regions, zones, nodes, and other user-defined top",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"label_selector": schema.SingleNestedAttribute{
									Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
									MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
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
														Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
											Description:         "matchLabels is a map of {key,value} pairs.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
									Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated.",
									MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"max_skew": schema.Int64Attribute{
									Description:         "MaxSkew describes the degree to which pods may be unevenly distributed.",
									MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"min_domains": schema.Int64Attribute{
									Description:         "MinDomains indicates a minimum number of eligible domains.",
									MarkdownDescription: "MinDomains indicates a minimum number of eligible domains.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"node_affinity_policy": schema.StringAttribute{
									Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew.",
									MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"node_taints_policy": schema.StringAttribute{
									Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew.",
									MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"topology_key": schema.StringAttribute{
									Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology.",
									MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"when_unsatisfiable": schema.StringAttribute{
									Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it.",
									MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it.",
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

					"upgrade_strategy": schema.StringAttribute{
						Description:         "UpgradeStrategy represents how the operator will handle upgrades to the CR when a newer version of the operator is deployed",
						MarkdownDescription: "UpgradeStrategy represents how the operator will handle upgrades to the CR when a newer version of the operator is deployed",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"volume_claim_templates": schema.ListNestedAttribute{
						Description:         "VolumeClaimTemplates will provide stable storage using PersistentVolumes. Only available when the mode=statefulset.",
						MarkdownDescription: "VolumeClaimTemplates will provide stable storage using PersistentVolumes. Only available when the mode=statefulset.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.",
									MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.",
									MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"metadata": schema.SingleNestedAttribute{
									Description:         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
									MarkdownDescription: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"finalizers": schema.ListAttribute{
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

								"spec": schema.SingleNestedAttribute{
									Description:         "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
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
											Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.",
											MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.",
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
											Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired.",
											MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired.",
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

												"namespace": schema.StringAttribute{
													Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.",
													MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.",
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
											Description:         "resources represents the minimum resources the volume should have.",
											MarkdownDescription: "resources represents the minimum resources the volume should have.",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
													MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
													Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required.",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
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
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
													Description:         "matchLabels is a map of {key,value} pairs.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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

								"status": schema.SingleNestedAttribute{
									Description:         "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									Attributes: map[string]schema.Attribute{
										"access_modes": schema.ListAttribute{
											Description:         "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"allocated_resource_statuses": schema.MapAttribute{
											Description:         "allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax.",
											MarkdownDescription: "allocatedResourceStatuses stores status of resource being resized for the given PVC. Key names follow standard Kubernetes label syntax.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"allocated_resources": schema.MapAttribute{
											Description:         "allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax.",
											MarkdownDescription: "allocatedResources tracks the resources allocated to a PVC including its capacity. Key names follow standard Kubernetes label syntax.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"capacity": schema.MapAttribute{
											Description:         "capacity represents the actual resources of the underlying volume.",
											MarkdownDescription: "capacity represents the actual resources of the underlying volume.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"conditions": schema.ListNestedAttribute{
											Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.",
											MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"last_probe_time": schema.StringAttribute{
														Description:         "lastProbeTime is the time we probed the condition.",
														MarkdownDescription: "lastProbeTime is the time we probed the condition.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"last_transition_time": schema.StringAttribute{
														Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
														MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"message": schema.StringAttribute{
														Description:         "message is the human-readable message indicating details about last transition.",
														MarkdownDescription: "message is the human-readable message indicating details about last transition.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"reason": schema.StringAttribute{
														Description:         "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition.",
														MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"status": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
														MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
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

										"phase": schema.StringAttribute{
											Description:         "phase represents the current phase of PersistentVolumeClaim.",
											MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",
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

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts represents the mount points to use in the underlying collector deployment(s)",
						MarkdownDescription: "VolumeMounts represents the mount points to use in the underlying collector deployment(s)",
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
									Description:         "Expanded path within the volume from which the container's volume should be mounted.",
									MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.",
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
						Description:         "Volumes represents which volumes to use in the underlying collector deployment(s).",
						MarkdownDescription: "Volumes represents which volumes to use in the underlying collector deployment(s).",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws_elastic_block_store": schema.SingleNestedAttribute{
									Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.",
									MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
											MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"partition": schema.Int64Attribute{
											Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'.",
											MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'.",
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
											Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set).",
											MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set).",
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
											Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
											Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
											MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"items": schema.ListNestedAttribute{
											Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.",
											MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.",
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
														Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
														MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
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
											Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls.",
											MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls.",
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
											Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default.",
											MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default.",
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
														Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
														MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
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
											Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory.",
											MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"size_limit": schema.StringAttribute{
											Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium.",
											MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium.",
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
									Description:         "ephemeral represents a volume that is handled by a cluster storage driver.",
									MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver.",
									Attributes: map[string]schema.Attribute{
										"volume_claim_template": schema.SingleNestedAttribute{
											Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e.",
											MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e.",
											Attributes: map[string]schema.Attribute{
												"metadata": schema.SingleNestedAttribute{
													Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
													MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
													Attributes: map[string]schema.Attribute{
														"annotations": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"finalizers": schema.ListAttribute{
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

												"spec": schema.SingleNestedAttribute{
													Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template.",
													MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template.",
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
															Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.",
															MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.",
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
															Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired.",
															MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired.",
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

																"namespace": schema.StringAttribute{
																	Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.",
																	MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.",
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
															Description:         "resources represents the minimum resources the volume should have.",
															MarkdownDescription: "resources represents the minimum resources the volume should have.",
															Attributes: map[string]schema.Attribute{
																"claims": schema.ListNestedAttribute{
																	Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
																	MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																				MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																	Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"requests": schema.MapAttribute{
																	Description:         "Requests describes the minimum amount of compute resources required.",
																	MarkdownDescription: "Requests describes the minimum amount of compute resources required.",
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
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty.",
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
																	Description:         "matchLabels is a map of {key,value} pairs.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs.",
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
											Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
											Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified.",
											MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified.",
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
									Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.",
									MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
											MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"partition": schema.Int64Attribute{
											Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'.",
											MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'.",
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
									Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated.",
									MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated.",
									Attributes: map[string]schema.Attribute{
										"directory": schema.StringAttribute{
											Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.",
											MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.",
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
									Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container.",
									MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container.",
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
											Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
											MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"initiator_name": schema.StringAttribute{
											Description:         "initiatorName is the custom iSCSI Initiator Name.",
											MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name.",
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
									Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.",
									MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.",
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
											Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
											MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
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
																Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.",
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
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
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
																			Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
																			MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
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
																Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value.",
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
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
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
																Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token.",
																MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"expiration_seconds": schema.Int64Attribute{
																Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token.",
																MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token.",
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
											Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
											MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'.",
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
											Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
											MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"items": schema.ListNestedAttribute{
											Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value.",
											MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value.",
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
														Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
														MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.",
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
											Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.",
											MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_opentelemetry_io_open_telemetry_collector_v1alpha1")

	var data OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "opentelemetry.io", Version: "v1alpha1", Resource: "OpenTelemetryCollector"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse OpentelemetryIoOpenTelemetryCollectorV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("opentelemetry.io/v1alpha1")
	data.Kind = pointer.String("OpenTelemetryCollector")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
