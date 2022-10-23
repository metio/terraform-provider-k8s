/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type CamelApacheOrgIntegrationV1Resource struct{}

var (
	_ resource.Resource = (*CamelApacheOrgIntegrationV1Resource)(nil)
)

type CamelApacheOrgIntegrationV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CamelApacheOrgIntegrationV1GoModel struct {
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
		Configuration *[]struct {
			ResourceKey *string `tfsdk:"resource_key" yaml:"resourceKey,omitempty"`

			ResourceMountPoint *string `tfsdk:"resource_mount_point" yaml:"resourceMountPoint,omitempty"`

			ResourceType *string `tfsdk:"resource_type" yaml:"resourceType,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"configuration" yaml:"configuration,omitempty"`

		Dependencies *[]string `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

		Flows *[]map[string]string `tfsdk:"flows" yaml:"flows,omitempty"`

		IntegrationKit *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

			Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
		} `tfsdk:"integration_kit" yaml:"integrationKit,omitempty"`

		Profile *string `tfsdk:"profile" yaml:"profile,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		Repositories *[]string `tfsdk:"repositories" yaml:"repositories,omitempty"`

		Resources *[]struct {
			Compression *bool `tfsdk:"compression" yaml:"compression,omitempty"`

			Content *string `tfsdk:"content" yaml:"content,omitempty"`

			ContentKey *string `tfsdk:"content_key" yaml:"contentKey,omitempty"`

			ContentRef *string `tfsdk:"content_ref" yaml:"contentRef,omitempty"`

			ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

			MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			RawContent *string `tfsdk:"raw_content" yaml:"rawContent,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

		Sources *[]struct {
			Compression *bool `tfsdk:"compression" yaml:"compression,omitempty"`

			Content *string `tfsdk:"content" yaml:"content,omitempty"`

			ContentKey *string `tfsdk:"content_key" yaml:"contentKey,omitempty"`

			ContentRef *string `tfsdk:"content_ref" yaml:"contentRef,omitempty"`

			ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

			Interceptors *[]string `tfsdk:"interceptors" yaml:"interceptors,omitempty"`

			Language *string `tfsdk:"language" yaml:"language,omitempty"`

			Loader *string `tfsdk:"loader" yaml:"loader,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Property_names *[]string `tfsdk:"property_names" yaml:"property-names,omitempty"`

			RawContent *string `tfsdk:"raw_content" yaml:"rawContent,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"sources" yaml:"sources,omitempty"`

		Template *struct {
			Spec *struct {
				ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" yaml:"activeDeadlineSeconds,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

				DnsPolicy *string `tfsdk:"dns_policy" yaml:"dnsPolicy,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

						Grpc *struct {
							Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

							Service *string `tfsdk:"service" yaml:"service,omitempty"`
						} `tfsdk:"grpc" yaml:"grpc,omitempty"`

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

				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				RestartPolicy *string `tfsdk:"restart_policy" yaml:"restartPolicy,omitempty"`

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

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

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
							Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

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

		Traits *struct {
			Threescale *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"threescale" yaml:"3scale,omitempty"`

			Addons utilities.Dynamic `tfsdk:"addons" yaml:"addons,omitempty"`

			Affinity *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				NodeAffinityLabels *[]string `tfsdk:"node_affinity_labels" yaml:"nodeAffinityLabels,omitempty"`

				PodAffinity *bool `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

				PodAffinityLabels *[]string `tfsdk:"pod_affinity_labels" yaml:"podAffinityLabels,omitempty"`

				PodAntiAffinity *bool `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`

				PodAntiAffinityLabels *[]string `tfsdk:"pod_anti_affinity_labels" yaml:"podAntiAffinityLabels,omitempty"`
			} `tfsdk:"affinity" yaml:"affinity,omitempty"`

			Builder *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Properties *[]string `tfsdk:"properties" yaml:"properties,omitempty"`

				Verbose *bool `tfsdk:"verbose" yaml:"verbose,omitempty"`
			} `tfsdk:"builder" yaml:"builder,omitempty"`

			Camel *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Properties *[]string `tfsdk:"properties" yaml:"properties,omitempty"`

				RuntimeVersion *string `tfsdk:"runtime_version" yaml:"runtimeVersion,omitempty"`
			} `tfsdk:"camel" yaml:"camel,omitempty"`

			Container *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Expose *bool `tfsdk:"expose" yaml:"expose,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

				LimitCPU *string `tfsdk:"limit_cpu" yaml:"limitCPU,omitempty"`

				LimitMemory *string `tfsdk:"limit_memory" yaml:"limitMemory,omitempty"`

				LivenessFailureThreshold *int64 `tfsdk:"liveness_failure_threshold" yaml:"livenessFailureThreshold,omitempty"`

				LivenessInitialDelay *int64 `tfsdk:"liveness_initial_delay" yaml:"livenessInitialDelay,omitempty"`

				LivenessPeriod *int64 `tfsdk:"liveness_period" yaml:"livenessPeriod,omitempty"`

				LivenessScheme *string `tfsdk:"liveness_scheme" yaml:"livenessScheme,omitempty"`

				LivenessSuccessThreshold *int64 `tfsdk:"liveness_success_threshold" yaml:"livenessSuccessThreshold,omitempty"`

				LivenessTimeout *int64 `tfsdk:"liveness_timeout" yaml:"livenessTimeout,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				PortName *string `tfsdk:"port_name" yaml:"portName,omitempty"`

				ProbesEnabled *bool `tfsdk:"probes_enabled" yaml:"probesEnabled,omitempty"`

				ReadinessFailureThreshold *int64 `tfsdk:"readiness_failure_threshold" yaml:"readinessFailureThreshold,omitempty"`

				ReadinessInitialDelay *int64 `tfsdk:"readiness_initial_delay" yaml:"readinessInitialDelay,omitempty"`

				ReadinessPeriod *int64 `tfsdk:"readiness_period" yaml:"readinessPeriod,omitempty"`

				ReadinessScheme *string `tfsdk:"readiness_scheme" yaml:"readinessScheme,omitempty"`

				ReadinessSuccessThreshold *int64 `tfsdk:"readiness_success_threshold" yaml:"readinessSuccessThreshold,omitempty"`

				ReadinessTimeout *int64 `tfsdk:"readiness_timeout" yaml:"readinessTimeout,omitempty"`

				RequestCPU *string `tfsdk:"request_cpu" yaml:"requestCPU,omitempty"`

				RequestMemory *string `tfsdk:"request_memory" yaml:"requestMemory,omitempty"`

				ServicePort *int64 `tfsdk:"service_port" yaml:"servicePort,omitempty"`

				ServicePortName *string `tfsdk:"service_port_name" yaml:"servicePortName,omitempty"`
			} `tfsdk:"container" yaml:"container,omitempty"`

			Cron *struct {
				ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" yaml:"activeDeadlineSeconds,omitempty"`

				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				BackoffLimit *int64 `tfsdk:"backoff_limit" yaml:"backoffLimit,omitempty"`

				Components *string `tfsdk:"components" yaml:"components,omitempty"`

				ConcurrencyPolicy *string `tfsdk:"concurrency_policy" yaml:"concurrencyPolicy,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Fallback *bool `tfsdk:"fallback" yaml:"fallback,omitempty"`

				Schedule *string `tfsdk:"schedule" yaml:"schedule,omitempty"`

				StartingDeadlineSeconds *int64 `tfsdk:"starting_deadline_seconds" yaml:"startingDeadlineSeconds,omitempty"`
			} `tfsdk:"cron" yaml:"cron,omitempty"`

			Dependencies *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"dependencies" yaml:"dependencies,omitempty"`

			Deployer *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				UseSSA *bool `tfsdk:"use_ssa" yaml:"useSSA,omitempty"`
			} `tfsdk:"deployer" yaml:"deployer,omitempty"`

			Deployment *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ProgressDeadlineSeconds *int64 `tfsdk:"progress_deadline_seconds" yaml:"progressDeadlineSeconds,omitempty"`

				RollingUpdateMaxSurge *int64 `tfsdk:"rolling_update_max_surge" yaml:"rollingUpdateMaxSurge,omitempty"`

				RollingUpdateMaxUnavailable *int64 `tfsdk:"rolling_update_max_unavailable" yaml:"rollingUpdateMaxUnavailable,omitempty"`

				Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`
			} `tfsdk:"deployment" yaml:"deployment,omitempty"`

			Environment *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				ContainerMeta *bool `tfsdk:"container_meta" yaml:"containerMeta,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				HttpProxy *bool `tfsdk:"http_proxy" yaml:"httpProxy,omitempty"`

				Vars *[]string `tfsdk:"vars" yaml:"vars,omitempty"`
			} `tfsdk:"environment" yaml:"environment,omitempty"`

			Error_handler *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Ref *string `tfsdk:"ref" yaml:"ref,omitempty"`
			} `tfsdk:"error_handler" yaml:"error-handler,omitempty"`

			Gc *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				DiscoveryCache *string `tfsdk:"discovery_cache" yaml:"discoveryCache,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"gc" yaml:"gc,omitempty"`

			Health *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				LivenessFailureThreshold *int64 `tfsdk:"liveness_failure_threshold" yaml:"livenessFailureThreshold,omitempty"`

				LivenessInitialDelay *int64 `tfsdk:"liveness_initial_delay" yaml:"livenessInitialDelay,omitempty"`

				LivenessPeriod *int64 `tfsdk:"liveness_period" yaml:"livenessPeriod,omitempty"`

				LivenessProbeEnabled *bool `tfsdk:"liveness_probe_enabled" yaml:"livenessProbeEnabled,omitempty"`

				LivenessScheme *string `tfsdk:"liveness_scheme" yaml:"livenessScheme,omitempty"`

				LivenessSuccessThreshold *int64 `tfsdk:"liveness_success_threshold" yaml:"livenessSuccessThreshold,omitempty"`

				LivenessTimeout *int64 `tfsdk:"liveness_timeout" yaml:"livenessTimeout,omitempty"`

				ReadinessFailureThreshold *int64 `tfsdk:"readiness_failure_threshold" yaml:"readinessFailureThreshold,omitempty"`

				ReadinessInitialDelay *int64 `tfsdk:"readiness_initial_delay" yaml:"readinessInitialDelay,omitempty"`

				ReadinessPeriod *int64 `tfsdk:"readiness_period" yaml:"readinessPeriod,omitempty"`

				ReadinessProbeEnabled *bool `tfsdk:"readiness_probe_enabled" yaml:"readinessProbeEnabled,omitempty"`

				ReadinessScheme *string `tfsdk:"readiness_scheme" yaml:"readinessScheme,omitempty"`

				ReadinessSuccessThreshold *int64 `tfsdk:"readiness_success_threshold" yaml:"readinessSuccessThreshold,omitempty"`

				ReadinessTimeout *int64 `tfsdk:"readiness_timeout" yaml:"readinessTimeout,omitempty"`
			} `tfsdk:"health" yaml:"health,omitempty"`

			Ingress *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`

			Istio *struct {
				Allow *string `tfsdk:"allow" yaml:"allow,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Inject *bool `tfsdk:"inject" yaml:"inject,omitempty"`
			} `tfsdk:"istio" yaml:"istio,omitempty"`

			Jolokia *struct {
				CACert *string `tfsdk:"ca_cert" yaml:"CACert,omitempty"`

				ClientPrincipal *[]string `tfsdk:"client_principal" yaml:"clientPrincipal,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				DiscoveryEnabled *bool `tfsdk:"discovery_enabled" yaml:"discoveryEnabled,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ExtendedClientCheck *bool `tfsdk:"extended_client_check" yaml:"extendedClientCheck,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

				Password *string `tfsdk:"password" yaml:"password,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

				UseSSLClientAuthentication *bool `tfsdk:"use_ssl_client_authentication" yaml:"useSSLClientAuthentication,omitempty"`

				User *string `tfsdk:"user" yaml:"user,omitempty"`
			} `tfsdk:"jolokia" yaml:"jolokia,omitempty"`

			Jvm *struct {
				Classpath *string `tfsdk:"classpath" yaml:"classpath,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Debug *bool `tfsdk:"debug" yaml:"debug,omitempty"`

				DebugAddress *string `tfsdk:"debug_address" yaml:"debugAddress,omitempty"`

				DebugSuspend *bool `tfsdk:"debug_suspend" yaml:"debugSuspend,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Options *[]string `tfsdk:"options" yaml:"options,omitempty"`

				PrintCommand *bool `tfsdk:"print_command" yaml:"printCommand,omitempty"`
			} `tfsdk:"jvm" yaml:"jvm,omitempty"`

			Kamelets *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				List *string `tfsdk:"list" yaml:"list,omitempty"`
			} `tfsdk:"kamelets" yaml:"kamelets,omitempty"`

			Keda *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"keda" yaml:"keda,omitempty"`

			Knative *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				ChannelSinks *[]string `tfsdk:"channel_sinks" yaml:"channelSinks,omitempty"`

				ChannelSources *[]string `tfsdk:"channel_sources" yaml:"channelSources,omitempty"`

				Config *string `tfsdk:"config" yaml:"config,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				EndpointSinks *[]string `tfsdk:"endpoint_sinks" yaml:"endpointSinks,omitempty"`

				EndpointSources *[]string `tfsdk:"endpoint_sources" yaml:"endpointSources,omitempty"`

				EventSinks *[]string `tfsdk:"event_sinks" yaml:"eventSinks,omitempty"`

				EventSources *[]string `tfsdk:"event_sources" yaml:"eventSources,omitempty"`

				FilterSourceChannels *bool `tfsdk:"filter_source_channels" yaml:"filterSourceChannels,omitempty"`

				SinkBinding *bool `tfsdk:"sink_binding" yaml:"sinkBinding,omitempty"`
			} `tfsdk:"knative" yaml:"knative,omitempty"`

			Knative_service *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				AutoscalingMetric *string `tfsdk:"autoscaling_metric" yaml:"autoscalingMetric,omitempty"`

				AutoscalingTarget *int64 `tfsdk:"autoscaling_target" yaml:"autoscalingTarget,omitempty"`

				Class *string `tfsdk:"class" yaml:"class,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				MaxScale *int64 `tfsdk:"max_scale" yaml:"maxScale,omitempty"`

				MinScale *int64 `tfsdk:"min_scale" yaml:"minScale,omitempty"`

				RolloutDuration *string `tfsdk:"rollout_duration" yaml:"rolloutDuration,omitempty"`

				Visibility *string `tfsdk:"visibility" yaml:"visibility,omitempty"`
			} `tfsdk:"knative_service" yaml:"knative-service,omitempty"`

			Logging *struct {
				Color *bool `tfsdk:"color" yaml:"color,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Format *string `tfsdk:"format" yaml:"format,omitempty"`

				Json *bool `tfsdk:"json" yaml:"json,omitempty"`

				JsonPrettyPrint *bool `tfsdk:"json_pretty_print" yaml:"jsonPrettyPrint,omitempty"`

				Level *string `tfsdk:"level" yaml:"level,omitempty"`
			} `tfsdk:"logging" yaml:"logging,omitempty"`

			Master *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"master" yaml:"master,omitempty"`

			Mount *struct {
				Configs *[]string `tfsdk:"configs" yaml:"configs,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Resources *[]string `tfsdk:"resources" yaml:"resources,omitempty"`

				Volumes *[]string `tfsdk:"volumes" yaml:"volumes,omitempty"`
			} `tfsdk:"mount" yaml:"mount,omitempty"`

			Openapi *struct {
				Configmaps *[]string `tfsdk:"configmaps" yaml:"configmaps,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"openapi" yaml:"openapi,omitempty"`

			Owner *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				TargetAnnotations *[]string `tfsdk:"target_annotations" yaml:"targetAnnotations,omitempty"`

				TargetLabels *[]string `tfsdk:"target_labels" yaml:"targetLabels,omitempty"`
			} `tfsdk:"owner" yaml:"owner,omitempty"`

			Pdb *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				MaxUnavailable *string `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

				MinAvailable *string `tfsdk:"min_available" yaml:"minAvailable,omitempty"`
			} `tfsdk:"pdb" yaml:"pdb,omitempty"`

			Platform *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				CreateDefault *bool `tfsdk:"create_default" yaml:"createDefault,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Global *bool `tfsdk:"global" yaml:"global,omitempty"`
			} `tfsdk:"platform" yaml:"platform,omitempty"`

			Pod *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"pod" yaml:"pod,omitempty"`

			Prometheus *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				PodMonitor *bool `tfsdk:"pod_monitor" yaml:"podMonitor,omitempty"`

				PodMonitorLabels *[]string `tfsdk:"pod_monitor_labels" yaml:"podMonitorLabels,omitempty"`
			} `tfsdk:"prometheus" yaml:"prometheus,omitempty"`

			Pull_secret *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				ImagePullerDelegation *bool `tfsdk:"image_puller_delegation" yaml:"imagePullerDelegation,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"pull_secret" yaml:"pull-secret,omitempty"`

			Quarkus *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				PackageTypes *[]string `tfsdk:"package_types" yaml:"packageTypes,omitempty"`
			} `tfsdk:"quarkus" yaml:"quarkus,omitempty"`

			Registry *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"registry" yaml:"registry,omitempty"`

			Route *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				TlsCACertificate *string `tfsdk:"tls_ca_certificate" yaml:"tlsCACertificate,omitempty"`

				TlsCACertificateSecret *string `tfsdk:"tls_ca_certificate_secret" yaml:"tlsCACertificateSecret,omitempty"`

				TlsCertificate *string `tfsdk:"tls_certificate" yaml:"tlsCertificate,omitempty"`

				TlsCertificateSecret *string `tfsdk:"tls_certificate_secret" yaml:"tlsCertificateSecret,omitempty"`

				TlsDestinationCACertificate *string `tfsdk:"tls_destination_ca_certificate" yaml:"tlsDestinationCACertificate,omitempty"`

				TlsDestinationCACertificateSecret *string `tfsdk:"tls_destination_ca_certificate_secret" yaml:"tlsDestinationCACertificateSecret,omitempty"`

				TlsInsecureEdgeTerminationPolicy *string `tfsdk:"tls_insecure_edge_termination_policy" yaml:"tlsInsecureEdgeTerminationPolicy,omitempty"`

				TlsKey *string `tfsdk:"tls_key" yaml:"tlsKey,omitempty"`

				TlsKeySecret *string `tfsdk:"tls_key_secret" yaml:"tlsKeySecret,omitempty"`

				TlsTermination *string `tfsdk:"tls_termination" yaml:"tlsTermination,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			Service *struct {
				Auto *bool `tfsdk:"auto" yaml:"auto,omitempty"`

				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				NodePort *bool `tfsdk:"node_port" yaml:"nodePort,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"service" yaml:"service,omitempty"`

			Service_binding *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Services *[]string `tfsdk:"services" yaml:"services,omitempty"`
			} `tfsdk:"service_binding" yaml:"service-binding,omitempty"`

			Strimzi *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"strimzi" yaml:"strimzi,omitempty"`

			Toleration *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Taints *[]string `tfsdk:"taints" yaml:"taints,omitempty"`
			} `tfsdk:"toleration" yaml:"toleration,omitempty"`

			Tracing *struct {
				Configuration utilities.Dynamic `tfsdk:"configuration" yaml:"configuration,omitempty"`
			} `tfsdk:"tracing" yaml:"tracing,omitempty"`
		} `tfsdk:"traits" yaml:"traits,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCamelApacheOrgIntegrationV1Resource() resource.Resource {
	return &CamelApacheOrgIntegrationV1Resource{}
}

