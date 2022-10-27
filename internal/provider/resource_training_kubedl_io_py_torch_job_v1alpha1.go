/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type TrainingKubedlIoPyTorchJobV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*TrainingKubedlIoPyTorchJobV1Alpha1Resource)(nil)
)

type TrainingKubedlIoPyTorchJobV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type TrainingKubedlIoPyTorchJobV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" yaml:"activeDeadlineSeconds,omitempty"`

		BackoffLimit *int64 `tfsdk:"backoff_limit" yaml:"backoffLimit,omitempty"`

		CacheBackend *struct {
			CacheEngine *struct {
				Fluid *struct {
					AlluxioRuntime *struct {
						Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

						TieredStorage *[]struct {
							CachePath *string `tfsdk:"cache_path" yaml:"cachePath,omitempty"`

							MediumType *string `tfsdk:"medium_type" yaml:"mediumType,omitempty"`

							Quota *string `tfsdk:"quota" yaml:"quota,omitempty"`
						} `tfsdk:"tiered_storage" yaml:"tieredStorage,omitempty"`
					} `tfsdk:"alluxio_runtime" yaml:"alluxioRuntime,omitempty"`
				} `tfsdk:"fluid" yaml:"fluid,omitempty"`
			} `tfsdk:"cache_engine" yaml:"cacheEngine,omitempty"`

			Dataset *struct {
				DataSources *[]struct {
					Location *string `tfsdk:"location" yaml:"location,omitempty"`

					SubDirName *string `tfsdk:"sub_dir_name" yaml:"subDirName,omitempty"`
				} `tfsdk:"data_sources" yaml:"dataSources,omitempty"`
			} `tfsdk:"dataset" yaml:"dataset,omitempty"`

			MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Options *struct {
				IdleTime *int64 `tfsdk:"idle_time" yaml:"idleTime,omitempty"`
			} `tfsdk:"options" yaml:"options,omitempty"`
		} `tfsdk:"cache_backend" yaml:"cacheBackend,omitempty"`

		CleanPodPolicy *string `tfsdk:"clean_pod_policy" yaml:"cleanPodPolicy,omitempty"`

		CronPolicy *struct {
			ConcurrencyPolicy *string `tfsdk:"concurrency_policy" yaml:"concurrencyPolicy,omitempty"`

			Deadline *string `tfsdk:"deadline" yaml:"deadline,omitempty"`

			HistoryLimit *int64 `tfsdk:"history_limit" yaml:"historyLimit,omitempty"`

			Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`

			Suspend *bool `tfsdk:"suspend" yaml:"suspend,omitempty"`
		} `tfsdk:"cron_policy" yaml:"cronPolicy,omitempty"`

		ModelVersion *struct {
			CreatedBy *string `tfsdk:"created_by" yaml:"createdBy,omitempty"`

			ImageRepo *string `tfsdk:"image_repo" yaml:"imageRepo,omitempty"`

			ImageTag *string `tfsdk:"image_tag" yaml:"imageTag,omitempty"`

			ModelName *string `tfsdk:"model_name" yaml:"modelName,omitempty"`

			Storage *struct {
				AWSEfs *struct {
					Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

					VolumeHandle *string `tfsdk:"volume_handle" yaml:"volumeHandle,omitempty"`
				} `tfsdk:"aws_efs" yaml:"AWSEfs,omitempty"`

				LocalStorage *struct {
					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					NodeName *string `tfsdk:"node_name" yaml:"nodeName,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"local_storage" yaml:"localStorage,omitempty"`

				Nfs *struct {
					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Server *string `tfsdk:"server" yaml:"server,omitempty"`
				} `tfsdk:"nfs" yaml:"nfs,omitempty"`
			} `tfsdk:"storage" yaml:"storage,omitempty"`
		} `tfsdk:"model_version" yaml:"modelVersion,omitempty"`

		PytorchReplicaSpecs *struct {
			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			RestartPolicy *string `tfsdk:"restart_policy" yaml:"restartPolicy,omitempty"`

			SpotReplicaSpec *struct {
				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

				SpotReplicaNumber *int64 `tfsdk:"spot_replica_number" yaml:"spotReplicaNumber,omitempty"`
			} `tfsdk:"spot_replica_spec" yaml:"spotReplicaSpec,omitempty"`

			Template *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`

				Spec *struct {
					ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" yaml:"activeDeadlineSeconds,omitempty"`

					Affinity *struct {
						NodeAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								Preference *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchFields *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
								} `tfsdk:"preference" yaml:"preference,omitempty"`

								Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

							RequiredDuringSchedulingIgnoredDuringExecution *struct {
								NodeSelectorTerms *[]struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchFields *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
								} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

						PodAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								PodAffinityTerm *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

									Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

									TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
								} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

								Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

							RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

						PodAntiAffinity *struct {
							PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
								PodAffinityTerm *struct {
									LabelSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

									Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

									TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
								} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

								Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
							} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

							RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" yaml:"affinity,omitempty"`

					AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" yaml:"automountServiceAccountToken,omitempty"`

					Containers *[]struct {
						Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						EnvFrom *[]struct {
							ConfigMapRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
						} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Lifecycle *struct {
							PostStart *struct {
								Exec *struct {
									Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
								} `tfsdk:"exec" yaml:"exec,omitempty"`

								HttpGet *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									HttpHeaders *[]struct {
										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

									Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
								} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

								TcpSocket *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
								} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
							} `tfsdk:"post_start" yaml:"postStart,omitempty"`

							PreStop *struct {
								Exec *struct {
									Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
								} `tfsdk:"exec" yaml:"exec,omitempty"`

								HttpGet *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									HttpHeaders *[]struct {
										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

									Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
								} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

								TcpSocket *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
								} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
							} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
						} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

						LivenessProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Ports *[]struct {
							ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

							HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

							HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
						} `tfsdk:"ports" yaml:"ports,omitempty"`

						ReadinessProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

						Resources *struct {
							Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

							Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`

						SecurityContext *struct {
							AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

							Capabilities *struct {
								Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

								Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
							} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

							Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

							ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

							ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

							RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

							RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

							RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

							SeLinuxOptions *struct {
								Level *string `tfsdk:"level" yaml:"level,omitempty"`

								Role *string `tfsdk:"role" yaml:"role,omitempty"`

								Type *string `tfsdk:"type" yaml:"type,omitempty"`

								User *string `tfsdk:"user" yaml:"user,omitempty"`
							} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

							SeccompProfile *struct {
								LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

								Type *string `tfsdk:"type" yaml:"type,omitempty"`
							} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

							WindowsOptions *struct {
								GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

								GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

								HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

								RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
							} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
						} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

						StartupProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

						Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

						StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

						TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

						TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

						Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

						VolumeDevices *[]struct {
							DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

						VolumeMounts *[]struct {
							MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

							MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

							SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
						} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

						WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
					} `tfsdk:"containers" yaml:"containers,omitempty"`

					DnsConfig *struct {
						Nameservers *[]string `tfsdk:"nameservers" yaml:"nameservers,omitempty"`

						Options *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"options" yaml:"options,omitempty"`

						Searches *[]string `tfsdk:"searches" yaml:"searches,omitempty"`
					} `tfsdk:"dns_config" yaml:"dnsConfig,omitempty"`

					DnsPolicy *string `tfsdk:"dns_policy" yaml:"dnsPolicy,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					EphemeralContainers *[]struct {
						Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						EnvFrom *[]struct {
							ConfigMapRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
						} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Lifecycle *struct {
							PostStart *struct {
								Exec *struct {
									Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
								} `tfsdk:"exec" yaml:"exec,omitempty"`

								HttpGet *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									HttpHeaders *[]struct {
										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

									Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
								} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

								TcpSocket *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
								} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
							} `tfsdk:"post_start" yaml:"postStart,omitempty"`

							PreStop *struct {
								Exec *struct {
									Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
								} `tfsdk:"exec" yaml:"exec,omitempty"`

								HttpGet *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									HttpHeaders *[]struct {
										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

									Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
								} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

								TcpSocket *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
								} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
							} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
						} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

						LivenessProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Ports *[]struct {
							ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

							HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

							HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
						} `tfsdk:"ports" yaml:"ports,omitempty"`

						ReadinessProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

						Resources *struct {
							Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

							Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`

						SecurityContext *struct {
							AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

							Capabilities *struct {
								Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

								Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
							} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

							Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

							ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

							ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

							RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

							RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

							RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

							SeLinuxOptions *struct {
								Level *string `tfsdk:"level" yaml:"level,omitempty"`

								Role *string `tfsdk:"role" yaml:"role,omitempty"`

								Type *string `tfsdk:"type" yaml:"type,omitempty"`

								User *string `tfsdk:"user" yaml:"user,omitempty"`
							} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

							SeccompProfile *struct {
								LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

								Type *string `tfsdk:"type" yaml:"type,omitempty"`
							} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

							WindowsOptions *struct {
								GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

								GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

								HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

								RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
							} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
						} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

						StartupProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

						Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

						StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

						TargetContainerName *string `tfsdk:"target_container_name" yaml:"targetContainerName,omitempty"`

						TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

						TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

						Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

						VolumeDevices *[]struct {
							DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

						VolumeMounts *[]struct {
							MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

							MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

							SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
						} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

						WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
					} `tfsdk:"ephemeral_containers" yaml:"ephemeralContainers,omitempty"`

					HostAliases *[]struct {
						Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

						Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
					} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

					HostIPC *bool `tfsdk:"host_ipc" yaml:"hostIPC,omitempty"`

					HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

					HostPID *bool `tfsdk:"host_pid" yaml:"hostPID,omitempty"`

					Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					InitContainers *[]struct {
						Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						EnvFrom *[]struct {
							ConfigMapRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

							Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
						} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Lifecycle *struct {
							PostStart *struct {
								Exec *struct {
									Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
								} `tfsdk:"exec" yaml:"exec,omitempty"`

								HttpGet *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									HttpHeaders *[]struct {
										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

									Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
								} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

								TcpSocket *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
								} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
							} `tfsdk:"post_start" yaml:"postStart,omitempty"`

							PreStop *struct {
								Exec *struct {
									Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
								} `tfsdk:"exec" yaml:"exec,omitempty"`

								HttpGet *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									HttpHeaders *[]struct {
										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Value *string `tfsdk:"value" yaml:"value,omitempty"`
									} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

									Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
								} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

								TcpSocket *struct {
									Host *string `tfsdk:"host" yaml:"host,omitempty"`

									Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
								} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
							} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
						} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

						LivenessProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Ports *[]struct {
							ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

							HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

							HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
						} `tfsdk:"ports" yaml:"ports,omitempty"`

						ReadinessProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

						Resources *struct {
							Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

							Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`

						SecurityContext *struct {
							AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

							Capabilities *struct {
								Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

								Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
							} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

							Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

							ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

							ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

							RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

							RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

							RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

							SeLinuxOptions *struct {
								Level *string `tfsdk:"level" yaml:"level,omitempty"`

								Role *string `tfsdk:"role" yaml:"role,omitempty"`

								Type *string `tfsdk:"type" yaml:"type,omitempty"`

								User *string `tfsdk:"user" yaml:"user,omitempty"`
							} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

							SeccompProfile *struct {
								LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

								Type *string `tfsdk:"type" yaml:"type,omitempty"`
							} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

							WindowsOptions *struct {
								GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

								GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

								HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

								RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
							} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
						} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

						StartupProbe *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

							HttpGet *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								HttpHeaders *[]struct {
									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Value *string `tfsdk:"value" yaml:"value,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

							PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

							SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

							TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

							TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
						} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

						Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

						StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

						TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

						TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

						Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

						VolumeDevices *[]struct {
							DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

						VolumeMounts *[]struct {
							MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

							MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

							SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
						} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

						WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
					} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`

					NodeName *string `tfsdk:"node_name" yaml:"nodeName,omitempty"`

					NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

					Overhead *map[string]string `tfsdk:"overhead" yaml:"overhead,omitempty"`

					PreemptionPolicy *string `tfsdk:"preemption_policy" yaml:"preemptionPolicy,omitempty"`

					Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

					ReadinessGates *[]struct {
						ConditionType *string `tfsdk:"condition_type" yaml:"conditionType,omitempty"`
					} `tfsdk:"readiness_gates" yaml:"readinessGates,omitempty"`

					RestartPolicy *string `tfsdk:"restart_policy" yaml:"restartPolicy,omitempty"`

					RuntimeClassName *string `tfsdk:"runtime_class_name" yaml:"runtimeClassName,omitempty"`

					SchedulerName *string `tfsdk:"scheduler_name" yaml:"schedulerName,omitempty"`

					SecurityContext *struct {
						FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

						FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" yaml:"fsGroupChangePolicy,omitempty"`

						RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

						RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

						SeLinuxOptions *struct {
							Level *string `tfsdk:"level" yaml:"level,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`
						} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

						SeccompProfile *struct {
							LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`
						} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

						SupplementalGroups *[]string `tfsdk:"supplemental_groups" yaml:"supplementalGroups,omitempty"`

						Sysctls *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"sysctls" yaml:"sysctls,omitempty"`

						WindowsOptions *struct {
							GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

							GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

							HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

							RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
						} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

					ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

					ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

					SetHostnameAsFQDN *bool `tfsdk:"set_hostname_as_fqdn" yaml:"setHostnameAsFQDN,omitempty"`

					ShareProcessNamespace *bool `tfsdk:"share_process_namespace" yaml:"shareProcessNamespace,omitempty"`

					Subdomain *string `tfsdk:"subdomain" yaml:"subdomain,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					Tolerations *[]struct {
						Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

					TopologySpreadConstraints *[]struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`

					Volumes *[]struct {
						AwsElasticBlockStore *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
						} `tfsdk:"aws_elastic_block_store" yaml:"awsElasticBlockStore,omitempty"`

						AzureDisk *struct {
							CachingMode *string `tfsdk:"caching_mode" yaml:"cachingMode,omitempty"`

							DiskName *string `tfsdk:"disk_name" yaml:"diskName,omitempty"`

							DiskURI *string `tfsdk:"disk_uri" yaml:"diskURI,omitempty"`

							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
						} `tfsdk:"azure_disk" yaml:"azureDisk,omitempty"`

						AzureFile *struct {
							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

							ShareName *string `tfsdk:"share_name" yaml:"shareName,omitempty"`
						} `tfsdk:"azure_file" yaml:"azureFile,omitempty"`

						Cephfs *struct {
							Monitors *[]string `tfsdk:"monitors" yaml:"monitors,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretFile *string `tfsdk:"secret_file" yaml:"secretFile,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`
						} `tfsdk:"cephfs" yaml:"cephfs,omitempty"`

						Cinder *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

							VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
						} `tfsdk:"cinder" yaml:"cinder,omitempty"`

						ConfigMap *struct {
							DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

							Items *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"items" yaml:"items,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map" yaml:"configMap,omitempty"`

						Csi *struct {
							Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							NodePublishSecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"node_publish_secret_ref" yaml:"nodePublishSecretRef,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							VolumeAttributes *map[string]string `tfsdk:"volume_attributes" yaml:"volumeAttributes,omitempty"`
						} `tfsdk:"csi" yaml:"csi,omitempty"`

						DownwardAPI *struct {
							DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

							Items *[]struct {
								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
							} `tfsdk:"items" yaml:"items,omitempty"`
						} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

						EmptyDir *struct {
							Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

							SizeLimit utilities.IntOrString `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
						} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

						Ephemeral *struct {
							VolumeClaimTemplate *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

									Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

									Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"metadata" yaml:"metadata,omitempty"`

								Spec *struct {
									AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

									DataSource *struct {
										ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

										Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"data_source" yaml:"dataSource,omitempty"`

									DataSourceRef *struct {
										ApiGroup *string `tfsdk:"api_group" yaml:"apiGroup,omitempty"`

										Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`
									} `tfsdk:"data_source_ref" yaml:"dataSourceRef,omitempty"`

									Resources *struct {
										Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

										Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
									} `tfsdk:"resources" yaml:"resources,omitempty"`

									Selector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"selector" yaml:"selector,omitempty"`

									StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

									VolumeMode *string `tfsdk:"volume_mode" yaml:"volumeMode,omitempty"`

									VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`
								} `tfsdk:"spec" yaml:"spec,omitempty"`
							} `tfsdk:"volume_claim_template" yaml:"volumeClaimTemplate,omitempty"`
						} `tfsdk:"ephemeral" yaml:"ephemeral,omitempty"`

						Fc *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							TargetWWNs *[]string `tfsdk:"target_ww_ns" yaml:"targetWWNs,omitempty"`

							Wwids *[]string `tfsdk:"wwids" yaml:"wwids,omitempty"`
						} `tfsdk:"fc" yaml:"fc,omitempty"`

						FlexVolume *struct {
							Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							Options *map[string]string `tfsdk:"options" yaml:"options,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
						} `tfsdk:"flex_volume" yaml:"flexVolume,omitempty"`

						Flocker *struct {
							DatasetName *string `tfsdk:"dataset_name" yaml:"datasetName,omitempty"`

							DatasetUUID *string `tfsdk:"dataset_uuid" yaml:"datasetUUID,omitempty"`
						} `tfsdk:"flocker" yaml:"flocker,omitempty"`

						GcePersistentDisk *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

							PdName *string `tfsdk:"pd_name" yaml:"pdName,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
						} `tfsdk:"gce_persistent_disk" yaml:"gcePersistentDisk,omitempty"`

						GitRepo *struct {
							Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`

							Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`

							Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`
						} `tfsdk:"git_repo" yaml:"gitRepo,omitempty"`

						Glusterfs *struct {
							Endpoints *string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
						} `tfsdk:"glusterfs" yaml:"glusterfs,omitempty"`

						HostPath *struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`
						} `tfsdk:"host_path" yaml:"hostPath,omitempty"`

						Iscsi *struct {
							ChapAuthDiscovery *bool `tfsdk:"chap_auth_discovery" yaml:"chapAuthDiscovery,omitempty"`

							ChapAuthSession *bool `tfsdk:"chap_auth_session" yaml:"chapAuthSession,omitempty"`

							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							InitiatorName *string `tfsdk:"initiator_name" yaml:"initiatorName,omitempty"`

							Iqn *string `tfsdk:"iqn" yaml:"iqn,omitempty"`

							IscsiInterface *string `tfsdk:"iscsi_interface" yaml:"iscsiInterface,omitempty"`

							Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

							Portals *[]string `tfsdk:"portals" yaml:"portals,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

							TargetPortal *string `tfsdk:"target_portal" yaml:"targetPortal,omitempty"`
						} `tfsdk:"iscsi" yaml:"iscsi,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Nfs *struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							Server *string `tfsdk:"server" yaml:"server,omitempty"`
						} `tfsdk:"nfs" yaml:"nfs,omitempty"`

						PersistentVolumeClaim *struct {
							ClaimName *string `tfsdk:"claim_name" yaml:"claimName,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
						} `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

						PhotonPersistentDisk *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							PdID *string `tfsdk:"pd_id" yaml:"pdID,omitempty"`
						} `tfsdk:"photon_persistent_disk" yaml:"photonPersistentDisk,omitempty"`

						PortworxVolume *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
						} `tfsdk:"portworx_volume" yaml:"portworxVolume,omitempty"`

						Projected *struct {
							DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

							Sources *[]struct {
								ConfigMap *struct {
									Items *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

										Path *string `tfsdk:"path" yaml:"path,omitempty"`
									} `tfsdk:"items" yaml:"items,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map" yaml:"configMap,omitempty"`

								DownwardAPI *struct {
									Items *[]struct {
										FieldRef *struct {
											ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

											FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
										} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

										Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

										Path *string `tfsdk:"path" yaml:"path,omitempty"`

										ResourceFieldRef *struct {
											ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

											Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

											Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
										} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
									} `tfsdk:"items" yaml:"items,omitempty"`
								} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

								Secret *struct {
									Items *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

										Path *string `tfsdk:"path" yaml:"path,omitempty"`
									} `tfsdk:"items" yaml:"items,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret" yaml:"secret,omitempty"`

								ServiceAccountToken *struct {
									Audience *string `tfsdk:"audience" yaml:"audience,omitempty"`

									ExpirationSeconds *int64 `tfsdk:"expiration_seconds" yaml:"expirationSeconds,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"service_account_token" yaml:"serviceAccountToken,omitempty"`
							} `tfsdk:"sources" yaml:"sources,omitempty"`
						} `tfsdk:"projected" yaml:"projected,omitempty"`

						Quobyte *struct {
							Group *string `tfsdk:"group" yaml:"group,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							Registry *string `tfsdk:"registry" yaml:"registry,omitempty"`

							Tenant *string `tfsdk:"tenant" yaml:"tenant,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`

							Volume *string `tfsdk:"volume" yaml:"volume,omitempty"`
						} `tfsdk:"quobyte" yaml:"quobyte,omitempty"`

						Rbd *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							Image *string `tfsdk:"image" yaml:"image,omitempty"`

							Keyring *string `tfsdk:"keyring" yaml:"keyring,omitempty"`

							Monitors *[]string `tfsdk:"monitors" yaml:"monitors,omitempty"`

							Pool *string `tfsdk:"pool" yaml:"pool,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`
						} `tfsdk:"rbd" yaml:"rbd,omitempty"`

						ScaleIO *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							Gateway *string `tfsdk:"gateway" yaml:"gateway,omitempty"`

							ProtectionDomain *string `tfsdk:"protection_domain" yaml:"protectionDomain,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

							SslEnabled *bool `tfsdk:"ssl_enabled" yaml:"sslEnabled,omitempty"`

							StorageMode *string `tfsdk:"storage_mode" yaml:"storageMode,omitempty"`

							StoragePool *string `tfsdk:"storage_pool" yaml:"storagePool,omitempty"`

							System *string `tfsdk:"system" yaml:"system,omitempty"`

							VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`
						} `tfsdk:"scale_io" yaml:"scaleIO,omitempty"`

						Secret *struct {
							DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

							Items *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"items" yaml:"items,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
						} `tfsdk:"secret" yaml:"secret,omitempty"`

						Storageos *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

							VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`

							VolumeNamespace *string `tfsdk:"volume_namespace" yaml:"volumeNamespace,omitempty"`
						} `tfsdk:"storageos" yaml:"storageos,omitempty"`

						VsphereVolume *struct {
							FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

							StoragePolicyID *string `tfsdk:"storage_policy_id" yaml:"storagePolicyID,omitempty"`

							StoragePolicyName *string `tfsdk:"storage_policy_name" yaml:"storagePolicyName,omitempty"`

							VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`
						} `tfsdk:"vsphere_volume" yaml:"vsphereVolume,omitempty"`
					} `tfsdk:"volumes" yaml:"volumes,omitempty"`
				} `tfsdk:"spec" yaml:"spec,omitempty"`
			} `tfsdk:"template" yaml:"template,omitempty"`
		} `tfsdk:"pytorch_replica_specs" yaml:"pytorchReplicaSpecs,omitempty"`

		SchedulingPolicy *struct {
			MinAvailable *int64 `tfsdk:"min_available" yaml:"minAvailable,omitempty"`

			Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

			PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

			Queue *string `tfsdk:"queue" yaml:"queue,omitempty"`
		} `tfsdk:"scheduling_policy" yaml:"schedulingPolicy,omitempty"`

		TtlSecondsAfterFinished *int64 `tfsdk:"ttl_seconds_after_finished" yaml:"ttlSecondsAfterFinished,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewTrainingKubedlIoPyTorchJobV1Alpha1Resource() resource.Resource {
	return &TrainingKubedlIoPyTorchJobV1Alpha1Resource{}
}

func (r *TrainingKubedlIoPyTorchJobV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_training_kubedl_io_py_torch_job_v1alpha1"
}

func (r *TrainingKubedlIoPyTorchJobV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"active_deadline_seconds": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backoff_limit": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cache_backend": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cache_engine": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fluid": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"alluxio_runtime": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"replicas": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tiered_storage": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"cache_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"medium_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"quota": {
																Description:         "",
																MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dataset": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"data_sources": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"location": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sub_dir_name": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount_path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"idle_time": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"clean_pod_policy": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cron_policy": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"concurrency_policy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"deadline": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.DateTime64Validator(),
								},
							},

							"history_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"schedule": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"suspend": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"model_version": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"created_by": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_repo": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_tag": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"model_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"aws_efs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"attributes": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_handle": {
												Description:         "",
												MarkdownDescription: "",

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

									"local_storage": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"mount_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "",
												MarkdownDescription: "",

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

									"nfs": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"mount_path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"server": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pytorch_replica_specs": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"replicas": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"restart_policy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spot_replica_spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority_class_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"spot_replica_number": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"template": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"finalizers": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

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

									"spec": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"active_deadline_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"affinity": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"node_affinity": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"preference": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"match_fields": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"weight": {
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

															"required_during_scheduling_ignored_during_execution": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"node_selector_terms": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"match_fields": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_affinity": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"pod_affinity_term": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"label_selector": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"operator": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"match_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace_selector": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"operator": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"match_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespaces": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"topology_key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"weight": {
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

															"required_during_scheduling_ignored_during_execution": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"match_labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespace_selector": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"match_labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespaces": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pod_anti_affinity": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"pod_affinity_term": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"label_selector": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"operator": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"match_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace_selector": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"operator": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"match_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespaces": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"topology_key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"weight": {
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

															"required_during_scheduling_ignored_during_execution": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"match_labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespace_selector": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"match_labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespaces": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"automount_service_account_token": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"containers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"args": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"command": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"env": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
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

																	"resource_field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
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

																	"secret_key_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"env_from": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lifecycle": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"post_start": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"command": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_get": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"http_headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
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

																			"path": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"scheme": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"tcp_socket": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"pre_stop": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"command": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_get": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"http_headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
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

																			"path": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"scheme": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"tcp_socket": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"liveness_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"ports": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"container_port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"host_ip": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"protocol": {
																Description:         "",
																MarkdownDescription: "",

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

													"readiness_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"requests": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"security_context": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_privilege_escalation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"capabilities": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"drop": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"privileged": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"proc_mount": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only_root_filesystem": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_group": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_non_root": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"se_linux_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"level": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"role": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"user": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"seccomp_profile": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"localhost_profile": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"type": {
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

															"windows_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"gmsa_credential_spec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"gmsa_credential_spec_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"host_process": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"run_as_user_name": {
																		Description:         "",
																		MarkdownDescription: "",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"startup_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"stdin": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"stdin_once": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_message_path": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_message_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tty": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_devices": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"device_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"volume_mounts": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"mount_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_propagation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sub_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sub_path_expr": {
																Description:         "",
																MarkdownDescription: "",

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

													"working_dir": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"dns_config": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"nameservers": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"options": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

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

													"searches": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dns_policy": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_service_links": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ephemeral_containers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"args": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"command": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"env": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
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

																	"resource_field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
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

																	"secret_key_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"env_from": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lifecycle": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"post_start": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"command": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_get": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"http_headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
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

																			"path": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"scheme": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"tcp_socket": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"pre_stop": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"command": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_get": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"http_headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
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

																			"path": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"scheme": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"tcp_socket": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"liveness_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"ports": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"container_port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"host_ip": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"protocol": {
																Description:         "",
																MarkdownDescription: "",

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

													"readiness_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"requests": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"security_context": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_privilege_escalation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"capabilities": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"drop": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"privileged": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"proc_mount": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only_root_filesystem": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_group": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_non_root": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"se_linux_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"level": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"role": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"user": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"seccomp_profile": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"localhost_profile": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"type": {
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

															"windows_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"gmsa_credential_spec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"gmsa_credential_spec_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"host_process": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"run_as_user_name": {
																		Description:         "",
																		MarkdownDescription: "",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"startup_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"stdin": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"stdin_once": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_container_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_message_path": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_message_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tty": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_devices": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"device_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"volume_mounts": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"mount_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_propagation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sub_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sub_path_expr": {
																Description:         "",
																MarkdownDescription: "",

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

													"working_dir": {
														Description:         "",
														MarkdownDescription: "",

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

											"host_aliases": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"hostnames": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ip": {
														Description:         "",
														MarkdownDescription: "",

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

											"host_ipc": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_network": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_pid": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hostname": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_pull_secrets": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

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

											"init_containers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"args": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"command": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"env": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
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

																	"resource_field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
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

																	"secret_key_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"env_from": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"prefix": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lifecycle": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"post_start": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"command": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_get": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"http_headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
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

																			"path": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"scheme": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"tcp_socket": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"pre_stop": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"exec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"command": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_get": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"http_headers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"value": {
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

																			"path": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"scheme": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"tcp_socket": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"host": {
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

																				Type: utilities.IntOrStringType{},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"liveness_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"ports": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"container_port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"host_ip": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"protocol": {
																Description:         "",
																MarkdownDescription: "",

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

													"readiness_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"requests": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"security_context": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_privilege_escalation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"capabilities": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"add": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"drop": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"privileged": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"proc_mount": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only_root_filesystem": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_group": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_non_root": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"se_linux_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"level": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"role": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"user": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"seccomp_profile": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"localhost_profile": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"type": {
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

															"windows_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"gmsa_credential_spec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"gmsa_credential_spec_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"host_process": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"run_as_user_name": {
																		Description:         "",
																		MarkdownDescription: "",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"startup_probe": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failure_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
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

																	"path": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"initial_delay_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"success_threshold": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
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

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"termination_grace_period_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"timeout_seconds": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"stdin": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"stdin_once": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_message_path": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination_message_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tty": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_devices": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"device_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"name": {
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

													"volume_mounts": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"mount_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_propagation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sub_path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sub_path_expr": {
																Description:         "",
																MarkdownDescription: "",

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

													"working_dir": {
														Description:         "",
														MarkdownDescription: "",

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

											"node_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selector": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"overhead": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preemption_policy": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"priority_class_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"readiness_gates": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"condition_type": {
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

											"restart_policy": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"runtime_class_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_group": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"fs_group_change_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_group": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_non_root": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"se_linux_options": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"level": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": {
																Description:         "",
																MarkdownDescription: "",

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

													"seccomp_profile": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"localhost_profile": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
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

													"supplemental_groups": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sysctls": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
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

													"windows_options": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"gmsa_credential_spec": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"gmsa_credential_spec_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_process": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user_name": {
																Description:         "",
																MarkdownDescription: "",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_account": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_account_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set_hostname_as_fqdn": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"share_process_namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subdomain": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_grace_period_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tolerations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"effect": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"operator": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"toleration_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

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

											"topology_spread_constraints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"label_selector": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"operator": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"topology_key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"when_unsatisfiable": {
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

											"volumes": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"aws_elastic_block_store": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"partition": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_id": {
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

													"azure_disk": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"caching_mode": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"disk_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"disk_uri": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"kind": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"azure_file": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"share_name": {
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

													"cephfs": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"monitors": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_file": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"user": {
																Description:         "",
																MarkdownDescription: "",

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

													"cinder": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"volume_id": {
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

													"config_map": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_mode": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"items": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"mode": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
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

															"name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"csi": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"driver": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"node_publish_secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_attributes": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"downward_api": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_mode": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"items": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
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

																	"mode": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"resource_field_ref": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
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
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"empty_dir": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"medium": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"size_limit": {
																Description:         "",
																MarkdownDescription: "",

																Type: utilities.IntOrStringType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ephemeral": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"volume_claim_template": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"annotations": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"finalizers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace": {
																				Description:         "",
																				MarkdownDescription: "",

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

																	"spec": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"access_modes": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"data_source": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"api_group": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"kind": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

																			"data_source_ref": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"api_group": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"kind": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
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

																			"resources": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"limits": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"requests": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"selector": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																							"key": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"operator": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.ListType{ElemType: types.StringType},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},
																						}),

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"match_labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"storage_class_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume_mode": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume_name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"fc": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"lun": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"target_ww_ns": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"wwids": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"flex_volume": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"driver": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"options": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"flocker": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"dataset_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"dataset_uuid": {
																Description:         "",
																MarkdownDescription: "",

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

													"gce_persistent_disk": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"partition": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"pd_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"git_repo": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"directory": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"repository": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"revision": {
																Description:         "",
																MarkdownDescription: "",

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

													"glusterfs": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"endpoints": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_path": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"type": {
																Description:         "",
																MarkdownDescription: "",

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

													"iscsi": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"chap_auth_discovery": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"chap_auth_session": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"initiator_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"iqn": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"iscsi_interface": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"lun": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"portals": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"target_portal": {
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

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"nfs": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"server": {
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

													"persistent_volume_claim": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"claim_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"photon_persistent_disk": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"pd_id": {
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

													"portworx_volume": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_id": {
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

													"projected": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_mode": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sources": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"items": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"mode": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.Int64Type,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"path": {
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

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"downward_api": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"items": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"field_ref": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"api_version": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"field_path": {
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

																					"mode": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.Int64Type,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"path": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"resource_field_ref": {
																						Description:         "",
																						MarkdownDescription: "",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"container_name": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"divisor": {
																								Description:         "",
																								MarkdownDescription: "",

																								Type: utilities.IntOrStringType{},

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"resource": {
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
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"items": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"mode": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.Int64Type,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"path": {
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

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"service_account_token": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"audience": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"expiration_seconds": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
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
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"quobyte": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"group": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"registry": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"tenant": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume": {
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

													"rbd": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"keyring": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"monitors": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"pool": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"user": {
																Description:         "",
																MarkdownDescription: "",

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

													"scale_io": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"gateway": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"protection_domain": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"ssl_enabled": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"storage_mode": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"storage_pool": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"system": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"volume_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"secret": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"default_mode": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"items": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"mode": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
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

															"optional": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_name": {
																Description:         "",
																MarkdownDescription: "",

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

													"storageos": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"read_only": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

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

															"volume_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_namespace": {
																Description:         "",
																MarkdownDescription: "",

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

													"vsphere_volume": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"storage_policy_id": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"storage_policy_name": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_path": {
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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"scheduling_policy": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"min_available": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"priority": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"priority_class_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"queue": {
								Description:         "",
								MarkdownDescription: "",

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

					"ttl_seconds_after_finished": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

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

func (r *TrainingKubedlIoPyTorchJobV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_training_kubedl_io_py_torch_job_v1alpha1")

	var state TrainingKubedlIoPyTorchJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TrainingKubedlIoPyTorchJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("training.kubedl.io/v1alpha1")
	goModel.Kind = utilities.Ptr("PyTorchJob")

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

func (r *TrainingKubedlIoPyTorchJobV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_training_kubedl_io_py_torch_job_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *TrainingKubedlIoPyTorchJobV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_training_kubedl_io_py_torch_job_v1alpha1")

	var state TrainingKubedlIoPyTorchJobV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TrainingKubedlIoPyTorchJobV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("training.kubedl.io/v1alpha1")
	goModel.Kind = utilities.Ptr("PyTorchJob")

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

func (r *TrainingKubedlIoPyTorchJobV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_training_kubedl_io_py_torch_job_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
