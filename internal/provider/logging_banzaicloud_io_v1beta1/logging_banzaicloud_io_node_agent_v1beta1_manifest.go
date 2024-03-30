/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package logging_banzaicloud_io_v1beta1

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
	_ datasource.DataSource = &LoggingBanzaicloudIoNodeAgentV1Beta1Manifest{}
)

func NewLoggingBanzaicloudIoNodeAgentV1Beta1Manifest() datasource.DataSource {
	return &LoggingBanzaicloudIoNodeAgentV1Beta1Manifest{}
}

type LoggingBanzaicloudIoNodeAgentV1Beta1Manifest struct{}

type LoggingBanzaicloudIoNodeAgentV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		LoggingRef *string `tfsdk:"logging_ref" json:"loggingRef,omitempty"`
		Metadata   *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"metadata" json:"metadata,omitempty"`
		NodeAgentFluentbit *struct {
			BufferStorage *struct {
				Storage_backlog_mem_limit           *string `tfsdk:"storage_backlog_mem_limit" json:"storage.backlog.mem_limit,omitempty"`
				Storage_checksum                    *string `tfsdk:"storage_checksum" json:"storage.checksum,omitempty"`
				Storage_delete_irrecoverable_chunks *string `tfsdk:"storage_delete_irrecoverable_chunks" json:"storage.delete_irrecoverable_chunks,omitempty"`
				Storage_metrics                     *string `tfsdk:"storage_metrics" json:"storage.metrics,omitempty"`
				Storage_path                        *string `tfsdk:"storage_path" json:"storage.path,omitempty"`
				Storage_sync                        *string `tfsdk:"storage_sync" json:"storage.sync,omitempty"`
			} `tfsdk:"buffer_storage" json:"bufferStorage,omitempty"`
			BufferStorageVolume *struct {
				EmptyDir *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				HostPath *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"host_path" json:"hostPath,omitempty"`
				Pvc *struct {
					Source *struct {
						ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"source" json:"source,omitempty"`
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
				} `tfsdk:"pvc" json:"pvc,omitempty"`
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
			} `tfsdk:"buffer_storage_volume" json:"bufferStorageVolume,omitempty"`
			ContainersPath     *string `tfsdk:"containers_path" json:"containersPath,omitempty"`
			CoroStackSize      *int64  `tfsdk:"coro_stack_size" json:"coroStackSize,omitempty"`
			CustomConfigSecret *string `tfsdk:"custom_config_secret" json:"customConfigSecret,omitempty"`
			DaemonSet          *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec *struct {
					MinReadySeconds      *int64 `tfsdk:"min_ready_seconds" json:"minReadySeconds,omitempty"`
					RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" json:"revisionHistoryLimit,omitempty"`
					Selector             *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
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
							NodeName          *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
							NodeSelector      *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
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
					UpdateStrategy *struct {
						RollingUpdate *struct {
							MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
							MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
						} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"daemon_set" json:"daemonSet,omitempty"`
			DisableKubernetesFilter *bool `tfsdk:"disable_kubernetes_filter" json:"disableKubernetesFilter,omitempty"`
			EnableUpstream          *bool `tfsdk:"enable_upstream" json:"enableUpstream,omitempty"`
			Enabled                 *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			ExtraVolumeMounts       *[]struct {
				Destination *string `tfsdk:"destination" json:"destination,omitempty"`
				ReadOnly    *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				Source      *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"extra_volume_mounts" json:"extraVolumeMounts,omitempty"`
			FilterAws *struct {
				Match             *string `tfsdk:"match" json:"Match,omitempty"`
				Account_id        *bool   `tfsdk:"account_id" json:"account_id,omitempty"`
				Ami_id            *bool   `tfsdk:"ami_id" json:"ami_id,omitempty"`
				Az                *bool   `tfsdk:"az" json:"az,omitempty"`
				Ec2_instance_id   *bool   `tfsdk:"ec2_instance_id" json:"ec2_instance_id,omitempty"`
				Ec2_instance_type *bool   `tfsdk:"ec2_instance_type" json:"ec2_instance_type,omitempty"`
				Hostname          *bool   `tfsdk:"hostname" json:"hostname,omitempty"`
				Imds_version      *string `tfsdk:"imds_version" json:"imds_version,omitempty"`
				Private_ip        *bool   `tfsdk:"private_ip" json:"private_ip,omitempty"`
				Vpc_id            *bool   `tfsdk:"vpc_id" json:"vpc_id,omitempty"`
			} `tfsdk:"filter_aws" json:"filterAws,omitempty"`
			FilterKubernetes *struct {
				Annotations                 *string `tfsdk:"annotations" json:"Annotations,omitempty"`
				Buffer_Size                 *string `tfsdk:"buffer__size" json:"Buffer_Size,omitempty"`
				Cache_Use_Docker_Id         *string `tfsdk:"cache__use__docker__id" json:"Cache_Use_Docker_Id,omitempty"`
				DNS_Retries                 *string `tfsdk:"dns__retries" json:"DNS_Retries,omitempty"`
				DNS_Wait_Time               *string `tfsdk:"dns__wait__time" json:"DNS_Wait_Time,omitempty"`
				Dummy_Meta                  *string `tfsdk:"dummy__meta" json:"Dummy_Meta,omitempty"`
				K8S_Logging_Exclude         *string `tfsdk:"k8_s__logging__exclude" json:"K8S-Logging.Exclude,omitempty"`
				K8S_Logging_Parser          *string `tfsdk:"k8_s__logging__parser" json:"K8S-Logging.Parser,omitempty"`
				Keep_Log                    *string `tfsdk:"keep__log" json:"Keep_Log,omitempty"`
				Kube_CA_File                *string `tfsdk:"kube_ca__file" json:"Kube_CA_File,omitempty"`
				Kube_CA_Path                *string `tfsdk:"kube_ca__path" json:"Kube_CA_Path,omitempty"`
				Kube_Meta_Cache_TTL         *string `tfsdk:"kube__meta__cache_ttl" json:"Kube_Meta_Cache_TTL,omitempty"`
				Kube_Tag_Prefix             *string `tfsdk:"kube__tag__prefix" json:"Kube_Tag_Prefix,omitempty"`
				Kube_Token_File             *string `tfsdk:"kube__token__file" json:"Kube_Token_File,omitempty"`
				Kube_Token_TTL              *string `tfsdk:"kube__token_ttl" json:"Kube_Token_TTL,omitempty"`
				Kube_URL                    *string `tfsdk:"kube__url" json:"Kube_URL,omitempty"`
				Kube_meta_preload_cache_dir *string `tfsdk:"kube_meta_preload_cache_dir" json:"Kube_meta_preload_cache_dir,omitempty"`
				Kubelet_Port                *string `tfsdk:"kubelet__port" json:"Kubelet_Port,omitempty"`
				Labels                      *string `tfsdk:"labels" json:"Labels,omitempty"`
				Match                       *string `tfsdk:"match" json:"Match,omitempty"`
				Merge_Log                   *string `tfsdk:"merge__log" json:"Merge_Log,omitempty"`
				Merge_Log_Key               *string `tfsdk:"merge__log__key" json:"Merge_Log_Key,omitempty"`
				Merge_Log_Trim              *string `tfsdk:"merge__log__trim" json:"Merge_Log_Trim,omitempty"`
				Merge_Parser                *string `tfsdk:"merge__parser" json:"Merge_Parser,omitempty"`
				Regex_Parser                *string `tfsdk:"regex__parser" json:"Regex_Parser,omitempty"`
				Use_Journal                 *string `tfsdk:"use__journal" json:"Use_Journal,omitempty"`
				Use_Kubelet                 *string `tfsdk:"use__kubelet" json:"Use_Kubelet,omitempty"`
				Tls_debug                   *string `tfsdk:"tls_debug" json:"tls.debug,omitempty"`
				Tls_verify                  *string `tfsdk:"tls_verify" json:"tls.verify,omitempty"`
			} `tfsdk:"filter_kubernetes" json:"filterKubernetes,omitempty"`
			Flush          *int64 `tfsdk:"flush" json:"flush,omitempty"`
			ForwardOptions *struct {
				Require_ack_response     *bool   `tfsdk:"require_ack_response" json:"Require_ack_response,omitempty"`
				Retry_Limit              *string `tfsdk:"retry__limit" json:"Retry_Limit,omitempty"`
				Send_options             *bool   `tfsdk:"send_options" json:"Send_options,omitempty"`
				Tag                      *string `tfsdk:"tag" json:"Tag,omitempty"`
				Time_as_Integer          *bool   `tfsdk:"time_as__integer" json:"Time_as_Integer,omitempty"`
				Workers                  *int64  `tfsdk:"workers" json:"Workers,omitempty"`
				Storage_total_limit_size *string `tfsdk:"storage_total_limit_size" json:"storage.total_limit_size,omitempty"`
			} `tfsdk:"forward_options" json:"forwardOptions,omitempty"`
			Grace     *int64 `tfsdk:"grace" json:"grace,omitempty"`
			InputTail *struct {
				Buffer_Chunk_Size  *string   `tfsdk:"buffer__chunk__size" json:"Buffer_Chunk_Size,omitempty"`
				Buffer_Max_Size    *string   `tfsdk:"buffer__max__size" json:"Buffer_Max_Size,omitempty"`
				DB                 *string   `tfsdk:"db" json:"DB,omitempty"`
				DB_journal_mode    *string   `tfsdk:"db_journal_mode" json:"DB.journal_mode,omitempty"`
				DB_locking         *bool     `tfsdk:"db_locking" json:"DB.locking,omitempty"`
				DB_Sync            *string   `tfsdk:"db__sync" json:"DB_Sync,omitempty"`
				Docker_Mode        *string   `tfsdk:"docker__mode" json:"Docker_Mode,omitempty"`
				Docker_Mode_Flush  *string   `tfsdk:"docker__mode__flush" json:"Docker_Mode_Flush,omitempty"`
				Docker_Mode_Parser *string   `tfsdk:"docker__mode__parser" json:"Docker_Mode_Parser,omitempty"`
				Exclude_Path       *string   `tfsdk:"exclude__path" json:"Exclude_Path,omitempty"`
				Ignore_Older       *string   `tfsdk:"ignore__older" json:"Ignore_Older,omitempty"`
				Key                *string   `tfsdk:"key" json:"Key,omitempty"`
				Mem_Buf_Limit      *string   `tfsdk:"mem__buf__limit" json:"Mem_Buf_Limit,omitempty"`
				Multiline          *string   `tfsdk:"multiline" json:"Multiline,omitempty"`
				Multiline_Flush    *string   `tfsdk:"multiline__flush" json:"Multiline_Flush,omitempty"`
				Parser             *string   `tfsdk:"parser" json:"Parser,omitempty"`
				Parser_Firstline   *string   `tfsdk:"parser__firstline" json:"Parser_Firstline,omitempty"`
				Parser_N           *[]string `tfsdk:"parser_n" json:"Parser_N,omitempty"`
				Path               *string   `tfsdk:"path" json:"Path,omitempty"`
				Path_Key           *string   `tfsdk:"path__key" json:"Path_Key,omitempty"`
				Read_From_Head     *bool     `tfsdk:"read__from__head" json:"Read_From_Head,omitempty"`
				Refresh_Interval   *string   `tfsdk:"refresh__interval" json:"Refresh_Interval,omitempty"`
				Rotate_Wait        *string   `tfsdk:"rotate__wait" json:"Rotate_Wait,omitempty"`
				Skip_Long_Lines    *string   `tfsdk:"skip__long__lines" json:"Skip_Long_Lines,omitempty"`
				Tag                *string   `tfsdk:"tag" json:"Tag,omitempty"`
				Tag_Regex          *string   `tfsdk:"tag__regex" json:"Tag_Regex,omitempty"`
				Multiline_parser   *[]string `tfsdk:"multiline_parser" json:"multiline.parser,omitempty"`
				Storage_type       *string   `tfsdk:"storage_type" json:"storage.type,omitempty"`
			} `tfsdk:"input_tail" json:"inputTail,omitempty"`
			LivenessDefaultCheck *bool   `tfsdk:"liveness_default_check" json:"livenessDefaultCheck,omitempty"`
			LogLevel             *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			Metrics              *struct {
				Interval              *string `tfsdk:"interval" json:"interval,omitempty"`
				Path                  *string `tfsdk:"path" json:"path,omitempty"`
				Port                  *int64  `tfsdk:"port" json:"port,omitempty"`
				PrometheusAnnotations *bool   `tfsdk:"prometheus_annotations" json:"prometheusAnnotations,omitempty"`
				PrometheusRules       *bool   `tfsdk:"prometheus_rules" json:"prometheusRules,omitempty"`
				ServiceMonitor        *bool   `tfsdk:"service_monitor" json:"serviceMonitor,omitempty"`
				ServiceMonitorConfig  *struct {
					AdditionalLabels  *map[string]string `tfsdk:"additional_labels" json:"additionalLabels,omitempty"`
					HonorLabels       *bool              `tfsdk:"honor_labels" json:"honorLabels,omitempty"`
					MetricRelabelings *[]struct {
						Action       *string   `tfsdk:"action" json:"action,omitempty"`
						Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
						Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
						Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
						Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
						SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
						TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
					} `tfsdk:"metric_relabelings" json:"metricRelabelings,omitempty"`
					Relabelings *[]struct {
						Action       *string   `tfsdk:"action" json:"action,omitempty"`
						Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
						Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
						Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
						Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
						SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
						TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
					} `tfsdk:"relabelings" json:"relabelings,omitempty"`
					Scheme    *string `tfsdk:"scheme" json:"scheme,omitempty"`
					TlsConfig *struct {
						Ca *struct {
							ConfigMap *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							Secret *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"ca" json:"ca,omitempty"`
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
							ConfigMap *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"config_map" json:"configMap,omitempty"`
							Secret *struct {
								Key      *string `tfsdk:"key" json:"key,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
						} `tfsdk:"cert" json:"cert,omitempty"`
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				} `tfsdk:"service_monitor_config" json:"serviceMonitorConfig,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			MetricsService *struct {
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
			} `tfsdk:"metrics_service" json:"metricsService,omitempty"`
			Network *struct {
				ConnectTimeout         *int64  `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
				ConnectTimeoutLogError *bool   `tfsdk:"connect_timeout_log_error" json:"connectTimeoutLogError,omitempty"`
				DnsMode                *string `tfsdk:"dns_mode" json:"dnsMode,omitempty"`
				DnsPreferIpv4          *bool   `tfsdk:"dns_prefer_ipv4" json:"dnsPreferIpv4,omitempty"`
				DnsResolver            *string `tfsdk:"dns_resolver" json:"dnsResolver,omitempty"`
				Keepalive              *bool   `tfsdk:"keepalive" json:"keepalive,omitempty"`
				KeepaliveIdleTimeout   *int64  `tfsdk:"keepalive_idle_timeout" json:"keepaliveIdleTimeout,omitempty"`
				KeepaliveMaxRecycle    *int64  `tfsdk:"keepalive_max_recycle" json:"keepaliveMaxRecycle,omitempty"`
				SourceAddress          *string `tfsdk:"source_address" json:"sourceAddress,omitempty"`
			} `tfsdk:"network" json:"network,omitempty"`
			PodPriorityClassName *string `tfsdk:"pod_priority_class_name" json:"podPriorityClassName,omitempty"`
			Positiondb           *struct {
				EmptyDir *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				HostPath *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"host_path" json:"hostPath,omitempty"`
				Pvc *struct {
					Source *struct {
						ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"source" json:"source,omitempty"`
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
				} `tfsdk:"pvc" json:"pvc,omitempty"`
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
			} `tfsdk:"positiondb" json:"positiondb,omitempty"`
			Security *struct {
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
				PodSecurityPolicyCreate      *bool `tfsdk:"pod_security_policy_create" json:"podSecurityPolicyCreate,omitempty"`
				RoleBasedAccessControlCreate *bool `tfsdk:"role_based_access_control_create" json:"roleBasedAccessControlCreate,omitempty"`
				SecurityContext              *struct {
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
				ServiceAccount *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			} `tfsdk:"security" json:"security,omitempty"`
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
			TargetHost *string `tfsdk:"target_host" json:"targetHost,omitempty"`
			TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			Tls        *struct {
				Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				SharedKey  *string `tfsdk:"shared_key" json:"sharedKey,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			VarLogsPath *string `tfsdk:"var_logs_path" json:"varLogsPath,omitempty"`
		} `tfsdk:"node_agent_fluentbit" json:"nodeAgentFluentbit,omitempty"`
		Profile *string `tfsdk:"profile" json:"profile,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LoggingBanzaicloudIoNodeAgentV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_logging_banzaicloud_io_node_agent_v1beta1_manifest"
}

func (r *LoggingBanzaicloudIoNodeAgentV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"logging_ref": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
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

					"node_agent_fluentbit": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"buffer_storage": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"storage_backlog_mem_limit": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_checksum": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_delete_irrecoverable_chunks": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_metrics": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_sync": schema.StringAttribute{
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

							"buffer_storage_volume": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
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

									"pvc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"source": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"containers_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"coro_stack_size": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_config_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"daemon_set": schema.SingleNestedAttribute{
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
											"min_ready_seconds": schema.Int64Attribute{
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

											"update_strategy": schema.SingleNestedAttribute{
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

							"disable_kubernetes_filter": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_upstream": schema.BoolAttribute{
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

							"extra_volume_mounts": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"destination": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^/.+$`), ""),
											},
										},

										"read_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^/.+$`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"filter_aws": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"match": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"account_id": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ami_id": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"az": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ec2_instance_id": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ec2_instance_type": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hostname": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"imds_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"private_ip": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vpc_id": schema.BoolAttribute{
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

							"filter_kubernetes": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"buffer__size": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cache__use__docker__id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns__retries": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns__wait__time": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dummy__meta": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"k8_s__logging__exclude": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"k8_s__logging__parser": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keep__log": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube_ca__file": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube_ca__path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube__meta__cache_ttl": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube__tag__prefix": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube__token__file": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube__token_ttl": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube__url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kube_meta_preload_cache_dir": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"kubelet__port": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"match": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"merge__log": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"merge__log__key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"merge__log__trim": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"merge__parser": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"regex__parser": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"use__journal": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"use__kubelet": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_debug": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_verify": schema.StringAttribute{
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

							"flush": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"forward_options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"require_ack_response": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry__limit": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"send_options": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"time_as__integer": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"workers": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_total_limit_size": schema.StringAttribute{
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

							"grace": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"input_tail": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"buffer__chunk__size": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"buffer__max__size": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"db": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"db_journal_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"db_locking": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"db__sync": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"docker__mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"docker__mode__flush": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"docker__mode__parser": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"exclude__path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ignore__older": schema.StringAttribute{
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

									"mem__buf__limit": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"multiline": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"multiline__flush": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parser": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parser__firstline": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"parser_n": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
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

									"path__key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"read__from__head": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"refresh__interval": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"rotate__wait": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"skip__long__lines": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tag__regex": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"multiline_parser": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage_type": schema.StringAttribute{
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

							"liveness_default_check": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_level": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
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

									"port": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"prometheus_annotations": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"prometheus_rules": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_monitor": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_monitor_config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"additional_labels": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"honor_labels": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metric_relabelings": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
															},
														},

														"modulus": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"regex": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"replacement": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"separator": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_labels": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_label": schema.StringAttribute{
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

											"relabelings": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
															},
														},

														"modulus": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"regex": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"replacement": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"separator": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source_labels": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_label": schema.StringAttribute{
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

											"scheme": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_config": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"ca": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.SingleNestedAttribute{
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

															"secret": schema.SingleNestedAttribute{
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

													"ca_file": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.SingleNestedAttribute{
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

															"secret": schema.SingleNestedAttribute{
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

													"cert_file": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_file": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_secret": schema.SingleNestedAttribute{
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

													"server_name": schema.StringAttribute{
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

									"timeout": schema.StringAttribute{
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

							"metrics_service": schema.SingleNestedAttribute{
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
											"allocate_load_balancer_node_ports": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_ip": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_i_ps": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_i_ps": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_traffic_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"health_check_node_port": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"internal_traffic_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_families": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_class": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_ip": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"load_balancer_source_ranges": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
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

														"node_port": schema.Int64Attribute{
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

														"target_port": schema.StringAttribute{
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

											"publish_not_ready_addresses": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"selector": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"session_affinity": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"network": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"connect_timeout": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_log_error": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_prefer_ipv4": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_resolver": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_idle_timeout": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"keepalive_max_recycle": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"source_address": schema.StringAttribute{
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

							"pod_priority_class_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"positiondb": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
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

									"pvc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"source": schema.SingleNestedAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"security": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"pod_security_context": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
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

									"pod_security_policy_create": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"role_based_access_control_create": schema.BoolAttribute{
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

									"service_account": schema.StringAttribute{
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

							"service_account": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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

									"secrets": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
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

												"resource_version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uid": schema.StringAttribute{
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

							"target_host": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"target_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"shared_key": schema.StringAttribute{
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

							"var_logs_path": schema.StringAttribute{
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

					"profile": schema.StringAttribute{
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
	}
}

func (r *LoggingBanzaicloudIoNodeAgentV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_logging_banzaicloud_io_node_agent_v1beta1_manifest")

	var model LoggingBanzaicloudIoNodeAgentV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("logging.banzaicloud.io/v1beta1")
	model.Kind = pointer.String("NodeAgent")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