func (r *CamelApacheOrgIntegrationV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_camel_apache_org_integration_v1"
}

func (r *CamelApacheOrgIntegrationV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Integration is the Schema for the integrations API",
		MarkdownDescription: "Integration is the Schema for the integrations API",
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
				Description:         "the desired Integration specification",
				MarkdownDescription: "the desired Integration specification",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"configuration": {
						Description:         "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes",
						MarkdownDescription: "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"resource_key": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_mount_point": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_type": {
								Description:         "Deprecated: no longer used",
								MarkdownDescription: "Deprecated: no longer used",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": {
								Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
								MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
								MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",

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

					"dependencies": {
						Description:         "the list of Camel or Maven dependencies required by the Integration",
						MarkdownDescription: "the list of Camel or Maven dependencies required by the Integration",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"flows": {
						Description:         "a source in YAML DSL language which contain the routes to run",
						MarkdownDescription: "a source in YAML DSL language which contain the routes to run",

						Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"integration_kit": {
						Description:         "the reference of the 'IntegrationKit' which is used for this Integration",
						MarkdownDescription: "the reference of the 'IntegrationKit' which is used for this Integration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"field_path": {
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_version": {
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"uid": {
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",

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

					"profile": {
						Description:         "the profile needed to run this Integration",
						MarkdownDescription: "the profile needed to run this Integration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": {
						Description:         "the number of 'Pods' needed for the running Integration",
						MarkdownDescription: "the number of 'Pods' needed for the running Integration",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"repositories": {
						Description:         "additional Maven repositories to be used",
						MarkdownDescription: "additional Maven repositories to be used",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Deprecated: Use mount trait (mount.resources) to manage resources Use openapi trait (openapi.configmaps) to manage OpenAPIs specifications",
						MarkdownDescription: "Deprecated: Use mount trait (mount.resources) to manage resources Use openapi trait (openapi.configmaps) to manage OpenAPIs specifications",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"compression": {
								Description:         "if the content is compressed (base64 encrypted)",
								MarkdownDescription: "if the content is compressed (base64 encrypted)",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content": {
								Description:         "the source code (plain text)",
								MarkdownDescription: "the source code (plain text)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_key": {
								Description:         "the confimap key holding the source content",
								MarkdownDescription: "the confimap key holding the source content",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_ref": {
								Description:         "the confimap reference holding the source content",
								MarkdownDescription: "the confimap reference holding the source content",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_type": {
								Description:         "the content type (tipically text or binary)",
								MarkdownDescription: "the content type (tipically text or binary)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount_path": {
								Description:         "the mount path on destination 'Pod'",
								MarkdownDescription: "the mount path on destination 'Pod'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "the name of the specification",
								MarkdownDescription: "the name of the specification",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "the path where the file is stored",
								MarkdownDescription: "the path where the file is stored",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"raw_content": {
								Description:         "the source code (binary)",
								MarkdownDescription: "the source code (binary)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.Base64Validator(),
								},
							},

							"type": {
								Description:         "the kind of data to expect",
								MarkdownDescription: "the kind of data to expect",

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

					"service_account_name": {
						Description:         "custom SA to use for the Integration",
						MarkdownDescription: "custom SA to use for the Integration",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sources": {
						Description:         "the sources which contain the Camel routes to run",
						MarkdownDescription: "the sources which contain the Camel routes to run",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"compression": {
								Description:         "if the content is compressed (base64 encrypted)",
								MarkdownDescription: "if the content is compressed (base64 encrypted)",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content": {
								Description:         "the source code (plain text)",
								MarkdownDescription: "the source code (plain text)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_key": {
								Description:         "the confimap key holding the source content",
								MarkdownDescription: "the confimap key holding the source content",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_ref": {
								Description:         "the confimap reference holding the source content",
								MarkdownDescription: "the confimap reference holding the source content",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"content_type": {
								Description:         "the content type (tipically text or binary)",
								MarkdownDescription: "the content type (tipically text or binary)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interceptors": {
								Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
								MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"language": {
								Description:         "specify which is the language (Camel DSL) used to interpret this source code",
								MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loader": {
								Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
								MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "the name of the specification",
								MarkdownDescription: "the name of the specification",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "the path where the file is stored",
								MarkdownDescription: "the path where the file is stored",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"property_names": {
								Description:         "List of property names defined in the source (e.g. if type is 'template')",
								MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"raw_content": {
								Description:         "the source code (binary)",
								MarkdownDescription: "the source code (binary)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									validators.Base64Validator(),
								},
							},

							"type": {
								Description:         "Type defines the kind of source described by this object",
								MarkdownDescription: "Type defines the kind of source described by this object",

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

					"template": {
						Description:         "Pod template customization",
						MarkdownDescription: "Pod template customization",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"spec": {
								Description:         "the specification",
								MarkdownDescription: "the specification",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"active_deadline_seconds": {
										Description:         "ActiveDeadlineSeconds",
										MarkdownDescription: "ActiveDeadlineSeconds",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"containers": {
										Description:         "Containers",
										MarkdownDescription: "Containers",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"args": {
												Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"env": {
												Description:         "List of environment variables to set in the container. Cannot be updated.",
												MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
														MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_key_ref": {
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

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
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"api_version": {
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_path": {
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",

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
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"container_name": {
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"divisor": {
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource": {
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",

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
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",

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
												Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
												MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",

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
														Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "The Secret to select from",
														MarkdownDescription: "The Secret to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",

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
												Description:         "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
												MarkdownDescription: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_pull_policy": {
												Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
												MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lifecycle": {
												Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
												MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"post_start": {
														Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
												Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
												MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"ports": {
												Description:         "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",
												MarkdownDescription: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"container_port": {
														Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host_ip": {
														Description:         "What host IP to bind the external port to.",
														MarkdownDescription: "What host IP to bind the external port to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_port": {
														Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
														MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
														MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",

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
												Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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
												Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
												MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allow_privilege_escalation": {
														Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"capabilities": {
														Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"add": {
																Description:         "Added capabilities",
																MarkdownDescription: "Added capabilities",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"drop": {
																Description:         "Removed capabilities",
																MarkdownDescription: "Removed capabilities",

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
														Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proc_mount": {
														Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only_root_filesystem": {
														Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_group": {
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_non_root": {
														Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"se_linux_options": {
														Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"level": {
																Description:         "Level is SELinux level label that applies to the container.",
																MarkdownDescription: "Level is SELinux level label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "Role is a SELinux role label that applies to the container.",
																MarkdownDescription: "Role is a SELinux role label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "Type is a SELinux type label that applies to the container.",
																MarkdownDescription: "Type is a SELinux type label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": {
																Description:         "User is a SELinux user label that applies to the container.",
																MarkdownDescription: "User is a SELinux user label that applies to the container.",

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
														Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"localhost_profile": {
																Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
																MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
																MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",

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
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"gmsa_credential_spec": {
																Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"gmsa_credential_spec_name": {
																Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_process": {
																Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
																MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user_name": {
																Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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
												Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
												MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stdin_once": {
												Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
												MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_path": {
												Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
												MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_policy": {
												Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
												MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tty": {
												Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
												MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_devices": {
												Description:         "volumeDevices is the list of block devices to be used by the container.",
												MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"device_path": {
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",

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
												Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
												MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"mount_path": {
														Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"mount_propagation": {
														Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "This must match the Name of a Volume.",
														MarkdownDescription: "This must match the Name of a Volume.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path": {
														Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path_expr": {
														Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",

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
												Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
												MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",

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

									"dns_policy": {
										Description:         "DNSPolicy",
										MarkdownDescription: "DNSPolicy",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ephemeral_containers": {
										Description:         "EphemeralContainers",
										MarkdownDescription: "EphemeralContainers",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"args": {
												Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"env": {
												Description:         "List of environment variables to set in the container. Cannot be updated.",
												MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
														MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_key_ref": {
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

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
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"api_version": {
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_path": {
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",

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
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"container_name": {
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"divisor": {
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource": {
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",

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
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",

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
												Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
												MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",

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
														Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "The Secret to select from",
														MarkdownDescription: "The Secret to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",

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
												Description:         "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",
												MarkdownDescription: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_pull_policy": {
												Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
												MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lifecycle": {
												Description:         "Lifecycle is not allowed for ephemeral containers.",
												MarkdownDescription: "Lifecycle is not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"post_start": {
														Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
												Description:         "Probes are not allowed for ephemeral containers.",
												MarkdownDescription: "Probes are not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",
												MarkdownDescription: "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"ports": {
												Description:         "Ports are not allowed for ephemeral containers.",
												MarkdownDescription: "Ports are not allowed for ephemeral containers.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"container_port": {
														Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host_ip": {
														Description:         "What host IP to bind the external port to.",
														MarkdownDescription: "What host IP to bind the external port to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_port": {
														Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
														MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
														MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",

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
												Description:         "Probes are not allowed for ephemeral containers.",
												MarkdownDescription: "Probes are not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
												MarkdownDescription: "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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
												Description:         "Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
												MarkdownDescription: "Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allow_privilege_escalation": {
														Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"capabilities": {
														Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"add": {
																Description:         "Added capabilities",
																MarkdownDescription: "Added capabilities",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"drop": {
																Description:         "Removed capabilities",
																MarkdownDescription: "Removed capabilities",

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
														Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proc_mount": {
														Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only_root_filesystem": {
														Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_group": {
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_non_root": {
														Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"se_linux_options": {
														Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"level": {
																Description:         "Level is SELinux level label that applies to the container.",
																MarkdownDescription: "Level is SELinux level label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "Role is a SELinux role label that applies to the container.",
																MarkdownDescription: "Role is a SELinux role label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "Type is a SELinux type label that applies to the container.",
																MarkdownDescription: "Type is a SELinux type label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": {
																Description:         "User is a SELinux user label that applies to the container.",
																MarkdownDescription: "User is a SELinux user label that applies to the container.",

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
														Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"localhost_profile": {
																Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
																MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
																MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",

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
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"gmsa_credential_spec": {
																Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"gmsa_credential_spec_name": {
																Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_process": {
																Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
																MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user_name": {
																Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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
												Description:         "Probes are not allowed for ephemeral containers.",
												MarkdownDescription: "Probes are not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
												MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stdin_once": {
												Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
												MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_container_name": {
												Description:         "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec.  The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.",
												MarkdownDescription: "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec.  The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_path": {
												Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
												MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_policy": {
												Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
												MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tty": {
												Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
												MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_devices": {
												Description:         "volumeDevices is the list of block devices to be used by the container.",
												MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"device_path": {
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",

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
												Description:         "Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.",
												MarkdownDescription: "Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"mount_path": {
														Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"mount_propagation": {
														Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "This must match the Name of a Volume.",
														MarkdownDescription: "This must match the Name of a Volume.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path": {
														Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path_expr": {
														Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",

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
												Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
												MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",

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
										Description:         "InitContainers",
										MarkdownDescription: "InitContainers",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"args": {
												Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"env": {
												Description:         "List of environment variables to set in the container. Cannot be updated.",
												MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
														MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_key_ref": {
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

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
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"api_version": {
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_path": {
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",

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
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"container_name": {
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"divisor": {
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource": {
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",

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
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",

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
												Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
												MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_ref": {
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",

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
														Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "The Secret to select from",
														MarkdownDescription: "The Secret to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"optional": {
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",

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
												Description:         "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
												MarkdownDescription: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_pull_policy": {
												Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
												MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lifecycle": {
												Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
												MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"post_start": {
														Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "Exec specifies the action to take.",
																MarkdownDescription: "Exec specifies the action to take.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"command": {
																		Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																		MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
																Description:         "HTTPGet specifies the http request to perform.",
																MarkdownDescription: "HTTPGet specifies the http request to perform.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

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
																		Description:         "Path to access on the HTTP server.",
																		MarkdownDescription: "Path to access on the HTTP server.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"scheme": {
																		Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																		MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
																Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
												Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
												MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"ports": {
												Description:         "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",
												MarkdownDescription: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"container_port": {
														Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host_ip": {
														Description:         "What host IP to bind the external port to.",
														MarkdownDescription: "What host IP to bind the external port to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_port": {
														Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
														MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
														MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",

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
												Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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
												Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
												MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allow_privilege_escalation": {
														Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"capabilities": {
														Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"add": {
																Description:         "Added capabilities",
																MarkdownDescription: "Added capabilities",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"drop": {
																Description:         "Removed capabilities",
																MarkdownDescription: "Removed capabilities",

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
														Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proc_mount": {
														Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only_root_filesystem": {
														Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_group": {
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_non_root": {
														Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"se_linux_options": {
														Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"level": {
																Description:         "Level is SELinux level label that applies to the container.",
																MarkdownDescription: "Level is SELinux level label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"role": {
																Description:         "Role is a SELinux role label that applies to the container.",
																MarkdownDescription: "Role is a SELinux role label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "Type is a SELinux type label that applies to the container.",
																MarkdownDescription: "Type is a SELinux type label that applies to the container.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"user": {
																Description:         "User is a SELinux user label that applies to the container.",
																MarkdownDescription: "User is a SELinux user label that applies to the container.",

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
														Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"localhost_profile": {
																Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
																MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
																MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",

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
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"gmsa_credential_spec": {
																Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"gmsa_credential_spec_name": {
																Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"host_process": {
																Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
																MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user_name": {
																Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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
												Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "Exec specifies the action to take.",
														MarkdownDescription: "Exec specifies the action to take.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"grpc": {
														Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
														MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"service": {
																Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
																MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_headers": {
																Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

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
																Description:         "Path to access on the HTTP server.",
																MarkdownDescription: "Path to access on the HTTP server.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: utilities.IntOrStringType{},

																Required: true,
																Optional: false,
																Computed: false,
															},

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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
														Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"period_seconds": {
														Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
														MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"success_threshold": {
														Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
														MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_socket": {
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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
														Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
														MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout_seconds": {
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
												Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
												MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stdin_once": {
												Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
												MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_path": {
												Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
												MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination_message_policy": {
												Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
												MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tty": {
												Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
												MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_devices": {
												Description:         "volumeDevices is the list of block devices to be used by the container.",
												MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"device_path": {
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",

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
												Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
												MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"mount_path": {
														Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"mount_propagation": {
														Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"name": {
														Description:         "This must match the Name of a Volume.",
														MarkdownDescription: "This must match the Name of a Volume.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path": {
														Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path_expr": {
														Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",

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
												Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
												MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",

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

									"node_selector": {
										Description:         "NodeSelector",
										MarkdownDescription: "NodeSelector",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"restart_policy": {
										Description:         "RestartPolicy",
										MarkdownDescription: "RestartPolicy",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"security_context": {
										Description:         "PodSecurityContext",
										MarkdownDescription: "PodSecurityContext",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_group": {
												Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fs_group_change_policy": {
												Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_group": {
												Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_non_root": {
												Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_user": {
												Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"se_linux_options": {
												Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"level": {
														Description:         "Level is SELinux level label that applies to the container.",
														MarkdownDescription: "Level is SELinux level label that applies to the container.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"role": {
														Description:         "Role is a SELinux role label that applies to the container.",
														MarkdownDescription: "Role is a SELinux role label that applies to the container.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "Type is a SELinux type label that applies to the container.",
														MarkdownDescription: "Type is a SELinux type label that applies to the container.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"user": {
														Description:         "User is a SELinux user label that applies to the container.",
														MarkdownDescription: "User is a SELinux user label that applies to the container.",

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
												Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"localhost_profile": {
														Description:         "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used. The profile must be preconfigured on the node to work. Must be a descending path, relative to the kubelet's configured seccomp profile location. Must only be set if type is 'Localhost'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied. Valid options are:  Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.",

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
												Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container. Note that this field cannot be set when spec.os.name is windows.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sysctls": {
												Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of a property to set",
														MarkdownDescription: "Name of a property to set",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value of a property to set",
														MarkdownDescription: "Value of a property to set",

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
												Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
												MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"gmsa_credential_spec": {
														Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"gmsa_credential_spec_name": {
														Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
														MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host_process": {
														Description:         "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",
														MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. This field is alpha-level and will only be honored by components that enable the WindowsHostProcessContainers feature flag. Setting this field without the feature flag will result in errors when validating the Pod. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).  In addition, if HostProcess is true then HostNetwork must also be set to true.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user_name": {
														Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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

									"termination_grace_period_seconds": {
										Description:         "TerminationGracePeriodSeconds",
										MarkdownDescription: "TerminationGracePeriodSeconds",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"topology_spread_constraints": {
										Description:         "TopologySpreadConstraints",
										MarkdownDescription: "TopologySpreadConstraints",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"label_selector": {
												Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
												MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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
												Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 1/1/0: | zone1 | zone2 | zone3 | |   P   |   P   |       | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 1/1/1; scheduling it onto zone1(zone2) would make the ActualSkew(2-0) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
												MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 1/1/0: | zone1 | zone2 | zone3 | |   P   |   P   |       | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 1/1/1; scheduling it onto zone1(zone2) would make the ActualSkew(2-0) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"topology_key": {
												Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. It's a required field.",
												MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. It's a required field.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"when_unsatisfiable": {
												Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,   but giving higher precedence to topologies that would help reduce the   skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
												MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,   but giving higher precedence to topologies that would help reduce the   skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",

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
										Description:         "Volumes",
										MarkdownDescription: "Volumes",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"aws_elastic_block_store": {
												Description:         "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"partition": {
														Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
														MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_id": {
														Description:         "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

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
												Description:         "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
												MarkdownDescription: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"caching_mode": {
														Description:         "Host Caching mode: None, Read Only, Read Write.",
														MarkdownDescription: "Host Caching mode: None, Read Only, Read Write.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"disk_name": {
														Description:         "The Name of the data disk in the blob storage",
														MarkdownDescription: "The Name of the data disk in the blob storage",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"disk_uri": {
														Description:         "The URI the data disk in the blob storage",
														MarkdownDescription: "The URI the data disk in the blob storage",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
														MarkdownDescription: "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

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
												Description:         "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
												MarkdownDescription: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"read_only": {
														Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "the name of secret that contains Azure Storage Account Name and Key",
														MarkdownDescription: "the name of secret that contains Azure Storage Account Name and Key",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"share_name": {
														Description:         "Share Name",
														MarkdownDescription: "Share Name",

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
												Description:         "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
												MarkdownDescription: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"monitors": {
														Description:         "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"path": {
														Description:         "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
														MarkdownDescription: "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_file": {
														Description:         "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
												Description:         "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "Optional: points to a secret object containing parameters used to connect to OpenStack.",
														MarkdownDescription: "Optional: points to a secret object containing parameters used to connect to OpenStack.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

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
												Description:         "ConfigMap represents a configMap that should populate this volume",
												MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key to project.",
																MarkdownDescription: "The key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"optional": {
														Description:         "Specify whether the ConfigMap or its keys must be defined",
														MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",

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
												Description:         "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
												MarkdownDescription: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"driver": {
														Description:         "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
														MarkdownDescription: "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
														MarkdownDescription: "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_publish_secret_ref": {
														Description:         "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
														MarkdownDescription: "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
														MarkdownDescription: "Specifies a read-only configuration for the volume. Defaults to false (read/write).",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_attributes": {
														Description:         "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
														MarkdownDescription: "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",

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
												Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
												MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "Items is a list of downward API volume file",
														MarkdownDescription: "Items is a list of downward API volume file",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"field_ref": {
																Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"api_version": {
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_path": {
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",

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
																Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"resource_field_ref": {
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"container_name": {
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"divisor": {
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource": {
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",

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
												Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"medium": {
														Description:         "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														MarkdownDescription: "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size_limit": {
														Description:         "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
														MarkdownDescription: "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",

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
												Description:         "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
												MarkdownDescription: "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"volume_claim_template": {
														Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
														MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"metadata": {
																Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"spec": {
																Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"access_modes": {
																		Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																		MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"data_source": {
																		Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																		MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_group": {
																				Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																				MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "Kind is the type of resource being referenced",
																				MarkdownDescription: "Kind is the type of resource being referenced",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name is the name of resource being referenced",
																				MarkdownDescription: "Name is the name of resource being referenced",

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
																		Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																		MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_group": {
																				Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																				MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "Kind is the type of resource being referenced",
																				MarkdownDescription: "Kind is the type of resource being referenced",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name is the name of resource being referenced",
																				MarkdownDescription: "Name is the name of resource being referenced",

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
																		Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																		MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"limits": {
																				Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																				MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"requests": {
																				Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																				MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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
																		Description:         "A label query over volumes to consider for binding.",
																		MarkdownDescription: "A label query over volumes to consider for binding.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"operator": {
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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
																				Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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
																		Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																		MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"volume_mode": {
																		Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																		MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"volume_name": {
																		Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
																		MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",

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
												Description:         "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
												MarkdownDescription: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lun": {
														Description:         "Optional: FC target lun number",
														MarkdownDescription: "Optional: FC target lun number",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_ww_ns": {
														Description:         "Optional: FC target worldwide names (WWNs)",
														MarkdownDescription: "Optional: FC target worldwide names (WWNs)",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"wwids": {
														Description:         "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
														MarkdownDescription: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",

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
												Description:         "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
												MarkdownDescription: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"driver": {
														Description:         "Driver is the name of the driver to use for this volume.",
														MarkdownDescription: "Driver is the name of the driver to use for this volume.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"options": {
														Description:         "Optional: Extra command options if any.",
														MarkdownDescription: "Optional: Extra command options if any.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
														MarkdownDescription: "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												Description:         "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
												MarkdownDescription: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dataset_name": {
														Description:         "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
														MarkdownDescription: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dataset_uuid": {
														Description:         "UUID of the dataset. This is unique identifier of a Flocker dataset",
														MarkdownDescription: "UUID of the dataset. This is unique identifier of a Flocker dataset",

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
												Description:         "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"partition": {
														Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pd_name": {
														Description:         "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

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
												Description:         "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
												MarkdownDescription: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"directory": {
														Description:         "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
														MarkdownDescription: "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"repository": {
														Description:         "Repository URL",
														MarkdownDescription: "Repository URL",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision": {
														Description:         "Commit hash for the specified revision.",
														MarkdownDescription: "Commit hash for the specified revision.",

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
												Description:         "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
												MarkdownDescription: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"endpoints": {
														Description:         "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"path": {
														Description:         "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

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
												Description:         "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
												MarkdownDescription: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"type": {
														Description:         "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

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
												Description:         "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
												MarkdownDescription: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"chap_auth_discovery": {
														Description:         "whether support iSCSI Discovery CHAP authentication",
														MarkdownDescription: "whether support iSCSI Discovery CHAP authentication",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"chap_auth_session": {
														Description:         "whether support iSCSI Session CHAP authentication",
														MarkdownDescription: "whether support iSCSI Session CHAP authentication",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"fs_type": {
														Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"initiator_name": {
														Description:         "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
														MarkdownDescription: "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"iqn": {
														Description:         "Target iSCSI Qualified Name.",
														MarkdownDescription: "Target iSCSI Qualified Name.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"iscsi_interface": {
														Description:         "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
														MarkdownDescription: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lun": {
														Description:         "iSCSI Target Lun number.",
														MarkdownDescription: "iSCSI Target Lun number.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"portals": {
														Description:         "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
														MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "CHAP Secret for iSCSI target and initiator authentication",
														MarkdownDescription: "CHAP Secret for iSCSI target and initiator authentication",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

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
												Description:         "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"nfs": {
												Description:         "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"server": {
														Description:         "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

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
												Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"claim_name": {
														Description:         "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
														MarkdownDescription: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "Will force the ReadOnly setting in VolumeMounts. Default false.",
														MarkdownDescription: "Will force the ReadOnly setting in VolumeMounts. Default false.",

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
												Description:         "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
												MarkdownDescription: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pd_id": {
														Description:         "ID that identifies Photon Controller persistent disk",
														MarkdownDescription: "ID that identifies Photon Controller persistent disk",

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
												Description:         "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
												MarkdownDescription: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_id": {
														Description:         "VolumeID uniquely identifies a Portworx volume",
														MarkdownDescription: "VolumeID uniquely identifies a Portworx volume",

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
												Description:         "Items for all in one resources secrets, configmaps, and downward API",
												MarkdownDescription: "Items for all in one resources secrets, configmaps, and downward API",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sources": {
														Description:         "list of volume projections",
														MarkdownDescription: "list of volume projections",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "information about the configMap data to project",
																MarkdownDescription: "information about the configMap data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key to project.",
																				MarkdownDescription: "The key to project.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"mode": {
																				Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the ConfigMap or its keys must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",

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
																Description:         "information about the downwardAPI data to project",
																MarkdownDescription: "information about the downwardAPI data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "Items is a list of DownwardAPIVolume file",
																		MarkdownDescription: "Items is a list of DownwardAPIVolume file",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"field_ref": {
																				Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																				MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"api_version": {
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"field_path": {
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",

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
																				Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																				MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"resource_field_ref": {
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"container_name": {
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"divisor": {
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"resource": {
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",

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
																Description:         "information about the secret data to project",
																MarkdownDescription: "information about the secret data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key to project.",
																				MarkdownDescription: "The key to project.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"mode": {
																				Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"optional": {
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",

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
																Description:         "information about the serviceAccountToken data to project",
																MarkdownDescription: "information about the serviceAccountToken data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"audience": {
																		Description:         "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																		MarkdownDescription: "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expiration_seconds": {
																		Description:         "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																		MarkdownDescription: "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "Path is the path relative to the mount point of the file to project the token into.",
																		MarkdownDescription: "Path is the path relative to the mount point of the file to project the token into.",

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
												Description:         "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
												MarkdownDescription: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"group": {
														Description:         "Group to map volume access to Default is no group",
														MarkdownDescription: "Group to map volume access to Default is no group",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
														MarkdownDescription: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"registry": {
														Description:         "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
														MarkdownDescription: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"tenant": {
														Description:         "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
														MarkdownDescription: "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"user": {
														Description:         "User to map volume access to Defaults to serivceaccount user",
														MarkdownDescription: "User to map volume access to Defaults to serivceaccount user",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume": {
														Description:         "Volume is a string that references an already created Quobyte volume by name.",
														MarkdownDescription: "Volume is a string that references an already created Quobyte volume by name.",

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
												Description:         "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
												MarkdownDescription: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"keyring": {
														Description:         "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"monitors": {
														Description:         "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"pool": {
														Description:         "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
												Description:         "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
												MarkdownDescription: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"gateway": {
														Description:         "The host address of the ScaleIO API Gateway.",
														MarkdownDescription: "The host address of the ScaleIO API Gateway.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"protection_domain": {
														Description:         "The name of the ScaleIO Protection Domain for the configured storage.",
														MarkdownDescription: "The name of the ScaleIO Protection Domain for the configured storage.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
														MarkdownDescription: "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "Flag to enable/disable SSL communication with Gateway, default false",
														MarkdownDescription: "Flag to enable/disable SSL communication with Gateway, default false",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_mode": {
														Description:         "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
														MarkdownDescription: "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_pool": {
														Description:         "The ScaleIO Storage Pool associated with the protection domain.",
														MarkdownDescription: "The ScaleIO Storage Pool associated with the protection domain.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"system": {
														Description:         "The name of the storage system as configured in ScaleIO.",
														MarkdownDescription: "The name of the storage system as configured in ScaleIO.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"volume_name": {
														Description:         "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
														MarkdownDescription: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",

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
												Description:         "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
												MarkdownDescription: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "The key to project.",
																MarkdownDescription: "The key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
														Description:         "Specify whether the Secret or its keys must be defined",
														MarkdownDescription: "Specify whether the Secret or its keys must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
														MarkdownDescription: "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

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
												Description:         "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
												MarkdownDescription: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
														MarkdownDescription: "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
														MarkdownDescription: "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_namespace": {
														Description:         "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
														MarkdownDescription: "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",

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
												Description:         "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
												MarkdownDescription: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_policy_id": {
														Description:         "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
														MarkdownDescription: "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_policy_name": {
														Description:         "Storage Policy Based Management (SPBM) profile name.",
														MarkdownDescription: "Storage Policy Based Management (SPBM) profile name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_path": {
														Description:         "Path that identifies vSphere volume vmdk",
														MarkdownDescription: "Path that identifies vSphere volume vmdk",

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

					"traits": {
						Description:         "the traits needed to run this Integration",
						MarkdownDescription: "the traits needed to run this Integration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"threescale": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"addons": {
								Description:         "The extension point with addon traits",
								MarkdownDescription: "The extension point with addon traits",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"affinity": {
								Description:         "The configuration of Affinity trait",
								MarkdownDescription: "The configuration of Affinity trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_affinity_labels": {
										Description:         "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
										MarkdownDescription: "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_affinity": {
										Description:         "Always co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Always co-locates multiple replicas of the integration in the same node (default *false*).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_affinity_labels": {
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_anti_affinity": {
										Description:         "Never co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Never co-locates multiple replicas of the integration in the same node (default *false*).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_anti_affinity_labels": {
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",

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

							"builder": {
								Description:         "The configuration of Builder trait",
								MarkdownDescription: "The configuration of Builder trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "A list of properties to be provided to the build task",
										MarkdownDescription: "A list of properties to be provided to the build task",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verbose": {
										Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",

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

							"camel": {
								Description:         "The configuration of Camel trait",
								MarkdownDescription: "The configuration of Camel trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"properties": {
										Description:         "A list of properties to be provided to the Integration runtime",
										MarkdownDescription: "A list of properties to be provided to the Integration runtime",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"runtime_version": {
										Description:         "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										MarkdownDescription: "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",

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

							"container": {
								Description:         "The configuration of Container trait",
								MarkdownDescription: "The configuration of Container trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically enable the trait",
										MarkdownDescription: "To automatically enable the trait",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expose": {
										Description:         "Can be used to enable/disable exposure via kubernetes Service.",
										MarkdownDescription: "Can be used to enable/disable exposure via kubernetes Service.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "The main container image",
										MarkdownDescription: "The main container image",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_pull_policy": {
										Description:         "The pull policy: Always|Never|IfNotPresent",
										MarkdownDescription: "The pull policy: Always|Never|IfNotPresent",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"limit_cpu": {
										Description:         "The maximum amount of CPU required.",
										MarkdownDescription: "The maximum amount of CPU required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"limit_memory": {
										Description:         "The maximum amount of memory required.",
										MarkdownDescription: "The maximum amount of memory required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_initial_delay": {
										Description:         "Number of seconds after the container has started before liveness probes are initiated. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_period": {
										Description:         "How often to perform the probe. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "How often to perform the probe. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_scheme": {
										Description:         "Scheme to use when connecting. Defaults to HTTP. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Scheme to use when connecting. Defaults to HTTP. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_timeout": {
										Description:         "Number of seconds after which the probe times out. Applies to the liveness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after which the probe times out. Applies to the liveness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "The main container name. It's named 'integration' by default.",
										MarkdownDescription: "The main container name. It's named 'integration' by default.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "To configure a different port exposed by the container (default '8080').",
										MarkdownDescription: "To configure a different port exposed by the container (default '8080').",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port_name": {
										Description:         "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
										MarkdownDescription: "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"probes_enabled": {
										Description:         "DeprecatedProbesEnabled enable/disable probes on the container (default 'false'). Deprecated: replaced by the health trait.",
										MarkdownDescription: "DeprecatedProbesEnabled enable/disable probes on the container (default 'false'). Deprecated: replaced by the health trait.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_initial_delay": {
										Description:         "Number of seconds after the container has started before readiness probes are initiated. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after the container has started before readiness probes are initiated. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_period": {
										Description:         "How often to perform the probe. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "How often to perform the probe. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_scheme": {
										Description:         "Scheme to use when connecting. Defaults to HTTP. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Scheme to use when connecting. Defaults to HTTP. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_timeout": {
										Description:         "Number of seconds after which the probe times out. Applies to the readiness probe. Deprecated: replaced by the health trait.",
										MarkdownDescription: "Number of seconds after which the probe times out. Applies to the readiness probe. Deprecated: replaced by the health trait.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_cpu": {
										Description:         "The minimum amount of CPU required.",
										MarkdownDescription: "The minimum amount of CPU required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_memory": {
										Description:         "The minimum amount of memory required.",
										MarkdownDescription: "The minimum amount of memory required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_port": {
										Description:         "To configure under which service port the container port is to be exposed (default '80').",
										MarkdownDescription: "To configure under which service port the container port is to be exposed (default '80').",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_port_name": {
										Description:         "To configure under which service port name the container port is to be exposed (default 'http').",
										MarkdownDescription: "To configure under which service port name the container port is to be exposed (default 'http').",

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

							"cron": {
								Description:         "The configuration of Cron trait",
								MarkdownDescription: "The configuration of Cron trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"active_deadline_seconds": {
										Description:         "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auto": {
										Description:         "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
										MarkdownDescription: "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"backoff_limit": {
										Description:         "Specifies the number of retries before marking the job failed. It defaults to 2.",
										MarkdownDescription: "Specifies the number of retries before marking the job failed. It defaults to 2.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"components": {
										Description:         "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",
										MarkdownDescription: "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"concurrency_policy": {
										Description:         "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
										MarkdownDescription: "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Allow", "Forbid", "Replace"),
										},
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fallback": {
										Description:         "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
										MarkdownDescription: "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"schedule": {
										Description:         "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
										MarkdownDescription: "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"starting_deadline_seconds": {
										Description:         "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",
										MarkdownDescription: "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",

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

							"dependencies": {
								Description:         "The configuration of Dependencies trait",
								MarkdownDescription: "The configuration of Dependencies trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

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

							"deployer": {
								Description:         "The configuration of Deployer trait",
								MarkdownDescription: "The configuration of Deployer trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
										MarkdownDescription: "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("deployment", "cron-job", "knative-service"),
										},
									},

									"use_ssa": {
										Description:         "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
										MarkdownDescription: "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",

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

							"deployment": {
								Description:         "The configuration of Deployment trait",
								MarkdownDescription: "The configuration of Deployment trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"progress_deadline_seconds": {
										Description:         "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rolling_update_max_surge": {
										Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rolling_update_max_unavailable": {
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"strategy": {
										Description:         "The deployment strategy to use to replace existing pods with new ones.",
										MarkdownDescription: "The deployment strategy to use to replace existing pods with new ones.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Recreate", "RollingUpdate"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"environment": {
								Description:         "The configuration of Environment trait",
								MarkdownDescription: "The configuration of Environment trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"container_meta": {
										Description:         "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
										MarkdownDescription: "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_proxy": {
										Description:         "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
										MarkdownDescription: "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"vars": {
										Description:         "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",
										MarkdownDescription: "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",

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

							"error_handler": {
								Description:         "The configuration of Error Handler trait",
								MarkdownDescription: "The configuration of Error Handler trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ref": {
										Description:         "The error handler ref name provided or found in application properties",
										MarkdownDescription: "The error handler ref name provided or found in application properties",

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

							"gc": {
								Description:         "The configuration of GC trait",
								MarkdownDescription: "The configuration of GC trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"discovery_cache": {
										Description:         "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
										MarkdownDescription: "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("disabled", "disk", "memory"),
										},
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

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

							"health": {
								Description:         "The configuration of Health trait",
								MarkdownDescription: "The configuration of Health trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_failure_threshold": {
										Description:         "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_initial_delay": {
										Description:         "Number of seconds after the container has started before the liveness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the liveness probe is initiated.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_period": {
										Description:         "How often to perform the liveness probe.",
										MarkdownDescription: "How often to perform the liveness probe.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_probe_enabled": {
										Description:         "Configures the liveness probe for the integration container (default 'false').",
										MarkdownDescription: "Configures the liveness probe for the integration container (default 'false').",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_scheme": {
										Description:         "Scheme to use when connecting to the liveness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the liveness probe (default 'HTTP').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_success_threshold": {
										Description:         "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_timeout": {
										Description:         "Number of seconds after which the liveness probe times out.",
										MarkdownDescription: "Number of seconds after which the liveness probe times out.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_failure_threshold": {
										Description:         "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_initial_delay": {
										Description:         "Number of seconds after the container has started before the readiness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the readiness probe is initiated.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_period": {
										Description:         "How often to perform the readiness probe.",
										MarkdownDescription: "How often to perform the readiness probe.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_probe_enabled": {
										Description:         "Configures the readiness probe for the integration container (default 'true').",
										MarkdownDescription: "Configures the readiness probe for the integration container (default 'true').",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_scheme": {
										Description:         "Scheme to use when connecting to the readiness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the readiness probe (default 'HTTP').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_success_threshold": {
										Description:         "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_timeout": {
										Description:         "Number of seconds after which the readiness probe times out.",
										MarkdownDescription: "Number of seconds after which the readiness probe times out.",

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

							"ingress": {
								Description:         "The configuration of Ingress trait",
								MarkdownDescription: "The configuration of Ingress trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
										MarkdownDescription: "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "To configure the host exposed by the ingress.",
										MarkdownDescription: "To configure the host exposed by the ingress.",

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

							"istio": {
								Description:         "The configuration of Istio trait",
								MarkdownDescription: "The configuration of Istio trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allow": {
										Description:         "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
										MarkdownDescription: "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"inject": {
										Description:         "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
										MarkdownDescription: "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",

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

							"jolokia": {
								Description:         "The configuration of Jolokia trait",
								MarkdownDescription: "The configuration of Jolokia trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_cert": {
										Description:         "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
										MarkdownDescription: "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_principal": {
										Description:         "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
										MarkdownDescription: "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"discovery_enabled": {
										Description:         "Listen for multicast requests (default 'false')",
										MarkdownDescription: "Listen for multicast requests (default 'false')",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"extended_client_check": {
										Description:         "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
										MarkdownDescription: "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
										MarkdownDescription: "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"options": {
										Description:         "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
										MarkdownDescription: "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password": {
										Description:         "The password used for authentication, applicable when the 'user' option is set.",
										MarkdownDescription: "The password used for authentication, applicable when the 'user' option is set.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "The Jolokia endpoint port (default '8778').",
										MarkdownDescription: "The Jolokia endpoint port (default '8778').",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol": {
										Description:         "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
										MarkdownDescription: "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_ssl_client_authentication": {
										Description:         "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
										MarkdownDescription: "Whether client certificates should be used for authentication (default 'true' for OpenShift).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": {
										Description:         "The user to be used for authentication",
										MarkdownDescription: "The user to be used for authentication",

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

							"jvm": {
								Description:         "The configuration of JVM trait",
								MarkdownDescription: "The configuration of JVM trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"classpath": {
										Description:         "Additional JVM classpath (use 'Linux' classpath separator)",
										MarkdownDescription: "Additional JVM classpath (use 'Linux' classpath separator)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"debug": {
										Description:         "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
										MarkdownDescription: "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"debug_address": {
										Description:         "Transport address at which to listen for the newly launched JVM (default '*:5005')",
										MarkdownDescription: "Transport address at which to listen for the newly launched JVM (default '*:5005')",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"debug_suspend": {
										Description:         "Suspends the target JVM immediately before the main class is loaded",
										MarkdownDescription: "Suspends the target JVM immediately before the main class is loaded",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"options": {
										Description:         "A list of JVM options",
										MarkdownDescription: "A list of JVM options",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"print_command": {
										Description:         "Prints the command used the start the JVM in the container logs (default 'true')",
										MarkdownDescription: "Prints the command used the start the JVM in the container logs (default 'true')",

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

							"kamelets": {
								Description:         "The configuration of Kamelets trait",
								MarkdownDescription: "The configuration of Kamelets trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
										MarkdownDescription: "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"list": {
										Description:         "Comma separated list of Kamelet names to load into the current integration",
										MarkdownDescription: "Comma separated list of Kamelet names to load into the current integration",

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

							"keda": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"knative": {
								Description:         "The configuration of Knative trait",
								MarkdownDescription: "The configuration of Knative trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Enable automatic discovery of all trait properties.",
										MarkdownDescription: "Enable automatic discovery of all trait properties.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"channel_sinks": {
										Description:         "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"channel_sources": {
										Description:         "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"config": {
										Description:         "Can be used to inject a Knative complete configuration in JSON format.",
										MarkdownDescription: "Can be used to inject a Knative complete configuration in JSON format.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"endpoint_sinks": {
										Description:         "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
										MarkdownDescription: "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"endpoint_sources": {
										Description:         "List of channels used as source of integration routes.",
										MarkdownDescription: "List of channels used as source of integration routes.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"event_sinks": {
										Description:         "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
										MarkdownDescription: "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"event_sources": {
										Description:         "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
										MarkdownDescription: "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"filter_source_channels": {
										Description:         "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
										MarkdownDescription: "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sink_binding": {
										Description:         "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
										MarkdownDescription: "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",

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

							"knative_service": {
								Description:         "The configuration of Knative Service trait",
								MarkdownDescription: "The configuration of Knative Service trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
										MarkdownDescription: "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"autoscaling_metric": {
										Description:         "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"autoscaling_target": {
										Description:         "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"class": {
										Description:         "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("kpa.autoscaling.knative.dev", "hpa.autoscaling.knative.dev"),
										},
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_scale": {
										Description:         "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_scale": {
										Description:         "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rollout_duration": {
										Description:         "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
										MarkdownDescription: "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"visibility": {
										Description:         "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("cluster-local"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": {
								Description:         "The configuration of Logging trait",
								MarkdownDescription: "The configuration of Logging trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"color": {
										Description:         "Colorize the log output",
										MarkdownDescription: "Colorize the log output",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"format": {
										Description:         "Logs message format",
										MarkdownDescription: "Logs message format",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json": {
										Description:         "Output the logs in JSON",
										MarkdownDescription: "Output the logs in JSON",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_pretty_print": {
										Description:         "Enable 'pretty printing' of the JSON logs",
										MarkdownDescription: "Enable 'pretty printing' of the JSON logs",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"level": {
										Description:         "Adjust the logging level (defaults to INFO)",
										MarkdownDescription: "Adjust the logging level (defaults to INFO)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("FATAL", "WARN", "INFO", "DEBUG", "TRACE"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"master": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount": {
								Description:         "The configuration of Mount trait",
								MarkdownDescription: "The configuration of Mount trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configs": {
										Description:         "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
										MarkdownDescription: "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
										MarkdownDescription: "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volumes": {
										Description:         "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
										MarkdownDescription: "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",

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

							"openapi": {
								Description:         "The configuration of OpenAPI trait",
								MarkdownDescription: "The configuration of OpenAPI trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configmaps": {
										Description:         "The configmaps holding the spec of the OpenAPI",
										MarkdownDescription: "The configmaps holding the spec of the OpenAPI",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

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

							"owner": {
								Description:         "The configuration of Owner trait",
								MarkdownDescription: "The configuration of Owner trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_annotations": {
										Description:         "The set of annotations to be transferred",
										MarkdownDescription: "The set of annotations to be transferred",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_labels": {
										Description:         "The set of labels to be transferred",
										MarkdownDescription: "The set of labels to be transferred",

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

							"pdb": {
								Description:         "The configuration of PDB trait",
								MarkdownDescription: "The configuration of PDB trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_unavailable": {
										Description:         "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_available": {
										Description:         "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",

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

							"platform": {
								Description:         "The configuration of Platform trait",
								MarkdownDescription: "The configuration of Platform trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",
										MarkdownDescription: "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"create_default": {
										Description:         "To create a default (empty) platform when the platform is missing.",
										MarkdownDescription: "To create a default (empty) platform when the platform is missing.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"global": {
										Description:         "Indicates if the platform should be created globally in the case of global operator (default true).",
										MarkdownDescription: "Indicates if the platform should be created globally in the case of global operator (default true).",

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

							"pod": {
								Description:         "The configuration of Pod trait",
								MarkdownDescription: "The configuration of Pod trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

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

							"prometheus": {
								Description:         "The configuration of Prometheus trait",
								MarkdownDescription: "The configuration of Prometheus trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_monitor": {
										Description:         "Whether a 'PodMonitor' resource is created (default 'true').",
										MarkdownDescription: "Whether a 'PodMonitor' resource is created (default 'true').",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_monitor_labels": {
										Description:         "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
										MarkdownDescription: "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",

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

							"pull_secret": {
								Description:         "The configuration of Pull Secret trait",
								MarkdownDescription: "The configuration of Pull Secret trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
										MarkdownDescription: "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_puller_delegation": {
										Description:         "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
										MarkdownDescription: "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
										MarkdownDescription: "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",

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

							"quarkus": {
								Description:         "The configuration of Quarkus trait",
								MarkdownDescription: "The configuration of Quarkus trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"package_types": {
										Description:         "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										MarkdownDescription: "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",

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

							"registry": {
								Description:         "The configuration of Registry trait",
								MarkdownDescription: "The configuration of Registry trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

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

							"route": {
								Description:         "The configuration of Route trait",
								MarkdownDescription: "The configuration of Route trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "To configure the host exposed by the route.",
										MarkdownDescription: "To configure the host exposed by the route.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_ca_certificate": {
										Description:         "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_ca_certificate_secret": {
										Description:         "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_certificate": {
										Description:         "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_certificate_secret": {
										Description:         "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_destination_ca_certificate": {
										Description:         "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_destination_ca_certificate_secret": {
										Description:         "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_insecure_edge_termination_policy": {
										Description:         "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("None", "Allow", "Redirect"),
										},
									},

									"tls_key": {
										Description:         "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_key_secret": {
										Description:         "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_termination": {
										Description:         "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("edge", "reencrypt", "passthrough"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service": {
								Description:         "The configuration of Service trait",
								MarkdownDescription: "The configuration of Service trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto": {
										Description:         "To automatically detect from the code if a Service needs to be created.",
										MarkdownDescription: "To automatically detect from the code if a Service needs to be created.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_port": {
										Description:         "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
										MarkdownDescription: "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
										MarkdownDescription: "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_binding": {
								Description:         "The configuration of Service Binding trait",
								MarkdownDescription: "The configuration of Service Binding trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"services": {
										Description:         "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
										MarkdownDescription: "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",

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

							"strimzi": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"toleration": {
								Description:         "The configuration of Toleration trait",
								MarkdownDescription: "The configuration of Toleration trait",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"taints": {
										Description:         "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
										MarkdownDescription: "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",

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

							"tracing": {
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"configuration": {
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",

										Type: utilities.DynamicType{},

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
		},
	}, nil
}

func (r *CamelApacheOrgIntegrationV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_camel_apache_org_integration_v1")

	var state CamelApacheOrgIntegrationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgIntegrationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("Integration")

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

func (r *CamelApacheOrgIntegrationV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_integration_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CamelApacheOrgIntegrationV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_camel_apache_org_integration_v1")

	var state CamelApacheOrgIntegrationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CamelApacheOrgIntegrationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("camel.apache.org/v1")
	goModel.Kind = utilities.Ptr("Integration")

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

func (r *CamelApacheOrgIntegrationV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_camel_apache_org_integration_v1")
	// NO-OP: Terraform removes the state automatically for us
}
