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

type SparkoperatorK8SIoSparkApplicationV1Beta2Resource struct{}

var (
	_ resource.Resource = (*SparkoperatorK8SIoSparkApplicationV1Beta2Resource)(nil)
)

type SparkoperatorK8SIoSparkApplicationV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SparkoperatorK8SIoSparkApplicationV1Beta2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		BatchSchedulerOptions *struct {
			PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

			Queue *string `tfsdk:"queue" yaml:"queue,omitempty"`

			Resources *map[string]string `tfsdk:"resources" yaml:"resources,omitempty"`
		} `tfsdk:"batch_scheduler_options" yaml:"batchSchedulerOptions,omitempty"`

		HadoopConf *map[string]string `tfsdk:"hadoop_conf" yaml:"hadoopConf,omitempty"`

		ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

		ImagePullSecrets *[]string `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

		SparkUIOptions *struct {
			IngressAnnotations *map[string]string `tfsdk:"ingress_annotations" yaml:"ingressAnnotations,omitempty"`

			IngressTLS *[]struct {
				Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"ingress_tls" yaml:"ingressTLS,omitempty"`

			ServiceAnnotations *map[string]string `tfsdk:"service_annotations" yaml:"serviceAnnotations,omitempty"`

			ServicePort *int64 `tfsdk:"service_port" yaml:"servicePort,omitempty"`

			ServicePortName *string `tfsdk:"service_port_name" yaml:"servicePortName,omitempty"`

			ServiceType *string `tfsdk:"service_type" yaml:"serviceType,omitempty"`
		} `tfsdk:"spark_ui_options" yaml:"sparkUIOptions,omitempty"`

		MemoryOverheadFactor *string `tfsdk:"memory_overhead_factor" yaml:"memoryOverheadFactor,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`

		Driver *struct {
			PodSecurityContext *struct {
				SeLinuxOptions *struct {
					Level *string `tfsdk:"level" yaml:"level,omitempty"`

					Role *string `tfsdk:"role" yaml:"role,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					User *string `tfsdk:"user" yaml:"user,omitempty"`
				} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

				SupplementalGroups *[]string `tfsdk:"supplemental_groups" yaml:"supplementalGroups,omitempty"`

				Sysctls *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"sysctls" yaml:"sysctls,omitempty"`

				WindowsOptions *struct {
					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

					RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`

					GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`
				} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

				FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

				RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

				RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

				RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
			} `tfsdk:"pod_security_context" yaml:"podSecurityContext,omitempty"`

			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Gpu *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Quantity *int64 `tfsdk:"quantity" yaml:"quantity,omitempty"`
			} `tfsdk:"gpu" yaml:"gpu,omitempty"`

			HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

			KubernetesMaster *string `tfsdk:"kubernetes_master" yaml:"kubernetesMaster,omitempty"`

			MemoryOverhead *string `tfsdk:"memory_overhead" yaml:"memoryOverhead,omitempty"`

			ServiceAnnotations *map[string]string `tfsdk:"service_annotations" yaml:"serviceAnnotations,omitempty"`

			ConfigMaps *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`
			} `tfsdk:"config_maps" yaml:"configMaps,omitempty"`

			EnvSecretKeyRefs *map[string]string `tfsdk:"env_secret_key_refs" yaml:"envSecretKeyRefs,omitempty"`

			ShareProcessNamespace *bool `tfsdk:"share_process_namespace" yaml:"shareProcessNamespace,omitempty"`

			Sidecars *[]struct {
				Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

				StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

				TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

				VolumeDevices *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`
				} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

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

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

				Lifecycle *struct {
					PreStop *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						HttpGet *struct {
							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`

					PostStart *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						HttpGet *struct {
							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"post_start" yaml:"postStart,omitempty"`
				} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

				ReadinessProbe *struct {
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`

					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					HttpGet *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				Resources *struct {
					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`

					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				StartupProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					HttpGet *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`
				} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

				TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

				WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`

				Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

				Env *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						SecretKeyRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

						ConfigMapKeyRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`
						} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

							FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

							Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

							Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"env" yaml:"env,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Ports *[]struct {
					HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

					HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

					ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				SecurityContext *struct {
					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

					Capabilities *struct {
						Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

						Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
					} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

					ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					WindowsOptions *struct {
						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

					Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

					ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					SeLinuxOptions *struct {
						Level *string `tfsdk:"level" yaml:"level,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

				Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

				VolumeMounts *[]struct {
					SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

					SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`

					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
				} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

				LivenessProbe *struct {
					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"sidecars" yaml:"sidecars,omitempty"`

			Tolerations *[]struct {
				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

			CoreLimit *string `tfsdk:"core_limit" yaml:"coreLimit,omitempty"`

			Cores *int64 `tfsdk:"cores" yaml:"cores,omitempty"`

			DnsConfig *struct {
				Nameservers *[]string `tfsdk:"nameservers" yaml:"nameservers,omitempty"`

				Options *[]struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"options" yaml:"options,omitempty"`

				Searches *[]string `tfsdk:"searches" yaml:"searches,omitempty"`
			} `tfsdk:"dns_config" yaml:"dnsConfig,omitempty"`

			EnvVars *map[string]string `tfsdk:"env_vars" yaml:"envVars,omitempty"`

			HostAliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

				Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
			} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

			Memory *string `tfsdk:"memory" yaml:"memory,omitempty"`

			CoreRequest *string `tfsdk:"core_request" yaml:"coreRequest,omitempty"`

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

			InitContainers *[]struct {
				Lifecycle *struct {
					PostStart *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						HttpGet *struct {
							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"post_start" yaml:"postStart,omitempty"`

					PreStop *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						HttpGet *struct {
							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
				} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

				StartupProbe *struct {
					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`
				} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

				Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

				VolumeDevices *[]struct {
					DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Env *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`
						} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

							FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

						ResourceFieldRef *struct {
							Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`

							ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

							Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`
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

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

				TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

				Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

				LivenessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					HttpGet *struct {
						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				ReadinessProbe *struct {
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				SecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

					Capabilities *struct {
						Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

						Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
					} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

					WindowsOptions *struct {
						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

					SeLinuxOptions *struct {
						Level *string `tfsdk:"level" yaml:"level,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

					Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

					ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

					ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

				Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

				VolumeMounts *[]struct {
					SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

					SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`

					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
				} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

				Ports *[]struct {
					ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

					HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

					HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

				TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

				WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`

			PodName *string `tfsdk:"pod_name" yaml:"podName,omitempty"`

			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

			Affinity *struct {
				PodAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						PodAffinityTerm *struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

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

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

				PodAntiAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						PodAffinityTerm *struct {
							TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

							Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`
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

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`

				NodeAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						Preference *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchFields *[]struct {
								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

								Key *string `tfsdk:"key" yaml:"key,omitempty"`
							} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
						} `tfsdk:"preference" yaml:"preference,omitempty"`

						Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
					} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

					RequiredDuringSchedulingIgnoredDuringExecution *struct {
						NodeSelectorTerms *[]struct {
							MatchExpressions *[]struct {
								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

								Key *string `tfsdk:"key" yaml:"key,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchFields *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
						} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`
			} `tfsdk:"affinity" yaml:"affinity,omitempty"`

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

						Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"env" yaml:"env,omitempty"`

			JavaOptions *string `tfsdk:"java_options" yaml:"javaOptions,omitempty"`

			Secrets *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				SecretType *string `tfsdk:"secret_type" yaml:"secretType,omitempty"`
			} `tfsdk:"secrets" yaml:"secrets,omitempty"`

			SecurityContext *struct {
				Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

				ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

				ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

				RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

				SeLinuxOptions *struct {
					User *string `tfsdk:"user" yaml:"user,omitempty"`

					Level *string `tfsdk:"level" yaml:"level,omitempty"`

					Role *string `tfsdk:"role" yaml:"role,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

				Capabilities *struct {
					Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`

					Add *[]string `tfsdk:"add" yaml:"add,omitempty"`
				} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

				RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

				RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

				WindowsOptions *struct {
					GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

					RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

				AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`
			} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			SchedulerName *string `tfsdk:"scheduler_name" yaml:"schedulerName,omitempty"`

			VolumeMounts *[]struct {
				MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

				SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`

				MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`
			} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			Lifecycle *struct {
				PostStart *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`

					HttpGet *struct {
						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
				} `tfsdk:"post_start" yaml:"postStart,omitempty"`

				PreStop *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`

					HttpGet *struct {
						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
				} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
			} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

			ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`
		} `tfsdk:"driver" yaml:"driver,omitempty"`

		NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		DynamicAllocation *struct {
			MinExecutors *int64 `tfsdk:"min_executors" yaml:"minExecutors,omitempty"`

			ShuffleTrackingTimeout *int64 `tfsdk:"shuffle_tracking_timeout" yaml:"shuffleTrackingTimeout,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			InitialExecutors *int64 `tfsdk:"initial_executors" yaml:"initialExecutors,omitempty"`

			MaxExecutors *int64 `tfsdk:"max_executors" yaml:"maxExecutors,omitempty"`
		} `tfsdk:"dynamic_allocation" yaml:"dynamicAllocation,omitempty"`

		RetryInterval *int64 `tfsdk:"retry_interval" yaml:"retryInterval,omitempty"`

		Volumes *[]struct {
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

						Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
				} `tfsdk:"items" yaml:"items,omitempty"`
			} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

			Cephfs *struct {
				User *string `tfsdk:"user" yaml:"user,omitempty"`

				Monitors *[]string `tfsdk:"monitors" yaml:"monitors,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SecretFile *string `tfsdk:"secret_file" yaml:"secretFile,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
			} `tfsdk:"cephfs" yaml:"cephfs,omitempty"`

			Cinder *struct {
				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
			} `tfsdk:"cinder" yaml:"cinder,omitempty"`

			GcePersistentDisk *struct {
				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

				PdName *string `tfsdk:"pd_name" yaml:"pdName,omitempty"`
			} `tfsdk:"gce_persistent_disk" yaml:"gcePersistentDisk,omitempty"`

			Iscsi *struct {
				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				InitiatorName *string `tfsdk:"initiator_name" yaml:"initiatorName,omitempty"`

				Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				TargetPortal *string `tfsdk:"target_portal" yaml:"targetPortal,omitempty"`

				ChapAuthDiscovery *bool `tfsdk:"chap_auth_discovery" yaml:"chapAuthDiscovery,omitempty"`

				ChapAuthSession *bool `tfsdk:"chap_auth_session" yaml:"chapAuthSession,omitempty"`

				Portals *[]string `tfsdk:"portals" yaml:"portals,omitempty"`

				Iqn *string `tfsdk:"iqn" yaml:"iqn,omitempty"`

				IscsiInterface *string `tfsdk:"iscsi_interface" yaml:"iscsiInterface,omitempty"`
			} `tfsdk:"iscsi" yaml:"iscsi,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

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

			Fc *struct {
				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				TargetWWNs *[]string `tfsdk:"target_ww_ns" yaml:"targetWWNs,omitempty"`

				Wwids *[]string `tfsdk:"wwids" yaml:"wwids,omitempty"`
			} `tfsdk:"fc" yaml:"fc,omitempty"`

			GitRepo *struct {
				Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`

				Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`

				Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`
			} `tfsdk:"git_repo" yaml:"gitRepo,omitempty"`

			HostPath *struct {
				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"host_path" yaml:"hostPath,omitempty"`

			ScaleIO *struct {
				ProtectionDomain *string `tfsdk:"protection_domain" yaml:"protectionDomain,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				SslEnabled *bool `tfsdk:"ssl_enabled" yaml:"sslEnabled,omitempty"`

				System *string `tfsdk:"system" yaml:"system,omitempty"`

				VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`

				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				Gateway *string `tfsdk:"gateway" yaml:"gateway,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				StorageMode *string `tfsdk:"storage_mode" yaml:"storageMode,omitempty"`

				StoragePool *string `tfsdk:"storage_pool" yaml:"storagePool,omitempty"`
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

			AzureFile *struct {
				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

				ShareName *string `tfsdk:"share_name" yaml:"shareName,omitempty"`
			} `tfsdk:"azure_file" yaml:"azureFile,omitempty"`

			Csi *struct {
				Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				NodePublishSecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"node_publish_secret_ref" yaml:"nodePublishSecretRef,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				VolumeAttributes *map[string]string `tfsdk:"volume_attributes" yaml:"volumeAttributes,omitempty"`
			} `tfsdk:"csi" yaml:"csi,omitempty"`

			Flocker *struct {
				DatasetName *string `tfsdk:"dataset_name" yaml:"datasetName,omitempty"`

				DatasetUUID *string `tfsdk:"dataset_uuid" yaml:"datasetUUID,omitempty"`
			} `tfsdk:"flocker" yaml:"flocker,omitempty"`

			Nfs *struct {
				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				Server *string `tfsdk:"server" yaml:"server,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`
			} `tfsdk:"nfs" yaml:"nfs,omitempty"`

			Projected *struct {
				Sources *[]struct {
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

								Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

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

					ConfigMap *struct {
						Items *[]struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`
						} `tfsdk:"items" yaml:"items,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`
				} `tfsdk:"sources" yaml:"sources,omitempty"`

				DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`
			} `tfsdk:"projected" yaml:"projected,omitempty"`

			AwsElasticBlockStore *struct {
				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
			} `tfsdk:"aws_elastic_block_store" yaml:"awsElasticBlockStore,omitempty"`

			AzureDisk *struct {
				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				CachingMode *string `tfsdk:"caching_mode" yaml:"cachingMode,omitempty"`

				DiskName *string `tfsdk:"disk_name" yaml:"diskName,omitempty"`

				DiskURI *string `tfsdk:"disk_uri" yaml:"diskURI,omitempty"`
			} `tfsdk:"azure_disk" yaml:"azureDisk,omitempty"`

			Glusterfs *struct {
				Endpoints *string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
			} `tfsdk:"glusterfs" yaml:"glusterfs,omitempty"`

			PersistentVolumeClaim *struct {
				ClaimName *string `tfsdk:"claim_name" yaml:"claimName,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
			} `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

			VsphereVolume *struct {
				VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`

				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				StoragePolicyID *string `tfsdk:"storage_policy_id" yaml:"storagePolicyID,omitempty"`

				StoragePolicyName *string `tfsdk:"storage_policy_name" yaml:"storagePolicyName,omitempty"`
			} `tfsdk:"vsphere_volume" yaml:"vsphereVolume,omitempty"`

			ConfigMap *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

				DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

				Items *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"items" yaml:"items,omitempty"`
			} `tfsdk:"config_map" yaml:"configMap,omitempty"`

			EmptyDir *struct {
				Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

				SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
			} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

			PortworxVolume *struct {
				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
			} `tfsdk:"portworx_volume" yaml:"portworxVolume,omitempty"`

			Quobyte *struct {
				Volume *string `tfsdk:"volume" yaml:"volume,omitempty"`

				Group *string `tfsdk:"group" yaml:"group,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				Registry *string `tfsdk:"registry" yaml:"registry,omitempty"`

				Tenant *string `tfsdk:"tenant" yaml:"tenant,omitempty"`

				User *string `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"quobyte" yaml:"quobyte,omitempty"`

			FlexVolume *struct {
				Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

				Options *map[string]string `tfsdk:"options" yaml:"options,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
			} `tfsdk:"flex_volume" yaml:"flexVolume,omitempty"`

			PhotonPersistentDisk *struct {
				PdID *string `tfsdk:"pd_id" yaml:"pdID,omitempty"`

				FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`
			} `tfsdk:"photon_persistent_disk" yaml:"photonPersistentDisk,omitempty"`
		} `tfsdk:"volumes" yaml:"volumes,omitempty"`

		Arguments *[]string `tfsdk:"arguments" yaml:"arguments,omitempty"`

		BatchScheduler *string `tfsdk:"batch_scheduler" yaml:"batchScheduler,omitempty"`

		Executor *struct {
			ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

			SchedulerName *string `tfsdk:"scheduler_name" yaml:"schedulerName,omitempty"`

			Secrets *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				SecretType *string `tfsdk:"secret_type" yaml:"secretType,omitempty"`
			} `tfsdk:"secrets" yaml:"secrets,omitempty"`

			DeleteOnTermination *bool `tfsdk:"delete_on_termination" yaml:"deleteOnTermination,omitempty"`

			DnsConfig *struct {
				Searches *[]string `tfsdk:"searches" yaml:"searches,omitempty"`

				Nameservers *[]string `tfsdk:"nameservers" yaml:"nameservers,omitempty"`

				Options *[]struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"options" yaml:"options,omitempty"`
			} `tfsdk:"dns_config" yaml:"dnsConfig,omitempty"`

			EnvFrom *[]struct {
				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				SecretRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

				ConfigMapRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`
			} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

			Gpu *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Quantity *int64 `tfsdk:"quantity" yaml:"quantity,omitempty"`
			} `tfsdk:"gpu" yaml:"gpu,omitempty"`

			CoreLimit *string `tfsdk:"core_limit" yaml:"coreLimit,omitempty"`

			EnvSecretKeyRefs *map[string]string `tfsdk:"env_secret_key_refs" yaml:"envSecretKeyRefs,omitempty"`

			EnvVars *map[string]string `tfsdk:"env_vars" yaml:"envVars,omitempty"`

			InitContainers *[]struct {
				TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

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

				StartupProbe *struct {
					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`

					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					HttpGet *struct {
						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`
				} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

				Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

				TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				SecurityContext *struct {
					Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

					ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

					ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					SeLinuxOptions *struct {
						Level *string `tfsdk:"level" yaml:"level,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

					WindowsOptions *struct {
						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`

						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					Capabilities *struct {
						Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

						Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
					} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

				WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`

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

							Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

							Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"env" yaml:"env,omitempty"`

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

				Lifecycle *struct {
					PreStop *struct {
						HttpGet *struct {
							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`
					} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`

					PostStart *struct {
						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						HttpGet *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"post_start" yaml:"postStart,omitempty"`
				} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

				LivenessProbe *struct {
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`

					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					HttpGet *struct {
						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				Ports *[]struct {
					ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

					HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

					HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				ReadinessProbe *struct {
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

				Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			MemoryOverhead *string `tfsdk:"memory_overhead" yaml:"memoryOverhead,omitempty"`

			SecurityContext *struct {
				SeLinuxOptions *struct {
					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					User *string `tfsdk:"user" yaml:"user,omitempty"`

					Level *string `tfsdk:"level" yaml:"level,omitempty"`

					Role *string `tfsdk:"role" yaml:"role,omitempty"`
				} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

				WindowsOptions *struct {
					GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

					RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

				AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

				Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

				ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

				RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

				Capabilities *struct {
					Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

					Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
				} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

				ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

				RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

				RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
			} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

			ConfigMaps *[]struct {
				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"config_maps" yaml:"configMaps,omitempty"`

			Cores *int64 `tfsdk:"cores" yaml:"cores,omitempty"`

			HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			JavaOptions *string `tfsdk:"java_options" yaml:"javaOptions,omitempty"`

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
									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

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

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

				PodAntiAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						PodAffinityTerm *struct {
							Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

							TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`
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

						Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
			} `tfsdk:"affinity" yaml:"affinity,omitempty"`

			Env *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				ValueFrom *struct {
					ResourceFieldRef *struct {
						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`

						ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

						Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`

					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					FieldRef *struct {
						FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`
					} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"env" yaml:"env,omitempty"`

			ShareProcessNamespace *bool `tfsdk:"share_process_namespace" yaml:"shareProcessNamespace,omitempty"`

			HostAliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

				Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
			} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

			Instances *int64 `tfsdk:"instances" yaml:"instances,omitempty"`

			Sidecars *[]struct {
				StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

				Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

				Ports *[]struct {
					HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

					HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

					ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				VolumeDevices *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`
				} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

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

							Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

							Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

						SecretKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"env" yaml:"env,omitempty"`

				SecurityContext *struct {
					ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

					ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					SeLinuxOptions *struct {
						Level *string `tfsdk:"level" yaml:"level,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

					WindowsOptions *struct {
						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

					Capabilities *struct {
						Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

						Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
					} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

				TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

				ReadinessProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`

					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					HttpGet *struct {
						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				StartupProbe *struct {
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`
				} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

				Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

				VolumeMounts *[]struct {
					MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

					SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

					SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`

					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`
				} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

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

							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"post_start" yaml:"postStart,omitempty"`

					PreStop *struct {
						HttpGet *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`
					} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
				} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

				LivenessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					HttpGet *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						HttpHeaders *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

					Exec *struct {
						Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
					} `tfsdk:"exec" yaml:"exec,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`

				TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

				Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"sidecars" yaml:"sidecars,omitempty"`

			Tolerations *[]struct {
				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`

				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`
			} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

			VolumeMounts *[]struct {
				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

				SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`

				MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

				MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

			CoreRequest *string `tfsdk:"core_request" yaml:"coreRequest,omitempty"`

			Memory *string `tfsdk:"memory" yaml:"memory,omitempty"`

			PodSecurityContext *struct {
				SeLinuxOptions *struct {
					User *string `tfsdk:"user" yaml:"user,omitempty"`

					Level *string `tfsdk:"level" yaml:"level,omitempty"`

					Role *string `tfsdk:"role" yaml:"role,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

				SupplementalGroups *[]string `tfsdk:"supplemental_groups" yaml:"supplementalGroups,omitempty"`

				Sysctls *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"sysctls" yaml:"sysctls,omitempty"`

				WindowsOptions *struct {
					GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

					GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

					RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

				FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

				RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

				RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

				RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
			} `tfsdk:"pod_security_context" yaml:"podSecurityContext,omitempty"`
		} `tfsdk:"executor" yaml:"executor,omitempty"`

		FailureRetries *int64 `tfsdk:"failure_retries" yaml:"failureRetries,omitempty"`

		MainClass *string `tfsdk:"main_class" yaml:"mainClass,omitempty"`

		Monitoring *struct {
			ExposeDriverMetrics *bool `tfsdk:"expose_driver_metrics" yaml:"exposeDriverMetrics,omitempty"`

			ExposeExecutorMetrics *bool `tfsdk:"expose_executor_metrics" yaml:"exposeExecutorMetrics,omitempty"`

			MetricsProperties *string `tfsdk:"metrics_properties" yaml:"metricsProperties,omitempty"`

			MetricsPropertiesFile *string `tfsdk:"metrics_properties_file" yaml:"metricsPropertiesFile,omitempty"`

			Prometheus *struct {
				ConfigFile *string `tfsdk:"config_file" yaml:"configFile,omitempty"`

				Configuration *string `tfsdk:"configuration" yaml:"configuration,omitempty"`

				JmxExporterJar *string `tfsdk:"jmx_exporter_jar" yaml:"jmxExporterJar,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				PortName *string `tfsdk:"port_name" yaml:"portName,omitempty"`
			} `tfsdk:"prometheus" yaml:"prometheus,omitempty"`
		} `tfsdk:"monitoring" yaml:"monitoring,omitempty"`

		ProxyUser *string `tfsdk:"proxy_user" yaml:"proxyUser,omitempty"`

		RestartPolicy *struct {
			OnSubmissionFailureRetries *int64 `tfsdk:"on_submission_failure_retries" yaml:"onSubmissionFailureRetries,omitempty"`

			OnSubmissionFailureRetryInterval *int64 `tfsdk:"on_submission_failure_retry_interval" yaml:"onSubmissionFailureRetryInterval,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			OnFailureRetries *int64 `tfsdk:"on_failure_retries" yaml:"onFailureRetries,omitempty"`

			OnFailureRetryInterval *int64 `tfsdk:"on_failure_retry_interval" yaml:"onFailureRetryInterval,omitempty"`
		} `tfsdk:"restart_policy" yaml:"restartPolicy,omitempty"`

		TimeToLiveSeconds *int64 `tfsdk:"time_to_live_seconds" yaml:"timeToLiveSeconds,omitempty"`

		Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

		PythonVersion *string `tfsdk:"python_version" yaml:"pythonVersion,omitempty"`

		SparkConfigMap *string `tfsdk:"spark_config_map" yaml:"sparkConfigMap,omitempty"`

		SparkVersion *string `tfsdk:"spark_version" yaml:"sparkVersion,omitempty"`

		HadoopConfigMap *string `tfsdk:"hadoop_config_map" yaml:"hadoopConfigMap,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		SparkConf *map[string]string `tfsdk:"spark_conf" yaml:"sparkConf,omitempty"`

		Deps *struct {
			Files *[]string `tfsdk:"files" yaml:"files,omitempty"`

			Jars *[]string `tfsdk:"jars" yaml:"jars,omitempty"`

			Packages *[]string `tfsdk:"packages" yaml:"packages,omitempty"`

			PyFiles *[]string `tfsdk:"py_files" yaml:"pyFiles,omitempty"`

			Repositories *[]string `tfsdk:"repositories" yaml:"repositories,omitempty"`

			ExcludePackages *[]string `tfsdk:"exclude_packages" yaml:"excludePackages,omitempty"`
		} `tfsdk:"deps" yaml:"deps,omitempty"`

		MainApplicationFile *string `tfsdk:"main_application_file" yaml:"mainApplicationFile,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSparkoperatorK8SIoSparkApplicationV1Beta2Resource() resource.Resource {
	return &SparkoperatorK8SIoSparkApplicationV1Beta2Resource{}
}

func (r *SparkoperatorK8SIoSparkApplicationV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sparkoperator_k8s_io_spark_application_v1beta2"
}

func (r *SparkoperatorK8SIoSparkApplicationV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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

					"batch_scheduler_options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

							"resources": {
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

					"hadoop_conf": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

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

					"image_pull_secrets": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"spark_ui_options": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ingress_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress_tls": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"hosts": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

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

							"service_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_port_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_type": {
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

					"memory_overhead_factor": {
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

					"driver": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"pod_security_context": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"gmsa_credential_spec_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"gmsa_credential_spec": {
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

									"fs_group": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gpu": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"quantity": {
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

							"host_network": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kubernetes_master": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"memory_overhead": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_maps": {
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

							"env_secret_key_refs": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

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

							"sidecars": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

									"volume_devices": {
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

											"device_path": {
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

																Type: types.StringType,

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

															"host": {
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

															"scheme": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

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

									"readiness_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

															"value": {
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

														Type: types.StringType,

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"requests": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"limits": {
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

									"startup_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

														Type: types.StringType,

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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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
										}),

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

									"working_dir": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"args": {
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

													"secret_key_ref": {
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

															"key": {
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

													"config_map_key_ref": {
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

															"key": {
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

																Type: types.StringType,

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

									"image": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

												Required: true,
												Optional: false,
												Computed: false,
											},

											"container_port": {
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

									"security_context": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"run_as_user": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

											"proc_mount": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

											"privileged": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

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
										}),

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

									"volume_mounts": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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
										}),

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

									"liveness_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

														Type: types.StringType,

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

							"core_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cores": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
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

											"value": {
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

							"env_vars": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

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

							"memory": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"core_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

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

							"init_containers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

															"scheme": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

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

																	"value": {
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

															"port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"host": {
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

									"startup_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

														Type: types.StringType,

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

									"image": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

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

															"key": {
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

															"resource": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

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

									"image_pull_policy": {
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

									"args": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

														Type: types.StringType,

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

													"host": {
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

													"port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host": {
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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

														Type: types.StringType,

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
										}),

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

									"volume_mounts": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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
										}),

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

									"working_dir": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_name": {
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

							"affinity": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																			"values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

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

															"topology_key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

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

															"namespaces": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

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

																	"key": {
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

																	"key": {
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
								}),

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

														Type: types.StringType,

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

							"java_options": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secrets": {
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

									"path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_type": {
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

							"security_context": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

									"run_as_non_root": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"se_linux_options": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"user": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"capabilities": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"drop": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"add": {
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

									"run_as_group": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

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

									"allow_privilege_escalation": {
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

							"node_selector": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

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

							"volume_mounts": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

									"mount_path": {
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

							"image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

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

													"port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

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

														Type: types.StringType,

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

							"service_account": {
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

					"node_selector": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dynamic_allocation": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"min_executors": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"shuffle_tracking_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_executors": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_executors": {
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

					"retry_interval": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volumes": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

														Type: types.StringType,

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

							"cephfs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"user": {
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

							"gce_persistent_disk": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"iscsi": {
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

									"initiator_name": {
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

									"portals": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

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

							"git_repo": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"revision": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

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

							"scale_io": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"protection_domain": {
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

									"read_only": {
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

							"nfs": {
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

									"server": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"projected": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"sources": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

																		Type: types.StringType,

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

											"config_map": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"items": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

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
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"default_mode": {
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

							"vsphere_volume": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"volume_path": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_map": {
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

							"quobyte": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"volume": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

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

							"photon_persistent_disk": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"pd_id": {
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

					"arguments": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"batch_scheduler": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"executor": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"service_account": {
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

							"scheduler_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secrets": {
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

									"path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_type": {
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

							"delete_on_termination": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"dns_config": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"searches": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

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

											"value": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gpu": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"quantity": {
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

							"core_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env_secret_key_refs": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env_vars": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"init_containers": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"termination_message_path": {
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

									"startup_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

													"port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

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

									"termination_message_policy": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

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

											"windows_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"run_as_user_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"allow_privilege_escalation": {
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

																Type: types.StringType,

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

											"pre_stop": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"http_get": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

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

																	"value": {
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

																Type: types.StringType,

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

									"liveness_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

													"scheme": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

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

												Required: true,
												Optional: false,
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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

														Type: types.StringType,

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
										}),

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

									"image": {
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
								}),

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

							"memory_overhead": {
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

									"se_linux_options": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

									"allow_privilege_escalation": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

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

									"proc_mount": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_maps": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"path": {
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

							"cores": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

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

							"node_selector": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

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

							"java_options": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

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

																			"values": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

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

											"resource_field_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"resource": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

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

													"field_path": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"api_version": {
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

							"share_process_namespace": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

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

							"instances": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sidecars": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"stdin_once": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

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

									"ports": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

												Required: true,
												Optional: false,
												Computed: false,
											},

											"container_port": {
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

									"volume_devices": {
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

											"device_path": {
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

																Type: types.StringType,

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

									"security_context": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"proc_mount": {
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

											"privileged": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

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
										}),

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

													"scheme": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

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

									"startup_probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host": {
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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

														Type: types.StringType,

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

									"volume_mounts": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

											"mount_path": {
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

																Type: types.StringType,

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

											"pre_stop": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																Type: types.StringType,

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

														Type: types.StringType,

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

											"timeout_seconds": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

							"tolerations": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_mounts": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"core_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"memory": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_security_context": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"se_linux_options": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"user": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

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

									"fs_group": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

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

					"failure_retries": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"main_class": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"expose_driver_metrics": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"expose_executor_metrics": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"metrics_properties": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics_properties_file": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"prometheus": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_file": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jmx_exporter_jar": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port_name": {
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

					"proxy_user": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"restart_policy": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"on_submission_failure_retries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"on_submission_failure_retry_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

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

							"on_failure_retries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"on_failure_retry_interval": {
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

					"time_to_live_seconds": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mode": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"python_version": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"spark_config_map": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"spark_version": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"hadoop_config_map": {
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

						Required: false,
						Optional: true,
						Computed: false,
					},

					"spark_conf": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"deps": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"files": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jars": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"packages": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"py_files": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"repositories": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"exclude_packages": {
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

					"main_application_file": {
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
		},
	}, nil
}

func (r *SparkoperatorK8SIoSparkApplicationV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_sparkoperator_k8s_io_spark_application_v1beta2")

	var state SparkoperatorK8SIoSparkApplicationV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SparkoperatorK8SIoSparkApplicationV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sparkoperator.k8s.io/v1beta2")
	goModel.Kind = utilities.Ptr("SparkApplication")

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

func (r *SparkoperatorK8SIoSparkApplicationV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sparkoperator_k8s_io_spark_application_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *SparkoperatorK8SIoSparkApplicationV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_sparkoperator_k8s_io_spark_application_v1beta2")

	var state SparkoperatorK8SIoSparkApplicationV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SparkoperatorK8SIoSparkApplicationV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("sparkoperator.k8s.io/v1beta2")
	goModel.Kind = utilities.Ptr("SparkApplication")

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

func (r *SparkoperatorK8SIoSparkApplicationV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_sparkoperator_k8s_io_spark_application_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
